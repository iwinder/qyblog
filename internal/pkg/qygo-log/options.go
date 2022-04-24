package qygo_log

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/pflag"
	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"
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
	jsonForm      = "json"
)

// Options 配置参数
type Options struct {
	Name              string   `json:"name" mapstructure:"name"`
	Level             string   `json:"level" mapstructure:"level"`
	Format            string   `json:"format" mapstructure:"format"`           // 格式
	Development       bool     `json:"development" mapstructure:"development"` // 标记是否为开发者模式，会和生产上有所不同
	EnableColor       bool     `json:"enable-color" mapstructure:"enable-color"`
	DisableCaller     bool     `json:"disable-caller" mapstructure:"disable-caller"`         // 是否开启行号和文件名显示功能
	DisableStacktrace bool     `json:"disable-stacktrace" mapstructure:"disable-stacktrace"` // 是否打印调用栈
	OutputPaths       []string `json:"output-paths" mapstructure:"output-paths"`             // 输出路径
	ErrorOutputPaths  []string `json:"error-output-paths" mapstructure:"error-output-paths"` // zap中出现的内部错误输出路径
}

func NewDefaultOption() *Options {
	return &Options{
		Level:             zapcore.InfoLevel.String(),
		Format:            consoleFormat,
		Development:       false,
		EnableColor:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}
}

// Validate 验证
func (o *Options) Validate() []error {
	var errs []error

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(o.Level)); err != nil {
		errs = append(errs, err)
	}

	formt := strings.ToLower(o.Format)
	if formt != consoleFormat && formt != jsonForm {
		errs = append(errs, fmt.Errorf("not a valid log format: %q", o.Format))
	}
	return errs
}

// AddFlags adds flags for log to the specified FlagSet object
func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Name, flagName, o.Name, "The name of the logger .（logger 名称）")
	fs.StringVar(&o.Level, flagLevel, o.Level, "Minimum log output `LEVEL` .（日常最小输出等级）")
	fs.StringVar(&o.Format, flagFormat, o.Format, "Log output `FORMAT`, support plain or json format.（输出格式，支持 plain 和 json）")
	fs.BoolVar(&o.Development, flagDevelopment, o.Development, "Development puts the logger in development mode, which changes "+
		"the behavior of DPanicLevel and takes stacktraces more liberally .(是否为开发者模式)")
	fs.BoolVar(&o.EnableColor, flagEnableColor, o.EnableColor, "Enable output ansi colors in plain format logs.(是否在plain模式开启彩色日志)")
	fs.BoolVar(&o.DisableCaller, flagDisableCaller, o.DisableCaller, "Disable output of caller information in the log.（是否开启行号和文件名显示功能）")
	fs.BoolVar(&o.DisableStacktrace, flagDisableStacktrace, o.DisableStacktrace, "Disable the log to record a stack trace for all messages at or above panic level.（是否打印调用栈）")
	fs.StringSliceVar(&o.OutputPaths, flagOutputPaths, o.OutputPaths, "Output paths of log. （日志输出路径）")
	fs.StringSliceVar(&o.ErrorOutputPaths, flagErrorOutputPaths, o.ErrorOutputPaths, "Error output paths of log. （日志内部错误输出路径）")
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}

// Build 构建一个公共的 zap logger
func (o *Options) Build() error {
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(o.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	encodeLeveL := zapcore.CapitalLevelEncoder
	if o.Format == consoleFormat && o.EnableColor {
		encodeLeveL = zapcore.CapitalColorLevelEncoder
	}

	zc := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       o.Development,
		DisableCaller:     o.DisableCaller,
		DisableStacktrace: o.DisableStacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: o.Format,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "timestamp",
			NameKey:        "logger",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    encodeLeveL,
			EncodeTime:     timeEncoder,
			EncodeDuration: milliSecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		},
		OutputPaths:      o.OutputPaths,
		ErrorOutputPaths: o.ErrorOutputPaths,
	}

	logger, err := zc.Build(zap.AddStacktrace(zapcore.PanicLevel))
	if err != nil {
		return err
	}
	zap.RedirectStdLog(logger.Named(o.Name))
	zap.ReplaceGlobals(logger)
	return nil
}
