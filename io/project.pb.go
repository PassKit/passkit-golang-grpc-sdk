//*
// Project
//
// Project manages the status of class object.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: io/common/project.proto

package io

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProjectStatus int32

const (
	// Zero based enum, don't use 0.
	ProjectStatus_NO_PROJECT_STATUS ProjectStatus = 0
	// Active projects allow for creation of records (turned on by default).
	ProjectStatus_PROJECT_ACTIVE_FOR_OBJECT_CREATION ProjectStatus = 1
	// Disabled projects do not allow for creation of records (i.e. temporarily disable enroling of members / creating of coupons, etc).
	ProjectStatus_PROJECT_DISABLED_FOR_OBJECT_CREATION ProjectStatus = 2
	// Default status for all projects. A draft project can do everything a published program can:
	// Cards can be created, stats will show up, etc.
	// Data belonging to draft projects is automatically purged every 7 days.
	// A draft project uses a PassKit Apple Developer Certificate, and has a PassKit legal disclaimer on the back.
	// Draft projects cannot be used for a production setup.
	ProjectStatus_PROJECT_DRAFT ProjectStatus = 4
	// A project can only be published if a custom Apple Developer Certificate is uploaded, if payment details have been provided,
	// and an account email address has been validated.
	// Published projects do not have a PassKit legal disclaimer on the back.
	// Published projects do not have their data purged on a 7 day basis.
	ProjectStatus_PROJECT_PUBLISHED ProjectStatus = 8
	// A project is private: records can only by created by authorized users / via the portal interface.
	ProjectStatus_PROJECT_PRIVATE ProjectStatus = 16
	// The project is currently over quota, and new records cannot be created (set by system only; cannot be set by user)
	ProjectStatus_PROJECT_OVER_QUOTA ProjectStatus = 32
	// A project which needs to be deleted. If delete flag is set, project will be deleted in the near future.
	ProjectStatus_PROJECT_DELETED ProjectStatus = 64
)

// Enum value maps for ProjectStatus.
var (
	ProjectStatus_name = map[int32]string{
		0:  "NO_PROJECT_STATUS",
		1:  "PROJECT_ACTIVE_FOR_OBJECT_CREATION",
		2:  "PROJECT_DISABLED_FOR_OBJECT_CREATION",
		4:  "PROJECT_DRAFT",
		8:  "PROJECT_PUBLISHED",
		16: "PROJECT_PRIVATE",
		32: "PROJECT_OVER_QUOTA",
		64: "PROJECT_DELETED",
	}
	ProjectStatus_value = map[string]int32{
		"NO_PROJECT_STATUS":                    0,
		"PROJECT_ACTIVE_FOR_OBJECT_CREATION":   1,
		"PROJECT_DISABLED_FOR_OBJECT_CREATION": 2,
		"PROJECT_DRAFT":                        4,
		"PROJECT_PUBLISHED":                    8,
		"PROJECT_PRIVATE":                      16,
		"PROJECT_OVER_QUOTA":                   32,
		"PROJECT_DELETED":                      64,
	}
)

func (x ProjectStatus) Enum() *ProjectStatus {
	p := new(ProjectStatus)
	*p = x
	return p
}

func (x ProjectStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_io_common_project_proto_enumTypes[0].Descriptor()
}

func (ProjectStatus) Type() protoreflect.EnumType {
	return &file_io_common_project_proto_enumTypes[0]
}

