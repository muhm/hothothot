// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

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

type IdReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdReq) Reset()         { *m = IdReq{} }
func (m *IdReq) String() string { return proto.CompactTextString(m) }
func (*IdReq) ProtoMessage()    {}
func (*IdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *IdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdReq.Unmarshal(m, b)
}
func (m *IdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdReq.Marshal(b, m, deterministic)
}
func (m *IdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdReq.Merge(m, src)
}
func (m *IdReq) XXX_Size() int {
	return xxx_messageInfo_IdReq.Size(m)
}
func (m *IdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IdReq.DiscardUnknown(m)
}

var xxx_messageInfo_IdReq proto.InternalMessageInfo

func (m *IdReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type NameReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameReq) Reset()         { *m = NameReq{} }
func (m *NameReq) String() string { return proto.CompactTextString(m) }
func (*NameReq) ProtoMessage()    {}
func (*NameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *NameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameReq.Unmarshal(m, b)
}
func (m *NameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameReq.Marshal(b, m, deterministic)
}
func (m *NameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameReq.Merge(m, src)
}
func (m *NameReq) XXX_Size() int {
	return xxx_messageInfo_NameReq.Size(m)
}
func (m *NameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NameReq.DiscardUnknown(m)
}

var xxx_messageInfo_NameReq proto.InternalMessageInfo

func (m *NameReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MailReq struct {
	Mail                 string   `protobuf:"bytes,1,opt,name=mail,proto3" json:"mail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MailReq) Reset()         { *m = MailReq{} }
func (m *MailReq) String() string { return proto.CompactTextString(m) }
func (*MailReq) ProtoMessage()    {}
func (*MailReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *MailReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MailReq.Unmarshal(m, b)
}
func (m *MailReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MailReq.Marshal(b, m, deterministic)
}
func (m *MailReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MailReq.Merge(m, src)
}
func (m *MailReq) XXX_Size() int {
	return xxx_messageInfo_MailReq.Size(m)
}
func (m *MailReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MailReq.DiscardUnknown(m)
}

var xxx_messageInfo_MailReq proto.InternalMessageInfo

func (m *MailReq) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

type UserInfoReply struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Gender               string   `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Mail                 string   `protobuf:"bytes,5,opt,name=mail,proto3" json:"mail,omitempty"`
	Password             string   `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	OtpSecret            string   `protobuf:"bytes,7,opt,name=otpSecret,proto3" json:"otpSecret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoReply) Reset()         { *m = UserInfoReply{} }
func (m *UserInfoReply) String() string { return proto.CompactTextString(m) }
func (*UserInfoReply) ProtoMessage()    {}
func (*UserInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserInfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoReply.Unmarshal(m, b)
}
func (m *UserInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoReply.Marshal(b, m, deterministic)
}
func (m *UserInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoReply.Merge(m, src)
}
func (m *UserInfoReply) XXX_Size() int {
	return xxx_messageInfo_UserInfoReply.Size(m)
}
func (m *UserInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoReply proto.InternalMessageInfo

func (m *UserInfoReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfoReply) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserInfoReply) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserInfoReply) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UserInfoReply) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

func (m *UserInfoReply) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserInfoReply) GetOtpSecret() string {
	if m != nil {
		return m.OtpSecret
	}
	return ""
}

type OtpReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Secret               string   `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OtpReq) Reset()         { *m = OtpReq{} }
func (m *OtpReq) String() string { return proto.CompactTextString(m) }
func (*OtpReq) ProtoMessage()    {}
func (*OtpReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *OtpReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtpReq.Unmarshal(m, b)
}
func (m *OtpReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtpReq.Marshal(b, m, deterministic)
}
func (m *OtpReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtpReq.Merge(m, src)
}
func (m *OtpReq) XXX_Size() int {
	return xxx_messageInfo_OtpReq.Size(m)
}
func (m *OtpReq) XXX_DiscardUnknown() {
	xxx_messageInfo_OtpReq.DiscardUnknown(m)
}

var xxx_messageInfo_OtpReq proto.InternalMessageInfo

func (m *OtpReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OtpReq) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

type OtpReply struct {
	Pong                 string   `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OtpReply) Reset()         { *m = OtpReply{} }
