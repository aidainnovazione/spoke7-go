// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: services/detection_section.service.proto

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

type CreateDetectionSectionRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	DatasourceName   string                 `protobuf:"bytes,1,opt,name=datasource_name,json=datasourceName,proto3" json:"datasource_name,omitempty"`
	DetectionSection *DetectionSection      `protobuf:"bytes,2,opt,name=detection_section,json=detectionSection,proto3" json:"detection_section,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateDetectionSectionRequest) Reset() {
	*x = CreateDetectionSectionRequest{}
	mi := &file_services_detection_section_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDetectionSectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDetectionSectionRequest) ProtoMessage() {}

func (x *CreateDetectionSectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_detection_section_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDetectionSectionRequest.ProtoReflect.Descriptor instead.
func (*CreateDetectionSectionRequest) Descriptor() ([]byte, []int) {
	return file_services_detection_section_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDetectionSectionRequest) GetDatasourceName() string {
	if x != nil {
		return x.DatasourceName
	}
	return ""
}

func (x *CreateDetectionSectionRequest) GetDetectionSection() *DetectionSection {
	if x != nil {
		return x.DetectionSection
	}
	return nil
}

type UpdateDetectionSectionRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	DatasourceName   string                 `protobuf:"bytes,1,opt,name=datasource_name,json=datasourceName,proto3" json:"datasource_name,omitempty"`
	DetectionSection *DetectionSection      `protobuf:"bytes,2,opt,name=detection_section,json=detectionSection,proto3" json:"detection_section,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *UpdateDetectionSectionRequest) Reset() {
	*x = UpdateDetectionSectionRequest{}
	mi := &file_services_detection_section_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDetectionSectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDetectionSectionRequest) ProtoMessage() {}

func (x *UpdateDetectionSectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_detection_section_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDetectionSectionRequest.ProtoReflect.Descriptor instead.
func (*UpdateDetectionSectionRequest) Descriptor() ([]byte, []int) {
	return file_services_detection_section_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateDetectionSectionRequest) GetDatasourceName() string {
	if x != nil {
		return x.DatasourceName
	}
	return ""
}

func (x *UpdateDetectionSectionRequest) GetDetectionSection() *DetectionSection {
	if x != nil {
		return x.DetectionSection
	}
	return nil
}

type DeleteDetectionSectionRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	DatasourceName string                 `protobuf:"bytes,1,opt,name=datasource_name,json=datasourceName,proto3" json:"datasource_name,omitempty"`
	Id             string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DeleteDetectionSectionRequest) Reset() {
	*x = DeleteDetectionSectionRequest{}
	mi := &file_services_detection_section_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDetectionSectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDetectionSectionRequest) ProtoMessage() {}

