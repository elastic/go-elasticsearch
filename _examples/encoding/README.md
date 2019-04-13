# Example: encoding

Examples in this folder demonstrate how to use third-party packages for JSON decoding and encoding.

## `tidwall/gjson`

The [`github.com/tidwall/gjson`](https://github.com/tidwall/gjson) package allows an
easy access to JSON properties, without parsing the payload into a data structure,
and is therefore convenient when accessing only selected parts of the response.

``` golang
var json = `{"foo":{"bar":"BAZ"}}`
fmt.Println(gjson.Get(json, "foo.bar"))
// => BAZ
```

The [**`gjson.go`**](./gjson.go) example displays data from the
[_Cluster Stats_](https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-stats.html) API.

```
go run gjson.go
```

![Screenshot](screenshot.png)

## `mailru/easyjson`

The [`github.com/mailru/easyjson`](https://github.com/mailru/easyjson) package uses code generation to
provide fast encoding and decoding of struct types.

The [**`easyjson.go`**](./easyjson.go) example uses custom types for representing simple`Article` documents
and index and search responses, including error responses.

```
make clean setup
go run easyjson.go
```

-----

### Benchmarks

Both `mailru/easyjson` and `tidwall/gjson` provide significant performance gains in different scenarios.

You can run the included benchmarks by executing the `make bench` command; example output below.

```
BenchmarkSearchResults/json           30000	     45679 ns/op	    8456 B/op	      60 allocs/op
BenchmarkSearchResults/easyjson      100000	     15612 ns/op	    7992 B/op	      52 allocs/op
BenchmarkClusterStats/json_-_map      20000	     96808 ns/op	   31971 B/op	     391 allocs/op
BenchmarkClusterStats/json_-_struct   30000	     59046 ns/op	   15048 B/op	      17 allocs/op
BenchmarkClusterStats/gjson         1000000	      2119 ns/op	     264 B/op	       3 allocs/op
```
