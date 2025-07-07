package dtos

import (
	"spoke7-go/internal/managment/models"
	"spoke7-go/internal/managment/pb"
)

func GroupInfoToModel(p *pb.GroupInfo) *models.GroupInfo {
	if p == nil {
		return nil
	}
	return &models.GroupInfo{
		Owner:        p.Owner,
		Name:         p.Name,
		CreatedTime:  p.CreatedTime,
		UpdatedTime:  p.UpdatedTime,
		DisplayName:  p.DisplayName,
		Manager:      p.Manager,
		ContactEmail: p.ContactEmail,
		Type:         p.Type,
		ParentId:     p.ParentId,
		IsTopGroup:   p.IsTopGroup,
		Users:        p.Users,
		Title:        p.Title,
		Key:          p.Key,
		//Children:     p.Children,
		IsEnabled: p.IsEnabled,
	}
}



func CreateGroupInfoToModel(p *pb.CreateGroupsRequest) *models.GroupInfo {
	if p == nil {
		return nil
	}
	return &models.GroupInfo{
		Owner:        p.Owner,
		Name:         p.Name,
		CreatedTime:  "",
		UpdatedTime:  "",
		DisplayName:  p.DisplayName,
		Manager:      p.Manager,
		ContactEmail: p.ContactEmail,
		Type:         p.Type,
		ParentId:     p.ParentId,
		IsTopGroup:   p.IsTopGroup,
		Users:        p.Users,
		Title:        p.Title,
		Key:          p.Key,
		//Children:     p.Children,
		IsEnabled: p.IsEnabled,
	}
}

func GroupInfoModelToProto(m *models.GroupInfo) *pb.GroupInfo {
	if m == nil {
		return nil
	}

	userInfoProtos := make([]*pb.UserInfo, 0, len(m.UsersInfo))
	for _, user := range m.UsersInfo {
		userInfoProto := UserInfoModelToProto(&user)
		userInfoProtos = append(userInfoProtos, userInfoProto)
	}
	return &pb.GroupInfo{
		Owner:        m.Owner,
		Name:         m.Name,
		CreatedTime:  m.CreatedTime,
		UpdatedTime:  m.UpdatedTime,
		DisplayName:  m.DisplayName,
		Manager:      m.Manager,
		ContactEmail: m.ContactEmail,
		Type:         m.Type,
		ParentId:     m.ParentId,
		IsTopGroup:   m.IsTopGroup,
		Users:        m.Users,
		UsersInfo:    userInfoProtos,
		Title:        m.Title,
		Key:          m.Key,
		IsEnabled:    m.IsEnabled,
		//		Children:     []*pb.GroupInfo{},
	}
}

func UpdateGroupInfoToModel(req *pb.UpdateGroupRequest) *models.GroupInfo {
	return &models.GroupInfo{
		Name:         req.Name,
		DisplayName:  req.DisplayName,
		Manager:      req.Manager,
		ContactEmail: req.ContactEmail,
		Type:         req.Type,
		ParentId:     req.ParentId,
		IsTopGroup:   req.IsTopGroup,
		Users:        req.Users,
		Title:        req.Title,
		Key:          req.Key,
		IsEnabled:    req.IsEnabled,
	}
}
