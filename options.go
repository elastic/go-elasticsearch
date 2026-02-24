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
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

// Option configures an Elasticsearch client. Use the With* functions to obtain
// Option values.
type Option struct {
	apply func(*clientOptions) error
}

// clientOptions accumulates the resolved client-level settings.
type clientOptions struct {
	addresses            []string
	cloudID              string
	discoverNodesOnStart bool
	compatibilityMode    bool
	disableMetaHeader    bool
	transportOptions     []elastictransport.Option
}

// --- Client-level options ---------------------------------------------------

// WithAddresses sets the Elasticsearch node addresses.
// Multiple addresses enable round-robin load balancing and failover.
//
// If neither WithAddresses nor WithCloudID is used, the client falls back to
// the ELASTICSEARCH_URL environment variable, then to http://localhost:9200.
func WithAddresses(addrs ...string) Option {
	return Option{apply: func(o *clientOptions) error {
		o.addresses = append(o.addresses, addrs...)
		return nil
	}}
}

// WithCloudID configures the client for Elastic Cloud using the given Cloud ID.
// See https://www.elastic.co/guide/en/cloud/current/ec-cloud-id.html.
//
// It is an error to set both WithCloudID and WithAddresses.
func WithCloudID(cloudID string) Option {
	return Option{apply: func(o *clientOptions) error {
		o.cloudID = cloudID
		return nil
	}}
}

// WithDiscoverNodesOnStart triggers node discovery when the client is created.
func WithDiscoverNodesOnStart() Option {
	return Option{apply: func(o *clientOptions) error {
		o.discoverNodesOnStart = true
		return nil
	}}
}

// WithCompatibilityMode enables the compatibility header on every request.
func WithCompatibilityMode() Option {
	return Option{apply: func(o *clientOptions) error {
		o.compatibilityMode = true
		return nil
	}}
}

// WithDisableMetaHeader disables the X-Elastic-Client-Meta header sent with
// every request.
func WithDisableMetaHeader() Option {
	return Option{apply: func(o *clientOptions) error {
		o.disableMetaHeader = true
		return nil
	}}
}

// --- Convenience transport options ------------------------------------------
// These delegate to elastictransport options so callers don't need to import
// the transport package for common settings.

// WithAPIKey configures API Key authentication. The key should be the
// base64-encoded value returned by the Elasticsearch create API key endpoint.
func WithAPIKey(apiKey string) Option {
	return withTransportOption(elastictransport.WithAPIKey(apiKey))
}

// WithBasicAuth configures HTTP Basic Authentication with the given username
// and password.
func WithBasicAuth(username, password string) Option {
	return withTransportOption(elastictransport.WithBasicAuth(username, password))
}

// WithServiceToken configures bearer-token authentication using the given
// service token.
func WithServiceToken(token string) Option {
	return withTransportOption(elastictransport.WithServiceToken(token))
}

// WithCACert sets PEM-encoded CA certificates used to verify the server's TLS
// certificate.
func WithCACert(cert []byte) Option {
	return withTransportOption(elastictransport.WithCACert(cert))
}

// WithCertificateFingerprint configures SHA-256 certificate fingerprint
// verification. The fingerprint should be a hex-encoded string.
func WithCertificateFingerprint(fingerprint string) Option {
	return withTransportOption(elastictransport.WithCertificateFingerprint(fingerprint))
}

// WithRetry configures retry behaviour with the given maximum number of retries
// and optional HTTP status codes that trigger a retry. If no status codes are
// provided the defaults (502, 503, 504) are used.
func WithRetry(maxRetries int, onStatus ...int) Option {
	return withTransportOption(elastictransport.WithRetry(maxRetries, onStatus...))
}

// WithCompression enables gzip compression for request bodies. An optional
// compression level may be provided (see compress/gzip constants). When omitted,
// gzip.DefaultCompression is used.
func WithCompression(level ...int) Option {
	return withTransportOption(elastictransport.WithCompression(level...))
}

