package logging

import (
	"io"
	"log/slog"

	"github.com/lmittmann/tint"
)

type SlogLogger struct {
	logger *slog.Logger
	level  *slog.LevelVar
}

func NewSlogLogger(writer io.Writer, level slog.Level) SlogLogger {
	levelVar := new(slog.LevelVar)
	levelVar.Set(level)
	return SlogLogger{
		logger: slog.New(tint.NewHandler(writer, &tint.Options{Level: levelVar})),
		level:  levelVar,
	}
}

func (s SlogLogger) Debug(msg string, args ...any) {
	s.logger.Debug(msg, args...)
}

func (s SlogLogger) Info(msg string, args ...any) {
	s.logger.Info(msg, args...)
}

func (s SlogLogger) Warn(msg string, args ...any) {
	s.logger.Warn(msg, args...)
}

func (s SlogLogger) Error(msg string, args ...any) {
	s.logger.Error(msg, args...)
}

func (s *SlogLogger) SetLevel(level slog.Level) {
	s.level.Set(level)
}
