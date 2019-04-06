# Example: Logging

## Default Loggers

The [`default.go`](default.go) example demonstrates how to use the bundled loggers:

* `estransport.TextLogger`
* `estransport.ColorLogger`
* `estransport.CurlLogger`
* `estransport.JSONLogger`

```bash
go run default.go

# ...
# ███ Color ██████████████████████████████████████████████████████████████████████
# DELETE http://localhost:9200/test/_doc/1 404 Not Found 9ms
# HEAD http://localhost:9200/test/_doc/1 404 Not Found 9ms
# ...
```
