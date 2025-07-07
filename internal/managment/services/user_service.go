package services

import (
	"context"
	"fmt"
	"strings"

	"spoke7-go/internal/managment/models"

	"github.com/casbin/casbin/v2"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type UserService interface {
	List(ctx context.Context) ([]*models.UserInfo, error)
	Create(ctx context.Context, user models.UserInfo) (*models.UserInfo, error)
	Get(ctx context.Context, userName string) (*models.UserInfo, error)
	Update(ctx context.Context, user models.UserInfo) (*models.UserInfo, error)
	Delete(ctx context.Context, userName string) error
	GetPermissions(name string, groups []string) (map[string][]string, error)
	SetPermissions(name string, permissions []string) (map[string][]string, error)
}
type userService struct {
	casdoorClient    *casdoorsdk.Client
	enforcer         casbin.IEnforcer
	organizationName string
}

func NewUserService(client *casdoorsdk.Client, _enforcer casbin.IEnforcer, _organizationName string) UserService {

	return &userService{
		casdoorClient:    client,
		enforcer:         _enforcer,
		organizationName: _organizationName,
	}
}

func (s *userService) List(ctx context.Context) ([]*models.UserInfo, error) {
	users, err := s.casdoorClient.GetUsers()
	if err != nil {
		return nil, err
	}

	userInfos := make([]*models.UserInfo, 0, len(users))
	for _, user := range users {
		userInfo := &models.UserInfo{
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
			Owner:       user.Owner,
		}
		userInfos = append(userInfos, userInfo)
	}

	return userInfos, nil
}

func (s *userService) Create(ctx context.Context, user models.UserInfo) (*models.UserInfo, error) {
	newUser := casdoorsdk.User{
		Owner:       user.Owner,
		Name:        user.Name,
		DisplayName: user.DisplayName,
		CreatedTime: casdoorsdk.GetCurrentTime(),
		UpdatedTime: casdoorsdk.GetCurrentTime(),
		Email:       user.Email,
		Address:     user.Address,
		Phone:       user.Phone,
		Avatar:      user.Avatar,
		Groups:      user.Groups,
	}

	_, err := s.casdoorClient.AddUser(&newUser)
	if err != nil {
		return nil, err
	}

	casdoorUser, err := s.casdoorClient.GetUser(newUser.Name)
	if err != nil {
		return nil, err
	}

	return &models.UserInfo{
		Sub:         casdoorUser.Id,
		Name:        casdoorUser.Name,
		DisplayName: casdoorUser.DisplayName,
		Email:       casdoorUser.Email,
		Avatar:      casdoorUser.Avatar,
		Address:     casdoorUser.Address,
		Phone:       casdoorUser.Phone,
		Groups:      casdoorUser.Groups,
		Owner:       casdoorUser.Owner,
	}, nil
}

func (s *userService) Get(ctx context.Context, userName string) (*models.UserInfo, error) {
	casdoorUser, err := s.casdoorClient.GetUser(userName)
	if err != nil {
		return nil, err
	}
	if casdoorUser == nil {
		return nil, fmt.Errorf("user not found %s", userName)
	}

	// Creazione di una nuova slice per i gruppi con eventuale trim del prefisso
	var groups []string

	for _, groupStr := range casdoorUser.Groups {
		groupStr = strings.TrimPrefix(groupStr, s.organizationName+"/")
		groups = append(groups, groupStr)

	}

	return &models.UserInfo{
		Sub:         casdoorUser.Id,
		Name:        casdoorUser.Name,
		DisplayName: casdoorUser.DisplayName,
		Email:       casdoorUser.Email,
		Avatar:      casdoorUser.Avatar,
		Address:     casdoorUser.Address,
		Phone:       casdoorUser.Phone,
		Groups:      groups,
		Owner:       casdoorUser.Owner,
	}, nil
}

func (s *userService) Update(ctx context.Context, user models.UserInfo) (*models.UserInfo, error) {
	newUser := casdoorsdk.User{
		Owner:       user.Owner,
		Name:        user.Name,
		DisplayName: user.DisplayName,
		CreatedTime: casdoorsdk.GetCurrentTime(),
		UpdatedTime: casdoorsdk.GetCurrentTime(),
		Email:       user.Email,
		Address:     user.Address,
		Phone:       user.Phone,
		Avatar:      user.Avatar,
		Groups:      user.Groups,
		Id:          user.Sub,
	}

	_, err := s.casdoorClient.UpdateUser(&newUser)
	if err != nil {
		return nil, err
	}

	casdoorUser, err := s.casdoorClient.GetUser(newUser.Name)
	if err != nil {
		return nil, err
	}

	return &models.UserInfo{
		Sub:         casdoorUser.Id,
		Name:        casdoorUser.Name,
		DisplayName: casdoorUser.DisplayName,
		Email:       casdoorUser.Email,
		Avatar:      casdoorUser.Avatar,
		Address:     casdoorUser.Address,
		Phone:       casdoorUser.Phone,
		Groups:      casdoorUser.Groups,
		Owner:       casdoorUser.Owner,
	}, nil
}

func (s *userService) Delete(ctx context.Context, userName string) error {
	//get user
	user, err := s.casdoorClient.GetUser(userName)
	if err != nil {
		return err
	}
	_, err = s.casdoorClient.DeleteUser(user)

	return err
}

func (s *userService) GetPermissions(name string, groups []string) (map[string][]string, error) {
	result := make(map[string][]string)

	// Get user permissions
	userPerms, err := s.enforcer.GetImplicitRolesForUser(name)
	if err != nil {
		return nil, err
	}
	result[name] = userPerms

	// Get group permissions
	for _, group := range groups {
		completeRoleName := fmt.Sprintf("role:%s", group)
		groupPerms, err := s.enforcer.GetImplicitRolesForUser(completeRoleName)
		if err != nil {
			return nil, err
		}
		result[group] = groupPerms
	}

	return result, nil
}
func (s *userService) SetPermissions(name string, permissions []string) (map[string][]string, error) {

	//1. get current user permissons
	currentPermissions, err := s.enforcer.GetFilteredGroupingPolicy(0, name)
	if err != nil {
		return nil, fmt.Errorf("failed to get current policies for role %s: %w", name, err)
	}

	//2. make a distinct of new permissions
	permissionSet := make(map[string]struct{}, len(permissions))
	for _, p := range permissions {
		permissionSet[p] = struct{}{}
	}

	//3. remove removed permission
	for _, p := range currentPermissions {
		if len(p) < 3 {
			continue // malformed policy
		}
		key := p[0]
		if _, ok := permissionSet[key]; !ok {
			// Remove policy: completeRoleName can no longer perform action on type
			_, err = s.enforcer.RemoveGroupingPolicy(name, p[0])
			if err != nil {
				return nil, fmt.Errorf("failed to remove policy for role %s: %w", name, err)
			}
		}
	}

	for _, p := range permissions {
		// Add policy: completeRoleName can perform action on type
		ok, err := s.enforcer.AddGroupingPolicy(name, p)
		if err != nil {
			return nil, fmt.Errorf("failed to add policy for role %s: %w", name, err)
		}
		if !ok {
			// Policy already exists
			continue
		}
	}

	// Optionally persist to storage (if not using AutoSave)
	if err := s.enforcer.SavePolicy(); err != nil {
		return nil, fmt.Errorf("failed to save policy: %w", err)
	}

	return map[string][]string{name: permissions}, nil

}
