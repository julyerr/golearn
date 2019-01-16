// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.megvii-inc.com/securitycore/face/cmd/capture-benchmark/capM/capture_message.proto

package capture_message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CaptureMessage struct {
	CropImage            []byte   `protobuf:"bytes,1,opt,name=CropImage,proto3" json:"CropImage,omitempty"`
	FaceImage            []byte   `protobuf:"bytes,2,opt,name=FaceImage,proto3" json:"FaceImage,omitempty"`
	TimeStamp            int64    `protobuf:"varint,3,opt,name=TimeStamp,proto3" json:"TimeStamp,omitempty"`
	Track                int64    `protobuf:"varint,4,opt,name=Track,proto3" json:"Track,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=Type,proto3" json:"Type,omitempty"`
	SequenceNumber       int64    `protobuf:"varint,6,opt,name=SequenceNumber,proto3" json:"SequenceNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CaptureMessage) Reset()         { *m = CaptureMessage{} }
func (m *CaptureMessage) String() string { return proto.CompactTextString(m) }
func (*CaptureMessage) ProtoMessage()    {}
func (*CaptureMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_04678640609360e7, []int{0}
}

func (m *CaptureMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CaptureMessage.Unmarshal(m, b)
}
func (m *CaptureMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CaptureMessage.Marshal(b, m, deterministic)
}
func (m *CaptureMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CaptureMessage.Merge(m, src)
}
func (m *CaptureMessage) XXX_Size() int {
	return xxx_messageInfo_CaptureMessage.Size(m)
}
func (m *CaptureMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CaptureMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CaptureMessage proto.InternalMessageInfo

func (m *CaptureMessage) GetCropImage() []byte {
	if m != nil {
		return m.CropImage
	}
	return nil
}

func (m *CaptureMessage) GetFaceImage() []byte {
	if m != nil {
		return m.FaceImage
	}
	return nil
}

func (m *CaptureMessage) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *CaptureMessage) GetTrack() int64 {
	if m != nil {
		return m.Track
	}
	return 0
}

func (m *CaptureMessage) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CaptureMessage) GetSequenceNumber() int64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*CaptureMessage)(nil), "capture.message.CaptureMessage")
}

func init() {
	proto.RegisterFile("go.megvii-inc.com/securitycore/face/cmd/capture-benchmark/capM/capture_message.proto", fileDescriptor_04678640609360e7)
}

var fileDescriptor_04678640609360e7 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xcf, 0xbf, 0x4a, 0xc4, 0x40,
	0x10, 0xc7, 0x71, 0xd6, 0xfb, 0x03, 0xb7, 0xc8, 0x09, 0x8b, 0xc5, 0x16, 0x16, 0xc1, 0x42, 0xd2,
	0x5c, 0x52, 0xf8, 0x08, 0x07, 0x82, 0xc5, 0x59, 0xe4, 0xd2, 0xcb, 0x66, 0xfc, 0x19, 0x97, 0x63,
	0xb2, 0xeb, 0x26, 0x11, 0xee, 0xdd, 0x7c, 0x38, 0xc9, 0xee, 0x69, 0xc0, 0x6e, 0xe6, 0xf3, 0x65,
	0x8a, 0x91, 0x75, 0xeb, 0x0a, 0x46, 0xfb, 0x65, 0xed, 0xce, 0x76, 0x54, 0x90, 0xe3, 0xb2, 0x07,
	0x8d, 0xc1, 0x0e, 0x67, 0x72, 0x01, 0xe5, 0xbb, 0x21, 0x94, 0xc4, 0x6f, 0x25, 0x19, 0x3f, 0x8c,
	0x01, 0xbb, 0x06, 0x1d, 0x7d, 0xb0, 0x09, 0xa7, 0x49, 0x0e, 0xbf, 0xfc, 0xca, 0xe8, 0x7b, 0xd3,
	0xa2, 0xf0, 0xc1, 0x0d, 0x4e, 0xdd, 0x5c, 0xb8, 0xb8, 0xf0, 0xfd, 0xb7, 0x90, 0xdb, 0x7d, 0xb2,
	0x43, 0x22, 0x75, 0x27, 0x37, 0xfb, 0xe0, 0xfc, 0x33, 0x9b, 0x16, 0x5a, 0x64, 0x22, 0xbf, 0xae,
	0x66, 0x98, 0xea, 0x93, 0x21, 0xa4, 0x7a, 0x95, 0xea, 0x1f, 0x4c, 0xb5, 0xb6, 0x8c, 0xe3, 0x60,
	0xd8, 0xeb, 0x45, 0x26, 0xf2, 0x45, 0x35, 0x83, 0xba, 0x95, 0xab, 0x3a, 0x18, 0x3a, 0xe9, 0x65,
	0x2c, 0x69, 0x51, 0x4a, 0x2e, 0xeb, 0xb3, 0x87, 0x5e, 0x65, 0x22, 0xdf, 0x54, 0x71, 0x56, 0x0f,
	0x72, 0x7b, 0xc4, 0xe7, 0x88, 0x8e, 0xf0, 0x32, 0x72, 0x83, 0xa0, 0xd7, 0xf1, 0xe4, 0x9f, 0x36,
	0xeb, 0xf8, 0xd6, 0xe3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0xb6, 0x0a, 0x36, 0x2e, 0x01,
	0x00, 0x00,
}
