// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package elasticsearch

import (
	"encoding/base64"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/v9/typedapi"

	"github.com/elastic/go-elasticsearch/v9/esapi"
	"github.com/elastic/go-elasticsearch/v9/internal/version"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	tpversion "github.com/elastic/elastic-transport-go/v8/elastictransport/version"
)

const (
	defaultURL = "http://localhost:9200"

	// Version returns the package version as a string.
	Version        = version.Client
	unknownProduct = "the client noticed that the server is not Elasticsearch and we do not support this unknown product"

	// HeaderClientMeta Key for the HTTP Header related to telemetry data sent with
	// each request to Elasticsearch.
	HeaderClientMeta = "x-elastic-client-meta"

	compatibilityHeader = "application/vnd.elasticsearch+json;compatible-with=9"
)

var (
	esCompatHeader = "ELASTIC_CLIENT_APIVERSIONING"
	userAgent      string
	reGoVersion    = regexp.MustCompile(`go(\d+\.\d+\..+)`)
	reMetaVersion  = regexp.MustCompile("([0-9.]+)(.*)")
)

func init() {
	userAgent = initUserAgent()
}

// Config represents the client configuration.
type Config struct {
	Addresses []string // A list of Elasticsearch nodes to use.
	Username  string   // Username for HTTP Basic Authentication.
	Password  string   // Password for HTTP Basic Authentication.

	CloudID                string // Endpoint for the Elastic Service (https://elastic.co/cloud).
	APIKey                 string // Base64-encoded token for authorization; if set, overrides username/password and service token.
	ServiceToken           string // Service token for authorization; if set, overrides username/password.
	CertificateFingerprint string // SHA256 hex fingerprint given by Elasticsearch on first launch.

	Header http.Header // Global HTTP request header.

	// PEM-encoded certificate authorities.
	// When set, an empty certificate pool will be created, and the certificates will be appended to it.
	// The option is only valid when the transport is not specified, or when it's http.Transport.
	CACert []byte

	RetryOnStatus []int                           // List of status codes for retry. Default: 502, 503, 504.
	DisableRetry  bool                            // Default: false.
	MaxRetries    int                             // Default: 3.
	RetryOnError  func(*http.Request, error) bool // Optional function allowing to indicate which error should be retried. Default: nil.

	CompressRequestBody      bool // Default: false.
	CompressRequestBodyLevel int  // Default: gzip.DefaultCompression.
	PoolCompressor           bool // If true, a sync.Pool based gzip writer is used. Default: false.

	DiscoverNodesOnStart  bool          // Discover nodes when initializing the client. Default: false.
	DiscoverNodesInterval time.Duration // Discover nodes periodically. Default: disabled.

	EnableMetrics           bool // Enable the metrics collection.
	EnableDebugLogger       bool // Enable the debug logging.
	EnableCompatibilityMode bool // Enable sends compatibility header

	DisableMetaHeader bool // Disable the additional "X-Elastic-Client-Meta" HTTP header.

	RetryBackoff func(attempt int) time.Duration // Optional backoff duration. Default: nil.

	Transport http.RoundTripper         // The HTTP transport object.
	Logger    elastictransport.Logger   // The logger object.
	Selector  elastictransport.Selector // The selector object.

	// Optional constructor function for a custom ConnectionPool. Default: nil.
	ConnectionPoolFunc func([]*elastictransport.Connection, elastictransport.Selector) elastictransport.ConnectionPool

	Instrumentation elastictransport.Instrumentation // Enable instrumentation throughout the client.
}

// NewOpenTelemetryInstrumentation provides the OpenTelemetry integration for both low-level and TypedAPI.
// provider is optional, if nil is passed the integration will retrieve the provider set globally by otel.
// captureSearchBody allows to define if the search queries body should be included in the span.
// Search endpoints are:
//
//	search
//	async_search.submit
//	msearch
//	eql.search
//	terms_enum
//	search_template
//	msearch_template
//	render_search_template
func NewOpenTelemetryInstrumentation(provider trace.TracerProvider, captureSearchBody bool) elastictransport.Instrumentation {
	return elastictransport.NewOtelInstrumentation(provider, captureSearchBody, Version)
}

