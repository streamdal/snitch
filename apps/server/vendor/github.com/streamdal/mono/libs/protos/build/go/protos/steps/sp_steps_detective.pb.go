// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: steps/sp_steps_detective.proto

package steps

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

type DetectiveType int32

const (
	DetectiveType_DETECTIVE_TYPE_UNKNOWN             DetectiveType = 0
	DetectiveType_DETECTIVE_TYPE_IS_EMPTY            DetectiveType = 1000
	DetectiveType_DETECTIVE_TYPE_HAS_FIELD           DetectiveType = 1001
	DetectiveType_DETECTIVE_TYPE_IS_TYPE             DetectiveType = 1002
	DetectiveType_DETECTIVE_TYPE_STRING_CONTAINS_ANY DetectiveType = 1003
	DetectiveType_DETECTIVE_TYPE_STRING_CONTAINS_ALL DetectiveType = 1004
	DetectiveType_DETECTIVE_TYPE_STRING_EQUAL        DetectiveType = 1005
	DetectiveType_DETECTIVE_TYPE_IPV4_ADDRESS        DetectiveType = 1006
	DetectiveType_DETECTIVE_TYPE_IPV6_ADDRESS        DetectiveType = 1007
	DetectiveType_DETECTIVE_TYPE_MAC_ADDRESS         DetectiveType = 1008
	DetectiveType_DETECTIVE_TYPE_REGEX               DetectiveType = 1009
	DetectiveType_DETECTIVE_TYPE_TIMESTAMP_RFC3339   DetectiveType = 1010
	DetectiveType_DETECTIVE_TYPE_TIMESTAMP_UNIX_NANO DetectiveType = 1011
	DetectiveType_DETECTIVE_TYPE_TIMESTAMP_UNIX      DetectiveType = 1012
	DetectiveType_DETECTIVE_TYPE_BOOLEAN_TRUE        DetectiveType = 1013
	DetectiveType_DETECTIVE_TYPE_BOOLEAN_FALSE       DetectiveType = 1014
	DetectiveType_DETECTIVE_TYPE_UUID                DetectiveType = 1015
	DetectiveType_DETECTIVE_TYPE_URL                 DetectiveType = 1016
	DetectiveType_DETECTIVE_TYPE_HOSTNAME            DetectiveType = 1017
	DetectiveType_DETECTIVE_TYPE_STRING_LENGTH_MIN   DetectiveType = 1018
	DetectiveType_DETECTIVE_TYPE_STRING_LENGTH_MAX   DetectiveType = 1019
	DetectiveType_DETECTIVE_TYPE_STRING_LENGTH_RANGE DetectiveType = 1020
	DetectiveType_DETECTIVE_TYPE_SEMVER              DetectiveType = 1021
	// / Payloads containing values with any PII - runs all PII matchers
	DetectiveType_DETECTIVE_TYPE_PII_ANY DetectiveType = 2000
	// Payloads containing values with a credit card number
	DetectiveType_DETECTIVE_TYPE_PII_CREDIT_CARD DetectiveType = 2001
	// Payloads containing values with a social security number
	DetectiveType_DETECTIVE_TYPE_PII_SSN DetectiveType = 2002
	// Payloads containing values with an email address
	DetectiveType_DETECTIVE_TYPE_PII_EMAIL DetectiveType = 2003
	// Payloads containing values with a phone number
	DetectiveType_DETECTIVE_TYPE_PII_PHONE DetectiveType = 2004
	// Payloads containing values with a driver's license
	DetectiveType_DETECTIVE_TYPE_PII_DRIVER_LICENSE DetectiveType = 2005
	// Payloads containing values with a passport ID
	DetectiveType_DETECTIVE_TYPE_PII_PASSPORT_ID DetectiveType = 2006
	// Payloads containing values with a VIN number
	DetectiveType_DETECTIVE_TYPE_PII_VIN_NUMBER DetectiveType = 2007
	// Payloads containing values with various serial number formats
	DetectiveType_DETECTIVE_TYPE_PII_SERIAL_NUMBER DetectiveType = 2008
	// Payloads containing fields named "login", "username", "user", "userid", "user_id", "user", "password", "pass", "passwd", "pwd"
	DetectiveType_DETECTIVE_TYPE_PII_LOGIN DetectiveType = 2009
	// Payloads containing fields named "taxpayer_id", "tax_id", "taxpayerid", "taxid"
	DetectiveType_DETECTIVE_TYPE_PII_TAXPAYER_ID DetectiveType = 2010
	// Payloads containing fields named "address", "street", "city", "state", "zip", "zipcode", "zip_code", "country"
	DetectiveType_DETECTIVE_TYPE_PII_ADDRESS DetectiveType = 2011
	// Payloads containing fields named "signature", "signature_image", "signature_image_url", "signature_image_uri"
	DetectiveType_DETECTIVE_TYPE_PII_SIGNATURE DetectiveType = 2012
	// Payloads containing values that contain GPS data or coordinates like "lat", "lon", "latitude", "longitude"
	DetectiveType_DETECTIVE_TYPE_PII_GEOLOCATION DetectiveType = 2013
	// Payloads containing fields like "school", "university", "college", "education"
	DetectiveType_DETECTIVE_TYPE_PII_EDUCATION DetectiveType = 2014
	// Payloads containing fields like "account", "bank", "credit", "debit", "financial", "finance"
	DetectiveType_DETECTIVE_TYPE_PII_FINANCIAL DetectiveType = 2015
	// Payloads containing fields like "patient", "health", "healthcare", "health care", "medical"
	DetectiveType_DETECTIVE_TYPE_PII_HEALTH            DetectiveType = 2016
	DetectiveType_DETECTIVE_TYPE_NUMERIC_EQUAL_TO      DetectiveType = 3000
	DetectiveType_DETECTIVE_TYPE_NUMERIC_GREATER_THAN  DetectiveType = 3001
	DetectiveType_DETECTIVE_TYPE_NUMERIC_GREATER_EQUAL DetectiveType = 3002
	DetectiveType_DETECTIVE_TYPE_NUMERIC_LESS_THAN     DetectiveType = 3003
	DetectiveType_DETECTIVE_TYPE_NUMERIC_LESS_EQUAL    DetectiveType = 3004
	DetectiveType_DETECTIVE_TYPE_NUMERIC_RANGE         DetectiveType = 3005
	DetectiveType_DETECTIVE_TYPE_NUMERIC_MIN           DetectiveType = 3006
	DetectiveType_DETECTIVE_TYPE_NUMERIC_MAX           DetectiveType = 3007
)

