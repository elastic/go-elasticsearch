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

// Returns the current global checkpoints for an index. This API is design for
// internal use by the fleet server project.
package globalcheckpoints

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

type GlobalCheckpoints struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	index string
}

// NewGlobalCheckpoints type alias for index.
type NewGlobalCheckpoints func(index string) *GlobalCheckpoints

// NewGlobalCheckpointsFunc returns a new instance of GlobalCheckpoints with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGlobalCheckpointsFunc(tp elastictransport.Interface) NewGlobalCheckpoints {
	return func(index string) *GlobalCheckpoints {
		n := New(tp)

		n.Index(index)

		return n
	}
}

// Returns the current global checkpoints for an index. This API is design for
// internal use by the fleet server project.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-global-checkpoints.html
func New(tp elastictransport.Interface) *GlobalCheckpoints {
	r := &GlobalCheckpoints{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GlobalCheckpoints) HttpRequest(ctx context.Context) (*http.Request, error) {
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
		path.WriteString("_fleet")
		path.WriteString("/")
		path.WriteString("global_checkpoints")

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

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r GlobalCheckpoints) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GlobalCheckpoints query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a globalcheckpoints.Response
func (r GlobalCheckpoints) Do(ctx context.Context) (*Response, error) {

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
func (r GlobalCheckpoints) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the GlobalCheckpoints headers map.
func (r *GlobalCheckpoints) Header(key, value string) *GlobalCheckpoints {
	r.headers.Set(key, value)

	return r
}

// Index A single index or index alias that resolves to a single index.
// API Name: index
func (r *GlobalCheckpoints) Index(v string) *GlobalCheckpoints {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// WaitForAdvance A boolean value which controls whether to wait (until the timeout) for the
// global checkpoints
// to advance past the provided `checkpoints`.
// API name: wait_for_advance
func (r *GlobalCheckpoints) WaitForAdvance(b bool) *GlobalCheckpoints {
	r.values.Set("wait_for_advance", strconv.FormatBool(b))

	return r
}

// WaitForIndex A boolean value which controls whether to wait (until the timeout) for the
// target index to exist
// and all primary shards be active. Can only be true when `wait_for_advance` is
// true.
// API name: wait_for_index
func (r *GlobalCheckpoints) WaitForIndex(b bool) *GlobalCheckpoints {
	r.values.Set("wait_for_index", strconv.FormatBool(b))

	return r
}

// Checkpoints A comma separated list of previous global checkpoints. When used in
// combination with `wait_for_advance`,
// the API will only return once the global checkpoints advances past the
// checkpoints. Providing an empty list
// will cause Elasticsearch to immediately return the current global
// checkpoints.
// API name: checkpoints
func (r *GlobalCheckpoints) Checkpoints(v string) *GlobalCheckpoints {
	r.values.Set("checkpoints", v)

	return r
}

// Timeout Period to wait for a global checkpoints to advance past `checkpoints`.
// API name: timeout
func (r *GlobalCheckpoints) Timeout(v string) *GlobalCheckpoints {
	r.values.Set("timeout", v)

	return r
}
