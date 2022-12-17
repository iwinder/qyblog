package qy_logger

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var _ log.Logger = (*ZapLogger)(nil)

type LoggerConfig struct {
	Mode       string
	Path       string
	Level      string
	MaxSize    int32
	MaxBackups int32
	MaxAge     int32
	Compress   bool
}
type ZapLogger struct {
	log  *zap.Logger
	Sync func() error
}

// Logger 配置zap日志,将zap日志库引入
func Logger(encoder zapcore.Encoder, logConfig LoggerConfig) log.Logger {
	//配置zap日志库的编码器

	//encoder = zapcore.EncoderConfig{
	////	TimeKey:        "time",
	////	LevelKey:       "level",
	////	NameKey:        "logger",
	////	CallerKey:      "caller",
	////	MessageKey:     "msg",
	////	StacktraceKey:  "stack",
	////	EncodeTime:     zapcore.ISO8601TimeEncoder,
	////	LineEnding:     zapcore.DefaultLineEnding,
	////	EncodeLevel:    zapcore.CapitalLevelEncoder,
	////	EncodeDuration: zapcore.SecondsDurationEncoder,
	////	EncodeCaller:   zapcore.FullCallerEncoder,
	////}

	return NewZapLogger(
		encoder,
		logConfig,
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
		zap.AddStacktrace(
			zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(2),
		zap.Development(),
	)
}

// 日志自动切割，采用 lumberjack 实现的
func getLogWriter(logConfig LoggerConfig) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logConfig.Path,            //指定日志存储位置
		MaxSize:    int(logConfig.MaxSize),    //日志的最大大小（M）
		MaxBackups: int(logConfig.MaxBackups), //日志的最大保存数量
		MaxAge:     int(logConfig.MaxAge),     //日志文件存储最大天数
		Compress:   logConfig.Compress,        //是否执行压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

// NewZapLogger return a zap logger.
func NewZapLogger(encoder zapcore.Encoder, logConfig LoggerConfig, level zap.AtomicLevel, opts ...zap.Option) *ZapLogger {
	//日志切割
	writeSyncer := getLogWriter(logConfig)
	//设置日志级别
	defLevel := zap.InfoLevel
	defLevel.Set(logConfig.Level)
	level.SetLevel(defLevel)
	var core zapcore.Core
	//开发模式下打印到标准输出
	// --根据配置文件判断输出到控制台还是日志文件--
	if logConfig.Mode == "dev" {
		core = zapcore.NewCore(
			encoder, // 编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台
			level, // 日志级别
		)
	} else {
		core = zapcore.NewCore(
			encoder, // 编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writeSyncer)), // 打印到控制台和文件
			level, // 日志级别
		)
	}
	zapLogger := zap.New(core, opts...)
	return &ZapLogger{log: zapLogger, Sync: zapLogger.Sync}
}

// Log 实现log接口
func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}

	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	case log.LevelFatal:
		l.log.Fatal("", data...)
	}
	return nil
}
