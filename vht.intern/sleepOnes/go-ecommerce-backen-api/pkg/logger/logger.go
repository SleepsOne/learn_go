package logger

import (
	"fmt"
	"go-ecommerce-backend-api/pkg/setting"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func (l *LoggerZap) Log(level zapcore.Level, msg string, fields ...zap.Field) {
	if l.Core().Enabled(level) {
		l.Check(level, msg).Write(fields...)
	}
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.Log_level
	fmt.Println(logLevel)

	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.File_log_name,
		MaxSize:    config.Max_size, // megabytes
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age,  //days
		Compress:   config.Compress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

// format logs a message
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	// 1716714967.877995 - 2024-05-26T16:16:07.877+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> Time
	encodeConfig.TimeKey = "time"
	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//"caller":"cli/main.log.go:24
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder // zao.Ne
	return zapcore.NewJSONEncoder(encodeConfig)
}
