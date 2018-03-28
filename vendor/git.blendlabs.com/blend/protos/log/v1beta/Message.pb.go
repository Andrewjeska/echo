// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log/v1beta/Message.proto

/*
Package v1beta is a generated protocol buffer package.

It is generated from these files:
	log/v1beta/Message.proto
	log/v1beta/SmokeSignalResult.proto

It has these top-level messages:
	Message
	SmokeSignalResult
*/
package v1beta

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import logv1 "git.blendlabs.com/blend/protos/log/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessageType int32

const (
	MessageType_SMOKE_SIGNAL MessageType = 0
)

var MessageType_name = map[int32]string{
	0: "SMOKE_SIGNAL",
}
var MessageType_value = map[string]int32{
	"SMOKE_SIGNAL": 0,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// [START messages]
type Message struct {
	// type is the type of the object
	// it is used principally in routing to processors.
	Type MessageType `protobuf:"varint,1,opt,name=type,enum=logv1beta.MessageType" json:"type,omitempty"`
	// meta contains extra metadata for the message like timestamp.
	Meta              *logv1.Meta        `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	SmokeSignalResult *SmokeSignalResult `protobuf:"bytes,200,opt,name=smokeSignalResult" json:"smokeSignalResult,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_SMOKE_SIGNAL
}

func (m *Message) GetMeta() *logv1.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Message) GetSmokeSignalResult() *SmokeSignalResult {
	if m != nil {
		return m.SmokeSignalResult
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "logv1beta.Message")
	proto.RegisterEnum("logv1beta.MessageType", MessageType_name, MessageType_value)
}

func init() { proto.RegisterFile("log/v1beta/Message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xc9, 0x4f, 0xd7,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0xd4, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0xc9, 0x4f, 0x87, 0x48, 0x48, 0x09, 0x42, 0x14, 0xe9,
	0xfb, 0xa6, 0x96, 0x24, 0x42, 0x64, 0xa5, 0x94, 0x90, 0xf4, 0x05, 0xe7, 0xe6, 0x67, 0xa7, 0x06,
	0x67, 0xa6, 0xe7, 0x25, 0xe6, 0x04, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x40, 0xd4, 0x28, 0x2d, 0x66,
	0xe4, 0x62, 0x87, 0x9a, 0x29, 0xa4, 0xc5, 0xc5, 0x52, 0x52, 0x59, 0x90, 0x2a, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x67, 0x24, 0xa6, 0x07, 0x37, 0x5c, 0x0f, 0xaa, 0x22, 0xa4, 0xb2, 0x20, 0x35, 0x08,
	0xac, 0x46, 0x48, 0x9e, 0x8b, 0x25, 0x37, 0xb5, 0x24, 0x51, 0x82, 0x49, 0x81, 0x51, 0x83, 0xdb,
	0x88, 0x1b, 0xa2, 0x56, 0x0f, 0x64, 0x79, 0x10, 0x58, 0x42, 0xc8, 0x9b, 0x4b, 0xb0, 0x18, 0xdd,
	0x4e, 0x89, 0x13, 0x8c, 0x60, 0xe5, 0x32, 0x48, 0x46, 0x63, 0x38, 0x2c, 0x08, 0x53, 0x9f, 0x96,
	0x3c, 0x17, 0x37, 0x92, 0x13, 0x84, 0x04, 0xb8, 0x78, 0x82, 0x7d, 0xfd, 0xbd, 0x5d, 0xe3, 0x83,
	0x3d, 0xdd, 0xfd, 0x1c, 0x7d, 0x04, 0x18, 0x9c, 0xb4, 0xa3, 0x34, 0xd3, 0x33, 0x4b, 0xf4, 0x92,
	0x72, 0x52, 0xf3, 0x52, 0x72, 0x12, 0x93, 0x8a, 0xf5, 0x92, 0xf3, 0x73, 0xf5, 0xc1, 0x3c, 0x7d,
	0xb0, 0x3f, 0x8b, 0xf5, 0x11, 0x61, 0x91, 0xc4, 0x06, 0x16, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x90, 0xe7, 0x3c, 0x5a, 0x58, 0x01, 0x00, 0x00,
}
