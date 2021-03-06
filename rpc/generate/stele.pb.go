// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: stele.proto

package generate

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type Prefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix []byte `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *Prefix) Reset() {
	*x = Prefix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stele_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prefix) ProtoMessage() {}

func (x *Prefix) ProtoReflect() protoreflect.Message {
	mi := &file_stele_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prefix.ProtoReflect.Descriptor instead.
func (*Prefix) Descriptor() ([]byte, []int) {
	return file_stele_proto_rawDescGZIP(), []int{0}
}

func (x *Prefix) GetPrefix() []byte {
	if x != nil {
		return x.Prefix
	}
	return nil
}

type SteleKS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys [][]byte `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *SteleKS) Reset() {
	*x = SteleKS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stele_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SteleKS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SteleKS) ProtoMessage() {}

func (x *SteleKS) ProtoReflect() protoreflect.Message {
	mi := &file_stele_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SteleKS.ProtoReflect.Descriptor instead.
func (*SteleKS) Descriptor() ([]byte, []int) {
	return file_stele_proto_rawDescGZIP(), []int{1}
}

func (x *SteleKS) GetKeys() [][]byte {
	if x != nil {
		return x.Keys
	}
	return nil
}

type SteleK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *SteleK) Reset() {
	*x = SteleK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stele_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SteleK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SteleK) ProtoMessage() {}

