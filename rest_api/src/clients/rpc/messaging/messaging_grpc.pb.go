// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package messaging

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessagingProtoInterfaceClient is the client API for MessagingProtoInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagingProtoInterfaceClient interface {
	CreateConversation(ctx context.Context, in *Conversation, opts ...grpc.CallOption) (*UuidMsg, error)
	//Changed, now instead of returning conversations alone it returns an object with a conversation an its participants inside.
	//Later should be fetched directly through JWT info.
	GetConversationsByUser(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*ArrayConversationResponse, error)
	UpdateConversationInfo(ctx context.Context, in *Conversation, opts ...grpc.CallOption) (*UpdateConversationResponse, error)
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*UuidMsg, error)
	//Later change userUuid to use JWT info.
	GetMessagesByConversation(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*ArrayMessageResponse, error)
	UpdateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageMsgResponse, error)
	CreateUserConversation(ctx context.Context, in *CreateUserConversationRequest, opts ...grpc.CallOption) (*SvrMsg, error)
}

type messagingProtoInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingProtoInterfaceClient(cc grpc.ClientConnInterface) MessagingProtoInterfaceClient {
	return &messagingProtoInterfaceClient{cc}
}

func (c *messagingProtoInterfaceClient) CreateConversation(ctx context.Context, in *Conversation, opts ...grpc.CallOption) (*UuidMsg, error) {
	out := new(UuidMsg)
	err := c.cc.Invoke(ctx, "/flydevs_chat_app_messaging.MessagingProtoInterface/CreateConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingProtoInterfaceClient) GetConversationsByUser(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*ArrayConversationResponse, error) {
	out := new(ArrayConversationResponse)
	err := c.cc.Invoke(ctx, "/flydevs_chat_app_messaging.MessagingProtoInterface/GetConversationsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingProtoInterfaceClient) UpdateConversationInfo(ctx context.Context, in *Conversation, opts ...grpc.CallOption) (*UpdateConversationResponse, error) {
	out := new(UpdateConversationResponse)
	err := c.cc.Invoke(ctx, "/flydevs_chat_app_messaging.MessagingProtoInterface/UpdateConversationInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingProtoInterfaceClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*UuidMsg, error) {
	out := new(UuidMsg)
	err := c.cc.Invoke(ctx, "/flydevs_chat_app_messaging.MessagingProtoInterface/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingProtoInterfaceClient) GetMessagesByConversation(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*ArrayMessageResponse, error) {
	out := new(ArrayMessageResponse)
	err := c.cc.Invoke(ctx, "/flydevs_chat_app_messaging.MessagingProtoInterface/GetMessagesByConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingProtoInterfaceClient) UpdateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageMsgResponse, error) {
	out := new(MessageMsgResponse)
	err := c.cc.Invoke(ctx, "/flydevs_chat_app_messaging.MessagingProtoInterface/UpdateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingProtoInterfaceClient) CreateUserConversation(ctx context.Context, in *CreateUserConversationRequest, opts ...grpc.CallOption) (*SvrMsg, error) {
	out := new(SvrMsg)
	err := c.cc.Invoke(ctx, "/flydevs_chat_app_messaging.MessagingProtoInterface/CreateUserConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagingProtoInterfaceServer is the server API for MessagingProtoInterface service.
// All implementations must embed UnimplementedMessagingProtoInterfaceServer
// for forward compatibility
type MessagingProtoInterfaceServer interface {
	CreateConversation(context.Context, *Conversation) (*UuidMsg, error)
	//Changed, now instead of returning conversations alone it returns an object with a conversation an its participants inside.
	//Later should be fetched directly through JWT info.
	GetConversationsByUser(context.Context, *Uuid) (*ArrayConversationResponse, error)
	UpdateConversationInfo(context.Context, *Conversation) (*UpdateConversationResponse, error)
	CreateMessage(context.Context, *CreateMessageRequest) (*UuidMsg, error)
	//Later change userUuid to use JWT info.
	GetMessagesByConversation(context.Context, *Uuid) (*ArrayMessageResponse, error)
	UpdateMessage(context.Context, *Message) (*MessageMsgResponse, error)
	CreateUserConversation(context.Context, *CreateUserConversationRequest) (*SvrMsg, error)
	mustEmbedUnimplementedMessagingProtoInterfaceServer()
}

// UnimplementedMessagingProtoInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedMessagingProtoInterfaceServer struct {
}

func (UnimplementedMessagingProtoInterfaceServer) CreateConversation(context.Context, *Conversation) (*UuidMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConversation not implemented")
}
func (UnimplementedMessagingProtoInterfaceServer) GetConversationsByUser(context.Context, *Uuid) (*ArrayConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversationsByUser not implemented")
}
func (UnimplementedMessagingProtoInterfaceServer) UpdateConversationInfo(context.Context, *Conversation) (*UpdateConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConversationInfo not implemented")
}
func (UnimplementedMessagingProtoInterfaceServer) CreateMessage(context.Context, *CreateMessageRequest) (*UuidMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedMessagingProtoInterfaceServer) GetMessagesByConversation(context.Context, *Uuid) (*ArrayMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessagesByConversation not implemented")
}
func (UnimplementedMessagingProtoInterfaceServer) UpdateMessage(context.Context, *Message) (*MessageMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedMessagingProtoInterfaceServer) CreateUserConversation(context.Context, *CreateUserConversationRequest) (*SvrMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserConversation not implemented")
}
func (UnimplementedMessagingProtoInterfaceServer) mustEmbedUnimplementedMessagingProtoInterfaceServer() {
}

// UnsafeMessagingProtoInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingProtoInterfaceServer will
// result in compilation errors.
type UnsafeMessagingProtoInterfaceServer interface {
	mustEmbedUnimplementedMessagingProtoInterfaceServer()
}

func RegisterMessagingProtoInterfaceServer(s grpc.ServiceRegistrar, srv MessagingProtoInterfaceServer) {
	s.RegisterService(&MessagingProtoInterface_ServiceDesc, srv)
}

func _MessagingProtoInterface_CreateConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Conversation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingProtoInterfaceServer).CreateConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flydevs_chat_app_messaging.MessagingProtoInterface/CreateConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingProtoInterfaceServer).CreateConversation(ctx, req.(*Conversation))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingProtoInterface_GetConversationsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingProtoInterfaceServer).GetConversationsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flydevs_chat_app_messaging.MessagingProtoInterface/GetConversationsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingProtoInterfaceServer).GetConversationsByUser(ctx, req.(*Uuid))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingProtoInterface_UpdateConversationInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Conversation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingProtoInterfaceServer).UpdateConversationInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flydevs_chat_app_messaging.MessagingProtoInterface/UpdateConversationInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingProtoInterfaceServer).UpdateConversationInfo(ctx, req.(*Conversation))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingProtoInterface_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingProtoInterfaceServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flydevs_chat_app_messaging.MessagingProtoInterface/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingProtoInterfaceServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingProtoInterface_GetMessagesByConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingProtoInterfaceServer).GetMessagesByConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flydevs_chat_app_messaging.MessagingProtoInterface/GetMessagesByConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingProtoInterfaceServer).GetMessagesByConversation(ctx, req.(*Uuid))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingProtoInterface_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingProtoInterfaceServer).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flydevs_chat_app_messaging.MessagingProtoInterface/UpdateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingProtoInterfaceServer).UpdateMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingProtoInterface_CreateUserConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingProtoInterfaceServer).CreateUserConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flydevs_chat_app_messaging.MessagingProtoInterface/CreateUserConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingProtoInterfaceServer).CreateUserConversation(ctx, req.(*CreateUserConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessagingProtoInterface_ServiceDesc is the grpc.ServiceDesc for MessagingProtoInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessagingProtoInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flydevs_chat_app_messaging.MessagingProtoInterface",
	HandlerType: (*MessagingProtoInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConversation",
			Handler:    _MessagingProtoInterface_CreateConversation_Handler,
		},
		{
			MethodName: "GetConversationsByUser",
			Handler:    _MessagingProtoInterface_GetConversationsByUser_Handler,
		},
		{
			MethodName: "UpdateConversationInfo",
			Handler:    _MessagingProtoInterface_UpdateConversationInfo_Handler,
		},
		{
			MethodName: "CreateMessage",
			Handler:    _MessagingProtoInterface_CreateMessage_Handler,
		},
		{
			MethodName: "GetMessagesByConversation",
			Handler:    _MessagingProtoInterface_GetMessagesByConversation_Handler,
		},
		{
			MethodName: "UpdateMessage",
			Handler:    _MessagingProtoInterface_UpdateMessage_Handler,
		},
		{
			MethodName: "CreateUserConversation",
			Handler:    _MessagingProtoInterface_CreateUserConversation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/clients/rpc/messaging/messaging.proto",
}
