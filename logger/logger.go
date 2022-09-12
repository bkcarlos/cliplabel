package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jtolds/gls"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// CmdbLoggerContext cmd
type CmdbLoggerContext struct {
	Username string
	UserID   int64
	TraceID  string
}

// GlsHook gls hook
type GlsHook struct{}

// Levels level
//  @receiver h
//  @return []logrus.Level
func (h *GlsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire fire
//  @receiver h
//  @param e
//  @return error
func (h *GlsHook) Fire(e *logrus.Entry) error {
	gls.EnsureGoroutineId(func(gid uint) {
		value, ok := LoggerContextMgr.GetValue(gid)
		if ok {
			context := value.(*CmdbLoggerContext)
			if context.UserID != 0 {
				e.Data["user_identifier"] = fmt.Sprintf("%s(%d)", context.Username, context.UserID)
			}
			if context.TraceID != "" {
				e.Data["trace_id"] = context.TraceID
			}
		}
	})
	return nil
}

// LoggerContextMgr gls 全局上下文管理器
var LoggerContextMgr = gls.NewContextManager()

// Logger 全局logger
var Logger *logrus.Logger

func matchLogLevel(levelStr string) (level logrus.Level) {
	//PANIC FATAL ERROR WARN INFO DEBUG TRACE
	switch levelStr {
	case "trace":
		level = logrus.TraceLevel
	case "debug":
		level = logrus.DebugLevel
	case "info":
		level = logrus.InfoLevel
	case "warn":
		level = logrus.WarnLevel
	case "error":
		level = logrus.ErrorLevel
	case "fatal":
		level = logrus.FatalLevel
	case "panic":
		level = logrus.PanicLevel
	default:
		level = logrus.InfoLevel
	}
	return
}

// GetLogLevel 获取日志等级
//  @return logrus.Level
func GetLogLevel() logrus.Level {
	return Logger.GetLevel()
}

// GetOutput 获取写io
//  @return io.Writer
func GetOutput() io.Writer {
	return Logger.Out
}

// InitLogger 初始化日志
//  @param filePath
//  @param logLevel
//  @param serverName
func InitLogger(filePath, logLevel, serverName string) {
	Logger = logrus.New()

	//设置日志级别
	Logger.SetLevel(matchLogLevel(logLevel))
	// Logger.SetReportCaller(true)

	// 没有配置文件路径，认为输出为标准输出
	if filePath == "" {
		Logger.SetOutput(os.Stdout)
		// 日志格式
		Logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05.495768536+08:00",
		})
	} else {
		writer, _ := rotatelogs.New(
			// 分割后的文件名称
			filePath+".%Y%m%d",

			// 生成软链，指向最新日志文件
			rotatelogs.WithLinkName(filePath),

			// 设置最大保存时间(60天)
			rotatelogs.WithMaxAge(60*24*time.Hour),

			// 设置日志切割时间间隔(1天)
			rotatelogs.WithRotationTime(24*time.Hour),
		)

		Logger.SetOutput(writer)
		// 日志格式
		Logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	Logger.AddHook(&GlsHook{})
}

// Info info
//  @param ctx
//  @param format
//  @param args
func Info(ctx context.Context, format string, args ...interface{}) {
	Logger.WithContext(ctx).Infof(format, args...)
}

// Debug debug
//  @param ctx
//  @param format
//  @param args
func Debug(ctx context.Context, format string, args ...interface{}) {
	Logger.WithContext(ctx).Debugf(format, args...)
}

// Error error
//  @param ctx
//  @param format
//  @param args
func Error(ctx context.Context, format string, args ...interface{}) {
	Logger.WithContext(ctx).Errorf(format, args...)
}

// Warn warn
//  @param ctx
//  @param format
//  @param args
func Warn(ctx context.Context, format string, args ...interface{}) {
	Logger.WithContext(ctx).Warnf(format, args...)
}

// Paninc panic
//  @param ctx
//  @param format
//  @param args
func Paninc(ctx context.Context, format string, args ...interface{}) {
	Logger.WithContext(ctx).Panicf(format, args...)
}

// GetTraceID 获取追踪ID
//  @param ctx
//  @return string
func GetTraceID(ctx context.Context) string {
	traceID := ""
	gls.EnsureGoroutineId(func(gid uint) {
		value, ok := LoggerContextMgr.GetValue(gid)
		if ok {
			context := value.(*CmdbLoggerContext)

			if context.TraceID != "" {
				traceID = context.TraceID
			}
		}
	})

	return traceID
}

// Infos info
//  @param ctx
//  @param format
//  @param args
func Infos(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

// Debugs debug
//  @param ctx
//  @param format
//  @param args
func Debugs(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

// Errors error
//  @param ctx
//  @param format
//  @param args
func Errors(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

// Warns warn
//  @param ctx
//  @param format
//  @param args
func Warns(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

// Panincs panic
//  @param ctx
//  @param format
//  @param args
func Panincs(format string, args ...interface{}) {
	Logger.Panicf(format, args...)
}
