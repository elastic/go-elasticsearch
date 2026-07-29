module github.com/elastic/go-elasticsearch/v9/_examples/bulk

go 1.25.0

toolchain go1.25.12

replace github.com/elastic/go-elasticsearch/v9 => ../..

require (
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/dustin/go-humanize v1.0.1
	github.com/elastic/elastic-transport-go/v8 v8.11.0
	github.com/elastic/go-elasticsearch/v9 v9.0.0-00010101000000-000000000000
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel v1.43.0 // indirect
	go.opentelemetry.io/otel/metric v1.43.0 // indirect
	go.opentelemetry.io/otel/trace v1.43.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
)
