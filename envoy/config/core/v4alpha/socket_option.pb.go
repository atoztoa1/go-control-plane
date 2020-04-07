// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v4alpha/socket_option.proto

package envoy_config_core_v4alpha

import (
	fmt "fmt"
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

type SocketOption_SocketState int32

const (
	SocketOption_STATE_PREBIND   SocketOption_SocketState = 0
	SocketOption_STATE_BOUND     SocketOption_SocketState = 1
	SocketOption_STATE_LISTENING SocketOption_SocketState = 2
)

var SocketOption_SocketState_name = map[int32]string{
	0: "STATE_PREBIND",
	1: "STATE_BOUND",
	2: "STATE_LISTENING",
}

var SocketOption_SocketState_value = map[string]int32{
	"STATE_PREBIND":   0,
	"STATE_BOUND":     1,
	"STATE_LISTENING": 2,
}

func (x SocketOption_SocketState) String() string {
	return proto.EnumName(SocketOption_SocketState_name, int32(x))
}

func (SocketOption_SocketState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1e5251252359b22c, []int{0, 0}
}

type SocketOption struct {
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Level       int64  `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Name        int64  `protobuf:"varint,3,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*SocketOption_IntValue
	//	*SocketOption_BufValue
	Value                isSocketOption_Value     `protobuf_oneof:"value"`
	State                SocketOption_SocketState `protobuf:"varint,6,opt,name=state,proto3,enum=envoy.config.core.v4alpha.SocketOption_SocketState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SocketOption) Reset()         { *m = SocketOption{} }
func (m *SocketOption) String() string { return proto.CompactTextString(m) }
func (*SocketOption) ProtoMessage()    {}
func (*SocketOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e5251252359b22c, []int{0}
}

func (m *SocketOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketOption.Unmarshal(m, b)
}
func (m *SocketOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketOption.Marshal(b, m, deterministic)
}
func (m *SocketOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketOption.Merge(m, src)
}
func (m *SocketOption) XXX_Size() int {
	return xxx_messageInfo_SocketOption.Size(m)
}
func (m *SocketOption) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketOption.DiscardUnknown(m)
}

var xxx_messageInfo_SocketOption proto.InternalMessageInfo

func (m *SocketOption) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SocketOption) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *SocketOption) GetName() int64 {
	if m != nil {
		return m.Name
	}
	return 0
}

type isSocketOption_Value interface {
	isSocketOption_Value()
}

type SocketOption_IntValue struct {
	IntValue int64 `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3,oneof"`
}

type SocketOption_BufValue struct {
	BufValue []byte `protobuf:"bytes,5,opt,name=buf_value,json=bufValue,proto3,oneof"`
}

func (*SocketOption_IntValue) isSocketOption_Value() {}

func (*SocketOption_BufValue) isSocketOption_Value() {}

func (m *SocketOption) GetValue() isSocketOption_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SocketOption) GetIntValue() int64 {
	if x, ok := m.GetValue().(*SocketOption_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *SocketOption) GetBufValue() []byte {
	if x, ok := m.GetValue().(*SocketOption_BufValue); ok {
		return x.BufValue
	}
	return nil
}

func (m *SocketOption) GetState() SocketOption_SocketState {
	if m != nil {
		return m.State
	}
	return SocketOption_STATE_PREBIND
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SocketOption) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SocketOption_IntValue)(nil),
		(*SocketOption_BufValue)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.config.core.v4alpha.SocketOption_SocketState", SocketOption_SocketState_name, SocketOption_SocketState_value)
	proto.RegisterType((*SocketOption)(nil), "envoy.config.core.v4alpha.SocketOption")
}

func init() {
	proto.RegisterFile("envoy/config/core/v4alpha/socket_option.proto", fileDescriptor_1e5251252359b22c)
}

