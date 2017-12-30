// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel.proto

/*
Package go_micro_srv_vessel is a generated protocol buffer package.

It is generated from these files:
	proto/vessel.proto

It has these top-level messages:
	Vessel
	Specification
	Response
*/
package go_micro_srv_vessel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Vessel struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Capacity  int32  `protobuf:"varint,2,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight int32  `protobuf:"varint,3,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Available bool   `protobuf:"varint,5,opt,name=available" json:"available,omitempty"`
	OwnerId   string `protobuf:"bytes,6,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
}

func (m *Vessel) Reset()                    { *m = Vessel{} }
func (m *Vessel) String() string            { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()               {}
func (*Vessel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity  int32 `protobuf:"varint,1,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight int32 `protobuf:"varint,2,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
}

func (m *Specification) Reset()                    { *m = Specification{} }
func (m *Specification) String() string            { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()               {}
func (*Specification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel  *Vessel   `protobuf:"bytes,1,opt,name=vessel" json:"vessel,omitempty"`
	Vessels []*Vessel `protobuf:"bytes,2,rep,name=vessels" json:"vessels,omitempty"`
	Created bool      `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func init() {
	proto.RegisterType((*Vessel)(nil), "go.micro.srv.vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "go.micro.srv.vessel.Specification")
	proto.RegisterType((*Response)(nil), "go.micro.srv.vessel.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...grpc.CallOption) (*Response, error)
	Create(ctx context.Context, in *Vessel, opts ...grpc.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	cc *grpc.ClientConn
}

func NewVesselServiceClient(cc *grpc.ClientConn) VesselServiceClient {
	return &vesselServiceClient{cc}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/go.micro.srv.vessel.VesselService/FindAvailable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vesselServiceClient) Create(ctx context.Context, in *Vessel, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/go.micro.srv.vessel.VesselService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceServer interface {
	FindAvailable(context.Context, *Specification) (*Response, error)
	Create(context.Context, *Vessel) (*Response, error)
}

func RegisterVesselServiceServer(s *grpc.Server, srv VesselServiceServer) {
	s.RegisterService(&_VesselService_serviceDesc, srv)
}

func _VesselService_FindAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Specification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VesselServiceServer).FindAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.vessel.VesselService/FindAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VesselServiceServer).FindAvailable(ctx, req.(*Specification))
	}
	return interceptor(ctx, in, info, handler)
}

func _VesselService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vessel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VesselServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.vessel.VesselService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VesselServiceServer).Create(ctx, req.(*Vessel))
	}
	return interceptor(ctx, in, info, handler)
}

var _VesselService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.vessel.VesselService",
	HandlerType: (*VesselServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAvailable",
			Handler:    _VesselService_FindAvailable_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _VesselService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vessel.proto",
}

func init() { proto.RegisterFile("proto/vessel.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x75, 0x0b, 0x94, 0x32, 0x06, 0x0f, 0xe3, 0x65, 0x45, 0x49, 0x9a, 0x9e, 0x38, 0xd5, 0x04,
	0xe2, 0x07, 0x18, 0x13, 0x12, 0x3d, 0x16, 0xa3, 0x47, 0xb2, 0x6c, 0x47, 0xdc, 0x84, 0x76, 0x9b,
	0x6e, 0x53, 0xf0, 0x2f, 0xfc, 0x04, 0x7f, 0xc0, 0x7f, 0x34, 0x99, 0x0a, 0x06, 0x43, 0xe0, 0x36,
	0x6f, 0xe6, 0xcd, 0x9b, 0xb7, 0x2f, 0x0b, 0x58, 0x94, 0xb6, 0xb2, 0xb7, 0x35, 0x39, 0x47, 0xab,
	0x98, 0x01, 0x5e, 0x2e, 0x6d, 0x9c, 0x19, 0x5d, 0xda, 0xd8, 0x95, 0x75, 0xdc, 0x8c, 0xa2, 0x2f,
	0x01, 0xfe, 0x0b, 0x97, 0x78, 0x01, 0x9e, 0x49, 0xa5, 0x08, 0xc5, 0xa8, 0x97, 0x78, 0x26, 0xc5,
	0x01, 0x04, 0x5a, 0x15, 0x4a, 0x9b, 0xea, 0x43, 0x7a, 0xa1, 0x18, 0x75, 0x92, 0x1d, 0xc6, 0x21,
	0x40, 0xa6, 0x36, 0xf3, 0x35, 0x99, 0xe5, 0x7b, 0x25, 0x5b, 0x3c, 0xed, 0x65, 0x6a, 0xf3, 0xca,
	0x0d, 0x44, 0x68, 0xe7, 0x2a, 0x23, 0xd9, 0x66, 0x31, 0xae, 0xf1, 0x06, 0x7a, 0xaa, 0x56, 0x66,
	0xa5, 0x16, 0x2b, 0x92, 0x9d, 0x50, 0x8c, 0x82, 0xe4, 0xaf, 0x81, 0x57, 0x10, 0xd8, 0x75, 0x4e,
	0xe5, 0xdc, 0xa4, 0xd2, 0xe7, 0xad, 0x2e, 0xe3, 0xc7, 0x34, 0x7a, 0x82, 0xfe, 0xac, 0x20, 0x6d,
	0xde, 0x8c, 0x56, 0x95, 0xb1, 0xf9, 0x9e, 0x31, 0x71, 0xd4, 0x98, 0xf7, 0xcf, 0x58, 0xf4, 0x29,
	0x20, 0x48, 0xc8, 0x15, 0x36, 0x77, 0x84, 0x13, 0xf0, 0x9b, 0x14, 0x58, 0xe5, 0x7c, 0x7c, 0x1d,
	0x1f, 0x48, 0x28, 0x6e, 0xd2, 0x49, 0x7e, 0xa9, 0x78, 0x07, 0xdd, 0xa6, 0x72, 0xd2, 0x0b, 0x5b,
	0xa7, 0xb6, 0xb6, 0x5c, 0x94, 0xd0, 0xd5, 0x25, 0xa9, 0x8a, 0x52, 0x4e, 0x2b, 0x48, 0xb6, 0x70,
	0xfc, 0x2d, 0xa0, 0xdf, 0xb0, 0x67, 0x54, 0xd6, 0x46, 0x13, 0x3e, 0x43, 0x7f, 0x6a, 0xf2, 0xf4,
	0x7e, 0x17, 0x4e, 0x74, 0xf0, 0xc4, 0x5e, 0x28, 0x83, 0xe1, 0x41, 0xce, 0xf6, 0xad, 0xd1, 0x19,
	0x4e, 0xc1, 0x7f, 0xe0, 0x93, 0x78, 0xcc, 0xf1, 0x49, 0x9d, 0x85, 0xcf, 0xbf, 0x69, 0xf2, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x38, 0x4e, 0xf6, 0x12, 0x63, 0x02, 0x00, 0x00,
}
