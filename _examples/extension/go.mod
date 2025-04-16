module github.com/elastic/go-elasticsearch/v8/_examples/extension

go 1.23

toolchain go1.24.2

replace github.com/elastic/go-elasticsearch/v8 => ../..

require (
	github.com/elastic/elastic-transport-go/v8 v8.7.0
	github.com/elastic/go-elasticsearch/v8 v8.0.0-00010101000000-000000000000
	github.com/elastic/go-elasticsearch/v9 v9.0.0-20250416133943-b5adae5afb7e
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
)
