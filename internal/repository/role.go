package repository

import (
	"booking-identity-management/config"
	"booking-identity-management/internal/common"
	"booking-identity-management/internal/helper"
	"booking-identity-management/internal/repository/model"
	"errors"

	// "booking-libs/log"
	"context"
	"strings"
	// "booking-libs/opentracing/jaeger"
)

type (
	// RoleRepository interface for role repository
	RoleRepository interface {
		CreateRole(ctx context.Context, roleName, roleDescription, roleManagerDomain, roleAlias string) (uint32, error)
		EditRole(ctx context.Context, roleID uint32, roleName, roleDescription, roleManagerDomain, roleAlias string) (uint32, error)
		DeleteRoleTx(ctx context.Context, tx *TransactionExt, roleID uint32) error
		GetAllRoles(ctx context.Context) ([]*model.Role, error)
		GetRoleDetail(ctx context.Context, roleID uint32) (*model.Role, bool, error)
	}
	roleRepository struct {
		cfg      *config.Config
		dbHelper helper.DBHelper
	}
)

// NewRoleRepository new instance for Role Repository
func NewRoleRepository(cfg *config.Config, dbHelper helper.DBHelper) RoleRepository {
	return &roleRepository{
		cfg:      cfg,
		dbHelper: dbHelper,
	}
}

func (r *roleRepository) CreateRole(ctx context.Context,
	roleName,
	roleDescription,
	roleManagerDomain,
	roleAlias string) (uint32, error) {

	stmt := `INSERT INTO IM_ROLE (ROLE_NAME, ROLE_DESCRIPTION, ROLE_MANAGER_DOMAIN, ROLE_ALIAS)
	VALUES (?, ?, ?, ?)`
	var args []interface{}
	args = append(args, roleName,
		roleDescription,
		roleManagerDomain,
		roleAlias,
	)
	tx, err := r.dbHelper.Begin()
	if err != nil {
		return 0, err
	}
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
	}

	insertedID, _ := result.LastInsertId()

	return uint32(insertedID), err
}
func (r *roleRepository) EditRole(ctx context.Context,
	roleID uint32,
	roleName, roleDescription, roleManagerDomain, roleAlias string) (uint32, error) {

	builder := strings.Builder{}
	builder.WriteString(`UPDATE IM_ROLE SET ROLE_NAME = ?
	, ROLE_DESCRIPTION = ?
	, ROLE_MANAGER_DOMAIN = ? 
	, ROLE_ALIAS = ? `)
	var queryString []string
	var args []interface{}

	args = append(args, roleName,
		roleDescription,
		roleManagerDomain,
		roleAlias,
	)
	builder.WriteString(strings.Join(queryString, " , "))
	builder.WriteString(" WHERE ID = ? ")
	args = append(args, roleID)
	tx, err := r.dbHelper.Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec(builder.String(), args...)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if num, err := result.RowsAffected(); num == 0 || err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return roleID, nil
}
func (r *roleRepository) DeleteRoleTx(ctx context.Context, tx *TransactionExt, roleID uint32) error {

	stmt := `DELETE FROM IM_ROLE WHERE ID = ? `

	result, err := tx.tx.Exec(stmt, roleID)
	if err != nil {
		return err
	}

	if num, err := result.RowsAffected(); num == 0 || err != nil {
		if num == 0 {
			return nil
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *roleRepository) GetAllRoles(ctx context.Context) ([]*model.Role, error) {

	stmt := `SELECT ID, ROLE_NAME, ROLE_DESCRIPTION, ROLE_MANAGER_DOMAIN, ROLE_ALIAS, IS_ACTIVE 
	FROM IM_ROLE 
	ORDER BY ID desc`
	rows, err := r.dbHelper.Open().Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var role *model.Role
	var roles []*model.Role
	for rows.Next() {
		role = &model.Role{}
		err = rows.Scan(&role.ID, &role.RoleName, &role.RoleDescription, &role.RoleManagerDomain, &role.RoleAlias, &role.IsActive)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if len(roles) < 1 {
		return nil, errors.New(common.ReasonNotFound.Code())
	}

	return roles, nil
}

func (r *roleRepository) GetRoleDetail(ctx context.Context, roleID uint32) (*model.Role, bool, error) {

	stmt := `SELECT ID, ROLE_NAME, ROLE_DESCRIPTION, ROLE_MANAGER_DOMAIN, ROLE_ALIAS, IS_ACTIVE  
		FROM IM_ROLE 
		WHERE ID = ?`
	role := &model.Role{}
	err := r.dbHelper.Open().QueryRow(stmt, roleID).Scan(&role.ID,
		&role.RoleName,
		&role.RoleDescription,
		&role.RoleManagerDomain,
		&role.RoleAlias,
		&role.IsActive)
	if err != nil {
		return nil, false, err
	}
	return role, true, err
}
