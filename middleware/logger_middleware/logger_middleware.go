package logger_middleware

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strconv"
	"time"
)

// LoggerMiddleware returns a middleware that logs HTTP requests.
func LoggerMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			resp := c.Response()

			latency := time.Now().Sub(start).String()
			reqSize, _ := strconv.ParseInt(req.Header.Get(echo.HeaderContentLength), 10, 0)

			fields := []zapcore.Field{
				zap.String("path", req.URL.Path),
				zap.String("method", req.Method),
				zap.Int("status", resp.Status),
				zap.Int64("request_size", reqSize),
				zap.Int64("response_size", resp.Size),
				zap.String("user_agent", req.UserAgent()),
				zap.String("remote_ip", c.RealIP()),
				zap.String("latency", latency),
			}

			status := resp.Status
			switch {
			case status >= 500:
				logger.Error("Server error", fields...)
			case status >= 400:
				logger.Info("Client error", fields...)
			case status >= 300:
				logger.Info("Redirection", fields...)
			default:
				logger.Info("Success", fields...)
			}

			return nil
		}
	}
}
