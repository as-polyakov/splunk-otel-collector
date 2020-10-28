module github.com/parallelstream/splunk-otel-collector

go 1.14

require github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter v0.13.1

require go.opentelemetry.io/collector v0.13.1-0.20201020175630-99cb5b244aad
replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.0.0-00010101000000-000000000000 => github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.13.1

//replace go.opentelemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter v0.13.1 => /Users/apolyakov/dev/opentelemetry-collector-contrib/exporter/splunkhecexporter

//replace github.com/open-telemetry/opentelemetry-collector-contrib v0.13.1 => /Users/apolyakov/dev/opentelemetry-collector-contrib
