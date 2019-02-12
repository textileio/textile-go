// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"
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

type Contact struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string               `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Username             string               `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Avatar               string               `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Inboxes              []*Cafe              `protobuf:"bytes,5,rep,name=inboxes,proto3" json:"inboxes,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_5017ecd419bd7aee, []int{0}
}
func (m *Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contact.Unmarshal(m, b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
}
func (dst *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(dst, src)
}
func (m *Contact) XXX_Size() int {
	return xxx_messageInfo_Contact.Size(m)
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func (m *Contact) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Contact) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Contact) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Contact) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Contact) GetInboxes() []*Cafe {
	if m != nil {
		return m.Inboxes
	}
	return nil
}

func (m *Contact) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Contact) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

type File struct {
	Mill                 string               `protobuf:"bytes,1,opt,name=mill,proto3" json:"mill,omitempty"`
	Checksum             string               `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Source               string               `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Opts                 string               `protobuf:"bytes,4,opt,name=opts,proto3" json:"opts,omitempty"`
	Hash                 string               `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	Key                  string               `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	Media                string               `protobuf:"bytes,7,opt,name=media,proto3" json:"media,omitempty"`
	Name                 string               `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Size                 int64                `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	Added                *timestamp.Timestamp `protobuf:"bytes,10,opt,name=added,proto3" json:"added,omitempty"`
	Meta                 *_struct.Struct      `protobuf:"bytes,11,opt,name=meta,proto3" json:"meta,omitempty"`
	Targets              []string             `protobuf:"bytes,12,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_5017ecd419bd7aee, []int{1}
}
func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (dst *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(dst, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetMill() string {
	if m != nil {
		return m.Mill
	}
	return ""
}

func (m *File) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *File) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *File) GetOpts() string {
	if m != nil {
		return m.Opts
	}
	return ""
}

func (m *File) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *File) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *File) GetMedia() string {
	if m != nil {
		return m.Media
	}
	return ""
}

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *File) GetAdded() *timestamp.Timestamp {
	if m != nil {
		return m.Added
	}
	return nil
}

func (m *File) GetMeta() *_struct.Struct {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *File) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type Cafe struct {
	Peer                 string   `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Api                  string   `protobuf:"bytes,3,opt,name=api,proto3" json:"api,omitempty"`
	Protocol             string   `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Node                 string   `protobuf:"bytes,5,opt,name=node,proto3" json:"node,omitempty"`
	Url                  string   `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Swarm                []string `protobuf:"bytes,7,rep,name=swarm,proto3" json:"swarm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cafe) Reset()         { *m = Cafe{} }
func (m *Cafe) String() string { return proto.CompactTextString(m) }
func (*Cafe) ProtoMessage()    {}
func (*Cafe) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_5017ecd419bd7aee, []int{2}
}
func (m *Cafe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cafe.Unmarshal(m, b)
}
func (m *Cafe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cafe.Marshal(b, m, deterministic)
}
func (dst *Cafe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cafe.Merge(dst, src)
}
func (m *Cafe) XXX_Size() int {
	return xxx_messageInfo_Cafe.Size(m)
}
func (m *Cafe) XXX_DiscardUnknown() {
	xxx_messageInfo_Cafe.DiscardUnknown(m)
}

var xxx_messageInfo_Cafe proto.InternalMessageInfo

func (m *Cafe) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

func (m *Cafe) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Cafe) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *Cafe) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Cafe) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *Cafe) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Cafe) GetSwarm() []string {
	if m != nil {
		return m.Swarm
	}
	return nil
}

func init() {
	proto.RegisterType((*Contact)(nil), "Contact")
	proto.RegisterType((*File)(nil), "File")
	proto.RegisterType((*Cafe)(nil), "Cafe")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_model_5017ecd419bd7aee) }

var fileDescriptor_model_5017ecd419bd7aee = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xdd, 0x8a, 0xdb, 0x3c,
	0x10, 0xc5, 0x7f, 0xf1, 0x66, 0xf2, 0xf1, 0xb1, 0x88, 0xd2, 0x8a, 0x50, 0x58, 0x93, 0xab, 0x40,
	0xc1, 0x5b, 0xb6, 0x7d, 0x82, 0x2e, 0xf4, 0x01, 0xdc, 0x5e, 0xf5, 0x4e, 0xb1, 0x66, 0x13, 0xb3,
	0x56, 0x64, 0x24, 0xb9, 0x7f, 0x0f, 0xd3, 0x57, 0xe9, 0x4b, 0xf5, 0x01, 0xca, 0x8c, 0xa5, 0x5c,
	0xb4, 0xd0, 0xbd, 0x3b, 0x67, 0x34, 0x67, 0x3c, 0xe7, 0x78, 0x60, 0x63, 0xac, 0xc6, 0xb1, 0x9d,
	0x9c, 0x0d, 0x76, 0x7b, 0x73, 0xb4, 0xf6, 0x38, 0xe2, 0x2d, 0xb3, 0xc3, 0xfc, 0x70, 0x1b, 0x06,
	0x83, 0x3e, 0x28, 0x33, 0xc5, 0x86, 0x97, 0x7f, 0x36, 0xf8, 0xe0, 0xe6, 0x3e, 0x2c, 0xaf, 0xbb,
	0x5f, 0x19, 0xd4, 0xf7, 0xf6, 0x1c, 0x54, 0x1f, 0xc4, 0xff, 0x90, 0x0f, 0x5a, 0x66, 0x4d, 0xb6,
	0x5f, 0x77, 0xf9, 0xa0, 0x85, 0x84, 0x5a, 0x69, 0xed, 0xd0, 0x7b, 0x99, 0x73, 0x31, 0x51, 0xb1,
	0x85, 0xab, 0xd9, 0xa3, 0x3b, 0x2b, 0x83, 0xb2, 0xe0, 0xa7, 0x0b, 0x17, 0xcf, 0x61, 0xa5, 0x3e,
	0xab, 0xa0, 0x9c, 0x2c, 0xf9, 0x25, 0x32, 0x71, 0x03, 0xf5, 0x70, 0x3e, 0xd8, 0xaf, 0xe8, 0x65,
	0xd5, 0x14, 0xfb, 0xcd, 0x5d, 0xd5, 0xde, 0xab, 0x07, 0xec, 0x52, 0x55, 0xbc, 0x85, 0xba, 0x77,
	0xa8, 0x02, 0x6a, 0xb9, 0x6a, 0xb2, 0xfd, 0xe6, 0x6e, 0xdb, 0x2e, 0xab, 0xb7, 0x69, 0xf5, 0xf6,
	0x63, 0xf2, 0xd6, 0xa5, 0x56, 0x52, 0xcd, 0x93, 0x66, 0x55, 0xfd, 0xb4, 0x2a, 0xb6, 0xee, 0x7e,
	0xe6, 0x50, 0xbe, 0x1f, 0x46, 0x14, 0x02, 0x4a, 0x33, 0x8c, 0x63, 0x74, 0xcd, 0x98, 0xdc, 0xf5,
	0x27, 0xec, 0x1f, 0xfd, 0x6c, 0xa2, 0xf1, 0x0b, 0x27, 0x77, 0xde, 0xce, 0xae, 0x4f, 0xbe, 0x23,
	0xa3, 0x39, 0x76, 0x0a, 0x3e, 0x7a, 0x66, 0x4c, 0xb5, 0x93, 0xf2, 0x27, 0x59, 0x2d, 0x35, 0xc2,
	0xe2, 0x1a, 0x8a, 0x47, 0xfc, 0xc6, 0x06, 0xd7, 0x1d, 0x41, 0xf1, 0x0c, 0x2a, 0x83, 0x7a, 0x50,
	0xbc, 0xfe, 0xba, 0x5b, 0x08, 0x69, 0x39, 0xdd, 0xab, 0x45, 0xcb, 0xc9, 0x0a, 0x28, 0xfd, 0xf0,
	0x1d, 0xe5, 0xba, 0xc9, 0xf6, 0x45, 0xc7, 0x58, 0xbc, 0x86, 0x4a, 0x69, 0x8d, 0x5a, 0xc2, 0x93,
	0xe6, 0x97, 0x46, 0xf1, 0x0a, 0x4a, 0x83, 0x41, 0xc9, 0x0d, 0x0b, 0x5e, 0xfc, 0x25, 0xf8, 0xc0,
	0xe7, 0xd1, 0x71, 0x13, 0x9d, 0x40, 0x50, 0xee, 0x88, 0xc1, 0xcb, 0xff, 0x9a, 0x82, 0x4e, 0x20,
	0xd2, 0xdd, 0x8f, 0x0c, 0x4a, 0xfa, 0x7f, 0xb4, 0xd5, 0x84, 0xe8, 0x52, 0x82, 0x84, 0xff, 0x71,
	0x39, 0xd7, 0x50, 0xa8, 0x69, 0x88, 0xe1, 0x11, 0xa4, 0xb4, 0xf9, 0xdb, 0xbd, 0x1d, 0x63, 0x7a,
	0x17, 0xce, 0x29, 0x58, 0x8d, 0x29, 0x41, 0xc2, 0x34, 0x61, 0x76, 0x63, 0x4a, 0x70, 0x76, 0x23,
	0x25, 0xe8, 0xbf, 0x28, 0x67, 0x64, 0xcd, 0x2b, 0x2e, 0xe4, 0x5d, 0xf9, 0x29, 0x9f, 0x0e, 0x87,
	0x15, 0xcf, 0x7a, 0xf3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x57, 0x1f, 0xc9, 0x63, 0x34, 0x03, 0x00,
	0x00,
}
