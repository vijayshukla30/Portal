// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/proto/account.proto

package user

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

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Mobile               int32    `protobuf:"varint,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Dob                  int64    `protobuf:"varint,6,opt,name=dob,proto3" json:"dob,omitempty"`
	Doj                  int64    `protobuf:"varint,7,opt,name=doj,proto3" json:"doj,omitempty"`
	Created              int64    `protobuf:"varint,8,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,9,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetMobile() int32 {
	if m != nil {
		return m.Mobile
	}
	return 0
}

func (m *User) GetDob() int64 {
	if m != nil {
		return m.Dob
	}
	return 0
}

func (m *User) GetDoj() int64 {
	if m != nil {
		return m.Doj
	}
	return 0
}

func (m *User) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *User) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type Session struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Expires              int64    `protobuf:"varint,4,opt,name=expires,proto3" json:"expires,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{1}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Session) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Session) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Session) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

//Create
type CreateRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{2}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{3}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//Read
type ReadRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{4}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{5}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

//Update
type UpdateRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{6}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{7}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//Delete
type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{8}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{9}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//Search
type SearchRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int64    `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{10}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SearchRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type SearchResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{11}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

//Update Password
type UpdatePasswordRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	OldPassword          string   `protobuf:"bytes,2,opt,name=oldPassword,proto3" json:"oldPassword,omitempty"`
	NewPassword          string   `protobuf:"bytes,3,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	ConfirmPassword      string   `protobuf:"bytes,4,opt,name=confirmPassword,proto3" json:"confirmPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePasswordRequest) Reset()         { *m = UpdatePasswordRequest{} }
func (m *UpdatePasswordRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePasswordRequest) ProtoMessage()    {}
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{12}
}

