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

// DeliverooServer is the server API for Deliveroo service.
// All implementations must embed UnimplementedDeliverooServer
// for forward compatibility
type DeliverooServer interface {
	GetData(context.Context, *DataInfoResquest) (*DataInfoRespone, error)
	PostData(context.Context, *DataPostResqest) (*DataPostRespone, error)
	UpdateData(context.Context, *DataUpdateResqest) (*DataUpdateRespone, error)
	ExportData(context.Context, *ExportDataResquest) (*ExportDataRespone, error)
	ImportData(context.Context, *ImportDataResquest) (*ImportDataRespone, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_deliveroo.proto",
}
