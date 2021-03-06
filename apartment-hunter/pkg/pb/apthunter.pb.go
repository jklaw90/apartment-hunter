// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apthunter.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type CreateOrUpdateRequest struct {
	Id                   string   `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,20,opt,name=address,proto3" json:"address,omitempty"`
	Url                  string   `protobuf:"bytes,30,opt,name=url,proto3" json:"url,omitempty"`
	Title                string   `protobuf:"bytes,40,opt,name=title,proto3" json:"title,omitempty"`
	Price                float32  `protobuf:"fixed32,50,opt,name=price,proto3" json:"price,omitempty"`
	Bedrooms             uint64   `protobuf:"varint,60,opt,name=bedrooms,proto3" json:"bedrooms,omitempty"`
	Bathrooms            float32  `protobuf:"fixed32,70,opt,name=bathrooms,proto3" json:"bathrooms,omitempty"`
	Sqft                 float32  `protobuf:"fixed32,80,opt,name=sqft,proto3" json:"sqft,omitempty"`
	AvailableDate        string   `protobuf:"bytes,90,opt,name=available_date,json=availableDate,proto3" json:"available_date,omitempty"`
	Cats                 bool     `protobuf:"varint,100,opt,name=cats,proto3" json:"cats,omitempty"`
	Dogs                 bool     `protobuf:"varint,110,opt,name=dogs,proto3" json:"dogs,omitempty"`
	HousingType          string   `protobuf:"bytes,120,opt,name=housing_type,json=housingType,proto3" json:"housing_type,omitempty"`
	WdType               string   `protobuf:"bytes,130,opt,name=wd_type,json=wdType,proto3" json:"wd_type,omitempty"`
	ParkingType          string   `protobuf:"bytes,140,opt,name=parking_type,json=parkingType,proto3" json:"parking_type,omitempty"`
	Images               []string `protobuf:"bytes,150,rep,name=images,proto3" json:"images,omitempty"`
	Body                 string   `protobuf:"bytes,160,opt,name=body,proto3" json:"body,omitempty"`
	Lng                  float32  `protobuf:"fixed32,170,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat                  float32  `protobuf:"fixed32,180,opt,name=lat,proto3" json:"lat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrUpdateRequest) Reset()         { *m = CreateOrUpdateRequest{} }
func (m *CreateOrUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrUpdateRequest) ProtoMessage()    {}
func (*CreateOrUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a2b52ad04f299cb, []int{0}
}

func (m *CreateOrUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrUpdateRequest.Unmarshal(m, b)
}
func (m *CreateOrUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrUpdateRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrUpdateRequest.Merge(m, src)
}
func (m *CreateOrUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrUpdateRequest.Size(m)
}
func (m *CreateOrUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrUpdateRequest proto.InternalMessageInfo

func (m *CreateOrUpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateOrUpdateRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CreateOrUpdateRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CreateOrUpdateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateOrUpdateRequest) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CreateOrUpdateRequest) GetBedrooms() uint64 {
	if m != nil {
		return m.Bedrooms
	}
	return 0
}

func (m *CreateOrUpdateRequest) GetBathrooms() float32 {
	if m != nil {
		return m.Bathrooms
	}
	return 0
}

func (m *CreateOrUpdateRequest) GetSqft() float32 {
	if m != nil {
		return m.Sqft
	}
	return 0
}

func (m *CreateOrUpdateRequest) GetAvailableDate() string {
	if m != nil {
		return m.AvailableDate
	}
	return ""
}

func (m *CreateOrUpdateRequest) GetCats() bool {
	if m != nil {
		return m.Cats
	}
	return false
}

func (m *CreateOrUpdateRequest) GetDogs() bool {
	if m != nil {
		return m.Dogs
	}
	return false
}

func (m *CreateOrUpdateRequest) GetHousingType() string {
	if m != nil {
		return m.HousingType
	}
	return ""
}

func (m *CreateOrUpdateRequest) GetWdType() string {
	if m != nil {
		return m.WdType
	}
	return ""
}

func (m *CreateOrUpdateRequest) GetParkingType() string {
	if m != nil {
		return m.ParkingType
	}
	return ""
}

func (m *CreateOrUpdateRequest) GetImages() []string {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *CreateOrUpdateRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *CreateOrUpdateRequest) GetLng() float32 {
	if m != nil {
		return m.Lng
	}
	return 0
}

func (m *CreateOrUpdateRequest) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

