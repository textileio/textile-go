// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

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

type Message_MessageType int32

const (
	Message_PING             Message_MessageType = 0
	Message_CHAT             Message_MessageType = 1
	Message_THREAD_BLOCK     Message_MessageType = 2
	Message_FOLLOW           Message_MessageType = 3
	Message_UNFOLLOW         Message_MessageType = 4
	Message_OFFLINE_ACK      Message_MessageType = 5
	Message_OFFLINE_RELAY    Message_MessageType = 6
	Message_MODERATOR_ADD    Message_MessageType = 7
	Message_MODERATOR_REMOVE Message_MessageType = 8
	Message_IPFS_BLOCK       Message_MessageType = 9
	Message_ERROR            Message_MessageType = 500
)

var Message_MessageType_name = map[int32]string{
	0:   "PING",
	1:   "CHAT",
	2:   "THREAD_BLOCK",
	3:   "FOLLOW",
	4:   "UNFOLLOW",
	5:   "OFFLINE_ACK",
	6:   "OFFLINE_RELAY",
	7:   "MODERATOR_ADD",
	8:   "MODERATOR_REMOVE",
	9:   "IPFS_BLOCK",
	500: "ERROR",
}
var Message_MessageType_value = map[string]int32{
	"PING":             0,
	"CHAT":             1,
	"THREAD_BLOCK":     2,
	"FOLLOW":           3,
	"UNFOLLOW":         4,
	"OFFLINE_ACK":      5,
	"OFFLINE_RELAY":    6,
	"MODERATOR_ADD":    7,
	"MODERATOR_REMOVE": 8,
	"IPFS_BLOCK":       9,
	"ERROR":            500,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (Message_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_ef4de6396bbb081f, []int{0, 0}
}

type Chat_Flag int32

const (
	Chat_MESSAGE Chat_Flag = 0
	Chat_TYPING  Chat_Flag = 1
	Chat_READ    Chat_Flag = 2
)

var Chat_Flag_name = map[int32]string{
	0: "MESSAGE",
	1: "TYPING",
	2: "READ",
}
var Chat_Flag_value = map[string]int32{
	"MESSAGE": 0,
	"TYPING":  1,
	"READ":    2,
}

func (x Chat_Flag) String() string {
	return proto.EnumName(Chat_Flag_name, int32(x))
}
func (Chat_Flag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_ef4de6396bbb081f, []int{2, 0}
}

type Message struct {
	MessageType          Message_MessageType `protobuf:"varint,1,opt,name=messageType,proto3,enum=Message_MessageType" json:"messageType,omitempty"`
	Payload              *any.Any            `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	RequestId            int32               `protobuf:"varint,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	IsResponse           bool                `protobuf:"varint,4,opt,name=isResponse,proto3" json:"isResponse,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_ef4de6396bbb081f, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessageType() Message_MessageType {
	if m != nil {
		return m.MessageType
	}
	return Message_PING
}

func (m *Message) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetRequestId() int32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *Message) GetIsResponse() bool {
	if m != nil {
		return m.IsResponse
	}
	return false
}

type Envelope struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Pubkey               []byte   `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_ef4de6396bbb081f, []int{1}
}
func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Envelope.Unmarshal(m, b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
}
func (dst *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(dst, src)
}
func (m *Envelope) XXX_Size() int {
	return xxx_messageInfo_Envelope.Size(m)
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Envelope) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *Envelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Chat struct {
	MessageId            string               `protobuf:"bytes,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Subject              string               `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Message              string               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Flag                 Chat_Flag            `protobuf:"varint,5,opt,name=flag,proto3,enum=Chat_Flag" json:"flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Chat) Reset()         { *m = Chat{} }
func (m *Chat) String() string { return proto.CompactTextString(m) }
func (*Chat) ProtoMessage()    {}
func (*Chat) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_ef4de6396bbb081f, []int{2}
}
func (m *Chat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chat.Unmarshal(m, b)
}
func (m *Chat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chat.Marshal(b, m, deterministic)
}
func (dst *Chat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chat.Merge(dst, src)
}
func (m *Chat) XXX_Size() int {
	return xxx_messageInfo_Chat.Size(m)
}
func (m *Chat) XXX_DiscardUnknown() {
	xxx_messageInfo_Chat.DiscardUnknown(m)
}

var xxx_messageInfo_Chat proto.InternalMessageInfo

func (m *Chat) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *Chat) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Chat) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Chat) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Chat) GetFlag() Chat_Flag {
	if m != nil {
		return m.Flag
	}
	return Chat_MESSAGE
}

