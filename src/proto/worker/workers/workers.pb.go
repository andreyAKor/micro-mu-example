// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workers.proto

package geo_proto_workers

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

// Событие на workers
type WorkersEvent struct {
	// id воркера в nats-сети
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Общее количество воркеров у воркера
	TotalWorkers uint32 `protobuf:"varint,2,opt,name=totalWorkers,proto3" json:"totalWorkers,omitempty"`
	// Количество свободных воркеров у воркера
	FreeWorkers uint32 `protobuf:"varint,3,opt,name=freeWorkers,proto3" json:"freeWorkers,omitempty"`
	// Количество занятых воркеров у воркера
	CountWorkers uint32 `protobuf:"varint,4,opt,name=countWorkers,proto3" json:"countWorkers,omitempty"`
	// Список данных занятых воркеров у воркера
	Workers []*WorkerData `protobuf:"bytes,5,rep,name=workers,proto3" json:"workers,omitempty"`
	// Временная метка операции
	Ts                   string   `protobuf:"bytes,6,opt,name=ts,proto3" json:"ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkersEvent) Reset()         { *m = WorkersEvent{} }
func (m *WorkersEvent) String() string { return proto.CompactTextString(m) }
func (*WorkersEvent) ProtoMessage()    {}
func (*WorkersEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_workers_80ab8d4b6b6fc76a, []int{0}
}
func (m *WorkersEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkersEvent.Unmarshal(m, b)
}
func (m *WorkersEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkersEvent.Marshal(b, m, deterministic)
}
func (dst *WorkersEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkersEvent.Merge(dst, src)
}
func (m *WorkersEvent) XXX_Size() int {
	return xxx_messageInfo_WorkersEvent.Size(m)
}
func (m *WorkersEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkersEvent.DiscardUnknown(m)
}

var xxx_messageInfo_WorkersEvent proto.InternalMessageInfo

func (m *WorkersEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WorkersEvent) GetTotalWorkers() uint32 {
	if m != nil {
		return m.TotalWorkers
	}
	return 0
}

func (m *WorkersEvent) GetFreeWorkers() uint32 {
	if m != nil {
		return m.FreeWorkers
	}
	return 0
}

func (m *WorkersEvent) GetCountWorkers() uint32 {
	if m != nil {
		return m.CountWorkers
	}
	return 0
}

func (m *WorkersEvent) GetWorkers() []*WorkerData {
	if m != nil {
		return m.Workers
	}
	return nil
}

func (m *WorkersEvent) GetTs() string {
	if m != nil {
		return m.Ts
	}
	return ""
}

// Данные воркера
type WorkerData struct {
	// Название
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkerData) Reset()         { *m = WorkerData{} }
func (m *WorkerData) String() string { return proto.CompactTextString(m) }
func (*WorkerData) ProtoMessage()    {}
func (*WorkerData) Descriptor() ([]byte, []int) {
	return fileDescriptor_workers_80ab8d4b6b6fc76a, []int{1}
}
func (m *WorkerData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerData.Unmarshal(m, b)
}
func (m *WorkerData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerData.Marshal(b, m, deterministic)
}
func (dst *WorkerData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerData.Merge(dst, src)
}
func (m *WorkerData) XXX_Size() int {
	return xxx_messageInfo_WorkerData.Size(m)
}
func (m *WorkerData) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerData.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerData proto.InternalMessageInfo

func (m *WorkerData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*WorkersEvent)(nil), "geo.proto.workers.WorkersEvent")
	proto.RegisterType((*WorkerData)(nil), "geo.proto.workers.WorkerData")
}

func init() { proto.RegisterFile("workers.proto", fileDescriptor_workers_80ab8d4b6b6fc76a) }

var fileDescriptor_workers_80ab8d4b6b6fc76a = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0xcf, 0x2f, 0xca,
	0x4e, 0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c, 0x4f, 0xcd, 0x87, 0x30,
	0xf5, 0xa0, 0x12, 0x4a, 0x17, 0x19, 0xb9, 0x78, 0xc2, 0x21, 0x6c, 0xd7, 0xb2, 0xd4, 0xbc, 0x12,
	0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x14,
	0x21, 0x25, 0x2e, 0x9e, 0x92, 0xfc, 0x92, 0xc4, 0x1c, 0xa8, 0x22, 0x09, 0x26, 0x05, 0x46, 0x0d,
	0xde, 0x20, 0x14, 0x31, 0x21, 0x05, 0x2e, 0xee, 0xb4, 0xa2, 0xd4, 0x54, 0x98, 0x12, 0x66, 0xb0,
	0x12, 0x64, 0x21, 0x90, 0x29, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x30, 0x25, 0x2c, 0x10, 0x53, 0x90,
	0xc5, 0x84, 0xcc, 0xb9, 0xd8, 0xa1, 0xae, 0x92, 0x60, 0x55, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd5,
	0xc3, 0x70, 0xaf, 0x1e, 0x44, 0xb1, 0x4b, 0x62, 0x49, 0x62, 0x10, 0x4c, 0x35, 0xc8, 0xc9, 0x25,
	0xc5, 0x12, 0x6c, 0x10, 0x27, 0x97, 0x14, 0x2b, 0x29, 0x70, 0x71, 0x21, 0x94, 0x09, 0x09, 0x71,
	0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x42, 0xbd, 0x04, 0x66, 0x27, 0xb1, 0x81, 0x0d, 0x35, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xa3, 0x8e, 0x74, 0x0d, 0x20, 0x01, 0x00, 0x00,
}