type CreateOrUpdateResponse struct {
	Id                   string   `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrUpdateResponse) Reset()         { *m = CreateOrUpdateResponse{} }
func (m *CreateOrUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOrUpdateResponse) ProtoMessage()    {}
func (*CreateOrUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a2b52ad04f299cb, []int{1}
}

func (m *CreateOrUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrUpdateResponse.Unmarshal(m, b)
}
func (m *CreateOrUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrUpdateResponse.Marshal(b, m, deterministic)
}
func (m *CreateOrUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrUpdateResponse.Merge(m, src)
}
func (m *CreateOrUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOrUpdateResponse.Size(m)
}
func (m *CreateOrUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrUpdateResponse proto.InternalMessageInfo

func (m *CreateOrUpdateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StreamRequest struct {
	Id                   string   `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	LastSeen             string   `protobuf:"bytes,20,opt,name=lastSeen,proto3" json:"lastSeen,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a2b52ad04f299cb, []int{2}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StreamRequest) GetLastSeen() string {
	if m != nil {
		return m.LastSeen
	}
	return ""
}

type AggregateEvent struct {
	Id                   string   `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	AggregateId          string   `protobuf:"bytes,20,opt,name=aggregateId,proto3" json:"aggregateId,omitempty"`
	Timestamp            int64    `protobuf:"varint,30,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Version              uint32   `protobuf:"varint,40,opt,name=version,proto3" json:"version,omitempty"`
	EventType            string   `protobuf:"bytes,50,opt,name=eventType,proto3" json:"eventType,omitempty"`
	Event                []byte   `protobuf:"bytes,60,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AggregateEvent) Reset()         { *m = AggregateEvent{} }
func (m *AggregateEvent) String() string { return proto.CompactTextString(m) }
func (*AggregateEvent) ProtoMessage()    {}
func (*AggregateEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a2b52ad04f299cb, []int{3}
}

func (m *AggregateEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateEvent.Unmarshal(m, b)
}
func (m *AggregateEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateEvent.Marshal(b, m, deterministic)
}
func (m *AggregateEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateEvent.Merge(m, src)
}
func (m *AggregateEvent) XXX_Size() int {
	return xxx_messageInfo_AggregateEvent.Size(m)
}
func (m *AggregateEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateEvent proto.InternalMessageInfo

func (m *AggregateEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AggregateEvent) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

func (m *AggregateEvent) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *AggregateEvent) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *AggregateEvent) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *AggregateEvent) GetEvent() []byte {
	if m != nil {
		return m.Event
	}
	return nil
}

type EndedRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndedRequest) Reset()         { *m = EndedRequest{} }
func (m *EndedRequest) String() string { return proto.CompactTextString(m) }
func (*EndedRequest) ProtoMessage()    {}
func (*EndedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a2b52ad04f299cb, []int{4}
}

func (m *EndedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndedRequest.Unmarshal(m, b)
}
func (m *EndedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndedRequest.Marshal(b, m, deterministic)
}
func (m *EndedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndedRequest.Merge(m, src)
}
func (m *EndedRequest) XXX_Size() int {
	return xxx_messageInfo_EndedRequest.Size(m)
}
func (m *EndedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndedRequest proto.InternalMessageInfo

type EndedResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndedResponse) Reset()         { *m = EndedResponse{} }
func (m *EndedResponse) String() string { return proto.CompactTextString(m) }
func (*EndedResponse) ProtoMessage()    {}
func (*EndedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a2b52ad04f299cb, []int{5}
}

func (m *EndedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndedResponse.Unmarshal(m, b)
}
func (m *EndedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndedResponse.Marshal(b, m, deterministic)
}
func (m *EndedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndedResponse.Merge(m, src)
}
func (m *EndedResponse) XXX_Size() int {
	return xxx_messageInfo_EndedResponse.Size(m)
}
func (m *EndedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EndedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EndedResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateOrUpdateRequest)(nil), "pb.CreateOrUpdateRequest")
	proto.RegisterType((*CreateOrUpdateResponse)(nil), "pb.CreateOrUpdateResponse")
	proto.RegisterType((*StreamRequest)(nil), "pb.StreamRequest")
	proto.RegisterType((*AggregateEvent)(nil), "pb.AggregateEvent")
	proto.RegisterType((*EndedRequest)(nil), "pb.EndedRequest")
	proto.RegisterType((*EndedResponse)(nil), "pb.EndedResponse")
}

func init() { proto.RegisterFile("apthunter.proto", fileDescriptor_3a2b52ad04f299cb) }

