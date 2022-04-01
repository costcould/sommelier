// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cork/v1/cork.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type Cork struct {
	// call body containing the ABI encoded bytes to send to the contract
	EncodedContractCall []byte `protobuf:"bytes,1,opt,name=encoded_contract_call,json=encodedContractCall,proto3" json:"encoded_contract_call,omitempty"`
	// address of the contract to send the call
	TargetContractAddress string `protobuf:"bytes,2,opt,name=target_contract_address,json=targetContractAddress,proto3" json:"target_contract_address,omitempty"`
}

func (m *Cork) Reset()         { *m = Cork{} }
func (m *Cork) String() string { return proto.CompactTextString(m) }
func (*Cork) ProtoMessage()    {}
func (*Cork) Descriptor() ([]byte, []int) {
	return fileDescriptor_79882ab39b78d896, []int{0}
}
func (m *Cork) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cork.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Cork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cork.Merge(m, src)
}
func (m *Cork) XXX_Size() int {
	return m.Size()
}
func (m *Cork) XXX_DiscardUnknown() {
	xxx_messageInfo_Cork.DiscardUnknown(m)
}

var xxx_messageInfo_Cork proto.InternalMessageInfo

func (m *Cork) GetEncodedContractCall() []byte {
	if m != nil {
		return m.EncodedContractCall
	}
	return nil
}

func (m *Cork) GetTargetContractAddress() string {
	if m != nil {
		return m.TargetContractAddress
	}
	return ""
}

type ValidatorCork struct {
	Cork      *Cork  `protobuf:"bytes,1,opt,name=cork,proto3" json:"cork,omitempty"`
	Validator string `protobuf:"bytes,2,opt,name=validator,proto3" json:"validator,omitempty"`
}

func (m *ValidatorCork) Reset()         { *m = ValidatorCork{} }
func (m *ValidatorCork) String() string { return proto.CompactTextString(m) }
func (*ValidatorCork) ProtoMessage()    {}
func (*ValidatorCork) Descriptor() ([]byte, []int) {
	return fileDescriptor_79882ab39b78d896, []int{1}
}
func (m *ValidatorCork) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorCork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorCork.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorCork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorCork.Merge(m, src)
}
func (m *ValidatorCork) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorCork) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorCork.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorCork proto.InternalMessageInfo

func (m *ValidatorCork) GetCork() *Cork {
	if m != nil {
		return m.Cork
	}
	return nil
}

func (m *ValidatorCork) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

type ScheduledCork struct {
	Cork        *Cork  `protobuf:"bytes,1,opt,name=cork,proto3" json:"cork,omitempty"`
	BlockHeight uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Validator   string `protobuf:"bytes,3,opt,name=validator,proto3" json:"validator,omitempty"`
}

func (m *ScheduledCork) Reset()         { *m = ScheduledCork{} }
func (m *ScheduledCork) String() string { return proto.CompactTextString(m) }
func (*ScheduledCork) ProtoMessage()    {}
func (*ScheduledCork) Descriptor() ([]byte, []int) {
	return fileDescriptor_79882ab39b78d896, []int{2}
}
func (m *ScheduledCork) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduledCork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScheduledCork.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScheduledCork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduledCork.Merge(m, src)
}
func (m *ScheduledCork) XXX_Size() int {
	return m.Size()
}
func (m *ScheduledCork) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduledCork.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduledCork proto.InternalMessageInfo

func (m *ScheduledCork) GetCork() *Cork {
	if m != nil {
		return m.Cork
	}
	return nil
}

func (m *ScheduledCork) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *ScheduledCork) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

type CellarIDSet struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (m *CellarIDSet) Reset()         { *m = CellarIDSet{} }
func (m *CellarIDSet) String() string { return proto.CompactTextString(m) }
func (*CellarIDSet) ProtoMessage()    {}
func (*CellarIDSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_79882ab39b78d896, []int{3}
}
func (m *CellarIDSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CellarIDSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CellarIDSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CellarIDSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellarIDSet.Merge(m, src)
}
func (m *CellarIDSet) XXX_Size() int {
	return m.Size()
}
func (m *CellarIDSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CellarIDSet.DiscardUnknown(m)
}

