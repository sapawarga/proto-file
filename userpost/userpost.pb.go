// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userpost.proto

package userpost

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

type ByID struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByID) Reset()         { *m = ByID{} }
func (m *ByID) String() string { return proto.CompactTextString(m) }
func (*ByID) ProtoMessage()    {}
func (*ByID) Descriptor() ([]byte, []int) {
	return fileDescriptor_44fc5843c1f19ec4, []int{0}
}

func (m *ByID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByID.Unmarshal(m, b)
}
func (m *ByID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByID.Marshal(b, m, deterministic)
}
func (m *ByID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByID.Merge(m, src)
}
func (m *ByID) XXX_Size() int {
	return xxx_messageInfo_ByID.Size(m)
}
func (m *ByID) XXX_DiscardUnknown() {
	xxx_messageInfo_ByID.DiscardUnknown(m)
}

var xxx_messageInfo_ByID proto.InternalMessageInfo

func (m *ByID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetListUserPostRequest struct {
	ActivityName         string   `protobuf:"bytes,1,opt,name=activity_name,json=activityName,proto3" json:"activity_name,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Category             string   `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Status               int64    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Page                 int64    `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int64    `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	SortBy               string   `protobuf:"bytes,7,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	OrderBy              string   `protobuf:"bytes,8,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListUserPostRequest) Reset()         { *m = GetListUserPostRequest{} }
func (m *GetListUserPostRequest) String() string { return proto.CompactTextString(m) }
func (*GetListUserPostRequest) ProtoMessage()    {}
func (*GetListUserPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44fc5843c1f19ec4, []int{1}
}

func (m *GetListUserPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListUserPostRequest.Unmarshal(m, b)
}
func (m *GetListUserPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListUserPostRequest.Marshal(b, m, deterministic)
}
func (m *GetListUserPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListUserPostRequest.Merge(m, src)
}
func (m *GetListUserPostRequest) XXX_Size() int {
	return xxx_messageInfo_GetListUserPostRequest.Size(m)
}
func (m *GetListUserPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListUserPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetListUserPostRequest proto.InternalMessageInfo

func (m *GetListUserPostRequest) GetActivityName() string {
	if m != nil {
		return m.ActivityName
	}
	return ""
}

func (m *GetListUserPostRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetListUserPostRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *GetListUserPostRequest) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetListUserPostRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetListUserPostRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetListUserPostRequest) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}

func (m *GetListUserPostRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

type GetListUserPostResponse struct {
	Data                 []*UserPost `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Metadata             *Metadata   `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetListUserPostResponse) Reset()         { *m = GetListUserPostResponse{} }
func (m *GetListUserPostResponse) String() string { return proto.CompactTextString(m) }
func (*GetListUserPostResponse) ProtoMessage()    {}
func (*GetListUserPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_44fc5843c1f19ec4, []int{2}
}

func (m *GetListUserPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListUserPostResponse.Unmarshal(m, b)
}
func (m *GetListUserPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListUserPostResponse.Marshal(b, m, deterministic)
}
func (m *GetListUserPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListUserPostResponse.Merge(m, src)
}
func (m *GetListUserPostResponse) XXX_Size() int {
	return xxx_messageInfo_GetListUserPostResponse.Size(m)
}
func (m *GetListUserPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListUserPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetListUserPostResponse proto.InternalMessageInfo

func (m *GetListUserPostResponse) GetData() []*UserPost {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetListUserPostResponse) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Metadata struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Total                int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	TotalPage            int64    `protobuf:"varint,3,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_44fc5843c1f19ec4, []int{3}
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

func (m *Metadata) GetTotalPage() int64 {
	if m != nil {
		return m.TotalPage
	}
	return 0
}

type UserPost struct {
	Id                    int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                 string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Tag                   string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	ImagePath             string   `protobuf:"bytes,4,opt,name=image_path,json=imagePath,proto3" json:"image_path,omitempty"`
	Images                string   `protobuf:"bytes,5,opt,name=images,proto3" json:"images,omitempty"`
	LastUserPostCommentId int64    `protobuf:"varint,6,opt,name=last_user_post_comment_id,json=lastUserPostCommentId,proto3" json:"last_user_post_comment_id,omitempty"`
	LastComment           *Comment `protobuf:"bytes,7,opt,name=last_comment,json=lastComment,proto3" json:"last_comment,omitempty"`
	LikesCount            int64    `protobuf:"varint,8,opt,name=likes_count,json=likesCount,proto3" json:"likes_count,omitempty"`
	IsLiked               bool     `protobuf:"varint,9,opt,name=is_liked,json=isLiked,proto3" json:"is_liked,omitempty"`
	CommentCounts         int64    `protobuf:"varint,10,opt,name=comment_counts,json=commentCounts,proto3" json:"comment_counts,omitempty"`
	Status                int64    `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
	Actor                 *Actor   `protobuf:"bytes,12,opt,name=actor,proto3" json:"actor,omitempty"`
	CreatedAt             string   `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt             string   `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *UserPost) Reset()         { *m = UserPost{} }
func (m *UserPost) String() string { return proto.CompactTextString(m) }
func (*UserPost) ProtoMessage()    {}
func (*UserPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_44fc5843c1f19ec4, []int{4}
}

func (m *UserPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPost.Unmarshal(m, b)
}
func (m *UserPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPost.Marshal(b, m, deterministic)
}
func (m *UserPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPost.Merge(m, src)
}
func (m *UserPost) XXX_Size() int {
	return xxx_messageInfo_UserPost.Size(m)
}
func (m *UserPost) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPost.DiscardUnknown(m)
}

var xxx_messageInfo_UserPost proto.InternalMessageInfo

func (m *UserPost) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserPost) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UserPost) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *UserPost) GetImagePath() string {
	if m != nil {
		return m.ImagePath
	}
	return ""
}

