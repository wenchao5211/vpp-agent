// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: models/vpp/ipsec/ipsec.proto

package vpp_ipsec // import "github.com/ligato/vpp-agent/api/models/vpp/ipsec"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SecurityPolicyDatabase_PolicyEntry_Action int32

const (
	SecurityPolicyDatabase_PolicyEntry_BYPASS  SecurityPolicyDatabase_PolicyEntry_Action = 0
	SecurityPolicyDatabase_PolicyEntry_DISCARD SecurityPolicyDatabase_PolicyEntry_Action = 1
	SecurityPolicyDatabase_PolicyEntry_RESOLVE SecurityPolicyDatabase_PolicyEntry_Action = 2
	SecurityPolicyDatabase_PolicyEntry_PROTECT SecurityPolicyDatabase_PolicyEntry_Action = 3
)

var SecurityPolicyDatabase_PolicyEntry_Action_name = map[int32]string{
	0: "BYPASS",
	1: "DISCARD",
	2: "RESOLVE",
	3: "PROTECT",
}
var SecurityPolicyDatabase_PolicyEntry_Action_value = map[string]int32{
	"BYPASS":  0,
	"DISCARD": 1,
	"RESOLVE": 2,
	"PROTECT": 3,
}

func (x SecurityPolicyDatabase_PolicyEntry_Action) String() string {
	return proto.EnumName(SecurityPolicyDatabase_PolicyEntry_Action_name, int32(x))
}
func (SecurityPolicyDatabase_PolicyEntry_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_5ddd9501648734e7, []int{0, 1, 0}
}

type SecurityAssociation_IPSecProtocol int32

const (
	SecurityAssociation_AH  SecurityAssociation_IPSecProtocol = 0
	SecurityAssociation_ESP SecurityAssociation_IPSecProtocol = 1
)

var SecurityAssociation_IPSecProtocol_name = map[int32]string{
	0: "AH",
	1: "ESP",
}
var SecurityAssociation_IPSecProtocol_value = map[string]int32{
	"AH":  0,
	"ESP": 1,
}

func (x SecurityAssociation_IPSecProtocol) String() string {
	return proto.EnumName(SecurityAssociation_IPSecProtocol_name, int32(x))
}
func (SecurityAssociation_IPSecProtocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_5ddd9501648734e7, []int{1, 0}
}

type SecurityAssociation_CryptoAlg int32

const (
	SecurityAssociation_NONE_CRYPTO SecurityAssociation_CryptoAlg = 0
	SecurityAssociation_AES_CBC_128 SecurityAssociation_CryptoAlg = 1
	SecurityAssociation_AES_CBC_192 SecurityAssociation_CryptoAlg = 2
	SecurityAssociation_AES_CBC_256 SecurityAssociation_CryptoAlg = 3
)

var SecurityAssociation_CryptoAlg_name = map[int32]string{
	0: "NONE_CRYPTO",
	1: "AES_CBC_128",
	2: "AES_CBC_192",
	3: "AES_CBC_256",
}
var SecurityAssociation_CryptoAlg_value = map[string]int32{
	"NONE_CRYPTO": 0,
	"AES_CBC_128": 1,
	"AES_CBC_192": 2,
	"AES_CBC_256": 3,
}

func (x SecurityAssociation_CryptoAlg) String() string {
	return proto.EnumName(SecurityAssociation_CryptoAlg_name, int32(x))
}
func (SecurityAssociation_CryptoAlg) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_5ddd9501648734e7, []int{1, 1}
}

type SecurityAssociation_IntegAlg int32

const (
	SecurityAssociation_NONE_INTEG  SecurityAssociation_IntegAlg = 0
	SecurityAssociation_MD5_96      SecurityAssociation_IntegAlg = 1
	SecurityAssociation_SHA1_96     SecurityAssociation_IntegAlg = 2
	SecurityAssociation_SHA_256_96  SecurityAssociation_IntegAlg = 3
	SecurityAssociation_SHA_256_128 SecurityAssociation_IntegAlg = 4
	SecurityAssociation_SHA_384_192 SecurityAssociation_IntegAlg = 5
	SecurityAssociation_SHA_512_256 SecurityAssociation_IntegAlg = 6
)