func (x *DeleteDetectionSectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_detection_section_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDetectionSectionRequest.ProtoReflect.Descriptor instead.
func (*DeleteDetectionSectionRequest) Descriptor() ([]byte, []int) {
	return file_services_detection_section_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteDetectionSectionRequest) GetDatasourceName() string {
	if x != nil {
		return x.DatasourceName
	}
	return ""
}

func (x *DeleteDetectionSectionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetDetectionSectionRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	DatasourceName string                 `protobuf:"bytes,1,opt,name=datasource_name,json=datasourceName,proto3" json:"datasource_name,omitempty"`
	Id             string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetDetectionSectionRequest) Reset() {
	*x = GetDetectionSectionRequest{}
	mi := &file_services_detection_section_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDetectionSectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetectionSectionRequest) ProtoMessage() {}

func (x *GetDetectionSectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_detection_section_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetectionSectionRequest.ProtoReflect.Descriptor instead.
func (*GetDetectionSectionRequest) Descriptor() ([]byte, []int) {
	return file_services_detection_section_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetDetectionSectionRequest) GetDatasourceName() string {
	if x != nil {
		return x.DatasourceName
	}
	return ""
}

func (x *GetDetectionSectionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListDetectionSectionsRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	DatasourceName string                 `protobuf:"bytes,1,opt,name=datasource_name,json=datasourceName,proto3" json:"datasource_name,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ListDetectionSectionsRequest) Reset() {
	*x = ListDetectionSectionsRequest{}
	mi := &file_services_detection_section_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDetectionSectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDetectionSectionsRequest) ProtoMessage() {}

func (x *ListDetectionSectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_detection_section_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDetectionSectionsRequest.ProtoReflect.Descriptor instead.
func (*ListDetectionSectionsRequest) Descriptor() ([]byte, []int) {
	return file_services_detection_section_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListDetectionSectionsRequest) GetDatasourceName() string {
	if x != nil {
		return x.DatasourceName
	}
	return ""
}

type ListDetectionSectionsResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	DatasourceName    string                 `protobuf:"bytes,1,opt,name=datasource_name,json=datasourceName,proto3" json:"datasource_name,omitempty"`
	DetectionSections []*DetectionSection    `protobuf:"bytes,2,rep,name=detection_sections,json=detectionSections,proto3" json:"detection_sections,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ListDetectionSectionsResponse) Reset() {
	*x = ListDetectionSectionsResponse{}
	mi := &file_services_detection_section_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDetectionSectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDetectionSectionsResponse) ProtoMessage() {}

func (x *ListDetectionSectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_detection_section_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDetectionSectionsResponse.ProtoReflect.Descriptor instead.
func (*ListDetectionSectionsResponse) Descriptor() ([]byte, []int) {
	return file_services_detection_section_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListDetectionSectionsResponse) GetDatasourceName() string {
	if x != nil {
		return x.DatasourceName
	}
	return ""
}

func (x *ListDetectionSectionsResponse) GetDetectionSections() []*DetectionSection {
	if x != nil {
		return x.DetectionSections
	}
	return nil
}

type BulkDetectionSectionRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	DatasourceName string                 `protobuf:"bytes,1,opt,name=datasource_name,json=datasourceName,proto3" json:"datasource_name,omitempty"`
	File           *File                  `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BulkDetectionSectionRequest) Reset() {
	*x = BulkDetectionSectionRequest{}
	mi := &file_services_detection_section_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BulkDetectionSectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkDetectionSectionRequest) ProtoMessage() {}

func (x *BulkDetectionSectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_detection_section_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkDetectionSectionRequest.ProtoReflect.Descriptor instead.
func (*BulkDetectionSectionRequest) Descriptor() ([]byte, []int) {
	return file_services_detection_section_service_proto_rawDescGZIP(), []int{6}
}

func (x *BulkDetectionSectionRequest) GetDatasourceName() string {
	if x != nil {
		return x.DatasourceName
	}
	return ""
}

func (x *BulkDetectionSectionRequest) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

type BulkDetectionSectionResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	DetectionSections []*DetectionSection    `protobuf:"bytes,1,rep,name=detection_sections,json=detectionSections,proto3" json:"detection_sections,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *BulkDetectionSectionResponse) Reset() {
	*x = BulkDetectionSectionResponse{}
	mi := &file_services_detection_section_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BulkDetectionSectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkDetectionSectionResponse) ProtoMessage() {}

func (x *BulkDetectionSectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_detection_section_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkDetectionSectionResponse.ProtoReflect.Descriptor instead.
func (*BulkDetectionSectionResponse) Descriptor() ([]byte, []int) {
	return file_services_detection_section_service_proto_rawDescGZIP(), []int{7}
}

func (x *BulkDetectionSectionResponse) GetDetectionSections() []*DetectionSection {
	if x != nil {
		return x.DetectionSections
	}
	return nil
}

var File_services_detection_section_service_proto protoreflect.FileDescriptor

const file_services_detection_section_service_proto_rawDesc = "" +
	"\n" +
	"(services/detection_section.service.proto\x12\x02pb\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cdtos/detection_section.proto\x1a\x0fdtos/file.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x8b\x01\n" +
	"\x1dCreateDetectionSectionRequest\x12'\n" +
	"\x0fdatasource_name\x18\x01 \x01(\tR\x0edatasourceName\x12A\n" +
	"\x11detection_section\x18\x02 \x01(\v2\x14.pb.DetectionSectionR\x10detectionSection\"\x8b\x01\n" +
	"\x1dUpdateDetectionSectionRequest\x12'\n" +
	"\x0fdatasource_name\x18\x01 \x01(\tR\x0edatasourceName\x12A\n" +
	"\x11detection_section\x18\x02 \x01(\v2\x14.pb.DetectionSectionR\x10detectionSection\"X\n" +
	"\x1dDeleteDetectionSectionRequest\x12'\n" +
	"\x0fdatasource_name\x18\x01 \x01(\tR\x0edatasourceName\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\tR\x02id\"U\n" +
	"\x1aGetDetectionSectionRequest\x12'\n" +
	"\x0fdatasource_name\x18\x01 \x01(\tR\x0edatasourceName\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\tR\x02id\"G\n" +
	"\x1cListDetectionSectionsRequest\x12'\n" +
	"\x0fdatasource_name\x18\x01 \x01(\tR\x0edatasourceName\"\x8d\x01\n" +
	"\x1dListDetectionSectionsResponse\x12'\n" +
	"\x0fdatasource_name\x18\x01 \x01(\tR\x0edatasourceName\x12C\n" +
	"\x12detection_sections\x18\x02 \x03(\v2\x14.pb.DetectionSectionR\x11detectionSections\"d\n" +
	"\x1bBulkDetectionSectionRequest\x12'\n" +
	"\x0fdatasource_name\x18\x01 \x01(\tR\x0edatasourceName\x12\x1c\n" +
	"\x04file\x18\x02 \x01(\v2\b.pb.FileR\x04file\"c\n" +
	"\x1cBulkDetectionSectionResponse\x12C\n" +
	"\x12detection_sections\x18\x01 \x03(\v2\x14.pb.DetectionSectionR\x11detectionSections2\xba\b\n" +
	"\x17DetectionSectionService\x12\xad\x01\n" +
	"\x15ListDetectionSections\x12 .pb.ListDetectionSectionsRequest\x1a!.pb.ListDetectionSectionsResponse\"O\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x028\x126/api/v1/datasource/{datasource_name}/detection_section\x12\xa1\x01\n" +
	"\x13GetDetectionSection\x12\x1e.pb.GetDetectionSectionRequest\x1a\x14.pb.DetectionSection\"T\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x02=\x12;/api/v1/datasource/{datasource_name}/detection_section/{id}\x12\xa5\x01\n" +
	"\x16CreateDetectionSection\x12!.pb.CreateDetectionSectionRequest\x1a\x14.pb.DetectionSection\"R\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x02;:\x01*\"6/api/v1/datasource/{datasource_name}/detection_section\x12\xa5\x01\n" +
	"\x16UpdateDetectionSection\x12!.pb.UpdateDetectionSectionRequest\x1a\x14.pb.DetectionSection\"R\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x02;:\x01*\x1a6/api/v1/datasource/{datasource_name}/detection_section\x12\xa9\x01\n" +
	"\x16DeleteDetectionSection\x12!.pb.DeleteDetectionSectionRequest\x1a\x16.google.protobuf.Empty\"T\x92A\x0eb\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x02=*;/api/v1/datasource/{datasource_name}/detection_section/{id}\x12\xce\x01\n" +
	"\x1bBulkCreateDetectionSections\x12\x1f.pb.BulkDetectionSectionRequest\x1a .pb.BulkDetectionSectionResponse\"l\x92A#2\x13multipart/form-datab\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00\x82\xd3\xe4\x93\x02@:\x01*\";/api/v1/datasource/{datasource_name}/detection_section/bulkB\x82\x01\x92AiZY\n" +
	"W\n" +
	"\x06bearer\x12M\b\x02\x128Authentication token, prefixed by Bearer: Bearer <token>\x1a\rAuthorization \x02b\f\n" +
	"\n" +
	"\n" +
	"\x06bearer\x12\x00Z\x14internal/metadata/pbb\x06proto3"

var (
	file_services_detection_section_service_proto_rawDescOnce sync.Once
	file_services_detection_section_service_proto_rawDescData []byte
)

func file_services_detection_section_service_proto_rawDescGZIP() []byte {
	file_services_detection_section_service_proto_rawDescOnce.Do(func() {
		file_services_detection_section_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_detection_section_service_proto_rawDesc), len(file_services_detection_section_service_proto_rawDesc)))
	})
	return file_services_detection_section_service_proto_rawDescData
}

