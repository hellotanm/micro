// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/service/storeexample/proto/example.proto

package srv_test_example

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

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_060cb6963cd2b8f6, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_060cb6963cd2b8f6, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "srv.test_example.Request")
	proto.RegisterType((*Response)(nil), "srv.test_example.Response")
}

func init() {
	proto.RegisterFile("test/service/storeexample/proto/example.proto", fileDescriptor_060cb6963cd2b8f6)
}

var fileDescriptor_060cb6963cd2b8f6 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2d, 0x49, 0x2d, 0x2e,
	0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4d, 0xad,
	0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x87, 0xf2, 0xf4, 0xc0,
	0x3c, 0x21, 0x81, 0xe2, 0xa2, 0x32, 0x3d, 0x90, 0x96, 0x78, 0xa8, 0xb8, 0x12, 0x27, 0x17, 0x7b,
	0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x92, 0x0c, 0x17, 0x47, 0x50, 0x6a, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x90, 0x00, 0x17, 0x73, 0x6e, 0x71, 0xba, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x88, 0x69, 0x14, 0xc0, 0xc5, 0xee, 0x0a, 0xd1, 0x23, 0xe4, 0xca, 0xc5, 0x15, 0x92, 0x5a,
	0x5c, 0xe2, 0x5a, 0x51, 0x90, 0x59, 0x54, 0x29, 0x24, 0xa9, 0x87, 0x6e, 0xa8, 0x1e, 0xd4, 0x44,
	0x29, 0x29, 0x6c, 0x52, 0x10, 0x1b, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x6e, 0x32, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x4c, 0xad, 0xe6, 0x25, 0xc4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleClient is the client API for Example service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleClient interface {
	TestExpiry(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type exampleClient struct {
	cc *grpc.ClientConn
}

func NewExampleClient(cc *grpc.ClientConn) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) TestExpiry(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/srv.test_example.Example/TestExpiry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServer is the server API for Example service.
type ExampleServer interface {
	TestExpiry(context.Context, *Request) (*Response, error)
}

// UnimplementedExampleServer can be embedded to have forward compatible implementations.
type UnimplementedExampleServer struct {
}

func (*UnimplementedExampleServer) TestExpiry(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestExpiry not implemented")
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_TestExpiry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).TestExpiry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/srv.test_example.Example/TestExpiry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).TestExpiry(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Example_serviceDesc = grpc.ServiceDesc{
	ServiceName: "srv.test_example.Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestExpiry",
			Handler:    _Example_TestExpiry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test/service/storeexample/proto/example.proto",
}
