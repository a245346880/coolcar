package server

import "go.uber.org/zap"

// NewZapLogger 创建一个zap日志对象
func NewZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()
}