// BaseClient represents the Elasticsearch client.
type BaseClient struct {
	Transport           elastictransport.Interface
	metaHeader          string
	compatibilityHeader bool

	disableMetaHeader   bool
	productCheckMu      sync.RWMutex
	productCheckSuccess bool
}

// Client represents the Functional Options API.
type Client struct {
	BaseClient
	*esapi.API
}

// TypedClient represents the Typed API.
type TypedClient struct {
	BaseClient
	*typedapi.MethodAPI
}

// NewBaseClient creates a new client free of any API.
func NewBaseClient(cfg Config) (*BaseClient, error) {
	tp, err := newTransport(cfg)
	if err != nil {
		return nil, err
	}

	compatHeaderEnv := os.Getenv(esCompatHeader)
	compatibilityHeader, _ := strconv.ParseBool(compatHeaderEnv)

	client := &BaseClient{
		Transport:           tp,
		disableMetaHeader:   cfg.DisableMetaHeader,
		metaHeader:          initMetaHeader(tp),
		compatibilityHeader: cfg.EnableCompatibilityMode || compatibilityHeader,
	}

	if cfg.DiscoverNodesOnStart {
		go client.DiscoverNodes()
	}

	return client, nil
}

// NewDefaultClient creates a new client with default options.
//
// It will use http://localhost:9200 as the default address.
//
// It will use the ELASTICSEARCH_URL environment variable, if set,
// to configure the addresses; use a comma to separate multiple URLs.
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
// If either cfg.Addresses or cfg.CloudID is set, the ELASTICSEARCH_URL
// environment variable is ignored.
//
// It's an error to set both cfg.Addresses and cfg.CloudID.
func NewClient(cfg Config) (*Client, error) {
	tp, err := newTransport(cfg)
	if err != nil {
		return nil, err
	}

	compatHeaderEnv := os.Getenv(esCompatHeader)
	compatibilityHeader, _ := strconv.ParseBool(compatHeaderEnv)

	client := &Client{
		BaseClient: BaseClient{
			Transport:           tp,
			disableMetaHeader:   cfg.DisableMetaHeader,
			metaHeader:          initMetaHeader(tp),
			compatibilityHeader: cfg.EnableCompatibilityMode || compatibilityHeader,
		},
	}
	client.API = esapi.New(client)

	if cfg.DiscoverNodesOnStart {
		go client.DiscoverNodes()
	}

	return client, nil
}

// NewTypedClient create a new client with the configuration from cfg.
//
// This version uses the same configuration as NewClient.
//
// It will return the client with the TypedAPI.
func NewTypedClient(cfg Config) (*TypedClient, error) {
	tp, err := newTransport(cfg)
	if err != nil {
		return nil, err
	}

	compatHeaderEnv := os.Getenv(esCompatHeader)
	compatibilityHeader, _ := strconv.ParseBool(compatHeaderEnv)

	metaHeader := strings.Join([]string{initMetaHeader(tp), "hl=1"}, ",")

	client := &TypedClient{
		BaseClient: BaseClient{
			Transport:           tp,
			disableMetaHeader:   cfg.DisableMetaHeader,
			metaHeader:          metaHeader,
			compatibilityHeader: cfg.EnableCompatibilityMode || compatibilityHeader,
		},
	}
	client.MethodAPI = typedapi.NewMethodAPI(client)

	if cfg.DiscoverNodesOnStart {
		go client.DiscoverNodes()
	}

	return client, nil
}

