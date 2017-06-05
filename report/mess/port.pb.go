// Code generated by protoc-gen-go.
// source: agent/report/mess/port.proto
// DO NOT EDIT!

package mess

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Port struct {
	PortNumber int64  `protobuf:"varint,1,opt,name=PortNumber" json:"PortNumber,omitempty"`
	Pid        int64  `protobuf:"varint,2,opt,name=Pid" json:"Pid,omitempty"`
	Protocol   string `protobuf:"bytes,3,opt,name=Protocol" json:"Protocol,omitempty"`
}

func (m *Port) Reset()                    { *m = Port{} }
func (m *Port) String() string            { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()               {}
func (*Port) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func init() {
	proto.RegisterType((*Port)(nil), "mess.Port")
}

func init() { proto.RegisterFile("agent/report/mess/port.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x4c, 0x4f, 0xcd,
	0x2b, 0xd1, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0xd1, 0xcf, 0x4d, 0x2d, 0x2e, 0xd6, 0x07, 0xb1,
	0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x58, 0x40, 0x02, 0x4a, 0x21, 0x5c, 0x2c, 0x01, 0x40,
	0x31, 0x21, 0x39, 0x2e, 0x2e, 0x10, 0xed, 0x57, 0x9a, 0x9b, 0x94, 0x5a, 0x24, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x1c, 0x84, 0x24, 0x22, 0x24, 0xc0, 0xc5, 0x1c, 0x90, 0x99, 0x22, 0xc1, 0x04, 0x96,
	0x00, 0x31, 0x85, 0xa4, 0xb8, 0x38, 0x02, 0x40, 0x06, 0x25, 0xe7, 0xe7, 0x48, 0x30, 0x03, 0x85,
	0x39, 0x83, 0xe0, 0xfc, 0x24, 0x36, 0xb0, 0x15, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b,
	0xd7, 0x8a, 0x33, 0x82, 0x00, 0x00, 0x00,
}
