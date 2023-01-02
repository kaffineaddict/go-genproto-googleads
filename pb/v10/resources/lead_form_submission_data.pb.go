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
// source: google/ads/googleads/v10/resources/lead_form_submission_data.proto

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

// Data from lead form submissions.
type LeadFormSubmissionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the lead form submission data.
	// Lead form submission data resource names have the form:
	//
	// `customers/{customer_id}/leadFormSubmissionData/{lead_form_submission_data_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. ID of this lead form submission.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. Asset associated with the submitted lead form.
	Asset string `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	// Output only. Campaign associated with the submitted lead form.
	Campaign string `protobuf:"bytes,4,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Output only. Submission data associated with a lead form.
	LeadFormSubmissionFields []*LeadFormSubmissionField `protobuf:"bytes,5,rep,name=lead_form_submission_fields,json=leadFormSubmissionFields,proto3" json:"lead_form_submission_fields,omitempty"`
	// Output only. AdGroup associated with the submitted lead form.
	AdGroup string `protobuf:"bytes,6,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Output only. AdGroupAd associated with the submitted lead form.
	AdGroupAd string `protobuf:"bytes,7,opt,name=ad_group_ad,json=adGroupAd,proto3" json:"ad_group_ad,omitempty"`
	// Output only. Google Click Id associated with the submissed lead form.
	Gclid string `protobuf:"bytes,8,opt,name=gclid,proto3" json:"gclid,omitempty"`
	// Output only. The date and time at which the lead form was submitted. The format is
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", e.g. "2019-01-01 12:32:45-08:00".
	SubmissionDateTime string `protobuf:"bytes,9,opt,name=submission_date_time,json=submissionDateTime,proto3" json:"submission_date_time,omitempty"`
}

func (x *LeadFormSubmissionData) Reset() {
	*x = LeadFormSubmissionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeadFormSubmissionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeadFormSubmissionData) ProtoMessage() {}

func (x *LeadFormSubmissionData) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeadFormSubmissionData.ProtoReflect.Descriptor instead.
func (*LeadFormSubmissionData) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_rawDescGZIP(), []int{0}
}

func (x *LeadFormSubmissionData) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *LeadFormSubmissionData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LeadFormSubmissionData) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *LeadFormSubmissionData) GetCampaign() string {
	if x != nil {
		return x.Campaign
	}
	return ""
}

func (x *LeadFormSubmissionData) GetLeadFormSubmissionFields() []*LeadFormSubmissionField {
	if x != nil {
		return x.LeadFormSubmissionFields
	}
	return nil
}

func (x *LeadFormSubmissionData) GetAdGroup() string {
	if x != nil {
		return x.AdGroup
	}
	return ""
}

func (x *LeadFormSubmissionData) GetAdGroupAd() string {
	if x != nil {
		return x.AdGroupAd
	}
	return ""
}

func (x *LeadFormSubmissionData) GetGclid() string {
	if x != nil {
		return x.Gclid
	}
	return ""
}

func (x *LeadFormSubmissionData) GetSubmissionDateTime() string {
	if x != nil {
		return x.SubmissionDateTime
	}
	return ""
}

// Fields in the submitted lead form.
type LeadFormSubmissionField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Field type for lead form fields.
	FieldType enums.LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType `protobuf:"varint,1,opt,name=field_type,json=fieldType,proto3,enum=google.ads.googleads.v10.enums.LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType" json:"field_type,omitempty"`
	// Output only. Field value for lead form fields.
	FieldValue string `protobuf:"bytes,2,opt,name=field_value,json=fieldValue,proto3" json:"field_value,omitempty"`
}

func (x *LeadFormSubmissionField) Reset() {
	*x = LeadFormSubmissionField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeadFormSubmissionField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeadFormSubmissionField) ProtoMessage() {}

func (x *LeadFormSubmissionField) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeadFormSubmissionField.ProtoReflect.Descriptor instead.
func (*LeadFormSubmissionField) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_rawDescGZIP(), []int{1}
}

func (x *LeadFormSubmissionField) GetFieldType() enums.LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType {
	if x != nil {
		return x.FieldType
	}
	return enums.LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType(0)
}

func (x *LeadFormSubmissionField) GetFieldValue() string {
	if x != nil {
		return x.FieldValue
	}
	return ""
}

var File_google_ads_googleads_v10_resources_lead_form_submission_data_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x66, 0x6f,
	0x72, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x05, 0x0a, 0x16, 0x4c,
	0x65, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x5c, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x37, 0xe0, 0x41,
	0x03, 0xfa, 0x41, 0x31, 0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c,
	0x65, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x20, 0x0a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52,
	0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x45, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x23,
	0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x7f, 0x0a,
	0x1b, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x18, 0x6c, 0x65, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x43,
	0x0a, 0x08, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x28, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x07, 0x61, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x4a, 0x0a, 0x0b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x24,
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x64, 0x52, 0x09, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x12,
	0x19, 0x0a, 0x05, 0x67, 0x63, 0x6c, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x05, 0x67, 0x63, 0x6c, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x14, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x3a, 0x84, 0x01, 0xea, 0x41, 0x80, 0x01, 0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4c, 0x65, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4d, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x2f, 0x7b, 0x6c, 0x65, 0x61, 0x64, 0x5f,
	0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x22, 0xbe, 0x01, 0x0a, 0x17, 0x4c, 0x65, 0x61,
	0x64, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x7d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x59, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x46, 0x6f,
	0x72, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x46, 0x6f, 0x72,
	0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x8d, 0x02, 0x0a, 0x26, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x42, 0x1b, 0x4c, 0x65, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x30, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_rawDescData = file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_rawDesc
)

func file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_rawDescData
}

var file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_goTypes = []interface{}{
	(*LeadFormSubmissionData)(nil),                                       // 0: google.ads.googleads.v10.resources.LeadFormSubmissionData
	(*LeadFormSubmissionField)(nil),                                      // 1: google.ads.googleads.v10.resources.LeadFormSubmissionField
	(enums.LeadFormFieldUserInputTypeEnum_LeadFormFieldUserInputType)(0), // 2: google.ads.googleads.v10.enums.LeadFormFieldUserInputTypeEnum.LeadFormFieldUserInputType
}
var file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v10.resources.LeadFormSubmissionData.lead_form_submission_fields:type_name -> google.ads.googleads.v10.resources.LeadFormSubmissionField
	2, // 1: google.ads.googleads.v10.resources.LeadFormSubmissionField.field_type:type_name -> google.ads.googleads.v10.enums.LeadFormFieldUserInputTypeEnum.LeadFormFieldUserInputType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_init() }
func file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_init() {
	if File_google_ads_googleads_v10_resources_lead_form_submission_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeadFormSubmissionData); i {
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
		file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeadFormSubmissionField); i {
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
			RawDescriptor: file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_resources_lead_form_submission_data_proto = out.File
	file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_rawDesc = nil
	file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_goTypes = nil
	file_google_ads_googleads_v10_resources_lead_form_submission_data_proto_depIdxs = nil
}
