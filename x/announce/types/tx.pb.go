// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hyperlane/announce/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// MsgAnnouncement Announces a validator signature storage location
type MsgAnnouncement struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// The validator (in eth address format) that is announcing its storage
	// location
	Validator []byte `protobuf:"bytes,2,opt,name=validator,proto3" json:"validator,omitempty"`
	// location where signatures will be stored
	StorageLocation string `protobuf:"bytes,3,opt,name=storage_location,json=storageLocation,proto3" json:"storage_location,omitempty"`
	// signed validator announcement
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *MsgAnnouncement) Reset()         { *m = MsgAnnouncement{} }
func (m *MsgAnnouncement) String() string { return proto.CompactTextString(m) }
func (*MsgAnnouncement) ProtoMessage()    {}
func (*MsgAnnouncement) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d512fd238195273, []int{0}
}

func (m *MsgAnnouncement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgAnnouncement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAnnouncement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgAnnouncement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAnnouncement.Merge(m, src)
}

func (m *MsgAnnouncement) XXX_Size() int {
	return m.Size()
}

func (m *MsgAnnouncement) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAnnouncement.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAnnouncement proto.InternalMessageInfo

// MsgAnnouncementResponse defines the MsgAnnouncementResponse response type.
type MsgAnnouncementResponse struct{}

func (m *MsgAnnouncementResponse) Reset()         { *m = MsgAnnouncementResponse{} }
func (m *MsgAnnouncementResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAnnouncementResponse) ProtoMessage()    {}
func (*MsgAnnouncementResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d512fd238195273, []int{1}
}

func (m *MsgAnnouncementResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgAnnouncementResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAnnouncementResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgAnnouncementResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAnnouncementResponse.Merge(m, src)
}

func (m *MsgAnnouncementResponse) XXX_Size() int {
	return m.Size()
}

func (m *MsgAnnouncementResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAnnouncementResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAnnouncementResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAnnouncement)(nil), "hyperlane.announce.v1.MsgAnnouncement")
	proto.RegisterType((*MsgAnnouncementResponse)(nil), "hyperlane.announce.v1.MsgAnnouncementResponse")
}

func init() { proto.RegisterFile("hyperlane/announce/v1/tx.proto", fileDescriptor_0d512fd238195273) }

