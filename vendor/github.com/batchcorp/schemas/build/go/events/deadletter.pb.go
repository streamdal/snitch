// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: deadletter.proto

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

type DeadLetterState int32

const (
	DeadLetterState_DEAD_LETTER_STATE_UNSET DeadLetterState = 0
	DeadLetterState_DEAD_LETTER_STATE_OK    DeadLetterState = 1
	DeadLetterState_DEAD_LETTER_STATE_ERROR DeadLetterState = 2
)

// Enum value maps for DeadLetterState.
var (
	DeadLetterState_name = map[int32]string{
		0: "DEAD_LETTER_STATE_UNSET",
		1: "DEAD_LETTER_STATE_OK",
		2: "DEAD_LETTER_STATE_ERROR",
	}
	DeadLetterState_value = map[string]int32{
		"DEAD_LETTER_STATE_UNSET": 0,
		"DEAD_LETTER_STATE_OK":    1,
		"DEAD_LETTER_STATE_ERROR": 2,
	}
)

func (x DeadLetterState) Enum() *DeadLetterState {
	p := new(DeadLetterState)
	*p = x
	return p
}

func (x DeadLetterState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeadLetterState) Descriptor() protoreflect.EnumDescriptor {
	return file_deadletter_proto_enumTypes[0].Descriptor()
}

func (DeadLetterState) Type() protoreflect.EnumType {
	return &file_deadletter_proto_enumTypes[0]
}

