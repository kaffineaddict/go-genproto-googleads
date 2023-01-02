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
// source: google/ads/googleads/v11/enums/display_upload_product_type.proto

package enums

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

// Enumerates display upload product types.
type DisplayUploadProductTypeEnum_DisplayUploadProductType int32

const (
	// Not specified.
	DisplayUploadProductTypeEnum_UNSPECIFIED DisplayUploadProductTypeEnum_DisplayUploadProductType = 0
	// The value is unknown in this version.
	DisplayUploadProductTypeEnum_UNKNOWN DisplayUploadProductTypeEnum_DisplayUploadProductType = 1
	// HTML5 upload ad. This product type requires the upload_media_bundle
	// field in DisplayUploadAdInfo to be set.
	DisplayUploadProductTypeEnum_HTML5_UPLOAD_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 2
	// Dynamic HTML5 education ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in an education campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_EDUCATION_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 3
	// Dynamic HTML5 flight ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a flight campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_FLIGHT_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 4
	// Dynamic HTML5 hotel and rental ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a hotel campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_HOTEL_RENTAL_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 5
	// Dynamic HTML5 job ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a job campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_JOB_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 6
	// Dynamic HTML5 local ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a local campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_LOCAL_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 7
	// Dynamic HTML5 real estate ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a real estate campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_REAL_ESTATE_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 8
	// Dynamic HTML5 custom ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a custom campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_CUSTOM_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 9
	// Dynamic HTML5 travel ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a travel campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_TRAVEL_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 10
	// Dynamic HTML5 hotel ad. This product type requires the
	// upload_media_bundle field in DisplayUploadAdInfo to be set. Can only be
	// used in a hotel campaign.
	DisplayUploadProductTypeEnum_DYNAMIC_HTML5_HOTEL_AD DisplayUploadProductTypeEnum_DisplayUploadProductType = 11
)

// Enum value maps for DisplayUploadProductTypeEnum_DisplayUploadProductType.
var (
	DisplayUploadProductTypeEnum_DisplayUploadProductType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "HTML5_UPLOAD_AD",
		3:  "DYNAMIC_HTML5_EDUCATION_AD",
		4:  "DYNAMIC_HTML5_FLIGHT_AD",
		5:  "DYNAMIC_HTML5_HOTEL_RENTAL_AD",
		6:  "DYNAMIC_HTML5_JOB_AD",
		7:  "DYNAMIC_HTML5_LOCAL_AD",
		8:  "DYNAMIC_HTML5_REAL_ESTATE_AD",
		9:  "DYNAMIC_HTML5_CUSTOM_AD",
		10: "DYNAMIC_HTML5_TRAVEL_AD",
		11: "DYNAMIC_HTML5_HOTEL_AD",
	}
	DisplayUploadProductTypeEnum_DisplayUploadProductType_value = map[string]int32{
		"UNSPECIFIED":                   0,
		"UNKNOWN":                       1,
		"HTML5_UPLOAD_AD":               2,
		"DYNAMIC_HTML5_EDUCATION_AD":    3,
		"DYNAMIC_HTML5_FLIGHT_AD":       4,
		"DYNAMIC_HTML5_HOTEL_RENTAL_AD": 5,
		"DYNAMIC_HTML5_JOB_AD":          6,
		"DYNAMIC_HTML5_LOCAL_AD":        7,
		"DYNAMIC_HTML5_REAL_ESTATE_AD":  8,
		"DYNAMIC_HTML5_CUSTOM_AD":       9,
		"DYNAMIC_HTML5_TRAVEL_AD":       10,
		"DYNAMIC_HTML5_HOTEL_AD":        11,
	}
)

func (x DisplayUploadProductTypeEnum_DisplayUploadProductType) Enum() *DisplayUploadProductTypeEnum_DisplayUploadProductType {
	p := new(DisplayUploadProductTypeEnum_DisplayUploadProductType)
	*p = x
	return p
}

func (x DisplayUploadProductTypeEnum_DisplayUploadProductType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DisplayUploadProductTypeEnum_DisplayUploadProductType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v11_enums_display_upload_product_type_proto_enumTypes[0].Descriptor()
}

func (DisplayUploadProductTypeEnum_DisplayUploadProductType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v11_enums_display_upload_product_type_proto_enumTypes[0]
}

