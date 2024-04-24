package sever

import "go.uber.org/zap"

func NewZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	return cfg.Build()
}
