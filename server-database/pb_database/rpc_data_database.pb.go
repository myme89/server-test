// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: rpc_data_database.proto

package pb_database

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

type SignUpAccResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userinfo *UserAccInfo `protobuf:"bytes,1,opt,name=userinfo,proto3" json:"userinfo,omitempty"`
}

func (x *SignUpAccResquest) Reset() {
	*x = SignUpAccResquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpAccResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpAccResquest) ProtoMessage() {}

func (x *SignUpAccResquest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpAccResquest.ProtoReflect.Descriptor instead.
func (*SignUpAccResquest) Descriptor() ([]byte, []int) {
	return file_rpc_data_database_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpAccResquest) GetUserinfo() *UserAccInfo {
	if x != nil {
		return x.Userinfo
	}
	return nil
}

type SignUpAccRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Noti string `protobuf:"bytes,1,opt,name=noti,proto3" json:"noti,omitempty"`
}

func (x *SignUpAccRespone) Reset() {
	*x = SignUpAccRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpAccRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpAccRespone) ProtoMessage() {}

func (x *SignUpAccRespone) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpAccRespone.ProtoReflect.Descriptor instead.
func (*SignUpAccRespone) Descriptor() ([]byte, []int) {
	return file_rpc_data_database_proto_rawDescGZIP(), []int{1}
}

func (x *SignUpAccRespone) GetNoti() string {
	if x != nil {
		return x.Noti
	}
	return ""
}

type LoginAccResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userinfo *UserAccInfo `protobuf:"bytes,1,opt,name=userinfo,proto3" json:"userinfo,omitempty"`
}

func (x *LoginAccResquest) Reset() {
	*x = LoginAccResquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginAccResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginAccResquest) ProtoMessage() {}

func (x *LoginAccResquest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginAccResquest.ProtoReflect.Descriptor instead.
func (*LoginAccResquest) Descriptor() ([]byte, []int) {
	return file_rpc_data_database_proto_rawDescGZIP(), []int{2}
}

func (x *LoginAccResquest) GetUserinfo() *UserAccInfo {
	if x != nil {
		return x.Userinfo
	}
	return nil
}

type LoginAccRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userinfo *UserAccInfo `protobuf:"bytes,1,opt,name=userinfo,proto3" json:"userinfo,omitempty"`
}

func (x *LoginAccRespone) Reset() {
	*x = LoginAccRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginAccRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginAccRespone) ProtoMessage() {}

func (x *LoginAccRespone) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginAccRespone.ProtoReflect.Descriptor instead.
func (*LoginAccRespone) Descriptor() ([]byte, []int) {
	return file_rpc_data_database_proto_rawDescGZIP(), []int{3}
}

func (x *LoginAccRespone) GetUserinfo() *UserAccInfo {
	if x != nil {
		return x.Userinfo
	}
	return nil
}

type UploadFileResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileUploadInfo *FileUploadInfo `protobuf:"bytes,1,opt,name=fileUploadInfo,proto3" json:"fileUploadInfo,omitempty"`
}

func (x *UploadFileResquest) Reset() {
	*x = UploadFileResquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResquest) ProtoMessage() {}

func (x *UploadFileResquest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResquest.ProtoReflect.Descriptor instead.
func (*UploadFileResquest) Descriptor() ([]byte, []int) {
	return file_rpc_data_database_proto_rawDescGZIP(), []int{4}
}

func (x *UploadFileResquest) GetFileUploadInfo() *FileUploadInfo {
	if x != nil {
		return x.FileUploadInfo
	}
	return nil
}

type UploadFileRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Noti string `protobuf:"bytes,1,opt,name=noti,proto3" json:"noti,omitempty"`
}

func (x *UploadFileRespone) Reset() {
	*x = UploadFileRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_database_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRespone) ProtoMessage() {}

func (x *UploadFileRespone) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_database_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRespone.ProtoReflect.Descriptor instead.
func (*UploadFileRespone) Descriptor() ([]byte, []int) {
	return file_rpc_data_database_proto_rawDescGZIP(), []int{5}
}

func (x *UploadFileRespone) GetNoti() string {
	if x != nil {
		return x.Noti
	}
	return ""
}

var File_rpc_data_database_proto protoreflect.FileDescriptor

var file_rpc_data_database_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x62, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x11, 0x53, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x41, 0x63, 0x63, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x70, 0x62, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x41, 0x63,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x69, 0x22, 0x48, 0x0a, 0x10,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x34, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x47, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41,
	0x63, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0x59, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x70, 0x62, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x27, 0x0a, 0x11, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x6f, 0x74, 0x69, 0x42, 0x29, 0x5a, 0x27, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x74, 0x65,
	0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x70, 0x62, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_data_database_proto_rawDescOnce sync.Once
	file_rpc_data_database_proto_rawDescData = file_rpc_data_database_proto_rawDesc
)

func file_rpc_data_database_proto_rawDescGZIP() []byte {
	file_rpc_data_database_proto_rawDescOnce.Do(func() {
		file_rpc_data_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_data_database_proto_rawDescData)
	})
	return file_rpc_data_database_proto_rawDescData
}

var file_rpc_data_database_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rpc_data_database_proto_goTypes = []interface{}{
	(*SignUpAccResquest)(nil),  // 0: pb_database.SignUpAccResquest
	(*SignUpAccRespone)(nil),   // 1: pb_database.SignUpAccRespone
	(*LoginAccResquest)(nil),   // 2: pb_database.LoginAccResquest
	(*LoginAccRespone)(nil),    // 3: pb_database.LoginAccRespone
	(*UploadFileResquest)(nil), // 4: pb_database.UploadFileResquest
	(*UploadFileRespone)(nil),  // 5: pb_database.UploadFileRespone
	(*UserAccInfo)(nil),        // 6: pb_database.UserAccInfo
	(*FileUploadInfo)(nil),     // 7: pb_database.FileUploadInfo
}
var file_rpc_data_database_proto_depIdxs = []int32{
	6, // 0: pb_database.SignUpAccResquest.userinfo:type_name -> pb_database.UserAccInfo
	6, // 1: pb_database.LoginAccResquest.userinfo:type_name -> pb_database.UserAccInfo
	6, // 2: pb_database.LoginAccRespone.userinfo:type_name -> pb_database.UserAccInfo
	7, // 3: pb_database.UploadFileResquest.fileUploadInfo:type_name -> pb_database.FileUploadInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_data_database_proto_init() }
func file_rpc_data_database_proto_init() {
	if File_rpc_data_database_proto != nil {
		return
	}
	file_user_acc_proto_init()
	file_file_upload_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_data_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpAccResquest); i {
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
		file_rpc_data_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpAccRespone); i {
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
		file_rpc_data_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginAccResquest); i {
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
		file_rpc_data_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginAccRespone); i {
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
		file_rpc_data_database_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileResquest); i {
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
		file_rpc_data_database_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRespone); i {
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
			RawDescriptor: file_rpc_data_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_data_database_proto_goTypes,
		DependencyIndexes: file_rpc_data_database_proto_depIdxs,
		MessageInfos:      file_rpc_data_database_proto_msgTypes,
	}.Build()
	File_rpc_data_database_proto = out.File
	file_rpc_data_database_proto_rawDesc = nil
	file_rpc_data_database_proto_goTypes = nil
	file_rpc_data_database_proto_depIdxs = nil
}
