// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.0
// source: io/core/a_rpc_images.proto

package io

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_io_core_a_rpc_images_proto protoreflect.FileDescriptor

var file_io_core_a_rpc_images_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x69, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x69, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x69, 0x6f, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x69,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc4, 0x1b, 0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x12, 0xa0, 0x01, 0x0a, 0x0f, 0x73, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x5e, 0x92, 0x41, 0x42, 0x12, 0x11, 0x53, 0x65, 0x74, 0x20, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x2d, 0x53, 0x65, 0x74,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x20, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x64, 0x20, 0x69, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x94, 0x01, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x07, 0x2e, 0x69, 0x6f, 0x2e, 0x55, 0x72, 0x6c, 0x22, 0x60, 0x92, 0x41, 0x47, 0x12, 0x11, 0x47,
	0x65, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x1a, 0x32, 0x47, 0x65, 0x74, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x55, 0x52, 0x4c, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x6f, 0x0a, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x69, 0x6f, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x0c, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x73, 0x22, 0x3b,
	0x92, 0x41, 0x26, 0x12, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x1a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x20, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a,
	0x01, 0x2a, 0x22, 0x07, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x72, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x69, 0x6f, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x0f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x22, 0x3c, 0x92, 0x41, 0x28, 0x12, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x1a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x20, 0x61, 0x6e,
	0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0b, 0x3a, 0x01, 0x2a, 0x1a, 0x06, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x5e, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x06,
	0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x07, 0x2e, 0x69, 0x6f, 0x2e, 0x55, 0x72, 0x6c, 0x22,
	0x3e, 0x92, 0x41, 0x28, 0x12, 0x0d, 0x47, 0x65, 0x74, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x20,
	0x55, 0x52, 0x4c, 0x1a, 0x17, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61,
	0x6e, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x75, 0x72, 0x6c, 0x2e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0xf1, 0x01, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x52, 0x4c, 0x12, 0x15, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x69, 0x6f,
	0x2e, 0x55, 0x72, 0x6c, 0x22, 0xbc, 0x01, 0x92, 0x41, 0x96, 0x01, 0x12, 0x13, 0x47, 0x65, 0x74,
	0x20, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x55, 0x52, 0x4c,
	0x1a, 0x3b, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61, 0x20, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x75, 0x72, 0x6c, 0x20, 0x77, 0x69,
	0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x20,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x4a, 0x22, 0x0a,
	0x03, 0x34, 0x30, 0x33, 0x12, 0x1b, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x20, 0x6c, 0x61, 0x63,
	0x6b, 0x73, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4a, 0x1e, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x20, 0x77, 0x61, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x7d, 0x12, 0x85, 0x01, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x6d, 0x70,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x06, 0x2e, 0x69, 0x6f,
	0x2e, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x50, 0x92, 0x41, 0x2d, 0x12, 0x10,
	0x47, 0x65, 0x74, 0x20, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x20, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x1a, 0x19, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61, 0x20, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x12, 0x18, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x96, 0x01, 0x0a, 0x16,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x6d,
	0x70, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x06, 0x2e, 0x69,
	0x6f, 0x2e, 0x49, 0x64, 0x22, 0x5e, 0x92, 0x41, 0x3d, 0x12, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x20, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x20, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x26,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x64, 0x20, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x1a,
	0x13, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0xc0, 0x01, 0x0a, 0x14, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x6d,
	0x70, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1c, 0x2e,
	0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x69, 0x6f,
	0x2e, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x22, 0x73, 0x92, 0x41, 0x51, 0x12, 0x17, 0x47, 0x65, 0x74, 0x20, 0x53, 0x74, 0x61,
	0x6d, 0x70, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x1a, 0x36, 0x47, 0x65, 0x74, 0x73, 0x20, 0x61, 0x20, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01,
	0x2a, 0x22, 0x14, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x9a, 0x01, 0x0a, 0x14, 0x67, 0x65, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c,
	0x12, 0x17, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x07, 0x2e, 0x69, 0x6f, 0x2e, 0x55,
	0x72, 0x6c, 0x22, 0x60, 0x92, 0x41, 0x3b, 0x12, 0x17, 0x47, 0x65, 0x74, 0x20, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x55, 0x52, 0x4c,
	0x1a, 0x20, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61, 0x20, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x55, 0x52,
	0x4c, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x7d, 0x12, 0x43, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x06, 0x2e, 0x69, 0x6f,
	0x2e, 0x49, 0x64, 0x1a, 0x07, 0x2e, 0x69, 0x6f, 0x2e, 0x55, 0x72, 0x6c, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x76, 0x0a, 0x0e, 0x67, 0x65, 0x74,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x06, 0x2e, 0x69, 0x6f,
	0x2e, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x22, 0x4b, 0x92, 0x41, 0x2e, 0x12, 0x10, 0x47, 0x65, 0x74, 0x20, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x20, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x1a, 0x1a, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x70, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x06, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x69, 0x6f, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x47, 0x92, 0x41, 0x2c, 0x12,
	0x0e, 0x47, 0x65, 0x74, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x44, 0x61, 0x74, 0x61, 0x1a,
	0x1a, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x12, 0x10, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x6c, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x06, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x3d, 0x92, 0x41, 0x27, 0x12, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x20, 0x61,
	0x6e, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0xa1, 0x01, 0x0a, 0x14, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x69, 0x6f, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x0f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x22, 0x5f, 0x92, 0x41, 0x3a, 0x12, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x20, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x1a, 0x20, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x2a, 0x1a, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x7d, 0x12, 0xa7, 0x01, 0x0a, 0x1b, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x65, 0x92, 0x41, 0x4e, 0x12, 0x17, 0x47, 0x65, 0x74,
	0x20, 0x41, 0x6c, 0x6c, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x20, 0x46, 0x6f, 0x72, 0x20,
	0x55, 0x73, 0x65, 0x72, 0x1a, 0x33, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20,
	0x61, 0x6c, 0x6c, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12,
	0x0c, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x30, 0x01, 0x12,
	0xa2, 0x01, 0x0a, 0x11, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x46, 0x6f,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x1a, 0x0f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x22, 0x6d, 0x92, 0x41, 0x4e, 0x12, 0x17, 0x47, 0x65, 0x74, 0x20, 0x41, 0x6c,
	0x6c, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x20, 0x46, 0x6f, 0x72, 0x20, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x33, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61, 0x6c, 0x6c,
	0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x20, 0x75,
	0x6e, 0x64, 0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22,
	0x11, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x30, 0x01, 0x12, 0x97, 0x01, 0x0a, 0x14, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x2e,
	0x69, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0f, 0x2e,
	0x69, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x5c,
	0x92, 0x41, 0x4a, 0x12, 0x18, 0x47, 0x65, 0x74, 0x20, 0x41, 0x6c, 0x6c, 0x20, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x2e, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x30, 0x01, 0x12, 0x92,
	0x01, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x0b, 0x2e,
	0x69, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x0f, 0x2e, 0x69, 0x6f, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x64, 0x92, 0x41, 0x4a,
	0x12, 0x18, 0x47, 0x65, 0x74, 0x20, 0x41, 0x6c, 0x6c, 0x20, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x2e, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x20, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x30, 0x01, 0x12, 0xb2, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x2e,
	0x69, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x09, 0x2e,
	0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7e, 0x92, 0x41, 0x66, 0x12, 0x26, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x20, 0x41, 0x6c, 0x6c, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x20,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x54, 0x68, 0x65,
	0x20, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x3c, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73,
	0x20, 0x61, 0x20, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0xa9, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x0b, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x09, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x81, 0x01, 0x92, 0x41, 0x66, 0x12, 0x26, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x41, 0x6c,
	0x6c, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x20, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x54, 0x68, 0x65, 0x20, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x3c,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61, 0x20, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x20, 0x6f, 0x66, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x20, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x64, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x27, 0x73, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0xf9, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x09, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0xbd, 0x01, 0x92, 0x41, 0x9f, 0x01, 0x12, 0x22, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x41,
	0x6c, 0x6c, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x20, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x20,
	0x62, 0x79, 0x20, 0x54, 0x68, 0x65, 0x20, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x35, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x4a, 0x22, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x1b, 0x0a, 0x19, 0x55, 0x73, 0x65,
	0x72, 0x20, 0x6c, 0x61, 0x63, 0x6b, 0x73, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x1e, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x17, 0x0a,
	0x15, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x20, 0x77, 0x61, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0xa9, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x1a, 0x09, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x7b, 0x92, 0x41, 0x5b, 0x12, 0x22, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x41, 0x6c, 0x6c, 0x20,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x20, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20,
	0x54, 0x68, 0x65, 0x20, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x35, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x73, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x20, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x64, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xe4, 0x06, 0x92,
	0x41, 0x99, 0x06, 0x12, 0xd4, 0x01, 0x0a, 0x12, 0x50, 0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x20,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x20, 0x41, 0x50, 0x49, 0x12, 0x43, 0x41, 0x50, 0x49, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x20, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x20, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x50, 0x61,
	0x73, 0x73, 0x20, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x20, 0x26, 0x20, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x20, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x1a,
	0x38, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x72, 0x6d, 0x73,
	0x2d, 0x6f, 0x66, 0x2d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x22, 0x3f, 0x0a, 0x0f, 0x50, 0x61, 0x73,
	0x73, 0x4b, 0x69, 0x74, 0x20, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b,
	0x69, 0x74, 0x2e, 0x69, 0x6f, 0x1a, 0x13, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x40, 0x70,
	0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x52, 0x39, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x32, 0x0a, 0x28, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x69, 0x73, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x66, 0x75, 0x6c, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x52, 0x34, 0x0a, 0x03,
	0x34, 0x30, 0x30, 0x12, 0x2d, 0x0a, 0x2b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20,
	0x77, 0x68, 0x65, 0x6e, 0x20, 0x77, 0x72, 0x6f, 0x6e, 0x67, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x69, 0x73, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x64, 0x2e, 0x52, 0x30, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x29, 0x0a, 0x27, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x2e, 0x52, 0x50, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x49, 0x0a, 0x47, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68,
	0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74,
	0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x34, 0x0a,
	0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a,
	0x02, 0x01, 0x07, 0x52, 0x3c, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x35, 0x0a, 0x2b, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x20, 0x69, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01,
	0x07, 0x52, 0x57, 0x0a, 0x03, 0x35, 0x30, 0x33, 0x12, 0x50, 0x0a, 0x4e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x20, 0x42, 0x61, 0x63, 0x6b, 0x20, 0x6f, 0x66, 0x66, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x32, 0x35, 0x30, 0x6d, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x20, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x2e, 0x5a, 0x3e, 0x0a, 0x3c, 0x0a, 0x0a,
	0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x2e, 0x08, 0x02, 0x12, 0x19,
	0x4a, 0x57, 0x54, 0x20, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a,
	0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x0a, 0x10, 0x63, 0x6f,
	0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5a, 0x24,
	0x73, 0x74, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67,
	0x6f, 0x2f, 0x69, 0x6f, 0xaa, 0x02, 0x0c, 0x50, 0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x2e, 0x47,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_io_core_a_rpc_images_proto_goTypes = []any{
	(*ProfileImageInput)(nil),        // 0: io.ProfileImageInput
	(*emptypb.Empty)(nil),            // 1: google.protobuf.Empty
	(*CreateImageInput)(nil),         // 2: io.CreateImageInput
	(*UpdateImageInput)(nil),         // 3: io.UpdateImageInput
	(*Id)(nil),                       // 4: io.Id
	(*StampImageRequest)(nil),        // 5: io.StampImageRequest
	(*StampImageConfig)(nil),         // 6: io.StampImageConfig
	(*StampImagePreviewRequest)(nil), // 7: io.StampImagePreviewRequest
	(*LocalizedImageInput)(nil),      // 8: io.LocalizedImageInput
	(*Pagination)(nil),               // 9: io.Pagination
	(*Filters)(nil),                  // 10: io.Filters
	(*Url)(nil),                      // 11: io.Url
	(*ImageIds)(nil),                 // 12: io.ImageIds
	(*ImageRecord)(nil),              // 13: io.ImageRecord
	(*StampImagePreview)(nil),        // 14: io.StampImagePreview
	(*ImageBundle)(nil),              // 15: io.ImageBundle
	(*Count)(nil),                    // 16: io.Count
}
var file_io_core_a_rpc_images_proto_depIdxs = []int32{
	0,  // 0: io.Images.setProfileImage:input_type -> io.ProfileImageInput
	1,  // 1: io.Images.getProfileImage:input_type -> google.protobuf.Empty
	2,  // 2: io.Images.createImages:input_type -> io.CreateImageInput
	3,  // 3: io.Images.updateImage:input_type -> io.UpdateImageInput
	4,  // 4: io.Images.getImageURL:input_type -> io.Id
	5,  // 5: io.Images.getStampImageURL:input_type -> io.StampImageRequest
	4,  // 6: io.Images.getStampImageConfig:input_type -> io.Id
	6,  // 7: io.Images.updateStampImageConfig:input_type -> io.StampImageConfig
	7,  // 8: io.Images.getStampImagePreview:input_type -> io.StampImagePreviewRequest
	8,  // 9: io.Images.getLocalizedImageURL:input_type -> io.LocalizedImageInput
	4,  // 10: io.Images.getProfileImageById:input_type -> io.Id
	4,  // 11: io.Images.getImageBundle:input_type -> io.Id
	4,  // 12: io.Images.getImageData:input_type -> io.Id
	4,  // 13: io.Images.deleteImage:input_type -> io.Id
	8,  // 14: io.Images.deleteLocalizedImage:input_type -> io.LocalizedImageInput
	9,  // 15: io.Images.listImagesForUserDeprecated:input_type -> io.Pagination
	10, // 16: io.Images.listImagesForUser:input_type -> io.Filters
	9,  // 17: io.Images.listImagesDeprecated:input_type -> io.Pagination
	10, // 18: io.Images.listImages:input_type -> io.Filters
	9,  // 19: io.Images.countImagesDeprecated:input_type -> io.Pagination
	10, // 20: io.Images.countImages:input_type -> io.Filters
	9,  // 21: io.Images.countImagesForUserDeprecated:input_type -> io.Pagination
	10, // 22: io.Images.countImagesForUser:input_type -> io.Filters
	1,  // 23: io.Images.setProfileImage:output_type -> google.protobuf.Empty
	11, // 24: io.Images.getProfileImage:output_type -> io.Url
	12, // 25: io.Images.createImages:output_type -> io.ImageIds
	13, // 26: io.Images.updateImage:output_type -> io.ImageRecord
	11, // 27: io.Images.getImageURL:output_type -> io.Url
	11, // 28: io.Images.getStampImageURL:output_type -> io.Url
	6,  // 29: io.Images.getStampImageConfig:output_type -> io.StampImageConfig
	4,  // 30: io.Images.updateStampImageConfig:output_type -> io.Id
	14, // 31: io.Images.getStampImagePreview:output_type -> io.StampImagePreview
	11, // 32: io.Images.getLocalizedImageURL:output_type -> io.Url
	11, // 33: io.Images.getProfileImageById:output_type -> io.Url
	15, // 34: io.Images.getImageBundle:output_type -> io.ImageBundle
	13, // 35: io.Images.getImageData:output_type -> io.ImageRecord
	1,  // 36: io.Images.deleteImage:output_type -> google.protobuf.Empty
	13, // 37: io.Images.deleteLocalizedImage:output_type -> io.ImageRecord
	13, // 38: io.Images.listImagesForUserDeprecated:output_type -> io.ImageRecord
	13, // 39: io.Images.listImagesForUser:output_type -> io.ImageRecord
	13, // 40: io.Images.listImagesDeprecated:output_type -> io.ImageRecord
	13, // 41: io.Images.listImages:output_type -> io.ImageRecord
	16, // 42: io.Images.countImagesDeprecated:output_type -> io.Count
	16, // 43: io.Images.countImages:output_type -> io.Count
	16, // 44: io.Images.countImagesForUserDeprecated:output_type -> io.Count
	16, // 45: io.Images.countImagesForUser:output_type -> io.Count
	23, // [23:46] is the sub-list for method output_type
	0,  // [0:23] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_io_core_a_rpc_images_proto_init() }
func file_io_core_a_rpc_images_proto_init() {
	if File_io_core_a_rpc_images_proto != nil {
		return
	}
	file_io_common_common_objects_proto_init()
	file_io_common_pagination_proto_init()
	file_io_image_image_proto_init()
	file_io_common_filter_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_io_core_a_rpc_images_proto_rawDesc), len(file_io_core_a_rpc_images_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_io_core_a_rpc_images_proto_goTypes,
		DependencyIndexes: file_io_core_a_rpc_images_proto_depIdxs,
	}.Build()
	File_io_core_a_rpc_images_proto = out.File
	file_io_core_a_rpc_images_proto_goTypes = nil
	file_io_core_a_rpc_images_proto_depIdxs = nil
}
