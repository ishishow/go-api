// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: ranking.proto

package ranking

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Rank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Rank     int32  `protobuf:"varint,3,opt,name=rank,proto3" json:"rank,omitempty"`
	Score    int32  `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Rank) Reset() {
	*x = Rank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ranking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rank) ProtoMessage() {}

func (x *Rank) ProtoReflect() protoreflect.Message {
	mi := &file_ranking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rank.ProtoReflect.Descriptor instead.
func (*Rank) Descriptor() ([]byte, []int) {
	return file_ranking_proto_rawDescGZIP(), []int{0}
}

func (x *Rank) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Rank) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Rank) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Rank) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type RankingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start string `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
}

func (x *RankingListRequest) Reset() {
	*x = RankingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ranking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankingListRequest) ProtoMessage() {}

func (x *RankingListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ranking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankingListRequest.ProtoReflect.Descriptor instead.
func (*RankingListRequest) Descriptor() ([]byte, []int) {
	return file_ranking_proto_rawDescGZIP(), []int{1}
}

func (x *RankingListRequest) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

type RankingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranks []*Rank `protobuf:"bytes,1,rep,name=ranks,proto3" json:"ranks,omitempty"`
}

func (x *RankingListResponse) Reset() {
	*x = RankingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ranking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankingListResponse) ProtoMessage() {}

func (x *RankingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ranking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankingListResponse.ProtoReflect.Descriptor instead.
func (*RankingListResponse) Descriptor() ([]byte, []int) {
	return file_ranking_proto_rawDescGZIP(), []int{2}
}

func (x *RankingListResponse) GetRanks() []*Rank {
	if x != nil {
		return x.Ranks
	}
	return nil
}

var File_ranking_proto protoreflect.FileDescriptor

var file_ranking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x2a, 0x0a, 0x12, 0x52,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x22, 0x37, 0x0a, 0x13, 0x52, 0x61, 0x6e, 0x6b, 0x69,
	0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x05, 0x72, 0x61, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x6b, 0x73,
	0x32, 0x6b, 0x0a, 0x0e, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x59, 0x0a, 0x0b, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d,
	0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x3b, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ranking_proto_rawDescOnce sync.Once
	file_ranking_proto_rawDescData = file_ranking_proto_rawDesc
)

func file_ranking_proto_rawDescGZIP() []byte {
	file_ranking_proto_rawDescOnce.Do(func() {
		file_ranking_proto_rawDescData = protoimpl.X.CompressGZIP(file_ranking_proto_rawDescData)
	})
	return file_ranking_proto_rawDescData
}

var file_ranking_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ranking_proto_goTypes = []interface{}{
	(*Rank)(nil),                // 0: user.Rank
	(*RankingListRequest)(nil),  // 1: user.RankingListRequest
	(*RankingListResponse)(nil), // 2: user.RankingListResponse
}
var file_ranking_proto_depIdxs = []int32{
	0, // 0: user.RankingListResponse.ranks:type_name -> user.Rank
	1, // 1: user.RankingService.RankingList:input_type -> user.RankingListRequest
	2, // 2: user.RankingService.RankingList:output_type -> user.RankingListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ranking_proto_init() }
func file_ranking_proto_init() {
	if File_ranking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ranking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rank); i {
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
		file_ranking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankingListRequest); i {
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
		file_ranking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankingListResponse); i {
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
			RawDescriptor: file_ranking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ranking_proto_goTypes,
		DependencyIndexes: file_ranking_proto_depIdxs,
		MessageInfos:      file_ranking_proto_msgTypes,
	}.Build()
	File_ranking_proto = out.File
	file_ranking_proto_rawDesc = nil
	file_ranking_proto_goTypes = nil
	file_ranking_proto_depIdxs = nil
}