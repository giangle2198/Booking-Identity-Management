package common

type ReasonCode string

const (
	ReasonGeneralError   ReasonCode = "-5"
	ReasonExpiredToken   ReasonCode = "-1000"
	ReasonInvalidToken   ReasonCode = "-10001"
	ReasonNotFound       ReasonCode = "-1"
	ReasonDBError        ReasonCode = "-2"
	ReasonNotValid       ReasonCode = "-3"
	ReasonNotCorrectPass ReasonCode = "-4"
	ReasonNotFoundUserMD ReasonCode = "-6"
	ReasonJWTExpired     ReasonCode = "-7"
	ReasonJWTInvalid     ReasonCode = "-8"

	// Reason User
	ReasonNotFoundDomain     ReasonCode = "-100"
	ReasonUserNotActive      ReasonCode = "-101"
	ReasonExistedUser        ReasonCode = "-102"
	ReasonFailedRegisterUser ReasonCode = "-103"

	// Reason JWT
	ReasonJWTError ReasonCode = "-200"
)

var reasonCodeValues = map[string]string{
	"-1":    "Not found any data",
	"-5":    "General Error",
	"-1000": "Token is expired",
	"-1001": "Invalid Token",
	"-2":    "DB Error",
	"-3":    "Data not valid",
	"-4":    "Password not correct",
	"-6":    "Not found user MD",
	"-7":    "JWT is expired",
	"-8":    "JWT invalid",
	"-100":  "Domain is not existed",
	"-101":  "Domain is unactive",
	"-102":  "Domain is existed",
	"-103":  "Failed while registering user",
	"-200":  "JWT Error",
}

// CheckReasonExisted check reason existed in predefined reason
func CheckReasonExisted(reason ReasonCode) bool {
	if _, ok := reasonCodeValues[string(reason)]; !ok {
		return false
	}
	return true
}

func (rc ReasonCode) Message() string {
	if value, ok := reasonCodeValues[rc.Code()]; ok {
		return value
	}
	return reasonCodeValues["-5"]
}

func (rc ReasonCode) Code() string {
	return string(rc)
}

func ParseError(err error) ReasonCode {
	return ReasonCode(err.Error())
}
