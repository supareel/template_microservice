// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: constants.proto

package __

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

type DBOPERATIONSTATUS int32

const (
	DBOPERATIONSTATUS_QUERIED DBOPERATIONSTATUS = 0
	DBOPERATIONSTATUS_CREATED DBOPERATIONSTATUS = 1
	DBOPERATIONSTATUS_UPDATED DBOPERATIONSTATUS = 2
	DBOPERATIONSTATUS_DELETED DBOPERATIONSTATUS = 3
	DBOPERATIONSTATUS_ERROR   DBOPERATIONSTATUS = 5
)

// Enum value maps for DBOPERATIONSTATUS.
var (
	DBOPERATIONSTATUS_name = map[int32]string{
		0: "QUERIED",
		1: "CREATED",
		2: "UPDATED",
		3: "DELETED",
		5: "ERROR",
	}
	DBOPERATIONSTATUS_value = map[string]int32{
		"QUERIED": 0,
		"CREATED": 1,
		"UPDATED": 2,
		"DELETED": 3,
		"ERROR":   5,
	}
)

func (x DBOPERATIONSTATUS) Enum() *DBOPERATIONSTATUS {
	p := new(DBOPERATIONSTATUS)
	*p = x
	return p
}

func (x DBOPERATIONSTATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DBOPERATIONSTATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_constants_proto_enumTypes[0].Descriptor()
}

func (DBOPERATIONSTATUS) Type() protoreflect.EnumType {
	return &file_constants_proto_enumTypes[0]
}

func (x DBOPERATIONSTATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DBOPERATIONSTATUS.Descriptor instead.
func (DBOPERATIONSTATUS) EnumDescriptor() ([]byte, []int) {
	return file_constants_proto_rawDescGZIP(), []int{0}
}

var File_constants_proto protoreflect.FileDescriptor

var file_constants_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x52, 0x0a, 0x11, 0x44, 0x42, 0x4f, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x0b, 0x0a,
	0x07, 0x51, 0x55, 0x45, 0x52, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x42, 0x03, 0x5a, 0x01,
	0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_constants_proto_rawDescOnce sync.Once
	file_constants_proto_rawDescData = file_constants_proto_rawDesc
)

func file_constants_proto_rawDescGZIP() []byte {
	file_constants_proto_rawDescOnce.Do(func() {
		file_constants_proto_rawDescData = protoimpl.X.CompressGZIP(file_constants_proto_rawDescData)
	})
	return file_constants_proto_rawDescData
}

var file_constants_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_constants_proto_goTypes = []interface{}{
	(DBOPERATIONSTATUS)(0), // 0: proto.DBOPERATIONSTATUS
}
var file_constants_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_constants_proto_init() }
func file_constants_proto_init() {
	if File_constants_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_constants_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_constants_proto_goTypes,
		DependencyIndexes: file_constants_proto_depIdxs,
		EnumInfos:         file_constants_proto_enumTypes,
	}.Build()
	File_constants_proto = out.File
	file_constants_proto_rawDesc = nil
	file_constants_proto_goTypes = nil
	file_constants_proto_depIdxs = nil
}