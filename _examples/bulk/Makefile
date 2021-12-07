GO_TEST_CMD = $(if $(shell which richgo),richgo test,go test)
export ELASTICSEARCH_URL=http://elastic:elastic@localhost:9200

test: test-default test-indexer

test-default:
	go run default.go

test-indexer:
	go run indexer.go

test-benchmarks: clean setup
	cd benchmarks && go run benchmarks.go

setup:
	go get -u github.com/mailru/easyjson/...
	go generate ./model

clean:
	@rm -f benchmarks/model/*_easyjson.go

.PHONY: test test-default test-indexer test-benchmarks setup clean
