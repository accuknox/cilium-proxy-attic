// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v2alpha1/route.proto

package v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	route "github.com/cilium/proxy/go/envoy/api/v2/route"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/lyft/protoc-gen-validate/validate"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// [#comment:next free field: 3]
type RouteConfiguration struct {
	// The name of the route configuration. Reserved for future use in asynchronous route discovery.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of routes that will be matched, in order, against incoming requests. The first route
	// that matches will be used.
	Routes               []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

// [#comment:next free field: 3]
type Route struct {
	// Route matching prarameters.
	Match *RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// Route request to some upstream cluster.
	Route                *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

// [#comment:next free field: 5]
type RouteMatch struct {
	// Types that are valid to be assigned to MatchSpecifier:
	//	*RouteMatch_MethodName
	//	*RouteMatch_ServiceName
	MatchSpecifier isRouteMatch_MatchSpecifier `protobuf_oneof:"match_specifier"`
	// Inverts whatever matching is done in the :ref:`method_name
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteMatch.method_name>` or
	// :ref:`service_name
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteMatch.service_name>` fields.
	// Cannot be combined with wildcard matching as that would result in routes never being matched.
	//
	// .. note::
	//
	//   This does not invert matching done as part of the :ref:`headers field
	//   <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteMatch.headers>` field. To
	//   invert header matching, see :ref:`invert_match
	//   <envoy_api_field_route.HeaderMatcher.invert_match>`.
	Invert bool `protobuf:"varint,3,opt,name=invert,proto3" json:"invert,omitempty"`
	// Specifies a set of headers that the route should match on. The router will check the request’s
	// headers against all the specified headers in the route config. A match will happen if all the
	// headers in the route are present in the request with the same values (or based on presence if
	// the value field is not in the config). Note that this only applies for Thrift transports and/or
	// protocols that support headers.
	Headers              []*route.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{2}
}

func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (m *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(m, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

type isRouteMatch_MatchSpecifier interface {
	isRouteMatch_MatchSpecifier()
}

type RouteMatch_MethodName struct {
	MethodName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3,oneof"`
}

type RouteMatch_ServiceName struct {
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3,oneof"`
}

func (*RouteMatch_MethodName) isRouteMatch_MatchSpecifier() {}

func (*RouteMatch_ServiceName) isRouteMatch_MatchSpecifier() {}

func (m *RouteMatch) GetMatchSpecifier() isRouteMatch_MatchSpecifier {
	if m != nil {
		return m.MatchSpecifier
	}
	return nil
}

func (m *RouteMatch) GetMethodName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_MethodName); ok {
		return x.MethodName
	}
	return ""
}

func (m *RouteMatch) GetServiceName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_ServiceName); ok {
		return x.ServiceName
	}
	return ""
}

func (m *RouteMatch) GetInvert() bool {
	if m != nil {
		return m.Invert
	}
	return false
}

func (m *RouteMatch) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RouteMatch) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RouteMatch_OneofMarshaler, _RouteMatch_OneofUnmarshaler, _RouteMatch_OneofSizer, []interface{}{
		(*RouteMatch_MethodName)(nil),
		(*RouteMatch_ServiceName)(nil),
	}
}

func _RouteMatch_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RouteMatch)
	// match_specifier
	switch x := m.MatchSpecifier.(type) {
	case *RouteMatch_MethodName:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.MethodName)
	case *RouteMatch_ServiceName:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ServiceName)
	case nil:
	default:
		return fmt.Errorf("RouteMatch.MatchSpecifier has unexpected type %T", x)
	}
	return nil
}

func _RouteMatch_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RouteMatch)
	switch tag {
	case 1: // match_specifier.method_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.MatchSpecifier = &RouteMatch_MethodName{x}
		return true, err
	case 2: // match_specifier.service_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.MatchSpecifier = &RouteMatch_ServiceName{x}
		return true, err
	default:
		return false, nil
	}
}

