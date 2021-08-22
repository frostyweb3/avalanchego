// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ghttpproto

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

// HTTPClient is the client API for HTTP service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HTTPClient interface {
	Handle(ctx context.Context, in *HTTPRequest, opts ...grpc.CallOption) (*HTTPResponse, error)
}

type hTTPClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPClient(cc grpc.ClientConnInterface) HTTPClient {
	return &hTTPClient{cc}
}

func (c *hTTPClient) Handle(ctx context.Context, in *HTTPRequest, opts ...grpc.CallOption) (*HTTPResponse, error) {
	out := new(HTTPResponse)
	err := c.cc.Invoke(ctx, "/ghttpproto.HTTP/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPServer is the server API for HTTP service.
// All implementations must embed UnimplementedHTTPServer
// for forward compatibility
type HTTPServer interface {
	Handle(context.Context, *HTTPRequest) (*HTTPResponse, error)
	mustEmbedUnimplementedHTTPServer()
}

// UnimplementedHTTPServer must be embedded to have forward compatible implementations.
type UnimplementedHTTPServer struct {
}

func (UnimplementedHTTPServer) Handle(context.Context, *HTTPRequest) (*HTTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}
func (UnimplementedHTTPServer) mustEmbedUnimplementedHTTPServer() {}

// UnsafeHTTPServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HTTPServer will
// result in compilation errors.
type UnsafeHTTPServer interface {
	mustEmbedUnimplementedHTTPServer()
}

func RegisterHTTPServer(s grpc.ServiceRegistrar, srv HTTPServer) {
	s.RegisterService(&HTTP_ServiceDesc, srv)
}

func _HTTP_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ghttpproto.HTTP/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPServer).Handle(ctx, req.(*HTTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HTTP_ServiceDesc is the grpc.ServiceDesc for HTTP service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HTTP_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ghttpproto.HTTP",
	HandlerType: (*HTTPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _HTTP_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ghttp.proto",
}
