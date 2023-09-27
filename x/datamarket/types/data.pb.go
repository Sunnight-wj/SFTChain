// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sft/datamarket/v1/data.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type DataSet struct {
	Urls []string `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (m *DataSet) Reset()         { *m = DataSet{} }
func (m *DataSet) String() string { return proto.CompactTextString(m) }
func (*DataSet) ProtoMessage()    {}
func (*DataSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_018cbeb6c89e83a6, []int{0}
}
func (m *DataSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSet.Merge(m, src)
}
func (m *DataSet) XXX_Size() int {
	return m.Size()
}
func (m *DataSet) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSet.DiscardUnknown(m)
}

var xxx_messageInfo_DataSet proto.InternalMessageInfo

func (m *DataSet) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

type Classes struct {
	Classes []string `protobuf:"bytes,1,rep,name=classes,proto3" json:"classes,omitempty"`
}

func (m *Classes) Reset()         { *m = Classes{} }
func (m *Classes) String() string { return proto.CompactTextString(m) }
func (*Classes) ProtoMessage()    {}
func (*Classes) Descriptor() ([]byte, []int) {
	return fileDescriptor_018cbeb6c89e83a6, []int{1}
}
func (m *Classes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Classes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Classes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Classes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Classes.Merge(m, src)
}
func (m *Classes) XXX_Size() int {
	return m.Size()
}
func (m *Classes) XXX_DiscardUnknown() {
	xxx_messageInfo_Classes.DiscardUnknown(m)
}

var xxx_messageInfo_Classes proto.InternalMessageInfo

func (m *Classes) GetClasses() []string {
	if m != nil {
		return m.Classes
	}
	return nil
}

type UploaderSet struct {
	UploaderSet []string `protobuf:"bytes,1,rep,name=uploaderSet,proto3" json:"uploaderSet,omitempty"`
}

func (m *UploaderSet) Reset()         { *m = UploaderSet{} }
func (m *UploaderSet) String() string { return proto.CompactTextString(m) }
func (*UploaderSet) ProtoMessage()    {}
func (*UploaderSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_018cbeb6c89e83a6, []int{2}
}
func (m *UploaderSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UploaderSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UploaderSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UploaderSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploaderSet.Merge(m, src)
}
func (m *UploaderSet) XXX_Size() int {
	return m.Size()
}
func (m *UploaderSet) XXX_DiscardUnknown() {
	xxx_messageInfo_UploaderSet.DiscardUnknown(m)
}

var xxx_messageInfo_UploaderSet proto.InternalMessageInfo

func (m *UploaderSet) GetUploaderSet() []string {
	if m != nil {
		return m.UploaderSet
	}
	return nil
}

func init() {
	proto.RegisterType((*DataSet)(nil), "sft.datamarket.v1.DataSet")
	proto.RegisterType((*Classes)(nil), "sft.datamarket.v1.Classes")
	proto.RegisterType((*UploaderSet)(nil), "sft.datamarket.v1.UploaderSet")
}

func init() { proto.RegisterFile("sft/datamarket/v1/data.proto", fileDescriptor_018cbeb6c89e83a6) }

var fileDescriptor_018cbeb6c89e83a6 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x4e, 0x2b, 0xd1,
	0x4f, 0x49, 0x2c, 0x49, 0xcc, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0xd1, 0x2f, 0x33, 0x04, 0xf3, 0xf4,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x8b, 0xd3, 0x4a, 0xf4, 0x10, 0xb2, 0x7a, 0x65, 0x86,
	0x4a, 0xb2, 0x5c, 0xec, 0x2e, 0x89, 0x25, 0x89, 0xc1, 0xa9, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0xa5,
	0x45, 0x39, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x32, 0x17, 0xbb,
	0x73, 0x4e, 0x62, 0x71, 0x71, 0x6a, 0xb1, 0x90, 0x04, 0x17, 0x7b, 0x32, 0x84, 0x09, 0x55, 0x01,
	0xe3, 0x2a, 0xe9, 0x73, 0x71, 0x87, 0x16, 0xe4, 0xe4, 0x27, 0xa6, 0xa4, 0x16, 0x81, 0xcc, 0x51,
	0xe0, 0xe2, 0x2e, 0x45, 0x70, 0xa1, 0x8a, 0x91, 0x85, 0x9c, 0x7c, 0x4e, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63,
	0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x28, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f,
	0x57, 0xdf, 0x39, 0xbf, 0x38, 0x37, 0xbf, 0xd8, 0x39, 0x3f, 0xaf, 0xa4, 0x28, 0x31, 0xb9, 0xa4,
	0x58, 0x3f, 0xab, 0x34, 0x2f, 0x5f, 0xbf, 0x02, 0xd9, 0x77, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49,
	0x6c, 0x60, 0xcf, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xcb, 0x8e, 0xc5, 0xfc, 0x00,
	0x00, 0x00,
}

func (m *DataSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Urls) > 0 {
		for iNdEx := len(m.Urls) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Urls[iNdEx])
			copy(dAtA[i:], m.Urls[iNdEx])
			i = encodeVarintData(dAtA, i, uint64(len(m.Urls[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Classes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Classes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Classes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Classes) > 0 {
		for iNdEx := len(m.Classes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Classes[iNdEx])
			copy(dAtA[i:], m.Classes[iNdEx])
			i = encodeVarintData(dAtA, i, uint64(len(m.Classes[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UploaderSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploaderSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UploaderSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UploaderSet) > 0 {
		for iNdEx := len(m.UploaderSet) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.UploaderSet[iNdEx])
			copy(dAtA[i:], m.UploaderSet[iNdEx])
			i = encodeVarintData(dAtA, i, uint64(len(m.UploaderSet[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintData(dAtA []byte, offset int, v uint64) int {
	offset -= sovData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Urls) > 0 {
		for _, s := range m.Urls {
			l = len(s)
			n += 1 + l + sovData(uint64(l))
		}
	}
	return n
}

func (m *Classes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Classes) > 0 {
		for _, s := range m.Classes {
			l = len(s)
			n += 1 + l + sovData(uint64(l))
		}
	}
	return n
}

func (m *UploaderSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.UploaderSet) > 0 {
		for _, s := range m.UploaderSet {
			l = len(s)
			n += 1 + l + sovData(uint64(l))
		}
	}
	return n
}

func sovData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozData(x uint64) (n int) {
	return sovData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
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
			return fmt.Errorf("proto: DataSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Urls", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Urls = append(m.Urls, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthData
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
func (m *Classes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
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
			return fmt.Errorf("proto: Classes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Classes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Classes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Classes = append(m.Classes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthData
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
func (m *UploaderSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
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
			return fmt.Errorf("proto: UploaderSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploaderSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UploaderSet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UploaderSet = append(m.UploaderSet, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthData
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
func skipData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowData
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
					return 0, ErrIntOverflowData
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
					return 0, ErrIntOverflowData
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
				return 0, ErrInvalidLengthData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupData = fmt.Errorf("proto: unexpected end of group")
)
