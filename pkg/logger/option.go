package logger

import "time"

type Option func(*Logger)

// WithName set file name
func WithName(name string) Option {
	return func(l *Logger) {
		l.conf.name = name
	}
}

// WithPath set file path
func WithPath(path string) Option {
	return func(l *Logger) {
		l.conf.path = path
	}
}

// WithMaxAge set max age
func WithMaxAge(maxAge time.Duration) Option {
	return func(l *Logger) {
		l.conf.maxAge = maxAge
	}
}

// WithRotationTime set rotation time
func WithRotationTime(rotationTime time.Duration) Option {
	return func(l *Logger) {
		l.conf.rotationTime = rotationTime
	}
}

// WithoutStd no output stdout
func WithoutStd() Option {
	return func(l *Logger) {
		l.conf.stdout = false
	}
}

// WithCallerFullPath log caller full path
func WithCallerFullPath() Option {
	return func(l *Logger) {
		l.conf.callerFullPath = true
	}
}