func (m *UpdatePasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePasswordRequest.Unmarshal(m, b)
}
func (m *UpdatePasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePasswordRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePasswordRequest.Merge(m, src)
}
func (m *UpdatePasswordRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePasswordRequest.Size(m)
}
func (m *UpdatePasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePasswordRequest proto.InternalMessageInfo

func (m *UpdatePasswordRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdatePasswordRequest) GetOldPassword() string {
	if m != nil {
		return m.OldPassword
	}
	return ""
}

func (m *UpdatePasswordRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

func (m *UpdatePasswordRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

type UpdatePasswordResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePasswordResponse) Reset()         { *m = UpdatePasswordResponse{} }
func (m *UpdatePasswordResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePasswordResponse) ProtoMessage()    {}
func (*UpdatePasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{13}
}

func (m *UpdatePasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePasswordResponse.Unmarshal(m, b)
}
func (m *UpdatePasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePasswordResponse.Marshal(b, m, deterministic)
}
func (m *UpdatePasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePasswordResponse.Merge(m, src)
}
func (m *UpdatePasswordResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePasswordResponse.Size(m)
}
func (m *UpdatePasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePasswordResponse proto.InternalMessageInfo

func (m *UpdatePasswordResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//Login
type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{14}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Session              *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{15}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

//Logout
type LogoutRequest struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{16}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type LogoutResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{17}
}

func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (m *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(m, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

func (m *LogoutResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//Read Session
type ReadSessionRequest struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadSessionRequest) Reset()         { *m = ReadSessionRequest{} }
func (m *ReadSessionRequest) String() string { return proto.CompactTextString(m) }
func (*ReadSessionRequest) ProtoMessage()    {}
func (*ReadSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{18}
}

func (m *ReadSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadSessionRequest.Unmarshal(m, b)
}
func (m *ReadSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadSessionRequest.Marshal(b, m, deterministic)
}
func (m *ReadSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadSessionRequest.Merge(m, src)
}
func (m *ReadSessionRequest) XXX_Size() int {
	return xxx_messageInfo_ReadSessionRequest.Size(m)
}
func (m *ReadSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadSessionRequest proto.InternalMessageInfo

func (m *ReadSessionRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type ReadSessionResponse struct {
	Session              *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadSessionResponse) Reset()         { *m = ReadSessionResponse{} }
func (m *ReadSessionResponse) String() string { return proto.CompactTextString(m) }
func (*ReadSessionResponse) ProtoMessage()    {}
func (*ReadSessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3a352d62e9c04dd, []int{19}
}

func (m *ReadSessionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadSessionResponse.Unmarshal(m, b)
}
func (m *ReadSessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadSessionResponse.Marshal(b, m, deterministic)
}
func (m *ReadSessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadSessionResponse.Merge(m, src)
}
func (m *ReadSessionResponse) XXX_Size() int {
	return xxx_messageInfo_ReadSessionResponse.Size(m)
}
func (m *ReadSessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadSessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadSessionResponse proto.InternalMessageInfo

func (m *ReadSessionResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Session)(nil), "user.Session")
	proto.RegisterType((*CreateRequest)(nil), "user.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "user.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "user.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "user.ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "user.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "user.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "user.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "user.DeleteResponse")
	proto.RegisterType((*SearchRequest)(nil), "user.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "user.SearchResponse")
	proto.RegisterType((*UpdatePasswordRequest)(nil), "user.UpdatePasswordRequest")
	proto.RegisterType((*UpdatePasswordResponse)(nil), "user.UpdatePasswordResponse")
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "user.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "user.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "user.LogoutResponse")
	proto.RegisterType((*ReadSessionRequest)(nil), "user.ReadSessionRequest")
	proto.RegisterType((*ReadSessionResponse)(nil), "user.ReadSessionResponse")
}

func init() { proto.RegisterFile("user/proto/account.proto", fileDescriptor_c3a352d62e9c04dd) }

var fileDescriptor_c3a352d62e9c04dd = []byte{
	// 680 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0xae, 0xeb, 0x1c, 0xea, 0x69, 0x9d, 0xff, 0x67, 0x5b, 0x2a, 0x63, 0x0a, 0x58, 0x7b, 0x43,
	0x54, 0x89, 0x44, 0x32, 0xaa, 0xc4, 0x15, 0x12, 0xa2, 0x42, 0x42, 0x1c, 0x54, 0xb9, 0xea, 0x05,
	0x97, 0x6e, 0xbc, 0x29, 0xae, 0x12, 0x6f, 0xf0, 0x3a, 0x2a, 0x4f, 0xc2, 0x1d, 0x6f, 0xc6, 0xc3,
	0xa0, 0x3d, 0x8c, 0xed, 0x35, 0x95, 0x92, 0xde, 0x65, 0xe6, 0x9b, 0xc3, 0x37, 0xb3, 0xf3, 0x39,
	0x10, 0xac, 0x05, 0x2b, 0xa7, 0xab, 0x92, 0x57, 0x7c, 0x9a, 0xce, 0x66, 0x7c, 0x5d, 0x54, 0x13,
	0x65, 0x91, 0x9e, 0x44, 0xe8, 0x1f, 0x07, 0x7a, 0x57, 0x82, 0x95, 0x64, 0x04, 0xbb, 0x79, 0x16,
	0x38, 0x91, 0x33, 0x76, 0x93, 0xdd, 0x3c, 0x23, 0x21, 0xec, 0xc9, 0x80, 0x22, 0x5d, 0xb2, 0x60,
	0x37, 0x72, 0xc6, 0x5e, 0x52, 0xdb, 0xe4, 0x04, 0xbc, 0x79, 0x5e, 0x8a, 0xea, 0xab, 0x04, 0x5d,
	0x05, 0x36, 0x0e, 0x99, 0xb9, 0x48, 0x0d, 0xd8, 0xd3, 0x99, 0x68, 0x93, 0x63, 0x18, 0x2c, 0xf9,
	0x75, 0xbe, 0x60, 0x41, 0x3f, 0x72, 0xc6, 0xfd, 0xc4, 0x58, 0xe4, 0x7f, 0x70, 0x33, 0x7e, 0x1d,
	0x0c, 0x54, 0x7b, 0xf9, 0x53, 0x7b, 0x6e, 0x83, 0x21, 0x7a, 0x6e, 0x49, 0x00, 0xc3, 0x59, 0xc9,
	0xd2, 0x8a, 0x65, 0xc1, 0x9e, 0xf2, 0xa2, 0x29, 0x91, 0xf5, 0x2a, 0x53, 0x88, 0xa7, 0x11, 0x63,
	0xd2, 0x1c, 0x86, 0x97, 0x4c, 0x88, 0x9c, 0x17, 0x0f, 0x1a, 0xb0, 0xd5, 0xca, 0xfd, 0xa7, 0x15,
	0xfb, 0xb9, 0xca, 0x4b, 0x26, 0xd4, 0x6c, 0x6e, 0x82, 0x26, 0xfd, 0x04, 0xfe, 0x7b, 0x15, 0x94,
	0xb0, 0x1f, 0x6b, 0x26, 0x2a, 0xf2, 0x1c, 0xd4, 0x8a, 0x55, 0xcb, 0xfd, 0x18, 0x26, 0xd2, 0x98,
	0xc8, 0x5d, 0x27, 0xca, 0x2f, 0x09, 0xac, 0x52, 0x21, 0xee, 0x78, 0x99, 0x21, 0x01, 0xb4, 0xe9,
	0x29, 0x8c, 0xb0, 0x98, 0x58, 0xf1, 0x42, 0x28, 0x4a, 0x4b, 0x26, 0x44, 0x7a, 0xc3, 0x54, 0x41,
	0x2f, 0x41, 0x93, 0x3e, 0x83, 0xfd, 0x84, 0xa5, 0x19, 0xb6, 0x6d, 0xe6, 0xf4, 0xe4, 0x9c, 0x74,
	0x02, 0x07, 0x1a, 0x36, 0x85, 0x36, 0xd0, 0xa2, 0x53, 0xf0, 0xaf, 0xd4, 0xf6, 0xb6, 0x9c, 0x43,
	0x72, 0xc5, 0x84, 0x8d, 0x5c, 0x5f, 0x80, 0x7f, 0xce, 0x16, 0xac, 0x29, 0xde, 0x65, 0x7b, 0x0a,
	0x23, 0x0c, 0xd8, 0x58, 0xec, 0x1b, 0xf8, 0x97, 0x2c, 0x2d, 0x67, 0xdf, 0xb1, 0x58, 0xfb, 0x49,
	0x9d, 0xce, 0x93, 0x1e, 0x41, 0x7f, 0x91, 0x2f, 0xf3, 0xca, 0x3c, 0xa8, 0x36, 0xe4, 0x3d, 0xf2,
	0xf9, 0x5c, 0xb0, 0xca, 0xbc, 0xa6, 0xb1, 0x68, 0x0c, 0x23, 0x2c, 0x6d, 0x68, 0x44, 0xd0, 0x97,
	0xb5, 0x44, 0xe0, 0x44, 0x6e, 0x67, 0x0d, 0x1a, 0xa0, 0xbf, 0x1d, 0x78, 0xac, 0x17, 0x71, 0x61,
	0x9e, 0x11, 0x79, 0x1d, 0xc3, 0x40, 0x86, 0x7c, 0xc4, 0x41, 0x8d, 0x45, 0x22, 0xd8, 0xe7, 0x8b,
	0xec, 0xc2, 0x3e, 0x82, 0xb6, 0x4b, 0x46, 0x14, 0xec, 0xae, 0x8e, 0xd0, 0x5a, 0x6b, 0xbb, 0xc8,
	0x18, 0xfe, 0x9b, 0xf1, 0x62, 0x9e, 0x97, 0xcb, 0x3a, 0x4a, 0x8b, 0xae, 0xeb, 0xa6, 0x31, 0x1c,
	0x77, 0xe9, 0x6d, 0x5c, 0xf1, 0x07, 0x38, 0xf8, 0xcc, 0x6f, 0xf2, 0x62, 0x9b, 0x0d, 0xb7, 0xef,
	0xd9, 0xed, 0xdc, 0xf3, 0x1b, 0xf0, 0x4d, 0x1d, 0xd3, 0xf2, 0x25, 0x0c, 0x85, 0x16, 0xa6, 0xb9,
	0x2b, 0x5f, 0x2f, 0xd4, 0xa8, 0x35, 0x41, 0x94, 0xbe, 0x52, 0x99, 0x7c, 0x5d, 0x21, 0x85, 0x13,
	0xf0, 0x0c, 0x56, 0xef, 0xb3, 0x71, 0xc8, 0xfb, 0xc1, 0xf0, 0x8d, 0xc3, 0xc5, 0x40, 0xa4, 0x32,
	0xb0, 0xe5, 0x56, 0xf5, 0xdf, 0xc2, 0xa1, 0x95, 0xf3, 0xc0, 0x71, 0xe2, 0x5f, 0x3d, 0x18, 0xbe,
	0xd3, 0xdf, 0x61, 0x72, 0x06, 0x03, 0x2d, 0x72, 0x72, 0xa8, 0xa3, 0xad, 0xef, 0x47, 0x78, 0x64,
	0x3b, 0x75, 0x27, 0xba, 0x43, 0xa6, 0xd0, 0x93, 0x14, 0xc8, 0x23, 0x8d, 0xb7, 0xb4, 0x1f, 0x92,
	0xb6, 0xab, 0x4e, 0x38, 0x83, 0x81, 0x7e, 0x78, 0xec, 0x63, 0xe9, 0x1b, 0xfb, 0xd8, 0x1a, 0xd6,
	0x69, 0x5a, 0x8a, 0x98, 0x66, 0x29, 0x17, 0xd3, 0x6c, 0xb5, 0xea, 0x34, 0x2d, 0x1d, 0x4c, 0xb3,
	0x34, 0x8a, 0x69, 0xb6, 0xba, 0xe8, 0x0e, 0xf9, 0x82, 0x5f, 0x91, 0xfa, 0xb2, 0x9f, 0xb6, 0x79,
	0x75, 0x24, 0x15, 0x9e, 0xdc, 0x0f, 0xd6, 0xe5, 0x62, 0xe8, 0xab, 0x83, 0x23, 0x66, 0x25, 0xed,
	0x2b, 0x0e, 0x0f, 0x2d, 0x5f, 0x9b, 0xb9, 0xbe, 0x1d, 0xd2, 0x04, 0x34, 0x87, 0x87, 0xcc, 0xed,
	0xf3, 0xa2, 0x3b, 0xe4, 0x5c, 0x7f, 0x7f, 0xf1, 0x7f, 0x26, 0x68, 0xde, 0xc0, 0xbe, 0xac, 0xf0,
	0xc9, 0x3d, 0x08, 0x56, 0xb9, 0x1e, 0xa8, 0x7f, 0xe5, 0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x90, 0xb6, 0x31, 0x94, 0xb1, 0x07, 0x00, 0x00,
}
