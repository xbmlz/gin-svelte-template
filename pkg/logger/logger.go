package logger

import (
	"time"

	"go.uber.org/zap"
)

// Logger logger interface
type ILogger interface {
	Debug(v ...any)
	Debugf(format string, v ...any)
	Info(v ...any)
	Infof(format string, v ...any)
	Warn(v ...any)
	Warnf(format string, v ...any)
	Error(v ...any)
	Errorf(format string, v ...any)
}

// LoggerConfig logger config
type LoggerConfig struct {
	// log level
	level Level
	// log file name
	name string
	// log file path, if it is empty no output file
	path string
	// if stdout is true will output stdout
	stdout bool
	// file max save duration, default is 7 days
	maxAge time.Duration
	// file rotation time, default is 1 file per day
	rotationTime time.Duration
	// if callerFullPath is true will output caller fullpath
	callerFullPath bool
}

// Logger logger
type Logger struct {
	conf LoggerConfig
	log  *zap.Logger
	slog *zap.SugaredLogger
}

// NewLogger new logger
func NewLogger(level Level, options ...Option) *Logger {
	logger := &Logger{
		conf: LoggerConfig{
			level:        level,
			name:         "logs",
			stdout:       true,
			maxAge:       7 * 24 * time.Hour,
			rotationTime: 24 * time.Hour,
		},
	}
	for _, option := range options {
		option(logger)
	}

	logger.log = InitZap(logger.conf)
	logger.slog = logger.log.Sugar()
	return logger
}

// Debug log
func (z *Logger) Debug(v ...any) {
	if z.conf.level <= DebugLevel {
		z.slog.Debug(v...)
	}
}

// Debugf log
func (z *Logger) Debugf(format string, v ...any) {
	if z.conf.level <= DebugLevel {
		z.slog.Debugf(format, v...)
	}
}

// Info log
func (z *Logger) Info(v ...any) {
	if z.conf.level <= InfoLevel {
		z.slog.Info(v...)
	}
}

// Infof log
func (z *Logger) Infof(format string, v ...any) {
	if z.conf.level <= InfoLevel {
		z.slog.Infof(format, v...)
	}
}

// Warn log
func (z *Logger) Warn(v ...any) {
	if z.conf.level <= WarnLevel {
		z.slog.Warn(v...)
	}
}

// Warnf log
func (z *Logger) Warnf(format string, v ...any) {
	if z.conf.level <= WarnLevel {
		z.slog.Warnf(format, v...)
	}
}

// Error log
func (z *Logger) Error(v ...any) {
	if z.conf.level <= ErrorLevel {
		z.slog.Error(v...)
	}
}

// Errorf log
func (z *Logger) Errorf(format string, v ...any) {
	if z.conf.level <= ErrorLevel {
		z.slog.Errorf(format, v...)
	}
}
