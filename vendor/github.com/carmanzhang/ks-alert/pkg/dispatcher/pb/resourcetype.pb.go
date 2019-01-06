// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resourcetype.proto

package pb

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

type ResourceType struct {
	ResourceTypeId    string            `protobuf:"bytes,1,opt,name=resource_type_id" json:"resource_type_id,omitempty"`
	ResourceTypeName  string            `protobuf:"bytes,2,opt,name=resource_type_name" json:"resource_type_name,omitempty"`
	ProductId         string            `protobuf:"bytes,3,opt,name=product_id" json:"product_id,omitempty"`
	MonitorCenterHost string            `protobuf:"bytes,4,opt,name=monitor_center_host" json:"monitor_center_host,omitempty"`
	MonitorCenterPort int32             `protobuf:"varint,5,opt,name=monitor_center_port" json:"monitor_center_port,omitempty"`
	ResourceUriTmpl   *ResourceUriTmpls `protobuf:"bytes,6,opt,name=resource_uri_tmpl" json:"resource_uri_tmpl,omitempty"`
	Enable            bool              `protobuf:"varint,7,opt,name=enable" json:"enable,omitempty"`
	Desc              string            `protobuf:"bytes,8,opt,name=desc" json:"desc,omitempty"`
}

func (m *ResourceType) Reset()                    { *m = ResourceType{} }
func (m *ResourceType) String() string            { return proto.CompactTextString(m) }
func (*ResourceType) ProtoMessage()               {}
func (*ResourceType) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *ResourceType) GetResourceTypeId() string {
	if m != nil {
		return m.ResourceTypeId
	}
	return ""
}

func (m *ResourceType) GetResourceTypeName() string {
	if m != nil {
		return m.ResourceTypeName
	}
	return ""
}

func (m *ResourceType) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *ResourceType) GetMonitorCenterHost() string {
	if m != nil {
		return m.MonitorCenterHost
	}
	return ""
}

func (m *ResourceType) GetMonitorCenterPort() int32 {
	if m != nil {
		return m.MonitorCenterPort
	}
	return 0
}

func (m *ResourceType) GetResourceUriTmpl() *ResourceUriTmpls {
	if m != nil {
		return m.ResourceUriTmpl
	}
	return nil
}

func (m *ResourceType) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *ResourceType) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type ResourceTypeID struct {
	ResourceTypeId string `protobuf:"bytes,1,opt,name=resource_type_id" json:"resource_type_id,omitempty"`
}

func (m *ResourceTypeID) Reset()                    { *m = ResourceTypeID{} }
func (m *ResourceTypeID) String() string            { return proto.CompactTextString(m) }
func (*ResourceTypeID) ProtoMessage()               {}
func (*ResourceTypeID) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *ResourceTypeID) GetResourceTypeId() string {
	if m != nil {
		return m.ResourceTypeId
	}
	return ""
}

