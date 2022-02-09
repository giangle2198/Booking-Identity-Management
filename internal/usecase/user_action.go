package usecase

import (
	"booking-identity-management/config"
	"booking-identity-management/internal/adapter"
	"booking-identity-management/internal/common"
	"booking-identity-management/internal/dto"
	"booking-identity-management/internal/repository"
	"booking-identity-management/internal/repository/model"
	"context"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type (
	UserActionUsecase interface {
		Login(ctx context.Context, req *dto.LoginRequest) (string, string, error)
		Logout(ctx context.Context, token string) error
		Register(ctx context.Context, req *dto.RegisterRequest) error
		VerifyToken(ctx context.Context, tokenStr, uid string) (dto.UserTokenDTO, error)
		SearchUsers(ctx context.Context, domain string) ([]*dto.UserDetailDTO, uint32, error)
		SetUserDetail(ctx context.Context, userDetail *dto.UserDetailDTO) error
		SetUserPassword(ctx context.Context, domain, oldPassword, password string) error
		GetUserDetail(ctx context.Context, domain string) (*dto.UserDetailDTO, error)
		DeleteUserDetail(ctx context.Context, userID uint32, domain string) error
	}

	userActionUsecase struct {
		cfg                     *config.Config
		userActionRepository    repository.UserActionRepository
		userAccountRepository   repository.UserAccountRepository
		jwtAdapter              adapter.JWTAdapter
		jwtManagementRepository repository.JWTManagementRepository
		pasetoAdapter           adapter.PasetoAdapter
	}
)

func NewUserActionUsecase(cfg *config.Config,
	userActionRepository repository.UserActionRepository,
	userAccountRepository repository.UserAccountRepository,
	jwtAdapter adapter.JWTAdapter,
	jwtManagementRepository repository.JWTManagementRepository,
	pasetoAdapter adapter.PasetoAdapter,
) UserActionUsecase {
	return &userActionUsecase{
		cfg:                     cfg,
		userActionRepository:    userActionRepository,
		userAccountRepository:   userAccountRepository,
		jwtAdapter:              jwtAdapter,
		pasetoAdapter:           pasetoAdapter,
		jwtManagementRepository: jwtManagementRepository,
	}
}

func (u *userActionUsecase) Login(ctx context.Context, req *dto.LoginRequest) (string, string, error) {

	// get user detail
	userDetail, err := u.userActionRepository.GetUserDetailByDomain(ctx, req.Domain)
	if err != nil {
		zap.S().Errorf("Error while getting user detail: %v", err)
		return "", "", err
	}
	if !userDetail.IsActive {
		zap.S().Errorf("User: %v is unactive status: %v", req.Domain, err)
		return "", "", errors.New(common.ReasonUserNotActive.Code())
	}

	isCorrectPass, err := u.userActionRepository.CheckCurrentPassword(ctx, req.Domain, req.Password)
	if err != nil || !isCorrectPass {
		zap.S().Errorf("Error while checking user password: %v", err)
		return "", "", errors.New(common.ReasonNotCorrectPass.Code())
	}

	token, _, err := u.jwtAdapter.GenerateToken(ctx, uint32(userDetail.UserID), userDetail.Domain)
	if err != nil {
		zap.S().Errorf("Error while generating token: %v", err)
		return "", "", errors.New(common.ReasonJWTError.Code())
	}

	_, err = u.jwtManagementRepository.AddToken(ctx, token, uint32(userDetail.UserID))
	if err != nil {
		zap.S().Errorf("Error while adding token to db: %v", err)
		return "", "", errors.New(common.ReasonJWTError.Code())
	}

	return token, "", nil
}

func (u *userActionUsecase) Logout(ctx context.Context, token string) error {

	claim, err := u.jwtAdapter.VerifyToken(ctx, token, "bim", "")
	if err != nil {
		zap.S().Errorf("Error while verifying token: %v", err)
		return errors.New(common.ReasonJWTError.Code())
	}

	// get user detail
	_, err = u.userActionRepository.GetUserDetailByDomain(ctx, claim.Domain)
	if err != nil {
		zap.S().Errorf("Error while getting user detail: %v", err)
		return err
	}

	err = u.jwtManagementRepository.DeleteToken(ctx, 0, token)
	if err != nil {
		zap.S().Errorf("Error while deleting user detail: %v", err)
		return errors.New(common.ReasonJWTError.Code())
	}

	return nil
}

func (u *userActionUsecase) Register(ctx context.Context, req *dto.RegisterRequest) error {

	// get user detail
	isExistedUser, err := u.userActionRepository.CheckExistedUser(ctx, req.Domain)
	if err != nil {
		zap.S().Errorf("Error while checking existed user detail: %v", err)
		return err
	}

	if isExistedUser {
		zap.S().Errorf("Found user existed in db: %v", err)
		return errors.New(common.ReasonExistedUser.Code())
	}

	err = u.userActionRepository.AddUser(ctx, &model.User{
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Domain:      req.Domain,
		Password:    req.Password,
		FullName:    req.FirstName + " " + req.LastName,
	}, false)
	if err != nil {
		zap.S().Errorf("Error while registing user to db: %v", err)
		return errors.New(common.ReasonFailedRegisterUser.Code())
	}
	return nil
}

func (u *userActionUsecase) SearchUsers(ctx context.Context, domain string) ([]*dto.UserDetailDTO, uint32, error) {

	listUser, err := u.userActionRepository.GetListUser(ctx, domain)
	if err != nil {
		zap.S().Errorf("Error while getting list user to db: %v", err)
		return nil, 0, errors.New(common.ReasonDBError.Code())
	}

	var listUserDTO []*dto.UserDetailDTO

	for _, user := range listUser {
		userItem := &dto.UserDetailDTO{
			UserID:      user.UserID,
			Domain:      user.Domain,
			FullName:    user.FullName,
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Address:     user.Address,
			PhoneNumber: user.PhoneNumber,
			IsActive:    user.IsActive,
		}
		listUserDTO = append(listUserDTO, userItem)
	}

	if len(listUser) < 1 {
		return nil, 0, errors.New(common.ReasonNotFound.Code())
	}

	return listUserDTO, uint32(len(listUserDTO)), nil
}

func (u *userActionUsecase) SetUserDetail(ctx context.Context, userDetail *dto.UserDetailDTO) error {

	err := u.userActionRepository.EditUser(ctx, &model.User{
		UserID:      userDetail.UserID,
		FullName:    userDetail.FullName,
		Email:       userDetail.Email,
		FirstName:   userDetail.FirstName,
		LastName:    userDetail.LastName,
		PhoneNumber: userDetail.PhoneNumber,
		Domain:      userDetail.Domain,
		IsActive:    userDetail.IsActive,
		Address:     userDetail.Address,
	})

	if err != nil {
		zap.S().Errorf("Error while setting user detail to db: %v", err)
		return errors.New(common.ReasonDBError.Code())
	}

	return nil
}

func (u *userActionUsecase) GetUserDetail(ctx context.Context, domain string) (*dto.UserDetailDTO, error) {

	if domain == "" {
		user, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("not existed user")
		}
		domain = user[common.DomainMDKey][0]
	}

	userDetail, err := u.userActionRepository.GetUserDetailByDomain(ctx, domain)
	if err != nil {
		zap.S().Errorf("Error while getting user detail: %v", err)
		return nil, err
	}

	return &dto.UserDetailDTO{
		UserID:      userDetail.UserID,
		FullName:    userDetail.FullName,
		Email:       userDetail.Email,
		FirstName:   userDetail.FirstName,
		LastName:    userDetail.LastName,
		PhoneNumber: userDetail.PhoneNumber,
		Domain:      userDetail.Domain,
		IsActive:    userDetail.IsActive,
		Address:     userDetail.Address,
	}, nil
}

