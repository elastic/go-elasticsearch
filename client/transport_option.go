package client

import "github.com/elastic/goelasticsearch/client/transport"

// TransportOption is a transport option.
type TransportOption func(t *transport.Transport)

// WithHosts connects to the specified hosts
func WithHosts(hosts []string) TransportOption {
	return func(t *transport.Transport) {
	}
}
