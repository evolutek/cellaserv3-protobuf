// Code generated by protoc-gen-go.
// source: cellaserv.proto
// DO NOT EDIT!

package cellaserv

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

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

type Reply_Error_Type int32

const (
	Reply_Error_NoSuchService         Reply_Error_Type = 0
	Reply_Error_InvalidIdentification Reply_Error_Type = 1
	Reply_Error_NoSuchMethod          Reply_Error_Type = 2
	Reply_Error_BadArguments          Reply_Error_Type = 3
	Reply_Error_Timeout               Reply_Error_Type = 4
	Reply_Error_Custom                Reply_Error_Type = 126
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

type Message struct {
	Type             *Message_MessageType `protobuf:"varint,1,req,name=type,enum=cellaserv.Message_MessageType" json:"type,omitempty"`
	Content          []byte               `protobuf:"bytes,2,req,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

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
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Identification   *string `protobuf:"bytes,2,opt,name=identification" json:"identification,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}

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
	ServiceName           *string `protobuf:"bytes,1,req,name=service_name" json:"service_name,omitempty"`
	ServiceIdentification *string `protobuf:"bytes,2,opt,name=service_identification" json:"service_identification,omitempty"`
	Method                *string `protobuf:"bytes,3,req,name=method" json:"method,omitempty"`
	Data                  []byte  `protobuf:"bytes,4,opt,name=data" json:"data,omitempty"`
	Id                    *uint64 `protobuf:"fixed64,99,req,name=id" json:"id,omitempty"`
	XXX_unrecognized      []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

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
	Data             []byte       `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Error            *Reply_Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Id               *uint64      `protobuf:"fixed64,99,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}

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
	Type             *Reply_Error_Type `protobuf:"varint,1,req,name=type,enum=cellaserv.Reply_Error_Type" json:"type,omitempty"`
	What             *string           `protobuf:"bytes,2,opt,name=what" json:"what,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Reply_Error) Reset()         { *m = Reply_Error{} }
func (m *Reply_Error) String() string { return proto.CompactTextString(m) }
func (*Reply_Error) ProtoMessage()    {}

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
	Event            *string `protobuf:"bytes,1,req,name=event" json:"event,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Subscribe) Reset()         { *m = Subscribe{} }
func (m *Subscribe) String() string { return proto.CompactTextString(m) }
func (*Subscribe) ProtoMessage()    {}

func (m *Subscribe) GetEvent() string {
	if m != nil && m.Event != nil {
		return *m.Event
	}
	return ""
}

type Publish struct {
	Event            *string `protobuf:"bytes,1,req,name=event" json:"event,omitempty"`
	Data             []byte  `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Publish) Reset()         { *m = Publish{} }
func (m *Publish) String() string { return proto.CompactTextString(m) }
func (*Publish) ProtoMessage()    {}

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
}
