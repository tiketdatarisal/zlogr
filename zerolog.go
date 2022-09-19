package zlogr

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
	"io"
	"os"
	"time"
)

func init() {
	Initialize()
}

// Initialize will init log engine with new configuration.
func Initialize(config ...Config) logr.Logger {
	cfg := configDefault(config...)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	zerolog.TimestampFunc = func() time.Time { return time.Now().In(cfg.Timezone) }

	zerologr.NameFieldName = "logger"
	zerologr.NameSeparator = "/"

	var writer io.Writer = os.Stdout
	if cfg.IsPrettyLogging {
		writer = zerolog.ConsoleWriter{Out: os.Stdout}
	}

	logger := zerolog.New(writer).Level(cfg.MinimumLevel).With().Caller().Timestamp().Logger()
	Log = zerologr.New(&logger)
	return Log
}
