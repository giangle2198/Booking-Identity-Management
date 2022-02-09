package repository

import (
	"booking-identity-management/config"
	"booking-identity-management/internal/common"
	"booking-identity-management/internal/helper"
	"booking-identity-management/internal/repository/model"
	"booking-identity-management/internal/util"
	"context"
	"database/sql"
	"errors"
	"strings"
)

type (
	UserActionRepository interface {
		// Add, Edit, Get User by ID or Domain, Check User Active by ID or Domain, Update First Login, Search Users, Search Users Total
		AddUser(ctx context.Context, user *model.User, isLogin bool) error
		EditUser(ctx context.Context, user *model.User) error
		CheckExistedUser(ctx context.Context, domain string) (bool, error)
		DeleteUser(ctx context.Context, userID uint32, domain string) error
		GetUserDetailByUserID(ctx context.Context, userID uint32) (*model.User, error)
		GetUserDetailByDomain(ctx context.Context, domain string) (*model.User, error)
		GetListUser(ctx context.Context, domain string) ([]*model.User, error)
		EditPassword(ctx context.Context, domain, password string) error
		CheckCurrentPassword(ctx context.Context, domain, currentPassword string) (bool, error)
	}

	userActionRepository struct {
		cfg *config.Config
		db  helper.DBHelper
	}
)

func NewUserActionRepository(cfg *config.Config,
	db helper.DBHelper) UserActionRepository {
	return &userActionRepository{
		cfg: cfg,
		db:  db,
	}
}

func (r *userActionRepository) AddUser(ctx context.Context, user *model.User, isLogin bool) error {
	var args []interface{}
	stmt := `INSERT INTO IM_USER (FULL_NAME, EMAIL, FIRST_NAME, LAST_NAME, PASSWORD, PHONE_NUMBER, DOMAIN, IS_ACTIVE)
		VALUES (?, ?, ?, ?, ?, ?, ?, '1')`
	tx, err := r.db.Begin()
	if err != nil {
		return errors.New(common.ReasonDBError.Code())
	}

	// hash passwords
	hashPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	args = append(args, user.FullName,
		user.Email,
		user.FirstName,
		user.LastName,
		hashPassword,
		user.PhoneNumber,
		user.Domain)
	result, err := tx.Exec(stmt, args...)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if num, err := result.RowsAffected(); num == 0 || err != nil {
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

func (r *userActionRepository) EditUser(ctx context.Context, user *model.User) error {
	var args []interface{}
	stmt := `UPDATE IM_USER  SET FULL_NAME=?
	, EMAIL = ?
	, FIRST_NAME = ?
	, LAST_NAME = ?
	, PHONE_NUMBER = ?
	, ADDRESS = ?
	, IS_ACTIVE = ?
	WHERE DOMAIN = ?`
	tx, err := r.db.Begin()
	if err != nil {
		return errors.New(common.ReasonDBError.Code())
	}
	args = append(args, user.FullName,
		user.Email,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Address,
		user.IsActive,
		user.Domain)
	result, err := tx.Exec(stmt, args...)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if num, err := result.RowsAffected(); num == 0 || err != nil {
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

func (r *userActionRepository) EditPassword(ctx context.Context, domain, password string) error {
	var args []interface{}
	stmt := `UPDATE IM_USER SET PASSWORD=?
	WHERE DOMAIN = ?`
	tx, err := r.db.Begin()
	if err != nil {
		return errors.New(common.ReasonDBError.Code())
	}

	// hash passwords
	hashPassword, err := util.HashPassword(password)
	if err != nil {
		return err
	}

	args = append(args,
		sql.Named("PASSWORD", hashPassword))
	result, err := tx.Exec(stmt, args...)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if num, err := result.RowsAffected(); num == 0 || err != nil {
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
func (r *userActionRepository) CheckExistedUser(ctx context.Context, domain string) (bool, error) {
	var num uint32
	stmt := `SELECT COUNT(ID) FROM IM_USER WHERE DOMAIN = ?`
	row := r.db.Open().QueryRow(stmt, domain)

	err := row.Scan(&num)
	if err != nil {
		return false, err
	}
	if num == 0 {
		return false, nil
	}
	return true, nil
}

func (r *userActionRepository) CheckCurrentPassword(ctx context.Context, domain, currentPassword string) (bool, error) {
	var pw string
	stmt := `SELECT PASSWORD FROM IM_USER WHERE DOMAIN = ? `
	row := r.db.Open().QueryRow(stmt, domain)

	err := row.Scan(&pw)
	if err != nil {
		return false, err
	}

	isCorrect := util.CheckPasswordHash(currentPassword, pw)
	if !isCorrect {
		return false, errors.New(common.ReasonNotCorrectPass.Code())
	}

	return true, nil
}
func (r *userActionRepository) DeleteUser(ctx context.Context, userID uint32, domain string) error {
	stmt := strings.Builder{}
	stmt.WriteString(`DELETE FROM IM_USER WHERE 1=1 `)
	var args []interface{}
	if userID != 0 {
		stmt.WriteString(` AND ID = ? `)
		args = append(args, userID)
	}
	if domain != "" {
		stmt.WriteString(` AND DOMAIN = ? `)
		args = append(args, domain)
	}
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec(stmt.String(), args...)
	if err != nil {
		return err
	}
	if num, err := result.RowsAffected(); err != nil || num == 0 {
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
func (r *userActionRepository) GetUserDetailByDomain(ctx context.Context, domain string) (user *model.User, err error) {
	stmt := `SELECT ID, FULL_NAME, EMAIL, FIRST_NAME, LAST_NAME, PHONE_NUMBER, DOMAIN, ADDRESS, IS_ACTIVE
	FROM IM_USER 
	WHERE DOMAIN = ? `
	row := r.db.Open().QueryRow(stmt, domain)

	user = &model.User{}
	err = row.Scan(&user.UserID,
		&user.FullName,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.Domain,
		&user.Address,
		&user.IsActive)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userActionRepository) GetUserDetailByUserID(ctx context.Context, userID uint32) (user *model.User, err error) {
	stmt := `SELECT ID, FULL_NAME, EMAIL, FIRST_NAME, LAST_NAME, PHONE_NUMBER, DOMAIN 
	FROM IM_USER 
	WHERE USER_ID = ?`
	row, err := r.db.Open().Query(stmt, userID)
	if err != nil {
		return nil, err
	}
	err = row.Scan(&user.UserID, &user.FullName, &user.Email, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.Domain)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userActionRepository) GetListUser(ctx context.Context, domain string) ([]*model.User, error) {
	queryStringBuilder := strings.Builder{}
	var args []interface{}
	queryStringBuilder.WriteString(`SELECT ID, FULL_NAME, EMAIL, FIRST_NAME, LAST_NAME, PHONE_NUMBER, DOMAIN 
	FROM IM_USER `)

	if domain != "" {
		queryStringBuilder.WriteString(` WHERE DOMAIN = ?`)
		args = append(args, domain)
	}

	rows, err := r.db.Open().Query(queryStringBuilder.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	listUser := make([]*model.User, 0)

	for rows.Next() {
		user := &model.User{}
		err = rows.Scan(&user.UserID, &user.FullName, &user.Email, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.Domain)
		if err != nil {
			return nil, err
		}
		listUser = append(listUser, user)
	}
	if len(listUser) < 1 {
		return nil, errors.New(common.ReasonNotFound.Code())
	}
	return listUser, nil
}
