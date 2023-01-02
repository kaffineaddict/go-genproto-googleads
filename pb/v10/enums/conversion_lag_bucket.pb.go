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
// source: google/ads/googleads/v10/enums/conversion_lag_bucket.proto

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

// Enum representing the number of days between impression and conversion.
type ConversionLagBucketEnum_ConversionLagBucket int32

const (
	// Not specified.
	ConversionLagBucketEnum_UNSPECIFIED ConversionLagBucketEnum_ConversionLagBucket = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionLagBucketEnum_UNKNOWN ConversionLagBucketEnum_ConversionLagBucket = 1
	// Conversion lag bucket from 0 to 1 day. 0 day is included, 1 day is not.
	ConversionLagBucketEnum_LESS_THAN_ONE_DAY ConversionLagBucketEnum_ConversionLagBucket = 2
	// Conversion lag bucket from 1 to 2 days. 1 day is included, 2 days is not.
	ConversionLagBucketEnum_ONE_TO_TWO_DAYS ConversionLagBucketEnum_ConversionLagBucket = 3
	// Conversion lag bucket from 2 to 3 days. 2 days is included,
	// 3 days is not.
	ConversionLagBucketEnum_TWO_TO_THREE_DAYS ConversionLagBucketEnum_ConversionLagBucket = 4
	// Conversion lag bucket from 3 to 4 days. 3 days is included,
	// 4 days is not.
	ConversionLagBucketEnum_THREE_TO_FOUR_DAYS ConversionLagBucketEnum_ConversionLagBucket = 5
	// Conversion lag bucket from 4 to 5 days. 4 days is included,
	// 5 days is not.
	ConversionLagBucketEnum_FOUR_TO_FIVE_DAYS ConversionLagBucketEnum_ConversionLagBucket = 6
	// Conversion lag bucket from 5 to 6 days. 5 days is included,
	// 6 days is not.
	ConversionLagBucketEnum_FIVE_TO_SIX_DAYS ConversionLagBucketEnum_ConversionLagBucket = 7
	// Conversion lag bucket from 6 to 7 days. 6 days is included,
	// 7 days is not.
	ConversionLagBucketEnum_SIX_TO_SEVEN_DAYS ConversionLagBucketEnum_ConversionLagBucket = 8
	// Conversion lag bucket from 7 to 8 days. 7 days is included,
	// 8 days is not.
	ConversionLagBucketEnum_SEVEN_TO_EIGHT_DAYS ConversionLagBucketEnum_ConversionLagBucket = 9
	// Conversion lag bucket from 8 to 9 days. 8 days is included,
	// 9 days is not.
	ConversionLagBucketEnum_EIGHT_TO_NINE_DAYS ConversionLagBucketEnum_ConversionLagBucket = 10
	// Conversion lag bucket from 9 to 10 days. 9 days is included,
	// 10 days is not.
	ConversionLagBucketEnum_NINE_TO_TEN_DAYS ConversionLagBucketEnum_ConversionLagBucket = 11
	// Conversion lag bucket from 10 to 11 days. 10 days is included,
	// 11 days is not.
	ConversionLagBucketEnum_TEN_TO_ELEVEN_DAYS ConversionLagBucketEnum_ConversionLagBucket = 12
	// Conversion lag bucket from 11 to 12 days. 11 days is included,
	// 12 days is not.
	ConversionLagBucketEnum_ELEVEN_TO_TWELVE_DAYS ConversionLagBucketEnum_ConversionLagBucket = 13
	// Conversion lag bucket from 12 to 13 days. 12 days is included,
	// 13 days is not.
	ConversionLagBucketEnum_TWELVE_TO_THIRTEEN_DAYS ConversionLagBucketEnum_ConversionLagBucket = 14
	// Conversion lag bucket from 13 to 14 days. 13 days is included,
	// 14 days is not.
	ConversionLagBucketEnum_THIRTEEN_TO_FOURTEEN_DAYS ConversionLagBucketEnum_ConversionLagBucket = 15
	// Conversion lag bucket from 14 to 21 days. 14 days is included,
	// 21 days is not.
	ConversionLagBucketEnum_FOURTEEN_TO_TWENTY_ONE_DAYS ConversionLagBucketEnum_ConversionLagBucket = 16
	// Conversion lag bucket from 21 to 30 days. 21 days is included,
	// 30 days is not.
	ConversionLagBucketEnum_TWENTY_ONE_TO_THIRTY_DAYS ConversionLagBucketEnum_ConversionLagBucket = 17
	// Conversion lag bucket from 30 to 45 days. 30 days is included,
	// 45 days is not.
	ConversionLagBucketEnum_THIRTY_TO_FORTY_FIVE_DAYS ConversionLagBucketEnum_ConversionLagBucket = 18
	// Conversion lag bucket from 45 to 60 days. 45 days is included,
	// 60 days is not.
	ConversionLagBucketEnum_FORTY_FIVE_TO_SIXTY_DAYS ConversionLagBucketEnum_ConversionLagBucket = 19
	// Conversion lag bucket from 60 to 90 days. 60 days is included,
	// 90 days is not.
	ConversionLagBucketEnum_SIXTY_TO_NINETY_DAYS ConversionLagBucketEnum_ConversionLagBucket = 20
)

