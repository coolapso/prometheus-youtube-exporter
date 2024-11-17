/*
Copyright © 2024 cool4pso
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/coolapso/prometheus-youtube-exporter/internal/collectors"
	"github.com/coolapso/prometheus-youtube-exporter/internal/httpServer"
	"github.com/coolapso/prometheus-youtube-exporter/internal/slogLogger"

	"github.com/prometheus/common/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "youtube-exporter",
	Short: "Prometheus Youtube Exporter",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return checkCoreSettings()
	},

	Run: func(cmd *cobra.Command, args []string) {
		exporter()
	},
}

const (
	defaultLogLevel    = "info"
	defaultLogFormat   = "text"
	defaultMetricsPath = "/metrics"
	defaultListenPort  = "10020"
	defaultAddress     = "localhost"
	Version			   = "DEV"
)

var (
	settings collectors.Settings
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("LOG_LEVEL", defaultLogLevel)
	viper.SetDefault("LOG_FORMAT", defaultLogFormat)
	viper.SetDefault("METRICS_PATH", defaultMetricsPath)
	viper.SetDefault("LISTEN_PORT", defaultListenPort)
	viper.SetDefault("ADDRESS", defaultAddress)
	viper.SetDefault("YT_CHANNEL_IDS", nil)
	viper.SetDefault("YT_API_KEY", "")

	rootCmd.Flags().StringVar(&settings.LogLevel, "log.level", defaultLogLevel, "Exporter log level")
	_ = viper.BindPFlag("log.level", rootCmd.Flags().Lookup("LOG_LEVEL"))

	rootCmd.Flags().StringVar(&settings.LogFormat, "log.format", defaultLogFormat, "Exporter log format, text or json")
	_ = viper.BindPFlag("log.format", rootCmd.Flags().Lookup("LOG_FORMAT"))

	rootCmd.Flags().StringVar(&settings.MetricsPath, "metrics.path", defaultMetricsPath, "Path to expose metrics at")
	_ = viper.BindPFlag("metrics.path", rootCmd.Flags().Lookup("METRICS_PATH"))

	rootCmd.Flags().StringVar(&settings.ListenPort, "listen.port", defaultListenPort, "Port to listen at")
	_ = viper.BindPFlag("listen.port", rootCmd.Flags().Lookup("LISTEN_PORT"))

	rootCmd.Flags().StringVar(&settings.Address, "address", defaultAddress, "The address to access the exporter used for oauth redirect uri")
	_ = viper.BindPFlag("address", rootCmd.Flags().Lookup("ADDRESS"))

	rootCmd.Flags().StringVar(&settings.ApiKey, "api.key", defaultAddress, "Youtube API Key")
	_ = viper.BindPFlag("api.key", rootCmd.Flags().Lookup("YT_API_KEY"))

	rootCmd.Flags().StringSliceVar(&settings.ChannelIds, "channel.ids", nil, "The ids of youttube channels to monitor")
	_ = viper.BindPFlag("channel.ids", rootCmd.Flags().Lookup("YT_CHANNEL_IDS"))

	settings.LogLevel = viper.GetString("LOG_LEVEL")
	settings.LogFormat = viper.GetString("LOG_FORMAT")
	settings.MetricsPath = viper.GetString("METRICS_PATH")
	settings.ListenPort = viper.GetString("LISTEN_PORT")
	settings.Address = viper.GetString("ADDRESS")
	settings.ChannelIds = viper.GetStringSlice("YT_CHANNEL_IDS")
	settings.ApiKey = viper.GetString("YT_API_KEY")
}

func checkCoreSettings() error {
	s := &settings

	if s.ApiKey == "" {
		return fmt.Errorf("Missing Youtube API Key")
	}

	if s.ChannelIds == nil {
		return fmt.Errorf("Missing Channel ids")
	}

	return nil
}

func exporter() {
	s := &settings

	logger, err := slogLogger.NewLogger(s.LogLevel, s.LogFormat)
	if err != nil {
		logger.Warn(err.Error())
	}

	logger.Info(fmt.Sprintf("starting prometheus exporter %v %v", Version, version.BuildContext()))
	exporter, err := collectors.NewExporter(s, logger)
	if err != nil {
		logger.Error("Failed to create new exporter", "err", err)
		os.Exit(1)
	}

	srv := httpServer.NewServer(exporter)
	logger.Info(fmt.Sprintf("Server ready and listening on port :%v", s.ListenPort))
	log.Fatal(srv.ListenAndServe())
}
