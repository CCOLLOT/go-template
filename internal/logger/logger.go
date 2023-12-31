package logger

import "go.uber.org/zap"

func New() (*zap.Logger, error) {
	log, err := zap.NewDevelopment()
	return log, err
}
