module github.com/elastic/go-elasticsearch/v9/_examples/bulk

go 1.24

toolchain go1.24.13

replace github.com/elastic/go-elasticsearch/v9 => ../..

require (
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/dustin/go-humanize v1.0.1
	github.com/elastic/go-elasticsearch/v9 v9.0.0-00010101000000-000000000000
)

require (
	github.com/elastic/elastic-transport-go/v8 v8.8.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
)
