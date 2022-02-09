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
	JWTManagementRepository interface {
		AddToken(ctx context.Context, token string, userID uint32) (uint32, error)
		DeleteToken(ctx context.Context, tokenID uint32, token string) error
		CheckExitstedToken(ctx context.Context, tokenID uint32, token string) (bool, error)
		GetListTokenBeforeDate(ctx context.Context, createdTime string) ([]*model.JWTManagment, error)
		DeleteTokensBeforeDate(ctx context.Context, createdTime string) error
	}

	jwtManagementRepository struct {
		cfg *config.Config
		db  helper.DBHelper
	}
)

func NewJWTManagement(cfg *config.Config,
	db helper.DBHelper) JWTManagementRepository {
	return &jwtManagementRepository{
		cfg: cfg,
		db:  db,
	}
}

func (r *jwtManagementRepository) AddToken(ctx context.Context, token string, userID uint32) (uint32, error) {
	var args []interface{}
	stmt := `INSERT INTO IM_JWT_MANAGEMENT (USER_ID, TOKEN, CREATED_TIME)
		VALUES (?, ?, CURRENT_TIMESTAMP) `
	tx, err := r.db.Begin()
	if err != nil {
		return 0, errors.New(common.ReasonDBError.Code())
	}
	var outID int64

	args = append(args, userID,
		token,
	)
	result, err := tx.Exec(stmt, args...)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	if num, err := result.RowsAffected(); num == 0 || err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	outID, _ = result.LastInsertId()
	return uint32(outID), nil
}

func (r *jwtManagementRepository) DeleteToken(ctx context.Context, tokenID uint32, token string) error {
	var args []interface{}
	stmt := `DELETE FROM IM_JWT_MANAGEMENT WHERE TOKEN = ?`

	if token == "" {
		return errors.New(common.ReasonNotFound.Code())
	}

	tx, err := r.db.Begin()
	if err != nil {
		return errors.New(common.ReasonDBError.Code())
	}

	args = append(args, token)
	_, err = tx.Exec(stmt, args...)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return nil
}

func (r *jwtManagementRepository) CheckExitstedToken(ctx context.Context, tokenID uint32, token string) (bool, error) {
	var args []interface{}
	stmt := `SELECT COUNT(ID) FROM IM_JWT_MANAGEMENT WHERE ID = ?`
	tx, err := r.db.Begin()
	if err != nil {
		return false, errors.New(common.ReasonDBError.Code())
	}

	args = append(args, tokenID)
	_, err = tx.Exec(stmt, args...)
	if err != nil {
		_ = tx.Rollback()
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return false, err
	}
	return true, nil
}

func (r *jwtManagementRepository) GetListTokenBeforeDate(ctx context.Context, createdTime string) ([]*model.JWTManagment, error) {

	stmt := `SELECT ID, USER_ID, TOKEN, CREATED_TIME 
	FROM IM_JWT_MANAGEMENT 
	WHERE TO_CHAR(CREATED_TIME, 'YYYYMMDD') < :1
	ORDER BY CREATED_TIME desc`
	rows, err := r.db.Open().Query(stmt)
	if err != nil {
		// TODO: log
		return nil, err
	}
	defer rows.Close()
	var tokens []*model.JWTManagment
	for rows.Next() {
		token := &model.JWTManagment{}
		err = rows.Scan(&token.ID, &token.UserID, &token.Token, &token.CreatedDate)
		if err != nil {
			// TODO: log
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (r *jwtManagementRepository) DeleteTokensBeforeDate(ctx context.Context, createdTime string) error {
	var args []interface{}
	stmt := `DELETE FROM IM_JWT_MANAGEMENT WHERE TO_CHAR(CREATED_TIME, 'YYYYMMDD') < :1`
	tx, err := r.db.Begin()
	if err != nil {
		return errors.New(common.ReasonDBError.Code())
	}

	args = append(args, createdTime)
	_, err = tx.Exec(stmt, args...)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return nil
}
