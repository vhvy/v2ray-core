package httphealthcheck

import (
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

// Config for HTTP health check server.
type HealthServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeout uint32 `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *HealthServerConfig) Reset() {
	*x = HealthServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_httphealthcheck_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthServerConfig) ProtoMessage() {}

func (x *HealthServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_httphealthcheck_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthServerConfig.ProtoReflect.Descriptor instead.
func (*HealthServerConfig) Descriptor() ([]byte, []int) {
	return file_proxy_httphealthcheck_config_proto_rawDescGZIP(), []int{0}
}

func (x *HealthServerConfig) GetTimeout() uint32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

var File_proxy_httphealthcheck_config_proto protoreflect.FileDescriptor

var file_proxy_httphealthcheck_config_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x2e, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x81, 0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x50,
	0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32,
	0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x35, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0xaa, 0x02, 0x20, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e,
	0x43, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proxy_httphealthcheck_config_proto_rawDescOnce sync.Once
	file_proxy_httphealthcheck_config_proto_rawDescData = file_proxy_httphealthcheck_config_proto_rawDesc
)

func file_proxy_httphealthcheck_config_proto_rawDescGZIP() []byte {
	file_proxy_httphealthcheck_config_proto_rawDescOnce.Do(func() {
		file_proxy_httphealthcheck_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_httphealthcheck_config_proto_rawDescData)
	})
	return file_proxy_httphealthcheck_config_proto_rawDescData
}

var file_proxy_httphealthcheck_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proxy_httphealthcheck_config_proto_goTypes = []interface{}{
	(*HealthServerConfig)(nil), // 0: v2ray.core.proxy.httphealthcheck.HealthServerConfig
}
var file_proxy_httphealthcheck_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proxy_httphealthcheck_config_proto_init() }
func file_proxy_httphealthcheck_config_proto_init() {
	if File_proxy_httphealthcheck_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_httphealthcheck_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthServerConfig); i {
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
			RawDescriptor: file_proxy_httphealthcheck_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_httphealthcheck_config_proto_goTypes,
		DependencyIndexes: file_proxy_httphealthcheck_config_proto_depIdxs,
		MessageInfos:      file_proxy_httphealthcheck_config_proto_msgTypes,
	}.Build()
	File_proxy_httphealthcheck_config_proto = out.File
	file_proxy_httphealthcheck_config_proto_rawDesc = nil
	file_proxy_httphealthcheck_config_proto_goTypes = nil
	file_proxy_httphealthcheck_config_proto_depIdxs = nil
}
