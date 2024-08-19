package slogLogger

import (
	"context"
	"log/slog"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name          string
		LogLevel      string
		LogFormat     string
		expectedLevel slog.Level
		expectedErr   string
	}{
		{
			name:          "Valid debug level",
			LogLevel:      "debug",
			LogFormat:     "text",
			expectedLevel: slog.LevelDebug,
			expectedErr:   "",
		},
		{
			name:          "Valid error level",
			LogLevel:      "error",
			LogFormat:     "json",
			expectedLevel: slog.LevelError,
			expectedErr:   "",
		},
		{
			name:          "Invalid log level",
			LogLevel:      "invalid",
			LogFormat:     "text",
			expectedLevel: slog.LevelInfo,
			expectedErr:   "Log level not recognized, falling back to info",
		},
		{
			name:          "Invalid log format",
			LogLevel:      "info",
			LogFormat:     "invalid",
			expectedLevel: slog.LevelInfo,
			expectedErr:   "Log format not recognized, falling back to text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := NewLogger(tt.LogLevel, tt.LogFormat)
			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("expected error: %v, got: %v", tt.expectedErr, err)
			}

			if err == nil && tt.expectedErr != "" {
				t.Errorf("expected error: %v, got: nil", tt.expectedErr)
			}

			if logger == nil {
				t.Fatalf("expected logger to be non-nil")
			}

			if !logger.Handler().Enabled(context.TODO(), tt.expectedLevel) {
				t.Fatalf("Logger not logging level: %v", tt.expectedLevel)
			}

			if tt.expectedLevel == slog.LevelError && logger.Handler().Enabled(context.TODO(), slog.LevelDebug) {
				t.Fatalf("Expected Handler logger to return false")
			}
		})
	}
}