type SignedData struct {
	SenderPubkey         []byte   `protobuf:"bytes,1,opt,name=senderPubkey,proto3" json:"senderPubkey,omitempty"`
	SerializedData       []byte   `protobuf:"bytes,2,opt,name=serializedData,proto3" json:"serializedData,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedData) Reset()         { *m = SignedData{} }
func (m *SignedData) String() string { return proto.CompactTextString(m) }
func (*SignedData) ProtoMessage()    {}
func (*SignedData) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_ef4de6396bbb081f, []int{3}
}
func (m *SignedData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedData.Unmarshal(m, b)
}
func (m *SignedData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedData.Marshal(b, m, deterministic)
}
func (dst *SignedData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedData.Merge(dst, src)
}
func (m *SignedData) XXX_Size() int {
	return xxx_messageInfo_SignedData.Size(m)
}
func (m *SignedData) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedData.DiscardUnknown(m)
}

var xxx_messageInfo_SignedData proto.InternalMessageInfo

func (m *SignedData) GetSenderPubkey() []byte {
	if m != nil {
		return m.SenderPubkey
	}
	return nil
}

func (m *SignedData) GetSerializedData() []byte {
	if m != nil {
		return m.SerializedData
	}
	return nil
}

func (m *SignedData) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type SignedData_Command struct {
	PeerID               string               `protobuf:"bytes,1,opt,name=peerID,proto3" json:"peerID,omitempty"`
	Type                 Message_MessageType  `protobuf:"varint,2,opt,name=type,proto3,enum=Message_MessageType" json:"type,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SignedData_Command) Reset()         { *m = SignedData_Command{} }
func (m *SignedData_Command) String() string { return proto.CompactTextString(m) }
func (*SignedData_Command) ProtoMessage()    {}
func (*SignedData_Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_ef4de6396bbb081f, []int{3, 0}
}
func (m *SignedData_Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedData_Command.Unmarshal(m, b)
}
func (m *SignedData_Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedData_Command.Marshal(b, m, deterministic)
}
func (dst *SignedData_Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedData_Command.Merge(dst, src)
}
func (m *SignedData_Command) XXX_Size() int {
	return xxx_messageInfo_SignedData_Command.Size(m)
}
func (m *SignedData_Command) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedData_Command.DiscardUnknown(m)
}

var xxx_messageInfo_SignedData_Command proto.InternalMessageInfo

func (m *SignedData_Command) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

func (m *SignedData_Command) GetType() Message_MessageType {
	if m != nil {
		return m.Type
	}
	return Message_PING
}

func (m *SignedData_Command) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type CidList struct {
	Cids                 []string `protobuf:"bytes,1,rep,name=cids,proto3" json:"cids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CidList) Reset()         { *m = CidList{} }
func (m *CidList) String() string { return proto.CompactTextString(m) }
func (*CidList) ProtoMessage()    {}
func (*CidList) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_ef4de6396bbb081f, []int{4}
}
func (m *CidList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CidList.Unmarshal(m, b)
}
func (m *CidList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CidList.Marshal(b, m, deterministic)
}
func (dst *CidList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CidList.Merge(dst, src)
}
func (m *CidList) XXX_Size() int {
	return xxx_messageInfo_CidList.Size(m)
}
func (m *CidList) XXX_DiscardUnknown() {
	xxx_messageInfo_CidList.DiscardUnknown(m)
}

var xxx_messageInfo_CidList proto.InternalMessageInfo

func (m *CidList) GetCids() []string {
	if m != nil {
		return m.Cids
	}
	return nil
}

type IPFSBlock struct {
	RawData              []byte   `protobuf:"bytes,1,opt,name=rawData,proto3" json:"rawData,omitempty"`
	Cid                  string   `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPFSBlock) Reset()         { *m = IPFSBlock{} }
func (m *IPFSBlock) String() string { return proto.CompactTextString(m) }
func (*IPFSBlock) ProtoMessage()    {}
func (*IPFSBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_ef4de6396bbb081f, []int{5}
}
func (m *IPFSBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPFSBlock.Unmarshal(m, b)
}
func (m *IPFSBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPFSBlock.Marshal(b, m, deterministic)
}
func (dst *IPFSBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPFSBlock.Merge(dst, src)
}
func (m *IPFSBlock) XXX_Size() int {
	return xxx_messageInfo_IPFSBlock.Size(m)
}
func (m *IPFSBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_IPFSBlock.DiscardUnknown(m)
}

var xxx_messageInfo_IPFSBlock proto.InternalMessageInfo

func (m *IPFSBlock) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

func (m *IPFSBlock) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

