package qy_log

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"sync"
)

type InfoLogger interface {
	Info(msg string, fields ...Field)
	Infof(format string, v ...interface{})
	Infow(msg string, keysAndValues ...interface{})

	Enabled() bool
}

type Logger interface {
	InfoLogger
	Debug(msg string, fields ...Field)
	Debugf(format string, v ...interface{})
	Debugw(msg string, keysAndValues ...interface{})

	Warn(msg string, fields ...Field)
	Warnf(format string, v ...interface{})
	Warnw(msg string, keysAndValues ...interface{})

	Error(msg string, fields ...Field)
	Errorf(format string, v ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	Panic(msg string, fields ...Field)
	Panicf(format string, v ...interface{})
	Panicw(msg string, keysAndValues ...interface{})

	Fatal(msg string, fields ...Field)
	Fatalf(format string, v ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	V(level int) InfoLogger
	Write(p []byte) (n int, err error)

	WithValues(keysAndValues ...interface{}) Logger

	WithName(name string) Logger

	WithContext(ctx context.Context) context.Context

	Flush()
}

//
type noopInfoLogger struct{}

func (l *noopInfoLogger) Enabled() bool {
	return false
}
func (l *noopInfoLogger) Info(_ string, _ ...Field)        {}
func (l *noopInfoLogger) Infof(_ string, _ ...interface{}) {}
func (l *noopInfoLogger) Infow(_ string, _ ...interface{}) {}

var disabledInfoLogger = &noopInfoLogger{}

type infoLogger struct {
	level zapcore.Level
	log   *zap.Logger
}

func (l *infoLogger) Enabled() bool {
	return true
}
func (l *infoLogger) Info(msg string, fields ...Field) {
	if checkedEntry := l.log.Check(l.level, msg); checkedEntry != nil {
		checkedEntry.Write(fields...)
	}
}
func (l *infoLogger) Infof(format string, args ...interface{}) {
	if checkedEntry := l.log.Check(l.level, fmt.Sprintf(format, args...)); checkedEntry != nil {
		checkedEntry.Write()
	}
}
func (l *infoLogger) Infow(msg string, keysAndValues ...interface{}) {

}

type zapLogger struct {
	zapLogger *zap.Logger
	infoLogger
}

var (
	std = New(NewOptions())
	mu  sync.Mutex
)

func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()
	std = New(opts)
}
func New(opts *Options) *zapLogger {
	if opts == nil {
		opts = NewOptions()
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	encodeLevel := zapcore.CapitalLevelEncoder
	if opts.Format == consoleFormat && opts.EnableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	encoderonfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	loggerConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       opts.Development,
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         opts.Format,
		EncoderConfig:    encoderonfig,
		OutputPaths:      opts.OutputPaths,
		ErrorOutputPaths: opts.ErrorOutputPaths,
	}

	l, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	logger := &zapLogger{
		zapLogger: l.Named(opts.Name),
		infoLogger: infoLogger{
			log:   l,
			level: zap.InfoLevel,
		},
	}

	zap.RedirectStdLog(l)

	return logger

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

func V(level int) InfoLogger {
	return std.V(level)
}
func (l *zapLogger) V(level int) InfoLogger {
	lvl := zapcore.Level(-1 * level)
	if l.zapLogger.Core().Enabled(lvl) {
		return &infoLogger{
			level: lvl,
			log:   l.zapLogger,
		}
	}
	return disabledInfoLogger
}

func ZapLogger() *zap.Logger {
	return std.zapLogger
}

// Debug
func Debug(msg string, fields ...Field) {
	std.zapLogger.Debug(msg, fields...)
}
func (l *zapLogger) Debug(msg string, fields ...Field) {
	l.zapLogger.Debug(msg, fields...)
}

func Debugf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Debugf(format, v...)
}
func (l *zapLogger) Debugf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Debugf(format, v...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Debugw(msg, keysAndValues...)
}
func (l *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Debugw(msg, keysAndValues...)
}

// Info
func Info(msg string, fields ...Field) {
	std.zapLogger.Info(msg, fields...)
}
func (l *zapLogger) Info(msg string, fields ...Field) {
	l.zapLogger.Info(msg, fields...)
}

func Infof(format string, v ...interface{}) {
	std.zapLogger.Sugar().Infof(format, v...)
}
func (l *zapLogger) Infof(format string, v ...interface{}) {
	l.zapLogger.Sugar().Infof(format, v...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Infow(msg, keysAndValues...)
}
func (l *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Infow(msg, keysAndValues...)
}

// Warn

func Warn(msg string, fields ...Field) {
	std.zapLogger.Warn(msg, fields...)
}
func (l *zapLogger) Warn(msg string, fields ...Field) {
	l.zapLogger.Warn(msg, fields...)
}
func Warnf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Warnf(format, v...)
}
func (l *zapLogger) Warnf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Warnf(format, v...)
}
func Warnw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Warnw(msg, keysAndValues...)
}
func (l *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Warnw(msg, keysAndValues...)
}

