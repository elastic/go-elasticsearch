package elasticsearch // import "github.com/elastic/go-elasticsearch"

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/estransport"
)

const (
	defaultURL = "http://localhost:9200"
)

// Config represents the client configuration.
//
type Config struct {
	Addresses []string          // A list of Elasticsearch nodes to use.
	Transport http.RoundTripper // The HTTP transport object.
}

// Client represents the Elasticsearch client.
//
type Client struct {
	*esapi.API // Embeds the API methods
	Transport  estransport.Interface
}

// NewDefaultClient creates a new client with default options.
//
// It will use http://localhost:9200 as the default address.
//
// It will use the ELASTICSEARCH_URL environment variable, if set,
// to configure the addresses; use a comma to separate multiple URLs.
//
func NewDefaultClient() (*Client, error) {
	return NewClient(Config{})
}

// NewClient creates a new client with configuration from cfg.
//
// It will use http://localhost:9200 as the default address.
//
// It will use the ELASTICSEARCH_URL environment variable, if set,
// to configure the addresses; use a comma to separate multiple URLs.
//
// It's an error to set both cfg.Addresses and the ELASTICSEARCH_URL
// environment variable.
//
func NewClient(cfg Config) (*Client, error) {
	var tran estransport.Interface

	urls, err := urlsFromEnvironment()
	if err != nil {
		return nil, fmt.Errorf("cannot create client: %s", err)
	}

	if len(urls) > 0 && len(cfg.Addresses) > 0 {
		return nil, fmt.Errorf("cannot create client: both ELASTICSEARCH_URL and Addresses are set")
	}

	for _, addr := range cfg.Addresses {
		u, err := url.Parse(strings.TrimRight(addr, "/"))
		if err != nil {
			return nil, fmt.Errorf("cannot create client: %s", err)
		}

		urls = append(urls, u)
	}

	if len(urls) < 1 {
		u, _ := url.Parse(defaultURL) // errcheck exclude
		urls = append(urls, u)
	}

	tran = estransport.New(estransport.Config{URLs: urls, Transport: cfg.Transport})

	return &Client{Transport: tran, API: esapi.New(tran)}, nil
}

// Perform delegates to Transport to execute a request and return a response.
//
func (c *Client) Perform(req *http.Request) (*http.Response, error) {
	return c.Transport.Perform(req)
}

// urlsFromEnvironment returns a list of urls by splitting
// the ELASTICSEARCH_URL environment variable with comma, or an empty list.
//
func urlsFromEnvironment() ([]*url.URL, error) {
	var addrs []*url.URL

	if envURLs, ok := os.LookupEnv("ELASTICSEARCH_URL"); ok && envURLs != "" {
		list := strings.Split(envURLs, ",")
		for _, u := range list {
			ur, err := url.Parse(strings.TrimRight(strings.TrimSpace(u), "/"))
			if err != nil {
				return nil, err
			}
			addrs = append(addrs, ur)
		}
	}

	return addrs, nil
}
