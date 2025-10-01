package gin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ginLogWriter struct{}

func (ginLogWriter) Write(p []byte) (int, error) {
	if len(p) > 0 {
		l := map[string]interface{}{}
		err := json.Unmarshal(p, &l)
		if err != nil {
			logrus.Debug("unmarshal log message: ", err.Error())
			logrus.Info(string(p))
		} else {
			level, err := logrus.ParseLevel(l["level"].(string))
			if err != nil {
				logrus.Debug("parse log level: ", err.Error())
				level = logrus.InfoLevel
			}
			if level == logrus.InfoLevel {
				logrus.Info(l["message"])
			} else if level == logrus.ErrorLevel {
				logrus.Error(l["message"])
			} else if level == logrus.WarnLevel {
				logrus.Warning(l["message"])
			} else if level == logrus.PanicLevel || level == logrus.FatalLevel {
				logrus.Error(l["message"])
			} else if level == logrus.DebugLevel {
				logrus.Debug(l["message"])
			} else {
				logrus.Trace(l["message"])
			}
		}
	}
	return len(p), nil
}

func Logger(skipPaths ...string) gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Output:    &ginLogWriter{},
		SkipPaths: skipPaths,
		Formatter: func(param gin.LogFormatterParams) string {
			var statusColor, methodColor, resetColor string

			if param.IsOutputColor() {
				statusColor = param.StatusCodeColor()
				methodColor = param.MethodColor()
				resetColor = param.ResetColor()
			}

			level := logrus.InfoLevel
			if param.StatusCode >= http.StatusInternalServerError {
				level = logrus.ErrorLevel
			} else if param.StatusCode >= http.StatusRequestTimeout {
				level = logrus.WarnLevel
			}
			msg := fmt.Sprintf("[GIN] %s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
				statusColor, param.StatusCode, resetColor,
				param.Latency,
				param.ClientIP,
				methodColor, param.Method, resetColor,
				param.Path,
				param.ErrorMessage,
			)
			l := map[string]interface{}{
				"level":   level,
				"message": msg,
			}
			t, err := json.Marshal(l)
			if err != nil {
				logrus.Debug("marshal logrus message: ", err.Error())
				return msg
			}
			return string(t)
		},
	})
}