var fileDescriptor_0d512fd238195273 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x31, 0x0f, 0xd2, 0x40,
	0x18, 0xed, 0x89, 0x92, 0x70, 0x21, 0x41, 0x1b, 0x0c, 0xa5, 0x31, 0x85, 0x30, 0x18, 0x24, 0xa1,
	0x27, 0xba, 0x18, 0x37, 0x98, 0x65, 0xa9, 0x9b, 0x0e, 0xe4, 0x68, 0xcf, 0xa3, 0x49, 0x7b, 0xd7,
	0xdc, 0x77, 0x34, 0xb0, 0x19, 0x27, 0xe3, 0xe4, 0x4f, 0xe0, 0x27, 0x30, 0xf8, 0x23, 0x1c, 0x89,
	0x93, 0xa3, 0x01, 0x13, 0xfc, 0x19, 0xa6, 0xf4, 0x80, 0x88, 0x0e, 0x2e, 0x4d, 0xef, 0xbd, 0x77,
	0xef, 0xe5, 0x7b, 0xdf, 0x61, 0x6f, 0xb1, 0xce, 0x98, 0x4a, 0xa8, 0x60, 0x84, 0x0a, 0x21, 0x97,
	0x22, 0x64, 0x24, 0x1f, 0x11, 0xbd, 0xf2, 0x33, 0x25, 0xb5, 0xb4, 0x1f, 0x5e, 0x78, 0xff, 0xcc,
	0xfb, 0xf9, 0xc8, 0x6d, 0x72, 0xc9, 0xe5, 0x49, 0x41, 0x8a, 0xbf, 0x52, 0xec, 0x3e, 0xa0, 0x69,
	0x2c, 0x24, 0x39, 0x7d, 0x0d, 0xd4, 0x0a, 0x25, 0xa4, 0x12, 0x48, 0x0a, 0xbc, 0xf0, 0x4d, 0x81,
	0x1b, 0xa2, 0x5d, 0x12, 0xb3, 0xd2, 0xa4, 0x3c, 0x94, 0x54, 0xef, 0x27, 0xc2, 0x8d, 0x29, 0xf0,
	0xb1, 0xc9, 0x4b, 0x99, 0xd0, 0xf6, 0x53, 0x5c, 0x05, 0x26, 0x22, 0xa6, 0x1c, 0xd4, 0x45, 0xfd,
	0xda, 0xc4, 0xf9, 0xf6, 0x65, 0xd8, 0x34, 0xb7, 0xc6, 0x51, 0xa4, 0x18, 0xc0, 0x6b, 0xad, 0x62,
	0xc1, 0x03, 0xa3, 0xb3, 0x1f, 0xe1, 0x5a, 0x4e, 0x93, 0x38, 0xa2, 0x5a, 0x2a, 0xe7, 0x4e, 0x17,
	0xf5, 0xeb, 0xc1, 0x15, 0xb0, 0x9f, 0xe0, 0xfb, 0xa0, 0xa5, 0xa2, 0x9c, 0xcd, 0x12, 0x19, 0x52,
	0x1d, 0x4b, 0xe1, 0x54, 0x0a, 0xe7, 0xa0, 0x61, 0xf0, 0x57, 0x06, 0x2e, 0x8c, 0x20, 0xe6, 0x82,
	0xea, 0xa5, 0x62, 0xce, 0xdd, 0xd2, 0xe8, 0x02, 0xbc, 0x7c, 0xf1, 0x71, 0xd3, 0xb1, 0x7e, 0x6d,
	0x3a, 0xd6, 0x87, 0xe3, 0x76, 0x60, 0xb2, 0x3f, 0x1d, 0xb7, 0x83, 0xee, 0xa5, 0xb8, 0xa1, 0xa9,
	0xe0, 0x66, 0xa4, 0x5e, 0x1b, 0xb7, 0x6e, 0xa0, 0x80, 0x41, 0x26, 0x05, 0xb0, 0x67, 0x1a, 0x57,
	0xa6, 0xc0, 0xed, 0x77, 0xb8, 0xfe, 0x47, 0x09, 0x8f, 0xfd, 0x7f, 0x6e, 0xc3, 0xbf, 0xb1, 0x71,
	0xfd, 0xff, 0xd3, 0x9d, 0xe3, 0xdc, 0x7b, 0xef, 0x8f, 0xdb, 0x01, 0x9a, 0xbc, 0xfd, 0xba, 0xf7,
	0xd0, 0x6e, 0xef, 0xa1, 0x1f, 0x7b, 0x0f, 0x7d, 0x3e, 0x78, 0xd6, 0xee, 0xe0, 0x59, 0xdf, 0x0f,
	0x9e, 0xf5, 0x66, 0xcc, 0x63, 0xbd, 0x58, 0xce, 0xfd, 0x50, 0xa6, 0x04, 0xb4, 0xa2, 0x82, 0xb3,
	0x44, 0xe6, 0x6c, 0x98, 0x33, 0x51, 0x94, 0x00, 0xe4, 0xaf, 0x61, 0x57, 0xd7, 0xe7, 0xa4, 0xd7,
	0x19, 0x83, 0x79, 0xf5, 0xb4, 0xdb, 0xe7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x14, 0x77, 0x54,
	0xcb, 0x71, 0x02, 0x00, 0x00,
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
	// Announces a validator signature storage location
	Announcement(ctx context.Context, in *MsgAnnouncement, opts ...grpc.CallOption) (*MsgAnnouncementResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Announcement(ctx context.Context, in *MsgAnnouncement, opts ...grpc.CallOption) (*MsgAnnouncementResponse, error) {
	out := new(MsgAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/hyperlane.announce.v1.Msg/Announcement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Announces a validator signature storage location
	Announcement(context.Context, *MsgAnnouncement) (*MsgAnnouncementResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct{}

func (*UnimplementedMsgServer) Announcement(ctx context.Context, req *MsgAnnouncement) (*MsgAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Announcement not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Announcement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAnnouncement)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Announcement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hyperlane.announce.v1.Msg/Announcement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Announcement(ctx, req.(*MsgAnnouncement))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hyperlane.announce.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Announcement",
			Handler:    _Msg_Announcement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hyperlane/announce/v1/tx.proto",
}

func (m *MsgAnnouncement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAnnouncement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAnnouncement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.StorageLocation) > 0 {
		i -= len(m.StorageLocation)
		copy(dAtA[i:], m.StorageLocation)
		i = encodeVarintTx(dAtA, i, uint64(len(m.StorageLocation)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAnnouncementResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAnnouncementResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAnnouncementResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *MsgAnnouncement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.StorageLocation)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAnnouncementResponse) Size() (n int) {
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

func (m *MsgAnnouncement) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAnnouncement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAnnouncement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = append(m.Validator[:0], dAtA[iNdEx:postIndex]...)
			if m.Validator == nil {
				m.Validator = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageLocation", wireType)
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
			m.StorageLocation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
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

func (m *MsgAnnouncementResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAnnouncementResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAnnouncementResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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