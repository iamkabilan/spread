// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/chunk.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChunkService_StoreChunk_FullMethodName = "/storage.ChunkService/StoreChunk"
	ChunkService_GetChunk_FullMethodName   = "/storage.ChunkService/GetChunk"
)

// ChunkServiceClient is the client API for ChunkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service for handling file chunks
type ChunkServiceClient interface {
	StoreChunk(ctx context.Context, in *StoreChunkRequest, opts ...grpc.CallOption) (*StoreChunkResponse, error)
	GetChunk(ctx context.Context, in *GetChunkRequest, opts ...grpc.CallOption) (*GetChunkResponse, error)
}

type chunkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChunkServiceClient(cc grpc.ClientConnInterface) ChunkServiceClient {
	return &chunkServiceClient{cc}
}

func (c *chunkServiceClient) StoreChunk(ctx context.Context, in *StoreChunkRequest, opts ...grpc.CallOption) (*StoreChunkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreChunkResponse)
	err := c.cc.Invoke(ctx, ChunkService_StoreChunk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServiceClient) GetChunk(ctx context.Context, in *GetChunkRequest, opts ...grpc.CallOption) (*GetChunkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChunkResponse)
	err := c.cc.Invoke(ctx, ChunkService_GetChunk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChunkServiceServer is the server API for ChunkService service.
// All implementations must embed UnimplementedChunkServiceServer
// for forward compatibility.
//
// Service for handling file chunks
type ChunkServiceServer interface {
	StoreChunk(context.Context, *StoreChunkRequest) (*StoreChunkResponse, error)
	GetChunk(context.Context, *GetChunkRequest) (*GetChunkResponse, error)
	mustEmbedUnimplementedChunkServiceServer()
}

// UnimplementedChunkServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChunkServiceServer struct{}

func (UnimplementedChunkServiceServer) StoreChunk(context.Context, *StoreChunkRequest) (*StoreChunkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreChunk not implemented")
}
func (UnimplementedChunkServiceServer) GetChunk(context.Context, *GetChunkRequest) (*GetChunkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChunk not implemented")
}
func (UnimplementedChunkServiceServer) mustEmbedUnimplementedChunkServiceServer() {}
func (UnimplementedChunkServiceServer) testEmbeddedByValue()                      {}

// UnsafeChunkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChunkServiceServer will
// result in compilation errors.
type UnsafeChunkServiceServer interface {
	mustEmbedUnimplementedChunkServiceServer()
}

func RegisterChunkServiceServer(s grpc.ServiceRegistrar, srv ChunkServiceServer) {
	// If the following call pancis, it indicates UnimplementedChunkServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChunkService_ServiceDesc, srv)
}

func _ChunkService_StoreChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServiceServer).StoreChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChunkService_StoreChunk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServiceServer).StoreChunk(ctx, req.(*StoreChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkService_GetChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServiceServer).GetChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChunkService_GetChunk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServiceServer).GetChunk(ctx, req.(*GetChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChunkService_ServiceDesc is the grpc.ServiceDesc for ChunkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChunkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.ChunkService",
	HandlerType: (*ChunkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreChunk",
			Handler:    _ChunkService_StoreChunk_Handler,
		},
		{
			MethodName: "GetChunk",
			Handler:    _ChunkService_GetChunk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/chunk.proto",
}
