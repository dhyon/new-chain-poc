// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/order_entry.proto

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

type OrderEntry struct {
	Price             int32    `protobuf:"varint,1,opt,name=price,proto3" json:"price,omitempty"`
	Quantity          int32    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	AllocationCreator []string `protobuf:"bytes,3,rep,name=allocationCreator,proto3" json:"allocationCreator,omitempty"`
	Allocation        []int32  `protobuf:"varint,4,rep,packed,name=allocation,proto3" json:"allocation,omitempty"`
}

func (m *OrderEntry) Reset()         { *m = OrderEntry{} }
func (m *OrderEntry) String() string { return proto.CompactTextString(m) }
func (*OrderEntry) ProtoMessage()    {}
func (*OrderEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_25878922effe12c2, []int{0}
}
func (m *OrderEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderEntry.Merge(m, src)
}
func (m *OrderEntry) XXX_Size() int {
	return m.Size()
}
func (m *OrderEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderEntry.DiscardUnknown(m)
}

var xxx_messageInfo_OrderEntry proto.InternalMessageInfo

func (m *OrderEntry) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderEntry) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *OrderEntry) GetAllocationCreator() []string {
	if m != nil {
		return m.AllocationCreator
	}
	return nil
}

func (m *OrderEntry) GetAllocation() []int32 {
	if m != nil {
		return m.Allocation
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderEntry)(nil), "codchen.newchainpoc.dex.OrderEntry")
}

func init() { proto.RegisterFile("dex/order_entry.proto", fileDescriptor_25878922effe12c2) }

var fileDescriptor_25878922effe12c2 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xcf, 0xbf, 0x4a, 0xc4, 0x30,
	0x1c, 0x07, 0xf0, 0xc6, 0x5a, 0xd1, 0x6c, 0x06, 0xc5, 0xe2, 0x10, 0x8a, 0x53, 0x87, 0xbb, 0x66,
	0xf0, 0x0d, 0x14, 0x5d, 0x85, 0x8e, 0x2e, 0x92, 0x4b, 0x7e, 0xd8, 0xc0, 0x99, 0x5f, 0x8c, 0xbf,
	0xe3, 0xda, 0x77, 0x70, 0xf0, 0xb1, 0x1c, 0x3b, 0x3a, 0x4a, 0xfb, 0x22, 0xd2, 0x2a, 0x2a, 0xdc,
	0xf8, 0xfd, 0xb3, 0x7c, 0xf8, 0xa9, 0x85, 0x56, 0x61, 0xb4, 0x10, 0x1f, 0xc0, 0x53, 0xec, 0xaa,
	0x10, 0x91, 0x50, 0x9c, 0x19, 0xb4, 0xa6, 0x01, 0x5f, 0x79, 0xd8, 0x9a, 0x46, 0x3b, 0x1f, 0xd0,
	0x54, 0x16, 0xda, 0x8b, 0x57, 0xc6, 0xf9, 0xdd, 0x74, 0xbf, 0x99, 0xde, 0xe2, 0x84, 0x67, 0x21,
	0x3a, 0x03, 0x39, 0x2b, 0x58, 0x99, 0xd5, 0xdf, 0x41, 0x9c, 0xf3, 0xc3, 0xe7, 0x8d, 0xf6, 0xe4,
	0xa8, 0xcb, 0xf7, 0xe6, 0xe1, 0x37, 0x8b, 0x05, 0x3f, 0xd6, 0xeb, 0x35, 0x1a, 0x4d, 0x0e, 0xfd,
	0x75, 0x04, 0x4d, 0x18, 0xf3, 0xb4, 0x48, 0xcb, 0xa3, 0x7a, 0x77, 0x10, 0x92, 0xf3, 0xbf, 0x32,
	0xdf, 0x2f, 0xd2, 0x32, 0xab, 0xff, 0x35, 0x57, 0xb7, 0xef, 0x83, 0x64, 0xfd, 0x20, 0xd9, 0xe7,
	0x20, 0xd9, 0xdb, 0x28, 0x93, 0x7e, 0x94, 0xc9, 0xc7, 0x28, 0x93, 0xfb, 0xc5, 0xa3, 0xa3, 0x66,
	0xb3, 0xaa, 0x0c, 0x3e, 0xa9, 0x1f, 0x8c, 0xf2, 0xb0, 0x5d, 0xce, 0x9a, 0x65, 0x40, 0xa3, 0x5a,
	0x35, 0xd9, 0xa9, 0x0b, 0xf0, 0xb2, 0x3a, 0x98, 0xd9, 0x97, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xb2, 0x54, 0x34, 0x09, 0x0f, 0x01, 0x00, 0x00,
}

func (m *OrderEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Allocation) > 0 {
		dAtA2 := make([]byte, len(m.Allocation)*10)
		var j1 int
		for _, num1 := range m.Allocation {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintOrderEntry(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AllocationCreator) > 0 {
		for iNdEx := len(m.AllocationCreator) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllocationCreator[iNdEx])
			copy(dAtA[i:], m.AllocationCreator[iNdEx])
			i = encodeVarintOrderEntry(dAtA, i, uint64(len(m.AllocationCreator[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Quantity != 0 {
		i = encodeVarintOrderEntry(dAtA, i, uint64(m.Quantity))
		i--
		dAtA[i] = 0x10
	}
	if m.Price != 0 {
		i = encodeVarintOrderEntry(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrderEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrderEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Price != 0 {
		n += 1 + sovOrderEntry(uint64(m.Price))
	}
	if m.Quantity != 0 {
		n += 1 + sovOrderEntry(uint64(m.Quantity))
	}
	if len(m.AllocationCreator) > 0 {
		for _, s := range m.AllocationCreator {
			l = len(s)
			n += 1 + l + sovOrderEntry(uint64(l))
		}
	}
	if len(m.Allocation) > 0 {
		l = 0
		for _, e := range m.Allocation {
			l += sovOrderEntry(uint64(e))
		}
		n += 1 + sovOrderEntry(uint64(l)) + l
	}
	return n
}

func sovOrderEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrderEntry(x uint64) (n int) {
	return sovOrderEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrderEntry
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
			return fmt.Errorf("proto: OrderEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			m.Quantity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Quantity |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationCreator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderEntry
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
				return ErrInvalidLengthOrderEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllocationCreator = append(m.AllocationCreator, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOrderEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Allocation = append(m.Allocation, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOrderEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthOrderEntry
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthOrderEntry
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Allocation) == 0 {
					m.Allocation = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOrderEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Allocation = append(m.Allocation, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocation", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrderEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrderEntry
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
func skipOrderEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrderEntry
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
					return 0, ErrIntOverflowOrderEntry
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
					return 0, ErrIntOverflowOrderEntry
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
				return 0, ErrInvalidLengthOrderEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrderEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrderEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrderEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrderEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrderEntry = fmt.Errorf("proto: unexpected end of group")
)
