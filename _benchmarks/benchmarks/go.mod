module github.com/elastic/go-elasticsearch/v8/benchmarks

go 1.21

toolchain go1.24.1

replace github.com/elastic/go-elasticsearch/v8 => ../../

require (
	github.com/elastic/elastic-transport-go/v8 v8.8.0
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20200408073057-6f36a473b19f
	github.com/fatih/color v1.7.0
	github.com/montanaflynn/stats v0.6.3
	github.com/tidwall/gjson v1.9.3
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	go.opentelemetry.io/otel v1.28.0 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.28.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
)
