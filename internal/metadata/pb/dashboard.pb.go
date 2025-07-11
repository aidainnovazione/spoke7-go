// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: dtos/dashboard.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type LayerType int32

const (
	LayerType_LAYER_TYPE_CURRENT_TRAFFIC_LANE    LayerType = 0
	LayerType_LAYER_TYPE_CURRENT_TRAFFIC_POINT   LayerType = 1
	LayerType_LAYER_TYPE_CURRENT_TRAFFIC_SECTION LayerType = 2
	LayerType_LAYER_TYPE_NETWORK                 LayerType = 3
	LayerType_LAYER_TYPE_SPIRA                   LayerType = 4
)

// Enum value maps for LayerType.
var (
	LayerType_name = map[int32]string{
		0: "LAYER_TYPE_CURRENT_TRAFFIC_LANE",
		1: "LAYER_TYPE_CURRENT_TRAFFIC_POINT",
		2: "LAYER_TYPE_CURRENT_TRAFFIC_SECTION",
		3: "LAYER_TYPE_NETWORK",
		4: "LAYER_TYPE_SPIRA",
	}
	LayerType_value = map[string]int32{
		"LAYER_TYPE_CURRENT_TRAFFIC_LANE":    0,
		"LAYER_TYPE_CURRENT_TRAFFIC_POINT":   1,
		"LAYER_TYPE_CURRENT_TRAFFIC_SECTION": 2,
		"LAYER_TYPE_NETWORK":                 3,
		"LAYER_TYPE_SPIRA":                   4,
	}
)

func (x LayerType) Enum() *LayerType {
	p := new(LayerType)
	*p = x
	return p
}

func (x LayerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LayerType) Descriptor() protoreflect.EnumDescriptor {
	return file_dtos_dashboard_proto_enumTypes[0].Descriptor()
}

func (LayerType) Type() protoreflect.EnumType {
	return &file_dtos_dashboard_proto_enumTypes[0]
}

func (x LayerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LayerType.Descriptor instead.
func (LayerType) EnumDescriptor() ([]byte, []int) {
	return file_dtos_dashboard_proto_rawDescGZIP(), []int{0}
}

type MapTheme int32

const (
	MapTheme_LIGHT MapTheme = 0
	MapTheme_DARK  MapTheme = 1
)

// Enum value maps for MapTheme.
var (
	MapTheme_name = map[int32]string{
		0: "LIGHT",
		1: "DARK",
	}
	MapTheme_value = map[string]int32{
		"LIGHT": 0,
		"DARK":  1,
	}
)

func (x MapTheme) Enum() *MapTheme {
	p := new(MapTheme)
	*p = x
	return p
}

func (x MapTheme) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MapTheme) Descriptor() protoreflect.EnumDescriptor {
	return file_dtos_dashboard_proto_enumTypes[1].Descriptor()
}

func (MapTheme) Type() protoreflect.EnumType {
	return &file_dtos_dashboard_proto_enumTypes[1]
}

func (x MapTheme) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MapTheme.Descriptor instead.
func (MapTheme) EnumDescriptor() ([]byte, []int) {
	return file_dtos_dashboard_proto_rawDescGZIP(), []int{1}
}

