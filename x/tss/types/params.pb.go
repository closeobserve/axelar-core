// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tss/v1beta1/params.proto

package types

import (
	fmt "fmt"
	utils "github.com/axelarnetwork/axelar-core/utils"
	exported "github.com/axelarnetwork/axelar-core/x/tss/exported"
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

// Params is the parameter set for this module
type Params struct {
	// Deprecated
	LockingPeriod int64 `protobuf:"varint,1,opt,name=locking_period,json=lockingPeriod,proto3" json:"locking_period,omitempty"`
	// MinKeygenThreshold defines the minimum % of stake that must be online
	// to authorize generation of a new key in the system.
	MinKeygenThreshold utils.Threshold `protobuf:"bytes,2,opt,name=min_keygen_threshold,json=minKeygenThreshold,proto3" json:"min_keygen_threshold"`
	// CorruptionThreshold defines the corruption threshold with which
	// we'll run keygen protocol.
	CorruptionThreshold utils.Threshold `protobuf:"bytes,3,opt,name=corruption_threshold,json=corruptionThreshold,proto3" json:"corruption_threshold"`
	// KeyRequirements defines the requirement of each key for each chain
	KeyRequirements []exported.KeyRequirement `protobuf:"bytes,4,rep,name=key_requirements,json=keyRequirements,proto3" json:"key_requirements"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c9a42e8b26dfec, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "tss.v1beta1.Params")
}

func init() { proto.RegisterFile("tss/v1beta1/params.proto", fileDescriptor_67c9a42e8b26dfec) }

var fileDescriptor_67c9a42e8b26dfec = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x5b, 0x20, 0x1c, 0x4a, 0xfc, 0x93, 0xca, 0xa1, 0x21, 0x71, 0x6d, 0x8c, 0x26, 0x5c,
	0x6c, 0x05, 0xdf, 0x80, 0x2b, 0x89, 0x41, 0xa2, 0x17, 0x2f, 0x4d, 0x29, 0x93, 0xb2, 0xe9, 0x9f,
	0xa9, 0xb3, 0x5b, 0xa5, 0x6f, 0xe1, 0x7b, 0xf8, 0x22, 0x1c, 0x39, 0x7a, 0x32, 0x0a, 0x2f, 0x62,
	0xba, 0x16, 0xaa, 0x37, 0x6f, 0xbb, 0xdf, 0x37, 0xf3, 0x9b, 0xc9, 0x7c, 0x86, 0x25, 0x85, 0x70,
	0x9f, 0x07, 0x33, 0x90, 0xfe, 0xc0, 0xcd, 0x7c, 0xf2, 0x13, 0xe1, 0x64, 0x84, 0x12, 0xcd, 0x8e,
	0x14, 0xc2, 0xa9, 0x9c, 0x5e, 0x37, 0xc4, 0x10, 0x95, 0xee, 0x96, 0xaf, 0x9f, 0x92, 0xde, 0x69,
	0x2e, 0x79, 0x5c, 0xb7, 0xcb, 0x05, 0x81, 0x58, 0x60, 0x3c, 0xaf, 0x6c, 0xbb, 0x64, 0xc3, 0x32,
	0x43, 0x92, 0x30, 0xaf, 0xab, 0x8a, 0x0c, 0xaa, 0x19, 0xe7, 0x6f, 0x0d, 0xa3, 0x3d, 0x51, 0x43,
	0xcd, 0x4b, 0xe3, 0x30, 0xc6, 0x20, 0xe2, 0x69, 0xe8, 0x65, 0x40, 0x1c, 0xe7, 0x96, 0x6e, 0xeb,
	0xfd, 0xe6, 0xf4, 0xa0, 0x52, 0x27, 0x4a, 0x34, 0x27, 0x46, 0x37, 0xe1, 0xa9, 0x17, 0x41, 0x11,
	0x42, 0xea, 0xed, 0x27, 0x5a, 0x0d, 0x5b, 0xef, 0x77, 0x86, 0x96, 0xa3, 0x36, 0xda, 0xad, 0xed,
	0xdc, 0xef, 0xfc, 0x51, 0x6b, 0xf5, 0x71, 0xa6, 0x4d, 0xcd, 0x84, 0xa7, 0x63, 0xd5, 0xba, 0x77,
	0xcc, 0x3b, 0xa3, 0x1b, 0x20, 0x51, 0x9e, 0x49, 0x8e, 0xbf, 0x89, 0xcd, 0x7f, 0x11, 0x4f, 0xea,
	0xde, 0x1a, 0xf9, 0x60, 0x1c, 0x47, 0x50, 0x78, 0x04, 0x4f, 0x39, 0x27, 0x48, 0x20, 0x95, 0xc2,
	0x6a, 0xd9, 0xcd, 0x7e, 0x67, 0x78, 0xe1, 0x94, 0x57, 0xdd, 0xdd, 0x64, 0x4f, 0x1d, 0x43, 0x31,
	0xad, 0x8b, 0x2b, 0xf4, 0x51, 0xf4, 0x47, 0x15, 0xa3, 0xdb, 0xd5, 0x17, 0xd3, 0x56, 0x1b, 0xa6,
	0xaf, 0x37, 0x4c, 0xff, 0xdc, 0x30, 0xfd, 0x75, 0xcb, 0xb4, 0xf5, 0x96, 0x69, 0xef, 0x5b, 0xa6,
	0x3d, 0x5e, 0x87, 0x5c, 0x2e, 0xf2, 0x99, 0x13, 0x60, 0xe2, 0xfa, 0x4b, 0x88, 0x7d, 0x4a, 0x41,
	0xbe, 0x20, 0x45, 0xd5, 0xef, 0x2a, 0x40, 0x02, 0x77, 0xe9, 0x96, 0xa1, 0xa8, 0x0c, 0x66, 0x6d,
	0x15, 0xc2, 0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0xa6, 0x27, 0xdf, 0x04, 0x02, 0x00,
	0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.KeyRequirements) > 0 {
		for iNdEx := len(m.KeyRequirements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.KeyRequirements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size, err := m.CorruptionThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.MinKeygenThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.LockingPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.LockingPeriod))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LockingPeriod != 0 {
		n += 1 + sovParams(uint64(m.LockingPeriod))
	}
	l = m.MinKeygenThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.CorruptionThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.KeyRequirements) > 0 {
		for _, e := range m.KeyRequirements {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockingPeriod", wireType)
			}
			m.LockingPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LockingPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinKeygenThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinKeygenThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorruptionThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CorruptionThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyRequirements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyRequirements = append(m.KeyRequirements, exported.KeyRequirement{})
			if err := m.KeyRequirements[len(m.KeyRequirements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
