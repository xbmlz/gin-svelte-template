package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/xbmlz/gin-svelte-template/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

type Logger struct {
	log  *zap.Logger
	slog *zap.SugaredLogger
}

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
		log:  logger,
		slog: logger.Sugar(),
	}
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
