package logging

import (
	"io"
	"log/slog"

	"github.com/lmittmann/tint"
)

type slogLogger struct {
	Logger *slog.Logger
}

func NewSlogLogger(writer io.Writer, level slog.Level) *slogLogger {
	return &slogLogger{
		Logger: slog.New(tint.NewHandler(writer, &tint.Options{Level: level})),
	}
}

func (s *slogLogger) Debug(msg string, args ...any) {
	s.Logger.Debug(msg, args...)
}

func (s *slogLogger) Info(msg string, args ...any) {
	s.Logger.Info(msg, args...)
}

func (s *slogLogger) Warn(msg string, args ...any) {
	s.Logger.Warn(msg, args...)
}

func (s *slogLogger) Error(msg string, args ...any) {
	s.Logger.Error(msg, args...)
}
