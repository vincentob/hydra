package logger

import (
	"os"

	"go.uber.org/zap"
)

const (
	Info  = "INFO"
	Error = "ERROR"
	Debug = "DEBUG"
)

var (
	AppLogger   *zap.Logger
	EventLogger *zap.Logger
	err         error
)

// init zap AppLogger and EventLogger
// for AppLogger, json format
// for EventLogger, text format
func InitZapLogger(svc string) {
	cfg := zap.NewProductionConfig()

	cfg.EncoderConfig.LevelKey = "log_level"
	cfg.InitialFields = map[string]interface{}{}
	cfg.InitialFields["service"] = svc
	cfg.InitialFields["hostname"], _ = os.Hostname()

	cfg.InitialFields["log_type"] = "app"
	AppLogger, err = cfg.Build()
	if err != nil {
		panic(err)
	}

	cfg.InitialFields["log_type"] = "event"
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	EventLogger, err = cfg.Build()
	if err != nil {
		panic(err)
	}
}
