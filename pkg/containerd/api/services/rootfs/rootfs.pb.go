// Code generated by protoc-gen-gogo.
// source: github.com/containerd/containerd/api/services/rootfs/rootfs.proto
// DO NOT EDIT!

/*
	Package rootfs is a generated protocol buffer package.

	It is generated from these files:
		github.com/containerd/containerd/api/services/rootfs/rootfs.proto

	It has these top-level messages:
		UnpackRequest
		UnpackResponse
		PrepareRequest
		MountsRequest
		MountResponse
*/
package rootfs

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import containerd_v1_types "github.com/containerd/containerd/api/types/mount"
import containerd_v1_types1 "github.com/containerd/containerd/api/types/descriptor"

import github_com_opencontainers_go_digest "github.com/opencontainers/go-digest"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

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

type UnpackRequest struct {
	Layers []*containerd_v1_types1.Descriptor `protobuf:"bytes,1,rep,name=layers" json:"layers,omitempty"`
}

func (m *UnpackRequest) Reset()                    { *m = UnpackRequest{} }
func (*UnpackRequest) ProtoMessage()               {}
func (*UnpackRequest) Descriptor() ([]byte, []int) { return fileDescriptorRootfs, []int{0} }

type UnpackResponse struct {
	ChainID github_com_opencontainers_go_digest.Digest `protobuf:"bytes,1,opt,name=chainid,proto3,customtype=github.com/opencontainers/go-digest.Digest" json:"chainid"`
}

func (m *UnpackResponse) Reset()                    { *m = UnpackResponse{} }
func (*UnpackResponse) ProtoMessage()               {}
func (*UnpackResponse) Descriptor() ([]byte, []int) { return fileDescriptorRootfs, []int{1} }

type PrepareRequest struct {
	Name     string                                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ChainID  github_com_opencontainers_go_digest.Digest `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3,customtype=github.com/opencontainers/go-digest.Digest" json:"chain_id"`
	Readonly bool                                       `protobuf:"varint,3,opt,name=readonly,proto3" json:"readonly,omitempty"`
}

func (m *PrepareRequest) Reset()                    { *m = PrepareRequest{} }
func (*PrepareRequest) ProtoMessage()               {}
func (*PrepareRequest) Descriptor() ([]byte, []int) { return fileDescriptorRootfs, []int{2} }

type MountsRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *MountsRequest) Reset()                    { *m = MountsRequest{} }
func (*MountsRequest) ProtoMessage()               {}
func (*MountsRequest) Descriptor() ([]byte, []int) { return fileDescriptorRootfs, []int{3} }

type MountResponse struct {
	Mounts []*containerd_v1_types.Mount `protobuf:"bytes,1,rep,name=mounts" json:"mounts,omitempty"`
}

func (m *MountResponse) Reset()                    { *m = MountResponse{} }
func (*MountResponse) ProtoMessage()               {}
func (*MountResponse) Descriptor() ([]byte, []int) { return fileDescriptorRootfs, []int{4} }

