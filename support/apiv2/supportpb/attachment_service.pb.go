// Copyright 2024 Google LLC
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
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: google/cloud/support/v2/attachment_service.proto

package supportpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message for the ListAttachments endpoint.
type ListAttachmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of Case object for which attachments should be
	// listed.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of attachments fetched with each request. If not
	// provided, the default is 10. The maximum page size that will be returned is
	// 100.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A token identifying the page of results to return. If unspecified, the
	// first page is retrieved.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListAttachmentsRequest) Reset() {
	*x = ListAttachmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_support_v2_attachment_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAttachmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttachmentsRequest) ProtoMessage() {}

func (x *ListAttachmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_support_v2_attachment_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAttachmentsRequest.ProtoReflect.Descriptor instead.
func (*ListAttachmentsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_support_v2_attachment_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListAttachmentsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListAttachmentsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAttachmentsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The response message for the ListAttachments endpoint.
type ListAttachmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of attachments associated with the given case.
	Attachments []*Attachment `protobuf:"bytes,1,rep,name=attachments,proto3" json:"attachments,omitempty"`
	// A token to retrieve the next page of results. This should be set in the
	// `page_token` field of subsequent `cases.attachments.list` requests. If
	// unspecified, there are no more results to retrieve.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListAttachmentsResponse) Reset() {
	*x = ListAttachmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_support_v2_attachment_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAttachmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAttachmentsResponse) ProtoMessage() {}

func (x *ListAttachmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_support_v2_attachment_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAttachmentsResponse.ProtoReflect.Descriptor instead.
func (*ListAttachmentsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_support_v2_attachment_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListAttachmentsResponse) GetAttachments() []*Attachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *ListAttachmentsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_google_cloud_support_v2_attachment_service_proto protoreflect.FileDescriptor

var file_google_cloud_support_v2_attachment_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x28, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x73, 0x65, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x88, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a,
	0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xd1, 0x02, 0x0a,
	0x15, 0x43, 0x61, 0x73, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xe6, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x70, 0xda,
	0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x61, 0x5a, 0x32,
	0x12, 0x30, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x2b, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x2a, 0x7d, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a,
	0x4f, 0xca, 0x41, 0x1b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2,
	0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x42, 0xbf, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32,
	0x42, 0x16, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0x3b, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x70,
	0x62, 0xaa, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x17, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x3a,
	0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_support_v2_attachment_service_proto_rawDescOnce sync.Once
	file_google_cloud_support_v2_attachment_service_proto_rawDescData = file_google_cloud_support_v2_attachment_service_proto_rawDesc
)

func file_google_cloud_support_v2_attachment_service_proto_rawDescGZIP() []byte {
	file_google_cloud_support_v2_attachment_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_support_v2_attachment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_support_v2_attachment_service_proto_rawDescData)
	})
	return file_google_cloud_support_v2_attachment_service_proto_rawDescData
}

var file_google_cloud_support_v2_attachment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_support_v2_attachment_service_proto_goTypes = []interface{}{
	(*ListAttachmentsRequest)(nil),  // 0: google.cloud.support.v2.ListAttachmentsRequest
	(*ListAttachmentsResponse)(nil), // 1: google.cloud.support.v2.ListAttachmentsResponse
	(*Attachment)(nil),              // 2: google.cloud.support.v2.Attachment
}
var file_google_cloud_support_v2_attachment_service_proto_depIdxs = []int32{
	2, // 0: google.cloud.support.v2.ListAttachmentsResponse.attachments:type_name -> google.cloud.support.v2.Attachment
	0, // 1: google.cloud.support.v2.CaseAttachmentService.ListAttachments:input_type -> google.cloud.support.v2.ListAttachmentsRequest
	1, // 2: google.cloud.support.v2.CaseAttachmentService.ListAttachments:output_type -> google.cloud.support.v2.ListAttachmentsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_support_v2_attachment_service_proto_init() }
func file_google_cloud_support_v2_attachment_service_proto_init() {
	if File_google_cloud_support_v2_attachment_service_proto != nil {
		return
	}
	file_google_cloud_support_v2_attachment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_support_v2_attachment_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAttachmentsRequest); i {
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
		file_google_cloud_support_v2_attachment_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAttachmentsResponse); i {
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
			RawDescriptor: file_google_cloud_support_v2_attachment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_support_v2_attachment_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_support_v2_attachment_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_support_v2_attachment_service_proto_msgTypes,
	}.Build()
	File_google_cloud_support_v2_attachment_service_proto = out.File
	file_google_cloud_support_v2_attachment_service_proto_rawDesc = nil
	file_google_cloud_support_v2_attachment_service_proto_goTypes = nil
	file_google_cloud_support_v2_attachment_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CaseAttachmentServiceClient is the client API for CaseAttachmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CaseAttachmentServiceClient interface {
	// Retrieve all attachments associated with a support case.
	ListAttachments(ctx context.Context, in *ListAttachmentsRequest, opts ...grpc.CallOption) (*ListAttachmentsResponse, error)
}

type caseAttachmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCaseAttachmentServiceClient(cc grpc.ClientConnInterface) CaseAttachmentServiceClient {
	return &caseAttachmentServiceClient{cc}
}

func (c *caseAttachmentServiceClient) ListAttachments(ctx context.Context, in *ListAttachmentsRequest, opts ...grpc.CallOption) (*ListAttachmentsResponse, error) {
	out := new(ListAttachmentsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.support.v2.CaseAttachmentService/ListAttachments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaseAttachmentServiceServer is the server API for CaseAttachmentService service.
type CaseAttachmentServiceServer interface {
	// Retrieve all attachments associated with a support case.
	ListAttachments(context.Context, *ListAttachmentsRequest) (*ListAttachmentsResponse, error)
}

// UnimplementedCaseAttachmentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCaseAttachmentServiceServer struct {
}

func (*UnimplementedCaseAttachmentServiceServer) ListAttachments(context.Context, *ListAttachmentsRequest) (*ListAttachmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAttachments not implemented")
}

func RegisterCaseAttachmentServiceServer(s *grpc.Server, srv CaseAttachmentServiceServer) {
	s.RegisterService(&_CaseAttachmentService_serviceDesc, srv)
}

func _CaseAttachmentService_ListAttachments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAttachmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseAttachmentServiceServer).ListAttachments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.support.v2.CaseAttachmentService/ListAttachments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseAttachmentServiceServer).ListAttachments(ctx, req.(*ListAttachmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CaseAttachmentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.support.v2.CaseAttachmentService",
	HandlerType: (*CaseAttachmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAttachments",
			Handler:    _CaseAttachmentService_ListAttachments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/support/v2/attachment_service.proto",
}
