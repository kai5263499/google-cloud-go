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
// source: google/cloud/video/stitcher/v1/events.proto

package stitcherpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Describes the event that occurred.
type Event_EventType int32

const (
	// The event type is unspecified.
	Event_EVENT_TYPE_UNSPECIFIED Event_EventType = 0
	// First frame of creative ad viewed.
	Event_CREATIVE_VIEW Event_EventType = 1
	// Creative ad started.
	Event_START Event_EventType = 2
	// Start of an ad break.
	Event_BREAK_START Event_EventType = 3
	// End of an ad break.
	Event_BREAK_END Event_EventType = 4
	// Impression.
	Event_IMPRESSION Event_EventType = 5
	// First quartile progress.
	Event_FIRST_QUARTILE Event_EventType = 6
	// Midpoint progress.
	Event_MIDPOINT Event_EventType = 7
	// Third quartile progress.
	Event_THIRD_QUARTILE Event_EventType = 8
	// Ad progress completed.
	Event_COMPLETE Event_EventType = 9
	// Specific progress event with an offset.
	Event_PROGRESS Event_EventType = 10
	// Player muted.
	Event_MUTE Event_EventType = 11
	// Player unmuted.
	Event_UNMUTE Event_EventType = 12
	// Player paused.
	Event_PAUSE Event_EventType = 13
	// Click event.
	Event_CLICK Event_EventType = 14
	// Click-through event.
	Event_CLICK_THROUGH Event_EventType = 15
	// Player rewinding.
	Event_REWIND Event_EventType = 16
	// Player resumed.
	Event_RESUME Event_EventType = 17
	// Error event.
	Event_ERROR Event_EventType = 18
	// Ad expanded to a larger size.
	Event_EXPAND Event_EventType = 21
	// Ad collapsed to a smaller size.
	Event_COLLAPSE Event_EventType = 22
	// Non-linear ad closed.
	Event_CLOSE Event_EventType = 24
	// Linear ad closed.
	Event_CLOSE_LINEAR Event_EventType = 25
	// Ad skipped.
	Event_SKIP Event_EventType = 26
	// Accept invitation event.
	Event_ACCEPT_INVITATION Event_EventType = 27
)

// Enum value maps for Event_EventType.
var (
	Event_EventType_name = map[int32]string{
		0:  "EVENT_TYPE_UNSPECIFIED",
		1:  "CREATIVE_VIEW",
		2:  "START",
		3:  "BREAK_START",
		4:  "BREAK_END",
		5:  "IMPRESSION",
		6:  "FIRST_QUARTILE",
		7:  "MIDPOINT",
		8:  "THIRD_QUARTILE",
		9:  "COMPLETE",
		10: "PROGRESS",
		11: "MUTE",
		12: "UNMUTE",
		13: "PAUSE",
		14: "CLICK",
		15: "CLICK_THROUGH",
		16: "REWIND",
		17: "RESUME",
		18: "ERROR",
		21: "EXPAND",
		22: "COLLAPSE",
		24: "CLOSE",
		25: "CLOSE_LINEAR",
		26: "SKIP",
		27: "ACCEPT_INVITATION",
	}
	Event_EventType_value = map[string]int32{
		"EVENT_TYPE_UNSPECIFIED": 0,
		"CREATIVE_VIEW":          1,
		"START":                  2,
		"BREAK_START":            3,
		"BREAK_END":              4,
		"IMPRESSION":             5,
		"FIRST_QUARTILE":         6,
		"MIDPOINT":               7,
		"THIRD_QUARTILE":         8,
		"COMPLETE":               9,
		"PROGRESS":               10,
		"MUTE":                   11,
		"UNMUTE":                 12,
		"PAUSE":                  13,
		"CLICK":                  14,
		"CLICK_THROUGH":          15,
		"REWIND":                 16,
		"RESUME":                 17,
		"ERROR":                  18,
		"EXPAND":                 21,
		"COLLAPSE":               22,
		"CLOSE":                  24,
		"CLOSE_LINEAR":           25,
		"SKIP":                   26,
		"ACCEPT_INVITATION":      27,
	}
)

func (x Event_EventType) Enum() *Event_EventType {
	p := new(Event_EventType)
	*p = x
	return p
}

func (x Event_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_video_stitcher_v1_events_proto_enumTypes[0].Descriptor()
}

func (Event_EventType) Type() protoreflect.EnumType {
	return &file_google_cloud_video_stitcher_v1_events_proto_enumTypes[0]
}

func (x Event_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_EventType.Descriptor instead.
func (Event_EventType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_events_proto_rawDescGZIP(), []int{0, 0}
}

// Describes an event and a trigger URI.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Describes the event that occurred.
	Type Event_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.video.stitcher.v1.Event_EventType" json:"type,omitempty"`
	// The URI to trigger for this event.
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// The ID of the event.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The offset in seconds if the event type is `PROGRESS`.
	Offset *durationpb.Duration `protobuf:"bytes,4,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetType() Event_EventType {
	if x != nil {
		return x.Type
	}
	return Event_EVENT_TYPE_UNSPECIFIED
}

func (x *Event) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetOffset() *durationpb.Duration {
	if x != nil {
		return x.Offset
	}
	return nil
}

// Indicates a time in which a list of events should be triggered
// during media playback.
type ProgressEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The time when the following tracking events occurs. The time is in
	// seconds relative to the start of the VOD asset.
	TimeOffset *durationpb.Duration `protobuf:"bytes,1,opt,name=time_offset,json=timeOffset,proto3" json:"time_offset,omitempty"`
	// The list of progress tracking events for the ad break. These can be of
	// the following IAB types: `BREAK_START`, `BREAK_END`, `IMPRESSION`,
	// `CREATIVE_VIEW`, `START`, `FIRST_QUARTILE`, `MIDPOINT`, `THIRD_QUARTILE`,
	// `COMPLETE`, `PROGRESS`.
	Events []*Event `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ProgressEvent) Reset() {
	*x = ProgressEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressEvent) ProtoMessage() {}

