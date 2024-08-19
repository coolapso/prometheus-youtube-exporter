package collectors

import (
	"github.com/coolapso/prometheus-youtube-exporter/internal/slogLogger"
	"strings"
	"testing"
)

func TestNewMetrics(t *testing.T) {
	metrics := newMetrics()

	t.Run("Test Sample metric A", func(t *testing.T) {
		expected := `Desc{fqName: "youtube_channel_isLive", help: "If Youtube channel live stream is broadcasting", constLabels: {}, variableLabels: {channel_name}}`
		got := metrics.IsLive.String()
		if !strings.Contains(got, expected) {
			t.Fatalf("Metric does not contain expected fqName, expected: %v, got %v", expected, got)
		}
	})
}

func TestNewExporter(t *testing.T) {
	logger, _ := slogLogger.NewLogger("info", "text")
	settings := &Settings{
		LogLevel:    "debug",
		LogFormat:   "text",
		ListenPort:  "8080",
		MetricsPath: "/metrics",
		Address:     "localhost",
		ApiKey:      "SomeAPIKey",
	}

	_, err := NewExporter(settings, logger)
	if err != nil {
		t.Fatal("Failed to create new exporter:", err)
	}
}
