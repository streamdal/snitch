// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: sp_sdk.proto

package protos

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

type ExecStatus int32

const (
	// Unset status. This should never be returned by the SDK. If it does, it is
	// probably a bug (and you should file an issue)
	ExecStatus_EXEC_STATUS_UNSET ExecStatus = 0
	// Indicates that the step execution evaluated to "true"
	ExecStatus_EXEC_STATUS_TRUE ExecStatus = 1
	// Indicates that the step execution evaluated to "false"
	ExecStatus_EXEC_STATUS_FALSE ExecStatus = 2
	// Indicates that the SDK encountered an error while trying to process the
	// request. Example error cases: SDK can't find the appropriate Wasm module,
	// Wasm function cannot alloc or dealloc memory, etc.
	ExecStatus_EXEC_STATUS_ERROR ExecStatus = 3
)

// Enum value maps for ExecStatus.
var (
	ExecStatus_name = map[int32]string{
		0: "EXEC_STATUS_UNSET",
		1: "EXEC_STATUS_TRUE",
		2: "EXEC_STATUS_FALSE",
		3: "EXEC_STATUS_ERROR",
	}
	ExecStatus_value = map[string]int32{
		"EXEC_STATUS_UNSET": 0,
		"EXEC_STATUS_TRUE":  1,
		"EXEC_STATUS_FALSE": 2,
		"EXEC_STATUS_ERROR": 3,
	}
)

func (x ExecStatus) Enum() *ExecStatus {
	p := new(ExecStatus)
	*p = x
	return p
}

func (x ExecStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExecStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_sp_sdk_proto_enumTypes[0].Descriptor()
}

func (ExecStatus) Type() protoreflect.EnumType {
	return &file_sp_sdk_proto_enumTypes[0]
}

