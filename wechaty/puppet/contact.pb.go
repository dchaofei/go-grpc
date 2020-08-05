// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: contact.proto

package puppet

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ContactGender int32

const (
	ContactGender_CONTACT_GENDER_UNSPECIFIED ContactGender = 0
	ContactGender_CONTACT_GENDER_MALE        ContactGender = 1
	ContactGender_CONTACT_GENDER_FEMALE      ContactGender = 2
)

// Enum value maps for ContactGender.
var (
	ContactGender_name = map[int32]string{
		0: "CONTACT_GENDER_UNSPECIFIED",
		1: "CONTACT_GENDER_MALE",
		2: "CONTACT_GENDER_FEMALE",
	}
	ContactGender_value = map[string]int32{
		"CONTACT_GENDER_UNSPECIFIED": 0,
		"CONTACT_GENDER_MALE":        1,
		"CONTACT_GENDER_FEMALE":      2,
	}
)

func (x ContactGender) Enum() *ContactGender {
	p := new(ContactGender)
	*p = x
	return p
}

func (x ContactGender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContactGender) Descriptor() protoreflect.EnumDescriptor {
	return file_contact_proto_enumTypes[0].Descriptor()
}

func (ContactGender) Type() protoreflect.EnumType {
	return &file_contact_proto_enumTypes[0]
}

func (x ContactGender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContactGender.Descriptor instead.
func (ContactGender) EnumDescriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{0}
}

type ContactType int32

const (
	ContactType_CONTACT_TYPE_UNSPECIFIED ContactType = 0
	ContactType_CONTACT_TYPE_PERSONAL    ContactType = 1
	ContactType_CONTACT_TYPE_OFFICIAL    ContactType = 2
)

// Enum value maps for ContactType.
var (
	ContactType_name = map[int32]string{
		0: "CONTACT_TYPE_UNSPECIFIED",
		1: "CONTACT_TYPE_PERSONAL",
		2: "CONTACT_TYPE_OFFICIAL",
	}
	ContactType_value = map[string]int32{
		"CONTACT_TYPE_UNSPECIFIED": 0,
		"CONTACT_TYPE_PERSONAL":    1,
		"CONTACT_TYPE_OFFICIAL":    2,
	}
)

func (x ContactType) Enum() *ContactType {
	p := new(ContactType)
	*p = x
	return p
}

func (x ContactType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContactType) Descriptor() protoreflect.EnumDescriptor {
	return file_contact_proto_enumTypes[1].Descriptor()
}

func (ContactType) Type() protoreflect.EnumType {
	return &file_contact_proto_enumTypes[1]
}

func (x ContactType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContactType.Descriptor instead.
func (ContactType) EnumDescriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{1}
}

type ContactListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContactListRequest) Reset() {
	*x = ContactListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactListRequest) ProtoMessage() {}

func (x *ContactListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactListRequest.ProtoReflect.Descriptor instead.
func (*ContactListRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{0}
}

type ContactListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ContactListResponse) Reset() {
	*x = ContactListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactListResponse) ProtoMessage() {}

func (x *ContactListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactListResponse.ProtoReflect.Descriptor instead.
func (*ContactListResponse) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{1}
}

func (x *ContactListResponse) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type ContactPayloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ContactPayloadRequest) Reset() {
	*x = ContactPayloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactPayloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactPayloadRequest) ProtoMessage() {}

