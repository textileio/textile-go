// Code generated by protoc-gen-go. DO NOT EDIT.
// source: thread.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
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

type ThreadBlock_Type int32

const (
	ThreadBlock_MERGE      ThreadBlock_Type = 0
	ThreadBlock_IGNORE     ThreadBlock_Type = 1
	ThreadBlock_JOIN       ThreadBlock_Type = 2
	ThreadBlock_ANNOUNCE   ThreadBlock_Type = 3
	ThreadBlock_LEAVE      ThreadBlock_Type = 4
	ThreadBlock_DATA       ThreadBlock_Type = 5
	ThreadBlock_ANNOTATION ThreadBlock_Type = 6
	ThreadBlock_INVITE     ThreadBlock_Type = 50
)

var ThreadBlock_Type_name = map[int32]string{
	0:  "MERGE",
	1:  "IGNORE",
	2:  "JOIN",
	3:  "ANNOUNCE",
	4:  "LEAVE",
	5:  "DATA",
	6:  "ANNOTATION",
	50: "INVITE",
}
var ThreadBlock_Type_value = map[string]int32{
	"MERGE":      0,
	"IGNORE":     1,
	"JOIN":       2,
	"ANNOUNCE":   3,
	"LEAVE":      4,
	"DATA":       5,
	"ANNOTATION": 6,
	"INVITE":     50,
}

func (x ThreadBlock_Type) String() string {
	return proto.EnumName(ThreadBlock_Type_name, int32(x))
}
func (ThreadBlock_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_thread_448bd09d2595bc87, []int{1, 0}
}

type ThreadData_Type int32

const (
	ThreadData_PHOTO ThreadData_Type = 0
	ThreadData_TEXT  ThreadData_Type = 1
)

var ThreadData_Type_name = map[int32]string{
	0: "PHOTO",
	1: "TEXT",
}
var ThreadData_Type_value = map[string]int32{
	"PHOTO": 0,
	"TEXT":  1,
}

func (x ThreadData_Type) String() string {
	return proto.EnumName(ThreadData_Type_name, int32(x))
}
func (ThreadData_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_thread_448bd09d2595bc87, []int{7, 0}
}

type ThreadAnnotation_Type int32

const (
	ThreadAnnotation_COMMENT ThreadAnnotation_Type = 0
	ThreadAnnotation_LIKE    ThreadAnnotation_Type = 1
)

var ThreadAnnotation_Type_name = map[int32]string{
	0: "COMMENT",
	1: "LIKE",
}
var ThreadAnnotation_Type_value = map[string]int32{
	"COMMENT": 0,
	"LIKE":    1,
}

func (x ThreadAnnotation_Type) String() string {
	return proto.EnumName(ThreadAnnotation_Type_name, int32(x))
}
func (ThreadAnnotation_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_thread_448bd09d2595bc87, []int{8, 0}
}

