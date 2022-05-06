package model

import "gorm.io/gorm"

// type UserAccount struct {
// 	Domain   string
// 	Password string
// }

type UserAccount struct {
	gorm.Model
	Domain   string
	Password string
}

// Table name for instance
func (r *UserAccount) TableName() string {
	return "IM_USER_ACCOUNT"
}
