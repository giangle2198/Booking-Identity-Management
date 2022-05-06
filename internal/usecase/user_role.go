package usecase

import (
	"booking-identity-management/config"
	"booking-identity-management/internal/dto"
	"booking-identity-management/internal/repository"
	"context"
)

type (
	UserRoleUsecase interface {
		GetUserRoles(ctx context.Context, domain string) (*dto.GetUserRolesResponseDTO, error)
		GetRoleUsers(ctx context.Context, roleID uint32) (*dto.GetRoleUsersResponseDTO, error)
		SetUserRoles(ctx context.Context, userRoles *dto.UserRoleDTO) error
		SetRoleUsers(ctx context.Context, userRoles *dto.RoleUserDTO) error
		UpdateUserRole(ctx context.Context, userRoleID uint32, isActive bool) error
		DeleteUserRole(ctx context.Context, userRoleID uint32) error
	}

	userRoleUsecase struct {
		baseUsecase           BaseUsecase
		cfg                   *config.Config
		roleRepository        repository.RoleRepository
		userAccountRepository repository.UserAccountRepository
		userRepository        repository.UserActionRepository
	}
)

func NewUserRoleUsecase(baseUsecase BaseUsecase,
	cfg *config.Config,
	roleRepository repository.RoleRepository,
	userAccountRepository repository.UserAccountRepository,
	userRepository repository.UserActionRepository) UserRoleUsecase {
	return &userRoleUsecase{
		baseUsecase:           baseUsecase,
		cfg:                   cfg,
		roleRepository:        roleRepository,
		userAccountRepository: userAccountRepository,
		userRepository:        userRepository,
	}
}

func (u *userRoleUsecase) GetUserRoles(ctx context.Context, domain string) (*dto.GetUserRolesResponseDTO, error) {
	return nil, nil
}
func (u *userRoleUsecase) GetRoleUsers(ctx context.Context, roleID uint32) (*dto.GetRoleUsersResponseDTO, error) {
	return nil, nil
}
func (u *userRoleUsecase) SetUserRoles(ctx context.Context, userRoles *dto.UserRoleDTO) error {
	return nil
}
func (u *userRoleUsecase) SetRoleUsers(ctx context.Context, userRoles *dto.RoleUserDTO) error {
	return nil
}
func (u *userRoleUsecase) UpdateUserRole(ctx context.Context, userRoleID uint32, isActive bool) error {
	return nil
}
func (u *userRoleUsecase) DeleteUserRole(ctx context.Context, userRoleID uint32) error {
	return nil
}
