package logger

import (
	"context"
	"fmt"
	"log"
	"os"
)

type LoggerLevel int

const (
	LoggerLevelTrace LoggerLevel = iota
	LoggerLevelDebug
	LoggerLevelInfo
	LoggerLevelNotice
	LoggerLevelWarn
	LoggerLevelError
	LoggerLevelFatal
)

type ILogger interface {
	Trace(v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Notice(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})

	Tracef(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Noticef(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})

	CtxTracef(ctx context.Context, format string, v ...interface{})
	CtxDebugf(ctx context.Context, format string, v ...interface{})
	CtxInfof(ctx context.Context, format string, v ...interface{})
	CtxNoticef(ctx context.Context, format string, v ...interface{})
	CtxWarnf(ctx context.Context, format string, v ...interface{})
	CtxErrorf(ctx context.Context, format string, v ...interface{})
	CtxFatalf(ctx context.Context, format string, v ...interface{})
}

var (
	_ ILogger = (*localLogger)(nil)
)

var level LoggerLevel

// SetLevel sets the level of logs below which logs will not be output.
// The default log level is LevelTrace.
func SetLevel(lv LoggerLevel) {
	if lv < LoggerLevelTrace || lv > LoggerLevelFatal {
		panic("invalid level")
	}
	level = lv
}

// Fatal calls the default logger's Fatal method and then os.Exit(1).
func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v)
}

// Error calls the default logger's Error method.
func Error(v ...interface{}) {
	if level > LoggerLevelError {
		return
	}
	defaultLogger.Error(v)
}

// Warn calls the default logger's Warn method.
func Warn(v ...interface{}) {
	if level > LoggerLevelWarn {
		return
	}
	defaultLogger.Warn(v)
}

// Notice calls the default logger's Notice method.
func Notice(v ...interface{}) {
	if level > LoggerLevelNotice {
		return
	}
	defaultLogger.Notice(v)
}

// Info calls the default logger's Info method.
func Info(v ...interface{}) {
	if level > LoggerLevelInfo {
		return
	}
	defaultLogger.Info(v)
}

// Debug calls the default logger's Debug method.
func Debug(v ...interface{}) {
	if level > LoggerLevelDebug {
		return
	}
	defaultLogger.Debug(v)
}

// Trace calls the default logger's Trace method.
func Trace(v ...interface{}) {
	if level > LoggerLevelTrace {
		return
	}
	defaultLogger.Trace(v)
}

// Fatalf calls the default logger's Fatalf method and then os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v)
}

// Errorf calls the default logger's Errorf method.
func Errorf(format string, v ...interface{}) {
	if level > LoggerLevelError {
		return
	}
	defaultLogger.Errorf(format, v)
}

// Warnf calls the default logger's Warnf method.
func Warnf(format string, v ...interface{}) {
	if level > LoggerLevelWarn {
		return
	}
	defaultLogger.Warnf(format, v)
}

// Noticef calls the default logger's Noticef method.
func Noticef(format string, v ...interface{}) {
	if level > LoggerLevelNotice {
		return
	}
	defaultLogger.Noticef(format, v)
}

// Infof calls the default logger's Infof method.
func Infof(format string, v ...interface{}) {
	if level > LoggerLevelInfo {
		return
	}
	defaultLogger.Infof(format, v)
}

// Debugf calls the default logger's Debugf method.
func Debugf(format string, v ...interface{}) {
	if level > LoggerLevelDebug {
		return
	}
	defaultLogger.Debugf(format, v)
}

// Tracef calls the default logger's Tracef method.
func Tracef(format string, v ...interface{}) {
	if level > LoggerLevelTrace {
		return
	}
	defaultLogger.Tracef(format, v)
}

// CtxFatalf calls the default logger's CtxFatalf method and then os.Exit(1).
func CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxFatalf(ctx, format, v)
}

// CtxErrorf calls the default logger's CtxErrorf method.
func CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	if level > LoggerLevelError {
		return
	}
	defaultLogger.CtxErrorf(ctx, format, v)
}

// CtxWarnf calls the default logger's CtxWarnf method.
func CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	if level > LoggerLevelWarn {
		return
	}
	defaultLogger.CtxWarnf(ctx, format, v)
}

