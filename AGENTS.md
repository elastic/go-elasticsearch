# AGENTS.md

This file provides guidance to AI coding agents working with code in this repository.
CLAUDE.md is a symlink to this file — they are the same file, do not update both.

Official Elasticsearch Go client. Module: `github.com/elastic/go-elasticsearch/v9`.

## Backwards Compatibility

This is a public library with external consumers. Maintain backwards compatibility:

- Do NOT remove or rename exported types, functions, methods, or fields.
- Do NOT change function/method signatures (parameter types, return types, receiver types).
- Do NOT change the behavior of existing public APIs in a breaking way.
- Adding new exported symbols, new fields to config structs, or new methods is fine.
- If a deprecation is needed, add a new API and mark the old one deprecated via godoc comment; do not remove it.

## Commands

```bash
make test-unit                  # unit tests
go test -v -run TestName ./pkg  # single test
make test-integ                 # integration tests (requires ES)
make test-bench                 # benchmarks
make lint                       # go vet
make go-mod-tidy-all            # tidy all modules
make cluster                    # start ES in Docker
make gen-api                    # regenerate esapi/ from specs
```

## Key Rules

- **DO NOT manually edit files in `esapi/` or `typedapi/`** — they are generated. Use `make gen-api`.
- All Go files must include the Apache 2.0 license header:

```go
// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
```

- Integration tests use build tag `//go:build integration`, unit tests use `//go:build !integration`.
- All exported types and functions must have godoc comments.
- Before finishing, always run: `make lint && make test-unit`.

## Regeneration

- `make gen-api` regenerates `esapi/` from the REST spec. Commit only this output in regeneration PRs.
- `make gen-tests` regenerates the YAML-driven tests in `esapi/test/` and `internal/build/cmd/generate/commands/gentests/api_registry.gen.go`. This runs in Buildkite CI; do not commit its output locally.
- Regenerating after a generator fix or long gap will include spec drift (new endpoints, reworded descriptions, additive fields). Mention it in the PR body alongside the motivating change.
- `make download-specs` uses the Makefile's default version. Override with `ELASTICSEARCH_BUILD_VERSION=X.Y.Z make download-specs` when regenerating a version branch from a checkout whose default targets a different version.

## Commits and Pull Requests

PRs are squash-merged into a single commit when landing on `main` or a version branch. Because of this, each PR should represent one distinct, self-contained change — a single feature, bug fix, or refactor. Mixing unrelated changes in one PR means they become inseparable in history, making reviews harder and reverts riskier (reverting one change would also revert the unrelated ones).

When planning work:

- Break larger features into a sequence of small, independently mergeable PRs.
- Keep each PR focused on a single logical change with a clear purpose.
- Avoid bundling unrelated fixes or refactors into a feature PR — send them separately.
- Prefer multiple small PRs over one large PR, even if they land in quick succession.

### Breaking changes

Commits containing breaking API changes need a `BREAKING-CHANGE:` footer (preceded by a blank line) on a `fix:` or `feat:` commit. release-please reads the footer from the squash-merged commit to flag the release accordingly. Do not mark a breaking change as `chore:` or `refactor:`.

## Architecture

### Two Client Types

Both clients share `BaseClient` (transport, product-check, meta headers) but expose different API surfaces:

- **`Client`** (functional API) — embeds `*esapi.API`. Created with `NewClient(Config)` or `New(opts ...Option)`. Each API endpoint is a function-type field on the `API` struct that accepts functional options (`func(*SearchRequest)`).
- **`TypedClient`** (typed API) — embeds `*typedapi.MethodAPI`. Created with `NewTypedClient(Config)` or `NewTyped(opts ...Option)`. Each endpoint returns a fluent builder struct with typed request/response structs.

### Package Layout

| Path               | Description                                      |
| ------------------ | ------------------------------------------------ |
| `elasticsearch.go` | Client construction, Config, transport setup     |
| `options.go`       | Functional-options API (`New(opts ...Option)`)   |
| `esapi/`           | Auto-generated flat API methods (DO NOT EDIT)    |
| `typedapi/`        | Auto-generated typed API builders (DO NOT EDIT)  |
| `esutil/`          | Utility helpers (BulkIndexer, JSONReader)         |
| `internal/build/`  | Code generator (cobra CLI, separate go.mod)      |
| `internal/testing/`| Integration/e2e tests (separate go.mod)          |
| `internal/version/`| Client version constant                          |

### Generated Code Patterns

**esapi** — each endpoint file (`api.<name>.go`) contains: a function-type alias (e.g. `type Search func(...) (*Response, error)`), a `SearchRequest` struct with all params, a `Do(ctx, transport)` method that builds the HTTP request, and `WithXxx` option functions.

**typedapi** — each endpoint is a separate package (`typedapi/core/search/`). Contains a builder struct with fluent setters, `HttpRequest(ctx)` to build the request, `Perform(ctx)` for raw response, and `Do(ctx)` for typed JSON-decoded response. Types live in `typedapi/types/` (~2000 generated files).

### Multi-Module Structure

The repo contains multiple Go modules (each with its own `go.mod`):
- Root: main client library
- `internal/build/`: `esapi` code generator (uses `replace` directive to reference root)
- `internal/testing/`: integration tests
- `esapi/test/`: generated YAML-driven API tests
- `_examples/*/`: each example is a standalone module

When modifying dependencies, use `make go-mod-tidy-all` to tidy all modules.

### Instrumentation & Interceptors

Two extensibility points, both from `elastic-transport-go`:

- **Instrumentation**: lifecycle hooks per API call (Start, BeforeRequest, AfterRequest, RecordError, Close). Enabled via `Config.Instrumentation`. OTel adapter: `NewOpenTelemetryInstrumentation(provider, captureSearchBody)`.
- **Interceptors**: `[]InterceptorFunc` middleware on the transport's RoundTripper. Set once at construction via `Config.Interceptors`. Applied left-to-right (outermost first).
