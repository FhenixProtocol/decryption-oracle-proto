// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: oracle/oracle.proto

package oracle

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

const (
	DecryptionOracle_Decrypt_FullMethodName     = "/oracle.DecryptionOracle/Decrypt"
	DecryptionOracle_Reencrypt_FullMethodName   = "/oracle.DecryptionOracle/Reencrypt"
	DecryptionOracle_AssertIsNil_FullMethodName = "/oracle.DecryptionOracle/AssertIsNil"
)

// DecryptionOracleClient is the client API for DecryptionOracle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DecryptionOracleClient interface {
	// Sends a greeting
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
	Reencrypt(ctx context.Context, in *ReencryptRequest, opts ...grpc.CallOption) (*ReencryptResponse, error)
	AssertIsNil(ctx context.Context, in *IsNilRequest, opts ...grpc.CallOption) (*IsNilResponse, error)
}

type decryptionOracleClient struct {
	cc grpc.ClientConnInterface
}

func NewDecryptionOracleClient(cc grpc.ClientConnInterface) DecryptionOracleClient {
	return &decryptionOracleClient{cc}
}

func (c *decryptionOracleClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	out := new(DecryptResponse)
	err := c.cc.Invoke(ctx, DecryptionOracle_Decrypt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decryptionOracleClient) Reencrypt(ctx context.Context, in *ReencryptRequest, opts ...grpc.CallOption) (*ReencryptResponse, error) {
	out := new(ReencryptResponse)
	err := c.cc.Invoke(ctx, DecryptionOracle_Reencrypt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decryptionOracleClient) AssertIsNil(ctx context.Context, in *IsNilRequest, opts ...grpc.CallOption) (*IsNilResponse, error) {
	out := new(IsNilResponse)
	err := c.cc.Invoke(ctx, DecryptionOracle_AssertIsNil_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DecryptionOracleServer is the server API for DecryptionOracle service.
// All implementations must embed UnimplementedDecryptionOracleServer
// for forward compatibility
type DecryptionOracleServer interface {
	// Sends a greeting
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
	Reencrypt(context.Context, *ReencryptRequest) (*ReencryptResponse, error)
	AssertIsNil(context.Context, *IsNilRequest) (*IsNilResponse, error)
	mustEmbedUnimplementedDecryptionOracleServer()
}

// UnimplementedDecryptionOracleServer must be embedded to have forward compatible implementations.
type UnimplementedDecryptionOracleServer struct {
}

func (UnimplementedDecryptionOracleServer) Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (UnimplementedDecryptionOracleServer) Reencrypt(context.Context, *ReencryptRequest) (*ReencryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reencrypt not implemented")
}
func (UnimplementedDecryptionOracleServer) AssertIsNil(context.Context, *IsNilRequest) (*IsNilResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssertIsNil not implemented")
}
func (UnimplementedDecryptionOracleServer) mustEmbedUnimplementedDecryptionOracleServer() {}

// UnsafeDecryptionOracleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DecryptionOracleServer will
// result in compilation errors.
type UnsafeDecryptionOracleServer interface {
	mustEmbedUnimplementedDecryptionOracleServer()
}

func RegisterDecryptionOracleServer(s grpc.ServiceRegistrar, srv DecryptionOracleServer) {
	s.RegisterService(&DecryptionOracle_ServiceDesc, srv)
}

func _DecryptionOracle_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecryptionOracleServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecryptionOracle_Decrypt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecryptionOracleServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecryptionOracle_Reencrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReencryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecryptionOracleServer).Reencrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecryptionOracle_Reencrypt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecryptionOracleServer).Reencrypt(ctx, req.(*ReencryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecryptionOracle_AssertIsNil_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsNilRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecryptionOracleServer).AssertIsNil(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecryptionOracle_AssertIsNil_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecryptionOracleServer).AssertIsNil(ctx, req.(*IsNilRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DecryptionOracle_ServiceDesc is the grpc.ServiceDesc for DecryptionOracle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DecryptionOracle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oracle.DecryptionOracle",
	HandlerType: (*DecryptionOracleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Decrypt",
			Handler:    _DecryptionOracle_Decrypt_Handler,
		},
		{
			MethodName: "Reencrypt",
			Handler:    _DecryptionOracle_Reencrypt_Handler,
		},
		{
			MethodName: "AssertIsNil",
			Handler:    _DecryptionOracle_AssertIsNil_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oracle/oracle.proto",
}