func (x DeadLetterState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeadLetterState.Descriptor instead.
func (DeadLetterState) EnumDescriptor() ([]byte, []int) {
	return file_deadletter_proto_rawDescGZIP(), []int{0}
}

type DeadLetterCopyType int32

const (
	DeadLetterCopyType_DEAD_LETTER_COPY_TYPE_UNSET    DeadLetterCopyType = 0
	DeadLetterCopyType_DEAD_LETTER_COPY_TYPE_FILTER   DeadLetterCopyType = 1
	DeadLetterCopyType_DEAD_LETTER_COPY_TYPE_EXPLICIT DeadLetterCopyType = 2
)

// Enum value maps for DeadLetterCopyType.
var (
	DeadLetterCopyType_name = map[int32]string{
		0: "DEAD_LETTER_COPY_TYPE_UNSET",
		1: "DEAD_LETTER_COPY_TYPE_FILTER",
		2: "DEAD_LETTER_COPY_TYPE_EXPLICIT",
	}
	DeadLetterCopyType_value = map[string]int32{
		"DEAD_LETTER_COPY_TYPE_UNSET":    0,
		"DEAD_LETTER_COPY_TYPE_FILTER":   1,
		"DEAD_LETTER_COPY_TYPE_EXPLICIT": 2,
	}
)

func (x DeadLetterCopyType) Enum() *DeadLetterCopyType {
	p := new(DeadLetterCopyType)
	*p = x
	return p
}

func (x DeadLetterCopyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeadLetterCopyType) Descriptor() protoreflect.EnumDescriptor {
	return file_deadletter_proto_enumTypes[1].Descriptor()
}

func (DeadLetterCopyType) Type() protoreflect.EnumType {
	return &file_deadletter_proto_enumTypes[1]
}

func (x DeadLetterCopyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeadLetterCopyType.Descriptor instead.
func (DeadLetterCopyType) EnumDescriptor() ([]byte, []int) {
	return file_deadletter_proto_rawDescGZIP(), []int{1}
}

type DeadLetterDeleteType int32

const (
	DeadLetterDeleteType_DEAD_LETTER_DELETE_TYPE_UNSET    DeadLetterDeleteType = 0
	DeadLetterDeleteType_DEAD_LETTER_DELETE_TYPE_EXPLICIT DeadLetterDeleteType = 1
	DeadLetterDeleteType_DEAD_LETTER_DELETE_TYPE_ALL      DeadLetterDeleteType = 2
	DeadLetterDeleteType_DEAD_LETTER_DELETE_TYPE_FILTER   DeadLetterDeleteType = 3
)

// Enum value maps for DeadLetterDeleteType.
var (
	DeadLetterDeleteType_name = map[int32]string{
		0: "DEAD_LETTER_DELETE_TYPE_UNSET",
		1: "DEAD_LETTER_DELETE_TYPE_EXPLICIT",
		2: "DEAD_LETTER_DELETE_TYPE_ALL",
		3: "DEAD_LETTER_DELETE_TYPE_FILTER",
	}
	DeadLetterDeleteType_value = map[string]int32{
		"DEAD_LETTER_DELETE_TYPE_UNSET":    0,
		"DEAD_LETTER_DELETE_TYPE_EXPLICIT": 1,
		"DEAD_LETTER_DELETE_TYPE_ALL":      2,
		"DEAD_LETTER_DELETE_TYPE_FILTER":   3,
	}
)

func (x DeadLetterDeleteType) Enum() *DeadLetterDeleteType {
	p := new(DeadLetterDeleteType)
	*p = x
	return p
}

func (x DeadLetterDeleteType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeadLetterDeleteType) Descriptor() protoreflect.EnumDescriptor {
	return file_deadletter_proto_enumTypes[2].Descriptor()
}

func (DeadLetterDeleteType) Type() protoreflect.EnumType {
	return &file_deadletter_proto_enumTypes[2]
}

func (x DeadLetterDeleteType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeadLetterDeleteType.Descriptor instead.
func (DeadLetterDeleteType) EnumDescriptor() ([]byte, []int) {
	return file_deadletter_proto_rawDescGZIP(), []int{2}
}

type DeadLetter_Type int32

const (
	DeadLetter_UNSET DeadLetter_Type = 0
	DeadLetter_COPY  DeadLetter_Type = 1
	// TODO: Remove after v2 launch
	//
	// Deprecated: Do not use.
	DeadLetter_IDS DeadLetter_Type = 2
	// TODO: Remove after v2 launch
	//
	// Deprecated: Do not use.
	DeadLetter_REPLAY   DeadLetter_Type = 3
	DeadLetter_FUNCTION DeadLetter_Type = 4
	DeadLetter_DELETE   DeadLetter_Type = 5
)

// Enum value maps for DeadLetter_Type.
var (
	DeadLetter_Type_name = map[int32]string{
		0: "UNSET",
		1: "COPY",
		2: "IDS",
		3: "REPLAY",
		4: "FUNCTION",
		5: "DELETE",
	}
	DeadLetter_Type_value = map[string]int32{
		"UNSET":    0,
		"COPY":     1,
		"IDS":      2,
		"REPLAY":   3,
		"FUNCTION": 4,
		"DELETE":   5,
	}
)

func (x DeadLetter_Type) Enum() *DeadLetter_Type {
	p := new(DeadLetter_Type)
	*p = x
	return p
}

func (x DeadLetter_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeadLetter_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_deadletter_proto_enumTypes[3].Descriptor()
}

func (DeadLetter_Type) Type() protoreflect.EnumType {
	return &file_deadletter_proto_enumTypes[3]
}

func (x DeadLetter_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeadLetter_Type.Descriptor instead.
func (DeadLetter_Type) EnumDescriptor() ([]byte, []int) {
	return file_deadletter_proto_rawDescGZIP(), []int{0, 0}
}

type DeadLetterIndex_Type int32

const (
	DeadLetterIndex_UNSET       DeadLetterIndex_Type = 0
	DeadLetterIndex_DEAD_LETTER DeadLetterIndex_Type = 1
	DeadLetterIndex_STAGE       DeadLetterIndex_Type = 2
)

// Enum value maps for DeadLetterIndex_Type.
var (
	DeadLetterIndex_Type_name = map[int32]string{
		0: "UNSET",
		1: "DEAD_LETTER",
		2: "STAGE",
	}
	DeadLetterIndex_Type_value = map[string]int32{
		"UNSET":       0,
		"DEAD_LETTER": 1,
		"STAGE":       2,
	}
)

func (x DeadLetterIndex_Type) Enum() *DeadLetterIndex_Type {
	p := new(DeadLetterIndex_Type)
	*p = x
	return p
}

func (x DeadLetterIndex_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeadLetterIndex_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_deadletter_proto_enumTypes[4].Descriptor()
}

func (DeadLetterIndex_Type) Type() protoreflect.EnumType {
	return &file_deadletter_proto_enumTypes[4]
}

func (x DeadLetterIndex_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeadLetterIndex_Type.Descriptor instead.
func (DeadLetterIndex_Type) EnumDescriptor() ([]byte, []int) {
	return file_deadletter_proto_rawDescGZIP(), []int{1, 0}
}

type DeadLetter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO: remove and reserve fields 1 to 10
	//
	// Deprecated: Do not use.
	Type DeadLetter_Type `protobuf:"varint,1,opt,name=type,proto3,enum=events.DeadLetter_Type" json:"type,omitempty"`
	// Deprecated: Do not use.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// Deprecated: Do not use.
	Ids []string `protobuf:"bytes,3,rep,name=ids,proto3" json:"ids,omitempty"`
	// Deprecated: Do not use.
	SourceCollectionId string `protobuf:"bytes,4,opt,name=source_collection_id,json=sourceCollectionId,proto3" json:"source_collection_id,omitempty"`
	// Deprecated: Do not use.
	DestinationCollectionId string `protobuf:"bytes,5,opt,name=destination_collection_id,json=destinationCollectionId,proto3" json:"destination_collection_id,omitempty"`
	// Deprecated: Do not use.
	RequestedAt int64 `protobuf:"varint,6,opt,name=requested_at,json=requestedAt,proto3" json:"requested_at,omitempty"`
	// Deprecated: Do not use.
	RequestedByAccount string `protobuf:"bytes,7,opt,name=requested_by_account,json=requestedByAccount,proto3" json:"requested_by_account,omitempty"`
	// Deprecated: Do not use.
	RequestedByTeam string `protobuf:"bytes,8,opt,name=requested_by_team,json=requestedByTeam,proto3" json:"requested_by_team,omitempty"`
	// Deprecated: Do not use.
	TransformFunction []byte `protobuf:"bytes,9,opt,name=transform_function,json=transformFunction,proto3" json:"transform_function,omitempty"`
	// Deprecated: Do not use.
	DestCollectionToken string `protobuf:"bytes,10,opt,name=dest_collection_token,json=destCollectionToken,proto3" json:"dest_collection_token,omitempty"`
	// TODO: rename to "type" once deprecated fields above are removed
	NewType DeadLetter_Type `protobuf:"varint,11,opt,name=new_type,json=newType,proto3,enum=events.DeadLetter_Type" json:"new_type,omitempty"`
	State   DeadLetterState `protobuf:"varint,12,opt,name=state,proto3,enum=events.DeadLetterState" json:"state,omitempty"`
	// Error messages returned from dead-letter service about the result of an operation
	StateMessages []string `protobuf:"bytes,13,rep,name=state_messages,json=stateMessages,proto3" json:"state_messages,omitempty"`
	// Which collection is associated with this dead letter operation
	// Currently only used for protobuf encoding when manually editing a payload
	CollectionId string `protobuf:"bytes,14,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	// Types that are assignable to Action:
	//
	//	*DeadLetter_Copy
	//	*DeadLetter_Function
	//	*DeadLetter_Delete
	Action isDeadLetter_Action `protobuf_oneof:"action"`
}

func (x *DeadLetter) Reset() {
	*x = DeadLetter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deadletter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeadLetter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeadLetter) ProtoMessage() {}

func (x *DeadLetter) ProtoReflect() protoreflect.Message {
	mi := &file_deadletter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeadLetter.ProtoReflect.Descriptor instead.
func (*DeadLetter) Descriptor() ([]byte, []int) {
	return file_deadletter_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
func (x *DeadLetter) GetType() DeadLetter_Type {
	if x != nil {
		return x.Type
	}
	return DeadLetter_UNSET
}

// Deprecated: Do not use.
func (x *DeadLetter) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// Deprecated: Do not use.
func (x *DeadLetter) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

// Deprecated: Do not use.
func (x *DeadLetter) GetSourceCollectionId() string {
	if x != nil {
		return x.SourceCollectionId
	}
	return ""
}

// Deprecated: Do not use.
func (x *DeadLetter) GetDestinationCollectionId() string {
	if x != nil {
		return x.DestinationCollectionId
	}
	return ""
}

// Deprecated: Do not use.
func (x *DeadLetter) GetRequestedAt() int64 {
	if x != nil {
		return x.RequestedAt
	}
	return 0
}

// Deprecated: Do not use.
func (x *DeadLetter) GetRequestedByAccount() string {
	if x != nil {
		return x.RequestedByAccount
	}
	return ""
}

// Deprecated: Do not use.
func (x *DeadLetter) GetRequestedByTeam() string {
	if x != nil {
		return x.RequestedByTeam
	}
	return ""
}

// Deprecated: Do not use.
func (x *DeadLetter) GetTransformFunction() []byte {
	if x != nil {
		return x.TransformFunction
	}
	return nil
}

// Deprecated: Do not use.
func (x *DeadLetter) GetDestCollectionToken() string {
	if x != nil {
		return x.DestCollectionToken
	}
	return ""
}

func (x *DeadLetter) GetNewType() DeadLetter_Type {
	if x != nil {
		return x.NewType
	}
	return DeadLetter_UNSET
}

func (x *DeadLetter) GetState() DeadLetterState {
	if x != nil {
		return x.State
	}
	return DeadLetterState_DEAD_LETTER_STATE_UNSET
}

func (x *DeadLetter) GetStateMessages() []string {
	if x != nil {
		return x.StateMessages
	}
	return nil
}

func (x *DeadLetter) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (m *DeadLetter) GetAction() isDeadLetter_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *DeadLetter) GetCopy() *DeadLetterCopy {
	if x, ok := x.GetAction().(*DeadLetter_Copy); ok {
		return x.Copy
	}
	return nil
}

func (x *DeadLetter) GetFunction() *DeadLetterFunction {
	if x, ok := x.GetAction().(*DeadLetter_Function); ok {
		return x.Function
	}
	return nil
}

func (x *DeadLetter) GetDelete() *DeadLetterDelete {
	if x, ok := x.GetAction().(*DeadLetter_Delete); ok {
		return x.Delete
	}
	return nil
}

type isDeadLetter_Action interface {
	isDeadLetter_Action()
}

type DeadLetter_Copy struct {
	Copy *DeadLetterCopy `protobuf:"bytes,100,opt,name=copy,proto3,oneof"`
}

type DeadLetter_Function struct {
	Function *DeadLetterFunction `protobuf:"bytes,101,opt,name=function,proto3,oneof"`
}

type DeadLetter_Delete struct {
	Delete *DeadLetterDelete `protobuf:"bytes,102,opt,name=delete,proto3,oneof"`
}

func (*DeadLetter_Copy) isDeadLetter_Action() {}

func (*DeadLetter_Function) isDeadLetter_Action() {}

func (*DeadLetter_Delete) isDeadLetter_Action() {}

type DeadLetterIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type DeadLetterIndex_Type `protobuf:"varint,1,opt,name=type,proto3,enum=events.DeadLetterIndex_Type" json:"type,omitempty"`
	Id   string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeadLetterIndex) Reset() {
	*x = DeadLetterIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deadletter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeadLetterIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeadLetterIndex) ProtoMessage() {}

func (x *DeadLetterIndex) ProtoReflect() protoreflect.Message {
	mi := &file_deadletter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeadLetterIndex.ProtoReflect.Descriptor instead.
func (*DeadLetterIndex) Descriptor() ([]byte, []int) {
	return file_deadletter_proto_rawDescGZIP(), []int{1}
}

func (x *DeadLetterIndex) GetType() DeadLetterIndex_Type {
	if x != nil {
		return x.Type
	}
	return DeadLetterIndex_UNSET
}

func (x *DeadLetterIndex) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeadLetterCopy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type             DeadLetterCopyType `protobuf:"varint,1,opt,name=type,proto3,enum=events.DeadLetterCopyType" json:"type,omitempty"`
	SourceIndex      *DeadLetterIndex   `protobuf:"bytes,2,opt,name=source_index,json=sourceIndex,proto3" json:"source_index,omitempty"`
	DestinationIndex *DeadLetterIndex   `protobuf:"bytes,3,opt,name=destination_index,json=destinationIndex,proto3" json:"destination_index,omitempty"`
	// CopyType influences which one of these will be filled out
	Filter string   `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	Ids    []string `protobuf:"bytes,5,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeadLetterCopy) Reset() {
	*x = DeadLetterCopy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deadletter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeadLetterCopy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeadLetterCopy) ProtoMessage() {}

func (x *DeadLetterCopy) ProtoReflect() protoreflect.Message {
	mi := &file_deadletter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeadLetterCopy.ProtoReflect.Descriptor instead.
func (*DeadLetterCopy) Descriptor() ([]byte, []int) {
	return file_deadletter_proto_rawDescGZIP(), []int{2}
}

func (x *DeadLetterCopy) GetType() DeadLetterCopyType {
	if x != nil {
		return x.Type
	}
	return DeadLetterCopyType_DEAD_LETTER_COPY_TYPE_UNSET
}

func (x *DeadLetterCopy) GetSourceIndex() *DeadLetterIndex {
	if x != nil {
		return x.SourceIndex
	}
	return nil
}

func (x *DeadLetterCopy) GetDestinationIndex() *DeadLetterIndex {
	if x != nil {
		return x.DestinationIndex
	}
	return nil
}

func (x *DeadLetterCopy) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *DeadLetterCopy) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeadLetterFunction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ES index we will be running the function on
	// Functions can be run on any index technically,
	// but the dead-letter service will only allow it on dead-letter-* or stage-* indices
	SourceIndex *DeadLetterIndex `protobuf:"bytes,1,opt,name=source_index,json=sourceIndex,proto3" json:"source_index,omitempty"`
	// Message contains database UUID of the suborbital function to run
	// and the ui-bff association_id for recording metrics for the function
	Function *FunctionAssoc `protobuf:"bytes,2,opt,name=function,proto3" json:"function,omitempty"`
	// Collection information is used by decoder to pull the collection
	// schema definition in order to decode the message
	CollectionId string `protobuf:"bytes,3,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
}

func (x *DeadLetterFunction) Reset() {
	*x = DeadLetterFunction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deadletter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeadLetterFunction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeadLetterFunction) ProtoMessage() {}

func (x *DeadLetterFunction) ProtoReflect() protoreflect.Message {
	mi := &file_deadletter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeadLetterFunction.ProtoReflect.Descriptor instead.
func (*DeadLetterFunction) Descriptor() ([]byte, []int) {
	return file_deadletter_proto_rawDescGZIP(), []int{3}
}

func (x *DeadLetterFunction) GetSourceIndex() *DeadLetterIndex {
	if x != nil {
		return x.SourceIndex
	}
	return nil
}

func (x *DeadLetterFunction) GetFunction() *FunctionAssoc {
	if x != nil {
		return x.Function
	}
	return nil
}

func (x *DeadLetterFunction) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

type DeadLetterDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  DeadLetterDeleteType `protobuf:"varint,1,opt,name=type,proto3,enum=events.DeadLetterDeleteType" json:"type,omitempty"`
	Index *DeadLetterIndex     `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	// Type influences which is filled out
	Filter string   `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	Ids    []string `protobuf:"bytes,4,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeadLetterDelete) Reset() {
	*x = DeadLetterDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deadletter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeadLetterDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeadLetterDelete) ProtoMessage() {}

func (x *DeadLetterDelete) ProtoReflect() protoreflect.Message {
	mi := &file_deadletter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeadLetterDelete.ProtoReflect.Descriptor instead.
func (*DeadLetterDelete) Descriptor() ([]byte, []int) {
	return file_deadletter_proto_rawDescGZIP(), []int{4}
}

func (x *DeadLetterDelete) GetType() DeadLetterDeleteType {
	if x != nil {
		return x.Type
	}
	return DeadLetterDeleteType_DEAD_LETTER_DELETE_TYPE_UNSET
}

func (x *DeadLetterDelete) GetIndex() *DeadLetterIndex {
	if x != nil {
		return x.Index
	}
	return nil
}

func (x *DeadLetterDelete) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *DeadLetterDelete) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_deadletter_proto protoreflect.FileDescriptor

var file_deadletter_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x0e, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x07, 0x0a, 0x0a, 0x44,
	0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x34, 0x0a, 0x14,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x12,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x3e, 0x0a, 0x19, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x17, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0b, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x34, 0x0a, 0x14, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x12, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2e, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f,
	0x74, 0x65, 0x61, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x12,
	0x31, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x15, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x13, 0x64, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x32, 0x0a, 0x08, 0x6e, 0x65,
	0x77, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x63, 0x6f, 0x70,
	0x79, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x70, 0x79, 0x48,
	0x00, 0x52, 0x04, 0x63, 0x6f, 0x70, 0x79, 0x12, 0x38, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x32, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x61, 0x64, 0x4c,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x00, 0x52, 0x06, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x52, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a,
	0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x4f, 0x50, 0x59,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x03, 0x49, 0x44, 0x53, 0x10, 0x02, 0x1a, 0x02, 0x08, 0x01, 0x12,
	0x0e, 0x0a, 0x06, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x59, 0x10, 0x03, 0x1a, 0x02, 0x08, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x0a, 0x0a,
	0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x05, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x82, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44,
	0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x44, 0x45, 0x41, 0x44, 0x5f, 0x4c, 0x45, 0x54, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x53, 0x54, 0x41, 0x47, 0x45, 0x10, 0x02, 0x22, 0xec, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x61,
	0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x70, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x70,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x61, 0x64, 0x4c,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x44, 0x0a, 0x11, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x61, 0x64,
	0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x10, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x61, 0x64,
	0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a,
	0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65,
	0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x0b, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x31, 0x0a, 0x08, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73,
	0x73, 0x6f, 0x63, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65,
	0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44,
	0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x2a, 0x65, 0x0a, 0x0f, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x45, 0x41, 0x44, 0x5f, 0x4c, 0x45,
	0x54, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54,
	0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x45, 0x41, 0x44, 0x5f, 0x4c, 0x45, 0x54, 0x54, 0x45,
	0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17,
	0x44, 0x45, 0x41, 0x44, 0x5f, 0x4c, 0x45, 0x54, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x2a, 0x7b, 0x0a, 0x12, 0x44, 0x65, 0x61,
	0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x70, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1f, 0x0a, 0x1b, 0x44, 0x45, 0x41, 0x44, 0x5f, 0x4c, 0x45, 0x54, 0x54, 0x45, 0x52, 0x5f, 0x43,
	0x4f, 0x50, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00,
	0x12, 0x20, 0x0a, 0x1c, 0x44, 0x45, 0x41, 0x44, 0x5f, 0x4c, 0x45, 0x54, 0x54, 0x45, 0x52, 0x5f,
	0x43, 0x4f, 0x50, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52,
	0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x44, 0x45, 0x41, 0x44, 0x5f, 0x4c, 0x45, 0x54, 0x54, 0x45,
	0x52, 0x5f, 0x43, 0x4f, 0x50, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x4c,
	0x49, 0x43, 0x49, 0x54, 0x10, 0x02, 0x2a, 0xa4, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x61, 0x64, 0x4c,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x21, 0x0a, 0x1d, 0x44, 0x45, 0x41, 0x44, 0x5f, 0x4c, 0x45, 0x54, 0x54, 0x45, 0x52, 0x5f, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54,
	0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x44, 0x45, 0x41, 0x44, 0x5f, 0x4c, 0x45, 0x54, 0x54, 0x45,
	0x52, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58,
	0x50, 0x4c, 0x49, 0x43, 0x49, 0x54, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x45, 0x41, 0x44,
	0x5f, 0x4c, 0x45, 0x54, 0x54, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x44, 0x45, 0x41,
	0x44, 0x5f, 0x4c, 0x45, 0x54, 0x54, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x10, 0x03, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deadletter_proto_rawDescOnce sync.Once
	file_deadletter_proto_rawDescData = file_deadletter_proto_rawDesc
)

func file_deadletter_proto_rawDescGZIP() []byte {
	file_deadletter_proto_rawDescOnce.Do(func() {
		file_deadletter_proto_rawDescData = protoimpl.X.CompressGZIP(file_deadletter_proto_rawDescData)
	})
	return file_deadletter_proto_rawDescData
}

var file_deadletter_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_deadletter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_deadletter_proto_goTypes = []interface{}{
	(DeadLetterState)(0),       // 0: events.DeadLetterState
	(DeadLetterCopyType)(0),    // 1: events.DeadLetterCopyType
	(DeadLetterDeleteType)(0),  // 2: events.DeadLetterDeleteType
	(DeadLetter_Type)(0),       // 3: events.DeadLetter.Type
	(DeadLetterIndex_Type)(0),  // 4: events.DeadLetterIndex.Type
	(*DeadLetter)(nil),         // 5: events.DeadLetter
	(*DeadLetterIndex)(nil),    // 6: events.DeadLetterIndex
	(*DeadLetterCopy)(nil),     // 7: events.DeadLetterCopy
	(*DeadLetterFunction)(nil), // 8: events.DeadLetterFunction
	(*DeadLetterDelete)(nil),   // 9: events.DeadLetterDelete
	(*FunctionAssoc)(nil),      // 10: events.FunctionAssoc
}
var file_deadletter_proto_depIdxs = []int32{
	3,  // 0: events.DeadLetter.type:type_name -> events.DeadLetter.Type
	3,  // 1: events.DeadLetter.new_type:type_name -> events.DeadLetter.Type
	0,  // 2: events.DeadLetter.state:type_name -> events.DeadLetterState
	7,  // 3: events.DeadLetter.copy:type_name -> events.DeadLetterCopy
	8,  // 4: events.DeadLetter.function:type_name -> events.DeadLetterFunction
	9,  // 5: events.DeadLetter.delete:type_name -> events.DeadLetterDelete
	4,  // 6: events.DeadLetterIndex.type:type_name -> events.DeadLetterIndex.Type
	1,  // 7: events.DeadLetterCopy.type:type_name -> events.DeadLetterCopyType
	6,  // 8: events.DeadLetterCopy.source_index:type_name -> events.DeadLetterIndex
	6,  // 9: events.DeadLetterCopy.destination_index:type_name -> events.DeadLetterIndex
	6,  // 10: events.DeadLetterFunction.source_index:type_name -> events.DeadLetterIndex
	10, // 11: events.DeadLetterFunction.function:type_name -> events.FunctionAssoc
	2,  // 12: events.DeadLetterDelete.type:type_name -> events.DeadLetterDeleteType
	6,  // 13: events.DeadLetterDelete.index:type_name -> events.DeadLetterIndex
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_deadletter_proto_init() }
func file_deadletter_proto_init() {
	if File_deadletter_proto != nil {
		return
	}
	file_function_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_deadletter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeadLetter); i {
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
		file_deadletter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeadLetterIndex); i {
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
		file_deadletter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeadLetterCopy); i {
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
		file_deadletter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeadLetterFunction); i {
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
		file_deadletter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeadLetterDelete); i {
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
	file_deadletter_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DeadLetter_Copy)(nil),
		(*DeadLetter_Function)(nil),
		(*DeadLetter_Delete)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_deadletter_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_deadletter_proto_goTypes,
		DependencyIndexes: file_deadletter_proto_depIdxs,
		EnumInfos:         file_deadletter_proto_enumTypes,
		MessageInfos:      file_deadletter_proto_msgTypes,
	}.Build()
	File_deadletter_proto = out.File
	file_deadletter_proto_rawDesc = nil
	file_deadletter_proto_goTypes = nil
	file_deadletter_proto_depIdxs = nil
}
