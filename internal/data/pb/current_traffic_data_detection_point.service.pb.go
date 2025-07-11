// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: services/current_traffic_data_detection_point.service.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type ListCurrentTrafficDataByDetectionPointsResponse struct {
	state                               protoimpl.MessageState                `protogen:"open.v1"`
	CurrentTrafficDataByDetectionPoints []*CurrentTrafficDataByDetectionPoint `protobuf:"bytes,1,rep,name=current_traffic_data_by_detection_points,json=currentTrafficDataByDetectionPoints,proto3" json:"current_traffic_data_by_detection_points,omitempty"`
	TotalCount                          uint32                                `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields                       protoimpl.UnknownFields
	sizeCache                           protoimpl.SizeCache
}

func (x *ListCurrentTrafficDataByDetectionPointsResponse) Reset() {
	*x = ListCurrentTrafficDataByDetectionPointsResponse{}
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCurrentTrafficDataByDetectionPointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCurrentTrafficDataByDetectionPointsResponse) ProtoMessage() {}

func (x *ListCurrentTrafficDataByDetectionPointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCurrentTrafficDataByDetectionPointsResponse.ProtoReflect.Descriptor instead.
func (*ListCurrentTrafficDataByDetectionPointsResponse) Descriptor() ([]byte, []int) {
	return file_services_current_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListCurrentTrafficDataByDetectionPointsResponse) GetCurrentTrafficDataByDetectionPoints() []*CurrentTrafficDataByDetectionPoint {
	if x != nil {
		return x.CurrentTrafficDataByDetectionPoints
	}
	return nil
}

func (x *ListCurrentTrafficDataByDetectionPointsResponse) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type GetCurrentTrafficDataByDetectionPointsResponse struct {
	state                              protoimpl.MessageState                `protogen:"open.v1"`
	CurrentTrafficDataByDetectionPoint []*CurrentTrafficDataByDetectionPoint `protobuf:"bytes,1,rep,name=current_traffic_data_by_detection_point,json=currentTrafficDataByDetectionPoint,proto3" json:"current_traffic_data_by_detection_point,omitempty"`
	TotalCount                         uint32                                `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields                      protoimpl.UnknownFields
	sizeCache                          protoimpl.SizeCache
}

func (x *GetCurrentTrafficDataByDetectionPointsResponse) Reset() {
	*x = GetCurrentTrafficDataByDetectionPointsResponse{}
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCurrentTrafficDataByDetectionPointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentTrafficDataByDetectionPointsResponse) ProtoMessage() {}

func (x *GetCurrentTrafficDataByDetectionPointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentTrafficDataByDetectionPointsResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentTrafficDataByDetectionPointsResponse) Descriptor() ([]byte, []int) {
	return file_services_current_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCurrentTrafficDataByDetectionPointsResponse) GetCurrentTrafficDataByDetectionPoint() []*CurrentTrafficDataByDetectionPoint {
	if x != nil {
		return x.CurrentTrafficDataByDetectionPoint
	}
	return nil
}

func (x *GetCurrentTrafficDataByDetectionPointsResponse) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type CreateCurrentTrafficDataByDetectionPointRequest struct {
	state                              protoimpl.MessageState              `protogen:"open.v1"`
	CurrentTrafficDataByDetectionPoint *CurrentTrafficDataByDetectionPoint `protobuf:"bytes,1,opt,name=current_traffic_data_by_detection_point,json=currentTrafficDataByDetectionPoint,proto3" json:"current_traffic_data_by_detection_point,omitempty"`
	unknownFields                      protoimpl.UnknownFields
	sizeCache                          protoimpl.SizeCache
}

func (x *CreateCurrentTrafficDataByDetectionPointRequest) Reset() {
	*x = CreateCurrentTrafficDataByDetectionPointRequest{}
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCurrentTrafficDataByDetectionPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCurrentTrafficDataByDetectionPointRequest) ProtoMessage() {}

