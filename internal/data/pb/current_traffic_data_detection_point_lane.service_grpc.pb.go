// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.1
// source: services/current_traffic_data_detection_point_lane.service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CurrentTrafficDataByDetectionPointByLaneService_ListCurrentTrafficDataByDetectionPointByLane_FullMethodName          = "/pb.CurrentTrafficDataByDetectionPointByLaneService/ListCurrentTrafficDataByDetectionPointByLane"
	CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLane_FullMethodName           = "/pb.CurrentTrafficDataByDetectionPointByLaneService/GetCurrentTrafficDataByDetectionPointByLane"
	CurrentTrafficDataByDetectionPointByLaneService_CreateCurrentTrafficDataByDetectionPointByLane_FullMethodName        = "/pb.CurrentTrafficDataByDetectionPointByLaneService/CreateCurrentTrafficDataByDetectionPointByLane"
	CurrentTrafficDataByDetectionPointByLaneService_UpdateCurrentTrafficDataByDetectionPointByLane_FullMethodName        = "/pb.CurrentTrafficDataByDetectionPointByLaneService/UpdateCurrentTrafficDataByDetectionPointByLane"
	CurrentTrafficDataByDetectionPointByLaneService_DeleteCurrentTrafficDataByDetectionPointByLane_FullMethodName        = "/pb.CurrentTrafficDataByDetectionPointByLaneService/DeleteCurrentTrafficDataByDetectionPointByLane"
	CurrentTrafficDataByDetectionPointByLaneService_BulkCreateCurrentTrafficDataByDetectionPointByLane_FullMethodName    = "/pb.CurrentTrafficDataByDetectionPointByLaneService/BulkCreateCurrentTrafficDataByDetectionPointByLane"
	CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLaneStatistics_FullMethodName = "/pb.CurrentTrafficDataByDetectionPointByLaneService/GetCurrentTrafficDataByDetectionPointByLaneStatistics"
	CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLaneAggregate_FullMethodName  = "/pb.CurrentTrafficDataByDetectionPointByLaneService/GetCurrentTrafficDataByDetectionPointByLaneAggregate"
	CurrentTrafficDataByDetectionPointByLaneService_ListCurrentTrafficDataByDetectionPointByLaneDaily_FullMethodName     = "/pb.CurrentTrafficDataByDetectionPointByLaneService/ListCurrentTrafficDataByDetectionPointByLaneDaily"
	CurrentTrafficDataByDetectionPointByLaneService_DownloadCurrentTrafficDataByDetectionPointByLane_FullMethodName      = "/pb.CurrentTrafficDataByDetectionPointByLaneService/DownloadCurrentTrafficDataByDetectionPointByLane"
)

