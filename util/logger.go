package util

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func init(){
	//todo add logger config to customize the logger
	logger, _ = zap.NewDevelopment()
}

func GetLogger() *zap.Logger{
	return logger
}


