/*
 * Licensed to Elasticsearch under one or more contributor
 * license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. Elasticsearch licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package client

import (
	"net/url"
	"strings"

	"github.com/elastic/go-elasticsearch/transport"
)

// Option is a client-wide option.
type Option func(c *Client) error

// WithHost connects to a single host. Supports RFC-1738 formatted URLs, valid usages include:
//	WithHost("localhost")
//	WithHost("localhost:9200")
//	WithHost("http://localhost:9200")
func WithHost(host string) Option {
	return func(c *Client) error {
		c.transport.URL = transport.DefaultURL
		u, err := url.Parse(host)
		if err != nil {
			return err
		}
		if u.Scheme != "" {
			c.transport.URL.Scheme = u.Scheme
		}
		c.transport.URL.Host = u.Host
		if !strings.Contains(c.transport.URL.Host, ":") {
			c.transport.URL.Host += ":" + string(transport.DefaultPort)
		}
		return nil
	}
}

/*
// WithHosts connects to the specified hosts. As WithHost, it supports RFC-1738 formatted URLs.
func WithHosts(hosts []string) Option {
	return func(c *Client) error {
		for _, host := range hosts {
			// TODO: use additional hosts
			if err := WithHost(host)(c); err != nil {
				return err
			}
		}
		return nil
	}
}

// WithCaFile sets the path to server TLS certificate file.
func WithCaFile(caFile string) Option {
	return func(c *Client) error {
		// TODO: implement
		return nil
	}
}

// WithCertFile sets the path to client TLS certificate file.
func WithCertFile(certFile string) Option {
	return func(c *Client) error {
		// TODO: implement,
		return nil
	}
}

// WithKeyFile sets the path to client TLS private key file.
func WithKeyFile(keyFile string) Option {
	return func(c *Client) error {
		// TODO: implement,
		return nil
	}
}

// WithSniffOnStart indicates whether to obtain a list of nodes from the cluser at startup time.
func WithSniffOnStart(sniffOnStart bool) Option {
	return func(c *Client) error {
		// TODO: implement,
		return nil
	}
}
*/
