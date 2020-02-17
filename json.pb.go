// Code generated by protoc-gen-go. DO NOT EDIT.
// source: json.proto

package proto_json // import "github.com/synerex/proto_json"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type JsonRecord struct {
	Tag                  string               `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Json                 string               `protobuf:"bytes,3,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *JsonRecord) Reset()         { *m = JsonRecord{} }
func (m *JsonRecord) String() string { return proto.CompactTextString(m) }
func (*JsonRecord) ProtoMessage()    {}
func (*JsonRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_json_9dcfc06972839d2a, []int{0}
}
func (m *JsonRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JsonRecord.Unmarshal(m, b)
}
func (m *JsonRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JsonRecord.Marshal(b, m, deterministic)
}
func (dst *JsonRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JsonRecord.Merge(dst, src)
}
func (m *JsonRecord) XXX_Size() int {
	return xxx_messageInfo_JsonRecord.Size(m)
}
func (m *JsonRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_JsonRecord.DiscardUnknown(m)
}

var xxx_messageInfo_JsonRecord proto.InternalMessageInfo

func (m *JsonRecord) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *JsonRecord) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *JsonRecord) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

func init() {
	proto.RegisterType((*JsonRecord)(nil), "proto.json.JsonRecord")
}

func init() { proto.RegisterFile("json.proto", fileDescriptor_json_9dcfc06972839d2a) }

var fileDescriptor_json_9dcfc06972839d2a = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x2a, 0xce, 0xcf,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0x53, 0x7a, 0x20, 0x11, 0x29, 0xf9, 0xf4,
	0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0xb0, 0x50, 0x52, 0x69, 0x9a, 0x7e, 0x49, 0x66, 0x6e, 0x6a,
	0x71, 0x49, 0x62, 0x6e, 0x01, 0x44, 0xb1, 0x52, 0x12, 0x17, 0x97, 0x57, 0x71, 0x7e, 0x5e, 0x50,
	0x6a, 0x72, 0x7e, 0x51, 0x8a, 0x90, 0x00, 0x17, 0x73, 0x49, 0x62, 0xba, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x88, 0x29, 0xa4, 0xc7, 0xc5, 0x02, 0xd2, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1,
	0x6d, 0x24, 0xa5, 0x07, 0x31, 0x4f, 0x0f, 0x66, 0x9e, 0x5e, 0x08, 0xcc, 0xbc, 0x20, 0xb0, 0x3a,
	0x21, 0x21, 0x2e, 0x16, 0x90, 0xc5, 0x12, 0xcc, 0x60, 0x23, 0xc0, 0x6c, 0x27, 0xf9, 0x28, 0xd9,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe2, 0xca, 0xbc, 0xd4, 0xa2,
	0xd4, 0x0a, 0x88, 0x93, 0xe2, 0x41, 0x0a, 0x92, 0xd8, 0xc0, 0x6c, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x55, 0x78, 0x85, 0x1f, 0xc6, 0x00, 0x00, 0x00,
}