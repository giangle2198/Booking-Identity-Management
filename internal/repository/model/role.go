package model

import "database/sql"

type Role struct {
	ID                uint32
	RoleName          string
	RoleDescription   sql.NullString
	RoleManagerDomain sql.NullString
	RoleAlias         sql.NullString
	IsActive          sql.NullBool
	CreatedBy         string
	CreatedTime       sql.NullTime
	UpdatedBy         string
	UpdatedTime       sql.NullTime
}

// Table name for instance
func (r *Role) Table() string {
	return "IM_ROLE"
}