var xxx_messageInfo_CellarIDSet proto.InternalMessageInfo

func (m *CellarIDSet) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*Cork)(nil), "cork.v1.Cork")
	proto.RegisterType((*ValidatorCork)(nil), "cork.v1.ValidatorCork")
	proto.RegisterType((*ScheduledCork)(nil), "cork.v1.ScheduledCork")
	proto.RegisterType((*CellarIDSet)(nil), "cork.v1.CellarIDSet")
}

func init() { proto.RegisterFile("cork/v1/cork.proto", fileDescriptor_79882ab39b78d896) }

var fileDescriptor_79882ab39b78d896 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x3b, 0x7f, 0xcb, 0x2f, 0x9d, 0xb6, 0x20, 0x91, 0x62, 0x15, 0x89, 0x6d, 0x57, 0x5d,
	0x48, 0x86, 0xb6, 0xe0, 0x5e, 0x23, 0xa2, 0x3b, 0x49, 0xc1, 0x85, 0x9b, 0x30, 0x9d, 0x19, 0x92,
	0xd8, 0x49, 0x6f, 0x99, 0x99, 0x06, 0xfb, 0x16, 0x3e, 0x96, 0xcb, 0x2e, 0x5d, 0x4a, 0xf3, 0x22,
	0x92, 0x49, 0x4a, 0xd1, 0x95, 0xab, 0xdc, 0x9c, 0x6f, 0xee, 0x39, 0x07, 0x2e, 0x76, 0x18, 0xa8,
	0x05, 0xc9, 0xc6, 0xa4, 0xf8, 0x7a, 0x2b, 0x05, 0x06, 0x9c, 0x23, 0x3b, 0x67, 0xe3, 0xf3, 0x33,
	0x06, 0x3a, 0x05, 0x1d, 0x5a, 0x99, 0x94, 0x3f, 0xe5, 0x9b, 0xa1, 0xc2, 0x0d, 0x1f, 0xd4, 0xc2,
	0x99, 0xe0, 0xae, 0x58, 0x32, 0xe0, 0x82, 0x87, 0x0c, 0x96, 0x46, 0x51, 0x66, 0x42, 0x46, 0xa5,
	0xec, 0xa1, 0x3e, 0x1a, 0xb5, 0x83, 0x93, 0x0a, 0xfa, 0x15, 0xf3, 0xa9, 0x94, 0xce, 0x35, 0x3e,
	0x35, 0x54, 0x45, 0xc2, 0x1c, 0x56, 0x28, 0xe7, 0x4a, 0x68, 0xdd, 0xfb, 0xd7, 0x47, 0xa3, 0x66,
	0xd0, 0x2d, 0xf1, 0x7e, 0xe9, 0xa6, 0x84, 0xc3, 0x27, 0xdc, 0x79, 0xa6, 0x32, 0xe1, 0xd4, 0x80,
	0xb2, 0xe1, 0x03, 0xdc, 0x28, 0xaa, 0xda, 0xac, 0xd6, 0xa4, 0xe3, 0x55, 0xbd, 0xbd, 0x02, 0x06,
	0x16, 0x39, 0x17, 0xb8, 0x99, 0xed, 0x77, 0x2a, 0xf7, 0x83, 0x30, 0xd4, 0xb8, 0x33, 0x63, 0xb1,
	0xe0, 0x6b, 0x59, 0x54, 0xfc, 0x9b, 0xe3, 0x00, 0xb7, 0xe7, 0x12, 0xd8, 0x22, 0x8c, 0x45, 0x12,
	0xc5, 0xc6, 0x9a, 0x36, 0x82, 0x96, 0xd5, 0x1e, 0xac, 0xf4, 0x33, 0xb4, 0xfe, 0x3b, 0xf4, 0x12,
	0xb7, 0x7c, 0x21, 0x25, 0x55, 0x8f, 0x77, 0x33, 0x61, 0x9c, 0x63, 0x5c, 0x4f, 0xb8, 0xee, 0xa1,
	0x7e, 0x7d, 0xd4, 0x0c, 0x8a, 0xf1, 0xf6, 0xfe, 0x63, 0xe7, 0xa2, 0xed, 0xce, 0x45, 0x5f, 0x3b,
	0x17, 0xbd, 0xe7, 0x6e, 0x6d, 0x9b, 0xbb, 0xb5, 0xcf, 0xdc, 0xad, 0xbd, 0x5c, 0x45, 0x89, 0x89,
	0xd7, 0x73, 0x8f, 0x41, 0x4a, 0x56, 0x22, 0x8a, 0x36, 0xaf, 0x19, 0xd1, 0x90, 0xa6, 0x42, 0x26,
	0x42, 0x91, 0x6c, 0x4a, 0xde, 0xec, 0x1d, 0x89, 0xd9, 0xac, 0x84, 0x9e, 0xff, 0xb7, 0xa7, 0x9a,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x53, 0xd2, 0xb5, 0x67, 0xe4, 0x01, 0x00, 0x00,
}