var file_services_detection_section_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_services_detection_section_service_proto_goTypes = []any{
	(*CreateDetectionSectionRequest)(nil), // 0: pb.CreateDetectionSectionRequest
	(*UpdateDetectionSectionRequest)(nil), // 1: pb.UpdateDetectionSectionRequest
	(*DeleteDetectionSectionRequest)(nil), // 2: pb.DeleteDetectionSectionRequest
	(*GetDetectionSectionRequest)(nil),    // 3: pb.GetDetectionSectionRequest
	(*ListDetectionSectionsRequest)(nil),  // 4: pb.ListDetectionSectionsRequest
	(*ListDetectionSectionsResponse)(nil), // 5: pb.ListDetectionSectionsResponse
	(*BulkDetectionSectionRequest)(nil),   // 6: pb.BulkDetectionSectionRequest
	(*BulkDetectionSectionResponse)(nil),  // 7: pb.BulkDetectionSectionResponse
	(*DetectionSection)(nil),              // 8: pb.DetectionSection
	(*File)(nil),                          // 9: pb.File
	(*emptypb.Empty)(nil),                 // 10: google.protobuf.Empty
}
var file_services_detection_section_service_proto_depIdxs = []int32{
	8,  // 0: pb.CreateDetectionSectionRequest.detection_section:type_name -> pb.DetectionSection
	8,  // 1: pb.UpdateDetectionSectionRequest.detection_section:type_name -> pb.DetectionSection
	8,  // 2: pb.ListDetectionSectionsResponse.detection_sections:type_name -> pb.DetectionSection
	9,  // 3: pb.BulkDetectionSectionRequest.file:type_name -> pb.File
	8,  // 4: pb.BulkDetectionSectionResponse.detection_sections:type_name -> pb.DetectionSection
	4,  // 5: pb.DetectionSectionService.ListDetectionSections:input_type -> pb.ListDetectionSectionsRequest
	3,  // 6: pb.DetectionSectionService.GetDetectionSection:input_type -> pb.GetDetectionSectionRequest
	0,  // 7: pb.DetectionSectionService.CreateDetectionSection:input_type -> pb.CreateDetectionSectionRequest
	1,  // 8: pb.DetectionSectionService.UpdateDetectionSection:input_type -> pb.UpdateDetectionSectionRequest
	2,  // 9: pb.DetectionSectionService.DeleteDetectionSection:input_type -> pb.DeleteDetectionSectionRequest
	6,  // 10: pb.DetectionSectionService.BulkCreateDetectionSections:input_type -> pb.BulkDetectionSectionRequest
	5,  // 11: pb.DetectionSectionService.ListDetectionSections:output_type -> pb.ListDetectionSectionsResponse
	8,  // 12: pb.DetectionSectionService.GetDetectionSection:output_type -> pb.DetectionSection
	8,  // 13: pb.DetectionSectionService.CreateDetectionSection:output_type -> pb.DetectionSection
	8,  // 14: pb.DetectionSectionService.UpdateDetectionSection:output_type -> pb.DetectionSection
	10, // 15: pb.DetectionSectionService.DeleteDetectionSection:output_type -> google.protobuf.Empty
	7,  // 16: pb.DetectionSectionService.BulkCreateDetectionSections:output_type -> pb.BulkDetectionSectionResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_services_detection_section_service_proto_init() }
func file_services_detection_section_service_proto_init() {
	if File_services_detection_section_service_proto != nil {
		return
	}
	file_dtos_detection_section_proto_init()
	file_dtos_file_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_detection_section_service_proto_rawDesc), len(file_services_detection_section_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_detection_section_service_proto_goTypes,
		DependencyIndexes: file_services_detection_section_service_proto_depIdxs,
		MessageInfos:      file_services_detection_section_service_proto_msgTypes,
	}.Build()
	File_services_detection_section_service_proto = out.File
	file_services_detection_section_service_proto_goTypes = nil
	file_services_detection_section_service_proto_depIdxs = nil
}
