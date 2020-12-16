// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/service/trace/v3/trace_service.proto

package envoy_service_trace_v3

import (
	context "context"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type StreamTracesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamTracesResponse) Reset() {
	*x = StreamTracesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_trace_v3_trace_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamTracesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamTracesResponse) ProtoMessage() {}

func (x *StreamTracesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_trace_v3_trace_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamTracesResponse.ProtoReflect.Descriptor instead.
func (*StreamTracesResponse) Descriptor() ([]byte, []int) {
	return file_envoy_service_trace_v3_trace_service_proto_rawDescGZIP(), []int{0}
}

type StreamTracesMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier data effectively is a structured metadata.
	// As a performance optimization this will only be sent in the first message
	// on the stream.
	Identifier *StreamTracesMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// A list of Span entries
	Spans []*v1.Span `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
}

func (x *StreamTracesMessage) Reset() {
	*x = StreamTracesMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_trace_v3_trace_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamTracesMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamTracesMessage) ProtoMessage() {}

func (x *StreamTracesMessage) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_trace_v3_trace_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamTracesMessage.ProtoReflect.Descriptor instead.
func (*StreamTracesMessage) Descriptor() ([]byte, []int) {
	return file_envoy_service_trace_v3_trace_service_proto_rawDescGZIP(), []int{1}
}

func (x *StreamTracesMessage) GetIdentifier() *StreamTracesMessage_Identifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *StreamTracesMessage) GetSpans() []*v1.Span {
	if x != nil {
		return x.Spans
	}
	return nil
}

type StreamTracesMessage_Identifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The node sending the access log messages over the stream.
	Node *v3.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *StreamTracesMessage_Identifier) Reset() {
	*x = StreamTracesMessage_Identifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_trace_v3_trace_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamTracesMessage_Identifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamTracesMessage_Identifier) ProtoMessage() {}

func (x *StreamTracesMessage_Identifier) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_trace_v3_trace_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamTracesMessage_Identifier.ProtoReflect.Descriptor instead.
func (*StreamTracesMessage_Identifier) Descriptor() ([]byte, []int) {
	return file_envoy_service_trace_v3_trace_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *StreamTracesMessage_Identifier) GetNode() *v3.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

var File_envoy_service_trace_v3_trace_service_proto protoreflect.FileDescriptor

var file_envoy_service_trace_v3_trace_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3a, 0x32, 0x9a,
	0xc5, 0x88, 0x1e, 0x2d, 0x0a, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xde, 0x02, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x56, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61,
	0x6e, 0x52, 0x05, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x1a, 0x84, 0x01, 0x0a, 0x0a, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x3a, 0x3c, 0x9a, 0xc5, 0x88, 0x1e, 0x37, 0x0a, 0x35, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x3a,
	0x31, 0x9a, 0xc5, 0x88, 0x1e, 0x2c, 0x0a, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x7d, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x73, 0x12, 0x2b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x42, 0x46, 0x0a, 0x24, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x11, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x88, 0x01,
	0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_service_trace_v3_trace_service_proto_rawDescOnce sync.Once
	file_envoy_service_trace_v3_trace_service_proto_rawDescData = file_envoy_service_trace_v3_trace_service_proto_rawDesc
)

func file_envoy_service_trace_v3_trace_service_proto_rawDescGZIP() []byte {
	file_envoy_service_trace_v3_trace_service_proto_rawDescOnce.Do(func() {
		file_envoy_service_trace_v3_trace_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_trace_v3_trace_service_proto_rawDescData)
	})
	return file_envoy_service_trace_v3_trace_service_proto_rawDescData
}

var file_envoy_service_trace_v3_trace_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_service_trace_v3_trace_service_proto_goTypes = []interface{}{
	(*StreamTracesResponse)(nil),           // 0: envoy.service.trace.v3.StreamTracesResponse
	(*StreamTracesMessage)(nil),            // 1: envoy.service.trace.v3.StreamTracesMessage
	(*StreamTracesMessage_Identifier)(nil), // 2: envoy.service.trace.v3.StreamTracesMessage.Identifier
	(*v1.Span)(nil),                        // 3: opencensus.proto.trace.v1.Span
	(*v3.Node)(nil),                        // 4: envoy.config.core.v3.Node
}
var file_envoy_service_trace_v3_trace_service_proto_depIdxs = []int32{
	2, // 0: envoy.service.trace.v3.StreamTracesMessage.identifier:type_name -> envoy.service.trace.v3.StreamTracesMessage.Identifier
	3, // 1: envoy.service.trace.v3.StreamTracesMessage.spans:type_name -> opencensus.proto.trace.v1.Span
	4, // 2: envoy.service.trace.v3.StreamTracesMessage.Identifier.node:type_name -> envoy.config.core.v3.Node
	1, // 3: envoy.service.trace.v3.TraceService.StreamTraces:input_type -> envoy.service.trace.v3.StreamTracesMessage
	0, // 4: envoy.service.trace.v3.TraceService.StreamTraces:output_type -> envoy.service.trace.v3.StreamTracesResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_service_trace_v3_trace_service_proto_init() }
func file_envoy_service_trace_v3_trace_service_proto_init() {
	if File_envoy_service_trace_v3_trace_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_trace_v3_trace_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamTracesResponse); i {
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
		file_envoy_service_trace_v3_trace_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamTracesMessage); i {
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
		file_envoy_service_trace_v3_trace_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamTracesMessage_Identifier); i {
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
			RawDescriptor: file_envoy_service_trace_v3_trace_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_trace_v3_trace_service_proto_goTypes,
		DependencyIndexes: file_envoy_service_trace_v3_trace_service_proto_depIdxs,
		MessageInfos:      file_envoy_service_trace_v3_trace_service_proto_msgTypes,
	}.Build()
	File_envoy_service_trace_v3_trace_service_proto = out.File
	file_envoy_service_trace_v3_trace_service_proto_rawDesc = nil
	file_envoy_service_trace_v3_trace_service_proto_goTypes = nil
	file_envoy_service_trace_v3_trace_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceServiceClient interface {
	// Envoy will connect and send StreamTracesMessage messages forever. It does
	// not expect any response to be sent as nothing would be done in the case
	// of failure.
	StreamTraces(ctx context.Context, opts ...grpc.CallOption) (TraceService_StreamTracesClient, error)
}

type traceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTraceServiceClient(cc grpc.ClientConnInterface) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) StreamTraces(ctx context.Context, opts ...grpc.CallOption) (TraceService_StreamTracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceService_serviceDesc.Streams[0], "/envoy.service.trace.v3.TraceService/StreamTraces", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceServiceStreamTracesClient{stream}
	return x, nil
}

type TraceService_StreamTracesClient interface {
	Send(*StreamTracesMessage) error
	CloseAndRecv() (*StreamTracesResponse, error)
	grpc.ClientStream
}

type traceServiceStreamTracesClient struct {
	grpc.ClientStream
}

func (x *traceServiceStreamTracesClient) Send(m *StreamTracesMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceServiceStreamTracesClient) CloseAndRecv() (*StreamTracesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamTracesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceServiceServer is the server API for TraceService service.
type TraceServiceServer interface {
	// Envoy will connect and send StreamTracesMessage messages forever. It does
	// not expect any response to be sent as nothing would be done in the case
	// of failure.
	StreamTraces(TraceService_StreamTracesServer) error
}

// UnimplementedTraceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTraceServiceServer struct {
}

func (*UnimplementedTraceServiceServer) StreamTraces(TraceService_StreamTracesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTraces not implemented")
}

func RegisterTraceServiceServer(s *grpc.Server, srv TraceServiceServer) {
	s.RegisterService(&_TraceService_serviceDesc, srv)
}

func _TraceService_StreamTraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceServiceServer).StreamTraces(&traceServiceStreamTracesServer{stream})
}

type TraceService_StreamTracesServer interface {
	SendAndClose(*StreamTracesResponse) error
	Recv() (*StreamTracesMessage, error)
	grpc.ServerStream
}

type traceServiceStreamTracesServer struct {
	grpc.ServerStream
}

func (x *traceServiceStreamTracesServer) SendAndClose(m *StreamTracesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceServiceStreamTracesServer) Recv() (*StreamTracesMessage, error) {
	m := new(StreamTracesMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.trace.v3.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTraces",
			Handler:       _TraceService_StreamTraces_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/trace/v3/trace_service.proto",
}
