// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: apis.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("apis.proto", fileDescriptor_b480fce41873d8e6) }

var fileDescriptor_b480fce41873d8e6 = []byte{
	// 941 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0xd6, 0xdb, 0x8e, 0xdb, 0x44,
	0x18, 0x07, 0x70, 0x1c, 0x50, 0x59, 0xa6, 0x9b, 0x76, 0x77, 0xb2, 0x39, 0x34, 0xdd, 0x8b, 0x91,
	0x2f, 0x03, 0x1b, 0xd3, 0xed, 0x05, 0x52, 0x57, 0xa2, 0x38, 0xb8, 0xac, 0x02, 0x54, 0xac, 0x92,
	0xb6, 0x48, 0x65, 0xa1, 0x72, 0xe2, 0xa9, 0x77, 0xba, 0x89, 0xc7, 0xd8, 0x93, 0x8d, 0xc2, 0x25,
	0x0f, 0x80, 0x84, 0x01, 0x71, 0x83, 0xb4, 0x5c, 0x20, 0x04, 0x12, 0xcf, 0xc1, 0x0b, 0xf0, 0x0a,
	0x3c, 0x08, 0x9a, 0x83, 0xe3, 0xb1, 0x9d, 0x55, 0xc2, 0xe9, 0x6a, 0x35, 0x7f, 0x7f, 0xf6, 0xfc,
	0xf6, 0xf3, 0xcc, 0x38, 0x00, 0xb8, 0x21, 0x89, 0xbb, 0x61, 0x44, 0x19, 0x85, 0x95, 0x70, 0xd4,
	0xde, 0xf7, 0x29, 0xf5, 0x27, 0xd8, 0x72, 0x43, 0x62, 0xb9, 0x41, 0x40, 0x99, 0xcb, 0x08, 0x0d,
	0x54, 0x45, 0xfb, 0x0d, 0xf1, 0x67, 0x7c, 0xe0, 0xe3, 0xe0, 0x20, 0x9e, 0xbb, 0xbe, 0x8f, 0x23,
	0x8b, 0x86, 0xa2, 0x62, 0x45, 0xf5, 0xee, 0x2c, 0xc6, 0xd1, 0x33, 0x77, 0xcc, 0xb3, 0x34, 0x8a,
	0xe8, 0x04, 0xe7, 0xa2, 0xc3, 0xdf, 0xeb, 0xe0, 0x5a, 0xaf, 0xff, 0xd0, 0x3e, 0xe9, 0xc3, 0x1f,
	0x0c, 0xf0, 0xaa, 0x43, 0x3f, 0xa4, 0x3e, 0x09, 0xe0, 0x5e, 0x37, 0x1c, 0x75, 0x1f, 0xc7, 0x38,
	0x12, 0xc3, 0x01, 0xfe, 0x7c, 0x86, 0x63, 0xd6, 0xae, 0x17, 0xd2, 0x38, 0xa4, 0x41, 0x8c, 0x4d,
	0x2f, 0xb1, 0x1d, 0x70, 0x83, 0xe7, 0xb6, 0x78, 0xb4, 0x7d, 0xd2, 0x8f, 0x61, 0x43, 0xd4, 0xa0,
	0xd1, 0x02, 0x79, 0x74, 0xea, 0x92, 0x00, 0xb9, 0x81, 0x87, 0xc2, 0xf9, 0xa8, 0x0d, 0x5a, 0x60,
	0xcf, 0x3e, 0xe9, 0x7f, 0x80, 0x17, 0xf6, 0x8c, 0x9d, 0xe1, 0x80, 0x91, 0xb1, 0x70, 0xc3, 0x97,
	0xbe, 0xfc, 0xe3, 0xcf, 0x6f, 0x2a, 0x4d, 0x13, 0x5a, 0x23, 0x32, 0xb5, 0x2e, 0xee, 0x58, 0xfc,
	0x1f, 0xb0, 0x26, 0xfc, 0x31, 0xf7, 0x8c, 0x0e, 0xbc, 0x34, 0x00, 0x70, 0xe8, 0x00, 0xfb, 0x24,
	0x66, 0x38, 0x82, 0xcd, 0xd4, 0x92, 0x26, 0x29, 0xb2, 0x55, 0xbe, 0xa0, 0x9c, 0x9f, 0x25, 0xf6,
	0x51, 0xc9, 0x59, 0x8d, 0x54, 0x19, 0xe2, 0xd3, 0xad, 0xe5, 0xb5, 0xcd, 0x7a, 0x8e, 0x97, 0xde,
	0xcd, 0x85, 0xdf, 0x1b, 0x60, 0x4b, 0x34, 0x90, 0xce, 0x18, 0xd4, 0x7b, 0x45, 0x67, 0x2c, 0xd5,
	0x35, 0x8a, 0xb1, 0xb2, 0x9d, 0xfe, 0x07, 0xb6, 0x96, 0x59, 0x2b, 0xb6, 0x8e, 0xce, 0x18, 0x97,
	0xfd, 0x66, 0x80, 0x9b, 0x0e, 0x3d, 0xc6, 0x8c, 0x4f, 0xe0, 0x60, 0xe6, 0x92, 0x09, 0x14, 0x7d,
	0xca, 0x45, 0xa9, 0xf1, 0xd6, 0x8a, 0x2b, 0xd9, 0xab, 0x7e, 0xaf, 0xc4, 0x6c, 0x1d, 0x63, 0x26,
	0x84, 0xc8, 0x13, 0xd5, 0x31, 0x7a, 0x1e, 0xd1, 0x29, 0x72, 0x7a, 0x6b, 0xc5, 0x75, 0x98, 0x17,
	0xcb, 0xfb, 0xf9, 0xab, 0xae, 0x3a, 0x74, 0x88, 0xdd, 0x68, 0x7c, 0xc6, 0xe7, 0x8a, 0x65, 0x37,
	0xb3, 0x20, 0xd7, 0x4d, 0x3d, 0x56, 0xcc, 0x67, 0x89, 0xfd, 0xa0, 0xc4, 0x6c, 0xca, 0x42, 0xc4,
	0x63, 0x34, 0x27, 0xec, 0x4c, 0xad, 0xcc, 0xb5, 0xca, 0x5d, 0x73, 0x5b, 0x57, 0xf2, 0x86, 0xfe,
	0x24, 0x1a, 0x3a, 0x2c, 0x37, 0x74, 0x78, 0x65, 0x43, 0x87, 0x2b, 0x1b, 0x7a, 0x9a, 0xd8, 0x6f,
	0x97, 0xa4, 0x3b, 0x0f, 0x3c, 0xc2, 0xa4, 0x53, 0x96, 0xaf, 0x6f, 0xa4, 0xb9, 0x93, 0x12, 0x89,
	0x3b, 0x5d, 0x32, 0x7f, 0x36, 0xc0, 0xb6, 0x43, 0x1d, 0x3c, 0xc1, 0x0c, 0xf3, 0xc7, 0xc1, 0xdb,
	0x5c, 0x92, 0x8d, 0xf3, 0xcc, 0xfd, 0xd5, 0x17, 0x95, 0xf4, 0x93, 0xc4, 0x7e, 0xa7, 0x24, 0x85,
	0xb2, 0xfc, 0x6f, 0x59, 0x61, 0xa7, 0x64, 0x85, 0xdf, 0x8a, 0xcd, 0x7d, 0x8c, 0xd9, 0x80, 0x4e,
	0x70, 0x0c, 0x6b, 0x6a, 0x05, 0x8a, 0x51, 0xca, 0xdb, 0xcb, 0x87, 0x19, 0xeb, 0x3e, 0xd8, 0xe2,
	0xd9, 0xbf, 0x7b, 0xc9, 0x35, 0xb8, 0xab, 0xab, 0x22, 0xe1, 0xb8, 0x34, 0xc0, 0x6b, 0x4b, 0xd6,
	0x72, 0xc7, 0xf0, 0xc1, 0xea, 0x1d, 0xa3, 0x5f, 0x51, 0xbe, 0x4f, 0xf9, 0xc6, 0xce, 0x7c, 0x35,
	0xe5, 0xe3, 0x81, 0xf4, 0x11, 0x6f, 0xad, 0xed, 0x16, 0x6c, 0x16, 0x6d, 0xe9, 0x56, 0xf9, 0x4e,
	0xbc, 0xe1, 0x77, 0x23, 0xec, 0x32, 0x2c, 0x90, 0x62, 0xa7, 0x64, 0xe3, 0xdc, 0x4e, 0xd1, 0x63,
	0xc5, 0x7b, 0x92, 0xd8, 0x6f, 0x69, 0xbc, 0x9b, 0xb2, 0x04, 0x05, 0x78, 0x8e, 0xf8, 0x34, 0x6b,
	0x69, 0x0d, 0xb3, 0xdc, 0x36, 0xbe, 0xf2, 0xbe, 0x16, 0xae, 0xc7, 0xa1, 0x97, 0x73, 0x65, 0xe3,
	0xfc, 0x79, 0xa8, 0xc5, 0xca, 0x35, 0x48, 0xec, 0xbb, 0x9a, 0xeb, 0xba, 0x2c, 0xd9, 0xd0, 0x74,
	0xb8, 0xda, 0xf4, 0x95, 0xb6, 0x1b, 0x32, 0x53, 0x36, 0xce, 0x99, 0xf4, 0x58, 0x99, 0x3e, 0x2a,
	0x98, 0xd4, 0xda, 0xdf, 0xc8, 0x54, 0xeb, 0xac, 0x58, 0x5e, 0xbf, 0x18, 0xe0, 0xc6, 0xf2, 0x58,
	0x96, 0x2b, 0xbf, 0xa9, 0x9d, 0xbd, 0xb9, 0xd5, 0xdf, 0x2a, 0x5f, 0x50, 0xac, 0x71, 0x62, 0xdb,
	0x60, 0x3b, 0xcd, 0x05, 0xad, 0x5e, 0xda, 0x05, 0x83, 0x4d, 0x90, 0x85, 0x75, 0xc6, 0x77, 0x66,
	0xac, 0xa8, 0xbf, 0xa6, 0x54, 0xfe, 0x20, 0x79, 0x26, 0x37, 0xb5, 0x45, 0x2f, 0x92, 0x22, 0x55,
	0xbb, 0x90, 0x7d, 0x3e, 0x7a, 0x05, 0x6a, 0x43, 0xa3, 0xc6, 0xff, 0xdc, 0x2a, 0x94, 0x52, 0xcc,
	0x4f, 0xbd, 0xea, 0xf2, 0x70, 0x16, 0x2f, 0xba, 0xa1, 0x1d, 0xc0, 0xfa, 0x9b, 0x6e, 0x96, 0xf2,
	0x0c, 0x7a, 0xbf, 0x00, 0xad, 0x0d, 0xd3, 0xaf, 0x9c, 0x40, 0x8a, 0x09, 0xd7, 0x2a, 0xf7, 0xcd,
	0xab, 0x3a, 0xaa, 0x8e, 0x67, 0x09, 0x4d, 0x3b, 0xb5, 0x84, 0xa6, 0x41, 0x11, 0x9a, 0xe5, 0xff,
	0x2f, 0x54, 0x6b, 0x27, 0x87, 0xfe, 0x68, 0x80, 0x1d, 0x87, 0x3e, 0xc1, 0x11, 0x79, 0xbe, 0x78,
	0xff, 0xe3, 0x47, 0x8f, 0xe8, 0x39, 0x0e, 0xa4, 0x55, 0x66, 0x22, 0xc8, 0x59, 0x73, 0x79, 0x76,
	0x54, 0x1f, 0x81, 0xeb, 0x0f, 0xdd, 0xc0, 0xf5, 0x71, 0x24, 0x3f, 0x74, 0xb2, 0x0c, 0xf9, 0xe4,
	0x02, 0x07, 0xe8, 0xc5, 0x9c, 0x6d, 0xf4, 0xf3, 0x70, 0x2c, 0x9d, 0x2f, 0xe6, 0xcc, 0xba, 0x10,
	0xf7, 0xdf, 0x33, 0x3a, 0xbd, 0xd3, 0xc4, 0x1e, 0x42, 0x13, 0xdc, 0xee, 0x51, 0x7a, 0x4e, 0x02,
	0x1f, 0xf5, 0x3d, 0x7e, 0x2f, 0x5b, 0x20, 0x39, 0xdf, 0x14, 0x07, 0xec, 0xf0, 0xe5, 0x3b, 0xdd,
	0x37, 0x9f, 0x1e, 0x80, 0xd7, 0xaf, 0x98, 0xa1, 0xb6, 0x55, 0x69, 0x57, 0x79, 0x46, 0x23, 0xf2,
	0x85, 0x88, 0x50, 0xe5, 0xe9, 0x2b, 0xdd, 0xa3, 0x70, 0x34, 0xba, 0x26, 0x7e, 0x2d, 0xdf, 0xfd,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0x84, 0x43, 0x63, 0xf3, 0xb1, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BIMAPIClient is the client API for BIMAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BIMAPIClient interface {
	// User
	DoLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	DoRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
	DoLogout(ctx context.Context, in *UserLogoutRequest, opts ...grpc.CallOption) (*UserLogoutResponse, error)
	DoGetUserDetail(ctx context.Context, in *GetUserDetailRequest, opts ...grpc.CallOption) (*GetUserDetailResponse, error)
	DoSearchUsers(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error)
	DoSetUserDetail(ctx context.Context, in *SetUserDetailRequest, opts ...grpc.CallOption) (*SetUserDetailResponse, error)
	DoDeleteUser(ctx context.Context, in *DeleteUserDetailRequest, opts ...grpc.CallOption) (*DeleteUserDetailResponse, error)
	DoGetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*GetRolesResponse, error)
	DoGetRole(ctx context.Context, in *GetRoleDetailRequest, opts ...grpc.CallOption) (*GetRoleDetailResponse, error)
	DoCreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error)
	DoUpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error)
	DoDeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*DeleteRoleResponse, error)
	// UserRoleAPIs
	DoGetUserRoles(ctx context.Context, in *GetUserRolesRequest, opts ...grpc.CallOption) (*GetUserRolesResponse, error)
	DoGetRoleUsers(ctx context.Context, in *GetRoleUsersRequest, opts ...grpc.CallOption) (*GetRoleUsersResponse, error)
	DoSetUserRole(ctx context.Context, in *SetUserRoleRequest, opts ...grpc.CallOption) (*SetUserRoleResponse, error)
	DoSetRoleUser(ctx context.Context, in *SetRoleUserRequest, opts ...grpc.CallOption) (*SetRoleUserResponse, error)
	// JWT
	DoVerifyJWTToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error)
}

