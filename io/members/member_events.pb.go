//*
// Membership Events
//
// Membership Events allow to track for member events and actions. Powerful when used in combination with scanning solutions like CodeREADr.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.0
// source: io/member/member_events.proto

package members

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

type MemberEvents int32

const (
	// Do not use
	MemberEvents_EVENT_MEMBER_DO_NOT_USE MemberEvents = 0
	// Used when a member is checked in by the checked in endpoint
	MemberEvents_EVENT_MEMBER_CHECKED_IN MemberEvents = 1
	// Used when a member is checked out by the checked out endpoint
	MemberEvents_EVENT_MEMBER_CHECKED_OUT MemberEvents = 2
	// Used when a member is verified by the verify endpoint
	MemberEvents_EVENT_MEMBER_VERIFIED MemberEvents = 3
	// Used when points are earned by the earn points endpoints
	MemberEvents_EVENT_MEMBER_POINTS_EARNED MemberEvents = 4
	// Used when points are added by the burn points endpoints
	MemberEvents_EVENT_MEMBER_POINTS_BURNED MemberEvents = 5
	// Used when points are set by the set points endpoints
	MemberEvents_EVENT_MEMBER_POINTS_SET MemberEvents = 6
	// Used when points the tier is changed
	MemberEvents_EVENT_MEMBER_TIER_CHANGED MemberEvents = 7
)

// Enum value maps for MemberEvents.
var (
	MemberEvents_name = map[int32]string{
		0: "EVENT_MEMBER_DO_NOT_USE",
		1: "EVENT_MEMBER_CHECKED_IN",
		2: "EVENT_MEMBER_CHECKED_OUT",
		3: "EVENT_MEMBER_VERIFIED",
		4: "EVENT_MEMBER_POINTS_EARNED",
		5: "EVENT_MEMBER_POINTS_BURNED",
		6: "EVENT_MEMBER_POINTS_SET",
		7: "EVENT_MEMBER_TIER_CHANGED",
	}
	MemberEvents_value = map[string]int32{
		"EVENT_MEMBER_DO_NOT_USE":    0,
		"EVENT_MEMBER_CHECKED_IN":    1,
		"EVENT_MEMBER_CHECKED_OUT":   2,
		"EVENT_MEMBER_VERIFIED":      3,
		"EVENT_MEMBER_POINTS_EARNED": 4,
		"EVENT_MEMBER_POINTS_BURNED": 5,
		"EVENT_MEMBER_POINTS_SET":    6,
		"EVENT_MEMBER_TIER_CHANGED":  7,
	}
)

func (x MemberEvents) Enum() *MemberEvents {
	p := new(MemberEvents)
	*p = x
	return p
}

func (x MemberEvents) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MemberEvents) Descriptor() protoreflect.EnumDescriptor {
	return file_io_member_member_events_proto_enumTypes[0].Descriptor()
}

func (MemberEvents) Type() protoreflect.EnumType {
	return &file_io_member_member_events_proto_enumTypes[0]
}

func (x MemberEvents) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MemberEvents.Descriptor instead.
func (MemberEvents) EnumDescriptor() ([]byte, []int) {
	return file_io_member_member_events_proto_rawDescGZIP(), []int{0}
}

