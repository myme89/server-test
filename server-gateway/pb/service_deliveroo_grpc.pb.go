// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: service_deliveroo.proto

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

// DeliverooClient is the client API for Deliveroo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliverooClient interface {
	GetData(ctx context.Context, in *DataInfoResquest, opts ...grpc.CallOption) (*DataInfoRespone, error)
	PostData(ctx context.Context, in *DataPostResqest, opts ...grpc.CallOption) (*DataPostRespone, error)
	UpdateData(ctx context.Context, in *DataUpdateResqest, opts ...grpc.CallOption) (*DataUpdateRespone, error)
	ExportData(ctx context.Context, in *ExportDataResquest, opts ...grpc.CallOption) (*ExportDataRespone, error)
	ImportData(ctx context.Context, in *ImportDataResquest, opts ...grpc.CallOption) (*ImportDataRespone, error)
	TestData1(ctx context.Context, in *TestResquest, opts ...grpc.CallOption) (*TestRespone, error)
	SignUp(ctx context.Context, in *SignUpResquest, opts ...grpc.CallOption) (*SignUpRespone, error)
	LogInAcc(ctx context.Context, in *SignInResquest, opts ...grpc.CallOption) (*SignInRespone, error)
	GetFileUploadInfo(ctx context.Context, in *FileUploadInfoResquest, opts ...grpc.CallOption) (*FileUploadInfoRespone, error)
	GetFileUploadShortInfo(ctx context.Context, in *FileUploadShortInfoResquest, opts ...grpc.CallOption) (*FileUploadShortInfoRespone, error)
	GetListServiceUpload(ctx context.Context, in *GetListServiceUploadResquest, opts ...grpc.CallOption) (*GetListServiceUploadRespone, error)
	GetListServiceProcess(ctx context.Context, in *GetListServiceProcessResquest, opts ...grpc.CallOption) (*GetListServiceProcessRespone, error)
	GetListFunctionProcess(ctx context.Context, in *GetListFunctionProcessResquest, opts ...grpc.CallOption) (*GetListFunctionProcessRespone, error)
}

type deliverooClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliverooClient(cc grpc.ClientConnInterface) DeliverooClient {
	return &deliverooClient{cc}
}

