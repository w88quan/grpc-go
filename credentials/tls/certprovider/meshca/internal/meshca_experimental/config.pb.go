// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/tls/provider/meshca/experimental/config.proto

// NOTE: This proto will very likely move to a different namespace and a
// different git repo in the future.

package meshca_experimental

import (
	fmt "fmt"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Type of key to be embedded in CSRs sent to the MeshCA.
type GoogleMeshCaConfig_KeyType int32

const (
	GoogleMeshCaConfig_KEY_TYPE_UNKNOWN GoogleMeshCaConfig_KeyType = 0
	GoogleMeshCaConfig_KEY_TYPE_RSA     GoogleMeshCaConfig_KeyType = 1
)

var GoogleMeshCaConfig_KeyType_name = map[int32]string{
	0: "KEY_TYPE_UNKNOWN",
	1: "KEY_TYPE_RSA",
}

var GoogleMeshCaConfig_KeyType_value = map[string]int32{
	"KEY_TYPE_UNKNOWN": 0,
	"KEY_TYPE_RSA":     1,
}

func (x GoogleMeshCaConfig_KeyType) String() string {
	return proto.EnumName(GoogleMeshCaConfig_KeyType_name, int32(x))
}

func (GoogleMeshCaConfig_KeyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f29f6dba7c86653, []int{0, 0}
}

// GoogleMeshCaConfig contains all configuration parameters required by the
// MeshCA CertificateProvider plugin implementation.
type GoogleMeshCaConfig struct {
	// GoogleMeshCA server endpoint to get CSRs signed via the *CreateCertificate*
	// unary call. This must have :ref:`api_type
	// <envoy_api_field_config.core.v3.ApiConfigSource.api_type>` :ref:`GRPC
	// <envoy_api_enum_value_config.core.v3.ApiConfigSource.ApiType.GRPC>`.
	// STS based call credentials need to be supplied in :ref:`call_credentials
	// <envoy_api_field_config.core.v3.GrpcService.GoogleGrpc.call_credentials>`.
	//
	// If :ref:`timeout envoy_api_field_config.core.v3.GrpcService.timeout` is
	// left unspecified, a default value of 10s will be used.
	Server *v3.ApiConfigSource `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	// Certificate lifetime to request in CSRs sent to the MeshCA.
	//
	// A default value of 24h will be used if left unspecified.
	CertificateLifetime *duration.Duration `protobuf:"bytes,2,opt,name=certificate_lifetime,json=certificateLifetime,proto3" json:"certificate_lifetime,omitempty"`
	// How long before certificate expiration should the certificate be renewed.
	//
	// A default value of 12h will be used if left unspecified.
	RenewalGracePeriod *duration.Duration `protobuf:"bytes,3,opt,name=renewal_grace_period,json=renewalGracePeriod,proto3" json:"renewal_grace_period,omitempty"`
	// Type of key.
	//
	// RSA keys will be used if left unspecified.
	KeyType GoogleMeshCaConfig_KeyType `protobuf:"varint,4,opt,name=key_type,json=keyType,proto3,enum=grpc.tls.provider.meshca.experimental.GoogleMeshCaConfig_KeyType" json:"key_type,omitempty"`
	// Size of the key in bits.
	//
	// 2048 bit keys will be used if left unspecified.
	KeySize uint32 `protobuf:"varint,5,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	// GCE location (region/zone) where the workload is located.
	//
	// GCE/GKE Metadata Server will be contacted if left unspecified.
	Location             string   `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoogleMeshCaConfig) Reset()         { *m = GoogleMeshCaConfig{} }
func (m *GoogleMeshCaConfig) String() string { return proto.CompactTextString(m) }
func (*GoogleMeshCaConfig) ProtoMessage()    {}
func (*GoogleMeshCaConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f29f6dba7c86653, []int{0}
}

func (m *GoogleMeshCaConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleMeshCaConfig.Unmarshal(m, b)
}
func (m *GoogleMeshCaConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleMeshCaConfig.Marshal(b, m, deterministic)
}
func (m *GoogleMeshCaConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleMeshCaConfig.Merge(m, src)
}
func (m *GoogleMeshCaConfig) XXX_Size() int {
	return xxx_messageInfo_GoogleMeshCaConfig.Size(m)
}
func (m *GoogleMeshCaConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleMeshCaConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleMeshCaConfig proto.InternalMessageInfo

func (m *GoogleMeshCaConfig) GetServer() *v3.ApiConfigSource {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *GoogleMeshCaConfig) GetCertificateLifetime() *duration.Duration {
	if m != nil {
		return m.CertificateLifetime
	}
	return nil
}

func (m *GoogleMeshCaConfig) GetRenewalGracePeriod() *duration.Duration {
	if m != nil {
		return m.RenewalGracePeriod
	}
	return nil
}

func (m *GoogleMeshCaConfig) GetKeyType() GoogleMeshCaConfig_KeyType {
	if m != nil {
		return m.KeyType
	}
	return GoogleMeshCaConfig_KEY_TYPE_UNKNOWN
}

func (m *GoogleMeshCaConfig) GetKeySize() uint32 {
	if m != nil {
		return m.KeySize
	}
	return 0
}

func (m *GoogleMeshCaConfig) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func init() {
	proto.RegisterEnum("grpc.tls.provider.meshca.experimental.GoogleMeshCaConfig_KeyType", GoogleMeshCaConfig_KeyType_name, GoogleMeshCaConfig_KeyType_value)
	proto.RegisterType((*GoogleMeshCaConfig)(nil), "grpc.tls.provider.meshca.experimental.GoogleMeshCaConfig")
}

func init() {
	proto.RegisterFile("grpc/tls/provider/meshca/experimental/config.proto", fileDescriptor_2f29f6dba7c86653)
}

var fileDescriptor_2f29f6dba7c86653 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x09, 0x83, 0x6e, 0x98, 0x3f, 0x2a, 0xa6, 0x87, 0xac, 0x07, 0x54, 0x4d, 0x9a, 0x94,
	0x93, 0x2d, 0xda, 0x33, 0x87, 0x6e, 0x4c, 0x3b, 0x74, 0x94, 0x2a, 0x1d, 0x9a, 0x86, 0x90, 0x22,
	0xcf, 0x7d, 0x9b, 0x59, 0x73, 0xed, 0xe8, 0x8d, 0x1b, 0xc8, 0x3e, 0x09, 0x9f, 0x84, 0xcf, 0x87,
	0xe2, 0x78, 0x53, 0x04, 0x07, 0x7a, 0xcb, 0x1b, 0xfb, 0xf7, 0x3c, 0x8f, 0x1e, 0xbf, 0x64, 0x9c,
	0x63, 0x21, 0xb9, 0xd3, 0x25, 0x2f, 0xd0, 0x56, 0x6a, 0x05, 0xc8, 0x37, 0x50, 0xde, 0x4a, 0xc1,
	0xe1, 0x67, 0x01, 0xa8, 0x36, 0x60, 0x9c, 0xd0, 0x5c, 0x5a, 0xb3, 0x56, 0x39, 0x2b, 0xd0, 0x3a,
	0x4b, 0x8f, 0x1b, 0x86, 0x39, 0x5d, 0xb2, 0x07, 0x86, 0xb5, 0x0c, 0xeb, 0x32, 0xc3, 0x04, 0x4c,
	0x65, 0xeb, 0x80, 0x72, 0x69, 0x11, 0x78, 0x35, 0x09, 0x63, 0x56, 0xda, 0x2d, 0x4a, 0x68, 0x05,
	0x87, 0xef, 0x73, 0x6b, 0x73, 0x0d, 0xdc, 0x4f, 0x37, 0xdb, 0x35, 0x5f, 0x6d, 0x51, 0x38, 0x65,
	0x4d, 0x7b, 0x7e, 0xf4, 0x7b, 0x8f, 0xd0, 0x73, 0x7f, 0xe5, 0x33, 0x94, 0xb7, 0xa7, 0xe2, 0xd4,
	0x6b, 0xd0, 0x8f, 0xa4, 0x57, 0x02, 0x56, 0x80, 0x71, 0x34, 0x8a, 0x92, 0x97, 0xe3, 0x63, 0xe6,
	0x1d, 0x59, 0x08, 0xdb, 0x38, 0xb2, 0x6a, 0xc2, 0xa6, 0x85, 0x6a, 0x81, 0xa5, 0xf7, 0x4c, 0x03,
	0x44, 0x2f, 0xc8, 0x40, 0x02, 0x3a, 0xb5, 0x56, 0x52, 0x38, 0xc8, 0xb4, 0x5a, 0x83, 0x53, 0x1b,
	0x88, 0x9f, 0x7a, 0xb1, 0x43, 0xd6, 0x86, 0x62, 0x0f, 0xa1, 0xd8, 0xa7, 0x10, 0x2a, 0x7d, 0xd7,
	0xc1, 0x2e, 0x02, 0x45, 0x67, 0x64, 0x80, 0x60, 0xe0, 0x87, 0xd0, 0x59, 0x8e, 0x42, 0x42, 0xd6,
	0x34, 0x61, 0x57, 0xf1, 0xde, 0xff, 0xd4, 0x68, 0xc0, 0xce, 0x1b, 0x6a, 0xe1, 0x21, 0xfa, 0x9d,
	0x1c, 0xdc, 0x41, 0x9d, 0xb9, 0xba, 0x80, 0xf8, 0xd9, 0x28, 0x4a, 0xde, 0x8c, 0xa7, 0x6c, 0xa7,
	0xd2, 0xd9, 0xbf, 0x35, 0xb1, 0x19, 0xd4, 0x97, 0x75, 0x01, 0xe9, 0xfe, 0x5d, 0xfb, 0x41, 0x0f,
	0x5b, 0xf5, 0x52, 0xdd, 0x43, 0xfc, 0x7c, 0x14, 0x25, 0xaf, 0xfd, 0xd1, 0x52, 0xdd, 0x03, 0x1d,
	0x92, 0x03, 0x6d, 0xa5, 0x0f, 0x16, 0xf7, 0x46, 0x51, 0xf2, 0x22, 0x7d, 0x9c, 0x8f, 0x3e, 0x90,
	0xfd, 0x20, 0x45, 0x07, 0xa4, 0x3f, 0x3b, 0xbb, 0xce, 0x2e, 0xaf, 0x17, 0x67, 0xd9, 0xd7, 0xf9,
	0x6c, 0xfe, 0xe5, 0x6a, 0xde, 0x7f, 0x42, 0xfb, 0xe4, 0xd5, 0xe3, 0xdf, 0x74, 0x39, 0xed, 0x47,
	0x27, 0xbf, 0x22, 0x92, 0x28, 0xbb, 0x5b, 0xf4, 0x93, 0xb7, 0xdd, 0xd4, 0x8b, 0xa6, 0xa7, 0x45,
	0xf4, 0xed, 0x2a, 0xf4, 0x96, 0x5b, 0x2d, 0x4c, 0xce, 0x2c, 0xe6, 0xdc, 0x6f, 0xac, 0x44, 0x58,
	0x81, 0x71, 0x4a, 0xe8, 0xd2, 0x6f, 0x6f, 0xf3, 0x24, 0x7f, 0x6f, 0xb0, 0x32, 0x0e, 0xd0, 0x08,
	0x1d, 0xe6, 0xac, 0xeb, 0x76, 0xd3, 0xf3, 0x2f, 0x31, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x72,
	0x44, 0x68, 0xc5, 0x01, 0x03, 0x00, 0x00,
}
