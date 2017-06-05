// Code generated by protoc-gen-go.
// source: agent/report/report.proto
// DO NOT EDIT!

package report

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import mess "agent/report/mess"
import mess1 "agent/report/mess"
import mess2 "agent/report/mess"
import mess3 "agent/report/mess"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Reply struct {
	Mess int64 `protobuf:"varint,1,opt,name=Mess" json:"Mess,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type ReplyMCSD struct {
	Mess int64 `protobuf:"varint,1,opt,name=Mess" json:"Mess,omitempty"`
}

func (m *ReplyMCSD) Reset()                    { *m = ReplyMCSD{} }
func (m *ReplyMCSD) String() string            { return proto.CompactTextString(m) }
func (*ReplyMCSD) ProtoMessage()               {}
func (*ReplyMCSD) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

type ReplyRT struct {
	Mess int64 `protobuf:"varint,1,opt,name=Mess" json:"Mess,omitempty"`
}

func (m *ReplyRT) Reset()                    { *m = ReplyRT{} }
func (m *ReplyRT) String() string            { return proto.CompactTextString(m) }
func (*ReplyRT) ProtoMessage()               {}
func (*ReplyRT) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

type Report struct {
	GUID       string     `protobuf:"bytes,1,opt,name=GUID" json:"GUID,omitempty"`
	ReportTime string     `protobuf:"bytes,2,opt,name=ReportTime" json:"ReportTime,omitempty"`
	Host       *Host      `protobuf:"bytes,3,opt,name=Host" json:"Host,omitempty"`
	Process    []*Process `protobuf:"bytes,4,rep,name=Process" json:"Process,omitempty"`
	NetWork    *NetWork   `protobuf:"bytes,5,opt,name=NetWork" json:"NetWork,omitempty"`
	Agent      *Agent     `protobuf:"bytes,6,opt,name=Agent" json:"Agent,omitempty"`
	Config     []*Config  `protobuf:"bytes,7,rep,name=Config" json:"Config,omitempty"`
}

func (m *Report) Reset()                    { *m = Report{} }
func (m *Report) String() string            { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()               {}
func (*Report) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *Report) GetHost() *Host {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Report) GetProcess() []*Process {
	if m != nil {
		return m.Process
	}
	return nil
}

func (m *Report) GetNetWork() *NetWork {
	if m != nil {
		return m.NetWork
	}
	return nil
}

func (m *Report) GetAgent() *Agent {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *Report) GetConfig() []*Config {
	if m != nil {
		return m.Config
	}
	return nil
}

type MCSD struct {
	GUID       string        `protobuf:"bytes,1,opt,name=GUID" json:"GUID,omitempty"`
	ReportTime string        `protobuf:"bytes,2,opt,name=ReportTime" json:"ReportTime,omitempty"`
	Mem        *mess.Mem     `protobuf:"bytes,3,opt,name=Mem" json:"Mem,omitempty"`
	Cpu        *mess1.Cpu    `protobuf:"bytes,4,opt,name=Cpu" json:"Cpu,omitempty"`
	Swap       *mess2.Swap   `protobuf:"bytes,5,opt,name=Swap" json:"Swap,omitempty"`
	Disk       []*mess3.Disk `protobuf:"bytes,6,rep,name=Disk" json:"Disk,omitempty"`
}

func (m *MCSD) Reset()                    { *m = MCSD{} }
func (m *MCSD) String() string            { return proto.CompactTextString(m) }
func (*MCSD) ProtoMessage()               {}
func (*MCSD) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *MCSD) GetMem() *mess.Mem {
	if m != nil {
		return m.Mem
	}
	return nil
}

func (m *MCSD) GetCpu() *mess1.Cpu {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *MCSD) GetSwap() *mess2.Swap {
	if m != nil {
		return m.Swap
	}
	return nil
}

func (m *MCSD) GetDisk() []*mess3.Disk {
	if m != nil {
		return m.Disk
	}
	return nil
}

type ProcessRT struct {
	StateRT string   `protobuf:"bytes,1,opt,name=StateRT" json:"StateRT,omitempty"`
	TimeRT  string   `protobuf:"bytes,2,opt,name=TimeRT" json:"TimeRT,omitempty"`
	Process *Process `protobuf:"bytes,3,opt,name=Process" json:"Process,omitempty"`
}

func (m *ProcessRT) Reset()                    { *m = ProcessRT{} }
func (m *ProcessRT) String() string            { return proto.CompactTextString(m) }
func (*ProcessRT) ProtoMessage()               {}
func (*ProcessRT) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *ProcessRT) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

func init() {
	proto.RegisterType((*Reply)(nil), "report.Reply")
	proto.RegisterType((*ReplyMCSD)(nil), "report.ReplyMCSD")
	proto.RegisterType((*ReplyRT)(nil), "report.ReplyRT")
	proto.RegisterType((*Report)(nil), "report.Report")
	proto.RegisterType((*MCSD)(nil), "report.MCSD")
	proto.RegisterType((*ProcessRT)(nil), "report.ProcessRT")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Greeter service

type GreeterClient interface {
	SendReport(ctx context.Context, in *Report, opts ...grpc.CallOption) (*Reply, error)
	SendMCSD(ctx context.Context, in *MCSD, opts ...grpc.CallOption) (*ReplyMCSD, error)
	SendReportRT(ctx context.Context, in *ProcessRT, opts ...grpc.CallOption) (*ReplyRT, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SendReport(ctx context.Context, in *Report, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/report.Greeter/SendReport", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SendMCSD(ctx context.Context, in *MCSD, opts ...grpc.CallOption) (*ReplyMCSD, error) {
	out := new(ReplyMCSD)
	err := grpc.Invoke(ctx, "/report.Greeter/SendMCSD", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SendReportRT(ctx context.Context, in *ProcessRT, opts ...grpc.CallOption) (*ReplyRT, error) {
	out := new(ReplyRT)
	err := grpc.Invoke(ctx, "/report.Greeter/SendReportRT", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	SendReport(context.Context, *Report) (*Reply, error)
	SendMCSD(context.Context, *MCSD) (*ReplyMCSD, error)
	SendReportRT(context.Context, *ProcessRT) (*ReplyRT, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SendReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Report)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SendReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.Greeter/SendReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SendReport(ctx, req.(*Report))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SendMCSD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MCSD)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SendMCSD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.Greeter/SendMCSD",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SendMCSD(ctx, req.(*MCSD))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SendReportRT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessRT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SendReportRT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.Greeter/SendReportRT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SendReportRT(ctx, req.(*ProcessRT))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "report.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendReport",
			Handler:    _Greeter_SendReport_Handler,
		},
		{
			MethodName: "SendMCSD",
			Handler:    _Greeter_SendMCSD_Handler,
		},
		{
			MethodName: "SendReportRT",
			Handler:    _Greeter_SendReportRT_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor5,
}

func init() { proto.RegisterFile("agent/report/report.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x9b, 0x38, 0xb1, 0xc9, 0xd0, 0x52, 0x75, 0x0f, 0xb0, 0x38, 0xfc, 0x89, 0x8c, 0x84,
	0xe0, 0x40, 0x90, 0x0a, 0x2f, 0x00, 0x89, 0xd4, 0x22, 0x11, 0x84, 0x26, 0x46, 0x3d, 0x87, 0x74,
	0x69, 0x23, 0x9a, 0xd8, 0xb2, 0x5d, 0x55, 0x3d, 0xf2, 0x20, 0x5c, 0x78, 0x06, 0x1e, 0x90, 0xd9,
	0xf1, 0x6c, 0xea, 0x2d, 0xc9, 0xa1, 0x97, 0x64, 0xe7, 0xfb, 0x7d, 0xd9, 0x9d, 0xfd, 0x66, 0x03,
	0x8f, 0x67, 0x67, 0x66, 0x55, 0xbd, 0x2d, 0x4c, 0x9e, 0x15, 0xee, 0x6b, 0x98, 0x17, 0x59, 0x95,
	0xa9, 0xb0, 0xae, 0xe2, 0x47, 0x9e, 0xe5, 0x3c, 0x2b, 0xc5, 0x10, 0xc7, 0x1e, 0x20, 0x6d, 0x6e,
	0xca, 0x72, 0x23, 0x5b, 0x99, 0xea, 0x2a, 0x2b, 0x7e, 0x0a, 0xd3, 0x1e, 0xe3, 0x42, 0x88, 0xdf,
	0xcd, 0x3c, 0x5b, 0xfd, 0x58, 0x9c, 0x09, 0xea, 0x7b, 0x68, 0x49, 0x27, 0xd1, 0xc7, 0x72, 0x3b,
	0x9c, 0xe7, 0x97, 0x02, 0x9f, 0xfc, 0x0f, 0xcb, 0xab, 0x59, 0xbe, 0x9d, 0x9e, 0x2e, 0x4a, 0x69,
	0x35, 0xe9, 0x43, 0x17, 0x4d, 0x7e, 0x71, 0xad, 0x14, 0x74, 0x26, 0xc4, 0x74, 0x6b, 0xd0, 0x7a,
	0x15, 0x20, 0xaf, 0x93, 0xe7, 0xd0, 0x63, 0x38, 0x19, 0x4d, 0xc7, 0x1b, 0x0d, 0x4f, 0x21, 0x62,
	0x03, 0xa6, 0x1b, 0xf1, 0xaf, 0x36, 0x84, 0xc8, 0xe7, 0x5a, 0x7c, 0xf4, 0xed, 0xd3, 0x98, 0x71,
	0x0f, 0x79, 0xad, 0x9e, 0x01, 0xd4, 0x34, 0x5d, 0x2c, 0x8d, 0x6e, 0x33, 0x69, 0x28, 0x6a, 0x00,
	0x9d, 0x63, 0x1a, 0x86, 0x0e, 0x88, 0xdc, 0x3f, 0xdc, 0x1d, 0xca, 0xf0, 0xac, 0x86, 0x4c, 0xd4,
	0x6b, 0x88, 0xbe, 0xd6, 0x53, 0xd1, 0x9d, 0x41, 0x40, 0xa6, 0x7d, 0x67, 0x12, 0x19, 0x1d, 0xb7,
	0xd6, 0x2f, 0xa6, 0x3a, 0xa1, 0x21, 0xe9, 0x2e, 0xef, 0xb7, 0xb6, 0x8a, 0x8c, 0x8e, 0xab, 0x17,
	0xd0, 0xfd, 0x60, 0x33, 0xd3, 0x21, 0x1b, 0xf7, 0x9c, 0x91, 0x45, 0xac, 0x99, 0x7a, 0x09, 0xe1,
	0x88, 0xc7, 0xa7, 0x23, 0x3e, 0xf9, 0x81, 0x73, 0xd5, 0x2a, 0x0a, 0x4d, 0xfe, 0xb6, 0x28, 0x18,
	0xc9, 0xef, 0xce, 0x09, 0xf4, 0x21, 0x98, 0x98, 0xa5, 0x04, 0xd0, 0x1b, 0xda, 0xe1, 0x0d, 0x49,
	0x40, 0xab, 0x5a, 0x38, 0xca, 0x2f, 0xe9, 0xe2, 0x0d, 0x48, 0x02, 0x5a, 0x95, 0x76, 0xee, 0x4c,
	0xe9, 0x0d, 0xc8, 0x5d, 0xa1, 0xa6, 0x56, 0x41, 0xd6, 0x2d, 0x1f, 0xd3, 0x2b, 0xa0, 0x2b, 0x06,
	0x37, 0xdc, 0x2a, 0xc8, 0x7a, 0x72, 0x0e, 0x3d, 0x17, 0x61, 0xaa, 0x34, 0x44, 0xd3, 0x6a, 0x56,
	0x19, 0x4c, 0xa5, 0x7b, 0x57, 0xaa, 0x87, 0x10, 0xda, 0x46, 0x09, 0xd4, 0xcd, 0x4b, 0xd5, 0x1c,
	0x4c, 0xe0, 0xa7, 0x7d, 0x7b, 0x30, 0x87, 0xbf, 0x5b, 0x10, 0x1d, 0x15, 0xc6, 0x54, 0xa6, 0x50,
	0x6f, 0x00, 0xa6, 0x66, 0x75, 0x2a, 0x6f, 0x66, 0x1d, 0x69, 0x5d, 0xc7, 0x7b, 0x8d, 0xfa, 0xe2,
	0x3a, 0xd9, 0x21, 0xfb, 0x3d, 0x6b, 0xe7, 0x78, 0xd7, 0xcf, 0xc3, 0x56, 0xf1, 0x81, 0x67, 0xb5,
	0x12, 0xd9, 0xdf, 0xc3, 0xee, 0xcd, 0xee, 0xd4, 0xe4, 0xc1, 0xed, 0x9e, 0xd2, 0x78, 0xdf, 0xfb,
	0x1d, 0xa6, 0xc9, 0xce, 0xc7, 0xf0, 0x4f, 0x3b, 0x38, 0xfe, 0x7c, 0xf2, 0x3d, 0xe4, 0x3f, 0xcc,
	0xbb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0xf4, 0xf9, 0x93, 0x51, 0x04, 0x00, 0x00,
}