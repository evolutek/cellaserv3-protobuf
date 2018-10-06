// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cellaserv.proto

package cellaserv

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message_MessageType int32

const (
	Message_Register  Message_MessageType = 0
	Message_Request   Message_MessageType = 1
	Message_Reply     Message_MessageType = 2
	Message_Subscribe Message_MessageType = 3
	Message_Publish   Message_MessageType = 4
)

var Message_MessageType_name = map[int32]string{
	0: "Register",
	1: "Request",
	2: "Reply",
	3: "Subscribe",
	4: "Publish",
}

var Message_MessageType_value = map[string]int32{
	"Register":  0,
	"Request":   1,
	"Reply":     2,
	"Subscribe": 3,
	"Publish":   4,
}

func (x Message_MessageType) Enum() *Message_MessageType {
	p := new(Message_MessageType)
	*p = x
	return p
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}

func (x *Message_MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Message_MessageType_value, data, "Message_MessageType")
	if err != nil {
		return err
	}
	*x = Message_MessageType(value)
	return nil
}

func (Message_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb5d5943edd42d61, []int{0, 0}
}

type Reply_Error_Type int32

const (
	Reply_Error_NoSuchService         Reply_Error_Type = 0
	Reply_Error_InvalidIdentification Reply_Error_Type = 1
	Reply_Error_NoSuchMethod          Reply_Error_Type = 2
	Reply_Error_BadArguments          Reply_Error_Type = 3
	Reply_Error_Timeout               Reply_Error_Type = 4
	// If you use this error type you should explain the error in the
	// 'what' field
	Reply_Error_Custom Reply_Error_Type = 126
)

var Reply_Error_Type_name = map[int32]string{
	0:   "NoSuchService",
	1:   "InvalidIdentification",
	2:   "NoSuchMethod",
	3:   "BadArguments",
	4:   "Timeout",
	126: "Custom",
}

var Reply_Error_Type_value = map[string]int32{
	"NoSuchService":         0,
	"InvalidIdentification": 1,
	"NoSuchMethod":          2,
	"BadArguments":          3,
	"Timeout":               4,
	"Custom":                126,
}

func (x Reply_Error_Type) Enum() *Reply_Error_Type {
	p := new(Reply_Error_Type)
	*p = x
	return p
}

func (x Reply_Error_Type) String() string {
	return proto.EnumName(Reply_Error_Type_name, int32(x))
}

func (x *Reply_Error_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Reply_Error_Type_value, data, "Reply_Error_Type")
	if err != nil {
		return err
	}
	*x = Reply_Error_Type(value)
	return nil
}

func (Reply_Error_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb5d5943edd42d61, []int{3, 0, 0}
}

type Message struct {
	Type *Message_MessageType `protobuf:"varint,1,req,name=type,enum=cellaserv.Message_MessageType" json:"type,omitempty"`
	// This field contains the serialized, actual message:  Register, Request...
	Content              []byte   `protobuf:"bytes,2,req,name=content" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb5d5943edd42d61, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() Message_MessageType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Message_Register
}

func (m *Message) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type Register struct {
	// Name of the service
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// Use this field if multiple instances of the service exist
	Identification       *string  `protobuf:"bytes,2,opt,name=identification" json:"identification,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}
func (*Register) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb5d5943edd42d61, []int{1}
}

func (m *Register) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register.Unmarshal(m, b)
}
func (m *Register) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register.Marshal(b, m, deterministic)
}
func (m *Register) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register.Merge(m, src)
}
func (m *Register) XXX_Size() int {
	return xxx_messageInfo_Register.Size(m)
}
func (m *Register) XXX_DiscardUnknown() {
	xxx_messageInfo_Register.DiscardUnknown(m)
}

var xxx_messageInfo_Register proto.InternalMessageInfo

func (m *Register) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Register) GetIdentification() string {
	if m != nil && m.Identification != nil {
		return *m.Identification
	}
	return ""
}

