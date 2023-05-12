// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hyperlane/mailbox/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgDispatch defines the request type for the Dispatch rpc.
type MsgDispatch struct {
	DestinationDomain uint32 `protobuf:"varint,1,opt,name=destination_domain,json=destinationDomain,proto3" json:"destination_domain,omitempty"`
	RecpientAddress   string `protobuf:"bytes,2,opt,name=recpient_address,json=recpientAddress,proto3" json:"recpient_address,omitempty"`
	MessageBody       string `protobuf:"bytes,3,opt,name=message_body,json=messageBody,proto3" json:"message_body,omitempty"`
}

func (m *MsgDispatch) Reset()         { *m = MsgDispatch{} }
func (m *MsgDispatch) String() string { return proto.CompactTextString(m) }
func (*MsgDispatch) ProtoMessage()    {}
func (*MsgDispatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_544026924916c03f, []int{0}
}

func (m *MsgDispatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgDispatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDispatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgDispatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDispatch.Merge(m, src)
}

func (m *MsgDispatch) XXX_Size() int {
	return m.Size()
}

func (m *MsgDispatch) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDispatch.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDispatch proto.InternalMessageInfo

func (m *MsgDispatch) GetDestinationDomain() uint32 {
	if m != nil {
		return m.DestinationDomain
	}
	return 0
}

func (m *MsgDispatch) GetRecpientAddress() string {
	if m != nil {
		return m.RecpientAddress
	}
	return ""
}

func (m *MsgDispatch) GetMessageBody() string {
	if m != nil {
		return m.MessageBody
	}
	return ""
}

