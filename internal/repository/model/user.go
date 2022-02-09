package model

type User struct {
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
func (r *User) Table() string {
	return "IM_USER"
}
