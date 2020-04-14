// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/conversion_adjustment_upload_error.proto

package errors

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

// Enum describing possible conversion adjustment upload errors.
type ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError int32

const (
	// Not specified.
	ConversionAdjustmentUploadErrorEnum_UNSPECIFIED ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 0
	// The received error code is not known in this version.
	ConversionAdjustmentUploadErrorEnum_UNKNOWN ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 1
	// The specified conversion action was created too recently.
	// Please try the upload again after 4-6 hours have passed since the
	// conversion action was created.
	ConversionAdjustmentUploadErrorEnum_TOO_RECENT_CONVERSION_ACTION ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 2
	// No conversion action of a supported ConversionActionType that matches the
	// provided information can be found for the customer.
	ConversionAdjustmentUploadErrorEnum_INVALID_CONVERSION_ACTION ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 3
	// A retraction was already reported for this conversion.
	ConversionAdjustmentUploadErrorEnum_CONVERSION_ALREADY_RETRACTED ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 4
	// A conversion for the supplied combination of conversion
	// action and conversion identifier could not be found.
	ConversionAdjustmentUploadErrorEnum_CONVERSION_NOT_FOUND ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 5
	// The specified conversion has already expired. Conversions expire after 55
	// days, after which adjustments cannot be reported against them.
	ConversionAdjustmentUploadErrorEnum_CONVERSION_EXPIRED ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 6
	// The supplied adjustment date time precedes that of the original
	// conversion.
	ConversionAdjustmentUploadErrorEnum_ADJUSTMENT_PRECEDES_CONVERSION ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 7
	// A restatement with a more recent adjustment date time was already
	// reported for this conversion.
	ConversionAdjustmentUploadErrorEnum_MORE_RECENT_RESTATEMENT_FOUND ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 8
	// The conversion was created too recently.
	ConversionAdjustmentUploadErrorEnum_TOO_RECENT_CONVERSION ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 9
	// Restatements cannot be reported for a conversion action that always uses
	// the default value.
	ConversionAdjustmentUploadErrorEnum_CANNOT_RESTATE_CONVERSION_ACTION_THAT_ALWAYS_USES_DEFAULT_CONVERSION_VALUE ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 10
	// The request contained more than 2000 adjustments.
	ConversionAdjustmentUploadErrorEnum_TOO_MANY_ADJUSTMENTS_IN_REQUEST ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 11
	// The conversion has been adjusted too many times.
	ConversionAdjustmentUploadErrorEnum_TOO_MANY_ADJUSTMENTS ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError = 12
)

var ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "TOO_RECENT_CONVERSION_ACTION",
	3:  "INVALID_CONVERSION_ACTION",
	4:  "CONVERSION_ALREADY_RETRACTED",
	5:  "CONVERSION_NOT_FOUND",
	6:  "CONVERSION_EXPIRED",
	7:  "ADJUSTMENT_PRECEDES_CONVERSION",
	8:  "MORE_RECENT_RESTATEMENT_FOUND",
	9:  "TOO_RECENT_CONVERSION",
	10: "CANNOT_RESTATE_CONVERSION_ACTION_THAT_ALWAYS_USES_DEFAULT_CONVERSION_VALUE",
	11: "TOO_MANY_ADJUSTMENTS_IN_REQUEST",
	12: "TOO_MANY_ADJUSTMENTS",
}

var ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError_value = map[string]int32{
	"UNSPECIFIED":                    0,
	"UNKNOWN":                        1,
	"TOO_RECENT_CONVERSION_ACTION":   2,
	"INVALID_CONVERSION_ACTION":      3,
	"CONVERSION_ALREADY_RETRACTED":   4,
	"CONVERSION_NOT_FOUND":           5,
	"CONVERSION_EXPIRED":             6,
	"ADJUSTMENT_PRECEDES_CONVERSION": 7,
	"MORE_RECENT_RESTATEMENT_FOUND":  8,
	"TOO_RECENT_CONVERSION":          9,
	"CANNOT_RESTATE_CONVERSION_ACTION_THAT_ALWAYS_USES_DEFAULT_CONVERSION_VALUE": 10,
	"TOO_MANY_ADJUSTMENTS_IN_REQUEST":                                            11,
	"TOO_MANY_ADJUSTMENTS":                                                       12,
}