// Enum value maps for DetectiveType.
var (
	DetectiveType_name = map[int32]string{
		0:    "DETECTIVE_TYPE_UNKNOWN",
		1000: "DETECTIVE_TYPE_IS_EMPTY",
		1001: "DETECTIVE_TYPE_HAS_FIELD",
		1002: "DETECTIVE_TYPE_IS_TYPE",
		1003: "DETECTIVE_TYPE_STRING_CONTAINS_ANY",
		1004: "DETECTIVE_TYPE_STRING_CONTAINS_ALL",
		1005: "DETECTIVE_TYPE_STRING_EQUAL",
		1006: "DETECTIVE_TYPE_IPV4_ADDRESS",
		1007: "DETECTIVE_TYPE_IPV6_ADDRESS",
		1008: "DETECTIVE_TYPE_MAC_ADDRESS",
		1009: "DETECTIVE_TYPE_REGEX",
		1010: "DETECTIVE_TYPE_TIMESTAMP_RFC3339",
		1011: "DETECTIVE_TYPE_TIMESTAMP_UNIX_NANO",
		1012: "DETECTIVE_TYPE_TIMESTAMP_UNIX",
		1013: "DETECTIVE_TYPE_BOOLEAN_TRUE",
		1014: "DETECTIVE_TYPE_BOOLEAN_FALSE",
		1015: "DETECTIVE_TYPE_UUID",
		1016: "DETECTIVE_TYPE_URL",
		1017: "DETECTIVE_TYPE_HOSTNAME",
		1018: "DETECTIVE_TYPE_STRING_LENGTH_MIN",
		1019: "DETECTIVE_TYPE_STRING_LENGTH_MAX",
		1020: "DETECTIVE_TYPE_STRING_LENGTH_RANGE",
		1021: "DETECTIVE_TYPE_SEMVER",
		2000: "DETECTIVE_TYPE_PII_ANY",
		2001: "DETECTIVE_TYPE_PII_CREDIT_CARD",
		2002: "DETECTIVE_TYPE_PII_SSN",
		2003: "DETECTIVE_TYPE_PII_EMAIL",
		2004: "DETECTIVE_TYPE_PII_PHONE",
		2005: "DETECTIVE_TYPE_PII_DRIVER_LICENSE",
		2006: "DETECTIVE_TYPE_PII_PASSPORT_ID",
		2007: "DETECTIVE_TYPE_PII_VIN_NUMBER",
		2008: "DETECTIVE_TYPE_PII_SERIAL_NUMBER",
		2009: "DETECTIVE_TYPE_PII_LOGIN",
		2010: "DETECTIVE_TYPE_PII_TAXPAYER_ID",
		2011: "DETECTIVE_TYPE_PII_ADDRESS",
		2012: "DETECTIVE_TYPE_PII_SIGNATURE",
		2013: "DETECTIVE_TYPE_PII_GEOLOCATION",
		2014: "DETECTIVE_TYPE_PII_EDUCATION",
		2015: "DETECTIVE_TYPE_PII_FINANCIAL",
		2016: "DETECTIVE_TYPE_PII_HEALTH",
		3000: "DETECTIVE_TYPE_NUMERIC_EQUAL_TO",
		3001: "DETECTIVE_TYPE_NUMERIC_GREATER_THAN",
		3002: "DETECTIVE_TYPE_NUMERIC_GREATER_EQUAL",
		3003: "DETECTIVE_TYPE_NUMERIC_LESS_THAN",
		3004: "DETECTIVE_TYPE_NUMERIC_LESS_EQUAL",
		3005: "DETECTIVE_TYPE_NUMERIC_RANGE",
		3006: "DETECTIVE_TYPE_NUMERIC_MIN",
		3007: "DETECTIVE_TYPE_NUMERIC_MAX",
	}
	DetectiveType_value = map[string]int32{
		"DETECTIVE_TYPE_UNKNOWN":               0,
		"DETECTIVE_TYPE_IS_EMPTY":              1000,
		"DETECTIVE_TYPE_HAS_FIELD":             1001,
		"DETECTIVE_TYPE_IS_TYPE":               1002,
		"DETECTIVE_TYPE_STRING_CONTAINS_ANY":   1003,
		"DETECTIVE_TYPE_STRING_CONTAINS_ALL":   1004,
		"DETECTIVE_TYPE_STRING_EQUAL":          1005,
		"DETECTIVE_TYPE_IPV4_ADDRESS":          1006,
		"DETECTIVE_TYPE_IPV6_ADDRESS":          1007,
		"DETECTIVE_TYPE_MAC_ADDRESS":           1008,
		"DETECTIVE_TYPE_REGEX":                 1009,
		"DETECTIVE_TYPE_TIMESTAMP_RFC3339":     1010,
		"DETECTIVE_TYPE_TIMESTAMP_UNIX_NANO":   1011,
		"DETECTIVE_TYPE_TIMESTAMP_UNIX":        1012,
		"DETECTIVE_TYPE_BOOLEAN_TRUE":          1013,
		"DETECTIVE_TYPE_BOOLEAN_FALSE":         1014,
		"DETECTIVE_TYPE_UUID":                  1015,
		"DETECTIVE_TYPE_URL":                   1016,
		"DETECTIVE_TYPE_HOSTNAME":              1017,
		"DETECTIVE_TYPE_STRING_LENGTH_MIN":     1018,
		"DETECTIVE_TYPE_STRING_LENGTH_MAX":     1019,
		"DETECTIVE_TYPE_STRING_LENGTH_RANGE":   1020,
		"DETECTIVE_TYPE_SEMVER":                1021,
		"DETECTIVE_TYPE_PII_ANY":               2000,
		"DETECTIVE_TYPE_PII_CREDIT_CARD":       2001,
		"DETECTIVE_TYPE_PII_SSN":               2002,
		"DETECTIVE_TYPE_PII_EMAIL":             2003,
		"DETECTIVE_TYPE_PII_PHONE":             2004,
		"DETECTIVE_TYPE_PII_DRIVER_LICENSE":    2005,
		"DETECTIVE_TYPE_PII_PASSPORT_ID":       2006,
		"DETECTIVE_TYPE_PII_VIN_NUMBER":        2007,
		"DETECTIVE_TYPE_PII_SERIAL_NUMBER":     2008,
		"DETECTIVE_TYPE_PII_LOGIN":             2009,
		"DETECTIVE_TYPE_PII_TAXPAYER_ID":       2010,
		"DETECTIVE_TYPE_PII_ADDRESS":           2011,
		"DETECTIVE_TYPE_PII_SIGNATURE":         2012,
		"DETECTIVE_TYPE_PII_GEOLOCATION":       2013,
		"DETECTIVE_TYPE_PII_EDUCATION":         2014,
		"DETECTIVE_TYPE_PII_FINANCIAL":         2015,
		"DETECTIVE_TYPE_PII_HEALTH":            2016,
		"DETECTIVE_TYPE_NUMERIC_EQUAL_TO":      3000,
		"DETECTIVE_TYPE_NUMERIC_GREATER_THAN":  3001,
		"DETECTIVE_TYPE_NUMERIC_GREATER_EQUAL": 3002,
		"DETECTIVE_TYPE_NUMERIC_LESS_THAN":     3003,
		"DETECTIVE_TYPE_NUMERIC_LESS_EQUAL":    3004,
		"DETECTIVE_TYPE_NUMERIC_RANGE":         3005,
		"DETECTIVE_TYPE_NUMERIC_MIN":           3006,
		"DETECTIVE_TYPE_NUMERIC_MAX":           3007,
	}
)

