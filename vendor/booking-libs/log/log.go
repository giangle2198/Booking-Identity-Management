package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *logger = &logger{}
)

type logger struct {
	*zap.SugaredLogger
	L *zap.Logger
}

func InitZap(app, env string, maskFields map[string]string) error {
	logLevel := zapcore.InfoLevel

	encoderConfig := zapcore.EncoderConfig{
		MessageKey: "message",

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,

		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,

		NameKey:    "app",
		EncodeName: zapcore.FullNameEncoder,
	}

	cfg := zap.Config{
		Encoding:         "custom-json",
		Level:            zap.NewAtomicLevelAt(logLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    encoderConfig,
	}

	zap.RegisterEncoder("custom-json", func(cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return zapcore.NewConsoleEncoder(cfg), nil
	})

	l, err := cfg.Build()
	if err != nil {
		return err
	}

	l = l.Named(app)
	zap.ReplaceGlobals(l)
	Logger = &logger{l.Sugar(), l}
	return nil
}
