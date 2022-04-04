// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: websocket_service.proto

package websocket_service

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

// WebSocketServiceClient is the client API for WebSocketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebSocketServiceClient interface {
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*SendMessageResponse, error)
}

type webSocketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebSocketServiceClient(cc grpc.ClientConnInterface) WebSocketServiceClient {
	return &webSocketServiceClient{cc}
}

func (c *webSocketServiceClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/genproto.WebSocketService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebSocketServiceServer is the server API for WebSocketService service.
// All implementations must embed UnimplementedWebSocketServiceServer
// for forward compatibility
type WebSocketServiceServer interface {
	SendMessage(context.Context, *Message) (*SendMessageResponse, error)
	mustEmbedUnimplementedWebSocketServiceServer()
}

// UnimplementedWebSocketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWebSocketServiceServer struct {
}

func (UnimplementedWebSocketServiceServer) SendMessage(context.Context, *Message) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedWebSocketServiceServer) mustEmbedUnimplementedWebSocketServiceServer() {}

// UnsafeWebSocketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebSocketServiceServer will
// result in compilation errors.
type UnsafeWebSocketServiceServer interface {
	mustEmbedUnimplementedWebSocketServiceServer()
}

func RegisterWebSocketServiceServer(s grpc.ServiceRegistrar, srv WebSocketServiceServer) {
	s.RegisterService(&WebSocketService_ServiceDesc, srv)
}

func _WebSocketService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebSocketServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.WebSocketService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebSocketServiceServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// WebSocketService_ServiceDesc is the grpc.ServiceDesc for WebSocketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebSocketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.WebSocketService",
	HandlerType: (*WebSocketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _WebSocketService_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "websocket_service.proto",
}
