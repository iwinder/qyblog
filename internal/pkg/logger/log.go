package logger

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
	V(level Level) InfoLogger
	Write(p []byte) (n int, err error)
	WithValues(keysAndValues ...interface{}) Logger
	WithName(name string) Logger
	WithContext(ctx context.Context) Logger
	Flush()
}

var _ Logger = &zapLogger{}

type infoLogger struct {
	level zapcore.Level
	log   *zap.Logger
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
	if checkedEntry := l.log.Check(l.level, msg); checkedEntry != nil {
		checkedEntry.Write(handleFields(l.log, keysAndValues)...)
	}
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
	// 日志级别
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	//  zapcore.LowercaseLevelEncoder 将日志级别字符串转化为小写
	encodeLevel := zapcore.CapitalLevelEncoder
	if opts.Format == consoleFormat && opts.EnableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	encoderConfig := zapcore.EncoderConfig{
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
		EncodeName:     zapcore.FullNameEncoder,
	}

	writes := getLogOutWriter(opts, false)
	errWrites := getLogOutWriter(opts, true)
	atomicLevel := zap.NewAtomicLevelAt(zapLevel)
	var encoder zapcore.Encoder
	if opts.Format == consoleFormat {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	debugLog := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.InfoLevel
	})
	infoLog := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})

	errorLog := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	var coreArry = make([]zapcore.Core, 0, 2)

	if atomicLevel.Level() <= zapcore.InfoLevel { // 保活日志
		if atomicLevel.Level() == zapcore.InfoLevel {
			coreArry = append(coreArry, zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writes...), infoLog))
		} else {
			coreArry = append(coreArry, zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writes...), debugLog))
		}
		coreArry = append(coreArry, zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errWrites...), errorLog))
	} else {
		coreArry = append(coreArry, zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errWrites...), atomicLevel))
	}

	core := zapcore.NewTee(
		// 错误日志
		coreArry...,
	)
	var zapOpts []zap.Option
	if !opts.DisableCaller { // 开启文件及行号
		zapOpts = append(zapOpts, zap.AddCaller())
	}
	if !opts.DisableStacktrace { // 堆栈跟踪
		zapOpts = append(zapOpts, zap.AddStacktrace(zapcore.WarnLevel))
	}
	if opts.Development { // 开发模式
		zapOpts = append(zapOpts, zap.Development())
	}

	l := zap.New(core, zapOpts...)

	logger := &zapLogger{
		zapLogger: l.Named(opts.Name),
		infoLogger: infoLogger{
			log:   l,
			level: zap.InfoLevel,
		},
	}
	//klog.InitLogger(l)
	zap.RedirectStdLog(l)

	return logger
}

func SugaredLogger() *zap.SugaredLogger {
	return std.zapLogger.Sugar()
}
func WithName(s string) Logger { return std.WithName(s) }
func Flush() {
	std.Flush()
}

func Debug(msg string, fields ...Field) {
	std.Debug(msg, fields...)
}
func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}
func Debugw(msg string, keysAndValues ...interface{}) {
	std.Debugw(msg, keysAndValues...)
}
func Info(msg string, fields ...Field) {
	std.Info(msg, fields...)
}

func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}
func Infow(msg string, keysAndValues ...interface{}) {
	std.Infow(msg, keysAndValues...)
}

func Warn(msg string, fields ...Field) {
	std.Warn(msg, fields...)
}
func Warnf(format string, v ...interface{}) {
	std.Warnf(format, v...)
}
func Warnw(msg string, keysAndValues ...interface{}) {
	std.Warnw(msg, keysAndValues...)
}

func Error(msg string, fields ...Field) {
	std.Error(msg, fields...)
}
func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}
func Errorw(msg string, keysAndValues ...interface{}) {
	std.Errorw(msg, keysAndValues...)
}
func Panic(msg string, fields ...Field) {
	std.Panic(msg, fields...)
}

func Panicf(format string, v ...interface{}) {
	std.Panicf(format, v...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	std.Panicw(msg, keysAndValues...)
}
func Fatal(msg string, fields ...Field) {
	std.Fatal(msg, fields...)
}
func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	std.Fatalw(msg, keysAndValues...)
}

func WithContext(ctx context.Context) Logger {
	return std.WithContext(ctx)
}
func (l *zapLogger) L(ctx context.Context) *zapLogger {
	lg := l.clone()

	if requestID := ctx.Value(KeyRequestID); requestID != nil {
		lg.zapLogger = lg.zapLogger.With(zap.Any(KeyRequestID, requestID))
	}
	if username := ctx.Value(KeyUsername); username != nil {
		lg.zapLogger = lg.zapLogger.With(zap.Any(KeyUsername, username))
	}
	if watcherName := ctx.Value(KeyWatcherName); watcherName != nil {
		lg.zapLogger = lg.zapLogger.With(zap.Any(KeyWatcherName, watcherName))
	}

	return lg
}
func (l *zapLogger) clone() *zapLogger {
	copy := *l

	return &copy
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
func handleFields(l *zap.Logger, args []interface{}, additional ...zap.Field) []zap.Field {
	// a slightly modified version of zap.SugaredLogger.sweetenFields
	if len(args) == 0 {
		// fast-return if we have no suggared fields.
		return additional
	}

	fields := make([]zap.Field, 0, len(args)/2+len(additional))
	for i := 0; i < len(args); {
		// check just in case for strongly-typed Zap fields, which is illegal (since
		// it breaks implementation agnosticism), so we can give a better error message.
		if _, ok := args[i].(zap.Field); ok {
			l.DPanic("strongly-typed Zap Field passed to logr", zap.Any("zap field", args[i]))

			break
		}

		if i == len(args)-1 {
			l.DPanic("odd number of arguments passed as key-value pairs for logging", zap.Any("ignored key", args[i]))

			break
		}

		key, val := args[i], args[i+1]
		keyStr, isString := key.(string)
		if !isString {
			// if the key isn't a string, DPanic and stop logging
			l.DPanic(
				"non-string key argument passed to logging, ignoring all later arguments",
				zap.Any("invalid key", key),
			)

			break
		}

		fields = append(fields, zap.Any(keyStr, val))
		i += 2
	}

	return append(fields, additional...)
}
