// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: odin/core/odin.proto

package pb

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

// OdinCoreClient is the client API for OdinCore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OdinCoreClient interface {
	CreateInstance(ctx context.Context, in *CreateInstanceReq, opts ...grpc.CallOption) (*CreateInstanceRsp, error)
}

type odinCoreClient struct {
	cc grpc.ClientConnInterface
}

func NewOdinCoreClient(cc grpc.ClientConnInterface) OdinCoreClient {
	return &odinCoreClient{cc}
}

func (c *odinCoreClient) CreateInstance(ctx context.Context, in *CreateInstanceReq, opts ...grpc.CallOption) (*CreateInstanceRsp, error) {
	out := new(CreateInstanceRsp)
	err := c.cc.Invoke(ctx, "/OdinCore/CreateInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OdinCoreServer is the server API for OdinCore service.
// All implementations must embed UnimplementedOdinCoreServer
// for forward compatibility
type OdinCoreServer interface {
	CreateInstance(context.Context, *CreateInstanceReq) (*CreateInstanceRsp, error)
	mustEmbedUnimplementedOdinCoreServer()
}

// UnimplementedOdinCoreServer must be embedded to have forward compatible implementations.
type UnimplementedOdinCoreServer struct {
}

func (UnimplementedOdinCoreServer) CreateInstance(context.Context, *CreateInstanceReq) (*CreateInstanceRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInstance not implemented")
}
func (UnimplementedOdinCoreServer) mustEmbedUnimplementedOdinCoreServer() {}

// UnsafeOdinCoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OdinCoreServer will
// result in compilation errors.
type UnsafeOdinCoreServer interface {
	mustEmbedUnimplementedOdinCoreServer()
}

func RegisterOdinCoreServer(s grpc.ServiceRegistrar, srv OdinCoreServer) {
	s.RegisterService(&OdinCore_ServiceDesc, srv)
}

func _OdinCore_CreateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OdinCoreServer).CreateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OdinCore/CreateInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OdinCoreServer).CreateInstance(ctx, req.(*CreateInstanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OdinCore_ServiceDesc is the grpc.ServiceDesc for OdinCore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OdinCore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OdinCore",
	HandlerType: (*OdinCoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInstance",
			Handler:    _OdinCore_CreateInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "odin/core/odin.proto",
}
