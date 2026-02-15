package logging

import "go.uber.org/zap"

func InitZap() (*zap.Logger, error) {
	// Production config outputs JSON, good for enterprise logs
	return zap.NewProduction()
}
