package logconfig

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	Logger *zerolog.Logger
}

// Configuration for logging
type LogConfig struct {
	// Enable console logging
	ConsoleLoggingEnabled bool

	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Filename is the name of the logfile which will be placed inside the directory
	FilenameAddr string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
}

//create a new logger
func New(lconfig LogConfig) *Logger {
	var writers []io.Writer

	if lconfig.ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if lconfig.FileLoggingEnabled {
		writers = append(writers, newRollingFile(lconfig))
	}
	mw := io.MultiWriter(writers...)

	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(mw).With().Timestamp().Logger()
	// logger.Info().
	// 	Bool("fileLogging", lconfig.FileLoggingEnabled).
	// 	Bool("jsonLogOutput", lconfig.EncodeLogsAsJson).
	// 	Str("logDirectory", lconfig.Directory).
	// 	Str("fileName", lconfig.Filename).
	// 	Int("maxSizeMB", lconfig.MaxSize).
	// 	Int("maxBackups", lconfig.MaxBackups).
	// 	Int("maxAgeInDays", lconfig.MaxAge).
	// 	Msg("logging configured")

	return &Logger{
		Logger: &logger,
	}
}

func newRollingFile(lconfig LogConfig) io.Writer {

	return &lumberjack.Logger{
		Filename:   lconfig.FilenameAddr,
		MaxBackups: lconfig.MaxBackups, // files
		MaxSize:    lconfig.MaxSize,    // megabytes
		MaxAge:     lconfig.MaxAge,     // days
	}
}
