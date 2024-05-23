package logger

import (
	"os"
	"path"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/xbmlz/gin-svelte-template/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// A Level is a logging priority. Higher levels are more important.
type Level int8

// ILogger Logger logger interface
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

type Logger struct {
	level Level
	log   *zap.Logger
	slog  *zap.SugaredLogger
}

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
)

func NewLogger(conf config.Config) Logger {
	logPath := conf.Log.Path
	logName := conf.Log.Name
	maxAge := time.Hour * time.Duration(conf.Log.MaxAge)
	rotationTime := time.Hour * time.Duration(conf.Log.RotationTime)
	callerFullPath := conf.Log.CallerFullPath

	cores := make([]zapcore.Core, 0)
	fileCores := createFileZapCore(logPath, logName, maxAge, rotationTime, callerFullPath)

	if len(fileCores) > 0 {
		cores = append(cores, fileCores...)
	}
	cores = append(cores, createConsoleCore(callerFullPath))

	core := zapcore.NewTee(cores...)
	caller := zap.AddCaller()
	callerSkip := zap.AddCallerSkip(2)
	logger := zap.New(core, caller, callerSkip, zap.Development())

	zap.ReplaceGlobals(logger)
	if _, err := zap.RedirectStdLogAt(logger, zapcore.ErrorLevel); err != nil {
		panic(err)
	}

	return Logger{
		level: ParseLevel(conf.Log.Level),
		log:   logger,
		slog:  logger.Sugar(),
	}
}

func (l Logger) GetZapLogger() *zap.Logger {
	return l.log
}

func createConsoleCore(callerFullPath bool) zapcore.Core {
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.EncodeTime = timeEncoder
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	if callerFullPath {
		consoleEncoderConfig.EncodeCaller = customCallerEncoder
	}
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)
	return zapcore.NewCore(consoleEncoder, consoleDebugging, zapcore.DebugLevel)
}

func createFileZapCore(logPath, logName string, maxAge, rotationTime time.Duration, callerFullPath bool) (cores []zapcore.Core) {
	if len(logPath) == 0 {
		return
	}
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		if err = os.MkdirAll(logPath, os.ModePerm); err != nil {
			panic(err)
		}
	}
	logPath = path.Join(logPath, logName)

	errWriter, err := rotatelogs.New(
		logPath+"_err_%Y-%m-%d.log",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		panic(err)
	}

	infoWriter, err := rotatelogs.New(
		logPath+"_info_%Y-%m-%d.log",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		panic(err)
	}

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	errorCore := zapcore.AddSync(errWriter)
	infoCore := zapcore.AddSync(infoWriter)
	fileEncodeConfig := zap.NewProductionEncoderConfig()
	fileEncodeConfig.EncodeTime = timeEncoder
	if callerFullPath {
		fileEncodeConfig.EncodeCaller = customCallerEncoder
	}
	fileEncoder := zapcore.NewConsoleEncoder(fileEncodeConfig)

	cores = make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(fileEncoder, errorCore, highPriority))
	cores = append(cores, zapcore.NewCore(fileEncoder, infoCore, lowPriority))
	return cores
}

// customCallerEncoder set caller fullpath
func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(caller.FullPath())
}

// timeEncoder format time
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	default:
		return ""
	}
}

// ParseLevel parses a level string into a logger Level value.
func ParseLevel(s string) Level {
	switch strings.ToUpper(s) {
	case "DEBUG":
		return DebugLevel
	case "INFO":
		return InfoLevel
	case "WARN":
		return WarnLevel
	case "ERROR":
		return ErrorLevel
	}
	return InfoLevel
}

// Debug log
func (z *Logger) Debug(v ...any) {
	if z.level <= DebugLevel {
		z.slog.Debug(v...)
	}
}

// Debugf log
func (z *Logger) Debugf(format string, v ...any) {
	if z.level <= DebugLevel {
		z.slog.Debugf(format, v...)
	}
}

// Info log
func (z *Logger) Info(v ...any) {
	if z.level <= InfoLevel {
		z.slog.Info(v...)
	}
}

// Infof log
func (z *Logger) Infof(format string, v ...any) {
	if z.level <= InfoLevel {
		z.slog.Infof(format, v...)
	}
}

// Warn log
func (z *Logger) Warn(v ...any) {
	if z.level <= WarnLevel {
		z.slog.Warn(v...)
	}
}

// Warnf log
func (z *Logger) Warnf(format string, v ...any) {
	if z.level <= WarnLevel {
		z.slog.Warnf(format, v...)
	}
}

// Error log
func (z *Logger) Error(v ...any) {
	if z.level <= ErrorLevel {
		z.slog.Error(v...)
	}
}

// Errorf log
func (z *Logger) Errorf(format string, v ...any) {
	if z.level <= ErrorLevel {
		z.slog.Errorf(format, v...)
	}
}
