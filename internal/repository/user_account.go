package repository

import (
	"booking-identity-management/config"
	"booking-identity-management/internal/common"
	"booking-identity-management/internal/helper"
	"booking-identity-management/internal/repository/model"
	"context"
	"errors"
)

type (
	UserAccountRepository interface {
		CheckExistedUser(ctx context.Context, domain string) error
		VerifyAccount(ctx context.Context, username, password string) (bool, error)
	}

	userAccountRepository struct {
		cfg *config.Config
		db  helper.DBHelper
	}
)

func NewUserAccount(cfg *config.Config,
	db helper.DBHelper) UserAccountRepository {
	return &userAccountRepository{
		cfg: cfg,
		db:  db,
	}
}

func (r *userAccountRepository) CheckExistedUser(ctx context.Context, domain string) error {
	var value int32
	statement := `select count(domain) from IM_USER_ACCOUNT where domain = :1`

	row := r.db.Open().QueryRow(statement, domain)

	err := row.Scan(&value)
	if err != nil {
		return errors.New(common.ReasonDBError.Code())
	}

	if value == 0 {
		return errors.New(common.ReasonNotFound.Code())
	}

	return nil
}

func (r *userAccountRepository) VerifyAccount(ctx context.Context, username, password string) (bool, error) {
	var currentUser *model.UserAccount
	statement := `select domain, password from IM_USER where domain = :1`

	row := r.db.Open().QueryRow(statement, username)

	err := row.Scan(&currentUser.Domain, &currentUser.Password)
	if err != nil {
		return false, errors.New(common.ReasonDBError.Code())
	}

	if currentUser.Password != password {
		return false, errors.New(common.ReasonNotValid.Code())
	}

	return true, nil
}
