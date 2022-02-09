package adapter

import (
	"booking-identity-management/config"
	"context"
	"crypto"
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/o1egl/paseto"
	"go.uber.org/zap"

	"booking-identity-management/internal/adapter/model"
	"booking-identity-management/internal/common"
)

type (
	PasetoAdapter interface {
		GenerateToken(ctx context.Context, userID uint32, domain string) (string, error)
		VerifyToken(ctx context.Context, token string) error
	}
	pasetoAdapter struct {
		cfg        *config.Config
		publicKey  crypto.PublicKey
		privateKey crypto.PrivateKey
	}
)

func NewPasetoAdapter(cfg *config.Config) PasetoAdapter {
	publicKeyStr, err := hex.DecodeString(cfg.Paseto.PublicKey)
	if err != nil {
		zap.S().Errorf("Error while converting public key code to string: %v", err)
	}
	publicKey := ed25519.PublicKey(publicKeyStr)
	privateKeyStr, err := hex.DecodeString(cfg.Paseto.PrivateKey)
	if err != nil {
		zap.S().Errorf("Error while converting private key code to string: %v", err)
	}
	privateKey := ed25519.PrivateKey(privateKeyStr)

	return &pasetoAdapter{
		cfg:        cfg,
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func (p *pasetoAdapter) VerifyToken(ctx context.Context, token string) error {
	zap.S().Infof("-------Verify token-------")
	var newJSONToken model.PasetoJSONToken
	var newFooter string

	inTimeSpan := func(start, end, check time.Time) bool {
		return check.Before(end) && check.After(start)
	}

	err := paseto.NewV2().Verify(token, p.publicKey, &newJSONToken, &newFooter)
	if err != nil {
		zap.S().Errorf("Error while verifing paseto token: %v", err)
		return errors.New(common.ReasonInvalidToken.Code())
	}

	if strings.EqualFold(common.Footer, newFooter) {
		zap.S().Errorf("Error while verifing paseto token: %v", err)
		return errors.New(common.ReasonInvalidToken.Code())
	}

	if !strings.EqualFold(newJSONToken.Issuer, p.cfg.Paseto.Issuer) {
		zap.S().Errorf("Error while verifing paseto token: %v", err)
		return errors.New(common.ReasonInvalidToken.Code())
	}

	if !inTimeSpan(newJSONToken.TimeInForce, newJSONToken.Expiration, time.Now()) {
		zap.S().Errorf("Error while verifing paseto token: %v", err)
		return errors.New(common.ReasonExpiredToken.Code())
	}

	return nil
}

func (p *pasetoAdapter) GenerateToken(ctx context.Context, userID uint32, domain string) (token string, err error) {
	zap.S().Infof("Start generated token by paseto")

	claim := model.PasetoJSONToken{
		UserID:      userID,
		Domain:      domain,
		TimeInForce: time.Now(),
		JSONToken: &paseto.JSONToken{
			Issuer:     p.cfg.Paseto.Issuer,
			Expiration: time.Now().Add(time.Second * time.Duration(p.cfg.Paseto.ExpriredTime)),
		},
	}

	// Sign data
	token, err = paseto.NewV2().Sign(p.privateKey, claim, common.Footer)
	if err != nil {
		zap.S().Errorf("new token is failed with err: %v", err)
		return "", err
	}
	return token, err
}
