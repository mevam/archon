// Code generated by protoc-gen-go.
// source: google/type/latlng.proto
// DO NOT EDIT!

package google_type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An object representing a latitude/longitude pair. This is expressed as a pair
// of doubles representing degrees latitude and degrees longitude. Unless
// specified otherwise, this must conform to the
// <a href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
// standard</a>. Values must be within normalized ranges.
type LatLng struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *LatLng) Reset()                    { *m = LatLng{} }
func (m *LatLng) String() string            { return proto.CompactTextString(m) }
func (*LatLng) ProtoMessage()               {}
func (*LatLng) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func init() {
	proto.RegisterType((*LatLng)(nil), "google.type.LatLng")
}

func init() { proto.RegisterFile("google/type/latlng.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x49, 0x2c, 0xc9, 0xc9, 0x4b, 0xd7, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0xc8, 0xe8, 0x81, 0x64, 0x94, 0x9c, 0xb8, 0xd8, 0x7c,
	0x12, 0x4b, 0x7c, 0xf2, 0xd2, 0x85, 0xa4, 0xb8, 0x38, 0x80, 0xca, 0x32, 0x4b, 0x4a, 0x53, 0x52,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x18, 0x83, 0xe0, 0x7c, 0x21, 0x19, 0x2e, 0xce, 0x9c, 0xfc, 0xbc,
	0x74, 0x88, 0x24, 0x13, 0x58, 0x12, 0x21, 0xe0, 0xa4, 0xcc, 0xc5, 0x9f, 0x9c, 0x9f, 0xab, 0x87,
	0x64, 0xac, 0x13, 0x37, 0xc4, 0xd0, 0x00, 0x90, 0x85, 0x01, 0x8c, 0x0b, 0x18, 0x19, 0x93, 0xd8,
	0xc0, 0x96, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x89, 0xd0, 0x63, 0x98, 0x00, 0x00,
	0x00,
}
