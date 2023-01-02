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
// source: google/ads/googleads/v10/resources/mobile_device_constant.proto

package resources

import (
	enums "github.com/dictav/go-genproto-googleads/pb/v10/enums"
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

// A mobile device constant.
type MobileDeviceConstant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the mobile device constant.
	// Mobile device constant resource names have the form:
	//
	// `mobileDeviceConstants/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the mobile device constant.
	Id *int64 `protobuf:"varint,7,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Output only. The name of the mobile device.
	Name *string `protobuf:"bytes,8,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Output only. The manufacturer of the mobile device.
	ManufacturerName *string `protobuf:"bytes,9,opt,name=manufacturer_name,json=manufacturerName,proto3,oneof" json:"manufacturer_name,omitempty"`
	// Output only. The operating system of the mobile device.
	OperatingSystemName *string `protobuf:"bytes,10,opt,name=operating_system_name,json=operatingSystemName,proto3,oneof" json:"operating_system_name,omitempty"`
	// Output only. The type of mobile device.
	Type enums.MobileDeviceTypeEnum_MobileDeviceType `protobuf:"varint,6,opt,name=type,proto3,enum=google.ads.googleads.v10.enums.MobileDeviceTypeEnum_MobileDeviceType" json:"type,omitempty"`
}

func (x *MobileDeviceConstant) Reset() {
	*x = MobileDeviceConstant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_resources_mobile_device_constant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MobileDeviceConstant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MobileDeviceConstant) ProtoMessage() {}

func (x *MobileDeviceConstant) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_resources_mobile_device_constant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MobileDeviceConstant.ProtoReflect.Descriptor instead.
func (*MobileDeviceConstant) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_resources_mobile_device_constant_proto_rawDescGZIP(), []int{0}
}

func (x *MobileDeviceConstant) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MobileDeviceConstant) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *MobileDeviceConstant) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *MobileDeviceConstant) GetManufacturerName() string {
	if x != nil && x.ManufacturerName != nil {
		return *x.ManufacturerName
	}
	return ""
}

func (x *MobileDeviceConstant) GetOperatingSystemName() string {
	if x != nil && x.OperatingSystemName != nil {
		return *x.OperatingSystemName
	}
	return ""
}

func (x *MobileDeviceConstant) GetType() enums.MobileDeviceTypeEnum_MobileDeviceType {
	if x != nil {
		return x.Type
	}
	return enums.MobileDeviceTypeEnum_MobileDeviceType(0)
}

var File_google_ads_googleads_v10_resources_mobile_device_constant_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_resources_mobile_device_constant_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x04, 0x0a, 0x14, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x12, 0x5a, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0xe0, 0x41, 0x03, 0xfa,
	0x41, 0x2f, 0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x11, 0x6d, 0x61, 0x6e, 0x75, 0x66,
	0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x10, 0x6d, 0x61, 0x6e, 0x75, 0x66,
	0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3c,
	0x0a, 0x15, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x48, 0x03, 0x52, 0x13, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x5e, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x58, 0xea, 0x41,
	0x55, 0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x12, 0x24, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x6d, 0x61, 0x6e, 0x75, 0x66,
	0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x18, 0x0a, 0x16,
	0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x8b, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x42, 0x19, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31,
	0x30, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_resources_mobile_device_constant_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_resources_mobile_device_constant_proto_rawDescData = file_google_ads_googleads_v10_resources_mobile_device_constant_proto_rawDesc
)

func file_google_ads_googleads_v10_resources_mobile_device_constant_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_resources_mobile_device_constant_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_resources_mobile_device_constant_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_resources_mobile_device_constant_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_resources_mobile_device_constant_proto_rawDescData
}

var file_google_ads_googleads_v10_resources_mobile_device_constant_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_resources_mobile_device_constant_proto_goTypes = []interface{}{
	(*MobileDeviceConstant)(nil),                     // 0: google.ads.googleads.v10.resources.MobileDeviceConstant
	(enums.MobileDeviceTypeEnum_MobileDeviceType)(0), // 1: google.ads.googleads.v10.enums.MobileDeviceTypeEnum.MobileDeviceType
}
var file_google_ads_googleads_v10_resources_mobile_device_constant_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v10.resources.MobileDeviceConstant.type:type_name -> google.ads.googleads.v10.enums.MobileDeviceTypeEnum.MobileDeviceType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_resources_mobile_device_constant_proto_init() }
func file_google_ads_googleads_v10_resources_mobile_device_constant_proto_init() {
	if File_google_ads_googleads_v10_resources_mobile_device_constant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_resources_mobile_device_constant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MobileDeviceConstant); i {
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
	file_google_ads_googleads_v10_resources_mobile_device_constant_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v10_resources_mobile_device_constant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_resources_mobile_device_constant_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_resources_mobile_device_constant_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v10_resources_mobile_device_constant_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_resources_mobile_device_constant_proto = out.File
	file_google_ads_googleads_v10_resources_mobile_device_constant_proto_rawDesc = nil
	file_google_ads_googleads_v10_resources_mobile_device_constant_proto_goTypes = nil
	file_google_ads_googleads_v10_resources_mobile_device_constant_proto_depIdxs = nil
}
