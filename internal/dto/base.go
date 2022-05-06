package dto

type (
	BaseResponse struct {
		TransactionID string
		StatusCode    string
		ReasonCode    string
		ReasonMessage string
	}

	UserBaseDTO struct {
		Domain   string
		Password string
	}
)
