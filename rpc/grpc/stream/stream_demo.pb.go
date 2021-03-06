// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream_demo.proto

package com_beyanger_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StreamReqData struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamReqData) Reset()         { *m = StreamReqData{} }
func (m *StreamReqData) String() string { return proto.CompactTextString(m) }
func (*StreamReqData) ProtoMessage()    {}
func (*StreamReqData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fea6c45b03f4b8c5, []int{0}
}

func (m *StreamReqData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamReqData.Unmarshal(m, b)
}
func (m *StreamReqData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamReqData.Marshal(b, m, deterministic)
}
func (m *StreamReqData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamReqData.Merge(m, src)
}
func (m *StreamReqData) XXX_Size() int {
	return xxx_messageInfo_StreamReqData.Size(m)
}
func (m *StreamReqData) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamReqData.DiscardUnknown(m)
}

var xxx_messageInfo_StreamReqData proto.InternalMessageInfo

func (m *StreamReqData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type StreamResData struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResData) Reset()         { *m = StreamResData{} }
func (m *StreamResData) String() string { return proto.CompactTextString(m) }
func (*StreamResData) ProtoMessage()    {}
func (*StreamResData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fea6c45b03f4b8c5, []int{1}
}

func (m *StreamResData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResData.Unmarshal(m, b)
}
func (m *StreamResData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResData.Marshal(b, m, deterministic)
}
func (m *StreamResData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResData.Merge(m, src)
}
func (m *StreamResData) XXX_Size() int {
	return xxx_messageInfo_StreamResData.Size(m)
}
func (m *StreamResData) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResData.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResData proto.InternalMessageInfo

func (m *StreamResData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamReqData)(nil), "com.beyanger.service.StreamReqData")
	proto.RegisterType((*StreamResData)(nil), "com.beyanger.service.StreamResData")
}

func init() { proto.RegisterFile("stream_demo.proto", fileDescriptor_fea6c45b03f4b8c5) }

var fileDescriptor_fea6c45b03f4b8c5 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0x8d, 0x4f, 0x49, 0xcd, 0xcd, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x49,
	0xce, 0xcf, 0xd5, 0x4b, 0x4a, 0xad, 0x4c, 0xcc, 0x4b, 0x4f, 0x2d, 0xd2, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x55, 0x52, 0xe6, 0xe2, 0x0d, 0x06, 0x2b, 0x0d, 0x4a, 0x2d, 0x74, 0x49, 0x2c,
	0x49, 0x14, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x02, 0xb3, 0x91, 0x15, 0x15, 0xe3, 0x52, 0x64, 0x34, 0x87, 0x89, 0x8b, 0xdd, 0xbd, 0x28, 0x35,
	0xb5, 0x24, 0xb5, 0x48, 0x28, 0x92, 0x8b, 0xd3, 0x3d, 0xb5, 0x04, 0xa2, 0x47, 0x48, 0x59, 0x0f,
	0x9b, 0xcd, 0x7a, 0x28, 0xd6, 0x4a, 0x11, 0x50, 0x04, 0xb6, 0x56, 0x89, 0xc1, 0x80, 0x11, 0x64,
	0x74, 0x40, 0x29, 0x0d, 0x8c, 0xd6, 0x60, 0x14, 0x8a, 0xe6, 0xe2, 0x74, 0xcc, 0xc9, 0xa1, 0x85,
	0xd1, 0x06, 0x8c, 0x49, 0x6c, 0xe0, 0x58, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x73,
	0x28, 0x77, 0x9a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	GetStream(ctx context.Context, in *StreamReqData, opts ...grpc.CallOption) (Greeter_GetStreamClient, error)
	PutStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_PutStreamClient, error)
	AllStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_AllStreamClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) GetStream(ctx context.Context, in *StreamReqData, opts ...grpc.CallOption) (Greeter_GetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[0], "/com.beyanger.service.Greeter/GetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterGetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_GetStreamClient interface {
	Recv() (*StreamResData, error)
	grpc.ClientStream
}

type greeterGetStreamClient struct {
	grpc.ClientStream
}

func (x *greeterGetStreamClient) Recv() (*StreamResData, error) {
	m := new(StreamResData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) PutStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_PutStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[1], "/com.beyanger.service.Greeter/PutStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterPutStreamClient{stream}
	return x, nil
}

type Greeter_PutStreamClient interface {
	Send(*StreamReqData) error
	CloseAndRecv() (*StreamResData, error)
	grpc.ClientStream
}

type greeterPutStreamClient struct {
	grpc.ClientStream
}

func (x *greeterPutStreamClient) Send(m *StreamReqData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterPutStreamClient) CloseAndRecv() (*StreamResData, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamResData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) AllStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_AllStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[2], "/com.beyanger.service.Greeter/AllStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterAllStreamClient{stream}
	return x, nil
}

type Greeter_AllStreamClient interface {
	Send(*StreamReqData) error
	Recv() (*StreamResData, error)
	grpc.ClientStream
}

type greeterAllStreamClient struct {
	grpc.ClientStream
}

func (x *greeterAllStreamClient) Send(m *StreamReqData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterAllStreamClient) Recv() (*StreamResData, error) {
	m := new(StreamResData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	GetStream(*StreamReqData, Greeter_GetStreamServer) error
	PutStream(Greeter_PutStreamServer) error
	AllStream(Greeter_AllStreamServer) error
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamReqData)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).GetStream(m, &greeterGetStreamServer{stream})
}

type Greeter_GetStreamServer interface {
	Send(*StreamResData) error
	grpc.ServerStream
}

type greeterGetStreamServer struct {
	grpc.ServerStream
}

func (x *greeterGetStreamServer) Send(m *StreamResData) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_PutStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).PutStream(&greeterPutStreamServer{stream})
}

type Greeter_PutStreamServer interface {
	SendAndClose(*StreamResData) error
	Recv() (*StreamReqData, error)
	grpc.ServerStream
}

type greeterPutStreamServer struct {
	grpc.ServerStream
}

func (x *greeterPutStreamServer) SendAndClose(m *StreamResData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterPutStreamServer) Recv() (*StreamReqData, error) {
	m := new(StreamReqData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_AllStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).AllStream(&greeterAllStreamServer{stream})
}

type Greeter_AllStreamServer interface {
	Send(*StreamResData) error
	Recv() (*StreamReqData, error)
	grpc.ServerStream
}

type greeterAllStreamServer struct {
	grpc.ServerStream
}

func (x *greeterAllStreamServer) Send(m *StreamResData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterAllStreamServer) Recv() (*StreamReqData, error) {
	m := new(StreamReqData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.beyanger.service.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _Greeter_GetStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PutStream",
			Handler:       _Greeter_PutStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AllStream",
			Handler:       _Greeter_AllStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream_demo.proto",
}
