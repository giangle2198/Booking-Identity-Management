package model

import "github.com/dgrijalva/jwt-go"

// JWTTokenClaim
type JWTTokenClaim struct {
	*jwt.StandardClaims
	UserID       uint32
	RefreshToken string
	Domain       string
}
