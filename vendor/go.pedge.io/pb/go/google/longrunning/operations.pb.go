// Code generated by protoc-gen-go.
// source: google/longrunning/operations.proto
// DO NOT EDIT!

package google_longrunning

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "go.pedge.io/pb/go/google/protobuf"
import _ "go.pedge.io/pb/go/google/protobuf"
import google_rpc "go.pedge.io/pb/go/google/rpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// The name of the operation resource, which is only unique within the same
	// service that originally returns it.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Some service-specific metadata associated with the operation.  It typically
	// contains progress information and common metadata such as create time.
	// Some services may not provide such metadata.  Any method that returns a
	// long-running operation should document the metadata type, if any.
	Metadata *google_protobuf1.Any `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
	// If the value is false, it means the operation is still in progress.
	// If true, the operation is completed and the `result` is available.
	Done bool `protobuf:"varint,3,opt,name=done" json:"done,omitempty"`
	// Types that are valid to be assigned to Result:
	//	*Operation_Error
	//	*Operation_Response
	Result isOperation_Result `protobuf_oneof:"result"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isOperation_Result interface {
	isOperation_Result()
}

type Operation_Error struct {
	Error *google_rpc.Status `protobuf:"bytes,4,opt,name=error,oneof"`
}
type Operation_Response struct {
	Response *google_protobuf1.Any `protobuf:"bytes,5,opt,name=response,oneof"`
}

func (*Operation_Error) isOperation_Result()    {}
func (*Operation_Response) isOperation_Result() {}

func (m *Operation) GetResult() isOperation_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Operation) GetMetadata() *google_protobuf1.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Operation) GetError() *google_rpc.Status {
	if x, ok := m.GetResult().(*Operation_Error); ok {
		return x.Error
	}
	return nil
}

func (m *Operation) GetResponse() *google_protobuf1.Any {
	if x, ok := m.GetResult().(*Operation_Response); ok {
		return x.Response
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operation_OneofMarshaler, _Operation_OneofUnmarshaler, _Operation_OneofSizer, []interface{}{
		(*Operation_Error)(nil),
		(*Operation_Response)(nil),
	}
}

func _Operation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operation)
	// result
	switch x := m.Result.(type) {
	case *Operation_Error:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *Operation_Response:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Response); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Operation.Result has unexpected type %T", x)
	}
	return nil
}

