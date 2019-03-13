// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ratelimit.proto

package ratelimit

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

type RateLimit_Unit int32

const (
	RateLimit_UNKNOWN RateLimit_Unit = 0
	RateLimit_SECOND  RateLimit_Unit = 1
	RateLimit_MINUTE  RateLimit_Unit = 2
	RateLimit_HOUR    RateLimit_Unit = 3
	RateLimit_DAY     RateLimit_Unit = 4
)

var RateLimit_Unit_name = map[int32]string{
	0: "UNKNOWN",
	1: "SECOND",
	2: "MINUTE",
	3: "HOUR",
	4: "DAY",
}
var RateLimit_Unit_value = map[string]int32{
	"UNKNOWN": 0,
	"SECOND":  1,
	"MINUTE":  2,
	"HOUR":    3,
	"DAY":     4,
}

func (x RateLimit_Unit) String() string {
	return proto.EnumName(RateLimit_Unit_name, int32(x))
}
func (RateLimit_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_ee79f6399c38160c, []int{2, 0}
}

type RateLimitResponse_Code int32

const (
	RateLimitResponse_UNKNOWN    RateLimitResponse_Code = 0
	RateLimitResponse_OK         RateLimitResponse_Code = 1
	RateLimitResponse_OVER_LIMIT RateLimitResponse_Code = 2
)

var RateLimitResponse_Code_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "OVER_LIMIT",
}
var RateLimitResponse_Code_value = map[string]int32{
	"UNKNOWN":    0,
	"OK":         1,
	"OVER_LIMIT": 2,
}

func (x RateLimitResponse_Code) String() string {
	return proto.EnumName(RateLimitResponse_Code_name, int32(x))
}
func (RateLimitResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_ee79f6399c38160c, []int{3, 0}
}