func (x DetectiveType) Enum() *DetectiveType {
	p := new(DetectiveType)
	*p = x
	return p
}

func (x DetectiveType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DetectiveType) Descriptor() protoreflect.EnumDescriptor {
	return file_steps_sp_steps_detective_proto_enumTypes[0].Descriptor()
}

func (DetectiveType) Type() protoreflect.EnumType {
	return &file_steps_sp_steps_detective_proto_enumTypes[0]
}

func (x DetectiveType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DetectiveType.Descriptor instead.
func (DetectiveType) EnumDescriptor() ([]byte, []int) {
	return file_steps_sp_steps_detective_proto_rawDescGZIP(), []int{0}
}

type DetectiveStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path   *string       `protobuf:"bytes,1,opt,name=path,proto3,oneof" json:"path,omitempty"`
	Args   []string      `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"` // args determined by match_type
	Negate *bool         `protobuf:"varint,3,opt,name=negate,proto3,oneof" json:"negate,omitempty"`
	Type   DetectiveType `protobuf:"varint,4,opt,name=type,proto3,enum=protos.steps.DetectiveType" json:"type,omitempty"`
}

func (x *DetectiveStep) Reset() {
	*x = DetectiveStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steps_sp_steps_detective_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectiveStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectiveStep) ProtoMessage() {}