func (x ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError) String() string {
	return proto.EnumName(ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError_name, int32(x))
}

func (ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_80983c668085e5a2, []int{0, 0}
}

// Container for enum describing possible conversion adjustment upload errors.
type ConversionAdjustmentUploadErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionAdjustmentUploadErrorEnum) Reset()         { *m = ConversionAdjustmentUploadErrorEnum{} }
func (m *ConversionAdjustmentUploadErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionAdjustmentUploadErrorEnum) ProtoMessage()    {}
func (*ConversionAdjustmentUploadErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_80983c668085e5a2, []int{0}
}

func (m *ConversionAdjustmentUploadErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionAdjustmentUploadErrorEnum.Unmarshal(m, b)
}
func (m *ConversionAdjustmentUploadErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionAdjustmentUploadErrorEnum.Marshal(b, m, deterministic)
}
func (m *ConversionAdjustmentUploadErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionAdjustmentUploadErrorEnum.Merge(m, src)
}
func (m *ConversionAdjustmentUploadErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionAdjustmentUploadErrorEnum.Size(m)
}
func (m *ConversionAdjustmentUploadErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionAdjustmentUploadErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionAdjustmentUploadErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError", ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError_name, ConversionAdjustmentUploadErrorEnum_ConversionAdjustmentUploadError_value)
	proto.RegisterType((*ConversionAdjustmentUploadErrorEnum)(nil), "google.ads.googleads.v3.errors.ConversionAdjustmentUploadErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/conversion_adjustment_upload_error.proto", fileDescriptor_80983c668085e5a2)
}

