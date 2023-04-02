// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: proto/contract.proto

package contractpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//
	//	*Message_ChallengeRequest
	//	*Message_ChallengeResponse
	//	*Message_ServiceRequest
	//	*Message_ServiceResponse
	Body isMessage_Body `protobuf_oneof:"body"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_contract_proto_rawDescGZIP(), []int{0}
}

func (m *Message) GetBody() isMessage_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Message) GetChallengeRequest() *ChallengeRequest {
	if x, ok := x.GetBody().(*Message_ChallengeRequest); ok {
		return x.ChallengeRequest
	}
	return nil
}

func (x *Message) GetChallengeResponse() *ChallengeResponse {
	if x, ok := x.GetBody().(*Message_ChallengeResponse); ok {
		return x.ChallengeResponse
	}
	return nil
}

func (x *Message) GetServiceRequest() *ServiceRequest {
	if x, ok := x.GetBody().(*Message_ServiceRequest); ok {
		return x.ServiceRequest
	}
	return nil
}

func (x *Message) GetServiceResponse() *ServiceResponse {
	if x, ok := x.GetBody().(*Message_ServiceResponse); ok {
		return x.ServiceResponse
	}
	return nil
}

type isMessage_Body interface {
	isMessage_Body()
}

type Message_ChallengeRequest struct {
	ChallengeRequest *ChallengeRequest `protobuf:"bytes,1,opt,name=challengeRequest,proto3,oneof"`
}

type Message_ChallengeResponse struct {
	ChallengeResponse *ChallengeResponse `protobuf:"bytes,2,opt,name=challengeResponse,proto3,oneof"`
}

type Message_ServiceRequest struct {
	ServiceRequest *ServiceRequest `protobuf:"bytes,3,opt,name=serviceRequest,proto3,oneof"`
}

type Message_ServiceResponse struct {
	ServiceResponse *ServiceResponse `protobuf:"bytes,4,opt,name=serviceResponse,proto3,oneof"`
}

func (*Message_ChallengeRequest) isMessage_Body() {}

func (*Message_ChallengeResponse) isMessage_Body() {}

func (*Message_ServiceRequest) isMessage_Body() {}

func (*Message_ServiceResponse) isMessage_Body() {}

type ChallengeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChallengeRequest) Reset() {
	*x = ChallengeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeRequest) ProtoMessage() {}

func (x *ChallengeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeRequest.ProtoReflect.Descriptor instead.
func (*ChallengeRequest) Descriptor() ([]byte, []int) {
	return file_proto_contract_proto_rawDescGZIP(), []int{1}
}

type ChallengeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Puzzle *Puzzle `protobuf:"bytes,1,opt,name=puzzle,proto3" json:"puzzle,omitempty"`
}

func (x *ChallengeResponse) Reset() {
	*x = ChallengeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeResponse) ProtoMessage() {}

func (x *ChallengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeResponse.ProtoReflect.Descriptor instead.
func (*ChallengeResponse) Descriptor() ([]byte, []int) {
	return file_proto_contract_proto_rawDescGZIP(), []int{2}
}

func (x *ChallengeResponse) GetPuzzle() *Puzzle {
	if x != nil {
		return x.Puzzle
	}
	return nil
}

type ServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PuzzleSolution *PuzzleSolution `protobuf:"bytes,1,opt,name=puzzleSolution,proto3" json:"puzzleSolution,omitempty"`
}

func (x *ServiceRequest) Reset() {
	*x = ServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceRequest) ProtoMessage() {}

func (x *ServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceRequest.ProtoReflect.Descriptor instead.
func (*ServiceRequest) Descriptor() ([]byte, []int) {
	return file_proto_contract_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceRequest) GetPuzzleSolution() *PuzzleSolution {
	if x != nil {
		return x.PuzzleSolution
	}
	return nil
}

type ServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote string `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (x *ServiceResponse) Reset() {
	*x = ServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceResponse) ProtoMessage() {}

func (x *ServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceResponse.ProtoReflect.Descriptor instead.
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return file_proto_contract_proto_rawDescGZIP(), []int{4}
}

func (x *ServiceResponse) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

