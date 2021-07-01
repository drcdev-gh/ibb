// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibb/borrowAccrued.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type BorrowAccrued struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id          uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	BlockHeight int32  `protobuf:"varint,3,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	Asset       string `protobuf:"bytes,4,opt,name=asset,proto3" json:"asset,omitempty"`
	Amount      int32  `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Denom       string `protobuf:"bytes,6,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *BorrowAccrued) Reset()         { *m = BorrowAccrued{} }
func (m *BorrowAccrued) String() string { return proto.CompactTextString(m) }
func (*BorrowAccrued) ProtoMessage()    {}
func (*BorrowAccrued) Descriptor() ([]byte, []int) {
	return fileDescriptor_674c2a307c2cb750, []int{0}
}
func (m *BorrowAccrued) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BorrowAccrued) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BorrowAccrued.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BorrowAccrued) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BorrowAccrued.Merge(m, src)
}
func (m *BorrowAccrued) XXX_Size() int {
	return m.Size()
}
func (m *BorrowAccrued) XXX_DiscardUnknown() {
	xxx_messageInfo_BorrowAccrued.DiscardUnknown(m)
}

var xxx_messageInfo_BorrowAccrued proto.InternalMessageInfo

func (m *BorrowAccrued) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *BorrowAccrued) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BorrowAccrued) GetBlockHeight() int32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *BorrowAccrued) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *BorrowAccrued) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *BorrowAccrued) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*BorrowAccrued)(nil), "sapienscosmos.ibb.ibb.BorrowAccrued")
}

func init() { proto.RegisterFile("ibb/borrowAccrued.proto", fileDescriptor_674c2a307c2cb750) }

var fileDescriptor_674c2a307c2cb750 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4e, 0xc4, 0x30,
	0x10, 0x45, 0xe3, 0xb0, 0x09, 0xc2, 0x08, 0x0a, 0x6b, 0x01, 0x8b, 0xc2, 0x8a, 0xa8, 0x42, 0x41,
	0x52, 0x70, 0x02, 0x96, 0x86, 0x3a, 0x25, 0x5d, 0xc6, 0xb1, 0xb2, 0x16, 0x24, 0x13, 0xd9, 0x8e,
	0x80, 0x5b, 0x70, 0x03, 0xae, 0x43, 0xb9, 0x25, 0x25, 0x4a, 0x2e, 0x82, 0x62, 0x2f, 0xd2, 0x16,
	0xb6, 0xfc, 0xac, 0x37, 0x5f, 0xfa, 0x43, 0xaf, 0x34, 0x40, 0x09, 0x68, 0x0c, 0xbe, 0x3d, 0x48,
	0x69, 0x46, 0xd5, 0x14, 0x83, 0x41, 0x87, 0xec, 0xc2, 0xd6, 0x83, 0x56, 0xbd, 0x95, 0x68, 0x3b,
	0xb4, 0x85, 0x06, 0x58, 0xce, 0xf5, 0xba, 0xc5, 0x16, 0xbd, 0x51, 0x2e, 0xaf, 0x20, 0xdf, 0x7c,
	0x11, 0x7a, 0xb6, 0x39, 0x0c, 0x61, 0x9c, 0x1e, 0x4b, 0xa3, 0x6a, 0x87, 0x86, 0x93, 0x8c, 0xe4,
	0x27, 0xd5, 0x3f, 0xb2, 0x73, 0x1a, 0xeb, 0x86, 0xc7, 0x19, 0xc9, 0x57, 0x55, 0xac, 0x1b, 0x96,
	0xd1, 0x53, 0x78, 0x45, 0xf9, 0xf2, 0xa4, 0x74, 0xbb, 0x75, 0xfc, 0x28, 0x23, 0x79, 0x52, 0x1d,
	0x7e, 0xb1, 0x35, 0x4d, 0x6a, 0x6b, 0x95, 0xe3, 0x2b, 0x9f, 0x14, 0x80, 0x5d, 0xd2, 0xb4, 0xee,
	0x70, 0xec, 0x1d, 0x4f, 0xfc, 0xc8, 0x9e, 0x16, 0xbb, 0x51, 0x3d, 0x76, 0x3c, 0x0d, 0xb6, 0x87,
	0xcd, 0xe3, 0xf7, 0x24, 0xc8, 0x6e, 0x12, 0xe4, 0x77, 0x12, 0xe4, 0x73, 0x16, 0xd1, 0x6e, 0x16,
	0xd1, 0xcf, 0x2c, 0xa2, 0xe7, 0xdb, 0x56, 0xbb, 0xed, 0x08, 0x85, 0xc4, 0xae, 0xdc, 0x77, 0xbe,
	0x0b, 0xa5, 0xcb, 0x65, 0x37, 0xef, 0xfe, 0x76, 0x1f, 0x83, 0xb2, 0x90, 0xfa, 0xb6, 0xf7, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x4c, 0x44, 0xee, 0x35, 0x01, 0x00, 0x00,
}

func (m *BorrowAccrued) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BorrowAccrued) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BorrowAccrued) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintBorrowAccrued(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x32
	}
	if m.Amount != 0 {
		i = encodeVarintBorrowAccrued(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintBorrowAccrued(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0x22
	}
	if m.BlockHeight != 0 {
		i = encodeVarintBorrowAccrued(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintBorrowAccrued(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBorrowAccrued(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBorrowAccrued(dAtA []byte, offset int, v uint64) int {
	offset -= sovBorrowAccrued(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BorrowAccrued) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBorrowAccrued(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovBorrowAccrued(uint64(m.Id))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovBorrowAccrued(uint64(m.BlockHeight))
	}
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovBorrowAccrued(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovBorrowAccrued(uint64(m.Amount))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovBorrowAccrued(uint64(l))
	}
	return n
}

func sovBorrowAccrued(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBorrowAccrued(x uint64) (n int) {
	return sovBorrowAccrued(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BorrowAccrued) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBorrowAccrued
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
			return fmt.Errorf("proto: BorrowAccrued: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BorrowAccrued: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrowAccrued
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
				return ErrInvalidLengthBorrowAccrued
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBorrowAccrued
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrowAccrued
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrowAccrued
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrowAccrued
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
				return ErrInvalidLengthBorrowAccrued
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBorrowAccrued
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrowAccrued
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrowAccrued
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
				return ErrInvalidLengthBorrowAccrued
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBorrowAccrued
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBorrowAccrued(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBorrowAccrued
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
func skipBorrowAccrued(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBorrowAccrued
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
					return 0, ErrIntOverflowBorrowAccrued
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
					return 0, ErrIntOverflowBorrowAccrued
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
				return 0, ErrInvalidLengthBorrowAccrued
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBorrowAccrued
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBorrowAccrued
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBorrowAccrued        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBorrowAccrued          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBorrowAccrued = fmt.Errorf("proto: unexpected end of group")
)
