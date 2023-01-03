// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.7
// source: google/ads/googleads/v12/errors/asset_set_error.proto

package errors

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

// Enum describing possible asset set errors.
type AssetSetErrorEnum_AssetSetError int32

const (
	// Enum unspecified.
	AssetSetErrorEnum_UNSPECIFIED AssetSetErrorEnum_AssetSetError = 0
	// The received error code is not known in this version.
	AssetSetErrorEnum_UNKNOWN AssetSetErrorEnum_AssetSetError = 1
	// The asset set name matches that of another enabled asset set.
	AssetSetErrorEnum_DUPLICATE_ASSET_SET_NAME AssetSetErrorEnum_AssetSetError = 2
	// The type of AssetSet.asset_set_source does not match the type of
	// AssetSet.location_set.source in its parent AssetSet.
	AssetSetErrorEnum_INVALID_PARENT_ASSET_SET_TYPE AssetSetErrorEnum_AssetSetError = 3
	// The asset set source doesn't match its parent AssetSet's data.
	AssetSetErrorEnum_ASSET_SET_SOURCE_INCOMPATIBLE_WITH_PARENT_ASSET_SET AssetSetErrorEnum_AssetSetError = 4
	// This AssetSet type cannot be linked to CustomerAssetSet.
	AssetSetErrorEnum_ASSET_SET_TYPE_CANNOT_BE_LINKED_TO_CUSTOMER AssetSetErrorEnum_AssetSetError = 5
	// The chain id(s) in ChainSet of a LOCATION_SYNC typed AssetSet is invalid.
	AssetSetErrorEnum_INVALID_CHAIN_IDS AssetSetErrorEnum_AssetSetError = 6
	// The relationship type in ChainSet of a LOCATION_SYNC typed AssetSet is
	// not supported.
	AssetSetErrorEnum_LOCATION_SYNC_ASSET_SET_DOES_NOT_SUPPORT_RELATIONSHIP_TYPE AssetSetErrorEnum_AssetSetError = 7
	// There is more than one enabled LocationSync typed AssetSet under one
	// customer.
	AssetSetErrorEnum_NOT_UNIQUE_ENABLED_LOCATION_SYNC_TYPED_ASSET_SET AssetSetErrorEnum_AssetSetError = 8
	// The place id(s) in a LocationSync typed AssetSet is invalid and can't be
	// decoded.
	AssetSetErrorEnum_INVALID_PLACE_IDS AssetSetErrorEnum_AssetSetError = 9
)

// Enum value maps for AssetSetErrorEnum_AssetSetError.
var (
	AssetSetErrorEnum_AssetSetError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "DUPLICATE_ASSET_SET_NAME",
		3: "INVALID_PARENT_ASSET_SET_TYPE",
		4: "ASSET_SET_SOURCE_INCOMPATIBLE_WITH_PARENT_ASSET_SET",
		5: "ASSET_SET_TYPE_CANNOT_BE_LINKED_TO_CUSTOMER",
		6: "INVALID_CHAIN_IDS",
		7: "LOCATION_SYNC_ASSET_SET_DOES_NOT_SUPPORT_RELATIONSHIP_TYPE",
		8: "NOT_UNIQUE_ENABLED_LOCATION_SYNC_TYPED_ASSET_SET",
		9: "INVALID_PLACE_IDS",
	}
	AssetSetErrorEnum_AssetSetError_value = map[string]int32{
		"UNSPECIFIED":                   0,
		"UNKNOWN":                       1,
		"DUPLICATE_ASSET_SET_NAME":      2,
		"INVALID_PARENT_ASSET_SET_TYPE": 3,
		"ASSET_SET_SOURCE_INCOMPATIBLE_WITH_PARENT_ASSET_SET": 4,
		"ASSET_SET_TYPE_CANNOT_BE_LINKED_TO_CUSTOMER":         5,
		"INVALID_CHAIN_IDS": 6,
		"LOCATION_SYNC_ASSET_SET_DOES_NOT_SUPPORT_RELATIONSHIP_TYPE": 7,
		"NOT_UNIQUE_ENABLED_LOCATION_SYNC_TYPED_ASSET_SET":           8,
		"INVALID_PLACE_IDS": 9,
	}
)

