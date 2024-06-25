// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: sp_command.proto

package protos

import (
	shared "github.com/streamdal/streamdal/libs/protos/build/go/protos/shared"
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

// Command is used by streamdal server for sending commands to SDKs
type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Who is this command intended for?
	// NOTE: Some commands (such as KeepAliveCommand, KVCommand) do NOT use audience and will ignore it
	Audience *Audience `protobuf:"bytes,1,opt,name=audience,proto3" json:"audience,omitempty"`
	// Types that are assignable to Command:
	//
	//	*Command_SetPipelines
	//	*Command_KeepAlive
	//	*Command_Kv
	//	*Command_Tail
	//	*Command_Delete
	Command isCommand_Command `protobuf_oneof:"command"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_sp_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_sp_command_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetAudience() *Audience {
	if x != nil {
		return x.Audience
	}
	return nil
}

func (m *Command) GetCommand() isCommand_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (x *Command) GetSetPipelines() *SetPipelinesCommand {
	if x, ok := x.GetCommand().(*Command_SetPipelines); ok {
		return x.SetPipelines
	}
	return nil
}

func (x *Command) GetKeepAlive() *KeepAliveCommand {
	if x, ok := x.GetCommand().(*Command_KeepAlive); ok {
		return x.KeepAlive
	}
	return nil
}

func (x *Command) GetKv() *KVCommand {
	if x, ok := x.GetCommand().(*Command_Kv); ok {
		return x.Kv
	}
	return nil
}

func (x *Command) GetTail() *TailCommand {
	if x, ok := x.GetCommand().(*Command_Tail); ok {
		return x.Tail
	}
	return nil
}

func (x *Command) GetDelete() *DeleteAudienceCommand {
	if x, ok := x.GetCommand().(*Command_Delete); ok {
		return x.Delete
	}
	return nil
}

type isCommand_Command interface {
	isCommand_Command()
}

type Command_SetPipelines struct {
	// Emitted by server when a user makes a pause, resume, delete or update
	// pipeline and set pipelines external grpc API call.
	// NOTE: This was introduced during ordered pipeline updates.
	SetPipelines *SetPipelinesCommand `protobuf:"bytes,100,opt,name=set_pipelines,json=setPipelines,proto3,oneof"`
}

type Command_KeepAlive struct {
	// Server sends this periodically to SDKs to keep the connection alive
	KeepAlive *KeepAliveCommand `protobuf:"bytes,101,opt,name=keep_alive,json=keepAlive,proto3,oneof"`
}

type Command_Kv struct {
	// Server will emit this when a user makes changes to the KV store
	// via the KV HTTP API.
	Kv *KVCommand `protobuf:"bytes,102,opt,name=kv,proto3,oneof"`
}

type Command_Tail struct {
	// Emitted by server when a user makes a Tail() call
	// Consumed by all server instances and by SDKs
	Tail *TailCommand `protobuf:"bytes,103,opt,name=tail,proto3,oneof"`
}

type Command_Delete struct {
	Delete *DeleteAudienceCommand `protobuf:"bytes,104,opt,name=delete,proto3,oneof"`
}

func (*Command_SetPipelines) isCommand_Command() {}

func (*Command_KeepAlive) isCommand_Command() {}

func (*Command_Kv) isCommand_Command() {}

func (*Command_Tail) isCommand_Command() {}

func (*Command_Delete) isCommand_Command() {}

type SetPipelinesCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pipelines []*Pipeline `protobuf:"bytes,1,rep,name=pipelines,proto3" json:"pipelines,omitempty"`
	// ID = wasm ID
	WasmModules map[string]*shared.WasmModule `protobuf:"bytes,2,rep,name=wasm_modules,json=wasmModules,proto3" json:"wasm_modules,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SetPipelinesCommand) Reset() {
	*x = SetPipelinesCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPipelinesCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPipelinesCommand) ProtoMessage() {}

