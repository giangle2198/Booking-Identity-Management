package api

import (
	"booking-identity-management/internal/common"
	"booking-identity-management/internal/dto"
	"booking-identity-management/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *api) DoLogin(ctx context.Context, req *pb.UserLoginRequest) (res *pb.UserLoginResponse, err error) {
	userDTO := &dto.LoginRequest{
		Domain:   req.Domain,
		Password: req.Password,
	}
	res = &pb.UserLoginResponse{}

	token, paseto, err := s.userActionUsecase.Login(ctx, userDTO)
	if err != nil {
		res.StatusCode = common.FAILED
		res.ReasonCode = common.ParseError(err).Code()
		res.ReasonMessage = common.ParseError(err).Message()
		return res, nil
	}
	res.Jwt = token
	res.Paseto = paseto
	res.StatusCode = common.DONE
	return res, nil
}

func (s *api) DoRegister(ctx context.Context, req *pb.UserRegisterRequest) (res *pb.UserRegisterResponse, err error) {
	var (
		reqDTO = &dto.RegisterRequest{}
	)
	res = &pb.UserRegisterResponse{}
	s.pbConverter.FromPb(reqDTO, req)
	err = s.userActionUsecase.Register(ctx, reqDTO)
	if err != nil {
		return &pb.UserRegisterResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res.StatusCode = common.DONE
	return res, nil
}

func (s *api) DoLogout(ctx context.Context, req *pb.UserLogoutRequest) (res *pb.UserLogoutResponse, err error) {

	err = s.userActionUsecase.Logout(ctx, req.Token)
	if err != nil {
		return &pb.UserLogoutResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res = &pb.UserLogoutResponse{}
	res.StatusCode = common.DONE
	return res, nil
}

func (s *api) DoGetUserDetail(ctx context.Context, req *pb.GetUserDetailRequest) (res *pb.GetUserDetailResponse, err error) {
	var (
		userDetail *dto.UserDetailDTO
	)
	userDetail, err = s.userActionUsecase.GetUserDetail(ctx, req.Domain)
	if err != nil {
		return &pb.GetUserDetailResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res = &pb.GetUserDetailResponse{}
	res.UserDetails = &pb.UserDetails{
		FullName:    userDetail.FullName,
		Email:       userDetail.Email,
		FirstName:   userDetail.FirstName,
		LastName:    userDetail.LastName,
		PhoneNumber: userDetail.PhoneNumber,
		Domain:      userDetail.Domain,
		IsActive:    userDetail.IsActive,
		Address:     userDetail.Address,
	}
	res.StatusCode = common.DONE
	return res, nil
}

func (s *api) DoSearchUsers(ctx context.Context, req *pb.SearchUserRequest) (res *pb.SearchUserResponse, err error) {
	var (
		userDetails []*dto.UserDetailDTO
	)
	userDetails, _, err = s.userActionUsecase.SearchUsers(ctx, req.Domain)
	if err != nil {
		return &pb.SearchUserResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res = &pb.SearchUserResponse{}
	for _, userDetail := range userDetails {
		res.UserDetails = append(res.UserDetails, &pb.UserDetails{
			Domain:      userDetail.Domain,
			Email:       userDetail.Email,
			PhoneNumber: userDetail.PhoneNumber,
			FirstName:   userDetail.FirstName,
			LastName:    userDetail.LastName,
			FullName:    userDetail.FullName,
			Address:     userDetail.Address,
			IsActive:    userDetail.IsActive,
		})
	}
	res.StatusCode = common.DONE
	return res, nil
}

func (s *api) DoSetUserDetail(ctx context.Context, req *pb.SetUserDetailRequest) (res *pb.SetUserDetailResponse, err error) {
	var (
		reqDTO = &dto.UserDetailDTO{}
	)
	s.pbConverter.FromPb(reqDTO, req)
	err = s.userActionUsecase.SetUserDetail(ctx, reqDTO)
	if err != nil {
		return &pb.SetUserDetailResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res = &pb.SetUserDetailResponse{}
	res.StatusCode = common.DONE
	return res, nil
}

func (s *api) DoDeleteUser(ctx context.Context, req *pb.DeleteUserDetailRequest) (res *pb.DeleteUserDetailResponse, err error) {

	err = s.userActionUsecase.DeleteUserDetail(ctx, uint32(req.UserID), req.Domain)
	if err != nil {
		return &pb.DeleteUserDetailResponse{
			StatusCode:    common.FAILED,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}, nil
	}
	res = &pb.DeleteUserDetailResponse{}
	res.StatusCode = common.DONE
	return res, nil
}

func (s *api) DoVerifyJWTToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	resp := &pb.VerifyTokenResponse{}
	userToken, uErr := s.userActionUsecase.VerifyToken(ctx, req.Jwt, req.Uid)
	if uErr != nil {
		reason := common.ParseError(uErr)
		resp.ReasonCode = reason.Code()
		resp.ReasonMessage = reason.Message()
		resp.StatusCode = common.FAILED
		return resp, status.Error(codes.Unauthenticated, resp.ReasonCode)
	}
	resp.Domain = userToken.Domain
	resp.StatusCode = common.DONE
	return resp, nil
}