// Error Level
func Error(msg string, fields ...Field) {
	std.zapLogger.Error(msg, fields...)
}
func (l *zapLogger) Error(msg string, fields ...Field) {
	l.zapLogger.Error(msg, fields...)
}
func Errorf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Errorf(format, v...)
}
func (l *zapLogger) Errorf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Errorf(format, v...)
}
func Errorw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Errorw(msg, keysAndValues...)
}
func (l *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Errorw(msg, keysAndValues...)
}

// Panic Level
func Panic(msg string, fields ...Field) {
	std.zapLogger.Panic(msg, fields...)
}
func (l *zapLogger) Panic(msg string, fields ...Field) {
	l.zapLogger.Panic(msg, fields...)
}
func Panicf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Panicf(format, v...)
}
func (l *zapLogger) Panicf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Panicf(format, v...)
}
func Panicw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Panicw(msg, keysAndValues...)
}
func (l *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Panicw(msg, keysAndValues...)
}

// Fatal Level
func Fatal(msg string, fields ...Field) {
	std.zapLogger.Fatal(msg, fields...)
}
func (l *zapLogger) Fatal(msg string, fields ...Field) {
	l.zapLogger.Fatal(msg, fields...)
}
func Fatalf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Panicf(format, v...)
}
func (l *zapLogger) Fatalf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Panicf(format, v...)
}
func Fatalw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Fatalw(msg, keysAndValues...)
}
func (l *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Fatalw(msg, keysAndValues...)
}

func (l *zapLogger) Write(p []byte) (n int, err error) {
	l.zapLogger.Info(string(p))
	return len(p), nil
}

func WithValues(keysAndValues ...interface{}) Logger {
	return std.WithValues(keysAndValues)
}
func (l *zapLogger) WithValues(keysAndValues ...interface{}) Logger {
	return nil
}

func WithName(name string) Logger {
	return std.WithName(name)
}
func (l *zapLogger) WithName(name string) Logger {
	newLogger := l.zapLogger.Named(name)
	return NewLogger(newLogger)
}

func WithContext(ctx context.Context) context.Context {
	return std.WithContext(ctx)
}
func (l *zapLogger) WithContext(ctx context.Context) context.Context {

	return context.WithValue(ctx, logContextKey, l)
}

func L(ctx context.Context) *zapLogger {
	return std.L(ctx)
}
func (l *zapLogger) L(ctx context.Context) *zapLogger {
	lg := l.clone()
	requestId, _ := ctx.Value(KeyRequestID).(string)
	username, _ := ctx.Value(KeyUsername).(string)
	lg.zapLogger = lg.zapLogger.With(zap.String(KeyRequestID, requestId), zap.String(KeyUsername, username))
	return lg
}

func Flush() {
	std.Flush()
}
func (l *zapLogger) Flush() {
	_ = l.zapLogger.Sync()
}
func (l *zapLogger) clone() *zapLogger {
	copy := *l
	return &copy
}
func NewLogger(l *zap.Logger) Logger {
	return &zapLogger{
		zapLogger: l,
		infoLogger: infoLogger{
			level: zap.InfoLevel,
			log:   l,
		},
	}
}
