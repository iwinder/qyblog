package logger

import (
	"encoding/json"
	"fmt"
	"github.com/iwinder/qyblog/internal/pkg/arrays"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/pflag"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

const (
	flagLevel             = "log.level"
	flagDisableCaller     = "log.disable-caller"
	flagDisableStacktrace = "log.disable-stacktrace"
	flagFormat            = "log.format"
	flagEnableColor       = "log.enable-color"
	flagOutputPaths       = "log.output-paths"
	flagErrorOutputPaths  = "log.error-output-paths"
	flagDevelopment       = "log.development"
	flagName              = "log.name"

	consoleFormat = "console"
	jsonFormat    = "json"
	consoleStdout = "stdout"
	consoleStderr = "stderr"
)

type Options struct {
	OutputPaths       []string `json:"output-paths"       mapstructure:"output-paths"`
	ErrorOutputPaths  []string `json:"error-output-paths" mapstructure:"error-output-paths"`
	Level             string   `json:"level"              mapstructure:"level"`
	Format            string   `json:"format"             mapstructure:"format"`
	DisableCaller     bool     `json:"disable-caller"     mapstructure:"disable-caller"`
	DisableStacktrace bool     `json:"disable-stacktrace" mapstructure:"disable-stacktrace"`
	EnableColor       bool     `json:"enable-color"       mapstructure:"enable-color"`
	Development       bool     `json:"development"        mapstructure:"development"`
	Name              string   `json:"name"               mapstructure:"name"`
	MaxSize           int32    `json:"max-size"           mapstructure:"max-size"`
	MaxBackups        int32    `json:"max-backups"        mapstructure:"max-backups"`
	MaxAge            int32    `json:"max-age"            mapstructure:"max-age"`
	Compress          bool     `json:"compress"           mapstructure:"compress"`
}

func NewOptions() *Options {
	return &Options{
		OutputPaths:       []string{consoleStdout},
		ErrorOutputPaths:  []string{consoleStderr},
		Level:             zapcore.InfoLevel.String(),
		Format:            consoleFormat,
		DisableCaller:     false,
		DisableStacktrace: false,
		EnableColor:       false,
		Development:       false,
		Name:              "default-log",
		MaxSize:           1,
		MaxBackups:        2,
		MaxAge:            1,
		Compress:          true,
	}
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Level, flagLevel, o.Level, "日志输出等级 `LEVEL`")
	fs.BoolVar(&o.DisableCaller, flagDisableCaller, o.DisableCaller, "Disable output of caller information in the log.")
	fs.BoolVar(&o.DisableStacktrace, flagDisableStacktrace,
		o.DisableStacktrace, "Disable the log to record a stack trace for all messages at or above panic level.")
	fs.StringVar(&o.Format, flagFormat, o.Format, "Log output `FORMAT`, support plain or json format.")
	fs.BoolVar(&o.EnableColor, flagEnableColor, o.EnableColor, "Enable output ansi colors in plain format logs.")
	fs.StringSliceVar(&o.OutputPaths, flagOutputPaths, o.OutputPaths, "Output paths of log.")
	fs.StringSliceVar(&o.ErrorOutputPaths, flagErrorOutputPaths, o.ErrorOutputPaths, "Error output paths of log.")
	fs.BoolVar(
		&o.Development,
		flagDevelopment,
		o.Development,
		"Development puts the logger in development mode, which changes "+
			"the behavior of DPanicLevel and takes stacktraces more liberally.",
	)
	fs.StringVar(&o.Name, flagName, o.Name, "The name of the logger.")
}

func (o *Options) Validate() []error {
	var errs []error

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(o.Level)); err != nil {
		errs = append(errs, err)
	}

	format := strings.ToLower(o.Format)
	if format != consoleFormat && format != jsonFormat {
		errs = append(errs, fmt.Errorf("not a valid log format: %q", o.Format))
	}

	return errs
}

//func (o *Options) Build() error {
//	var zapLevel zapcore.Level
//	if err := zapLevel.UnmarshalText([]byte(o.Level)); err != nil {
//		zapLevel = zapcore.InfoLevel
//	}
//	encodeLevel := zapcore.CapitalLevelEncoder
//	if o.Format == consoleFormat && o.EnableColor {
//		encodeLevel = zapcore.CapitalColorLevelEncoder
//	}
//	zc := &zap.Config{
//		Level:             zap.NewAtomicLevelAt(zapLevel),
//		Development:       o.Development,
//		DisableCaller:     o.DisableCaller,
//		DisableStacktrace: o.DisableStacktrace,
//		Sampling: &zap.SamplingConfig{
//			Initial:    100,
//			Thereafter: 100,
//		},
//		Encoding: o.Format,
//		EncoderConfig: zapcore.EncoderConfig{
//			MessageKey:     "message",
//			LevelKey:       "level",
//			TimeKey:        "timestamp",
//			NameKey:        "logger",
//			CallerKey:      "caller",
//			StacktraceKey:  "stacktrace",
//			LineEnding:     zapcore.DefaultLineEnding,
//			EncodeLevel:    encodeLevel,
//			EncodeTime:     timeEncoder,
//			EncodeDuration: milliSecondsDurationEncoder,
//			EncodeCaller:   zapcore.ShortCallerEncoder,
//			EncodeName:     zapcore.FullNameEncoder,
//		},
//		OutputPaths:      o.OutputPaths,
//		ErrorOutputPaths: o.ErrorOutputPaths,
//	}
//	logger, err := zc.Build(zap.AddStacktrace(zapcore.PanicLevel))
//	if err != nil {
//		return err
//	}
//	zap.RedirectStdLog(logger.Named(o.Name))
//	zap.ReplaceGlobals(logger)
//	return nil
//}

// 日志自动切割，采用 lumberjack 实现的
func getLogOutWriter(opts *Options, errFlag bool) []zapcore.WriteSyncer {
	var writes = make([]zapcore.WriteSyncer, 0, 2)
	if opts.Development || (!errFlag && arrays.ContainsString(opts.OutputPaths, consoleStdout) >= 0) {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}
	if opts.Development || (errFlag && arrays.ContainsString(opts.ErrorOutputPaths, consoleStderr) >= 0) {
		writes = append(writes, zapcore.AddSync(os.Stderr))
	}
	var outPath = ""
	if errFlag {
		for _, path := range opts.ErrorOutputPaths {
			if path != consoleStderr {
				outPath = path
				break
			}
		}
	} else {
		for _, path := range opts.OutputPaths {
			if path != consoleStdout {
				outPath = path
				break
			}
		}
	}

	if len(outPath) == 0 {
		if len(writes) == 0 {
			if errFlag {
				writes = append(writes, zapcore.AddSync(os.Stderr))
			} else {
				writes = append(writes, zapcore.AddSync(os.Stdout))
			}
		}
		return writes
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   outPath,              //指定日志存储位置
		MaxSize:    int(opts.MaxSize),    //日志的最大大小（M）
		MaxBackups: int(opts.MaxBackups), //日志的最大保存数量
		MaxAge:     int(opts.MaxAge),     //日志文件存储最大天数
		Compress:   opts.Compress,        //是否执行压缩
	}
	writes = append(writes, zapcore.AddSync(lumberJackLogger))
	return writes
}