func (x *SetPipelinesCommand) ProtoReflect() protoreflect.Message {
	mi := &file_sp_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPipelinesCommand.ProtoReflect.Descriptor instead.
func (*SetPipelinesCommand) Descriptor() ([]byte, []int) {
	return file_sp_command_proto_rawDescGZIP(), []int{1}
}

func (x *SetPipelinesCommand) GetPipelines() []*Pipeline {
	if x != nil {
		return x.Pipelines
	}
	return nil
}

func (x *SetPipelinesCommand) GetWasmModules() map[string]*shared.WasmModule {
	if x != nil {
		return x.WasmModules
	}
	return nil
}

type KeepAliveCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KeepAliveCommand) Reset() {
	*x = KeepAliveCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeepAliveCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeepAliveCommand) ProtoMessage() {}

func (x *KeepAliveCommand) ProtoReflect() protoreflect.Message {
	mi := &file_sp_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeepAliveCommand.ProtoReflect.Descriptor instead.
func (*KeepAliveCommand) Descriptor() ([]byte, []int) {
	return file_sp_command_proto_rawDescGZIP(), []int{2}
}

// Sent by server on Register channel(s) to live SDKs
type KVCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instructions []*KVInstruction `protobuf:"bytes,1,rep,name=instructions,proto3" json:"instructions,omitempty"`
	// Create & Update specific setting that will cause the Create or Update to
	// work as an upsert.
	Overwrite bool `protobuf:"varint,2,opt,name=overwrite,proto3" json:"overwrite,omitempty"`
}

func (x *KVCommand) Reset() {
	*x = KVCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVCommand) ProtoMessage() {}

