// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: grpc-echo.proto

package whereami

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type All struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind     string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Version  string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Region   string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Cluster  string `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Hostname string `protobuf:"bytes,5,opt,name=hostname,proto3" json:"hostname,omitempty"`
	SourceIp string `protobuf:"bytes,6,opt,name=source_ip,json=sourceIp,proto3" json:"source_ip,omitempty"`
}

func (x *All) Reset() {
	*x = All{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *All) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*All) ProtoMessage() {}

func (x *All) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use All.ProtoReflect.Descriptor instead.
func (*All) Descriptor() ([]byte, []int) {
	return file_grpc_echo_proto_rawDescGZIP(), []int{0}
}

func (x *All) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *All) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *All) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *All) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *All) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *All) GetSourceIp() string {
	if x != nil {
		return x.SourceIp
	}
	return ""
}

type Kind struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *Kind) Reset() {
	*x = Kind{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kind) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kind) ProtoMessage() {}

func (x *Kind) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kind.ProtoReflect.Descriptor instead.
func (*Kind) Descriptor() ([]byte, []int) {
	return file_grpc_echo_proto_rawDescGZIP(), []int{1}
}

func (x *Kind) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_echo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_echo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_grpc_echo_proto_rawDescGZIP(), []int{2}
}

func (x *Version) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *Region) Reset() {
	*x = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_echo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_echo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_grpc_echo_proto_rawDescGZIP(), []int{3}
}

func (x *Region) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_echo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_echo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_grpc_echo_proto_rawDescGZIP(), []int{4}
}

func (x *Cluster) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

type Hostname struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *Hostname) Reset() {
	*x = Hostname{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_echo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hostname) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hostname) ProtoMessage() {}

func (x *Hostname) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_echo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hostname.ProtoReflect.Descriptor instead.
func (*Hostname) Descriptor() ([]byte, []int) {
	return file_grpc_echo_proto_rawDescGZIP(), []int{5}
}

func (x *Hostname) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type SourceIp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceIp string `protobuf:"bytes,1,opt,name=source_ip,json=sourceIp,proto3" json:"source_ip,omitempty"`
}

func (x *SourceIp) Reset() {
	*x = SourceIp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_echo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceIp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceIp) ProtoMessage() {}

func (x *SourceIp) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_echo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceIp.ProtoReflect.Descriptor instead.
func (*SourceIp) Descriptor() ([]byte, []int) {
	return file_grpc_echo_proto_rawDescGZIP(), []int{6}
}

func (x *SourceIp) GetSourceIp() string {
	if x != nil {
		return x.SourceIp
	}
	return ""
}

type HeaderName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestHeaderName string `protobuf:"bytes,1,opt,name=request_header_name,json=requestHeaderName,proto3" json:"request_header_name,omitempty"`
}

func (x *HeaderName) Reset() {
	*x = HeaderName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_echo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderName) ProtoMessage() {}

func (x *HeaderName) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_echo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderName.ProtoReflect.Descriptor instead.
func (*HeaderName) Descriptor() ([]byte, []int) {
	return file_grpc_echo_proto_rawDescGZIP(), []int{7}
}

func (x *HeaderName) GetRequestHeaderName() string {
	if x != nil {
		return x.RequestHeaderName
	}
	return ""
}

type HeaderValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestHeaderValue string `protobuf:"bytes,1,opt,name=request_header_value,json=requestHeaderValue,proto3" json:"request_header_value,omitempty"`
}

func (x *HeaderValue) Reset() {
	*x = HeaderValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_echo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValue) ProtoMessage() {}

func (x *HeaderValue) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_echo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValue.ProtoReflect.Descriptor instead.
func (*HeaderValue) Descriptor() ([]byte, []int) {
	return file_grpc_echo_proto_rawDescGZIP(), []int{8}
}

func (x *HeaderValue) GetRequestHeaderValue() string {
	if x != nil {
		return x.RequestHeaderValue
	}
	return ""
}

var File_grpc_echo_proto protoreflect.FileDescriptor

var file_grpc_echo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e,
	0x01, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x70, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70, 0x22,
	0x1a, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x23, 0x0a, 0x07, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x20, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x22, 0x23, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x27, 0x0a, 0x08, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70, 0x22, 0x3c, 0x0a, 0x0a, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xf9, 0x02, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x04, 0x2e, 0x41, 0x6c, 0x6c, 0x12,
	0x28, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x05, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x08, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x07,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x48, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x09,
	0x2e, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x09, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70, 0x12, 0x26, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x0c, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x61, 0x7a, 0x73, 0x68, 0x69, 0x6e, 0x6f, 0x68, 0x61, 0x72, 0x61, 0x2f, 0x70,
	0x62, 0x2f, 0x77, 0x68, 0x65, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_grpc_echo_proto_rawDescOnce sync.Once
	file_grpc_echo_proto_rawDescData = file_grpc_echo_proto_rawDesc
)

func file_grpc_echo_proto_rawDescGZIP() []byte {
	file_grpc_echo_proto_rawDescOnce.Do(func() {
		file_grpc_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_echo_proto_rawDescData)
	})
	return file_grpc_echo_proto_rawDescData
}

var file_grpc_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_grpc_echo_proto_goTypes = []interface{}{
	(*All)(nil),           // 0: All
	(*Kind)(nil),          // 1: Kind
	(*Version)(nil),       // 2: Version
	(*Region)(nil),        // 3: Region
	(*Cluster)(nil),       // 4: Cluster
	(*Hostname)(nil),      // 5: Hostname
	(*SourceIp)(nil),      // 6: SourceIp
	(*HeaderName)(nil),    // 7: HeaderName
	(*HeaderValue)(nil),   // 8: HeaderValue
	(*emptypb.Empty)(nil), // 9: google.protobuf.Empty
}
var file_grpc_echo_proto_depIdxs = []int32{
	9, // 0: EchoService.GetAll:input_type -> google.protobuf.Empty
	9, // 1: EchoService.GetKind:input_type -> google.protobuf.Empty
	9, // 2: EchoService.GetVersion:input_type -> google.protobuf.Empty
	9, // 3: EchoService.GetRegion:input_type -> google.protobuf.Empty
	9, // 4: EchoService.GetCluster:input_type -> google.protobuf.Empty
	9, // 5: EchoService.GetHostname:input_type -> google.protobuf.Empty
	9, // 6: EchoService.GetSourceIp:input_type -> google.protobuf.Empty
	7, // 7: EchoService.GetHeader:input_type -> HeaderName
	0, // 8: EchoService.GetAll:output_type -> All
	1, // 9: EchoService.GetKind:output_type -> Kind
	2, // 10: EchoService.GetVersion:output_type -> Version
	3, // 11: EchoService.GetRegion:output_type -> Region
	4, // 12: EchoService.GetCluster:output_type -> Cluster
	5, // 13: EchoService.GetHostname:output_type -> Hostname
	6, // 14: EchoService.GetSourceIp:output_type -> SourceIp
	8, // 15: EchoService.GetHeader:output_type -> HeaderValue
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_echo_proto_init() }
func file_grpc_echo_proto_init() {
	if File_grpc_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*All); i {
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
		file_grpc_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kind); i {
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
		file_grpc_echo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_grpc_echo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
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
		file_grpc_echo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
		file_grpc_echo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hostname); i {
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
		file_grpc_echo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceIp); i {
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
		file_grpc_echo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderName); i {
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
		file_grpc_echo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderValue); i {
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
			RawDescriptor: file_grpc_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_echo_proto_goTypes,
		DependencyIndexes: file_grpc_echo_proto_depIdxs,
		MessageInfos:      file_grpc_echo_proto_msgTypes,
	}.Build()
	File_grpc_echo_proto = out.File
	file_grpc_echo_proto_rawDesc = nil
	file_grpc_echo_proto_goTypes = nil
	file_grpc_echo_proto_depIdxs = nil
}
