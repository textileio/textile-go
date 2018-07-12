// Code generated by protoc-gen-go. DO NOT EDIT.
// source: thread_block.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ThreadInvite_Type int32

const (
	ThreadInvite_INTERNAL ThreadInvite_Type = 0
	ThreadInvite_EXTERNAL ThreadInvite_Type = 1
)

var ThreadInvite_Type_name = map[int32]string{
	0: "INTERNAL",
	1: "EXTERNAL",
}
var ThreadInvite_Type_value = map[string]int32{
	"INTERNAL": 0,
	"EXTERNAL": 1,
}

func (x ThreadInvite_Type) String() string {
	return proto.EnumName(ThreadInvite_Type_name, int32(x))
}
func (ThreadInvite_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_thread_block_93e8a7803ac50cdd, []int{2, 0}
}

type ThreadBlock struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Target               string               `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Parents              []string             `protobuf:"bytes,3,rep,name=parents,proto3" json:"parents,omitempty"`
	TargetKey            []byte               `protobuf:"bytes,4,opt,name=targetKey,proto3" json:"targetKey,omitempty"`
	ThreadPubKey         string               `protobuf:"bytes,5,opt,name=threadPubKey,proto3" json:"threadPubKey,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ThreadBlock) Reset()         { *m = ThreadBlock{} }
func (m *ThreadBlock) String() string { return proto.CompactTextString(m) }
func (*ThreadBlock) ProtoMessage()    {}
func (*ThreadBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_block_93e8a7803ac50cdd, []int{0}
}
func (m *ThreadBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadBlock.Unmarshal(m, b)
}
func (m *ThreadBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadBlock.Marshal(b, m, deterministic)
}
func (dst *ThreadBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadBlock.Merge(dst, src)
}
func (m *ThreadBlock) XXX_Size() int {
	return xxx_messageInfo_ThreadBlock.Size(m)
}
func (m *ThreadBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadBlock.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadBlock proto.InternalMessageInfo

func (m *ThreadBlock) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ThreadBlock) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ThreadBlock) GetParents() []string {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *ThreadBlock) GetTargetKey() []byte {
	if m != nil {
		return m.TargetKey
	}
	return nil
}

func (m *ThreadBlock) GetThreadPubKey() string {
	if m != nil {
		return m.ThreadPubKey
	}
	return ""
}

func (m *ThreadBlock) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type SignedThreadBlock struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	ThreadId             string   `protobuf:"bytes,4,opt,name=threadId,proto3" json:"threadId,omitempty"`
	ThreadName           string   `protobuf:"bytes,5,opt,name=threadName,proto3" json:"threadName,omitempty"`
	IssuerPubKey         []byte   `protobuf:"bytes,6,opt,name=issuerPubKey,proto3" json:"issuerPubKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedThreadBlock) Reset()         { *m = SignedThreadBlock{} }
func (m *SignedThreadBlock) String() string { return proto.CompactTextString(m) }
func (*SignedThreadBlock) ProtoMessage()    {}
func (*SignedThreadBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_block_93e8a7803ac50cdd, []int{1}
}
func (m *SignedThreadBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedThreadBlock.Unmarshal(m, b)
}
func (m *SignedThreadBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedThreadBlock.Marshal(b, m, deterministic)
}
func (dst *SignedThreadBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedThreadBlock.Merge(dst, src)
}
func (m *SignedThreadBlock) XXX_Size() int {
	return xxx_messageInfo_SignedThreadBlock.Size(m)
}
func (m *SignedThreadBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedThreadBlock.DiscardUnknown(m)
}

var xxx_messageInfo_SignedThreadBlock proto.InternalMessageInfo

func (m *SignedThreadBlock) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SignedThreadBlock) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SignedThreadBlock) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignedThreadBlock) GetThreadId() string {
	if m != nil {
		return m.ThreadId
	}
	return ""
}

func (m *SignedThreadBlock) GetThreadName() string {
	if m != nil {
		return m.ThreadName
	}
	return ""
}

func (m *SignedThreadBlock) GetIssuerPubKey() []byte {
	if m != nil {
		return m.IssuerPubKey
	}
	return nil
}

type ThreadInvite struct {
	Block                *ThreadBlock      `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Type                 ThreadInvite_Type `protobuf:"varint,2,opt,name=type,proto3,enum=ThreadInvite_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ThreadInvite) Reset()         { *m = ThreadInvite{} }
func (m *ThreadInvite) String() string { return proto.CompactTextString(m) }
func (*ThreadInvite) ProtoMessage()    {}
func (*ThreadInvite) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_block_93e8a7803ac50cdd, []int{2}
}
func (m *ThreadInvite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadInvite.Unmarshal(m, b)
}
func (m *ThreadInvite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadInvite.Marshal(b, m, deterministic)
}
func (dst *ThreadInvite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadInvite.Merge(dst, src)
}
func (m *ThreadInvite) XXX_Size() int {
	return xxx_messageInfo_ThreadInvite.Size(m)
}
func (m *ThreadInvite) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadInvite.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadInvite proto.InternalMessageInfo

func (m *ThreadInvite) GetBlock() *ThreadBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *ThreadInvite) GetType() ThreadInvite_Type {
	if m != nil {
		return m.Type
	}
	return ThreadInvite_INTERNAL
}

type ThreadData struct {
	Block                *ThreadBlock `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ThreadData) Reset()         { *m = ThreadData{} }
