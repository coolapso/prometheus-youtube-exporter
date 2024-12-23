# PROMETHEUS YOUTUBE EXPORTER

[![release](https://github.com/coolapso/prometheus-youtube-exporter/actions/workflows/release.yaml/badge.svg)](https://github.com/coolapso/prometheus-youtube-exporter/actions/workflows/release.yaml)
![GitHub Tag](https://img.shields.io/github/v/tag/coolapso/prometheus-youtube-exporter?logo=semver&label=semver&labelColor=gray&color=green)
[![Docker image version](https://img.shields.io/docker/v/coolapso/youtube-exporter/latest?logo=docker)](https://hub.docker.com/r/coolapso/youtube-exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/coolapso/prometheus-youtube-exporter)](https://goreportcard.com/report/github.com/coolapso/prometheus-youtube-exporter)
![GitHub Sponsors](https://img.shields.io/github/sponsors/coolapso?style=flat&logo=githubsponsors)

A Prometheus exporter to scrape metrics from the YouTube API.

I currently use this exporter solely to monitor the status of my 24/7 live stream @thearcticskies, so it is the only feature available at the moment. If you need additional features, feel free to open a feature request or submit a pull request with your contribution.

## Install

Currently you can only grab one of the binaries provided in the releases page, or run it using docker. Check each use case examples for more details.

## Exported Metrics

| Metric | Meaning | Labels | type |
| ------ | ------- | ------ | ---- |
| youtube_channel_is_live | If Youtube channel live stream is broadcasting | channel_name | gauge |

## Usage
* Create an API key. For instructions, see [Google API Support](https://support.google.com/googleapi/answer/6158862?hl=en&ref_topic=7013279).
* Start the the exporter providing at least the API Key and a list of channels to monitor 

```
Prometheus Youtube Exporter

Usage:
  youtube-exporter [flags]

Flags:
      --address string        The address to access the exporter used for oauth redirect uri (default "localhost")
      --api.key string        Youtube API Key (default "localhost")
      --channel.ids strings   The ids of youttube channels to monitor
  -h, --help                  help for youtube-exporter
      --listen.port string    Port to listen at (default "9101")
      --log.format string     Exporter log format, text or json (default "text")
      --log.level string      Exporter log level (default "info")
      --metrics.path string   Path to expose metrics at (default "/metrics")
```

You can also use environment variables. The most accurate list for them is available [here](cmd/root.go).

### Example 

*Using flags:*
```
./youtube-exporter --api.key <ApiKey> --channel.ids "foo,bar"
```

*Using Environment Variables:*
```
YT_CHANNEL_IDS="foo bar" YT_API_KEY="APIKEY" ./youtube-exporter
```

*With docker, using flags:*
```
docker run -d -p 10020:10020 \
        coolapso/youtube-exporter \
        --api.key=apiKey \
        --channel.ids="chan1,chan2,chan3"
```

*With Docker, using environment variables:*
```
docker run -d -p 10020:10020 \
        -e YT_API_KEY=<ApiKey> \
        -e YT_CHANNEL_IDS="chan1 chan2 chan3" \
        coolapso/youtube-exporter
```

## Troubleshooting

### Exporter reports broadcast is offline even when it is actively broacasting

Make sure you have the option: "Unlist live replay once stream ends" Turned off when you start the live stream, Youtube API returns different values than expected when this setting is turnned on. More details on [issue #19](https://github.com/coolapso/prometheus-youtube-exporter/issues/19)



# Contributions

Improvements and suggestions are always welcome. Feel free to check for any open issues, or open a new issue or pull request.

If you like this project and want to support / contribute in a different way you can always [:heart: Sponsor Me](https://github.com/sponsors/coolapso) or

<a href="https://www.buymeacoffee.com/coolapso" target="_blank">
  <img src="https://cdn.buymeacoffee.com/buttons/default-yellow.png" alt="Buy Me A Coffee" style="height: 51px !important;width: 217px !important;" />
</a>
