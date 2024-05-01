// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: github.com/openconfig/gnsi/certz/certz.proto

package cert

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

// CertzClient is the client API for Certz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertzClient interface {
	Rotate(ctx context.Context, opts ...grpc.CallOption) (Certz_RotateClient, error)
	AddProfile(ctx context.Context, in *AddProfileRequest, opts ...grpc.CallOption) (*AddProfileResponse, error)
	DeleteProfile(ctx context.Context, in *DeleteProfileRequest, opts ...grpc.CallOption) (*DeleteProfileResponse, error)
	GetProfileList(ctx context.Context, in *GetProfileListRequest, opts ...grpc.CallOption) (*GetProfileListResponse, error)
	CanGenerateCSR(ctx context.Context, in *CanGenerateCSRRequest, opts ...grpc.CallOption) (*CanGenerateCSRResponse, error)
}

type certzClient struct {
	cc grpc.ClientConnInterface
}

func NewCertzClient(cc grpc.ClientConnInterface) CertzClient {
	return &certzClient{cc}
}

func (c *certzClient) Rotate(ctx context.Context, opts ...grpc.CallOption) (Certz_RotateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Certz_ServiceDesc.Streams[0], "/gnsi.certz.v1.Certz/Rotate", opts...)
	if err != nil {
		return nil, err
	}
	x := &certzRotateClient{stream}
	return x, nil
}

type Certz_RotateClient interface {
	Send(*RotateCertificateRequest) error
	Recv() (*RotateCertificateResponse, error)
	grpc.ClientStream
}

type certzRotateClient struct {
	grpc.ClientStream
}

func (x *certzRotateClient) Send(m *RotateCertificateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *certzRotateClient) Recv() (*RotateCertificateResponse, error) {
	m := new(RotateCertificateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *certzClient) AddProfile(ctx context.Context, in *AddProfileRequest, opts ...grpc.CallOption) (*AddProfileResponse, error) {
	out := new(AddProfileResponse)
	err := c.cc.Invoke(ctx, "/gnsi.certz.v1.Certz/AddProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certzClient) DeleteProfile(ctx context.Context, in *DeleteProfileRequest, opts ...grpc.CallOption) (*DeleteProfileResponse, error) {
	out := new(DeleteProfileResponse)
	err := c.cc.Invoke(ctx, "/gnsi.certz.v1.Certz/DeleteProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certzClient) GetProfileList(ctx context.Context, in *GetProfileListRequest, opts ...grpc.CallOption) (*GetProfileListResponse, error) {
	out := new(GetProfileListResponse)
	err := c.cc.Invoke(ctx, "/gnsi.certz.v1.Certz/GetProfileList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certzClient) CanGenerateCSR(ctx context.Context, in *CanGenerateCSRRequest, opts ...grpc.CallOption) (*CanGenerateCSRResponse, error) {
	out := new(CanGenerateCSRResponse)
	err := c.cc.Invoke(ctx, "/gnsi.certz.v1.Certz/CanGenerateCSR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertzServer is the server API for Certz service.
// All implementations must embed UnimplementedCertzServer
// for forward compatibility
type CertzServer interface {
	Rotate(Certz_RotateServer) error
	AddProfile(context.Context, *AddProfileRequest) (*AddProfileResponse, error)
	DeleteProfile(context.Context, *DeleteProfileRequest) (*DeleteProfileResponse, error)
	GetProfileList(context.Context, *GetProfileListRequest) (*GetProfileListResponse, error)
	CanGenerateCSR(context.Context, *CanGenerateCSRRequest) (*CanGenerateCSRResponse, error)
	mustEmbedUnimplementedCertzServer()
}

// UnimplementedCertzServer must be embedded to have forward compatible implementations.
type UnimplementedCertzServer struct {
}

func (UnimplementedCertzServer) Rotate(Certz_RotateServer) error {
	return status.Errorf(codes.Unimplemented, "method Rotate not implemented")
}
func (UnimplementedCertzServer) AddProfile(context.Context, *AddProfileRequest) (*AddProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProfile not implemented")
}
func (UnimplementedCertzServer) DeleteProfile(context.Context, *DeleteProfileRequest) (*DeleteProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProfile not implemented")
}
func (UnimplementedCertzServer) GetProfileList(context.Context, *GetProfileListRequest) (*GetProfileListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileList not implemented")
}
func (UnimplementedCertzServer) CanGenerateCSR(context.Context, *CanGenerateCSRRequest) (*CanGenerateCSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanGenerateCSR not implemented")
}
func (UnimplementedCertzServer) mustEmbedUnimplementedCertzServer() {}

// UnsafeCertzServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertzServer will
// result in compilation errors.
type UnsafeCertzServer interface {
	mustEmbedUnimplementedCertzServer()
}

func RegisterCertzServer(s grpc.ServiceRegistrar, srv CertzServer) {
	s.RegisterService(&Certz_ServiceDesc, srv)
}

func _Certz_Rotate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CertzServer).Rotate(&certzRotateServer{stream})
}

type Certz_RotateServer interface {
	Send(*RotateCertificateResponse) error
	Recv() (*RotateCertificateRequest, error)
	grpc.ServerStream
}

type certzRotateServer struct {
	grpc.ServerStream
}

func (x *certzRotateServer) Send(m *RotateCertificateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *certzRotateServer) Recv() (*RotateCertificateRequest, error) {
	m := new(RotateCertificateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Certz_AddProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertzServer).AddProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnsi.certz.v1.Certz/AddProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertzServer).AddProfile(ctx, req.(*AddProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certz_DeleteProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertzServer).DeleteProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnsi.certz.v1.Certz/DeleteProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertzServer).DeleteProfile(ctx, req.(*DeleteProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certz_GetProfileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertzServer).GetProfileList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnsi.certz.v1.Certz/GetProfileList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertzServer).GetProfileList(ctx, req.(*GetProfileListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certz_CanGenerateCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CanGenerateCSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertzServer).CanGenerateCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnsi.certz.v1.Certz/CanGenerateCSR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertzServer).CanGenerateCSR(ctx, req.(*CanGenerateCSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Certz_ServiceDesc is the grpc.ServiceDesc for Certz service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Certz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnsi.certz.v1.Certz",
	HandlerType: (*CertzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProfile",
			Handler:    _Certz_AddProfile_Handler,
		},
		{
			MethodName: "DeleteProfile",
			Handler:    _Certz_DeleteProfile_Handler,
		},
		{
			MethodName: "GetProfileList",
			Handler:    _Certz_GetProfileList_Handler,
		},
		{
			MethodName: "CanGenerateCSR",
			Handler:    _Certz_CanGenerateCSR_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Rotate",
			Handler:       _Certz_Rotate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/openconfig/gnsi/certz/certz.proto",
}
