// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tunnel.proto

package camo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type TCPPacket struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TCPPacket) Reset()         { *m = TCPPacket{} }
func (m *TCPPacket) String() string { return proto.CompactTextString(m) }
func (*TCPPacket) ProtoMessage()    {}
func (*TCPPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{0}
}
func (m *TCPPacket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TCPPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TCPPacket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TCPPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TCPPacket.Merge(m, src)
}
func (m *TCPPacket) XXX_Size() int {
	return m.Size()
}
func (m *TCPPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_TCPPacket.DiscardUnknown(m)
}

var xxx_messageInfo_TCPPacket proto.InternalMessageInfo

func (m *TCPPacket) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type UDPPacket struct {
	Dst                  string   `protobuf:"bytes,1,opt,name=dst,proto3" json:"dst,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UDPPacket) Reset()         { *m = UDPPacket{} }
func (m *UDPPacket) String() string { return proto.CompactTextString(m) }
func (*UDPPacket) ProtoMessage()    {}
func (*UDPPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{1}
}
func (m *UDPPacket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UDPPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UDPPacket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UDPPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UDPPacket.Merge(m, src)
}
func (m *UDPPacket) XXX_Size() int {
	return m.Size()
}
func (m *UDPPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_UDPPacket.DiscardUnknown(m)
}

var xxx_messageInfo_UDPPacket proto.InternalMessageInfo

func (m *UDPPacket) GetDst() string {
	if m != nil {
		return m.Dst
	}
	return ""
}

func (m *UDPPacket) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*TCPPacket)(nil), "camo.TCPPacket")
	proto.RegisterType((*UDPPacket)(nil), "camo.UDPPacket")
}

func init() { proto.RegisterFile("tunnel.proto", fileDescriptor_6f51ddaa7891a711) }

var fileDescriptor_6f51ddaa7891a711 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x29, 0xcd, 0xcb,
	0x4b, 0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x4e, 0xcc, 0xcd, 0x57, 0x92,
	0xe7, 0xe2, 0x0c, 0x71, 0x0e, 0x08, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0x11, 0x12, 0xe2, 0x62, 0x49,
	0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x95, 0x0c, 0xb9, 0x38,
	0x43, 0x5d, 0x60, 0x0a, 0x04, 0xb8, 0x98, 0x53, 0x8a, 0x4b, 0xc0, 0xf2, 0x9c, 0x41, 0x20, 0x26,
	0x5c, 0x0b, 0x13, 0x42, 0x8b, 0x51, 0x1e, 0x17, 0x5b, 0x08, 0xd8, 0x26, 0x21, 0x43, 0x2e, 0x76,
	0xe7, 0xfc, 0xbc, 0xbc, 0xd4, 0xe4, 0x12, 0x21, 0x7e, 0x3d, 0x90, 0x7d, 0x7a, 0x70, 0xcb, 0xa4,
	0xd0, 0x05, 0x94, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x41, 0x5a, 0xfc, 0x0b, 0x52, 0xf3, 0x42, 0x5d,
	0x02, 0x60, 0x5a, 0xe0, 0xd6, 0x4b, 0xa1, 0x0b, 0x40, 0xb4, 0x38, 0x09, 0x9c, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81,
	0xbd, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x43, 0x38, 0xa2, 0xf2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TunnelClient is the client API for Tunnel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TunnelClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (Tunnel_ConnectClient, error)
	OpenUDP(ctx context.Context, opts ...grpc.CallOption) (Tunnel_OpenUDPClient, error)
}

type tunnelClient struct {
	cc *grpc.ClientConn
}

func NewTunnelClient(cc *grpc.ClientConn) TunnelClient {
	return &tunnelClient{cc}
}

func (c *tunnelClient) Connect(ctx context.Context, opts ...grpc.CallOption) (Tunnel_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tunnel_serviceDesc.Streams[0], "/camo.Tunnel/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelConnectClient{stream}
	return x, nil
}

type Tunnel_ConnectClient interface {
	Send(*TCPPacket) error
	Recv() (*TCPPacket, error)
	grpc.ClientStream
}

type tunnelConnectClient struct {
	grpc.ClientStream
}

func (x *tunnelConnectClient) Send(m *TCPPacket) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelConnectClient) Recv() (*TCPPacket, error) {
	m := new(TCPPacket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tunnelClient) OpenUDP(ctx context.Context, opts ...grpc.CallOption) (Tunnel_OpenUDPClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tunnel_serviceDesc.Streams[1], "/camo.Tunnel/OpenUDP", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelOpenUDPClient{stream}
	return x, nil
}

type Tunnel_OpenUDPClient interface {
	Send(*UDPPacket) error
	Recv() (*UDPPacket, error)
	grpc.ClientStream
}

type tunnelOpenUDPClient struct {
	grpc.ClientStream
}

func (x *tunnelOpenUDPClient) Send(m *UDPPacket) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelOpenUDPClient) Recv() (*UDPPacket, error) {
	m := new(UDPPacket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TunnelServer is the server API for Tunnel service.
type TunnelServer interface {
	Connect(Tunnel_ConnectServer) error
	OpenUDP(Tunnel_OpenUDPServer) error
}

// UnimplementedTunnelServer can be embedded to have forward compatible implementations.
type UnimplementedTunnelServer struct {
}

func (*UnimplementedTunnelServer) Connect(srv Tunnel_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedTunnelServer) OpenUDP(srv Tunnel_OpenUDPServer) error {
	return status.Errorf(codes.Unimplemented, "method OpenUDP not implemented")
}

func RegisterTunnelServer(s *grpc.Server, srv TunnelServer) {
	s.RegisterService(&_Tunnel_serviceDesc, srv)
}

func _Tunnel_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelServer).Connect(&tunnelConnectServer{stream})
}

type Tunnel_ConnectServer interface {
	Send(*TCPPacket) error
	Recv() (*TCPPacket, error)
	grpc.ServerStream
}

type tunnelConnectServer struct {
	grpc.ServerStream
}

func (x *tunnelConnectServer) Send(m *TCPPacket) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelConnectServer) Recv() (*TCPPacket, error) {
	m := new(TCPPacket)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tunnel_OpenUDP_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelServer).OpenUDP(&tunnelOpenUDPServer{stream})
}

type Tunnel_OpenUDPServer interface {
	Send(*UDPPacket) error
	Recv() (*UDPPacket, error)
	grpc.ServerStream
}

type tunnelOpenUDPServer struct {
	grpc.ServerStream
}

func (x *tunnelOpenUDPServer) Send(m *UDPPacket) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelOpenUDPServer) Recv() (*UDPPacket, error) {
	m := new(UDPPacket)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Tunnel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "camo.Tunnel",
	HandlerType: (*TunnelServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Tunnel_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "OpenUDP",
			Handler:       _Tunnel_OpenUDP_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "tunnel.proto",
}

func (m *TCPPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TCPPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TCPPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTunnel(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UDPPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UDPPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UDPPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTunnel(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Dst) > 0 {
		i -= len(m.Dst)
		copy(dAtA[i:], m.Dst)
		i = encodeVarintTunnel(dAtA, i, uint64(len(m.Dst)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTunnel(dAtA []byte, offset int, v uint64) int {
	offset -= sovTunnel(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TCPPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTunnel(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UDPPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Dst)
	if l > 0 {
		n += 1 + l + sovTunnel(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTunnel(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTunnel(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTunnel(x uint64) (n int) {
	return sovTunnel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TCPPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTunnel
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
			return fmt.Errorf("proto: TCPPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TCPPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTunnel
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
				return ErrInvalidLengthTunnel
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTunnel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTunnel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTunnel
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTunnel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UDPPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTunnel
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
			return fmt.Errorf("proto: UDPPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UDPPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dst", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTunnel
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
				return ErrInvalidLengthTunnel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTunnel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dst = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTunnel
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
				return ErrInvalidLengthTunnel
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTunnel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTunnel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTunnel
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTunnel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTunnel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTunnel
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
					return 0, ErrIntOverflowTunnel
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTunnel
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
				return 0, ErrInvalidLengthTunnel
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTunnel
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTunnel
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTunnel(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTunnel
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTunnel = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTunnel   = fmt.Errorf("proto: integer overflow")
)