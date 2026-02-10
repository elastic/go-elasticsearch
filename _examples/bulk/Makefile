GO_TEST_CMD = $(if $(shell which richgo),richgo test,go test)
export ELASTICSEARCH_URL=http://elastic:elastic@localhost:9200

test: test-default test-indexer test-typed

test-default:
	go run default.go

test-indexer:
	go run indexer.go

test-typed:
	go run -tags bulk_typed typed.go

test-benchmarks: clean setup
	cd benchmarks && go run benchmarks.go

setup:
	go install github.com/mailru/easyjson/easyjson@latest
	cd benchmarks && go generate ./model

clean:
	@rm -f benchmarks/model/*_easyjson.go

.PHONY: test test-default test-indexer test-typed test-benchmarks setup clean
