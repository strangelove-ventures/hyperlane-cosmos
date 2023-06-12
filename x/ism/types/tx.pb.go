// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hyperlane/ism/v1/tx.proto

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

// MsgSetDefaultIsm defines the request type for the SetDefaultIsm rpc.
type MsgSetDefaultIsm struct {
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Isms   []*Ism `protobuf:"bytes,2,rep,name=isms,proto3" json:"isms,omitempty"`
}

func (m *MsgSetDefaultIsm) Reset()         { *m = MsgSetDefaultIsm{} }
func (m *MsgSetDefaultIsm) String() string { return proto.CompactTextString(m) }
func (*MsgSetDefaultIsm) ProtoMessage()    {}
func (*MsgSetDefaultIsm) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b013040feeda308, []int{0}
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

func (m *MsgSetDefaultIsm) GetIsms() []*Ism {
	if m != nil {
		return m.Isms
	}
	return nil
}

// MsgSetDefaultIsmResponse defines the Msg/SetDefaultIsm response type
type MsgSetDefaultIsmResponse struct{}

func (m *MsgSetDefaultIsmResponse) Reset()         { *m = MsgSetDefaultIsmResponse{} }
func (m *MsgSetDefaultIsmResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetDefaultIsmResponse) ProtoMessage()    {}
func (*MsgSetDefaultIsmResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b013040feeda308, []int{1}
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
	proto.RegisterType((*MsgSetDefaultIsm)(nil), "hyperlane.ism.v1.MsgSetDefaultIsm")
	proto.RegisterType((*MsgSetDefaultIsmResponse)(nil), "hyperlane.ism.v1.MsgSetDefaultIsmResponse")
}

func init() { proto.RegisterFile("hyperlane/ism/v1/tx.proto", fileDescriptor_7b013040feeda308) }

var fileDescriptor_7b013040feeda308 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xa8, 0x2c, 0x48,
	0x2d, 0xca, 0x49, 0xcc, 0x4b, 0xd5, 0xcf, 0x2c, 0xce, 0xd5, 0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0x4b, 0xe9, 0x65, 0x16, 0xe7, 0xea, 0x95, 0x19,
	0x4a, 0x49, 0x61, 0x28, 0x06, 0x49, 0x80, 0x55, 0x2b, 0x85, 0x72, 0x09, 0xf8, 0x16, 0xa7, 0x07,
	0xa7, 0x96, 0xb8, 0xa4, 0xa6, 0x25, 0x96, 0xe6, 0x94, 0x78, 0x16, 0xe7, 0x0a, 0x89, 0x71, 0xb1,
	0x15, 0x67, 0xa6, 0xe7, 0xa5, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x42,
	0x9a, 0x5c, 0x2c, 0x99, 0xc5, 0xb9, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xa2, 0x7a,
	0xe8, 0x16, 0xe9, 0x79, 0x16, 0xe7, 0x06, 0x81, 0x95, 0x28, 0x49, 0x71, 0x49, 0xa0, 0x1b, 0x1b,
	0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x6a, 0x94, 0xc6, 0xc5, 0xec, 0x5b, 0x9c, 0x2e, 0x14,
	0xcf, 0xc5, 0x8b, 0x6a, 0xad, 0x12, 0xa6, 0x81, 0xe8, 0x66, 0x48, 0x69, 0x11, 0x56, 0x03, 0xb3,
	0xc7, 0x29, 0xec, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c,
	0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x6c, 0xd2, 0x33,
	0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x8b, 0x4b, 0x8a, 0x12, 0xf3, 0xd2, 0x53,
	0x73, 0xf2, 0xcb, 0x52, 0x75, 0xcb, 0x52, 0xf3, 0x4a, 0x4a, 0x8b, 0x52, 0x8b, 0xf5, 0xe1, 0x96,
	0xe8, 0x26, 0xe7, 0x17, 0xe7, 0xe6, 0x17, 0xeb, 0x57, 0x80, 0x43, 0xae, 0xa4, 0xb2, 0x20, 0xb5,
	0x38, 0x89, 0x0d, 0x1c, 0x72, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0xd6, 0x66, 0x48,
	0x84, 0x01, 0x00, 0x00,
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
	// SetDefaultIsm defines a rpc handler method for MsgSetDefaultIsm.
	SetDefaultIsm(ctx context.Context, in *MsgSetDefaultIsm, opts ...grpc.CallOption) (*MsgSetDefaultIsmResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetDefaultIsm(ctx context.Context, in *MsgSetDefaultIsm, opts ...grpc.CallOption) (*MsgSetDefaultIsmResponse, error) {
	out := new(MsgSetDefaultIsmResponse)
	err := c.cc.Invoke(ctx, "/hyperlane.ism.v1.Msg/SetDefaultIsm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SetDefaultIsm defines a rpc handler method for MsgSetDefaultIsm.
	SetDefaultIsm(context.Context, *MsgSetDefaultIsm) (*MsgSetDefaultIsmResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct{}

func (*UnimplementedMsgServer) SetDefaultIsm(ctx context.Context, req *MsgSetDefaultIsm) (*MsgSetDefaultIsmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefaultIsm not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
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
		FullMethod: "/hyperlane.ism.v1.Msg/SetDefaultIsm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetDefaultIsm(ctx, req.(*MsgSetDefaultIsm))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hyperlane.ism.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetDefaultIsm",
			Handler:    _Msg_SetDefaultIsm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hyperlane/ism/v1/tx.proto",
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
	if len(m.Isms) > 0 {
		for iNdEx := len(m.Isms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Isms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
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
	if len(m.Isms) > 0 {
		for _, e := range m.Isms {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Isms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Isms = append(m.Isms, &Ism{})
			if err := m.Isms[len(m.Isms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
