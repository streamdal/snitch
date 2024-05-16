// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: sp_bus.proto

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

// Type used by `server` for broadcasting events to other nodes
type BusEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// This _should_ contain request messages - each server can determine
	// how to interpret and handle the message.
	//
	// NOTE: The bus _should not_ be used for transmitting commands to SDKs. The
	// consumer in each SDK should receive a request and potentially craft a new
	// command to send to the appropriate SDK(s).
	//
	// Types that are assignable to Event:
	//
	//	*BusEvent_RegisterRequest
	//	*BusEvent_DeregisterRequest
	//	*BusEvent_CreatePipelineRequest
	//	*BusEvent_DeletePipelineRequest
	//	*BusEvent_UpdatePipelineRequest
	//	*BusEvent_PausePipelineRequest
	//	*BusEvent_ResumePipelineRequest
	//	*BusEvent_MetricsRequest
	//	*BusEvent_KvRequest
	//	*BusEvent_DeleteAudienceRequest
	//	*BusEvent_NewAudienceRequest
	//	*BusEvent_TailRequest
	//	*BusEvent_TailResponse
	//	*BusEvent_SetPipelinesRequest
	Event isBusEvent_Event `protobuf_oneof:"event"`
	// All gRPC metadata is stored in ctx; when request goes outside of gRPC
	// bounds, we will translate ctx metadata into this field.
	//
	// Example:
	//  1. Request comes into server via external gRPC to set new pipeline
	//  2. server has to send SetPipeline cmd to SDK via gRPC - it passes
	//     on original metadata in request.
	//  3. server has to broadcast SetPipeline cmd to other services via bus
	//  4. Since this is not a gRPC call, server translates ctx metadata to
	//     this field and includes it in the bus event.
	XMetadata map[string]string `protobuf:"bytes,1000,rep,name=_metadata,json=Metadata,proto3" json:"_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
}

func (x *BusEvent) Reset() {
	*x = BusEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_bus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusEvent) ProtoMessage() {}

func (x *BusEvent) ProtoReflect() protoreflect.Message {
	mi := &file_sp_bus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusEvent.ProtoReflect.Descriptor instead.
func (*BusEvent) Descriptor() ([]byte, []int) {
	return file_sp_bus_proto_rawDescGZIP(), []int{0}
}

func (x *BusEvent) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (m *BusEvent) GetEvent() isBusEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *BusEvent) GetRegisterRequest() *RegisterRequest {
	if x, ok := x.GetEvent().(*BusEvent_RegisterRequest); ok {
		return x.RegisterRequest
	}
	return nil
}

func (x *BusEvent) GetDeregisterRequest() *DeregisterRequest {
	if x, ok := x.GetEvent().(*BusEvent_DeregisterRequest); ok {
		return x.DeregisterRequest
	}
	return nil
}

func (x *BusEvent) GetCreatePipelineRequest() *CreatePipelineRequest {
	if x, ok := x.GetEvent().(*BusEvent_CreatePipelineRequest); ok {
		return x.CreatePipelineRequest
	}
	return nil
}

func (x *BusEvent) GetDeletePipelineRequest() *DeletePipelineRequest {
	if x, ok := x.GetEvent().(*BusEvent_DeletePipelineRequest); ok {
		return x.DeletePipelineRequest
	}
	return nil
}

func (x *BusEvent) GetUpdatePipelineRequest() *UpdatePipelineRequest {
	if x, ok := x.GetEvent().(*BusEvent_UpdatePipelineRequest); ok {
		return x.UpdatePipelineRequest
	}
	return nil
}

func (x *BusEvent) GetPausePipelineRequest() *PausePipelineRequest {
	if x, ok := x.GetEvent().(*BusEvent_PausePipelineRequest); ok {
		return x.PausePipelineRequest
	}
	return nil
}

func (x *BusEvent) GetResumePipelineRequest() *ResumePipelineRequest {
	if x, ok := x.GetEvent().(*BusEvent_ResumePipelineRequest); ok {
		return x.ResumePipelineRequest
	}
	return nil
}

func (x *BusEvent) GetMetricsRequest() *MetricsRequest {
	if x, ok := x.GetEvent().(*BusEvent_MetricsRequest); ok {
		return x.MetricsRequest
	}
	return nil
}

func (x *BusEvent) GetKvRequest() *KVRequest {
	if x, ok := x.GetEvent().(*BusEvent_KvRequest); ok {
		return x.KvRequest
	}
	return nil
}

func (x *BusEvent) GetDeleteAudienceRequest() *DeleteAudienceRequest {
	if x, ok := x.GetEvent().(*BusEvent_DeleteAudienceRequest); ok {
		return x.DeleteAudienceRequest
	}
	return nil
}

func (x *BusEvent) GetNewAudienceRequest() *NewAudienceRequest {
	if x, ok := x.GetEvent().(*BusEvent_NewAudienceRequest); ok {
		return x.NewAudienceRequest
	}
	return nil
}

func (x *BusEvent) GetTailRequest() *TailRequest {
	if x, ok := x.GetEvent().(*BusEvent_TailRequest); ok {
		return x.TailRequest
	}
	return nil
}

func (x *BusEvent) GetTailResponse() *TailResponse {
	if x, ok := x.GetEvent().(*BusEvent_TailResponse); ok {
		return x.TailResponse
	}
	return nil
}

func (x *BusEvent) GetSetPipelinesRequest() *SetPipelinesRequest {
	if x, ok := x.GetEvent().(*BusEvent_SetPipelinesRequest); ok {
		return x.SetPipelinesRequest
	}
	return nil
}

func (x *BusEvent) GetXMetadata() map[string]string {
	if x != nil {
		return x.XMetadata
	}
	return nil
}

type isBusEvent_Event interface {
	isBusEvent_Event()
}

type BusEvent_RegisterRequest struct {
	RegisterRequest *RegisterRequest `protobuf:"bytes,100,opt,name=register_request,json=registerRequest,proto3,oneof"`
}

type BusEvent_DeregisterRequest struct {
	DeregisterRequest *DeregisterRequest `protobuf:"bytes,101,opt,name=deregister_request,json=deregisterRequest,proto3,oneof"`
}

type BusEvent_CreatePipelineRequest struct {
	CreatePipelineRequest *CreatePipelineRequest `protobuf:"bytes,102,opt,name=create_pipeline_request,json=createPipelineRequest,proto3,oneof"`
}

type BusEvent_DeletePipelineRequest struct {
	DeletePipelineRequest *DeletePipelineRequest `protobuf:"bytes,103,opt,name=delete_pipeline_request,json=deletePipelineRequest,proto3,oneof"`
}

type BusEvent_UpdatePipelineRequest struct {
	UpdatePipelineRequest *UpdatePipelineRequest `protobuf:"bytes,104,opt,name=update_pipeline_request,json=updatePipelineRequest,proto3,oneof"`
}

type BusEvent_PausePipelineRequest struct {
	PausePipelineRequest *PausePipelineRequest `protobuf:"bytes,107,opt,name=pause_pipeline_request,json=pausePipelineRequest,proto3,oneof"`
}

type BusEvent_ResumePipelineRequest struct {
	ResumePipelineRequest *ResumePipelineRequest `protobuf:"bytes,108,opt,name=resume_pipeline_request,json=resumePipelineRequest,proto3,oneof"`
}

type BusEvent_MetricsRequest struct {
	MetricsRequest *MetricsRequest `protobuf:"bytes,109,opt,name=metrics_request,json=metricsRequest,proto3,oneof"`
}

type BusEvent_KvRequest struct {
	KvRequest *KVRequest `protobuf:"bytes,110,opt,name=kv_request,json=kvRequest,proto3,oneof"`
}

type BusEvent_DeleteAudienceRequest struct {
	DeleteAudienceRequest *DeleteAudienceRequest `protobuf:"bytes,111,opt,name=delete_audience_request,json=deleteAudienceRequest,proto3,oneof"`
}

type BusEvent_NewAudienceRequest struct {
	NewAudienceRequest *NewAudienceRequest `protobuf:"bytes,112,opt,name=new_audience_request,json=newAudienceRequest,proto3,oneof"`
}

type BusEvent_TailRequest struct {
	TailRequest *TailRequest `protobuf:"bytes,113,opt,name=tail_request,json=tailRequest,proto3,oneof"`
}

type BusEvent_TailResponse struct {
	TailResponse *TailResponse `protobuf:"bytes,114,opt,name=tail_response,json=tailResponse,proto3,oneof"`
}

type BusEvent_SetPipelinesRequest struct {
	SetPipelinesRequest *SetPipelinesRequest `protobuf:"bytes,115,opt,name=set_pipelines_request,json=setPipelinesRequest,proto3,oneof"`
}

func (*BusEvent_RegisterRequest) isBusEvent_Event() {}

func (*BusEvent_DeregisterRequest) isBusEvent_Event() {}

func (*BusEvent_CreatePipelineRequest) isBusEvent_Event() {}

func (*BusEvent_DeletePipelineRequest) isBusEvent_Event() {}

func (*BusEvent_UpdatePipelineRequest) isBusEvent_Event() {}

func (*BusEvent_PausePipelineRequest) isBusEvent_Event() {}

func (*BusEvent_ResumePipelineRequest) isBusEvent_Event() {}

func (*BusEvent_MetricsRequest) isBusEvent_Event() {}

func (*BusEvent_KvRequest) isBusEvent_Event() {}

func (*BusEvent_DeleteAudienceRequest) isBusEvent_Event() {}

func (*BusEvent_NewAudienceRequest) isBusEvent_Event() {}

func (*BusEvent_TailRequest) isBusEvent_Event() {}

func (*BusEvent_TailResponse) isBusEvent_Event() {}

func (*BusEvent_SetPipelinesRequest) isBusEvent_Event() {}

var File_sp_bus_proto protoreflect.FileDescriptor

var file_sp_bus_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x70, 0x5f, 0x62, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x0f, 0x73, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x70, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x70, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x73,
	0x70, 0x5f, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x09, 0x0a, 0x08, 0x42,
	0x75, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x44, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x12, 0x64, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x11,
	0x64, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x57, 0x0a, 0x17, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x66, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x17, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x15, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x17, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x68,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x15, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x16,
	0x70, 0x61, 0x75, 0x73, 0x65, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x75, 0x73, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x14, 0x70, 0x61,
	0x75, 0x73, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x57, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x6c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x15, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x6d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x0a, 0x6b, 0x76, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x6e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4b, 0x56, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x6b, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x57, 0x0a, 0x17, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x61, 0x75, 0x64,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x6f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x15, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x14, 0x6e,
	0x65, 0x77, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x12, 0x6e, 0x65, 0x77, 0x41, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x74,
	0x61, 0x69, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x71, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0d, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x72, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x51, 0x0a, 0x15, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x73, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x13, 0x73, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0xe8, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4a, 0x04, 0x08, 0x69, 0x10, 0x6a, 0x4a,
	0x04, 0x08, 0x6a, 0x10, 0x6b, 0x42, 0x50, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0xea, 0x02, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x3a,
	0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sp_bus_proto_rawDescOnce sync.Once
	file_sp_bus_proto_rawDescData = file_sp_bus_proto_rawDesc
)

func file_sp_bus_proto_rawDescGZIP() []byte {
	file_sp_bus_proto_rawDescOnce.Do(func() {
		file_sp_bus_proto_rawDescData = protoimpl.X.CompressGZIP(file_sp_bus_proto_rawDescData)
	})
	return file_sp_bus_proto_rawDescData
}

var file_sp_bus_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sp_bus_proto_goTypes = []interface{}{
	(*BusEvent)(nil),              // 0: protos.BusEvent
	nil,                           // 1: protos.BusEvent.MetadataEntry
	(*RegisterRequest)(nil),       // 2: protos.RegisterRequest
	(*DeregisterRequest)(nil),     // 3: protos.DeregisterRequest
	(*CreatePipelineRequest)(nil), // 4: protos.CreatePipelineRequest
	(*DeletePipelineRequest)(nil), // 5: protos.DeletePipelineRequest
	(*UpdatePipelineRequest)(nil), // 6: protos.UpdatePipelineRequest
	(*PausePipelineRequest)(nil),  // 7: protos.PausePipelineRequest
	(*ResumePipelineRequest)(nil), // 8: protos.ResumePipelineRequest
	(*MetricsRequest)(nil),        // 9: protos.MetricsRequest
	(*KVRequest)(nil),             // 10: protos.KVRequest
	(*DeleteAudienceRequest)(nil), // 11: protos.DeleteAudienceRequest
	(*NewAudienceRequest)(nil),    // 12: protos.NewAudienceRequest
	(*TailRequest)(nil),           // 13: protos.TailRequest
	(*TailResponse)(nil),          // 14: protos.TailResponse
	(*SetPipelinesRequest)(nil),   // 15: protos.SetPipelinesRequest
}
var file_sp_bus_proto_depIdxs = []int32{
	2,  // 0: protos.BusEvent.register_request:type_name -> protos.RegisterRequest
	3,  // 1: protos.BusEvent.deregister_request:type_name -> protos.DeregisterRequest
	4,  // 2: protos.BusEvent.create_pipeline_request:type_name -> protos.CreatePipelineRequest
	5,  // 3: protos.BusEvent.delete_pipeline_request:type_name -> protos.DeletePipelineRequest
	6,  // 4: protos.BusEvent.update_pipeline_request:type_name -> protos.UpdatePipelineRequest
	7,  // 5: protos.BusEvent.pause_pipeline_request:type_name -> protos.PausePipelineRequest
	8,  // 6: protos.BusEvent.resume_pipeline_request:type_name -> protos.ResumePipelineRequest
	9,  // 7: protos.BusEvent.metrics_request:type_name -> protos.MetricsRequest
	10, // 8: protos.BusEvent.kv_request:type_name -> protos.KVRequest
	11, // 9: protos.BusEvent.delete_audience_request:type_name -> protos.DeleteAudienceRequest
	12, // 10: protos.BusEvent.new_audience_request:type_name -> protos.NewAudienceRequest
	13, // 11: protos.BusEvent.tail_request:type_name -> protos.TailRequest
	14, // 12: protos.BusEvent.tail_response:type_name -> protos.TailResponse
	15, // 13: protos.BusEvent.set_pipelines_request:type_name -> protos.SetPipelinesRequest
	1,  // 14: protos.BusEvent._metadata:type_name -> protos.BusEvent.MetadataEntry
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_sp_bus_proto_init() }
func file_sp_bus_proto_init() {
	if File_sp_bus_proto != nil {
		return
	}
	file_sp_common_proto_init()
	file_sp_external_proto_init()
	file_sp_internal_proto_init()
	file_sp_kv_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sp_bus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusEvent); i {
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
	file_sp_bus_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BusEvent_RegisterRequest)(nil),
		(*BusEvent_DeregisterRequest)(nil),
		(*BusEvent_CreatePipelineRequest)(nil),
		(*BusEvent_DeletePipelineRequest)(nil),
		(*BusEvent_UpdatePipelineRequest)(nil),
		(*BusEvent_PausePipelineRequest)(nil),
		(*BusEvent_ResumePipelineRequest)(nil),
		(*BusEvent_MetricsRequest)(nil),
		(*BusEvent_KvRequest)(nil),
		(*BusEvent_DeleteAudienceRequest)(nil),
		(*BusEvent_NewAudienceRequest)(nil),
		(*BusEvent_TailRequest)(nil),
		(*BusEvent_TailResponse)(nil),
		(*BusEvent_SetPipelinesRequest)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sp_bus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sp_bus_proto_goTypes,
		DependencyIndexes: file_sp_bus_proto_depIdxs,
		MessageInfos:      file_sp_bus_proto_msgTypes,
	}.Build()
	File_sp_bus_proto = out.File
	file_sp_bus_proto_rawDesc = nil
	file_sp_bus_proto_goTypes = nil
	file_sp_bus_proto_depIdxs = nil
}