func _Operation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Operation)
	switch tag {
	case 4: // result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_rpc.Status)
		err := b.DecodeMessage(msg)
		m.Result = &Operation_Error{msg}
		return true, err
	case 5: // result.response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Any)
		err := b.DecodeMessage(msg)
		m.Result = &Operation_Response{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Operation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Operation)
	// result
	switch x := m.Result.(type) {
	case *Operation_Error:
		s := proto.Size(x.Error)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Response:
		s := proto.Size(x.Response)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The request message for [Operations.GetOperation][google.longrunning.Operations.GetOperation].
type GetOperationRequest struct {
	// The name of the operation resource.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetOperationRequest) Reset()                    { *m = GetOperationRequest{} }
func (m *GetOperationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOperationRequest) ProtoMessage()               {}
func (*GetOperationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// The request message for [Operations.ListOperations][google.longrunning.Operations.ListOperations].
type ListOperationsRequest struct {
	// The name of the operation collection.
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	// The standard List filter.
	Filter string `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	// The standard List page size.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// The standard List page token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListOperationsRequest) Reset()                    { *m = ListOperationsRequest{} }
func (m *ListOperationsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListOperationsRequest) ProtoMessage()               {}
func (*ListOperationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// The response message for [Operations.ListOperations][google.longrunning.Operations.ListOperations].
type ListOperationsResponse struct {
	// A list of operations that match the specified filter in the request.
	Operations []*Operation `protobuf:"bytes,1,rep,name=operations" json:"operations,omitempty"`
	// The standard List next-page token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListOperationsResponse) Reset()                    { *m = ListOperationsResponse{} }
func (m *ListOperationsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListOperationsResponse) ProtoMessage()               {}
func (*ListOperationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListOperationsResponse) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// The request message for [Operations.CancelOperation][google.longrunning.Operations.CancelOperation].
type CancelOperationRequest struct {
	// The name of the operation resource to be cancelled.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CancelOperationRequest) Reset()                    { *m = CancelOperationRequest{} }
func (m *CancelOperationRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelOperationRequest) ProtoMessage()               {}
func (*CancelOperationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// The request message for [Operations.DeleteOperation][google.longrunning.Operations.DeleteOperation].
type DeleteOperationRequest struct {
	// The name of the operation resource to be deleted.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteOperationRequest) Reset()                    { *m = DeleteOperationRequest{} }
func (m *DeleteOperationRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteOperationRequest) ProtoMessage()               {}
func (*DeleteOperationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*Operation)(nil), "google.longrunning.Operation")
	proto.RegisterType((*GetOperationRequest)(nil), "google.longrunning.GetOperationRequest")
	proto.RegisterType((*ListOperationsRequest)(nil), "google.longrunning.ListOperationsRequest")
	proto.RegisterType((*ListOperationsResponse)(nil), "google.longrunning.ListOperationsResponse")
	proto.RegisterType((*CancelOperationRequest)(nil), "google.longrunning.CancelOperationRequest")
	proto.RegisterType((*DeleteOperationRequest)(nil), "google.longrunning.DeleteOperationRequest")
}

func init() { proto.RegisterFile("google/longrunning/operations.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xfd, 0x9c, 0x26, 0x51, 0x32, 0x1f, 0x10, 0x69, 0xa0, 0xa9, 0x71, 0x89, 0xa8, 0x0c, 0x02,
	0x6a, 0x90, 0x0d, 0xe1, 0x56, 0xa9, 0x07, 0x02, 0x08, 0x0e, 0x48, 0x44, 0x2e, 0xf7, 0xca, 0x4d,
	0xa7, 0x91, 0x85, 0xb3, 0x6b, 0xec, 0x35, 0xb4, 0xa0, 0x0a, 0xc1, 0x81, 0x13, 0x37, 0x7e, 0x18,
	0x07, 0xfe, 0x02, 0x3f, 0x84, 0xf5, 0xc6, 0x89, 0x8d, 0xbb, 0x41, 0xb9, 0xed, 0xee, 0xbc, 0x7d,
	0xf3, 0xe6, 0xbd, 0x5d, 0xb8, 0x35, 0xe5, 0x7c, 0x1a, 0x91, 0x17, 0x71, 0x36, 0x4d, 0x32, 0xc6,
	0x42, 0x36, 0xf5, 0x78, 0x4c, 0x49, 0x20, 0x42, 0xce, 0x52, 0x37, 0x4e, 0xb8, 0xe0, 0x88, 0x73,
	0x90, 0x5b, 0x01, 0x59, 0x37, 0x8a, 0x8b, 0x41, 0x1c, 0x7a, 0x01, 0x63, 0x5c, 0x54, 0x6f, 0x58,
	0xd7, 0x8b, 0xaa, 0xda, 0x1d, 0x65, 0x27, 0x12, 0x72, 0x56, 0x94, 0xb6, 0xeb, 0x25, 0x9a, 0xc5,
	0x62, 0x51, 0xdc, 0x2a, 0x8a, 0x49, 0x3c, 0xf1, 0x52, 0x49, 0x99, 0x15, 0x84, 0xf6, 0x4f, 0x03,
	0xba, 0xaf, 0x17, 0xba, 0x10, 0xa1, 0xc9, 0x82, 0x19, 0x99, 0xc6, 0x8e, 0x71, 0xaf, 0xeb, 0xab,
	0x35, 0x3e, 0x84, 0xce, 0x8c, 0x44, 0x70, 0x1c, 0x88, 0xc0, 0x6c, 0xc8, 0xf3, 0xff, 0x87, 0xd7,
	0xdc, 0x42, 0xf7, 0xa2, 0x95, 0xfb, 0x84, 0x9d, 0xf9, 0x4b, 0x54, 0xce, 0x72, 0xcc, 0x19, 0x99,
	0x1b, 0x12, 0xdd, 0xf1, 0xd5, 0x1a, 0x1d, 0x68, 0x51, 0x92, 0xf0, 0xc4, 0x6c, 0x2a, 0x0a, 0x5c,
	0x50, 0x48, 0x41, 0xee, 0x81, 0x12, 0xf4, 0xf2, 0x3f, 0x7f, 0x0e, 0xc1, 0x21, 0x74, 0x12, 0x4a,
	0x63, 0x39, 0x35, 0x99, 0xad, 0xd5, 0x1d, 0xe5, 0x85, 0x25, 0x6e, 0xd4, 0x81, 0xb6, 0x5c, 0x67,
	0x91, 0xb0, 0x77, 0xe1, 0xea, 0x0b, 0x12, 0xcb, 0x99, 0x7c, 0x7a, 0x97, 0x51, 0x2a, 0x74, 0xa3,
	0xd9, 0x9f, 0x61, 0xf3, 0x55, 0x98, 0x96, 0xd8, 0xb4, 0x0e, 0x6e, 0x56, 0x7c, 0xe8, 0x43, 0xfb,
	0x24, 0x8c, 0x04, 0x25, 0x05, 0x45, 0xb1, 0xc3, 0x6d, 0xe8, 0xc6, 0xc1, 0x94, 0x0e, 0xd3, 0xf0,
	0x23, 0x29, 0x83, 0x5a, 0x7e, 0x27, 0x3f, 0x38, 0x90, 0x7b, 0x1c, 0x00, 0xa8, 0xa2, 0xe0, 0x6f,
	0x89, 0x29, 0x43, 0xba, 0xbe, 0x82, 0xbf, 0xc9, 0x0f, 0xa4, 0x80, 0x7e, 0x5d, 0xc0, 0x7c, 0x1e,
	0xdc, 0x07, 0x28, 0x9f, 0x8b, 0xec, 0xb8, 0x21, 0x5d, 0x18, 0xb8, 0x17, 0xdf, 0x8b, 0x5b, 0x0e,
	0x5a, 0xb9, 0x80, 0x77, 0xa0, 0xc7, 0xe8, 0x54, 0x1c, 0x56, 0x9a, 0x37, 0x54, 0xf3, 0xcb, 0xf9,
	0xf1, 0x78, 0x29, 0xe0, 0x01, 0xf4, 0x9f, 0x06, 0x6c, 0x42, 0xd1, 0x5a, 0x7e, 0x49, 0xf4, 0x33,
	0x8a, 0x48, 0xd0, 0x3a, 0xe8, 0xe1, 0xf7, 0x26, 0x40, 0x39, 0x19, 0x9e, 0xc2, 0xa5, 0x6a, 0x2e,
	0x78, 0x57, 0x37, 0x8d, 0x26, 0x39, 0xeb, 0xdf, 0x63, 0xdb, 0x3b, 0x5f, 0x7f, 0xfd, 0xfe, 0xd1,
	0xb0, 0xd0, 0xf4, 0xde, 0x3f, 0xf2, 0x3e, 0xe5, 0x9d, 0xf7, 0x4b, 0x23, 0x3c, 0xc7, 0x39, 0xc7,
	0x6f, 0x06, 0x5c, 0xf9, 0xdb, 0x66, 0xdc, 0xd5, 0x71, 0x6a, 0xdf, 0x82, 0xe5, 0xac, 0x03, 0x9d,
	0xa7, 0x66, 0x0f, 0x94, 0x96, 0x2d, 0xdc, 0xd4, 0x69, 0x39, 0xc7, 0x2f, 0x06, 0xf4, 0x6a, 0x76,
	0xa3, 0x96, 0x5e, 0x9f, 0x89, 0xd5, 0xbf, 0xf0, 0x0d, 0x9e, 0xe7, 0x7f, 0xdc, 0x76, 0x54, 0xdb,
	0xdb, 0xf6, 0xcd, 0x55, 0x16, 0xec, 0x4d, 0x14, 0xe1, 0x9e, 0xe1, 0xe0, 0x07, 0xe8, 0xd5, 0x32,
	0xd4, 0x4b, 0xd0, 0x07, 0xbd, 0x52, 0x42, 0x91, 0x82, 0xb3, 0x32, 0x85, 0xd1, 0x7d, 0xe8, 0x4f,
	0xf8, 0x4c, 0xd3, 0x6a, 0xd4, 0x2b, 0x9d, 0x1c, 0xe7, 0xac, 0x63, 0xe3, 0xa8, 0xad, 0xe8, 0x1f,
	0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x2c, 0x30, 0xa8, 0x47, 0x05, 0x00, 0x00,
}
