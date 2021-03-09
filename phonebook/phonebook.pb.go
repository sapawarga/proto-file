// Code generated by protoc-gen-go. DO NOT EDIT.
// source: phonebook.proto

package phonebook

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetListRequest struct {
	Search               string   `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	RegencyId            int64    `protobuf:"varint,2,opt,name=regency_id,json=regencyId,proto3" json:"regency_id,omitempty"`
	DistrictId           int64    `protobuf:"varint,3,opt,name=district_id,json=districtId,proto3" json:"district_id,omitempty"`
	VillageId            int64    `protobuf:"varint,4,opt,name=village_id,json=villageId,proto3" json:"village_id,omitempty"`
	Status               int64    `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	Limit                int64    `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int64    `protobuf:"varint,7,opt,name=page,proto3" json:"page,omitempty"`
	Longitude            string   `protobuf:"bytes,8,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude             string   `protobuf:"bytes,9,opt,name=latitude,proto3" json:"latitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListRequest) Reset()         { *m = GetListRequest{} }
func (m *GetListRequest) String() string { return proto.CompactTextString(m) }
func (*GetListRequest) ProtoMessage()    {}
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_34db5399df65ad55, []int{0}
}

func (m *GetListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListRequest.Unmarshal(m, b)
}
func (m *GetListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListRequest.Marshal(b, m, deterministic)
}
func (m *GetListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListRequest.Merge(m, src)
}
func (m *GetListRequest) XXX_Size() int {
	return xxx_messageInfo_GetListRequest.Size(m)
}
func (m *GetListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetListRequest proto.InternalMessageInfo

func (m *GetListRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

func (m *GetListRequest) GetRegencyId() int64 {
	if m != nil {
		return m.RegencyId
	}
	return 0
}

func (m *GetListRequest) GetDistrictId() int64 {
	if m != nil {
		return m.DistrictId
	}
	return 0
}

func (m *GetListRequest) GetVillageId() int64 {
	if m != nil {
		return m.VillageId
	}
	return 0
}

func (m *GetListRequest) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetListRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetListRequest) GetLongitude() string {
	if m != nil {
		return m.Longitude
	}
	return ""
}

func (m *GetListRequest) GetLatitude() string {
	if m != nil {
		return m.Latitude
	}
	return ""
}

type GetListResponse struct {
	Data                 []*PhoneBook `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Metadata             *Metadata    `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetListResponse) Reset()         { *m = GetListResponse{} }
func (m *GetListResponse) String() string { return proto.CompactTextString(m) }
func (*GetListResponse) ProtoMessage()    {}
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_34db5399df65ad55, []int{1}
}

func (m *GetListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListResponse.Unmarshal(m, b)
}
func (m *GetListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListResponse.Marshal(b, m, deterministic)
}
func (m *GetListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListResponse.Merge(m, src)
}
func (m *GetListResponse) XXX_Size() int {
	return xxx_messageInfo_GetListResponse.Size(m)
}
func (m *GetListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetListResponse proto.InternalMessageInfo

func (m *GetListResponse) GetData() []*PhoneBook {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetListResponse) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Metadata struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Total                int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_34db5399df65ad55, []int{2}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Metadata) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type PhoneBook struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PhoneNumbers         string   `protobuf:"bytes,2,opt,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Latitude             string   `protobuf:"bytes,6,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            string   `protobuf:"bytes,7,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Status               int64    `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	Category             string   `protobuf:"bytes,9,opt,name=category,proto3" json:"category,omitempty"`
	Distance             string   `protobuf:"bytes,10,opt,name=distance,proto3" json:"distance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhoneBook) Reset()         { *m = PhoneBook{} }
func (m *PhoneBook) String() string { return proto.CompactTextString(m) }
func (*PhoneBook) ProtoMessage()    {}
func (*PhoneBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_34db5399df65ad55, []int{3}
}