type bIMAPIClient struct {
	cc *grpc.ClientConn
}

func NewBIMAPIClient(cc *grpc.ClientConn) BIMAPIClient {
	return &bIMAPIClient{cc}
}

func (c *bIMAPIClient) DoLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoLogout(ctx context.Context, in *UserLogoutRequest, opts ...grpc.CallOption) (*UserLogoutResponse, error) {
	out := new(UserLogoutResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoGetUserDetail(ctx context.Context, in *GetUserDetailRequest, opts ...grpc.CallOption) (*GetUserDetailResponse, error) {
	out := new(GetUserDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoGetUserDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoSearchUsers(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error) {
	out := new(SearchUserResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoSearchUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoSetUserDetail(ctx context.Context, in *SetUserDetailRequest, opts ...grpc.CallOption) (*SetUserDetailResponse, error) {
	out := new(SetUserDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoSetUserDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoDeleteUser(ctx context.Context, in *DeleteUserDetailRequest, opts ...grpc.CallOption) (*DeleteUserDetailResponse, error) {
	out := new(DeleteUserDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoDeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoGetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*GetRolesResponse, error) {
	out := new(GetRolesResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoGetRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoGetRole(ctx context.Context, in *GetRoleDetailRequest, opts ...grpc.CallOption) (*GetRoleDetailResponse, error) {
	out := new(GetRoleDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoGetRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoCreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error) {
	out := new(CreateRoleResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoCreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoUpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error) {
	out := new(UpdateRoleResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoUpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoDeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*DeleteRoleResponse, error) {
	out := new(DeleteRoleResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoDeleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoGetUserRoles(ctx context.Context, in *GetUserRolesRequest, opts ...grpc.CallOption) (*GetUserRolesResponse, error) {
	out := new(GetUserRolesResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoGetUserRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoGetRoleUsers(ctx context.Context, in *GetRoleUsersRequest, opts ...grpc.CallOption) (*GetRoleUsersResponse, error) {
	out := new(GetRoleUsersResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoGetRoleUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoSetUserRole(ctx context.Context, in *SetUserRoleRequest, opts ...grpc.CallOption) (*SetUserRoleResponse, error) {
	out := new(SetUserRoleResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoSetUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoSetRoleUser(ctx context.Context, in *SetRoleUserRequest, opts ...grpc.CallOption) (*SetRoleUserResponse, error) {
	out := new(SetRoleUserResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoSetRoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bIMAPIClient) DoVerifyJWTToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error) {
	out := new(VerifyTokenResponse)
	err := c.cc.Invoke(ctx, "/pb.BIMAPI/DoVerifyJWTToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BIMAPIServer is the server API for BIMAPI service.
type BIMAPIServer interface {
	// User
	DoLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	DoRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
	DoLogout(context.Context, *UserLogoutRequest) (*UserLogoutResponse, error)
	DoGetUserDetail(context.Context, *GetUserDetailRequest) (*GetUserDetailResponse, error)
	DoSearchUsers(context.Context, *SearchUserRequest) (*SearchUserResponse, error)
	DoSetUserDetail(context.Context, *SetUserDetailRequest) (*SetUserDetailResponse, error)
	DoDeleteUser(context.Context, *DeleteUserDetailRequest) (*DeleteUserDetailResponse, error)
	DoGetRoles(context.Context, *GetRolesRequest) (*GetRolesResponse, error)
	DoGetRole(context.Context, *GetRoleDetailRequest) (*GetRoleDetailResponse, error)
	DoCreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error)
	DoUpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error)
	DoDeleteRole(context.Context, *DeleteRoleRequest) (*DeleteRoleResponse, error)
	// UserRoleAPIs
	DoGetUserRoles(context.Context, *GetUserRolesRequest) (*GetUserRolesResponse, error)
	DoGetRoleUsers(context.Context, *GetRoleUsersRequest) (*GetRoleUsersResponse, error)
	DoSetUserRole(context.Context, *SetUserRoleRequest) (*SetUserRoleResponse, error)
	DoSetRoleUser(context.Context, *SetRoleUserRequest) (*SetRoleUserResponse, error)
	// JWT
	DoVerifyJWTToken(context.Context, *VerifyTokenRequest) (*VerifyTokenResponse, error)
}

// UnimplementedBIMAPIServer can be embedded to have forward compatible implementations.
type UnimplementedBIMAPIServer struct {
}

func (*UnimplementedBIMAPIServer) DoLogin(ctx context.Context, req *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoLogin not implemented")
}
func (*UnimplementedBIMAPIServer) DoRegister(ctx context.Context, req *UserRegisterRequest) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoRegister not implemented")
}
func (*UnimplementedBIMAPIServer) DoLogout(ctx context.Context, req *UserLogoutRequest) (*UserLogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoLogout not implemented")
}
func (*UnimplementedBIMAPIServer) DoGetUserDetail(ctx context.Context, req *GetUserDetailRequest) (*GetUserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoGetUserDetail not implemented")
}
func (*UnimplementedBIMAPIServer) DoSearchUsers(ctx context.Context, req *SearchUserRequest) (*SearchUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSearchUsers not implemented")
}
func (*UnimplementedBIMAPIServer) DoSetUserDetail(ctx context.Context, req *SetUserDetailRequest) (*SetUserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSetUserDetail not implemented")
}
func (*UnimplementedBIMAPIServer) DoDeleteUser(ctx context.Context, req *DeleteUserDetailRequest) (*DeleteUserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoDeleteUser not implemented")
}
func (*UnimplementedBIMAPIServer) DoGetRoles(ctx context.Context, req *GetRolesRequest) (*GetRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoGetRoles not implemented")
}
func (*UnimplementedBIMAPIServer) DoGetRole(ctx context.Context, req *GetRoleDetailRequest) (*GetRoleDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoGetRole not implemented")
}
func (*UnimplementedBIMAPIServer) DoCreateRole(ctx context.Context, req *CreateRoleRequest) (*CreateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoCreateRole not implemented")
}
func (*UnimplementedBIMAPIServer) DoUpdateRole(ctx context.Context, req *UpdateRoleRequest) (*UpdateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoUpdateRole not implemented")
}
func (*UnimplementedBIMAPIServer) DoDeleteRole(ctx context.Context, req *DeleteRoleRequest) (*DeleteRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoDeleteRole not implemented")
}
func (*UnimplementedBIMAPIServer) DoGetUserRoles(ctx context.Context, req *GetUserRolesRequest) (*GetUserRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoGetUserRoles not implemented")
}
func (*UnimplementedBIMAPIServer) DoGetRoleUsers(ctx context.Context, req *GetRoleUsersRequest) (*GetRoleUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoGetRoleUsers not implemented")
}
func (*UnimplementedBIMAPIServer) DoSetUserRole(ctx context.Context, req *SetUserRoleRequest) (*SetUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSetUserRole not implemented")
}
func (*UnimplementedBIMAPIServer) DoSetRoleUser(ctx context.Context, req *SetRoleUserRequest) (*SetRoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSetRoleUser not implemented")
}
func (*UnimplementedBIMAPIServer) DoVerifyJWTToken(ctx context.Context, req *VerifyTokenRequest) (*VerifyTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoVerifyJWTToken not implemented")
}

func RegisterBIMAPIServer(s *grpc.Server, srv BIMAPIServer) {
	s.RegisterService(&_BIMAPI_serviceDesc, srv)
}

func _BIMAPI_DoLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoRegister(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoLogout(ctx, req.(*UserLogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoGetUserDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoGetUserDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoGetUserDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoGetUserDetail(ctx, req.(*GetUserDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoSearchUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoSearchUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoSearchUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoSearchUsers(ctx, req.(*SearchUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoSetUserDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoSetUserDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoSetUserDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoSetUserDetail(ctx, req.(*SetUserDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoDeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoDeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoDeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoDeleteUser(ctx, req.(*DeleteUserDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoGetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoGetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoGetRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoGetRoles(ctx, req.(*GetRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoGetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoGetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoGetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoGetRole(ctx, req.(*GetRoleDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoCreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoCreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoCreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoCreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoUpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoUpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoUpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoUpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoDeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoDeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoDeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoDeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoGetUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoGetUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoGetUserRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoGetUserRoles(ctx, req.(*GetUserRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoGetRoleUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoGetRoleUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoGetRoleUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoGetRoleUsers(ctx, req.(*GetRoleUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoSetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoSetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoSetUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoSetUserRole(ctx, req.(*SetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoSetRoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoSetRoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoSetRoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoSetRoleUser(ctx, req.(*SetRoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BIMAPI_DoVerifyJWTToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BIMAPIServer).DoVerifyJWTToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BIMAPI/DoVerifyJWTToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BIMAPIServer).DoVerifyJWTToken(ctx, req.(*VerifyTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BIMAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BIMAPI",
	HandlerType: (*BIMAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoLogin",
			Handler:    _BIMAPI_DoLogin_Handler,
		},
		{
			MethodName: "DoRegister",
			Handler:    _BIMAPI_DoRegister_Handler,
		},
		{
			MethodName: "DoLogout",
			Handler:    _BIMAPI_DoLogout_Handler,
		},
		{
			MethodName: "DoGetUserDetail",
			Handler:    _BIMAPI_DoGetUserDetail_Handler,
		},
		{
			MethodName: "DoSearchUsers",
			Handler:    _BIMAPI_DoSearchUsers_Handler,
		},
		{
			MethodName: "DoSetUserDetail",
			Handler:    _BIMAPI_DoSetUserDetail_Handler,
		},
		{
			MethodName: "DoDeleteUser",
			Handler:    _BIMAPI_DoDeleteUser_Handler,
		},
		{
			MethodName: "DoGetRoles",
			Handler:    _BIMAPI_DoGetRoles_Handler,
		},
		{
			MethodName: "DoGetRole",
			Handler:    _BIMAPI_DoGetRole_Handler,
		},
		{
			MethodName: "DoCreateRole",
			Handler:    _BIMAPI_DoCreateRole_Handler,
		},
		{
			MethodName: "DoUpdateRole",
			Handler:    _BIMAPI_DoUpdateRole_Handler,
		},
		{
			MethodName: "DoDeleteRole",
			Handler:    _BIMAPI_DoDeleteRole_Handler,
		},
		{
			MethodName: "DoGetUserRoles",
			Handler:    _BIMAPI_DoGetUserRoles_Handler,
		},
		{
			MethodName: "DoGetRoleUsers",
			Handler:    _BIMAPI_DoGetRoleUsers_Handler,
		},
		{
			MethodName: "DoSetUserRole",
			Handler:    _BIMAPI_DoSetUserRole_Handler,
		},
		{
			MethodName: "DoSetRoleUser",
			Handler:    _BIMAPI_DoSetRoleUser_Handler,
		},
		{
			MethodName: "DoVerifyJWTToken",
			Handler:    _BIMAPI_DoVerifyJWTToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apis.proto",
}