// Main message for a rate limit request. The rate limit service is designed to be fully generic
// in the sense that it can operate on arbitrary hierarchical key/value pairs. The loaded
// configuration will parse the request and find the most specific limit to apply. In addition,
// a RateLimitRequest can contain multiple "descriptors" to limit on. When multiple descriptors
// are provided, the server will limit on *ALL* of them and return an OVER_LIMIT response if any
// of them are over limit. This enables more complex application level rate limiting scenarios
// if desired.
type RateLimitRequest struct {
	// All rate limit requests must specify a domain. This enables the configuration to be per
	// application without fear of overlap. E.g., "envoy".
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// All rate limit requests must specify at least one RateLimitDescriptor. Each descriptor is
	// processed by the service (see below). If any of the descriptors are over limit, the entire
	// request is considered to be over limit.
	Descriptors          []*RateLimitDescriptor `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RateLimitRequest) Reset()         { *m = RateLimitRequest{} }
func (m *RateLimitRequest) String() string { return proto.CompactTextString(m) }
func (*RateLimitRequest) ProtoMessage()    {}
func (*RateLimitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_ee79f6399c38160c, []int{0}
}
func (m *RateLimitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitRequest.Unmarshal(m, b)
}
func (m *RateLimitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitRequest.Marshal(b, m, deterministic)
}
func (dst *RateLimitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitRequest.Merge(dst, src)
}
func (m *RateLimitRequest) XXX_Size() int {
	return xxx_messageInfo_RateLimitRequest.Size(m)
}
func (m *RateLimitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitRequest proto.InternalMessageInfo

func (m *RateLimitRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitRequest) GetDescriptors() []*RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

// A RateLimitDescriptor is a list of hierarchical entries that are used by the service to
// determine the final rate limit key and overall allowed limit. Here are some examples of how
// they might be used for the domain "envoy".
// 1) ["authenticated": "false"], ["ip_address": "10.0.0.1"]
//    What it does: Limits all unauthenticated traffic for the IP address 10.0.0.1. The
//    configuration supplies a default limit for the ip_address field. If there is a desire to raise
//    the limit for 10.0.0.1 or block it entirely it can be specified directly in the
//    configuration.
// 2) ["authenticated": "false"], ["path": "/foo/bar"]
//    What it does: Limits all unauthenticated traffic globally for a specific path (or prefix if
//    configured that way in the service).
// 3) ["authenticated": "false"], ["path": "/foo/bar"], ["ip_address": "10.0.0.1"]
//    What it does: Limits unauthenticated traffic to a specific path for a specific IP address.
//    Like (1) we can raise/block specific IP addresses if we want with an override configuration.
// 4) ["authenticated": "true"], ["client_id": "foo"]
//    What it does: Limits all traffic for an authenticated client "foo"
// 5) ["authenticated": "true"], ["client_id": "foo"], ["path": "/foo/bar"]
//    What it does: Limits traffic to a specific path for an authenticated client "foo"
//
// The idea behind the API is that (1)/(2)/(3) and (4)/(5) can be sent in 1 request if desired.
// This enables building complex application scenarios with a generic backend.
type RateLimitDescriptor struct {
	Entries              []*RateLimitDescriptor_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RateLimitDescriptor) Reset()         { *m = RateLimitDescriptor{} }
func (m *RateLimitDescriptor) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptor) ProtoMessage()    {}
func (*RateLimitDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_ee79f6399c38160c, []int{1}
}
func (m *RateLimitDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitDescriptor.Unmarshal(m, b)
}
func (m *RateLimitDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitDescriptor.Marshal(b, m, deterministic)
}
func (dst *RateLimitDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptor.Merge(dst, src)
}
func (m *RateLimitDescriptor) XXX_Size() int {
	return xxx_messageInfo_RateLimitDescriptor.Size(m)
}
func (m *RateLimitDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptor proto.InternalMessageInfo

func (m *RateLimitDescriptor) GetEntries() []*RateLimitDescriptor_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type RateLimitDescriptor_Entry struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitDescriptor_Entry) Reset()         { *m = RateLimitDescriptor_Entry{} }
func (m *RateLimitDescriptor_Entry) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptor_Entry) ProtoMessage()    {}
func (*RateLimitDescriptor_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_ee79f6399c38160c, []int{1, 0}
}
func (m *RateLimitDescriptor_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Unmarshal(m, b)
}
func (m *RateLimitDescriptor_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Marshal(b, m, deterministic)
}
func (dst *RateLimitDescriptor_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptor_Entry.Merge(dst, src)
}
func (m *RateLimitDescriptor_Entry) XXX_Size() int {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Size(m)
}
func (m *RateLimitDescriptor_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptor_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptor_Entry proto.InternalMessageInfo

func (m *RateLimitDescriptor_Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RateLimitDescriptor_Entry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Defines an actual rate limit in terms of requests per unit of time and the unit itself.
type RateLimit struct {
	RequestsPerUnit      uint32         `protobuf:"varint,1,opt,name=requests_per_unit,json=requestsPerUnit,proto3" json:"requests_per_unit,omitempty"`
	Unit                 RateLimit_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=pb.lyft.ratelimit.RateLimit_Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_ee79f6399c38160c, []int{2}
}
func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (dst *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(dst, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetRequestsPerUnit() uint32 {
	if m != nil {
		return m.RequestsPerUnit
	}
	return 0
}

func (m *RateLimit) GetUnit() RateLimit_Unit {
	if m != nil {
		return m.Unit
	}
	return RateLimit_UNKNOWN
}

// A response from a ShouldRateLimit call.
type RateLimitResponse struct {
	// The overall response code which takes into account all of the descriptors that were passed
	// in the RateLimitRequest message.
	OverallCode RateLimitResponse_Code `protobuf:"varint,1,opt,name=overall_code,json=overallCode,proto3,enum=pb.lyft.ratelimit.RateLimitResponse_Code" json:"overall_code,omitempty"`
	// A list of DescriptorStatus messages which matches the length of the descriptor list passed
	// in the RateLimitRequest. This can be used by the caller to determine which individual
	// descriptors failed and/or what the currently configured limits are for all of them.
	Statuses             []*RateLimitResponse_DescriptorStatus `protobuf:"bytes,2,rep,name=statuses,proto3" json:"statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *RateLimitResponse) Reset()         { *m = RateLimitResponse{} }
func (m *RateLimitResponse) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse) ProtoMessage()    {}
func (*RateLimitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_ee79f6399c38160c, []int{3}
}
func (m *RateLimitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse.Unmarshal(m, b)
}
func (m *RateLimitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse.Marshal(b, m, deterministic)
}
func (dst *RateLimitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse.Merge(dst, src)
}
func (m *RateLimitResponse) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse.Size(m)
}
func (m *RateLimitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse proto.InternalMessageInfo

func (m *RateLimitResponse) GetOverallCode() RateLimitResponse_Code {
	if m != nil {
		return m.OverallCode
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse) GetStatuses() []*RateLimitResponse_DescriptorStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

type RateLimitResponse_DescriptorStatus struct {
	// The response code for an individual descriptor.
	Code RateLimitResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=pb.lyft.ratelimit.RateLimitResponse_Code" json:"code,omitempty"`
	// The current limit as configured by the server. Useful for debugging, etc.
	CurrentLimit *RateLimit `protobuf:"bytes,2,opt,name=current_limit,json=currentLimit,proto3" json:"current_limit,omitempty"`
	// The limit remaining in the current time unit.
	LimitRemaining       uint32   `protobuf:"varint,3,opt,name=limit_remaining,json=limitRemaining,proto3" json:"limit_remaining,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitResponse_DescriptorStatus) Reset()         { *m = RateLimitResponse_DescriptorStatus{} }
func (m *RateLimitResponse_DescriptorStatus) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_DescriptorStatus) ProtoMessage()    {}
func (*RateLimitResponse_DescriptorStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_ee79f6399c38160c, []int{3, 0}
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Unmarshal(m, b)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Marshal(b, m, deterministic)
}
func (dst *RateLimitResponse_DescriptorStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.Merge(dst, src)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Size(m)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_DescriptorStatus proto.InternalMessageInfo

func (m *RateLimitResponse_DescriptorStatus) GetCode() RateLimitResponse_Code {
	if m != nil {
		return m.Code
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse_DescriptorStatus) GetCurrentLimit() *RateLimit {
	if m != nil {
		return m.CurrentLimit
	}
	return nil
}

func (m *RateLimitResponse_DescriptorStatus) GetLimitRemaining() uint32 {
	if m != nil {
		return m.LimitRemaining
	}
	return 0
}

func init() {
	proto.RegisterType((*RateLimitRequest)(nil), "pb.lyft.ratelimit.RateLimitRequest")
	proto.RegisterType((*RateLimitDescriptor)(nil), "pb.lyft.ratelimit.RateLimitDescriptor")
	proto.RegisterType((*RateLimitDescriptor_Entry)(nil), "pb.lyft.ratelimit.RateLimitDescriptor.Entry")
	proto.RegisterType((*RateLimit)(nil), "pb.lyft.ratelimit.RateLimit")
	proto.RegisterType((*RateLimitResponse)(nil), "pb.lyft.ratelimit.RateLimitResponse")
	proto.RegisterType((*RateLimitResponse_DescriptorStatus)(nil), "pb.lyft.ratelimit.RateLimitResponse.DescriptorStatus")
	proto.RegisterEnum("pb.lyft.ratelimit.RateLimit_Unit", RateLimit_Unit_name, RateLimit_Unit_value)
	proto.RegisterEnum("pb.lyft.ratelimit.RateLimitResponse_Code", RateLimitResponse_Code_name, RateLimitResponse_Code_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RateLimitServiceClient is the client API for RateLimitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitServiceClient interface {
	// Determine whether rate limiting should take place.
	ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error)
}

type rateLimitServiceClient struct {
	cc *grpc.ClientConn
}

func NewRateLimitServiceClient(cc *grpc.ClientConn) RateLimitServiceClient {
	return &rateLimitServiceClient{cc}
}

func (c *rateLimitServiceClient) ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error) {
	out := new(RateLimitResponse)
	err := c.cc.Invoke(ctx, "/pb.lyft.ratelimit.RateLimitService/ShouldRateLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitServiceServer is the server API for RateLimitService service.
type RateLimitServiceServer interface {
	// Determine whether rate limiting should take place.
	ShouldRateLimit(context.Context, *RateLimitRequest) (*RateLimitResponse, error)
}

func RegisterRateLimitServiceServer(s *grpc.Server, srv RateLimitServiceServer) {
	s.RegisterService(&_RateLimitService_serviceDesc, srv)
}

func _RateLimitService_ShouldRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.lyft.ratelimit.RateLimitService/ShouldRateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, req.(*RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.lyft.ratelimit.RateLimitService",
	HandlerType: (*RateLimitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShouldRateLimit",
			Handler:    _RateLimitService_ShouldRateLimit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ratelimit.proto",
}

func init() { proto.RegisterFile("ratelimit.proto", fileDescriptor_ratelimit_ee79f6399c38160c) }

var fileDescriptor_ratelimit_ee79f6399c38160c = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x2f, 0x4d, 0x9a, 0x71, 0x9b, 0x6c, 0x16, 0x84, 0xac, 0x88, 0x87, 0x62, 0x10, 0x94,
	0x8b, 0x8c, 0x64, 0xd4, 0x47, 0x90, 0x4a, 0x13, 0xd4, 0xa8, 0xa9, 0x0d, 0x9b, 0x06, 0x04, 0x0f,
	0x58, 0x6e, 0x3c, 0x80, 0x85, 0x6b, 0x9b, 0xf5, 0x3a, 0x52, 0x7e, 0x82, 0x7f, 0xe0, 0x07, 0xf8,
	0x03, 0xfe, 0x0d, 0x79, 0xe3, 0xb8, 0xe1, 0xa2, 0x28, 0xea, 0xdb, 0xce, 0xcc, 0x39, 0x67, 0x67,
	0xe7, 0xcc, 0x42, 0x87, 0x07, 0x02, 0xe3, 0xe8, 0x32, 0x12, 0x76, 0xc6, 0x53, 0x91, 0xd2, 0x6e,
	0x76, 0x61, 0xc7, 0xf3, 0x4f, 0xc2, 0xae, 0x0b, 0x96, 0x00, 0xc2, 0x02, 0x81, 0xa3, 0x32, 0x60,
	0xf8, 0xad, 0xc0, 0x5c, 0xd0, 0x5b, 0xd0, 0x08, 0xd3, 0xcb, 0x20, 0x4a, 0x4c, 0x65, 0x5f, 0x39,
	0x68, 0xb1, 0x2a, 0xa2, 0x27, 0x60, 0x84, 0x98, 0x4f, 0x79, 0x94, 0x89, 0x94, 0xe7, 0xa6, 0xba,
	0xaf, 0x1d, 0x18, 0xce, 0x7d, 0xfb, 0x1f, 0x51, 0xbb, 0x56, 0xec, 0xd7, 0x70, 0xb6, 0x4a, 0xb5,
	0xbe, 0x2b, 0x70, 0xe3, 0x3f, 0x20, 0xfa, 0x0a, 0x9a, 0x98, 0x08, 0x1e, 0x61, 0x6e, 0x2a, 0x52,
	0xfd, 0xc9, 0x66, 0xea, 0xf6, 0x20, 0x11, 0x7c, 0xce, 0x96, 0xe4, 0xde, 0x53, 0xd8, 0x96, 0x19,
	0x4a, 0x40, 0xfb, 0x8a, 0xf3, 0xea, 0x1d, 0xe5, 0x91, 0xde, 0x84, 0xed, 0x59, 0x10, 0x17, 0x68,
	0xaa, 0x32, 0xb7, 0x08, 0xac, 0x9f, 0x0a, 0xb4, 0x6a, 0x5d, 0xfa, 0x08, 0xba, 0x7c, 0x31, 0x8b,
	0xdc, 0xcf, 0x90, 0xfb, 0x45, 0x12, 0x09, 0xa9, 0xb1, 0xc7, 0x3a, 0xcb, 0xc2, 0x6b, 0xe4, 0x93,
	0x24, 0x12, 0xf4, 0x10, 0x74, 0x59, 0x2e, 0xe5, 0xda, 0xce, 0x9d, 0x75, 0xfd, 0xda, 0x25, 0x81,
	0x49, 0xb8, 0xf5, 0x02, 0x74, 0x49, 0x37, 0xa0, 0x39, 0x71, 0x4f, 0x5d, 0xef, 0x9d, 0x4b, 0xb6,
	0x28, 0x40, 0x63, 0x3c, 0x38, 0xf6, 0xdc, 0x3e, 0x51, 0xca, 0xf3, 0xd9, 0xd0, 0x9d, 0x9c, 0x0f,
	0x88, 0x4a, 0x77, 0x40, 0x3f, 0xf1, 0x26, 0x8c, 0x68, 0xb4, 0x09, 0x5a, 0xff, 0xe8, 0x3d, 0xd1,
	0xad, 0x1f, 0x1a, 0x74, 0x57, 0x8c, 0xcb, 0xb3, 0x34, 0xc9, 0x91, 0x8e, 0x60, 0x37, 0x9d, 0x21,
	0x0f, 0xe2, 0xd8, 0x9f, 0xa6, 0x21, 0xca, 0x9e, 0xdb, 0xce, 0xc3, 0x75, 0x4d, 0x2d, 0xb9, 0xf6,
	0x71, 0x1a, 0x22, 0x33, 0x2a, 0x7a, 0x19, 0xd0, 0x37, 0xb0, 0x93, 0x8b, 0x40, 0x14, 0x39, 0x2e,
	0xcd, 0x3e, 0xdc, 0x48, 0xe9, 0xca, 0x97, 0xb1, 0xa4, 0xb3, 0x5a, 0xa6, 0xf7, 0x4b, 0x01, 0xf2,
	0x77, 0x99, 0x3e, 0x07, 0xfd, 0x7a, 0xdd, 0x4a, 0x1a, 0x3d, 0x82, 0xbd, 0x69, 0xc1, 0x39, 0x26,
	0xc2, 0x97, 0x68, 0x69, 0x85, 0xe1, 0xdc, 0x5e, 0xab, 0xb3, 0x5b, 0x51, 0x16, 0x86, 0x3f, 0x80,
	0x8e, 0x04, 0xf8, 0x1c, 0xcb, 0x4d, 0x8f, 0x92, 0xcf, 0xa6, 0x26, 0xed, 0x6e, 0xc7, 0x8b, 0x5b,
	0xab, 0xac, 0xf5, 0x18, 0x74, 0x39, 0x9a, 0x3f, 0x6c, 0x6b, 0x80, 0xea, 0x9d, 0x12, 0x85, 0xb6,
	0x01, 0xbc, 0xb7, 0x03, 0xe6, 0x8f, 0x86, 0x67, 0xc3, 0x73, 0xa2, 0x3a, 0x7c, 0xe5, 0x6f, 0x8d,
	0x91, 0xcf, 0xa2, 0x29, 0xd2, 0x8f, 0xd0, 0x19, 0x7f, 0x49, 0x8b, 0x38, 0xbc, 0xda, 0xb6, 0xbb,
	0xeb, 0x1f, 0x2c, 0xd7, 0xad, 0x77, 0x6f, 0x93, 0xa9, 0x58, 0x5b, 0x2f, 0x8d, 0x0f, 0xad, 0x1a,
	0x70, 0xd1, 0x90, 0xdf, 0xfe, 0xd9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x6c, 0xc4, 0xe3,
	0x09, 0x04, 0x00, 0x00,
}
