GO_TEST_CMD = $(if $(shell which richgo),richgo test,go test)
export ELASTICSEARCH_URL=http://elastic:elastic@localhost:9200

test: clean setup ## Run tests
	go run gjson.go
	go generate ./model && go run easyjson.go
	go run jsonreader.go

bench: clean setup # Run benchmarks
	go test -run=none -bench=. -benchmem benchmark_test.go

setup: ## Install & Setup
	go get -u github.com/mailru/easyjson/...

clean: ## Remove artifacts
	@rm -f model/*_easyjson.go

.PHONY: clean bench setup test
