package config

// BaseConfig the common env use for all projects.
type BaseConfig struct {
	LogLevel  string `env:"LOG_LEVEL"  envDefault:"INFO" ` // log level defined
	SentryDSN string `env:"SENTRY_DSN" envDefault:"" `     // dsn for sentry
}