func newTransport(cfg Config) (*elastictransport.Client, error) {
	var addrs []string

	if len(cfg.Addresses) == 0 && cfg.CloudID == "" {
		addrs = addrsFromEnvironment()
	} else {
		if len(cfg.Addresses) > 0 && cfg.CloudID != "" {
			return nil, errors.New("cannot create client: both Addresses and CloudID are set")
		}

		if cfg.CloudID != "" {
			cloudAddr, err := addrFromCloudID(cfg.CloudID)
			if err != nil {
				return nil, fmt.Errorf("cannot create client: cannot parse CloudID: %s", err)
			}
			addrs = append(addrs, cloudAddr)
		}

		if len(cfg.Addresses) > 0 {
			addrs = append(addrs, cfg.Addresses...)
		}
	}

	urls, err := addrsToURLs(addrs)
	if err != nil {
		return nil, fmt.Errorf("cannot create client: %s", err)
	}

	if len(urls) == 0 {
		u, _ := url.Parse(defaultURL) // errcheck exclude
		urls = append(urls, u)
	}

	// TODO(karmi): Refactor
	if urls[0].User != nil {
		cfg.Username = urls[0].User.Username()
		pw, _ := urls[0].User.Password()
		cfg.Password = pw
	}

	tpConfig := elastictransport.Config{
		UserAgent: userAgent,

		URLs:                   urls,
		Username:               cfg.Username,
		Password:               cfg.Password,
		APIKey:                 cfg.APIKey,
		ServiceToken:           cfg.ServiceToken,
		CertificateFingerprint: cfg.CertificateFingerprint,

		Header: cfg.Header,
		CACert: cfg.CACert,

		RetryOnStatus: cfg.RetryOnStatus,
		DisableRetry:  cfg.DisableRetry,
		RetryOnError:  cfg.RetryOnError,
		MaxRetries:    cfg.MaxRetries,
		RetryBackoff:  cfg.RetryBackoff,

		CompressRequestBody:      cfg.CompressRequestBody,
		CompressRequestBodyLevel: cfg.CompressRequestBodyLevel,
		PoolCompressor:           cfg.PoolCompressor,

		EnableMetrics:     cfg.EnableMetrics,
		EnableDebugLogger: cfg.EnableDebugLogger,

		DiscoverNodesInterval: cfg.DiscoverNodesInterval,

		Transport:          cfg.Transport,
		Logger:             cfg.Logger,
		Selector:           cfg.Selector,
		ConnectionPoolFunc: cfg.ConnectionPoolFunc,

		Instrumentation: cfg.Instrumentation,
	}

	tp, err := elastictransport.New(tpConfig)
	if err != nil {
		return nil, fmt.Errorf("error creating transport: %s", err)
	}

	return tp, nil
}

// Perform delegates to Transport to execute a request and return a response.
func (c *BaseClient) Perform(req *http.Request) (*http.Response, error) {
	// Compatibility Header
	if c.compatibilityHeader {
		if req.Body != nil {
			req.Header.Set("Content-Type", compatibilityHeader)
		}
		req.Header.Set("Accept", compatibilityHeader)
	}

	if !c.disableMetaHeader {
		existingMetaHeader := req.Header.Get(HeaderClientMeta)
		if existingMetaHeader != "" {
			req.Header.Set(HeaderClientMeta, strings.Join([]string{c.metaHeader, existingMetaHeader}, ","))
		} else {
			req.Header.Add(HeaderClientMeta, c.metaHeader)
		}
	} else {
		req.Header.Del(HeaderClientMeta)
	}

	// Retrieve the original request.
	res, err := c.Transport.Perform(req)
	if err != nil {
		return nil, err
	}

	// ResponseCheck, we run the header check on the first answer from ES.
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		checkHeader := func() error { return genuineCheckHeader(res.Header) }
		if err := c.doProductCheck(checkHeader); err != nil {
			res.Body.Close()
			return nil, err
		}
	}

	return res, nil
}

// InstrumentationEnabled propagates back to the client the Instrumentation provided by the transport.
func (c *BaseClient) InstrumentationEnabled() elastictransport.Instrumentation {
	if tp, ok := c.Transport.(elastictransport.Instrumented); ok {
		return tp.InstrumentationEnabled()
	}
	return nil
}