func (m *PhoneBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhoneBook.Unmarshal(m, b)
}
func (m *PhoneBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhoneBook.Marshal(b, m, deterministic)
}
func (m *PhoneBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhoneBook.Merge(m, src)
}
func (m *PhoneBook) XXX_Size() int {
	return xxx_messageInfo_PhoneBook.Size(m)
}
func (m *PhoneBook) XXX_DiscardUnknown() {
	xxx_messageInfo_PhoneBook.DiscardUnknown(m)
}

var xxx_messageInfo_PhoneBook proto.InternalMessageInfo

func (m *PhoneBook) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PhoneBook) GetPhoneNumbers() string {
	if m != nil {
		return m.PhoneNumbers
	}
	return ""
}

func (m *PhoneBook) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PhoneBook) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PhoneBook) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PhoneBook) GetLatitude() string {
	if m != nil {
		return m.Latitude
	}
	return ""
}

func (m *PhoneBook) GetLongitude() string {
	if m != nil {
		return m.Longitude
	}
	return ""
}

func (m *PhoneBook) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PhoneBook) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *PhoneBook) GetDistance() string {
	if m != nil {
		return m.Distance
	}
	return ""
}

type GetDetailRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDetailRequest) Reset()         { *m = GetDetailRequest{} }
func (m *GetDetailRequest) String() string { return proto.CompactTextString(m) }
func (*GetDetailRequest) ProtoMessage()    {}
func (*GetDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_34db5399df65ad55, []int{4}
}