func (c *deliverooClient) GetData(ctx context.Context, in *DataInfoResquest, opts ...grpc.CallOption) (*DataInfoRespone, error) {
	out := new(DataInfoRespone)
	err := c.cc.Invoke(ctx, "/pb.Deliveroo/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverooClient) PostData(ctx context.Context, in *DataPostResqest, opts ...grpc.CallOption) (*DataPostRespone, error) {
	out := new(DataPostRespone)
	err := c.cc.Invoke(ctx, "/pb.Deliveroo/PostData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverooClient) UpdateData(ctx context.Context, in *DataUpdateResqest, opts ...grpc.CallOption) (*DataUpdateRespone, error) {
	out := new(DataUpdateRespone)
	err := c.cc.Invoke(ctx, "/pb.Deliveroo/UpdateData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverooClient) ExportData(ctx context.Context, in *ExportDataResquest, opts ...grpc.CallOption) (*ExportDataRespone, error) {
	out := new(ExportDataRespone)
	err := c.cc.Invoke(ctx, "/pb.Deliveroo/ExportData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverooClient) ImportData(ctx context.Context, in *ImportDataResquest, opts ...grpc.CallOption) (*ImportDataRespone, error) {
	out := new(ImportDataRespone)
	err := c.cc.Invoke(ctx, "/pb.Deliveroo/ImportData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverooClient) TestData1(ctx context.Context, in *TestResquest, opts ...grpc.CallOption) (*TestRespone, error) {
	out := new(TestRespone)
	err := c.cc.Invoke(ctx, "/pb.Deliveroo/TestData1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverooClient) SignUp(ctx context.Context, in *SignUpResquest, opts ...grpc.CallOption) (*SignUpRespone, error) {
	out := new(SignUpRespone)
	err := c.cc.Invoke(ctx, "/pb.Deliveroo/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverooClient) LogInAcc(ctx context.Context, in *SignInResquest, opts ...grpc.CallOption) (*SignInRespone, error) {
	out := new(SignInRespone)
	err := c.cc.Invoke(ctx, "/pb.Deliveroo/LogInAcc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverooClient) GetFileUploadInfo(ctx context.Context, in *FileUploadInfoResquest, opts ...grpc.CallOption) (*FileUploadInfoRespone, error) {
	out := new(FileUploadInfoRespone)
	err := c.cc.Invoke(ctx, "/pb.Deliveroo/GetFileUploadInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverooClient) GetFileUploadShortInfo(ctx context.Context, in *FileUploadShortInfoResquest, opts ...grpc.CallOption) (*FileUploadShortInfoRespone, error) {
	out := new(FileUploadShortInfoRespone)
	err := c.cc.Invoke(ctx, "/pb.Deliveroo/GetFileUploadShortInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverooClient) GetListServiceUpload(ctx context.Context, in *GetListServiceUploadResquest, opts ...grpc.CallOption) (*GetListServiceUploadRespone, error) {
	out := new(GetListServiceUploadRespone)
	err := c.cc.Invoke(ctx, "/pb.Deliveroo/GetListServiceUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverooClient) GetListServiceProcess(ctx context.Context, in *GetListServiceProcessResquest, opts ...grpc.CallOption) (*GetListServiceProcessRespone, error) {
	out := new(GetListServiceProcessRespone)
	err := c.cc.Invoke(ctx, "/pb.Deliveroo/GetListServiceProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliverooClient) GetListFunctionProcess(ctx context.Context, in *GetListFunctionProcessResquest, opts ...grpc.CallOption) (*GetListFunctionProcessRespone, error) {
	out := new(GetListFunctionProcessRespone)
	err := c.cc.Invoke(ctx, "/pb.Deliveroo/GetListFunctionProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliverooServer is the server API for Deliveroo service.
// All implementations must embed UnimplementedDeliverooServer
// for forward compatibility
type DeliverooServer interface {
	GetData(context.Context, *DataInfoResquest) (*DataInfoRespone, error)
	PostData(context.Context, *DataPostResqest) (*DataPostRespone, error)
	UpdateData(context.Context, *DataUpdateResqest) (*DataUpdateRespone, error)
	ExportData(context.Context, *ExportDataResquest) (*ExportDataRespone, error)
	ImportData(context.Context, *ImportDataResquest) (*ImportDataRespone, error)
	TestData1(context.Context, *TestResquest) (*TestRespone, error)
	SignUp(context.Context, *SignUpResquest) (*SignUpRespone, error)
	LogInAcc(context.Context, *SignInResquest) (*SignInRespone, error)
	GetFileUploadInfo(context.Context, *FileUploadInfoResquest) (*FileUploadInfoRespone, error)
	GetFileUploadShortInfo(context.Context, *FileUploadShortInfoResquest) (*FileUploadShortInfoRespone, error)
	GetListServiceUpload(context.Context, *GetListServiceUploadResquest) (*GetListServiceUploadRespone, error)
	GetListServiceProcess(context.Context, *GetListServiceProcessResquest) (*GetListServiceProcessRespone, error)
	GetListFunctionProcess(context.Context, *GetListFunctionProcessResquest) (*GetListFunctionProcessRespone, error)
	mustEmbedUnimplementedDeliverooServer()
}

// UnimplementedDeliverooServer must be embedded to have forward compatible implementations.
type UnimplementedDeliverooServer struct {
}

func (UnimplementedDeliverooServer) GetData(context.Context, *DataInfoResquest) (*DataInfoRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedDeliverooServer) PostData(context.Context, *DataPostResqest) (*DataPostRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostData not implemented")
}
func (UnimplementedDeliverooServer) UpdateData(context.Context, *DataUpdateResqest) (*DataUpdateRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateData not implemented")
}
func (UnimplementedDeliverooServer) ExportData(context.Context, *ExportDataResquest) (*ExportDataRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportData not implemented")
}
func (UnimplementedDeliverooServer) ImportData(context.Context, *ImportDataResquest) (*ImportDataRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportData not implemented")
}
func (UnimplementedDeliverooServer) TestData1(context.Context, *TestResquest) (*TestRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestData1 not implemented")
}
func (UnimplementedDeliverooServer) SignUp(context.Context, *SignUpResquest) (*SignUpRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedDeliverooServer) LogInAcc(context.Context, *SignInResquest) (*SignInRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogInAcc not implemented")
}
func (UnimplementedDeliverooServer) GetFileUploadInfo(context.Context, *FileUploadInfoResquest) (*FileUploadInfoRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileUploadInfo not implemented")
}
func (UnimplementedDeliverooServer) GetFileUploadShortInfo(context.Context, *FileUploadShortInfoResquest) (*FileUploadShortInfoRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileUploadShortInfo not implemented")
}
func (UnimplementedDeliverooServer) GetListServiceUpload(context.Context, *GetListServiceUploadResquest) (*GetListServiceUploadRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListServiceUpload not implemented")
}
func (UnimplementedDeliverooServer) GetListServiceProcess(context.Context, *GetListServiceProcessResquest) (*GetListServiceProcessRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListServiceProcess not implemented")
}
func (UnimplementedDeliverooServer) GetListFunctionProcess(context.Context, *GetListFunctionProcessResquest) (*GetListFunctionProcessRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListFunctionProcess not implemented")
}
func (UnimplementedDeliverooServer) mustEmbedUnimplementedDeliverooServer() {}

// UnsafeDeliverooServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeliverooServer will
// result in compilation errors.
type UnsafeDeliverooServer interface {
	mustEmbedUnimplementedDeliverooServer()
}

func RegisterDeliverooServer(s grpc.ServiceRegistrar, srv DeliverooServer) {
	s.RegisterService(&Deliveroo_ServiceDesc, srv)
}

func _Deliveroo_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataInfoResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverooServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Deliveroo/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverooServer).GetData(ctx, req.(*DataInfoResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliveroo_PostData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataPostResqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverooServer).PostData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Deliveroo/PostData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverooServer).PostData(ctx, req.(*DataPostResqest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliveroo_UpdateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataUpdateResqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverooServer).UpdateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Deliveroo/UpdateData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverooServer).UpdateData(ctx, req.(*DataUpdateResqest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliveroo_ExportData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportDataResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverooServer).ExportData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Deliveroo/ExportData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverooServer).ExportData(ctx, req.(*ExportDataResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliveroo_ImportData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportDataResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverooServer).ImportData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Deliveroo/ImportData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverooServer).ImportData(ctx, req.(*ImportDataResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliveroo_TestData1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverooServer).TestData1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Deliveroo/TestData1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverooServer).TestData1(ctx, req.(*TestResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliveroo_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverooServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Deliveroo/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverooServer).SignUp(ctx, req.(*SignUpResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliveroo_LogInAcc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverooServer).LogInAcc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Deliveroo/LogInAcc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverooServer).LogInAcc(ctx, req.(*SignInResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliveroo_GetFileUploadInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileUploadInfoResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverooServer).GetFileUploadInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Deliveroo/GetFileUploadInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverooServer).GetFileUploadInfo(ctx, req.(*FileUploadInfoResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliveroo_GetFileUploadShortInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileUploadShortInfoResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverooServer).GetFileUploadShortInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Deliveroo/GetFileUploadShortInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverooServer).GetFileUploadShortInfo(ctx, req.(*FileUploadShortInfoResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliveroo_GetListServiceUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListServiceUploadResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverooServer).GetListServiceUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Deliveroo/GetListServiceUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverooServer).GetListServiceUpload(ctx, req.(*GetListServiceUploadResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliveroo_GetListServiceProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListServiceProcessResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverooServer).GetListServiceProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Deliveroo/GetListServiceProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverooServer).GetListServiceProcess(ctx, req.(*GetListServiceProcessResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deliveroo_GetListFunctionProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListFunctionProcessResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverooServer).GetListFunctionProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Deliveroo/GetListFunctionProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverooServer).GetListFunctionProcess(ctx, req.(*GetListFunctionProcessResquest))
	}
	return interceptor(ctx, in, info, handler)
}

// Deliveroo_ServiceDesc is the grpc.ServiceDesc for Deliveroo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Deliveroo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Deliveroo",
	HandlerType: (*DeliverooServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _Deliveroo_GetData_Handler,
		},
		{
			MethodName: "PostData",
			Handler:    _Deliveroo_PostData_Handler,
		},
		{
			MethodName: "UpdateData",
			Handler:    _Deliveroo_UpdateData_Handler,
		},
		{
			MethodName: "ExportData",
			Handler:    _Deliveroo_ExportData_Handler,
		},
		{
			MethodName: "ImportData",
			Handler:    _Deliveroo_ImportData_Handler,
		},
		{
			MethodName: "TestData1",
			Handler:    _Deliveroo_TestData1_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _Deliveroo_SignUp_Handler,
		},
		{
			MethodName: "LogInAcc",
			Handler:    _Deliveroo_LogInAcc_Handler,
		},
		{
			MethodName: "GetFileUploadInfo",
			Handler:    _Deliveroo_GetFileUploadInfo_Handler,
		},
		{
			MethodName: "GetFileUploadShortInfo",
			Handler:    _Deliveroo_GetFileUploadShortInfo_Handler,
		},
		{
			MethodName: "GetListServiceUpload",
			Handler:    _Deliveroo_GetListServiceUpload_Handler,
		},
		{
			MethodName: "GetListServiceProcess",
			Handler:    _Deliveroo_GetListServiceProcess_Handler,
		},
		{
			MethodName: "GetListFunctionProcess",
			Handler:    _Deliveroo_GetListFunctionProcess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_deliveroo.proto",
}
