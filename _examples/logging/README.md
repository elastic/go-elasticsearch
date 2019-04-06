# Example: Logging

## Default Loggers

The [`default.go`](default.go) example demonstrates how to use the bundled loggers.

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

## Custom Logger

The [`custom.go`](custom.go) example demonstrates how to implement a custom logger, using the `rs/zerolog` package.

```bash
go run custom.go

# 1:41PM WRN http://localhost:9200/test/_doc/1 duration=19.001 method=DELETE req_bytes=0 res_bytes=155 status_code=404
# 1:41PM WRN http://localhost:9200/test/_doc/1 duration=5.898 method=HEAD req_bytes=0 res_bytes=0 status_code=404
# 1:41PM INF http://localhost:9200/test/_doc?refresh=true duration=68.841 method=POST req_bytes=21 res_bytes=194 status_code=201
# 1:41PM WRN http://localhost:9200/_search?q=%7BFAIL duration=17.276 method=GET req_bytes=0 res_bytes=765 status_code=400
# 1:41PM INF http://localhost:9200/test/_search?size=1 duration=9.429 method=GET req_bytes=49 res_bytes=280 status_code=200
```