type Puzzle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp        *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	CoveredPreImage  []byte                 `protobuf:"bytes,2,opt,name=coveredPreImage,proto3" json:"coveredPreImage,omitempty"`
	CoveredBitsCount int32                  `protobuf:"varint,3,opt,name=coveredBitsCount,proto3" json:"coveredBitsCount,omitempty"`
	Hash             []byte                 `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Puzzle) Reset() {
	*x = Puzzle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Puzzle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Puzzle) ProtoMessage() {}

func (x *Puzzle) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Puzzle.ProtoReflect.Descriptor instead.
func (*Puzzle) Descriptor() ([]byte, []int) {
	return file_proto_contract_proto_rawDescGZIP(), []int{5}
}

func (x *Puzzle) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Puzzle) GetCoveredPreImage() []byte {
	if x != nil {
		return x.CoveredPreImage
	}
	return nil
}

func (x *Puzzle) GetCoveredBitsCount() int32 {
	if x != nil {
		return x.CoveredBitsCount
	}
	return 0
}

func (x *Puzzle) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type PuzzleSolution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PreImage  []byte                 `protobuf:"bytes,2,opt,name=preImage,proto3" json:"preImage,omitempty"`
}

func (x *PuzzleSolution) Reset() {
	*x = PuzzleSolution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contract_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PuzzleSolution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PuzzleSolution) ProtoMessage() {}

func (x *PuzzleSolution) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contract_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PuzzleSolution.ProtoReflect.Descriptor instead.
func (*PuzzleSolution) Descriptor() ([]byte, []int) {
	return file_proto_contract_proto_rawDescGZIP(), []int{6}
}

func (x *PuzzleSolution) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *PuzzleSolution) GetPreImage() []byte {
	if x != nil {
		return x.PreImage
	}
	return nil
}

var File_proto_contract_proto protoreflect.FileDescriptor

var file_proto_contract_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x10, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x11, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00,
	0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a,
	0x11, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x75, 0x7a, 0x7a, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x75, 0x7a, 0x7a, 0x6c, 0x65, 0x52, 0x06, 0x70, 0x75, 0x7a,
	0x7a, 0x6c, 0x65, 0x22, 0x49, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0e, 0x70, 0x75, 0x7a, 0x7a, 0x6c, 0x65, 0x53,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x50, 0x75, 0x7a, 0x7a, 0x6c, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e,
	0x70, 0x75, 0x7a, 0x7a, 0x6c, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27,
	0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x06, 0x50, 0x75, 0x7a, 0x7a,
	0x6c, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x28, 0x0a, 0x0f,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x50, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x50, 0x72,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65,
	0x64, 0x42, 0x69, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x10, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x42, 0x69, 0x74, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x66, 0x0a, 0x0e, 0x50, 0x75, 0x7a, 0x7a, 0x6c, 0x65,
	0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x0e,
	0x5a, 0x0c, 0x2e, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_contract_proto_rawDescOnce sync.Once
	file_proto_contract_proto_rawDescData = file_proto_contract_proto_rawDesc
)

func file_proto_contract_proto_rawDescGZIP() []byte {
	file_proto_contract_proto_rawDescOnce.Do(func() {
		file_proto_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_contract_proto_rawDescData)
	})
	return file_proto_contract_proto_rawDescData
}

var file_proto_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_contract_proto_goTypes = []interface{}{
	(*Message)(nil),               // 0: Message
	(*ChallengeRequest)(nil),      // 1: ChallengeRequest
	(*ChallengeResponse)(nil),     // 2: ChallengeResponse
	(*ServiceRequest)(nil),        // 3: ServiceRequest
	(*ServiceResponse)(nil),       // 4: ServiceResponse
	(*Puzzle)(nil),                // 5: Puzzle
	(*PuzzleSolution)(nil),        // 6: PuzzleSolution
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_proto_contract_proto_depIdxs = []int32{
	1, // 0: Message.challengeRequest:type_name -> ChallengeRequest
	2, // 1: Message.challengeResponse:type_name -> ChallengeResponse
	3, // 2: Message.serviceRequest:type_name -> ServiceRequest
	4, // 3: Message.serviceResponse:type_name -> ServiceResponse
	5, // 4: ChallengeResponse.puzzle:type_name -> Puzzle
	6, // 5: ServiceRequest.puzzleSolution:type_name -> PuzzleSolution
	7, // 6: Puzzle.timestamp:type_name -> google.protobuf.Timestamp
	7, // 7: PuzzleSolution.timestamp:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_proto_contract_proto_init() }
func file_proto_contract_proto_init() {
	if File_proto_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_proto_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeRequest); i {
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
		file_proto_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeResponse); i {
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
		file_proto_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceRequest); i {
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
		file_proto_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceResponse); i {
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
		file_proto_contract_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Puzzle); i {
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
		file_proto_contract_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PuzzleSolution); i {
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
	file_proto_contract_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_ChallengeRequest)(nil),
		(*Message_ChallengeResponse)(nil),
		(*Message_ServiceRequest)(nil),
		(*Message_ServiceResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_contract_proto_goTypes,
		DependencyIndexes: file_proto_contract_proto_depIdxs,
		MessageInfos:      file_proto_contract_proto_msgTypes,
	}.Build()
	File_proto_contract_proto = out.File
	file_proto_contract_proto_rawDesc = nil
	file_proto_contract_proto_goTypes = nil
	file_proto_contract_proto_depIdxs = nil
}
