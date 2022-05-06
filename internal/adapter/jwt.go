package adapter

import (
	"booking-identity-management/config"
	"booking-identity-management/internal/adapter/model"
	"booking-identity-management/internal/common"
	"context"
	cryptoRand "crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

type (
	// JWTAdapter
	JWTAdapter interface {
		GenerateToken(ctx context.Context, userID uint32, domain string) (a, b string, c error)
		VerifyToken(ctx context.Context, token, uid, key string) (model.JWTTokenClaim, error)
	}

	jwtAdapter struct {
		cfg           *config.Config
		signingMethod jwt.SigningMethod
		publicKey     *rsa.PublicKey
		privateKey    *rsa.PrivateKey
	}
)

// NewJWTVerifyAdapter
func NewJWTVerifyAdapter(publicKeyStr, signingMethod string) JWTAdapter {
	publicStr, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		zap.S().Panicf("Error at init public key for JWT: %v", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicStr)
	if err != nil {
		zap.S().Panicf("Error at init public key for JWT: %v", err)
	}
	return &jwtAdapter{
		publicKey:     publicKey,
		signingMethod: jwt.GetSigningMethod(signingMethod),
	}
}

// NewJWTAdapter
func NewJWTAdapter(cfg *config.Config) JWTAdapter {
	privateStr, err := base64.StdEncoding.DecodeString(cfg.JWT.PrivateKey)
	if err != nil {
		zap.S().Panicf("Error at init private key for JWT: %v", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateStr)
	if err != nil {
		zap.S().Panicf("Error at init private key for JWT: %v", err)
	}
	publicStr, err := base64.StdEncoding.DecodeString(cfg.JWT.PublicKey)
	if err != nil {
		zap.S().Panicf("Error at init public key for JWT: %v", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicStr)
	if err != nil {
		zap.S().Panicf("Error at init public key for JWT: %v", err)
	}
	return &jwtAdapter{
		cfg:           cfg,
		signingMethod: jwt.GetSigningMethod(cfg.JWT.SigningMethod),
		privateKey:    privateKey,
		publicKey:     publicKey,
	}
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) (str string) {
	b := make([]rune, n)
	for i := range b {
		randIndex, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return str
		}
		b[i] = letters[randIndex.Int64()]
	}
	return string(b)
}

func nowAsUnixSecond() int64 {
	return time.Now().UnixNano() / 1e9
}

func getRefreshToken() string {
	dest, _ := hex.DecodeString(fmt.Sprintf("%d", nowAsUnixSecond()))
	var id strings.Builder
	encode := base64.StdEncoding.EncodeToString(dest)
	rand.Seed(time.Now().UnixNano())
	id.WriteString(encode)
	id.WriteString(randString(4))
	return strings.Replace(id.String(), "=", randString(1), 1)
}

// GenerateToken generate new token
func (j *jwtAdapter) GenerateToken(ctx context.Context, userID uint32, domain string) (a, b string, c error) {
	token := jwt.New(j.signingMethod)
	var claim model.JWTTokenClaim

	if j.cfg.JWT.IsUsageRefreshToken {
		refreshToken := getRefreshToken()
		claim = model.JWTTokenClaim{
			UserID:       userID,
			RefreshToken: refreshToken,
			Domain:       domain,
			StandardClaims: &jwt.StandardClaims{
				Issuer:    j.cfg.JWT.Issuer,
				ExpiresAt: time.Now().Add(time.Second * time.Duration(j.cfg.JWT.ShortLongTokenExpireTime)).Unix(),
			},
		}
	}

	if !j.cfg.JWT.IsUsageRefreshToken {
		claim = model.JWTTokenClaim{
			UserID: userID,
			Domain: domain,
			StandardClaims: &jwt.StandardClaims{
				Issuer:    j.cfg.JWT.Issuer,
				ExpiresAt: time.Now().Add(time.Second * time.Duration(j.cfg.JWT.LongLiveTokenExpireTime)).Unix(),
			},
		}
	}

	token.Claims = claim

	tokenStr, err := token.SignedString(j.privateKey)
	if err != nil {
		return "", "", err
	}
	return tokenStr, claim.RefreshToken, nil
}

func (j *jwtAdapter) VerifyToken(ctx context.Context, tokenStr, uid, key string) (model.JWTTokenClaim, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return j.publicKey, nil
	}
	var claim model.JWTTokenClaim
	token, err := jwt.ParseWithClaims(tokenStr, &claim, keyFunc)
	v, _ := err.(*jwt.ValidationError)
	if v != nil && v.Errors == jwt.ValidationErrorExpired && claim.RefreshToken != "" {
		err = errors.New(common.ReasonJWTExpired.Code())
		return claim, err
	}
	if err != nil || !token.Valid {
		err = errors.New(common.ReasonJWTInvalid.Code())
		return claim, err
	}
	return claim, nil
}