// MsgDispatchResponse defines the Dispatch response type.
type MsgDispatchResponse struct {
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (m *MsgDispatchResponse) Reset()         { *m = MsgDispatchResponse{} }
func (m *MsgDispatchResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDispatchResponse) ProtoMessage()    {}
func (*MsgDispatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_544026924916c03f, []int{1}
}

func (m *MsgDispatchResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgDispatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDispatchResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgDispatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDispatchResponse.Merge(m, src)
}

func (m *MsgDispatchResponse) XXX_Size() int {
	return m.Size()
}

func (m *MsgDispatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDispatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDispatchResponse proto.InternalMessageInfo

func (m *MsgDispatchResponse) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

// MsgProcess defines the request type for the Process rpc.
type MsgProcess struct {
	Metadata string `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *MsgProcess) Reset()         { *m = MsgProcess{} }
func (m *MsgProcess) String() string { return proto.CompactTextString(m) }
func (*MsgProcess) ProtoMessage()    {}
func (*MsgProcess) Descriptor() ([]byte, []int) {
	return fileDescriptor_544026924916c03f, []int{2}
}

func (m *MsgProcess) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgProcess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProcess.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgProcess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProcess.Merge(m, src)
}

func (m *MsgProcess) XXX_Size() int {
	return m.Size()
}

func (m *MsgProcess) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProcess.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProcess proto.InternalMessageInfo

func (m *MsgProcess) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *MsgProcess) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// MsgProcessResponse defines the Process response type.
type MsgProcessResponse struct{}

func (m *MsgProcessResponse) Reset()         { *m = MsgProcessResponse{} }
func (m *MsgProcessResponse) String() string { return proto.CompactTextString(m) }
func (*MsgProcessResponse) ProtoMessage()    {}
func (*MsgProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_544026924916c03f, []int{3}
}

func (m *MsgProcessResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProcessResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProcessResponse.Merge(m, src)
}

func (m *MsgProcessResponse) XXX_Size() int {
	return m.Size()
}

func (m *MsgProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProcessResponse proto.InternalMessageInfo

// MsgSetDefaultIsm defines the request type for the SetDefaultIsm rpc.
type MsgSetDefaultIsm struct {
	Signer  string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	IsmHash string `protobuf:"bytes,2,opt,name=ism_hash,json=ismHash,proto3" json:"ism_hash,omitempty"`
}

func (m *MsgSetDefaultIsm) Reset()         { *m = MsgSetDefaultIsm{} }
func (m *MsgSetDefaultIsm) String() string { return proto.CompactTextString(m) }
func (*MsgSetDefaultIsm) ProtoMessage()    {}
func (*MsgSetDefaultIsm) Descriptor() ([]byte, []int) {
	return fileDescriptor_544026924916c03f, []int{4}
}

func (m *MsgSetDefaultIsm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgSetDefaultIsm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetDefaultIsm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgSetDefaultIsm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetDefaultIsm.Merge(m, src)
}

func (m *MsgSetDefaultIsm) XXX_Size() int {
	return m.Size()
}

func (m *MsgSetDefaultIsm) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetDefaultIsm.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetDefaultIsm proto.InternalMessageInfo

func (m *MsgSetDefaultIsm) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgSetDefaultIsm) GetIsmHash() string {
	if m != nil {
		return m.IsmHash
	}
	return ""
}

// MsgSetDefaultIsmResponse defines the Msg/SetDefaultIsm response type
type MsgSetDefaultIsmResponse struct{}

func (m *MsgSetDefaultIsmResponse) Reset()         { *m = MsgSetDefaultIsmResponse{} }
func (m *MsgSetDefaultIsmResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetDefaultIsmResponse) ProtoMessage()    {}
func (*MsgSetDefaultIsmResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_544026924916c03f, []int{5}
}

func (m *MsgSetDefaultIsmResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgSetDefaultIsmResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetDefaultIsmResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgSetDefaultIsmResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetDefaultIsmResponse.Merge(m, src)
}

func (m *MsgSetDefaultIsmResponse) XXX_Size() int {
	return m.Size()
}

func (m *MsgSetDefaultIsmResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetDefaultIsmResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetDefaultIsmResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgDispatch)(nil), "hyperlane.mailbox.v1.MsgDispatch")
	proto.RegisterType((*MsgDispatchResponse)(nil), "hyperlane.mailbox.v1.MsgDispatchResponse")
	proto.RegisterType((*MsgProcess)(nil), "hyperlane.mailbox.v1.MsgProcess")
	proto.RegisterType((*MsgProcessResponse)(nil), "hyperlane.mailbox.v1.MsgProcessResponse")
	proto.RegisterType((*MsgSetDefaultIsm)(nil), "hyperlane.mailbox.v1.MsgSetDefaultIsm")
	proto.RegisterType((*MsgSetDefaultIsmResponse)(nil), "hyperlane.mailbox.v1.MsgSetDefaultIsmResponse")
}

func init() { proto.RegisterFile("hyperlane/mailbox/v1/tx.proto", fileDescriptor_544026924916c03f) }

var fileDescriptor_544026924916c03f = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0xcf, 0xb6, 0xd0, 0x26, 0xaf, 0x16, 0xeb, 0x58, 0x64, 0x5d, 0xe8, 0x92, 0xee, 0x41, 0xd2,
	0x43, 0x76, 0xa9, 0xfa, 0x01, 0x34, 0x44, 0xb0, 0x87, 0x80, 0x44, 0x04, 0xe9, 0x25, 0x4c, 0x76,
	0x9e, 0xb3, 0x03, 0x99, 0x99, 0x65, 0xdf, 0x24, 0x24, 0x5f, 0xc1, 0x93, 0xdf, 0xc0, 0xaf, 0xe3,
	0xb1, 0x47, 0x8f, 0x92, 0x7c, 0x11, 0x31, 0xce, 0xae, 0xa9, 0x58, 0xea, 0xf1, 0xbd, 0xdf, 0xbf,
	0xc7, 0x6f, 0x18, 0x38, 0x2b, 0x56, 0x25, 0x56, 0x33, 0x6e, 0x30, 0xd3, 0x5c, 0xcd, 0xa6, 0x76,
	0x99, 0x2d, 0x2e, 0x33, 0xb7, 0x4c, 0xcb, 0xca, 0x3a, 0xcb, 0x4e, 0x1b, 0x38, 0xf5, 0x70, 0xba,
	0xb8, 0x4c, 0x3e, 0x07, 0x70, 0x34, 0x22, 0x39, 0x54, 0x54, 0x72, 0x97, 0x17, 0xac, 0x0f, 0x4c,
	0x20, 0x39, 0x65, 0xb8, 0x53, 0xd6, 0x4c, 0x84, 0xd5, 0x5c, 0x99, 0x30, 0xe8, 0x06, 0xbd, 0xe3,
	0xf1, 0xa3, 0x1d, 0x64, 0xb8, 0x05, 0xd8, 0x05, 0x9c, 0x54, 0x98, 0x97, 0x0a, 0x8d, 0x9b, 0x70,
	0x21, 0x2a, 0x24, 0x0a, 0xf7, 0xba, 0x41, 0xaf, 0x33, 0x7e, 0x58, 0xef, 0x5f, 0xff, 0x5e, 0xb3,
	0x73, 0x78, 0xa0, 0x91, 0x88, 0x4b, 0x9c, 0x4c, 0xad, 0x58, 0x85, 0xfb, 0x5b, 0xda, 0x91, 0xdf,
	0x0d, 0xac, 0x58, 0x25, 0x2f, 0xe1, 0xf1, 0xce, 0x2d, 0x63, 0xa4, 0xd2, 0x1a, 0x42, 0x76, 0x06,
	0x50, 0x2b, 0x95, 0xd8, 0xde, 0xd2, 0x19, 0x77, 0xfc, 0xe6, 0x4a, 0x24, 0x03, 0x80, 0x11, 0xc9,
	0x77, 0x95, 0xcd, 0x7f, 0xc5, 0x44, 0xd0, 0xd6, 0xe8, 0xb8, 0xe0, 0x8e, 0x7b, 0x6a, 0x33, 0xb3,
	0x10, 0x0e, 0xbd, 0xcc, 0x1f, 0x59, 0x8f, 0xc9, 0x29, 0xb0, 0x3f, 0x1e, 0x75, 0x70, 0xf2, 0x06,
	0x4e, 0x46, 0x24, 0xdf, 0xa3, 0x1b, 0xe2, 0x27, 0x3e, 0x9f, 0xb9, 0x2b, 0xd2, 0xec, 0x09, 0x1c,
	0x90, 0x92, 0x06, 0x2b, 0xef, 0xee, 0x27, 0xf6, 0x14, 0xda, 0x8a, 0xf4, 0xa4, 0xe0, 0x54, 0xd4,
	0xe6, 0x8a, 0xf4, 0x5b, 0x4e, 0x45, 0x12, 0x41, 0xf8, 0xb7, 0x4d, 0x1d, 0xf1, 0xfc, 0xeb, 0x1e,
	0xec, 0x8f, 0x48, 0xb2, 0x8f, 0xd0, 0x6e, 0xde, 0xe0, 0x3c, 0xfd, 0xd7, 0x53, 0xa5, 0x3b, 0xd5,
	0x44, 0x17, 0xf7, 0x52, 0x9a, 0xf6, 0x3e, 0xc0, 0x61, 0xdd, 0x4d, 0xf7, 0x4e, 0x95, 0x67, 0x44,
	0xbd, 0xfb, 0x18, 0x8d, 0xad, 0x84, 0xe3, 0xdb, 0xc5, 0x3c, 0xbb, 0x53, 0x7a, 0x8b, 0x17, 0xa5,
	0xff, 0xc7, 0xab, 0x83, 0x06, 0xd7, 0xdf, 0xd6, 0x71, 0x70, 0xb3, 0x8e, 0x83, 0x1f, 0xeb, 0x38,
	0xf8, 0xb2, 0x89, 0x5b, 0x37, 0x9b, 0xb8, 0xf5, 0x7d, 0x13, 0xb7, 0xae, 0x5f, 0x49, 0xe5, 0x8a,
	0xf9, 0x34, 0xcd, 0xad, 0xce, 0xc8, 0x55, 0xdc, 0x48, 0x9c, 0xd9, 0x05, 0xf6, 0x17, 0x68, 0xdc,
	0xbc, 0x42, 0xca, 0x9a, 0xa0, 0x7e, 0x6e, 0x49, 0x5b, 0xca, 0x96, 0xcd, 0xcf, 0x70, 0xab, 0x12,
	0x69, 0x7a, 0xb0, 0xfd, 0x1a, 0x2f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xfe, 0x34, 0x92,
	0x3b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Dispatch sends interchain messages
	Dispatch(ctx context.Context, in *MsgDispatch, opts ...grpc.CallOption) (*MsgDispatchResponse, error)
	// Process delivers interchain messages
	Process(ctx context.Context, in *MsgProcess, opts ...grpc.CallOption) (*MsgProcessResponse, error)
	// SetDefaultIsm defines a rpc handler method for MsgSetDefaultIsm.
	SetDefaultIsm(ctx context.Context, in *MsgSetDefaultIsm, opts ...grpc.CallOption) (*MsgSetDefaultIsmResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Dispatch(ctx context.Context, in *MsgDispatch, opts ...grpc.CallOption) (*MsgDispatchResponse, error) {
	out := new(MsgDispatchResponse)
	err := c.cc.Invoke(ctx, "/hyperlane.mailbox.v1.Msg/Dispatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Process(ctx context.Context, in *MsgProcess, opts ...grpc.CallOption) (*MsgProcessResponse, error) {
	out := new(MsgProcessResponse)
	err := c.cc.Invoke(ctx, "/hyperlane.mailbox.v1.Msg/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetDefaultIsm(ctx context.Context, in *MsgSetDefaultIsm, opts ...grpc.CallOption) (*MsgSetDefaultIsmResponse, error) {
	out := new(MsgSetDefaultIsmResponse)
	err := c.cc.Invoke(ctx, "/hyperlane.mailbox.v1.Msg/SetDefaultIsm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Dispatch sends interchain messages
	Dispatch(context.Context, *MsgDispatch) (*MsgDispatchResponse, error)
	// Process delivers interchain messages
	Process(context.Context, *MsgProcess) (*MsgProcessResponse, error)
	// SetDefaultIsm defines a rpc handler method for MsgSetDefaultIsm.
	SetDefaultIsm(context.Context, *MsgSetDefaultIsm) (*MsgSetDefaultIsmResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct{}

func (*UnimplementedMsgServer) Dispatch(ctx context.Context, req *MsgDispatch) (*MsgDispatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dispatch not implemented")
}

func (*UnimplementedMsgServer) Process(ctx context.Context, req *MsgProcess) (*MsgProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}

func (*UnimplementedMsgServer) SetDefaultIsm(ctx context.Context, req *MsgSetDefaultIsm) (*MsgSetDefaultIsmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefaultIsm not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Dispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDispatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Dispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hyperlane.mailbox.v1.Msg/Dispatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Dispatch(ctx, req.(*MsgDispatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgProcess)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hyperlane.mailbox.v1.Msg/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Process(ctx, req.(*MsgProcess))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetDefaultIsm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetDefaultIsm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetDefaultIsm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hyperlane.mailbox.v1.Msg/SetDefaultIsm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetDefaultIsm(ctx, req.(*MsgSetDefaultIsm))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hyperlane.mailbox.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dispatch",
			Handler:    _Msg_Dispatch_Handler,
		},
		{
			MethodName: "Process",
			Handler:    _Msg_Process_Handler,
		},
		{
			MethodName: "SetDefaultIsm",
			Handler:    _Msg_SetDefaultIsm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hyperlane/mailbox/v1/tx.proto",
}

func (m *MsgDispatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDispatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDispatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MessageBody) > 0 {
		i -= len(m.MessageBody)
		copy(dAtA[i:], m.MessageBody)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MessageBody)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RecpientAddress) > 0 {
		i -= len(m.RecpientAddress)
		copy(dAtA[i:], m.RecpientAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RecpientAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.DestinationDomain != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.DestinationDomain))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgDispatchResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDispatchResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDispatchResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MessageId) > 0 {
		i -= len(m.MessageId)
		copy(dAtA[i:], m.MessageId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MessageId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgProcess) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProcess) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProcess) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgProcessResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProcessResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProcessResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSetDefaultIsm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetDefaultIsm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetDefaultIsm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IsmHash) > 0 {
		i -= len(m.IsmHash)
		copy(dAtA[i:], m.IsmHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.IsmHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetDefaultIsmResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetDefaultIsmResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetDefaultIsmResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *MsgDispatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DestinationDomain != 0 {
		n += 1 + sovTx(uint64(m.DestinationDomain))
	}
	l = len(m.RecpientAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MessageBody)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDispatchResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MessageId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgProcess) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgProcessResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSetDefaultIsm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.IsmHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSetDefaultIsmResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *MsgDispatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDispatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDispatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationDomain", wireType)
			}
			m.DestinationDomain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DestinationDomain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecpientAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecpientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageBody", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageBody = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *MsgDispatchResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDispatchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDispatchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *MsgProcess) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgProcess: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProcess: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *MsgProcessResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgProcessResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProcessResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *MsgSetDefaultIsm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetDefaultIsm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetDefaultIsm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsmHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IsmHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *MsgSetDefaultIsmResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetDefaultIsmResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetDefaultIsmResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