// CurrentTrafficDataByDetectionPointByLaneServiceClient is the client API for CurrentTrafficDataByDetectionPointByLaneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrentTrafficDataByDetectionPointByLaneServiceClient interface {
	ListCurrentTrafficDataByDetectionPointByLane(ctx context.Context, in *ListTrafficDataByDetectionPointByLaneRequest, opts ...grpc.CallOption) (*ListCurrentTrafficDataByDetectionPointByLaneResponse, error)
	GetCurrentTrafficDataByDetectionPointByLane(ctx context.Context, in *GetTrafficDataByDetectionPointByLaneRequest, opts ...grpc.CallOption) (*GetCurrentTrafficDataByDetectionPointByLaneResponse, error)
	CreateCurrentTrafficDataByDetectionPointByLane(ctx context.Context, in *CreateCurrentTrafficDataByDetectionPointByLaneRequest, opts ...grpc.CallOption) (*CreateCurrentTrafficDataByDetectionPointByLaneResponse, error)
	UpdateCurrentTrafficDataByDetectionPointByLane(ctx context.Context, in *UpdateCurrentTrafficDataByDetectionPointByLaneRequest, opts ...grpc.CallOption) (*UpdateCurrentTrafficDataByDetectionPointByLaneResponse, error)
	DeleteCurrentTrafficDataByDetectionPointByLane(ctx context.Context, in *DeleteTrafficDataByDetectionPointByLaneRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BulkCreateCurrentTrafficDataByDetectionPointByLane(ctx context.Context, in *BulkCreateTrafficDataByDetectionPointByLaneRequest, opts ...grpc.CallOption) (*BulkCreateTrafficDataByDetectionPointByLaneResponse, error)
	GetCurrentTrafficDataByDetectionPointByLaneStatistics(ctx context.Context, in *GetTrafficDataByDetectionPointByLaneStatisticsRequest, opts ...grpc.CallOption) (*GetCurrentTrafficDataByDetectionPointByLaneStatisticsResponse, error)
	GetCurrentTrafficDataByDetectionPointByLaneAggregate(ctx context.Context, in *GetTrafficDataByDetectionPointByLaneAggregateRequest, opts ...grpc.CallOption) (*GetCurrentTrafficDataByDetectionPointByLaneAggregateResponse, error)
	// Get Traffic Data Group by Day
	ListCurrentTrafficDataByDetectionPointByLaneDaily(ctx context.Context, in *ListTrafficDataByDetectionPointByLaneDailyRequest, opts ...grpc.CallOption) (*ListCurrentTrafficDataByDetectionPointByLaneDailyResponse, error)
	DownloadCurrentTrafficDataByDetectionPointByLane(ctx context.Context, in *DownloadTrafficDataByDetectionPointByLaneRequest, opts ...grpc.CallOption) (*DownloadResponse, error)
}

type currentTrafficDataByDetectionPointByLaneServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrentTrafficDataByDetectionPointByLaneServiceClient(cc grpc.ClientConnInterface) CurrentTrafficDataByDetectionPointByLaneServiceClient {
	return &currentTrafficDataByDetectionPointByLaneServiceClient{cc}
}

