GO_TEST_CMD = $(if $(shell which richgo),richgo test,go test)

test:  ## Run tests
	go run default.go
	go run custom.go

.PHONY: test
