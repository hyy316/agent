// Code generated by protoc-gen-go.
// source: agent/report/mess/procTime.proto
// DO NOT EDIT!

package mess

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProcTime struct {
	StartTime uint64 `protobuf:"varint,1,opt,name=StartTime" json:"StartTime,omitempty"`
	User      uint64 `protobuf:"varint,2,opt,name=User" json:"User,omitempty"`
	Sys       uint64 `protobuf:"varint,3,opt,name=Sys" json:"Sys,omitempty"`
	Total     uint64 `protobuf:"varint,4,opt,name=Total" json:"Total,omitempty"`
}

func (m *ProcTime) Reset()                    { *m = ProcTime{} }
func (m *ProcTime) String() string            { return proto.CompactTextString(m) }
func (*ProcTime) ProtoMessage()               {}
func (*ProcTime) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func init() {
	proto.RegisterType((*ProcTime)(nil), "mess.ProcTime")
}

func init() { proto.RegisterFile("agent/report/mess/procTime.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x48, 0x4c, 0x4f, 0xcd,
	0x2b, 0xd1, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0xd1, 0xcf, 0x4d, 0x2d, 0x2e, 0xd6, 0x2f, 0x28,
	0xca, 0x4f, 0x0e, 0xc9, 0xcc, 0x4d, 0xd5, 0x03, 0x32, 0x4a, 0xf2, 0x85, 0x58, 0x40, 0x82, 0x4a,
	0x29, 0x5c, 0x1c, 0x01, 0x50, 0x71, 0x21, 0x19, 0x2e, 0xce, 0xe0, 0x92, 0xc4, 0xa2, 0x12, 0x10,
	0x47, 0x82, 0x51, 0x81, 0x51, 0x83, 0x25, 0x08, 0x21, 0x20, 0x24, 0xc4, 0xc5, 0x12, 0x5a, 0x9c,
	0x5a, 0x24, 0xc1, 0x04, 0x96, 0x00, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x83, 0x2b, 0x8b, 0x25, 0x98,
	0xc1, 0x42, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x48, 0x7e, 0x49, 0x62, 0x8e, 0x04, 0x0b, 0x58,
	0x0c, 0xc2, 0x49, 0x62, 0x03, 0x5b, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x80, 0xe2,
	0x2c, 0x96, 0x00, 0x00, 0x00,
}