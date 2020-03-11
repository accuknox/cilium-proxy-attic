// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/wasm/v2/wasm.proto

package envoy_config_filter_network_wasm_v2

import (
	fmt "fmt"
	v2 "github.com/cilium/proxy/go/envoy/config/wasm/v2"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type Wasm struct {
	// General Plugin configuration.
	Config               *v2.PluginConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Wasm) Reset()         { *m = Wasm{} }
func (m *Wasm) String() string { return proto.CompactTextString(m) }
func (*Wasm) ProtoMessage()    {}
func (*Wasm) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b24e213c46f788, []int{0}
}

func (m *Wasm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Wasm.Unmarshal(m, b)
}
func (m *Wasm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Wasm.Marshal(b, m, deterministic)
}
func (m *Wasm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wasm.Merge(m, src)
}
func (m *Wasm) XXX_Size() int {
	return xxx_messageInfo_Wasm.Size(m)
}
func (m *Wasm) XXX_DiscardUnknown() {
	xxx_messageInfo_Wasm.DiscardUnknown(m)
}

var xxx_messageInfo_Wasm proto.InternalMessageInfo

func (m *Wasm) GetConfig() *v2.PluginConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*Wasm)(nil), "envoy.config.filter.network.wasm.v2.Wasm")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/wasm/v2/wasm.proto", fileDescriptor_67b24e213c46f788)
}

var fileDescriptor_67b24e213c46f788 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2f, 0x4f, 0x2c, 0xce, 0xd5, 0x2f, 0x33, 0x02,
	0xd3, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xca, 0x60, 0xf5, 0x7a, 0x10, 0xf5, 0x7a, 0x10,
	0xf5, 0x7a, 0x50, 0xf5, 0x7a, 0x60, 0x75, 0x65, 0x46, 0x52, 0xf2, 0x28, 0x86, 0x62, 0x9a, 0x22,
	0x25, 0x57, 0x9a, 0x52, 0x90, 0xa8, 0x9f, 0x98, 0x97, 0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f,
	0x57, 0xac, 0x9f, 0x9b, 0x99, 0x5e, 0x94, 0x58, 0x92, 0x0a, 0x95, 0x17, 0x2f, 0x4b, 0xcc, 0xc9,
	0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x12, 0x4a, 0x4e, 0x5c, 0x2c, 0xe1, 0x89, 0xc5,
	0xb9, 0x42, 0x56, 0x5c, 0x6c, 0x10, 0xd3, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x94, 0xf4,
	0x50, 0xdc, 0x05, 0x75, 0x88, 0x5e, 0x40, 0x4e, 0x69, 0x7a, 0x66, 0x9e, 0x33, 0x58, 0x30, 0x08,
	0xaa, 0xc3, 0xa9, 0xe0, 0xd3, 0x8c, 0x7f, 0xfd, 0xac, 0x5a, 0x42, 0x1a, 0x10, 0x2d, 0xa9, 0x15,
	0x25, 0xa9, 0x79, 0xc5, 0x20, 0x47, 0x40, 0xbd, 0x53, 0x8c, 0xe6, 0x1f, 0x63, 0x2e, 0xc3, 0xcc,
	0x7c, 0x88, 0xf9, 0x05, 0x45, 0xf9, 0x15, 0x95, 0x7a, 0x44, 0x04, 0x81, 0x13, 0x27, 0xc8, 0x99,
	0x01, 0x20, 0x37, 0x07, 0x30, 0x26, 0xb1, 0x81, 0x1d, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x41, 0x47, 0x46, 0x73, 0x6d, 0x01, 0x00, 0x00,
}
