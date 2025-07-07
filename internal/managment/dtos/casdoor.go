package dtos

import (
	"spoke7-go/internal/managment/models"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

// toGroupInfos converts a slice of casdoorsdk.Group to a slice of models.GroupInfo.
func CasdoorGroupToModel(group *casdoorsdk.Group) *models.GroupInfo {
	return &models.GroupInfo{
		Owner:        group.Owner,
		Name:         group.Name,
		CreatedTime:  group.CreatedTime,
		UpdatedTime:  group.UpdatedTime,
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
}

// toGroupInfos converts a slice of casdoorsdk.Group to a slice of models.GroupInfo.
func CasdoorGroupsToModels(groups []*casdoorsdk.Group) []*models.GroupInfo {
	res := make([]*models.GroupInfo, 0, len(groups))

	for i := range groups {
		res = append(res, CasdoorGroupToModel(groups[i]))
	}

	return res

}
