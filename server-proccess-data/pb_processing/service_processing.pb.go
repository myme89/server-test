// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: service_processing.proto

package pb_processing

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_processing_proto protoreflect.FileDescriptor

var file_service_processing_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x62, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x19, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x87, 0x03, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x12, 0x52, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x32, 0x12, 0x21, 0x2e,
	0x70, 0x62, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74, 0x31,
	0x1a, 0x20, 0x2e, 0x70, 0x62, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x65, 0x31, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x12, 0x22, 0x2e, 0x70, 0x62,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x70, 0x62, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x18, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x1e, 0x2e, 0x70, 0x62, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x22,
	0x00, 0x12, 0x68, 0x0a, 0x13, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x66, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x27, 0x2e, 0x70, 0x62, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x70, 0x62, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70,
	0x62, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_service_processing_proto_goTypes = []interface{}{
	(*DataInfoTestResquest1)(nil),       // 0: pb_storage.DataInfoTestResquest1
	(*ProcessingFileResquest)(nil),      // 1: pb_storage.ProcessingFileResquest
	(*ExportFileResquest)(nil),          // 2: pb_storage.ExportFileResquest
	(*DownloadFileProcessResquest)(nil), // 3: pb_storage.DownloadFileProcessResquest
	(*DataInfoTestRespone1)(nil),        // 4: pb_storage.DataInfoTestRespone1
	(*ProcessingFileRespone)(nil),       // 5: pb_storage.ProcessingFileRespone
	(*ExportFileRespone)(nil),           // 6: pb_storage.ExportFileRespone
	(*DownloadFileProcessRespone)(nil),  // 7: pb_storage.DownloadFileProcessRespone
}
var file_service_processing_proto_depIdxs = []int32{
	0, // 0: pb_storage.Processing.TestData2:input_type -> pb_storage.DataInfoTestResquest1
	1, // 1: pb_storage.Processing.ProcessingFileExcel:input_type -> pb_storage.ProcessingFileResquest
	2, // 2: pb_storage.Processing.ExportTemplateFileUpload:input_type -> pb_storage.ExportFileResquest
	3, // 3: pb_storage.Processing.DownloafFileProcess:input_type -> pb_storage.DownloadFileProcessResquest
	4, // 4: pb_storage.Processing.TestData2:output_type -> pb_storage.DataInfoTestRespone1
	5, // 5: pb_storage.Processing.ProcessingFileExcel:output_type -> pb_storage.ProcessingFileRespone
	6, // 6: pb_storage.Processing.ExportTemplateFileUpload:output_type -> pb_storage.ExportFileRespone
	7, // 7: pb_storage.Processing.DownloafFileProcess:output_type -> pb_storage.DownloadFileProcessRespone
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_processing_proto_init() }
func file_service_processing_proto_init() {
	if File_service_processing_proto != nil {
		return
	}
	file_rpc_data_processing_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_processing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_processing_proto_goTypes,
		DependencyIndexes: file_service_processing_proto_depIdxs,
	}.Build()
	File_service_processing_proto = out.File
	file_service_processing_proto_rawDesc = nil
	file_service_processing_proto_goTypes = nil
	file_service_processing_proto_depIdxs = nil
}