var fileDescriptor_1e5251252359b22c = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x31, 0x8f, 0xd3, 0x30,
	0x1c, 0xc5, 0xcf, 0xc9, 0xa5, 0xba, 0x73, 0x0b, 0x97, 0x33, 0x48, 0x84, 0x93, 0x0e, 0xa5, 0x65,
	0x20, 0x0b, 0x89, 0x44, 0x99, 0xd8, 0xb0, 0xda, 0x42, 0x25, 0x94, 0x56, 0x49, 0x61, 0xad, 0xdc,
	0xd4, 0x2d, 0x16, 0xc1, 0x8e, 0x12, 0x27, 0xa2, 0x1b, 0x62, 0xe2, 0x33, 0xf0, 0x51, 0xd8, 0x91,
	0x58, 0xf9, 0x20, 0xec, 0x88, 0x09, 0xd9, 0x0e, 0x52, 0x10, 0xd7, 0xcd, 0xff, 0xf7, 0x7b, 0x7f,
	0xeb, 0x3d, 0x1b, 0x3e, 0xa6, 0xbc, 0x11, 0x87, 0x28, 0x13, 0x7c, 0xc7, 0xf6, 0x51, 0x26, 0x4a,
	0x1a, 0x35, 0x4f, 0x49, 0x5e, 0xbc, 0x25, 0x51, 0x25, 0xb2, 0x77, 0x54, 0xae, 0x45, 0x21, 0x99,
	0xe0, 0x61, 0x51, 0x0a, 0x29, 0xd0, 0x7d, 0x6d, 0x0f, 0x8d, 0x3d, 0x54, 0xf6, 0xb0, 0xb5, 0x5f,
	0x5d, 0xd7, 0xdb, 0x82, 0x44, 0x84, 0x73, 0x21, 0x89, 0xda, 0xa8, 0xa2, 0x4a, 0x12, 0x59, 0x57,
	0x66, 0xf3, 0x6a, 0xf8, 0x1f, 0x6e, 0x68, 0x59, 0x31, 0xc1, 0x19, 0xdf, 0xb7, 0x96, 0x7b, 0x0d,
	0xc9, 0xd9, 0x96, 0x48, 0x1a, 0xfd, 0x3d, 0x18, 0x30, 0xfa, 0x69, 0xc1, 0x41, 0xaa, 0xd3, 0x2c,
	0x74, 0x18, 0xe4, 0xc3, 0xfe, 0x96, 0x56, 0x59, 0xc9, 0xf4, 0xe8, 0x01, 0x1f, 0x04, 0xe7, 0x49,
	0x57, 0x42, 0x77, 0xa1, 0x93, 0xd3, 0x86, 0xe6, 0x9e, 0xe5, 0x83, 0xc0, 0x4e, 0xcc, 0x80, 0x10,
	0x3c, 0xe5, 0xe4, 0x3d, 0xf5, 0x6c, 0x2d, 0xea, 0x33, 0xba, 0x86, 0xe7, 0x8c, 0xcb, 0x75, 0x43,
	0xf2, 0x9a, 0x7a, 0xa7, 0x0a, 0xbc, 0x3c, 0x49, 0xce, 0x18, 0x97, 0x6f, 0x94, 0xa2, 0xf0, 0xa6,
	0xde, 0xb5, 0xd8, 0xf1, 0x41, 0x30, 0x50, 0x78, 0x53, 0xef, 0x0c, 0x4e, 0xa1, 0xa3, 0x6a, 0x52,
	0xaf, 0xe7, 0x83, 0xe0, 0xf6, 0x93, 0x71, 0x78, 0xf4, 0x81, 0xc2, 0x6e, 0x83, 0x76, 0x48, 0xd5,
	0x2a, 0x3e, 0xfb, 0x8d, 0x9d, 0x4f, 0xc0, 0x72, 0x41, 0x62, 0xee, 0x1a, 0xcd, 0x60, 0xbf, 0xc3,
	0xd1, 0x25, 0xbc, 0x95, 0xae, 0x9e, 0xaf, 0xa6, 0xeb, 0x65, 0x32, 0xc5, 0xf3, 0x78, 0xe2, 0x9e,
	0xa0, 0x0b, 0xd8, 0x37, 0x12, 0x5e, 0xbc, 0x8e, 0x27, 0x2e, 0x40, 0x77, 0xe0, 0x85, 0x11, 0x5e,
	0xcd, 0xd3, 0xd5, 0x34, 0x9e, 0xc7, 0x2f, 0x5c, 0xeb, 0x59, 0xf0, 0xe5, 0xdb, 0xe7, 0x07, 0x0f,
	0xe1, 0xf0, 0x86, 0x4c, 0xe3, 0x7f, 0xe2, 0xe0, 0x01, 0x74, 0x74, 0x43, 0x64, 0xff, 0xc2, 0x00,
	0xcf, 0xbe, 0x7e, 0xfc, 0xfe, 0xa3, 0x67, 0xb9, 0x36, 0x7c, 0xc4, 0x84, 0x69, 0x54, 0x94, 0xe2,
	0xc3, 0xe1, 0x78, 0x39, 0x7c, 0xd9, 0xbd, 0x6e, 0xa9, 0x7e, 0x6d, 0x09, 0x36, 0x3d, 0xfd, 0x7d,
	0xe3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x2c, 0xd4, 0x86, 0x65, 0x02, 0x00, 0x00,
}