func (x *DetectiveStep) ProtoReflect() protoreflect.Message {
	mi := &file_steps_sp_steps_detective_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectiveStep.ProtoReflect.Descriptor instead.
func (*DetectiveStep) Descriptor() ([]byte, []int) {
	return file_steps_sp_steps_detective_proto_rawDescGZIP(), []int{0}
}

func (x *DetectiveStep) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *DetectiveStep) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *DetectiveStep) GetNegate() bool {
	if x != nil && x.Negate != nil {
		return *x.Negate
	}
	return false
}

func (x *DetectiveStep) GetType() DetectiveType {
	if x != nil {
		return x.Type
	}
	return DetectiveType_DETECTIVE_TYPE_UNKNOWN
}

var File_steps_sp_steps_detective_proto protoreflect.FileDescriptor

var file_steps_sp_steps_detective_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2f, 0x73, 0x70, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73,
	0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x73, 0x22, 0x9e,
	0x01, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x65, 0x70,
	0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x1b, 0x0a,
	0x06, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52,
	0x06, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2a,
	0x8d, 0x0d, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1c, 0x0a,
	0x17, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x49, 0x53, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0xe8, 0x07, 0x12, 0x1d, 0x0a, 0x18, 0x44,
	0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x41,
	0x53, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0xe9, 0x07, 0x12, 0x1b, 0x0a, 0x16, 0x44, 0x45,
	0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x53, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0xea, 0x07, 0x12, 0x27, 0x0a, 0x22, 0x44, 0x45, 0x54, 0x45, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47,
	0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x5f, 0x41, 0x4e, 0x59, 0x10, 0xeb, 0x07,
	0x12, 0x27, 0x0a, 0x22, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49,
	0x4e, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0xec, 0x07, 0x12, 0x20, 0x0a, 0x1b, 0x44, 0x45, 0x54,
	0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0xed, 0x07, 0x12, 0x20, 0x0a, 0x1b, 0x44,
	0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x50,
	0x56, 0x34, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0xee, 0x07, 0x12, 0x20, 0x0a,
	0x1b, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x49, 0x50, 0x56, 0x36, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0xef, 0x07, 0x12,
	0x1f, 0x0a, 0x1a, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x41, 0x43, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0xf0, 0x07,
	0x12, 0x19, 0x0a, 0x14, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x45, 0x58, 0x10, 0xf1, 0x07, 0x12, 0x25, 0x0a, 0x20, 0x44,
	0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x52, 0x46, 0x43, 0x33, 0x33, 0x33, 0x39, 0x10,
	0xf2, 0x07, 0x12, 0x27, 0x0a, 0x22, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x55,
	0x4e, 0x49, 0x58, 0x5f, 0x4e, 0x41, 0x4e, 0x4f, 0x10, 0xf3, 0x07, 0x12, 0x22, 0x0a, 0x1d, 0x44,
	0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x55, 0x4e, 0x49, 0x58, 0x10, 0xf4, 0x07, 0x12,
	0x20, 0x0a, 0x1b, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x5f, 0x54, 0x52, 0x55, 0x45, 0x10, 0xf5,
	0x07, 0x12, 0x21, 0x0a, 0x1c, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x5f, 0x46, 0x41, 0x4c, 0x53,
	0x45, 0x10, 0xf6, 0x07, 0x12, 0x18, 0x0a, 0x13, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x55, 0x49, 0x44, 0x10, 0xf7, 0x07, 0x12, 0x17,
	0x0a, 0x12, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x52, 0x4c, 0x10, 0xf8, 0x07, 0x12, 0x1c, 0x0a, 0x17, 0x44, 0x45, 0x54, 0x45, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x4e, 0x41,
	0x4d, 0x45, 0x10, 0xf9, 0x07, 0x12, 0x25, 0x0a, 0x20, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4c,
	0x45, 0x4e, 0x47, 0x54, 0x48, 0x5f, 0x4d, 0x49, 0x4e, 0x10, 0xfa, 0x07, 0x12, 0x25, 0x0a, 0x20,
	0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53,
	0x54, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x45, 0x4e, 0x47, 0x54, 0x48, 0x5f, 0x4d, 0x41, 0x58,
	0x10, 0xfb, 0x07, 0x12, 0x27, 0x0a, 0x22, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x45, 0x4e,
	0x47, 0x54, 0x48, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0xfc, 0x07, 0x12, 0x1a, 0x0a, 0x15,
	0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53,
	0x45, 0x4d, 0x56, 0x45, 0x52, 0x10, 0xfd, 0x07, 0x12, 0x1b, 0x0a, 0x16, 0x44, 0x45, 0x54, 0x45,
	0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x41,
	0x4e, 0x59, 0x10, 0xd0, 0x0f, 0x12, 0x23, 0x0a, 0x1e, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x43, 0x52, 0x45, 0x44,
	0x49, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0xd1, 0x0f, 0x12, 0x1b, 0x0a, 0x16, 0x44, 0x45,
	0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49,
	0x5f, 0x53, 0x53, 0x4e, 0x10, 0xd2, 0x0f, 0x12, 0x1d, 0x0a, 0x18, 0x44, 0x45, 0x54, 0x45, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x45, 0x4d,
	0x41, 0x49, 0x4c, 0x10, 0xd3, 0x0f, 0x12, 0x1d, 0x0a, 0x18, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54,
	0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x50, 0x48, 0x4f,
	0x4e, 0x45, 0x10, 0xd4, 0x0f, 0x12, 0x26, 0x0a, 0x21, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x44, 0x52, 0x49, 0x56,
	0x45, 0x52, 0x5f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x10, 0xd5, 0x0f, 0x12, 0x23, 0x0a,
	0x1e, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x50, 0x49, 0x49, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x49, 0x44, 0x10,
	0xd6, 0x0f, 0x12, 0x22, 0x0a, 0x1d, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x56, 0x49, 0x4e, 0x5f, 0x4e, 0x55, 0x4d,
	0x42, 0x45, 0x52, 0x10, 0xd7, 0x0f, 0x12, 0x25, 0x0a, 0x20, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54,
	0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x53, 0x45, 0x52,
	0x49, 0x41, 0x4c, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0xd8, 0x0f, 0x12, 0x1d, 0x0a,
	0x18, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x50, 0x49, 0x49, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0xd9, 0x0f, 0x12, 0x23, 0x0a, 0x1e,
	0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x49, 0x49, 0x5f, 0x54, 0x41, 0x58, 0x50, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x10, 0xda,
	0x0f, 0x12, 0x1f, 0x0a, 0x1a, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10,
	0xdb, 0x0f, 0x12, 0x21, 0x0a, 0x1c, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55,
	0x52, 0x45, 0x10, 0xdc, 0x0f, 0x12, 0x23, 0x0a, 0x1e, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x47, 0x45, 0x4f, 0x4c,
	0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xdd, 0x0f, 0x12, 0x21, 0x0a, 0x1c, 0x44, 0x45,
	0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49,
	0x5f, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xde, 0x0f, 0x12, 0x21, 0x0a,
	0x1c, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x50, 0x49, 0x49, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4e, 0x43, 0x49, 0x41, 0x4c, 0x10, 0xdf, 0x0f,
	0x12, 0x1e, 0x0a, 0x19, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x10, 0xe0, 0x0f,
	0x12, 0x24, 0x0a, 0x1f, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c,
	0x5f, 0x54, 0x4f, 0x10, 0xb8, 0x17, 0x12, 0x28, 0x0a, 0x23, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54,
	0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x49, 0x43,
	0x5f, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10, 0xb9, 0x17,
	0x12, 0x29, 0x0a, 0x24, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x47, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x52, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0xba, 0x17, 0x12, 0x25, 0x0a, 0x20, 0x44,
	0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55,
	0x4d, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10,
	0xbb, 0x17, 0x12, 0x26, 0x0a, 0x21, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x4c, 0x45, 0x53,
	0x53, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0xbc, 0x17, 0x12, 0x21, 0x0a, 0x1c, 0x44, 0x45,
	0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4d,
	0x45, 0x52, 0x49, 0x43, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0xbd, 0x17, 0x12, 0x1f, 0x0a,
	0x1a, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4e, 0x55, 0x4d, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x4d, 0x49, 0x4e, 0x10, 0xbe, 0x17, 0x12, 0x1f,
	0x0a, 0x1a, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x4d, 0x41, 0x58, 0x10, 0xbf, 0x17, 0x42,
	0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x2f, 0x6c, 0x69, 0x62,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_steps_sp_steps_detective_proto_rawDescOnce sync.Once
	file_steps_sp_steps_detective_proto_rawDescData = file_steps_sp_steps_detective_proto_rawDesc
)