type Error struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_ef4de6396bbb081f, []int{6}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*Envelope)(nil), "Envelope")
	proto.RegisterType((*Chat)(nil), "Chat")
	proto.RegisterType((*SignedData)(nil), "SignedData")
	proto.RegisterType((*SignedData_Command)(nil), "SignedData.Command")
	proto.RegisterType((*CidList)(nil), "CidList")
	proto.RegisterType((*IPFSBlock)(nil), "IPFSBlock")
	proto.RegisterType((*Error)(nil), "Error")
	proto.RegisterEnum("Message_MessageType", Message_MessageType_name, Message_MessageType_value)
	proto.RegisterEnum("Chat_Flag", Chat_Flag_name, Chat_Flag_value)
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_ef4de6396bbb081f) }

var fileDescriptor_message_ef4de6396bbb081f = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x13, 0x27, 0x8e, 0xc7, 0x69, 0x59, 0x56, 0x15, 0x32, 0x15, 0x94, 0xc8, 0x07, 0x14,
	0x2e, 0xae, 0x14, 0x24, 0xe0, 0x86, 0xdc, 0xc4, 0x69, 0xa3, 0x26, 0x75, 0xb5, 0x09, 0xa0, 0x72,
	0xa9, 0x36, 0xf1, 0xd6, 0x98, 0x3a, 0xb6, 0xb1, 0x1d, 0x50, 0xb8, 0xf3, 0x4b, 0xfc, 0x0d, 0x77,
	0x3e, 0x80, 0x33, 0x42, 0xbb, 0x5e, 0x93, 0xb6, 0x48, 0x95, 0x38, 0x79, 0xe6, 0xbd, 0xf1, 0xcc,
	0xf8, 0xbd, 0x31, 0x6c, 0x2f, 0x59, 0x9e, 0xd3, 0x80, 0xd9, 0x69, 0x96, 0x14, 0xc9, 0xde, 0xc3,
	0x20, 0x49, 0x82, 0x88, 0x1d, 0x88, 0x6c, 0xbe, 0xba, 0x3c, 0xa0, 0xf1, 0x5a, 0x52, 0x4f, 0x6e,
	0x53, 0x45, 0xb8, 0x64, 0x79, 0x41, 0x97, 0x69, 0x59, 0x60, 0xfd, 0xac, 0x81, 0x36, 0x29, 0xbb,
	0xe1, 0x17, 0x60, 0xc8, 0xc6, 0xb3, 0x75, 0xca, 0x4c, 0xa5, 0xa3, 0x74, 0x77, 0x7a, 0xbb, 0xb6,
	0xa4, 0xab, 0x27, 0xe7, 0xc8, 0xf5, 0x42, 0x6c, 0x83, 0x96, 0xd2, 0x75, 0x94, 0x50, 0xdf, 0xac,
	0x75, 0x94, 0xae, 0xd1, 0xdb, 0xb5, 0xcb, 0xb1, 0x76, 0x35, 0xd6, 0x76, 0xe2, 0x35, 0xa9, 0x8a,
	0xf0, 0x23, 0xd0, 0x33, 0xf6, 0x69, 0xc5, 0xf2, 0x62, 0xe4, 0x9b, 0xf5, 0x8e, 0xd2, 0x6d, 0x90,
	0x0d, 0x80, 0xf7, 0x01, 0xc2, 0x9c, 0xb0, 0x3c, 0x4d, 0xe2, 0x9c, 0x99, 0x6a, 0x47, 0xe9, 0xb6,
	0xc8, 0x35, 0xc4, 0xfa, 0xae, 0x80, 0x71, 0x6d, 0x15, 0xdc, 0x02, 0xf5, 0x6c, 0x74, 0x7a, 0x84,
	0xb6, 0x78, 0xd4, 0x3f, 0x76, 0x66, 0x48, 0xc1, 0x08, 0xda, 0xb3, 0x63, 0xe2, 0x3a, 0x83, 0x8b,
	0xc3, 0xb1, 0xd7, 0x3f, 0x41, 0x35, 0x0c, 0xd0, 0x1c, 0x7a, 0xe3, 0xb1, 0xf7, 0x0e, 0xd5, 0x71,
	0x1b, 0x5a, 0x6f, 0x4e, 0x65, 0xa6, 0xe2, 0x7b, 0x60, 0x78, 0xc3, 0xe1, 0x78, 0x74, 0xea, 0x5e,
	0x38, 0xfd, 0x13, 0xd4, 0xc0, 0xf7, 0x61, 0xbb, 0x02, 0x88, 0x3b, 0x76, 0xce, 0x51, 0x93, 0x43,
	0x13, 0x6f, 0xe0, 0x12, 0x67, 0xe6, 0x91, 0x0b, 0x67, 0x30, 0x40, 0x1a, 0xde, 0x05, 0xb4, 0x81,
	0x88, 0x3b, 0xf1, 0xde, 0xba, 0xa8, 0x85, 0x77, 0x00, 0x46, 0x67, 0xc3, 0xa9, 0x1c, 0xab, 0x63,
	0x80, 0x86, 0x4b, 0x88, 0x47, 0xd0, 0xaf, 0xba, 0xe5, 0x43, 0xcb, 0x8d, 0x3f, 0xb3, 0x28, 0x49,
	0x19, 0xb6, 0x40, 0x93, 0x0a, 0x0a, 0x99, 0x8d, 0x5e, 0xab, 0x92, 0x97, 0x54, 0x04, 0x7e, 0x00,
	0xcd, 0x74, 0x35, 0xbf, 0x62, 0x6b, 0xa1, 0x6a, 0x9b, 0xc8, 0x8c, 0xcb, 0x97, 0x87, 0x41, 0x4c,
	0x8b, 0x55, 0xc6, 0x84, 0x7c, 0x6d, 0xb2, 0x01, 0xac, 0x1f, 0x0a, 0xa8, 0xfd, 0x0f, 0xb4, 0xe0,
	0x65, 0xb2, 0xd3, 0xc8, 0x17, 0x43, 0x74, 0xb2, 0x01, 0xb0, 0x09, 0x5a, 0xbe, 0x9a, 0x7f, 0x64,
	0x8b, 0x42, 0x74, 0xd7, 0x49, 0x95, 0x72, 0xa6, 0x5a, 0xad, 0x5e, 0x32, 0xd5, 0x42, 0xaf, 0x40,
	0xff, 0x7b, 0x3e, 0xc2, 0x18, 0xa3, 0xb7, 0xf7, 0x8f, 0xd3, 0xb3, 0xaa, 0x82, 0x6c, 0x8a, 0xf1,
	0x3e, 0xa8, 0x97, 0x11, 0x0d, 0xcc, 0x86, 0x38, 0x29, 0xb0, 0xf9, 0x82, 0xf6, 0x30, 0xa2, 0x01,
	0x11, 0xb8, 0xf5, 0x0c, 0x54, 0x9e, 0x61, 0x03, 0xb4, 0x89, 0x3b, 0x9d, 0x3a, 0x47, 0x2e, 0xda,
	0xe2, 0x96, 0xcd, 0xce, 0x85, 0xb5, 0x0a, 0xb7, 0x96, 0xdb, 0x89, 0x6a, 0xd6, 0x6f, 0x05, 0x60,
	0x1a, 0x06, 0x31, 0xf3, 0x07, 0xb4, 0xa0, 0xd8, 0x82, 0x76, 0xce, 0x62, 0x9f, 0x65, 0x67, 0xa5,
	0x54, 0x8a, 0xd0, 0xe3, 0x06, 0x86, 0x9f, 0xc2, 0x4e, 0xce, 0xb2, 0x90, 0x46, 0xe1, 0xd7, 0xf2,
	0x2d, 0x29, 0xe8, 0x2d, 0xf4, 0x6e, 0x61, 0xf7, 0xbe, 0x29, 0xa0, 0xf5, 0x93, 0xe5, 0x92, 0xc6,
	0xbe, 0xb0, 0x86, 0xb1, 0x6c, 0x34, 0x90, 0xc2, 0xca, 0x0c, 0x77, 0x41, 0x2d, 0xf8, 0xaf, 0x53,
	0xbb, 0xe3, 0xd7, 0x11, 0x15, 0x37, 0xb5, 0xac, 0xff, 0x87, 0x96, 0xd6, 0x63, 0xd0, 0xfa, 0xa1,
	0x3f, 0x0e, 0xf3, 0x02, 0x63, 0x50, 0x17, 0xa1, 0x9f, 0x9b, 0x4a, 0xa7, 0xde, 0xd5, 0x89, 0x88,
	0xad, 0x97, 0xa0, 0xf3, 0x0b, 0x3c, 0x8c, 0x92, 0xc5, 0x15, 0xf7, 0x32, 0xa3, 0x5f, 0xc4, 0x27,
	0x97, 0xc2, 0x54, 0x29, 0x46, 0x50, 0x5f, 0x84, 0xbe, 0xf4, 0x9e, 0x87, 0xd6, 0x6b, 0x68, 0xb8,
	0x59, 0x96, 0x64, 0xa2, 0x6b, 0xe2, 0x97, 0x87, 0xb9, 0x4d, 0x44, 0xcc, 0x65, 0x66, 0x9c, 0x94,
	0x1f, 0x22, 0xdf, 0xbb, 0x81, 0x1d, 0xaa, 0xef, 0x6b, 0xe9, 0x7c, 0xde, 0x14, 0xdb, 0x3f, 0xff,
	0x13, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x86, 0xa4, 0xfc, 0xa4, 0x04, 0x00, 0x00,
}
