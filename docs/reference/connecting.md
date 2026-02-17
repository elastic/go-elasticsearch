---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/connecting.html
---

# Connecting [connecting]

This page contains the information you need to connect and use the Client with {{es}}.

## Connecting to Elastic Cloud [connecting-to-elastic-cloud]

If you are using [Elastic Cloud](https://www.elastic.co/cloud), the client offers an easy way to connect to it. You must pass the Cloud ID that you can find in the cloud console and the corresponding API key.

```go
cfg := elasticsearch.Config{
        CloudID: "CLOUD_ID",
        APIKey: "API_KEY"
}
es, err := elasticsearch.NewClient(cfg)
```

::::{important}
You need to copy and store the `API key` in a secure place since you will not be able to view it again in Elastic Cloud.
::::

## Connecting to a self-managed cluster [connecting-to-self-managed]

Starting from version 8.0, {{es}} offers security by default with authentication and TLS enabled.

To connect to the {{es}} cluster you need to configure the client to use the generated CA certificate. If you're just getting started with {{es}} we recommend reading the documentation on configuring and starting {{es}} to ensure your cluster is running as expected.

When you start {{es}} for the first time you'll see a distinct block like the one below in the output from {{es}} (you may have to scroll up if it's been a while):

```sh
----------------------------------------------------------------
-> Elasticsearch security features have been automatically configured!
-> Authentication is enabled and cluster connections are encrypted.
->  Password for the elastic user (reset with `bin/elasticsearch-reset-password -u elastic`):
  lhQpLELkjkrawaBoaz0Q
->  HTTP CA certificate SHA-256 fingerprint:
  a52dd93511e8c6045e21f16654b77c9ee0f34aea26d9f40320b531c474676228
...
----------------------------------------------------------------
```

Note down the `elastic` user password and HTTP CA fingerprint for the next sections. In the examples below they will be stored in the variables `ELASTIC_PASSWORD` and `CERT_FINGERPRINT` respectively.

Depending on the circumstances there are two options for verifying the HTTPS connection, either verifying with the CA certificate itself or via the HTTP CA certificate fingerprint.

## Verifying HTTPS with CA certificates [verifying-with-ca]

The generated root CA certificate can be found in the `certs` directory in your {{es}} config location (`$ES_CONF_PATH/certs/http_ca.crt`). If you're running {{es}} in Docker there is [additional documentation for retrieving the CA certificate](docs-content://deploy-manage/deploy/self-managed/install-elasticsearch-with-docker.md).

Once you have the `http_ca.crt` file somewhere accessible pass the content of the file to the client via `CACert`:

```go
cert, _ := os.ReadFile("/path/to/http_ca.crt")

cfg := elasticsearch.Config{
        Addresses: []string{
            "https://localhost:9200",
        },
        Username: "elastic",
        Password: ELASTIC_PASSWORD,
        CACert:   cert,
}
es, err := elasticsearch.NewClient(cfg)
```

## Verifying HTTPS with certificate fingerprint [verifying-with-fingerprint]

This method of verifying the HTTPS connection takes advantage of the certificate fingerprint value noted down earlier. Take this SHA256 fingerprint value and pass it to the Go {{es}} client via `CertificateFingerprint`:

```go
cfg := elasticsearch.Config{
        Addresses: []string{
            "https://localhost:9200",
        },
        Username: "elastic",
        Password: ELASTIC_PASSWORD,
        CertificateFingerprint: CERT_FINGERPRINT,
}
es, err := elasticsearch.NewClient(cfg)
```

The certificate fingerprint can be calculated using OpenSSL x509 with the certificate file:

```sh
openssl x509 -fingerprint -sha256 -noout -in /path/to/http_ca.crt
```

If you don't have access to the generated CA file from {{es}} you can use the following script to output the root CA fingerprint of the {{es}} instance with `openssl s_client`:

```sh
# Replace the values of 'localhost' and '9200' to the
# corresponding host and port values for the cluster.
openssl s_client -connect localhost:9200 -servername localhost -showcerts </dev/null 2>/dev/null \
  | openssl x509 -fingerprint -sha256 -noout -in /dev/stdin
```

The output of `openssl x509` will look something like this:

```sh
SHA256 Fingerprint=A5:2D:D9:35:11:E8:C6:04:5E:21:F1:66:54:B7:7C:9E:E0:F3:4A:EA:26:D9:F4:03:20:B5:31:C4:74:67:62:28
```

## Connecting without security enabled [connecting-without-security]

::::{warning}
Running {{es}} without security enabled is not recommended.
::::