// for wire transport
type ThreadEnvelope struct {
	Thread               string   `protobuf:"bytes,1,opt,name=thread,proto3" json:"thread,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	CipherBlock          []byte   `protobuf:"bytes,3,opt,name=cipherBlock,proto3" json:"cipherBlock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadEnvelope) Reset()         { *m = ThreadEnvelope{} }
func (m *ThreadEnvelope) String() string { return proto.CompactTextString(m) }
func (*ThreadEnvelope) ProtoMessage()    {}
func (*ThreadEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_448bd09d2595bc87, []int{0}
}
func (m *ThreadEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadEnvelope.Unmarshal(m, b)
}
func (m *ThreadEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadEnvelope.Marshal(b, m, deterministic)
}
func (dst *ThreadEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadEnvelope.Merge(dst, src)
}
func (m *ThreadEnvelope) XXX_Size() int {
	return xxx_messageInfo_ThreadEnvelope.Size(m)
}
func (m *ThreadEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadEnvelope proto.InternalMessageInfo

func (m *ThreadEnvelope) GetThread() string {
	if m != nil {
		return m.Thread
	}
	return ""
}

func (m *ThreadEnvelope) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *ThreadEnvelope) GetCipherBlock() []byte {
	if m != nil {
		return m.CipherBlock
	}
	return nil
}

type ThreadBlock struct {
	Header               *ThreadBlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Type                 ThreadBlock_Type   `protobuf:"varint,2,opt,name=type,proto3,enum=ThreadBlock_Type" json:"type,omitempty"`
	Payload              *any.Any           `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ThreadBlock) Reset()         { *m = ThreadBlock{} }
func (m *ThreadBlock) String() string { return proto.CompactTextString(m) }
func (*ThreadBlock) ProtoMessage()    {}
func (*ThreadBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_448bd09d2595bc87, []int{1}
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

func (m *ThreadBlock) GetHeader() *ThreadBlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ThreadBlock) GetType() ThreadBlock_Type {
	if m != nil {
		return m.Type
	}
	return ThreadBlock_MERGE
}

func (m *ThreadBlock) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type ThreadBlockHeader struct {
	Date                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Parents              []string             `protobuf:"bytes,2,rep,name=parents,proto3" json:"parents,omitempty"`
	Author               string               `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ThreadBlockHeader) Reset()         { *m = ThreadBlockHeader{} }
func (m *ThreadBlockHeader) String() string { return proto.CompactTextString(m) }
func (*ThreadBlockHeader) ProtoMessage()    {}
func (*ThreadBlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_448bd09d2595bc87, []int{2}
}
func (m *ThreadBlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadBlockHeader.Unmarshal(m, b)
}
func (m *ThreadBlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadBlockHeader.Marshal(b, m, deterministic)
}
func (dst *ThreadBlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadBlockHeader.Merge(dst, src)
}
func (m *ThreadBlockHeader) XXX_Size() int {
	return xxx_messageInfo_ThreadBlockHeader.Size(m)
}
func (m *ThreadBlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadBlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadBlockHeader proto.InternalMessageInfo

func (m *ThreadBlockHeader) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *ThreadBlockHeader) GetParents() []string {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *ThreadBlockHeader) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type ThreadIgnore struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadIgnore) Reset()         { *m = ThreadIgnore{} }
func (m *ThreadIgnore) String() string { return proto.CompactTextString(m) }
func (*ThreadIgnore) ProtoMessage()    {}
func (*ThreadIgnore) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_448bd09d2595bc87, []int{3}
}
func (m *ThreadIgnore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadIgnore.Unmarshal(m, b)
}
func (m *ThreadIgnore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadIgnore.Marshal(b, m, deterministic)
}
func (dst *ThreadIgnore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadIgnore.Merge(dst, src)
}
func (m *ThreadIgnore) XXX_Size() int {
	return xxx_messageInfo_ThreadIgnore.Size(m)
}
func (m *ThreadIgnore) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadIgnore.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadIgnore proto.InternalMessageInfo

func (m *ThreadIgnore) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ThreadInvite struct {
	Sk                   []byte   `protobuf:"bytes,1,opt,name=sk,proto3" json:"sk,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadInvite) Reset()         { *m = ThreadInvite{} }
func (m *ThreadInvite) String() string { return proto.CompactTextString(m) }
func (*ThreadInvite) ProtoMessage()    {}
func (*ThreadInvite) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_448bd09d2595bc87, []int{4}
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

func (m *ThreadInvite) GetSk() []byte {
	if m != nil {
		return m.Sk
	}
	return nil
}

func (m *ThreadInvite) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ThreadJoin struct {
	Inviter              string   `protobuf:"bytes,1,opt,name=inviter,proto3" json:"inviter,omitempty"`
	Invite               string   `protobuf:"bytes,2,opt,name=invite,proto3" json:"invite,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Inboxes              []string `protobuf:"bytes,4,rep,name=inboxes,proto3" json:"inboxes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadJoin) Reset()         { *m = ThreadJoin{} }
func (m *ThreadJoin) String() string { return proto.CompactTextString(m) }
func (*ThreadJoin) ProtoMessage()    {}
func (*ThreadJoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_448bd09d2595bc87, []int{5}
}
func (m *ThreadJoin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadJoin.Unmarshal(m, b)
}
func (m *ThreadJoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadJoin.Marshal(b, m, deterministic)
}
func (dst *ThreadJoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadJoin.Merge(dst, src)
}
func (m *ThreadJoin) XXX_Size() int {
	return xxx_messageInfo_ThreadJoin.Size(m)
}
func (m *ThreadJoin) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadJoin.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadJoin proto.InternalMessageInfo

