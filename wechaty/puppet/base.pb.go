// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package puppet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StartRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}

func (m *StartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartRequest.Unmarshal(m, b)
}
func (m *StartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartRequest.Marshal(b, m, deterministic)
}
func (m *StartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartRequest.Merge(m, src)
}
func (m *StartRequest) XXX_Size() int {
	return xxx_messageInfo_StartRequest.Size(m)
}
func (m *StartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartRequest proto.InternalMessageInfo

type StartResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}
func (*StartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{1}
}

func (m *StartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResponse.Unmarshal(m, b)
}
func (m *StartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResponse.Marshal(b, m, deterministic)
}
func (m *StartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResponse.Merge(m, src)
}
func (m *StartResponse) XXX_Size() int {
	return xxx_messageInfo_StartResponse.Size(m)
}
func (m *StartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartResponse proto.InternalMessageInfo

type StopRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}
func (*StopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{2}
}

func (m *StopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRequest.Unmarshal(m, b)
}
func (m *StopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRequest.Marshal(b, m, deterministic)
}
func (m *StopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRequest.Merge(m, src)
}
func (m *StopRequest) XXX_Size() int {
	return xxx_messageInfo_StopRequest.Size(m)
}
func (m *StopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopRequest proto.InternalMessageInfo

type StopResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}
func (*StopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{3}
}

func (m *StopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopResponse.Unmarshal(m, b)
}
func (m *StopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopResponse.Marshal(b, m, deterministic)
}
func (m *StopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopResponse.Merge(m, src)
}
func (m *StopResponse) XXX_Size() int {
	return xxx_messageInfo_StopResponse.Size(m)
}
func (m *StopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopResponse proto.InternalMessageInfo

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{4}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{5}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type LogoutRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{6}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

type LogoutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{7}
}

func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (m *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(m, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

type DingRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DingRequest) Reset()         { *m = DingRequest{} }
func (m *DingRequest) String() string { return proto.CompactTextString(m) }
func (*DingRequest) ProtoMessage()    {}
func (*DingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{8}
}

func (m *DingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DingRequest.Unmarshal(m, b)
}
func (m *DingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DingRequest.Marshal(b, m, deterministic)
}
func (m *DingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DingRequest.Merge(m, src)
}
func (m *DingRequest) XXX_Size() int {
	return xxx_messageInfo_DingRequest.Size(m)
}
func (m *DingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DingRequest proto.InternalMessageInfo

func (m *DingRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type DingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DingResponse) Reset()         { *m = DingResponse{} }
func (m *DingResponse) String() string { return proto.CompactTextString(m) }
func (*DingResponse) ProtoMessage()    {}
func (*DingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{9}
}

func (m *DingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DingResponse.Unmarshal(m, b)
}
func (m *DingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DingResponse.Marshal(b, m, deterministic)
}
func (m *DingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DingResponse.Merge(m, src)
}
func (m *DingResponse) XXX_Size() int {
	return xxx_messageInfo_DingResponse.Size(m)
}
func (m *DingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StartRequest)(nil), "wechaty.puppet.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "wechaty.puppet.StartResponse")
	proto.RegisterType((*StopRequest)(nil), "wechaty.puppet.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "wechaty.puppet.StopResponse")
	proto.RegisterType((*VersionRequest)(nil), "wechaty.puppet.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "wechaty.puppet.VersionResponse")
	proto.RegisterType((*LogoutRequest)(nil), "wechaty.puppet.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "wechaty.puppet.LogoutResponse")
	proto.RegisterType((*DingRequest)(nil), "wechaty.puppet.DingRequest")
	proto.RegisterType((*DingResponse)(nil), "wechaty.puppet.DingResponse")
}

func init() {
	proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150)
}

var fileDescriptor_db1b6b0986796150 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x41, 0x4f, 0xc4, 0x20,
	0x10, 0x85, 0xb3, 0x89, 0xd1, 0x38, 0xeb, 0xb2, 0x1b, 0x4e, 0x3d, 0x6a, 0x4f, 0x8d, 0x26, 0xe5,
	0xe0, 0x3f, 0x30, 0x1e, 0x3d, 0xb5, 0x89, 0x07, 0x6f, 0xb4, 0x4e, 0x28, 0x07, 0x3b, 0x08, 0x83,
	0xc6, 0x7f, 0x6f, 0x28, 0xd0, 0x1b, 0xef, 0xcd, 0x07, 0x8f, 0x79, 0x00, 0x93, 0x0e, 0xd8, 0x3b,
	0x4f, 0x4c, 0x52, 0xfc, 0xe2, 0xbc, 0x68, 0xfe, 0xeb, 0x5d, 0x74, 0x0e, 0xb9, 0x15, 0x70, 0x37,
	0xb2, 0xf6, 0x3c, 0xe0, 0x77, 0xc4, 0xc0, 0xed, 0x19, 0x4e, 0x45, 0x07, 0x47, 0x6b, 0xc0, 0xf6,
	0x04, 0xc7, 0x91, 0xc9, 0xd5, 0xf9, 0xc6, 0x27, 0x59, 0xc6, 0x17, 0x10, 0xef, 0xe8, 0x83, 0xa5,
	0xb5, 0x12, 0x4f, 0x70, 0xde, 0x9d, 0x0c, 0xc9, 0x06, 0x6e, 0x7e, 0xb2, 0xd5, 0x1c, 0xee, 0x0f,
	0xdd, 0xed, 0x50, 0x65, 0x8a, 0x7b, 0x23, 0x43, 0x71, 0xcf, 0xbf, 0x80, 0xa8, 0x46, 0x49, 0x78,
	0x80, 0xe3, 0xab, 0x5d, 0x4d, 0x01, 0xa4, 0x84, 0xab, 0x4f, 0xcd, 0xba, 0x3c, 0xb4, 0x9d, 0xd3,
	0xa7, 0x32, 0x92, 0xaf, 0xbc, 0x3c, 0x7e, 0x74, 0xc6, 0xf2, 0x12, 0xa7, 0x7e, 0xa6, 0x2f, 0x95,
	0xf6, 0xb5, 0xa8, 0x8c, 0x77, 0xb3, 0x32, 0xa4, 0x4a, 0x01, 0x2a, 0x17, 0x30, 0x5d, 0x6f, 0xbd,
	0x3c, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x27, 0x4a, 0xb6, 0xa0, 0x25, 0x01, 0x00, 0x00,
}
