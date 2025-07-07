package dtos

import (
	"spoke7-go/internal/managment/models"
	"spoke7-go/internal/managment/pb"
)

func RoleToModel(r *pb.RoleInfo) *models.Role {
	if r == nil {
		return nil
	}
	return &models.Role{
		Name:        r.Name,
		Permissions: PermissionsToModels(r.Permissions),
	}
}

func RoleToProto(r *models.Role) *pb.RoleInfo {
	if r == nil {
		return nil
	}
	return &pb.RoleInfo{
		Name:        r.Name,
		Permissions: PermissionsToProtos(r.Permissions),
	}
}

func PermissionFeatureToModel(p *pb.PermissionFeature) *models.PermissionFeatures {
	if p == nil {
		return nil
	}
	return &models.PermissionFeatures{
		Feature:     p.Feature,
		Permissions: PermissionsToModels(p.Permissions),
	}
}

func PermissionsToModels(permissionsList []*pb.Permission) []models.Permission {
	permissions := make([]models.Permission, 0, len(permissionsList))
	for _, c := range permissionsList {
		permissions = append(permissions, PermissionToModel(c))
	}
	return permissions
}

func PermissionToModel(p *pb.Permission) models.Permission {
	if p == nil {
		return models.Permission{}
	}
	return models.Permission{
		Type:           p.Resource,
		Action:         p.Action,
		PermissionName: p.PermissionName,
	}
}

func PermissionFeatureToProto(p models.PermissionFeatures) *pb.PermissionFeature {

	return &pb.PermissionFeature{
		Feature:     p.Feature,
		Permissions: PermissionsToProtos(p.Permissions),
	}
}

func PermissionsToProtos(permissionsList []models.Permission) []*pb.Permission {
	permissions := make([]*pb.Permission, 0, len(permissionsList))
	for _, c := range permissionsList {
		permissions = append(permissions, PermissionToProto(c))
	}
	return permissions
}

func PermissionToProto(p models.Permission) *pb.Permission {
	return &pb.Permission{
		Resource:       p.Type,
		Action:         p.Action,
		PermissionName: p.PermissionName,
	}
}
