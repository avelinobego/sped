package configs

import (
	"context"
	"io"
	"log/slog"
	"os"
)

const LevelFatal = slog.Level(12) // Error is 8, so 12 is higher

type Logger struct {
	log *slog.Logger
}

func NewLogger(writers ...io.Writer) *Logger {

	opts := &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// 1. Rename 'level' to 'status' or 'severity' for Datadog
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				a.Key = "status"
				if level == LevelFatal {
					a.Value = slog.StringValue("FATAL")
				} else {
					a.Value = slog.StringValue(level.String())
				}
				return slog.Attr{Key: "status", Value: a.Value}
			}
			// 2. Rename 'msg' to 'message' (Datadog's default message key)
			if a.Key == slog.MessageKey {
				return slog.Attr{Key: "message", Value: a.Value}
			}

			return a
		},
	}

	result := &Logger{}

	set := make(map[io.Writer]struct{}, len(writers))

	set[os.Stdout] = struct{}{}

	for _, w := range writers {
		set[w] = struct{}{}
	}

	writers = writers[:0]
	for w := range set {
		writers = append(writers, w)
	}

	multiWriter := io.MultiWriter(writers...)
	json_handler := slog.NewJSONHandler(multiWriter, opts)
	instance := slog.New(json_handler)

	result.log = instance

	return result
}

func (l *Logger) Info(msg string, attrs ...any) {
	l.log.Info(msg, attrs...)
}

func (l *Logger) Fatal(msg string, attrs ...any) {
	l.log.Log(context.Background(), LevelFatal, msg, attrs...)
}

func (l *Logger) Error(msg string, attrs ...any) {
	l.log.Error(msg, attrs...)
}

func (l *Logger) Warn(msg string, attrs ...any) {
	l.log.Warn(msg, attrs...)
}

func (l *Logger) Debug(msg string, attrs ...any) {
	l.log.Debug(msg, attrs...)
}
