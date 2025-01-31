// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: PropExtraInfo.proto

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

type PropExtraInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueInfo *PropRogueInfo `protobuf:"bytes,11,opt,name=rogue_info,json=rogueInfo,proto3" json:"rogue_info,omitempty"`
}

func (x *PropExtraInfo) Reset() {
	*x = PropExtraInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PropExtraInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropExtraInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropExtraInfo) ProtoMessage() {}

func (x *PropExtraInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PropExtraInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropExtraInfo.ProtoReflect.Descriptor instead.
func (*PropExtraInfo) Descriptor() ([]byte, []int) {
	return file_PropExtraInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PropExtraInfo) GetRogueInfo() *PropRogueInfo {
	if x != nil {
		return x.RogueInfo
	}
	return nil
}

var File_PropExtraInfo_proto protoreflect.FileDescriptor

var file_PropExtraInfo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x50, 0x72, 0x6f, 0x70, 0x45, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x50, 0x72,
	0x6f, 0x70, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x44, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x70, 0x45, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x33, 0x0a, 0x0a, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x72, 0x6f,
	0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PropExtraInfo_proto_rawDescOnce sync.Once
	file_PropExtraInfo_proto_rawDescData = file_PropExtraInfo_proto_rawDesc
)

func file_PropExtraInfo_proto_rawDescGZIP() []byte {
	file_PropExtraInfo_proto_rawDescOnce.Do(func() {
		file_PropExtraInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PropExtraInfo_proto_rawDescData)
	})
	return file_PropExtraInfo_proto_rawDescData
}

var file_PropExtraInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PropExtraInfo_proto_goTypes = []interface{}{
	(*PropExtraInfo)(nil), // 0: proto.PropExtraInfo
	(*PropRogueInfo)(nil), // 1: proto.PropRogueInfo
}
var file_PropExtraInfo_proto_depIdxs = []int32{
	1, // 0: proto.PropExtraInfo.rogue_info:type_name -> proto.PropRogueInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PropExtraInfo_proto_init() }
func file_PropExtraInfo_proto_init() {
	if File_PropExtraInfo_proto != nil {
		return
	}
	file_PropRogueInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PropExtraInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropExtraInfo); i {
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
			RawDescriptor: file_PropExtraInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PropExtraInfo_proto_goTypes,
		DependencyIndexes: file_PropExtraInfo_proto_depIdxs,
		MessageInfos:      file_PropExtraInfo_proto_msgTypes,
	}.Build()
	File_PropExtraInfo_proto = out.File
	file_PropExtraInfo_proto_rawDesc = nil
	file_PropExtraInfo_proto_goTypes = nil
	file_PropExtraInfo_proto_depIdxs = nil
}
