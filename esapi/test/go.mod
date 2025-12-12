module github.com/elastic/go-elasticsearch/v8/esapi/test

go 1.21

toolchain go1.22.5

replace github.com/elastic/go-elasticsearch/v8 => ../../

require (
	github.com/elastic/elastic-transport-go/v8 v8.7.0
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210817150010-57d659deaca7
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
)
