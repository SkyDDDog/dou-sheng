// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: cosService.proto

package service

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

// CosServiceClient is the client API for CosService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CosServiceClient interface {
	UploadFile(ctx context.Context, in *CosUploadRequest, opts ...grpc.CallOption) (*CosUploadResponse, error)
}

type cosServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCosServiceClient(cc grpc.ClientConnInterface) CosServiceClient {
	return &cosServiceClient{cc}
}

func (c *cosServiceClient) UploadFile(ctx context.Context, in *CosUploadRequest, opts ...grpc.CallOption) (*CosUploadResponse, error) {
	out := new(CosUploadResponse)
	err := c.cc.Invoke(ctx, "/pb.CosService/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CosServiceServer is the server API for CosService service.
// All implementations must embed UnimplementedCosServiceServer
// for forward compatibility
type CosServiceServer interface {
	UploadFile(context.Context, *CosUploadRequest) (*CosUploadResponse, error)
	mustEmbedUnimplementedCosServiceServer()
}

// UnimplementedCosServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCosServiceServer struct {
}

func (UnimplementedCosServiceServer) UploadFile(context.Context, *CosUploadRequest) (*CosUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedCosServiceServer) mustEmbedUnimplementedCosServiceServer() {}

// UnsafeCosServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CosServiceServer will
// result in compilation errors.
type UnsafeCosServiceServer interface {
	mustEmbedUnimplementedCosServiceServer()
}

func RegisterCosServiceServer(s grpc.ServiceRegistrar, srv CosServiceServer) {
	s.RegisterService(&CosService_ServiceDesc, srv)
}

func _CosService_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CosUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServiceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CosService/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServiceServer).UploadFile(ctx, req.(*CosUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CosService_ServiceDesc is the grpc.ServiceDesc for CosService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CosService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CosService",
	HandlerType: (*CosServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadFile",
			Handler:    _CosService_UploadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosService.proto",
}