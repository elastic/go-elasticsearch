GO_TEST_CMD = $(if $(shell which richgo),richgo test,go test)

test:  ## Run tests
	$(GO_TEST_CMD) -v ./...

setup:
	@go mod tidy
	@go mod download

.PHONY: test setup
