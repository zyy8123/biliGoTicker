// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: proto/master.proto

package pb

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
	TicketMaster_RegisterWorker_FullMethodName = "/worker.TicketMaster/RegisterWorker"
	TicketMaster_CancelTask_FullMethodName     = "/worker.TicketMaster/CancelTask"
)

// TicketMasterClient is the client API for TicketMaster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// protoc --go_out=. --go-grpc_out=. proto/master.proto
type TicketMasterClient interface {
	RegisterWorker(ctx context.Context, in *WorkerInfo, opts ...grpc.CallOption) (*RegisterReply, error)
	CancelTask(ctx context.Context, in *CancelTaskInfo, opts ...grpc.CallOption) (*CancelReply, error)
}

type ticketMasterClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketMasterClient(cc grpc.ClientConnInterface) TicketMasterClient {
	return &ticketMasterClient{cc}
}

func (c *ticketMasterClient) RegisterWorker(ctx context.Context, in *WorkerInfo, opts ...grpc.CallOption) (*RegisterReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, TicketMaster_RegisterWorker_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketMasterClient) CancelTask(ctx context.Context, in *CancelTaskInfo, opts ...grpc.CallOption) (*CancelReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelReply)
	err := c.cc.Invoke(ctx, TicketMaster_CancelTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketMasterServer is the server API for TicketMaster service.
// All implementations must embed UnimplementedTicketMasterServer
// for forward compatibility.
//
// protoc --go_out=. --go-grpc_out=. proto/master.proto
type TicketMasterServer interface {
	RegisterWorker(context.Context, *WorkerInfo) (*RegisterReply, error)
	CancelTask(context.Context, *CancelTaskInfo) (*CancelReply, error)
	mustEmbedUnimplementedTicketMasterServer()
}

// UnimplementedTicketMasterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTicketMasterServer struct{}

func (UnimplementedTicketMasterServer) RegisterWorker(context.Context, *WorkerInfo) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWorker not implemented")
}
func (UnimplementedTicketMasterServer) CancelTask(context.Context, *CancelTaskInfo) (*CancelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTask not implemented")
}
func (UnimplementedTicketMasterServer) mustEmbedUnimplementedTicketMasterServer() {}
func (UnimplementedTicketMasterServer) testEmbeddedByValue()                      {}

// UnsafeTicketMasterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketMasterServer will
// result in compilation errors.
type UnsafeTicketMasterServer interface {
	mustEmbedUnimplementedTicketMasterServer()
}

func RegisterTicketMasterServer(s grpc.ServiceRegistrar, srv TicketMasterServer) {
	// If the following call pancis, it indicates UnimplementedTicketMasterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TicketMaster_ServiceDesc, srv)
}

func _TicketMaster_RegisterWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketMasterServer).RegisterWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketMaster_RegisterWorker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketMasterServer).RegisterWorker(ctx, req.(*WorkerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketMaster_CancelTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTaskInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketMasterServer).CancelTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketMaster_CancelTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketMasterServer).CancelTask(ctx, req.(*CancelTaskInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketMaster_ServiceDesc is the grpc.ServiceDesc for TicketMaster service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketMaster_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "worker.TicketMaster",
	HandlerType: (*TicketMasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterWorker",
			Handler:    _TicketMaster_RegisterWorker_Handler,
		},
		{
			MethodName: "CancelTask",
			Handler:    _TicketMaster_CancelTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/master.proto",
}
