package controller

import (
	"context"
	"spoke7-go/internal/managment/dtos"
	"spoke7-go/internal/managment/pb"
	"spoke7-go/internal/managment/services"
	"spoke7-go/pkg/authz"
	"spoke7-go/pkg/logger"

	"google.golang.org/grpc"
)

var (
	_ pb.RoleServiceServer = (*roleController)(nil)
	_ authz.AuthResolver   = (*roleController)(nil)
)

type roleController struct {
	service services.RoleService
	logger  logger.Logger
	pb.UnimplementedRoleServiceServer
}

// GetObjectAndActionFromRequest implements authz.AuthResolver.
func (u *roleController) GetObjectAndActionFromRequest(ctx context.Context, req any, info *grpc.UnaryServerInfo) (*authz.Object, string, error) {
	return authz.GenerateObjectActionFromProto(info.FullMethod)
}

func NewRoleController(service services.RoleService, logger logger.Logger) pb.RoleServiceServer {
	return &roleController{
		service: service,
		logger:  logger,
	}
}

func (u *roleController) ListRoles(ctx context.Context, req *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	data, err := u.service.ListRoles()

	if err != nil {
		return nil, err
	}

	return &pb.ListRolesResponse{
		Roles: data,
	}, nil
}

func (u *roleController) GetRole(ctx context.Context, req *pb.GetRoleRequest) (*pb.RoleInfo, error) {
	data, err := u.service.GetRole(req.Name)
	if err != nil {
		return nil, err
	}

	roleDto := dtos.RoleToProto(data)
	return roleDto, nil
}

func (u *roleController) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.RoleInfo, error) {
	permissions := dtos.PermissionsToModels(req.Permissions)

	role, err := u.service.CreateRole(ctx, req.Name, permissions)
	if err != nil {
		return nil, err
	}

	return &pb.RoleInfo{
		Name:        role.Name,
		Permissions: dtos.PermissionsToProtos(role.Permissions),
	}, nil
}

func (u *roleController) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.RoleInfo, error) {
	permissions := dtos.PermissionsToModels(req.Permissions)

	role, err := u.service.UpdateRole(ctx, req.Name, permissions)
	if err != nil {
		return nil, err
	}

	return &pb.RoleInfo{
		Name:        role.Name,
		Permissions: dtos.PermissionsToProtos(role.Permissions),
	}, nil
}

func (u *roleController) ListPermissions(ctx context.Context, req *pb.ListPermissionsRequest) (*pb.ListPermissionsResponse, error) {
	data, err := u.service.ListPermissions()
	if err != nil {
		return nil, err
	}

	permissionFeatureProtos := make([]*pb.PermissionFeature, 0)
	for _, feature := range data {
		permissionFeatureProtos = append(permissionFeatureProtos, dtos.PermissionFeatureToProto(feature))
	}

	return &pb.ListPermissionsResponse{
		Permissions: permissionFeatureProtos,
	}, nil
}
