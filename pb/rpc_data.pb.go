// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: rpc_data.proto

package pb

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

type DataInfoResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DataInfoResquest) Reset() {
	*x = DataInfoResquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataInfoResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataInfoResquest) ProtoMessage() {}

func (x *DataInfoResquest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataInfoResquest.ProtoReflect.Descriptor instead.
func (*DataInfoResquest) Descriptor() ([]byte, []int) {
	return file_rpc_data_proto_rawDescGZIP(), []int{0}
}

type DataInfoRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*DataInfo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *DataInfoRespone) Reset() {
	*x = DataInfoRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataInfoRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataInfoRespone) ProtoMessage() {}

func (x *DataInfoRespone) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataInfoRespone.ProtoReflect.Descriptor instead.
func (*DataInfoRespone) Descriptor() ([]byte, []int) {
	return file_rpc_data_proto_rawDescGZIP(), []int{1}
}

func (x *DataInfoRespone) GetData() []*DataInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type DataPostResqest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Fullname string `protobuf:"bytes,3,opt,name=fullname,proto3" json:"fullname,omitempty"`
}

func (x *DataPostResqest) Reset() {
	*x = DataPostResqest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPostResqest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPostResqest) ProtoMessage() {}

func (x *DataPostResqest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPostResqest.ProtoReflect.Descriptor instead.
func (*DataPostResqest) Descriptor() ([]byte, []int) {
	return file_rpc_data_proto_rawDescGZIP(), []int{2}
}

func (x *DataPostResqest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DataPostResqest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataPostResqest) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

type DataPostRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notice string `protobuf:"bytes,1,opt,name=notice,proto3" json:"notice,omitempty"`
}

func (x *DataPostRespone) Reset() {
	*x = DataPostRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPostRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPostRespone) ProtoMessage() {}

func (x *DataPostRespone) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPostRespone.ProtoReflect.Descriptor instead.
func (*DataPostRespone) Descriptor() ([]byte, []int) {
	return file_rpc_data_proto_rawDescGZIP(), []int{3}
}

func (x *DataPostRespone) GetNotice() string {
	if x != nil {
		return x.Notice
	}
	return ""
}

type DataUpdateResqest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oldname     string `protobuf:"bytes,1,opt,name=oldname,proto3" json:"oldname,omitempty"`
	Newname     string `protobuf:"bytes,2,opt,name=newname,proto3" json:"newname,omitempty"`
	Newfullname string `protobuf:"bytes,3,opt,name=newfullname,proto3" json:"newfullname,omitempty"`
}

func (x *DataUpdateResqest) Reset() {
	*x = DataUpdateResqest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataUpdateResqest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataUpdateResqest) ProtoMessage() {}

func (x *DataUpdateResqest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataUpdateResqest.ProtoReflect.Descriptor instead.
func (*DataUpdateResqest) Descriptor() ([]byte, []int) {
	return file_rpc_data_proto_rawDescGZIP(), []int{4}
}

func (x *DataUpdateResqest) GetOldname() string {
	if x != nil {
		return x.Oldname
	}
	return ""
}

func (x *DataUpdateResqest) GetNewname() string {
	if x != nil {
		return x.Newname
	}
	return ""
}

func (x *DataUpdateResqest) GetNewfullname() string {
	if x != nil {
		return x.Newfullname
	}
	return ""
}

type DataUpdateRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notice string `protobuf:"bytes,1,opt,name=notice,proto3" json:"notice,omitempty"`
}

func (x *DataUpdateRespone) Reset() {
	*x = DataUpdateRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataUpdateRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataUpdateRespone) ProtoMessage() {}

func (x *DataUpdateRespone) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataUpdateRespone.ProtoReflect.Descriptor instead.
func (*DataUpdateRespone) Descriptor() ([]byte, []int) {
	return file_rpc_data_proto_rawDescGZIP(), []int{5}
}

