// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/extensions/clusters/dynamic_forward_proxy/v3/cluster.proto

package envoy_extensions_clusters_dynamic_forward_proxy_v3

import (
	v3 "github.com/cilium/proxy/go/envoy/extensions/common/dynamic_forward_proxy/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration for the dynamic forward proxy cluster. See the :ref:`architecture overview
// <arch_overview_http_dynamic_forward_proxy>` for more information.
// [#extension: envoy.clusters.dynamic_forward_proxy]
type ClusterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The DNS cache configuration that the cluster will attach to. Note this configuration must
	// match that of associated :ref:`dynamic forward proxy HTTP filter configuration
	// <envoy_api_field_extensions.filters.http.dynamic_forward_proxy.v3.FilterConfig.dns_cache_config>`.
	DnsCacheConfig *v3.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
	// If true allow the cluster configuration to disable the auto_sni and auto_san_validation options
	// in the :ref:`cluster's upstream_http_protocol_options
	// <envoy_api_field_config.cluster.v3.Cluster.upstream_http_protocol_options>`
	AllowInsecureClusterOptions bool `protobuf:"varint,2,opt,name=allow_insecure_cluster_options,json=allowInsecureClusterOptions,proto3" json:"allow_insecure_cluster_options,omitempty"`
}

func (x *ClusterConfig) Reset() {
	*x = ClusterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterConfig) ProtoMessage() {}

func (x *ClusterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterConfig.ProtoReflect.Descriptor instead.
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterConfig) GetDnsCacheConfig() *v3.DnsCacheConfig {
	if x != nil {
		return x.DnsCacheConfig
	}
	return nil
}

func (x *ClusterConfig) GetAllowInsecureClusterOptions() bool {
	if x != nil {
		return x.AllowInsecureClusterOptions
	}
	return false
}

var File_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto protoreflect.FileDescriptor

var file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDesc = []byte{
	0x0a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x1a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x74, 0x0a, 0x10, 0x64, 0x6e, 0x73, 0x5f, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x40, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x33, 0x2e, 0x44, 0x6e, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x64, 0x6e, 0x73,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a, 0x1e, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x1b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x3a, 0x47, 0x9a, 0xc5, 0x88, 0x1e, 0x42, 0x0a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x5a, 0x0a, 0x40, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x0c, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescOnce sync.Once
	file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescData = file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDesc
)

func file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescGZIP() []byte {
	file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescData)
	})
	return file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDescData
}

var file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_goTypes = []interface{}{
	(*ClusterConfig)(nil),     // 0: envoy.extensions.clusters.dynamic_forward_proxy.v3.ClusterConfig
	(*v3.DnsCacheConfig)(nil), // 1: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig
}
var file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.clusters.dynamic_forward_proxy.v3.ClusterConfig.dns_cache_config:type_name -> envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_init() }
func file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_init() {
	if File_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_msgTypes,
	}.Build()
	File_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto = out.File
	file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_rawDesc = nil
	file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_goTypes = nil
	file_envoy_extensions_clusters_dynamic_forward_proxy_v3_cluster_proto_depIdxs = nil
}
