// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: GetChallengeScRsp.proto

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

type GetChallengeScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeList       []*Challenge       `protobuf:"bytes,8,rep,name=challenge_list,json=challengeList,proto3" json:"challenge_list,omitempty"`
	Retcode             uint32             `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ChallengeRewardList []*ChallengeReward `protobuf:"bytes,9,rep,name=challenge_reward_list,json=challengeRewardList,proto3" json:"challenge_reward_list,omitempty"`
}

func (x *GetChallengeScRsp) Reset() {
	*x = GetChallengeScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetChallengeScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChallengeScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChallengeScRsp) ProtoMessage() {}

func (x *GetChallengeScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetChallengeScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChallengeScRsp.ProtoReflect.Descriptor instead.
func (*GetChallengeScRsp) Descriptor() ([]byte, []int) {
	return file_GetChallengeScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetChallengeScRsp) GetChallengeList() []*Challenge {
	if x != nil {
		return x.ChallengeList
	}
	return nil
}

func (x *GetChallengeScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetChallengeScRsp) GetChallengeRewardList() []*ChallengeReward {
	if x != nil {
		return x.ChallengeRewardList
	}
	return nil
}

var File_GetChallengeScRsp_proto protoreflect.FileDescriptor

var file_GetChallengeScRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x37,
	0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x4a, 0x0a, 0x15, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x13, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetChallengeScRsp_proto_rawDescOnce sync.Once
	file_GetChallengeScRsp_proto_rawDescData = file_GetChallengeScRsp_proto_rawDesc
)

func file_GetChallengeScRsp_proto_rawDescGZIP() []byte {
	file_GetChallengeScRsp_proto_rawDescOnce.Do(func() {
		file_GetChallengeScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetChallengeScRsp_proto_rawDescData)
	})
	return file_GetChallengeScRsp_proto_rawDescData
}

var file_GetChallengeScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetChallengeScRsp_proto_goTypes = []interface{}{
	(*GetChallengeScRsp)(nil), // 0: proto.GetChallengeScRsp
	(*Challenge)(nil),         // 1: proto.Challenge
	(*ChallengeReward)(nil),   // 2: proto.ChallengeReward
}
var file_GetChallengeScRsp_proto_depIdxs = []int32{
	1, // 0: proto.GetChallengeScRsp.challenge_list:type_name -> proto.Challenge
	2, // 1: proto.GetChallengeScRsp.challenge_reward_list:type_name -> proto.ChallengeReward
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetChallengeScRsp_proto_init() }
func file_GetChallengeScRsp_proto_init() {
	if File_GetChallengeScRsp_proto != nil {
		return
	}
	file_Challenge_proto_init()
	file_ChallengeReward_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetChallengeScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChallengeScRsp); i {
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
			RawDescriptor: file_GetChallengeScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetChallengeScRsp_proto_goTypes,
		DependencyIndexes: file_GetChallengeScRsp_proto_depIdxs,
		MessageInfos:      file_GetChallengeScRsp_proto_msgTypes,
	}.Build()
	File_GetChallengeScRsp_proto = out.File
	file_GetChallengeScRsp_proto_rawDesc = nil
	file_GetChallengeScRsp_proto_goTypes = nil
	file_GetChallengeScRsp_proto_depIdxs = nil
}