func (m *ThreadData) String() string { return proto.CompactTextString(m) }
func (*ThreadData) ProtoMessage()    {}
func (*ThreadData) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_block_93e8a7803ac50cdd, []int{3}
}
func (m *ThreadData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadData.Unmarshal(m, b)
}
func (m *ThreadData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadData.Marshal(b, m, deterministic)
}
func (dst *ThreadData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadData.Merge(dst, src)
}
func (m *ThreadData) XXX_Size() int {
	return xxx_messageInfo_ThreadData.Size(m)
}
func (m *ThreadData) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadData.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadData proto.InternalMessageInfo

func (m *ThreadData) GetBlock() *ThreadBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

type ThreadAnnotation struct {
	Block                *ThreadBlock `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ThreadAnnotation) Reset()         { *m = ThreadAnnotation{} }
func (m *ThreadAnnotation) String() string { return proto.CompactTextString(m) }
func (*ThreadAnnotation) ProtoMessage()    {}
func (*ThreadAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_block_93e8a7803ac50cdd, []int{4}
}
func (m *ThreadAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadAnnotation.Unmarshal(m, b)
}
func (m *ThreadAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadAnnotation.Marshal(b, m, deterministic)
}
func (dst *ThreadAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadAnnotation.Merge(dst, src)
}
func (m *ThreadAnnotation) XXX_Size() int {
	return xxx_messageInfo_ThreadAnnotation.Size(m)
}
func (m *ThreadAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadAnnotation proto.InternalMessageInfo

func (m *ThreadAnnotation) GetBlock() *ThreadBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*ThreadBlock)(nil), "ThreadBlock")
	proto.RegisterType((*SignedThreadBlock)(nil), "SignedThreadBlock")
	proto.RegisterType((*ThreadInvite)(nil), "ThreadInvite")
	proto.RegisterType((*ThreadData)(nil), "ThreadData")
	proto.RegisterType((*ThreadAnnotation)(nil), "ThreadAnnotation")
	proto.RegisterEnum("ThreadInvite_Type", ThreadInvite_Type_name, ThreadInvite_Type_value)
}

func init() { proto.RegisterFile("thread_block.proto", fileDescriptor_thread_block_93e8a7803ac50cdd) }

var fileDescriptor_thread_block_93e8a7803ac50cdd = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xbf, 0xa4, 0xe9, 0x9f, 0xdc, 0x86, 0xd2, 0x6f, 0x16, 0x12, 0x8a, 0x68, 0x98, 0x85,
	0x64, 0x35, 0x95, 0x0a, 0xee, 0x5b, 0xec, 0xa2, 0x28, 0x45, 0xc6, 0x2c, 0xc4, 0x8d, 0x4c, 0xcc,
	0x18, 0x07, 0xdb, 0x24, 0x24, 0x13, 0xa1, 0x1b, 0xdf, 0xca, 0x07, 0xf0, 0xcd, 0x64, 0x66, 0xd2,
	0xda, 0x6e, 0xa4, 0xbb, 0x9c, 0x33, 0xf7, 0xdc, 0x9c, 0xdf, 0x0c, 0x20, 0xf9, 0x56, 0x72, 0x96,
	0x3c, 0xc7, 0xab, 0xfc, 0xe5, 0x9d, 0x14, 0x65, 0x2e, 0xf3, 0xd1, 0x79, 0x9a, 0xe7, 0xe9, 0x8a,
	0x8f, 0xb5, 0x8a, 0xeb, 0xd7, 0xb1, 0x14, 0x6b, 0x5e, 0x49, 0xb6, 0x2e, 0xcc, 0x00, 0xfe, 0xb6,
	0xa0, 0x1f, 0xe9, 0xdc, 0x4c, 0xc5, 0xd0, 0x00, 0x6c, 0x91, 0xf8, 0x56, 0x60, 0x85, 0x2e, 0xb5,
	0x45, 0x82, 0x4e, 0xa0, 0x23, 0x59, 0x99, 0x72, 0xe9, 0xdb, 0xda, 0x6b, 0x14, 0xf2, 0xa1, 0x5b,
	0xb0, 0x92, 0x67, 0xb2, 0xf2, 0x5b, 0x41, 0x2b, 0x74, 0xe9, 0x56, 0xa2, 0x53, 0x70, 0xcd, 0xcc,
	0x2d, 0xdf, 0xf8, 0x4e, 0x60, 0x85, 0x1e, 0xfd, 0x35, 0x10, 0x06, 0xcf, 0xd4, 0xbc, 0xaf, 0x63,
	0x35, 0xd0, 0xd6, 0x5b, 0x0f, 0x3c, 0x44, 0xc0, 0x49, 0x98, 0xe4, 0x7e, 0x37, 0xb0, 0xc2, 0xfe,
	0x64, 0x44, 0x0c, 0x03, 0xd9, 0x32, 0x90, 0x68, 0xcb, 0x40, 0xf5, 0x1c, 0xfe, 0xb2, 0xe0, 0xff,
	0x83, 0x48, 0x33, 0x9e, 0xfc, 0x45, 0x82, 0xf4, 0x56, 0xa6, 0x39, 0x3c, 0x9d, 0x64, 0xaa, 0x6b,
	0x25, 0xd2, 0x8c, 0xc9, 0xba, 0xe4, 0x7e, 0xcb, 0x74, 0xdd, 0x19, 0x68, 0x04, 0x3d, 0xd3, 0x6b,
	0x91, 0x68, 0x10, 0x97, 0xee, 0x34, 0x3a, 0x03, 0x30, 0xdf, 0x4b, 0xb6, 0xe6, 0x0d, 0xc5, 0x9e,
	0xa3, 0x38, 0x45, 0x55, 0xd5, 0xbc, 0x6c, 0x38, 0x3b, 0x7a, 0xf9, 0x81, 0x87, 0x3f, 0xc1, 0x33,
	0x85, 0x17, 0xd9, 0x87, 0x90, 0x2a, 0xd3, 0xd6, 0x6f, 0xa7, 0x4b, 0xf7, 0x27, 0x1e, 0xd9, 0xc3,
	0xa1, 0xe6, 0x08, 0x5d, 0x80, 0x23, 0x37, 0x05, 0xd7, 0x14, 0x83, 0x09, 0x22, 0xfb, 0x0b, 0x48,
	0xb4, 0x29, 0x38, 0xd5, 0xe7, 0x18, 0x83, 0xa3, 0x14, 0xf2, 0xa0, 0xb7, 0x58, 0x46, 0x73, 0xba,
	0x9c, 0xde, 0x0d, 0xff, 0x29, 0x35, 0x7f, 0x6c, 0x94, 0x85, 0x2f, 0x01, 0x4c, 0xfc, 0x46, 0xdd,
	0xc5, 0x11, 0x7f, 0xc7, 0xd7, 0x30, 0x34, 0xee, 0x34, 0xcb, 0x72, 0xc9, 0xa4, 0xc8, 0xb3, 0x63,
	0x72, 0x33, 0xe7, 0xc9, 0x2e, 0xe2, 0xb8, 0xa3, 0x5f, 0xf0, 0xea, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x1e, 0xd1, 0xbb, 0x58, 0xa9, 0x02, 0x00, 0x00,
}