type Request struct {
	ServiceName           *string  `protobuf:"bytes,1,req,name=service_name,json=serviceName" json:"service_name,omitempty"`
	ServiceIdentification *string  `protobuf:"bytes,2,opt,name=service_identification,json=serviceIdentification" json:"service_identification,omitempty"`
	Method                *string  `protobuf:"bytes,3,req,name=method" json:"method,omitempty"`
	Data                  []byte   `protobuf:"bytes,4,opt,name=data" json:"data,omitempty"`
	Id                    *uint64  `protobuf:"fixed64,99,req,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb5d5943edd42d61, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *Request) GetServiceIdentification() string {
	if m != nil && m.ServiceIdentification != nil {
		return *m.ServiceIdentification
	}
	return ""
}

func (m *Request) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *Request) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Request) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type Reply struct {
	Data                 []byte       `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Error                *Reply_Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Id                   *uint64      `protobuf:"fixed64,99,req,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb5d5943edd42d61, []int{3}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Reply) GetError() *Reply_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Reply) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type Reply_Error struct {
	Type                 *Reply_Error_Type `protobuf:"varint,1,req,name=type,enum=cellaserv.Reply_Error_Type" json:"type,omitempty"`
	What                 *string           `protobuf:"bytes,2,opt,name=what" json:"what,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Reply_Error) Reset()         { *m = Reply_Error{} }
func (m *Reply_Error) String() string { return proto.CompactTextString(m) }
func (*Reply_Error) ProtoMessage()    {}
func (*Reply_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb5d5943edd42d61, []int{3, 0}
}

func (m *Reply_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply_Error.Unmarshal(m, b)
}
func (m *Reply_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply_Error.Marshal(b, m, deterministic)
}
func (m *Reply_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply_Error.Merge(m, src)
}
func (m *Reply_Error) XXX_Size() int {
	return xxx_messageInfo_Reply_Error.Size(m)
}
func (m *Reply_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Reply_Error proto.InternalMessageInfo

func (m *Reply_Error) GetType() Reply_Error_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Reply_Error_NoSuchService
}

func (m *Reply_Error) GetWhat() string {
	if m != nil && m.What != nil {
		return *m.What
	}
	return ""
}

type Subscribe struct {
	// The name of the event the sender of this message is subscribing to
	Event                *string  `protobuf:"bytes,1,req,name=event" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subscribe) Reset()         { *m = Subscribe{} }
func (m *Subscribe) String() string { return proto.CompactTextString(m) }
func (*Subscribe) ProtoMessage()    {}
func (*Subscribe) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb5d5943edd42d61, []int{4}
}

func (m *Subscribe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscribe.Unmarshal(m, b)
}
func (m *Subscribe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscribe.Marshal(b, m, deterministic)
}
func (m *Subscribe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscribe.Merge(m, src)
}
func (m *Subscribe) XXX_Size() int {
	return xxx_messageInfo_Subscribe.Size(m)
}
func (m *Subscribe) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscribe.DiscardUnknown(m)
}

var xxx_messageInfo_Subscribe proto.InternalMessageInfo

func (m *Subscribe) GetEvent() string {
	if m != nil && m.Event != nil {
		return *m.Event
	}
	return ""
}

type Publish struct {
	// The name of the event the sender is publishing
	Event                *string  `protobuf:"bytes,1,req,name=event" json:"event,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Publish) Reset()         { *m = Publish{} }
func (m *Publish) String() string { return proto.CompactTextString(m) }
func (*Publish) ProtoMessage()    {}
func (*Publish) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb5d5943edd42d61, []int{5}
}

func (m *Publish) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Publish.Unmarshal(m, b)
}
func (m *Publish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Publish.Marshal(b, m, deterministic)
}
func (m *Publish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Publish.Merge(m, src)
}
func (m *Publish) XXX_Size() int {
	return xxx_messageInfo_Publish.Size(m)
}
func (m *Publish) XXX_DiscardUnknown() {
	xxx_messageInfo_Publish.DiscardUnknown(m)
}

var xxx_messageInfo_Publish proto.InternalMessageInfo

func (m *Publish) GetEvent() string {
	if m != nil && m.Event != nil {
		return *m.Event
	}
	return ""
}

func (m *Publish) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("cellaserv.Message_MessageType", Message_MessageType_name, Message_MessageType_value)
	proto.RegisterEnum("cellaserv.Reply_Error_Type", Reply_Error_Type_name, Reply_Error_Type_value)
	proto.RegisterType((*Message)(nil), "cellaserv.Message")
	proto.RegisterType((*Register)(nil), "cellaserv.Register")
	proto.RegisterType((*Request)(nil), "cellaserv.Request")
	proto.RegisterType((*Reply)(nil), "cellaserv.Reply")
	proto.RegisterType((*Reply_Error)(nil), "cellaserv.Reply.Error")
	proto.RegisterType((*Subscribe)(nil), "cellaserv.Subscribe")
	proto.RegisterType((*Publish)(nil), "cellaserv.Publish")
}

func init() { proto.RegisterFile("cellaserv.proto", fileDescriptor_bb5d5943edd42d61) }