func (m *ThreadJoin) GetInviter() string {
	if m != nil {
		return m.Inviter
	}
	return ""
}

func (m *ThreadJoin) GetInvite() string {
	if m != nil {
		return m.Invite
	}
	return ""
}

func (m *ThreadJoin) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ThreadJoin) GetInboxes() []string {
	if m != nil {
		return m.Inboxes
	}
	return nil
}

type ThreadAnnounce struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Inboxes              []string `protobuf:"bytes,2,rep,name=inboxes,proto3" json:"inboxes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadAnnounce) Reset()         { *m = ThreadAnnounce{} }
func (m *ThreadAnnounce) String() string { return proto.CompactTextString(m) }
func (*ThreadAnnounce) ProtoMessage()    {}
func (*ThreadAnnounce) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_448bd09d2595bc87, []int{6}
}
func (m *ThreadAnnounce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadAnnounce.Unmarshal(m, b)
}
func (m *ThreadAnnounce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadAnnounce.Marshal(b, m, deterministic)
}
func (dst *ThreadAnnounce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadAnnounce.Merge(dst, src)
}
func (m *ThreadAnnounce) XXX_Size() int {
	return xxx_messageInfo_ThreadAnnounce.Size(m)
}
func (m *ThreadAnnounce) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadAnnounce.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadAnnounce proto.InternalMessageInfo

func (m *ThreadAnnounce) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ThreadAnnounce) GetInboxes() []string {
	if m != nil {
		return m.Inboxes
	}
	return nil
}

type ThreadData struct {
	Type                 ThreadData_Type `protobuf:"varint,1,opt,name=type,proto3,enum=ThreadData_Type" json:"type,omitempty"`
	Data                 string          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Key                  string          `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Caption              string          `protobuf:"bytes,4,opt,name=caption,proto3" json:"caption,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ThreadData) Reset()         { *m = ThreadData{} }
func (m *ThreadData) String() string { return proto.CompactTextString(m) }
func (*ThreadData) ProtoMessage()    {}
func (*ThreadData) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_448bd09d2595bc87, []int{7}
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

func (m *ThreadData) GetType() ThreadData_Type {
	if m != nil {
		return m.Type
	}
	return ThreadData_PHOTO
}

func (m *ThreadData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ThreadData) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ThreadData) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

type ThreadAnnotation struct {
	Type                 ThreadAnnotation_Type `protobuf:"varint,1,opt,name=type,proto3,enum=ThreadAnnotation_Type" json:"type,omitempty"`
	Data                 string                `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Caption              string                `protobuf:"bytes,3,opt,name=caption,proto3" json:"caption,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ThreadAnnotation) Reset()         { *m = ThreadAnnotation{} }
func (m *ThreadAnnotation) String() string { return proto.CompactTextString(m) }
func (*ThreadAnnotation) ProtoMessage()    {}
func (*ThreadAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_448bd09d2595bc87, []int{8}
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

func (m *ThreadAnnotation) GetType() ThreadAnnotation_Type {
	if m != nil {
		return m.Type
	}
	return ThreadAnnotation_COMMENT
}

func (m *ThreadAnnotation) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ThreadAnnotation) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func init() {
	proto.RegisterType((*ThreadEnvelope)(nil), "ThreadEnvelope")
	proto.RegisterType((*ThreadBlock)(nil), "ThreadBlock")
	proto.RegisterType((*ThreadBlockHeader)(nil), "ThreadBlockHeader")
	proto.RegisterType((*ThreadIgnore)(nil), "ThreadIgnore")
	proto.RegisterType((*ThreadInvite)(nil), "ThreadInvite")
	proto.RegisterType((*ThreadJoin)(nil), "ThreadJoin")
	proto.RegisterType((*ThreadAnnounce)(nil), "ThreadAnnounce")
	proto.RegisterType((*ThreadData)(nil), "ThreadData")
	proto.RegisterType((*ThreadAnnotation)(nil), "ThreadAnnotation")
	proto.RegisterEnum("ThreadBlock_Type", ThreadBlock_Type_name, ThreadBlock_Type_value)
	proto.RegisterEnum("ThreadData_Type", ThreadData_Type_name, ThreadData_Type_value)
	proto.RegisterEnum("ThreadAnnotation_Type", ThreadAnnotation_Type_name, ThreadAnnotation_Type_value)
}

func init() { proto.RegisterFile("thread.proto", fileDescriptor_thread_448bd09d2595bc87) }

var fileDescriptor_thread_448bd09d2595bc87 = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x8d, 0x64, 0xc5, 0x89, 0xc7, 0xc6, 0x6c, 0x96, 0x12, 0xd4, 0x94, 0x52, 0x23, 0x5a, 0x08,
	0x39, 0x28, 0xe0, 0xfe, 0x02, 0x25, 0x51, 0x13, 0xa5, 0x89, 0x54, 0x84, 0x1a, 0x4a, 0x0f, 0x85,
	0xb5, 0xbd, 0x89, 0x84, 0xed, 0x5d, 0x21, 0xad, 0x43, 0xf5, 0x1b, 0x7a, 0xe8, 0xdf, 0xed, 0xb1,
	0xec, 0x87, 0x62, 0x3b, 0xa1, 0xf4, 0x36, 0x4f, 0xf3, 0xe6, 0xed, 0x7b, 0x9a, 0x81, 0x81, 0xc8,
	0x2b, 0x4a, 0x66, 0x7e, 0x59, 0x71, 0xc1, 0x8f, 0x5e, 0x3f, 0x70, 0xfe, 0xb0, 0xa0, 0xa7, 0x0a,
	0x4d, 0x56, 0xf7, 0xa7, 0x84, 0x35, 0xa6, 0xf5, 0xee, 0x79, 0x4b, 0x14, 0x4b, 0x5a, 0x0b, 0xb2,
	0x2c, 0x35, 0xc1, 0xfb, 0x01, 0xc3, 0x4c, 0x69, 0x85, 0xec, 0x91, 0x2e, 0x78, 0x49, 0xf1, 0x21,
	0x74, 0xb5, 0xba, 0x6b, 0x8d, 0xac, 0xe3, 0x5e, 0x6a, 0x10, 0xc6, 0xe0, 0xe4, 0xa4, 0xce, 0x5d,
	0x5b, 0x7d, 0x55, 0x35, 0x1e, 0x41, 0x7f, 0x5a, 0x94, 0x39, 0xad, 0xce, 0x16, 0x7c, 0x3a, 0x77,
	0x3b, 0x23, 0xeb, 0x78, 0x90, 0x6e, 0x7e, 0xf2, 0xfe, 0x58, 0xd0, 0xd7, 0x0f, 0x28, 0x8c, 0x4f,
	0xa0, 0x9b, 0x53, 0x32, 0xa3, 0x95, 0x52, 0xef, 0x8f, 0xb1, 0xbf, 0xd1, 0xbd, 0x52, 0x9d, 0xd4,
	0x30, 0xf0, 0x07, 0x70, 0x44, 0x53, 0x52, 0xf5, 0xe2, 0x70, 0x7c, 0xb0, 0xc9, 0xf4, 0xb3, 0xa6,
	0xa4, 0xa9, 0x6a, 0x63, 0x1f, 0xf6, 0x4a, 0xd2, 0x2c, 0x38, 0x99, 0x29, 0x03, 0xfd, 0xf1, 0x2b,
	0x5f, 0xa7, 0xf6, 0xdb, 0xd4, 0x7e, 0xc0, 0x9a, 0xb4, 0x25, 0x79, 0xf7, 0xe0, 0xc8, 0x69, 0xdc,
	0x83, 0xdd, 0xdb, 0x30, 0xbd, 0x0c, 0xd1, 0x0e, 0x06, 0xe8, 0x46, 0x97, 0x71, 0x92, 0x86, 0xc8,
	0xc2, 0xfb, 0xe0, 0x5c, 0x27, 0x51, 0x8c, 0x6c, 0x3c, 0x80, 0xfd, 0x20, 0x8e, 0x93, 0xaf, 0xf1,
	0x79, 0x88, 0x3a, 0x92, 0x7e, 0x13, 0x06, 0x77, 0x21, 0x72, 0x24, 0xe5, 0x22, 0xc8, 0x02, 0xb4,
	0x8b, 0x87, 0x00, 0x92, 0x92, 0x05, 0x59, 0x94, 0xc4, 0xa8, 0xab, 0x84, 0xe2, 0xbb, 0x28, 0x0b,
	0xd1, 0xd8, 0x5b, 0xc1, 0xc1, 0x8b, 0x6c, 0xd8, 0x07, 0x67, 0x46, 0x04, 0x35, 0xe9, 0x8f, 0x5e,
	0x38, 0xcd, 0xda, 0xfd, 0xa4, 0x8a, 0x87, 0x5d, 0x19, 0xae, 0xa2, 0x4c, 0xd4, 0xae, 0x3d, 0xea,
	0x1c, 0xf7, 0xd2, 0x16, 0xca, 0x3d, 0x91, 0x95, 0xc8, 0x79, 0xa5, 0x52, 0xf7, 0x52, 0x83, 0x3c,
	0x0f, 0x06, 0xfa, 0xd9, 0xe8, 0x81, 0xf1, 0x8a, 0xca, 0xbd, 0xcd, 0x88, 0x20, 0x66, 0x9b, 0xaa,
	0xf6, 0xc6, 0x4f, 0x1c, 0xf6, 0x58, 0x08, 0x8a, 0x87, 0x60, 0xd7, 0x73, 0xc5, 0x18, 0xa4, 0x76,
	0x3d, 0x97, 0x33, 0x8c, 0x2c, 0x69, 0xbb, 0x6b, 0x59, 0x7b, 0x02, 0x40, 0xcf, 0x5c, 0xf3, 0x82,
	0x49, 0x5f, 0x85, 0x9a, 0xad, 0x8c, 0x70, 0x0b, 0xa5, 0x2f, 0x5d, 0x9a, 0x69, 0x83, 0xf0, 0x11,
	0xec, 0xaf, 0x6a, 0x5a, 0x29, 0x5d, 0xed, 0xf8, 0x09, 0x6b, 0xb5, 0x09, 0xff, 0x49, 0x6b, 0xd7,
	0xd1, 0x29, 0x0d, 0xf4, 0x3e, 0xb5, 0xf7, 0x19, 0x30, 0xc6, 0x57, 0x6c, 0xba, 0xad, 0x63, 0xfd,
	0x5b, 0xc7, 0xde, 0xd6, 0xf9, 0x6d, 0xb5, 0xf6, 0x2f, 0x88, 0x20, 0xf8, 0xbd, 0x39, 0x2d, 0x4b,
	0x9d, 0x16, 0xf2, 0xd7, 0xad, 0xcd, 0xcb, 0x6a, 0x7f, 0x9d, 0xbd, 0xfe, 0x75, 0x18, 0x41, 0x67,
	0x4e, 0x1b, 0x93, 0x40, 0x96, 0xf2, 0xd1, 0x29, 0x29, 0x45, 0xc1, 0x99, 0xeb, 0xe8, 0x5f, 0x61,
	0xa0, 0xf7, 0x66, 0x7d, 0x69, 0x5f, 0xae, 0x92, 0x2c, 0x41, 0x3b, 0xf2, 0x74, 0xb2, 0xf0, 0x5b,
	0x86, 0x2c, 0xef, 0x97, 0x05, 0x68, 0x1d, 0x4d, 0x10, 0x39, 0x81, 0x4f, 0xb6, 0x7c, 0x1d, 0xfa,
	0xcf, 0x09, 0xff, 0x73, 0xb7, 0xe1, 0xa5, 0xb3, 0xed, 0xe5, 0xad, 0xf1, 0xd2, 0x87, 0xbd, 0xf3,
	0xe4, 0xf6, 0x36, 0x8c, 0x33, 0xed, 0xe6, 0x26, 0xfa, 0x1c, 0x22, 0xeb, 0xcc, 0xf9, 0x6e, 0x97,
	0x93, 0x49, 0x57, 0xdd, 0xe1, 0xc7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x4a, 0xf2, 0xd8,
	0x60, 0x04, 0x00, 0x00,
}