// Enum value maps for ConversionLagBucketEnum_ConversionLagBucket.
var (
	ConversionLagBucketEnum_ConversionLagBucket_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "LESS_THAN_ONE_DAY",
		3:  "ONE_TO_TWO_DAYS",
		4:  "TWO_TO_THREE_DAYS",
		5:  "THREE_TO_FOUR_DAYS",
		6:  "FOUR_TO_FIVE_DAYS",
		7:  "FIVE_TO_SIX_DAYS",
		8:  "SIX_TO_SEVEN_DAYS",
		9:  "SEVEN_TO_EIGHT_DAYS",
		10: "EIGHT_TO_NINE_DAYS",
		11: "NINE_TO_TEN_DAYS",
		12: "TEN_TO_ELEVEN_DAYS",
		13: "ELEVEN_TO_TWELVE_DAYS",
		14: "TWELVE_TO_THIRTEEN_DAYS",
		15: "THIRTEEN_TO_FOURTEEN_DAYS",
		16: "FOURTEEN_TO_TWENTY_ONE_DAYS",
		17: "TWENTY_ONE_TO_THIRTY_DAYS",
		18: "THIRTY_TO_FORTY_FIVE_DAYS",
		19: "FORTY_FIVE_TO_SIXTY_DAYS",
		20: "SIXTY_TO_NINETY_DAYS",
	}
	ConversionLagBucketEnum_ConversionLagBucket_value = map[string]int32{
		"UNSPECIFIED":                 0,
		"UNKNOWN":                     1,
		"LESS_THAN_ONE_DAY":           2,
		"ONE_TO_TWO_DAYS":             3,
		"TWO_TO_THREE_DAYS":           4,
		"THREE_TO_FOUR_DAYS":          5,
		"FOUR_TO_FIVE_DAYS":           6,
		"FIVE_TO_SIX_DAYS":            7,
		"SIX_TO_SEVEN_DAYS":           8,
		"SEVEN_TO_EIGHT_DAYS":         9,
		"EIGHT_TO_NINE_DAYS":          10,
		"NINE_TO_TEN_DAYS":            11,
		"TEN_TO_ELEVEN_DAYS":          12,
		"ELEVEN_TO_TWELVE_DAYS":       13,
		"TWELVE_TO_THIRTEEN_DAYS":     14,
		"THIRTEEN_TO_FOURTEEN_DAYS":   15,
		"FOURTEEN_TO_TWENTY_ONE_DAYS": 16,
		"TWENTY_ONE_TO_THIRTY_DAYS":   17,
		"THIRTY_TO_FORTY_FIVE_DAYS":   18,
		"FORTY_FIVE_TO_SIXTY_DAYS":    19,
		"SIXTY_TO_NINETY_DAYS":        20,
	}
)

func (x ConversionLagBucketEnum_ConversionLagBucket) Enum() *ConversionLagBucketEnum_ConversionLagBucket {
	p := new(ConversionLagBucketEnum_ConversionLagBucket)
	*p = x
	return p
}

func (x ConversionLagBucketEnum_ConversionLagBucket) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversionLagBucketEnum_ConversionLagBucket) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_enumTypes[0].Descriptor()
}

func (ConversionLagBucketEnum_ConversionLagBucket) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_enumTypes[0]
}

func (x ConversionLagBucketEnum_ConversionLagBucket) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversionLagBucketEnum_ConversionLagBucket.Descriptor instead.
func (ConversionLagBucketEnum_ConversionLagBucket) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum representing the number of days between impression and
// conversion.
type ConversionLagBucketEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConversionLagBucketEnum) Reset() {
	*x = ConversionLagBucketEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionLagBucketEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionLagBucketEnum) ProtoMessage() {}

