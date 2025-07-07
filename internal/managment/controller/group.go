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
	_ pb.GroupServiceServer = (*GroupController)(nil)
	_ authz.AuthResolver    = (*GroupController)(nil)
)

type GroupController struct {
	service services.GroupService
	logger  logger.Logger

	pb.UnimplementedGroupServiceServer
}

func NewGroupController(service services.GroupService, logger logger.Logger) *GroupController {
	return &GroupController{
		service: service,
		logger:  logger,
	}
}

// Get Method and Object Requirment

// func (u *GroupController) GetServiceName() string {
// 	//pb.GroupService
// 	return pb.GroupService_ServiceDesc.ServiceName
// }

// func (u *GroupController) GetPermission() []authz.Permission {

// 	// permission := make([]authz.Permission, 0)
// 	// for _, item := range pb.MethodSecurityScope_name {
// 	// 	//split the item from underscore
// 	// 	item := strings.Split(item, "_")
// 	// 	permission = append(permission, authz.Permission{
// 	// 		Type:   item[0],
// 	// 		Action: item[1],
// 	// 	})

// 	// }
// 	//return permission
// 	// enforcer.AddPolicy("role:admin", item[0], item[1])
// 	// 	logger.Infof("add role:admin policy with : %s %s", item[0], item[1])

// 	return []authz.Permission{
// 		PERMISSION_GROUP_LIST, PERMISSION_GROUP_CREATE, PERMISSION_GROUP_DELETE,
// 	}

// }

type namedRequest interface {
	GetName() string
}

func (u *GroupController) GetObjectAndActionFromRequest(ctx context.Context, req any, info *grpc.UnaryServerInfo) (*authz.Object, string, error) {
	restricted, err := authz.IsRestricted(info.FullMethod)
	if err != nil {
		return nil, "", err
	}

	if restricted {
		namedReq := req.(namedRequest)
		groupInfo, err := u.service.Get(ctx, namedReq.GetName())
		if err != nil {
			return nil, "", err
		}
		return authz.GenerateObjectWithOwnerActionFromProto(info.FullMethod, groupInfo.Manager, []string{groupInfo.Name})
	}

	return authz.GenerateObjectActionFromProto(info.FullMethod)

}

func (u *GroupController) ListGroup(ctx context.Context, req *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	user, err := authz.GetUserFromContext(ctx)
	if err != nil {
		u.logger.Errorf("Failed to get user ")
	}

	data, err := u.service.ListByMember(ctx, user.Username, user.Groups)

	if err != nil {
		return nil, err
	}

	protoRes := make([]*pb.GroupInfo, len(data))
	for key, item := range data {

		protoRes[key] = dtos.GroupInfoModelToProto(item)

	}

	return &pb.ListGroupsResponse{
		Groups: protoRes,
	}, nil
}

func (u *GroupController) ListAllGroup(ctx context.Context, req *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	data, err := u.service.ListAll(ctx)

	if err != nil {
		return nil, err
	}

	protoRes := make([]*pb.GroupInfo, len(data))
	for key, item := range data {

		protoRes[key] = dtos.GroupInfoModelToProto(item)

	}

	return &pb.ListGroupsResponse{
		Groups: protoRes,
	}, nil

}

func (u *GroupController) CreateGroup(ctx context.Context, req *pb.CreateGroupsRequest) (*pb.GroupInfo, error) {
	/*1. convert request dtos to model */
	reqGroup := dtos.CreateGroupInfoToModel(req)

	user, err := authz.GetUserFromContext(ctx)
	reqGroup.Owner = user.Username
	reqGroup.Manager = user.Username

	/*2. call service to create a groups*/
	group, err := u.service.Create(ctx, *reqGroup)
	/*3. convert response model to dtos*/
	if err != nil {
		return nil, err
	}

	return dtos.GroupInfoModelToProto(group), nil
}

func (u *GroupController) UpdateGroup(ctx context.Context, req *pb.UpdateGroupRequest) (*pb.GroupInfo, error) {
	// Convert incoming protobuf request to internal model
	groupModel := dtos.UpdateGroupInfoToModel(req)

	// Update the group using the service
	updatedGroup, err := u.service.Update(ctx, *groupModel)
	if err != nil {
		u.logger.Errorf("Failed to update group %s: %v", groupModel.Name, err)
		return nil, err
	}

	// Return the updated group as a protobuf response
	return dtos.GroupInfoModelToProto(updatedGroup), nil
}

func (u *GroupController) DeleteGroup(ctx context.Context, req *pb.DeleteGroupsRequest) (*pb.DeleteGroupsResponse, error) {
	err := u.service.Delete(ctx, req.Name)
	if err != nil {
		u.logger.Errorf("Failed to delete group %s: %v", req.Name, err)
		return nil, err
	}

	return &pb.DeleteGroupsResponse{
		Ok: true,
	}, nil
}

func (u *GroupController) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.GroupInfo, error) {
	group, err := u.service.Get(ctx, req.Name)
	if err != nil {
		u.logger.Errorf("Failed to get group %s: %v", req.Name, err)
		return nil, err
	}

	return dtos.GroupInfoModelToProto(group), nil
}
