// Code generated by protoc-gen-go.
// source: agent/report/mess/procMem.proto
// DO NOT EDIT!

package mess

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProcMem struct {
	Size        uint64 `protobuf:"varint,1,opt,name=Size" json:"Size,omitempty"`
	Resident    uint64 `protobuf:"varint,2,opt,name=Resident" json:"Resident,omitempty"`
	Share       uint64 `protobuf:"varint,3,opt,name=Share" json:"Share,omitempty"`
	MinorFaults uint64 `protobuf:"varint,4,opt,name=MinorFaults" json:"MinorFaults,omitempty"`
	MajorFaults uint64 `protobuf:"varint,5,opt,name=MajorFaults" json:"MajorFaults,omitempty"`
	PageFaults  uint64 `protobuf:"varint,6,opt,name=PageFaults" json:"PageFaults,omitempty"`
}

func (m *ProcMem) Reset()                    { *m = ProcMem{} }
func (m *ProcMem) String() string            { return proto.CompactTextString(m) }
func (*ProcMem) ProtoMessage()               {}
func (*ProcMem) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func init() {
	proto.RegisterType((*ProcMem)(nil), "mess.ProcMem")
}

func init() { proto.RegisterFile("agent/report/mess/procMem.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x4c, 0x4f, 0xcd,
	0x2b, 0xd1, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0xd1, 0xcf, 0x4d, 0x2d, 0x2e, 0xd6, 0x2f, 0x28,
	0xca, 0x4f, 0xf6, 0x4d, 0xcd, 0xd5, 0x03, 0xd2, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0x31, 0xa5, 0xcd,
	0x8c, 0x5c, 0xec, 0x01, 0x10, 0x71, 0x21, 0x21, 0x2e, 0x96, 0xe0, 0xcc, 0xaa, 0x54, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0x96, 0x20, 0x30, 0x5b, 0x48, 0x8a, 0x8b, 0x23, 0x28, 0xb5, 0x38, 0x33, 0x05,
	0x68, 0x96, 0x04, 0x13, 0x58, 0x1c, 0xce, 0x17, 0x12, 0xe1, 0x62, 0x0d, 0xce, 0x48, 0x2c, 0x4a,
	0x95, 0x60, 0x06, 0x4b, 0x40, 0x38, 0x42, 0x0a, 0x5c, 0xdc, 0xbe, 0x99, 0x79, 0xf9, 0x45, 0x6e,
	0x89, 0xa5, 0x39, 0x25, 0xc5, 0x12, 0x2c, 0x60, 0x39, 0x64, 0x21, 0xb0, 0x8a, 0xc4, 0x2c, 0xb8,
	0x0a, 0x56, 0xa8, 0x0a, 0x84, 0x90, 0x90, 0x1c, 0x17, 0x57, 0x00, 0xd0, 0xfd, 0x50, 0x05, 0x6c,
	0x60, 0x05, 0x48, 0x22, 0x49, 0x6c, 0x60, 0x2f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9f,
	0x9e, 0x2a, 0x50, 0xe5, 0x00, 0x00, 0x00,
}
