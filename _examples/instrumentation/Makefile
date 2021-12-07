export ELASTICSEARCH_URL=http://elastic:elastic@localhost:9200

test:
	go build -o /dev/null expvar.go
	go build -o /dev/null opencensus.go
	go build -o /dev/null apmelasticsearch.go

setup:
	@go mod download

.PHONY: test setup
