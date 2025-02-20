// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: sso/protobuf.proto

package services_sso_test

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type RequestRegister struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Username      string                 `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequestRegister) Reset() {
	*x = RequestRegister{}
	mi := &file_sso_protobuf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRegister) ProtoMessage() {}

func (x *RequestRegister) ProtoReflect() protoreflect.Message {
	mi := &file_sso_protobuf_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRegister.ProtoReflect.Descriptor instead.
func (*RequestRegister) Descriptor() ([]byte, []int) {
	return file_sso_protobuf_proto_rawDescGZIP(), []int{0}
}

func (x *RequestRegister) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RequestRegister) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RequestRegister) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ResponseRegister struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResponseRegister) Reset() {
	*x = ResponseRegister{}
	mi := &file_sso_protobuf_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRegister) ProtoMessage() {}

func (x *ResponseRegister) ProtoReflect() protoreflect.Message {
	mi := &file_sso_protobuf_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRegister.ProtoReflect.Descriptor instead.
func (*ResponseRegister) Descriptor() ([]byte, []int) {
	return file_sso_protobuf_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseRegister) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ResponseRegister) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RequestLogin struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequestLogin) Reset() {
	*x = RequestLogin{}
	mi := &file_sso_protobuf_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLogin) ProtoMessage() {}

func (x *RequestLogin) ProtoReflect() protoreflect.Message {
	mi := &file_sso_protobuf_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLogin.ProtoReflect.Descriptor instead.
func (*RequestLogin) Descriptor() ([]byte, []int) {
	return file_sso_protobuf_proto_rawDescGZIP(), []int{2}
}

func (x *RequestLogin) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RequestLogin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ResponseLogin struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResponseLogin) Reset() {
	*x = ResponseLogin{}
	mi := &file_sso_protobuf_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseLogin) ProtoMessage() {}

func (x *ResponseLogin) ProtoReflect() protoreflect.Message {
	mi := &file_sso_protobuf_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseLogin.ProtoReflect.Descriptor instead.
func (*ResponseLogin) Descriptor() ([]byte, []int) {
	return file_sso_protobuf_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseLogin) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RequestIsAdmin struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequestIsAdmin) Reset() {
	*x = RequestIsAdmin{}
	mi := &file_sso_protobuf_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestIsAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestIsAdmin) ProtoMessage() {}

func (x *RequestIsAdmin) ProtoReflect() protoreflect.Message {
	mi := &file_sso_protobuf_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestIsAdmin.ProtoReflect.Descriptor instead.
func (*RequestIsAdmin) Descriptor() ([]byte, []int) {
	return file_sso_protobuf_proto_rawDescGZIP(), []int{4}
}

func (x *RequestIsAdmin) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ResponseIsAdmin struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Response      bool                   `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResponseIsAdmin) Reset() {
	*x = ResponseIsAdmin{}
	mi := &file_sso_protobuf_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseIsAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseIsAdmin) ProtoMessage() {}

func (x *ResponseIsAdmin) ProtoReflect() protoreflect.Message {
	mi := &file_sso_protobuf_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseIsAdmin.ProtoReflect.Descriptor instead.
func (*ResponseIsAdmin) Descriptor() ([]byte, []int) {
	return file_sso_protobuf_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseIsAdmin) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

type ApplyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	SecreteCode   string                 `protobuf:"bytes,2,opt,name=secrete_code,json=secreteCode,proto3" json:"secrete_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApplyRequest) Reset() {
	*x = ApplyRequest{}
	mi := &file_sso_protobuf_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyRequest) ProtoMessage() {}

func (x *ApplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_protobuf_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyRequest.ProtoReflect.Descriptor instead.
func (*ApplyRequest) Descriptor() ([]byte, []int) {
	return file_sso_protobuf_proto_rawDescGZIP(), []int{6}
}

func (x *ApplyRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ApplyRequest) GetSecreteCode() string {
	if x != nil {
		return x.SecreteCode
	}
	return ""
}

type ApplyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Apply         bool                   `protobuf:"varint,1,opt,name=apply,proto3" json:"apply,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApplyResponse) Reset() {
	*x = ApplyResponse{}
	mi := &file_sso_protobuf_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyResponse) ProtoMessage() {}

func (x *ApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_protobuf_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyResponse.ProtoReflect.Descriptor instead.
func (*ApplyResponse) Descriptor() ([]byte, []int) {
	return file_sso_protobuf_proto_rawDescGZIP(), []int{7}
}

func (x *ApplyResponse) GetApply() bool {
	if x != nil {
		return x.Apply
	}
	return false
}

type RetryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RetryRequest) Reset() {
	*x = RetryRequest{}
	mi := &file_sso_protobuf_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RetryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryRequest) ProtoMessage() {}