// doProductCheck calls f if there as not been a prior successful call to doProductCheck,
// returning nil otherwise.
func (c *BaseClient) doProductCheck(f func() error) error {
	c.productCheckMu.RLock()
	productCheckSuccess := c.productCheckSuccess
	c.productCheckMu.RUnlock()

	if productCheckSuccess {
		return nil
	}

	c.productCheckMu.Lock()
	defer c.productCheckMu.Unlock()

	if c.productCheckSuccess {
		return nil
	}

	if err := f(); err != nil {
		return err
	}

	c.productCheckSuccess = true

	return nil
}

// genuineCheckHeader validates the presence of the X-Elastic-Product header
func genuineCheckHeader(header http.Header) error {
	if header.Get("X-Elastic-Product") != "Elasticsearch" {
		return errors.New(unknownProduct)
	}
	return nil
}

// Metrics returns the client metrics.
func (c *BaseClient) Metrics() (elastictransport.Metrics, error) {
	if mt, ok := c.Transport.(elastictransport.Measurable); ok {
		return mt.Metrics()
	}
	return elastictransport.Metrics{}, errors.New("transport is missing method Metrics()")
}

// DiscoverNodes reloads the client connections by fetching information from the cluster.
func (c *BaseClient) DiscoverNodes() error {
	if dt, ok := c.Transport.(elastictransport.Discoverable); ok {
		return dt.DiscoverNodes()
	}
	return errors.New("transport is missing method DiscoverNodes()")
}

// addrsFromEnvironment returns a list of addresses by splitting
// the ELASTICSEARCH_URL environment variable with comma, or an empty list.
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

// addrFromCloudID extracts the Elasticsearch URL from CloudID.
// See: https://www.elastic.co/guide/en/cloud/current/ec-cloud-id.html
func addrFromCloudID(input string) (string, error) {
	var scheme = "https://"

	values := strings.Split(input, ":")
	if len(values) != 2 {
		return "", fmt.Errorf("unexpected format: %q", input)
	}
	data, err := base64.StdEncoding.DecodeString(values[1])
	if err != nil {
		return "", err
	}
	parts := strings.Split(string(data), "$")

	if len(parts) < 2 {
		return "", fmt.Errorf("invalid encoded value: %s", parts)
	}

	return fmt.Sprintf("%s%s.%s", scheme, parts[1], parts[0]), nil
}

func initUserAgent() string {
	var b strings.Builder

	b.WriteString("go-elasticsearch")
	b.WriteRune('/')
	b.WriteString(Version)
	b.WriteRune(' ')
	b.WriteRune('(')
	b.WriteString(runtime.GOOS)
	b.WriteRune(' ')
	b.WriteString(runtime.GOARCH)
	b.WriteString("; ")
	b.WriteString("Go ")
	if v := reGoVersion.ReplaceAllString(runtime.Version(), "$1"); v != "" {
		b.WriteString(v)
	} else {
		b.WriteString(runtime.Version())
	}
	b.WriteRune(')')

	return b.String()
}

func initMetaHeader(transport interface{}) string {
	var b strings.Builder
	var strippedGoVersion string
	var strippedEsVersion string
	var strippedTransportVersion string

	strippedEsVersion = buildStrippedVersion(Version)
	strippedGoVersion = buildStrippedVersion(runtime.Version())

	if _, ok := transport.(*elastictransport.Client); ok {
		strippedTransportVersion = buildStrippedVersion(tpversion.Transport)
	} else {
		strippedTransportVersion = strippedEsVersion
	}

	var duos = [][]string{
		{
			"es",
			strippedEsVersion,
		},
		{
			"go",
			strippedGoVersion,
		},
		{
			"t",
			strippedTransportVersion,
		},
		{
			"hc",
			strippedGoVersion,
		},
	}

	var arr []string
	for _, duo := range duos {
		arr = append(arr, strings.Join(duo, "="))
	}
	b.WriteString(strings.Join(arr, ","))

	return b.String()
}

func buildStrippedVersion(version string) string {
	v := reMetaVersion.FindStringSubmatch(version)

	if len(v) == 3 && !strings.Contains(version, "devel") {
		switch {
		case v[2] != "":
			return v[1] + "p"
		default:
			return v[1]
		}
	}

	return "0.0p"
}
