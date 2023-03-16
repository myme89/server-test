// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: rpc_data_storage.proto

package pb_storage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DataInfoTestResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataTest string `protobuf:"bytes,1,opt,name=dataTest,proto3" json:"dataTest,omitempty"`
}

func (x *DataInfoTestResquest) Reset() {
	*x = DataInfoTestResquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataInfoTestResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataInfoTestResquest) ProtoMessage() {}

func (x *DataInfoTestResquest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataInfoTestResquest.ProtoReflect.Descriptor instead.
func (*DataInfoTestResquest) Descriptor() ([]byte, []int) {
	return file_rpc_data_storage_proto_rawDescGZIP(), []int{0}
}

func (x *DataInfoTestResquest) GetDataTest() string {
	if x != nil {
		return x.DataTest
	}
	return ""
}

type DataInfoTestRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataResp string `protobuf:"bytes,1,opt,name=dataResp,proto3" json:"dataResp,omitempty"`
}

func (x *DataInfoTestRespone) Reset() {
	*x = DataInfoTestRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataInfoTestRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataInfoTestRespone) ProtoMessage() {}

func (x *DataInfoTestRespone) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataInfoTestRespone.ProtoReflect.Descriptor instead.
func (*DataInfoTestRespone) Descriptor() ([]byte, []int) {
	return file_rpc_data_storage_proto_rawDescGZIP(), []int{1}
}

func (x *DataInfoTestRespone) GetDataResp() string {
	if x != nil {
		return x.DataResp
	}
	return ""
}

var File_rpc_data_storage_proto protoreflect.FileDescriptor

var file_rpc_data_storage_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x62, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x22, 0x32, 0x0a, 0x14, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x54, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x54, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x13, 0x44, 0x61, 0x74, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x42, 0x27, 0x5a, 0x25, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x62, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_data_storage_proto_rawDescOnce sync.Once
	file_rpc_data_storage_proto_rawDescData = file_rpc_data_storage_proto_rawDesc
)

func file_rpc_data_storage_proto_rawDescGZIP() []byte {
	file_rpc_data_storage_proto_rawDescOnce.Do(func() {
		file_rpc_data_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_data_storage_proto_rawDescData)
	})
	return file_rpc_data_storage_proto_rawDescData
}

var file_rpc_data_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_data_storage_proto_goTypes = []interface{}{
	(*DataInfoTestResquest)(nil), // 0: pb_storage.DataInfoTestResquest
	(*DataInfoTestRespone)(nil),  // 1: pb_storage.DataInfoTestRespone
}
var file_rpc_data_storage_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_data_storage_proto_init() }
func file_rpc_data_storage_proto_init() {
	if File_rpc_data_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_data_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataInfoTestResquest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_data_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataInfoTestRespone); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_data_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_data_storage_proto_goTypes,
		DependencyIndexes: file_rpc_data_storage_proto_depIdxs,
		MessageInfos:      file_rpc_data_storage_proto_msgTypes,
	}.Build()
	File_rpc_data_storage_proto = out.File
	file_rpc_data_storage_proto_rawDesc = nil
	file_rpc_data_storage_proto_goTypes = nil
	file_rpc_data_storage_proto_depIdxs = nil
}
