package logging

import (
	"io"
	"log/slog"

	"github.com/lmittmann/tint"
)

type slogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger(writer io.Writer, level slog.Level) *slogLogger {
	return &slogLogger{
		logger: slog.New(tint.NewHandler(writer, &tint.Options{Level: level})),
	}
}

func (s *slogLogger) Debug(msg string, args ...any) {
	s.logger.Debug(msg, args...)
}

func (s *slogLogger) Info(msg string, args ...any) {
	s.logger.Info(msg, args...)
}

func (s *slogLogger) Warn(msg string, args ...any) {
	s.logger.Warn(msg, args...)
}

func (s *slogLogger) Error(msg string, args ...any) {
	s.logger.Error(msg, args...)
}