type ResourceTypeResponse struct {
	ResourceTypeId *ResourceTypeID `protobuf:"bytes,1,opt,name=resource_type_id" json:"resource_type_id,omitempty"`
	ResourceType   *ResourceType   `protobuf:"bytes,2,opt,name=resource_type" json:"resource_type,omitempty"`
	Error          *Error          `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *ResourceTypeResponse) Reset()                    { *m = ResourceTypeResponse{} }
func (m *ResourceTypeResponse) String() string            { return proto.CompactTextString(m) }
func (*ResourceTypeResponse) ProtoMessage()               {}
func (*ResourceTypeResponse) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *ResourceTypeResponse) GetResourceTypeId() *ResourceTypeID {
	if m != nil {
		return m.ResourceTypeId
	}
	return nil
}

func (m *ResourceTypeResponse) GetResourceType() *ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return nil
}

func (m *ResourceTypeResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceType)(nil), "pb.ResourceType")
	proto.RegisterType((*ResourceTypeID)(nil), "pb.ResourceTypeID")
	proto.RegisterType((*ResourceTypeResponse)(nil), "pb.ResourceTypeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourceTypeHandler service

type ResourceTypeHandlerClient interface {
	// resource type
	CreateResourceType(ctx context.Context, in *ResourceType, opts ...grpc.CallOption) (*ResourceTypeResponse, error)
	DeleteResourceType(ctx context.Context, in *ResourceTypeID, opts ...grpc.CallOption) (*ResourceTypeResponse, error)
	UpdateResourceType(ctx context.Context, in *ResourceType, opts ...grpc.CallOption) (*ResourceTypeResponse, error)
	GetResourceType(ctx context.Context, in *ResourceTypeID, opts ...grpc.CallOption) (*ResourceTypeResponse, error)
}

type resourceTypeHandlerClient struct {
	cc *grpc.ClientConn
}

func NewResourceTypeHandlerClient(cc *grpc.ClientConn) ResourceTypeHandlerClient {
	return &resourceTypeHandlerClient{cc}
}

func (c *resourceTypeHandlerClient) CreateResourceType(ctx context.Context, in *ResourceType, opts ...grpc.CallOption) (*ResourceTypeResponse, error) {
	out := new(ResourceTypeResponse)
	err := grpc.Invoke(ctx, "/pb.ResourceTypeHandler/CreateResourceType", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceTypeHandlerClient) DeleteResourceType(ctx context.Context, in *ResourceTypeID, opts ...grpc.CallOption) (*ResourceTypeResponse, error) {
	out := new(ResourceTypeResponse)
	err := grpc.Invoke(ctx, "/pb.ResourceTypeHandler/DeleteResourceType", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceTypeHandlerClient) UpdateResourceType(ctx context.Context, in *ResourceType, opts ...grpc.CallOption) (*ResourceTypeResponse, error) {
	out := new(ResourceTypeResponse)
	err := grpc.Invoke(ctx, "/pb.ResourceTypeHandler/UpdateResourceType", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceTypeHandlerClient) GetResourceType(ctx context.Context, in *ResourceTypeID, opts ...grpc.CallOption) (*ResourceTypeResponse, error) {
	out := new(ResourceTypeResponse)
	err := grpc.Invoke(ctx, "/pb.ResourceTypeHandler/GetResourceType", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceTypeHandler service

type ResourceTypeHandlerServer interface {
	// resource type
	CreateResourceType(context.Context, *ResourceType) (*ResourceTypeResponse, error)
	DeleteResourceType(context.Context, *ResourceTypeID) (*ResourceTypeResponse, error)
	UpdateResourceType(context.Context, *ResourceType) (*ResourceTypeResponse, error)
	GetResourceType(context.Context, *ResourceTypeID) (*ResourceTypeResponse, error)
}

func RegisterResourceTypeHandlerServer(s *grpc.Server, srv ResourceTypeHandlerServer) {
	s.RegisterService(&_ResourceTypeHandler_serviceDesc, srv)
}

func _ResourceTypeHandler_CreateResourceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTypeHandlerServer).CreateResourceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceTypeHandler/CreateResourceType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTypeHandlerServer).CreateResourceType(ctx, req.(*ResourceType))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceTypeHandler_DeleteResourceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceTypeID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTypeHandlerServer).DeleteResourceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceTypeHandler/DeleteResourceType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTypeHandlerServer).DeleteResourceType(ctx, req.(*ResourceTypeID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceTypeHandler_UpdateResourceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTypeHandlerServer).UpdateResourceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceTypeHandler/UpdateResourceType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTypeHandlerServer).UpdateResourceType(ctx, req.(*ResourceType))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceTypeHandler_GetResourceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceTypeID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTypeHandlerServer).GetResourceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceTypeHandler/GetResourceType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTypeHandlerServer).GetResourceType(ctx, req.(*ResourceTypeID))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceTypeHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ResourceTypeHandler",
	HandlerType: (*ResourceTypeHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResourceType",
			Handler:    _ResourceTypeHandler_CreateResourceType_Handler,
		},
		{
			MethodName: "DeleteResourceType",
			Handler:    _ResourceTypeHandler_DeleteResourceType_Handler,
		},
		{
			MethodName: "UpdateResourceType",
			Handler:    _ResourceTypeHandler_UpdateResourceType_Handler,
		},
		{
			MethodName: "GetResourceType",
			Handler:    _ResourceTypeHandler_GetResourceType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resourcetype.proto",
}

func init() { proto.RegisterFile("resourcetype.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x6d, 0x05, 0x84, 0x29, 0x22, 0x0e, 0x98, 0x6c, 0xea, 0xa5, 0xe9, 0xc5, 0xc6, 0x18,
	0x4c, 0xea, 0x13, 0x88, 0x18, 0xf5, 0x4a, 0xe0, 0xe2, 0xa5, 0xe9, 0x9f, 0x49, 0x6c, 0xd2, 0x76,
	0x37, 0xdb, 0xed, 0x81, 0x27, 0xf0, 0x4d, 0x7c, 0x2f, 0xdf, 0xc4, 0xb4, 0x50, 0x2d, 0x28, 0x5e,
	0xb8, 0x7e, 0xf3, 0xeb, 0x37, 0xdf, 0x37, 0x5d, 0x40, 0x49, 0x39, 0x2f, 0x64, 0x48, 0x6a, 0x25,
	0x68, 0x22, 0x24, 0x57, 0x1c, 0x75, 0x11, 0x98, 0x06, 0x49, 0xc9, 0xe5, 0x5a, 0x30, 0x2f, 0x6a,
	0xa8, 0x90, 0xb1, 0x4a, 0x45, 0xb2, 0x96, 0xed, 0x4f, 0x0d, 0xfa, 0xf3, 0xcd, 0x64, 0xb1, 0x12,
	0x84, 0x0c, 0x86, 0x35, 0xe9, 0x95, 0x7e, 0x5e, 0x1c, 0x31, 0xcd, 0xd2, 0x9c, 0x1e, 0x9a, 0x3f,
	0x8b, 0xd6, 0x93, 0xcc, 0x4f, 0x89, 0xe9, 0xd5, 0x0c, 0x01, 0x84, 0xe4, 0x51, 0x11, 0xaa, 0x92,
	0x3f, 0xae, 0xb4, 0x4b, 0x18, 0xa5, 0x3c, 0x8b, 0x15, 0x97, 0x5e, 0x48, 0x99, 0x22, 0xe9, 0xbd,
	0xf1, 0x5c, 0xb1, 0xd6, 0x9e, 0xa1, 0xe0, 0x52, 0xb1, 0xb6, 0xa5, 0x39, 0x6d, 0xbc, 0x85, 0xf3,
	0xef, 0x4d, 0x85, 0x8c, 0xbd, 0x32, 0x2f, 0xeb, 0x58, 0x9a, 0x63, 0xb8, 0xe3, 0x89, 0x08, 0x26,
	0x75, 0xe0, 0xa5, 0x8c, 0x17, 0xa9, 0x48, 0x72, 0x1c, 0x40, 0x87, 0x32, 0x3f, 0x48, 0x88, 0x9d,
	0x58, 0x9a, 0xd3, 0xc5, 0x3e, 0xb4, 0x22, 0xca, 0x43, 0xd6, 0x2d, 0x77, 0xd9, 0xd7, 0x30, 0x68,
	0x56, 0x7c, 0x99, 0xed, 0x2f, 0x69, 0xbf, 0x6b, 0x30, 0x6e, 0xc2, 0x73, 0xca, 0x05, 0xcf, 0x72,
	0xc2, 0x9b, 0x3d, 0x9f, 0x18, 0x2e, 0x36, 0x23, 0x6d, 0x16, 0x5c, 0xc1, 0xe9, 0x16, 0x5d, 0x9d,
	0xc9, 0x70, 0x87, 0xbb, 0x28, 0x32, 0x68, 0x57, 0x7f, 0xa9, 0xba, 0x99, 0xe1, 0xf6, 0x4a, 0xe0,
	0xb1, 0x14, 0xdc, 0x0f, 0x1d, 0x46, 0x4d, 0xf4, 0xd9, 0xcf, 0xa2, 0x84, 0x24, 0x4e, 0x01, 0x1f,
	0x24, 0xf9, 0x8a, 0xb6, 0x7c, 0x7e, 0x39, 0x9b, 0x6c, 0x57, 0xa9, 0xab, 0xd8, 0x47, 0x38, 0x03,
	0x9c, 0x51, 0x42, 0x3b, 0x1e, 0x7f, 0x14, 0xf9, 0xd7, 0x65, 0x0a, 0xb8, 0x14, 0xd1, 0x61, 0x49,
	0xee, 0xe1, 0xec, 0x89, 0xd4, 0x21, 0x31, 0xa6, 0xad, 0x57, 0x5d, 0x04, 0x41, 0xa7, 0x7a, 0xcf,
	0x77, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xad, 0x51, 0xaa, 0x0d, 0x03, 0x00, 0x00,
}