package util

import (
	"os"
	"path/filepath"

	"github.com/HoronLee/GinHub/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger 日志记录器接口
type Logger struct {
	*zap.Logger
}

// NewLogger 创建日志记录器
func NewLogger(cfg *config.AppConfig) *Logger {
	var cores []zapcore.Core
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if cfg.Server.Mode == "debug" {
		// Debug模式：控制台输出，带颜色，非JSON
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		cores = append(cores, zapcore.NewCore(
			consoleEncoder,
			zapcore.AddSync(os.Stdout),
			zapcore.DebugLevel,
		))
	} else if cfg.Server.Mode == "release" {
		// Release模式：控制台+文件，JSON格式，无颜色
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

		// 控制台输出
		cores = append(cores, zapcore.NewCore(
			jsonEncoder,
			zapcore.AddSync(os.Stdout),
			zapcore.InfoLevel,
		))

		// 文件输出
		logDir := "logs"
		os.MkdirAll(logDir, 0755)
		writer := &lumberjack.Logger{
			Filename:   filepath.Join(logDir, "app.log"),
			MaxSize:    100,
			MaxBackups: 5,
			MaxAge:     30,
			Compress:   true,
		}
		cores = append(cores, zapcore.NewCore(
			jsonEncoder,
			zapcore.AddSync(writer),
			zapcore.InfoLevel,
		))
	} else {
		panic("unsupport app mode")
	}

	core := zapcore.NewTee(cores...)
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return &Logger{zapLogger}
}
