//*
// Ticket Type holds details about the ticket type, and links to the before & after redeem Pass Template Designs.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.0
// source: io/event_tickets/ticket_type.proto

package event_tickets

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	io "github.com/PassKit/passkit-golang-grpc-sdk/io"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// The Ticket Type Details
type TicketType struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// PassKit generated ticket type id (22 characters).
	// @tag: validateGeneric:"required_without=Uid" validateCreate:"isdefault" validateUpdate:"required_without=Uid"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateGeneric:"required_without=Uid" validateCreate:"isdefault" validateUpdate:"required_without=Uid"`
	// User generated ticket type id; unique within the Production.
	// @tag: validateGeneric:"required_without=Id" validateCreate:"omitempty" validateUpdate:"required_without=Id"
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty" validateGeneric:"required_without=Id" validateCreate:"omitempty" validateUpdate:"required_without=Id"`
	// The Production the ticket type belongs to
	// @tag: validateGeneric:"required_without=Id" validateCreate:"required" validateUpdate:"required_without=Id"
	ProductionId string `protobuf:"bytes,3,opt,name=productionId,proto3" json:"productionId,omitempty" validateGeneric:"required_without=Id" validateCreate:"required" validateUpdate:"required_without=Id"`
	// Name of the ticket type. Internal use only.
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"`
	// Localized name of the ticket type.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	LocalizedName *io.LocalizedString `protobuf:"bytes,5,opt,name=localizedName,proto3" json:"localizedName,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Terms & conditions specifically for this ticket type.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	TicketTypeConditions string `protobuf:"bytes,6,opt,name=ticketTypeConditions,proto3" json:"ticketTypeConditions,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Localized ticket conditions for this ticket type.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	LocalizedTicketTypeConditions *io.LocalizedString `protobuf:"bytes,7,opt,name=localizedTicketTypeConditions,proto3" json:"localizedTicketTypeConditions,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// The pass template design ID that tickets will use when initially issued.
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"
	BeforeRedeemPassTemplateId string `protobuf:"bytes,8,opt,name=beforeRedeemPassTemplateId,proto3" json:"beforeRedeemPassTemplateId,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"`
	// Optional pass template ID that tickets will use after the ticket holder checked into the venue (when the ticket switches to its redeemed state).
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	AfterRedeemPassTemplateId string `protobuf:"bytes,9,opt,name=afterRedeemPassTemplateId,proto3" json:"afterRedeemPassTemplateId,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// The date the ticket type was created. Cannot be set via the API.
	// @tag: validateGeneric:"isdefault" validateCreate:"isdefault" validateUpdate:"isdefault"
	Created *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created,proto3" json:"created,omitempty" validateGeneric:"isdefault" validateCreate:"isdefault" validateUpdate:"isdefault"`
	// The date the ticket type updated. Cannot be set via the API.
	// @tag: validateGeneric:"isdefault" validateCreate:"isdefault" validateUpdate:"isdefault"
	Updated       *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated,proto3" json:"updated,omitempty" validateGeneric:"isdefault" validateCreate:"isdefault" validateUpdate:"isdefault"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TicketType) Reset() {
	*x = TicketType{}
	mi := &file_io_event_tickets_ticket_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TicketType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketType) ProtoMessage() {}

func (x *TicketType) ProtoReflect() protoreflect.Message {
	mi := &file_io_event_tickets_ticket_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketType.ProtoReflect.Descriptor instead.
func (*TicketType) Descriptor() ([]byte, []int) {
	return file_io_event_tickets_ticket_type_proto_rawDescGZIP(), []int{0}
}

func (x *TicketType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TicketType) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *TicketType) GetProductionId() string {
	if x != nil {
		return x.ProductionId
	}
	return ""
}

func (x *TicketType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TicketType) GetLocalizedName() *io.LocalizedString {
	if x != nil {
		return x.LocalizedName
	}
	return nil
}

func (x *TicketType) GetTicketTypeConditions() string {
	if x != nil {
		return x.TicketTypeConditions
	}
	return ""
}

func (x *TicketType) GetLocalizedTicketTypeConditions() *io.LocalizedString {
	if x != nil {
		return x.LocalizedTicketTypeConditions
	}
	return nil
}

func (x *TicketType) GetBeforeRedeemPassTemplateId() string {
	if x != nil {
		return x.BeforeRedeemPassTemplateId
	}
	return ""
}