func (m *Cork) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cork) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Cork) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TargetContractAddress) > 0 {
		i -= len(m.TargetContractAddress)
		copy(dAtA[i:], m.TargetContractAddress)
		i = encodeVarintCork(dAtA, i, uint64(len(m.TargetContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.EncodedContractCall) > 0 {
		i -= len(m.EncodedContractCall)
		copy(dAtA[i:], m.EncodedContractCall)
		i = encodeVarintCork(dAtA, i, uint64(len(m.EncodedContractCall)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorCork) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorCork) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorCork) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintCork(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Cork != nil {
		{
			size, err := m.Cork.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCork(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ScheduledCork) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduledCork) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduledCork) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintCork(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintCork(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.Cork != nil {
		{
			size, err := m.Cork.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCork(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CellarIDSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CellarIDSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CellarIDSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ids) > 0 {
		for iNdEx := len(m.Ids) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Ids[iNdEx])
			copy(dAtA[i:], m.Ids[iNdEx])
			i = encodeVarintCork(dAtA, i, uint64(len(m.Ids[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCork(dAtA []byte, offset int, v uint64) int {
	offset -= sovCork(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Cork) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EncodedContractCall)
	if l > 0 {
		n += 1 + l + sovCork(uint64(l))
	}
	l = len(m.TargetContractAddress)
	if l > 0 {
		n += 1 + l + sovCork(uint64(l))
	}
	return n
}

func (m *ValidatorCork) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cork != nil {
		l = m.Cork.Size()
		n += 1 + l + sovCork(uint64(l))
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovCork(uint64(l))
	}
	return n
}

func (m *ScheduledCork) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cork != nil {
		l = m.Cork.Size()
		n += 1 + l + sovCork(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovCork(uint64(m.BlockHeight))
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovCork(uint64(l))
	}
	return n
}

func (m *CellarIDSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ids) > 0 {
		for _, s := range m.Ids {
			l = len(s)
			n += 1 + l + sovCork(uint64(l))
		}
	}
	return n
}

func sovCork(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCork(x uint64) (n int) {
	return sovCork(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Cork) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCork
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
			return fmt.Errorf("proto: Cork: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cork: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncodedContractCall", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncodedContractCall = append(m.EncodedContractCall[:0], dAtA[iNdEx:postIndex]...)
			if m.EncodedContractCall == nil {
				m.EncodedContractCall = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCork
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
func (m *ValidatorCork) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCork
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
			return fmt.Errorf("proto: ValidatorCork: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorCork: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cork", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cork == nil {
				m.Cork = &Cork{}
			}
			if err := m.Cork.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCork
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
func (m *ScheduledCork) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCork
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
			return fmt.Errorf("proto: ScheduledCork: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduledCork: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cork", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cork == nil {
				m.Cork = &Cork{}
			}
			if err := m.Cork.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCork
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
func (m *CellarIDSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCork
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
			return fmt.Errorf("proto: CellarIDSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CellarIDSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ids", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCork
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
				return ErrInvalidLengthCork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ids = append(m.Ids, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCork
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
func skipCork(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCork
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
					return 0, ErrIntOverflowCork
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
					return 0, ErrIntOverflowCork
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
				return 0, ErrInvalidLengthCork
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCork
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCork
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCork        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCork          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCork = fmt.Errorf("proto: unexpected end of group")
)
