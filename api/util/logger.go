package util

import (
	"time"

	"github.com/jerryshell/my-flomo/api/config"
	"go.uber.org/zap"
)

// Logger 日志工具类
type Logger struct {
	*zap.Logger
}

// NewLogger 创建新的logger
func NewLogger(service string) *Logger {
	logger := config.GetLogger(service)
	return &Logger{Logger: logger}
}

// Debug 输出调试日志
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.Logger.Debug(msg, fields...)
}

// Info 输出信息日志
func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.Logger.Info(msg, fields...)
}

// Warn 输出警告日志
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.Logger.Warn(msg, fields...)
}

// Error 输出错误日志
func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.Logger.Error(msg, fields...)
}

// Fatal 输出致命错误日志并退出
func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.Logger.Fatal(msg, fields...)
}

// Debugf 格式化调试日志
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Logger.Sugar().Debugf(format, args...)
}

// Infof 格式化信息日志
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Logger.Sugar().Infof(format, args...)
}

// Warnf 格式化警告日志
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Logger.Sugar().Warnf(format, args...)
}

// Errorf 格式化错误日志
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Logger.Sugar().Errorf(format, args...)
}

// Fatalf 格式化致命错误日志并退出
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Sugar().Fatalf(format, args...)
}

// With 添加字段到logger
func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{Logger: l.Logger.With(fields...)}
}

// Sync 同步日志缓冲区
func (l *Logger) Sync() {
	_ = l.Logger.Sync()
}

// 全局logger实例
var (
	defaultLogger *Logger
)

// init 初始化全局logger
func init() {
	defaultLogger = NewLogger("app")
}

// GetLogger 获取全局logger
func GetLogger() *Logger {
	return defaultLogger
}

// Debug 全局调试日志
func Debug(msg string, fields ...zap.Field) {
	defaultLogger.Debug(msg, fields...)
}

// Info 全局信息日志
func Info(msg string, fields ...zap.Field) {
	defaultLogger.Info(msg, fields...)
}

// Warn 全局警告日志
func Warn(msg string, fields ...zap.Field) {
	defaultLogger.Warn(msg, fields...)
}

// Error 全局错误日志
func Error(msg string, fields ...zap.Field) {
	defaultLogger.Error(msg, fields...)
}

// Fatal 全局致命错误日志
func Fatal(msg string, fields ...zap.Field) {
	defaultLogger.Fatal(msg, fields...)
}

// Debugf 全局格式化调试日志
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Infof 全局格式化信息日志
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Warnf 全局格式化警告日志
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Errorf 全局格式化错误日志
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Fatalf 全局格式化致命错误日志
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// 常用字段快捷方法

// StringField 字符串字段
func StringField(key, value string) zap.Field {
	return zap.String(key, value)
}

// IntField 整数字段
func IntField(key string, value int) zap.Field {
	return zap.Int(key, value)
}

// ErrorField 错误字段
func ErrorField(err error) zap.Field {
	return zap.Error(err)
}

// AnyField 任意类型字段
func AnyField(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

// TimeField 时间字段
func TimeField(key string, value time.Time) zap.Field {
	return zap.Time(key, value)
}