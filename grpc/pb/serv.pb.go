// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serv.proto

package serv

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

type Fileinfo struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Cid                  string   `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fileinfo) Reset()         { *m = Fileinfo{} }
func (m *Fileinfo) String() string { return proto.CompactTextString(m) }
func (*Fileinfo) ProtoMessage()    {}
func (*Fileinfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_20d2027059fc23fa, []int{0}
}

func (m *Fileinfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fileinfo.Unmarshal(m, b)
}
func (m *Fileinfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fileinfo.Marshal(b, m, deterministic)
}
func (m *Fileinfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fileinfo.Merge(m, src)
}
func (m *Fileinfo) XXX_Size() int {
	return xxx_messageInfo_Fileinfo.Size(m)
}
func (m *Fileinfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Fileinfo.DiscardUnknown(m)
}

var xxx_messageInfo_Fileinfo proto.InternalMessageInfo

func (m *Fileinfo) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Fileinfo) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

type Id struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_20d2027059fc23fa, []int{1}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Peer struct {
	Cid                  string   `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_20d2027059fc23fa, []int{2}
}

func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func init() {
	proto.RegisterType((*Fileinfo)(nil), "Fileinfo")
	proto.RegisterType((*Id)(nil), "Id")
	proto.RegisterType((*Peer)(nil), "peer")
}

func init() { proto.RegisterFile("serv.proto", fileDescriptor_20d2027059fc23fa) }

var fileDescriptor_20d2027059fc23fa = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x4e, 0x2d, 0x2a,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2, 0xe0, 0xe2, 0x70, 0xcb, 0xcc, 0x49, 0xcd, 0xcc,
	0x4b, 0xcb, 0x17, 0x92, 0xe2, 0xe2, 0x48, 0xcb, 0xcc, 0x49, 0xcd, 0x4b, 0xcc, 0x4d, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85, 0x04, 0xb8, 0x98, 0x93, 0x33, 0x53, 0x24, 0x98,
	0xc0, 0xc2, 0x20, 0xa6, 0x92, 0x08, 0x17, 0x93, 0x67, 0x8a, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x0a,
	0x54, 0x35, 0x53, 0x66, 0x8a, 0x92, 0x04, 0x17, 0x4b, 0x41, 0x6a, 0x6a, 0x11, 0x4c, 0x3d, 0x23,
	0x5c, 0xbd, 0x91, 0x3d, 0x17, 0x57, 0x62, 0x4a, 0x8a, 0x73, 0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x89,
	0x90, 0x24, 0x17, 0xbb, 0x63, 0x4a, 0x0a, 0xc8, 0x6a, 0x21, 0x4e, 0x3d, 0x98, 0x0b, 0xa4, 0x98,
	0xf5, 0x3c, 0x53, 0x84, 0xc4, 0x40, 0x4e, 0xca, 0x4b, 0x09, 0x00, 0x19, 0xc3, 0xaa, 0x07, 0x32,
	0x0d, 0x2c, 0x9e, 0xc4, 0x06, 0x76, 0xb1, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x36, 0x03,
	0xf3, 0xbf, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddContentClient is the client API for AddContent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddContentClient interface {
	AddFile(ctx context.Context, in *Fileinfo, opts ...grpc.CallOption) (*Id, error)
	FindPeer(ctx context.Context, in *Peer, opts ...grpc.CallOption) (*Id, error)
}

type addContentClient struct {
	cc *grpc.ClientConn
}

func NewAddContentClient(cc *grpc.ClientConn) AddContentClient {
	return &addContentClient{cc}
}

func (c *addContentClient) AddFile(ctx context.Context, in *Fileinfo, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/addContent/AddFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addContentClient) FindPeer(ctx context.Context, in *Peer, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/addContent/FindPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddContentServer is the server API for AddContent service.
type AddContentServer interface {
	AddFile(context.Context, *Fileinfo) (*Id, error)
	FindPeer(context.Context, *Peer) (*Id, error)
}

// UnimplementedAddContentServer can be embedded to have forward compatible implementations.
type UnimplementedAddContentServer struct {
}

func (*UnimplementedAddContentServer) AddFile(ctx context.Context, req *Fileinfo) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFile not implemented")
}
func (*UnimplementedAddContentServer) FindPeer(ctx context.Context, req *Peer) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPeer not implemented")
}

func RegisterAddContentServer(s *grpc.Server, srv AddContentServer) {
	s.RegisterService(&_AddContent_serviceDesc, srv)
}

func _AddContent_AddFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Fileinfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddContentServer).AddFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addContent/AddFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddContentServer).AddFile(ctx, req.(*Fileinfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddContent_FindPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Peer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddContentServer).FindPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addContent/FindPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddContentServer).FindPeer(ctx, req.(*Peer))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddContent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "addContent",
	HandlerType: (*AddContentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFile",
			Handler:    _AddContent_AddFile_Handler,
		},
		{
			MethodName: "FindPeer",
			Handler:    _AddContent_FindPeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serv.proto",
}
