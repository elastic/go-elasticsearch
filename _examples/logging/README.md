# Example: Logging

## Default Loggers

The [`default.go`](default.go) example demonstrates how to use the loggers included with the package.

* `estransport.TextLogger`
* `estransport.ColorLogger`
* `estransport.CurlLogger`
* `estransport.JSONLogger`

```bash
go run default.go

# ...
# DELETE http://localhost:9200/test/_doc/1 404 Not Found 8ms
#   HEAD http://localhost:9200/test/_doc/1 404 Not Found 16ms
#   POST http://localhost:9200/test/_doc?filter_path=result,_id&pretty=true&refresh=true 201 Created 45ms
#    GET http://localhost:9200/_search?q=[FAIL 400 Bad Request 21ms
#    GET http://localhost:9200/test/_search?filter_path=took,hits.hits&pretty=true&size=1 200 OK 11ms
# ...
```

## Custom Logger

The [`custom.go`](custom.go) example demonstrates how to implement a custom logger and use it with the client for structured logging.

```bash
go run custom.go

# 1:41PM WRN http://localhost:9200/test/_doc/1 duration=19.001 method=DELETE req_bytes=0 res_bytes=155 status_code=404
# 1:41PM WRN http://localhost:9200/test/_doc/1 duration=5.898 method=HEAD req_bytes=0 res_bytes=0 status_code=404
# 1:41PM INF http://localhost:9200/test/_doc?refresh=true duration=68.841 method=POST req_bytes=21 res_bytes=194 status_code=201
# 1:41PM WRN http://localhost:9200/_search?q=%7BFAIL duration=17.276 method=GET req_bytes=0 res_bytes=765 status_code=400
# 1:41PM INF http://localhost:9200/test/_search?size=1 duration=9.429 method=GET req_bytes=49 res_bytes=280 status_code=200
```
