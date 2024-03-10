package logger

import (
	"context"
	"go.uber.org/zap"
)

type key int

const (
	logContextKey key = iota
)

type zapLogger struct {
	zapLogger *zap.Logger
	infoLogger
}

func NewLogger(l *zap.Logger) Logger {
	return &zapLogger{
		zapLogger: l,
		infoLogger: infoLogger{
			log:   l,
			level: zap.InfoLevel,
		},
	}
}

func (z *zapLogger) Info(msg string, fields ...Field) {
	z.zapLogger.Info(msg, fields...)
}

func (z *zapLogger) Infof(format string, v ...interface{}) {
	z.zapLogger.Sugar().Infof(format, v...)
}

func (z *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	z.zapLogger.Sugar().Infow(msg, keysAndValues...)
}

func (z *zapLogger) Debug(msg string, fields ...Field) {
	z.zapLogger.Debug(msg, fields...)
}

func (z *zapLogger) Debugf(format string, v ...interface{}) {
	z.zapLogger.Sugar().Debugf(format, v...)
}

func (z *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	z.zapLogger.Sugar().Debugw(msg, keysAndValues...)
}

func (z *zapLogger) Warn(msg string, fields ...Field) {
	z.zapLogger.Warn(msg, fields...)
}

func (z *zapLogger) Warnf(format string, v ...interface{}) {
	z.zapLogger.Sugar().Warnf(format, v...)
}

func (z *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	z.zapLogger.Sugar().Warnw(msg, keysAndValues...)
}

func (z *zapLogger) Error(msg string, fields ...Field) {
	z.zapLogger.Error(msg, fields...)
}

func (z *zapLogger) Errorf(format string, v ...interface{}) {
	z.zapLogger.Sugar().Errorf(format, v...)
}

func (z *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	z.zapLogger.Sugar().Errorw(msg, keysAndValues...)
}

func (z *zapLogger) Panic(msg string, fields ...Field) {
	z.zapLogger.Panic(msg, fields...)
}

func (z *zapLogger) Panicf(format string, v ...interface{}) {
	z.zapLogger.Sugar().Panicf(format, v...)
}

func (z *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	z.zapLogger.Sugar().Panicw(msg, keysAndValues...)
}

func (z *zapLogger) Fatal(msg string, fields ...Field) {
	z.zapLogger.Fatal(msg, fields...)
}

func (z *zapLogger) Fatalf(format string, v ...interface{}) {
	z.zapLogger.Sugar().Fatalf(format, v...)
}

func (z *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	z.zapLogger.Sugar().Fatalw(msg, keysAndValues...)
}

func (z *zapLogger) V(level Level) InfoLogger {
	if z.zapLogger.Core().Enabled(level) {
		return &infoLogger{
			level: level,
			log:   z.zapLogger,
		}
	}

	return nil
}

func (z *zapLogger) Write(p []byte) (n int, err error) {
	z.zapLogger.Info(string(p))

	return len(p), nil
}

func (z *zapLogger) WithValues(keysAndValues ...interface{}) Logger {
	newLogger := z.zapLogger.With(handleFields(z.zapLogger, keysAndValues)...)

	return NewLogger(newLogger)
}

func (z *zapLogger) WithName(name string) Logger {
	newLogger := z.zapLogger.Named(name)

	return NewLogger(newLogger)
}

func (z *zapLogger) WithContext(ctx context.Context) Logger {
	actx := context.WithValue(ctx, logContextKey, z)
	if actx != nil {
		logger := actx.Value(logContextKey)
		if logger != nil {
			return logger.(Logger)
		}
	}

	return WithName("Unknown-Context")
}

func (z *zapLogger) Flush() {
	_ = z.zapLogger.Sync()
}
