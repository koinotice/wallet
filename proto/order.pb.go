// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package gravatar

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

type CloseRequest struct {
	Roomid               int64    `protobuf:"varint,1,opt,name=roomid,proto3" json:"roomid,omitempty"`
	Orderid              int64    `protobuf:"varint,2,opt,name=orderid,proto3" json:"orderid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseRequest) Reset()         { *m = CloseRequest{} }
func (m *CloseRequest) String() string { return proto.CompactTextString(m) }
func (*CloseRequest) ProtoMessage()    {}
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *CloseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseRequest.Unmarshal(m, b)
}
func (m *CloseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseRequest.Marshal(b, m, deterministic)
}
func (m *CloseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseRequest.Merge(m, src)
}
func (m *CloseRequest) XXX_Size() int {
	return xxx_messageInfo_CloseRequest.Size(m)
}
func (m *CloseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloseRequest proto.InternalMessageInfo

func (m *CloseRequest) GetRoomid() int64 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

func (m *CloseRequest) GetOrderid() int64 {
	if m != nil {
		return m.Orderid
	}
	return 0
}

type CloseResponse struct {
	Txhash               string   `protobuf:"bytes,2,opt,name=txhash,proto3" json:"txhash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseResponse) Reset()         { *m = CloseResponse{} }
func (m *CloseResponse) String() string { return proto.CompactTextString(m) }
func (*CloseResponse) ProtoMessage()    {}
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *CloseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseResponse.Unmarshal(m, b)
}
func (m *CloseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseResponse.Marshal(b, m, deterministic)
}
func (m *CloseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseResponse.Merge(m, src)
}
func (m *CloseResponse) XXX_Size() int {
	return xxx_messageInfo_CloseResponse.Size(m)
}
func (m *CloseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloseResponse proto.InternalMessageInfo

func (m *CloseResponse) GetTxhash() string {
	if m != nil {
		return m.Txhash
	}
	return ""
}

func init() {
	proto.RegisterType((*CloseRequest)(nil), "gravatar.CloseRequest")
	proto.RegisterType((*CloseResponse)(nil), "gravatar.CloseResponse")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x2f, 0x4a, 0x49,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x2f, 0x4a, 0x2c, 0x4b, 0x2c, 0x49,
	0x2c, 0x52, 0x72, 0xe0, 0xe2, 0x71, 0xce, 0xc9, 0x2f, 0x4e, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x12, 0xe3, 0x62, 0x2b, 0xca, 0xcf, 0xcf, 0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x0e, 0x82, 0xf2, 0x84, 0x24, 0xb8, 0xd8, 0xc1, 0x06, 0x64, 0xa6, 0x48, 0x30, 0x81, 0x25,
	0x60, 0x5c, 0x25, 0x75, 0x2e, 0x5e, 0xa8, 0x09, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x20, 0x23,
	0x4a, 0x2a, 0x32, 0x12, 0x8b, 0x33, 0xc0, 0x2a, 0x39, 0x83, 0xa0, 0x3c, 0x23, 0x37, 0x2e, 0x0e,
	0xa7, 0xd4, 0x12, 0x7f, 0x90, 0x36, 0x21, 0x2b, 0x2e, 0x56, 0xb0, 0x26, 0x21, 0x31, 0x3d, 0x98,
	0x53, 0xf4, 0x90, 0xdd, 0x21, 0x25, 0x8e, 0x21, 0x0e, 0x31, 0x5d, 0x89, 0x21, 0x89, 0x0d, 0xec,
	0x07, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xac, 0xc5, 0x1e, 0xd2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BetOrderClient is the client API for BetOrder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BetOrderClient interface {
	Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error)
}

type betOrderClient struct {
	cc *grpc.ClientConn
}

func NewBetOrderClient(cc *grpc.ClientConn) BetOrderClient {
	return &betOrderClient{cc}
}

func (c *betOrderClient) Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error) {
	out := new(CloseResponse)
	err := c.cc.Invoke(ctx, "/gravatar.BetOrder/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BetOrderServer is the server API for BetOrder service.
type BetOrderServer interface {
	Close(context.Context, *CloseRequest) (*CloseResponse, error)
}

func RegisterBetOrderServer(s *grpc.Server, srv BetOrderServer) {
	s.RegisterService(&_BetOrder_serviceDesc, srv)
}

func _BetOrder_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetOrderServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gravatar.BetOrder/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetOrderServer).Close(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BetOrder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gravatar.BetOrder",
	HandlerType: (*BetOrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Close",
			Handler:    _BetOrder_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