// CtxNoticef calls the default logger's CtxNoticef method.
func CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	if level > LoggerLevelNotice {
		return
	}
	defaultLogger.CtxNoticef(ctx, format, v)
}

// CtxInfof calls the default logger's CtxInfof method.
func CtxInfof(ctx context.Context, format string, v ...interface{}) {
	if level > LoggerLevelInfo {
		return
	}
	defaultLogger.CtxInfof(ctx, format, v)
}

// CtxDebugf calls the default logger's CtxDebugf method.
func CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	if level > LoggerLevelDebug {
		return
	}
	defaultLogger.CtxDebugf(ctx, format, v)
}

// CtxTracef calls the default logger's CtxTracef method.
func CtxTracef(ctx context.Context, format string, v ...interface{}) {
	if level > LoggerLevelTrace {
		return
	}
	defaultLogger.CtxTracef(ctx, format, v)
}

var strs = []string{
	"[Trace] ",
	"[Debug] ",
	"[Info] ",
	"[Notice] ",
	"[Warn] ",
	"[Error] ",
	"[Fatal] ",
}

func (lv LoggerLevel) toString() string {
	if lv >= LoggerLevelTrace && lv <= LoggerLevelFatal {
		return strs[lv]
	}
	return fmt.Sprintf("[?%d] ", lv)
}

func SetDefaultLogger(l ILogger) {
	if l == nil {
		panic("logger must not be nil")
	}

}

var defaultLogger ILogger = &localLogger{
	logger: log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
}

type localLogger struct {
	logger *log.Logger
}

func (ll *localLogger) logf(lvl LoggerLevel, format *string, v ...interface{}) {
	if level > lvl {
		return
	}
	msg := lvl.toString()
	if format != nil {
		msg += fmt.Sprintf(*format, v...)
	} else {
		msg += fmt.Sprint(v...)
	}
	ll.logger.Output(3, msg)
	if lvl == LoggerLevelFatal {
		os.Exit(1)
	}
}

func (ll *localLogger) Fatal(v ...interface{}) {
	ll.logf(LoggerLevelFatal, nil, v...)
}

func (ll *localLogger) Error(v ...interface{}) {
	ll.logf(LoggerLevelError, nil, v...)
}

func (ll *localLogger) Warn(v ...interface{}) {
	ll.logf(LoggerLevelWarn, nil, v...)
}

func (ll *localLogger) Notice(v ...interface{}) {
	ll.logf(LoggerLevelNotice, nil, v...)
}

func (ll *localLogger) Info(v ...interface{}) {
	ll.logf(LoggerLevelInfo, nil, v...)
}

func (ll *localLogger) Debug(v ...interface{}) {
	ll.logf(LoggerLevelDebug, nil, v...)
}

func (ll *localLogger) Trace(v ...interface{}) {
	ll.logf(LoggerLevelTrace, nil, v...)
}

func (ll *localLogger) Fatalf(format string, v ...interface{}) {
	ll.logf(LoggerLevelFatal, &format, v...)
}

func (ll *localLogger) Errorf(format string, v ...interface{}) {
	ll.logf(LoggerLevelError, &format, v...)
}

func (ll *localLogger) Warnf(format string, v ...interface{}) {
	ll.logf(LoggerLevelWarn, &format, v...)
}

func (ll *localLogger) Noticef(format string, v ...interface{}) {
	ll.logf(LoggerLevelNotice, &format, v...)
}

func (ll *localLogger) Infof(format string, v ...interface{}) {
	ll.logf(LoggerLevelInfo, &format, v...)
}

func (ll *localLogger) Debugf(format string, v ...interface{}) {
	ll.logf(LoggerLevelDebug, &format, v...)
}

func (ll *localLogger) Tracef(format string, v ...interface{}) {
	ll.logf(LoggerLevelTrace, &format, v...)
}

func (ll *localLogger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LoggerLevelFatal, &format, v...)
}

func (ll *localLogger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LoggerLevelError, &format, v...)
}

func (ll *localLogger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LoggerLevelWarn, &format, v...)
}

func (ll *localLogger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LoggerLevelNotice, &format, v...)
}

func (ll *localLogger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LoggerLevelInfo, &format, v...)
}

func (ll *localLogger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LoggerLevelDebug, &format, v...)
}

func (ll *localLogger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LoggerLevelTrace, &format, v...)
}