var fileDescriptor_bb5d5943edd42d61 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xeb, 0xad, 0x93, 0xe0, 0x89, 0x1b, 0x96, 0x11, 0x8d, 0x0c, 0x48, 0xc8, 0xf5, 0x01,
	0xf9, 0x80, 0x82, 0x14, 0xc4, 0x03, 0x00, 0x02, 0xa9, 0x87, 0x16, 0xb4, 0xe9, 0x1d, 0x39, 0xf6,
	0x90, 0xac, 0xe4, 0x3f, 0xa9, 0x77, 0x1d, 0x94, 0x0b, 0x2f, 0xc2, 0x81, 0x2b, 0x6f, 0xc1, 0xab,
	0x55, 0x5e, 0x3b, 0x8e, 0x53, 0xe5, 0xe4, 0xdd, 0x99, 0x9f, 0xbf, 0xdd, 0xfd, 0xbe, 0x81, 0xa7,
	0x31, 0xa5, 0x69, 0xa4, 0xa8, 0xdc, 0xce, 0x36, 0x65, 0xa1, 0x0b, 0x74, 0xba, 0x42, 0xf0, 0xcf,
	0x82, 0xd1, 0x0d, 0x29, 0x15, 0xad, 0x08, 0xe7, 0x60, 0xeb, 0xdd, 0x86, 0x3c, 0xcb, 0x67, 0xe1,
	0x64, 0xfe, 0x7a, 0x76, 0xf8, 0xad, 0x25, 0xf6, 0xdf, 0xbb, 0xdd, 0x86, 0x84, 0x61, 0xd1, 0x83,
	0x51, 0x5c, 0xe4, 0x9a, 0x72, 0xed, 0x31, 0x9f, 0x85, 0xae, 0xd8, 0x6f, 0x83, 0x6f, 0x30, 0xee,
	0xe1, 0xe8, 0xc2, 0x13, 0x41, 0x2b, 0xa9, 0x34, 0x95, 0xfc, 0x0c, 0xc7, 0x30, 0x12, 0x74, 0x5f,
	0x91, 0xd2, 0xdc, 0x42, 0x07, 0x06, 0x82, 0x36, 0xe9, 0x8e, 0x33, 0xbc, 0x00, 0x67, 0x51, 0x2d,
	0x55, 0x5c, 0xca, 0x25, 0xf1, 0xf3, 0x1a, 0xfb, 0x5e, 0x2d, 0x53, 0xa9, 0xd6, 0xdc, 0x0e, 0xbe,
	0x1e, 0x14, 0x10, 0xc1, 0xce, 0xa3, 0xac, 0xb9, 0xaa, 0x23, 0xcc, 0x1a, 0xdf, 0xc0, 0x44, 0x26,
	0x94, 0x6b, 0xf9, 0x53, 0xc6, 0x91, 0x96, 0x45, 0xee, 0x31, 0xdf, 0x0a, 0x1d, 0xf1, 0xa8, 0x1a,
	0xfc, 0xb5, 0xba, 0xc3, 0xf1, 0x0a, 0xdc, 0xfa, 0x81, 0x32, 0xa6, 0x1f, 0x3d, 0xbd, 0x71, 0x5b,
	0xbb, 0xad, 0x65, 0x3f, 0xc0, 0x74, 0x8f, 0x9c, 0x94, 0xbf, 0x6c, 0xbb, 0xd7, 0x47, 0x4d, 0x9c,
	0xc2, 0x30, 0x23, 0xbd, 0x2e, 0x12, 0xef, 0xdc, 0x68, 0xb6, 0xbb, 0xfa, 0xe6, 0x49, 0xa4, 0x23,
	0xcf, 0xf6, 0xad, 0xd0, 0x15, 0x66, 0x8d, 0x13, 0x60, 0x32, 0xf1, 0x62, 0x9f, 0x85, 0x43, 0xc1,
	0x64, 0x12, 0xfc, 0x61, 0xad, 0x23, 0x1d, 0x6d, 0xf5, 0xe8, 0xb7, 0x30, 0xa0, 0xb2, 0x2c, 0x4a,
	0x73, 0xfe, 0x78, 0x3e, 0xed, 0xe5, 0x64, 0x7e, 0x9a, 0x7d, 0xa9, 0xbb, 0xa2, 0x81, 0x1e, 0x6b,
	0xbf, 0xfc, 0x6f, 0xc1, 0xc0, 0x00, 0xf8, 0xee, 0x28, 0xee, 0x57, 0xa7, 0x65, 0x66, 0xbd, 0xac,
	0x11, 0xec, 0x5f, 0xeb, 0x48, 0xb7, 0xef, 0x36, 0xeb, 0xe0, 0x1e, 0x6c, 0x13, 0xef, 0x33, 0xb8,
	0xb8, 0x2d, 0x16, 0x55, 0xbc, 0x5e, 0x34, 0x6e, 0xf0, 0x33, 0x7c, 0x01, 0x97, 0xd7, 0xf9, 0x36,
	0x4a, 0x65, 0x72, 0x6c, 0x0d, 0xb7, 0x90, 0x83, 0xdb, 0xd0, 0x37, 0xc6, 0x14, 0xce, 0xea, 0xca,
	0xa7, 0x28, 0xf9, 0x58, 0xae, 0xaa, 0x8c, 0x72, 0xad, 0x9a, 0xec, 0xef, 0x64, 0x46, 0x45, 0xa5,
	0xb9, 0x8d, 0x00, 0xc3, 0xcf, 0x95, 0xd2, 0x45, 0xc6, 0x7f, 0x07, 0x57, 0xbd, 0x19, 0xc1, 0xe7,
	0x30, 0xa0, 0x6d, 0x3d, 0x7d, 0x4d, 0x72, 0xcd, 0x26, 0x78, 0xdf, 0xcd, 0xcd, 0x69, 0xa0, 0xf3,
	0x95, 0x1d, 0x7c, 0x7d, 0x08, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x31, 0x78, 0x56, 0x27, 0x03, 0x00,
	0x00,
}
