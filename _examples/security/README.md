# Example: Security

This example demonstrates how to use Transport Layer Security (TLS) for
encrypting and verifying the communication with an Elasticsearch cluster
by passing a custom certificate authority to the client.

## Creating the certificates for the cluster nodes

Generate the certificates using the
[`elasticsearch-certutil`](https://www.elastic.co/guide/en/elasticsearch/reference/current/certutil.html)
tool:

```bash
make certificates
```

> See the [_Encrypting communications in an Elasticsearch Docker Container_](https://www.elastic.co/guide/en/elasticsearch/reference/current/configuring-tls-docker.html) tutorial for a complete overview.

Start the cluster with full security configuration:

```bash
make cluster
```

See the [`elasticsearch-cluster.yml`](elasticsearch-cluster.yml) file for details.

Use `curl` to verify access to the cluster:

```
curl --cacert certificates/ca/ca.crt https://elastic:elastic@localhost:9200
```

> NOTE: On Mac OS X, you may need to add the certificate to the Keychain with `security add-trusted-cert -p ssl certificates/ca/ca.crt`. To remove it, run `security remove-trusted-cert certificates/ca/ca.crt`.


## Using the client configuration option

To pass the certificate authority (CA) to the client, so it can verify the server certificate,
use the `elasticsearch.Config.CACert` configuration option:

```go
// --> Read the certificate from file
cert, _ := ioutil.ReadFile(*cacert)

es, _ := elasticsearch.NewClient(
	elasticsearch.Config{
		// ...

		// --> Pass the certificate to the client
		CACert: cert,
	})
```

Run the full example:

```bash
go run tls_with_ca.go

# [200 OK] {
# ...
```


## Manual transport configuration

To configure the transport manually, use the
`(*http.Transport).TLSClientConfig.RootCAs.AppendCertsFromPEM()` method.

Run the full example:

```bash
go run tls_configure_ca.go

# [200 OK] {
# ...
```
