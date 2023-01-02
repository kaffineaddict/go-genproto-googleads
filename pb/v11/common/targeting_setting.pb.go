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
// source: google/ads/googleads/v11/common/targeting_setting.proto

package common

import (
	enums "github.com/dictav/go-genproto-googleads/pb/v11/enums"
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

// The operator.
type TargetRestrictionOperation_Operator int32

const (
	// Unspecified.
	TargetRestrictionOperation_UNSPECIFIED TargetRestrictionOperation_Operator = 0
	// Used for return value only. Represents value unknown in this version.
	TargetRestrictionOperation_UNKNOWN TargetRestrictionOperation_Operator = 1
	// Add the restriction to the existing restrictions.
	TargetRestrictionOperation_ADD TargetRestrictionOperation_Operator = 2
	// Remove the restriction from the existing restrictions.
	TargetRestrictionOperation_REMOVE TargetRestrictionOperation_Operator = 3
)

// Enum value maps for TargetRestrictionOperation_Operator.
var (
	TargetRestrictionOperation_Operator_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "ADD",
		3: "REMOVE",
	}
	TargetRestrictionOperation_Operator_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"ADD":         2,
		"REMOVE":      3,
	}
)

func (x TargetRestrictionOperation_Operator) Enum() *TargetRestrictionOperation_Operator {
	p := new(TargetRestrictionOperation_Operator)
	*p = x
	return p
}

func (x TargetRestrictionOperation_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TargetRestrictionOperation_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v11_common_targeting_setting_proto_enumTypes[0].Descriptor()
}

func (TargetRestrictionOperation_Operator) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v11_common_targeting_setting_proto_enumTypes[0]
}

func (x TargetRestrictionOperation_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TargetRestrictionOperation_Operator.Descriptor instead.
func (TargetRestrictionOperation_Operator) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_common_targeting_setting_proto_rawDescGZIP(), []int{2, 0}
}

// Settings for the targeting-related features, at the campaign and ad group
// levels. For more details about the targeting setting, visit
// https://support.google.com/google-ads/answer/7365594
type TargetingSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The per-targeting-dimension setting to restrict the reach of your campaign
	// or ad group.
	TargetRestrictions []*TargetRestriction `protobuf:"bytes,1,rep,name=target_restrictions,json=targetRestrictions,proto3" json:"target_restrictions,omitempty"`
	// The list of operations changing the target restrictions.
	//
	// Adding a target restriction with a targeting dimension that already exists
	// causes the existing target restriction to be replaced with the new value.
	TargetRestrictionOperations []*TargetRestrictionOperation `protobuf:"bytes,2,rep,name=target_restriction_operations,json=targetRestrictionOperations,proto3" json:"target_restriction_operations,omitempty"`
}

func (x *TargetingSetting) Reset() {
	*x = TargetingSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v11_common_targeting_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetingSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetingSetting) ProtoMessage() {}

func (x *TargetingSetting) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v11_common_targeting_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetingSetting.ProtoReflect.Descriptor instead.
func (*TargetingSetting) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_common_targeting_setting_proto_rawDescGZIP(), []int{0}
}

func (x *TargetingSetting) GetTargetRestrictions() []*TargetRestriction {
	if x != nil {
		return x.TargetRestrictions
	}
	return nil
}

func (x *TargetingSetting) GetTargetRestrictionOperations() []*TargetRestrictionOperation {
	if x != nil {
		return x.TargetRestrictionOperations
	}
	return nil
}

// The list of per-targeting-dimension targeting settings.
type TargetRestriction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The targeting dimension that these settings apply to.
	TargetingDimension enums.TargetingDimensionEnum_TargetingDimension `protobuf:"varint,1,opt,name=targeting_dimension,json=targetingDimension,proto3,enum=google.ads.googleads.v11.enums.TargetingDimensionEnum_TargetingDimension" json:"targeting_dimension,omitempty"`
	// Indicates whether to restrict your ads to show only for the criteria you
	// have selected for this targeting_dimension, or to target all values for
	// this targeting_dimension and show ads based on your targeting in other
	// TargetingDimensions. A value of `true` means that these criteria will only
	// apply bid modifiers, and not affect targeting. A value of `false` means
	// that these criteria will restrict targeting as well as applying bid
	// modifiers.
	BidOnly *bool `protobuf:"varint,3,opt,name=bid_only,json=bidOnly,proto3,oneof" json:"bid_only,omitempty"`
}

func (x *TargetRestriction) Reset() {
	*x = TargetRestriction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v11_common_targeting_setting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetRestriction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetRestriction) ProtoMessage() {}

func (x *TargetRestriction) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v11_common_targeting_setting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetRestriction.ProtoReflect.Descriptor instead.
func (*TargetRestriction) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_common_targeting_setting_proto_rawDescGZIP(), []int{1}
}

func (x *TargetRestriction) GetTargetingDimension() enums.TargetingDimensionEnum_TargetingDimension {
	if x != nil {
		return x.TargetingDimension
	}
	return enums.TargetingDimensionEnum_TargetingDimension(0)
}

