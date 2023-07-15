// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: team.proto

package events

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Dimension_Type int32

const (
	Dimension_UNSET                Dimension_Type = 0
	Dimension_COLLECTION           Dimension_Type = 1
	Dimension_INGEST               Dimension_Type = 2
	Dimension_TRANSFORM            Dimension_Type = 3
	Dimension_MONITOR              Dimension_Type = 4
	Dimension_STORAGE_STS          Dimension_Type = 5
	Dimension_STORAGE_LTS_INTERNAL Dimension_Type = 6
	Dimension_STORAGE_LTS_EXTERNAL Dimension_Type = 7
	Dimension_REPLAY               Dimension_Type = 8
	Dimension_SOURCE               Dimension_Type = 9
)

// Enum value maps for Dimension_Type.
var (
	Dimension_Type_name = map[int32]string{
		0: "UNSET",
		1: "COLLECTION",
		2: "INGEST",
		3: "TRANSFORM",
		4: "MONITOR",
		5: "STORAGE_STS",
		6: "STORAGE_LTS_INTERNAL",
		7: "STORAGE_LTS_EXTERNAL",
		8: "REPLAY",
		9: "SOURCE",
	}
	Dimension_Type_value = map[string]int32{
		"UNSET":                0,
		"COLLECTION":           1,
		"INGEST":               2,
		"TRANSFORM":            3,
		"MONITOR":              4,
		"STORAGE_STS":          5,
		"STORAGE_LTS_INTERNAL": 6,
		"STORAGE_LTS_EXTERNAL": 7,
		"REPLAY":               8,
		"SOURCE":               9,
	}
)

func (x Dimension_Type) Enum() *Dimension_Type {
	p := new(Dimension_Type)
	*p = x
	return p
}

func (x Dimension_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Dimension_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_team_proto_enumTypes[0].Descriptor()
}

func (Dimension_Type) Type() protoreflect.EnumType {
	return &file_team_proto_enumTypes[0]
}

func (x Dimension_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Dimension_Type.Descriptor instead.
func (Dimension_Type) EnumDescriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{2, 0}
}

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Account *Account `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	Pricing *Pricing `protobuf:"bytes,4,opt,name=pricing,proto3" json:"pricing,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Team) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *Team) GetPricing() *Pricing {
	if x != nil {
		return x.Pricing
	}
	return nil
}

type Pricing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dimensions map[string]*Dimension `protobuf:"bytes,1,rep,name=dimensions,proto3" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Pricing) Reset() {
	*x = Pricing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pricing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pricing) ProtoMessage() {}

func (x *Pricing) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pricing.ProtoReflect.Descriptor instead.
func (*Pricing) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{1}
}

func (x *Pricing) GetDimensions() map[string]*Dimension {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

type Dimension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Dimension_Type `protobuf:"varint,1,opt,name=type,proto3,enum=events.Dimension_Type" json:"type,omitempty"`
	// Whether this feature is enabled and should be billed for the user
	Enabled bool `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Price per byte or per unit of usage
	// Double = float64 in go
	Rate float64 `protobuf:"fixed64,3,opt,name=rate,proto3" json:"rate,omitempty"`
	// How many units are included per collection
	UnitsIncluded int64 `protobuf:"varint,4,opt,name=units_included,json=unitsIncluded,proto3" json:"units_included,omitempty"`
}

func (x *Dimension) Reset() {
	*x = Dimension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dimension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dimension) ProtoMessage() {}

func (x *Dimension) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dimension.ProtoReflect.Descriptor instead.
func (*Dimension) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{2}
}

func (x *Dimension) GetType() Dimension_Type {
	if x != nil {
		return x.Type
	}
	return Dimension_UNSET
}

func (x *Dimension) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Dimension) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *Dimension) GetUnitsIncluded() int64 {
	if x != nil {
		return x.UnitsIncluded
	}
	return 0
}

var File_team_proto protoreflect.FileDescriptor

var file_team_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x1a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x29, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x70,
	0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x70,
	0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x22, 0x9c, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x12, 0x3f, 0x0a, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x50, 0x0a, 0x0f, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb5, 0x02, 0x0a, 0x09, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x49, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x64, 0x22, 0xa6, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a,
	0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4c, 0x4c,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4e, 0x47, 0x45,
	0x53, 0x54, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x4f, 0x52,
	0x4d, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x04,
	0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x53, 0x10,
	0x05, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x4c, 0x54, 0x53,
	0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x53,
	0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x4c, 0x54, 0x53, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52,
	0x4e, 0x41, 0x4c, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x59, 0x10,
	0x08, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x09, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_team_proto_rawDescOnce sync.Once
	file_team_proto_rawDescData = file_team_proto_rawDesc
)

func file_team_proto_rawDescGZIP() []byte {
	file_team_proto_rawDescOnce.Do(func() {
		file_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_team_proto_rawDescData)
	})
	return file_team_proto_rawDescData
}

var file_team_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_team_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_team_proto_goTypes = []interface{}{
	(Dimension_Type)(0), // 0: events.Dimension.Type
	(*Team)(nil),        // 1: events.Team
	(*Pricing)(nil),     // 2: events.Pricing
	(*Dimension)(nil),   // 3: events.Dimension
	nil,                 // 4: events.Pricing.DimensionsEntry
	(*Account)(nil),     // 5: events.Account
}
var file_team_proto_depIdxs = []int32{
	5, // 0: events.Team.account:type_name -> events.Account
	2, // 1: events.Team.pricing:type_name -> events.Pricing
	4, // 2: events.Pricing.dimensions:type_name -> events.Pricing.DimensionsEntry
	0, // 3: events.Dimension.type:type_name -> events.Dimension.Type
	3, // 4: events.Pricing.DimensionsEntry.value:type_name -> events.Dimension
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_team_proto_init() }
func file_team_proto_init() {
	if File_team_proto != nil {
		return
	}
	file_account_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_team_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Team); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_team_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pricing); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_team_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dimension); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_team_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_team_proto_goTypes,
		DependencyIndexes: file_team_proto_depIdxs,
		EnumInfos:         file_team_proto_enumTypes,
		MessageInfos:      file_team_proto_msgTypes,
	}.Build()
	File_team_proto = out.File
	file_team_proto_rawDesc = nil
	file_team_proto_goTypes = nil
	file_team_proto_depIdxs = nil
}