func (x *ContactPayloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactPayloadRequest.ProtoReflect.Descriptor instead.
func (*ContactPayloadRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{2}
}

func (x *ContactPayloadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ContactPayloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Gender    ContactGender `protobuf:"varint,2,opt,name=gender,proto3,enum=wechaty.puppet.ContactGender" json:"gender,omitempty"`
	Type      ContactType   `protobuf:"varint,3,opt,name=type,proto3,enum=wechaty.puppet.ContactType" json:"type,omitempty"`
	Name      string        `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Avatar    string        `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Address   string        `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Alias     string        `protobuf:"bytes,7,opt,name=alias,proto3" json:"alias,omitempty"`
	City      string        `protobuf:"bytes,8,opt,name=city,proto3" json:"city,omitempty"`
	Friend    bool          `protobuf:"varint,9,opt,name=friend,proto3" json:"friend,omitempty"`
	Province  string        `protobuf:"bytes,10,opt,name=province,proto3" json:"province,omitempty"`
	Signature string        `protobuf:"bytes,11,opt,name=signature,proto3" json:"signature,omitempty"`
	Star      bool          `protobuf:"varint,12,opt,name=star,proto3" json:"star,omitempty"`
	Weixin    string        `protobuf:"bytes,13,opt,name=weixin,proto3" json:"weixin,omitempty"`
}

func (x *ContactPayloadResponse) Reset() {
	*x = ContactPayloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactPayloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactPayloadResponse) ProtoMessage() {}

func (x *ContactPayloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactPayloadResponse.ProtoReflect.Descriptor instead.
func (*ContactPayloadResponse) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{3}
}

func (x *ContactPayloadResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContactPayloadResponse) GetGender() ContactGender {
	if x != nil {
		return x.Gender
	}
	return ContactGender_CONTACT_GENDER_UNSPECIFIED
}

func (x *ContactPayloadResponse) GetType() ContactType {
	if x != nil {
		return x.Type
	}
	return ContactType_CONTACT_TYPE_UNSPECIFIED
}

func (x *ContactPayloadResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContactPayloadResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *ContactPayloadResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ContactPayloadResponse) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *ContactPayloadResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *ContactPayloadResponse) GetFriend() bool {
	if x != nil {
		return x.Friend
	}
	return false
}

func (x *ContactPayloadResponse) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *ContactPayloadResponse) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *ContactPayloadResponse) GetStar() bool {
	if x != nil {
		return x.Star
	}
	return false
}

func (x *ContactPayloadResponse) GetWeixin() string {
	if x != nil {
		return x.Weixin
	}
	return ""
}

type ContactSelfQRCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContactSelfQRCodeRequest) Reset() {
	*x = ContactSelfQRCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactSelfQRCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactSelfQRCodeRequest) ProtoMessage() {}

func (x *ContactSelfQRCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactSelfQRCodeRequest.ProtoReflect.Descriptor instead.
func (*ContactSelfQRCodeRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{4}
}

type ContactSelfQRCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qrcode string `protobuf:"bytes,1,opt,name=qrcode,proto3" json:"qrcode,omitempty"`
}

func (x *ContactSelfQRCodeResponse) Reset() {
	*x = ContactSelfQRCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactSelfQRCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactSelfQRCodeResponse) ProtoMessage() {}

func (x *ContactSelfQRCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactSelfQRCodeResponse.ProtoReflect.Descriptor instead.
func (*ContactSelfQRCodeResponse) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{5}
}

func (x *ContactSelfQRCodeResponse) GetQrcode() string {
	if x != nil {
		return x.Qrcode
	}
	return ""
}

type ContactSelfNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ContactSelfNameRequest) Reset() {
	*x = ContactSelfNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactSelfNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactSelfNameRequest) ProtoMessage() {}

func (x *ContactSelfNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactSelfNameRequest.ProtoReflect.Descriptor instead.
func (*ContactSelfNameRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{6}
}

func (x *ContactSelfNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ContactSelfNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContactSelfNameResponse) Reset() {
	*x = ContactSelfNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactSelfNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactSelfNameResponse) ProtoMessage() {}

func (x *ContactSelfNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactSelfNameResponse.ProtoReflect.Descriptor instead.
func (*ContactSelfNameResponse) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{7}
}

type ContactSelfSignatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ContactSelfSignatureRequest) Reset() {
	*x = ContactSelfSignatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactSelfSignatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactSelfSignatureRequest) ProtoMessage() {}

func (x *ContactSelfSignatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactSelfSignatureRequest.ProtoReflect.Descriptor instead.
func (*ContactSelfSignatureRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{8}
}

func (x *ContactSelfSignatureRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type ContactSelfSignatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContactSelfSignatureResponse) Reset() {
	*x = ContactSelfSignatureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactSelfSignatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactSelfSignatureResponse) ProtoMessage() {}

func (x *ContactSelfSignatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactSelfSignatureResponse.ProtoReflect.Descriptor instead.
func (*ContactSelfSignatureResponse) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{9}
}

type ContactAliasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// nullable
	Alias *wrappers.StringValue `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
}

func (x *ContactAliasRequest) Reset() {
	*x = ContactAliasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactAliasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactAliasRequest) ProtoMessage() {}

func (x *ContactAliasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactAliasRequest.ProtoReflect.Descriptor instead.
func (*ContactAliasRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{10}
}

func (x *ContactAliasRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContactAliasRequest) GetAlias() *wrappers.StringValue {
	if x != nil {
		return x.Alias
	}
	return nil
}

type ContactAliasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alias *wrappers.StringValue `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
}

func (x *ContactAliasResponse) Reset() {
	*x = ContactAliasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactAliasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactAliasResponse) ProtoMessage() {}

func (x *ContactAliasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactAliasResponse.ProtoReflect.Descriptor instead.
func (*ContactAliasResponse) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{11}
}

func (x *ContactAliasResponse) GetAlias() *wrappers.StringValue {
	if x != nil {
		return x.Alias
	}
	return nil
}

type ContactAvatarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// nullable
	Filebox *wrappers.StringValue `protobuf:"bytes,2,opt,name=filebox,proto3" json:"filebox,omitempty"`
}

func (x *ContactAvatarRequest) Reset() {
	*x = ContactAvatarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactAvatarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactAvatarRequest) ProtoMessage() {}

func (x *ContactAvatarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactAvatarRequest.ProtoReflect.Descriptor instead.
func (*ContactAvatarRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{12}
}

func (x *ContactAvatarRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContactAvatarRequest) GetFilebox() *wrappers.StringValue {
	if x != nil {
		return x.Filebox
	}
	return nil
}

type ContactAvatarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filebox *wrappers.StringValue `protobuf:"bytes,1,opt,name=filebox,proto3" json:"filebox,omitempty"`
}

func (x *ContactAvatarResponse) Reset() {
	*x = ContactAvatarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactAvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactAvatarResponse) ProtoMessage() {}

func (x *ContactAvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactAvatarResponse.ProtoReflect.Descriptor instead.
func (*ContactAvatarResponse) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{13}
}

func (x *ContactAvatarResponse) GetFilebox() *wrappers.StringValue {
	if x != nil {
		return x.Filebox
	}
	return nil
}

var File_contact_proto protoreflect.FileDescriptor

var file_contact_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2e, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x14, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x27,
	0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xfe, 0x02, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x35, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2e, 0x70, 0x75, 0x70,
	0x70, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x47, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74,
	0x79, 0x2e, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x74, 0x61, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x74, 0x61, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x22, 0x1a, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x53, 0x65, 0x6c, 0x66, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53,
	0x65, 0x6c, 0x66, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2c, 0x0a, 0x16, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x53, 0x65, 0x6c, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x53, 0x65, 0x6c, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x3b, 0x0a, 0x1b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x65, 0x6c,
	0x66, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22,
	0x1e, 0x0a, 0x1c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x65, 0x6c, 0x66, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x59, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0x4a, 0x0a, 0x14, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0x5e, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36,
	0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x6f, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x62, 0x6f, 0x78, 0x22, 0x4f, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07,
	0x66, 0x69, 0x6c, 0x65, 0x62, 0x6f, 0x78, 0x2a, 0x63, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4e, 0x54,
	0x41, 0x43, 0x54, 0x5f, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4e, 0x54,
	0x41, 0x43, 0x54, 0x5f, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10,
	0x01, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x5f, 0x47, 0x45, 0x4e,
	0x44, 0x45, 0x52, 0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x2a, 0x61, 0x0a, 0x0b,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x43,
	0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4e,
	0x54, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e,
	0x41, 0x4c, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x46, 0x46, 0x49, 0x43, 0x49, 0x41, 0x4c, 0x10, 0x02, 0x42,
	0x4a, 0x0a, 0x1d, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x77, 0x65, 0x63,
	0x68, 0x61, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x63,
	0x68, 0x61, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x65, 0x63,
	0x68, 0x61, 0x74, 0x79, 0x2f, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_contact_proto_rawDescOnce sync.Once
	file_contact_proto_rawDescData = file_contact_proto_rawDesc
)

func file_contact_proto_rawDescGZIP() []byte {
	file_contact_proto_rawDescOnce.Do(func() {
		file_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_contact_proto_rawDescData)
	})
	return file_contact_proto_rawDescData
}

var file_contact_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_contact_proto_goTypes = []interface{}{
	(ContactGender)(0),                   // 0: wechaty.puppet.ContactGender
	(ContactType)(0),                     // 1: wechaty.puppet.ContactType
	(*ContactListRequest)(nil),           // 2: wechaty.puppet.ContactListRequest
	(*ContactListResponse)(nil),          // 3: wechaty.puppet.ContactListResponse
	(*ContactPayloadRequest)(nil),        // 4: wechaty.puppet.ContactPayloadRequest
	(*ContactPayloadResponse)(nil),       // 5: wechaty.puppet.ContactPayloadResponse
	(*ContactSelfQRCodeRequest)(nil),     // 6: wechaty.puppet.ContactSelfQRCodeRequest
	(*ContactSelfQRCodeResponse)(nil),    // 7: wechaty.puppet.ContactSelfQRCodeResponse
	(*ContactSelfNameRequest)(nil),       // 8: wechaty.puppet.ContactSelfNameRequest
	(*ContactSelfNameResponse)(nil),      // 9: wechaty.puppet.ContactSelfNameResponse
	(*ContactSelfSignatureRequest)(nil),  // 10: wechaty.puppet.ContactSelfSignatureRequest
	(*ContactSelfSignatureResponse)(nil), // 11: wechaty.puppet.ContactSelfSignatureResponse
	(*ContactAliasRequest)(nil),          // 12: wechaty.puppet.ContactAliasRequest
	(*ContactAliasResponse)(nil),         // 13: wechaty.puppet.ContactAliasResponse
	(*ContactAvatarRequest)(nil),         // 14: wechaty.puppet.ContactAvatarRequest
	(*ContactAvatarResponse)(nil),        // 15: wechaty.puppet.ContactAvatarResponse
	(*wrappers.StringValue)(nil),         // 16: google.protobuf.StringValue
}
var file_contact_proto_depIdxs = []int32{
	0,  // 0: wechaty.puppet.ContactPayloadResponse.gender:type_name -> wechaty.puppet.ContactGender
	1,  // 1: wechaty.puppet.ContactPayloadResponse.type:type_name -> wechaty.puppet.ContactType
	16, // 2: wechaty.puppet.ContactAliasRequest.alias:type_name -> google.protobuf.StringValue
	16, // 3: wechaty.puppet.ContactAliasResponse.alias:type_name -> google.protobuf.StringValue
	16, // 4: wechaty.puppet.ContactAvatarRequest.filebox:type_name -> google.protobuf.StringValue
	16, // 5: wechaty.puppet.ContactAvatarResponse.filebox:type_name -> google.protobuf.StringValue
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_contact_proto_init() }
func file_contact_proto_init() {
	if File_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactListRequest); i {
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
		file_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactListResponse); i {
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
		file_contact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactPayloadRequest); i {
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
		file_contact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactPayloadResponse); i {
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
		file_contact_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactSelfQRCodeRequest); i {
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
		file_contact_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactSelfQRCodeResponse); i {
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
		file_contact_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactSelfNameRequest); i {
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
		file_contact_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactSelfNameResponse); i {
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
		file_contact_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactSelfSignatureRequest); i {
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
		file_contact_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactSelfSignatureResponse); i {
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
		file_contact_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactAliasRequest); i {
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
		file_contact_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactAliasResponse); i {
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
		file_contact_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactAvatarRequest); i {
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
		file_contact_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactAvatarResponse); i {
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
			RawDescriptor: file_contact_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contact_proto_goTypes,
		DependencyIndexes: file_contact_proto_depIdxs,
		EnumInfos:         file_contact_proto_enumTypes,
		MessageInfos:      file_contact_proto_msgTypes,
	}.Build()
	File_contact_proto = out.File
	file_contact_proto_rawDesc = nil
	file_contact_proto_goTypes = nil
	file_contact_proto_depIdxs = nil
}
