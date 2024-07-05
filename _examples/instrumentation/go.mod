module github.com/elastic/go-elasticsearch/v8/_examples/instrumentation/opencensus

go 1.21
toolchain go1.21.0

replace github.com/elastic/go-elasticsearch/v8 => ../..

require (
	github.com/elastic/elastic-transport-go/v8 v8.6.0
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20191002063538-b491ce54d752
	github.com/fatih/color v1.7.0
	go.elastic.co/apm v1.11.0
	go.elastic.co/apm/module/apmelasticsearch v1.5.0
	go.opencensus.io v0.19.0
	golang.org/x/crypto v0.17.0
)

require (
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/elastic/go-sysinfo v1.1.1 // indirect
	github.com/elastic/go-windows v1.0.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/joeshaw/multierror v0.0.0-20140124173710-69b34d4ec901 // indirect
	github.com/mattn/go-colorable v0.1.0 // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/prometheus/procfs v0.0.3 // indirect
	github.com/santhosh-tekuri/jsonschema v1.2.4 // indirect
	go.elastic.co/apm/module/apmhttp v1.5.0 // indirect
	go.elastic.co/fastjson v1.1.0 // indirect
	go.opentelemetry.io/otel v1.28.0 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.28.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/term v0.15.0 // indirect
	howett.net/plist v0.0.0-20181124034731-591f970eefbb // indirect
)