func (x DisplayUploadProductTypeEnum_DisplayUploadProductType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DisplayUploadProductTypeEnum_DisplayUploadProductType.Descriptor instead.
func (DisplayUploadProductTypeEnum_DisplayUploadProductType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_enums_display_upload_product_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for display upload product types. Product types that have the word
// "DYNAMIC" in them must be associated with a campaign that has a dynamic
// remarketing feed. See https://support.google.com/google-ads/answer/6053288
// for more info about dynamic remarketing. Other product types are regarded
// as "static" and do not have this requirement.
type DisplayUploadProductTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisplayUploadProductTypeEnum) Reset() {
	*x = DisplayUploadProductTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v11_enums_display_upload_product_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisplayUploadProductTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisplayUploadProductTypeEnum) ProtoMessage() {}

func (x *DisplayUploadProductTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v11_enums_display_upload_product_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisplayUploadProductTypeEnum.ProtoReflect.Descriptor instead.
func (*DisplayUploadProductTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_enums_display_upload_product_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v11_enums_display_upload_product_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v11_enums_display_upload_product_type_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x22, 0xfc, 0x02, 0x0a, 0x1c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x22, 0xdb, 0x02, 0x0a, 0x18, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x48, 0x54, 0x4d, 0x4c, 0x35, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x41,
	0x44, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x48,
	0x54, 0x4d, 0x4c, 0x35, 0x5f, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41,
	0x44, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x48,
	0x54, 0x4d, 0x4c, 0x35, 0x5f, 0x46, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x41, 0x44, 0x10, 0x04,
	0x12, 0x21, 0x0a, 0x1d, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x48, 0x54, 0x4d, 0x4c,
	0x35, 0x5f, 0x48, 0x4f, 0x54, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x4e, 0x54, 0x41, 0x4c, 0x5f, 0x41,
	0x44, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x48,
	0x54, 0x4d, 0x4c, 0x35, 0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x41, 0x44, 0x10, 0x06, 0x12, 0x1a, 0x0a,
	0x16, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x35, 0x5f, 0x4c,
	0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x41, 0x44, 0x10, 0x07, 0x12, 0x20, 0x0a, 0x1c, 0x44, 0x59, 0x4e,
	0x41, 0x4d, 0x49, 0x43, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x35, 0x5f, 0x52, 0x45, 0x41, 0x4c, 0x5f,
	0x45, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x44, 0x10, 0x08, 0x12, 0x1b, 0x0a, 0x17, 0x44,
	0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x35, 0x5f, 0x43, 0x55, 0x53,
	0x54, 0x4f, 0x4d, 0x5f, 0x41, 0x44, 0x10, 0x09, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x59, 0x4e, 0x41,
	0x4d, 0x49, 0x43, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x35, 0x5f, 0x54, 0x52, 0x41, 0x56, 0x45, 0x4c,
	0x5f, 0x41, 0x44, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43,
	0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x35, 0x5f, 0x48, 0x4f, 0x54, 0x45, 0x4c, 0x5f, 0x41, 0x44, 0x10,
	0x0b, 0x42, 0xf7, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1d, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x31, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x31,
	0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x31, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v11_enums_display_upload_product_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v11_enums_display_upload_product_type_proto_rawDescData = file_google_ads_googleads_v11_enums_display_upload_product_type_proto_rawDesc
)

func file_google_ads_googleads_v11_enums_display_upload_product_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v11_enums_display_upload_product_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v11_enums_display_upload_product_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v11_enums_display_upload_product_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v11_enums_display_upload_product_type_proto_rawDescData
}

var file_google_ads_googleads_v11_enums_display_upload_product_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v11_enums_display_upload_product_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v11_enums_display_upload_product_type_proto_goTypes = []interface{}{
	(DisplayUploadProductTypeEnum_DisplayUploadProductType)(0), // 0: google.ads.googleads.v11.enums.DisplayUploadProductTypeEnum.DisplayUploadProductType
	(*DisplayUploadProductTypeEnum)(nil),                       // 1: google.ads.googleads.v11.enums.DisplayUploadProductTypeEnum
}
var file_google_ads_googleads_v11_enums_display_upload_product_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v11_enums_display_upload_product_type_proto_init() }
func file_google_ads_googleads_v11_enums_display_upload_product_type_proto_init() {
	if File_google_ads_googleads_v11_enums_display_upload_product_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v11_enums_display_upload_product_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisplayUploadProductTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v11_enums_display_upload_product_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v11_enums_display_upload_product_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v11_enums_display_upload_product_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v11_enums_display_upload_product_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v11_enums_display_upload_product_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v11_enums_display_upload_product_type_proto = out.File
	file_google_ads_googleads_v11_enums_display_upload_product_type_proto_rawDesc = nil
	file_google_ads_googleads_v11_enums_display_upload_product_type_proto_goTypes = nil
	file_google_ads_googleads_v11_enums_display_upload_product_type_proto_depIdxs = nil
}