func (x AssetSetErrorEnum_AssetSetError) Enum() *AssetSetErrorEnum_AssetSetError {
	p := new(AssetSetErrorEnum_AssetSetError)
	*p = x
	return p
}

func (x AssetSetErrorEnum_AssetSetError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetSetErrorEnum_AssetSetError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v12_errors_asset_set_error_proto_enumTypes[0].Descriptor()
}

func (AssetSetErrorEnum_AssetSetError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v12_errors_asset_set_error_proto_enumTypes[0]
}

func (x AssetSetErrorEnum_AssetSetError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetSetErrorEnum_AssetSetError.Descriptor instead.
func (AssetSetErrorEnum_AssetSetError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_errors_asset_set_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible asset set errors.
type AssetSetErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssetSetErrorEnum) Reset() {
	*x = AssetSetErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_errors_asset_set_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetSetErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetSetErrorEnum) ProtoMessage() {}

func (x *AssetSetErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_errors_asset_set_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetSetErrorEnum.ProtoReflect.Descriptor instead.
func (*AssetSetErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_errors_asset_set_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v12_errors_asset_set_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_errors_asset_set_error_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x32, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x92, 0x03, 0x0a, 0x11, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x53, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xfc,
	0x02, 0x0a, 0x0d, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1c,
	0x0a, 0x18, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45,
	0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x41,
	0x53, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x03, 0x12,
	0x37, 0x0a, 0x33, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45,
	0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x53, 0x53,
	0x45, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x04, 0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x53, 0x53, 0x45,
	0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f,
	0x54, 0x5f, 0x42, 0x45, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x43,
	0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x53, 0x10, 0x06,
	0x12, 0x3e, 0x0a, 0x3a, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x4e,
	0x43, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x44, 0x4f, 0x45, 0x53,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x4c,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x07,
	0x12, 0x34, 0x0a, 0x30, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x45,
	0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x59, 0x4e, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x44, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54,
	0x5f, 0x53, 0x45, 0x54, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x53, 0x10, 0x09, 0x42, 0xf2, 0x01,
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x12, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e,
	0x56, 0x31, 0x32, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x31, 0x32, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x32, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v12_errors_asset_set_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_errors_asset_set_error_proto_rawDescData = file_google_ads_googleads_v12_errors_asset_set_error_proto_rawDesc
)

func file_google_ads_googleads_v12_errors_asset_set_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_errors_asset_set_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_errors_asset_set_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_errors_asset_set_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_errors_asset_set_error_proto_rawDescData
}

var file_google_ads_googleads_v12_errors_asset_set_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v12_errors_asset_set_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v12_errors_asset_set_error_proto_goTypes = []interface{}{
	(AssetSetErrorEnum_AssetSetError)(0), // 0: google.ads.googleads.v12.errors.AssetSetErrorEnum.AssetSetError
	(*AssetSetErrorEnum)(nil),            // 1: google.ads.googleads.v12.errors.AssetSetErrorEnum
}
var file_google_ads_googleads_v12_errors_asset_set_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v12_errors_asset_set_error_proto_init() }
func file_google_ads_googleads_v12_errors_asset_set_error_proto_init() {
	if File_google_ads_googleads_v12_errors_asset_set_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_errors_asset_set_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetSetErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v12_errors_asset_set_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v12_errors_asset_set_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_errors_asset_set_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v12_errors_asset_set_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v12_errors_asset_set_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_errors_asset_set_error_proto = out.File
	file_google_ads_googleads_v12_errors_asset_set_error_proto_rawDesc = nil
	file_google_ads_googleads_v12_errors_asset_set_error_proto_goTypes = nil
	file_google_ads_googleads_v12_errors_asset_set_error_proto_depIdxs = nil
}