func file_steps_sp_steps_detective_proto_rawDescGZIP() []byte {
	file_steps_sp_steps_detective_proto_rawDescOnce.Do(func() {
		file_steps_sp_steps_detective_proto_rawDescData = protoimpl.X.CompressGZIP(file_steps_sp_steps_detective_proto_rawDescData)
	})
	return file_steps_sp_steps_detective_proto_rawDescData
}

var file_steps_sp_steps_detective_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_steps_sp_steps_detective_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_steps_sp_steps_detective_proto_goTypes = []interface{}{
	(DetectiveType)(0),    // 0: protos.steps.DetectiveType
	(*DetectiveStep)(nil), // 1: protos.steps.DetectiveStep
}
var file_steps_sp_steps_detective_proto_depIdxs = []int32{
	0, // 0: protos.steps.DetectiveStep.type:type_name -> protos.steps.DetectiveType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_steps_sp_steps_detective_proto_init() }
func file_steps_sp_steps_detective_proto_init() {
	if File_steps_sp_steps_detective_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_steps_sp_steps_detective_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectiveStep); i {
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
	file_steps_sp_steps_detective_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_steps_sp_steps_detective_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_steps_sp_steps_detective_proto_goTypes,
		DependencyIndexes: file_steps_sp_steps_detective_proto_depIdxs,
		EnumInfos:         file_steps_sp_steps_detective_proto_enumTypes,
		MessageInfos:      file_steps_sp_steps_detective_proto_msgTypes,
	}.Build()
	File_steps_sp_steps_detective_proto = out.File
	file_steps_sp_steps_detective_proto_rawDesc = nil
	file_steps_sp_steps_detective_proto_goTypes = nil
	file_steps_sp_steps_detective_proto_depIdxs = nil
}
