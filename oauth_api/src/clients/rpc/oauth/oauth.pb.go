// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: src/clients/rpc/oauth/oauth.proto

package oauth

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_src_clients_rpc_oauth_oauth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Uuid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Uuid) Reset() {
	*x = Uuid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uuid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uuid) ProtoMessage() {}

func (x *Uuid) ProtoReflect() protoreflect.Message {
	mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uuid.ProtoReflect.Descriptor instead.
func (*Uuid) Descriptor() ([]byte, []int) {
	return file_src_clients_rpc_oauth_oauth_proto_rawDescGZIP(), []int{1}
}

func (x *Uuid) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ServiceKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *ServiceKey) Reset() {
	*x = ServiceKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceKey) ProtoMessage() {}

func (x *ServiceKey) ProtoReflect() protoreflect.Message {
	mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceKey.ProtoReflect.Descriptor instead.
func (*ServiceKey) Descriptor() ([]byte, []int) {
	return file_src_clients_rpc_oauth_oauth_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type JWT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt string `protobuf:"bytes,1,opt,name=Jwt,proto3" json:"Jwt,omitempty"`
}

func (x *JWT) Reset() {
	*x = JWT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWT) ProtoMessage() {}

func (x *JWT) ProtoReflect() protoreflect.Message {
	mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWT.ProtoReflect.Descriptor instead.
func (*JWT) Descriptor() ([]byte, []int) {
	return file_src_clients_rpc_oauth_oauth_proto_rawDescGZIP(), []int{3}
}

func (x *JWT) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type JWTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt      string  `protobuf:"bytes,1,opt,name=Jwt,proto3" json:"Jwt,omitempty"`
	Response *SvrMsg `protobuf:"bytes,2,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *JWTResponse) Reset() {
	*x = JWTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTResponse) ProtoMessage() {}

func (x *JWTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTResponse.ProtoReflect.Descriptor instead.
func (*JWTResponse) Descriptor() ([]byte, []int) {
	return file_src_clients_rpc_oauth_oauth_proto_rawDescGZIP(), []int{4}
}

func (x *JWTResponse) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *JWTResponse) GetResponse() *SvrMsg {
	if x != nil {
		return x.Response
	}
	return nil
}

type JWTAndUuidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt      string  `protobuf:"bytes,1,opt,name=Jwt,proto3" json:"Jwt,omitempty"`
	Uuid     *Uuid   `protobuf:"bytes,2,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	Response *SvrMsg `protobuf:"bytes,3,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *JWTAndUuidResponse) Reset() {
	*x = JWTAndUuidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTAndUuidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTAndUuidResponse) ProtoMessage() {}

func (x *JWTAndUuidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTAndUuidResponse.ProtoReflect.Descriptor instead.
func (*JWTAndUuidResponse) Descriptor() ([]byte, []int) {
	return file_src_clients_rpc_oauth_oauth_proto_rawDescGZIP(), []int{5}
}

func (x *JWTAndUuidResponse) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *JWTAndUuidResponse) GetUuid() *Uuid {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *JWTAndUuidResponse) GetResponse() *SvrMsg {
	if x != nil {
		return x.Response
	}
	return nil
}

type EntityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        *Uuid   `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	Permissions int32   `protobuf:"varint,2,opt,name=Permissions,proto3" json:"Permissions,omitempty"`
	Response    *SvrMsg `protobuf:"bytes,3,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *EntityResponse) Reset() {
	*x = EntityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityResponse) ProtoMessage() {}

func (x *EntityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityResponse.ProtoReflect.Descriptor instead.
func (*EntityResponse) Descriptor() ([]byte, []int) {
	return file_src_clients_rpc_oauth_oauth_proto_rawDescGZIP(), []int{6}
}

func (x *EntityResponse) GetUuid() *Uuid {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *EntityResponse) GetPermissions() int32 {
	if x != nil {
		return x.Permissions
	}
	return 0
}

func (x *EntityResponse) GetResponse() *SvrMsg {
	if x != nil {
		return x.Response
	}
	return nil
}

type SvrMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SvrMsg) Reset() {
	*x = SvrMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SvrMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SvrMsg) ProtoMessage() {}

func (x *SvrMsg) ProtoReflect() protoreflect.Message {
	mi := &file_src_clients_rpc_oauth_oauth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SvrMsg.ProtoReflect.Descriptor instead.
func (*SvrMsg) Descriptor() ([]byte, []int) {
	return file_src_clients_rpc_oauth_oauth_proto_rawDescGZIP(), []int{7}
}

func (x *SvrMsg) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SvrMsg) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_src_clients_rpc_oauth_oauth_proto protoreflect.FileDescriptor