If your cluster is configured with [security explicitly disabled](elasticsearch://reference/elasticsearch/configuration-reference/security-settings.md) then you can connect via HTTP:

```go
cfg := elasticsearch.Config{
        Addresses: []string{
            "http://localhost:9200",
        },
}
es, err := elasticsearch.NewClient(cfg)
```

## Connecting to multiple nodes [connecting-multiple-nodes]

The Go {{es}} client supports sending API requests to multiple nodes in the cluster. This means that work will be more evenly spread across the cluster instead of hammering the same node over and over with requests. To configure the client with multiple nodes you can pass a list of URLs, each URL will be used as a separate node in the pool.

```go
cfg := elasticsearch.Config{
  Addresses: []string{
    "https://localhost:9200",
    "https://localhost:9201",
  },
  CACert:   caCert,
  Username: "elastic",
  Password: ELASTIC_PASSWORD,
}
es, err := elasticsearch.NewClient(cfg)
```

By default nodes are selected using round-robin, but alternate node selection strategies can be implemented via the `elastictransport.Selector` interface and provided to the client configuration.

::::{note}
If your {{es}} cluster is behind a load balancer like when using Elastic Cloud you won't need to configure multiple nodes. Instead use the load balancer host and port.
::::

## Authentication [auth-reference]

This section contains code snippets to show you how to authenticate with {{es}}.

### Basic authentication [auth-basic]

To set the cluster endpoints, the username, and the password programmatically, pass a configuration object to the `elasticsearch.NewClient()` function.

```go
cfg := elasticsearch.Config{
  Addresses: []string{
    "https://localhost:9200",
    "https://localhost:9201",
  },
  Username: "foo",
  Password: "bar",
}
es, err := elasticsearch.NewClient(cfg)
```

You can also include the username and password in the endpoint URL:

```text
'https://username:password@localhost:9200'
```

### HTTP Bearer authentication [auth-token]

HTTP Bearer authentication uses the `ServiceToken` parameter by passing the token as a string. This authentication method is used by [Service Account Tokens](https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-security-create-service-token) and [Bearer Tokens](https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-security-get-token).

```go
cfg := elasticsearch.Config{
    Addresses: []string{
        "https://localhost:9200",
    },
    ServiceToken: "token-value",
}
es, err := elasticsearch.NewClient(cfg)
```

## Compatibility mode [compatibility-mode]

The {{es}} server version 8.0 is introducing a new compatibility mode that allows you a smoother upgrade experience from 7 to 8. In a nutshell, you can use the latest 7.x `go-elasticsearch` Elasticsearch client with an 8.x Elasticsearch server, giving more room to coordinate the upgrade of your codebase to the next major version.

If you want to leverage this functionality, please make sure that you are using the latest 7.x `go-elasticsearch` client and set the environment variable `ELASTIC_CLIENT_APIVERSIONING` to `true` or the configuration option `config.EnableCompatibilityMode` in the client `Config`. The client is handling the rest internally.
For every 8.0 and beyond `go-elasticsearch` client, you're all set! The compatibility mode is enabled by default.

## Using the client [client-usage]

The {{es}} package ties together two separate packages for calling the Elasticsearch APIs and transferring data over HTTP: `esapi` and `estransport`, respectively.

Use the `elasticsearch.NewDefaultClient()` function to create the client with the default settings.

```go
es, err := elasticsearch.NewDefaultClient()
if err != nil {
  log.Fatalf("Error creating the client: %s", err)
}
defer func() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := es.Close(ctx); err != nil {
        log.Fatalf("Error closing the client: %s", err)
    }
} ()

res, err := es.Info()
if err != nil {
  log.Fatalf("Error getting response: %s", err)
}

defer res.Body.Close()
log.Println(res)
```

## Using the client in a function-as-a-service environment [connecting-faas]

This section illustrates the best practices for leveraging the {{es}} client in a Function-as-a-Service (FaaS) environment.
The most influential optimization is to initialize the client outside of the function, the global scope.
This practice does not only improve performance but also enables background functionality as – for example – [sniffing](https://www.elastic.co/blog/elasticsearch-sniffing-best-practices-what-when-why-how).
The following examples provide a skeleton for the best practices.

### GCP Cloud Functions [connecting-faas-gcp]

```go
package httpexample

import (
    "context"
    "net/http"
    "time"
    "log"

    "github.com/elastic/go-elasticsearch/v9"
)

var client *elasticsearch.Client

func init() {
    var err error

    ... # Client configuration
    client, err = elasticsearch.NewClient(cfg)
    if err != nil {
        log.Fatalf("elasticsearch.NewClient: %v", err)
    }
}

func HttpExample(w http.ResponseWriter, r *http.Request) {
    ... # Client usage
}
```

### AWS Lambda [connecting-faas-aws]

```go
package httpexample

import (
    "log"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/elastic/go-elasticsearch/v9"
)

var client *elasticsearch.Client

func init() {
    var err error

    ... # Client configuration
    client, err = elasticsearch.NewClient(cfg)
    if err != nil {
        log.Fatalf("elasticsearch.NewClient: %v", err)
    }
}

func HttpExample() {
    ... # Client usage
}

func main() {
    lambda.Start(HttpExample)
}
```

Resources used to assess these recommendations:

- [GCP Cloud Functions: Tips & Tricks](https://cloud.google.com/functions/docs/bestpractices/tips#use_global_variables_to_reuse_objects_in_future_invocations)
- [Best practices for working with AWS Lambda functions](https://docs.aws.amazon.com/lambda/latest/dg/best-practices.html)
- [AWS Lambda: Comparing the effect of global scope](https://docs.aws.amazon.com/lambda/latest/operatorguide/global-scope.html)
