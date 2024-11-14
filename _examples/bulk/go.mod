module github.com/elastic/go-elasticsearch/v8/_examples/bulk

go 1.22

toolchain go1.22.0

replace github.com/elastic/go-elasticsearch/v8 => ../..

require (
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/dustin/go-humanize v1.0.1
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20211123103400-5f8a17a2322f
)

require (
	github.com/elastic/elastic-transport-go/v8 v8.6.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel v1.28.0 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.28.0 // indirect
)
