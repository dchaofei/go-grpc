// Code generated by protoc-gen-go. DO NOT EDIT.
// source: puppet/file_box.proto

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

type FileBoxChunk struct {
	// Types that are valid to be assigned to Payload:
	//	*FileBoxChunk_Data
	//	*FileBoxChunk_Name
	Payload              isFileBoxChunk_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *FileBoxChunk) Reset()         { *m = FileBoxChunk{} }
func (m *FileBoxChunk) String() string { return proto.CompactTextString(m) }
func (*FileBoxChunk) ProtoMessage()    {}
func (*FileBoxChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7afdbae050741bb, []int{0}
}

func (m *FileBoxChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileBoxChunk.Unmarshal(m, b)
}
func (m *FileBoxChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileBoxChunk.Marshal(b, m, deterministic)
}
func (m *FileBoxChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileBoxChunk.Merge(m, src)
}
func (m *FileBoxChunk) XXX_Size() int {
	return xxx_messageInfo_FileBoxChunk.Size(m)
}
func (m *FileBoxChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_FileBoxChunk.DiscardUnknown(m)
}

var xxx_messageInfo_FileBoxChunk proto.InternalMessageInfo

type isFileBoxChunk_Payload interface {
	isFileBoxChunk_Payload()
}

type FileBoxChunk_Data struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3,oneof"`
}

type FileBoxChunk_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*FileBoxChunk_Data) isFileBoxChunk_Payload() {}

func (*FileBoxChunk_Name) isFileBoxChunk_Payload() {}

func (m *FileBoxChunk) GetPayload() isFileBoxChunk_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *FileBoxChunk) GetData() []byte {
	if x, ok := m.GetPayload().(*FileBoxChunk_Data); ok {
		return x.Data
	}
	return nil
}

func (m *FileBoxChunk) GetName() string {
	if x, ok := m.GetPayload().(*FileBoxChunk_Name); ok {
		return x.Name
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FileBoxChunk) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FileBoxChunk_Data)(nil),
		(*FileBoxChunk_Name)(nil),
	}
}

func init() {
	proto.RegisterType((*FileBoxChunk)(nil), "wechaty.puppet.FileBoxChunk")
}

func init() { proto.RegisterFile("puppet/file_box.proto", fileDescriptor_a7afdbae050741bb) }

var fileDescriptor_a7afdbae050741bb = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0x2d, 0x28,
	0x48, 0x2d, 0xd1, 0x4f, 0xcb, 0xcc, 0x49, 0x8d, 0x4f, 0xca, 0xaf, 0xd0, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x2b, 0x4f, 0x4d, 0xce, 0x48, 0x2c, 0xa9, 0xd4, 0x83, 0x48, 0x2b, 0xb9, 0x72,
	0xf1, 0xb8, 0x65, 0xe6, 0xa4, 0x3a, 0xe5, 0x57, 0x38, 0x67, 0x94, 0xe6, 0x65, 0x0b, 0x89, 0x70,
	0xb1, 0xa4, 0x24, 0x96, 0x24, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x78, 0x30, 0x04, 0x81, 0x79,
	0x20, 0xd1, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x4e, 0x90, 0x28, 0x88, 0xe7,
	0xc4, 0xc9, 0xc5, 0x5e, 0x90, 0x58, 0x99, 0x93, 0x9f, 0x98, 0xe2, 0xa4, 0x1d, 0xa5, 0x99, 0x9e,
	0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0xb5, 0x43, 0x3f, 0x3d, 0x5f, 0x37,
	0xbd, 0xa8, 0x20, 0x19, 0xce, 0x87, 0xd8, 0x99, 0xc4, 0x06, 0x76, 0x8a, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xc2, 0x5d, 0x20, 0xbb, 0xa3, 0x00, 0x00, 0x00,
}
