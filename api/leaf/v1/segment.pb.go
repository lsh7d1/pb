// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: leaf/v1/segment.proto

package v1

import (
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

type GenIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizTag string `protobuf:"bytes,1,opt,name=bizTag,proto3" json:"bizTag,omitempty"`
	Count  int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GenIdsRequest) Reset() {
	*x = GenIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaf_v1_segment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenIdsRequest) ProtoMessage() {}

func (x *GenIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_leaf_v1_segment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenIdsRequest.ProtoReflect.Descriptor instead.
func (*GenIdsRequest) Descriptor() ([]byte, []int) {
	return file_leaf_v1_segment_proto_rawDescGZIP(), []int{0}
}

func (x *GenIdsRequest) GetBizTag() string {
	if x != nil {
		return x.BizTag
	}
	return ""
}

func (x *GenIdsRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GenIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GenIdsResponse) Reset() {
	*x = GenIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaf_v1_segment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenIdsResponse) ProtoMessage() {}

func (x *GenIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_leaf_v1_segment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenIdsResponse.ProtoReflect.Descriptor instead.
func (*GenIdsResponse) Descriptor() ([]byte, []int) {
	return file_leaf_v1_segment_proto_rawDescGZIP(), []int{1}
}

func (x *GenIdsResponse) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_leaf_v1_segment_proto protoreflect.FileDescriptor

var file_leaf_v1_segment_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x65, 0x61, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x65, 0x61, 0x66, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d,
	0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x54, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x69, 0x7a, 0x54, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x22, 0x0a,
	0x0e, 0x47, 0x65, 0x6e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64,
	0x73, 0x32, 0x7d, 0x0a, 0x12, 0x4c, 0x65, 0x61, 0x66, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x6c, 0x65, 0x61, 0x66, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x6e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x6c, 0x65, 0x61, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x49, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x62, 0x03, 0x69, 0x64, 0x73, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x62, 0x69, 0x7a, 0x54, 0x61, 0x67, 0x7d,
	0x42, 0x10, 0x5a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x65, 0x61, 0x66, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_leaf_v1_segment_proto_rawDescOnce sync.Once
	file_leaf_v1_segment_proto_rawDescData = file_leaf_v1_segment_proto_rawDesc
)

func file_leaf_v1_segment_proto_rawDescGZIP() []byte {
	file_leaf_v1_segment_proto_rawDescOnce.Do(func() {
		file_leaf_v1_segment_proto_rawDescData = protoimpl.X.CompressGZIP(file_leaf_v1_segment_proto_rawDescData)
	})
	return file_leaf_v1_segment_proto_rawDescData
}

var file_leaf_v1_segment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_leaf_v1_segment_proto_goTypes = []interface{}{
	(*GenIdsRequest)(nil),  // 0: leaf.v1.GenIdsRequest
	(*GenIdsResponse)(nil), // 1: leaf.v1.GenIdsResponse
}
var file_leaf_v1_segment_proto_depIdxs = []int32{
	0, // 0: leaf.v1.LeafSegmentService.GetSegmentId:input_type -> leaf.v1.GenIdsRequest
	1, // 1: leaf.v1.LeafSegmentService.GetSegmentId:output_type -> leaf.v1.GenIdsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_leaf_v1_segment_proto_init() }
func file_leaf_v1_segment_proto_init() {
	if File_leaf_v1_segment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_leaf_v1_segment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenIdsRequest); i {
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
		file_leaf_v1_segment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenIdsResponse); i {
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
			RawDescriptor: file_leaf_v1_segment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_leaf_v1_segment_proto_goTypes,
		DependencyIndexes: file_leaf_v1_segment_proto_depIdxs,
		MessageInfos:      file_leaf_v1_segment_proto_msgTypes,
	}.Build()
	File_leaf_v1_segment_proto = out.File
	file_leaf_v1_segment_proto_rawDesc = nil
	file_leaf_v1_segment_proto_goTypes = nil
	file_leaf_v1_segment_proto_depIdxs = nil
}