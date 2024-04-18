// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: com.proto

package proto

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

// ComServiceClient is the client API for ComService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComServiceClient interface {
	// Unary communication
	Healthz(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*StatusResponse, error)
	// Server side streaming
	ComServerStreaming(ctx context.Context, in *ContentList, opts ...grpc.CallOption) (ComService_ComServerStreamingClient, error)
	// Client side streaming
	ComClientStreaming(ctx context.Context, opts ...grpc.CallOption) (ComService_ComClientStreamingClient, error)
	// Bidirectional streaming
	ComBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (ComService_ComBidirectionalStreamingClient, error)
}

type comServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComServiceClient(cc grpc.ClientConnInterface) ComServiceClient {
	return &comServiceClient{cc}
}

func (c *comServiceClient) Healthz(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/communicatio_service.ComService/Healthz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comServiceClient) ComServerStreaming(ctx context.Context, in *ContentList, opts ...grpc.CallOption) (ComService_ComServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &ComService_ServiceDesc.Streams[0], "/communicatio_service.ComService/ComServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &comServiceComServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ComService_ComServerStreamingClient interface {
	Recv() (*ComResponse, error)
	grpc.ClientStream
}

type comServiceComServerStreamingClient struct {
	grpc.ClientStream
}

func (x *comServiceComServerStreamingClient) Recv() (*ComResponse, error) {
	m := new(ComResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *comServiceClient) ComClientStreaming(ctx context.Context, opts ...grpc.CallOption) (ComService_ComClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &ComService_ServiceDesc.Streams[1], "/communicatio_service.ComService/ComClientStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &comServiceComClientStreamingClient{stream}
	return x, nil
}

type ComService_ComClientStreamingClient interface {
	Send(*ComRequest) error
	CloseAndRecv() (*MessagesList, error)
	grpc.ClientStream
}

type comServiceComClientStreamingClient struct {
	grpc.ClientStream
}

func (x *comServiceComClientStreamingClient) Send(m *ComRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *comServiceComClientStreamingClient) CloseAndRecv() (*MessagesList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessagesList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *comServiceClient) ComBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (ComService_ComBidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &ComService_ServiceDesc.Streams[2], "/communicatio_service.ComService/ComBidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &comServiceComBidirectionalStreamingClient{stream}
	return x, nil
}

type ComService_ComBidirectionalStreamingClient interface {
	Send(*ComRequest) error
	Recv() (*ComResponse, error)
	grpc.ClientStream
}

type comServiceComBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *comServiceComBidirectionalStreamingClient) Send(m *ComRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *comServiceComBidirectionalStreamingClient) Recv() (*ComResponse, error) {
	m := new(ComResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ComServiceServer is the server API for ComService service.
// All implementations must embed UnimplementedComServiceServer
// for forward compatibility
type ComServiceServer interface {
	// Unary communication
	Healthz(context.Context, *NoParam) (*StatusResponse, error)
	// Server side streaming
	ComServerStreaming(*ContentList, ComService_ComServerStreamingServer) error
	// Client side streaming
	ComClientStreaming(ComService_ComClientStreamingServer) error
	// Bidirectional streaming
	ComBidirectionalStreaming(ComService_ComBidirectionalStreamingServer) error
	mustEmbedUnimplementedComServiceServer()
}

// UnimplementedComServiceServer must be embedded to have forward compatible implementations.
type UnimplementedComServiceServer struct {
}

func (UnimplementedComServiceServer) Healthz(context.Context, *NoParam) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthz not implemented")
}
func (UnimplementedComServiceServer) ComServerStreaming(*ContentList, ComService_ComServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ComServerStreaming not implemented")
}
func (UnimplementedComServiceServer) ComClientStreaming(ComService_ComClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ComClientStreaming not implemented")
}
func (UnimplementedComServiceServer) ComBidirectionalStreaming(ComService_ComBidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ComBidirectionalStreaming not implemented")
}
func (UnimplementedComServiceServer) mustEmbedUnimplementedComServiceServer() {}

// UnsafeComServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComServiceServer will
// result in compilation errors.
type UnsafeComServiceServer interface {
	mustEmbedUnimplementedComServiceServer()
}

func RegisterComServiceServer(s grpc.ServiceRegistrar, srv ComServiceServer) {
	s.RegisterService(&ComService_ServiceDesc, srv)
}

func _ComService_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComServiceServer).Healthz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communicatio_service.ComService/Healthz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComServiceServer).Healthz(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComService_ComServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ContentList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ComServiceServer).ComServerStreaming(m, &comServiceComServerStreamingServer{stream})
}

type ComService_ComServerStreamingServer interface {
	Send(*ComResponse) error
	grpc.ServerStream
}

type comServiceComServerStreamingServer struct {
	grpc.ServerStream
}

func (x *comServiceComServerStreamingServer) Send(m *ComResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ComService_ComClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ComServiceServer).ComClientStreaming(&comServiceComClientStreamingServer{stream})
}

type ComService_ComClientStreamingServer interface {
	SendAndClose(*MessagesList) error
	Recv() (*ComRequest, error)
	grpc.ServerStream
}

type comServiceComClientStreamingServer struct {
	grpc.ServerStream
}

func (x *comServiceComClientStreamingServer) SendAndClose(m *MessagesList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *comServiceComClientStreamingServer) Recv() (*ComRequest, error) {
	m := new(ComRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ComService_ComBidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ComServiceServer).ComBidirectionalStreaming(&comServiceComBidirectionalStreamingServer{stream})
}

type ComService_ComBidirectionalStreamingServer interface {
	Send(*ComResponse) error
	Recv() (*ComRequest, error)
	grpc.ServerStream
}

type comServiceComBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *comServiceComBidirectionalStreamingServer) Send(m *ComResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *comServiceComBidirectionalStreamingServer) Recv() (*ComRequest, error) {
	m := new(ComRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ComService_ServiceDesc is the grpc.ServiceDesc for ComService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "communicatio_service.ComService",
	HandlerType: (*ComServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthz",
			Handler:    _ComService_Healthz_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ComServerStreaming",
			Handler:       _ComService_ComServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComClientStreaming",
			Handler:       _ComService_ComClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ComBidirectionalStreaming",
			Handler:       _ComService_ComBidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "com.proto",
}