func (x *ConversionLagBucketEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionLagBucketEnum.ProtoReflect.Descriptor instead.
func (*ConversionLagBucketEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_enums_conversion_lag_bucket_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x67, 0x5f,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xbb, 0x04, 0x0a,
	0x17, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x67, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x9f, 0x04, 0x0a, 0x13, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x67, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x15,
	0x0a, 0x11, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x4f, 0x4e, 0x45, 0x5f,
	0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x4e, 0x45, 0x5f, 0x54, 0x4f, 0x5f,
	0x54, 0x57, 0x4f, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x57,
	0x4f, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x45, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10,
	0x04, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x48, 0x52, 0x45, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x46, 0x4f,
	0x55, 0x52, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x4f, 0x55,
	0x52, 0x5f, 0x54, 0x4f, 0x5f, 0x46, 0x49, 0x56, 0x45, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x06,
	0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x49, 0x58, 0x5f,
	0x44, 0x41, 0x59, 0x53, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x49, 0x58, 0x5f, 0x54, 0x4f,
	0x5f, 0x53, 0x45, 0x56, 0x45, 0x4e, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x08, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x45, 0x56, 0x45, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x45, 0x49, 0x47, 0x48, 0x54, 0x5f,
	0x44, 0x41, 0x59, 0x53, 0x10, 0x09, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x49, 0x47, 0x48, 0x54, 0x5f,
	0x54, 0x4f, 0x5f, 0x4e, 0x49, 0x4e, 0x45, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x0a, 0x12, 0x14,
	0x0a, 0x10, 0x4e, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x45, 0x4e, 0x5f, 0x44, 0x41,
	0x59, 0x53, 0x10, 0x0b, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x45, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x45,
	0x4c, 0x45, 0x56, 0x45, 0x4e, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x0c, 0x12, 0x19, 0x0a, 0x15,
	0x45, 0x4c, 0x45, 0x56, 0x45, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x57, 0x45, 0x4c, 0x56, 0x45,
	0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x0d, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x57, 0x45, 0x4c, 0x56,
	0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x54, 0x45, 0x45, 0x4e, 0x5f, 0x44, 0x41,
	0x59, 0x53, 0x10, 0x0e, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x48, 0x49, 0x52, 0x54, 0x45, 0x45, 0x4e,
	0x5f, 0x54, 0x4f, 0x5f, 0x46, 0x4f, 0x55, 0x52, 0x54, 0x45, 0x45, 0x4e, 0x5f, 0x44, 0x41, 0x59,
	0x53, 0x10, 0x0f, 0x12, 0x1f, 0x0a, 0x1b, 0x46, 0x4f, 0x55, 0x52, 0x54, 0x45, 0x45, 0x4e, 0x5f,
	0x54, 0x4f, 0x5f, 0x54, 0x57, 0x45, 0x4e, 0x54, 0x59, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x44, 0x41,
	0x59, 0x53, 0x10, 0x10, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x57, 0x45, 0x4e, 0x54, 0x59, 0x5f, 0x4f,
	0x4e, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x54, 0x59, 0x5f, 0x44, 0x41, 0x59,
	0x53, 0x10, 0x11, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x48, 0x49, 0x52, 0x54, 0x59, 0x5f, 0x54, 0x4f,
	0x5f, 0x46, 0x4f, 0x52, 0x54, 0x59, 0x5f, 0x46, 0x49, 0x56, 0x45, 0x5f, 0x44, 0x41, 0x59, 0x53,
	0x10, 0x12, 0x12, 0x1c, 0x0a, 0x18, 0x46, 0x4f, 0x52, 0x54, 0x59, 0x5f, 0x46, 0x49, 0x56, 0x45,
	0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x49, 0x58, 0x54, 0x59, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x13,
	0x12, 0x18, 0x0a, 0x14, 0x53, 0x49, 0x58, 0x54, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x49, 0x4e,
	0x45, 0x54, 0x59, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x14, 0x42, 0xf2, 0x01, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x42, 0x18, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x67,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e,
	0x56, 0x31, 0x30, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_rawDescData = file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_rawDesc
)

func file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_rawDescData
}

var file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_goTypes = []interface{}{
	(ConversionLagBucketEnum_ConversionLagBucket)(0), // 0: google.ads.googleads.v10.enums.ConversionLagBucketEnum.ConversionLagBucket
	(*ConversionLagBucketEnum)(nil),                  // 1: google.ads.googleads.v10.enums.ConversionLagBucketEnum
}
var file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_init() }
func file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_init() {
	if File_google_ads_googleads_v10_enums_conversion_lag_bucket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionLagBucketEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_enums_conversion_lag_bucket_proto = out.File
	file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_rawDesc = nil
	file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_goTypes = nil
	file_google_ads_googleads_v10_enums_conversion_lag_bucket_proto_depIdxs = nil
}