// The Member Event record
type MemberEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// PassKit generated Event ID. Not writable.
	// @tag: validateCreate:"omitempty"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateCreate:"omitempty"`
	// Member that the event belongs to. This does not contain the full member details, only the key data.
	// @tag: validateCreate:"required"
	Member *MemberMininmal `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty" validateCreate:"required"`
	// Event type.
	// @tag: validateCreate:"required"
	EventType MemberEvents `protobuf:"varint,3,opt,name=eventType,proto3,enum=members.MemberEvents" json:"eventType,omitempty" validateCreate:"required"`
	// Address the event took place.
	// @tag: validateCreate:"omitempty,max=255"
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty" validateCreate:"omitempty,max=255"`
	// Latitude the event took place.
	// @tag: validateCreate:"omitempty"
	Lat float64 `protobuf:"fixed64,5,opt,name=lat,proto3" json:"lat,omitempty" validateCreate:"omitempty"`
	// Longitude the event took place.
	// @tag: validateCreate:"omitempty"
	Lon float64 `protobuf:"fixed64,6,opt,name=lon,proto3" json:"lon,omitempty" validateCreate:"omitempty"`
	// Altitude the event took place (in metres).
	// @tag: validateCreate:"omitempty"
	Alt int32 `protobuf:"varint,7,opt,name=alt,proto3" json:"alt,omitempty" validateCreate:"omitempty"`
	// External unique ID of the event.
	// @tag: validateCreate:"omitempty,max=255"
	ExternalId string `protobuf:"bytes,8,opt,name=externalId,proto3" json:"externalId,omitempty" validateCreate:"omitempty,max=255"`
	// External user ID of the logged in user that captured the event (for example when using an external scanning app).
	// @tag: validateCreate:"omitempty,max=255"
	ExternalUserId string `protobuf:"bytes,9,opt,name=externalUserId,proto3" json:"externalUserId,omitempty" validateCreate:"omitempty,max=255"`
	// External device ID of the device that was used to capture the event (for example when using an external scanning app).
	// @tag: validateCreate:"omitempty,max=255"
	ExternalDeviceId string `protobuf:"bytes,10,opt,name=externalDeviceId,proto3" json:"externalDeviceId,omitempty" validateCreate:"omitempty,max=255"`
	// External service ID of the service that was used for capturing the event (for example when using an external scanning app).
	// @tag: validateCreate:"omitempty,max=255"
	ExternalServiceId string `protobuf:"bytes,11,opt,name=externalServiceId,proto3" json:"externalServiceId,omitempty" validateCreate:"omitempty,max=255"`
	// Any meta data (for example gathered on scanning) that is relevant to the event (# of points earner, bill spent, device meta-data, etc).
	// @tag: validateCreate:"omitempty"
	MetaData map[string]string `protobuf:"bytes,12,rep,name=metaData,proto3" json:"metaData,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value" validateCreate:"omitempty"`
	// Any relevant notes for the event.
	// @tag: validateCreate:"omitempty"
	Notes string `protobuf:"bytes,13,opt,name=notes,proto3" json:"notes,omitempty" validateCreate:"omitempty"`
	// The event date.
	// @tag: validateCreate:"required"
	Date *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=date,proto3" json:"date,omitempty" validateCreate:"required"`
	// The date until the event is retained in the PassKit database.
	// @tag: validateCreate:"omitempty"
	RetainedUntilDate *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=retainedUntilDate,proto3" json:"retainedUntilDate,omitempty" validateCreate:"omitempty"`
	// The date the event record was created in the PassKit database.
	// @tag: validateCreate:"omitempty"
	Created       *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=created,proto3" json:"created,omitempty" validateCreate:"omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MemberEvent) Reset() {
	*x = MemberEvent{}
	mi := &file_io_member_member_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemberEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberEvent) ProtoMessage() {}

func (x *MemberEvent) ProtoReflect() protoreflect.Message {
	mi := &file_io_member_member_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberEvent.ProtoReflect.Descriptor instead.
func (*MemberEvent) Descriptor() ([]byte, []int) {
	return file_io_member_member_events_proto_rawDescGZIP(), []int{0}
}

func (x *MemberEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MemberEvent) GetMember() *MemberMininmal {
	if x != nil {
		return x.Member
	}
	return nil
}

func (x *MemberEvent) GetEventType() MemberEvents {
	if x != nil {
		return x.EventType
	}
	return MemberEvents_EVENT_MEMBER_DO_NOT_USE
}

func (x *MemberEvent) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MemberEvent) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *MemberEvent) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *MemberEvent) GetAlt() int32 {
	if x != nil {
		return x.Alt
	}
	return 0
}

func (x *MemberEvent) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *MemberEvent) GetExternalUserId() string {
	if x != nil {
		return x.ExternalUserId
	}
	return ""
}

func (x *MemberEvent) GetExternalDeviceId() string {
	if x != nil {
		return x.ExternalDeviceId
	}
	return ""
}

func (x *MemberEvent) GetExternalServiceId() string {
	if x != nil {
		return x.ExternalServiceId
	}
	return ""
}

func (x *MemberEvent) GetMetaData() map[string]string {
	if x != nil {
		return x.MetaData
	}
	return nil
}

func (x *MemberEvent) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *MemberEvent) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *MemberEvent) GetRetainedUntilDate() *timestamppb.Timestamp {
	if x != nil {
		return x.RetainedUntilDate
	}
	return nil
}

func (x *MemberEvent) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