// WithInstrumentation sets the Instrumentation used for tracing and metrics
// propagation (e.g. OpenTelemetry).
func WithInstrumentation(i elastictransport.Instrumentation) Option {
	return withTransportOption(elastictransport.WithInstrumentation(i))
}

// WithLogger sets the Logger used to log request and response information.
func WithLogger(l elastictransport.Logger) Option {
	return withTransportOption(elastictransport.WithLogger(l))
}

// --- Transport passthrough --------------------------------------------------

// WithTransportOptions appends raw transport options for settings not covered
// by the convenience helpers above. Options are applied in order after all
// client-level options have been resolved.
func WithTransportOptions(opts ...elastictransport.Option) Option {
	return Option{apply: func(o *clientOptions) error {
		o.transportOptions = append(o.transportOptions, opts...)
		return nil
	}}
}

// --- internal helpers -------------------------------------------------------

func withTransportOption(to elastictransport.Option) Option {
	return Option{apply: func(o *clientOptions) error {
		o.transportOptions = append(o.transportOptions, to)
		return nil
	}}
}

// resolvedOptions holds the resolved settings produced by applying client
// options. It is used by the public constructors to build the final client
// without copying a sync.RWMutex.
type resolvedOptions struct {
	transport            *elastictransport.Client
	metaHeader           string
	disableMetaHeader    bool
	compatibilityHeader  bool
	discoverNodesOnStart bool
}

// resolveOptions applies opts, resolves addresses, creates the transport via
// elastictransport.NewClient, and returns the resolved settings.
//
// metaHeaderSuffix is appended to the meta header (e.g. "hl=1" for TypedClient).
func resolveOptions(opts []Option, metaHeaderSuffix string) (*resolvedOptions, error) {
	var co clientOptions
	var errs []error
	for _, opt := range opts {
		if err := opt.apply(&co); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return nil, fmt.Errorf("client options: %w", errors.Join(errs...))
	}

	if len(co.addresses) > 0 && co.cloudID != "" {
		return nil, errors.New("cannot create client: both Addresses and CloudID are set")
	}

	var addrs []string
	switch {
	case co.cloudID != "":
		cloudAddr, err := addrFromCloudID(co.cloudID)
		if err != nil {
			return nil, fmt.Errorf("cannot create client: cannot parse CloudID: %s", err)
		}
		addrs = append(addrs, cloudAddr)
	case len(co.addresses) > 0:
		addrs = append(addrs, co.addresses...)
	default:
		addrs = addrsFromEnvironment()
	}

	urls, err := addrsToURLs(addrs)
	if err != nil {
		return nil, fmt.Errorf("cannot create client: %s", err)
	}

	if len(urls) == 0 {
		u, parseErr := parseDefaultURL()
		if parseErr != nil {
			return nil, parseErr
		}
		urls = append(urls, u)
	}

	tpOpts := []elastictransport.Option{
		elastictransport.WithUserAgent(userAgent),
		elastictransport.WithURLs(urls...),
	}

	if urls[0].User != nil {
		username := urls[0].User.Username()
		password, _ := urls[0].User.Password()
		tpOpts = append(tpOpts, elastictransport.WithBasicAuth(username, password))
	}

	tpOpts = append(tpOpts, co.transportOptions...)

	tp, err := elastictransport.NewClient(tpOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating transport: %s", err)
	}

	compatHeaderEnv := os.Getenv(esCompatHeader)
	envCompat, _ := strconv.ParseBool(compatHeaderEnv)

	meta := initMetaHeader(tp)
	if metaHeaderSuffix != "" {
		meta = strings.Join([]string{meta, metaHeaderSuffix}, ",")
	}

	return &resolvedOptions{
		transport:            tp,
		metaHeader:           meta,
		disableMetaHeader:    co.disableMetaHeader,
		compatibilityHeader:  co.compatibilityMode || envCompat,
		discoverNodesOnStart: co.discoverNodesOnStart,
	}, nil
}