func (x *DataUpdateRespone) GetNotice() string {
	if x != nil {
		return x.Notice
	}
	return ""
}

type ExportDataResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExportDataResquest) Reset() {
	*x = ExportDataResquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportDataResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportDataResquest) ProtoMessage() {}

func (x *ExportDataResquest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportDataResquest.ProtoReflect.Descriptor instead.
func (*ExportDataResquest) Descriptor() ([]byte, []int) {
	return file_rpc_data_proto_rawDescGZIP(), []int{6}
}

type ExportDataRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PathExport string `protobuf:"bytes,1,opt,name=pathExport,proto3" json:"pathExport,omitempty"`
}

func (x *ExportDataRespone) Reset() {
	*x = ExportDataRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportDataRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportDataRespone) ProtoMessage() {}

func (x *ExportDataRespone) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportDataRespone.ProtoReflect.Descriptor instead.
func (*ExportDataRespone) Descriptor() ([]byte, []int) {
	return file_rpc_data_proto_rawDescGZIP(), []int{7}
}

func (x *ExportDataRespone) GetPathExport() string {
	if x != nil {
		return x.PathExport
	}
	return ""
}

type ImportDataResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ImportDataResquest) Reset() {
	*x = ImportDataResquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportDataResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportDataResquest) ProtoMessage() {}

func (x *ImportDataResquest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportDataResquest.ProtoReflect.Descriptor instead.
func (*ImportDataResquest) Descriptor() ([]byte, []int) {
	return file_rpc_data_proto_rawDescGZIP(), []int{8}
}

func (x *ImportDataResquest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ImportDataResquest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ImportDataRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notice string `protobuf:"bytes,1,opt,name=notice,proto3" json:"notice,omitempty"`
}

func (x *ImportDataRespone) Reset() {
	*x = ImportDataRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportDataRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportDataRespone) ProtoMessage() {}

func (x *ImportDataRespone) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportDataRespone.ProtoReflect.Descriptor instead.
func (*ImportDataRespone) Descriptor() ([]byte, []int) {
	return file_rpc_data_proto_rawDescGZIP(), []int{9}
}

func (x *ImportDataRespone) GetNotice() string {
	if x != nil {
		return x.Notice
	}
	return ""
}

type TestResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Testdata string `protobuf:"bytes,1,opt,name=testdata,proto3" json:"testdata,omitempty"`
}

func (x *TestResquest) Reset() {
	*x = TestResquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResquest) ProtoMessage() {}

func (x *TestResquest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResquest.ProtoReflect.Descriptor instead.
func (*TestResquest) Descriptor() ([]byte, []int) {
	return file_rpc_data_proto_rawDescGZIP(), []int{10}
}

func (x *TestResquest) GetTestdata() string {
	if x != nil {
		return x.Testdata
	}
	return ""
}

type TestRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notice string `protobuf:"bytes,1,opt,name=notice,proto3" json:"notice,omitempty"`
}

func (x *TestRespone) Reset() {
	*x = TestRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRespone) ProtoMessage() {}

func (x *TestRespone) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRespone.ProtoReflect.Descriptor instead.
func (*TestRespone) Descriptor() ([]byte, []int) {
	return file_rpc_data_proto_rawDescGZIP(), []int{11}
}

func (x *TestRespone) GetNotice() string {
	if x != nil {
		return x.Notice
	}
	return ""
}

type SignUpResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignUpResquest) Reset() {
	*x = SignUpResquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResquest) ProtoMessage() {}

func (x *SignUpResquest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResquest.ProtoReflect.Descriptor instead.
func (*SignUpResquest) Descriptor() ([]byte, []int) {
	return file_rpc_data_proto_rawDescGZIP(), []int{12}
}

type SignUpRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *SignUpRespone) Reset() {
	*x = SignUpRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_data_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRespone) ProtoMessage() {}

