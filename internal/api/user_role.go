package api

import (
	"booking-identity-management/internal/common"
	"booking-identity-management/internal/dto"
	"booking-identity-management/pb"
	"context"
)

func (s *api) DoGetUserRoles(ctx context.Context, req *pb.GetUserRolesRequest) (res *pb.GetUserRolesResponse, err error) {
	var (
		resDTO *dto.GetUserRolesResponseDTO
	)
	resDTO, err = s.userRoleUsecase.GetUserRoles(ctx, req.Domain)
	if err != nil {
		return &pb.GetUserRolesResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res = &pb.GetUserRolesResponse{}
	s.pbConverter.ToPb(res, resDTO)
	return res, nil
}

func (s *api) DoGetRoleUsers(ctx context.Context, req *pb.GetRoleUsersRequest) (res *pb.GetRoleUsersResponse, err error) {
	var (
		resDTO *dto.GetRoleUsersResponseDTO
	)
	resDTO, err = s.userRoleUsecase.GetRoleUsers(ctx, req.RoleID)
	if err != nil {
		return &pb.GetRoleUsersResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res = &pb.GetRoleUsersResponse{}
	s.pbConverter.ToPb(res, resDTO)
	return res, nil
}

func (s *api) DoSetUserRole(ctx context.Context, req *pb.SetUserRoleRequest) (res *pb.SetUserRoleResponse, err error) {
	var (
		userRoles *dto.UserRoleDTO
	)

	err = s.userRoleUsecase.SetUserRoles(ctx, userRoles)
	if err != nil {
		return &pb.SetUserRoleResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res = &pb.SetUserRoleResponse{}
	res.StatusCode = common.DONE
	return res, nil
}

func (s *api) DoSetRoleUser(ctx context.Context, req *pb.SetRoleUserRequest) (res *pb.SetRoleUserResponse, err error) {
	var (
		userRoles *dto.RoleUserDTO
	)

	err = s.userRoleUsecase.SetRoleUsers(ctx, userRoles)
	if err != nil {
		return &pb.SetRoleUserResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res = &pb.SetRoleUserResponse{}
	res.StatusCode = common.DONE
	return res, nil
}