func (x *KVCommand) ProtoReflect() protoreflect.Message {
	mi := &file_sp_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVCommand.ProtoReflect.Descriptor instead.
func (*KVCommand) Descriptor() ([]byte, []int) {
	return file_sp_command_proto_rawDescGZIP(), []int{3}
}

func (x *KVCommand) GetInstructions() []*KVInstruction {
	if x != nil {
		return x.Instructions
	}
	return nil
}

func (x *KVCommand) GetOverwrite() bool {
	if x != nil {
		return x.Overwrite
	}
	return false
}

type TailCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *TailRequest `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *TailCommand) Reset() {
	*x = TailCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TailCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TailCommand) ProtoMessage() {}

func (x *TailCommand) ProtoReflect() protoreflect.Message {
	mi := &file_sp_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TailCommand.ProtoReflect.Descriptor instead.
func (*TailCommand) Descriptor() ([]byte, []int) {
	return file_sp_command_proto_rawDescGZIP(), []int{4}
}

func (x *TailCommand) GetRequest() *TailRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type DeleteAudienceCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Audience *Audience `protobuf:"bytes,1,opt,name=audience,proto3" json:"audience,omitempty"`
}

func (x *DeleteAudienceCommand) Reset() {
	*x = DeleteAudienceCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_command_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAudienceCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAudienceCommand) ProtoMessage() {}

func (x *DeleteAudienceCommand) ProtoReflect() protoreflect.Message {
	mi := &file_sp_command_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAudienceCommand.ProtoReflect.Descriptor instead.
func (*DeleteAudienceCommand) Descriptor() ([]byte, []int) {
	return file_sp_command_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAudienceCommand) GetAudience() *Audience {
	if x != nil {
		return x.Audience
	}
	return nil
}

var File_sp_command_proto protoreflect.FileDescriptor

var file_sp_command_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x16, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2f, 0x73, 0x70, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x73, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x73, 0x70, 0x5f, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x73, 0x70, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xca, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x2c, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x42, 0x0a,
	0x0d, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x65,
	0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x73, 0x12, 0x39, 0x0a, 0x0a, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4b,
	0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48,
	0x00, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x23, 0x0a, 0x02,
	0x6b, 0x76, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x4b, 0x56, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x02, 0x6b,
	0x76, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x04, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x37, 0x0a, 0x06,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x68, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x06, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x22, 0xf1, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2e, 0x0a, 0x09, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x09, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x0c, 0x77, 0x61, 0x73, 0x6d,
	0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x57, 0x61, 0x73, 0x6d,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x77, 0x61,
	0x73, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x59, 0x0a, 0x10, 0x57, 0x61, 0x73,
	0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x57,
	0x61, 0x73, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x12, 0x0a, 0x10, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x64, 0x0a, 0x09, 0x4b, 0x56, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x39, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4b, 0x56, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x22, 0x3c,
	0x0a, 0x0b, 0x54, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2d, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x42, 0x50, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0xea, 0x02, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x3a, 0x3a, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sp_command_proto_rawDescOnce sync.Once
	file_sp_command_proto_rawDescData = file_sp_command_proto_rawDesc
)

func file_sp_command_proto_rawDescGZIP() []byte {
	file_sp_command_proto_rawDescOnce.Do(func() {
		file_sp_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_sp_command_proto_rawDescData)
	})
	return file_sp_command_proto_rawDescData
}

var file_sp_command_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_sp_command_proto_goTypes = []interface{}{
	(*Command)(nil),               // 0: protos.Command
	(*SetPipelinesCommand)(nil),   // 1: protos.SetPipelinesCommand
	(*KeepAliveCommand)(nil),      // 2: protos.KeepAliveCommand
	(*KVCommand)(nil),             // 3: protos.KVCommand
	(*TailCommand)(nil),           // 4: protos.TailCommand
	(*DeleteAudienceCommand)(nil), // 5: protos.DeleteAudienceCommand
	nil,                           // 6: protos.SetPipelinesCommand.WasmModulesEntry
	(*Audience)(nil),              // 7: protos.Audience
	(*Pipeline)(nil),              // 8: protos.Pipeline
	(*KVInstruction)(nil),         // 9: protos.KVInstruction
	(*TailRequest)(nil),           // 10: protos.TailRequest
	(*shared.WasmModule)(nil),     // 11: protos.shared.WasmModule
}
var file_sp_command_proto_depIdxs = []int32{
	7,  // 0: protos.Command.audience:type_name -> protos.Audience
	1,  // 1: protos.Command.set_pipelines:type_name -> protos.SetPipelinesCommand
	2,  // 2: protos.Command.keep_alive:type_name -> protos.KeepAliveCommand
	3,  // 3: protos.Command.kv:type_name -> protos.KVCommand
	4,  // 4: protos.Command.tail:type_name -> protos.TailCommand
	5,  // 5: protos.Command.delete:type_name -> protos.DeleteAudienceCommand
	8,  // 6: protos.SetPipelinesCommand.pipelines:type_name -> protos.Pipeline
	6,  // 7: protos.SetPipelinesCommand.wasm_modules:type_name -> protos.SetPipelinesCommand.WasmModulesEntry
	9,  // 8: protos.KVCommand.instructions:type_name -> protos.KVInstruction
	10, // 9: protos.TailCommand.request:type_name -> protos.TailRequest
	7,  // 10: protos.DeleteAudienceCommand.audience:type_name -> protos.Audience
	11, // 11: protos.SetPipelinesCommand.WasmModulesEntry.value:type_name -> protos.shared.WasmModule
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_sp_command_proto_init() }
func file_sp_command_proto_init() {
	if File_sp_command_proto != nil {
		return
	}
	file_sp_common_proto_init()
	file_sp_kv_proto_init()
	file_sp_pipeline_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sp_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_sp_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPipelinesCommand); i {
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
		file_sp_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeepAliveCommand); i {
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
		file_sp_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVCommand); i {
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
		file_sp_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TailCommand); i {
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
		file_sp_command_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAudienceCommand); i {
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
	file_sp_command_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Command_SetPipelines)(nil),
		(*Command_KeepAlive)(nil),
		(*Command_Kv)(nil),
		(*Command_Tail)(nil),
		(*Command_Delete)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sp_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sp_command_proto_goTypes,
		DependencyIndexes: file_sp_command_proto_depIdxs,
		MessageInfos:      file_sp_command_proto_msgTypes,
	}.Build()
	File_sp_command_proto = out.File
	file_sp_command_proto_rawDesc = nil
	file_sp_command_proto_goTypes = nil
	file_sp_command_proto_depIdxs = nil
}
