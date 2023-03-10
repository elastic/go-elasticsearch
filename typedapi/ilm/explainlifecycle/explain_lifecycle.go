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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

// Retrieves information about the index's current lifecycle state, such as the
// currently executing phase, action, and step.
package explainlifecycle

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ExplainLifecycle struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	index string
}

// NewExplainLifecycle type alias for index.
type NewExplainLifecycle func(index string) *ExplainLifecycle

// NewExplainLifecycleFunc returns a new instance of ExplainLifecycle with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewExplainLifecycleFunc(tp elastictransport.Interface) NewExplainLifecycle {
	return func(index string) *ExplainLifecycle {
		n := New(tp)

		n.Index(index)

		return n
	}
}

// Retrieves information about the index's current lifecycle state, such as the
// currently executing phase, action, and step.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-explain-lifecycle.html
func New(tp elastictransport.Interface) *ExplainLifecycle {
	r := &ExplainLifecycle{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ExplainLifecycle) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_ilm")
		path.WriteString("/")
		path.WriteString("explain")

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

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r ExplainLifecycle) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the ExplainLifecycle query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a explainlifecycle.Response
func (r ExplainLifecycle) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil

	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	return nil, errorResponse
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r ExplainLifecycle) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Perform(ctx)

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

// Header set a key, value pair in the ExplainLifecycle headers map.
func (r *ExplainLifecycle) Header(key, value string) *ExplainLifecycle {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and aliases to target.
// Supports wildcards (`*`).
// To target all data streams and indices, use `*` or `_all`.
// API Name: index
func (r *ExplainLifecycle) Index(v string) *ExplainLifecycle {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// OnlyErrors Filters the returned indices to only indices that are managed by ILM and are
// in an error state, either due to an encountering an error while executing the
// policy, or attempting to use a policy that does not exist.
// API name: only_errors
func (r *ExplainLifecycle) OnlyErrors(b bool) *ExplainLifecycle {
	r.values.Set("only_errors", strconv.FormatBool(b))

	return r
}

// OnlyManaged Filters the returned indices to only indices that are managed by ILM.
// API name: only_managed
func (r *ExplainLifecycle) OnlyManaged(b bool) *ExplainLifecycle {
	r.values.Set("only_managed", strconv.FormatBool(b))

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *ExplainLifecycle) MasterTimeout(v string) *ExplainLifecycle {
	r.values.Set("master_timeout", v)

	return r
}

// Timeout Period to wait for a response. If no response is received before the timeout
// expires, the request fails and returns an error.
// API name: timeout
func (r *ExplainLifecycle) Timeout(v string) *ExplainLifecycle {
	r.values.Set("timeout", v)

	return r
}
