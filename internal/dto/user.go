package dto

type (
	LoginRequest struct {
		Domain   string
		Password string
	}

	RegisterRequest struct {
		FirstName   string
		LastName    string
		Email       string
		PhoneNumber string
		Domain      string
		Password    string
		Address     string
	}

	UserDetailDTO struct {
		UserID      int
		Domain      string
		FullName    string
		Email       string
		FirstName   string
		LastName    string
		Address     string
		PhoneNumber string
		IsActive    bool
	}

	SearchUserDTO struct {
		Domain string
	}

	GetRoleUsersResponseDTO struct {
		BaseResponse
		UserDetails []*UserDetailDTO
	}

	UserTokenDTO struct {
		UserID uint32
		Domain string
	}
)
