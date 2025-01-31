// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: ActivityScheduleInfo.proto

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

type ActivityScheduleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId uint32 `protobuf:"varint,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	EndTime    int64  `protobuf:"varint,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	ModuleId   uint32 `protobuf:"varint,5,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	BeginTime  int64  `protobuf:"varint,8,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
}

func (x *ActivityScheduleInfo) Reset() {
	*x = ActivityScheduleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActivityScheduleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityScheduleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityScheduleInfo) ProtoMessage() {}

func (x *ActivityScheduleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ActivityScheduleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityScheduleInfo.ProtoReflect.Descriptor instead.
func (*ActivityScheduleInfo) Descriptor() ([]byte, []int) {
	return file_ActivityScheduleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityScheduleInfo) GetActivityId() uint32 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

func (x *ActivityScheduleInfo) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *ActivityScheduleInfo) GetModuleId() uint32 {
	if x != nil {
		return x.ModuleId
	}
	return 0
}

func (x *ActivityScheduleInfo) GetBeginTime() int64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

var File_ActivityScheduleInfo_proto protoreflect.FileDescriptor

var file_ActivityScheduleInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x14, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ActivityScheduleInfo_proto_rawDescOnce sync.Once
	file_ActivityScheduleInfo_proto_rawDescData = file_ActivityScheduleInfo_proto_rawDesc
)

func file_ActivityScheduleInfo_proto_rawDescGZIP() []byte {
	file_ActivityScheduleInfo_proto_rawDescOnce.Do(func() {
		file_ActivityScheduleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ActivityScheduleInfo_proto_rawDescData)
	})
	return file_ActivityScheduleInfo_proto_rawDescData
}

var file_ActivityScheduleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ActivityScheduleInfo_proto_goTypes = []interface{}{
	(*ActivityScheduleInfo)(nil), // 0: proto.ActivityScheduleInfo
}
var file_ActivityScheduleInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ActivityScheduleInfo_proto_init() }
func file_ActivityScheduleInfo_proto_init() {
	if File_ActivityScheduleInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ActivityScheduleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityScheduleInfo); i {
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
			RawDescriptor: file_ActivityScheduleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ActivityScheduleInfo_proto_goTypes,
		DependencyIndexes: file_ActivityScheduleInfo_proto_depIdxs,
		MessageInfos:      file_ActivityScheduleInfo_proto_msgTypes,
	}.Build()
	File_ActivityScheduleInfo_proto = out.File
	file_ActivityScheduleInfo_proto_rawDesc = nil
	file_ActivityScheduleInfo_proto_goTypes = nil
	file_ActivityScheduleInfo_proto_depIdxs = nil
}
