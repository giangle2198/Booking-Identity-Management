package model

import (
	"database/sql"

	"gorm.io/gorm"
)

// type Role struct {
// 	ID                uint32
// 	RoleName          string
// 	RoleDescription   sql.NullString
// 	RoleManagerDomain sql.NullString
// 	RoleAlias         sql.NullString
// 	IsActive          sql.NullBool
// 	CreatedBy         string
// 	CreatedTime       sql.NullTime
// 	UpdatedBy         string
// 	UpdatedTime       sql.NullTime
// }

type Role struct {
	gorm.Model
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
func (r *Role) TableName() string {
	return "IM_ROLE"
}
