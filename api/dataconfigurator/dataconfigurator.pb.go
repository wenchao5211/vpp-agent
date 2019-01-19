// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dataconfigurator/dataconfigurator.proto

package dataconfigurator // import "github.com/ligato/vpp-agent/api/dataconfigurator"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import linux "github.com/ligato/vpp-agent/api/models/linux"
import vpp "github.com/ligato/vpp-agent/api/models/vpp"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Data is a container for all known model data.
type Data struct {
	VppData              *vpp.Data   `protobuf:"bytes,1,opt,name=vpp_data,json=vppData,proto3" json:"vpp_data,omitempty"`
	LinuxData            *linux.Data `protobuf:"bytes,2,opt,name=linux_data,json=linuxData,proto3" json:"linux_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataconfigurator_82b99ef3bb04c26e, []int{0}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (dst *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(dst, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetVppData() *vpp.Data {
	if m != nil {
		return m.VppData
	}
	return nil
}

func (m *Data) GetLinuxData() *linux.Data {
	if m != nil {
		return m.LinuxData
	}
	return nil
}

type UpdateRequest struct {
	// Update describes config data to be updated.
	Update *Data `protobuf:"bytes,1,opt,name=update,proto3" json:"update,omitempty"`
	// The full_resync option can be used
	// to overwrite all existing data.
	FullResync           bool     `protobuf:"varint,2,opt,name=full_resync,json=fullResync,proto3" json:"full_resync,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataconfigurator_82b99ef3bb04c26e, []int{1}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(dst, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetUpdate() *Data {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *UpdateRequest) GetFullResync() bool {
	if m != nil {
		return m.FullResync
	}
	return false
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataconfigurator_82b99ef3bb04c26e, []int{2}
}
func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(dst, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

type DeleteRequest struct {
	// Delete describes config data to be deleted.
	Delete               *Data    `protobuf:"bytes,1,opt,name=delete,proto3" json:"delete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataconfigurator_82b99ef3bb04c26e, []int{3}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetDelete() *Data {
	if m != nil {
		return m.Delete
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataconfigurator_82b99ef3bb04c26e, []int{4}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(dst, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataconfigurator_82b99ef3bb04c26e, []int{5}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type GetResponse struct {
	// Config describes current config data retrieved from agent.
	Config               *Data    `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataconfigurator_82b99ef3bb04c26e, []int{6}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetConfig() *Data {
	if m != nil {
		return m.Config
	}
	return nil
}

type DumpRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpRequest) Reset()         { *m = DumpRequest{} }
func (m *DumpRequest) String() string { return proto.CompactTextString(m) }
func (*DumpRequest) ProtoMessage()    {}
func (*DumpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataconfigurator_82b99ef3bb04c26e, []int{7}
}
func (m *DumpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DumpRequest.Unmarshal(m, b)
}
func (m *DumpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DumpRequest.Marshal(b, m, deterministic)
}
func (dst *DumpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpRequest.Merge(dst, src)
}
func (m *DumpRequest) XXX_Size() int {
	return xxx_messageInfo_DumpRequest.Size(m)
}
func (m *DumpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DumpRequest proto.InternalMessageInfo

type DumpResponse struct {
	// State describes running state data that was dumped from southbound.
	State                *Data    `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpResponse) Reset()         { *m = DumpResponse{} }
func (m *DumpResponse) String() string { return proto.CompactTextString(m) }
func (*DumpResponse) ProtoMessage()    {}
func (*DumpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataconfigurator_82b99ef3bb04c26e, []int{8}
}
func (m *DumpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DumpResponse.Unmarshal(m, b)
}
func (m *DumpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DumpResponse.Marshal(b, m, deterministic)
}
func (dst *DumpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpResponse.Merge(dst, src)
}
func (m *DumpResponse) XXX_Size() int {
	return xxx_messageInfo_DumpResponse.Size(m)
}
func (m *DumpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DumpResponse proto.InternalMessageInfo

func (m *DumpResponse) GetState() *Data {
	if m != nil {
		return m.State
	}
	return nil
}

func init() {
	proto.RegisterType((*Data)(nil), "dataconfigurator.Data")
	proto.RegisterType((*UpdateRequest)(nil), "dataconfigurator.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "dataconfigurator.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "dataconfigurator.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "dataconfigurator.DeleteResponse")
	proto.RegisterType((*GetRequest)(nil), "dataconfigurator.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "dataconfigurator.GetResponse")
	proto.RegisterType((*DumpRequest)(nil), "dataconfigurator.DumpRequest")
	proto.RegisterType((*DumpResponse)(nil), "dataconfigurator.DumpResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataConfiguratorClient is the client API for DataConfigurator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataConfiguratorClient interface {
	// Update is used for updating config data.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete is used for deleting config data.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Get is used for listing current config data.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Dump is used for dumping real data state.
	Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (*DumpResponse, error)
}

type dataConfiguratorClient struct {
	cc *grpc.ClientConn
}

func NewDataConfiguratorClient(cc *grpc.ClientConn) DataConfiguratorClient {
	return &dataConfiguratorClient{cc}
}

func (c *dataConfiguratorClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/dataconfigurator.DataConfigurator/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataConfiguratorClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/dataconfigurator.DataConfigurator/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataConfiguratorClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/dataconfigurator.DataConfigurator/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataConfiguratorClient) Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (*DumpResponse, error) {
	out := new(DumpResponse)
	err := c.cc.Invoke(ctx, "/dataconfigurator.DataConfigurator/Dump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataConfiguratorServer is the server API for DataConfigurator service.
type DataConfiguratorServer interface {
	// Update is used for updating config data.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete is used for deleting config data.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Get is used for listing current config data.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Dump is used for dumping real data state.
	Dump(context.Context, *DumpRequest) (*DumpResponse, error)
}

func RegisterDataConfiguratorServer(s *grpc.Server, srv DataConfiguratorServer) {
	s.RegisterService(&_DataConfigurator_serviceDesc, srv)
}

func _DataConfigurator_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataConfiguratorServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataconfigurator.DataConfigurator/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataConfiguratorServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataConfigurator_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataConfiguratorServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataconfigurator.DataConfigurator/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataConfiguratorServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataConfigurator_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataConfiguratorServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataconfigurator.DataConfigurator/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataConfiguratorServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataConfigurator_Dump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataConfiguratorServer).Dump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataconfigurator.DataConfigurator/Dump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataConfiguratorServer).Dump(ctx, req.(*DumpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataConfigurator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataconfigurator.DataConfigurator",
	HandlerType: (*DataConfiguratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _DataConfigurator_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DataConfigurator_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DataConfigurator_Get_Handler,
		},
		{
			MethodName: "Dump",
			Handler:    _DataConfigurator_Dump_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataconfigurator/dataconfigurator.proto",
}

func init() {
	proto.RegisterFile("dataconfigurator/dataconfigurator.proto", fileDescriptor_dataconfigurator_82b99ef3bb04c26e)
}

var fileDescriptor_dataconfigurator_82b99ef3bb04c26e = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x5f, 0x6b, 0xe2, 0x40,
	0x14, 0xc5, 0xd1, 0x75, 0xb3, 0x7a, 0xa3, 0x8b, 0x0c, 0xcb, 0x22, 0xa1, 0x7f, 0x24, 0x14, 0x5a,
	0x4a, 0x9b, 0x80, 0x7d, 0x6c, 0x4b, 0xa9, 0x5a, 0x7c, 0xf0, 0x2d, 0x20, 0x94, 0xbe, 0xd8, 0xd1,
	0x8c, 0x69, 0x20, 0x66, 0xa6, 0xc9, 0x8c, 0xb4, 0x9f, 0xac, 0x5f, 0xaf, 0xcc, 0x1f, 0xa9, 0x31,
	0x8a, 0x0f, 0xca, 0xcc, 0xb9, 0xe7, 0xfe, 0x6e, 0x72, 0xc2, 0x85, 0xf3, 0x10, 0x73, 0x3c, 0xa7,
	0xe9, 0x22, 0x8e, 0x44, 0x86, 0x39, 0xcd, 0xfc, 0x6d, 0xc1, 0x63, 0x19, 0xe5, 0x14, 0xb5, 0xb7,
	0x75, 0xe7, 0xdf, 0x92, 0x86, 0x24, 0xc9, 0xfd, 0x15, 0x63, 0xf2, 0xa7, 0x7d, 0x4e, 0xc7, 0xa8,
	0x49, 0x9c, 0x8a, 0x0f, 0xfd, 0xaf, 0x2b, 0xee, 0x33, 0xd4, 0x86, 0x98, 0x63, 0x74, 0x06, 0xf5,
	0x15, 0x63, 0x53, 0xc9, 0xeb, 0x54, 0xba, 0x95, 0x0b, 0xbb, 0xd7, 0xf0, 0x64, 0xbf, 0x2c, 0x06,
	0x7f, 0x56, 0x8c, 0x29, 0xd7, 0x25, 0x80, 0x6a, 0xd6, 0xbe, 0xaa, 0xf2, 0xd9, 0x9e, 0xe6, 0x29,
	0x67, 0x43, 0x9d, 0xe5, 0xd1, 0x7d, 0x85, 0xd6, 0x84, 0x85, 0x98, 0x93, 0x80, 0xbc, 0x0b, 0x92,
	0x73, 0xe4, 0x81, 0x25, 0x94, 0x60, 0x06, 0xfc, 0xf7, 0x4a, 0x6f, 0xa5, 0x18, 0xc6, 0x85, 0x4e,
	0xc1, 0x5e, 0x88, 0x24, 0x99, 0x66, 0x24, 0xff, 0x4c, 0xe7, 0x6a, 0x5a, 0x3d, 0x00, 0x29, 0x05,
	0x4a, 0x71, 0xdb, 0xf0, 0x77, 0x3d, 0x21, 0x67, 0x34, 0xcd, 0x89, 0xfb, 0x00, 0xad, 0x21, 0x49,
	0x48, 0x61, 0x66, 0xa8, 0x84, 0x43, 0x33, 0xb5, 0x4b, 0x22, 0xd7, 0x00, 0x83, 0x6c, 0x02, 0x8c,
	0x08, 0x37, 0x3c, 0xf7, 0x1e, 0x6c, 0x75, 0xd3, 0x45, 0x89, 0xd7, 0xac, 0x43, 0x78, 0x2d, 0xb9,
	0x2d, 0xb0, 0x87, 0x62, 0xc9, 0xd6, 0xb4, 0x3b, 0x68, 0xea, 0xab, 0xc1, 0x5d, 0xc1, 0xef, 0x9c,
	0x1f, 0x0e, 0x48, 0x9b, 0x7a, 0x5f, 0x55, 0x68, 0xcb, 0xfb, 0x60, 0xc3, 0x80, 0xc6, 0x60, 0x4d,
	0x4c, 0x7c, 0xe5, 0xee, 0xc2, 0xf7, 0x70, 0xba, 0xfb, 0x0d, 0xe6, 0x79, 0xc6, 0x60, 0xe9, 0x34,
	0x76, 0xc1, 0x0a, 0x41, 0xef, 0x82, 0x15, 0x83, 0x44, 0x7d, 0xf8, 0x35, 0x22, 0x1c, 0x1d, 0x95,
	0x8d, 0x3f, 0xf9, 0x3a, 0xc7, 0x7b, 0xaa, 0x86, 0xf1, 0x04, 0x35, 0x19, 0x18, 0xda, 0x61, 0xdb,
	0xc8, 0xd5, 0x39, 0xd9, 0x57, 0xd6, 0x98, 0xfe, 0xe0, 0xe5, 0x31, 0x8a, 0xf9, 0x9b, 0x98, 0x79,
	0x73, 0xba, 0xf4, 0x93, 0x38, 0xc2, 0x9c, 0xca, 0x6d, 0xb9, 0xc6, 0x11, 0x49, 0xb9, 0x8f, 0x59,
	0x5c, 0x5a, 0xb8, 0xdb, 0x6d, 0x61, 0x66, 0xa9, 0x05, 0xba, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x40, 0x94, 0x5b, 0x58, 0xad, 0x03, 0x00, 0x00,
}
