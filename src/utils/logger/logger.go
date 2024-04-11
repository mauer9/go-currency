package logger

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"os"
)

// NewServiceLogger
// svc logger constructor
func NewServiceLogger(svcName string) log.Logger {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = level.NewFilter(logger, level.AllowDebug())
	logger = log.With(logger,
		"svc", svcName,
		"ts", log.DefaultTimestampUTC,
		"clr", log.DefaultCaller,
	)

	return logger
}
