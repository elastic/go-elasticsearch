GO_TEST_CMD = $(if $(shell which richgo),richgo test,go test)

test: clean setup test-default test-indexer

test-default:
	go run default.go

test-indexer:
	go run indexer.go

setup:
	@go get -u github.com/mailru/easyjson/...
	go generate ./model

clean:
	@rm -f model/*_easyjson.go

.PHONY: test test-default test-indexer setup clean
