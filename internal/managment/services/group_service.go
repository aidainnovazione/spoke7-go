package services

import (
	"context"
	"fmt"
	"slices"
	"spoke7-go/internal/managment/dtos"
	"spoke7-go/internal/managment/models"
	"strings"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type GroupService interface {
	ListAll(ctx context.Context) ([]*models.GroupInfo, error)
	ListByManager(ctx context.Context, user string) ([]*models.GroupInfo, error)
	ListByMember(ctx context.Context, user string, groups []string) ([]*models.GroupInfo, error)
	Create(ctx context.Context, group models.GroupInfo) (*models.GroupInfo, error)
	Get(ctx context.Context, groupName string) (*models.GroupInfo, error)
	Update(ctx context.Context, group models.GroupInfo) (*models.GroupInfo, error)
	Delete(ctx context.Context, groupName string) error
}
type groupService struct {
	organizationName string
	casdoorClient    *casdoorsdk.Client
}

func NewGroupService(organizationName string, client *casdoorsdk.Client) GroupService {

	return &groupService{organizationName: organizationName, casdoorClient: client}
}

// ListAll retrieves all groups. This typically requires administrative permissions
// or a broad scope of access.
func (s *groupService) ListAll(ctx context.Context) ([]*models.GroupInfo, error) {
	groups, err := s.casdoorClient.GetGroups()
	if err != nil {
		return nil, fmt.Errorf("failed to get all groups from Casdoor: %w", err)
	}

	return dtos.CasdoorGroupsToModels(groups), nil
}

// ListByManager retrieves groups where the specified user is the manager.
func (s *groupService) ListByManager(ctx context.Context, managerID string) ([]*models.GroupInfo, error) {
	if managerID == "" {
		return nil, fmt.Errorf("managerID cannot be empty for ListByManager")
	}

	query := map[string]string{
		"field": "manager",
		"value": managerID,
	}
	groups, _, err := s.casdoorClient.GetPaginationGroups(0, 1000, query) // Adjust pagination as needed
	if err != nil {
		return nil, fmt.Errorf("failed to get groups by manager from Casdoor: %w", err)
	}

	return dtos.CasdoorGroupsToModels(groups), nil
}

// ListByMember retrieves groups where the specified user is a member.
// Note: Casdoor's SDK currently doesn't directly support querying groups by member ID.
// This implementation fetches all groups and then filters them in memory.
// For large numbers of groups, consider adding a specific Casdoor SDK method
// or a more efficient database query if available.
func (s *groupService) ListByMember(ctx context.Context, userID string, groups []string) ([]*models.GroupInfo, error) {
	if userID == "" {
		return nil, fmt.Errorf("userID cannot be empty for ListByMember")
	}

	// Fetch all groups as Casdoor SDK doesn't have a direct 'GetGroupsByMember'
	allGroups, err := s.casdoorClient.GetGroups()
	if err != nil {
		return nil, fmt.Errorf("failed to get all groups for member filtering: %w", err)
	}

	filteredGroups := make([]*casdoorsdk.Group, 0, len(allGroups))
	for _, group := range allGroups {
		// Check if the userID is present in the group's Users slice
		if group.Manager == userID || slices.Contains(groups, group.Name) {
			filteredGroups = append(filteredGroups, group)
		}

	}

	return dtos.CasdoorGroupsToModels(filteredGroups), nil
}

func (s *groupService) Create(ctx context.Context, group models.GroupInfo) (*models.GroupInfo, error) {

	newGroup := casdoorsdk.Group{
		Owner:        group.Owner,
		Name:         group.Name,
		CreatedTime:  casdoorsdk.GetCurrentTime(),
		UpdatedTime:  casdoorsdk.GetCurrentTime(),
		DisplayName:  group.DisplayName,
		Manager:      group.Manager,
		ContactEmail: "",
		Type:         "Virtual",
		ParentId:     s.organizationName,
		IsTopGroup:   false,
		Users:        []string{group.Owner},
		Title:        "",
		Key:          "",
		Children:     []*casdoorsdk.Group{},
		IsEnabled:    true,
	}

	_, err := s.casdoorClient.AddGroup(&newGroup)
	if err != nil {
		return nil, err
	}

	groupInfo := &models.GroupInfo{
		Owner:        newGroup.Owner,
		Name:         newGroup.Name,
		CreatedTime:  newGroup.CreatedTime,
		UpdatedTime:  newGroup.UpdatedTime,
		DisplayName:  newGroup.DisplayName,
		Manager:      newGroup.Manager,
		ContactEmail: newGroup.ContactEmail,
		Type:         newGroup.Type,
		ParentId:     newGroup.ParentId,
		IsTopGroup:   newGroup.IsTopGroup,
		Users:        newGroup.Users,
		Title:        newGroup.Title,
		Key:          newGroup.Key,
		IsEnabled:    newGroup.IsEnabled,
	}

	return groupInfo, nil

}

func (s *groupService) Get(ctx context.Context, groupName string) (*models.GroupInfo, error) {
	casdoorGroup, err := s.casdoorClient.GetGroup(groupName)
	if err != nil {
		return nil, err
	}
	if casdoorGroup == nil {
		return nil, nil
	}

	var casdoorUsers []*casdoorsdk.User
	if casdoorGroup.Users != nil && len(casdoorGroup.Users) > 0 {
		casdoorUsers = make([]*casdoorsdk.User, 0, len(casdoorGroup.Users))
		for _, user := range casdoorGroup.Users {
			split := strings.SplitN(user, "/", 2)
			if len(split) > 1 {
				userName := split[1]
				casdoorUser, err := s.casdoorClient.GetUser(userName)
				if err != nil {
					return nil, err
				}
				casdoorUsers = append(casdoorUsers, casdoorUser)
			}
		}
	}

	userInfos := make([]models.UserInfo, 0, len(casdoorUsers))
	for _, user := range casdoorUsers {
		userInfo := models.UserInfo{
			Name:        user.Name,
			Email:       user.Email,
			Sub:         user.Id,
			Iss:         "",
			Aud:         "",
			DisplayName: user.DisplayName,
			Avatar:      "",
			Address:     []string{},
			Phone:       user.Phone,
			Groups:      user.Groups,
		}
		userInfos = append(userInfos, userInfo)
	}

	return &models.GroupInfo{
		Owner:        casdoorGroup.Owner,
		Name:         casdoorGroup.Name,
		CreatedTime:  casdoorGroup.CreatedTime,
		UpdatedTime:  casdoorGroup.UpdatedTime,
		DisplayName:  casdoorGroup.DisplayName,
		Manager:      casdoorGroup.Manager,
		ContactEmail: casdoorGroup.ContactEmail,
		Type:         casdoorGroup.Type,
		ParentId:     casdoorGroup.ParentId,
		IsTopGroup:   casdoorGroup.IsTopGroup,
		Users:        casdoorGroup.Users,
		UsersInfo:    userInfos,
		Title:        casdoorGroup.Title,
		Key:          casdoorGroup.Key,
		IsEnabled:    casdoorGroup.IsEnabled,
	}, nil
}

func (s *groupService) Update(ctx context.Context, group models.GroupInfo) (*models.GroupInfo, error) {
	// Prepare updated group
	updatedGroup := casdoorsdk.Group{
		Owner:        group.Owner,
		Name:         group.Name,
		DisplayName:  group.DisplayName,
		Manager:      group.Manager,
		ContactEmail: group.ContactEmail,
		Type:         group.Type,
		ParentId:     group.ParentId,
		IsTopGroup:   group.IsTopGroup,
		Users:        group.Users,
		Title:        group.Title,
		Key:          group.Key,
		IsEnabled:    group.IsEnabled,
	}

	_, err := s.casdoorClient.UpdateGroup(&updatedGroup)
	if err != nil {
		return nil, err
	}

	// Return the updated group info
	return &models.GroupInfo{
		Owner:        updatedGroup.Owner,
		Name:         updatedGroup.Name,
		CreatedTime:  updatedGroup.CreatedTime,
		UpdatedTime:  updatedGroup.UpdatedTime,
		DisplayName:  updatedGroup.DisplayName,
		Manager:      updatedGroup.Manager,
		ContactEmail: updatedGroup.ContactEmail,
		Type:         updatedGroup.Type,
		ParentId:     updatedGroup.ParentId,
		IsTopGroup:   updatedGroup.IsTopGroup,
		Users:        updatedGroup.Users,
		Title:        updatedGroup.Title,
		Key:          updatedGroup.Key,
		IsEnabled:    updatedGroup.IsEnabled,
	}, nil
}

func (s *groupService) Delete(ctx context.Context, groupName string) error {
	//get group
	group, err := s.casdoorClient.GetGroup(groupName)
	if err != nil {
		return err
	}
	_, err = s.casdoorClient.DeleteGroup(group)

	return err
}
