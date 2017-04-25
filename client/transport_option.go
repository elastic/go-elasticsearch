package client

// TransportOption is a transport option for the Client.
type TransportOption func(c *Client)

// WithHosts connects to the specified hosts
func WithHosts(hosts []string) TransportOption {
	return func(c *Client) {
	}
}
