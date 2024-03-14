module github.com/elastic/go-elasticsearch/v8/_examples/fasthttp

go 1.21

toolchain go1.21.8

replace github.com/elastic/go-elasticsearch/v8 => ../..

require (
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210817150010-57d659deaca7
	github.com/valyala/fasthttp v1.34.0
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/elastic/elastic-transport-go/v8 v8.4.0 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	go.opentelemetry.io/otel v1.21.0 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.21.0 // indirect
)