func (x *SteleK) ProtoReflect() protoreflect.Message {
	mi := &file_stele_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SteleK.ProtoReflect.Descriptor instead.
func (*SteleK) Descriptor() ([]byte, []int) {
	return file_stele_proto_rawDescGZIP(), []int{2}
}

func (x *SteleK) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type SteleVal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []byte `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *SteleVal) Reset() {
	*x = SteleVal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stele_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SteleVal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SteleVal) ProtoMessage() {}

func (x *SteleVal) ProtoReflect() protoreflect.Message {
	mi := &file_stele_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SteleVal.ProtoReflect.Descriptor instead.
func (*SteleVal) Descriptor() ([]byte, []int) {
	return file_stele_proto_rawDescGZIP(), []int{3}
}

func (x *SteleVal) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

type SteleKV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val []byte `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
	Ttl int64  `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *SteleKV) Reset() {
	*x = SteleKV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stele_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SteleKV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SteleKV) ProtoMessage() {}

func (x *SteleKV) ProtoReflect() protoreflect.Message {
	mi := &file_stele_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SteleKV.ProtoReflect.Descriptor instead.
func (*SteleKV) Descriptor() ([]byte, []int) {
	return file_stele_proto_rawDescGZIP(), []int{4}
}

func (x *SteleKV) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *SteleKV) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *SteleKV) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type BatchSetKVs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kvs []*SteleKV `protobuf:"bytes,1,rep,name=kvs,proto3" json:"kvs,omitempty"`
}

func (x *BatchSetKVs) Reset() {
	*x = BatchSetKVs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stele_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchSetKVs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchSetKVs) ProtoMessage() {}

func (x *BatchSetKVs) ProtoReflect() protoreflect.Message {
	mi := &file_stele_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchSetKVs.ProtoReflect.Descriptor instead.
func (*BatchSetKVs) Descriptor() ([]byte, []int) {
	return file_stele_proto_rawDescGZIP(), []int{5}
}

func (x *BatchSetKVs) GetKvs() []*SteleKV {
	if x != nil {
		return x.Kvs
	}
	return nil
}

type SteleStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SteleStatus) Reset() {
	*x = SteleStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stele_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SteleStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SteleStatus) ProtoMessage() {}

func (x *SteleStatus) ProtoReflect() protoreflect.Message {
	mi := &file_stele_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SteleStatus.ProtoReflect.Descriptor instead.
func (*SteleStatus) Descriptor() ([]byte, []int) {
	return file_stele_proto_rawDescGZIP(), []int{6}
}

type SteleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SteleRequest) Reset() {
	*x = SteleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stele_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SteleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SteleRequest) ProtoMessage() {}

func (x *SteleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stele_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SteleRequest.ProtoReflect.Descriptor instead.
func (*SteleRequest) Descriptor() ([]byte, []int) {
	return file_stele_proto_rawDescGZIP(), []int{7}
}

var File_stele_proto protoreflect.FileDescriptor

var file_stele_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x65, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x1d, 0x0a, 0x07, 0x53, 0x74, 0x65,
	0x6c, 0x65, 0x4b, 0x53, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x65, 0x6c,
	0x65, 0x4b, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x1c, 0x0a, 0x08, 0x53, 0x74, 0x65, 0x6c, 0x65, 0x56, 0x61, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x76,
	0x61, 0x6c, 0x22, 0x3f, 0x0a, 0x07, 0x53, 0x74, 0x65, 0x6c, 0x65, 0x4b, 0x56, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x74, 0x74, 0x6c, 0x22, 0x32, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x74, 0x4b,
	0x56, 0x73, 0x12, 0x23, 0x0a, 0x03, 0x6b, 0x76, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x65, 0x6c, 0x65,
	0x4b, 0x56, 0x52, 0x03, 0x6b, 0x76, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x74, 0x65, 0x6c, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x74, 0x65, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xc0, 0x03, 0x0a, 0x05, 0x53, 0x74, 0x65, 0x6c, 0x65,
	0x12, 0x2f, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x2e, 0x53, 0x74, 0x65, 0x6c, 0x65, 0x4b, 0x56, 0x1a, 0x15, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x65, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x65, 0x6c, 0x65, 0x4b, 0x1a, 0x12, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x65, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x31,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x65, 0x6c, 0x65, 0x4b, 0x1a, 0x15, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x65, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x38, 0x0a, 0x08, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x74, 0x12, 0x15, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65,
	0x74, 0x4b, 0x56, 0x73, 0x1a, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e,
	0x53, 0x74, 0x65, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x08, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x2e, 0x53, 0x74, 0x65, 0x6c, 0x65, 0x4b, 0x53, 0x1a, 0x15, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x74, 0x4b, 0x56,
	0x73, 0x12, 0x38, 0x0a, 0x0b, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x65, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x65, 0x6c, 0x65, 0x4b, 0x53, 0x12, 0x45, 0x0a, 0x14, 0x49,
	0x74, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x41, 0x6e, 0x64, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x53,
	0x74, 0x65, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x74, 0x4b,
	0x56, 0x73, 0x12, 0x35, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x53, 0x63, 0x61, 0x6e,
	0x12, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x1a, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x65, 0x74, 0x4b, 0x56, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stele_proto_rawDescOnce sync.Once
	file_stele_proto_rawDescData = file_stele_proto_rawDesc
)

func file_stele_proto_rawDescGZIP() []byte {
	file_stele_proto_rawDescOnce.Do(func() {
		file_stele_proto_rawDescData = protoimpl.X.CompressGZIP(file_stele_proto_rawDescData)
	})
	return file_stele_proto_rawDescData
}

var file_stele_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_stele_proto_goTypes = []interface{}{
	(*Prefix)(nil),       // 0: generate.Prefix
	(*SteleKS)(nil),      // 1: generate.SteleKS
	(*SteleK)(nil),       // 2: generate.SteleK
	(*SteleVal)(nil),     // 3: generate.SteleVal
	(*SteleKV)(nil),      // 4: generate.SteleKV
	(*BatchSetKVs)(nil),  // 5: generate.BatchSetKVs
	(*SteleStatus)(nil),  // 6: generate.SteleStatus
	(*SteleRequest)(nil), // 7: generate.SteleRequest
}
var file_stele_proto_depIdxs = []int32{
	4, // 0: generate.BatchSetKVs.kvs:type_name -> generate.SteleKV
	4, // 1: generate.Stele.Set:input_type -> generate.SteleKV
	2, // 2: generate.Stele.Get:input_type -> generate.SteleK
	2, // 3: generate.Stele.Delete:input_type -> generate.SteleK
	5, // 4: generate.Stele.BatchSet:input_type -> generate.BatchSetKVs
	1, // 5: generate.Stele.BatchGet:input_type -> generate.SteleKS
	7, // 6: generate.Stele.IterateKeys:input_type -> generate.SteleRequest
	7, // 7: generate.Stele.IterateKeysAndValues:input_type -> generate.SteleRequest
	0, // 8: generate.Stele.PrefixScan:input_type -> generate.Prefix
	6, // 9: generate.Stele.Set:output_type -> generate.SteleStatus
	3, // 10: generate.Stele.Get:output_type -> generate.SteleVal
	6, // 11: generate.Stele.Delete:output_type -> generate.SteleStatus
	6, // 12: generate.Stele.BatchSet:output_type -> generate.SteleStatus
	5, // 13: generate.Stele.BatchGet:output_type -> generate.BatchSetKVs
	1, // 14: generate.Stele.IterateKeys:output_type -> generate.SteleKS
	5, // 15: generate.Stele.IterateKeysAndValues:output_type -> generate.BatchSetKVs
	5, // 16: generate.Stele.PrefixScan:output_type -> generate.BatchSetKVs
	9, // [9:17] is the sub-list for method output_type
	1, // [1:9] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_stele_proto_init() }
func file_stele_proto_init() {
	if File_stele_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stele_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prefix); i {
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
		file_stele_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SteleKS); i {
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
		file_stele_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SteleK); i {
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
		file_stele_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SteleVal); i {
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
		file_stele_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SteleKV); i {
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
		file_stele_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchSetKVs); i {
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
		file_stele_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SteleStatus); i {
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
		file_stele_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SteleRequest); i {
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
			RawDescriptor: file_stele_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stele_proto_goTypes,
		DependencyIndexes: file_stele_proto_depIdxs,
		MessageInfos:      file_stele_proto_msgTypes,
	}.Build()
	File_stele_proto = out.File
	file_stele_proto_rawDesc = nil
	file_stele_proto_goTypes = nil
	file_stele_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SteleClient is the client API for Stele service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SteleClient interface {
	Set(ctx context.Context, in *SteleKV, opts ...grpc.CallOption) (*SteleStatus, error)
	Get(ctx context.Context, in *SteleK, opts ...grpc.CallOption) (*SteleVal, error)
	Delete(ctx context.Context, in *SteleK, opts ...grpc.CallOption) (*SteleStatus, error)
	// Batch Insertion Failure Rollback. 批量插入 失败 回滚
	BatchSet(ctx context.Context, in *BatchSetKVs, opts ...grpc.CallOption) (*SteleStatus, error)
	// Batch Get. 批量查询
	BatchGet(ctx context.Context, in *SteleKS, opts ...grpc.CallOption) (*BatchSetKVs, error)
	// Iterate over all keys. 遍历所有Key
	IterateKeys(ctx context.Context, in *SteleRequest, opts ...grpc.CallOption) (*SteleKS, error)
	// Iterate over keys and values. 遍历Key和value
	IterateKeysAndValues(ctx context.Context, in *SteleRequest, opts ...grpc.CallOption) (*BatchSetKVs, error)
	// Prefix Scan. 前缀扫描
	PrefixScan(ctx context.Context, in *Prefix, opts ...grpc.CallOption) (*BatchSetKVs, error)
}

type steleClient struct {
	cc grpc.ClientConnInterface
}

func NewSteleClient(cc grpc.ClientConnInterface) SteleClient {
	return &steleClient{cc}
}

func (c *steleClient) Set(ctx context.Context, in *SteleKV, opts ...grpc.CallOption) (*SteleStatus, error) {
	out := new(SteleStatus)
	err := c.cc.Invoke(ctx, "/generate.Stele/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *steleClient) Get(ctx context.Context, in *SteleK, opts ...grpc.CallOption) (*SteleVal, error) {
	out := new(SteleVal)
	err := c.cc.Invoke(ctx, "/generate.Stele/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *steleClient) Delete(ctx context.Context, in *SteleK, opts ...grpc.CallOption) (*SteleStatus, error) {
	out := new(SteleStatus)
	err := c.cc.Invoke(ctx, "/generate.Stele/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *steleClient) BatchSet(ctx context.Context, in *BatchSetKVs, opts ...grpc.CallOption) (*SteleStatus, error) {
	out := new(SteleStatus)
	err := c.cc.Invoke(ctx, "/generate.Stele/BatchSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *steleClient) BatchGet(ctx context.Context, in *SteleKS, opts ...grpc.CallOption) (*BatchSetKVs, error) {
	out := new(BatchSetKVs)
	err := c.cc.Invoke(ctx, "/generate.Stele/BatchGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *steleClient) IterateKeys(ctx context.Context, in *SteleRequest, opts ...grpc.CallOption) (*SteleKS, error) {
	out := new(SteleKS)
	err := c.cc.Invoke(ctx, "/generate.Stele/IterateKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *steleClient) IterateKeysAndValues(ctx context.Context, in *SteleRequest, opts ...grpc.CallOption) (*BatchSetKVs, error) {
	out := new(BatchSetKVs)
	err := c.cc.Invoke(ctx, "/generate.Stele/IterateKeysAndValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *steleClient) PrefixScan(ctx context.Context, in *Prefix, opts ...grpc.CallOption) (*BatchSetKVs, error) {
	out := new(BatchSetKVs)
	err := c.cc.Invoke(ctx, "/generate.Stele/PrefixScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SteleServer is the server API for Stele service.
type SteleServer interface {
	Set(context.Context, *SteleKV) (*SteleStatus, error)
	Get(context.Context, *SteleK) (*SteleVal, error)
	Delete(context.Context, *SteleK) (*SteleStatus, error)
	// Batch Insertion Failure Rollback. 批量插入 失败 回滚
	BatchSet(context.Context, *BatchSetKVs) (*SteleStatus, error)
	// Batch Get. 批量查询
	BatchGet(context.Context, *SteleKS) (*BatchSetKVs, error)
	// Iterate over all keys. 遍历所有Key
	IterateKeys(context.Context, *SteleRequest) (*SteleKS, error)
	// Iterate over keys and values. 遍历Key和value
	IterateKeysAndValues(context.Context, *SteleRequest) (*BatchSetKVs, error)
	// Prefix Scan. 前缀扫描
	PrefixScan(context.Context, *Prefix) (*BatchSetKVs, error)
}

// UnimplementedSteleServer can be embedded to have forward compatible implementations.
type UnimplementedSteleServer struct {
}

func (*UnimplementedSteleServer) Set(context.Context, *SteleKV) (*SteleStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedSteleServer) Get(context.Context, *SteleK) (*SteleVal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedSteleServer) Delete(context.Context, *SteleK) (*SteleStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedSteleServer) BatchSet(context.Context, *BatchSetKVs) (*SteleStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchSet not implemented")
}
func (*UnimplementedSteleServer) BatchGet(context.Context, *SteleKS) (*BatchSetKVs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGet not implemented")
}
func (*UnimplementedSteleServer) IterateKeys(context.Context, *SteleRequest) (*SteleKS, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IterateKeys not implemented")
}
func (*UnimplementedSteleServer) IterateKeysAndValues(context.Context, *SteleRequest) (*BatchSetKVs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IterateKeysAndValues not implemented")
}
func (*UnimplementedSteleServer) PrefixScan(context.Context, *Prefix) (*BatchSetKVs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrefixScan not implemented")
}

func RegisterSteleServer(s *grpc.Server, srv SteleServer) {
	s.RegisterService(&_Stele_serviceDesc, srv)
}

func _Stele_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SteleKV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SteleServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generate.Stele/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SteleServer).Set(ctx, req.(*SteleKV))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stele_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SteleK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SteleServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generate.Stele/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SteleServer).Get(ctx, req.(*SteleK))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stele_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SteleK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SteleServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generate.Stele/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SteleServer).Delete(ctx, req.(*SteleK))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stele_BatchSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchSetKVs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SteleServer).BatchSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generate.Stele/BatchSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SteleServer).BatchSet(ctx, req.(*BatchSetKVs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stele_BatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SteleKS)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SteleServer).BatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generate.Stele/BatchGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SteleServer).BatchGet(ctx, req.(*SteleKS))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stele_IterateKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SteleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SteleServer).IterateKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generate.Stele/IterateKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SteleServer).IterateKeys(ctx, req.(*SteleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stele_IterateKeysAndValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SteleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SteleServer).IterateKeysAndValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generate.Stele/IterateKeysAndValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SteleServer).IterateKeysAndValues(ctx, req.(*SteleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stele_PrefixScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Prefix)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SteleServer).PrefixScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generate.Stele/PrefixScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SteleServer).PrefixScan(ctx, req.(*Prefix))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stele_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generate.Stele",
	HandlerType: (*SteleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Stele_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Stele_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Stele_Delete_Handler,
		},
		{
			MethodName: "BatchSet",
			Handler:    _Stele_BatchSet_Handler,
		},
		{
			MethodName: "BatchGet",
			Handler:    _Stele_BatchGet_Handler,
		},
		{
			MethodName: "IterateKeys",
			Handler:    _Stele_IterateKeys_Handler,
		},
		{
			MethodName: "IterateKeysAndValues",
			Handler:    _Stele_IterateKeysAndValues_Handler,
		},
		{
			MethodName: "PrefixScan",
			Handler:    _Stele_PrefixScan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stele.proto",
}
