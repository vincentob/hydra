package logger

import (
	"github.com/sirupsen/logrus"
)

const (
	RFC666 = "2006-01-02 15:04:05.000"
)

// InitLogrusConsoleLogger
//
// log level:
//   Panic > Fatal > Error > Warn > Info > Debug > Trace
func InitLogrusConsoleLogger() {
	formatter := &logrus.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		TimestampFormat:           RFC666,
		DisableLevelTruncation:    false,
		QuoteEmptyFields:          true,
		DisableColors:             false,
		DisableTimestamp:          false,
		DisableSorting:            false,
	}

	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.InfoLevel)
}

// InitLogrusJSONLogger
func InitLogrusJSONLogger() {
	formatter := &logrus.JSONFormatter{
		TimestampFormat:   RFC666,
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "message",
		PrettyPrint:       false,
	}

	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.InfoLevel)
}

// IsLevelEnabled check the log level enable or not.
func IsLevelEnabled(lvl string) bool {
	l, err := logrus.ParseLevel(lvl)
	if err != nil {
		return false
	}

	return logrus.IsLevelEnabled(l)
}
