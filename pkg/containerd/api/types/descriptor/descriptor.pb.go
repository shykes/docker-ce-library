// Code generated by protoc-gen-gogo.
// source: github.com/docker/containerd/api/types/descriptor/descriptor.proto
// DO NOT EDIT!

/*
	Package descriptor is a generated protocol buffer package.

	It is generated from these files:
		github.com/docker/containerd/api/types/descriptor/descriptor.proto

	It has these top-level messages:
		Descriptor
*/
package descriptor

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_opencontainers_go_digest "github.com/opencontainers/go-digest"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Descriptor struct {
	MediaType string                                     `protobuf:"bytes,1,opt,name=mediaType,proto3" json:"mediaType,omitempty"`
	Digest    github_com_opencontainers_go_digest.Digest `protobuf:"bytes,2,opt,name=digest,proto3,customtype=github.com/opencontainers/go-digest.Digest" json:"digest"`
	MediaSize int64                                      `protobuf:"varint,3,opt,name=mediaSize,proto3" json:"mediaSize,omitempty"`
}

func (m *Descriptor) Reset()                    { *m = Descriptor{} }
func (*Descriptor) ProtoMessage()               {}
func (*Descriptor) Descriptor() ([]byte, []int) { return fileDescriptorDescriptor, []int{0} }

func init() {
	proto.RegisterType((*Descriptor)(nil), "containerd.v1.types.Descriptor")
}
func (m *Descriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Descriptor) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.MediaType) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDescriptor(dAtA, i, uint64(len(m.MediaType)))
		i += copy(dAtA[i:], m.MediaType)
	}
	if len(m.Digest) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDescriptor(dAtA, i, uint64(len(m.Digest)))
		i += copy(dAtA[i:], m.Digest)
	}
	if m.MediaSize != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintDescriptor(dAtA, i, uint64(m.MediaSize))
	}
	return i, nil
}

func encodeFixed64Descriptor(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Descriptor(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDescriptor(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Descriptor) Size() (n int) {
	var l int
	_ = l
	l = len(m.MediaType)
	if l > 0 {
		n += 1 + l + sovDescriptor(uint64(l))
	}
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovDescriptor(uint64(l))
	}
	if m.MediaSize != 0 {
		n += 1 + sovDescriptor(uint64(m.MediaSize))
	}
	return n
}

func sovDescriptor(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDescriptor(x uint64) (n int) {
	return sovDescriptor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Descriptor) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Descriptor{`,
		`MediaType:` + fmt.Sprintf("%v", this.MediaType) + `,`,
		`Digest:` + fmt.Sprintf("%v", this.Digest) + `,`,
		`MediaSize:` + fmt.Sprintf("%v", this.MediaSize) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDescriptor(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Descriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDescriptor
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
			return fmt.Errorf("proto: Descriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Descriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MediaType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescriptor
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
				return ErrInvalidLengthDescriptor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MediaType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Digest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescriptor
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
				return ErrInvalidLengthDescriptor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Digest = github_com_opencontainers_go_digest.Digest(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MediaSize", wireType)
			}
			m.MediaSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescriptor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MediaSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDescriptor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDescriptor
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
func skipDescriptor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDescriptor
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
					return 0, ErrIntOverflowDescriptor
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
					return 0, ErrIntOverflowDescriptor
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
				return 0, ErrInvalidLengthDescriptor
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDescriptor
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
				next, err := skipDescriptor(dAtA[start:])
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
	ErrInvalidLengthDescriptor = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDescriptor   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/docker/containerd/api/types/descriptor/descriptor.proto", fileDescriptorDescriptor)
}

var fileDescriptorDescriptor = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xc9, 0x4f, 0xce, 0x4e, 0x2d, 0xd2, 0x4f, 0xce,
	0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d, 0x4a, 0xd1, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0xa9, 0x2c,
	0x48, 0x2d, 0xd6, 0x4f, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0x2f, 0x42, 0x62, 0xea,
	0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0x23, 0x74, 0xe8, 0x95, 0x19, 0xea, 0x81, 0x35, 0x48,
	0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0xe5, 0xf5, 0x41, 0x2c, 0x88, 0x52, 0xa5, 0x29, 0x8c, 0x5c,
	0x5c, 0x2e, 0x70, 0xfd, 0x42, 0x32, 0x5c, 0x9c, 0xb9, 0xa9, 0x29, 0x99, 0x89, 0x21, 0x95, 0x05,
	0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x08, 0x01, 0x21, 0x2f, 0x2e, 0xb6, 0x94, 0xcc,
	0xf4, 0xd4, 0xe2, 0x12, 0x09, 0x26, 0x90, 0x94, 0x93, 0xd1, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee,
	0xc9, 0x6b, 0x21, 0xb9, 0x39, 0xbf, 0x20, 0x35, 0x0f, 0x6e, 0x7d, 0xb1, 0x7e, 0x7a, 0xbe, 0x2e,
	0x44, 0x8b, 0x9e, 0x0b, 0x98, 0x0a, 0x82, 0x9a, 0x00, 0xb7, 0x29, 0x38, 0xb3, 0x2a, 0x55, 0x82,
	0x59, 0x81, 0x51, 0x83, 0x39, 0x08, 0x21, 0xe0, 0x24, 0x71, 0xe2, 0xa1, 0x1c, 0xc3, 0x8d, 0x87,
	0x72, 0x0c, 0x0d, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1,
	0x23, 0x39, 0xc6, 0x24, 0x36, 0xb0, 0xbb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x3c,
	0x86, 0x2d, 0x28, 0x01, 0x00, 0x00,
}
