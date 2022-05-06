package model

import (
	"gorm.io/gorm"
)

// type User struct {
// 	UserID      int
// 	FullName    string
// 	Email       string
// 	FirstName   string
// 	LastName    string
// 	PhoneNumber string
// 	Domain      string
// 	IsActive    bool
// 	Address     string
// 	Password    string
// }

type User struct {
	gorm.Model
	UserID      int
	FullName    string
	Email       string
	FirstName   string
	LastName    string
	PhoneNumber string
	Domain      string
	IsActive    bool
	Address     string
	Password    string
}

// Table name for instance
func (r *User) TableName() string {
	return "IM_USER"
}
