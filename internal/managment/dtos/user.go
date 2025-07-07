package dtos

import (
	"spoke7-go/internal/managment/models"
	"spoke7-go/internal/managment/pb"
)

func UserInfoToModel(p *pb.UserInfo) *models.UserInfo {
	if p == nil {
		return nil
	}
	return &models.UserInfo{
		Sub:         p.GetSub(),
		Iss:         p.GetIss(),
		Aud:         p.GetAud(),
		Name:        p.GetName(),
		DisplayName: p.GetDisplayName(),
		Email:       p.GetEmail(),
		Avatar:      p.GetAvatar(),
		Address:     p.GetAddress(),
		Phone:       p.GetPhone(),
		Groups:      p.GetGroups(),
		Owner:       p.GetOwner(),
	}
}

func UserInfoModelToProto(m *models.UserInfo) *pb.UserInfo {
	if m == nil {
		return nil
	}
	return &pb.UserInfo{
		Sub:         m.Sub,
		Iss:         m.Iss,
		Aud:         m.Aud,
		Name:        m.Name,
		DisplayName: m.DisplayName,
		Email:       m.Email,
		Avatar:      m.Avatar,
		Address:     m.Address,
		Phone:       m.Phone,
		Groups:      m.Groups,
		Owner:       m.Owner,
	}
}
