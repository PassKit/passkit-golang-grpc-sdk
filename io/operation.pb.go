//*
// Operations
//
// Change the operator depending on the request

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: io/common/operation.proto

package io

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Operation int32

const (
	// Will use the default operator PATCH.
	Operation_OPERATION_DO_NOT_USE Operation = 0
	// PUT will replace all values provided in the request.
	Operation_OPERATION_PUT Operation = 1
	// PATCH will ignore missing or falsey values.
	Operation_OPERATION_PATCH Operation = 2
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0: "OPERATION_DO_NOT_USE",
		1: "OPERATION_PUT",
		2: "OPERATION_PATCH",
	}
	Operation_value = map[string]int32{
		"OPERATION_DO_NOT_USE": 0,
		"OPERATION_PUT":        1,
		"OPERATION_PATCH":      2,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_io_common_operation_proto_enumTypes[0].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_io_common_operation_proto_enumTypes[0]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_io_common_operation_proto_rawDescGZIP(), []int{0}
}

var File_io_common_operation_proto protoreflect.FileDescriptor

var file_io_common_operation_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69, 0x6f, 0x2a,
	0x4d, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x14,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x4f, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x55, 0x53, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x55, 0x54, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x02, 0x42, 0x47,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5a, 0x24, 0x73, 0x74, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69,
	0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0xaa, 0x02, 0x0c, 0x50, 0x61, 0x73, 0x73, 0x4b,
	0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_io_common_operation_proto_rawDescOnce sync.Once
	file_io_common_operation_proto_rawDescData = file_io_common_operation_proto_rawDesc
)

func file_io_common_operation_proto_rawDescGZIP() []byte {
	file_io_common_operation_proto_rawDescOnce.Do(func() {
		file_io_common_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_io_common_operation_proto_rawDescData)
	})
	return file_io_common_operation_proto_rawDescData
}

var file_io_common_operation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_io_common_operation_proto_goTypes = []interface{}{
	(Operation)(0), // 0: io.Operation
}
var file_io_common_operation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_io_common_operation_proto_init() }
func file_io_common_operation_proto_init() {
	if File_io_common_operation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_io_common_operation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_common_operation_proto_goTypes,
		DependencyIndexes: file_io_common_operation_proto_depIdxs,
		EnumInfos:         file_io_common_operation_proto_enumTypes,
	}.Build()
	File_io_common_operation_proto = out.File
	file_io_common_operation_proto_rawDesc = nil
	file_io_common_operation_proto_goTypes = nil
	file_io_common_operation_proto_depIdxs = nil
}
