GO_TEST_CMD = $(if $(shell which richgo),richgo test,go test)

test:  ## Run tests
	go run default.go
	go run custom.go

setup:
	@go mod download
	@go mod download github.com/rs/zerolog

.PHONY: test setup
