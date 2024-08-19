package httpServer

import (
	"github.com/coolapso/prometheus-youtube-exporter/internal/collectors"
	"github.com/prometheus/client_golang/prometheus"
	promCollectors "github.com/prometheus/client_golang/prometheus/collectors"
	versioncollector "github.com/prometheus/client_golang/prometheus/collectors/version"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"html/template"
	"net/http"
)

const (
	rootTemplate string = `<html>
	 <head><title>Prometheus Youtube exporter</title></head>
	 <body>
		 <h1>Prometheus Youtube Exporter</h1>
		 <p>Metrics at: <a href='{{ .MetricsPath }}'>{{ .MetricsPath }}</a></p>
		 <p>Source: <a href='https://github.com/coolapso/prometheus-youtube-exporter'>github.com/coolapso/prometheus-youtube-exporter</a></p>
	 </body>
	 </html>`
)

// Serves root page with html template on root page
// Serves Metrics on settings.MetricsPath
func NewServer(e *collectors.Exporter) *http.Server {
	s := e.Settings
	t := template.Must(template.New("root").Parse(rootTemplate))

	reg := prometheus.NewRegistry()
	reg.MustRegister(
		e,
		versioncollector.NewCollector("exporter"),
		promCollectors.NewBuildInfoCollector(),
		promCollectors.NewGoCollector(),
	)

	promHandlerOpts := promhttp.HandlerOpts{
		Registry: reg,
	}

	// Metrics handler
	http.Handle(s.MetricsPath, promhttp.HandlerFor(reg, promHandlerOpts))

	// Root Page handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, e.Settings)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	return &http.Server{Addr: ":" + s.ListenPort}
}
