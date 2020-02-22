GO_TEST_CMD = $(if $(shell which richgo),richgo test,go test)

test: test-default test-benchmarks

test-default:
	go run default.go

test-benchmarks: clean setup
	cd benchmarks && go run main.go --runs=1 --count=10000

setup:
	@go get -u github.com/mailru/easyjson/...
	cd benchmarks && go generate ./model

clean:
	@rm -f benchmarks/model/*_easyjson.go

.PHONY: test test-default test-benchmarks setup clean
