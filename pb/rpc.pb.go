// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

/*
Package protocols_p2p is a generated protocol buffer package.

It is generated from these files:
	rpc.proto
	minimal.proto

It has these top-level messages:
	RPCAddPeerReq
	RPCSubscribeShardReq
	RPCUnsubscribeShardReq
	RPCGetSubscribedShardReq
	RPCGetSubscribedShardReply
	RPCBroadcastCollationReq
	RPCReply
	AddPeerRequest
	AddPeerResponse
	Collation
	CollationRequest
	CollationResponse
	NotifyShardsRequest
*/
package protocols_p2p

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AddPeer
type RPCAddPeerReq struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	Seed int64  `protobuf:"varint,3,opt,name=seed" json:"seed,omitempty"`
}

func (m *RPCAddPeerReq) Reset()                    { *m = RPCAddPeerReq{} }
func (m *RPCAddPeerReq) String() string            { return proto.CompactTextString(m) }
func (*RPCAddPeerReq) ProtoMessage()               {}
func (*RPCAddPeerReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RPCAddPeerReq) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *RPCAddPeerReq) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *RPCAddPeerReq) GetSeed() int64 {
	if m != nil {
		return m.Seed
	}
	return 0
}

// SubscribeShard
type RPCSubscribeShardReq struct {
	ShardIDs []int64 `protobuf:"varint,1,rep,packed,name=shardIDs" json:"shardIDs,omitempty"`
}

func (m *RPCSubscribeShardReq) Reset()                    { *m = RPCSubscribeShardReq{} }
func (m *RPCSubscribeShardReq) String() string            { return proto.CompactTextString(m) }
func (*RPCSubscribeShardReq) ProtoMessage()               {}
func (*RPCSubscribeShardReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RPCSubscribeShardReq) GetShardIDs() []int64 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

// UnsubscribeShard
type RPCUnsubscribeShardReq struct {
	ShardIDs []int64 `protobuf:"varint,1,rep,packed,name=shardIDs" json:"shardIDs,omitempty"`
}

func (m *RPCUnsubscribeShardReq) Reset()                    { *m = RPCUnsubscribeShardReq{} }
func (m *RPCUnsubscribeShardReq) String() string            { return proto.CompactTextString(m) }
func (*RPCUnsubscribeShardReq) ProtoMessage()               {}
func (*RPCUnsubscribeShardReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RPCUnsubscribeShardReq) GetShardIDs() []int64 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

// GetSubscribedShard
type RPCGetSubscribedShardReq struct {
}

func (m *RPCGetSubscribedShardReq) Reset()                    { *m = RPCGetSubscribedShardReq{} }
func (m *RPCGetSubscribedShardReq) String() string            { return proto.CompactTextString(m) }
func (*RPCGetSubscribedShardReq) ProtoMessage()               {}
func (*RPCGetSubscribedShardReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type RPCGetSubscribedShardReply struct {
	ShardIDs []int64 `protobuf:"varint,1,rep,packed,name=shardIDs" json:"shardIDs,omitempty"`
	Status   bool    `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
}

func (m *RPCGetSubscribedShardReply) Reset()                    { *m = RPCGetSubscribedShardReply{} }
func (m *RPCGetSubscribedShardReply) String() string            { return proto.CompactTextString(m) }
func (*RPCGetSubscribedShardReply) ProtoMessage()               {}
func (*RPCGetSubscribedShardReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RPCGetSubscribedShardReply) GetShardIDs() []int64 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

func (m *RPCGetSubscribedShardReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

// BroadcastCollation
type RPCBroadcastCollationReq struct {
	ShardID int64 `protobuf:"varint,1,opt,name=shardID" json:"shardID,omitempty"`
	Number  int32 `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
	Size    int32 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Period  int32 `protobuf:"varint,4,opt,name=period" json:"period,omitempty"`
}

func (m *RPCBroadcastCollationReq) Reset()                    { *m = RPCBroadcastCollationReq{} }
func (m *RPCBroadcastCollationReq) String() string            { return proto.CompactTextString(m) }
func (*RPCBroadcastCollationReq) ProtoMessage()               {}
func (*RPCBroadcastCollationReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RPCBroadcastCollationReq) GetShardID() int64 {
	if m != nil {
		return m.ShardID
	}
	return 0
}

func (m *RPCBroadcastCollationReq) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *RPCBroadcastCollationReq) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *RPCBroadcastCollationReq) GetPeriod() int32 {
	if m != nil {
		return m.Period
	}
	return 0
}

type RPCReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Status  bool   `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
}

func (m *RPCReply) Reset()                    { *m = RPCReply{} }
func (m *RPCReply) String() string            { return proto.CompactTextString(m) }
func (*RPCReply) ProtoMessage()               {}
func (*RPCReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RPCReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RPCReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*RPCAddPeerReq)(nil), "protocols.p2p.RPCAddPeerReq")
	proto.RegisterType((*RPCSubscribeShardReq)(nil), "protocols.p2p.RPCSubscribeShardReq")
	proto.RegisterType((*RPCUnsubscribeShardReq)(nil), "protocols.p2p.RPCUnsubscribeShardReq")
	proto.RegisterType((*RPCGetSubscribedShardReq)(nil), "protocols.p2p.RPCGetSubscribedShardReq")
	proto.RegisterType((*RPCGetSubscribedShardReply)(nil), "protocols.p2p.RPCGetSubscribedShardReply")
	proto.RegisterType((*RPCBroadcastCollationReq)(nil), "protocols.p2p.RPCBroadcastCollationReq")
	proto.RegisterType((*RPCReply)(nil), "protocols.p2p.RPCReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Poc service

type PocClient interface {
	// Sends a greeting
	AddPeer(ctx context.Context, in *RPCAddPeerReq, opts ...grpc.CallOption) (*RPCReply, error)
	SubscribeShard(ctx context.Context, in *RPCSubscribeShardReq, opts ...grpc.CallOption) (*RPCReply, error)
	UnsubscribeShard(ctx context.Context, in *RPCUnsubscribeShardReq, opts ...grpc.CallOption) (*RPCReply, error)
	GetSubscribedShard(ctx context.Context, in *RPCGetSubscribedShardReq, opts ...grpc.CallOption) (*RPCGetSubscribedShardReply, error)
	BroadcastCollation(ctx context.Context, in *RPCBroadcastCollationReq, opts ...grpc.CallOption) (*RPCReply, error)
}

type pocClient struct {
	cc *grpc.ClientConn
}

func NewPocClient(cc *grpc.ClientConn) PocClient {
	return &pocClient{cc}
}

func (c *pocClient) AddPeer(ctx context.Context, in *RPCAddPeerReq, opts ...grpc.CallOption) (*RPCReply, error) {
	out := new(RPCReply)
	err := grpc.Invoke(ctx, "/protocols.p2p.Poc/AddPeer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) SubscribeShard(ctx context.Context, in *RPCSubscribeShardReq, opts ...grpc.CallOption) (*RPCReply, error) {
	out := new(RPCReply)
	err := grpc.Invoke(ctx, "/protocols.p2p.Poc/SubscribeShard", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) UnsubscribeShard(ctx context.Context, in *RPCUnsubscribeShardReq, opts ...grpc.CallOption) (*RPCReply, error) {
	out := new(RPCReply)
	err := grpc.Invoke(ctx, "/protocols.p2p.Poc/UnsubscribeShard", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) GetSubscribedShard(ctx context.Context, in *RPCGetSubscribedShardReq, opts ...grpc.CallOption) (*RPCGetSubscribedShardReply, error) {
	out := new(RPCGetSubscribedShardReply)
	err := grpc.Invoke(ctx, "/protocols.p2p.Poc/GetSubscribedShard", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pocClient) BroadcastCollation(ctx context.Context, in *RPCBroadcastCollationReq, opts ...grpc.CallOption) (*RPCReply, error) {
	out := new(RPCReply)
	err := grpc.Invoke(ctx, "/protocols.p2p.Poc/BroadcastCollation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Poc service

type PocServer interface {
	// Sends a greeting
	AddPeer(context.Context, *RPCAddPeerReq) (*RPCReply, error)
	SubscribeShard(context.Context, *RPCSubscribeShardReq) (*RPCReply, error)
	UnsubscribeShard(context.Context, *RPCUnsubscribeShardReq) (*RPCReply, error)
	GetSubscribedShard(context.Context, *RPCGetSubscribedShardReq) (*RPCGetSubscribedShardReply, error)
	BroadcastCollation(context.Context, *RPCBroadcastCollationReq) (*RPCReply, error)
}

func RegisterPocServer(s *grpc.Server, srv PocServer) {
	s.RegisterService(&_Poc_serviceDesc, srv)
}

func _Poc_AddPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCAddPeerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).AddPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocols.p2p.Poc/AddPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).AddPeer(ctx, req.(*RPCAddPeerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_SubscribeShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCSubscribeShardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).SubscribeShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocols.p2p.Poc/SubscribeShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).SubscribeShard(ctx, req.(*RPCSubscribeShardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_UnsubscribeShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCUnsubscribeShardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).UnsubscribeShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocols.p2p.Poc/UnsubscribeShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).UnsubscribeShard(ctx, req.(*RPCUnsubscribeShardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_GetSubscribedShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCGetSubscribedShardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).GetSubscribedShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocols.p2p.Poc/GetSubscribedShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).GetSubscribedShard(ctx, req.(*RPCGetSubscribedShardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poc_BroadcastCollation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCBroadcastCollationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocServer).BroadcastCollation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocols.p2p.Poc/BroadcastCollation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocServer).BroadcastCollation(ctx, req.(*RPCBroadcastCollationReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Poc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocols.p2p.Poc",
	HandlerType: (*PocServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPeer",
			Handler:    _Poc_AddPeer_Handler,
		},
		{
			MethodName: "SubscribeShard",
			Handler:    _Poc_SubscribeShard_Handler,
		},
		{
			MethodName: "UnsubscribeShard",
			Handler:    _Poc_UnsubscribeShard_Handler,
		},
		{
			MethodName: "GetSubscribedShard",
			Handler:    _Poc_GetSubscribedShard_Handler,
		},
		{
			MethodName: "BroadcastCollation",
			Handler:    _Poc_BroadcastCollation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x35, 0x4d, 0x3f, 0x07, 0x5a, 0x64, 0x91, 0xba, 0x04, 0x0f, 0x21, 0x22, 0xc6, 0x4b, 0x0e,
	0xd5, 0xa3, 0x17, 0x1b, 0xa1, 0x78, 0x5b, 0xb6, 0x0a, 0x5e, 0xf3, 0xb1, 0x68, 0x24, 0xed, 0xae,
	0xbb, 0x29, 0xa8, 0x7f, 0xcc, 0xbf, 0x27, 0xd9, 0xa4, 0x85, 0xba, 0x69, 0xd0, 0x53, 0xe6, 0x65,
	0x66, 0xde, 0x9b, 0x9d, 0x37, 0x30, 0x92, 0x22, 0x09, 0x84, 0xe4, 0x05, 0x47, 0x63, 0xfd, 0x49,
	0x78, 0xae, 0x02, 0x31, 0x13, 0xde, 0x02, 0xc6, 0x94, 0x84, 0x77, 0x69, 0x4a, 0x18, 0x93, 0x94,
	0xbd, 0xa3, 0x09, 0x74, 0x32, 0x81, 0x2d, 0xd7, 0xf2, 0x47, 0xb4, 0x93, 0x09, 0x84, 0xa0, 0x2b,
	0xb8, 0x2c, 0x70, 0xc7, 0xb5, 0xfc, 0x1e, 0xd5, 0x71, 0xf9, 0x4f, 0x31, 0x96, 0x62, 0xdb, 0xb5,
	0x7c, 0x9b, 0xea, 0xd8, 0x9b, 0xc1, 0x09, 0x25, 0xe1, 0x72, 0x13, 0xab, 0x44, 0x66, 0x31, 0x5b,
	0xbe, 0x46, 0x32, 0x2d, 0xf9, 0x1c, 0x18, 0xaa, 0x32, 0x7e, 0xb8, 0x57, 0xd8, 0x72, 0x6d, 0xdf,
	0xa6, 0x3b, 0xec, 0xdd, 0xc0, 0x94, 0x92, 0xf0, 0x69, 0xad, 0xfe, 0xd5, 0xe5, 0x00, 0xa6, 0x24,
	0x5c, 0xb0, 0x62, 0x27, 0x96, 0x6e, 0xfb, 0x3c, 0x02, 0xce, 0x81, 0x9c, 0xc8, 0x3f, 0xdb, 0x58,
	0xd1, 0x14, 0xfa, 0xaa, 0x88, 0x8a, 0x8d, 0xd2, 0x2f, 0x1d, 0xd2, 0x1a, 0x79, 0x1f, 0x5a, 0x6d,
	0x2e, 0x79, 0x94, 0x26, 0x91, 0x2a, 0x42, 0x9e, 0xe7, 0x51, 0x91, 0xf1, 0x75, 0x39, 0x25, 0x86,
	0x41, 0xdd, 0xaf, 0x17, 0x66, 0xd3, 0x2d, 0x2c, 0xd9, 0xd6, 0x9b, 0x55, 0xcc, 0x64, 0xbd, 0xb7,
	0x1a, 0xe9, 0xcd, 0x65, 0x5f, 0x4c, 0x6f, 0xae, 0x47, 0x75, 0x5c, 0xd6, 0x0a, 0x26, 0x33, 0x9e,
	0xe2, 0x6e, 0x55, 0x5b, 0x21, 0xef, 0x16, 0x86, 0x94, 0x84, 0xd5, 0xe4, 0x18, 0x06, 0x2b, 0xa6,
	0x54, 0xf4, 0xc2, 0x6a, 0x6b, 0xb6, 0xf0, 0xd0, 0xdc, 0xb3, 0x6f, 0x1b, 0x6c, 0xc2, 0x13, 0x34,
	0x87, 0x41, 0xed, 0x2e, 0x3a, 0x0b, 0xf6, 0xbc, 0x0f, 0xf6, 0x8c, 0x77, 0x4e, 0xcd, 0xac, 0xd6,
	0xf6, 0x8e, 0x10, 0x81, 0xc9, 0xbe, 0xb1, 0xe8, 0xdc, 0x2c, 0x36, 0xac, 0x6f, 0x63, 0x7c, 0x84,
	0xe3, 0xdf, 0xb6, 0xa3, 0x0b, 0xb3, 0xbc, 0xe1, 0x34, 0xda, 0x58, 0xdf, 0x00, 0x99, 0xd6, 0xa3,
	0x4b, 0xb3, 0xa1, 0xf1, 0x78, 0x9c, 0xab, 0xbf, 0x15, 0x56, 0x5a, 0xcf, 0x80, 0xcc, 0xa3, 0x68,
	0xd2, 0x6a, 0x3c, 0x9d, 0x96, 0x57, 0xc4, 0x7d, 0x9d, 0xb9, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xb5, 0x15, 0x52, 0x29, 0xb5, 0x03, 0x00, 0x00,
}
