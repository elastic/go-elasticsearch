module github.com/elastic/go-elasticsearch/v8/_examples/security

go 1.22

toolchain go1.22.0

replace github.com/elastic/go-elasticsearch/v8 => ../..

require github.com/elastic/go-elasticsearch/v8 v8.0.0-00010101000000-000000000000

require (
	github.com/elastic/elastic-transport-go/v8 v8.6.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel v1.28.0 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.28.0 // indirect
)