func init() {
	proto.RegisterType((*UnpackRequest)(nil), "containerd.v1.UnpackRequest")
	proto.RegisterType((*UnpackResponse)(nil), "containerd.v1.UnpackResponse")
	proto.RegisterType((*PrepareRequest)(nil), "containerd.v1.PrepareRequest")
	proto.RegisterType((*MountsRequest)(nil), "containerd.v1.MountsRequest")
	proto.RegisterType((*MountResponse)(nil), "containerd.v1.MountResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RootFS service

type RootFSClient interface {
	Unpack(ctx context.Context, in *UnpackRequest, opts ...grpc.CallOption) (*UnpackResponse, error)
	Prepare(ctx context.Context, in *PrepareRequest, opts ...grpc.CallOption) (*MountResponse, error)
	Mounts(ctx context.Context, in *MountsRequest, opts ...grpc.CallOption) (*MountResponse, error)
}

type rootFSClient struct {
	cc *grpc.ClientConn
}

func NewRootFSClient(cc *grpc.ClientConn) RootFSClient {
	return &rootFSClient{cc}
}

func (c *rootFSClient) Unpack(ctx context.Context, in *UnpackRequest, opts ...grpc.CallOption) (*UnpackResponse, error) {
	out := new(UnpackResponse)
	err := grpc.Invoke(ctx, "/containerd.v1.RootFS/Unpack", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rootFSClient) Prepare(ctx context.Context, in *PrepareRequest, opts ...grpc.CallOption) (*MountResponse, error) {
	out := new(MountResponse)
	err := grpc.Invoke(ctx, "/containerd.v1.RootFS/Prepare", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rootFSClient) Mounts(ctx context.Context, in *MountsRequest, opts ...grpc.CallOption) (*MountResponse, error) {
	out := new(MountResponse)
	err := grpc.Invoke(ctx, "/containerd.v1.RootFS/Mounts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RootFS service

type RootFSServer interface {
	Unpack(context.Context, *UnpackRequest) (*UnpackResponse, error)
	Prepare(context.Context, *PrepareRequest) (*MountResponse, error)
	Mounts(context.Context, *MountsRequest) (*MountResponse, error)
}

func RegisterRootFSServer(s *grpc.Server, srv RootFSServer) {
	s.RegisterService(&_RootFS_serviceDesc, srv)
}

func _RootFS_Unpack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnpackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootFSServer).Unpack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.v1.RootFS/Unpack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootFSServer).Unpack(ctx, req.(*UnpackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RootFS_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootFSServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.v1.RootFS/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootFSServer).Prepare(ctx, req.(*PrepareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RootFS_Mounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootFSServer).Mounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.v1.RootFS/Mounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootFSServer).Mounts(ctx, req.(*MountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RootFS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.v1.RootFS",
	HandlerType: (*RootFSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unpack",
			Handler:    _RootFS_Unpack_Handler,
		},
		{
			MethodName: "Prepare",
			Handler:    _RootFS_Prepare_Handler,
		},
		{
			MethodName: "Mounts",
			Handler:    _RootFS_Mounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/containerd/containerd/api/services/rootfs/rootfs.proto",
}

func (m *UnpackRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnpackRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Layers) > 0 {
		for _, msg := range m.Layers {
			dAtA[i] = 0xa
			i++
			i = encodeVarintRootfs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *UnpackResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnpackResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ChainID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRootfs(dAtA, i, uint64(len(m.ChainID)))
		i += copy(dAtA[i:], m.ChainID)
	}
	return i, nil
}

func (m *PrepareRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrepareRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRootfs(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.ChainID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRootfs(dAtA, i, uint64(len(m.ChainID)))
		i += copy(dAtA[i:], m.ChainID)
	}
	if m.Readonly {
		dAtA[i] = 0x18
		i++
		if m.Readonly {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *MountsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MountsRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRootfs(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *MountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MountResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Mounts) > 0 {
		for _, msg := range m.Mounts {
			dAtA[i] = 0xa
			i++
			i = encodeVarintRootfs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Rootfs(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Rootfs(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRootfs(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UnpackRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.Layers) > 0 {
		for _, e := range m.Layers {
			l = e.Size()
			n += 1 + l + sovRootfs(uint64(l))
		}
	}
	return n
}

func (m *UnpackResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovRootfs(uint64(l))
	}
	return n
}

func (m *PrepareRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRootfs(uint64(l))
	}
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovRootfs(uint64(l))
	}
	if m.Readonly {
		n += 2
	}
	return n
}

func (m *MountsRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRootfs(uint64(l))
	}
	return n
}

func (m *MountResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Mounts) > 0 {
		for _, e := range m.Mounts {
			l = e.Size()
			n += 1 + l + sovRootfs(uint64(l))
		}
	}
	return n
}

func sovRootfs(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRootfs(x uint64) (n int) {
	return sovRootfs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *UnpackRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UnpackRequest{`,
		`Layers:` + strings.Replace(fmt.Sprintf("%v", this.Layers), "Descriptor", "containerd_v1_types1.Descriptor", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UnpackResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UnpackResponse{`,
		`ChainID:` + fmt.Sprintf("%v", this.ChainID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PrepareRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PrepareRequest{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`ChainID:` + fmt.Sprintf("%v", this.ChainID) + `,`,
		`Readonly:` + fmt.Sprintf("%v", this.Readonly) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MountsRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MountsRequest{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MountResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MountResponse{`,
		`Mounts:` + strings.Replace(fmt.Sprintf("%v", this.Mounts), "Mount", "containerd_v1_types.Mount", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRootfs(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *UnpackRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRootfs
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
			return fmt.Errorf("proto: UnpackRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnpackRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Layers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootfs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRootfs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Layers = append(m.Layers, &containerd_v1_types1.Descriptor{})
			if err := m.Layers[len(m.Layers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRootfs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRootfs
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
func (m *UnpackResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRootfs
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
			return fmt.Errorf("proto: UnpackResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnpackResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootfs
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
				return ErrInvalidLengthRootfs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = github_com_opencontainers_go_digest.Digest(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRootfs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRootfs
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
func (m *PrepareRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRootfs
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
			return fmt.Errorf("proto: PrepareRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrepareRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootfs
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
				return ErrInvalidLengthRootfs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootfs
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
				return ErrInvalidLengthRootfs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = github_com_opencontainers_go_digest.Digest(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Readonly", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootfs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Readonly = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRootfs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRootfs
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
func (m *MountsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRootfs
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
			return fmt.Errorf("proto: MountsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MountsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootfs
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
				return ErrInvalidLengthRootfs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRootfs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRootfs
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
func (m *MountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRootfs
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
			return fmt.Errorf("proto: MountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRootfs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRootfs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mounts = append(m.Mounts, &containerd_v1_types.Mount{})
			if err := m.Mounts[len(m.Mounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRootfs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRootfs
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
func skipRootfs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRootfs
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
					return 0, ErrIntOverflowRootfs
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
					return 0, ErrIntOverflowRootfs
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
				return 0, ErrInvalidLengthRootfs
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRootfs
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
				next, err := skipRootfs(dAtA[start:])
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
	ErrInvalidLengthRootfs = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRootfs   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/services/rootfs/rootfs.proto", fileDescriptorRootfs)
}

var fileDescriptorRootfs = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x52, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0xed, 0xf8, 0x24, 0xad, 0x23, 0x7d, 0x8b, 0xc1, 0x45, 0x08, 0x9a, 0x94, 0xb8, 0x29, 0x82,
	0x09, 0xd6, 0x85, 0x1b, 0x5d, 0xf8, 0x5e, 0x2c, 0xbe, 0x85, 0x20, 0x11, 0xd1, 0x9d, 0x4c, 0x93,
	0x31, 0x1d, 0x6c, 0xe7, 0x8e, 0x33, 0xd3, 0x42, 0x77, 0xfe, 0x0e, 0x7f, 0x51, 0x97, 0x2e, 0x45,
	0xb0, 0xd8, 0xfc, 0x12, 0x69, 0xbe, 0x6c, 0x4a, 0x95, 0x0a, 0x6f, 0x93, 0xb9, 0x61, 0xce, 0x39,
	0xf7, 0xdc, 0x73, 0x07, 0x3f, 0xcf, 0xb8, 0x99, 0x2e, 0x26, 0x41, 0x02, 0xf3, 0x30, 0x01, 0x61,
	0x28, 0x17, 0x4c, 0xa5, 0xfb, 0x25, 0x95, 0x3c, 0xd4, 0x4c, 0x2d, 0x79, 0xc2, 0x74, 0xa8, 0x00,
	0xcc, 0xc7, 0xfa, 0x08, 0xa4, 0x02, 0x03, 0xa4, 0xff, 0x07, 0x1c, 0x2c, 0x1f, 0x39, 0x77, 0x32,
	0xc8, 0xa0, 0xb8, 0x09, 0x77, 0x55, 0x09, 0x72, 0x9e, 0x9e, 0xd4, 0xc7, 0xac, 0x24, 0xd3, 0xe1,
	0x1c, 0x16, 0xc2, 0x94, 0xdf, 0x8a, 0x3d, 0xfe, 0x0f, 0x76, 0xca, 0x74, 0xa2, 0xb8, 0x34, 0xa0,
	0xf6, 0xca, 0x52, 0xc7, 0x7f, 0x89, 0xfb, 0x6f, 0x85, 0xa4, 0xc9, 0xa7, 0x98, 0x7d, 0x5e, 0x30,
	0x6d, 0xc8, 0x13, 0x6c, 0xcd, 0xe8, 0x8a, 0x29, 0x6d, 0xa3, 0xc1, 0xd9, 0xf0, 0xf6, 0xc8, 0x0b,
	0x5a, 0xc3, 0x04, 0x85, 0x64, 0x10, 0x35, 0x3a, 0x71, 0x05, 0xf7, 0x39, 0x3e, 0xaf, 0x95, 0xb4,
	0x04, 0xa1, 0x19, 0x79, 0x87, 0xbb, 0xc9, 0x94, 0x72, 0xc1, 0x53, 0x1b, 0x0d, 0xd0, 0xf0, 0xd6,
	0xc5, 0xb3, 0xf5, 0xc6, 0xeb, 0xfc, 0xd8, 0x78, 0x0f, 0xf6, 0xcc, 0x83, 0x64, 0xa2, 0xe9, 0xa0,
	0xc3, 0x0c, 0x1e, 0xa6, 0x3c, 0x63, 0xda, 0x04, 0x51, 0x71, 0xe4, 0x1b, 0xaf, 0x7b, 0xb9, 0x13,
	0xb9, 0x8a, 0xe2, 0x5a, 0xcd, 0xff, 0x8a, 0xf0, 0xf9, 0x6b, 0xc5, 0x24, 0x55, 0xac, 0xb6, 0x4d,
	0xf0, 0x4d, 0x41, 0xe7, 0xac, 0x6c, 0x14, 0x17, 0x35, 0x79, 0x8f, 0x7b, 0x05, 0xe3, 0x03, 0x4f,
	0xed, 0x1b, 0xd7, 0x67, 0xe0, 0x2a, 0x25, 0x0e, 0xee, 0x29, 0x46, 0x53, 0x10, 0xb3, 0x95, 0x7d,
	0x36, 0x40, 0xc3, 0x5e, 0xdc, 0xfc, 0xfb, 0xf7, 0x71, 0xff, 0xd5, 0x6e, 0x51, 0xfa, 0x1f, 0xd6,
	0xfc, 0xcb, 0x0a, 0xd4, 0x64, 0x35, 0xc2, 0x56, 0xb1, 0xde, 0x3a, 0x76, 0xe7, 0x68, 0xec, 0x25,
	0xa7, 0x42, 0x8e, 0x7e, 0x22, 0x6c, 0xc5, 0x00, 0x66, 0xfc, 0x86, 0xbc, 0xc0, 0x56, 0x19, 0x3e,
	0xb9, 0x7b, 0x40, 0x6c, 0x6d, 0xd7, 0xb9, 0xf7, 0x97, 0xdb, 0xca, 0xc5, 0x18, 0x77, 0xab, 0x5c,
	0xc9, 0x21, 0xb2, 0x9d, 0xb7, 0x73, 0xd8, 0xa6, 0x3d, 0x4d, 0x84, 0xad, 0x32, 0x03, 0x72, 0x14,
	0xa7, 0x4f, 0x52, 0xb9, 0xb0, 0xd7, 0x5b, 0xb7, 0xf3, 0x7d, 0xeb, 0x76, 0xbe, 0xe4, 0x2e, 0x5a,
	0xe7, 0x2e, 0xfa, 0x96, 0xbb, 0xe8, 0x57, 0xee, 0xa2, 0x89, 0x55, 0x3c, 0xde, 0xc7, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x3c, 0xba, 0xae, 0xaa, 0xac, 0x03, 0x00, 0x00,
}
