// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: file_upload.proto

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

type FileUploadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string  `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Size     float32 `protobuf:"fixed32,2,opt,name=size,proto3" json:"size,omitempty"`
	Typefile string  `protobuf:"bytes,3,opt,name=typefile,proto3" json:"typefile,omitempty"`
	Iduser   string  `protobuf:"bytes,4,opt,name=iduser,proto3" json:"iduser,omitempty"`
	Link     string  `protobuf:"bytes,5,opt,name=link,proto3" json:"link,omitempty"`
	Createat string  `protobuf:"bytes,6,opt,name=createat,proto3" json:"createat,omitempty"`
}

func (x *FileUploadInfo) Reset() {
	*x = FileUploadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadInfo) ProtoMessage() {}

func (x *FileUploadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_file_upload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadInfo.ProtoReflect.Descriptor instead.
func (*FileUploadInfo) Descriptor() ([]byte, []int) {
	return file_file_upload_proto_rawDescGZIP(), []int{0}
}

func (x *FileUploadInfo) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileUploadInfo) GetSize() float32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileUploadInfo) GetTypefile() string {
	if x != nil {
		return x.Typefile
	}
	return ""
}

func (x *FileUploadInfo) GetIduser() string {
	if x != nil {
		return x.Iduser
	}
	return ""
}

func (x *FileUploadInfo) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *FileUploadInfo) GetCreateat() string {
	if x != nil {
		return x.Createat
	}
	return ""
}

type ImportFileExcel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Content  []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ImportFileExcel) Reset() {
	*x = ImportFileExcel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportFileExcel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportFileExcel) ProtoMessage() {}

func (x *ImportFileExcel) ProtoReflect() protoreflect.Message {
	mi := &file_file_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportFileExcel.ProtoReflect.Descriptor instead.
func (*ImportFileExcel) Descriptor() ([]byte, []int) {
	return file_file_upload_proto_rawDescGZIP(), []int{1}
}

func (x *ImportFileExcel) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ImportFileExcel) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type TemplateFilePersonInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lastname    string `protobuf:"bytes,1,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Firstname   string `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Fullname    string `protobuf:"bytes,3,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Phonenumber string `protobuf:"bytes,4,opt,name=phonenumber,proto3" json:"phonenumber,omitempty"`
	Address     string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *TemplateFilePersonInfo) Reset() {
	*x = TemplateFilePersonInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_upload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateFilePersonInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateFilePersonInfo) ProtoMessage() {}

func (x *TemplateFilePersonInfo) ProtoReflect() protoreflect.Message {
	mi := &file_file_upload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateFilePersonInfo.ProtoReflect.Descriptor instead.
func (*TemplateFilePersonInfo) Descriptor() ([]byte, []int) {
	return file_file_upload_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateFilePersonInfo) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *TemplateFilePersonInfo) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *TemplateFilePersonInfo) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *TemplateFilePersonInfo) GetPhonenumber() string {
	if x != nil {
		return x.Phonenumber
	}
	return ""
}

func (x *TemplateFilePersonInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_file_upload_proto protoreflect.FileDescriptor

var file_file_upload_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x62, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x22, 0xa4, 0x01, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x64, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x64, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x61, 0x74, 0x22, 0x47, 0x0a, 0x0f, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0xaa, 0x01, 0x0a, 0x16, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x29, 0x5a,
	0x27, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x62, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_upload_proto_rawDescOnce sync.Once
	file_file_upload_proto_rawDescData = file_file_upload_proto_rawDesc
)

func file_file_upload_proto_rawDescGZIP() []byte {
	file_file_upload_proto_rawDescOnce.Do(func() {
		file_file_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_upload_proto_rawDescData)
	})
	return file_file_upload_proto_rawDescData
}

var file_file_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_file_upload_proto_goTypes = []interface{}{
	(*FileUploadInfo)(nil),         // 0: pb_database.FileUploadInfo
	(*ImportFileExcel)(nil),        // 1: pb_database.ImportFileExcel
	(*TemplateFilePersonInfo)(nil), // 2: pb_database.TemplateFilePersonInfo
}
var file_file_upload_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_upload_proto_init() }
func file_file_upload_proto_init() {
	if File_file_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadInfo); i {
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
		file_file_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportFileExcel); i {
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
		file_file_upload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateFilePersonInfo); i {
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
			RawDescriptor: file_file_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_file_upload_proto_goTypes,
		DependencyIndexes: file_file_upload_proto_depIdxs,
		MessageInfos:      file_file_upload_proto_msgTypes,
	}.Build()
	File_file_upload_proto = out.File
	file_file_upload_proto_rawDesc = nil
	file_file_upload_proto_goTypes = nil
	file_file_upload_proto_depIdxs = nil
}
