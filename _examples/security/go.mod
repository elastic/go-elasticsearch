module github.com/elastic/go-elasticsearch/v8/_examples/security

go 1.21

toolchain go1.21.0

replace github.com/elastic/go-elasticsearch/v8 => ../..

require github.com/elastic/go-elasticsearch/v8 v8.0.0-00010101000000-000000000000

require (
	github.com/elastic/elastic-transport-go/v8 v8.9.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel v1.29.0 // indirect
	go.opentelemetry.io/otel/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.29.0 // indirect
)
