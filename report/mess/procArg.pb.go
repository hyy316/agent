// Code generated by protoc-gen-go.
// source: agent/report/mess/procArg.proto
// DO NOT EDIT!

package mess

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProcArg struct {
	StartDir string `protobuf:"bytes,1,opt,name=StartDir" json:"StartDir,omitempty"`
}

func (m *ProcArg) Reset()                    { *m = ProcArg{} }
func (m *ProcArg) String() string            { return proto.CompactTextString(m) }
func (*ProcArg) ProtoMessage()               {}
func (*ProcArg) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func init() {
	proto.RegisterType((*ProcArg)(nil), "mess.ProcArg")
}

func init() { proto.RegisterFile("agent/report/mess/procArg.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 92 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x4c, 0x4f, 0xcd,
	0x2b, 0xd1, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0xd1, 0xcf, 0x4d, 0x2d, 0x2e, 0xd6, 0x2f, 0x28,
	0xca, 0x4f, 0x76, 0x2c, 0x4a, 0xd7, 0x03, 0xd2, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0x31, 0x25, 0x55,
	0x2e, 0xf6, 0x00, 0x88, 0xb0, 0x90, 0x14, 0x17, 0x47, 0x70, 0x49, 0x62, 0x51, 0x89, 0x4b, 0x66,
	0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0x9f, 0xc4, 0x06, 0xd6, 0x63, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x70, 0x30, 0x30, 0x0a, 0x56, 0x00, 0x00, 0x00,
}