package logger

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// GormLogger gorm loggger
type GormLogger struct {
	slowThreshold                                      time.Duration
	ignoreRecordNotFoundError                          bool
	logLevel                                           logrus.Level
	infoStr, warnStr, errStr                           string
	traceStr, traceErrStr, traceWarnStr                string
	colorInfoStr, colorWarnStr, colorErrStr            string
	colorTraceStr, colorTraceErrStr, colorTraceWarnStr string
	isDebug                                            bool
}

// NewGormLogger 获取grom logger
//  @param logLevel
//  @return *GormLogger
func NewGormLogger(logLevel logrus.Level) *GormLogger {
	var (
		infoStr           = "%s\n[info] "
		warnStr           = "%s\n[warn] "
		errStr            = "%s\n[error] "
		traceStr          = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr      = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr       = "%s %s\n[%.3fms] [rows:%v] %s"
		colorInfoStr      = "%s\n[info] "
		colorWarnStr      = "%s\n[warn] "
		colorErrStr       = "%s\n[error] "
		colorTraceStr     = "%s\n[%.3fms] [rows:%v] %s"
		colorTraceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		colorTraceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	colorInfoStr = logger.Green + "%s\n" + logger.Reset + logger.Green + "[info] " + logger.Reset
	colorWarnStr = logger.BlueBold + "%s\n" + logger.Reset + logger.Magenta + "[warn] " + logger.Reset
	colorErrStr = logger.Magenta + "%s\n" + logger.Reset + logger.Red + "[error] " + logger.Reset
	colorTraceStr = logger.Green + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
	colorTraceWarnStr = logger.Green + "%s " + logger.Yellow + "%s\n" + logger.Reset + logger.RedBold + "[%.3fms] " + logger.Yellow + "[rows:%v]" + logger.Magenta + " %s" + logger.Reset
	colorTraceErrStr = logger.RedBold + "%s " + logger.MagentaBold + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"

	return &GormLogger{
		logLevel:                  logLevel,
		slowThreshold:             200 * time.Millisecond,
		ignoreRecordNotFoundError: true,
		infoStr:                   infoStr,
		warnStr:                   warnStr,
		errStr:                    errStr,
		traceStr:                  traceStr,
		traceWarnStr:              traceWarnStr,
		traceErrStr:               traceErrStr,
		colorInfoStr:              colorInfoStr,
		colorWarnStr:              colorWarnStr,
		colorErrStr:               colorErrStr,
		colorTraceStr:             colorTraceStr,
		colorTraceWarnStr:         colorTraceWarnStr,
		colorTraceErrStr:          colorTraceErrStr,
	}
}

// LogMode log mode
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.logLevel = logrus.Level(level)
	return l
}

// Info print info
func (l GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logrus.InfoLevel {
		l.printf(logrus.InfoLevel, l.infoStr+msg, l.colorInfoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logrus.WarnLevel {
		l.printf(logrus.WarnLevel, l.warnStr+msg, l.colorWarnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (l GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logrus.ErrorLevel {
		l.printf(logrus.ErrorLevel, l.errStr+msg, l.colorErrStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.logLevel >= logrus.ErrorLevel && (!errors.Is(err, logger.ErrRecordNotFound) || !l.ignoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.printf(logrus.ErrorLevel, l.traceErrStr, l.colorTraceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.printf(logrus.ErrorLevel, l.traceErrStr, l.colorTraceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.slowThreshold && l.slowThreshold != 0 && l.logLevel >= logrus.WarnLevel:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.slowThreshold)
		if rows == -1 {
			l.printf(logrus.WarnLevel, l.traceWarnStr, l.colorTraceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.printf(logrus.WarnLevel, l.traceWarnStr, l.colorTraceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.logLevel >= logrus.InfoLevel:
		sql, rows := fc()
		if rows == -1 {
			l.printf(logrus.InfoLevel, l.traceStr, l.colorTraceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.printf(logrus.InfoLevel, l.traceStr, l.colorTraceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}

func (l GormLogger) printf(level logrus.Level, msg string, colorfulMsg string, data ...interface{}) {
	if l.isDebug {
		msg = colorfulMsg
	}
	msg = strings.ReplaceAll(msg, "\n", " ")
	switch level {
	case logrus.TraceLevel:
		Logger.Tracef(msg, data...)
	case logrus.DebugLevel:
		Logger.Debugf(msg, data...)
	case logrus.InfoLevel:
		Logger.Infof(msg, data...)
	case logrus.WarnLevel:
		Logger.Warnf(msg, data...)
	case logrus.ErrorLevel:
		logrus.Errorf(msg, data...)
	}
}