func (x *TicketType) GetAfterRedeemPassTemplateId() string {
	if x != nil {
		return x.AfterRedeemPassTemplateId
	}
	return ""
}

func (x *TicketType) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *TicketType) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

type GetByUidRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The production id of the ticket type.
	// @tag: validateGeneric:"required" validateCreate:"required" validateUpdate:"required"
	ProductionId string `protobuf:"bytes,1,opt,name=productionId,proto3" json:"productionId,omitempty" validateGeneric:"required" validateCreate:"required" validateUpdate:"required"`
	// User defined id of the ticket type.
	// @tag: validateGeneric:"required" validateCreate:"required" validateUpdate:"required"
	Uid           string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty" validateGeneric:"required" validateCreate:"required" validateUpdate:"required"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetByUidRequest) Reset() {
	*x = GetByUidRequest{}
	mi := &file_io_event_tickets_ticket_type_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetByUidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUidRequest) ProtoMessage() {}

func (x *GetByUidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_io_event_tickets_ticket_type_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByUidRequest.ProtoReflect.Descriptor instead.
func (*GetByUidRequest) Descriptor() ([]byte, []int) {
	return file_io_event_tickets_ticket_type_proto_rawDescGZIP(), []int{1}
}

func (x *GetByUidRequest) GetProductionId() string {
	if x != nil {
		return x.ProductionId
	}
	return ""
}

func (x *GetByUidRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type TicketTypeLimitedFields struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// PassKit generated ticket type id (22 characters).
	// @tag: validateGeneric:"omitempty" validateCreate:"isdefault" validateUpdate:"required_without=Uid"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateGeneric:"omitempty" validateCreate:"isdefault" validateUpdate:"required_without=Uid"`
	// User generated ticket type id; unique within the Production.
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required_without=Id"
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required_without=Id"`
	// Name of the ticket type. Internal use only.
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"omitempty"
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TicketTypeLimitedFields) Reset() {
	*x = TicketTypeLimitedFields{}
	mi := &file_io_event_tickets_ticket_type_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TicketTypeLimitedFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketTypeLimitedFields) ProtoMessage() {}

func (x *TicketTypeLimitedFields) ProtoReflect() protoreflect.Message {
	mi := &file_io_event_tickets_ticket_type_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketTypeLimitedFields.ProtoReflect.Descriptor instead.
func (*TicketTypeLimitedFields) Descriptor() ([]byte, []int) {
	return file_io_event_tickets_ticket_type_proto_rawDescGZIP(), []int{2}
}

func (x *TicketTypeLimitedFields) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TicketTypeLimitedFields) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *TicketTypeLimitedFields) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TicketTypeListRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// @tag: validateGeneric:"required" validateCreate:"required" validateUpdate:"required"
	ProductionId string `protobuf:"bytes,1,opt,name=productionId,proto3" json:"productionId,omitempty" validateGeneric:"required" validateCreate:"required" validateUpdate:"required"`
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	Filters       *io.Filters `protobuf:"bytes,2,opt,name=filters,proto3" json:"filters,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TicketTypeListRequest) Reset() {
	*x = TicketTypeListRequest{}
	mi := &file_io_event_tickets_ticket_type_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TicketTypeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketTypeListRequest) ProtoMessage() {}

func (x *TicketTypeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_io_event_tickets_ticket_type_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketTypeListRequest.ProtoReflect.Descriptor instead.
func (*TicketTypeListRequest) Descriptor() ([]byte, []int) {
	return file_io_event_tickets_ticket_type_proto_rawDescGZIP(), []int{3}
}

func (x *TicketTypeListRequest) GetProductionId() string {
	if x != nil {
		return x.ProductionId
	}
	return ""
}

func (x *TicketTypeListRequest) GetFilters() *io.Filters {
	if x != nil {
		return x.Filters
	}
	return nil
}

var File_io_event_tickets_ticket_type_proto protoreflect.FileDescriptor

var file_io_event_tickets_ticket_type_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x69, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x05, 0x0a, 0x0a, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x92, 0x41, 0x02, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x59, 0x0a, 0x1d, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x1d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3e, 0x0a, 0x1a, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x64, 0x65, 0x65, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x19, 0x61, 0x66, 0x74, 0x65, 0x72, 0x52, 0x65, 0x64,
	0x65, 0x65, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x61, 0x66, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x64, 0x65, 0x65, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x05, 0x92, 0x41, 0x02, 0x40, 0x01, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x3b, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x05, 0x92, 0x41,
	0x02, 0x40, 0x01, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x3a, 0xb7, 0x01, 0x92,
	0x41, 0xb3, 0x01, 0x0a, 0xb0, 0x01, 0x2a, 0x0b, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x54,
	0x79, 0x70, 0x65, 0x32, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x54, 0x79, 0x70, 0x65,
	0x20, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x20, 0x61,
	0x62, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x20,
	0x74, 0x79, 0x70, 0x65, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x20,
	0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x20, 0x26, 0x20,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x20, 0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x20, 0x50, 0x61, 0x73,
	0x73, 0x20, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x44, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x73, 0x2e, 0xd2, 0x01, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0xd2, 0x01, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x1a, 0x62, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55,
	0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0xe2, 0x01, 0x0a, 0x17, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x3a, 0x90, 0x01, 0x92, 0x41, 0x8c, 0x01, 0x0a, 0x89, 0x01, 0x2a, 0x17, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x20, 0x54, 0x79, 0x70, 0x65, 0x20, 0x28, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x20,
	0x76, 0x65, 0x72, 0x29, 0x32, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x54, 0x79, 0x70,
	0x65, 0x20, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x20,
	0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x20, 0x74, 0x79, 0x70, 0x65, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x20, 0x26,
	0x20, 0x61, 0x66, 0x74, 0x65, 0x72, 0x20, 0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x20, 0x50, 0x61,
	0x73, 0x73, 0x20, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x44, 0x65, 0x73, 0x69,
	0x67, 0x6e, 0x73, 0x2e, 0x22, 0x62, 0x0a, 0x15, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x42, 0x6f, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5a, 0x32, 0x73, 0x74, 0x61, 0x73, 0x68,
	0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0xaa, 0x02, 0x19,
	0x50, 0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_io_event_tickets_ticket_type_proto_rawDescOnce sync.Once
	file_io_event_tickets_ticket_type_proto_rawDescData []byte
)

func file_io_event_tickets_ticket_type_proto_rawDescGZIP() []byte {
	file_io_event_tickets_ticket_type_proto_rawDescOnce.Do(func() {
		file_io_event_tickets_ticket_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_io_event_tickets_ticket_type_proto_rawDesc), len(file_io_event_tickets_ticket_type_proto_rawDesc)))
	})
	return file_io_event_tickets_ticket_type_proto_rawDescData
}

var file_io_event_tickets_ticket_type_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_io_event_tickets_ticket_type_proto_goTypes = []any{
	(*TicketType)(nil),              // 0: event_tickets.TicketType
	(*GetByUidRequest)(nil),         // 1: event_tickets.GetByUidRequest
	(*TicketTypeLimitedFields)(nil), // 2: event_tickets.TicketTypeLimitedFields
	(*TicketTypeListRequest)(nil),   // 3: event_tickets.TicketTypeListRequest
	(*io.LocalizedString)(nil),      // 4: io.LocalizedString
	(*timestamppb.Timestamp)(nil),   // 5: google.protobuf.Timestamp
	(*io.Filters)(nil),              // 6: io.Filters
}
var file_io_event_tickets_ticket_type_proto_depIdxs = []int32{
	4, // 0: event_tickets.TicketType.localizedName:type_name -> io.LocalizedString
	4, // 1: event_tickets.TicketType.localizedTicketTypeConditions:type_name -> io.LocalizedString
	5, // 2: event_tickets.TicketType.created:type_name -> google.protobuf.Timestamp
	5, // 3: event_tickets.TicketType.updated:type_name -> google.protobuf.Timestamp
	6, // 4: event_tickets.TicketTypeListRequest.filters:type_name -> io.Filters
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_io_event_tickets_ticket_type_proto_init() }
func file_io_event_tickets_ticket_type_proto_init() {
	if File_io_event_tickets_ticket_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_io_event_tickets_ticket_type_proto_rawDesc), len(file_io_event_tickets_ticket_type_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_event_tickets_ticket_type_proto_goTypes,
		DependencyIndexes: file_io_event_tickets_ticket_type_proto_depIdxs,
		MessageInfos:      file_io_event_tickets_ticket_type_proto_msgTypes,
	}.Build()
	File_io_event_tickets_ticket_type_proto = out.File
	file_io_event_tickets_ticket_type_proto_goTypes = nil
	file_io_event_tickets_ticket_type_proto_depIdxs = nil
}
