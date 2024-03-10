package logger

import (
	"testing"
)

func Test_WithOptions(t *testing.T) {
	defer Flush() // used for record logger printer
	opt := &Options{
		OutputPaths:       []string{consoleStdout, "/home/windcoder/work/logs/info.log"},
		ErrorOutputPaths:  []string{consoleStderr, "/home/windcoder/work/logs/err.log"},
		Level:             "info",
		Format:            consoleFormat,
		DisableCaller:     false,
		DisableStacktrace: false,
		EnableColor:       false, // 开启 console 颜色时，console 格式的文件会乱码
		Development:       false,
		Name:              "default-log",
		MaxSize:           1,
		MaxBackups:        2,
		MaxAge:            1,
		Compress:          true,
	}
	Init(opt)
	Error("dddddddddddddddddd 大大的")
	Infow("infow Hello world!", "foo", "bar")   // structed logger
	Debugw("debugw Hello world!", "foo", "bar") // structed logger
	Errorw("errorW Hello world!", "foo", "bar") // structed logger
	Error("errorW Hello world!")                // structed logger
	Warn("wwwc")
	//Fatal("xxxx")
}
