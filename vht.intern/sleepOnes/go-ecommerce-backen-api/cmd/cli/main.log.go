package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello world, %s", "Duc")

	// //logger
	// logger := zap.NewExample()
	// logger.Info("Hello world", zap.String("name", "Duc"))

	//2
	/*
		logger, _ := zap.NewProduction()
		logger := zap.NewExample()
		logger.Info("Hello new Example")

		//development
		logger, _ = zap.NewDevelopment()
		logger.Info("Hello new Development")

		//production
		logger, _ = zap.NewProduction()
		logger.Info("Hello new Production")
	*/
	//3 custom log

	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	logger := zap.New(core)

	logger.Info("Hello world",
		zap.Int("line", 1),
		zap.String("name", "Duc"))

	logger.Error("Error log",
		zap.Int("line", 2))
}

func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
