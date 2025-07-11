// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: services/history_day_traffic_data_detection_point.service.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListHistoryDayTrafficDataByDetectionPointResponse struct {
	state                                  protoimpl.MessageState                   `protogen:"open.v1"`
	HistoryDayTrafficDataByDetectionPoints []*HistoryDayTrafficDataByDetectionPoint `protobuf:"bytes,1,rep,name=history_day_traffic_data_by_detection_points,json=historyDayTrafficDataByDetectionPoints,proto3" json:"history_day_traffic_data_by_detection_points,omitempty"`
	TotalCount                             uint32                                   `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields                          protoimpl.UnknownFields
	sizeCache                              protoimpl.SizeCache
}

func (x *ListHistoryDayTrafficDataByDetectionPointResponse) Reset() {
	*x = ListHistoryDayTrafficDataByDetectionPointResponse{}
	mi := &file_services_history_day_traffic_data_detection_point_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListHistoryDayTrafficDataByDetectionPointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHistoryDayTrafficDataByDetectionPointResponse) ProtoMessage() {}

func (x *ListHistoryDayTrafficDataByDetectionPointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_history_day_traffic_data_detection_point_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHistoryDayTrafficDataByDetectionPointResponse.ProtoReflect.Descriptor instead.
func (*ListHistoryDayTrafficDataByDetectionPointResponse) Descriptor() ([]byte, []int) {
	return file_services_history_day_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListHistoryDayTrafficDataByDetectionPointResponse) GetHistoryDayTrafficDataByDetectionPoints() []*HistoryDayTrafficDataByDetectionPoint {
	if x != nil {
		return x.HistoryDayTrafficDataByDetectionPoints
	}
	return nil
}

func (x *ListHistoryDayTrafficDataByDetectionPointResponse) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type GetHistoryDayTrafficDataByDetectionPointResponse struct {
	state                                 protoimpl.MessageState                   `protogen:"open.v1"`
	HistoryDayTrafficDataByDetectionPoint []*HistoryDayTrafficDataByDetectionPoint `protobuf:"bytes,1,rep,name=history_day_traffic_data_by_detection_point,json=historyDayTrafficDataByDetectionPoint,proto3" json:"history_day_traffic_data_by_detection_point,omitempty"`
	TotalCount                            uint32                                   `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields                         protoimpl.UnknownFields
	sizeCache                             protoimpl.SizeCache
}

func (x *GetHistoryDayTrafficDataByDetectionPointResponse) Reset() {
	*x = GetHistoryDayTrafficDataByDetectionPointResponse{}
	mi := &file_services_history_day_traffic_data_detection_point_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHistoryDayTrafficDataByDetectionPointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoryDayTrafficDataByDetectionPointResponse) ProtoMessage() {}

