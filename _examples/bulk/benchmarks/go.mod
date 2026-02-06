module github.com/elastic/go-elasticsearch/v9/_examples/bulk/benchmarks

go 1.24.0

toolchain go1.24.13

replace github.com/elastic/go-elasticsearch/v9 => ../../..

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/elastic/elastic-transport-go/v8 v8.8.0
	github.com/elastic/go-elasticsearch/v9 v9.0.0-20210817150010-57d659deaca7
	github.com/mailru/easyjson v0.9.1
	github.com/montanaflynn/stats v0.6.3
	github.com/valyala/fasthttp v1.69.0
)

require (
	github.com/andybalholm/brotli v1.2.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.18.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
)