func _RouteMatch_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RouteMatch)
	// match_specifier
	switch x := m.MatchSpecifier.(type) {
	case *RouteMatch_MethodName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.MethodName)))
		n += len(x.MethodName)
	case *RouteMatch_ServiceName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ServiceName)))
		n += len(x.ServiceName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// [#comment:next free field: 5]
type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	ClusterSpecifier isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	// Optional endpoint metadata match criteria used by the subset load balancer. Only endpoints in
	// the upstream cluster with metadata matching what is set in this field will be considered.
	// Note that this will be merged with what's provided in :ref: `WeightedCluster.MetadataMatch
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.WeightedCluster.ClusterWeight.metadata_match>`,
	// with values there taking precedence. Keys and values should be provided under the "envoy.lb"
	// metadata key.
	MetadataMatch *core.Metadata `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	// Specifies a set of rate limit configurations that could be applied to the route.
	// N.B. Thrift service or method name matching can be achieved by specifying a RequestHeaders
	// action with the header name ":method-name".
	RateLimits           []*route.RateLimit `protobuf:"bytes,4,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{3}
}

func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (m *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(m, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedClusters() *WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *RouteAction) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *RouteAction) GetRateLimits() []*route.RateLimit {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RouteAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RouteAction_OneofMarshaler, _RouteAction_OneofUnmarshaler, _RouteAction_OneofSizer, []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
	}
}

func _RouteAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RouteAction)
	// cluster_specifier
	switch x := m.ClusterSpecifier.(type) {
	case *RouteAction_Cluster:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Cluster)
	case *RouteAction_WeightedClusters:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WeightedClusters); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RouteAction.ClusterSpecifier has unexpected type %T", x)
	}
	return nil
}

func _RouteAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RouteAction)
	switch tag {
	case 1: // cluster_specifier.cluster
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ClusterSpecifier = &RouteAction_Cluster{x}
		return true, err
	case 2: // cluster_specifier.weighted_clusters
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WeightedCluster)
		err := b.DecodeMessage(msg)
		m.ClusterSpecifier = &RouteAction_WeightedClusters{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RouteAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RouteAction)
	// cluster_specifier
	switch x := m.ClusterSpecifier.(type) {
	case *RouteAction_Cluster:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Cluster)))
		n += len(x.Cluster)
	case *RouteAction_WeightedClusters:
		s := proto.Size(x.WeightedClusters)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Allows for specification of multiple upstream clusters along with weights that indicate the
// percentage of traffic to be forwarded to each cluster. The router selects an upstream cluster
// based on these weights.
type WeightedCluster struct {
	// Specifies one or more upstream clusters associated with the route.
	Clusters             []*WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *WeightedCluster) Reset()         { *m = WeightedCluster{} }
func (m *WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster) ProtoMessage()    {}
func (*WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{4}
}

func (m *WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster.Unmarshal(m, b)
}
func (m *WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster.Merge(m, src)
}
func (m *WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster.Size(m)
}
func (m *WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster proto.InternalMessageInfo

func (m *WeightedCluster) GetClusters() []*WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type WeightedCluster_ClusterWeight struct {
	// Name of the upstream cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When a request matches the route, the choice of an upstream cluster is determined by its
	// weight. The sum of weights across all entries in the clusters array determines the total
	// weight.
	Weight *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
	// Optional endpoint metadata match criteria used by the subset load balancer. Only endpoints in
	// the upstream cluster with metadata matching what is set in this field, combined with what's
	// provided in :ref: `RouteAction's metadata_match
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteAction.metadata_match>`,
	// will be considered. Values here will take precedence. Keys and values should be provided
	// under the "envoy.lb" metadata key.
	MetadataMatch        *core.Metadata `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WeightedCluster_ClusterWeight) Reset()         { *m = WeightedCluster_ClusterWeight{} }
func (m *WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{4, 0}
}

func (m *WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Size(m)
}
func (m *WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WeightedCluster_ClusterWeight) GetWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *WeightedCluster_ClusterWeight) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.RouteAction")
	proto.RegisterType((*WeightedCluster)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.WeightedCluster")
	proto.RegisterType((*WeightedCluster_ClusterWeight)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.WeightedCluster.ClusterWeight")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/thrift_proxy/v2alpha1/route.proto", fileDescriptor_3de6ac0eae6369f5)
}

