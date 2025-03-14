//*
// Links
//
// You can render different types of link on digital pass.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.0
// source: io/common/links.proto

package io

import (
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

// Used to specify the type of link for link field. Each type has different icon on Google Pay.
type LinkType int32

const (
	// A link to website.
	LinkType_URI_DO_NOT_USE LinkType = 0
	// A link to website.
	LinkType_URI_WEB LinkType = 1
	// A phone number.
	LinkType_URI_TEL LinkType = 2
	// An email address.
	LinkType_URI_EMAIL LinkType = 3
	// A location address.
	LinkType_URI_LOCATION LinkType = 4
	// A calendar event.
	LinkType_URI_CALENDAR LinkType = 5
)

// Enum value maps for LinkType.
var (
	LinkType_name = map[int32]string{
		0: "URI_DO_NOT_USE",
		1: "URI_WEB",
		2: "URI_TEL",
		3: "URI_EMAIL",
		4: "URI_LOCATION",
		5: "URI_CALENDAR",
	}
	LinkType_value = map[string]int32{
		"URI_DO_NOT_USE": 0,
		"URI_WEB":        1,
		"URI_TEL":        2,
		"URI_EMAIL":      3,
		"URI_LOCATION":   4,
		"URI_CALENDAR":   5,
	}
)

func (x LinkType) Enum() *LinkType {
	p := new(LinkType)
	*p = x
	return p
}

func (x LinkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LinkType) Descriptor() protoreflect.EnumDescriptor {
	return file_io_common_links_proto_enumTypes[0].Descriptor()
}

func (LinkType) Type() protoreflect.EnumType {
	return &file_io_common_links_proto_enumTypes[0]
}

func (x LinkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LinkType.Descriptor instead.
func (LinkType) EnumDescriptor() ([]byte, []int) {
	return file_io_common_links_proto_rawDescGZIP(), []int{0}
}

// Used to specify links put on the back of the pass.
type Link struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Link Id. Not writable.
	// @tag: validateCreate:"omitempty,uuidCompressedString" validateUpdate:"required,uuidCompressedString"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateCreate:"omitempty,uuidCompressedString" validateUpdate:"required,uuidCompressedString"`
	// A link text.
	// @tag: validateGeneric:"required_without=Id,uri" validateCreate:"required_without=Id,uri" validateUpdate:"required_without=Id,uri"
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty" validateGeneric:"required_without=Id,uri" validateCreate:"required_without=Id,uri" validateUpdate:"required_without=Id,uri"`
	// Title for the link.
	// @tag: validateGeneric:"required_without=Id" validateCreate:"required_without=Id" validateUpdate:"required_without=Id"
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty" validateGeneric:"required_without=Id" validateCreate:"required_without=Id" validateUpdate:"required_without=Id"`
	// Type of link data.
	// @tag: validateGeneric:"max=5,min=1" validateCreate:"max=5,min=1" validateUpdate:"max=5,min=1"
	Type LinkType `protobuf:"varint,4,opt,name=type,proto3,enum=io.LinkType" json:"type,omitempty" validateGeneric:"max=5,min=1" validateCreate:"max=5,min=1" validateUpdate:"max=5,min=1"`
	// This customises link text for different languages. Ignored by Google Pay passes.
	LocalizedLink *LocalizedString `protobuf:"bytes,5,opt,name=localizedLink,proto3" json:"localizedLink,omitempty"`
	// This translates link title in different languages.
	LocalizedTitle *LocalizedString `protobuf:"bytes,6,opt,name=localizedTitle,proto3" json:"localizedTitle,omitempty"`
	// Indicates which wallets the link should be rendered for (this allows to hide certain link on Apple Wallet, and vise versa).
	// @tag: validateGeneric:"required,min=1" validateCreate:"required,min=1" validateUpdate:"required,min=1"
	Usage []UsageType `protobuf:"varint,7,rep,packed,name=usage,proto3,enum=io.UsageType" json:"usage,omitempty" validateGeneric:"required,min=1" validateCreate:"required,min=1" validateUpdate:"required,min=1"`
	// Links will be rendered in order of their position, from lowest to highest. Position can be any value, E.g. 3 links with positions 30, 10, 20 would render 10 first, 20 second and 30 third.  If no position is provided, the order of the links cannot be guaranteed.
	Position      uint32 `protobuf:"varint,8,opt,name=position,proto3" json:"position,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Link) Reset() {
	*x = Link{}
	mi := &file_io_common_links_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_links_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_io_common_links_proto_rawDescGZIP(), []int{0}
}

func (x *Link) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Link) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Link) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Link) GetType() LinkType {
	if x != nil {
		return x.Type
	}
	return LinkType_URI_DO_NOT_USE
}

func (x *Link) GetLocalizedLink() *LocalizedString {
	if x != nil {
		return x.LocalizedLink
	}
	return nil
}

func (x *Link) GetLocalizedTitle() *LocalizedString {
	if x != nil {
		return x.LocalizedTitle
	}
	return nil
}

func (x *Link) GetUsage() []UsageType {
	if x != nil {
		return x.Usage
	}
	return nil
}

func (x *Link) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

type DbLink struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Links         []*Link                `protobuf:"bytes,1,rep,name=links,proto3" json:"links,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DbLink) Reset() {
	*x = DbLink{}
	mi := &file_io_common_links_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DbLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbLink) ProtoMessage() {}

func (x *DbLink) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_links_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbLink.ProtoReflect.Descriptor instead.
func (*DbLink) Descriptor() ([]byte, []int) {
	return file_io_common_links_proto_rawDescGZIP(), []int{1}
}

func (x *DbLink) GetLinks() []*Link {
	if x != nil {
		return x.Links
	}
	return nil
}

var File_io_common_links_proto protoreflect.FileDescriptor

var file_io_common_links_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x6e, 0x6b,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69, 0x6f, 0x1a, 0x1e, 0x69, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x69, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x02, 0x0a, 0x04, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x69,
	0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0d,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x3b, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x69, 0x6f, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x28, 0x0a, 0x06, 0x44, 0x62, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x1e, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2a,
	0x6b, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x55,
	0x52, 0x49, 0x5f, 0x44, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x52, 0x49, 0x5f, 0x57, 0x45, 0x42, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x52, 0x49, 0x5f, 0x54, 0x45, 0x4c, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x52, 0x49,
	0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x52, 0x49, 0x5f,
	0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x52,
	0x49, 0x5f, 0x43, 0x41, 0x4c, 0x45, 0x4e, 0x44, 0x41, 0x52, 0x10, 0x05, 0x42, 0x47, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5a, 0x24, 0x73, 0x74, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0xaa, 0x02, 0x0c, 0x50, 0x61, 0x73, 0x73, 0x4b, 0x69, 0x74,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_io_common_links_proto_rawDescOnce sync.Once
	file_io_common_links_proto_rawDescData []byte
)