type Dashboard struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description    string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DataSourceName string                 `protobuf:"bytes,4,opt,name=data_source_name,json=dataSourceName,proto3" json:"data_source_name,omitempty"`
	Sections       []*Section             `protobuf:"bytes,5,rep,name=sections,proto3" json:"sections,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ModifiedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
	Owner          string                 `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	Groups         []string               `protobuf:"bytes,9,rep,name=groups,proto3" json:"groups,omitempty"`
	Bottombar      *Bottombar             `protobuf:"bytes,10,opt,name=bottombar,proto3" json:"bottombar,omitempty"`
	Sidebar        *Sidebar               `protobuf:"bytes,11,opt,name=sidebar,proto3" json:"sidebar,omitempty"`
	BaseMap        *BaseMap               `protobuf:"bytes,12,opt,name=base_map,json=baseMap,proto3" json:"base_map,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Dashboard) Reset() {
	*x = Dashboard{}
	mi := &file_dtos_dashboard_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Dashboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dashboard) ProtoMessage() {}

func (x *Dashboard) ProtoReflect() protoreflect.Message {
	mi := &file_dtos_dashboard_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dashboard.ProtoReflect.Descriptor instead.
func (*Dashboard) Descriptor() ([]byte, []int) {
	return file_dtos_dashboard_proto_rawDescGZIP(), []int{0}
}

func (x *Dashboard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Dashboard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dashboard) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Dashboard) GetDataSourceName() string {
	if x != nil {
		return x.DataSourceName
	}
	return ""
}

func (x *Dashboard) GetSections() []*Section {
	if x != nil {
		return x.Sections
	}
	return nil
}

func (x *Dashboard) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Dashboard) GetModifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedAt
	}
	return nil
}

func (x *Dashboard) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Dashboard) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *Dashboard) GetBottombar() *Bottombar {
	if x != nil {
		return x.Bottombar
	}
	return nil
}

func (x *Dashboard) GetSidebar() *Sidebar {
	if x != nil {
		return x.Sidebar
	}
	return nil
}

func (x *Dashboard) GetBaseMap() *BaseMap {
	if x != nil {
		return x.BaseMap
	}
	return nil
}

type Section struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsVisible     bool                   `protobuf:"varint,2,opt,name=is_visible,json=isVisible,proto3" json:"is_visible,omitempty"`
	IsExpanded    bool                   `protobuf:"varint,3,opt,name=is_expanded,json=isExpanded,proto3" json:"is_expanded,omitempty"`
	LayerType     LayerType              `protobuf:"varint,4,opt,name=layer_type,json=layerType,proto3,enum=pb.LayerType" json:"layer_type,omitempty"`
	Params        *structpb.Struct       `protobuf:"bytes,5,opt,name=params,proto3" json:"params,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Section) Reset() {
	*x = Section{}
	mi := &file_dtos_dashboard_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Section) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Section) ProtoMessage() {}

func (x *Section) ProtoReflect() protoreflect.Message {
	mi := &file_dtos_dashboard_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Section.ProtoReflect.Descriptor instead.
func (*Section) Descriptor() ([]byte, []int) {
	return file_dtos_dashboard_proto_rawDescGZIP(), []int{1}
}

func (x *Section) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Section) GetIsVisible() bool {
	if x != nil {
		return x.IsVisible
	}
	return false
}

func (x *Section) GetIsExpanded() bool {
	if x != nil {
		return x.IsExpanded
	}
	return false
}

func (x *Section) GetLayerType() LayerType {
	if x != nil {
		return x.LayerType
	}
	return LayerType_LAYER_TYPE_CURRENT_TRAFFIC_LANE
}

func (x *Section) GetParams() *structpb.Struct {
	if x != nil {
		return x.Params
	}
	return nil
}

type Sidebar struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	IsOpen            bool                   `protobuf:"varint,1,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	InjectedComponent string                 `protobuf:"bytes,2,opt,name=injected_component,json=injectedComponent,proto3" json:"injected_component,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Sidebar) Reset() {
	*x = Sidebar{}
	mi := &file_dtos_dashboard_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sidebar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sidebar) ProtoMessage() {}

func (x *Sidebar) ProtoReflect() protoreflect.Message {
	mi := &file_dtos_dashboard_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sidebar.ProtoReflect.Descriptor instead.
func (*Sidebar) Descriptor() ([]byte, []int) {
	return file_dtos_dashboard_proto_rawDescGZIP(), []int{2}
}

func (x *Sidebar) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *Sidebar) GetInjectedComponent() string {
	if x != nil {
		return x.InjectedComponent
	}
	return ""
}

