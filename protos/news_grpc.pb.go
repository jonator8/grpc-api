// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: protos/news.proto

package protos

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

// NewsApiClient is the client API for NewsApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsApiClient interface {
	GetNews(ctx context.Context, in *GetNewsRequest, opts ...grpc.CallOption) (*NewsResponse, error)
	GetNew(ctx context.Context, in *GetNewRequest, opts ...grpc.CallOption) (*New, error)
}

type newsApiClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsApiClient(cc grpc.ClientConnInterface) NewsApiClient {
	return &newsApiClient{cc}
}

func (c *newsApiClient) GetNews(ctx context.Context, in *GetNewsRequest, opts ...grpc.CallOption) (*NewsResponse, error) {
	out := new(NewsResponse)
	err := c.cc.Invoke(ctx, "/news.api.NewsApi/GetNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsApiClient) GetNew(ctx context.Context, in *GetNewRequest, opts ...grpc.CallOption) (*New, error) {
	out := new(New)
	err := c.cc.Invoke(ctx, "/news.api.NewsApi/GetNew", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsApiServer is the server API for NewsApi service.
// All implementations must embed UnimplementedNewsApiServer
// for forward compatibility
type NewsApiServer interface {
	GetNews(context.Context, *GetNewsRequest) (*NewsResponse, error)
	GetNew(context.Context, *GetNewRequest) (*New, error)
	mustEmbedUnimplementedNewsApiServer()
}

// UnimplementedNewsApiServer must be embedded to have forward compatible implementations.
type UnimplementedNewsApiServer struct {
}

func (UnimplementedNewsApiServer) GetNews(context.Context, *GetNewsRequest) (*NewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNews not implemented")
}
func (UnimplementedNewsApiServer) GetNew(context.Context, *GetNewRequest) (*New, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNew not implemented")
}
func (UnimplementedNewsApiServer) mustEmbedUnimplementedNewsApiServer() {}

// UnsafeNewsApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsApiServer will
// result in compilation errors.
type UnsafeNewsApiServer interface {
	mustEmbedUnimplementedNewsApiServer()
}

func RegisterNewsApiServer(s grpc.ServiceRegistrar, srv NewsApiServer) {
	s.RegisterService(&NewsApi_ServiceDesc, srv)
}

func _NewsApi_GetNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsApiServer).GetNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.api.NewsApi/GetNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsApiServer).GetNews(ctx, req.(*GetNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsApi_GetNew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsApiServer).GetNew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.api.NewsApi/GetNew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsApiServer).GetNew(ctx, req.(*GetNewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NewsApi_ServiceDesc is the grpc.ServiceDesc for NewsApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewsApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "news.api.NewsApi",
	HandlerType: (*NewsApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNews",
			Handler:    _NewsApi_GetNews_Handler,
		},
		{
			MethodName: "GetNew",
			Handler:    _NewsApi_GetNew_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/news.proto",
}
