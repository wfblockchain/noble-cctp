// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: noble/cctp/v1/token_pair.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// *
// TokenPair is used to look up the Noble token (i.e. "uusdc") from a remote domain token address
// Multiple remote_domain + remote_token pairs can map to the same local_token
//
// @param remote_domain the remote domain_id corresponding to the token
// @param remote_token the remote token address
// @param local_token the corresponding Noble token denom in uunits
type TokenPair struct {
	RemoteDomain uint32 `protobuf:"varint,1,opt,name=remote_domain,json=remoteDomain,proto3" json:"remote_domain,omitempty"`
	RemoteToken  string `protobuf:"bytes,2,opt,name=remote_token,json=remoteToken,proto3" json:"remote_token,omitempty"`
	LocalToken   string `protobuf:"bytes,3,opt,name=local_token,json=localToken,proto3" json:"local_token,omitempty"`
}

func (m *TokenPair) Reset()         { *m = TokenPair{} }
func (m *TokenPair) String() string { return proto.CompactTextString(m) }
func (*TokenPair) ProtoMessage()    {}
func (*TokenPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a764822a752ee9, []int{0}
}
func (m *TokenPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenPair.Merge(m, src)
}
func (m *TokenPair) XXX_Size() int {
	return m.Size()
}
func (m *TokenPair) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenPair.DiscardUnknown(m)
}

var xxx_messageInfo_TokenPair proto.InternalMessageInfo

func (m *TokenPair) GetRemoteDomain() uint32 {
	if m != nil {
		return m.RemoteDomain
	}
	return 0
}

func (m *TokenPair) GetRemoteToken() string {
	if m != nil {
		return m.RemoteToken
	}
	return ""
}

func (m *TokenPair) GetLocalToken() string {
	if m != nil {
		return m.LocalToken
	}
	return ""
}

func init() {
	proto.RegisterType((*TokenPair)(nil), "noble.cctp.v1.TokenPair")
}

func init() { proto.RegisterFile("noble/cctp/v1/token_pair.proto", fileDescriptor_73a764822a752ee9) }

var fileDescriptor_73a764822a752ee9 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0xcb, 0x4f, 0xca,
	0x49, 0xd5, 0x4f, 0x4e, 0x2e, 0x29, 0xd0, 0x2f, 0x33, 0xd4, 0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0x8b,
	0x2f, 0x48, 0xcc, 0x2c, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x05, 0xcb, 0xeb, 0x81,
	0xe4, 0xf5, 0xca, 0x0c, 0x95, 0x4a, 0xb8, 0x38, 0x43, 0x40, 0x4a, 0x02, 0x12, 0x33, 0x8b, 0x84,
	0x94, 0xb9, 0x78, 0x8b, 0x52, 0x73, 0xf3, 0x4b, 0x52, 0xe3, 0x53, 0xf2, 0x73, 0x13, 0x33, 0xf3,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0x78, 0x20, 0x82, 0x2e, 0x60, 0x31, 0x21, 0x45, 0x2e,
	0x28, 0x3f, 0x1e, 0x6c, 0xb6, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x37, 0x44, 0x0c, 0x6c,
	0x96, 0x90, 0x3c, 0x17, 0x77, 0x4e, 0x7e, 0x72, 0x62, 0x0e, 0x54, 0x05, 0x33, 0x58, 0x05, 0x17,
	0x58, 0x08, 0xac, 0xc0, 0xc9, 0xed, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2,
	0x74, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0x33, 0x8b, 0x92,
	0x73, 0x52, 0xd3, 0x32, 0xf3, 0xf4, 0xc1, 0x6e, 0xd6, 0x05, 0xfb, 0xa9, 0x02, 0xe2, 0xb5, 0x92,
	0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x9f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8d,
	0xb4, 0x0e, 0xe6, 0xf5, 0x00, 0x00, 0x00,
}

func (m *TokenPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LocalToken) > 0 {
		i -= len(m.LocalToken)
		copy(dAtA[i:], m.LocalToken)
		i = encodeVarintTokenPair(dAtA, i, uint64(len(m.LocalToken)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RemoteToken) > 0 {
		i -= len(m.RemoteToken)
		copy(dAtA[i:], m.RemoteToken)
		i = encodeVarintTokenPair(dAtA, i, uint64(len(m.RemoteToken)))
		i--
		dAtA[i] = 0x12
	}
	if m.RemoteDomain != 0 {
		i = encodeVarintTokenPair(dAtA, i, uint64(m.RemoteDomain))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTokenPair(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokenPair(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RemoteDomain != 0 {
		n += 1 + sovTokenPair(uint64(m.RemoteDomain))
	}
	l = len(m.RemoteToken)
	if l > 0 {
		n += 1 + l + sovTokenPair(uint64(l))
	}
	l = len(m.LocalToken)
	if l > 0 {
		n += 1 + l + sovTokenPair(uint64(l))
	}
	return n
}

func sovTokenPair(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokenPair(x uint64) (n int) {
	return sovTokenPair(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenPair
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
			return fmt.Errorf("proto: TokenPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteDomain", wireType)
			}
			m.RemoteDomain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemoteDomain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenPair
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
				return ErrInvalidLengthTokenPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenPair
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
				return ErrInvalidLengthTokenPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LocalToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenPair
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
func skipTokenPair(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokenPair
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
					return 0, ErrIntOverflowTokenPair
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
					return 0, ErrIntOverflowTokenPair
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
				return 0, ErrInvalidLengthTokenPair
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokenPair
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokenPair
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokenPair        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokenPair          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokenPair = fmt.Errorf("proto: unexpected end of group")
)
