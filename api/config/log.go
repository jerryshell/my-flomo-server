package config

import (
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogConfig 日志配置
var LogConfig *zap.Config

// InitLog 初始化日志配置
func InitLog() {
	// 设置默认日志级别为Info
	logLevel := zapcore.InfoLevel
	development := false

	// 根据LogLevel配置设置日志级别
	switch Data.LogLevel {
	case "debug":
		logLevel = zapcore.DebugLevel
		development = true
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	case "dpanic":
		logLevel = zapcore.DPanicLevel
	case "panic":
		logLevel = zapcore.PanicLevel
	case "fatal":
		logLevel = zapcore.FatalLevel
	}

	LogConfig = &zap.Config{
		Level:       zap.NewAtomicLevelAt(logLevel),
		Development: development,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// 如果配置了日志文件路径，则添加文件输出
	if Data.LogPath != "" {
		LogConfig.OutputPaths = append(LogConfig.OutputPaths, Data.LogPath)
		LogConfig.ErrorOutputPaths = append(LogConfig.ErrorOutputPaths, Data.LogPath)
	}
}

// GetLogger 获取指定名称的logger
func GetLogger(name string) *zap.Logger {
	if LogConfig == nil {
		InitLog()
	}

	logger, err := LogConfig.Build()
	if err != nil {
		// 如果配置失败，使用默认的sugared logger
		defaultLogger, _ := zap.NewProduction()
		return defaultLogger.Named(name)
	}

	return logger.Named(name)
}

// GetDefaultLogger 获取默认logger
func GetDefaultLogger() *zap.Logger {
	return GetLogger("app")
}

// SyncLogger 同步日志缓冲区
func SyncLogger(logger *zap.Logger) {
	_ = logger.Sync()
}

// LogData 日志数据结构
type LogData struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Service   string    `json:"service"`
	Message   string    `json:"message"`
	Data      any       `json:"data,omitempty"`
	Error     string    `json:"error,omitempty"`
}

// SetupLogFile 设置日志文件
func SetupLogFile() error {
	if Data.LogPath == "" {
		return nil
	}

	// 创建日志目录
	dir := filepath.Dir(Data.LogPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// 创建或打开日志文件
	file, err := os.OpenFile(Data.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
