package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"
)

type slogLogEntry struct {
	base   *slogLogger
	fields Fields
}

type slogLogger struct {
	logger *slog.Logger
}

// NewSlogLogger erstellt einen neuen Logger mit slog Logger
func NewSlogLogger(logger *slog.Logger) (Logger, error) {
	return &slogLogger{
		logger: logger,
	}, nil
}

func (l *slogLogger) doLog(level slog.Level, msg string, attrs ...slog.Attr) {
	if !l.logger.Enabled(context.Background(), level) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, log]
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	if len(attrs) > 0 {
		r.AddAttrs(attrs...)
	}
	_ = l.logger.Handler().Handle(context.Background(), r)
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *slogLogger) Debug(args ...any) {
	l.doLog(slog.LevelDebug, fmt.Sprint(args...))
}

// Info uses fmt.Sprint to construct and log a message.
func (l *slogLogger) Info(args ...any) {
	l.doLog(slog.LevelInfo, fmt.Sprint(args...))
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *slogLogger) Warn(args ...any) {
	l.doLog(slog.LevelWarn, fmt.Sprint(args...))
}

// Error uses fmt.Sprint to construct and log a message.
func (l *slogLogger) Error(args ...any) {
	l.doLog(slog.LevelError, fmt.Sprint(args...))
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *slogLogger) Panic(args ...any) {
	msg := fmt.Sprint(args...)
	l.doLog(slog.LevelError, msg)
	panic(msg)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *slogLogger) Fatal(args ...any) {
	l.doLog(slog.LevelError, fmt.Sprint(args...))
	os.Exit(1)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *slogLogger) Debugf(template string, args ...any) {
	l.doLog(slog.LevelDebug, fmt.Sprintf(template, args...))
}

func (l *slogLogger) Infof(template string, args ...any) {
	l.doLog(slog.LevelInfo, fmt.Sprintf(template, args...))
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *slogLogger) Warnf(template string, args ...any) {
	l.doLog(slog.LevelWarn, fmt.Sprintf(template, args...))
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *slogLogger) Errorf(template string, args ...any) {
	l.doLog(slog.LevelError, fmt.Sprintf(template, args...))
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *slogLogger) Panicf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.doLog(slog.LevelError, msg)
	panic(msg)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *slogLogger) Fatalf(template string, args ...any) {
	l.doLog(slog.LevelError, fmt.Sprintf(template, args...))
	os.Exit(1)
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func (l *slogLogger) WithFields(fields Fields) Logger {
	return &slogLogEntry{
		base:   l,
		fields: fields,
	}
}

func (l *slogLogEntry) doLogWithFields(level slog.Level, msg string) {
	var attrs []slog.Attr
	for k, v := range l.fields {
		attrs = append(attrs, slog.Attr{Key: k, Value: slog.AnyValue(v)})
	}
	l.base.doLog(level, msg, attrs...)
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *slogLogEntry) Debug(args ...any) {
	l.doLogWithFields(slog.LevelDebug, fmt.Sprint(args...))
}

// Info uses fmt.Sprint to construct and log a message.
func (l *slogLogEntry) Info(args ...any) {
	l.doLogWithFields(slog.LevelInfo, fmt.Sprint(args...))
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *slogLogEntry) Warn(args ...any) {
	l.doLogWithFields(slog.LevelWarn, fmt.Sprint(args...))
}

// Error uses fmt.Sprint to construct and log a message.
func (l *slogLogEntry) Error(args ...any) {
	l.doLogWithFields(slog.LevelError, fmt.Sprint(args...))
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *slogLogEntry) Panic(args ...any) {
	msg := fmt.Sprint(args...)
	l.doLogWithFields(slog.LevelError, msg)
	panic(msg)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *slogLogEntry) Fatal(args ...any) {
	l.doLogWithFields(slog.LevelError, fmt.Sprint(args...))
	os.Exit(1)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *slogLogEntry) Debugf(template string, args ...any) {
	l.doLogWithFields(slog.LevelDebug, fmt.Sprintf(template, args...))
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *slogLogEntry) Infof(template string, args ...any) {
	l.doLogWithFields(slog.LevelInfo, fmt.Sprintf(template, args...))
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *slogLogEntry) Warnf(template string, args ...any) {
	l.doLogWithFields(slog.LevelWarn, fmt.Sprintf(template, args...))
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *slogLogEntry) Errorf(template string, args ...any) {
	l.doLogWithFields(slog.LevelError, fmt.Sprintf(template, args...))
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *slogLogEntry) Panicf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.doLogWithFields(slog.LevelError, msg)
	panic(msg)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *slogLogEntry) Fatalf(template string, args ...any) {
	l.doLogWithFields(slog.LevelError, fmt.Sprintf(template, args...))
}

// WithFields adds fields to the logging context
func (l *slogLogEntry) WithFields(fields Fields) Logger {
	var allFields = make(Fields, len(l.fields)+len(fields))
	for k, v := range fields {
		allFields[k] = v
	}
	for k, v := range l.fields {
		allFields[k] = v
	}
	return l.base.WithFields(allFields)
}
