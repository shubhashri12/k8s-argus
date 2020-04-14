// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/ad_schedule_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// An ad schedule view summarizes the performance of campaigns by
// AdSchedule criteria.
type AdScheduleView struct {
	// The resource name of the ad schedule view.
	// AdSchedule view resource names have the form:
	//
	// `customers/{customer_id}/adScheduleViews/{campaign_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdScheduleView) Reset()         { *m = AdScheduleView{} }
func (m *AdScheduleView) String() string { return proto.CompactTextString(m) }
func (*AdScheduleView) ProtoMessage()    {}
func (*AdScheduleView) Descriptor() ([]byte, []int) {
	return fileDescriptor_99aaa568f8d5a785, []int{0}
}

func (m *AdScheduleView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdScheduleView.Unmarshal(m, b)
}
func (m *AdScheduleView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdScheduleView.Marshal(b, m, deterministic)
}
func (m *AdScheduleView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdScheduleView.Merge(m, src)
}
func (m *AdScheduleView) XXX_Size() int {
	return xxx_messageInfo_AdScheduleView.Size(m)
}
func (m *AdScheduleView) XXX_DiscardUnknown() {
	xxx_messageInfo_AdScheduleView.DiscardUnknown(m)
}

var xxx_messageInfo_AdScheduleView proto.InternalMessageInfo

func (m *AdScheduleView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*AdScheduleView)(nil), "google.ads.googleads.v3.resources.AdScheduleView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/ad_schedule_view.proto", fileDescriptor_99aaa568f8d5a785)
}

var fileDescriptor_99aaa568f8d5a785 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0xb9, 0x13, 0x04, 0x0f, 0xb5, 0x88, 0x8d, 0x06, 0x0b, 0xa3, 0x04, 0xad, 0x76, 0x8b,
	0x2b, 0x94, 0xb5, 0xda, 0x34, 0x01, 0x0b, 0x09, 0x11, 0xae, 0x90, 0x83, 0xb0, 0xb9, 0x1d, 0xce,
	0x83, 0xdc, 0x4e, 0xb8, 0xb9, 0x24, 0x45, 0x08, 0xf8, 0x20, 0x56, 0x96, 0x3e, 0x8a, 0x8f, 0x92,
	0xa7, 0x90, 0x64, 0xb3, 0x1b, 0x93, 0x42, 0xbb, 0x8f, 0x9b, 0x6f, 0xfe, 0x99, 0xb9, 0x8d, 0x1e,
	0x72, 0xc4, 0x7c, 0x04, 0x5c, 0x69, 0xe2, 0x16, 0x57, 0x34, 0x8d, 0x79, 0x05, 0x84, 0x93, 0x2a,
	0x03, 0xe2, 0x4a, 0x0f, 0x28, 0x7b, 0x03, 0x3d, 0x19, 0xc1, 0x60, 0x5a, 0xc0, 0x8c, 0x8d, 0x2b,
	0xac, 0xb1, 0xd1, 0xb2, 0x3a, 0x53, 0x9a, 0x98, 0xef, 0x64, 0xd3, 0x98, 0xf9, 0xce, 0xe6, 0x85,
	0x0b, 0x1f, 0x17, 0x3e, 0xcf, 0x76, 0x37, 0x2f, 0x7f, 0x95, 0x94, 0x31, 0x58, 0xab, 0xba, 0x40,
	0x43, 0xb6, 0x7a, 0xfd, 0x11, 0x44, 0xa7, 0x52, 0xbf, 0x6c, 0xa6, 0x26, 0x05, 0xcc, 0x1a, 0x37,
	0xd1, 0x89, 0x8b, 0x18, 0x18, 0x55, 0xc2, 0x79, 0x70, 0x15, 0xdc, 0x1d, 0xf5, 0x8f, 0xdd, 0xc7,
	0x67, 0x55, 0x82, 0x80, 0xa5, 0x1c, 0x46, 0xb7, 0xdb, 0x6d, 0x36, 0x34, 0x2e, 0x88, 0x65, 0x58,
	0xf2, 0xbd, 0xc8, 0xfb, 0x6c, 0x42, 0x35, 0x96, 0x50, 0x11, 0x9f, 0x3b, 0x5c, 0x70, 0xb5, 0x23,
	0x11, 0x9f, 0xef, 0xdf, 0xbf, 0xe8, 0xbc, 0x87, 0x51, 0x3b, 0xc3, 0x92, 0xfd, 0xfb, 0x07, 0x3a,
	0x67, 0xbb, 0x23, 0x7b, 0xab, 0xeb, 0x7a, 0xc1, 0xeb, 0xd3, 0xa6, 0x33, 0xc7, 0x91, 0x32, 0x39,
	0xc3, 0x2a, 0xe7, 0x39, 0x98, 0xf5, 0xed, 0x7c, 0xbb, 0xf3, 0x1f, 0x8f, 0xf2, 0xe8, 0xe9, 0x33,
	0x3c, 0xe8, 0x4a, 0xf9, 0x15, 0xb6, 0xba, 0x36, 0x52, 0x6a, 0x62, 0x16, 0x57, 0x94, 0xc4, 0xac,
	0xef, 0xcc, 0x6f, 0xe7, 0xa4, 0x52, 0x53, 0xea, 0x9d, 0x34, 0x89, 0x53, 0xef, 0x2c, 0xc3, 0xb6,
	0x2d, 0x08, 0x21, 0x35, 0x09, 0xe1, 0x2d, 0x21, 0x92, 0x58, 0x08, 0xef, 0x0d, 0x0f, 0xd7, 0xcb,
	0xc6, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x6d, 0xce, 0x2d, 0x40, 0x02, 0x00, 0x00,
}