func (x *GetHistoryDayTrafficDataByDetectionPointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_history_day_traffic_data_detection_point_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoryDayTrafficDataByDetectionPointResponse.ProtoReflect.Descriptor instead.
func (*GetHistoryDayTrafficDataByDetectionPointResponse) Descriptor() ([]byte, []int) {
	return file_services_history_day_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetHistoryDayTrafficDataByDetectionPointResponse) GetHistoryDayTrafficDataByDetectionPoint() []*HistoryDayTrafficDataByDetectionPoint {
	if x != nil {
		return x.HistoryDayTrafficDataByDetectionPoint
	}
	return nil
}

func (x *GetHistoryDayTrafficDataByDetectionPointResponse) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type CreateHistoryDayTrafficDataByDetectionPointRequest struct {
	state                                 protoimpl.MessageState                 `protogen:"open.v1"`
	HistoryDayTrafficDataByDetectionPoint *HistoryDayTrafficDataByDetectionPoint `protobuf:"bytes,1,opt,name=history_day_traffic_data_by_detection_point,json=historyDayTrafficDataByDetectionPoint,proto3" json:"history_day_traffic_data_by_detection_point,omitempty"`
	unknownFields                         protoimpl.UnknownFields
	sizeCache                             protoimpl.SizeCache
}

func (x *CreateHistoryDayTrafficDataByDetectionPointRequest) Reset() {
	*x = CreateHistoryDayTrafficDataByDetectionPointRequest{}
	mi := &file_services_history_day_traffic_data_detection_point_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateHistoryDayTrafficDataByDetectionPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHistoryDayTrafficDataByDetectionPointRequest) ProtoMessage() {}

func (x *CreateHistoryDayTrafficDataByDetectionPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_history_day_traffic_data_detection_point_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHistoryDayTrafficDataByDetectionPointRequest.ProtoReflect.Descriptor instead.
func (*CreateHistoryDayTrafficDataByDetectionPointRequest) Descriptor() ([]byte, []int) {
	return file_services_history_day_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateHistoryDayTrafficDataByDetectionPointRequest) GetHistoryDayTrafficDataByDetectionPoint() *HistoryDayTrafficDataByDetectionPoint {
	if x != nil {
		return x.HistoryDayTrafficDataByDetectionPoint
	}
	return nil
}

type CreateHistoryDayTrafficDataByDetectionPointResponse struct {
	state                                 protoimpl.MessageState                 `protogen:"open.v1"`
	HistoryDayTrafficDataByDetectionPoint *HistoryDayTrafficDataByDetectionPoint `protobuf:"bytes,1,opt,name=history_day_traffic_data_by_detection_point,json=historyDayTrafficDataByDetectionPoint,proto3" json:"history_day_traffic_data_by_detection_point,omitempty"`
	unknownFields                         protoimpl.UnknownFields
	sizeCache                             protoimpl.SizeCache
}

func (x *CreateHistoryDayTrafficDataByDetectionPointResponse) Reset() {
	*x = CreateHistoryDayTrafficDataByDetectionPointResponse{}
	mi := &file_services_history_day_traffic_data_detection_point_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateHistoryDayTrafficDataByDetectionPointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHistoryDayTrafficDataByDetectionPointResponse) ProtoMessage() {}

func (x *CreateHistoryDayTrafficDataByDetectionPointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_history_day_traffic_data_detection_point_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHistoryDayTrafficDataByDetectionPointResponse.ProtoReflect.Descriptor instead.
func (*CreateHistoryDayTrafficDataByDetectionPointResponse) Descriptor() ([]byte, []int) {
	return file_services_history_day_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateHistoryDayTrafficDataByDetectionPointResponse) GetHistoryDayTrafficDataByDetectionPoint() *HistoryDayTrafficDataByDetectionPoint {
	if x != nil {
		return x.HistoryDayTrafficDataByDetectionPoint
	}
	return nil
}

type UpdateHistoryDayTrafficDataByDetectionPointRequest struct {
	state                                 protoimpl.MessageState                 `protogen:"open.v1"`
	HistoryDayTrafficDataByDetectionPoint *HistoryDayTrafficDataByDetectionPoint `protobuf:"bytes,1,opt,name=history_day_traffic_data_by_detection_point,json=historyDayTrafficDataByDetectionPoint,proto3" json:"history_day_traffic_data_by_detection_point,omitempty"`
	unknownFields                         protoimpl.UnknownFields
	sizeCache                             protoimpl.SizeCache
}

func (x *UpdateHistoryDayTrafficDataByDetectionPointRequest) Reset() {
	*x = UpdateHistoryDayTrafficDataByDetectionPointRequest{}
	mi := &file_services_history_day_traffic_data_detection_point_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateHistoryDayTrafficDataByDetectionPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHistoryDayTrafficDataByDetectionPointRequest) ProtoMessage() {}

func (x *UpdateHistoryDayTrafficDataByDetectionPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_history_day_traffic_data_detection_point_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHistoryDayTrafficDataByDetectionPointRequest.ProtoReflect.Descriptor instead.
func (*UpdateHistoryDayTrafficDataByDetectionPointRequest) Descriptor() ([]byte, []int) {
	return file_services_history_day_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateHistoryDayTrafficDataByDetectionPointRequest) GetHistoryDayTrafficDataByDetectionPoint() *HistoryDayTrafficDataByDetectionPoint {
	if x != nil {
		return x.HistoryDayTrafficDataByDetectionPoint
	}
	return nil
}

type UpdateHistoryDayTrafficDataByDetectionPointResponse struct {
	state                                 protoimpl.MessageState                 `protogen:"open.v1"`
	HistoryDayTrafficDataByDetectionPoint *HistoryDayTrafficDataByDetectionPoint `protobuf:"bytes,1,opt,name=history_day_traffic_data_by_detection_point,json=historyDayTrafficDataByDetectionPoint,proto3" json:"history_day_traffic_data_by_detection_point,omitempty"`
	unknownFields                         protoimpl.UnknownFields
	sizeCache                             protoimpl.SizeCache
}

func (x *UpdateHistoryDayTrafficDataByDetectionPointResponse) Reset() {
	*x = UpdateHistoryDayTrafficDataByDetectionPointResponse{}
	mi := &file_services_history_day_traffic_data_detection_point_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateHistoryDayTrafficDataByDetectionPointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHistoryDayTrafficDataByDetectionPointResponse) ProtoMessage() {}

func (x *UpdateHistoryDayTrafficDataByDetectionPointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_history_day_traffic_data_detection_point_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHistoryDayTrafficDataByDetectionPointResponse.ProtoReflect.Descriptor instead.
func (*UpdateHistoryDayTrafficDataByDetectionPointResponse) Descriptor() ([]byte, []int) {
	return file_services_history_day_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateHistoryDayTrafficDataByDetectionPointResponse) GetHistoryDayTrafficDataByDetectionPoint() *HistoryDayTrafficDataByDetectionPoint {
	if x != nil {
		return x.HistoryDayTrafficDataByDetectionPoint
	}
	return nil
}

type ListHistoryDayTrafficDataByDetectionPointDailyResponse struct {
	state         protoimpl.MessageState                   `protogen:"open.v1"`
	DailyStats    []*HistoryDayTrafficDataByDetectionPoint `protobuf:"bytes,1,rep,name=daily_stats,json=dailyStats,proto3" json:"daily_stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListHistoryDayTrafficDataByDetectionPointDailyResponse) Reset() {
	*x = ListHistoryDayTrafficDataByDetectionPointDailyResponse{}
	mi := &file_services_history_day_traffic_data_detection_point_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListHistoryDayTrafficDataByDetectionPointDailyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHistoryDayTrafficDataByDetectionPointDailyResponse) ProtoMessage() {}

func (x *ListHistoryDayTrafficDataByDetectionPointDailyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_history_day_traffic_data_detection_point_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHistoryDayTrafficDataByDetectionPointDailyResponse.ProtoReflect.Descriptor instead.
func (*ListHistoryDayTrafficDataByDetectionPointDailyResponse) Descriptor() ([]byte, []int) {
	return file_services_history_day_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListHistoryDayTrafficDataByDetectionPointDailyResponse) GetDailyStats() []*HistoryDayTrafficDataByDetectionPoint {
	if x != nil {
		return x.DailyStats
	}
	return nil
}

var File_services_history_day_traffic_data_detection_point_service_proto protoreflect.FileDescriptor

const file_services_history_day_traffic_data_detection_point_service_proto_rawDesc = "" +
	"\n" +
	"?services/history_day_traffic_data_detection_point.service.proto\x12\x02pb\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a3dtos/history_day_traffic_data_detection_point.proto\x1a\x1bdtos/traffic_requests.proto\x1a\x1cdtos/traffic_responses.proto\"\xde\x01\n" +
	"1ListHistoryDayTrafficDataByDetectionPointResponse\x12\x87\x01\n" +
	",history_day_traffic_data_by_detection_points\x18\x01 \x03(\v2).pb.HistoryDayTrafficDataByDetectionPointR&historyDayTrafficDataByDetectionPoints\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\rR\n" +
	"totalCount\"\xdb\x01\n" +
	"0GetHistoryDayTrafficDataByDetectionPointResponse\x12\x85\x01\n" +
	"+history_day_traffic_data_by_detection_point\x18\x01 \x03(\v2).pb.HistoryDayTrafficDataByDetectionPointR%historyDayTrafficDataByDetectionPoint\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\rR\n" +
	"totalCount\"\xbc\x01\n" +
	"2CreateHistoryDayTrafficDataByDetectionPointRequest\x12\x85\x01\n" +
	"+history_day_traffic_data_by_detection_point\x18\x01 \x01(\v2).pb.HistoryDayTrafficDataByDetectionPointR%historyDayTrafficDataByDetectionPoint\"\xbd\x01\n" +
	"3CreateHistoryDayTrafficDataByDetectionPointResponse\x12\x85\x01\n" +
	"+history_day_traffic_data_by_detection_point\x18\x01 \x01(\v2).pb.HistoryDayTrafficDataByDetectionPointR%historyDayTrafficDataByDetectionPoint\"\xbc\x01\n" +
	"2UpdateHistoryDayTrafficDataByDetectionPointRequest\x12\x85\x01\n" +
	"+history_day_traffic_data_by_detection_point\x18\x01 \x01(\v2).pb.HistoryDayTrafficDataByDetectionPointR%historyDayTrafficDataByDetectionPoint\"\xbd\x01\n" +
	"3UpdateHistoryDayTrafficDataByDetectionPointResponse\x12\x85\x01\n" +
	"+history_day_traffic_data_by_detection_point\x18\x01 \x01(\v2).pb.HistoryDayTrafficDataByDetectionPointR%historyDayTrafficDataByDetectionPoint\"\x84\x01\n" +
	"6ListHistoryDayTrafficDataByDetectionPointDailyResponse\x12J\n" +
	"\vdaily_stats\x18\x01 \x03(\v2).pb.HistoryDayTrafficDataByDetectionPointR\n" +
	"dailyStats2\x82\x10\n" +
	",HistoryDayTrafficDataByDetectionPointService\x12\xd9\x01\n" +
	")ListHistoryDayTrafficDataByDetectionPoint\x12*.pb.ListTrafficDataByDetectionPointRequest\x1a5.pb.ListHistoryDayTrafficDataByDetectionPointResponse\"I\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x022\x120/api/v1/traffic/history_day/detection_point/list\x12\xd8\x01\n" +
	"(GetHistoryDayTrafficDataByDetectionPoint\x12).pb.GetTrafficDataByDetectionPointRequest\x1a4.pb.GetHistoryDayTrafficDataByDetectionPointResponse\"K\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x024\x122/api/v1/traffic/history_day/detection_point/single\x12\xe7\x01\n" +
	"+CreateHistoryDayTrafficDataByDetectionPoint\x126.pb.CreateHistoryDayTrafficDataByDetectionPointRequest\x1a7.pb.CreateHistoryDayTrafficDataByDetectionPointResponse\"G\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x020:\x01*\"+/api/v1/traffic/history_day/detection_point\x12\xe7\x01\n" +
	"+UpdateHistoryDayTrafficDataByDetectionPoint\x126.pb.UpdateHistoryDayTrafficDataByDetectionPointRequest\x1a7.pb.UpdateHistoryDayTrafficDataByDetectionPointResponse\"G\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x020:\x01*2+/api/v1/traffic/history_day/detection_point\x12\xb9\x01\n" +
	"+DeleteHistoryDayTrafficDataByDetectionPoint\x12,.pb.DeleteTrafficDataByDetectionPointRequest\x1a\x16.google.protobuf.Empty\"D\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x02-*+/api/v1/traffic/history_day/detection_point\x12\xe4\x01\n" +
	"/BulkCreateHistoryDayTrafficDataByDetectionPoint\x120.pb.BulkCreateTrafficDataByDetectionPointRequest\x1a1.pb.BulkCreateTrafficDataByDetectionPointResponse\"L\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x025:\x01*\"0/api/v1/traffic/history_day/detection_point/bulk\x12\xf0\x01\n" +
	"2GetHistoryDayTrafficDataByDetectionPointStatistics\x123.pb.GetTrafficDataByDetectionPointStatisticsRequest\x1a4.pb.GetTrafficDataByDetectionPointStatisticsResponse\"O\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x028\x126/api/v1/traffic/history_day/detection_point/statistics\x12\xe9\x01\n" +
	".ListHistoryDayTrafficDataByDetectionPointDaily\x12/.pb.ListTrafficDataByDetectionPointDailyRequest\x1a:.pb.ListHistoryDayTrafficDataByDetectionPointDailyResponse\"J\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x023\x121/api/v1/traffic/history_day/detection_point/daily\x12\xc4\x01\n" +
	"-DownloadHistoryDayTrafficDataByDetectionPoint\x12..pb.DownloadTrafficDataByDetectionPointRequest\x1a\x14.pb.DownloadResponse\"M\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x026\x124/api/v1/traffic/history_day/detection_point/downloadB~\x92AiZY\n" +
	"W\n" +
	"\x06bearer\x12M\b\x02\x128Authentication token, prefixed by Bearer: Bearer <token>\x1a\rAuthorization \x02b\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00Z\x10internal/data/pbb\x06proto3"

var (
	file_services_history_day_traffic_data_detection_point_service_proto_rawDescOnce sync.Once
	file_services_history_day_traffic_data_detection_point_service_proto_rawDescData []byte
)

func file_services_history_day_traffic_data_detection_point_service_proto_rawDescGZIP() []byte {
	file_services_history_day_traffic_data_detection_point_service_proto_rawDescOnce.Do(func() {
		file_services_history_day_traffic_data_detection_point_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_history_day_traffic_data_detection_point_service_proto_rawDesc), len(file_services_history_day_traffic_data_detection_point_service_proto_rawDesc)))
	})
	return file_services_history_day_traffic_data_detection_point_service_proto_rawDescData
}

var file_services_history_day_traffic_data_detection_point_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_services_history_day_traffic_data_detection_point_service_proto_goTypes = []any{
	(*ListHistoryDayTrafficDataByDetectionPointResponse)(nil),      // 0: pb.ListHistoryDayTrafficDataByDetectionPointResponse
	(*GetHistoryDayTrafficDataByDetectionPointResponse)(nil),       // 1: pb.GetHistoryDayTrafficDataByDetectionPointResponse
	(*CreateHistoryDayTrafficDataByDetectionPointRequest)(nil),     // 2: pb.CreateHistoryDayTrafficDataByDetectionPointRequest
	(*CreateHistoryDayTrafficDataByDetectionPointResponse)(nil),    // 3: pb.CreateHistoryDayTrafficDataByDetectionPointResponse
	(*UpdateHistoryDayTrafficDataByDetectionPointRequest)(nil),     // 4: pb.UpdateHistoryDayTrafficDataByDetectionPointRequest
	(*UpdateHistoryDayTrafficDataByDetectionPointResponse)(nil),    // 5: pb.UpdateHistoryDayTrafficDataByDetectionPointResponse
	(*ListHistoryDayTrafficDataByDetectionPointDailyResponse)(nil), // 6: pb.ListHistoryDayTrafficDataByDetectionPointDailyResponse
	(*HistoryDayTrafficDataByDetectionPoint)(nil),                  // 7: pb.HistoryDayTrafficDataByDetectionPoint
	(*ListTrafficDataByDetectionPointRequest)(nil),                 // 8: pb.ListTrafficDataByDetectionPointRequest
	(*GetTrafficDataByDetectionPointRequest)(nil),                  // 9: pb.GetTrafficDataByDetectionPointRequest
	(*DeleteTrafficDataByDetectionPointRequest)(nil),               // 10: pb.DeleteTrafficDataByDetectionPointRequest
	(*BulkCreateTrafficDataByDetectionPointRequest)(nil),           // 11: pb.BulkCreateTrafficDataByDetectionPointRequest
	(*GetTrafficDataByDetectionPointStatisticsRequest)(nil),        // 12: pb.GetTrafficDataByDetectionPointStatisticsRequest
	(*ListTrafficDataByDetectionPointDailyRequest)(nil),            // 13: pb.ListTrafficDataByDetectionPointDailyRequest
	(*DownloadTrafficDataByDetectionPointRequest)(nil),             // 14: pb.DownloadTrafficDataByDetectionPointRequest
	(*emptypb.Empty)(nil), // 15: google.protobuf.Empty
	(*BulkCreateTrafficDataByDetectionPointResponse)(nil),    // 16: pb.BulkCreateTrafficDataByDetectionPointResponse
	(*GetTrafficDataByDetectionPointStatisticsResponse)(nil), // 17: pb.GetTrafficDataByDetectionPointStatisticsResponse
	(*DownloadResponse)(nil),                                 // 18: pb.DownloadResponse
}
var file_services_history_day_traffic_data_detection_point_service_proto_depIdxs = []int32{
	7,  // 0: pb.ListHistoryDayTrafficDataByDetectionPointResponse.history_day_traffic_data_by_detection_points:type_name -> pb.HistoryDayTrafficDataByDetectionPoint
	7,  // 1: pb.GetHistoryDayTrafficDataByDetectionPointResponse.history_day_traffic_data_by_detection_point:type_name -> pb.HistoryDayTrafficDataByDetectionPoint
	7,  // 2: pb.CreateHistoryDayTrafficDataByDetectionPointRequest.history_day_traffic_data_by_detection_point:type_name -> pb.HistoryDayTrafficDataByDetectionPoint
	7,  // 3: pb.CreateHistoryDayTrafficDataByDetectionPointResponse.history_day_traffic_data_by_detection_point:type_name -> pb.HistoryDayTrafficDataByDetectionPoint
	7,  // 4: pb.UpdateHistoryDayTrafficDataByDetectionPointRequest.history_day_traffic_data_by_detection_point:type_name -> pb.HistoryDayTrafficDataByDetectionPoint
	7,  // 5: pb.UpdateHistoryDayTrafficDataByDetectionPointResponse.history_day_traffic_data_by_detection_point:type_name -> pb.HistoryDayTrafficDataByDetectionPoint
	7,  // 6: pb.ListHistoryDayTrafficDataByDetectionPointDailyResponse.daily_stats:type_name -> pb.HistoryDayTrafficDataByDetectionPoint
	8,  // 7: pb.HistoryDayTrafficDataByDetectionPointService.ListHistoryDayTrafficDataByDetectionPoint:input_type -> pb.ListTrafficDataByDetectionPointRequest
	9,  // 8: pb.HistoryDayTrafficDataByDetectionPointService.GetHistoryDayTrafficDataByDetectionPoint:input_type -> pb.GetTrafficDataByDetectionPointRequest
	2,  // 9: pb.HistoryDayTrafficDataByDetectionPointService.CreateHistoryDayTrafficDataByDetectionPoint:input_type -> pb.CreateHistoryDayTrafficDataByDetectionPointRequest
	4,  // 10: pb.HistoryDayTrafficDataByDetectionPointService.UpdateHistoryDayTrafficDataByDetectionPoint:input_type -> pb.UpdateHistoryDayTrafficDataByDetectionPointRequest
	10, // 11: pb.HistoryDayTrafficDataByDetectionPointService.DeleteHistoryDayTrafficDataByDetectionPoint:input_type -> pb.DeleteTrafficDataByDetectionPointRequest
	11, // 12: pb.HistoryDayTrafficDataByDetectionPointService.BulkCreateHistoryDayTrafficDataByDetectionPoint:input_type -> pb.BulkCreateTrafficDataByDetectionPointRequest
	12, // 13: pb.HistoryDayTrafficDataByDetectionPointService.GetHistoryDayTrafficDataByDetectionPointStatistics:input_type -> pb.GetTrafficDataByDetectionPointStatisticsRequest
	13, // 14: pb.HistoryDayTrafficDataByDetectionPointService.ListHistoryDayTrafficDataByDetectionPointDaily:input_type -> pb.ListTrafficDataByDetectionPointDailyRequest
	14, // 15: pb.HistoryDayTrafficDataByDetectionPointService.DownloadHistoryDayTrafficDataByDetectionPoint:input_type -> pb.DownloadTrafficDataByDetectionPointRequest
	0,  // 16: pb.HistoryDayTrafficDataByDetectionPointService.ListHistoryDayTrafficDataByDetectionPoint:output_type -> pb.ListHistoryDayTrafficDataByDetectionPointResponse
	1,  // 17: pb.HistoryDayTrafficDataByDetectionPointService.GetHistoryDayTrafficDataByDetectionPoint:output_type -> pb.GetHistoryDayTrafficDataByDetectionPointResponse
	3,  // 18: pb.HistoryDayTrafficDataByDetectionPointService.CreateHistoryDayTrafficDataByDetectionPoint:output_type -> pb.CreateHistoryDayTrafficDataByDetectionPointResponse
	5,  // 19: pb.HistoryDayTrafficDataByDetectionPointService.UpdateHistoryDayTrafficDataByDetectionPoint:output_type -> pb.UpdateHistoryDayTrafficDataByDetectionPointResponse
	15, // 20: pb.HistoryDayTrafficDataByDetectionPointService.DeleteHistoryDayTrafficDataByDetectionPoint:output_type -> google.protobuf.Empty
	16, // 21: pb.HistoryDayTrafficDataByDetectionPointService.BulkCreateHistoryDayTrafficDataByDetectionPoint:output_type -> pb.BulkCreateTrafficDataByDetectionPointResponse
	17, // 22: pb.HistoryDayTrafficDataByDetectionPointService.GetHistoryDayTrafficDataByDetectionPointStatistics:output_type -> pb.GetTrafficDataByDetectionPointStatisticsResponse
	6,  // 23: pb.HistoryDayTrafficDataByDetectionPointService.ListHistoryDayTrafficDataByDetectionPointDaily:output_type -> pb.ListHistoryDayTrafficDataByDetectionPointDailyResponse
	18, // 24: pb.HistoryDayTrafficDataByDetectionPointService.DownloadHistoryDayTrafficDataByDetectionPoint:output_type -> pb.DownloadResponse
	16, // [16:25] is the sub-list for method output_type
	7,  // [7:16] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_services_history_day_traffic_data_detection_point_service_proto_init() }
func file_services_history_day_traffic_data_detection_point_service_proto_init() {
	if File_services_history_day_traffic_data_detection_point_service_proto != nil {
		return
	}
	file_dtos_history_day_traffic_data_detection_point_proto_init()
	file_dtos_traffic_requests_proto_init()
	file_dtos_traffic_responses_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_history_day_traffic_data_detection_point_service_proto_rawDesc), len(file_services_history_day_traffic_data_detection_point_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_history_day_traffic_data_detection_point_service_proto_goTypes,
		DependencyIndexes: file_services_history_day_traffic_data_detection_point_service_proto_depIdxs,
		MessageInfos:      file_services_history_day_traffic_data_detection_point_service_proto_msgTypes,
	}.Build()
	File_services_history_day_traffic_data_detection_point_service_proto = out.File
	file_services_history_day_traffic_data_detection_point_service_proto_goTypes = nil
	file_services_history_day_traffic_data_detection_point_service_proto_depIdxs = nil
}
