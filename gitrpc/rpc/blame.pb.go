// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: blame.proto

package rpc

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

type BlameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base   *ReadRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	GitRef string       `protobuf:"bytes,2,opt,name=git_ref,json=gitRef,proto3" json:"git_ref,omitempty"`
	Path   string       `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Range  *LineRange   `protobuf:"bytes,4,opt,name=range,proto3" json:"range,omitempty"`
}

func (x *BlameRequest) Reset() {
	*x = BlameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlameRequest) ProtoMessage() {}

func (x *BlameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlameRequest.ProtoReflect.Descriptor instead.
func (*BlameRequest) Descriptor() ([]byte, []int) {
	return file_blame_proto_rawDescGZIP(), []int{0}
}

func (x *BlameRequest) GetBase() *ReadRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *BlameRequest) GetGitRef() string {
	if x != nil {
		return x.GitRef
	}
	return ""
}

func (x *BlameRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *BlameRequest) GetRange() *LineRange {
	if x != nil {
		return x.Range
	}
	return nil
}

type LineRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From int32 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To   int32 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *LineRange) Reset() {
	*x = LineRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blame_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineRange) ProtoMessage() {}

func (x *LineRange) ProtoReflect() protoreflect.Message {
	mi := &file_blame_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineRange.ProtoReflect.Descriptor instead.
func (*LineRange) Descriptor() ([]byte, []int) {
	return file_blame_proto_rawDescGZIP(), []int{1}
}

func (x *LineRange) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *LineRange) GetTo() int32 {
	if x != nil {
		return x.To
	}
	return 0
}

type BlamePart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commit *Commit  `protobuf:"bytes,1,opt,name=commit,proto3" json:"commit,omitempty"`
	Lines  [][]byte `protobuf:"bytes,2,rep,name=lines,proto3" json:"lines,omitempty"`
}

func (x *BlamePart) Reset() {
	*x = BlamePart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blame_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlamePart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlamePart) ProtoMessage() {}

func (x *BlamePart) ProtoReflect() protoreflect.Message {
	mi := &file_blame_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlamePart.ProtoReflect.Descriptor instead.
func (*BlamePart) Descriptor() ([]byte, []int) {
	return file_blame_proto_rawDescGZIP(), []int{2}
}

func (x *BlamePart) GetCommit() *Commit {
	if x != nil {
		return x.Commit
	}
	return nil
}

func (x *BlamePart) GetLines() [][]byte {
	if x != nil {
		return x.Lines
	}
	return nil
}

var File_blame_proto protoreflect.FileDescriptor

var file_blame_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6c, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72,
	0x70, 0x63, 0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x87, 0x01, 0x0a, 0x0c, 0x42, 0x6c, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x69, 0x74, 0x5f, 0x72,
	0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x69, 0x74, 0x52, 0x65, 0x66,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x2f, 0x0a, 0x09, 0x4c, 0x69,
	0x6e, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x09, 0x42,
	0x6c, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x32, 0x3c, 0x0a, 0x0c, 0x42, 0x6c, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x42, 0x6c, 0x61, 0x6d, 0x65, 0x12, 0x11, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x42, 0x6c, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x6c, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x74, 0x30,
	0x01, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x2f,
	0x67, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_blame_proto_rawDescOnce sync.Once
	file_blame_proto_rawDescData = file_blame_proto_rawDesc
)

func file_blame_proto_rawDescGZIP() []byte {
	file_blame_proto_rawDescOnce.Do(func() {
		file_blame_proto_rawDescData = protoimpl.X.CompressGZIP(file_blame_proto_rawDescData)
	})
	return file_blame_proto_rawDescData
}

var file_blame_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_blame_proto_goTypes = []interface{}{
	(*BlameRequest)(nil), // 0: rpc.BlameRequest
	(*LineRange)(nil),    // 1: rpc.LineRange
	(*BlamePart)(nil),    // 2: rpc.BlamePart
	(*ReadRequest)(nil),  // 3: rpc.ReadRequest
	(*Commit)(nil),       // 4: rpc.Commit
}
var file_blame_proto_depIdxs = []int32{
	3, // 0: rpc.BlameRequest.base:type_name -> rpc.ReadRequest
	1, // 1: rpc.BlameRequest.range:type_name -> rpc.LineRange
	4, // 2: rpc.BlamePart.commit:type_name -> rpc.Commit
	0, // 3: rpc.BlameService.Blame:input_type -> rpc.BlameRequest
	2, // 4: rpc.BlameService.Blame:output_type -> rpc.BlamePart
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_blame_proto_init() }
func file_blame_proto_init() {
	if File_blame_proto != nil {
		return
	}
	file_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_blame_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlameRequest); i {
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
		file_blame_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineRange); i {
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
		file_blame_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlamePart); i {
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
			RawDescriptor: file_blame_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blame_proto_goTypes,
		DependencyIndexes: file_blame_proto_depIdxs,
		MessageInfos:      file_blame_proto_msgTypes,
	}.Build()
	File_blame_proto = out.File
	file_blame_proto_rawDesc = nil
	file_blame_proto_goTypes = nil
	file_blame_proto_depIdxs = nil
}