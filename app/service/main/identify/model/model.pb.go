// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

/*
	Package model is a generated protocol buffer package.

	It is generated from these files:
		model.proto

	It has these top-level messages:
		IdentifyInfo
*/
package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// result of http identify info.
type IdentifyInfo struct {
	Mid     int64  `protobuf:"varint,1,opt,name=mid,proto3" json:"mid"`
	Csrf    string `protobuf:"bytes,2,opt,name=csrf,proto3" json:"csrfToken"`
	Expires int32  `protobuf:"varint,3,opt,name=expires,proto3" json:"expires"`
}

func (m *IdentifyInfo) Reset()                    { *m = IdentifyInfo{} }
func (m *IdentifyInfo) String() string            { return proto.CompactTextString(m) }
func (*IdentifyInfo) ProtoMessage()               {}
func (*IdentifyInfo) Descriptor() ([]byte, []int) { return fileDescriptorModel, []int{0} }

func (m *IdentifyInfo) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *IdentifyInfo) GetCsrf() string {
	if m != nil {
		return m.Csrf
	}
	return ""
}

func (m *IdentifyInfo) GetExpires() int32 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func init() {
	proto.RegisterType((*IdentifyInfo)(nil), "passport.service.identify.IdentifyInfo")
}
func (m *IdentifyInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdentifyInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Mid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintModel(dAtA, i, uint64(m.Mid))
	}
	if len(m.Csrf) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintModel(dAtA, i, uint64(len(m.Csrf)))
		i += copy(dAtA[i:], m.Csrf)
	}
	if m.Expires != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintModel(dAtA, i, uint64(m.Expires))
	}
	return i, nil
}

func encodeVarintModel(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *IdentifyInfo) Size() (n int) {
	var l int
	_ = l
	if m.Mid != 0 {
		n += 1 + sovModel(uint64(m.Mid))
	}
	l = len(m.Csrf)
	if l > 0 {
		n += 1 + l + sovModel(uint64(l))
	}
	if m.Expires != 0 {
		n += 1 + sovModel(uint64(m.Expires))
	}
	return n
}

func sovModel(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozModel(x uint64) (n int) {
	return sovModel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IdentifyInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IdentifyInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdentifyInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			m.Mid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Csrf", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Csrf = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			m.Expires = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expires |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModel
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
func skipModel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModel
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
					return 0, ErrIntOverflowModel
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
					return 0, ErrIntOverflowModel
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthModel
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowModel
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
				next, err := skipModel(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthModel = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModel   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("model.proto", fileDescriptorModel) }

var fileDescriptorModel = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2c, 0x48, 0x2c, 0x2e, 0x2e, 0xc8, 0x2f,
	0x2a, 0xd1, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x49, 0xcd, 0x2b, 0xc9,
	0x4c, 0xab, 0x94, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f,
	0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xeb, 0x48, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0x62,
	0x92, 0x52, 0x31, 0x17, 0x8f, 0x27, 0x54, 0xab, 0x67, 0x5e, 0x5a, 0xbe, 0x90, 0x24, 0x17, 0x73,
	0x6e, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb3, 0x13, 0xfb, 0xab, 0x7b, 0xf2, 0x20, 0x6e,
	0x10, 0x88, 0x10, 0x52, 0xe4, 0x62, 0x49, 0x2e, 0x2e, 0x4a, 0x93, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x74, 0xe2, 0x7d, 0x75, 0x4f, 0x9e, 0x13, 0xc4, 0x0f, 0xc9, 0xcf, 0x4e, 0xcd, 0x0b, 0x02, 0x4b,
	0x09, 0xa9, 0x72, 0xb1, 0xa7, 0x56, 0x14, 0x64, 0x16, 0xa5, 0x16, 0x4b, 0x30, 0x2b, 0x30, 0x6a,
	0xb0, 0x3a, 0x71, 0xbf, 0xba, 0x27, 0x0f, 0x13, 0x0a, 0x82, 0x31, 0x9c, 0xc4, 0x4f, 0x3c, 0x92,
	0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0xa2, 0x58,
	0xc1, 0xbe, 0x4b, 0x62, 0x03, 0x3b, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x73, 0x88, 0x08,
	0x23, 0xed, 0x00, 0x00, 0x00,
}
