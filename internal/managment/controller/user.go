package controller

import (
	"context"
	"spoke7-go/internal/managment/dtos"
	"spoke7-go/internal/managment/pb"
	"spoke7-go/internal/managment/services"
	"spoke7-go/pkg/authz"
	"spoke7-go/pkg/logger"

	"github.com/casbin/casbin/v2"
	"google.golang.org/grpc"
)

var (
	_ pb.UserServiceServer = (*userControlller)(nil)
	_ authz.AuthResolver   = (*userControlller)(nil)
)

type userControlller struct {
	service  services.UserService
	enforcer casbin.IEnforcer
	logger   logger.Logger
	pb.UnimplementedUserServiceServer
}

// GetObjectAndActionFromRequest implements authz.AuthResolver.
func (u *userControlller) GetObjectAndActionFromRequest(ctx context.Context, req any, info *grpc.UnaryServerInfo) (*authz.Object, string, error) {
	return authz.GenerateObjectActionFromProto(info.FullMethod)
}

func NewUserController(service services.UserService, enforcer casbin.IEnforcer, logger logger.Logger) pb.UserServiceServer {

	return &userControlller{
		service:  service,
		enforcer: enforcer,
		logger:   logger,
	}
}

// TODO CHECK PERMISSIONS

func (u *userControlller) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {

	//object := proto.GetExtension( , pb.E_MyMethodOptions).(string)
	//u.logger.Infof("object: %s", object)
	// fd := pb.File_services_users_proto
	// methodDesc := fd.Services().ByName("UserService").Methods().ByName("ListUsers").Options()
	// u.logger.Infof("methodDesc: %+v", methodDesc)

	// if ok, err := isAllowed(ctx, u.enforcer, "*", "user", "list"); !ok {
	// 	return nil, err
	// }

	data, err := u.service.List(ctx)

	if err != nil {
		return nil, err
	}

	protoRes := make([]*pb.UserInfo, len(data))
	for key, item := range data {

		protoRes[key] = dtos.UserInfoModelToProto(item)

	}

	return &pb.ListUsersResponse{
		Users: protoRes,
	}, nil
}

// TODO check if the user has permission to list users

func (u *userControlller) CreateUser(ctx context.Context, req *pb.UserInfo) (*pb.UserInfo, error) {
	/*1. convert request dtos to model */
	reqUser := dtos.UserInfoToModel(req)

	userFromContext, err := authz.GetUserFromContext(ctx)
	reqUser.Owner = userFromContext.Username

	/*2. call service to create a users*/
	user, err := u.service.Create(ctx, *reqUser)
	/*3. convert response model to dtos*/
	if err != nil {
		return nil, err
	}

	return dtos.UserInfoModelToProto(user), nil
}

func (u *userControlller) UpdateUser(ctx context.Context, req *pb.UserInfo) (*pb.UserInfo, error) {
	// Convert incoming protobuf request to internal model
	userModel := dtos.UserInfoToModel(req)

	// Update the user using the service
	updatedUser, err := u.service.Update(ctx, *userModel)
	if err != nil {
		u.logger.Errorf("Failed to update user %s: %v", userModel.Name, err)
		return nil, err
	}

	// Return the updated user as a protobuf response
	return dtos.UserInfoModelToProto(updatedUser), nil
}

func (u *userControlller) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := u.service.Delete(ctx, req.Name)
	if err != nil {
		u.logger.Errorf("Failed to delete user %s: %v", req.Name, err)
		return nil, err
	}

	return &pb.DeleteUserResponse{
		Ok: true,
	}, nil
}

func (u *userControlller) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserInfo, error) {
	user, err := u.service.Get(ctx, req.Name)
	if err != nil {
		u.logger.Errorf("Failed to get user %s: %v", req.Name, err)
		return nil, err
	}

	return dtos.UserInfoModelToProto(user), nil
}

// GetUserPermissions implements pb.UserServiceServer.
func (u *userControlller) GetUserPermissions(ctx context.Context, req *pb.GetUserPermissionRequest) (*pb.GetUserPermissionsResponse, error) {
	user, err := u.service.Get(ctx, req.Name)
	if err != nil {
		u.logger.Errorf("Failed to get user: %v", err)
		return nil, err
	}

	permissions, err := u.service.GetPermissions(user.Name, user.Groups)
	if err != nil {
		u.logger.Errorf("Failed to get permissions for user %s: %v", user.Name, err)
		return nil, err
	}

	entries := []*pb.FeaturePermissionEntry{}
	for key, values := range permissions {
		entries = append(entries, &pb.FeaturePermissionEntry{
			Key:    key,
			Values: values,
		})
	}

	return &pb.GetUserPermissionsResponse{
		Permissions: entries,
	}, nil
}

// GetUserPermissions implements pb.UserServiceServer.
func (u *userControlller) GetSelfPermissions(ctx context.Context, _ *pb.GetUserCasbinContextRequest) (*pb.GetUserPermissionsResponse, error) {
	user, err := authz.GetUserFromContext(ctx)
	if err != nil {
		u.logger.Errorf("Failed to get user from context: %v", err)
		return nil, err
	}

	permissions, err := u.service.GetPermissions(user.Username, user.Groups)
	if err != nil {
		u.logger.Errorf("Failed to get permissions for user %s: %v", user.Username, err)
		return nil, err
	}

	entries := []*pb.FeaturePermissionEntry{}
	for key, values := range permissions {
		entries = append(entries, &pb.FeaturePermissionEntry{
			Key:    key,
			Values: values,
		})
	}

	return &pb.GetUserPermissionsResponse{
		Permissions: entries,
	}, nil
}

func (u *userControlller) SetUserPermissions(ctx context.Context, req *pb.SetUserPermissionRequest) (*pb.SetUserPermissionsResponse, error) {
	//check if user exists
	user, err := u.service.Get(ctx, req.Name)
	if err != nil {
		u.logger.Errorf("Failed to get user: %v", err)
		return nil, err
	}

	permissions, err := u.service.SetPermissions(user.Name, req.Permissions)
	if err != nil {
		u.logger.Errorf("Failed to get permissions for user %s: %v", user.Name, err)
		return nil, err
	}

	entries := []*pb.FeaturePermissionEntry{}
	for key, values := range permissions {
		entries = append(entries, &pb.FeaturePermissionEntry{
			Key:    key,
			Values: values,
		})
	}

	return &pb.SetUserPermissionsResponse{
		Permissions: entries,
	}, nil
}
