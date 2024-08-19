package httpServer

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/coolapso/prometheus-youtube-exporter/internal/collectors"
	"github.com/coolapso/prometheus-youtube-exporter/internal/slogLogger"
)

func TestNewServer(t *testing.T) {
	logger, err := slogLogger.NewLogger("info", "text")
	if err != nil {
		t.Fatal("Failed to create new logger:", err)
	}

	settings := &collectors.Settings{
		LogLevel:    "debug",
		LogFormat:   "text",
		ListenPort:  "8080",
		MetricsPath: "/metrics",
		Address:     "localhost",
		ApiKey:      "SomeAPIKey",
	}

	exporter, err := collectors.NewExporter(settings, logger)
	if err != nil {
		t.Fatalf("Failed to create new exporter: %v", err)
	}

	NewServer(exporter) // This binds handlers to the default mux.

	// Testing the root handler
	testRootHandler(t)

	// Testing the metrics handler
	testMetricsHandler(t)
}

func testRootHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Root handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedBodyContains := "<title>Prometheus Youtube exporter</title>"
	if body := recorder.Body.String(); !strings.Contains(body, expectedBodyContains) {
		t.Errorf("Root handler returned unexpected body: does not contain %v", expectedBodyContains)
	}
}

func testMetricsHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "/metrics", nil)
	recorder := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Metrics handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if body := recorder.Body.String(); len(body) == 0 {
		t.Error("Metrics handler returned empty body")
	}
}
