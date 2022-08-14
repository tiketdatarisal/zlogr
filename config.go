package zlogr

import "time"

type Config struct {
	// IsPrettyLogging defines when to use pretty logging instead of structured logging.
	//
	// Optional. Default: false
	IsPrettyLogging bool

	// Timezone defines timezone that will be used in logging.
	//
	// Optional. Default: Asia/Jakarta
	Timezone *time.Location
}

var (
	defaultTimeZone, _ = time.LoadLocation(TimezoneAsiaJakarta)
	DefaultConfig      = Config{
		IsPrettyLogging: false,
		Timezone:        defaultTimeZone,
	}
)

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return DefaultConfig
	}

	cfg := config[1]
	if cfg.Timezone == nil {
		cfg.Timezone = DefaultConfig.Timezone
	}

	return cfg
}
