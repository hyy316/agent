// Code generated by protoc-gen-go.
// source: agent/report/mess/procState.proto
// DO NOT EDIT!

package mess

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProcState struct {
	Name      string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=Username" json:"Username,omitempty"`
	State     int64  `protobuf:"varint,3,opt,name=State" json:"State,omitempty"`
	Ppid      int64  `protobuf:"varint,4,opt,name=Ppid" json:"Ppid,omitempty"`
	Pgid      int64  `protobuf:"varint,5,opt,name=Pgid" json:"Pgid,omitempty"`
	Tty       int64  `protobuf:"varint,6,opt,name=Tty" json:"Tty,omitempty"`
	Priority  int64  `protobuf:"varint,7,opt,name=Priority" json:"Priority,omitempty"`
	Nice      int64  `protobuf:"varint,8,opt,name=Nice" json:"Nice,omitempty"`
	Processor int64  `protobuf:"varint,9,opt,name=Processor" json:"Processor,omitempty"`
}

func (m *ProcState) Reset()                    { *m = ProcState{} }
func (m *ProcState) String() string            { return proto.CompactTextString(m) }
func (*ProcState) ProtoMessage()               {}
func (*ProcState) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func init() {
	proto.RegisterType((*ProcState)(nil), "mess.ProcState")
}

func init() { proto.RegisterFile("agent/report/mess/procState.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0xcf, 0x3b, 0x0e, 0xc2, 0x30,
	0x0c, 0x06, 0x60, 0x95, 0x3e, 0x68, 0x33, 0x21, 0x8b, 0x21, 0x42, 0x0c, 0xc0, 0xc4, 0x44, 0x07,
	0xee, 0x81, 0xaa, 0x02, 0x07, 0x28, 0xc5, 0xaa, 0x32, 0xb4, 0xa9, 0xdc, 0x2c, 0x1c, 0x95, 0xdb,
	0xe0, 0xb8, 0x10, 0xb6, 0xdf, 0x9f, 0x15, 0xc7, 0x56, 0xfb, 0xa6, 0xc3, 0xc1, 0x95, 0x84, 0xa3,
	0x25, 0x57, 0xf6, 0x38, 0x4d, 0xe5, 0x48, 0xb6, 0xbd, 0xba, 0xc6, 0xe1, 0x89, 0x93, 0xb3, 0x90,
	0x78, 0x3d, 0xbc, 0x23, 0x55, 0x54, 0xbf, 0x0e, 0x80, 0x4a, 0x2e, 0x4d, 0x8f, 0x3a, 0xda, 0x45,
	0xc7, 0xa2, 0x96, 0x0c, 0x1b, 0x95, 0xdf, 0x27, 0xa4, 0xc1, 0xfb, 0x42, 0x3c, 0xd4, 0xb0, 0x56,
	0xa9, 0x3c, 0xd4, 0x31, 0x37, 0xe2, 0x3a, 0x0d, 0x53, 0xaa, 0xd1, 0x3c, 0x75, 0x22, 0x28, 0x59,
	0xac, 0x63, 0x4b, 0xbf, 0xc6, 0x19, 0x56, 0x2a, 0xbe, 0xb9, 0x97, 0xce, 0x84, 0x7c, 0xf4, 0x7f,
	0x55, 0x64, 0x2c, 0x19, 0xe6, 0xa5, 0x70, 0xa8, 0x65, 0x37, 0xd3, 0xa2, 0xce, 0xe7, 0x09, 0x3e,
	0xc3, 0x76, 0x5e, 0x9e, 0x0f, 0xb1, 0xa4, 0x0b, 0x69, 0xfc, 0xe1, 0x91, 0xc9, 0xa1, 0xe7, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x90, 0xc7, 0xc9, 0x0d, 0x01, 0x00, 0x00,
}