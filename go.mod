module github.com/elastic/go-elasticsearch/v8

go 1.21

toolchain go1.21.0

require (
	github.com/elastic/elastic-transport-go/v8 v8.7.0
	go.opentelemetry.io/otel/trace v1.28.0
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel v1.28.0 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
)

retract v8.18.0