var file_src_clients_rpc_oauth_oauth_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6c, 0x79, 0x64, 0x65, 0x76, 0x5f, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x1a,
	0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x0a, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x22, 0x17, 0x0a, 0x03, 0x4a, 0x57,
	0x54, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4a, 0x77, 0x74, 0x22, 0x57, 0x0a, 0x0b, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4a, 0x77, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x66, 0x6c, 0x79, 0x64, 0x65, 0x76, 0x5f,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x76, 0x72, 0x5f, 0x6d,
	0x73, 0x67, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8b, 0x01, 0x0a,
	0x12, 0x4a, 0x57, 0x54, 0x41, 0x6e, 0x64, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4a, 0x77, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x6c, 0x79, 0x64, 0x65, 0x76, 0x5f, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x04, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x66, 0x6c, 0x79, 0x64, 0x65, 0x76, 0x5f, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x76, 0x72, 0x5f, 0x6d, 0x73, 0x67,
	0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x0e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x6c,
	0x79, 0x64, 0x65, 0x76, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x55, 0x75, 0x69, 0x64, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x66, 0x6c, 0x79, 0x64, 0x65, 0x76, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x53, 0x76, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x0a, 0x07, 0x53, 0x76, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x83, 0x02, 0x0a, 0x13, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x66, 0x6c, 0x79, 0x64, 0x65, 0x76, 0x5f,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x66, 0x6c, 0x79, 0x64, 0x65, 0x76,
	0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4a, 0x57, 0x54, 0x41,
	0x6e, 0x64, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4e, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x2e, 0x66, 0x6c, 0x79, 0x64, 0x65, 0x76, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x1a, 0x1e,
	0x2e, 0x66, 0x6c, 0x79, 0x64, 0x65, 0x76, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x16, 0x2e, 0x66, 0x6c, 0x79,
	0x64, 0x65, 0x76, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4a,
	0x57, 0x54, 0x1a, 0x21, 0x2e, 0x66, 0x6c, 0x79, 0x64, 0x65, 0x76, 0x5f, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x64, 0x65, 0x76, 0x73, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x2d, 0x61, 0x70, 0x70, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_src_clients_rpc_oauth_oauth_proto_rawDescOnce sync.Once
	file_src_clients_rpc_oauth_oauth_proto_rawDescData = file_src_clients_rpc_oauth_oauth_proto_rawDesc
)

func file_src_clients_rpc_oauth_oauth_proto_rawDescGZIP() []byte {
	file_src_clients_rpc_oauth_oauth_proto_rawDescOnce.Do(func() {
		file_src_clients_rpc_oauth_oauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_clients_rpc_oauth_oauth_proto_rawDescData)
	})
	return file_src_clients_rpc_oauth_oauth_proto_rawDescData
}

var file_src_clients_rpc_oauth_oauth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_src_clients_rpc_oauth_oauth_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),       // 0: flydev_chat_oauth.LoginRequest
	(*Uuid)(nil),               // 1: flydev_chat_oauth.Uuid
	(*ServiceKey)(nil),         // 2: flydev_chat_oauth.ServiceKey
	(*JWT)(nil),                // 3: flydev_chat_oauth.JWT
	(*JWTResponse)(nil),        // 4: flydev_chat_oauth.JWTResponse
	(*JWTAndUuidResponse)(nil), // 5: flydev_chat_oauth.JWTAndUuidResponse
	(*EntityResponse)(nil),     // 6: flydev_chat_oauth.EntityResponse
	(*SvrMsg)(nil),             // 7: flydev_chat_oauth.Svr_msg
}
var file_src_clients_rpc_oauth_oauth_proto_depIdxs = []int32{
	7, // 0: flydev_chat_oauth.JWTResponse.Response:type_name -> flydev_chat_oauth.Svr_msg
	1, // 1: flydev_chat_oauth.JWTAndUuidResponse.Uuid:type_name -> flydev_chat_oauth.Uuid
	7, // 2: flydev_chat_oauth.JWTAndUuidResponse.Response:type_name -> flydev_chat_oauth.Svr_msg
	1, // 3: flydev_chat_oauth.EntityResponse.Uuid:type_name -> flydev_chat_oauth.Uuid
	7, // 4: flydev_chat_oauth.EntityResponse.Response:type_name -> flydev_chat_oauth.Svr_msg
	0, // 5: flydev_chat_oauth.OauthProtoInterface.LoginUser:input_type -> flydev_chat_oauth.LoginRequest
	2, // 6: flydev_chat_oauth.OauthProtoInterface.LoginClient:input_type -> flydev_chat_oauth.ServiceKey
	3, // 7: flydev_chat_oauth.OauthProtoInterface.Verify:input_type -> flydev_chat_oauth.JWT
	5, // 8: flydev_chat_oauth.OauthProtoInterface.LoginUser:output_type -> flydev_chat_oauth.JWTAndUuidResponse
	4, // 9: flydev_chat_oauth.OauthProtoInterface.LoginClient:output_type -> flydev_chat_oauth.JWTResponse
	6, // 10: flydev_chat_oauth.OauthProtoInterface.Verify:output_type -> flydev_chat_oauth.EntityResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_src_clients_rpc_oauth_oauth_proto_init() }
func file_src_clients_rpc_oauth_oauth_proto_init() {
	if File_src_clients_rpc_oauth_oauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_clients_rpc_oauth_oauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_src_clients_rpc_oauth_oauth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uuid); i {
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
		file_src_clients_rpc_oauth_oauth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceKey); i {
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
		file_src_clients_rpc_oauth_oauth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWT); i {
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
		file_src_clients_rpc_oauth_oauth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTResponse); i {
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
		file_src_clients_rpc_oauth_oauth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTAndUuidResponse); i {
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
		file_src_clients_rpc_oauth_oauth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityResponse); i {
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
		file_src_clients_rpc_oauth_oauth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SvrMsg); i {
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
			RawDescriptor: file_src_clients_rpc_oauth_oauth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_clients_rpc_oauth_oauth_proto_goTypes,
		DependencyIndexes: file_src_clients_rpc_oauth_oauth_proto_depIdxs,
		MessageInfos:      file_src_clients_rpc_oauth_oauth_proto_msgTypes,
	}.Build()
	File_src_clients_rpc_oauth_oauth_proto = out.File
	file_src_clients_rpc_oauth_oauth_proto_rawDesc = nil
	file_src_clients_rpc_oauth_oauth_proto_goTypes = nil
	file_src_clients_rpc_oauth_oauth_proto_depIdxs = nil
}