var SecurityAssociation_IntegAlg_name = map[int32]string{
	0: "NONE_INTEG",
	1: "MD5_96",
	2: "SHA1_96",
	3: "SHA_256_96",
	4: "SHA_256_128",
	5: "SHA_384_192",
	6: "SHA_512_256",
}
var SecurityAssociation_IntegAlg_value = map[string]int32{
	"NONE_INTEG":  0,
	"MD5_96":      1,
	"SHA1_96":     2,
	"SHA_256_96":  3,
	"SHA_256_128": 4,
	"SHA_384_192": 5,
	"SHA_512_256": 6,
}

func (x SecurityAssociation_IntegAlg) String() string {
	return proto.EnumName(SecurityAssociation_IntegAlg_name, int32(x))
}
func (SecurityAssociation_IntegAlg) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_5ddd9501648734e7, []int{1, 2}
}

// Security Policy Database (SPD)
type SecurityPolicyDatabase struct {
	Index                string                                `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Interfaces           []*SecurityPolicyDatabase_Interface   `protobuf:"bytes,2,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	PolicyEntries        []*SecurityPolicyDatabase_PolicyEntry `protobuf:"bytes,3,rep,name=policy_entries,json=policyEntries,proto3" json:"policy_entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *SecurityPolicyDatabase) Reset()         { *m = SecurityPolicyDatabase{} }
func (m *SecurityPolicyDatabase) String() string { return proto.CompactTextString(m) }
func (*SecurityPolicyDatabase) ProtoMessage()    {}
func (*SecurityPolicyDatabase) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_5ddd9501648734e7, []int{0}
}
func (m *SecurityPolicyDatabase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityPolicyDatabase.Unmarshal(m, b)
}
func (m *SecurityPolicyDatabase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityPolicyDatabase.Marshal(b, m, deterministic)
}
func (dst *SecurityPolicyDatabase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityPolicyDatabase.Merge(dst, src)
}
func (m *SecurityPolicyDatabase) XXX_Size() int {
	return xxx_messageInfo_SecurityPolicyDatabase.Size(m)
}
func (m *SecurityPolicyDatabase) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityPolicyDatabase.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityPolicyDatabase proto.InternalMessageInfo

func (m *SecurityPolicyDatabase) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *SecurityPolicyDatabase) GetInterfaces() []*SecurityPolicyDatabase_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *SecurityPolicyDatabase) GetPolicyEntries() []*SecurityPolicyDatabase_PolicyEntry {
	if m != nil {
		return m.PolicyEntries
	}
	return nil
}

func (*SecurityPolicyDatabase) XXX_MessageName() string {
	return "vpp.ipsec.SecurityPolicyDatabase"
}

