// Code generated by protoc-gen-go. DO NOT EDIT.
// source: time/time.proto

package time

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"

	common "github.com/talos-systems/talos/api/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The response message containing the ntp server
type TimeRequest struct {
	Server               string   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeRequest) Reset()         { *m = TimeRequest{} }
func (m *TimeRequest) String() string { return proto.CompactTextString(m) }
func (*TimeRequest) ProtoMessage()    {}
func (*TimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ed1ef5b20ef4ce, []int{0}
}

func (m *TimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRequest.Unmarshal(m, b)
}

func (m *TimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRequest.Marshal(b, m, deterministic)
}

func (m *TimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRequest.Merge(m, src)
}

func (m *TimeRequest) XXX_Size() int {
	return xxx_messageInfo_TimeRequest.Size(m)
}

func (m *TimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRequest proto.InternalMessageInfo

func (m *TimeRequest) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

type Time struct {
	Metadata             *common.Metadata     `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Server               string               `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	Localtime            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=localtime,proto3" json:"localtime,omitempty"`
	Remotetime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=remotetime,proto3" json:"remotetime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Time) Reset()         { *m = Time{} }
func (m *Time) String() string { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()    {}
func (*Time) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ed1ef5b20ef4ce, []int{1}
}

func (m *Time) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Time.Unmarshal(m, b)
}

func (m *Time) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Time.Marshal(b, m, deterministic)
}

func (m *Time) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Time.Merge(m, src)
}

func (m *Time) XXX_Size() int {
	return xxx_messageInfo_Time.Size(m)
}

func (m *Time) XXX_DiscardUnknown() {
	xxx_messageInfo_Time.DiscardUnknown(m)
}

var xxx_messageInfo_Time proto.InternalMessageInfo

func (m *Time) GetMetadata() *common.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Time) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *Time) GetLocaltime() *timestamp.Timestamp {
	if m != nil {
		return m.Localtime
	}
	return nil
}

func (m *Time) GetRemotetime() *timestamp.Timestamp {
	if m != nil {
		return m.Remotetime
	}
	return nil
}

