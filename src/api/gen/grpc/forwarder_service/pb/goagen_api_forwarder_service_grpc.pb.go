// Code generated with goa v3.19.1, DO NOT EDIT.
//
// ForwarderService protocol buffer definition
//
// Command:
// $ goa gen
// github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/api/design
// --output ./src/api

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: goagen_api_forwarder_service.proto

package forwarder_servicepb

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
	ForwarderService_Forward_FullMethodName = "/forwarder_service.ForwarderService/Forward"
)

// ForwarderServiceClient is the client API for ForwarderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The forwarder service forwards warnings to the appropriate channel
type ForwarderServiceClient interface {
	// Forwards a warning to the appropriate channel
	Forward(ctx context.Context, in *ForwardRequest, opts ...grpc.CallOption) (*ForwardResponse, error)
}

type forwarderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewForwarderServiceClient(cc grpc.ClientConnInterface) ForwarderServiceClient {
	return &forwarderServiceClient{cc}
}

func (c *forwarderServiceClient) Forward(ctx context.Context, in *ForwardRequest, opts ...grpc.CallOption) (*ForwardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ForwardResponse)
	err := c.cc.Invoke(ctx, ForwarderService_Forward_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForwarderServiceServer is the server API for ForwarderService service.
// All implementations must embed UnimplementedForwarderServiceServer
// for forward compatibility.
//
// The forwarder service forwards warnings to the appropriate channel
type ForwarderServiceServer interface {
	// Forwards a warning to the appropriate channel
	Forward(context.Context, *ForwardRequest) (*ForwardResponse, error)
	mustEmbedUnimplementedForwarderServiceServer()
}

// UnimplementedForwarderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedForwarderServiceServer struct{}

func (UnimplementedForwarderServiceServer) Forward(context.Context, *ForwardRequest) (*ForwardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Forward not implemented")
}
func (UnimplementedForwarderServiceServer) mustEmbedUnimplementedForwarderServiceServer() {}
func (UnimplementedForwarderServiceServer) testEmbeddedByValue()                          {}

// UnsafeForwarderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ForwarderServiceServer will
// result in compilation errors.
type UnsafeForwarderServiceServer interface {
	mustEmbedUnimplementedForwarderServiceServer()
}

func RegisterForwarderServiceServer(s grpc.ServiceRegistrar, srv ForwarderServiceServer) {
	// If the following call pancis, it indicates UnimplementedForwarderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ForwarderService_ServiceDesc, srv)
}

func _ForwarderService_Forward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForwardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwarderServiceServer).Forward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ForwarderService_Forward_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwarderServiceServer).Forward(ctx, req.(*ForwardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ForwarderService_ServiceDesc is the grpc.ServiceDesc for ForwarderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ForwarderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "forwarder_service.ForwarderService",
	HandlerType: (*ForwarderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Forward",
			Handler:    _ForwarderService_Forward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goagen_api_forwarder_service.proto",
}
