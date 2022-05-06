package model

import "gorm.io/gorm"

// type JWTManagment struct {
// 	ID          uint64
// 	UserID      uint32
// 	Token       string
// 	CreatedDate string
// }

type JWTManagment struct {
	gorm.Model
	ID          uint64
	UserID      uint32
	Token       string
	CreatedTime string
}

// Table name for instance
func (r *JWTManagment) TableName() string {
	return "IM_JWT_MANAGEMENT"
}
