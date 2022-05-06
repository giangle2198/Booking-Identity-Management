package dto

import (
	"time"
)

type (
	RoleDTO struct {
		ID                uint32
		RoleName          string
		RoleDescription   string
		RoleManagerDomain string
		RoleAlias         string
		IsActive          bool
		CreatedBy         string
		CreatedTime       time.Time
		UpdatedBy         string
		UpdatedTime       time.Time
	}

	UserRoleDTO struct {
		Domain  string
		RoleIDs []uint32
	}

	RoleUserDTO struct {
		RoleID  uint32
		UserIDs []uint32
	}

	GetUserRolesResponseDTO struct {
		BaseResponse
		RoleDetails []*RoleDTO
	}
)
