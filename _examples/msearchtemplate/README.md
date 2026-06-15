# Example: MSearchTemplate

The [`default.go`](default.go) example demonstrates two ways to call the
[`_msearch/template`](https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-msearch-template)
API with the Go client:

- `esapi` (`Client`) with a raw NDJSON payload.
- `typedapi` (`TypedClient`) with `msearchtemplate.Request`.

```bash
make test
# or
go run -tags msearchtemplate_default default.go
```

## Notes

- The request body must be NDJSON with alternating lines: header, template, header, template.
- Always end the body with a trailing newline (`\n`).
- For the low-level `esapi` call, set `Content-Type: application/x-ndjson`.