func file_io_common_links_proto_rawDescGZIP() []byte {
	file_io_common_links_proto_rawDescOnce.Do(func() {
		file_io_common_links_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_io_common_links_proto_rawDesc), len(file_io_common_links_proto_rawDesc)))
	})
	return file_io_common_links_proto_rawDescData
}

var file_io_common_links_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_io_common_links_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_io_common_links_proto_goTypes = []any{
	(LinkType)(0),           // 0: io.LinkType
	(*Link)(nil),            // 1: io.Link
	(*DbLink)(nil),          // 2: io.DbLink
	(*LocalizedString)(nil), // 3: io.LocalizedString
	(UsageType)(0),          // 4: io.UsageType
}
var file_io_common_links_proto_depIdxs = []int32{
	0, // 0: io.Link.type:type_name -> io.LinkType
	3, // 1: io.Link.localizedLink:type_name -> io.LocalizedString
	3, // 2: io.Link.localizedTitle:type_name -> io.LocalizedString
	4, // 3: io.Link.usage:type_name -> io.UsageType
	1, // 4: io.DbLink.links:type_name -> io.Link
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_io_common_links_proto_init() }
func file_io_common_links_proto_init() {
	if File_io_common_links_proto != nil {
		return
	}
	file_io_common_common_objects_proto_init()
	file_io_common_localization_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_io_common_links_proto_rawDesc), len(file_io_common_links_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_common_links_proto_goTypes,
		DependencyIndexes: file_io_common_links_proto_depIdxs,
		EnumInfos:         file_io_common_links_proto_enumTypes,
		MessageInfos:      file_io_common_links_proto_msgTypes,
	}.Build()
	File_io_common_links_proto = out.File
	file_io_common_links_proto_goTypes = nil
	file_io_common_links_proto_depIdxs = nil
}
