// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: SyncRogueFinishScNotify.proto

package proto

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

type SyncRogueFinishScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinishInfo *RogueFinishInfo `protobuf:"bytes,11,opt,name=finish_info,json=finishInfo,proto3" json:"finish_info,omitempty"`
}

func (x *SyncRogueFinishScNotify) Reset() {
	*x = SyncRogueFinishScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SyncRogueFinishScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRogueFinishScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRogueFinishScNotify) ProtoMessage() {}

func (x *SyncRogueFinishScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SyncRogueFinishScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRogueFinishScNotify.ProtoReflect.Descriptor instead.
func (*SyncRogueFinishScNotify) Descriptor() ([]byte, []int) {
	return file_SyncRogueFinishScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SyncRogueFinishScNotify) GetFinishInfo() *RogueFinishInfo {
	if x != nil {
		return x.FinishInfo
	}
	return nil
}

var File_SyncRogueFinishScNotify_proto protoreflect.FileDescriptor

var file_SyncRogueFinishScNotify_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a,
	0x17, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x37, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SyncRogueFinishScNotify_proto_rawDescOnce sync.Once
	file_SyncRogueFinishScNotify_proto_rawDescData = file_SyncRogueFinishScNotify_proto_rawDesc
)

func file_SyncRogueFinishScNotify_proto_rawDescGZIP() []byte {
	file_SyncRogueFinishScNotify_proto_rawDescOnce.Do(func() {
		file_SyncRogueFinishScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SyncRogueFinishScNotify_proto_rawDescData)
	})
	return file_SyncRogueFinishScNotify_proto_rawDescData
}

var file_SyncRogueFinishScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SyncRogueFinishScNotify_proto_goTypes = []interface{}{
	(*SyncRogueFinishScNotify)(nil), // 0: proto.SyncRogueFinishScNotify
	(*RogueFinishInfo)(nil),         // 1: proto.RogueFinishInfo
}
var file_SyncRogueFinishScNotify_proto_depIdxs = []int32{
	1, // 0: proto.SyncRogueFinishScNotify.finish_info:type_name -> proto.RogueFinishInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SyncRogueFinishScNotify_proto_init() }
func file_SyncRogueFinishScNotify_proto_init() {
	if File_SyncRogueFinishScNotify_proto != nil {
		return
	}
	file_RogueFinishInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SyncRogueFinishScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRogueFinishScNotify); i {
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
			RawDescriptor: file_SyncRogueFinishScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SyncRogueFinishScNotify_proto_goTypes,
		DependencyIndexes: file_SyncRogueFinishScNotify_proto_depIdxs,
		MessageInfos:      file_SyncRogueFinishScNotify_proto_msgTypes,
	}.Build()
	File_SyncRogueFinishScNotify_proto = out.File
	file_SyncRogueFinishScNotify_proto_rawDesc = nil
	file_SyncRogueFinishScNotify_proto_goTypes = nil
	file_SyncRogueFinishScNotify_proto_depIdxs = nil
}
