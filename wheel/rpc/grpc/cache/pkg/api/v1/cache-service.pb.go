// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache-service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type StoreReq struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val                  []byte   `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreReq) Reset()         { *m = StoreReq{} }
func (m *StoreReq) String() string { return proto.CompactTextString(m) }
func (*StoreReq) ProtoMessage()    {}
func (*StoreReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0ee6ef8b54c4e4, []int{0}
}

func (m *StoreReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreReq.Unmarshal(m, b)
}
func (m *StoreReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreReq.Marshal(b, m, deterministic)
}
func (m *StoreReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreReq.Merge(m, src)
}
func (m *StoreReq) XXX_Size() int {
	return xxx_messageInfo_StoreReq.Size(m)
}
func (m *StoreReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreReq.DiscardUnknown(m)
}

var xxx_messageInfo_StoreReq proto.InternalMessageInfo

func (m *StoreReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StoreReq) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

type StoreResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreResp) Reset()         { *m = StoreResp{} }
func (m *StoreResp) String() string { return proto.CompactTextString(m) }
func (*StoreResp) ProtoMessage()    {}
func (*StoreResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0ee6ef8b54c4e4, []int{1}
}

func (m *StoreResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreResp.Unmarshal(m, b)
}
func (m *StoreResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreResp.Marshal(b, m, deterministic)
}
func (m *StoreResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResp.Merge(m, src)
}
func (m *StoreResp) XXX_Size() int {
	return xxx_messageInfo_StoreResp.Size(m)
}
func (m *StoreResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResp.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResp proto.InternalMessageInfo

type GetReq struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReq) Reset()         { *m = GetReq{} }
func (m *GetReq) String() string { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()    {}
func (*GetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0ee6ef8b54c4e4, []int{2}
}

func (m *GetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReq.Unmarshal(m, b)
}
func (m *GetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReq.Marshal(b, m, deterministic)
}
func (m *GetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReq.Merge(m, src)
}
func (m *GetReq) XXX_Size() int {
	return xxx_messageInfo_GetReq.Size(m)
}
func (m *GetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetReq proto.InternalMessageInfo

func (m *GetReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResp struct {
	Val                  []byte   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResp) Reset()         { *m = GetResp{} }
func (m *GetResp) String() string { return proto.CompactTextString(m) }
func (*GetResp) ProtoMessage()    {}
func (*GetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0ee6ef8b54c4e4, []int{3}
}

func (m *GetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResp.Unmarshal(m, b)
}
func (m *GetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResp.Marshal(b, m, deterministic)
}
func (m *GetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResp.Merge(m, src)
}
func (m *GetResp) XXX_Size() int {
	return xxx_messageInfo_GetResp.Size(m)
}
func (m *GetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetResp proto.InternalMessageInfo

func (m *GetResp) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

func init() {
	proto.RegisterType((*StoreReq)(nil), "v1.StoreReq")
	proto.RegisterType((*StoreResp)(nil), "v1.StoreResp")
	proto.RegisterType((*GetReq)(nil), "v1.GetReq")
	proto.RegisterType((*GetResp)(nil), "v1.GetResp")
}

func init() { proto.RegisterFile("cache-service.proto", fileDescriptor_9b0ee6ef8b54c4e4) }

var fileDescriptor_9b0ee6ef8b54c4e4 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x4e, 0x4c, 0xce,
	0x48, 0xd5, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2a, 0x33, 0x94, 0x92, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x24, 0x95,
	0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0x14, 0x29, 0xe9, 0x71,
	0x71, 0x04, 0x97, 0xe4, 0x17, 0xa5, 0x06, 0xa5, 0x16, 0x0a, 0x09, 0x70, 0x31, 0x67, 0xa7, 0x56,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x20, 0x91, 0xb2, 0xc4, 0x1c, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0x9e, 0x20, 0x10, 0x53, 0x89, 0x9b, 0x8b, 0x13, 0xaa, 0xbe, 0xb8, 0x40, 0x49,
	0x8a, 0x8b, 0xcd, 0x3d, 0xb5, 0x04, 0xab, 0x56, 0x25, 0x69, 0x2e, 0x76, 0xb0, 0x5c, 0x71, 0x01,
	0xcc, 0x14, 0x46, 0xb8, 0x29, 0x46, 0x81, 0x5c, 0xac, 0xce, 0x20, 0x17, 0x0b, 0xa9, 0x71, 0xb1,
	0x82, 0x8d, 0x13, 0xe2, 0xd1, 0x2b, 0x33, 0xd4, 0x83, 0xb9, 0x44, 0x8a, 0x17, 0x89, 0x57, 0x5c,
	0xa0, 0xc4, 0x20, 0xa4, 0xc0, 0xc5, 0xec, 0x9e, 0x5a, 0x22, 0xc4, 0x05, 0x12, 0x87, 0x58, 0x29,
	0xc5, 0x0d, 0x67, 0x83, 0x54, 0x24, 0xb1, 0x81, 0xfd, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xe0, 0xc6, 0x35, 0x85, 0x0b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CacheClient is the client API for Cache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CacheClient interface {
	Store(ctx context.Context, in *StoreReq, opts ...grpc.CallOption) (*StoreResp, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
}

type cacheClient struct {
	cc *grpc.ClientConn
}

func NewCacheClient(cc *grpc.ClientConn) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) Store(ctx context.Context, in *StoreReq, opts ...grpc.CallOption) (*StoreResp, error) {
	out := new(StoreResp)
	err := c.cc.Invoke(ctx, "/v1.Cache/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, "/v1.Cache/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServer is the server API for Cache service.
type CacheServer interface {
	Store(context.Context, *StoreReq) (*StoreResp, error)
	Get(context.Context, *GetReq) (*GetResp, error)
}

// UnimplementedCacheServer can be embedded to have forward compatible implementations.
type UnimplementedCacheServer struct {
}

func (*UnimplementedCacheServer) Store(ctx context.Context, req *StoreReq) (*StoreResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedCacheServer) Get(ctx context.Context, req *GetReq) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterCacheServer(s *grpc.Server, srv CacheServer) {
	s.RegisterService(&_Cache_serviceDesc, srv)
}

func _Cache_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Cache/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Store(ctx, req.(*StoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Cache/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _Cache_Store_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Cache_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache-service.proto",
}
