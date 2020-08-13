// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: cilium/api/bpf_metadata.proto

package cilium

import (
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

type BpfMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// File system root for bpf. Defaults to "/sys/fs/bpf" if left empty.
	BpfRoot string `protobuf:"bytes,1,opt,name=bpf_root,json=bpfRoot,proto3" json:"bpf_root,omitempty"`
	// 'true' if the filter is on ingress listener, 'false' for egress listener.
	IsIngress bool `protobuf:"varint,2,opt,name=is_ingress,json=isIngress,proto3" json:"is_ingress,omitempty"`
	// Use of the original source address requires kernel datapath support which
	// may or may not be available. 'true' if original source address
	// functionality is availeble. Original source address use may still be
	// skipped in scenarios where it is knows to not work.
	MayUseOriginalSourceAddress bool `protobuf:"varint,3,opt,name=may_use_original_source_address,json=mayUseOriginalSourceAddress,proto3" json:"may_use_original_source_address,omitempty"`
}

func (x *BpfMetadata) Reset() {
	*x = BpfMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cilium_api_bpf_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BpfMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BpfMetadata) ProtoMessage() {}

func (x *BpfMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_cilium_api_bpf_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BpfMetadata.ProtoReflect.Descriptor instead.
func (*BpfMetadata) Descriptor() ([]byte, []int) {
	return file_cilium_api_bpf_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *BpfMetadata) GetBpfRoot() string {
	if x != nil {
		return x.BpfRoot
	}
	return ""
}

func (x *BpfMetadata) GetIsIngress() bool {
	if x != nil {
		return x.IsIngress
	}
	return false
}

func (x *BpfMetadata) GetMayUseOriginalSourceAddress() bool {
	if x != nil {
		return x.MayUseOriginalSourceAddress
	}
	return false
}

var File_cilium_api_bpf_metadata_proto protoreflect.FileDescriptor

var file_cilium_api_bpf_metadata_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x70, 0x66,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x22, 0x8d, 0x01, 0x0a, 0x0b, 0x42, 0x70, 0x66, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x70, 0x66, 0x5f, 0x72,
	0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x70, 0x66, 0x52, 0x6f,
	0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x44, 0x0a, 0x1f, 0x6d, 0x61, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x6d, 0x61, 0x79, 0x55,
	0x73, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cilium_api_bpf_metadata_proto_rawDescOnce sync.Once
	file_cilium_api_bpf_metadata_proto_rawDescData = file_cilium_api_bpf_metadata_proto_rawDesc
)

func file_cilium_api_bpf_metadata_proto_rawDescGZIP() []byte {
	file_cilium_api_bpf_metadata_proto_rawDescOnce.Do(func() {
		file_cilium_api_bpf_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_cilium_api_bpf_metadata_proto_rawDescData)
	})
	return file_cilium_api_bpf_metadata_proto_rawDescData
}

var file_cilium_api_bpf_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cilium_api_bpf_metadata_proto_goTypes = []interface{}{
	(*BpfMetadata)(nil), // 0: cilium.BpfMetadata
}
var file_cilium_api_bpf_metadata_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cilium_api_bpf_metadata_proto_init() }
func file_cilium_api_bpf_metadata_proto_init() {
	if File_cilium_api_bpf_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cilium_api_bpf_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BpfMetadata); i {
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
			RawDescriptor: file_cilium_api_bpf_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cilium_api_bpf_metadata_proto_goTypes,
		DependencyIndexes: file_cilium_api_bpf_metadata_proto_depIdxs,
		MessageInfos:      file_cilium_api_bpf_metadata_proto_msgTypes,
	}.Build()
	File_cilium_api_bpf_metadata_proto = out.File
	file_cilium_api_bpf_metadata_proto_rawDesc = nil
	file_cilium_api_bpf_metadata_proto_goTypes = nil
	file_cilium_api_bpf_metadata_proto_depIdxs = nil
}