func (x ExecStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExecStatus.Descriptor instead.
func (ExecStatus) EnumDescriptor() ([]byte, []int) {
	return file_sp_sdk_proto_rawDescGZIP(), []int{0}
}

// SDKErrorMode is used to alter the error behavior of a shim library
// instrumented with the Streamdal SDK at runtime.
//
// NOTE: This structure is usually used when the SDK is used via a shim/wrapper
// library where you have less control over SDK behavior. Read more about shims
// here: https://docs.streamdal.com/en/core-components/libraries-shims/
//
// protolint:disable ENUM_FIELD_NAMES_PREFIX
type SDKErrorMode int32

const (
	// This instructs the shim to IGNORE any non-recoverable errors that the SDK
	// might run into. If the SDK runs into an error, the shim will NOT pass the
	// error back to the user - it will instead return the whatever the upstream
	// library would normally return to the user.
	//
	// *** This is the default behavior ***
	//
	// Example with Redis Shim
	// ------------------------
	// Under normal conditions, a Redis shim would work in the following way when
	// user is performing a read operation:
	//
	// 1. The shim would call the upstream Redis library to perform the read operation
	// 2. Upstream library returns results to the shim
	// 3. Shim passes the result to the integrated Streamdal SDK for processing
	// 4. SDK returns (potentially) modified data to the shim
	// 5. Shim returns the modified data to the user
	//
	// This setting tells the shim that IF it runs into a non-recoverable error
	// while calling the SDK (step 3), it will side-step steps 4 and 5 and instead
	// return the _original_ payload (read during step 1) to the user.
	SDKErrorMode_SDK_ERROR_MODE_UNSET SDKErrorMode = 0
	// This instructs the shim to ABORT execution if the SDK runs into any
	// non-recoverable errors. Upon aborting, the shim will return the error that
	// the SDK ran into and the error will be passed all the way back to the user.
	SDKErrorMode_SDK_ERROR_MODE_STRICT SDKErrorMode = 1
)

// Enum value maps for SDKErrorMode.
var (
	SDKErrorMode_name = map[int32]string{
		0: "SDK_ERROR_MODE_UNSET",
		1: "SDK_ERROR_MODE_STRICT",
	}
	SDKErrorMode_value = map[string]int32{
		"SDK_ERROR_MODE_UNSET":  0,
		"SDK_ERROR_MODE_STRICT": 1,
	}
)

func (x SDKErrorMode) Enum() *SDKErrorMode {
	p := new(SDKErrorMode)
	*p = x
	return p
}

func (x SDKErrorMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SDKErrorMode) Descriptor() protoreflect.EnumDescriptor {
	return file_sp_sdk_proto_enumTypes[1].Descriptor()
}

func (SDKErrorMode) Type() protoreflect.EnumType {
	return &file_sp_sdk_proto_enumTypes[1]
}

func (x SDKErrorMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SDKErrorMode.Descriptor instead.
func (SDKErrorMode) EnumDescriptor() ([]byte, []int) {
	return file_sp_sdk_proto_rawDescGZIP(), []int{1}
}

// Common return response used by all SDKs
type SDKResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains (potentially) modified input data
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Execution status of the last step
	Status ExecStatus `protobuf:"varint,2,opt,name=status,proto3,enum=protos.ExecStatus" json:"status,omitempty"`
	// Optional message accompanying the exec status for the last step
	StatusMessage *string `protobuf:"bytes,3,opt,name=status_message,json=statusMessage,proto3,oneof" json:"status_message,omitempty"`
	// An array of pipelines that the SDK executed and the status of each step
	PipelineStatus []*PipelineStatus `protobuf:"bytes,4,rep,name=pipeline_status,json=pipelineStatus,proto3" json:"pipeline_status,omitempty"`
	// Includes any metadata that the step(s) may want to pass back to the user.
	//
	// NOTE: Metadata is aggregated across all steps in the pipeline, so if two
	// steps both set a key "foo" to different values, the value of "foo" in the
	// response will be the value set by the last step in the pipeline.
	//
	// To learn more about "metadata", see SDK Spec V2 doc "Pipeline Step & Error
	// Behavior" section.
	Metadata map[string]string `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SDKResponse) Reset() {
	*x = SDKResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_sdk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SDKResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SDKResponse) ProtoMessage() {}

func (x *SDKResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sp_sdk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SDKResponse.ProtoReflect.Descriptor instead.
func (*SDKResponse) Descriptor() ([]byte, []int) {
	return file_sp_sdk_proto_rawDescGZIP(), []int{0}
}

func (x *SDKResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SDKResponse) GetStatus() ExecStatus {
	if x != nil {
		return x.Status
	}
	return ExecStatus_EXEC_STATUS_UNSET
}

func (x *SDKResponse) GetStatusMessage() string {
	if x != nil && x.StatusMessage != nil {
		return *x.StatusMessage
	}
	return ""
}

func (x *SDKResponse) GetPipelineStatus() []*PipelineStatus {
	if x != nil {
		return x.PipelineStatus
	}
	return nil
}

func (x *SDKResponse) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type PipelineStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the pipeline
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the pipeline
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The status of each step in the pipeline
	StepStatus []*StepStatus `protobuf:"bytes,3,rep,name=step_status,json=stepStatus,proto3" json:"step_status,omitempty"`
}

func (x *PipelineStatus) Reset() {
	*x = PipelineStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_sdk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineStatus) ProtoMessage() {}

func (x *PipelineStatus) ProtoReflect() protoreflect.Message {
	mi := &file_sp_sdk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineStatus.ProtoReflect.Descriptor instead.
func (*PipelineStatus) Descriptor() ([]byte, []int) {
	return file_sp_sdk_proto_rawDescGZIP(), []int{1}
}

func (x *PipelineStatus) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PipelineStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PipelineStatus) GetStepStatus() []*StepStatus {
	if x != nil {
		return x.StepStatus
	}
	return nil
}

type StepStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the step
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Execution outcome status of the step
	Status ExecStatus `protobuf:"varint,2,opt,name=status,proto3,enum=protos.ExecStatus" json:"status,omitempty"`
	// Optional message accompanying the exec status
	StatusMessage *string `protobuf:"bytes,3,opt,name=status_message,json=statusMessage,proto3,oneof" json:"status_message,omitempty"`
	// Indicates if current or all future pipelines were aborted.
	//
	// IMPORTANT: The SDK running into an error does not automatically abort
	// current or all future pipelines - the user must define the abort conditions
	// for "on_error".
	AbortCondition AbortCondition `protobuf:"varint,4,opt,name=abort_condition,json=abortCondition,proto3,enum=protos.AbortCondition" json:"abort_condition,omitempty"`
}

func (x *StepStatus) Reset() {
	*x = StepStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_sdk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepStatus) ProtoMessage() {}

func (x *StepStatus) ProtoReflect() protoreflect.Message {
	mi := &file_sp_sdk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepStatus.ProtoReflect.Descriptor instead.
func (*StepStatus) Descriptor() ([]byte, []int) {
	return file_sp_sdk_proto_rawDescGZIP(), []int{2}
}

func (x *StepStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StepStatus) GetStatus() ExecStatus {
	if x != nil {
		return x.Status
	}
	return ExecStatus_EXEC_STATUS_UNSET
}

func (x *StepStatus) GetStatusMessage() string {
	if x != nil && x.StatusMessage != nil {
		return *x.StatusMessage
	}
	return ""
}

func (x *StepStatus) GetAbortCondition() AbortCondition {
	if x != nil {
		return x.AbortCondition
	}
	return AbortCondition_ABORT_CONDITION_UNSET
}

// SDKStartupConfig is the configuration structure that is used in Streamdal SDKs
// to configure the client at startup. Some SDKs may expose additional config
// options aside from these baseline options.
type SDKStartupConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL for the Streamdal server gRPC API. Example: "streamdal-server-address:8082"
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Auth token used to authenticate with the Streamdal server (NOTE: should be
	// the same as the token used for running the Streamdal server).
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// Service name used for identifying the SDK client in the Streamdal server and console
	ServiceName string `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// How long to wait for a pipeline execution to complete before timing out
	PipelineTimeoutSeconds *int32 `protobuf:"varint,4,opt,name=pipeline_timeout_seconds,json=pipelineTimeoutSeconds,proto3,oneof" json:"pipeline_timeout_seconds,omitempty"`
	// How long to wait for a step execution to complete before timing out
	StepTimeoutSeconds *int32 `protobuf:"varint,5,opt,name=step_timeout_seconds,json=stepTimeoutSeconds,proto3,oneof" json:"step_timeout_seconds,omitempty"`
	// Tells the SDK how to behave when it runs into an error. This setting has
	// no effect if the SDK is NOT used within a shim/wrapper library. Read more
	// about shims here: https://docs.streamdal.com/en/core-components/libraries-shims/
	ErrorMode *SDKErrorMode `protobuf:"varint,6,opt,name=error_mode,json=errorMode,proto3,enum=protos.SDKErrorMode,oneof" json:"error_mode,omitempty"`
}

func (x *SDKStartupConfig) Reset() {
	*x = SDKStartupConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_sdk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SDKStartupConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SDKStartupConfig) ProtoMessage() {}

func (x *SDKStartupConfig) ProtoReflect() protoreflect.Message {
	mi := &file_sp_sdk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SDKStartupConfig.ProtoReflect.Descriptor instead.
func (*SDKStartupConfig) Descriptor() ([]byte, []int) {
	return file_sp_sdk_proto_rawDescGZIP(), []int{3}
}

func (x *SDKStartupConfig) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SDKStartupConfig) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SDKStartupConfig) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *SDKStartupConfig) GetPipelineTimeoutSeconds() int32 {
	if x != nil && x.PipelineTimeoutSeconds != nil {
		return *x.PipelineTimeoutSeconds
	}
	return 0
}