func (x *ProgressEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressEvent.ProtoReflect.Descriptor instead.
func (*ProgressEvent) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_events_proto_rawDescGZIP(), []int{1}
}

func (x *ProgressEvent) GetTimeOffset() *durationpb.Duration {
	if x != nil {
		return x.TimeOffset
	}
	return nil
}

func (x *ProgressEvent) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_google_cloud_video_stitcher_v1_events_proto protoreflect.FileDescriptor

var file_google_cloud_video_stitcher_v1_events_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x04,
	0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x22, 0x81, 0x03, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x16, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x52, 0x45,
	0x41, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x52,
	0x45, 0x41, 0x4b, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4d, 0x50,
	0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x49, 0x52,
	0x53, 0x54, 0x5f, 0x51, 0x55, 0x41, 0x52, 0x54, 0x49, 0x4c, 0x45, 0x10, 0x06, 0x12, 0x0c, 0x0a,
	0x08, 0x4d, 0x49, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x54,
	0x48, 0x49, 0x52, 0x44, 0x5f, 0x51, 0x55, 0x41, 0x52, 0x54, 0x49, 0x4c, 0x45, 0x10, 0x08, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x09, 0x12, 0x0c, 0x0a,
	0x08, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04, 0x4d,
	0x55, 0x54, 0x45, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4d, 0x55, 0x54, 0x45, 0x10,
	0x0c, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x55, 0x53, 0x45, 0x10, 0x0d, 0x12, 0x09, 0x0a, 0x05,
	0x43, 0x4c, 0x49, 0x43, 0x4b, 0x10, 0x0e, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4c, 0x49, 0x43, 0x4b,
	0x5f, 0x54, 0x48, 0x52, 0x4f, 0x55, 0x47, 0x48, 0x10, 0x0f, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45,
	0x57, 0x49, 0x4e, 0x44, 0x10, 0x10, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x53, 0x55, 0x4d, 0x45,
	0x10, 0x11, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x12, 0x12, 0x0a, 0x0a,
	0x06, 0x45, 0x58, 0x50, 0x41, 0x4e, 0x44, 0x10, 0x15, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4c,
	0x4c, 0x41, 0x50, 0x53, 0x45, 0x10, 0x16, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x53, 0x45,
	0x10, 0x18, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x5f, 0x4c, 0x49, 0x4e, 0x45,
	0x41, 0x52, 0x10, 0x19, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4b, 0x49, 0x50, 0x10, 0x1a, 0x12, 0x15,
	0x0a, 0x11, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x1b, 0x22, 0x8a, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x42, 0x73, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x74, 0x69,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_video_stitcher_v1_events_proto_rawDescOnce sync.Once
	file_google_cloud_video_stitcher_v1_events_proto_rawDescData = file_google_cloud_video_stitcher_v1_events_proto_rawDesc
)

func file_google_cloud_video_stitcher_v1_events_proto_rawDescGZIP() []byte {
	file_google_cloud_video_stitcher_v1_events_proto_rawDescOnce.Do(func() {
		file_google_cloud_video_stitcher_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_video_stitcher_v1_events_proto_rawDescData)
	})
	return file_google_cloud_video_stitcher_v1_events_proto_rawDescData
}

var file_google_cloud_video_stitcher_v1_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_video_stitcher_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_video_stitcher_v1_events_proto_goTypes = []interface{}{
	(Event_EventType)(0),        // 0: google.cloud.video.stitcher.v1.Event.EventType
	(*Event)(nil),               // 1: google.cloud.video.stitcher.v1.Event
	(*ProgressEvent)(nil),       // 2: google.cloud.video.stitcher.v1.ProgressEvent
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
}
var file_google_cloud_video_stitcher_v1_events_proto_depIdxs = []int32{
	0, // 0: google.cloud.video.stitcher.v1.Event.type:type_name -> google.cloud.video.stitcher.v1.Event.EventType
	3, // 1: google.cloud.video.stitcher.v1.Event.offset:type_name -> google.protobuf.Duration
	3, // 2: google.cloud.video.stitcher.v1.ProgressEvent.time_offset:type_name -> google.protobuf.Duration
	1, // 3: google.cloud.video.stitcher.v1.ProgressEvent.events:type_name -> google.cloud.video.stitcher.v1.Event
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_video_stitcher_v1_events_proto_init() }
func file_google_cloud_video_stitcher_v1_events_proto_init() {
	if File_google_cloud_video_stitcher_v1_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_video_stitcher_v1_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_google_cloud_video_stitcher_v1_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressEvent); i {
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
			RawDescriptor: file_google_cloud_video_stitcher_v1_events_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_video_stitcher_v1_events_proto_goTypes,
		DependencyIndexes: file_google_cloud_video_stitcher_v1_events_proto_depIdxs,
		EnumInfos:         file_google_cloud_video_stitcher_v1_events_proto_enumTypes,
		MessageInfos:      file_google_cloud_video_stitcher_v1_events_proto_msgTypes,
	}.Build()
	File_google_cloud_video_stitcher_v1_events_proto = out.File
	file_google_cloud_video_stitcher_v1_events_proto_rawDesc = nil
	file_google_cloud_video_stitcher_v1_events_proto_goTypes = nil
	file_google_cloud_video_stitcher_v1_events_proto_depIdxs = nil
}
