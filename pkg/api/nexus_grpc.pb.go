// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NexusClient is the client API for Nexus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NexusClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error)
	Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error)
	AddNode(ctx context.Context, in *AddNodeRequest, opts ...grpc.CallOption) (*Status, error)
	RemoveNode(ctx context.Context, in *RemoveNodeRequest, opts ...grpc.CallOption) (*Status, error)
	ListNodes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListNodesResponse, error)
}

type nexusClient struct {
	cc grpc.ClientConnInterface
}

func NewNexusClient(cc grpc.ClientConnInterface) NexusClient {
	return &nexusClient{cc}
}

func (c *nexusClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/nexus.api.Nexus/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nexusClient) Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, "/nexus.api.Nexus/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nexusClient) Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error) {
	out := new(LoadResponse)
	err := c.cc.Invoke(ctx, "/nexus.api.Nexus/Load", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nexusClient) AddNode(ctx context.Context, in *AddNodeRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/nexus.api.Nexus/AddNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nexusClient) RemoveNode(ctx context.Context, in *RemoveNodeRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/nexus.api.Nexus/RemoveNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nexusClient) ListNodes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListNodesResponse, error) {
	out := new(ListNodesResponse)
	err := c.cc.Invoke(ctx, "/nexus.api.Nexus/ListNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NexusServer is the server API for Nexus service.
// All implementations should embed UnimplementedNexusServer
// for forward compatibility
type NexusServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	Save(context.Context, *SaveRequest) (*SaveResponse, error)
	Load(context.Context, *LoadRequest) (*LoadResponse, error)
	AddNode(context.Context, *AddNodeRequest) (*Status, error)
	RemoveNode(context.Context, *RemoveNodeRequest) (*Status, error)
	ListNodes(context.Context, *emptypb.Empty) (*ListNodesResponse, error)
}

// UnimplementedNexusServer should be embedded to have forward compatible implementations.
type UnimplementedNexusServer struct {
}

func (UnimplementedNexusServer) Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedNexusServer) Save(context.Context, *SaveRequest) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedNexusServer) Load(context.Context, *LoadRequest) (*LoadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (UnimplementedNexusServer) AddNode(context.Context, *AddNodeRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNode not implemented")
}
func (UnimplementedNexusServer) RemoveNode(context.Context, *RemoveNodeRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveNode not implemented")
}
func (UnimplementedNexusServer) ListNodes(context.Context, *emptypb.Empty) (*ListNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodes not implemented")
}

// UnsafeNexusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NexusServer will
// result in compilation errors.
type UnsafeNexusServer interface {
	mustEmbedUnimplementedNexusServer()
}

func RegisterNexusServer(s grpc.ServiceRegistrar, srv NexusServer) {
	s.RegisterService(&Nexus_ServiceDesc, srv)
}

func _Nexus_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexusServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.api.Nexus/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexusServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nexus_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexusServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.api.Nexus/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexusServer).Save(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nexus_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexusServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.api.Nexus/Load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexusServer).Load(ctx, req.(*LoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nexus_AddNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexusServer).AddNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.api.Nexus/AddNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexusServer).AddNode(ctx, req.(*AddNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nexus_RemoveNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexusServer).RemoveNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.api.Nexus/RemoveNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexusServer).RemoveNode(ctx, req.(*RemoveNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nexus_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NexusServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nexus.api.Nexus/ListNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NexusServer).ListNodes(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Nexus_ServiceDesc is the grpc.ServiceDesc for Nexus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nexus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nexus.api.Nexus",
	HandlerType: (*NexusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Nexus_Check_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _Nexus_Save_Handler,
		},
		{
			MethodName: "Load",
			Handler:    _Nexus_Load_Handler,
		},
		{
			MethodName: "AddNode",
			Handler:    _Nexus_AddNode_Handler,
		},
		{
			MethodName: "RemoveNode",
			Handler:    _Nexus_RemoveNode_Handler,
		},
		{
			MethodName: "ListNodes",
			Handler:    _Nexus_ListNodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/nexus.proto",
}