func (x *SDKStartupConfig) GetStepTimeoutSeconds() int32 {
	if x != nil && x.StepTimeoutSeconds != nil {
		return *x.StepTimeoutSeconds
	}
	return 0
}

func (x *SDKStartupConfig) GetErrorMode() SDKErrorMode {
	if x != nil && x.ErrorMode != nil {
		return *x.ErrorMode
	}
	return SDKErrorMode_SDK_ERROR_MODE_UNSET
}

// SDKRuntimeConfig is the configuration structure that is used in SDKs to
// configure how the SDK behaves at runtime. It is most often exposed as an
// optional parameter that you can pass to an upstream library's read or write
// operation. Ie. kafkaProducer.Write(data, &streamdal.SDKRuntimeConfig{...})
//
// NOTE: This structure is usually used when the SDK is used via a shim/wrapper
// library where you have less control over SDK behavior. Read more about shims
// here: https://docs.streamdal.com/en/core-components/libraries-shims/
type SDKRuntimeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies how the shim should behave if it runs into any errors when calling the SDK
	ErrorMode *SDKErrorMode `protobuf:"varint,1,opt,name=error_mode,json=errorMode,proto3,enum=protos.SDKErrorMode,oneof" json:"error_mode,omitempty"`
	// Audience that will be used by shim when calling SDK.Process()
	Audience *Audience `protobuf:"bytes,2,opt,name=audience,proto3,oneof" json:"audience,omitempty"`
}

func (x *SDKRuntimeConfig) Reset() {
	*x = SDKRuntimeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_sdk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SDKRuntimeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SDKRuntimeConfig) ProtoMessage() {}