func (x ProjectStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectStatus.Descriptor instead.
func (ProjectStatus) EnumDescriptor() ([]byte, []int) {
	return file_io_common_project_proto_rawDescGZIP(), []int{0}
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The PassProtocol the project implements.
	Protocol PassProtocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=io.PassProtocol" json:"protocol,omitempty"`
	// The class ID that the projects refers to (highest level protocol object; i.e. member program id, coupon campaign id, etc).
	ClassId string `protobuf:"bytes,2,opt,name=classId,proto3" json:"classId,omitempty"`
	// The project name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The project short-code; used for creating public project URL's.
	ShortCode string `protobuf:"bytes,4,opt,name=shortCode,proto3" json:"shortCode,omitempty"`
	// The timestamp when the project was created.
	Created *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	// A shared secret used for creating signed links.
	// @tag: validateGeneric:"-" validateCreate:"-" validateUpdate:"-"
	Secret string `protobuf:"bytes,6,opt,name=secret,proto3" json:"secret,omitempty" validateGeneric:"-" validateCreate:"-" validateUpdate:"-"`
	// A key used to create encrypted, signed links.
	// @tag: validateGeneric:"-" validateCreate:"-" validateUpdate:"-"
	Key string `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty" validateGeneric:"-" validateCreate:"-" validateUpdate:"-"`
	// Encrypted class ID. Can be used in integrations where the Class ID needs to be publicly exposed.
	// @tag: validateGeneric:"-" validateCreate:"-" validateUpdate:"-"
	EncryptedClassId string `protobuf:"bytes,9,opt,name=encryptedClassId,proto3" json:"encryptedClassId,omitempty" validateGeneric:"-" validateCreate:"-" validateUpdate:"-"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_common_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_io_common_project_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetProtocol() PassProtocol {
	if x != nil {
		return x.Protocol
	}
	return PassProtocol_PASS_PROTOCOL_DO_NOT_USE
}

func (x *Project) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetShortCode() string {
	if x != nil {
		return x.ShortCode
	}
	return ""
}

func (x *Project) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Project) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Project) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Project) GetEncryptedClassId() string {
	if x != nil {
		return x.EncryptedClassId
	}
	return ""
}

type ProjectByShortCodeResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project  *Project      `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Template *PassTemplate `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *ProjectByShortCodeResult) Reset() {
	*x = ProjectByShortCodeResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_common_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectByShortCodeResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectByShortCodeResult) ProtoMessage() {}

func (x *ProjectByShortCodeResult) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectByShortCodeResult.ProtoReflect.Descriptor instead.
func (*ProjectByShortCodeResult) Descriptor() ([]byte, []int) {
	return file_io_common_project_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectByShortCodeResult) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *ProjectByShortCodeResult) GetTemplate() *PassTemplate {
	if x != nil {
		return x.Template
	}
	return nil
}

type ProjectStatusFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ProjectStatus `protobuf:"varint,1,opt,name=status,proto3,enum=io.ProjectStatus" json:"status,omitempty"`
}

func (x *ProjectStatusFilter) Reset() {
	*x = ProjectStatusFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_common_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectStatusFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectStatusFilter) ProtoMessage() {}

func (x *ProjectStatusFilter) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectStatusFilter.ProtoReflect.Descriptor instead.
func (*ProjectStatusFilter) Descriptor() ([]byte, []int) {
	return file_io_common_project_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectStatusFilter) GetStatus() ProjectStatus {
	if x != nil {
		return x.Status
	}
	return ProjectStatus_NO_PROJECT_STATUS
}

var File_io_common_project_proto protoreflect.FileDescriptor

var file_io_common_project_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x69, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x42, 0x79, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x25, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6f, 0x2e,
	0x50, 0x61, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x40, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x29, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x69, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0xe4, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x4f,
	0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10,
	0x00, 0x12, 0x26, 0x0a, 0x22, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x28, 0x0a, 0x24, 0x50, 0x52, 0x4f,
	0x4a, 0x45, 0x43, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x5f, 0x46, 0x4f,
	0x52, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x44,
	0x52, 0x41, 0x46, 0x54, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43,
	0x54, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x08, 0x12, 0x13, 0x0a,
	0x0f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45,
	0x10, 0x10, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4f, 0x56,
	0x45, 0x52, 0x5f, 0x51, 0x55, 0x4f, 0x54, 0x41, 0x10, 0x20, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52,
	0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x40, 0x42,
	0x47, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5a, 0x24, 0x73, 0x74, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b,
	0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0xaa, 0x02, 0x0c, 0x50, 0x61, 0x73, 0x73,
	0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_io_common_project_proto_rawDescOnce sync.Once
	file_io_common_project_proto_rawDescData = file_io_common_project_proto_rawDesc
)

func file_io_common_project_proto_rawDescGZIP() []byte {
	file_io_common_project_proto_rawDescOnce.Do(func() {
		file_io_common_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_io_common_project_proto_rawDescData)
	})
	return file_io_common_project_proto_rawDescData
}

var file_io_common_project_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_io_common_project_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_io_common_project_proto_goTypes = []interface{}{
	(ProjectStatus)(0),               // 0: io.ProjectStatus
	(*Project)(nil),                  // 1: io.Project
	(*ProjectByShortCodeResult)(nil), // 2: io.ProjectByShortCodeResult
	(*ProjectStatusFilter)(nil),      // 3: io.ProjectStatusFilter
	(PassProtocol)(0),                // 4: io.PassProtocol
	(*timestamppb.Timestamp)(nil),    // 5: google.protobuf.Timestamp
	(*PassTemplate)(nil),             // 6: io.PassTemplate
}
var file_io_common_project_proto_depIdxs = []int32{
	4, // 0: io.Project.protocol:type_name -> io.PassProtocol
	5, // 1: io.Project.created:type_name -> google.protobuf.Timestamp
	1, // 2: io.ProjectByShortCodeResult.project:type_name -> io.Project
	6, // 3: io.ProjectByShortCodeResult.template:type_name -> io.PassTemplate
	0, // 4: io.ProjectStatusFilter.status:type_name -> io.ProjectStatus
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_io_common_project_proto_init() }
func file_io_common_project_proto_init() {
	if File_io_common_project_proto != nil {
		return
	}
	file_io_common_protocols_proto_init()
	file_io_common_template_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_io_common_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_io_common_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectByShortCodeResult); i {
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
		file_io_common_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectStatusFilter); i {
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
			RawDescriptor: file_io_common_project_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_common_project_proto_goTypes,
		DependencyIndexes: file_io_common_project_proto_depIdxs,
		EnumInfos:         file_io_common_project_proto_enumTypes,
		MessageInfos:      file_io_common_project_proto_msgTypes,
	}.Build()
	File_io_common_project_proto = out.File
	file_io_common_project_proto_rawDesc = nil
	file_io_common_project_proto_goTypes = nil
	file_io_common_project_proto_depIdxs = nil
}
