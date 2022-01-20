// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: proto/constants.proto

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

type DBOPERATIONSTATUS int32

const (
	DBOPERATIONSTATUS_CREATED_RECORD_TO_DB     DBOPERATIONSTATUS = 0
	DBOPERATIONSTATUS_UPDATED_RECORD_FROM_DB   DBOPERATIONSTATUS = 1
	DBOPERATIONSTATUS_QUERIED_RECORD_FROM_DB   DBOPERATIONSTATUS = 2
	DBOPERATIONSTATUS_DELETED_RECORD_FROM_DB   DBOPERATIONSTATUS = 3
	DBOPERATIONSTATUS_ERROR_WHILE_SAVING_IN_DB DBOPERATIONSTATUS = 4
)

// Enum value maps for DBOPERATIONSTATUS.
var (
	DBOPERATIONSTATUS_name = map[int32]string{
		0: "CREATED_RECORD_TO_DB",
		1: "UPDATED_RECORD_FROM_DB",
		2: "QUERIED_RECORD_FROM_DB",
		3: "DELETED_RECORD_FROM_DB",
		4: "ERROR_WHILE_SAVING_IN_DB",
	}
	DBOPERATIONSTATUS_value = map[string]int32{
		"CREATED_RECORD_TO_DB":     0,
		"UPDATED_RECORD_FROM_DB":   1,
		"QUERIED_RECORD_FROM_DB":   2,
		"DELETED_RECORD_FROM_DB":   3,
		"ERROR_WHILE_SAVING_IN_DB": 4,
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
	return file_proto_constants_proto_enumTypes[0].Descriptor()
}

func (DBOPERATIONSTATUS) Type() protoreflect.EnumType {
	return &file_proto_constants_proto_enumTypes[0]
}

func (x DBOPERATIONSTATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DBOPERATIONSTATUS.Descriptor instead.
func (DBOPERATIONSTATUS) EnumDescriptor() ([]byte, []int) {
	return file_proto_constants_proto_rawDescGZIP(), []int{0}
}

var File_proto_constants_proto protoreflect.FileDescriptor

var file_proto_constants_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2a, 0x9f, 0x01, 0x0a, 0x11, 0x44, 0x42, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x44, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x44, 0x42, 0x10, 0x00,
	0x12, 0x1a, 0x0a, 0x16, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x43, 0x4f,
	0x52, 0x44, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x44, 0x42, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16,
	0x51, 0x55, 0x45, 0x52, 0x49, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x46,
	0x52, 0x4f, 0x4d, 0x5f, 0x44, 0x42, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f,
	0x44, 0x42, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x57, 0x48,
	0x49, 0x4c, 0x45, 0x5f, 0x53, 0x41, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4e, 0x5f, 0x44, 0x42,
	0x10, 0x04, 0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x6f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_constants_proto_rawDescOnce sync.Once
	file_proto_constants_proto_rawDescData = file_proto_constants_proto_rawDesc
)

func file_proto_constants_proto_rawDescGZIP() []byte {
	file_proto_constants_proto_rawDescOnce.Do(func() {
		file_proto_constants_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_constants_proto_rawDescData)
	})
	return file_proto_constants_proto_rawDescData
}

var file_proto_constants_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_constants_proto_goTypes = []interface{}{
	(DBOPERATIONSTATUS)(0), // 0: account.DBOPERATIONSTATUS
}
var file_proto_constants_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_constants_proto_init() }
func file_proto_constants_proto_init() {
	if File_proto_constants_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_constants_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_constants_proto_goTypes,
		DependencyIndexes: file_proto_constants_proto_depIdxs,
		EnumInfos:         file_proto_constants_proto_enumTypes,
	}.Build()
	File_proto_constants_proto = out.File
	file_proto_constants_proto_rawDesc = nil
	file_proto_constants_proto_goTypes = nil
	file_proto_constants_proto_depIdxs = nil
}