func (x *CreateCurrentTrafficDataByDetectionPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCurrentTrafficDataByDetectionPointRequest.ProtoReflect.Descriptor instead.
func (*CreateCurrentTrafficDataByDetectionPointRequest) Descriptor() ([]byte, []int) {
	return file_services_current_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCurrentTrafficDataByDetectionPointRequest) GetCurrentTrafficDataByDetectionPoint() *CurrentTrafficDataByDetectionPoint {
	if x != nil {
		return x.CurrentTrafficDataByDetectionPoint
	}
	return nil
}

type CreateCurrentTrafficDataByDetectionPointResponse struct {
	state                              protoimpl.MessageState              `protogen:"open.v1"`
	CurrentTrafficDataByDetectionPoint *CurrentTrafficDataByDetectionPoint `protobuf:"bytes,1,opt,name=current_traffic_data_by_detection_point,json=currentTrafficDataByDetectionPoint,proto3" json:"current_traffic_data_by_detection_point,omitempty"`
	unknownFields                      protoimpl.UnknownFields
	sizeCache                          protoimpl.SizeCache
}

func (x *CreateCurrentTrafficDataByDetectionPointResponse) Reset() {
	*x = CreateCurrentTrafficDataByDetectionPointResponse{}
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCurrentTrafficDataByDetectionPointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCurrentTrafficDataByDetectionPointResponse) ProtoMessage() {}

func (x *CreateCurrentTrafficDataByDetectionPointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCurrentTrafficDataByDetectionPointResponse.ProtoReflect.Descriptor instead.
func (*CreateCurrentTrafficDataByDetectionPointResponse) Descriptor() ([]byte, []int) {
	return file_services_current_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCurrentTrafficDataByDetectionPointResponse) GetCurrentTrafficDataByDetectionPoint() *CurrentTrafficDataByDetectionPoint {
	if x != nil {
		return x.CurrentTrafficDataByDetectionPoint
	}
	return nil
}

type UpdateCurrentTrafficDataByDetectionPointRequest struct {
	state                              protoimpl.MessageState              `protogen:"open.v1"`
	CurrentTrafficDataByDetectionPoint *CurrentTrafficDataByDetectionPoint `protobuf:"bytes,1,opt,name=current_traffic_data_by_detection_point,json=currentTrafficDataByDetectionPoint,proto3" json:"current_traffic_data_by_detection_point,omitempty"`
	unknownFields                      protoimpl.UnknownFields
	sizeCache                          protoimpl.SizeCache
}

func (x *UpdateCurrentTrafficDataByDetectionPointRequest) Reset() {
	*x = UpdateCurrentTrafficDataByDetectionPointRequest{}
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCurrentTrafficDataByDetectionPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrentTrafficDataByDetectionPointRequest) ProtoMessage() {}

func (x *UpdateCurrentTrafficDataByDetectionPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrentTrafficDataByDetectionPointRequest.ProtoReflect.Descriptor instead.
func (*UpdateCurrentTrafficDataByDetectionPointRequest) Descriptor() ([]byte, []int) {
	return file_services_current_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCurrentTrafficDataByDetectionPointRequest) GetCurrentTrafficDataByDetectionPoint() *CurrentTrafficDataByDetectionPoint {
	if x != nil {
		return x.CurrentTrafficDataByDetectionPoint
	}
	return nil
}

type UpdateCurrentTrafficDataByDetectionPointResponse struct {
	state                              protoimpl.MessageState              `protogen:"open.v1"`
	CurrentTrafficDataByDetectionPoint *CurrentTrafficDataByDetectionPoint `protobuf:"bytes,1,opt,name=current_traffic_data_by_detection_point,json=currentTrafficDataByDetectionPoint,proto3" json:"current_traffic_data_by_detection_point,omitempty"`
	unknownFields                      protoimpl.UnknownFields
	sizeCache                          protoimpl.SizeCache
}

func (x *UpdateCurrentTrafficDataByDetectionPointResponse) Reset() {
	*x = UpdateCurrentTrafficDataByDetectionPointResponse{}
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCurrentTrafficDataByDetectionPointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrentTrafficDataByDetectionPointResponse) ProtoMessage() {}

func (x *UpdateCurrentTrafficDataByDetectionPointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrentTrafficDataByDetectionPointResponse.ProtoReflect.Descriptor instead.
func (*UpdateCurrentTrafficDataByDetectionPointResponse) Descriptor() ([]byte, []int) {
	return file_services_current_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCurrentTrafficDataByDetectionPointResponse) GetCurrentTrafficDataByDetectionPoint() *CurrentTrafficDataByDetectionPoint {
	if x != nil {
		return x.CurrentTrafficDataByDetectionPoint
	}
	return nil
}

type GetCurrentTrafficDataByDetectionPointStatisticsResponse struct {
	state                      protoimpl.MessageState                          `protogen:"open.v1"`
	DataSourceName             string                                          `protobuf:"bytes,1,opt,name=data_source_name,json=dataSourceName,proto3" json:"data_source_name,omitempty"`
	RecordsCount               uint32                                          `protobuf:"varint,2,opt,name=records_count,json=recordsCount,proto3" json:"records_count,omitempty"`
	FirstRecordTimestamp       *timestamppb.Timestamp                          `protobuf:"bytes,3,opt,name=first_record_timestamp,json=firstRecordTimestamp,proto3" json:"first_record_timestamp,omitempty"`
	LastRecordTimestamp        *timestamppb.Timestamp                          `protobuf:"bytes,4,opt,name=last_record_timestamp,json=lastRecordTimestamp,proto3" json:"last_record_timestamp,omitempty"`
	StatisticsByDetectionPoint []*CurrentTrafficDataByDetectionPointStatistics `protobuf:"bytes,5,rep,name=statistics_by_detection_point,json=statisticsByDetectionPoint,proto3" json:"statistics_by_detection_point,omitempty"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *GetCurrentTrafficDataByDetectionPointStatisticsResponse) Reset() {
	*x = GetCurrentTrafficDataByDetectionPointStatisticsResponse{}
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCurrentTrafficDataByDetectionPointStatisticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentTrafficDataByDetectionPointStatisticsResponse) ProtoMessage() {}

func (x *GetCurrentTrafficDataByDetectionPointStatisticsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentTrafficDataByDetectionPointStatisticsResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentTrafficDataByDetectionPointStatisticsResponse) Descriptor() ([]byte, []int) {
	return file_services_current_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetCurrentTrafficDataByDetectionPointStatisticsResponse) GetDataSourceName() string {
	if x != nil {
		return x.DataSourceName
	}
	return ""
}

func (x *GetCurrentTrafficDataByDetectionPointStatisticsResponse) GetRecordsCount() uint32 {
	if x != nil {
		return x.RecordsCount
	}
	return 0
}

func (x *GetCurrentTrafficDataByDetectionPointStatisticsResponse) GetFirstRecordTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.FirstRecordTimestamp
	}
	return nil
}

func (x *GetCurrentTrafficDataByDetectionPointStatisticsResponse) GetLastRecordTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.LastRecordTimestamp
	}
	return nil
}

func (x *GetCurrentTrafficDataByDetectionPointStatisticsResponse) GetStatisticsByDetectionPoint() []*CurrentTrafficDataByDetectionPointStatistics {
	if x != nil {
		return x.StatisticsByDetectionPoint
	}
	return nil
}

type ListCurrentTrafficDataByDetectionPointDailyResponse struct {
	state         protoimpl.MessageState                `protogen:"open.v1"`
	DailyStats    []*CurrentTrafficDataByDetectionPoint `protobuf:"bytes,1,rep,name=daily_stats,json=dailyStats,proto3" json:"daily_stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCurrentTrafficDataByDetectionPointDailyResponse) Reset() {
	*x = ListCurrentTrafficDataByDetectionPointDailyResponse{}
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCurrentTrafficDataByDetectionPointDailyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCurrentTrafficDataByDetectionPointDailyResponse) ProtoMessage() {}

func (x *ListCurrentTrafficDataByDetectionPointDailyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_current_traffic_data_detection_point_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCurrentTrafficDataByDetectionPointDailyResponse.ProtoReflect.Descriptor instead.
func (*ListCurrentTrafficDataByDetectionPointDailyResponse) Descriptor() ([]byte, []int) {
	return file_services_current_traffic_data_detection_point_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListCurrentTrafficDataByDetectionPointDailyResponse) GetDailyStats() []*CurrentTrafficDataByDetectionPoint {
	if x != nil {
		return x.DailyStats
	}
	return nil
}

var File_services_current_traffic_data_detection_point_service_proto protoreflect.FileDescriptor

const file_services_current_traffic_data_detection_point_service_proto_rawDesc = "" +
	"\n" +
	";services/current_traffic_data_detection_point.service.proto\x12\x02pb\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a/dtos/current_traffic_data_detection_point.proto\x1a\x1bdtos/traffic_requests.proto\x1a\x1ddtos/traffic_statistics.proto\x1a\x1cdtos/traffic_responses.proto\"\xd1\x01\n" +
	"/ListCurrentTrafficDataByDetectionPointsResponse\x12}\n" +
	"(current_traffic_data_by_detection_points\x18\x01 \x03(\v2&.pb.CurrentTrafficDataByDetectionPointR#currentTrafficDataByDetectionPoints\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\rR\n" +
	"totalCount\"\xce\x01\n" +
	".GetCurrentTrafficDataByDetectionPointsResponse\x12{\n" +
	"'current_traffic_data_by_detection_point\x18\x01 \x03(\v2&.pb.CurrentTrafficDataByDetectionPointR\"currentTrafficDataByDetectionPoint\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\rR\n" +
	"totalCount\"\xae\x01\n" +
	"/CreateCurrentTrafficDataByDetectionPointRequest\x12{\n" +
	"'current_traffic_data_by_detection_point\x18\x01 \x01(\v2&.pb.CurrentTrafficDataByDetectionPointR\"currentTrafficDataByDetectionPoint\"\xaf\x01\n" +
	"0CreateCurrentTrafficDataByDetectionPointResponse\x12{\n" +
	"'current_traffic_data_by_detection_point\x18\x01 \x01(\v2&.pb.CurrentTrafficDataByDetectionPointR\"currentTrafficDataByDetectionPoint\"\xae\x01\n" +
	"/UpdateCurrentTrafficDataByDetectionPointRequest\x12{\n" +
	"'current_traffic_data_by_detection_point\x18\x01 \x01(\v2&.pb.CurrentTrafficDataByDetectionPointR\"currentTrafficDataByDetectionPoint\"\xaf\x01\n" +
	"0UpdateCurrentTrafficDataByDetectionPointResponse\x12{\n" +
	"'current_traffic_data_by_detection_point\x18\x01 \x01(\v2&.pb.CurrentTrafficDataByDetectionPointR\"currentTrafficDataByDetectionPoint\"\x9f\x03\n" +
	"7GetCurrentTrafficDataByDetectionPointStatisticsResponse\x12(\n" +
	"\x10data_source_name\x18\x01 \x01(\tR\x0edataSourceName\x12#\n" +
	"\rrecords_count\x18\x02 \x01(\rR\frecordsCount\x12P\n" +
	"\x16first_record_timestamp\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\x14firstRecordTimestamp\x12N\n" +
	"\x15last_record_timestamp\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\x13lastRecordTimestamp\x12s\n" +
	"\x1dstatistics_by_detection_point\x18\x05 \x03(\v20.pb.CurrentTrafficDataByDetectionPointStatisticsR\x1astatisticsByDetectionPoint\"~\n" +
	"3ListCurrentTrafficDataByDetectionPointDailyResponse\x12G\n" +
	"\vdaily_stats\x18\x01 \x03(\v2&.pb.CurrentTrafficDataByDetectionPointR\n" +
	"dailyStats2\xb5\x0f\n" +
	")CurrentTrafficDataByDetectionPointService\x12\xd1\x01\n" +
	"'ListCurrentTrafficDataByDetectionPoints\x12*.pb.ListTrafficDataByDetectionPointRequest\x1a3.pb.ListCurrentTrafficDataByDetectionPointsResponse\"E\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x02.\x12,/api/v1/traffic/current/detection_point/list\x12\xcf\x01\n" +
	"%GetCurrentTrafficDataByDetectionPoint\x12).pb.GetTrafficDataByDetectionPointRequest\x1a2.pb.GetCurrentTrafficDataByDetectionPointsResponse\"G\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x020\x12./api/v1/traffic/current/detection_point/single\x12\xda\x01\n" +
	"(CreateCurrentTrafficDataByDetectionPoint\x123.pb.CreateCurrentTrafficDataByDetectionPointRequest\x1a4.pb.CreateCurrentTrafficDataByDetectionPointResponse\"C\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x02,:\x01*\"'/api/v1/traffic/current/detection_point\x12\xda\x01\n" +
	"(UpdateCurrentTrafficDataByDetectionPoint\x123.pb.UpdateCurrentTrafficDataByDetectionPointRequest\x1a4.pb.UpdateCurrentTrafficDataByDetectionPointResponse\"C\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x02,:\x01*2'/api/v1/traffic/current/detection_point\x12\xb2\x01\n" +
	"(DeleteCurrentTrafficDataByDetectionPoint\x12,.pb.DeleteTrafficDataByDetectionPointRequest\x1a\x16.google.protobuf.Empty\"@\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x02)*'/api/v1/traffic/current/detection_point\x12\xdd\x01\n" +
	",BulkCreateCurrentTrafficDataByDetectionPoint\x120.pb.BulkCreateTrafficDataByDetectionPointRequest\x1a1.pb.BulkCreateTrafficDataByDetectionPointResponse\"H\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x021:\x01*\",/api/v1/traffic/current/detection_point/bulk\x12\xf0\x01\n" +
	"/GetCurrentTrafficDataByDetectionPointStatistics\x123.pb.GetTrafficDataByDetectionPointStatisticsRequest\x1a;.pb.GetCurrentTrafficDataByDetectionPointStatisticsResponse\"K\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x024\x122/api/v1/traffic/current/detection_point/statistics\x12\xdf\x01\n" +
	"+ListCurrentTrafficDataByDetectionPointDaily\x12/.pb.ListTrafficDataByDetectionPointDailyRequest\x1a7.pb.ListCurrentTrafficDataByDetectionPointDailyResponse\"F\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x02/\x12-/api/v1/traffic/current/detection_point/daily\x12\xbd\x01\n" +
	"*DownloadCurrentTrafficDataByDetectionPoint\x12..pb.DownloadTrafficDataByDetectionPointRequest\x1a\x14.pb.DownloadResponse\"I\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x022\x120/api/v1/traffic/current/detection_point/downloadB~\x92AiZY\n" +
	"W\n" +
	"\x06bearer\x12M\b\x02\x128Authentication token, prefixed by Bearer: Bearer <token>\x1a\rAuthorization \x02b\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00Z\x10internal/data/pbb\x06proto3"

var (
	file_services_current_traffic_data_detection_point_service_proto_rawDescOnce sync.Once
	file_services_current_traffic_data_detection_point_service_proto_rawDescData []byte
)

func file_services_current_traffic_data_detection_point_service_proto_rawDescGZIP() []byte {
	file_services_current_traffic_data_detection_point_service_proto_rawDescOnce.Do(func() {
		file_services_current_traffic_data_detection_point_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_current_traffic_data_detection_point_service_proto_rawDesc), len(file_services_current_traffic_data_detection_point_service_proto_rawDesc)))
	})
	return file_services_current_traffic_data_detection_point_service_proto_rawDescData
}

var file_services_current_traffic_data_detection_point_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_services_current_traffic_data_detection_point_service_proto_goTypes = []any{
	(*ListCurrentTrafficDataByDetectionPointsResponse)(nil),         // 0: pb.ListCurrentTrafficDataByDetectionPointsResponse
	(*GetCurrentTrafficDataByDetectionPointsResponse)(nil),          // 1: pb.GetCurrentTrafficDataByDetectionPointsResponse
	(*CreateCurrentTrafficDataByDetectionPointRequest)(nil),         // 2: pb.CreateCurrentTrafficDataByDetectionPointRequest
	(*CreateCurrentTrafficDataByDetectionPointResponse)(nil),        // 3: pb.CreateCurrentTrafficDataByDetectionPointResponse
	(*UpdateCurrentTrafficDataByDetectionPointRequest)(nil),         // 4: pb.UpdateCurrentTrafficDataByDetectionPointRequest
	(*UpdateCurrentTrafficDataByDetectionPointResponse)(nil),        // 5: pb.UpdateCurrentTrafficDataByDetectionPointResponse
	(*GetCurrentTrafficDataByDetectionPointStatisticsResponse)(nil), // 6: pb.GetCurrentTrafficDataByDetectionPointStatisticsResponse
	(*ListCurrentTrafficDataByDetectionPointDailyResponse)(nil),     // 7: pb.ListCurrentTrafficDataByDetectionPointDailyResponse
	(*CurrentTrafficDataByDetectionPoint)(nil),                      // 8: pb.CurrentTrafficDataByDetectionPoint
	(*timestamppb.Timestamp)(nil),                                   // 9: google.protobuf.Timestamp
	(*CurrentTrafficDataByDetectionPointStatistics)(nil),            // 10: pb.CurrentTrafficDataByDetectionPointStatistics
	(*ListTrafficDataByDetectionPointRequest)(nil),                  // 11: pb.ListTrafficDataByDetectionPointRequest
	(*GetTrafficDataByDetectionPointRequest)(nil),                   // 12: pb.GetTrafficDataByDetectionPointRequest
	(*DeleteTrafficDataByDetectionPointRequest)(nil),                // 13: pb.DeleteTrafficDataByDetectionPointRequest
	(*BulkCreateTrafficDataByDetectionPointRequest)(nil),            // 14: pb.BulkCreateTrafficDataByDetectionPointRequest
	(*GetTrafficDataByDetectionPointStatisticsRequest)(nil),         // 15: pb.GetTrafficDataByDetectionPointStatisticsRequest
	(*ListTrafficDataByDetectionPointDailyRequest)(nil),             // 16: pb.ListTrafficDataByDetectionPointDailyRequest
	(*DownloadTrafficDataByDetectionPointRequest)(nil),              // 17: pb.DownloadTrafficDataByDetectionPointRequest
	(*emptypb.Empty)(nil),                                           // 18: google.protobuf.Empty
	(*BulkCreateTrafficDataByDetectionPointResponse)(nil),           // 19: pb.BulkCreateTrafficDataByDetectionPointResponse
	(*DownloadResponse)(nil),                                        // 20: pb.DownloadResponse
}
var file_services_current_traffic_data_detection_point_service_proto_depIdxs = []int32{
	8,  // 0: pb.ListCurrentTrafficDataByDetectionPointsResponse.current_traffic_data_by_detection_points:type_name -> pb.CurrentTrafficDataByDetectionPoint
	8,  // 1: pb.GetCurrentTrafficDataByDetectionPointsResponse.current_traffic_data_by_detection_point:type_name -> pb.CurrentTrafficDataByDetectionPoint
	8,  // 2: pb.CreateCurrentTrafficDataByDetectionPointRequest.current_traffic_data_by_detection_point:type_name -> pb.CurrentTrafficDataByDetectionPoint
	8,  // 3: pb.CreateCurrentTrafficDataByDetectionPointResponse.current_traffic_data_by_detection_point:type_name -> pb.CurrentTrafficDataByDetectionPoint
	8,  // 4: pb.UpdateCurrentTrafficDataByDetectionPointRequest.current_traffic_data_by_detection_point:type_name -> pb.CurrentTrafficDataByDetectionPoint
	8,  // 5: pb.UpdateCurrentTrafficDataByDetectionPointResponse.current_traffic_data_by_detection_point:type_name -> pb.CurrentTrafficDataByDetectionPoint
	9,  // 6: pb.GetCurrentTrafficDataByDetectionPointStatisticsResponse.first_record_timestamp:type_name -> google.protobuf.Timestamp
	9,  // 7: pb.GetCurrentTrafficDataByDetectionPointStatisticsResponse.last_record_timestamp:type_name -> google.protobuf.Timestamp
	10, // 8: pb.GetCurrentTrafficDataByDetectionPointStatisticsResponse.statistics_by_detection_point:type_name -> pb.CurrentTrafficDataByDetectionPointStatistics
	8,  // 9: pb.ListCurrentTrafficDataByDetectionPointDailyResponse.daily_stats:type_name -> pb.CurrentTrafficDataByDetectionPoint
	11, // 10: pb.CurrentTrafficDataByDetectionPointService.ListCurrentTrafficDataByDetectionPoints:input_type -> pb.ListTrafficDataByDetectionPointRequest
	12, // 11: pb.CurrentTrafficDataByDetectionPointService.GetCurrentTrafficDataByDetectionPoint:input_type -> pb.GetTrafficDataByDetectionPointRequest
	2,  // 12: pb.CurrentTrafficDataByDetectionPointService.CreateCurrentTrafficDataByDetectionPoint:input_type -> pb.CreateCurrentTrafficDataByDetectionPointRequest
	4,  // 13: pb.CurrentTrafficDataByDetectionPointService.UpdateCurrentTrafficDataByDetectionPoint:input_type -> pb.UpdateCurrentTrafficDataByDetectionPointRequest
	13, // 14: pb.CurrentTrafficDataByDetectionPointService.DeleteCurrentTrafficDataByDetectionPoint:input_type -> pb.DeleteTrafficDataByDetectionPointRequest
	14, // 15: pb.CurrentTrafficDataByDetectionPointService.BulkCreateCurrentTrafficDataByDetectionPoint:input_type -> pb.BulkCreateTrafficDataByDetectionPointRequest
	15, // 16: pb.CurrentTrafficDataByDetectionPointService.GetCurrentTrafficDataByDetectionPointStatistics:input_type -> pb.GetTrafficDataByDetectionPointStatisticsRequest
	16, // 17: pb.CurrentTrafficDataByDetectionPointService.ListCurrentTrafficDataByDetectionPointDaily:input_type -> pb.ListTrafficDataByDetectionPointDailyRequest
	17, // 18: pb.CurrentTrafficDataByDetectionPointService.DownloadCurrentTrafficDataByDetectionPoint:input_type -> pb.DownloadTrafficDataByDetectionPointRequest
	0,  // 19: pb.CurrentTrafficDataByDetectionPointService.ListCurrentTrafficDataByDetectionPoints:output_type -> pb.ListCurrentTrafficDataByDetectionPointsResponse
	1,  // 20: pb.CurrentTrafficDataByDetectionPointService.GetCurrentTrafficDataByDetectionPoint:output_type -> pb.GetCurrentTrafficDataByDetectionPointsResponse
	3,  // 21: pb.CurrentTrafficDataByDetectionPointService.CreateCurrentTrafficDataByDetectionPoint:output_type -> pb.CreateCurrentTrafficDataByDetectionPointResponse
	5,  // 22: pb.CurrentTrafficDataByDetectionPointService.UpdateCurrentTrafficDataByDetectionPoint:output_type -> pb.UpdateCurrentTrafficDataByDetectionPointResponse
	18, // 23: pb.CurrentTrafficDataByDetectionPointService.DeleteCurrentTrafficDataByDetectionPoint:output_type -> google.protobuf.Empty
	19, // 24: pb.CurrentTrafficDataByDetectionPointService.BulkCreateCurrentTrafficDataByDetectionPoint:output_type -> pb.BulkCreateTrafficDataByDetectionPointResponse
	6,  // 25: pb.CurrentTrafficDataByDetectionPointService.GetCurrentTrafficDataByDetectionPointStatistics:output_type -> pb.GetCurrentTrafficDataByDetectionPointStatisticsResponse
	7,  // 26: pb.CurrentTrafficDataByDetectionPointService.ListCurrentTrafficDataByDetectionPointDaily:output_type -> pb.ListCurrentTrafficDataByDetectionPointDailyResponse
	20, // 27: pb.CurrentTrafficDataByDetectionPointService.DownloadCurrentTrafficDataByDetectionPoint:output_type -> pb.DownloadResponse
	19, // [19:28] is the sub-list for method output_type
	10, // [10:19] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_services_current_traffic_data_detection_point_service_proto_init() }
func file_services_current_traffic_data_detection_point_service_proto_init() {
	if File_services_current_traffic_data_detection_point_service_proto != nil {
		return
	}
	file_dtos_current_traffic_data_detection_point_proto_init()
	file_dtos_traffic_requests_proto_init()
	file_dtos_traffic_statistics_proto_init()
	file_dtos_traffic_responses_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_current_traffic_data_detection_point_service_proto_rawDesc), len(file_services_current_traffic_data_detection_point_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_current_traffic_data_detection_point_service_proto_goTypes,
		DependencyIndexes: file_services_current_traffic_data_detection_point_service_proto_depIdxs,
		MessageInfos:      file_services_current_traffic_data_detection_point_service_proto_msgTypes,
	}.Build()
	File_services_current_traffic_data_detection_point_service_proto = out.File
	file_services_current_traffic_data_detection_point_service_proto_goTypes = nil
	file_services_current_traffic_data_detection_point_service_proto_depIdxs = nil
}
