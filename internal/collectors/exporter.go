package collectors

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"log/slog"
)

// Namespace constant value prefixed on metrics boilerplate_
const (
	namespace = "youtube_channel"
)

// exporter settings, more settings can be added if needed
type Settings struct {
	LogLevel    string
	LogFormat   string
	MetricsPath string
	ListenPort  string
	Address     string
	ChannelIds  []string
	ApiKey      string
}

type metrics struct {
	IsLive *prometheus.Desc
}

type Exporter struct {
	client   *youtube.Service
	metrics  *metrics
	Settings *Settings
	Logger   *slog.Logger
}

// Describes all metrics
// implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.metrics.IsLive
}

// Collects metrics configured and returns them as prometheus metrics
// implements prometheus.Collector
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	for _, ytChan := range e.Settings.ChannelIds {
		ch <- prometheus.MustNewConstMetric(
			e.metrics.IsLive,
			prometheus.GaugeValue,
			e.isLive(ytChan),
			"isLive",
		)
	}
}

// Functions that gather and return the metric values
func (e *Exporter) isLive(chId string) float64 {
	call := e.client.Search.List([]string{"snippet"}).Type("video").ChannelId(chId).Order("date")
	resp, err := call.Do()
	if err != nil {
		e.Logger.Error("Failed to request videos", "err", err)
		return 0
	}

	if resp.HTTPStatusCode != 200 {
		e.Logger.Error("Failed to request videos", "statusCode", resp.HTTPStatusCode, "err", resp.ServerResponse)
		return 0
	}

	for _, video := range resp.Items {
		if video.Snippet.LiveBroadcastContent == "live" {
			return 1
		}
	}

	return 0
}

// Initializes the metrics
func newMetrics() *metrics {
	return &metrics{
		IsLive: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "isLive"),
			"If Youtube channel live stream is broadcasting",
			[]string{"channel_name"}, nil,
		),
	}
}

// Initializes the exporter
func NewExporter(s *Settings, logger *slog.Logger) (*Exporter, error) {
	client, err := youtube.NewService(context.Background(), option.WithAPIKey(s.ApiKey))
	if err != nil {
		log.Fatalf("Failed to create youtube client: %v", err)
	}
	metrics := newMetrics()

	exporter := &Exporter{
		client:   client,
		metrics:  metrics,
		Settings: s,
		Logger:   logger,
	}

	return exporter, nil
}
