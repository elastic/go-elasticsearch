package client

// TransportOption is a transport option.
type TransportOption func(c *Client)

// WithHost connects to a single host. Supports RFC-1738 formatted URLs, valid usages include:
//	WithHost("localhost")
//	WithHost("localhost:9200")
//	WithHost("http://localhost:9200")
func WithHost(host string) TransportOption {
	return func(c *Client) {
		// TODO: implement
	}
}

// WithHosts connects to the specified hosts. As WithHost, it supports RFC-1738 formatted URLs.
func WithHosts(hosts []string) TransportOption {
	return func(c *Client) {
		for _, host := range hosts {
			WithHost(host)(c)
		}
	}
}

// WithCaFile sets the path to server TLS certificate file.
func WithCaFile(caFile string) TransportOption {
	return func(c *Client) {
		// TODO: implement
	}
}

// WithCertFile sets the path to client TLS certificate file.
func WithCertFile(certFile string) TransportOption {
	return func(c *Client) {
		// TODO: implement
	}
}

// WithKeyFile sets the path to client TLS private key file.
func WithKeyFile(keyFile string) TransportOption {
	return func(c *Client) {
		// TODO: implement
	}
}

// WithSniffOnStart indicates whether to obtain a list of nodes from the cluser at startup time.
func WithSniffOnStart(sniffOnStart bool) TransportOption {
	return func(c *Client) {
		// TODO: implement
	}
}
