module github.com/elastic/go-elasticsearch/v8/_examples/xkcdsearch

go 1.17

replace github.com/elastic/go-elasticsearch/v8 => ../..

require (
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210817150010-57d659deaca7
	github.com/rs/zerolog v1.11.0
	github.com/spf13/cobra v0.0.3
	golang.org/x/crypto v0.1.0
)

require (
	github.com/elastic/elastic-transport-go/v8 v8.3.1-0.20231219154625-903383caf56b // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.3 // indirect
	go.opentelemetry.io/otel v1.21.0 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.21.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/term v0.1.0 // indirect
)
