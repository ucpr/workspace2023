// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: proto/example/content/v1/content.proto

package contentv1

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

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt int64  `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_example_content_v1_content_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_proto_example_content_v1_content_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_proto_example_content_v1_content_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Content) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Content) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Content) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_proto_example_content_v1_content_proto protoreflect.FileDescriptor

var file_proto_example_content_v1_content_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x22, 0x68, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0xb5,
	0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x63,
	0x70, 0x72, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x32, 0x30, 0x32, 0x33,
	0x2f, 0x62, 0x75, 0x66, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_example_content_v1_content_proto_rawDescOnce sync.Once
	file_proto_example_content_v1_content_proto_rawDescData = file_proto_example_content_v1_content_proto_rawDesc
)

func file_proto_example_content_v1_content_proto_rawDescGZIP() []byte {
	file_proto_example_content_v1_content_proto_rawDescOnce.Do(func() {
		file_proto_example_content_v1_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_example_content_v1_content_proto_rawDescData)
	})
	return file_proto_example_content_v1_content_proto_rawDescData
}

var file_proto_example_content_v1_content_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_example_content_v1_content_proto_goTypes = []interface{}{
	(*Content)(nil), // 0: content.v1.Content
}
var file_proto_example_content_v1_content_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_example_content_v1_content_proto_init() }
func file_proto_example_content_v1_content_proto_init() {
	if File_proto_example_content_v1_content_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_example_content_v1_content_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
			RawDescriptor: file_proto_example_content_v1_content_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_example_content_v1_content_proto_goTypes,
		DependencyIndexes: file_proto_example_content_v1_content_proto_depIdxs,
		MessageInfos:      file_proto_example_content_v1_content_proto_msgTypes,
	}.Build()
	File_proto_example_content_v1_content_proto = out.File
	file_proto_example_content_v1_content_proto_rawDesc = nil
	file_proto_example_content_v1_content_proto_goTypes = nil
	file_proto_example_content_v1_content_proto_depIdxs = nil
}