type MemberMininmal struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Id assigned by PassKit to represent the member record. It will be used as the serial number in Apple Wallet and as the Object identifier for Google Wallet. This field is not writable.
	// @tag: validateCreate:"required,max=255"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateCreate:"required,max=255"`
	// This can be used to set the 'external' ID of the member (i.e. the member ID as it's being used in your system). If provided then this can be used to query & update members. This field will be treated as unique within the program, and cannot be updated at a later stage.
	ExternalId string `protobuf:"bytes,2,opt,name=externalId,proto3" json:"externalId,omitempty"`
	// Grouping Identifier can be used to group members under the same membership (i.e. couple).
	GroupingIdentifier string `protobuf:"bytes,3,opt,name=groupingIdentifier,proto3" json:"groupingIdentifier,omitempty"`
	// Indicates the ID of the tier this member is in.
	TierId string `protobuf:"bytes,4,opt,name=tierId,proto3" json:"tierId,omitempty"`
	// Indicates the ID of the program this member is in.
	ProgramId string `protobuf:"bytes,5,opt,name=programId,proto3" json:"programId,omitempty"`
	// Personal details of the member.
	Person        *io.Person `protobuf:"bytes,6,opt,name=person,proto3" json:"person,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MemberMininmal) Reset() {
	*x = MemberMininmal{}
	mi := &file_io_member_member_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemberMininmal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberMininmal) ProtoMessage() {}

func (x *MemberMininmal) ProtoReflect() protoreflect.Message {
	mi := &file_io_member_member_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberMininmal.ProtoReflect.Descriptor instead.
func (*MemberMininmal) Descriptor() ([]byte, []int) {
	return file_io_member_member_events_proto_rawDescGZIP(), []int{1}
}

func (x *MemberMininmal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MemberMininmal) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *MemberMininmal) GetGroupingIdentifier() string {
	if x != nil {
		return x.GroupingIdentifier
	}
	return ""
}

func (x *MemberMininmal) GetTierId() string {
	if x != nil {
		return x.TierId
	}
	return ""
}

func (x *MemberMininmal) GetProgramId() string {
	if x != nil {
		return x.ProgramId
	}
	return ""
}

func (x *MemberMininmal) GetPerson() *io.Person {
	if x != nil {
		return x.Person
	}
	return nil
}

var File_io_member_member_events_proto protoreflect.FileDescriptor

var file_io_member_member_events_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x69, 0x6f, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x69, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x06, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x6d, 0x61, 0x6c, 0x52, 0x06, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x09,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x48, 0x0a, 0x11, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x55, 0x6e, 0x74,
	0x69, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x7b,
	0x92, 0x41, 0x78, 0x0a, 0x76, 0x2a, 0x0d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x32, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x74, 0x68, 0x61,
	0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x68, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x20, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x20, 0x41, 0x50, 0x49, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x61, 0x72, 0x65,
	0x20, 0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x22, 0xca, 0x01, 0x0a, 0x0e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x6d, 0x61, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x2e,
	0x0a, 0x12, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x69, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2a, 0xfd, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x44, 0x4f, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x55, 0x53, 0x45, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x49,
	0x4e, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x4d,
	0x42, 0x45, 0x52, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x4f, 0x55, 0x54, 0x10,
	0x02, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45,
	0x52, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x50, 0x4f, 0x49,
	0x4e, 0x54, 0x53, 0x5f, 0x45, 0x41, 0x52, 0x4e, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x50, 0x4f, 0x49,
	0x4e, 0x54, 0x53, 0x5f, 0x42, 0x55, 0x52, 0x4e, 0x45, 0x44, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x50, 0x4f, 0x49,
	0x4e, 0x54, 0x53, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x06, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x07, 0x42, 0x5f, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x5a, 0x2c, 0x73, 0x74, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x61, 0x73, 0x73,
	0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0xaa, 0x02, 0x14, 0x50, 0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70,
	0x63, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_io_member_member_events_proto_rawDescOnce sync.Once
	file_io_member_member_events_proto_rawDescData []byte
)

func file_io_member_member_events_proto_rawDescGZIP() []byte {
	file_io_member_member_events_proto_rawDescOnce.Do(func() {
		file_io_member_member_events_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_io_member_member_events_proto_rawDesc), len(file_io_member_member_events_proto_rawDesc)))
	})
	return file_io_member_member_events_proto_rawDescData
}

var file_io_member_member_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_io_member_member_events_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_io_member_member_events_proto_goTypes = []any{
	(MemberEvents)(0),             // 0: members.MemberEvents
	(*MemberEvent)(nil),           // 1: members.MemberEvent
	(*MemberMininmal)(nil),        // 2: members.MemberMininmal
	nil,                           // 3: members.MemberEvent.MetaDataEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*io.Person)(nil),             // 5: io.Person
}
var file_io_member_member_events_proto_depIdxs = []int32{
	2, // 0: members.MemberEvent.member:type_name -> members.MemberMininmal
	0, // 1: members.MemberEvent.eventType:type_name -> members.MemberEvents
	3, // 2: members.MemberEvent.metaData:type_name -> members.MemberEvent.MetaDataEntry
	4, // 3: members.MemberEvent.date:type_name -> google.protobuf.Timestamp
	4, // 4: members.MemberEvent.retainedUntilDate:type_name -> google.protobuf.Timestamp
	4, // 5: members.MemberEvent.created:type_name -> google.protobuf.Timestamp
	5, // 6: members.MemberMininmal.person:type_name -> io.Person
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_io_member_member_events_proto_init() }
func file_io_member_member_events_proto_init() {
	if File_io_member_member_events_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_io_member_member_events_proto_rawDesc), len(file_io_member_member_events_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_member_member_events_proto_goTypes,
		DependencyIndexes: file_io_member_member_events_proto_depIdxs,
		EnumInfos:         file_io_member_member_events_proto_enumTypes,
		MessageInfos:      file_io_member_member_events_proto_msgTypes,
	}.Build()
	File_io_member_member_events_proto = out.File
	file_io_member_member_events_proto_goTypes = nil
	file_io_member_member_events_proto_depIdxs = nil
}
