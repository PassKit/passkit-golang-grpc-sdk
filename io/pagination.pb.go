//*
// Pagination
//
// Filter your data with multiple parameters.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.0
// source: io/common/pagination.proto

package io

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Pagination struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Limit the number of records returned. If not specified, a default of 25 is used.  Enter -1 for all records.
	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Allows you to offset the first record returned by the limit.
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Array of column names to filter results by.
	FilterField []string `protobuf:"bytes,3,rep,name=filterField,proto3" json:"filterField,omitempty"`
	// Array of values to test against the filter fields.
	FilterValue []string `protobuf:"bytes,4,rep,name=filterValue,proto3" json:"filterValue,omitempty"`
	// will be whitelisted operators in io core
	FilterOperator []string `protobuf:"bytes,5,rep,name=filterOperator,proto3" json:"filterOperator,omitempty"`
	// Field to order results by.
	OrderBy string `protobuf:"bytes,6,opt,name=orderBy,proto3" json:"orderBy,omitempty"`
	// Will return in ascending order if true, or descending order if false.
	OrderAsc      bool `protobuf:"varint,7,opt,name=orderAsc,proto3" json:"orderAsc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	mi := &file_io_common_pagination_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_pagination_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_io_common_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *Pagination) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Pagination) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Pagination) GetFilterField() []string {
	if x != nil {
		return x.FilterField
	}
	return nil
}

func (x *Pagination) GetFilterValue() []string {
	if x != nil {
		return x.FilterValue
	}
	return nil
}

func (x *Pagination) GetFilterOperator() []string {
	if x != nil {
		return x.FilterOperator
	}
	return nil
}

func (x *Pagination) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *Pagination) GetOrderAsc() bool {
	if x != nil {
		return x.OrderAsc
	}
	return false
}

type Filter struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Array of column names to filter results by.
	FilterField []string `protobuf:"bytes,1,rep,name=filterField,proto3" json:"filterField,omitempty"`
	// Array of values to test against the filter fields.
	FilterValue []string `protobuf:"bytes,2,rep,name=filterValue,proto3" json:"filterValue,omitempty"`
	// will be whitelisted operators in io core
	FilterOperator []string `protobuf:"bytes,3,rep,name=filterOperator,proto3" json:"filterOperator,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Filter) Reset() {
	*x = Filter{}
	mi := &file_io_common_pagination_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_pagination_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_io_common_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *Filter) GetFilterField() []string {
	if x != nil {
		return x.FilterField
	}
	return nil
}

func (x *Filter) GetFilterValue() []string {
	if x != nil {
		return x.FilterValue
	}
	return nil
}

func (x *Filter) GetFilterOperator() []string {
	if x != nil {
		return x.FilterOperator
	}
	return nil
}

var File_io_common_pagination_proto protoreflect.FileDescriptor

var file_io_common_pagination_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xac, 0x02, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x73, 0x63, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x73, 0x63, 0x3a,
	0x4e, 0x92, 0x41, 0x4b, 0x0a, 0x49, 0x2a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0x3b, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x70,
	0x6c, 0x69, 0x74, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x20, 0x69, 0x6e, 0x74, 0x6f,
	0x20, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x20, 0x70, 0x61, 0x72, 0x74, 0x73, 0x2e, 0x22,
	0x74, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x47, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73,
	0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5a, 0x24, 0x73, 0x74, 0x61, 0x73, 0x68,
	0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0xaa,
	0x02, 0x0c, 0x50, 0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_io_common_pagination_proto_rawDescOnce sync.Once
	file_io_common_pagination_proto_rawDescData []byte
)

func file_io_common_pagination_proto_rawDescGZIP() []byte {
	file_io_common_pagination_proto_rawDescOnce.Do(func() {
		file_io_common_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_io_common_pagination_proto_rawDesc), len(file_io_common_pagination_proto_rawDesc)))
	})
	return file_io_common_pagination_proto_rawDescData
}

var file_io_common_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_io_common_pagination_proto_goTypes = []any{
	(*Pagination)(nil), // 0: io.Pagination
	(*Filter)(nil),     // 1: io.Filter
}
var file_io_common_pagination_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_io_common_pagination_proto_init() }
func file_io_common_pagination_proto_init() {
	if File_io_common_pagination_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_io_common_pagination_proto_rawDesc), len(file_io_common_pagination_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_common_pagination_proto_goTypes,
		DependencyIndexes: file_io_common_pagination_proto_depIdxs,
		MessageInfos:      file_io_common_pagination_proto_msgTypes,
	}.Build()
	File_io_common_pagination_proto = out.File
	file_io_common_pagination_proto_goTypes = nil
	file_io_common_pagination_proto_depIdxs = nil
}