type Bottombar struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsOpen        bool                   `protobuf:"varint,1,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	SelectedTime  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=selected_time,json=selectedTime,proto3" json:"selected_time,omitempty"`
	StartTime     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Interval      int64                  `protobuf:"varint,5,opt,name=interval,proto3" json:"interval,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bottombar) Reset() {
	*x = Bottombar{}
	mi := &file_dtos_dashboard_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bottombar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bottombar) ProtoMessage() {}

func (x *Bottombar) ProtoReflect() protoreflect.Message {
	mi := &file_dtos_dashboard_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bottombar.ProtoReflect.Descriptor instead.
func (*Bottombar) Descriptor() ([]byte, []int) {
	return file_dtos_dashboard_proto_rawDescGZIP(), []int{3}
}

func (x *Bottombar) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *Bottombar) GetSelectedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SelectedTime
	}
	return nil
}

func (x *Bottombar) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Bottombar) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Bottombar) GetInterval() int64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

type BaseMap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MapTheme      MapTheme               `protobuf:"varint,1,opt,name=map_theme,json=mapTheme,proto3,enum=pb.MapTheme" json:"map_theme,omitempty"`
	Center        []float64              `protobuf:"fixed64,2,rep,packed,name=center,proto3" json:"center,omitempty"` // [latitude, longitude]
	Zoom          float64                `protobuf:"fixed64,3,opt,name=zoom,proto3" json:"zoom,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BaseMap) Reset() {
	*x = BaseMap{}
	mi := &file_dtos_dashboard_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseMap) ProtoMessage() {}

func (x *BaseMap) ProtoReflect() protoreflect.Message {
	mi := &file_dtos_dashboard_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseMap.ProtoReflect.Descriptor instead.
func (*BaseMap) Descriptor() ([]byte, []int) {
	return file_dtos_dashboard_proto_rawDescGZIP(), []int{4}
}

func (x *BaseMap) GetMapTheme() MapTheme {
	if x != nil {
		return x.MapTheme
	}
	return MapTheme_LIGHT
}

func (x *BaseMap) GetCenter() []float64 {
	if x != nil {
		return x.Center
	}
	return nil
}

func (x *BaseMap) GetZoom() float64 {
	if x != nil {
		return x.Zoom
	}
	return 0
}

var File_dtos_dashboard_proto protoreflect.FileDescriptor

const file_dtos_dashboard_proto_rawDesc = "" +
	"\n" +
	"\x14dtos/dashboard.proto\x12\x02pb\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xc6\x03\n" +
	"\tDashboard\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12(\n" +
	"\x10data_source_name\x18\x04 \x01(\tR\x0edataSourceName\x12'\n" +
	"\bsections\x18\x05 \x03(\v2\v.pb.SectionR\bsections\x129\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12;\n" +
	"\vmodified_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"modifiedAt\x12\x14\n" +
	"\x05owner\x18\b \x01(\tR\x05owner\x12\x16\n" +
	"\x06groups\x18\t \x03(\tR\x06groups\x12+\n" +
	"\tbottombar\x18\n" +
	" \x01(\v2\r.pb.BottombarR\tbottombar\x12%\n" +
	"\asidebar\x18\v \x01(\v2\v.pb.SidebarR\asidebar\x12&\n" +
	"\bbase_map\x18\f \x01(\v2\v.pb.BaseMapR\abaseMap\"\xbc\x01\n" +
	"\aSection\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1d\n" +
	"\n" +
	"is_visible\x18\x02 \x01(\bR\tisVisible\x12\x1f\n" +
	"\vis_expanded\x18\x03 \x01(\bR\n" +
	"isExpanded\x12,\n" +
	"\n" +
	"layer_type\x18\x04 \x01(\x0e2\r.pb.LayerTypeR\tlayerType\x12/\n" +
	"\x06params\x18\x05 \x01(\v2\x17.google.protobuf.StructR\x06params\"Q\n" +
	"\aSidebar\x12\x17\n" +
	"\ais_open\x18\x01 \x01(\bR\x06isOpen\x12-\n" +
	"\x12injected_component\x18\x02 \x01(\tR\x11injectedComponent\"\xf3\x01\n" +
	"\tBottombar\x12\x17\n" +
	"\ais_open\x18\x01 \x01(\bR\x06isOpen\x12?\n" +
	"\rselected_time\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\fselectedTime\x129\n" +
	"\n" +
	"start_time\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tstartTime\x125\n" +
	"\bend_time\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\aendTime\x12\x1a\n" +
	"\binterval\x18\x05 \x01(\x03R\binterval\"`\n" +
	"\aBaseMap\x12)\n" +
	"\tmap_theme\x18\x01 \x01(\x0e2\f.pb.MapThemeR\bmapTheme\x12\x16\n" +
	"\x06center\x18\x02 \x03(\x01R\x06center\x12\x12\n" +
	"\x04zoom\x18\x03 \x01(\x01R\x04zoom*\xac\x01\n" +
	"\tLayerType\x12#\n" +
	"\x1fLAYER_TYPE_CURRENT_TRAFFIC_LANE\x10\x00\x12$\n" +
	" LAYER_TYPE_CURRENT_TRAFFIC_POINT\x10\x01\x12&\n" +
	"\"LAYER_TYPE_CURRENT_TRAFFIC_SECTION\x10\x02\x12\x16\n" +
	"\x12LAYER_TYPE_NETWORK\x10\x03\x12\x14\n" +
	"\x10LAYER_TYPE_SPIRA\x10\x04*\x1f\n" +
	"\bMapTheme\x12\t\n" +
	"\x05LIGHT\x10\x00\x12\b\n" +
	"\x04DARK\x10\x01B\x16Z\x14internal/metadata/pbb\x06proto3"

var (
	file_dtos_dashboard_proto_rawDescOnce sync.Once
	file_dtos_dashboard_proto_rawDescData []byte
)

func file_dtos_dashboard_proto_rawDescGZIP() []byte {
	file_dtos_dashboard_proto_rawDescOnce.Do(func() {
		file_dtos_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_dtos_dashboard_proto_rawDesc), len(file_dtos_dashboard_proto_rawDesc)))
	})
	return file_dtos_dashboard_proto_rawDescData
}

