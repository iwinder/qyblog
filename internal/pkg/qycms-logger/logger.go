package qycms_logger

import (
	"context"
	"fmt"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qygo-log"
	gormlogger "gorm.io/gorm/logger"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Writer log writer interface.
type Writer interface {
	Printf(string, ...interface{})
}

// Config 定义 gorm logger configuration
type Config struct {
	SlowThreshold time.Duration
	Colorful      bool
	LogLevel      gormlogger.LogLevel
}
type logger struct {
	Writer
	Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func (l *logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= GormInfo {
		l.Printf(l.infoStr+msg, append([]interface{}{fileWithLineNum()}, data...)...)
	}
}

func (l *logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= GormWarn {
		l.Printf(l.warnStr+msg, append([]interface{}{fileWithLineNum()}, data...)...)
	}
}

func (l *logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= GormError {
		l.Printf(l.errStr+msg, append([]interface{}{fileWithLineNum()}, data...)...)
	}
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= 0 {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= GormError:
		sql, rows := fc()
		if rows == -1 {
			l.Printf(l.traceErrStr, fileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceErrStr, fileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= GormWarn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Printf(l.traceWarnStr, fileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceWarnStr, fileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel >= GormInfo:
		sql, rows := fc()
		if rows == -1 {
			l.Printf(l.traceStr, fileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceStr, fileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}

}

func (l *logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func fileWithLineNum() string {
	for i := 4; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)

		if ok && !strings.HasSuffix(file, "_test.go") {
			dir, f := filepath.Split(file)

			return filepath.Join(filepath.Base(dir), f) + ":" + strconv.FormatInt(int64(line), 10)
		}
	}
	return ""
}

// New 创建一个新的 gorm logger instance
func New(level int) gormlogger.Interface {
	var (
		infoStr      = "%s[info] "
		warnStr      = "%s[warn] "
		errStr       = "%s[error] "
		traceStr     = "[%s][%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s[%.3fms] [rows:%v] %s"
	)

	config := Config{
		SlowThreshold: 200 * time.Millisecond,
		Colorful:      false,
		LogLevel:      gormlogger.LogLevel(level),
	}

	if config.Colorful {
		infoStr = Green + "%s " + Reset + Green + "[info] " + Reset
		warnStr = BlueBold + "%s " + Reset + Magenta + "[warn] " + Reset
		errStr = Magenta + "%s " + Reset + Red + "[error] " + Reset
		traceStr = Green + "%s " + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
		traceWarnStr = Green + "%s " + Yellow + "%s " + Reset + RedBold + "[%.3fms] " + Yellow + "[rows:%v]" + Magenta + " %s" + Reset
		traceErrStr = RedBold + "%s " + MagentaBold + "%s " + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
	}

	return &logger{
		Writer:       log.StdInfoLogger(),
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceErrStr:  traceErrStr,
		traceWarnStr: traceWarnStr,
	}
}