func (x *SDKRuntimeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_sp_sdk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SDKRuntimeConfig.ProtoReflect.Descriptor instead.
func (*SDKRuntimeConfig) Descriptor() ([]byte, []int) {
	return file_sp_sdk_proto_rawDescGZIP(), []int{4}
}

func (x *SDKRuntimeConfig) GetErrorMode() SDKErrorMode {
	if x != nil && x.ErrorMode != nil {
		return *x.ErrorMode
	}
	return SDKErrorMode_SDK_ERROR_MODE_UNSET
}

func (x *SDKRuntimeConfig) GetAudience() *Audience {
	if x != nil {
		return x.Audience
	}
	return nil
}

var File_sp_sdk_proto protoreflect.FileDescriptor

var file_sp_sdk_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x70, 0x5f, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x0f, 0x73, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x70, 0x5f, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x02, 0x0a, 0x0b, 0x53,
	0x44, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x0e, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x0f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x53, 0x44, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x69, 0x0a, 0x0e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x0b,
	0x73, 0x74, 0x65, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x73, 0x74, 0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0xcc, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x78,
	0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x2a, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x0f,
	0x61, 0x62, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41,
	0x62, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x61,
	0x62, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x11, 0x0a,
	0x0f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0xd2, 0x02, 0x0a, 0x10, 0x53, 0x44, 0x4b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x3d, 0x0a, 0x18, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x16, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x35, 0x0a, 0x14, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52,
	0x12, 0x73, 0x74, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x53, 0x44, 0x4b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x6f, 0x64, 0x65,
	0x48, 0x02, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x1b, 0x0a, 0x19, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x42, 0x17, 0x0a,
	0x15, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x9b, 0x01, 0x0a, 0x10, 0x53, 0x44, 0x4b, 0x52, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x38, 0x0a, 0x0a, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x44, 0x4b, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x6f, 0x64,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x48, 0x01, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x2a, 0x67, 0x0a, 0x0a, 0x45, 0x78, 0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x58, 0x45, 0x43, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x58, 0x45, 0x43,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x52, 0x55, 0x45, 0x10, 0x01, 0x12, 0x15,
	0x0a, 0x11, 0x45, 0x58, 0x45, 0x43, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41,
	0x4c, 0x53, 0x45, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x58, 0x45, 0x43, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x2a, 0x43, 0x0a, 0x0c,
	0x53, 0x44, 0x4b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x14,
	0x53, 0x44, 0x4b, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x44, 0x4b, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x10,
	0x01, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x64, 0x61, 0x6c, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sp_sdk_proto_rawDescOnce sync.Once
	file_sp_sdk_proto_rawDescData = file_sp_sdk_proto_rawDesc
)

func file_sp_sdk_proto_rawDescGZIP() []byte {
	file_sp_sdk_proto_rawDescOnce.Do(func() {
		file_sp_sdk_proto_rawDescData = protoimpl.X.CompressGZIP(file_sp_sdk_proto_rawDescData)
	})
	return file_sp_sdk_proto_rawDescData
}

var file_sp_sdk_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_sp_sdk_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sp_sdk_proto_goTypes = []interface{}{
	(ExecStatus)(0),          // 0: protos.ExecStatus
	(SDKErrorMode)(0),        // 1: protos.SDKErrorMode
	(*SDKResponse)(nil),      // 2: protos.SDKResponse
	(*PipelineStatus)(nil),   // 3: protos.PipelineStatus
	(*StepStatus)(nil),       // 4: protos.StepStatus
	(*SDKStartupConfig)(nil), // 5: protos.SDKStartupConfig
	(*SDKRuntimeConfig)(nil), // 6: protos.SDKRuntimeConfig
	nil,                      // 7: protos.SDKResponse.MetadataEntry
	(AbortCondition)(0),      // 8: protos.AbortCondition
	(*Audience)(nil),         // 9: protos.Audience
}
var file_sp_sdk_proto_depIdxs = []int32{
	0, // 0: protos.SDKResponse.status:type_name -> protos.ExecStatus
	3, // 1: protos.SDKResponse.pipeline_status:type_name -> protos.PipelineStatus
	7, // 2: protos.SDKResponse.metadata:type_name -> protos.SDKResponse.MetadataEntry
	4, // 3: protos.PipelineStatus.step_status:type_name -> protos.StepStatus
	0, // 4: protos.StepStatus.status:type_name -> protos.ExecStatus
	8, // 5: protos.StepStatus.abort_condition:type_name -> protos.AbortCondition
	1, // 6: protos.SDKStartupConfig.error_mode:type_name -> protos.SDKErrorMode
	1, // 7: protos.SDKRuntimeConfig.error_mode:type_name -> protos.SDKErrorMode
	9, // 8: protos.SDKRuntimeConfig.audience:type_name -> protos.Audience
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_sp_sdk_proto_init() }
func file_sp_sdk_proto_init() {
	if File_sp_sdk_proto != nil {
		return
	}
	file_sp_common_proto_init()
	file_sp_pipeline_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sp_sdk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SDKResponse); i {
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
		file_sp_sdk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineStatus); i {
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
		file_sp_sdk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepStatus); i {
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
		file_sp_sdk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SDKStartupConfig); i {
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
		file_sp_sdk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SDKRuntimeConfig); i {
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
	file_sp_sdk_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_sp_sdk_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_sp_sdk_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_sp_sdk_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sp_sdk_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sp_sdk_proto_goTypes,
		DependencyIndexes: file_sp_sdk_proto_depIdxs,
		EnumInfos:         file_sp_sdk_proto_enumTypes,
		MessageInfos:      file_sp_sdk_proto_msgTypes,
	}.Build()
	File_sp_sdk_proto = out.File
	file_sp_sdk_proto_rawDesc = nil
	file_sp_sdk_proto_goTypes = nil
	file_sp_sdk_proto_depIdxs = nil
}