// The response message containing the ntp server, time, and offset
type TimeResponse struct {
	Messages             []*Time  `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeResponse) Reset()         { *m = TimeResponse{} }
func (m *TimeResponse) String() string { return proto.CompactTextString(m) }
func (*TimeResponse) ProtoMessage()    {}
func (*TimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ed1ef5b20ef4ce, []int{2}
}

func (m *TimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeResponse.Unmarshal(m, b)
}

func (m *TimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeResponse.Marshal(b, m, deterministic)
}

func (m *TimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeResponse.Merge(m, src)
}

func (m *TimeResponse) XXX_Size() int {
	return xxx_messageInfo_TimeResponse.Size(m)
}

func (m *TimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TimeResponse proto.InternalMessageInfo

func (m *TimeResponse) GetMessages() []*Time {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*TimeRequest)(nil), "time.TimeRequest")
	proto.RegisterType((*Time)(nil), "time.Time")
	proto.RegisterType((*TimeResponse)(nil), "time.TimeResponse")
}

func init() { proto.RegisterFile("time/time.proto", fileDescriptor_e7ed1ef5b20ef4ce) }

var fileDescriptor_e7ed1ef5b20ef4ce = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xa9, 0x1b, 0xd3, 0xbd, 0x0d, 0xd4, 0x08, 0x63, 0xd4, 0x83, 0x63, 0xa0, 0xee, 0xa0,
	0x2d, 0x54, 0x10, 0xf1, 0xe6, 0xc4, 0xa3, 0x20, 0x75, 0x27, 0x6f, 0x59, 0x7d, 0x76, 0xc1, 0x66,
	0x89, 0x7d, 0xe9, 0x60, 0xff, 0x9e, 0x7f, 0x99, 0x24, 0xe9, 0x66, 0x51, 0xc1, 0xcb, 0xd6, 0xf7,
	0xde, 0xef, 0x7d, 0xf9, 0xf2, 0x05, 0xf6, 0x8d, 0x90, 0x18, 0xdb, 0x9f, 0x48, 0x97, 0xca, 0x28,
	0xd6, 0xb6, 0xdf, 0xe1, 0x71, 0xae, 0x54, 0x5e, 0x60, 0xec, 0x7a, 0xf3, 0xea, 0x2d, 0x46, 0xa9,
	0xcd, 0xda, 0x23, 0xe1, 0xc9, 0xcf, 0xa1, 0x5d, 0x21, 0xc3, 0xa5, 0xae, 0x81, 0xa3, 0x4c, 0x49,
	0xa9, 0x96, 0xb1, 0xff, 0xf3, 0xcd, 0xf1, 0x29, 0xf4, 0x66, 0x42, 0x62, 0x8a, 0x1f, 0x15, 0x92,
	0x61, 0x03, 0xe8, 0x10, 0x96, 0x2b, 0x2c, 0x87, 0xc1, 0x28, 0x98, 0x74, 0xd3, 0xba, 0x1a, 0x7f,
	0x06, 0xd0, 0xb6, 0x1c, 0xbb, 0x80, 0x3d, 0x89, 0x86, 0xbf, 0x72, 0xc3, 0x1d, 0xd2, 0x4b, 0x0e,
	0xa2, 0x5a, 0xf0, 0xb1, 0xee, 0xa7, 0x5b, 0xa2, 0x21, 0xb7, 0xd3, 0x94, 0x63, 0x37, 0xd0, 0x2d,
	0x54, 0xc6, 0x0b, 0x6b, 0x71, 0xd8, 0x72, 0x32, 0x61, 0xe4, 0xfd, 0x47, 0x1b, 0xff, 0xd1, 0x6c,
	0xe3, 0x3f, 0xfd, 0x86, 0xd9, 0x2d, 0x40, 0x89, 0x52, 0x19, 0x74, 0xab, 0xed, 0x7f, 0x57, 0x1b,
	0xf4, 0xf8, 0x1a, 0xfa, 0xfe, 0xae, 0xa4, 0xd5, 0x92, 0x90, 0x9d, 0xd9, 0xbb, 0x10, 0xf1, 0x1c,
	0x69, 0x18, 0x8c, 0x5a, 0x93, 0x5e, 0x02, 0x91, 0xcb, 0xdc, 0x51, 0xdb, 0x59, 0x52, 0xf9, 0x8c,
	0x9e, 0xb1, 0x5c, 0x89, 0x0c, 0x59, 0x52, 0x47, 0x31, 0xf8, 0x75, 0xec, 0x83, 0x7d, 0x8e, 0x90,
	0x35, 0x44, 0x36, 0x47, 0x25, 0xd0, 0xb5, 0xf5, 0xfd, 0x02, 0xb3, 0x77, 0x76, 0xd8, 0x04, 0x5c,
	0xee, 0x7f, 0xed, 0x4c, 0xa7, 0xd0, 0xcf, 0x94, 0xf4, 0x03, 0xae, 0xc5, 0x74, 0xd7, 0x4e, 0xef,
	0xb4, 0x78, 0x0a, 0x5e, 0xce, 0x73, 0x61, 0x16, 0xd5, 0xdc, 0x26, 0x1f, 0x1b, 0x5e, 0x28, 0xba,
	0xa4, 0x35, 0x19, 0x94, 0xe4, 0xab, 0x98, 0x6b, 0xe1, 0x9e, 0x7f, 0xde, 0x71, 0xde, 0xae, 0xbe,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xc9, 0x4e, 0xf0, 0x51, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TimeServiceClient is the client API for TimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimeServiceClient interface {
	Time(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TimeResponse, error)
	TimeCheck(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*TimeResponse, error)
}

type timeServiceClient struct {
	cc *grpc.ClientConn
}

func NewTimeServiceClient(cc *grpc.ClientConn) TimeServiceClient {
	return &timeServiceClient{cc}
}

func (c *timeServiceClient) Time(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TimeResponse, error) {
	out := new(TimeResponse)
	err := c.cc.Invoke(ctx, "/time.TimeService/Time", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeServiceClient) TimeCheck(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*TimeResponse, error) {
	out := new(TimeResponse)
	err := c.cc.Invoke(ctx, "/time.TimeService/TimeCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeServiceServer is the server API for TimeService service.
type TimeServiceServer interface {
	Time(context.Context, *empty.Empty) (*TimeResponse, error)
	TimeCheck(context.Context, *TimeRequest) (*TimeResponse, error)
}

func RegisterTimeServiceServer(s *grpc.Server, srv TimeServiceServer) {
	s.RegisterService(&_TimeService_serviceDesc, srv)
}

func _TimeService_Time_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).Time(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/time.TimeService/Time",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).Time(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeService_TimeCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).TimeCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/time.TimeService/TimeCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).TimeCheck(ctx, req.(*TimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TimeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "time.TimeService",
	HandlerType: (*TimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Time",
			Handler:    _TimeService_Time_Handler,
		},
		{
			MethodName: "TimeCheck",
			Handler:    _TimeService_TimeCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "time/time.proto",
}