type SecurityPolicyDatabase_Interface struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecurityPolicyDatabase_Interface) Reset()         { *m = SecurityPolicyDatabase_Interface{} }
func (m *SecurityPolicyDatabase_Interface) String() string { return proto.CompactTextString(m) }
func (*SecurityPolicyDatabase_Interface) ProtoMessage()    {}
func (*SecurityPolicyDatabase_Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_5ddd9501648734e7, []int{0, 0}
}
func (m *SecurityPolicyDatabase_Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityPolicyDatabase_Interface.Unmarshal(m, b)
}
func (m *SecurityPolicyDatabase_Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityPolicyDatabase_Interface.Marshal(b, m, deterministic)
}
func (dst *SecurityPolicyDatabase_Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityPolicyDatabase_Interface.Merge(dst, src)
}
func (m *SecurityPolicyDatabase_Interface) XXX_Size() int {
	return xxx_messageInfo_SecurityPolicyDatabase_Interface.Size(m)
}
func (m *SecurityPolicyDatabase_Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityPolicyDatabase_Interface.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityPolicyDatabase_Interface proto.InternalMessageInfo

func (m *SecurityPolicyDatabase_Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (*SecurityPolicyDatabase_Interface) XXX_MessageName() string {
	return "vpp.ipsec.SecurityPolicyDatabase.Interface"
}

type SecurityPolicyDatabase_PolicyEntry struct {
	SaIndex              string                                    `protobuf:"bytes,1,opt,name=sa_index,json=saIndex,proto3" json:"sa_index,omitempty"`
	Priority             int32                                     `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	IsOutbound           bool                                      `protobuf:"varint,3,opt,name=is_outbound,json=isOutbound,proto3" json:"is_outbound,omitempty"`
	RemoteAddrStart      string                                    `protobuf:"bytes,4,opt,name=remote_addr_start,json=remoteAddrStart,proto3" json:"remote_addr_start,omitempty"`
	RemoteAddrStop       string                                    `protobuf:"bytes,5,opt,name=remote_addr_stop,json=remoteAddrStop,proto3" json:"remote_addr_stop,omitempty"`
	LocalAddrStart       string                                    `protobuf:"bytes,6,opt,name=local_addr_start,json=localAddrStart,proto3" json:"local_addr_start,omitempty"`
	LocalAddrStop        string                                    `protobuf:"bytes,7,opt,name=local_addr_stop,json=localAddrStop,proto3" json:"local_addr_stop,omitempty"`
	Protocol             uint32                                    `protobuf:"varint,8,opt,name=protocol,proto3" json:"protocol,omitempty"`
	RemotePortStart      uint32                                    `protobuf:"varint,9,opt,name=remote_port_start,json=remotePortStart,proto3" json:"remote_port_start,omitempty"`
	RemotePortStop       uint32                                    `protobuf:"varint,10,opt,name=remote_port_stop,json=remotePortStop,proto3" json:"remote_port_stop,omitempty"`
	LocalPortStart       uint32                                    `protobuf:"varint,11,opt,name=local_port_start,json=localPortStart,proto3" json:"local_port_start,omitempty"`
	LocalPortStop        uint32                                    `protobuf:"varint,12,opt,name=local_port_stop,json=localPortStop,proto3" json:"local_port_stop,omitempty"`
	Action               SecurityPolicyDatabase_PolicyEntry_Action `protobuf:"varint,13,opt,name=action,proto3,enum=vpp.ipsec.SecurityPolicyDatabase_PolicyEntry_Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *SecurityPolicyDatabase_PolicyEntry) Reset()         { *m = SecurityPolicyDatabase_PolicyEntry{} }
func (m *SecurityPolicyDatabase_PolicyEntry) String() string { return proto.CompactTextString(m) }
func (*SecurityPolicyDatabase_PolicyEntry) ProtoMessage()    {}
func (*SecurityPolicyDatabase_PolicyEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_5ddd9501648734e7, []int{0, 1}
}
func (m *SecurityPolicyDatabase_PolicyEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityPolicyDatabase_PolicyEntry.Unmarshal(m, b)
}
func (m *SecurityPolicyDatabase_PolicyEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityPolicyDatabase_PolicyEntry.Marshal(b, m, deterministic)
}
func (dst *SecurityPolicyDatabase_PolicyEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityPolicyDatabase_PolicyEntry.Merge(dst, src)
}
func (m *SecurityPolicyDatabase_PolicyEntry) XXX_Size() int {
	return xxx_messageInfo_SecurityPolicyDatabase_PolicyEntry.Size(m)
}
func (m *SecurityPolicyDatabase_PolicyEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityPolicyDatabase_PolicyEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityPolicyDatabase_PolicyEntry proto.InternalMessageInfo

func (m *SecurityPolicyDatabase_PolicyEntry) GetSaIndex() string {
	if m != nil {
		return m.SaIndex
	}
	return ""
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetIsOutbound() bool {
	if m != nil {
		return m.IsOutbound
	}
	return false
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetRemoteAddrStart() string {
	if m != nil {
		return m.RemoteAddrStart
	}
	return ""
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetRemoteAddrStop() string {
	if m != nil {
		return m.RemoteAddrStop
	}
	return ""
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetLocalAddrStart() string {
	if m != nil {
		return m.LocalAddrStart
	}
	return ""
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetLocalAddrStop() string {
	if m != nil {
		return m.LocalAddrStop
	}
	return ""
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetRemotePortStart() uint32 {
	if m != nil {
		return m.RemotePortStart
	}
	return 0
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetRemotePortStop() uint32 {
	if m != nil {
		return m.RemotePortStop
	}
	return 0
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetLocalPortStart() uint32 {
	if m != nil {
		return m.LocalPortStart
	}
	return 0
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetLocalPortStop() uint32 {
	if m != nil {
		return m.LocalPortStop
	}
	return 0
}

func (m *SecurityPolicyDatabase_PolicyEntry) GetAction() SecurityPolicyDatabase_PolicyEntry_Action {
	if m != nil {
		return m.Action
	}
	return SecurityPolicyDatabase_PolicyEntry_BYPASS
}

func (*SecurityPolicyDatabase_PolicyEntry) XXX_MessageName() string {
	return "vpp.ipsec.SecurityPolicyDatabase.PolicyEntry"
}

// Security Association (SA)
type SecurityAssociation struct {
	Index                string                            `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Spi                  uint32                            `protobuf:"varint,2,opt,name=spi,proto3" json:"spi,omitempty"`
	Protocol             SecurityAssociation_IPSecProtocol `protobuf:"varint,3,opt,name=protocol,proto3,enum=vpp.ipsec.SecurityAssociation_IPSecProtocol" json:"protocol,omitempty"`
	CryptoAlg            SecurityAssociation_CryptoAlg     `protobuf:"varint,4,opt,name=crypto_alg,json=cryptoAlg,proto3,enum=vpp.ipsec.SecurityAssociation_CryptoAlg" json:"crypto_alg,omitempty"`
	CryptoKey            string                            `protobuf:"bytes,5,opt,name=crypto_key,json=cryptoKey,proto3" json:"crypto_key,omitempty"`
	IntegAlg             SecurityAssociation_IntegAlg      `protobuf:"varint,6,opt,name=integ_alg,json=integAlg,proto3,enum=vpp.ipsec.SecurityAssociation_IntegAlg" json:"integ_alg,omitempty"`
	IntegKey             string                            `protobuf:"bytes,7,opt,name=integ_key,json=integKey,proto3" json:"integ_key,omitempty"`
	UseEsn               bool                              `protobuf:"varint,8,opt,name=use_esn,json=useEsn,proto3" json:"use_esn,omitempty"`
	UseAntiReplay        bool                              `protobuf:"varint,9,opt,name=use_anti_replay,json=useAntiReplay,proto3" json:"use_anti_replay,omitempty"`
	TunnelSrcAddr        string                            `protobuf:"bytes,10,opt,name=tunnel_src_addr,json=tunnelSrcAddr,proto3" json:"tunnel_src_addr,omitempty"`
	TunnelDstAddr        string                            `protobuf:"bytes,11,opt,name=tunnel_dst_addr,json=tunnelDstAddr,proto3" json:"tunnel_dst_addr,omitempty"`
	EnableUdpEncap       bool                              `protobuf:"varint,12,opt,name=enable_udp_encap,json=enableUdpEncap,proto3" json:"enable_udp_encap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *SecurityAssociation) Reset()         { *m = SecurityAssociation{} }
func (m *SecurityAssociation) String() string { return proto.CompactTextString(m) }
func (*SecurityAssociation) ProtoMessage()    {}
func (*SecurityAssociation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ipsec_5ddd9501648734e7, []int{1}
}
func (m *SecurityAssociation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityAssociation.Unmarshal(m, b)
}
func (m *SecurityAssociation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityAssociation.Marshal(b, m, deterministic)
}
func (dst *SecurityAssociation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityAssociation.Merge(dst, src)
}
func (m *SecurityAssociation) XXX_Size() int {
	return xxx_messageInfo_SecurityAssociation.Size(m)
}
func (m *SecurityAssociation) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityAssociation.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityAssociation proto.InternalMessageInfo

func (m *SecurityAssociation) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *SecurityAssociation) GetSpi() uint32 {
	if m != nil {
		return m.Spi
	}
	return 0
}

func (m *SecurityAssociation) GetProtocol() SecurityAssociation_IPSecProtocol {
	if m != nil {
		return m.Protocol
	}
	return SecurityAssociation_AH
}

func (m *SecurityAssociation) GetCryptoAlg() SecurityAssociation_CryptoAlg {
	if m != nil {
		return m.CryptoAlg
	}
	return SecurityAssociation_NONE_CRYPTO
}

func (m *SecurityAssociation) GetCryptoKey() string {
	if m != nil {
		return m.CryptoKey
	}
	return ""
}

func (m *SecurityAssociation) GetIntegAlg() SecurityAssociation_IntegAlg {
	if m != nil {
		return m.IntegAlg
	}
	return SecurityAssociation_NONE_INTEG
}

func (m *SecurityAssociation) GetIntegKey() string {
	if m != nil {
		return m.IntegKey
	}
	return ""
}

func (m *SecurityAssociation) GetUseEsn() bool {
	if m != nil {
		return m.UseEsn
	}
	return false
}

func (m *SecurityAssociation) GetUseAntiReplay() bool {
	if m != nil {
		return m.UseAntiReplay
	}
	return false
}

func (m *SecurityAssociation) GetTunnelSrcAddr() string {
	if m != nil {
		return m.TunnelSrcAddr
	}
	return ""
}

func (m *SecurityAssociation) GetTunnelDstAddr() string {
	if m != nil {
		return m.TunnelDstAddr
	}
	return ""
}

func (m *SecurityAssociation) GetEnableUdpEncap() bool {
	if m != nil {
		return m.EnableUdpEncap
	}
	return false
}

func (*SecurityAssociation) XXX_MessageName() string {
	return "vpp.ipsec.SecurityAssociation"
}
func init() {
	proto.RegisterType((*SecurityPolicyDatabase)(nil), "vpp.ipsec.SecurityPolicyDatabase")
	proto.RegisterType((*SecurityPolicyDatabase_Interface)(nil), "vpp.ipsec.SecurityPolicyDatabase.Interface")
	proto.RegisterType((*SecurityPolicyDatabase_PolicyEntry)(nil), "vpp.ipsec.SecurityPolicyDatabase.PolicyEntry")
	proto.RegisterType((*SecurityAssociation)(nil), "vpp.ipsec.SecurityAssociation")
	proto.RegisterEnum("vpp.ipsec.SecurityPolicyDatabase_PolicyEntry_Action", SecurityPolicyDatabase_PolicyEntry_Action_name, SecurityPolicyDatabase_PolicyEntry_Action_value)
	proto.RegisterEnum("vpp.ipsec.SecurityAssociation_IPSecProtocol", SecurityAssociation_IPSecProtocol_name, SecurityAssociation_IPSecProtocol_value)
	proto.RegisterEnum("vpp.ipsec.SecurityAssociation_CryptoAlg", SecurityAssociation_CryptoAlg_name, SecurityAssociation_CryptoAlg_value)
	proto.RegisterEnum("vpp.ipsec.SecurityAssociation_IntegAlg", SecurityAssociation_IntegAlg_name, SecurityAssociation_IntegAlg_value)
}

func init() { proto.RegisterFile("models/vpp/ipsec/ipsec.proto", fileDescriptor_ipsec_5ddd9501648734e7) }

var fileDescriptor_ipsec_5ddd9501648734e7 = []byte{
	// 874 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0xae, 0x93, 0xc6, 0x49, 0x5e, 0x48, 0xd6, 0x0c, 0x08, 0x4c, 0xf9, 0xb1, 0x51, 0x0e, 0x60,
	0x01, 0x4d, 0xd4, 0xec, 0x76, 0xb5, 0xcb, 0x5e, 0x70, 0x13, 0x6b, 0x1b, 0xed, 0xd2, 0x44, 0x76,
	0x40, 0x5a, 0x2e, 0xd6, 0xc4, 0x9e, 0x0d, 0x23, 0x5c, 0xcf, 0xc8, 0x33, 0xae, 0xc8, 0x7f, 0xc8,
	0x95, 0x3b, 0x37, 0xfe, 0x09, 0x8e, 0x68, 0xc6, 0x4e, 0x3a, 0x5d, 0x55, 0xaa, 0xb8, 0x44, 0xef,
	0x7d, 0xf3, 0xbd, 0xef, 0xcb, 0xb3, 0xbf, 0x91, 0xe1, 0x8b, 0x6b, 0x96, 0x92, 0x4c, 0x4c, 0x6e,
	0x38, 0x9f, 0x50, 0x2e, 0x48, 0x52, 0xfd, 0x8e, 0x79, 0xc1, 0x24, 0x43, 0xdd, 0x1b, 0xce, 0xc7,
	0x1a, 0x38, 0x39, 0xdd, 0x52, 0xf9, 0x5b, 0xb9, 0x19, 0x27, 0xec, 0x7a, 0xb2, 0x65, 0x5b, 0x36,
	0xd1, 0x8c, 0x4d, 0xf9, 0x4e, 0x77, 0xba, 0xd1, 0x55, 0x35, 0x39, 0xfa, 0xcb, 0x86, 0x4f, 0x22,
	0x92, 0x94, 0x05, 0x95, 0xbb, 0x15, 0xcb, 0x68, 0xb2, 0x9b, 0x63, 0x89, 0x37, 0x58, 0x10, 0xf4,
	0x31, 0xb4, 0x68, 0x9e, 0x92, 0x3f, 0x5c, 0x6b, 0x68, 0x79, 0xdd, 0xb0, 0x6a, 0xd0, 0x6b, 0x00,
	0x9a, 0x4b, 0x52, 0xbc, 0xc3, 0x09, 0x11, 0x6e, 0x63, 0xd8, 0xf4, 0x7a, 0xd3, 0xef, 0xc6, 0x07,
	0xff, 0xf1, 0xfd, 0x62, 0xe3, 0xc5, 0x7e, 0x26, 0x34, 0xc6, 0xd1, 0x1a, 0x06, 0x5c, 0xf3, 0x62,
	0x92, 0xcb, 0x82, 0x12, 0xe1, 0x36, 0xb5, 0xe0, 0xe9, 0xc3, 0x82, 0x55, 0x1b, 0xe4, 0xb2, 0xd8,
	0x85, 0x7d, 0x7e, 0x68, 0x28, 0x11, 0x27, 0x8f, 0xa1, 0x7b, 0xb0, 0x43, 0x08, 0x8e, 0x73, 0x7c,
	0x4d, 0xea, 0x25, 0x74, 0x7d, 0xf2, 0xf7, 0x31, 0xf4, 0x8c, 0x79, 0xf4, 0x19, 0x74, 0x04, 0x8e,
	0xcd, 0x65, 0xdb, 0x02, 0x2f, 0xf4, 0xba, 0x27, 0xd0, 0xe1, 0x05, 0x65, 0xea, 0x0f, 0xb8, 0x8d,
	0xa1, 0xe5, 0xb5, 0xc2, 0x43, 0x8f, 0x1e, 0x43, 0x8f, 0x8a, 0x98, 0x95, 0x72, 0xc3, 0xca, 0x3c,
	0x75, 0x9b, 0x43, 0xcb, 0xeb, 0x84, 0x40, 0xc5, 0xb2, 0x46, 0xd0, 0xb7, 0xf0, 0x61, 0x41, 0xae,
	0x99, 0x24, 0x31, 0x4e, 0xd3, 0x22, 0x16, 0x12, 0x17, 0xd2, 0x3d, 0xd6, 0x06, 0x8f, 0xaa, 0x03,
	0x3f, 0x4d, 0x8b, 0x48, 0xc1, 0xc8, 0x03, 0xe7, 0x2e, 0x97, 0x71, 0xb7, 0xa5, 0xa9, 0x03, 0x93,
	0xca, 0xb8, 0x62, 0x66, 0x2c, 0xc1, 0x99, 0x29, 0x6a, 0x57, 0x4c, 0x8d, 0xdf, 0x6a, 0x7e, 0x0d,
	0x8f, 0xee, 0x30, 0x19, 0x77, 0xdb, 0x9a, 0xd8, 0x37, 0x88, 0x8c, 0x57, 0x4b, 0x32, 0xc9, 0x12,
	0x96, 0xb9, 0x9d, 0xa1, 0xe5, 0xf5, 0xc3, 0x43, 0x6f, 0xec, 0xc0, 0x59, 0x21, 0x6b, 0xbb, 0xae,
	0x26, 0xd5, 0x3b, 0xac, 0x58, 0x21, 0xdf, 0xdf, 0xa1, 0xe6, 0x32, 0xee, 0x82, 0xa6, 0x0e, 0x4c,
	0xaa, 0xb9, 0x83, 0x21, 0xda, 0xab, 0x98, 0x1a, 0xbf, 0xd5, 0x3c, 0xec, 0x70, 0x2b, 0xf9, 0x81,
	0x26, 0xf6, 0x0d, 0x22, 0xe3, 0xe8, 0x0d, 0xd8, 0x38, 0x91, 0x94, 0xe5, 0x6e, 0x7f, 0x68, 0x79,
	0x83, 0xe9, 0xd3, 0xff, 0x15, 0xa1, 0xb1, 0xaf, 0x67, 0xc3, 0x5a, 0x63, 0xf4, 0x12, 0xec, 0x0a,
	0x41, 0x00, 0xf6, 0xc5, 0xdb, 0x95, 0x1f, 0x45, 0xce, 0x11, 0xea, 0x41, 0x7b, 0xbe, 0x88, 0x66,
	0x7e, 0x38, 0x77, 0x2c, 0xd5, 0x84, 0x41, 0xb4, 0x7c, 0xf3, 0x4b, 0xe0, 0x34, 0x54, 0xb3, 0x0a,
	0x97, 0xeb, 0x60, 0xb6, 0x76, 0x9a, 0xa3, 0x7f, 0x5b, 0xf0, 0xd1, 0xde, 0xd2, 0x17, 0x82, 0x25,
	0x14, 0x6b, 0xa9, 0xfb, 0x2f, 0x94, 0x03, 0x4d, 0xc1, 0xa9, 0x0e, 0x57, 0x3f, 0x54, 0x25, 0xba,
	0x34, 0x5e, 0x47, 0x53, 0x2f, 0xf3, 0xfd, 0x3d, 0xcb, 0x18, 0xca, 0xe3, 0xc5, 0x2a, 0x22, 0xc9,
	0xaa, 0x9e, 0x31, 0x5e, 0xde, 0x2b, 0x80, 0xa4, 0xd8, 0x71, 0xc9, 0x62, 0x9c, 0x6d, 0x75, 0xf2,
	0x06, 0x53, 0xef, 0x01, 0xad, 0x99, 0x1e, 0xf0, 0xb3, 0x6d, 0xd8, 0x4d, 0xf6, 0x25, 0xfa, 0xf2,
	0x20, 0xf4, 0x3b, 0xd9, 0xd5, 0xb9, 0xac, 0x8f, 0x5f, 0x93, 0x1d, 0x9a, 0x43, 0x57, 0xdd, 0xea,
	0xad, 0xb6, 0xb1, 0xb5, 0xcd, 0x37, 0x0f, 0xfd, 0x65, 0xc5, 0x57, 0x2e, 0x1d, 0x5a, 0x57, 0xe8,
	0xf3, 0xbd, 0x8a, 0xf2, 0xa8, 0x82, 0x5a, 0x1d, 0x2a, 0x8b, 0x4f, 0xa1, 0x5d, 0x0a, 0x12, 0x13,
	0x91, 0xeb, 0x88, 0x76, 0x42, 0xbb, 0x14, 0x24, 0x10, 0xb9, 0x0a, 0x88, 0x3a, 0xc0, 0xb9, 0xa4,
	0x71, 0x41, 0x78, 0x86, 0x77, 0x3a, 0x9e, 0x9d, 0xb0, 0x5f, 0x0a, 0xe2, 0xe7, 0x92, 0x86, 0x1a,
	0x54, 0x3c, 0x59, 0xe6, 0x39, 0xc9, 0x62, 0x51, 0x24, 0xfa, 0x46, 0xe8, 0x6c, 0x76, 0xc3, 0x7e,
	0x05, 0x47, 0x45, 0xa2, 0x2e, 0x84, 0xc1, 0x4b, 0x85, 0xac, 0x78, 0x3d, 0x93, 0x37, 0x17, 0x52,
	0xf3, 0x3c, 0x70, 0x48, 0x8e, 0x37, 0x19, 0x89, 0xcb, 0x94, 0xc7, 0x24, 0x4f, 0x70, 0x95, 0xcc,
	0x4e, 0x38, 0xa8, 0xf0, 0x9f, 0x53, 0x1e, 0x28, 0x74, 0x34, 0x84, 0xfe, 0x9d, 0x17, 0x84, 0x6c,
	0x68, 0xf8, 0x97, 0xce, 0x11, 0x6a, 0x43, 0x33, 0x88, 0x56, 0x8e, 0x35, 0x5a, 0x42, 0xf7, 0xf0,
	0xd8, 0xd1, 0x23, 0xe8, 0x5d, 0x2d, 0xaf, 0x82, 0x78, 0x16, 0xbe, 0x5d, 0xad, 0x97, 0xce, 0x91,
	0x02, 0xfc, 0x20, 0x8a, 0x67, 0x17, 0xb3, 0xf8, 0x6c, 0xfa, 0xdc, 0xb1, 0xee, 0x00, 0x2f, 0xa6,
	0x4e, 0xc3, 0x04, 0xa6, 0xe7, 0xcf, 0x9c, 0xe6, 0xe8, 0x06, 0x3a, 0xfb, 0x07, 0x8c, 0x06, 0x00,
	0x5a, 0x6f, 0x71, 0xb5, 0x0e, 0x5e, 0x39, 0x47, 0x2a, 0xd1, 0x3f, 0xcd, 0xcf, 0xe3, 0x17, 0xcf,
	0xaa, 0x10, 0x47, 0x97, 0xfe, 0x99, 0x6a, 0x1a, 0x8a, 0x18, 0x5d, 0xfa, 0x4a, 0x41, 0xf5, 0x4d,
	0xa5, 0xba, 0xef, 0x95, 0xef, 0xf1, 0x1e, 0x78, 0xf2, 0xfc, 0xa9, 0xf6, 0x6d, 0xed, 0x81, 0xf3,
	0xb3, 0xa9, 0xf6, 0xb5, 0x2f, 0x7e, 0xfc, 0xf3, 0x9f, 0xaf, 0xac, 0x5f, 0x7f, 0x30, 0xbe, 0x41,
	0x19, 0xdd, 0x62, 0xc9, 0xd4, 0x77, 0xeb, 0x14, 0x6f, 0x49, 0x2e, 0x27, 0x98, 0xd3, 0xc9, 0xfb,
	0x1f, 0xb3, 0x97, 0x37, 0x9c, 0xc7, 0xba, 0xda, 0xd8, 0x3a, 0xbc, 0x4f, 0xfe, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0x90, 0xaa, 0x2e, 0x39, 0xf1, 0x06, 0x00, 0x00,
}
