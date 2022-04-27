package audit

import (
	"windows/config"
	"windows/logging/logconfig"

	"github.com/rs/zerolog"
)

var logger = &logconfig.Logger{}

func Logger() *zerolog.Logger {
	return logger.Logger
}

func CreateAuditLogger() {
	lc := logconfig.LogConfig{
		ConsoleLoggingEnabled: true,
		EncodeLogsAsJson:      true,
		FileLoggingEnabled:    true,
		MaxSize:               10,
		MaxBackups:            10,
		MaxAge:                10,
	}

	if config.Main().AAAS.AuditingStateBool == "false" || config.Main().AAAS.AuditingLogAddress == "" {
		lc.ConsoleLoggingEnabled = false
		lc.FileLoggingEnabled = false
		logger = logconfig.New(lc)
		return
	}

	if config.Main().AAAS.AuditingLogAddress != "" {
		lc.FilenameAddr = config.Main().AAAS.AuditingLogAddress
	}
	logger = logconfig.New(lc)

}
