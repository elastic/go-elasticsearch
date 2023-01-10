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


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Returns low-level information about REST actions usage on nodes.
package usage

import (
	gobytes "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nodeidMask = iota + 1

	metricMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Usage struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	nodeid string
	metric string
}

// NewUsage type alias for index.
type NewUsage func() *Usage

// NewUsageFunc returns a new instance of Usage with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUsageFunc(tp elastictransport.Interface) NewUsage {
	return func() *Usage {
		n := New(tp)

		return n
	}
}

// Returns low-level information about REST actions usage on nodes.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/cluster-nodes-usage.html
func New(tp elastictransport.Interface) *Usage {
	r := &Usage{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Usage) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")
		path.WriteString("usage")

		method = http.MethodGet
	case r.paramSet == nodeidMask:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")

		path.WriteString(r.nodeid)
		path.WriteString("/")
		path.WriteString("usage")

		method = http.MethodGet
	case r.paramSet == metricMask:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")
		path.WriteString("usage")
		path.WriteString("/")

		path.WriteString(r.metric)

		method = http.MethodGet
	case r.paramSet == nodeidMask|metricMask:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")

		path.WriteString(r.nodeid)
		path.WriteString("/")
		path.WriteString("usage")
		path.WriteString("/")

		path.WriteString(r.metric)

		method = http.MethodGet
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Do runs the http.Request through the provided transport.
func (r Usage) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Usage query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r Usage) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Do(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	return false, nil
}

// Header set a key, value pair in the Usage headers map.
func (r *Usage) Header(key, value string) *Usage {
	r.headers.Set(key, value)

	return r
}

// NodeId A comma-separated list of node IDs or names to limit the returned
// information; use `_local` to return information from the node you're
// connecting to, leave empty to get information from all nodes
// API Name: nodeid
func (r *Usage) NodeId(v string) *Usage {
	r.paramSet |= nodeidMask
	r.nodeid = v

	return r
}

// Metric Limit the information returned to the specified metrics
// API Name: metric
func (r *Usage) Metric(v string) *Usage {
	r.paramSet |= metricMask
	r.metric = v

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *Usage) Timeout(value string) *Usage {
	r.values.Set("timeout", value)

	return r
}