var fileDescriptor_80983c668085e5a2 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x59, 0x0b, 0x1b, 0xb8, 0x48, 0x58, 0x16, 0x20, 0x86, 0xb6, 0x0e, 0x3a, 0xae, 0x93,
	0x8b, 0xdc, 0x85, 0x2b, 0x2f, 0x39, 0x2d, 0x19, 0xad, 0x53, 0x12, 0x27, 0xa3, 0xa8, 0x92, 0x15,
	0x96, 0x2a, 0x2a, 0x6a, 0xe3, 0x2a, 0x6e, 0xfb, 0x40, 0x5c, 0xf2, 0x1e, 0xdc, 0xf0, 0x28, 0x48,
	0x3c, 0x03, 0xc8, 0xf1, 0x5a, 0x8a, 0x28, 0xec, 0x2a, 0x47, 0xf6, 0x77, 0xfe, 0xff, 0xf7, 0xc9,
	0x41, 0xbd, 0x42, 0xca, 0x62, 0x36, 0xb1, 0xb3, 0x5c, 0xd9, 0xa6, 0xd4, 0xd5, 0xda, 0xb1, 0x27,
	0x55, 0x25, 0x2b, 0x65, 0x5f, 0xcb, 0x72, 0x3d, 0xa9, 0xd4, 0x54, 0x96, 0x22, 0xcb, 0x3f, 0xad,
	0xd4, 0x72, 0x3e, 0x29, 0x97, 0x62, 0xb5, 0x98, 0xc9, 0x2c, 0x17, 0x35, 0x63, 0x2d, 0x2a, 0xb9,
	0x94, 0xa4, 0x6d, 0xba, 0xad, 0x2c, 0x57, 0xd6, 0x56, 0xc8, 0x5a, 0x3b, 0x96, 0x11, 0x7a, 0x7e,
	0xb2, 0x31, 0x5a, 0x4c, 0xed, 0xac, 0x2c, 0xe5, 0x32, 0x5b, 0x4e, 0x65, 0xa9, 0x4c, 0x77, 0xe7,
	0x47, 0x13, 0x9d, 0x7b, 0x5b, 0x2b, 0xba, 0x75, 0x4a, 0x6a, 0x23, 0xd0, 0x12, 0x50, 0xae, 0xe6,
	0x9d, 0xaf, 0x4d, 0x74, 0x76, 0x0b, 0x47, 0x1e, 0xa1, 0x56, 0xc2, 0xe2, 0x21, 0x78, 0x41, 0x37,
	0x00, 0x1f, 0xdf, 0x21, 0x2d, 0x74, 0x94, 0xb0, 0xb7, 0x2c, 0xbc, 0x62, 0xf8, 0x80, 0xbc, 0x40,
	0x27, 0x3c, 0x0c, 0x45, 0x04, 0x1e, 0x30, 0x2e, 0xbc, 0x90, 0xa5, 0x10, 0xc5, 0x41, 0xc8, 0x04,
	0xf5, 0x78, 0x10, 0x32, 0xdc, 0x20, 0xa7, 0xe8, 0x38, 0x60, 0x29, 0xed, 0x07, 0xfe, 0x9e, 0xeb,
	0xa6, 0x16, 0xd8, 0x3d, 0xee, 0x47, 0x40, 0xfd, 0x91, 0x88, 0x80, 0x47, 0xd4, 0xe3, 0xe0, 0xe3,
	0xbb, 0xe4, 0x19, 0x7a, 0xbc, 0x43, 0xb0, 0x90, 0x8b, 0x6e, 0x98, 0x30, 0x1f, 0xdf, 0x23, 0x4f,
	0x11, 0xd9, 0xb9, 0x81, 0xf7, 0xc3, 0x20, 0x02, 0x1f, 0x1f, 0x92, 0x0e, 0x6a, 0x53, 0xff, 0x32,
	0x89, 0xf9, 0x40, 0x87, 0x1a, 0xea, 0x70, 0x3e, 0xc4, 0x3b, 0xf6, 0xf8, 0x88, 0xbc, 0x44, 0xa7,
	0x83, 0x30, 0x82, 0x4d, 0xf2, 0x08, 0x62, 0x4e, 0x39, 0xd4, 0x0d, 0x46, 0xfe, 0x3e, 0x39, 0x46,
	0x4f, 0xf6, 0xbe, 0x0d, 0x3f, 0x20, 0x0c, 0x5d, 0x7a, 0x94, 0xe9, 0x2c, 0x37, 0x8d, 0x7f, 0xbf,
	0x4d, 0xf0, 0x37, 0x94, 0x0b, 0xda, 0xbf, 0xa2, 0xa3, 0x58, 0x24, 0x31, 0xc4, 0xc2, 0x87, 0x2e,
	0x4d, 0xfa, 0x7f, 0x4c, 0x29, 0xa5, 0xfd, 0x04, 0x30, 0x22, 0xe7, 0xe8, 0x4c, 0x5b, 0x0d, 0x28,
	0x1b, 0x89, 0xdf, 0xd1, 0x63, 0x11, 0x30, 0x11, 0xc1, 0xbb, 0x04, 0x62, 0x8e, 0x5b, 0x7a, 0x10,
	0xfb, 0x20, 0xfc, 0xf0, 0xe2, 0xe7, 0x01, 0xea, 0x5c, 0xcb, 0xb9, 0xf5, 0xff, 0xa5, 0xb9, 0x78,
	0x75, 0xcb, 0xbf, 0x1e, 0xea, 0xe5, 0x19, 0x1e, 0x7c, 0xf0, 0x6f, 0x74, 0x0a, 0x39, 0xcb, 0xca,
	0xc2, 0x92, 0x55, 0x61, 0x17, 0x93, 0xb2, 0x5e, 0xad, 0xcd, 0x56, 0x2f, 0xa6, 0xea, 0x5f, 0x4b,
	0xfe, 0xda, 0x7c, 0x3e, 0x37, 0x9a, 0x3d, 0x4a, 0xbf, 0x34, 0xda, 0x3d, 0x23, 0x46, 0x73, 0x65,
	0x99, 0x52, 0x57, 0xa9, 0x63, 0xd5, 0x96, 0xea, 0xdb, 0x06, 0x18, 0xd3, 0x5c, 0x8d, 0xb7, 0xc0,
	0x38, 0x75, 0xc6, 0x06, 0xf8, 0xde, 0xe8, 0x98, 0x53, 0xd7, 0xa5, 0xb9, 0x72, 0xdd, 0x2d, 0xe2,
	0xba, 0xa9, 0xe3, 0xba, 0x06, 0xfa, 0x78, 0x58, 0xa7, 0x73, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x6c, 0xec, 0x4a, 0x41, 0x81, 0x03, 0x00, 0x00,
}