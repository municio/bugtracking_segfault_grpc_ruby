// Code generated by protoc-gen-go.
// source: munic/type/message.proto
// DO NOT EDIT!

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_type "google.golang.org/genproto/googleapis/type/latlng"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Message struct {
	// Channel is the name of the communication channel.
	Channel string `protobuf:"bytes,1,opt,name=channel" json:"channel,omitempty"`
	// The actual content of the message.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// True if the message must be acknowledged by the recipient.
	// This field is set only when the recipient is an asset.
	WantAck bool `protobuf:"varint,3,opt,name=want_ack,json=wantAck" json:"want_ack,omitempty"`
	// The interpretation of this field depends on wether the recipient of the message
	// is an asset or an agent.
	//
	//  - recipient is an asset :	this field exists iff `want_ack` is true
	// 				and corresponds to the time when the message
	// 				is destroyed by the Binary server if no ack
	// 				is recieved from the asset.
	//  - recipient is the cloud :	TODO
	//
	TimeoutAt *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=timeout_at,json=timeoutAt" json:"timeout_at,omitempty"`
	// Last known position of the `related_asset` when this message was recorded
	// This field is set only when the sender is an asset.
	Position *google_type.LatLng `protobuf:"bytes,5,opt,name=position" json:"position,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Message) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Message) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetWantAck() bool {
	if m != nil {
		return m.WantAck
	}
	return false
}

func (m *Message) GetTimeoutAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.TimeoutAt
	}
	return nil
}

func (m *Message) GetPosition() *google_type.LatLng {
	if m != nil {
		return m.Position
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "munic.type.Message")
}

func init() { proto.RegisterFile("munic/type/message.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0x86, 0xe5, 0xef, 0x03, 0xda, 0x9a, 0x4e, 0x81, 0xc1, 0x74, 0x21, 0x62, 0xca, 0x52, 0x5b,
	0x82, 0x89, 0xb1, 0x9d, 0x8b, 0x54, 0x45, 0x4c, 0x2c, 0xd5, 0xc5, 0x31, 0xae, 0x55, 0xdb, 0x17,
	0x35, 0x17, 0x41, 0xff, 0x0d, 0x3f, 0x86, 0x1f, 0x86, 0xe2, 0x34, 0x30, 0xbe, 0x7a, 0x9f, 0xbb,
	0x3c, 0x17, 0x73, 0x11, 0xba, 0xe8, 0xb4, 0xa2, 0x53, 0x63, 0x54, 0x30, 0x6d, 0x0b, 0xd6, 0xc8,
	0xe6, 0x88, 0x84, 0x19, 0x4f, 0x8d, 0xec, 0x9b, 0xc5, 0xbd, 0x45, 0xb4, 0xde, 0xa8, 0xd4, 0x54,
	0xdd, 0xbb, 0x22, 0x17, 0x4c, 0x4b, 0x10, 0x9a, 0x01, 0x5e, 0x88, 0x33, 0x90, 0xf6, 0x78, 0x20,
	0x1f, 0xed, 0xd0, 0x3c, 0x7c, 0x33, 0x3e, 0x79, 0x19, 0x16, 0x67, 0x82, 0x4f, 0xf4, 0x1e, 0x62,
	0x34, 0x5e, 0xb0, 0x9c, 0x15, 0xb3, 0x72, 0x8c, 0x7d, 0xd3, 0xc0, 0xc9, 0x23, 0xd4, 0xe2, 0x5f,
	0xce, 0x8a, 0x79, 0x39, 0xc6, 0xec, 0x8e, 0x4f, 0x3f, 0x20, 0xd2, 0x0e, 0xf4, 0x41, 0xfc, 0xcf,
	0x59, 0x31, 0x2d, 0x27, 0x7d, 0x5e, 0xe9, 0x43, 0xf6, 0xcc, 0x79, 0xef, 0x81, 0x1d, 0xed, 0x80,
	0xc4, 0x45, 0xce, 0x8a, 0xeb, 0xc7, 0x85, 0x1c, 0x4c, 0xe4, 0xa8, 0x2a, 0x5f, 0x47, 0xd5, 0x72,
	0x76, 0xa6, 0x57, 0x94, 0x29, 0x3e, 0x6d, 0xb0, 0x75, 0xe4, 0x30, 0x8a, 0xcb, 0x34, 0x78, 0x33,
	0x0e, 0xf6, 0x27, 0xc8, 0x0d, 0xd0, 0x26, 0xda, 0xf2, 0x17, 0x5a, 0x7b, 0x7e, 0xab, 0x31, 0xc8,
	0xe1, 0x9f, 0x04, 0xac, 0x8d, 0x4f, 0xe0, 0x7a, 0x7e, 0xbe, 0x6d, 0xdb, 0x7f, 0x6e, 0xcb, 0xde,
	0x56, 0xd6, 0x91, 0x87, 0x4a, 0x06, 0xac, 0x9c, 0x37, 0x4b, 0x17, 0xe9, 0x08, 0x52, 0x63, 0x50,
	0xda, 0x63, 0x57, 0x2f, 0xa3, 0xf9, 0x24, 0xd5, 0xea, 0xbd, 0x09, 0xa0, 0x34, 0x86, 0xc6, 0x79,
	0x53, 0x2b, 0x8b, 0xea, 0xef, 0x15, 0xda, 0x2f, 0xc6, 0xaa, 0xab, 0x64, 0xff, 0xf4, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x6e, 0xfe, 0x46, 0xa1, 0x9e, 0x01, 0x00, 0x00,
}