package usecase

import (
	"booking-identity-management/config"
	"booking-identity-management/internal/common"
	"booking-identity-management/internal/repository"
	"context"
	"errors"

	"go.uber.org/zap"
)

type (
	BaseUsecase interface {
		DeleteRole(ctx context.Context, roleID uint32) error
	}

	baseUsecase struct {
		cfg                   *config.Config
		baseRepository        repository.BaseRepository
		roleRepository        repository.RoleRepository
		userAccountRepository repository.UserAccountRepository
		userRepository        repository.UserActionRepository
	}
)

func NewBaseUsecase(cfg *config.Config,
	baseRepository repository.BaseRepository,
	roleRepository repository.RoleRepository,
	userAccountRepository repository.UserAccountRepository,
	userRepository repository.UserActionRepository) BaseUsecase {
	return &baseUsecase{
		cfg:                   cfg,
		baseRepository:        baseRepository,
		roleRepository:        roleRepository,
		userAccountRepository: userAccountRepository,
		userRepository:        userRepository,
	}
}

func (u *baseUsecase) DeleteRole(ctx context.Context, roleID uint32) error {

	txHandler, err := u.baseRepository.BuildTransactions(ctx, "")
	if err != nil {
		zap.S().Errorf("Error while building transaction for db: %v", err)
		return errors.New(common.ReasonDBError.Code())
	}

	deleteRoleStep := func(tx *repository.TransactionExt) error {
		return u.roleRepository.DeleteRoleTx(ctx, tx, roleID)
	}

	err = txHandler.SetTransactionStep(deleteRoleStep)
	if err != nil {
		zap.S().Errorf("Error while running transation in deleting role detail to db: %v", err)
		return errors.New(common.ReasonDBError.Code())
	}

	err = txHandler.ExecThenCommit()
	if err != nil {
		zap.S().Errorf("Error while committing transaction: %v", err)
		return errors.New(common.ReasonDBError.Code())
	}

	return nil
}