func (m *UserPost) GetImages() string {
	if m != nil {
		return m.Images
	}
	return ""
}

func (m *UserPost) GetLastUserPostCommentId() int64 {
	if m != nil {
		return m.LastUserPostCommentId
	}
	return 0
}

func (m *UserPost) GetLastComment() *Comment {
	if m != nil {
		return m.LastComment
	}
	return nil
}

func (m *UserPost) GetLikesCount() int64 {
	if m != nil {
		return m.LikesCount
	}
	return 0
}

func (m *UserPost) GetIsLiked() bool {
	if m != nil {
		return m.IsLiked
	}
	return false
}

func (m *UserPost) GetCommentCounts() int64 {
	if m != nil {
		return m.CommentCounts
	}
	return 0
}

func (m *UserPost) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UserPost) GetActor() *Actor {
	if m != nil {
		return m.Actor
	}
	return nil
}

func (m *UserPost) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UserPost) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Comment struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserPostId           int64    `protobuf:"varint,2,opt,name=user_post_id,json=userPostId,proto3" json:"user_post_id,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	CreatedAt            string   `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedBy            *Actor   `protobuf:"bytes,6,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedBy            *Actor   `protobuf:"bytes,7,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_44fc5843c1f19ec4, []int{5}
}

func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Comment) GetUserPostId() int64 {
	if m != nil {
		return m.UserPostId
	}
	return 0
}

func (m *Comment) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Comment) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Comment) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Comment) GetCreatedBy() *Actor {
	if m != nil {
		return m.CreatedBy
	}
	return nil
}

func (m *Comment) GetUpdatedBy() *Actor {
	if m != nil {
		return m.UpdatedBy
	}
	return nil
}