func (x *RetryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_protobuf_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryRequest.ProtoReflect.Descriptor instead.
func (*RetryRequest) Descriptor() ([]byte, []int) {
	return file_sso_protobuf_proto_rawDescGZIP(), []int{8}
}

func (x *RetryRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RetryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Response      bool                   `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RetryResponse) Reset() {
	*x = RetryResponse{}
	mi := &file_sso_protobuf_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RetryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryResponse) ProtoMessage() {}

func (x *RetryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_protobuf_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryResponse.ProtoReflect.Descriptor instead.
func (*RetryResponse) Descriptor() ([]byte, []int) {
	return file_sso_protobuf_proto_rawDescGZIP(), []int{9}
}

func (x *RetryResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

var File_sso_protobuf_proto protoreflect.FileDescriptor

var file_sso_protobuf_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x73, 0x73, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x73, 0x6f, 0x22, 0x5f, 0x0a, 0x0f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x10, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x40, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0x2d, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x73, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x47, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70,
	0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x70, 0x70, 0x6c, 0x79,
	0x22, 0x24, 0x0a, 0x0c, 0x52, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2b, 0x0a, 0x0d, 0x52, 0x65, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xa5, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x37, 0x0a, 0x08,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x15,
	0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x11,
	0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x1a, 0x12, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x13, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x73,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x14, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x32, 0x6c, 0x0a, 0x0a, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2e, 0x0a, 0x05, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x12, 0x11, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x52, 0x65, 0x74,
	0x72, 0x79, 0x12, 0x11, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x6d, 0x79, 0x5f,
	0x73, 0x73, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x3a, 0x73, 0x73, 0x6f, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_sso_protobuf_proto_rawDescOnce sync.Once
	file_sso_protobuf_proto_rawDescData []byte
)

func file_sso_protobuf_proto_rawDescGZIP() []byte {
	file_sso_protobuf_proto_rawDescOnce.Do(func() {
		file_sso_protobuf_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sso_protobuf_proto_rawDesc), len(file_sso_protobuf_proto_rawDesc)))
	})
	return file_sso_protobuf_proto_rawDescData
}

var file_sso_protobuf_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_sso_protobuf_proto_goTypes = []any{
	(*RequestRegister)(nil),  // 0: sso.RequestRegister
	(*ResponseRegister)(nil), // 1: sso.ResponseRegister
	(*RequestLogin)(nil),     // 2: sso.RequestLogin
	(*ResponseLogin)(nil),    // 3: sso.ResponseLogin
	(*RequestIsAdmin)(nil),   // 4: sso.RequestIsAdmin
	(*ResponseIsAdmin)(nil),  // 5: sso.ResponseIsAdmin
	(*ApplyRequest)(nil),     // 6: sso.ApplyRequest
	(*ApplyResponse)(nil),    // 7: sso.ApplyResponse
	(*RetryRequest)(nil),     // 8: sso.RetryRequest
	(*RetryResponse)(nil),    // 9: sso.RetryResponse
}
var file_sso_protobuf_proto_depIdxs = []int32{
	0, // 0: sso.Auth.Register:input_type -> sso.RequestRegister
	2, // 1: sso.Auth.Login:input_type -> sso.RequestLogin
	4, // 2: sso.Auth.IsAdmin:input_type -> sso.RequestIsAdmin
	6, // 3: sso.ApplyEmail.Apply:input_type -> sso.ApplyRequest
	8, // 4: sso.ApplyEmail.Retry:input_type -> sso.RetryRequest
	1, // 5: sso.Auth.Register:output_type -> sso.ResponseRegister
	3, // 6: sso.Auth.Login:output_type -> sso.ResponseLogin
	5, // 7: sso.Auth.IsAdmin:output_type -> sso.ResponseIsAdmin
	7, // 8: sso.ApplyEmail.Apply:output_type -> sso.ApplyResponse
	9, // 9: sso.ApplyEmail.Retry:output_type -> sso.RetryResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sso_protobuf_proto_init() }
func file_sso_protobuf_proto_init() {
	if File_sso_protobuf_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sso_protobuf_proto_rawDesc), len(file_sso_protobuf_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_sso_protobuf_proto_goTypes,
		DependencyIndexes: file_sso_protobuf_proto_depIdxs,
		MessageInfos:      file_sso_protobuf_proto_msgTypes,
	}.Build()
	File_sso_protobuf_proto = out.File
	file_sso_protobuf_proto_goTypes = nil
	file_sso_protobuf_proto_depIdxs = nil
}
