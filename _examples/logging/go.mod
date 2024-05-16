module github.com/elastic/go-elasticsearch/v8/_examples/logging

go 1.21

toolchain go1.21.8

replace github.com/elastic/go-elasticsearch/v8 => ../..

require (
	github.com/elastic/elastic-transport-go/v8 v8.5.0
	github.com/elastic/go-elasticsearch/v8 v8.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.32.0
)

require (
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
)
