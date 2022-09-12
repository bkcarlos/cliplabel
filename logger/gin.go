package logger

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var defaultLogFormatter = func(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.Latency > time.Minute {
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006-01-02 15:04:05.495768536+08:00"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

// NewGinLogger 获取gin loggger
//  @return gin.HandlerFunc
func NewGinLogger() gin.HandlerFunc {
	conf := gin.LoggerConfig{}
	formatter := conf.Formatter
	if formatter == nil {
		formatter = defaultLogFormatter
	}

	out := conf.Output
	if out == nil {
		out = GetOutput()
	}

	notlogged := conf.SkipPaths

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		// TODO:: 打印 req body
		{
			data, err := c.GetRawData()
			if err != nil {
				Logger.Errorf("get request raw data failed, err: %v", err)
				c.Abort()
				return
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			if len(data) > 0 {
				Logger.Infof("%v req body :%v", path, string(data))
			}
		}

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			param := gin.LogFormatterParams{
				Request: c.Request,
				Keys:    c.Keys,
			}
			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)

			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}
			param.Path = path
			if strings.HasPrefix(param.Path, "/ping") {
				return
			}
			requestInfoStr := strings.ReplaceAll(strings.TrimSpace(formatter(param)), "\"", "'")
			Logger.Infof(requestInfoStr)
		}
	}
}
