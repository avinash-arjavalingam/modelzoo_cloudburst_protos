// Code generated by protoc-gen-go. DO NOT EDIT.
// source: causal.proto

package common

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A message representing an individual tuple when the system is operating in
// causal mode.
type CausalTuple struct {
	// The name of the key-value pair this tuple represents.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The serialized payload of this key-value pair.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Any errors generated while retrieving this tuple (see CausalError).
	Error                AnnaError `protobuf:"varint,3,opt,name=error,proto3,enum=AnnaError" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CausalTuple) Reset()         { *m = CausalTuple{} }
func (m *CausalTuple) String() string { return proto.CompactTextString(m) }
func (*CausalTuple) ProtoMessage()    {}
func (*CausalTuple) Descriptor() ([]byte, []int) {
	return fileDescriptor_0056cb66ec996b89, []int{0}
}

func (m *CausalTuple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CausalTuple.Unmarshal(m, b)
}
func (m *CausalTuple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CausalTuple.Marshal(b, m, deterministic)
}
func (m *CausalTuple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CausalTuple.Merge(m, src)
}
func (m *CausalTuple) XXX_Size() int {
	return xxx_messageInfo_CausalTuple.Size(m)
}
func (m *CausalTuple) XXX_DiscardUnknown() {
	xxx_messageInfo_CausalTuple.DiscardUnknown(m)
}

var xxx_messageInfo_CausalTuple proto.InternalMessageInfo

func (m *CausalTuple) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CausalTuple) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CausalTuple) GetError() AnnaError {
	if m != nil {
		return m.Error
	}
	return AnnaError_NO_ERROR
}

// A request for one or more causal tuples, made from an executor and answered
// by a cache.
type CausalRequest struct {
	// The consistency mode that the request is being run in -- this should only
	// be SINGLE or MULTI.
	Consistency ConsistencyType `protobuf:"varint,1,opt,name=consistency,proto3,enum=ConsistencyType" json:"consistency,omitempty"`
	// The unique ID of the client that is making this request -- this may be
	// independent from the physical node on which the request is running.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The set of CausalTuples which are being requested from the cache.
	Tuples []*CausalTuple `protobuf:"bytes,3,rep,name=tuples,proto3" json:"tuples,omitempty"`
	// The set of keys that will be read by downstream function executors.
	FutureReadSet []string `protobuf:"bytes,4,rep,name=future_read_set,json=futureReadSet,proto3" json:"future_read_set,omitempty"`
	// The locations of upstream caches at which particular key versions are
	// stored -- this is used to optimize version retrieval.
	KeyVersionLocations map[string]*KeyVersionList `protobuf:"bytes,5,rep,name=key_version_locations,json=keyVersionLocations,proto3" json:"key_version_locations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The address at which the client will be waiting for a response for this
	// particular request.
	ResponseAddress      string   `protobuf:"bytes,6,opt,name=response_address,json=responseAddress,proto3" json:"response_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CausalRequest) Reset()         { *m = CausalRequest{} }
func (m *CausalRequest) String() string { return proto.CompactTextString(m) }
func (*CausalRequest) ProtoMessage()    {}
func (*CausalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0056cb66ec996b89, []int{1}
}

func (m *CausalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CausalRequest.Unmarshal(m, b)
}
func (m *CausalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CausalRequest.Marshal(b, m, deterministic)
}
func (m *CausalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CausalRequest.Merge(m, src)
}
func (m *CausalRequest) XXX_Size() int {
	return xxx_messageInfo_CausalRequest.Size(m)
}
func (m *CausalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CausalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CausalRequest proto.InternalMessageInfo

func (m *CausalRequest) GetConsistency() ConsistencyType {
	if m != nil {
		return m.Consistency
	}
	return ConsistencyType_NORMAL
}

func (m *CausalRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CausalRequest) GetTuples() []*CausalTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

func (m *CausalRequest) GetFutureReadSet() []string {
	if m != nil {
		return m.FutureReadSet
	}
	return nil
}

func (m *CausalRequest) GetKeyVersionLocations() map[string]*KeyVersionList {
	if m != nil {
		return m.KeyVersionLocations
	}
	return nil
}

func (m *CausalRequest) GetResponseAddress() string {
	if m != nil {
		return m.ResponseAddress
	}
	return ""
}

// A response to a particular CausalRequest.
type CausalResponse struct {
	// The CausalTuples that respond to an individual request. These form a
	// 1-to-1 mapping with the CausalTuples in the corresponding CausalRequest.
	Tuples []*CausalTuple `protobuf:"bytes,1,rep,name=tuples,proto3" json:"tuples,omitempty"`
	// The address of the cache that serviced this request -- this used to
	// propagate causal metadata to downstream executors.
	KeyVersionQueryAddress string `protobuf:"bytes,2,opt,name=key_version_query_address,json=keyVersionQueryAddress,proto3" json:"key_version_query_address,omitempty"`
	// The set of keys relevant to this DAG that are being cached at the above
	// address to service queries by downstream caches.
	KeyVersions          []*KeyVersion `protobuf:"bytes,3,rep,name=key_versions,json=keyVersions,proto3" json:"key_versions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CausalResponse) Reset()         { *m = CausalResponse{} }
func (m *CausalResponse) String() string { return proto.CompactTextString(m) }
func (*CausalResponse) ProtoMessage()    {}
func (*CausalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0056cb66ec996b89, []int{2}
}

func (m *CausalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CausalResponse.Unmarshal(m, b)
}
func (m *CausalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CausalResponse.Marshal(b, m, deterministic)
}
func (m *CausalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CausalResponse.Merge(m, src)
}
func (m *CausalResponse) XXX_Size() int {
	return xxx_messageInfo_CausalResponse.Size(m)
}
func (m *CausalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CausalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CausalResponse proto.InternalMessageInfo

func (m *CausalResponse) GetTuples() []*CausalTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

func (m *CausalResponse) GetKeyVersionQueryAddress() string {
	if m != nil {
		return m.KeyVersionQueryAddress
	}
	return ""
}

func (m *CausalResponse) GetKeyVersions() []*KeyVersion {
	if m != nil {
		return m.KeyVersions
	}
	return nil
}

// A request that is passed from one cache to another to retrieve a specific
// version of a particular key. The key versions are tracked by the request ID
// that they are associated with.
type KeyVersionRequest struct {
	// The request ID for which we are retrieving key versions.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The IP port pair that will receive the asynchronous response.
	ResponseAddress string `protobuf:"bytes,2,opt,name=response_address,json=responseAddress,proto3" json:"response_address,omitempty"`
	// The list of keys that we are requesting.
	Keys                 []string `protobuf:"bytes,3,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyVersionRequest) Reset()         { *m = KeyVersionRequest{} }
func (m *KeyVersionRequest) String() string { return proto.CompactTextString(m) }
func (*KeyVersionRequest) ProtoMessage()    {}
func (*KeyVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0056cb66ec996b89, []int{3}
}

func (m *KeyVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyVersionRequest.Unmarshal(m, b)
}
func (m *KeyVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyVersionRequest.Marshal(b, m, deterministic)
}
func (m *KeyVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyVersionRequest.Merge(m, src)
}
func (m *KeyVersionRequest) XXX_Size() int {
	return xxx_messageInfo_KeyVersionRequest.Size(m)
}
func (m *KeyVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyVersionRequest proto.InternalMessageInfo

func (m *KeyVersionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *KeyVersionRequest) GetResponseAddress() string {
	if m != nil {
		return m.ResponseAddress
	}
	return ""
}

func (m *KeyVersionRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

// A response to a partticular KeyVersionRequest.
type KeyVersionResponse struct {
	// The ID of the request for which this response is inteded.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The list of CausalTuples that was requested. This is a 1-to-1 mapping with
	// the keys field in the KeyVersionRequest.
	Tuples               []*CausalTuple `protobuf:"bytes,2,rep,name=tuples,proto3" json:"tuples,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KeyVersionResponse) Reset()         { *m = KeyVersionResponse{} }
func (m *KeyVersionResponse) String() string { return proto.CompactTextString(m) }
func (*KeyVersionResponse) ProtoMessage()    {}
func (*KeyVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0056cb66ec996b89, []int{4}
}

func (m *KeyVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyVersionResponse.Unmarshal(m, b)
}
func (m *KeyVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyVersionResponse.Marshal(b, m, deterministic)
}
func (m *KeyVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyVersionResponse.Merge(m, src)
}
func (m *KeyVersionResponse) XXX_Size() int {
	return xxx_messageInfo_KeyVersionResponse.Size(m)
}
func (m *KeyVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KeyVersionResponse proto.InternalMessageInfo

func (m *KeyVersionResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *KeyVersionResponse) GetTuples() []*CausalTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

func init() {
	proto.RegisterType((*CausalTuple)(nil), "CausalTuple")
	proto.RegisterType((*CausalRequest)(nil), "CausalRequest")
	proto.RegisterMapType((map[string]*KeyVersionList)(nil), "CausalRequest.KeyVersionLocationsEntry")
	proto.RegisterType((*CausalResponse)(nil), "CausalResponse")
	proto.RegisterType((*KeyVersionRequest)(nil), "KeyVersionRequest")
	proto.RegisterType((*KeyVersionResponse)(nil), "KeyVersionResponse")
}

func init() {
	proto.RegisterFile("causal.proto", fileDescriptor_0056cb66ec996b89)
}

var fileDescriptor_0056cb66ec996b89 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xed, 0x26, 0x90, 0xb1, 0x9b, 0x84, 0x45, 0x20, 0xd3, 0x93, 0x65, 0xf1, 0x61, 0x2e,
	0x3e, 0x98, 0x0b, 0x70, 0x2b, 0x55, 0x2f, 0xc0, 0x85, 0xa5, 0x02, 0x89, 0x1e, 0xac, 0x8d, 0x3d,
	0x08, 0x2b, 0xee, 0xae, 0xbb, 0x1f, 0x95, 0x7c, 0xe5, 0x9f, 0xf0, 0x4f, 0x51, 0xbc, 0x76, 0xb3,
	0x40, 0xca, 0x6d, 0xf7, 0xbd, 0x9d, 0x79, 0xf3, 0xde, 0xd8, 0x10, 0x55, 0xcc, 0x28, 0xd6, 0xe6,
	0x9d, 0x14, 0x5a, 0x9c, 0x00, 0xe3, 0x9c, 0x8d, 0xe7, 0x75, 0xd5, 0x0a, 0x53, 0x6f, 0x8c, 0x54,
	0x7a, 0x44, 0x22, 0xf5, 0x83, 0x49, 0xac, 0xed, 0x2d, 0xbd, 0x84, 0xf0, 0x6c, 0xa8, 0xbd, 0x30,
	0x5d, 0x8b, 0x64, 0x0d, 0xc1, 0x16, 0xfb, 0xd8, 0x4b, 0xbc, 0x6c, 0x41, 0x77, 0x47, 0x12, 0xc3,
	0xbd, 0x8e, 0xf5, 0xad, 0x60, 0x75, 0xec, 0x27, 0x5e, 0x16, 0xd1, 0xe9, 0x4a, 0x12, 0x98, 0xa1,
	0x94, 0x42, 0xc6, 0x41, 0xe2, 0x65, 0xcb, 0x02, 0xf2, 0x53, 0xce, 0xd9, 0xf9, 0x0e, 0xa1, 0x96,
	0x48, 0x7f, 0x06, 0x70, 0x6c, 0xbb, 0x53, 0xbc, 0x36, 0xa8, 0x34, 0x29, 0x20, 0xac, 0x04, 0x57,
	0x8d, 0xd2, 0xc8, 0x2b, 0xab, 0xb3, 0x2c, 0xd6, 0xf9, 0xd9, 0x1e, 0xbb, 0xe8, 0x3b, 0xa4, 0xee,
	0x23, 0xb2, 0x04, 0xbf, 0xb1, 0xe2, 0x0b, 0xea, 0x37, 0x35, 0x79, 0x0a, 0x73, 0xbd, 0x1b, 0x56,
	0xc5, 0x41, 0x12, 0x64, 0x61, 0x11, 0xe5, 0x8e, 0x03, 0x3a, 0x72, 0xe4, 0x39, 0xac, 0xbe, 0x1b,
	0x6d, 0x24, 0x96, 0x12, 0x59, 0x5d, 0x2a, 0xd4, 0xf1, 0x51, 0x12, 0x64, 0x0b, 0x7a, 0x6c, 0x61,
	0x8a, 0xac, 0xfe, 0x8c, 0x9a, 0x5c, 0xc2, 0xa3, 0x2d, 0xf6, 0xe5, 0x0d, 0x4a, 0xd5, 0x08, 0x5e,
	0xb6, 0xa2, 0x62, 0xba, 0x11, 0x5c, 0xc5, 0xb3, 0xa1, 0xf9, 0x8b, 0xfc, 0x0f, 0x03, 0xf9, 0x07,
	0xec, 0xbf, 0xd8, 0xa7, 0x1f, 0xa7, 0x97, 0xe7, 0x5c, 0xcb, 0x9e, 0x3e, 0xdc, 0xfe, 0xcb, 0x90,
	0x97, 0xb0, 0x96, 0xa8, 0x3a, 0xc1, 0x15, 0x96, 0xac, 0xae, 0x25, 0x2a, 0x15, 0xcf, 0x07, 0x23,
	0xab, 0x09, 0x3f, 0xb5, 0xf0, 0xc9, 0x57, 0x88, 0xef, 0xea, 0x7d, 0x60, 0x2b, 0xcf, 0x60, 0x76,
	0xc3, 0x5a, 0x83, 0x43, 0x2c, 0x61, 0xb1, 0x72, 0xe7, 0x6a, 0x94, 0xa6, 0x96, 0x7d, 0xeb, 0xbf,
	0xf6, 0xd2, 0x5f, 0x1e, 0x2c, 0x27, 0x0f, 0x56, 0xd2, 0x49, 0xd0, 0xfb, 0x4f, 0x82, 0x6f, 0xe0,
	0x89, 0x9b, 0xcc, 0xb5, 0x41, 0xd9, 0xdf, 0xba, 0xb0, 0xeb, 0x78, 0xbc, 0x37, 0xfd, 0x69, 0x47,
	0x8f, 0x66, 0x48, 0x0e, 0x91, 0x53, 0x3a, 0x2d, 0x2a, 0x74, 0xa6, 0xa4, 0xe1, 0xbe, 0x54, 0xa5,
	0x1b, 0x78, 0xe0, 0x50, 0xe3, 0xb7, 0x62, 0xf7, 0xee, 0xdd, 0xee, 0xfd, 0x50, 0x98, 0xfe, 0xc1,
	0x30, 0x09, 0x81, 0xa3, 0x2d, 0xf6, 0x56, 0x77, 0x41, 0x87, 0x73, 0xfa, 0x1e, 0x88, 0xab, 0x31,
	0x46, 0xf1, 0xb7, 0xc8, 0x3e, 0x1a, 0xff, 0xee, 0x68, 0xde, 0xdd, 0xff, 0x36, 0xaf, 0xc4, 0xd5,
	0x95, 0xe0, 0x9b, 0xf9, 0xf0, 0x1b, 0xbd, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x56, 0x5d,
	0xf6, 0x82, 0x03, 0x00, 0x00,
}
