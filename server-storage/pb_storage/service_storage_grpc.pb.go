// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: service_storage.proto

package pb_storage

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

// StogareClient is the client API for Stogare service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StogareClient interface {
	TestData(ctx context.Context, in *DataInfoTestResquest, opts ...grpc.CallOption) (*DataInfoTestRespone, error)
	UploadFile(ctx context.Context, in *FileInfoResquest, opts ...grpc.CallOption) (*FileInfoRespone, error)
	GetListFileUpload(ctx context.Context, in *GetListFileUploadResquest, opts ...grpc.CallOption) (*GetListFileUploadRespone, error)
	DownloafFile(ctx context.Context, in *DownloadFileResquest, opts ...grpc.CallOption) (*DownloadFileRespone, error)
	UpdateStatusUploadFile(ctx context.Context, in *UpdateStatusResquest, opts ...grpc.CallOption) (*UpdateStatusRespone, error)
	GetShortInfoFileUpload(ctx context.Context, in *GetShortInfoFileUploadResquest, opts ...grpc.CallOption) (*GetShortInfoFileUploadRespone, error)
}

type stogareClient struct {
	cc grpc.ClientConnInterface
}

func NewStogareClient(cc grpc.ClientConnInterface) StogareClient {
	return &stogareClient{cc}
}

func (c *stogareClient) TestData(ctx context.Context, in *DataInfoTestResquest, opts ...grpc.CallOption) (*DataInfoTestRespone, error) {
	out := new(DataInfoTestRespone)
	err := c.cc.Invoke(ctx, "/pb_storage.Stogare/TestData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stogareClient) UploadFile(ctx context.Context, in *FileInfoResquest, opts ...grpc.CallOption) (*FileInfoRespone, error) {
	out := new(FileInfoRespone)
	err := c.cc.Invoke(ctx, "/pb_storage.Stogare/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stogareClient) GetListFileUpload(ctx context.Context, in *GetListFileUploadResquest, opts ...grpc.CallOption) (*GetListFileUploadRespone, error) {
	out := new(GetListFileUploadRespone)
	err := c.cc.Invoke(ctx, "/pb_storage.Stogare/GetListFileUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stogareClient) DownloafFile(ctx context.Context, in *DownloadFileResquest, opts ...grpc.CallOption) (*DownloadFileRespone, error) {
	out := new(DownloadFileRespone)
	err := c.cc.Invoke(ctx, "/pb_storage.Stogare/DownloafFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stogareClient) UpdateStatusUploadFile(ctx context.Context, in *UpdateStatusResquest, opts ...grpc.CallOption) (*UpdateStatusRespone, error) {
	out := new(UpdateStatusRespone)
	err := c.cc.Invoke(ctx, "/pb_storage.Stogare/UpdateStatusUploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stogareClient) GetShortInfoFileUpload(ctx context.Context, in *GetShortInfoFileUploadResquest, opts ...grpc.CallOption) (*GetShortInfoFileUploadRespone, error) {
	out := new(GetShortInfoFileUploadRespone)
	err := c.cc.Invoke(ctx, "/pb_storage.Stogare/GetShortInfoFileUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StogareServer is the server API for Stogare service.
// All implementations must embed UnimplementedStogareServer
// for forward compatibility
type StogareServer interface {
	TestData(context.Context, *DataInfoTestResquest) (*DataInfoTestRespone, error)
	UploadFile(context.Context, *FileInfoResquest) (*FileInfoRespone, error)
	GetListFileUpload(context.Context, *GetListFileUploadResquest) (*GetListFileUploadRespone, error)
	DownloafFile(context.Context, *DownloadFileResquest) (*DownloadFileRespone, error)
	UpdateStatusUploadFile(context.Context, *UpdateStatusResquest) (*UpdateStatusRespone, error)
	GetShortInfoFileUpload(context.Context, *GetShortInfoFileUploadResquest) (*GetShortInfoFileUploadRespone, error)
	mustEmbedUnimplementedStogareServer()
}

// UnimplementedStogareServer must be embedded to have forward compatible implementations.
type UnimplementedStogareServer struct {
}

func (UnimplementedStogareServer) TestData(context.Context, *DataInfoTestResquest) (*DataInfoTestRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestData not implemented")
}
func (UnimplementedStogareServer) UploadFile(context.Context, *FileInfoResquest) (*FileInfoRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedStogareServer) GetListFileUpload(context.Context, *GetListFileUploadResquest) (*GetListFileUploadRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListFileUpload not implemented")
}
func (UnimplementedStogareServer) DownloafFile(context.Context, *DownloadFileResquest) (*DownloadFileRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloafFile not implemented")
}
func (UnimplementedStogareServer) UpdateStatusUploadFile(context.Context, *UpdateStatusResquest) (*UpdateStatusRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatusUploadFile not implemented")
}
func (UnimplementedStogareServer) GetShortInfoFileUpload(context.Context, *GetShortInfoFileUploadResquest) (*GetShortInfoFileUploadRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortInfoFileUpload not implemented")
}
func (UnimplementedStogareServer) mustEmbedUnimplementedStogareServer() {}

// UnsafeStogareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StogareServer will
// result in compilation errors.
type UnsafeStogareServer interface {
	mustEmbedUnimplementedStogareServer()
}

func RegisterStogareServer(s grpc.ServiceRegistrar, srv StogareServer) {
	s.RegisterService(&Stogare_ServiceDesc, srv)
}

func _Stogare_TestData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataInfoTestResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StogareServer).TestData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_storage.Stogare/TestData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StogareServer).TestData(ctx, req.(*DataInfoTestResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stogare_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileInfoResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StogareServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_storage.Stogare/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StogareServer).UploadFile(ctx, req.(*FileInfoResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stogare_GetListFileUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListFileUploadResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StogareServer).GetListFileUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_storage.Stogare/GetListFileUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StogareServer).GetListFileUpload(ctx, req.(*GetListFileUploadResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stogare_DownloafFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadFileResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StogareServer).DownloafFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_storage.Stogare/DownloafFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StogareServer).DownloafFile(ctx, req.(*DownloadFileResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stogare_UpdateStatusUploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StogareServer).UpdateStatusUploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_storage.Stogare/UpdateStatusUploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StogareServer).UpdateStatusUploadFile(ctx, req.(*UpdateStatusResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stogare_GetShortInfoFileUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShortInfoFileUploadResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StogareServer).GetShortInfoFileUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_storage.Stogare/GetShortInfoFileUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StogareServer).GetShortInfoFileUpload(ctx, req.(*GetShortInfoFileUploadResquest))
	}
	return interceptor(ctx, in, info, handler)
}

// Stogare_ServiceDesc is the grpc.ServiceDesc for Stogare service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stogare_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_storage.Stogare",
	HandlerType: (*StogareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestData",
			Handler:    _Stogare_TestData_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _Stogare_UploadFile_Handler,
		},
		{
			MethodName: "GetListFileUpload",
			Handler:    _Stogare_GetListFileUpload_Handler,
		},
		{
			MethodName: "DownloafFile",
			Handler:    _Stogare_DownloafFile_Handler,
		},
		{
			MethodName: "UpdateStatusUploadFile",
			Handler:    _Stogare_UpdateStatusUploadFile_Handler,
		},
		{
			MethodName: "GetShortInfoFileUpload",
			Handler:    _Stogare_GetShortInfoFileUpload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_storage.proto",
}