func (u *userActionUsecase) SetUserPassword(ctx context.Context, domain, oldPassword, newPassword string) error {

	_, err := u.userActionRepository.CheckCurrentPassword(ctx, domain, oldPassword)
	if err != nil {
		zap.S().Errorf("Error while checking password user of %v: %v", domain, err)
		return err
	}

	err = u.userActionRepository.EditPassword(ctx, domain, newPassword)
	if err != nil {
		zap.S().Errorf("Error while setting password user of %v: %v", domain, err)
		return errors.New(common.ReasonDBError.Code())
	}

	return nil
}

func (u *userActionUsecase) DeleteUserDetail(ctx context.Context, userID uint32, domain string) error {

	err := u.userActionRepository.DeleteUser(ctx, userID, domain)
	if err != nil {
		zap.S().Errorf("Error while deleting password user of %v: %v", userID, err)
		return err
	}

	return nil
}

func (u *userActionUsecase) VerifyToken(ctx context.Context, tokenStr, uid string) (dto.UserTokenDTO, error) {

	var userToken dto.UserTokenDTO
	claim, err := u.jwtAdapter.VerifyToken(ctx, tokenStr, uid, "")
	if err != nil {
		return userToken, err
	}
	return dto.UserTokenDTO{
		UserID: claim.UserID,
		Domain: claim.Domain,
	}, nil
}