func (c *currentTrafficDataByDetectionPointByLaneServiceClient) ListCurrentTrafficDataByDetectionPointByLane(ctx context.Context, in *ListTrafficDataByDetectionPointByLaneRequest, opts ...grpc.CallOption) (*ListCurrentTrafficDataByDetectionPointByLaneResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCurrentTrafficDataByDetectionPointByLaneResponse)
	err := c.cc.Invoke(ctx, CurrentTrafficDataByDetectionPointByLaneService_ListCurrentTrafficDataByDetectionPointByLane_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentTrafficDataByDetectionPointByLaneServiceClient) GetCurrentTrafficDataByDetectionPointByLane(ctx context.Context, in *GetTrafficDataByDetectionPointByLaneRequest, opts ...grpc.CallOption) (*GetCurrentTrafficDataByDetectionPointByLaneResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCurrentTrafficDataByDetectionPointByLaneResponse)
	err := c.cc.Invoke(ctx, CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLane_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentTrafficDataByDetectionPointByLaneServiceClient) CreateCurrentTrafficDataByDetectionPointByLane(ctx context.Context, in *CreateCurrentTrafficDataByDetectionPointByLaneRequest, opts ...grpc.CallOption) (*CreateCurrentTrafficDataByDetectionPointByLaneResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCurrentTrafficDataByDetectionPointByLaneResponse)
	err := c.cc.Invoke(ctx, CurrentTrafficDataByDetectionPointByLaneService_CreateCurrentTrafficDataByDetectionPointByLane_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentTrafficDataByDetectionPointByLaneServiceClient) UpdateCurrentTrafficDataByDetectionPointByLane(ctx context.Context, in *UpdateCurrentTrafficDataByDetectionPointByLaneRequest, opts ...grpc.CallOption) (*UpdateCurrentTrafficDataByDetectionPointByLaneResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCurrentTrafficDataByDetectionPointByLaneResponse)
	err := c.cc.Invoke(ctx, CurrentTrafficDataByDetectionPointByLaneService_UpdateCurrentTrafficDataByDetectionPointByLane_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentTrafficDataByDetectionPointByLaneServiceClient) DeleteCurrentTrafficDataByDetectionPointByLane(ctx context.Context, in *DeleteTrafficDataByDetectionPointByLaneRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CurrentTrafficDataByDetectionPointByLaneService_DeleteCurrentTrafficDataByDetectionPointByLane_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentTrafficDataByDetectionPointByLaneServiceClient) BulkCreateCurrentTrafficDataByDetectionPointByLane(ctx context.Context, in *BulkCreateTrafficDataByDetectionPointByLaneRequest, opts ...grpc.CallOption) (*BulkCreateTrafficDataByDetectionPointByLaneResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BulkCreateTrafficDataByDetectionPointByLaneResponse)
	err := c.cc.Invoke(ctx, CurrentTrafficDataByDetectionPointByLaneService_BulkCreateCurrentTrafficDataByDetectionPointByLane_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentTrafficDataByDetectionPointByLaneServiceClient) GetCurrentTrafficDataByDetectionPointByLaneStatistics(ctx context.Context, in *GetTrafficDataByDetectionPointByLaneStatisticsRequest, opts ...grpc.CallOption) (*GetCurrentTrafficDataByDetectionPointByLaneStatisticsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCurrentTrafficDataByDetectionPointByLaneStatisticsResponse)
	err := c.cc.Invoke(ctx, CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLaneStatistics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentTrafficDataByDetectionPointByLaneServiceClient) GetCurrentTrafficDataByDetectionPointByLaneAggregate(ctx context.Context, in *GetTrafficDataByDetectionPointByLaneAggregateRequest, opts ...grpc.CallOption) (*GetCurrentTrafficDataByDetectionPointByLaneAggregateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCurrentTrafficDataByDetectionPointByLaneAggregateResponse)
	err := c.cc.Invoke(ctx, CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLaneAggregate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentTrafficDataByDetectionPointByLaneServiceClient) ListCurrentTrafficDataByDetectionPointByLaneDaily(ctx context.Context, in *ListTrafficDataByDetectionPointByLaneDailyRequest, opts ...grpc.CallOption) (*ListCurrentTrafficDataByDetectionPointByLaneDailyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCurrentTrafficDataByDetectionPointByLaneDailyResponse)
	err := c.cc.Invoke(ctx, CurrentTrafficDataByDetectionPointByLaneService_ListCurrentTrafficDataByDetectionPointByLaneDaily_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentTrafficDataByDetectionPointByLaneServiceClient) DownloadCurrentTrafficDataByDetectionPointByLane(ctx context.Context, in *DownloadTrafficDataByDetectionPointByLaneRequest, opts ...grpc.CallOption) (*DownloadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DownloadResponse)
	err := c.cc.Invoke(ctx, CurrentTrafficDataByDetectionPointByLaneService_DownloadCurrentTrafficDataByDetectionPointByLane_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrentTrafficDataByDetectionPointByLaneServiceServer is the server API for CurrentTrafficDataByDetectionPointByLaneService service.
// All implementations must embed UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer
// for forward compatibility.
type CurrentTrafficDataByDetectionPointByLaneServiceServer interface {
	ListCurrentTrafficDataByDetectionPointByLane(context.Context, *ListTrafficDataByDetectionPointByLaneRequest) (*ListCurrentTrafficDataByDetectionPointByLaneResponse, error)
	GetCurrentTrafficDataByDetectionPointByLane(context.Context, *GetTrafficDataByDetectionPointByLaneRequest) (*GetCurrentTrafficDataByDetectionPointByLaneResponse, error)
	CreateCurrentTrafficDataByDetectionPointByLane(context.Context, *CreateCurrentTrafficDataByDetectionPointByLaneRequest) (*CreateCurrentTrafficDataByDetectionPointByLaneResponse, error)
	UpdateCurrentTrafficDataByDetectionPointByLane(context.Context, *UpdateCurrentTrafficDataByDetectionPointByLaneRequest) (*UpdateCurrentTrafficDataByDetectionPointByLaneResponse, error)
	DeleteCurrentTrafficDataByDetectionPointByLane(context.Context, *DeleteTrafficDataByDetectionPointByLaneRequest) (*emptypb.Empty, error)
	BulkCreateCurrentTrafficDataByDetectionPointByLane(context.Context, *BulkCreateTrafficDataByDetectionPointByLaneRequest) (*BulkCreateTrafficDataByDetectionPointByLaneResponse, error)
	GetCurrentTrafficDataByDetectionPointByLaneStatistics(context.Context, *GetTrafficDataByDetectionPointByLaneStatisticsRequest) (*GetCurrentTrafficDataByDetectionPointByLaneStatisticsResponse, error)
	GetCurrentTrafficDataByDetectionPointByLaneAggregate(context.Context, *GetTrafficDataByDetectionPointByLaneAggregateRequest) (*GetCurrentTrafficDataByDetectionPointByLaneAggregateResponse, error)
	// Get Traffic Data Group by Day
	ListCurrentTrafficDataByDetectionPointByLaneDaily(context.Context, *ListTrafficDataByDetectionPointByLaneDailyRequest) (*ListCurrentTrafficDataByDetectionPointByLaneDailyResponse, error)
	DownloadCurrentTrafficDataByDetectionPointByLane(context.Context, *DownloadTrafficDataByDetectionPointByLaneRequest) (*DownloadResponse, error)
	mustEmbedUnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer()
}

// UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer struct{}

func (UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer) ListCurrentTrafficDataByDetectionPointByLane(context.Context, *ListTrafficDataByDetectionPointByLaneRequest) (*ListCurrentTrafficDataByDetectionPointByLaneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCurrentTrafficDataByDetectionPointByLane not implemented")
}
func (UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer) GetCurrentTrafficDataByDetectionPointByLane(context.Context, *GetTrafficDataByDetectionPointByLaneRequest) (*GetCurrentTrafficDataByDetectionPointByLaneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentTrafficDataByDetectionPointByLane not implemented")
}
func (UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer) CreateCurrentTrafficDataByDetectionPointByLane(context.Context, *CreateCurrentTrafficDataByDetectionPointByLaneRequest) (*CreateCurrentTrafficDataByDetectionPointByLaneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCurrentTrafficDataByDetectionPointByLane not implemented")
}
func (UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer) UpdateCurrentTrafficDataByDetectionPointByLane(context.Context, *UpdateCurrentTrafficDataByDetectionPointByLaneRequest) (*UpdateCurrentTrafficDataByDetectionPointByLaneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrentTrafficDataByDetectionPointByLane not implemented")
}
func (UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer) DeleteCurrentTrafficDataByDetectionPointByLane(context.Context, *DeleteTrafficDataByDetectionPointByLaneRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCurrentTrafficDataByDetectionPointByLane not implemented")
}
func (UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer) BulkCreateCurrentTrafficDataByDetectionPointByLane(context.Context, *BulkCreateTrafficDataByDetectionPointByLaneRequest) (*BulkCreateTrafficDataByDetectionPointByLaneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkCreateCurrentTrafficDataByDetectionPointByLane not implemented")
}
func (UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer) GetCurrentTrafficDataByDetectionPointByLaneStatistics(context.Context, *GetTrafficDataByDetectionPointByLaneStatisticsRequest) (*GetCurrentTrafficDataByDetectionPointByLaneStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentTrafficDataByDetectionPointByLaneStatistics not implemented")
}
func (UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer) GetCurrentTrafficDataByDetectionPointByLaneAggregate(context.Context, *GetTrafficDataByDetectionPointByLaneAggregateRequest) (*GetCurrentTrafficDataByDetectionPointByLaneAggregateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentTrafficDataByDetectionPointByLaneAggregate not implemented")
}
func (UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer) ListCurrentTrafficDataByDetectionPointByLaneDaily(context.Context, *ListTrafficDataByDetectionPointByLaneDailyRequest) (*ListCurrentTrafficDataByDetectionPointByLaneDailyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCurrentTrafficDataByDetectionPointByLaneDaily not implemented")
}
func (UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer) DownloadCurrentTrafficDataByDetectionPointByLane(context.Context, *DownloadTrafficDataByDetectionPointByLaneRequest) (*DownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadCurrentTrafficDataByDetectionPointByLane not implemented")
}
func (UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer) mustEmbedUnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer() {
}
func (UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer) testEmbeddedByValue() {}

// UnsafeCurrentTrafficDataByDetectionPointByLaneServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrentTrafficDataByDetectionPointByLaneServiceServer will
// result in compilation errors.
type UnsafeCurrentTrafficDataByDetectionPointByLaneServiceServer interface {
	mustEmbedUnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer()
}

func RegisterCurrentTrafficDataByDetectionPointByLaneServiceServer(s grpc.ServiceRegistrar, srv CurrentTrafficDataByDetectionPointByLaneServiceServer) {
	// If the following call pancis, it indicates UnimplementedCurrentTrafficDataByDetectionPointByLaneServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CurrentTrafficDataByDetectionPointByLaneService_ServiceDesc, srv)
}

func _CurrentTrafficDataByDetectionPointByLaneService_ListCurrentTrafficDataByDetectionPointByLane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTrafficDataByDetectionPointByLaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).ListCurrentTrafficDataByDetectionPointByLane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentTrafficDataByDetectionPointByLaneService_ListCurrentTrafficDataByDetectionPointByLane_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).ListCurrentTrafficDataByDetectionPointByLane(ctx, req.(*ListTrafficDataByDetectionPointByLaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrafficDataByDetectionPointByLaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).GetCurrentTrafficDataByDetectionPointByLane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLane_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).GetCurrentTrafficDataByDetectionPointByLane(ctx, req.(*GetTrafficDataByDetectionPointByLaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentTrafficDataByDetectionPointByLaneService_CreateCurrentTrafficDataByDetectionPointByLane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCurrentTrafficDataByDetectionPointByLaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).CreateCurrentTrafficDataByDetectionPointByLane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentTrafficDataByDetectionPointByLaneService_CreateCurrentTrafficDataByDetectionPointByLane_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).CreateCurrentTrafficDataByDetectionPointByLane(ctx, req.(*CreateCurrentTrafficDataByDetectionPointByLaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentTrafficDataByDetectionPointByLaneService_UpdateCurrentTrafficDataByDetectionPointByLane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurrentTrafficDataByDetectionPointByLaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).UpdateCurrentTrafficDataByDetectionPointByLane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentTrafficDataByDetectionPointByLaneService_UpdateCurrentTrafficDataByDetectionPointByLane_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).UpdateCurrentTrafficDataByDetectionPointByLane(ctx, req.(*UpdateCurrentTrafficDataByDetectionPointByLaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentTrafficDataByDetectionPointByLaneService_DeleteCurrentTrafficDataByDetectionPointByLane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTrafficDataByDetectionPointByLaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).DeleteCurrentTrafficDataByDetectionPointByLane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentTrafficDataByDetectionPointByLaneService_DeleteCurrentTrafficDataByDetectionPointByLane_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).DeleteCurrentTrafficDataByDetectionPointByLane(ctx, req.(*DeleteTrafficDataByDetectionPointByLaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentTrafficDataByDetectionPointByLaneService_BulkCreateCurrentTrafficDataByDetectionPointByLane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkCreateTrafficDataByDetectionPointByLaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).BulkCreateCurrentTrafficDataByDetectionPointByLane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentTrafficDataByDetectionPointByLaneService_BulkCreateCurrentTrafficDataByDetectionPointByLane_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).BulkCreateCurrentTrafficDataByDetectionPointByLane(ctx, req.(*BulkCreateTrafficDataByDetectionPointByLaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLaneStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrafficDataByDetectionPointByLaneStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).GetCurrentTrafficDataByDetectionPointByLaneStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLaneStatistics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).GetCurrentTrafficDataByDetectionPointByLaneStatistics(ctx, req.(*GetTrafficDataByDetectionPointByLaneStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLaneAggregate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrafficDataByDetectionPointByLaneAggregateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).GetCurrentTrafficDataByDetectionPointByLaneAggregate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLaneAggregate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).GetCurrentTrafficDataByDetectionPointByLaneAggregate(ctx, req.(*GetTrafficDataByDetectionPointByLaneAggregateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentTrafficDataByDetectionPointByLaneService_ListCurrentTrafficDataByDetectionPointByLaneDaily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTrafficDataByDetectionPointByLaneDailyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).ListCurrentTrafficDataByDetectionPointByLaneDaily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentTrafficDataByDetectionPointByLaneService_ListCurrentTrafficDataByDetectionPointByLaneDaily_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).ListCurrentTrafficDataByDetectionPointByLaneDaily(ctx, req.(*ListTrafficDataByDetectionPointByLaneDailyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentTrafficDataByDetectionPointByLaneService_DownloadCurrentTrafficDataByDetectionPointByLane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadTrafficDataByDetectionPointByLaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).DownloadCurrentTrafficDataByDetectionPointByLane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentTrafficDataByDetectionPointByLaneService_DownloadCurrentTrafficDataByDetectionPointByLane_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentTrafficDataByDetectionPointByLaneServiceServer).DownloadCurrentTrafficDataByDetectionPointByLane(ctx, req.(*DownloadTrafficDataByDetectionPointByLaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CurrentTrafficDataByDetectionPointByLaneService_ServiceDesc is the grpc.ServiceDesc for CurrentTrafficDataByDetectionPointByLaneService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrentTrafficDataByDetectionPointByLaneService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CurrentTrafficDataByDetectionPointByLaneService",
	HandlerType: (*CurrentTrafficDataByDetectionPointByLaneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCurrentTrafficDataByDetectionPointByLane",
			Handler:    _CurrentTrafficDataByDetectionPointByLaneService_ListCurrentTrafficDataByDetectionPointByLane_Handler,
		},
		{
			MethodName: "GetCurrentTrafficDataByDetectionPointByLane",
			Handler:    _CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLane_Handler,
		},
		{
			MethodName: "CreateCurrentTrafficDataByDetectionPointByLane",
			Handler:    _CurrentTrafficDataByDetectionPointByLaneService_CreateCurrentTrafficDataByDetectionPointByLane_Handler,
		},
		{
			MethodName: "UpdateCurrentTrafficDataByDetectionPointByLane",
			Handler:    _CurrentTrafficDataByDetectionPointByLaneService_UpdateCurrentTrafficDataByDetectionPointByLane_Handler,
		},
		{
			MethodName: "DeleteCurrentTrafficDataByDetectionPointByLane",
			Handler:    _CurrentTrafficDataByDetectionPointByLaneService_DeleteCurrentTrafficDataByDetectionPointByLane_Handler,
		},
		{
			MethodName: "BulkCreateCurrentTrafficDataByDetectionPointByLane",
			Handler:    _CurrentTrafficDataByDetectionPointByLaneService_BulkCreateCurrentTrafficDataByDetectionPointByLane_Handler,
		},
		{
			MethodName: "GetCurrentTrafficDataByDetectionPointByLaneStatistics",
			Handler:    _CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLaneStatistics_Handler,
		},
		{
			MethodName: "GetCurrentTrafficDataByDetectionPointByLaneAggregate",
			Handler:    _CurrentTrafficDataByDetectionPointByLaneService_GetCurrentTrafficDataByDetectionPointByLaneAggregate_Handler,
		},
		{
			MethodName: "ListCurrentTrafficDataByDetectionPointByLaneDaily",
			Handler:    _CurrentTrafficDataByDetectionPointByLaneService_ListCurrentTrafficDataByDetectionPointByLaneDaily_Handler,
		},
		{
			MethodName: "DownloadCurrentTrafficDataByDetectionPointByLane",
			Handler:    _CurrentTrafficDataByDetectionPointByLaneService_DownloadCurrentTrafficDataByDetectionPointByLane_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/current_traffic_data_detection_point_lane.service.proto",
}
