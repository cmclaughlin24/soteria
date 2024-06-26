// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: iam.proto

package iam

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

// IamServiceClient is the client API for IamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IamServiceClient interface {
	VerifyAccessToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*AccessTokenClaimsReponse, error)
	VerifyApiKey(ctx context.Context, in *VerifyApiKeyRequest, opts ...grpc.CallOption) (*ApiKeyClaimsReponse, error)
}

type iamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIamServiceClient(cc grpc.ClientConnInterface) IamServiceClient {
	return &iamServiceClient{cc}
}

func (c *iamServiceClient) VerifyAccessToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*AccessTokenClaimsReponse, error) {
	out := new(AccessTokenClaimsReponse)
	err := c.cc.Invoke(ctx, "/iam.IamService/VerifyAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamServiceClient) VerifyApiKey(ctx context.Context, in *VerifyApiKeyRequest, opts ...grpc.CallOption) (*ApiKeyClaimsReponse, error) {
	out := new(ApiKeyClaimsReponse)
	err := c.cc.Invoke(ctx, "/iam.IamService/VerifyApiKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IamServiceServer is the server API for IamService service.
// All implementations must embed UnimplementedIamServiceServer
// for forward compatibility
type IamServiceServer interface {
	VerifyAccessToken(context.Context, *VerifyTokenRequest) (*AccessTokenClaimsReponse, error)
	VerifyApiKey(context.Context, *VerifyApiKeyRequest) (*ApiKeyClaimsReponse, error)
	mustEmbedUnimplementedIamServiceServer()
}

// UnimplementedIamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIamServiceServer struct {
}

func (UnimplementedIamServiceServer) VerifyAccessToken(context.Context, *VerifyTokenRequest) (*AccessTokenClaimsReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAccessToken not implemented")
}
func (UnimplementedIamServiceServer) VerifyApiKey(context.Context, *VerifyApiKeyRequest) (*ApiKeyClaimsReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyApiKey not implemented")
}
func (UnimplementedIamServiceServer) mustEmbedUnimplementedIamServiceServer() {}

// UnsafeIamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IamServiceServer will
// result in compilation errors.
type UnsafeIamServiceServer interface {
	mustEmbedUnimplementedIamServiceServer()
}

func RegisterIamServiceServer(s grpc.ServiceRegistrar, srv IamServiceServer) {
	s.RegisterService(&IamService_ServiceDesc, srv)
}

func _IamService_VerifyAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).VerifyAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.IamService/VerifyAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).VerifyAccessToken(ctx, req.(*VerifyTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IamService_VerifyApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServiceServer).VerifyApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.IamService/VerifyApiKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServiceServer).VerifyApiKey(ctx, req.(*VerifyApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IamService_ServiceDesc is the grpc.ServiceDesc for IamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iam.IamService",
	HandlerType: (*IamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyAccessToken",
			Handler:    _IamService_VerifyAccessToken_Handler,
		},
		{
			MethodName: "VerifyApiKey",
			Handler:    _IamService_VerifyApiKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iam.proto",
}
