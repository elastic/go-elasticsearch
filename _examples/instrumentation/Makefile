test:
	go build -o /dev/null expvar.go
	go build -o /dev/null opencensus.go
	go build -o /dev/null apmelasticsearch.go

setup:
	@go mod tidy
	@go mod download

.PHONY: test setup
