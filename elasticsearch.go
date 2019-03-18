package elasticsearch // import "github.com/elastic/go-elasticsearch"

import (
	"errors"
	"fmt"
	"io"
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
	// A list of Elasticsearch nodes to use.
	Addresses []string

	// The HTTP transport object.
	Transport http.RoundTripper

	// Logging output.
	LogOutput io.Writer
	// Logging format. One of: text (default), curl, json.
	LogFormat estransport.LogFormat
	// Custom logging function.
	LoggerFunc func(*http.Request, *http.Response)
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
	envAddrs := addrsFromEnvironment()

	if len(envAddrs) > 0 && len(cfg.Addresses) > 0 {
		return nil, errors.New("cannot create client: both ELASTICSEARCH_URL and Addresses are set")
	}

	addrs := append(envAddrs, cfg.Addresses...)

	urls, err := addrsToURLs(addrs)
	if err != nil {
		return nil, fmt.Errorf("cannot create client: %s", err)
	}

	if len(urls) == 0 {
		u, _ := url.Parse(defaultURL) // errcheck exclude
		urls = append(urls, u)
	}

	tran := estransport.New(estransport.Config{
		URLs: urls, Transport: cfg.Transport,
		LogOutput: cfg.LogOutput, LogFormat: cfg.LogFormat, LoggerFunc: cfg.LoggerFunc,
	})

	return &Client{Transport: tran, API: esapi.New(tran)}, nil
}

// Perform delegates to Transport to execute a request and return a response.
//
func (c *Client) Perform(req *http.Request) (*http.Response, error) {
	return c.Transport.Perform(req)
}

// addrsFromEnvironment returns a list of addresses by splitting
// the ELASTICSEARCH_URL environment variable with comma, or an empty list.
//
func addrsFromEnvironment() []string {
	var addrs []string

	if envURLs, ok := os.LookupEnv("ELASTICSEARCH_URL"); ok && envURLs != "" {
		list := strings.Split(envURLs, ",")
		for _, u := range list {
			addrs = append(addrs, strings.TrimSpace(u))
		}
	}

	return addrs
}

// addrsToURLs creates a list of url.URL structures from url list.
//
func addrsToURLs(addrs []string) ([]*url.URL, error) {
	var urls []*url.URL
	for _, addr := range addrs {
		u, err := url.Parse(strings.TrimRight(addr, "/"))
		if err != nil {
			return nil, fmt.Errorf("cannot parse url: %v", err)
		}

		urls = append(urls, u)
	}
	return urls, nil
}
