# Example: Extending the Client

This example demonstrates how to extend the client, in order to call custom APIs, for example added by a [plugin](https://github.com/elastic/elasticsearch/blob/master/plugins/examples/rest-handler/src/main/java/org/elasticsearch/example/resthandler/ExampleCatAction.java).

The [`main.go`](main.go) example defines a custom type, which embeds the regular Elasticsearch client, and adds a `Custom` namespace with an `Example()` method.

To run the example:

```bash
go run main.go

#    GET http://localhost:9209/_cat/health 200 OK 25ms
#        « 1555252476 14:34:36 go-elasticsearch green 1 1 0 0 0 0 0 0 - 100.0%
#
#    GET http://localhost:9209/_cat/example 200 OK 0s
#        « Hello from Cat Example action
```