func (m *GetDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDetailRequest.Unmarshal(m, b)
}
func (m *GetDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDetailRequest.Marshal(b, m, deterministic)
}
func (m *GetDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDetailRequest.Merge(m, src)
}
func (m *GetDetailRequest) XXX_Size() int {
	return xxx_messageInfo_GetDetailRequest.Size(m)
}
func (m *GetDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDetailRequest proto.InternalMessageInfo

func (m *GetDetailRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetDetailResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PhoneNumbers         string   `protobuf:"bytes,2,opt,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Latitude             string   `protobuf:"bytes,6,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            string   `protobuf:"bytes,7,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Status               int64    `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            string   `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CategoryId           int64    `protobuf:"varint,11,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	CategoryName         string   `protobuf:"bytes,12,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	Distance             string   `protobuf:"bytes,13,opt,name=distance,proto3" json:"distance,omitempty"`
	RegencyId            int64    `protobuf:"varint,14,opt,name=regency_id,json=regencyId,proto3" json:"regency_id,omitempty"`
	RegencyName          string   `protobuf:"bytes,15,opt,name=regency_name,json=regencyName,proto3" json:"regency_name,omitempty"`
	DistrictId           int64    `protobuf:"varint,16,opt,name=district_id,json=districtId,proto3" json:"district_id,omitempty"`
	DistrictName         string   `protobuf:"bytes,17,opt,name=district_name,json=districtName,proto3" json:"district_name,omitempty"`
	VillageId            int64    `protobuf:"varint,18,opt,name=village_id,json=villageId,proto3" json:"village_id,omitempty"`
	VillageName          string   `protobuf:"bytes,19,opt,name=village_name,json=villageName,proto3" json:"village_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDetailResponse) Reset()         { *m = GetDetailResponse{} }
func (m *GetDetailResponse) String() string { return proto.CompactTextString(m) }
func (*GetDetailResponse) ProtoMessage()    {}
func (*GetDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_34db5399df65ad55, []int{5}
}

func (m *GetDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDetailResponse.Unmarshal(m, b)
}
func (m *GetDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDetailResponse.Marshal(b, m, deterministic)
}
func (m *GetDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDetailResponse.Merge(m, src)
}
func (m *GetDetailResponse) XXX_Size() int {
	return xxx_messageInfo_GetDetailResponse.Size(m)
}
func (m *GetDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDetailResponse proto.InternalMessageInfo

func (m *GetDetailResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetDetailResponse) GetPhoneNumbers() string {
	if m != nil {
		return m.PhoneNumbers
	}
	return ""
}

func (m *GetDetailResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetDetailResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetDetailResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GetDetailResponse) GetLatitude() string {
	if m != nil {
		return m.Latitude
	}
	return ""
}

func (m *GetDetailResponse) GetLongitude() string {
	if m != nil {
		return m.Longitude
	}
	return ""
}

func (m *GetDetailResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetDetailResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *GetDetailResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *GetDetailResponse) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *GetDetailResponse) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *GetDetailResponse) GetDistance() string {
	if m != nil {
		return m.Distance
	}
	return ""
}

func (m *GetDetailResponse) GetRegencyId() int64 {
	if m != nil {
		return m.RegencyId
	}
	return 0
}

func (m *GetDetailResponse) GetRegencyName() string {
	if m != nil {
		return m.RegencyName
	}
	return ""
}

func (m *GetDetailResponse) GetDistrictId() int64 {
	if m != nil {
		return m.DistrictId
	}
	return 0
}

func (m *GetDetailResponse) GetDistrictName() string {
	if m != nil {
		return m.DistrictName
	}
	return ""
}

func (m *GetDetailResponse) GetVillageId() int64 {
	if m != nil {
		return m.VillageId
	}
	return 0
}

func (m *GetDetailResponse) GetVillageName() string {
	if m != nil {
		return m.VillageName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetListRequest)(nil), "phonebook.GetListRequest")
	proto.RegisterType((*GetListResponse)(nil), "phonebook.GetListResponse")
	proto.RegisterType((*Metadata)(nil), "phonebook.Metadata")
	proto.RegisterType((*PhoneBook)(nil), "phonebook.PhoneBook")
	proto.RegisterType((*GetDetailRequest)(nil), "phonebook.GetDetailRequest")
	proto.RegisterType((*GetDetailResponse)(nil), "phonebook.GetDetailResponse")
}

func init() { proto.RegisterFile("phonebook.proto", fileDescriptor_34db5399df65ad55) }

var fileDescriptor_34db5399df65ad55 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xba, 0xae, 0x8d, 0x4f, 0xbb, 0x3f, 0xaf, 0x42, 0xa1, 0x6c, 0xa2, 0xcb, 0x6e, 0x7a,
	0x35, 0xa4, 0xc2, 0x03, 0x30, 0x84, 0x18, 0x95, 0x60, 0x42, 0x79, 0x81, 0xca, 0x8d, 0xad, 0xce,
	0x5a, 0x1a, 0x87, 0xc4, 0x45, 0xda, 0x35, 0x6f, 0x81, 0x78, 0x2c, 0x1e, 0x08, 0xf9, 0xd8, 0x4e,
	0x93, 0xb0, 0x27, 0xe0, 0xce, 0xe7, 0xfb, 0x8e, 0x8f, 0xed, 0xef, 0x3b, 0x3e, 0x70, 0x52, 0x3c,
	0xa8, 0x5c, 0xac, 0x95, 0x7a, 0xbc, 0x29, 0x4a, 0xa5, 0x15, 0x25, 0x35, 0x10, 0xff, 0xec, 0xc1,
	0xf1, 0x9d, 0xd0, 0x5f, 0x64, 0xa5, 0x13, 0xf1, 0x7d, 0x27, 0x2a, 0x4d, 0x5f, 0xc0, 0xa0, 0x12,
	0xac, 0x4c, 0x1f, 0xa2, 0x60, 0x16, 0xcc, 0x49, 0xe2, 0x22, 0x7a, 0x09, 0x50, 0x8a, 0x8d, 0xc8,
	0xd3, 0xa7, 0x95, 0xe4, 0x51, 0x6f, 0x16, 0xcc, 0x0f, 0x12, 0xe2, 0x90, 0x25, 0xa7, 0xaf, 0x61,
	0xc4, 0x65, 0xa5, 0x4b, 0x99, 0x6a, 0xc3, 0x1f, 0x20, 0x0f, 0x1e, 0x5a, 0x72, 0xb3, 0xff, 0x87,
	0xcc, 0x32, 0xb6, 0x11, 0x86, 0xef, 0xdb, 0xfd, 0x0e, 0x59, 0x72, 0x3c, 0x56, 0x33, 0xbd, 0xab,
	0xa2, 0x43, 0xa4, 0x5c, 0x44, 0x27, 0x70, 0x98, 0xc9, 0xad, 0xd4, 0xd1, 0x00, 0x61, 0x1b, 0x50,
	0x0a, 0xfd, 0x82, 0x6d, 0x44, 0x34, 0x44, 0x10, 0xd7, 0xf4, 0x02, 0x48, 0xa6, 0xf2, 0x8d, 0xd4,
	0x3b, 0x2e, 0xa2, 0x10, 0xef, 0xbe, 0x07, 0xe8, 0x14, 0xc2, 0x8c, 0x69, 0x4b, 0x12, 0x24, 0xeb,
	0x38, 0xce, 0xe0, 0xa4, 0x16, 0xa1, 0x2a, 0x54, 0x5e, 0x09, 0x3a, 0x87, 0x3e, 0x67, 0x9a, 0x45,
	0xc1, 0xec, 0x60, 0x3e, 0x5a, 0x4c, 0x6e, 0xf6, 0x1a, 0x7e, 0x33, 0xab, 0x0f, 0x4a, 0x3d, 0x26,
	0x98, 0x41, 0xdf, 0x40, 0xb8, 0x15, 0x9a, 0x61, 0xb6, 0x51, 0x65, 0xb4, 0x38, 0x6f, 0x64, 0x7f,
	0x75, 0x54, 0x52, 0x27, 0xc5, 0xef, 0x20, 0xf4, 0x68, 0xfd, 0x8e, 0xa0, 0xf1, 0x8e, 0x09, 0x1c,
	0x6a, 0xa5, 0x59, 0xe6, 0x34, 0xb6, 0x41, 0xfc, 0xab, 0x07, 0xa4, 0x3e, 0x9a, 0x1e, 0x43, 0x4f,
	0x72, 0xb7, 0xab, 0x27, 0x39, 0xbd, 0x86, 0x23, 0x3c, 0x73, 0x95, 0xef, 0xb6, 0x6b, 0x51, 0x56,
	0xb8, 0x97, 0x24, 0x63, 0x04, 0xef, 0x2d, 0x46, 0x67, 0x30, 0xe2, 0xa2, 0x4a, 0x4b, 0x59, 0x68,
	0xa9, 0x72, 0xb4, 0x88, 0x24, 0x4d, 0xc8, 0x5c, 0x27, 0x67, 0x5b, 0x81, 0xee, 0x90, 0x04, 0xd7,
	0x34, 0x82, 0x21, 0xe3, 0xbc, 0x14, 0x95, 0x75, 0x86, 0x24, 0x3e, 0x6c, 0x49, 0x3a, 0x68, 0x4b,
	0xda, 0x36, 0x63, 0xd8, 0x35, 0x63, 0x6f, 0x76, 0xd8, 0x32, 0x7b, 0x0a, 0x61, 0xca, 0xb4, 0xd8,
	0xa8, 0xf2, 0xc9, 0x9b, 0xe4, 0x63, 0xc3, 0x99, 0x6e, 0x62, 0x79, 0x2a, 0x22, 0xb0, 0x9c, 0x8f,
	0xe3, 0x18, 0x4e, 0xef, 0x84, 0xfe, 0x28, 0x34, 0x93, 0x99, 0xef, 0xe3, 0x8e, 0x44, 0xf1, 0x9f,
	0x3e, 0x9c, 0x35, 0x92, 0x9c, 0xcf, 0xff, 0xa9, 0x90, 0x97, 0x00, 0x69, 0x29, 0x98, 0x16, 0x7c,
	0xc5, 0xb4, 0x93, 0x92, 0x38, 0xe4, 0x56, 0x1b, 0x7a, 0x57, 0x70, 0x4f, 0x5b, 0x35, 0x89, 0x43,
	0x6e, 0xb5, 0xf9, 0xcb, 0x5e, 0x76, 0xf3, 0x57, 0x47, 0xf6, 0x2f, 0x7b, 0x68, 0x89, 0x2a, 0xd5,
	0x09, 0xf8, 0xce, 0xb1, 0x55, 0xc9, 0x83, 0xf7, 0xe6, 0xbd, 0x4d, 0xc3, 0x8e, 0xda, 0x86, 0x75,
	0x86, 0xc9, 0x71, 0x77, 0x98, 0x5c, 0xc1, 0xd8, 0xd3, 0x58, 0xfe, 0xc4, 0x2a, 0xec, 0x30, 0xac,
	0xde, 0x99, 0x37, 0xa7, 0xff, 0xcc, 0x9b, 0x6b, 0x38, 0xaa, 0x13, 0xb0, 0xc8, 0x99, 0xbd, 0xa3,
	0x07, 0xb1, 0x4a, 0x7b, 0x28, 0xd1, 0xee, 0x50, 0xba, 0x82, 0xb1, 0xa7, 0xb1, 0xc4, 0xb9, 0xbd,
	0x87, 0xc3, 0x4c, 0x85, 0xc5, 0xef, 0x00, 0x4e, 0xeb, 0x7f, 0xf9, 0x99, 0xe5, 0x3c, 0x13, 0x25,
	0x7d, 0x0f, 0x43, 0x37, 0x50, 0xe8, 0xcb, 0xc6, 0x30, 0x68, 0x4f, 0xda, 0xe9, 0xf4, 0x39, 0xca,
	0xf5, 0xe5, 0x27, 0x20, 0x75, 0xb3, 0xd2, 0x57, 0xed, 0xc4, 0x56, 0x9f, 0x4f, 0x2f, 0x9e, 0x27,
	0x6d, 0x9d, 0xf5, 0x00, 0x47, 0xfe, 0xdb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x91, 0xa1, 0xbb,
	0x9e, 0x05, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PhoneBookHandlerClient is the client API for PhoneBookHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PhoneBookHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc.CallOption) (*GetDetailResponse, error)
}

type phoneBookHandlerClient struct {
	cc *grpc.ClientConn
}

func NewPhoneBookHandlerClient(cc *grpc.ClientConn) PhoneBookHandlerClient {
	return &phoneBookHandlerClient{cc}
}

func (c *phoneBookHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/phonebook.PhoneBookHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phoneBookHandlerClient) GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc.CallOption) (*GetDetailResponse, error) {
	out := new(GetDetailResponse)
	err := c.cc.Invoke(ctx, "/phonebook.PhoneBookHandler/GetDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhoneBookHandlerServer is the server API for PhoneBookHandler service.
type PhoneBookHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	GetDetail(context.Context, *GetDetailRequest) (*GetDetailResponse, error)
}

// UnimplementedPhoneBookHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedPhoneBookHandlerServer struct {
}

func (*UnimplementedPhoneBookHandlerServer) GetList(ctx context.Context, req *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (*UnimplementedPhoneBookHandlerServer) GetDetail(ctx context.Context, req *GetDetailRequest) (*GetDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetail not implemented")
}

func RegisterPhoneBookHandlerServer(s *grpc.Server, srv PhoneBookHandlerServer) {
	s.RegisterService(&_PhoneBookHandler_serviceDesc, srv)
}

func _PhoneBookHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneBookHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phonebook.PhoneBookHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneBookHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhoneBookHandler_GetDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneBookHandlerServer).GetDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phonebook.PhoneBookHandler/GetDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneBookHandlerServer).GetDetail(ctx, req.(*GetDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PhoneBookHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "phonebook.PhoneBookHandler",
	HandlerType: (*PhoneBookHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _PhoneBookHandler_GetList_Handler,
		},
		{
			MethodName: "GetDetail",
			Handler:    _PhoneBookHandler_GetDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "phonebook.proto",
}