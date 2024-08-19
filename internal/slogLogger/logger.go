package slogLogger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func NewLogger(level, format string) (logger *slog.Logger, err error) {
	var slogHandlerOptions slog.HandlerOptions
	var slogHandler slog.Handler

	switch strings.ToLower(level) {
	case "debug":
		slogHandlerOptions = slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
	case "info":
		slogHandlerOptions = slog.HandlerOptions{
			Level: slog.LevelInfo,
		}
	case "warn":
		slogHandlerOptions = slog.HandlerOptions{
			Level: slog.LevelWarn,
		}
	case "error":
		slogHandlerOptions = slog.HandlerOptions{
			Level: slog.LevelError,
		}
	default:
		slogHandlerOptions = slog.HandlerOptions{
			Level: slog.LevelInfo,
		}
		err = fmt.Errorf("Log level not recognized, falling back to info")
	}

	switch strings.ToLower(format) {
	case "text":
		slogHandler = slog.NewTextHandler(os.Stdout, &slogHandlerOptions)
	case "json":
		slogHandler = slog.NewJSONHandler(os.Stdout, &slogHandlerOptions)
	default:
		slogHandler = slog.NewTextHandler(os.Stdout, &slogHandlerOptions)
		err = fmt.Errorf("Log format not recognized, falling back to text")
	}

	return slog.New(slogHandler), err
}
