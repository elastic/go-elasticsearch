module github.com/elastic/go-elasticsearch/v8/_examples/bulk/benchmarks

go 1.21

toolchain go1.21.8

replace github.com/elastic/go-elasticsearch/v8 => ../../..

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210817150010-57d659deaca7
	github.com/mailru/easyjson v0.7.1
	github.com/montanaflynn/stats v0.6.3
)

require (
	github.com/elastic/elastic-transport-go/v8 v8.5.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
)
