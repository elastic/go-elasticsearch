GO_TEST_CMD = $(if $(shell which richgo),richgo test,go test)
export ELASTICSEARCH_URL=http://elastic:elastic@localhost:9200

test:  ## Run tests
	$(GO_TEST_CMD) -v ./...

setup:

.PHONY: test setup
