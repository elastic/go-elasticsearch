module github.com/elastic/go-elasticsearch/v9/_examples/interceptor

go 1.25

toolchain go1.25.7

replace github.com/elastic/go-elasticsearch/v9 => ../..

require (
	github.com/elastic/elastic-transport-go/v8 v8.9.0
	github.com/elastic/go-elasticsearch/v9 v9.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel v1.40.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.38.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.38.0
	go.opentelemetry.io/otel/metric v1.40.0
	go.opentelemetry.io/otel/sdk v1.40.0
	go.opentelemetry.io/otel/sdk/metric v1.40.0
	go.opentelemetry.io/otel/trace v1.40.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	golang.org/x/sys v0.40.0 // indirect
)