func (x *SignUpRespone) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_data_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRespone.ProtoReflect.Descriptor instead.
func (*SignUpRespone) Descriptor() ([]byte, []int) {
	return file_rpc_data_proto_rawDescGZIP(), []int{13}
}

func (x *SignUpRespone) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_rpc_data_proto protoreflect.FileDescriptor

var file_rpc_data_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x12, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x51, 0x0a, 0x0f, 0x44, 0x61, 0x74,
	0x61, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x71, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0f,
	0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x22, 0x69, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x71, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x6c, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x6c, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x22,
	0x14, 0x0a, 0x12, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x11, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61,
	0x74, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x61, 0x74, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x3c, 0x0a, 0x12, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x11, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x22, 0x2a, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x25, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x0d, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x10, 0x5a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_data_proto_rawDescOnce sync.Once
	file_rpc_data_proto_rawDescData = file_rpc_data_proto_rawDesc
)

func file_rpc_data_proto_rawDescGZIP() []byte {
	file_rpc_data_proto_rawDescOnce.Do(func() {
		file_rpc_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_data_proto_rawDescData)
	})
	return file_rpc_data_proto_rawDescData
}

var file_rpc_data_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_rpc_data_proto_goTypes = []interface{}{
	(*DataInfoResquest)(nil),   // 0: pb.DataInfoResquest
	(*DataInfoRespone)(nil),    // 1: pb.DataInfoRespone
	(*DataPostResqest)(nil),    // 2: pb.DataPostResqest
	(*DataPostRespone)(nil),    // 3: pb.DataPostRespone
	(*DataUpdateResqest)(nil),  // 4: pb.DataUpdateResqest
	(*DataUpdateRespone)(nil),  // 5: pb.DataUpdateRespone
	(*ExportDataResquest)(nil), // 6: pb.ExportDataResquest
	(*ExportDataRespone)(nil),  // 7: pb.ExportDataRespone
	(*ImportDataResquest)(nil), // 8: pb.ImportDataResquest
	(*ImportDataRespone)(nil),  // 9: pb.ImportDataRespone
	(*TestResquest)(nil),       // 10: pb.TestResquest
	(*TestRespone)(nil),        // 11: pb.TestRespone
	(*SignUpResquest)(nil),     // 12: pb.SignUpResquest
	(*SignUpRespone)(nil),      // 13: pb.SignUpRespone
	(*DataInfo)(nil),           // 14: pb.DataInfo
}
var file_rpc_data_proto_depIdxs = []int32{
	14, // 0: pb.DataInfoRespone.data:type_name -> pb.DataInfo
	1,  // [1:1] is the sub-list for method output_type
	1,  // [1:1] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_data_proto_init() }
func file_rpc_data_proto_init() {
	if File_rpc_data_proto != nil {
		return
	}
	file_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataInfoResquest); i {
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
		file_rpc_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataInfoRespone); i {
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
		file_rpc_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPostResqest); i {
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
		file_rpc_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPostRespone); i {
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
		file_rpc_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataUpdateResqest); i {
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
		file_rpc_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataUpdateRespone); i {
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
		file_rpc_data_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportDataResquest); i {
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
		file_rpc_data_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportDataRespone); i {
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
		file_rpc_data_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportDataResquest); i {
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
		file_rpc_data_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportDataRespone); i {
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
		file_rpc_data_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResquest); i {
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
		file_rpc_data_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRespone); i {
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
		file_rpc_data_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResquest); i {
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
		file_rpc_data_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRespone); i {
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
			RawDescriptor: file_rpc_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_data_proto_goTypes,
		DependencyIndexes: file_rpc_data_proto_depIdxs,
		MessageInfos:      file_rpc_data_proto_msgTypes,
	}.Build()
	File_rpc_data_proto = out.File
	file_rpc_data_proto_rawDesc = nil
	file_rpc_data_proto_goTypes = nil
	file_rpc_data_proto_depIdxs = nil
}
