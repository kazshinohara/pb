// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: whereami.proto

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

type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind     string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Version  string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Region   string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Cluster  string `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Hostname string `protobuf:"bytes,5,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_whereami_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_whereami_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_whereami_proto_rawDescGZIP(), []int{0}
}

func (x *ServerInfo) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ServerInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServerInfo) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ServerInfo) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *ServerInfo) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

var File_whereami_proto protoreflect.FileDescriptor

var file_whereami_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x77, 0x68, 0x65, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01,
	0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x42, 0x0a, 0x08, 0x57, 0x68, 0x65, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x12, 0x36, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x7a, 0x73, 0x68,
	0x69, 0x6e, 0x6f, 0x68, 0x61, 0x72, 0x61, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x68, 0x65, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_whereami_proto_rawDescOnce sync.Once
	file_whereami_proto_rawDescData = file_whereami_proto_rawDesc
)

func file_whereami_proto_rawDescGZIP() []byte {
	file_whereami_proto_rawDescOnce.Do(func() {
		file_whereami_proto_rawDescData = protoimpl.X.CompressGZIP(file_whereami_proto_rawDescData)
	})
	return file_whereami_proto_rawDescData
}

var file_whereami_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_whereami_proto_goTypes = []interface{}{
	(*ServerInfo)(nil),    // 0: ServerInfo
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_whereami_proto_depIdxs = []int32{
	1, // 0: Whereami.GetServerInfo:input_type -> google.protobuf.Empty
	0, // 1: Whereami.GetServerInfo:output_type -> ServerInfo
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_whereami_proto_init() }
func file_whereami_proto_init() {
	if File_whereami_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_whereami_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfo); i {
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
			RawDescriptor: file_whereami_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_whereami_proto_goTypes,
		DependencyIndexes: file_whereami_proto_depIdxs,
		MessageInfos:      file_whereami_proto_msgTypes,
	}.Build()
	File_whereami_proto = out.File
	file_whereami_proto_rawDesc = nil
	file_whereami_proto_goTypes = nil
	file_whereami_proto_depIdxs = nil
}
