package qygo_log

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"sync"
)

//

var disabledInfoLogger = &noopInfoLogger{}

var (
	std = New(NewDefaultOption())
	mu  sync.Mutex
)

func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()
	std = New(opts)
}

func SugaredLogger() *zap.SugaredLogger {
	return std.zapLogger.Sugar()
}

func StdErrLogger() *log.Logger {
	if std == nil {
		return nil
	}
	if l, err := zap.NewStdLogAt(std.zapLogger, zapcore.ErrorLevel); err == nil {
		return l
	}
	return nil
}

func StdInfoLogger() *log.Logger {
	if std == nil {
		return nil
	}
	if l, err := zap.NewStdLogAt(std.zapLogger, zapcore.InfoLevel); err == nil {
		return l
	}
	return nil
}

func ZapLogger() *zap.Logger {
	return std.zapLogger
}

// Debug
func Debug(msg string, fields ...Field) {
	std.zapLogger.Debug(msg, fields...)
}

func Debugf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Debugf(format, v...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Debugw(msg, keysAndValues...)
}

// Info
func Info(msg string, fields ...Field) {
	std.zapLogger.Info(msg, fields...)
}

func Infof(format string, v ...interface{}) {
	std.zapLogger.Sugar().Infof(format, v...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Infow(msg, keysAndValues...)
}

// Warn
func Warn(msg string, fields ...Field) {
	std.zapLogger.Warn(msg, fields...)
}
func Warnf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Warnf(format, v...)
}
func Warnw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Warnw(msg, keysAndValues...)
}

// Error Level
func Error(msg string, fields ...Field) {
	std.zapLogger.Error(msg, fields...)
}
func Errorf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Errorf(format, v...)
}
func Errorw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Errorw(msg, keysAndValues...)
}

// Panic Level
func Panic(msg string, fields ...Field) {
	std.zapLogger.Panic(msg, fields...)
}

func Panicf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Panicf(format, v...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Panicw(msg, keysAndValues...)
}

// Fatal Level
func Fatal(msg string, fields ...Field) {
	std.zapLogger.Fatal(msg, fields...)
}

func Fatalf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Panicf(format, v...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Fatalw(msg, keysAndValues...)
}

func V(level int) InfoLogger {
	return std.V(level)
}

func WithValues(keysAndValues ...interface{}) Logger {
	return std.WithValues(keysAndValues)
}

func WithName(name string) Logger {
	return std.WithName(name)
}

func WithContext(ctx context.Context) context.Context {
	return std.WithContext(ctx)
}

func L(ctx context.Context) *zapLogger {
	return std.L(ctx)
}

func Flush() {
	std.Flush()
}
