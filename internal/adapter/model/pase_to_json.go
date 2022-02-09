package model

import (
	"time"

	"github.com/o1egl/paseto"
)

type PasetoJSONToken struct {
	*paseto.JSONToken
	UserID       uint32
	RefreshToken string
	Domain       string
	TimeInForce  time.Time
}