func (m *OtpReply) String() string { return proto.CompactTextString(m) }
func (*OtpReply) ProtoMessage()    {}
func (*OtpReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *OtpReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtpReply.Unmarshal(m, b)
}
func (m *OtpReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtpReply.Marshal(b, m, deterministic)
}
func (m *OtpReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtpReply.Merge(m, src)
}
func (m *OtpReply) XXX_Size() int {
	return xxx_messageInfo_OtpReply.Size(m)
}
func (m *OtpReply) XXX_DiscardUnknown() {
	xxx_messageInfo_OtpReply.DiscardUnknown(m)
}

var xxx_messageInfo_OtpReply proto.InternalMessageInfo

func (m *OtpReply) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

func init() {
	proto.RegisterType((*IdReq)(nil), "user.IdReq")
	proto.RegisterType((*NameReq)(nil), "user.NameReq")
	proto.RegisterType((*MailReq)(nil), "user.MailReq")
	proto.RegisterType((*UserInfoReply)(nil), "user.UserInfoReply")
	proto.RegisterType((*OtpReq)(nil), "user.OtpReq")
	proto.RegisterType((*OtpReply)(nil), "user.OtpReply")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x53, 0x28, 0x05, 0x06, 0xcb, 0x61, 0x4d, 0xb4, 0x21, 0x6a, 0x4c, 0x4f, 0x5e, 0x44,
	0x23, 0x6f, 0xe0, 0x8d, 0x83, 0x92, 0xd4, 0xf8, 0x00, 0x2b, 0x3b, 0x92, 0xc6, 0xd2, 0x5d, 0x77,
	0xab, 0x86, 0x87, 0xf3, 0x1d, 0x7c, 0x24, 0x33, 0xb3, 0x5d, 0x50, 0x43, 0xe2, 0x6d, 0x7e, 0xbe,
	0xfd, 0x87, 0xe1, 0xff, 0x01, 0x78, 0x73, 0x68, 0xa7, 0xc6, 0xea, 0x46, 0x8b, 0x98, 0xe6, 0xfc,
	0x18, 0x7a, 0x73, 0x55, 0xe0, 0xab, 0x18, 0x43, 0xa7, 0x54, 0x59, 0x74, 0x1e, 0x5d, 0x74, 0x8b,
	0x4e, 0xa9, 0xf2, 0x53, 0xe8, 0xdf, 0xcb, 0x35, 0x12, 0x12, 0x10, 0xd7, 0x72, 0x8d, 0x0c, 0x87,
	0x05, 0xcf, 0x84, 0xef, 0x64, 0x59, 0xb5, 0x78, 0x2d, 0xcb, 0x2a, 0x60, 0x9a, 0xf3, 0xcf, 0x08,
	0xd2, 0x47, 0x87, 0x76, 0x5e, 0x3f, 0xeb, 0x02, 0x4d, 0xb5, 0xf9, 0xbb, 0x5f, 0x4c, 0x60, 0x40,
	0x07, 0xf0, 0xe2, 0x0e, 0x3b, 0xb7, 0x9a, 0x58, 0x5d, 0x2e, 0x5f, 0x98, 0x75, 0x3d, 0x0b, 0x5a,
	0x1c, 0x41, 0xb2, 0xc2, 0x5a, 0xa1, 0xcd, 0x62, 0x26, 0xad, 0xda, 0x5e, 0xd1, 0xdb, 0x5d, 0x41,
	0x7b, 0x8c, 0x74, 0xee, 0x43, 0x5b, 0x95, 0x25, 0x7e, 0x4f, 0xd0, 0xe2, 0x04, 0x86, 0xba, 0x31,
	0x0f, 0xb8, 0xb4, 0xd8, 0x64, 0x7d, 0x86, 0xbb, 0x0f, 0xf2, 0x6b, 0x48, 0x16, 0x8d, 0xd9, 0x93,
	0x0b, 0x7d, 0xbf, 0xf3, 0x26, 0x7f, 0x75, 0xab, 0xf2, 0x33, 0x18, 0xb0, 0x83, 0x7e, 0xab, 0x80,
	0xd8, 0xe8, 0x7a, 0x15, 0x12, 0xa1, 0xf9, 0xe6, 0x2b, 0x02, 0x4e, 0x5c, 0x5c, 0xc1, 0x68, 0x85,
	0x0d, 0x85, 0x73, 0xbb, 0x99, 0x2b, 0x31, 0x9a, 0x72, 0x27, 0x5c, 0xc2, 0xe4, 0xd0, 0x8b, 0xdf,
	0xc9, 0xcd, 0x20, 0xdd, 0x1a, 0xa8, 0x12, 0x91, 0xfa, 0x57, 0x6d, 0x3d, 0xff, 0x9b, 0xa8, 0xa8,
	0x60, 0x6a, 0x4b, 0xdb, 0x6f, 0xba, 0x84, 0xd4, 0xc9, 0x77, 0x5c, 0x84, 0x18, 0xc4, 0x81, 0x7f,
	0xe5, 0xa3, 0x98, 0x8c, 0x7f, 0x28, 0x53, 0x6d, 0x9e, 0x12, 0xfe, 0x23, 0xcd, 0xbe, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xbc, 0x6d, 0xf1, 0xdf, 0x56, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetUserById(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserInfoReply, error)
	GetUserByName(ctx context.Context, in *NameReq, opts ...grpc.CallOption) (*UserInfoReply, error)
	GetUserByMail(ctx context.Context, in *MailReq, opts ...grpc.CallOption) (*UserInfoReply, error)
	SaveOtpSecret(ctx context.Context, in *OtpReq, opts ...grpc.CallOption) (*OtpReply, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserById(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, "/user.user/getUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserByName(ctx context.Context, in *NameReq, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, "/user.user/getUserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserByMail(ctx context.Context, in *MailReq, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, "/user.user/getUserByMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SaveOtpSecret(ctx context.Context, in *OtpReq, opts ...grpc.CallOption) (*OtpReply, error) {
	out := new(OtpReply)
	err := c.cc.Invoke(ctx, "/user.user/saveOtpSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetUserById(context.Context, *IdReq) (*UserInfoReply, error)
	GetUserByName(context.Context, *NameReq) (*UserInfoReply, error)
	GetUserByMail(context.Context, *MailReq) (*UserInfoReply, error)
	SaveOtpSecret(context.Context, *OtpReq) (*OtpReply, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUserById(ctx context.Context, req *IdReq) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (*UnimplementedUserServer) GetUserByName(ctx context.Context, req *NameReq) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByName not implemented")
}
func (*UnimplementedUserServer) GetUserByMail(ctx context.Context, req *MailReq) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByMail not implemented")
}
func (*UnimplementedUserServer) SaveOtpSecret(ctx context.Context, req *OtpReq) (*OtpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveOtpSecret not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.user/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserById(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.user/GetUserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByName(ctx, req.(*NameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserByMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.user/GetUserByMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByMail(ctx, req.(*MailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SaveOtpSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OtpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SaveOtpSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.user/SaveOtpSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SaveOtpSecret(ctx, req.(*OtpReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.user",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getUserById",
			Handler:    _User_GetUserById_Handler,
		},
		{
			MethodName: "getUserByName",
			Handler:    _User_GetUserByName_Handler,
		},
		{
			MethodName: "getUserByMail",
			Handler:    _User_GetUserByMail_Handler,
		},
		{
			MethodName: "saveOtpSecret",
			Handler:    _User_SaveOtpSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
