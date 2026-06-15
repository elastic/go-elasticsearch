GO_TEST_CMD = $(if $(shell which richgo),richgo test,go test)
export ELASTICSEARCH_URL=http://elastic:elastic@localhost:9200

test:  ## Run tests
	go run ./cmd/auth_provider/main.go
	go run ./cmd/context_auth/main.go
	go run ./cmd/custom_auth/main.go
	go run ./cmd/custom_observability/main.go

setup:

.PHONY: test setup
