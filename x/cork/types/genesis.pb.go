// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cork/v1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState - all cork state that must be provided at genesis
type GenesisState struct {
	Params    Params      `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	CellarIds CellarIDSet `protobuf:"bytes,2,opt,name=cellar_ids,json=cellarIds,proto3" json:"cellar_ids"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec09bf5ff5694398, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetCellarIds() CellarIDSet {
	if m != nil {
		return m.CellarIds
	}
	return CellarIDSet{}
}

// Params cork parameters
type Params struct {
	// VotePeriod defines the number of blocks to wait for votes before attempting to tally
	VotePeriod int64 `protobuf:"varint,1,opt,name=vote_period,json=votePeriod,proto3" json:"vote_period,omitempty" yaml:"vote_period"`
	// VoteThreshold defines the percentage of bonded stake required to vote each period
	VoteThreshold github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=vote_threshold,json=voteThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"vote_threshold" yaml:"vote_threshold"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec09bf5ff5694398, []int{1}
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

func (m *Params) GetVotePeriod() int64 {
	if m != nil {
		return m.VotePeriod
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cork.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "cork.v1.Params")
}

func init() { proto.RegisterFile("cork/v1/genesis.proto", fileDescriptor_ec09bf5ff5694398) }

var fileDescriptor_ec09bf5ff5694398 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x4f, 0x4e, 0xc2, 0x40,
	0x14, 0xc6, 0x5b, 0x35, 0x18, 0x06, 0xff, 0x65, 0x02, 0x86, 0xb0, 0x68, 0x4d, 0x17, 0xc6, 0x85,
	0x74, 0x82, 0x2c, 0x8c, 0x2e, 0x2b, 0x91, 0xb0, 0x23, 0xe0, 0xca, 0x0d, 0x29, 0xed, 0x4b, 0xa9,
	0xb4, 0x4c, 0x33, 0x33, 0x36, 0x70, 0x0b, 0xaf, 0xe2, 0x2d, 0x58, 0xb2, 0x34, 0x2e, 0x1a, 0x03,
	0x37, 0xe0, 0x04, 0xa6, 0x33, 0x85, 0xb0, 0x9a, 0xf7, 0xbe, 0xef, 0xbd, 0xef, 0x37, 0x99, 0x41,
	0x35, 0x8f, 0xb2, 0x29, 0x49, 0x5b, 0x24, 0x80, 0x19, 0xf0, 0x90, 0xdb, 0x09, 0xa3, 0x82, 0xe2,
	0xd3, 0x5c, 0xb6, 0xd3, 0x56, 0xe3, 0x6a, 0xe7, 0x8b, 0xb9, 0xb2, 0x1a, 0x78, 0xa7, 0xc8, 0x11,
	0xa5, 0x55, 0x03, 0x1a, 0x50, 0x59, 0x92, 0xbc, 0x52, 0xaa, 0x35, 0x47, 0x67, 0x5d, 0x95, 0x3a,
	0x14, 0xae, 0x00, 0xdc, 0x44, 0xa5, 0xc4, 0x65, 0x6e, 0xcc, 0xeb, 0xfa, 0x8d, 0x7e, 0x57, 0x79,
	0xb8, 0xb4, 0x0b, 0x8a, 0xdd, 0x97, 0xb2, 0x73, 0xb2, 0xcc, 0x4c, 0x6d, 0x50, 0x0c, 0xe1, 0x27,
	0x84, 0x3c, 0x88, 0x22, 0x97, 0x8d, 0x42, 0x9f, 0xd7, 0x8f, 0xe4, 0x4a, 0x75, 0xbf, 0xf2, 0x22,
	0xad, 0x5e, 0x67, 0x08, 0xa2, 0xd8, 0x2b, 0xab, 0xe9, 0x9e, 0xcf, 0xad, 0x6f, 0x1d, 0x95, 0x54,
	0x26, 0x7e, 0x44, 0x95, 0x94, 0x0a, 0x18, 0x25, 0xc0, 0x42, 0xea, 0x4b, 0xf2, 0xb1, 0x73, 0xbd,
	0xcd, 0x4c, 0xbc, 0x70, 0xe3, 0xe8, 0xd9, 0x3a, 0x30, 0xad, 0x01, 0xca, 0xbb, 0xbe, 0x6c, 0xf0,
	0x0c, 0x5d, 0x48, 0x4f, 0x4c, 0x18, 0xf0, 0x09, 0x8d, 0x7c, 0x79, 0x85, 0xb2, 0xd3, 0xcd, 0x61,
	0xbf, 0x99, 0x79, 0x1b, 0x84, 0x62, 0xf2, 0x39, 0xb6, 0x3d, 0x1a, 0x13, 0x8f, 0xf2, 0x98, 0xf2,
	0xe2, 0x68, 0x72, 0x7f, 0x4a, 0xc4, 0x22, 0x01, 0x6e, 0x77, 0xc0, 0xdb, 0x66, 0x66, 0xed, 0x80,
	0xb4, 0x4f, 0xb3, 0x06, 0xe7, 0xb9, 0xf0, 0xb6, 0xeb, 0x9d, 0xd7, 0xe5, 0xda, 0xd0, 0x57, 0x6b,
	0x43, 0xff, 0x5b, 0x1b, 0xfa, 0xd7, 0xc6, 0xd0, 0x56, 0x1b, 0x43, 0xfb, 0xd9, 0x18, 0xda, 0xfb,
	0xfd, 0x01, 0x29, 0x81, 0x20, 0x58, 0x7c, 0xa4, 0x84, 0xd3, 0x38, 0x86, 0x28, 0x04, 0x46, 0xd2,
	0x36, 0x99, 0xcb, 0xbf, 0x50, 0xcc, 0x71, 0x49, 0x3e, 0x7e, 0xfb, 0x3f, 0x00, 0x00, 0xff, 0xff,
	0xc2, 0x8e, 0x31, 0x7e, 0xda, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CellarIds.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
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
	{
		size := m.VoteThreshold.Size()
		i -= size
		if _, err := m.VoteThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.VotePeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.VotePeriod))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.CellarIds.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VotePeriod != 0 {
		n += 1 + sovGenesis(uint64(m.VotePeriod))
	}
	l = m.VoteThreshold.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CellarIds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CellarIds.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
				return fmt.Errorf("proto: wrong wireType = %d for field VotePeriod", wireType)
			}
			m.VotePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotePeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VoteThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)