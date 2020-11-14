package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dantin-s/hydra/logger"
)

// Message is the common message log for gin http server.
type Message struct {
	Origin  string
	RawHost string
	Scope   string
	Status  int
	Elapse  time.Duration
}

// MarshalLogObject
func (s *Message) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("Origin", s.Origin)
	enc.AddString("Raw-Host", s.RawHost)
	enc.AddString("Scope", s.Scope)
	enc.AddInt("Status", s.Status)
	enc.AddDuration("Elapse", s.Elapse)

	return nil
}

// AccessLog log json format access log using zap.
func AccessLog(logLevel string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if logger.IsLevelEnabled(logLevel) {
			start := time.Now()
			c.Next()
			end := time.Now()
			latency := end.Sub(start)

			if len(c.Errors) > 0 {
				for _, e := range c.Errors.Errors() {
					logger.EventLogger.Error(e)
				}
			}

			logger.AppLogger.Info("access log",
				zap.String("time", time.Now().Format(time.RFC3339)),
				zap.String("request_id", c.Request.Header.Get("Request-Id")),
				zap.Object("message", &Message{
					Origin: c.Request.Header.Get("Origin"),
					Status: c.Writer.Status(),
					Elapse: latency,
				}),
			)
		}
	}
}