func (x *TargetRestriction) GetBidOnly() bool {
	if x != nil && x.BidOnly != nil {
		return *x.BidOnly
	}
	return false
}

// Operation to be performed on a target restriction list in a mutate.
type TargetRestrictionOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of list operation to perform.
	Operator TargetRestrictionOperation_Operator `protobuf:"varint,1,opt,name=operator,proto3,enum=google.ads.googleads.v11.common.TargetRestrictionOperation_Operator" json:"operator,omitempty"`
	// The target restriction being added to or removed from the list.
	Value *TargetRestriction `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TargetRestrictionOperation) Reset() {
	*x = TargetRestrictionOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v11_common_targeting_setting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetRestrictionOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetRestrictionOperation) ProtoMessage() {}

func (x *TargetRestrictionOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v11_common_targeting_setting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetRestrictionOperation.ProtoReflect.Descriptor instead.
func (*TargetRestrictionOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_common_targeting_setting_proto_rawDescGZIP(), []int{2}
}

func (x *TargetRestrictionOperation) GetOperator() TargetRestrictionOperation_Operator {
	if x != nil {
		return x.Operator
	}
	return TargetRestrictionOperation_UNSPECIFIED
}

func (x *TargetRestrictionOperation) GetValue() *TargetRestriction {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_google_ads_googleads_v11_common_targeting_setting_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v11_common_targeting_setting_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x10, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x63, 0x0a, 0x13, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x7f,
	0x0a, 0x1d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x1b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0xbc, 0x01, 0x0a, 0x11, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7a, 0x0a, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x49, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x6d,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x08, 0x62, 0x69, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x62, 0x69, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x88, 0x01,
	0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x22, 0x87,
	0x02, 0x0a, 0x1a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x60, 0x0a,
	0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x48, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3d, 0x0a, 0x08, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x03, 0x42, 0xf5, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x42, 0x15, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v11_common_targeting_setting_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v11_common_targeting_setting_proto_rawDescData = file_google_ads_googleads_v11_common_targeting_setting_proto_rawDesc
)

func file_google_ads_googleads_v11_common_targeting_setting_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v11_common_targeting_setting_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v11_common_targeting_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v11_common_targeting_setting_proto_rawDescData)
	})
	return file_google_ads_googleads_v11_common_targeting_setting_proto_rawDescData
}

var file_google_ads_googleads_v11_common_targeting_setting_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v11_common_targeting_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_ads_googleads_v11_common_targeting_setting_proto_goTypes = []interface{}{
	(TargetRestrictionOperation_Operator)(0),             // 0: google.ads.googleads.v11.common.TargetRestrictionOperation.Operator
	(*TargetingSetting)(nil),                             // 1: google.ads.googleads.v11.common.TargetingSetting
	(*TargetRestriction)(nil),                            // 2: google.ads.googleads.v11.common.TargetRestriction
	(*TargetRestrictionOperation)(nil),                   // 3: google.ads.googleads.v11.common.TargetRestrictionOperation
	(enums.TargetingDimensionEnum_TargetingDimension)(0), // 4: google.ads.googleads.v11.enums.TargetingDimensionEnum.TargetingDimension
}
var file_google_ads_googleads_v11_common_targeting_setting_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v11.common.TargetingSetting.target_restrictions:type_name -> google.ads.googleads.v11.common.TargetRestriction
	3, // 1: google.ads.googleads.v11.common.TargetingSetting.target_restriction_operations:type_name -> google.ads.googleads.v11.common.TargetRestrictionOperation
	4, // 2: google.ads.googleads.v11.common.TargetRestriction.targeting_dimension:type_name -> google.ads.googleads.v11.enums.TargetingDimensionEnum.TargetingDimension
	0, // 3: google.ads.googleads.v11.common.TargetRestrictionOperation.operator:type_name -> google.ads.googleads.v11.common.TargetRestrictionOperation.Operator
	2, // 4: google.ads.googleads.v11.common.TargetRestrictionOperation.value:type_name -> google.ads.googleads.v11.common.TargetRestriction
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v11_common_targeting_setting_proto_init() }
func file_google_ads_googleads_v11_common_targeting_setting_proto_init() {
	if File_google_ads_googleads_v11_common_targeting_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v11_common_targeting_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetingSetting); i {
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
		file_google_ads_googleads_v11_common_targeting_setting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetRestriction); i {
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
		file_google_ads_googleads_v11_common_targeting_setting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetRestrictionOperation); i {
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
	file_google_ads_googleads_v11_common_targeting_setting_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v11_common_targeting_setting_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v11_common_targeting_setting_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v11_common_targeting_setting_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v11_common_targeting_setting_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v11_common_targeting_setting_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v11_common_targeting_setting_proto = out.File
	file_google_ads_googleads_v11_common_targeting_setting_proto_rawDesc = nil
	file_google_ads_googleads_v11_common_targeting_setting_proto_goTypes = nil
	file_google_ads_googleads_v11_common_targeting_setting_proto_depIdxs = nil
}
