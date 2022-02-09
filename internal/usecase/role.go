package usecase

import (
	"booking-identity-management/config"
	"booking-identity-management/internal/common"
	"booking-identity-management/internal/dto"
	"booking-identity-management/internal/repository"
	"booking-identity-management/internal/repository/model"
	"context"
	"errors"

	"go.uber.org/zap"
)

type (
	RoleUsecase interface {
		CreateRole(ctx context.Context, role *dto.RoleDTO) (uint32, error)
		EditRole(ctx context.Context, roleDTO *dto.RoleDTO) error
		GetAllRoles(ctx context.Context) ([]*model.Role, error)
		GetRoleDetail(ctx context.Context, roleID uint32) (*model.Role, bool, error)
	}

	roleUsecase struct {
		cfg            *config.Config
		roleRepository repository.RoleRepository
	}
)

func NewRoleUsecase(cfg *config.Config,
	roleRepository repository.RoleRepository) RoleUsecase {
	return &roleUsecase{
		cfg:            cfg,
		roleRepository: roleRepository,
	}
}

func (u *roleUsecase) CreateRole(ctx context.Context, role *dto.RoleDTO) (uint32, error) {

	outID, err := u.roleRepository.CreateRole(ctx, role.RoleName, role.RoleDescription, role.RoleManagerDomain, role.RoleAlias)
	if err != nil {
		zap.S().Errorf("Error while creating role detail to db: %v", err)
		return outID, errors.New(common.ReasonDBError.Code())
	}

	return outID, nil
}
func (u *roleUsecase) EditRole(ctx context.Context, roleDTO *dto.RoleDTO) error {
	_, err := u.roleRepository.EditRole(ctx, roleDTO.ID, roleDTO.RoleName, roleDTO.RoleDescription, roleDTO.RoleManagerDomain, roleDTO.RoleAlias)
	if err != nil {
		zap.S().Errorf("Error while creating role detail to db: %v", err)
		return errors.New(common.ReasonDBError.Code())
	}

	return nil
}
func (u *roleUsecase) GetAllRoles(ctx context.Context) ([]*model.Role, error) {

	listRole, err := u.roleRepository.GetAllRoles(ctx)
	if err != nil {
		zap.S().Errorf("Error while getting list role to db: %v", err)
		return nil, errors.New(common.ReasonDBError.Code())
	}

	return listRole, nil
}

func (u *roleUsecase) GetRoleDetail(ctx context.Context, roleID uint32) (*model.Role, bool, error) {

	roleDetail, isFound, err := u.roleRepository.GetRoleDetail(ctx, roleID)
	if err != nil {
		zap.S().Errorf("Error while getting list role to db: %v", err)
		return nil, isFound, errors.New(common.ReasonDBError.Code())
	}

	return roleDetail, true, nil
}
