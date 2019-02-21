// Code generated by protoc-gen-go. DO NOT EDIT.
// source: check_task.proto

package geo_proto_check_task

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Запрос на CheckTask
type CheckTaskRpcRequest struct {
	// Название
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckTaskRpcRequest) Reset()         { *m = CheckTaskRpcRequest{} }
func (m *CheckTaskRpcRequest) String() string { return proto.CompactTextString(m) }
func (*CheckTaskRpcRequest) ProtoMessage()    {}
func (*CheckTaskRpcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_check_task_083df35a8756236f, []int{0}
}
func (m *CheckTaskRpcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckTaskRpcRequest.Unmarshal(m, b)
}
func (m *CheckTaskRpcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckTaskRpcRequest.Marshal(b, m, deterministic)
}
func (dst *CheckTaskRpcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckTaskRpcRequest.Merge(dst, src)
}
func (m *CheckTaskRpcRequest) XXX_Size() int {
	return xxx_messageInfo_CheckTaskRpcRequest.Size(m)
}
func (m *CheckTaskRpcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckTaskRpcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckTaskRpcRequest proto.InternalMessageInfo

func (m *CheckTaskRpcRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Ответ на CheckTask
type CheckTaskRpcResponse struct {
	// Признак пробиваемой станции
	InProgress bool `protobuf:"varint,1,opt,name=inProgress,proto3" json:"inProgress,omitempty"`
	// Описание ошибки
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// Временная метка операции
	Ts                   string   `protobuf:"bytes,3,opt,name=ts,proto3" json:"ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckTaskRpcResponse) Reset()         { *m = CheckTaskRpcResponse{} }
func (m *CheckTaskRpcResponse) String() string { return proto.CompactTextString(m) }
func (*CheckTaskRpcResponse) ProtoMessage()    {}
func (*CheckTaskRpcResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_check_task_083df35a8756236f, []int{1}
}
func (m *CheckTaskRpcResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckTaskRpcResponse.Unmarshal(m, b)
}
func (m *CheckTaskRpcResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckTaskRpcResponse.Marshal(b, m, deterministic)
}
func (dst *CheckTaskRpcResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckTaskRpcResponse.Merge(dst, src)
}
func (m *CheckTaskRpcResponse) XXX_Size() int {
	return xxx_messageInfo_CheckTaskRpcResponse.Size(m)
}
func (m *CheckTaskRpcResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckTaskRpcResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckTaskRpcResponse proto.InternalMessageInfo

func (m *CheckTaskRpcResponse) GetInProgress() bool {
	if m != nil {
		return m.InProgress
	}
	return false
}

func (m *CheckTaskRpcResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CheckTaskRpcResponse) GetTs() string {
	if m != nil {
		return m.Ts
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckTaskRpcRequest)(nil), "geo.proto.check_task.CheckTaskRpcRequest")
	proto.RegisterType((*CheckTaskRpcResponse)(nil), "geo.proto.check_task.CheckTaskRpcResponse")
}

func init() { proto.RegisterFile("check_task.proto", fileDescriptor_check_task_083df35a8756236f) }

var fileDescriptor_check_task_083df35a8756236f = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xce, 0x48, 0x4d,
	0xce, 0x8e, 0x2f, 0x49, 0x2c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x49, 0x4f,
	0xcd, 0x87, 0x30, 0xf5, 0x10, 0x72, 0x4a, 0x9a, 0x5c, 0xc2, 0xce, 0x20, 0x5e, 0x48, 0x62, 0x71,
	0x76, 0x50, 0x41, 0x72, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x5e,
	0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xad, 0x14, 0xc3, 0x25, 0x82,
	0xaa, 0xb4, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8e, 0x8b, 0x2b, 0x33, 0x2f, 0xa0, 0x28,
	0x3f, 0xbd, 0x28, 0xb5, 0xb8, 0x18, 0xac, 0x83, 0x23, 0x08, 0x49, 0x44, 0x48, 0x84, 0x8b, 0x35,
	0xb5, 0xa8, 0x28, 0xbf, 0x48, 0x82, 0x09, 0x6c, 0x18, 0x84, 0x23, 0xc4, 0xc7, 0xc5, 0x54, 0x52,
	0x2c, 0xc1, 0x0c, 0x16, 0x62, 0x2a, 0x29, 0x36, 0x2a, 0xe4, 0xe2, 0x84, 0x9b, 0x2e, 0x94, 0x82,
	0xcc, 0xd1, 0xd4, 0xc3, 0xe6, 0x72, 0x3d, 0x2c, 0xce, 0x96, 0xd2, 0x22, 0x46, 0x29, 0xc4, 0xd9,
	0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x85, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x24, 0xe9,
	0xc6, 0x2c, 0x01, 0x00, 0x00,
}