var fileDescriptor_3a2b52ad04f299cb = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x3f, 0x6f, 0x13, 0x41,
	0x10, 0xc5, 0x73, 0xb6, 0x31, 0xf6, 0xf8, 0x4f, 0xc8, 0x12, 0x60, 0xb1, 0x10, 0x3a, 0x4e, 0x42,
	0xba, 0xca, 0x42, 0x49, 0x09, 0x0d, 0x82, 0x20, 0xa5, 0x02, 0x5d, 0x40, 0x48, 0x34, 0xd1, 0x5e,
	0x76, 0x38, 0x9f, 0xb8, 0x3f, 0x9b, 0xdd, 0xb5, 0x83, 0x5b, 0x2a, 0x0a, 0x6a, 0x6a, 0x6a, 0x0a,
	0x2a, 0xbe, 0x08, 0xdf, 0x08, 0xed, 0xec, 0xd9, 0xc6, 0x21, 0xe9, 0xe6, 0xfd, 0x66, 0xdf, 0xda,
	0xfb, 0x66, 0x0e, 0x76, 0x85, 0xb2, 0xb3, 0x79, 0x65, 0x51, 0x4f, 0x95, 0xae, 0x6d, 0xcd, 0x5a,
	0x2a, 0x8d, 0xfe, 0xb4, 0xe1, 0xce, 0x0b, 0x8d, 0xc2, 0xe2, 0x6b, 0xfd, 0x4e, 0x49, 0x61, 0x31,
	0xc1, 0xf3, 0x39, 0x1a, 0xcb, 0xc6, 0xd0, 0xca, 0x25, 0x87, 0x30, 0x88, 0xfb, 0x49, 0x2b, 0x97,
	0x8c, 0xc3, 0x4d, 0x21, 0xa5, 0x46, 0x63, 0xf8, 0x3e, 0xc1, 0x95, 0x64, 0xb7, 0xa0, 0x3d, 0xd7,
	0x05, 0x7f, 0x48, 0xd4, 0x95, 0x6c, 0x1f, 0x6e, 0xd8, 0xdc, 0x16, 0xc8, 0x63, 0x62, 0x5e, 0x38,
	0xaa, 0x74, 0x7e, 0x86, 0xfc, 0x20, 0x0c, 0xe2, 0x56, 0xe2, 0x05, 0x9b, 0x40, 0x2f, 0x45, 0xa9,
	0xeb, 0xba, 0x34, 0xfc, 0x59, 0x18, 0xc4, 0x9d, 0x64, 0xad, 0xd9, 0x03, 0xe8, 0xa7, 0xc2, 0xce,
	0x7c, 0xf3, 0x15, 0xb9, 0x36, 0x80, 0x31, 0xe8, 0x98, 0xf3, 0x8f, 0x96, 0xbf, 0xa1, 0x06, 0xd5,
	0xec, 0x31, 0x8c, 0xc5, 0x42, 0xe4, 0x85, 0x48, 0x0b, 0x3c, 0x75, 0xcf, 0xe1, 0x1f, 0xe8, 0x2f,
	0x8c, 0xd6, 0xf4, 0xa5, 0xb0, 0xe8, 0xac, 0x67, 0xc2, 0x1a, 0x2e, 0xc3, 0x20, 0xee, 0x25, 0x54,
	0x3b, 0x26, 0xeb, 0xcc, 0xf0, 0xca, 0x33, 0x57, 0xb3, 0x47, 0x30, 0x9c, 0xd5, 0x73, 0x93, 0x57,
	0xd9, 0xa9, 0x5d, 0x2a, 0xe4, 0x9f, 0xe9, 0xb2, 0x41, 0xc3, 0xde, 0x2e, 0x15, 0xba, 0x5c, 0x2e,
	0xa4, 0xef, 0x7e, 0x09, 0xa8, 0xdd, 0xbd, 0x90, 0xd4, 0x89, 0x60, 0xa8, 0x84, 0xfe, 0xb4, 0x36,
	0x7f, 0xf3, 0xed, 0x41, 0x03, 0xe9, 0xcc, 0x3d, 0xe8, 0xe6, 0xa5, 0xc8, 0xd0, 0xf0, 0xef, 0x41,
	0xd8, 0x76, 0x66, 0x2f, 0xd9, 0x6d, 0xe8, 0xa4, 0xb5, 0x5c, 0xf2, 0x1f, 0xde, 0x44, 0x82, 0xed,
	0x41, 0xbb, 0xa8, 0x32, 0xfe, 0x33, 0xa0, 0x17, 0xbb, 0x9a, 0x90, 0xb0, 0xfc, 0xf7, 0x0a, 0x09,
	0x1b, 0xc5, 0x70, 0xf7, 0xf2, 0x48, 0x8d, 0xaa, 0x2b, 0x83, 0x97, 0x67, 0x1a, 0x3d, 0x85, 0xd1,
	0x89, 0xd5, 0x28, 0xca, 0xeb, 0x86, 0x3e, 0x81, 0x5e, 0x21, 0x8c, 0x3d, 0x41, 0xac, 0x9a, 0xa9,
	0xaf, 0x75, 0xf4, 0x2b, 0x80, 0xf1, 0xf3, 0x2c, 0xd3, 0x98, 0x09, 0x8b, 0x47, 0x0b, 0xac, 0xfe,
	0xb7, 0x87, 0x30, 0x10, 0xab, 0x13, 0xc7, 0xb2, 0xb9, 0xe1, 0x5f, 0xe4, 0x26, 0x6c, 0xf3, 0x12,
	0x8d, 0x15, 0xa5, 0xa2, 0x0d, 0x6a, 0x27, 0x1b, 0xe0, 0xb2, 0x5d, 0xa0, 0x36, 0x79, 0x5d, 0xd1,
	0x26, 0x8d, 0x92, 0x95, 0x74, 0x3e, 0x74, 0x3f, 0xe9, 0x42, 0xa4, 0x7d, 0xea, 0x27, 0x1b, 0xe0,
	0x36, 0x8d, 0x04, 0x2d, 0xd4, 0x30, 0xf1, 0x22, 0x1a, 0xc3, 0xf0, 0xa8, 0x92, 0x28, 0x9b, 0xc7,
	0x46, 0xbb, 0x30, 0x6a, 0xb4, 0x8f, 0xe7, 0xe0, 0x6b, 0x00, 0xdd, 0xf7, 0x3a, 0xb7, 0xa8, 0xd9,
	0x31, 0x8c, 0xb7, 0x33, 0x64, 0xf7, 0xa7, 0x2a, 0x9d, 0x5e, 0xf9, 0xa9, 0x4c, 0x26, 0x57, 0xb5,
	0xfc, 0x9d, 0xd1, 0x0e, 0x3b, 0x84, 0xae, 0x0f, 0x99, 0xed, 0xb9, 0x73, 0x5b, 0x81, 0x4f, 0x98,
	0x43, 0xdb, 0x29, 0x46, 0x3b, 0x4f, 0x82, 0xb4, 0x4b, 0x9f, 0xe8, 0xe1, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe4, 0xfb, 0x7e, 0x58, 0xb5, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WriterClient is the client API for Writer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WriterClient interface {
	CreateOrUpdate(ctx context.Context, in *CreateOrUpdateRequest, opts ...grpc.CallOption) (*CreateOrUpdateResponse, error)
	Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Writer_StreamClient, error)
}

