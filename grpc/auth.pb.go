// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetUserRequest struct {
	// Username.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Plain text password.
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GetSuperuserRequest struct {
	// Username.
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSuperuserRequest) Reset()         { *m = GetSuperuserRequest{} }
func (m *GetSuperuserRequest) String() string { return proto.CompactTextString(m) }
func (*GetSuperuserRequest) ProtoMessage()    {}
func (*GetSuperuserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *GetSuperuserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSuperuserRequest.Unmarshal(m, b)
}
func (m *GetSuperuserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSuperuserRequest.Marshal(b, m, deterministic)
}
func (m *GetSuperuserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSuperuserRequest.Merge(m, src)
}
func (m *GetSuperuserRequest) XXX_Size() int {
	return xxx_messageInfo_GetSuperuserRequest.Size(m)
}
func (m *GetSuperuserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSuperuserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSuperuserRequest proto.InternalMessageInfo

func (m *GetSuperuserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type CheckAclRequest struct {
	// Username.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Topic to be checked for.
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	// The client connection's id.
	Clientid string `protobuf:"bytes,3,opt,name=clientid,proto3" json:"clientid,omitempty"`
	// Topic access.
	Acc                  int32    `protobuf:"varint,4,opt,name=acc,proto3" json:"acc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAclRequest) Reset()         { *m = CheckAclRequest{} }
func (m *CheckAclRequest) String() string { return proto.CompactTextString(m) }
func (*CheckAclRequest) ProtoMessage()    {}
func (*CheckAclRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *CheckAclRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAclRequest.Unmarshal(m, b)
}
func (m *CheckAclRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAclRequest.Marshal(b, m, deterministic)
}
func (m *CheckAclRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAclRequest.Merge(m, src)
}
func (m *CheckAclRequest) XXX_Size() int {
	return xxx_messageInfo_CheckAclRequest.Size(m)
}
func (m *CheckAclRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAclRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAclRequest proto.InternalMessageInfo

func (m *CheckAclRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CheckAclRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *CheckAclRequest) GetClientid() string {
	if m != nil {
		return m.Clientid
	}
	return ""
}

func (m *CheckAclRequest) GetAcc() int32 {
	if m != nil {
		return m.Acc
	}
	return 0
}

type AuthResponse struct {
	// If the user is authorized/authenticated.
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type NameResponse struct {
	// The name of the gRPC backend.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameResponse) Reset()         { *m = NameResponse{} }
func (m *NameResponse) String() string { return proto.CompactTextString(m) }
func (*NameResponse) ProtoMessage()    {}
func (*NameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *NameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameResponse.Unmarshal(m, b)
}
func (m *NameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameResponse.Marshal(b, m, deterministic)
}
func (m *NameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameResponse.Merge(m, src)
}
func (m *NameResponse) XXX_Size() int {
	return xxx_messageInfo_NameResponse.Size(m)
}
func (m *NameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NameResponse proto.InternalMessageInfo

func (m *NameResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "grpc.GetUserRequest")
	proto.RegisterType((*GetSuperuserRequest)(nil), "grpc.GetSuperuserRequest")
	proto.RegisterType((*CheckAclRequest)(nil), "grpc.CheckAclRequest")
	proto.RegisterType((*AuthResponse)(nil), "grpc.AuthResponse")
	proto.RegisterType((*NameResponse)(nil), "grpc.NameResponse")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x4e, 0xc2, 0x30,
	0x14, 0xc6, 0x61, 0x0c, 0xc5, 0x23, 0x41, 0x53, 0xd1, 0x4c, 0x4c, 0x0c, 0xe9, 0x15, 0x57, 0x23,
	0x6a, 0x8c, 0xde, 0x19, 0x62, 0x0c, 0x5c, 0x79, 0x31, 0xe2, 0x03, 0x8c, 0x72, 0x84, 0x85, 0x41,
	0x4b, 0xff, 0x68, 0x7c, 0x2c, 0xdf, 0xd0, 0x74, 0x65, 0xcb, 0x24, 0x59, 0xc2, 0x5d, 0xcf, 0x77,
	0xfa, 0x9d, 0xaf, 0xfd, 0xe5, 0x00, 0xc4, 0x46, 0x2f, 0x43, 0x21, 0xb9, 0xe6, 0xc4, 0x5f, 0x48,
	0xc1, 0x7a, 0x37, 0x0b, 0xce, 0x17, 0x29, 0x0e, 0x33, 0x6d, 0x66, 0x3e, 0x87, 0xb8, 0x16, 0xfa,
	0xc7, 0x5d, 0xa1, 0x13, 0xe8, 0x8c, 0x51, 0x7f, 0x28, 0x94, 0x11, 0x6e, 0x0d, 0x2a, 0x4d, 0x7a,
	0xd0, 0x32, 0x0a, 0xe5, 0x26, 0x5e, 0x63, 0x50, 0xef, 0xd7, 0x07, 0x27, 0x51, 0x51, 0xdb, 0x9e,
	0x88, 0x95, 0xfa, 0xe6, 0x72, 0x1e, 0x78, 0xae, 0x97, 0xd7, 0xf4, 0x0e, 0x2e, 0xc6, 0xa8, 0xa7,
	0x46, 0xa0, 0x34, 0x87, 0x8d, 0xa3, 0x5b, 0x38, 0x7b, 0x5d, 0x22, 0x5b, 0x8d, 0x58, 0x7a, 0x48,
	0x7a, 0x17, 0x9a, 0x9a, 0x8b, 0x84, 0xed, 0xa2, 0x5d, 0x61, 0x1d, 0x2c, 0x4d, 0x70, 0xa3, 0x93,
	0x79, 0xd0, 0x70, 0x8e, 0xbc, 0x26, 0xe7, 0xd0, 0x88, 0x19, 0x0b, 0xfc, 0x7e, 0x7d, 0xd0, 0x8c,
	0xec, 0x91, 0xde, 0x42, 0x7b, 0x64, 0xf4, 0x32, 0x42, 0x25, 0xf8, 0x46, 0x21, 0xe9, 0x80, 0xc7,
	0x57, 0x59, 0x52, 0x2b, 0xf2, 0xf8, 0x8a, 0x52, 0x68, 0xbf, 0xc7, 0x6b, 0x2c, 0xfa, 0x04, 0xfc,
	0xd2, 0x5b, 0xb2, 0xf3, 0xfd, 0xaf, 0x07, 0xa7, 0x76, 0xc8, 0x14, 0xe5, 0x57, 0xc2, 0x90, 0x3c,
	0xc2, 0xf1, 0x8e, 0x21, 0xe9, 0x86, 0x16, 0x79, 0xf8, 0x1f, 0x69, 0x8f, 0x38, 0xb5, 0x1c, 0x4c,
	0x6b, 0xe4, 0x05, 0xda, 0x65, 0x60, 0xe4, 0xba, 0xf0, 0xee, 0x43, 0xac, 0x18, 0xf0, 0x04, 0xad,
	0x1c, 0x1f, 0xb9, 0x74, 0x37, 0xf6, 0x70, 0x56, 0x1a, 0xed, 0x83, 0xed, 0x3f, 0xc9, 0x55, 0xe8,
	0xb6, 0x23, 0xcc, 0xb7, 0x23, 0x7c, 0xb3, 0xdb, 0x91, 0x1b, 0xcb, 0x2c, 0x68, 0x8d, 0x3c, 0x83,
	0x3f, 0x89, 0x53, 0x5d, 0xe9, 0xaa, 0xd0, 0x69, 0x6d, 0x76, 0x94, 0x29, 0x0f, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xc9, 0x8a, 0xd9, 0x7b, 0x9f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	// GetUser tries to authenticate a user.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// GetSuperuser checks if a user is a superuser.
	GetSuperuser(ctx context.Context, in *GetSuperuserRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// CheckAcl checks user's authorization for the given topic.
	CheckAcl(ctx context.Context, in *CheckAclRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// GetName retrieves the name of the backend.
	GetName(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NameResponse, error)
	// Halt signals the backend to halt.
	Halt(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/grpc.AuthService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetSuperuser(ctx context.Context, in *GetSuperuserRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/grpc.AuthService/GetSuperuser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckAcl(ctx context.Context, in *CheckAclRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/grpc.AuthService/CheckAcl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetName(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NameResponse, error) {
	out := new(NameResponse)
	err := c.cc.Invoke(ctx, "/grpc.AuthService/GetName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Halt(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc.AuthService/Halt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	// GetUser tries to authenticate a user.
	GetUser(context.Context, *GetUserRequest) (*AuthResponse, error)
	// GetSuperuser checks if a user is a superuser.
	GetSuperuser(context.Context, *GetSuperuserRequest) (*AuthResponse, error)
	// CheckAcl checks user's authorization for the given topic.
	CheckAcl(context.Context, *CheckAclRequest) (*AuthResponse, error)
	// GetName retrieves the name of the backend.
	GetName(context.Context, *empty.Empty) (*NameResponse, error)
	// Halt signals the backend to halt.
	Halt(context.Context, *empty.Empty) (*empty.Empty, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetSuperuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSuperuserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetSuperuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthService/GetSuperuser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetSuperuser(ctx, req.(*GetSuperuserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAclRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthService/CheckAcl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckAcl(ctx, req.(*CheckAclRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthService/GetName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetName(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Halt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Halt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthService/Halt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Halt(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _AuthService_GetUser_Handler,
		},
		{
			MethodName: "GetSuperuser",
			Handler:    _AuthService_GetSuperuser_Handler,
		},
		{
			MethodName: "CheckAcl",
			Handler:    _AuthService_CheckAcl_Handler,
		},
		{
			MethodName: "GetName",
			Handler:    _AuthService_GetName_Handler,
		},
		{
			MethodName: "Halt",
			Handler:    _AuthService_Halt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
