package api

import (
	"booking-identity-management/internal/common"
	"booking-identity-management/internal/dto"
	"booking-identity-management/internal/repository/model"
	"booking-identity-management/pb"
	"context"
)

func (s *api) DoGetRoles(ctx context.Context, req *pb.GetRolesRequest) (res *pb.GetRolesResponse, err error) {
	var (
		roleDetails []*model.Role
	)
	roleDetails, err = s.roleUsecase.GetAllRoles(ctx)
	if err != nil {
		return &pb.GetRolesResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res = &pb.GetRolesResponse{}
	for _, roleDetail := range roleDetails {
		res.Roles = append(res.Roles, &pb.RoleDetail{
			RoleID:            roleDetail.ID,
			RoleName:          roleDetail.RoleName,
			RoleDescription:   roleDetail.RoleDescription.String,
			RoleManagerDomain: roleDetail.RoleManagerDomain.String,
			RoleAlias:         roleDetail.RoleAlias.String,
			IsActive:          roleDetail.IsActive.Bool,
		})
	}
	res.StatusCode = common.DONE
	return res, nil
}

func (s *api) DoCreateRole(ctx context.Context, req *pb.CreateRoleRequest) (res *pb.CreateRoleResponse, err error) {
	var (
		reqDTO = &dto.RoleDTO{}
		id     uint32
	)
	res = &pb.CreateRoleResponse{}

	s.pbConverter.FromPb(reqDTO, req)
	id, err = s.roleUsecase.CreateRole(ctx, reqDTO)
	if err != nil {
		return &pb.CreateRoleResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res.StatusCode = common.DONE
	res.RoleID = id
	return res, nil
}

func (s *api) DoUpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (res *pb.UpdateRoleResponse, err error) {
	var (
		reqDTO = &dto.RoleDTO{}
	)
	res = &pb.UpdateRoleResponse{}

	s.pbConverter.FromPb(reqDTO, req)
	err = s.roleUsecase.EditRole(ctx, reqDTO)
	if err != nil {
		return &pb.UpdateRoleResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res.StatusCode = common.DONE
	return res, nil
}

func (s *api) DoDeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (res *pb.DeleteRoleResponse, err error) {

	err = s.baseUsecase.DeleteRole(ctx, req.ID)
	if err != nil {
		return &pb.DeleteRoleResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res = &pb.DeleteRoleResponse{}
	res.StatusCode = common.DONE
	return res, nil
}

func (s *api) DoGetRole(ctx context.Context, req *pb.GetRoleDetailRequest) (res *pb.GetRoleDetailResponse, err error) {
	var (
		roleDetail *model.Role
	)
	roleDetail, _, err = s.roleUsecase.GetRoleDetail(ctx, req.RoleID)
	if err != nil {
		return &pb.GetRoleDetailResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res = &pb.GetRoleDetailResponse{
		RoleDetail: &pb.RoleDetail{},
	}
	s.pbConverter.ToPb(res.RoleDetail, roleDetail)
	res.RoleDetail.RoleID = roleDetail.ID
	res.StatusCode = common.DONE
	return res, nil
}
