package model

type JWTManagment struct {
	ID          uint64
	UserID      uint32
	Token       string
	CreatedDate string
}

// Table name for instance
func (r *JWTManagment) Table() string {
	return "IM_JWT_MANAGEMENT"
}
