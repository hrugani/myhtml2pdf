package app

import (
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	router = gin.Default()
)

func StartApplication(port int) {

	logger, _ := createMyZapLogger()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	listeningString := fmt.Sprintf(":%d", port)
	mapUrls()
	router.Run(listeningString)
}

func createMyZapLogger() (*zap.Logger, error) {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	zapLogger := zap.New(core, zap.AddCaller())

	return zapLogger, nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./mypdfservice_debug.log",
		MaxSize:    10, // Megabytes
		MaxBackups: 20, // max number of old log files
		MaxAge:     90, // days
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