var file_dtos_dashboard_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_dtos_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dtos_dashboard_proto_goTypes = []any{
	(LayerType)(0),                // 0: pb.LayerType
	(MapTheme)(0),                 // 1: pb.MapTheme
	(*Dashboard)(nil),             // 2: pb.Dashboard
	(*Section)(nil),               // 3: pb.Section
	(*Sidebar)(nil),               // 4: pb.Sidebar
	(*Bottombar)(nil),             // 5: pb.Bottombar
	(*BaseMap)(nil),               // 6: pb.BaseMap
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*structpb.Struct)(nil),       // 8: google.protobuf.Struct
}
var file_dtos_dashboard_proto_depIdxs = []int32{
	3,  // 0: pb.Dashboard.sections:type_name -> pb.Section
	7,  // 1: pb.Dashboard.created_at:type_name -> google.protobuf.Timestamp
	7,  // 2: pb.Dashboard.modified_at:type_name -> google.protobuf.Timestamp
	5,  // 3: pb.Dashboard.bottombar:type_name -> pb.Bottombar
	4,  // 4: pb.Dashboard.sidebar:type_name -> pb.Sidebar
	6,  // 5: pb.Dashboard.base_map:type_name -> pb.BaseMap
	0,  // 6: pb.Section.layer_type:type_name -> pb.LayerType
	8,  // 7: pb.Section.params:type_name -> google.protobuf.Struct
	7,  // 8: pb.Bottombar.selected_time:type_name -> google.protobuf.Timestamp
	7,  // 9: pb.Bottombar.start_time:type_name -> google.protobuf.Timestamp
	7,  // 10: pb.Bottombar.end_time:type_name -> google.protobuf.Timestamp
	1,  // 11: pb.BaseMap.map_theme:type_name -> pb.MapTheme
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_dtos_dashboard_proto_init() }
func file_dtos_dashboard_proto_init() {
	if File_dtos_dashboard_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_dtos_dashboard_proto_rawDesc), len(file_dtos_dashboard_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dtos_dashboard_proto_goTypes,
		DependencyIndexes: file_dtos_dashboard_proto_depIdxs,
		EnumInfos:         file_dtos_dashboard_proto_enumTypes,
		MessageInfos:      file_dtos_dashboard_proto_msgTypes,
	}.Build()
	File_dtos_dashboard_proto = out.File
	file_dtos_dashboard_proto_goTypes = nil
	file_dtos_dashboard_proto_depIdxs = nil
}