type Actor struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhotoUrl             string   `protobuf:"bytes,3,opt,name=photo_url,json=photoUrl,proto3" json:"photo_url,omitempty"`
	Role                 int64    `protobuf:"varint,4,opt,name=role,proto3" json:"role,omitempty"`
	Regency              string   `protobuf:"bytes,5,opt,name=regency,proto3" json:"regency,omitempty"`
	District             string   `protobuf:"bytes,6,opt,name=district,proto3" json:"district,omitempty"`
	Village              string   `protobuf:"bytes,7,opt,name=village,proto3" json:"village,omitempty"`
	Rw                   string   `protobuf:"bytes,8,opt,name=rw,proto3" json:"rw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Actor) Reset()         { *m = Actor{} }
func (m *Actor) String() string { return proto.CompactTextString(m) }
func (*Actor) ProtoMessage()    {}
func (*Actor) Descriptor() ([]byte, []int) {
	return fileDescriptor_44fc5843c1f19ec4, []int{6}
}

func (m *Actor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Actor.Unmarshal(m, b)
}
func (m *Actor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Actor.Marshal(b, m, deterministic)
}
func (m *Actor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Actor.Merge(m, src)
}
func (m *Actor) XXX_Size() int {
	return xxx_messageInfo_Actor.Size(m)
}
func (m *Actor) XXX_DiscardUnknown() {
	xxx_messageInfo_Actor.DiscardUnknown(m)
}

var xxx_messageInfo_Actor proto.InternalMessageInfo

func (m *Actor) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Actor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Actor) GetPhotoUrl() string {
	if m != nil {
		return m.PhotoUrl
	}
	return ""
}

func (m *Actor) GetRole() int64 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *Actor) GetRegency() string {
	if m != nil {
		return m.Regency
	}
	return ""
}

func (m *Actor) GetDistrict() string {
	if m != nil {
		return m.District
	}
	return ""
}

func (m *Actor) GetVillage() string {
	if m != nil {
		return m.Village
	}
	return ""
}

func (m *Actor) GetRw() string {
	if m != nil {
		return m.Rw
	}
	return ""
}

func init() {
	proto.RegisterType((*ByID)(nil), "userpost.ByID")
	proto.RegisterType((*GetListUserPostRequest)(nil), "userpost.GetListUserPostRequest")
	proto.RegisterType((*GetListUserPostResponse)(nil), "userpost.GetListUserPostResponse")
	proto.RegisterType((*Metadata)(nil), "userpost.Metadata")
	proto.RegisterType((*UserPost)(nil), "userpost.UserPost")
	proto.RegisterType((*Comment)(nil), "userpost.Comment")
	proto.RegisterType((*Actor)(nil), "userpost.Actor")
}

func init() { proto.RegisterFile("userpost.proto", fileDescriptor_44fc5843c1f19ec4) }

var fileDescriptor_44fc5843c1f19ec4 = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdd, 0x4e, 0xdb, 0x4a,
	0x10, 0x96, 0xe3, 0xfc, 0x79, 0x12, 0xc2, 0x61, 0x75, 0x0e, 0x18, 0x8e, 0x8e, 0x4e, 0x9a, 0x8a,
	0x8a, 0xab, 0x5c, 0xd0, 0x4a, 0xed, 0x2d, 0x01, 0x89, 0x22, 0xd1, 0x0a, 0xb9, 0xa2, 0xb7, 0xd6,
	0x62, 0xaf, 0xc2, 0xaa, 0x4e, 0x36, 0xec, 0x4e, 0x40, 0x7e, 0x9d, 0x3e, 0x47, 0x5f, 0xa0, 0xcf,
	0xd2, 0xcb, 0xbe, 0x40, 0xb5, 0xe3, 0xdd, 0x18, 0x42, 0xc4, 0xdd, 0x7e, 0xdf, 0x37, 0xb3, 0x99,
	0x99, 0x6f, 0x3d, 0x81, 0xc1, 0xd2, 0x08, 0xbd, 0x50, 0x06, 0xc7, 0x0b, 0xad, 0x50, 0xb1, 0xae,
	0xc7, 0xa3, 0x5d, 0x68, 0x4e, 0xca, 0x8b, 0x33, 0x36, 0x80, 0x86, 0xcc, 0xe3, 0x60, 0x18, 0x1c,
	0x85, 0x49, 0x43, 0xe6, 0xa3, 0x5f, 0x01, 0xec, 0x9e, 0x0b, 0xbc, 0x94, 0x06, 0xaf, 0x8d, 0xd0,
	0x57, 0xca, 0x60, 0x22, 0xee, 0x96, 0xc2, 0x20, 0x7b, 0x0d, 0x5b, 0x3c, 0x43, 0x79, 0x2f, 0xb1,
	0x4c, 0xe7, 0x7c, 0x26, 0x28, 0x2b, 0x4a, 0xfa, 0x9e, 0xfc, 0xcc, 0x67, 0x82, 0x1d, 0x00, 0xfd,
	0x06, 0xe9, 0x0d, 0xd2, 0x57, 0xd8, 0x6a, 0x19, 0x47, 0x31, 0x55, 0xba, 0x8c, 0xc3, 0x4a, 0xf3,
	0x98, 0xed, 0x42, 0xdb, 0x20, 0xc7, 0xa5, 0x89, 0x9b, 0x54, 0x8b, 0x43, 0x8c, 0x41, 0x73, 0xc1,
	0xa7, 0x22, 0x6e, 0x11, 0x4b, 0x67, 0xf6, 0x37, 0xb4, 0x0a, 0x39, 0x93, 0x18, 0xb7, 0x89, 0xac,
	0x00, 0xdb, 0x83, 0x8e, 0x51, 0x1a, 0xd3, 0x9b, 0x32, 0xee, 0xd0, 0xe5, 0x6d, 0x0b, 0x27, 0x25,
	0xdb, 0x87, 0xae, 0xd2, 0xb9, 0xd0, 0x56, 0xe9, 0x92, 0xd2, 0x21, 0x3c, 0x29, 0x47, 0x77, 0xb0,
	0xf7, 0xac, 0x59, 0xb3, 0x50, 0x73, 0x23, 0xd8, 0x1b, 0x68, 0xe6, 0x1c, 0x79, 0x1c, 0x0c, 0xc3,
	0xa3, 0xde, 0x31, 0x1b, 0xaf, 0x26, 0xb9, 0x8a, 0x24, 0x9d, 0x8d, 0xa1, 0x3b, 0x13, 0xc8, 0x29,
	0xd6, 0x36, 0xfc, 0x24, 0xf6, 0x93, 0x53, 0x92, 0x55, 0xcc, 0xe8, 0x0b, 0x74, 0x3d, 0xbb, 0x6a,
	0x2e, 0x78, 0xda, 0x1c, 0x2a, 0xe4, 0x05, 0x5d, 0x16, 0x26, 0x15, 0x60, 0xff, 0x01, 0xd0, 0x21,
	0xa5, 0xf8, 0x90, 0xa4, 0x88, 0x98, 0x2b, 0x3e, 0x15, 0xa3, 0x9f, 0x21, 0x74, 0x7d, 0x5d, 0xeb,
	0x96, 0xd2, 0x8d, 0x12, 0x0b, 0xef, 0x47, 0x05, 0xd8, 0x5f, 0x10, 0x22, 0x9f, 0x3a, 0x1f, 0xec,
	0xd1, 0xfe, 0x86, 0x9c, 0xf1, 0xa9, 0x48, 0x17, 0x1c, 0x6f, 0xc9, 0x86, 0x28, 0x89, 0x88, 0xb9,
	0xe2, 0x78, 0x6b, 0x1d, 0x22, 0x60, 0xc8, 0x8b, 0x28, 0x71, 0x88, 0x7d, 0x80, 0xfd, 0x82, 0x1b,
	0x4c, 0x6d, 0xd3, 0xa9, 0xed, 0x3a, 0xcd, 0xd4, 0x6c, 0x26, 0xe6, 0x98, 0xca, 0xdc, 0x39, 0xf4,
	0x8f, 0x0d, 0xf0, 0xf5, 0x9d, 0x56, 0xea, 0x45, 0xce, 0xde, 0x41, 0x9f, 0x32, 0x5d, 0x3c, 0xd9,
	0xd6, 0x3b, 0xde, 0xa9, 0xc7, 0xe7, 0x42, 0x93, 0x9e, 0x0d, 0x73, 0x80, 0xfd, 0x0f, 0xbd, 0x42,
	0x7e, 0x13, 0x26, 0xcd, 0xd4, 0x72, 0x8e, 0xe4, 0x68, 0x98, 0x00, 0x51, 0xa7, 0x96, 0xb1, 0x7e,
	0x4b, 0x93, 0x5a, 0x22, 0x8f, 0xa3, 0x61, 0x70, 0xd4, 0x4d, 0x3a, 0xd2, 0x5c, 0x5a, 0xc8, 0x0e,
	0x61, 0xe0, 0x8b, 0xa3, 0x6c, 0x13, 0x03, 0xa5, 0x6f, 0x39, 0x96, 0x2e, 0x30, 0x8f, 0x1e, 0x63,
	0xef, 0xc9, 0x63, 0x3c, 0x84, 0x16, 0xcf, 0x50, 0xe9, 0xb8, 0x4f, 0x95, 0x6e, 0xd7, 0x95, 0x9e,
	0x58, 0x3a, 0xa9, 0x54, 0x3b, 0xc8, 0x4c, 0x0b, 0x8e, 0x22, 0x4f, 0x39, 0xc6, 0x5b, 0xd5, 0x20,
	0x1d, 0x73, 0x82, 0x56, 0x5e, 0x2e, 0x72, 0x2f, 0x0f, 0x2a, 0xd9, 0x31, 0x27, 0x38, 0xfa, 0x1d,
	0x40, 0xc7, 0xf7, 0xba, 0x6e, 0xe5, 0x10, 0xfa, 0xf5, 0x98, 0x65, 0xee, 0xde, 0x08, 0x2c, 0xdd,
	0x68, 0x2f, 0x72, 0x16, 0x43, 0xc7, 0x8f, 0xb3, 0xb2, 0xd6, 0xc3, 0xb5, 0xaa, 0x9a, 0x2f, 0x57,
	0xd5, 0x5a, 0xab, 0x8a, 0x8d, 0xeb, 0xec, 0x9b, 0x92, 0x6c, 0xdd, 0xd0, 0xbf, 0xbf, 0x6e, 0x52,
	0xda, 0x78, 0x7f, 0x9d, 0xfb, 0x20, 0x37, 0xc5, 0xbb, 0x90, 0x49, 0x39, 0xfa, 0x11, 0x40, 0x8b,
	0xc8, 0x67, 0x3d, 0x33, 0x68, 0x3e, 0xda, 0x26, 0x74, 0x66, 0xff, 0x42, 0xb4, 0xb8, 0x55, 0xa8,
	0xd2, 0xa5, 0x2e, 0xfc, 0x2a, 0x21, 0xe2, 0x5a, 0x17, 0x36, 0x41, 0xab, 0x42, 0xb8, 0x45, 0x42,
	0x67, 0x3b, 0x16, 0x2d, 0xa6, 0x62, 0x9e, 0x95, 0xae, 0x35, 0x0f, 0xed, 0x52, 0xca, 0xa5, 0x41,
	0x2d, 0xb3, 0x6a, 0x9f, 0x44, 0xc9, 0x0a, 0xdb, 0xac, 0x7b, 0x59, 0x14, 0xf6, 0x93, 0xab, 0x56,
	0x8a, 0x87, 0xb6, 0x48, 0xfd, 0xe0, 0xb6, 0x49, 0x43, 0x3f, 0x1c, 0x7f, 0x0f, 0x60, 0xdb, 0x3f,
	0xf0, 0x8f, 0x7c, 0x9e, 0x17, 0x42, 0xb3, 0xaf, 0xb0, 0xbd, 0xb6, 0x5c, 0xd8, 0xb0, 0x9e, 0xc0,
	0xe6, 0x25, 0x7b, 0xf0, 0xea, 0x85, 0x08, 0xb7, 0x99, 0xde, 0xc3, 0xce, 0xb9, 0xc0, 0x33, 0x81,
	0x5c, 0x16, 0xf5, 0x47, 0x5f, 0xe7, 0xd9, 0xbd, 0x7e, 0xb0, 0x61, 0x61, 0xdd, 0xb4, 0xe9, 0x4f,
	0xe0, 0xed, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x8f, 0x6b, 0x62, 0x16, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserPostHandlerClient is the client API for UserPostHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserPostHandlerClient interface {
	GetListUserPost(ctx context.Context, in *GetListUserPostRequest, opts ...grpc.CallOption) (*GetListUserPostResponse, error)
	GetDetailUserPost(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*UserPost, error)
}

type userPostHandlerClient struct {
	cc *grpc.ClientConn
}

func NewUserPostHandlerClient(cc *grpc.ClientConn) UserPostHandlerClient {
	return &userPostHandlerClient{cc}
}

func (c *userPostHandlerClient) GetListUserPost(ctx context.Context, in *GetListUserPostRequest, opts ...grpc.CallOption) (*GetListUserPostResponse, error) {
	out := new(GetListUserPostResponse)
	err := c.cc.Invoke(ctx, "/userpost.UserPostHandler/GetListUserPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostHandlerClient) GetDetailUserPost(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*UserPost, error) {
	out := new(UserPost)
	err := c.cc.Invoke(ctx, "/userpost.UserPostHandler/GetDetailUserPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserPostHandlerServer is the server API for UserPostHandler service.
type UserPostHandlerServer interface {
	GetListUserPost(context.Context, *GetListUserPostRequest) (*GetListUserPostResponse, error)
	GetDetailUserPost(context.Context, *ByID) (*UserPost, error)
}

// UnimplementedUserPostHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedUserPostHandlerServer struct {
}

func (*UnimplementedUserPostHandlerServer) GetListUserPost(ctx context.Context, req *GetListUserPostRequest) (*GetListUserPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListUserPost not implemented")
}
func (*UnimplementedUserPostHandlerServer) GetDetailUserPost(ctx context.Context, req *ByID) (*UserPost, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetailUserPost not implemented")
}

func RegisterUserPostHandlerServer(s *grpc.Server, srv UserPostHandlerServer) {
	s.RegisterService(&_UserPostHandler_serviceDesc, srv)
}

func _UserPostHandler_GetListUserPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListUserPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostHandlerServer).GetListUserPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpost.UserPostHandler/GetListUserPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostHandlerServer).GetListUserPost(ctx, req.(*GetListUserPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPostHandler_GetDetailUserPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostHandlerServer).GetDetailUserPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpost.UserPostHandler/GetDetailUserPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostHandlerServer).GetDetailUserPost(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserPostHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userpost.UserPostHandler",
	HandlerType: (*UserPostHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListUserPost",
			Handler:    _UserPostHandler_GetListUserPost_Handler,
		},
		{
			MethodName: "GetDetailUserPost",
			Handler:    _UserPostHandler_GetDetailUserPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userpost.proto",
}
