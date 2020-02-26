// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mouselive.proto

package mouselive

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ba8147a9cd784f, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Additional           string   `protobuf:"bytes,2,opt,name=additional,proto3" json:"additional,omitempty"`
	Number               int64    `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ba8147a9cd784f, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloReply) GetAdditional() string {
	if m != nil {
		return m.Additional
	}
	return ""
}

func (m *HelloReply) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "mouselive.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "mouselive.HelloReply")
}

func init() { proto.RegisterFile("mouselive.proto", fileDescriptor_d2ba8147a9cd784f) }

var fileDescriptor_d2ba8147a9cd784f = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0xcd, 0x2f, 0x2d,
	0x4e, 0xcd, 0xc9, 0x2c, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x29, 0x71, 0xf1, 0x78, 0xa4, 0xe6, 0xe4, 0xe4, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08,
	0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9,
	0x4a, 0x71, 0x5c, 0x5c, 0x50, 0x35, 0x05, 0x39, 0x95, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5,
	0xc5, 0x89, 0xe9, 0x30, 0x45, 0x30, 0xae, 0x90, 0x1c, 0x17, 0x57, 0x62, 0x4a, 0x4a, 0x66, 0x49,
	0x66, 0x7e, 0x5e, 0x62, 0x8e, 0x04, 0x13, 0x58, 0x12, 0x49, 0x44, 0x48, 0x8c, 0x8b, 0x2d, 0xaf,
	0x34, 0x37, 0x29, 0xb5, 0x48, 0x82, 0x59, 0x81, 0x51, 0x83, 0x39, 0x08, 0xca, 0x33, 0xea, 0x62,
	0xe4, 0x62, 0x77, 0x0f, 0x0a, 0x70, 0x76, 0x2c, 0xc8, 0x14, 0xb2, 0xe1, 0xe2, 0x08, 0x4e, 0xac,
	0x04, 0x5b, 0x27, 0x24, 0xae, 0x87, 0x70, 0x38, 0xb2, 0x23, 0xa5, 0x44, 0x31, 0x25, 0x0a, 0x72,
	0x2a, 0x95, 0x18, 0x84, 0x1c, 0xb9, 0x78, 0x61, 0xba, 0x1d, 0xd3, 0x13, 0x33, 0xf3, 0x48, 0x37,
	0x22, 0x89, 0x0d, 0x1c, 0x44, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xd1, 0x37, 0x7a,
	0x35, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GRPCApiClient is the client API for GRPCApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GRPCApiClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type gRPCApiClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCApiClient(cc grpc.ClientConnInterface) GRPCApiClient {
	return &gRPCApiClient{cc}
}

func (c *gRPCApiClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/mouselive.GRPCApi/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCApiClient) SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/mouselive.GRPCApi/SayHelloAgain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCApiServer is the server API for GRPCApi service.
type GRPCApiServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHelloAgain(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedGRPCApiServer can be embedded to have forward compatible implementations.
type UnimplementedGRPCApiServer struct {
}

func (*UnimplementedGRPCApiServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedGRPCApiServer) SayHelloAgain(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloAgain not implemented")
}

func RegisterGRPCApiServer(s *grpc.Server, srv GRPCApiServer) {
	s.RegisterService(&_GRPCApi_serviceDesc, srv)
}

func _GRPCApi_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCApiServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mouselive.GRPCApi/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCApiServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCApi_SayHelloAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCApiServer).SayHelloAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mouselive.GRPCApi/SayHelloAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCApiServer).SayHelloAgain(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GRPCApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mouselive.GRPCApi",
	HandlerType: (*GRPCApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GRPCApi_SayHello_Handler,
		},
		{
			MethodName: "SayHelloAgain",
			Handler:    _GRPCApi_SayHelloAgain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mouselive.proto",
}