type writerClient struct {
	cc *grpc.ClientConn
}

func NewWriterClient(cc *grpc.ClientConn) WriterClient {
	return &writerClient{cc}
}

func (c *writerClient) CreateOrUpdate(ctx context.Context, in *CreateOrUpdateRequest, opts ...grpc.CallOption) (*CreateOrUpdateResponse, error) {
	out := new(CreateOrUpdateResponse)
	err := c.cc.Invoke(ctx, "/pb.Writer/CreateOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerClient) Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Writer_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Writer_serviceDesc.Streams[0], "/pb.Writer/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &writerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Writer_StreamClient interface {
	Recv() (*AggregateEvent, error)
	grpc.ClientStream
}

type writerStreamClient struct {
	grpc.ClientStream
}

func (x *writerStreamClient) Recv() (*AggregateEvent, error) {
	m := new(AggregateEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WriterServer is the server API for Writer service.
type WriterServer interface {
	CreateOrUpdate(context.Context, *CreateOrUpdateRequest) (*CreateOrUpdateResponse, error)
	Stream(*StreamRequest, Writer_StreamServer) error
}

func RegisterWriterServer(s *grpc.Server, srv WriterServer) {
	s.RegisterService(&_Writer_serviceDesc, srv)
}

func _Writer_CreateOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterServer).CreateOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Writer/CreateOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterServer).CreateOrUpdate(ctx, req.(*CreateOrUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Writer_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WriterServer).Stream(m, &writerStreamServer{stream})
}

type Writer_StreamServer interface {
	Send(*AggregateEvent) error
	grpc.ServerStream
}

type writerStreamServer struct {
	grpc.ServerStream
}

func (x *writerStreamServer) Send(m *AggregateEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _Writer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Writer",
	HandlerType: (*WriterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdate",
			Handler:    _Writer_CreateOrUpdate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Writer_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "apthunter.proto",
}
