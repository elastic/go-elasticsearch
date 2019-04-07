module github.com/elastic/go-elasticsearch/v5/_examples/instrumentation/opencensus

go 1.11

replace github.com/elastic/go-elasticsearch/v5 => ../..

require (
	github.com/elastic/go-elasticsearch/v5 5.x

	github.com/fatih/color v1.7.0
	github.com/mattn/go-colorable v0.1.0 // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect

	go.elastic.co/apm v1.2.1-0.20190212105052-525c5c425aa9
	go.elastic.co/apm/module/apmelasticsearch v1.2.1-0.20190212105052-525c5c425aa9

	go.opencensus.io v0.19.0

	golang.org/x/crypto v0.0.0-20190211182817-74369b46fc67
)