var fileDescriptor_3de6ac0eae6369f5 = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x8d, 0x37, 0x1f, 0x6d, 0x27, 0xf4, 0xcb, 0x42, 0x10, 0x85, 0xb6, 0xa4, 0x41, 0x48, 0x11,
	0x07, 0xaf, 0x9a, 0x5e, 0x90, 0x50, 0x8b, 0xd8, 0x5e, 0x82, 0x44, 0x11, 0xb2, 0x44, 0x91, 0xb8,
	0x44, 0xee, 0xc6, 0x49, 0x2c, 0x92, 0xf5, 0xe2, 0x75, 0x36, 0xed, 0x8d, 0x33, 0x47, 0xfe, 0x0a,
	0x17, 0xc4, 0xa9, 0x88, 0x03, 0x9c, 0xf9, 0x01, 0xdc, 0xfb, 0x2f, 0xd0, 0xda, 0xde, 0x36, 0xa1,
	0x5c, 0x28, 0x9c, 0x76, 0x76, 0x3c, 0x6f, 0x9e, 0xdf, 0x1b, 0xdb, 0xb0, 0xc7, 0xa3, 0x54, 0x9e,
	0xfa, 0xa1, 0x8c, 0xfa, 0x62, 0xe0, 0xf7, 0xc5, 0x48, 0x73, 0xe5, 0x47, 0x5c, 0x4f, 0xa5, 0x7a,
	0xe3, 0xeb, 0xa1, 0x12, 0x7d, 0xdd, 0x8d, 0x95, 0x3c, 0x39, 0xf5, 0xd3, 0x36, 0x1b, 0xc5, 0x43,
	0xb6, 0xe3, 0x2b, 0x39, 0xd1, 0x9c, 0xc4, 0x4a, 0x6a, 0x89, 0x77, 0x0c, 0x9c, 0x58, 0x38, 0xb1,
	0x70, 0xe2, 0xe0, 0x64, 0x16, 0x4e, 0x72, 0x78, 0x7d, 0xc3, 0x32, 0xb2, 0x58, 0xf8, 0x69, 0xdb,
	0x0f, 0xa5, 0xe2, 0xfe, 0x31, 0x4b, 0x5c, 0xc3, 0xfa, 0xd6, 0xdc, 0xaa, 0xa1, 0x9a, 0x25, 0xac,
	0x6f, 0x0d, 0xa4, 0x1c, 0x8c, 0xb8, 0x6f, 0xfe, 0x8e, 0x27, 0x7d, 0x7f, 0xaa, 0x58, 0x1c, 0x73,
	0x95, 0xb8, 0xf5, 0xdb, 0x29, 0x1b, 0x89, 0x1e, 0xd3, 0xdc, 0xcf, 0x03, 0xb7, 0x70, 0x73, 0x20,
	0x07, 0xd2, 0x84, 0x7e, 0x16, 0xd9, 0x6c, 0xf3, 0x1d, 0x02, 0x4c, 0xb3, 0xf6, 0x07, 0x46, 0xc1,
	0x44, 0x31, 0x2d, 0x64, 0x84, 0x31, 0x94, 0x22, 0x36, 0xe6, 0x35, 0xd4, 0x40, 0xad, 0x25, 0x6a,
	0x62, 0x7c, 0x04, 0x15, 0xb3, 0x91, 0xa4, 0xe6, 0x35, 0x8a, 0xad, 0x6a, 0xfb, 0x21, 0xf9, 0x6b,
	0xed, 0xc4, 0x50, 0x05, 0xa5, 0xef, 0x3f, 0xef, 0x16, 0xa8, 0xeb, 0xd6, 0xfc, 0x81, 0xa0, 0x6c,
	0xf2, 0x38, 0x84, 0xf2, 0x98, 0xe9, 0x70, 0x68, 0x68, 0xab, 0xed, 0xbd, 0xeb, 0x12, 0x1c, 0x66,
	0x4d, 0x82, 0x95, 0x8c, 0xe5, 0xf3, 0xf9, 0x59, 0xb1, 0xfc, 0x1e, 0x79, 0x6b, 0x88, 0xda, 0xde,
	0xb8, 0x07, 0x65, 0x43, 0x5c, 0xf3, 0x0c, 0xc9, 0xfe, 0x75, 0x49, 0x9e, 0x84, 0x99, 0x53, 0x57,
	0x59, 0x4c, 0xf3, 0xe6, 0x57, 0x04, 0x70, 0xb9, 0x17, 0xbc, 0x0d, 0xd5, 0x31, 0xd7, 0x43, 0xd9,
	0xeb, 0x5e, 0xda, 0xda, 0x29, 0x50, 0xb0, 0xc9, 0xe7, 0x99, 0xbd, 0xf7, 0xe0, 0x46, 0xc2, 0x55,
	0x2a, 0x42, 0x6e, 0x6b, 0x3c, 0x57, 0x53, 0x75, 0x59, 0x53, 0x74, 0x0b, 0x2a, 0x22, 0x4a, 0xb9,
	0xd2, 0xb5, 0x62, 0x03, 0xb5, 0x16, 0xa9, 0xfb, 0xc3, 0x8f, 0x60, 0x61, 0xc8, 0x59, 0x8f, 0xab,
	0xa4, 0x56, 0x32, 0xc3, 0xd9, 0x76, 0xb2, 0x58, 0x2c, 0x48, 0xda, 0x26, 0xf6, 0x04, 0x75, 0x4c,
	0x89, 0xd9, 0x11, 0x57, 0x34, 0x47, 0x04, 0x35, 0x58, 0x35, 0xd6, 0x74, 0x93, 0x98, 0x87, 0xa2,
	0x2f, 0xb8, 0xc2, 0xe5, 0x4f, 0xe7, 0x67, 0x45, 0xd4, 0xfc, 0xe6, 0x41, 0x75, 0x46, 0x2c, 0xbe,
	0x0f, 0x0b, 0xe1, 0x68, 0x92, 0x68, 0xae, 0xac, 0x84, 0x60, 0x29, 0x53, 0x5e, 0x52, 0x5e, 0x03,
	0x75, 0x0a, 0x34, 0x5f, 0xc3, 0x6f, 0x61, 0x7d, 0xca, 0xc5, 0x60, 0xa8, 0x79, 0xaf, 0xeb, 0x72,
	0x89, 0xb3, 0x3b, 0xb8, 0x86, 0xdd, 0xaf, 0x5c, 0xaf, 0x03, 0xdb, 0xaa, 0x53, 0xa0, 0x6b, 0xd3,
	0xf9, 0x54, 0x82, 0x03, 0x58, 0x19, 0x73, 0xcd, 0x7a, 0x4c, 0xb3, 0xae, 0x3d, 0x43, 0x45, 0xc3,
	0x77, 0x67, 0xde, 0x87, 0xec, 0xb6, 0x91, 0x43, 0x57, 0x48, 0x97, 0x73, 0x88, 0x1d, 0xd2, 0x3e,
	0x54, 0x15, 0xd3, 0xbc, 0x3b, 0x12, 0x63, 0xa1, 0x73, 0x23, 0x37, 0xff, 0x64, 0x24, 0x65, 0x9a,
	0x3f, 0xcb, 0xaa, 0x28, 0xa8, 0x3c, 0x4c, 0x82, 0x3a, 0xac, 0x3b, 0xb5, 0x57, 0x9d, 0xfc, 0xe2,
	0xc1, 0xea, 0x6f, 0x3a, 0xf0, 0x09, 0x2c, 0x5e, 0xb8, 0x83, 0x0c, 0xd9, 0x8b, 0x7f, 0x77, 0x87,
	0xb8, 0xaf, 0x4d, 0x07, 0x60, 0x8e, 0xe6, 0x07, 0xe4, 0x2d, 0x22, 0x7a, 0xc1, 0x56, 0xff, 0x88,
	0x60, 0x79, 0xae, 0x0e, 0x6f, 0xce, 0x5e, 0xf8, 0x99, 0xb1, 0xba, 0xbb, 0xff, 0x18, 0x2a, 0xd6,
	0x72, 0x37, 0xc6, 0x0d, 0x62, 0x9f, 0x21, 0x92, 0x3f, 0x43, 0xe4, 0xe5, 0xd3, 0x48, 0xef, 0xb6,
	0x8f, 0xd8, 0x68, 0xc2, 0x1d, 0xfc, 0x81, 0xd7, 0x42, 0xd4, 0xc1, 0xfe, 0xc7, 0x7c, 0x82, 0xd2,
	0x6b, 0x2f, 0x6d, 0x1f, 0x57, 0x0c, 0xe5, 0xee, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x51,
	0xf0, 0x35, 0xb9, 0x05, 0x00, 0x00,
}