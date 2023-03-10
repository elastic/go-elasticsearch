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

// Allows to copy documents from one index to another, optionally filtering the
// source
// documents by a query, changing the destination index settings, or fetching
// the
// documents from a remote cluster.
package reindex

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Reindex struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw io.Reader

	paramSet int
}

// NewReindex type alias for index.
type NewReindex func() *Reindex

// NewReindexFunc returns a new instance of Reindex with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewReindexFunc(tp elastictransport.Interface) NewReindex {
	return func() *Reindex {
		n := New(tp)

		return n
	}
}

// Allows to copy documents from one index to another, optionally filtering the
// source
// documents by a query, changing the destination index settings, or fetching
// the
// documents from a remote cluster.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-reindex.html
func New(tp elastictransport.Interface) *Reindex {
	r := &Reindex{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Reindex) Raw(raw io.Reader) *Reindex {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Reindex) Request(req *Request) *Reindex {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Reindex) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Reindex: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_reindex")

		method = http.MethodPost
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
func (r Reindex) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Reindex query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a reindex.Response
func (r Reindex) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the Reindex headers map.
func (r *Reindex) Header(key, value string) *Reindex {
	r.headers.Set(key, value)

	return r
}

// Refresh Should the affected indexes be refreshed?
// API name: refresh
func (r *Reindex) Refresh(b bool) *Reindex {
	r.values.Set("refresh", strconv.FormatBool(b))

	return r
}

// RequestsPerSecond The throttle to set on this request in sub-requests per second. -1 means no
// throttle.
// API name: requests_per_second
func (r *Reindex) RequestsPerSecond(v string) *Reindex {
	r.values.Set("requests_per_second", v)

	return r
}

// Scroll Control how long to keep the search context alive
// API name: scroll
func (r *Reindex) Scroll(v string) *Reindex {
	r.values.Set("scroll", v)

	return r
}

// Slices The number of slices this task should be divided into. Defaults to 1, meaning
// the task isn't sliced into subtasks. Can be set to `auto`.
// API name: slices
func (r *Reindex) Slices(v string) *Reindex {
	r.values.Set("slices", v)

	return r
}

// Timeout Time each individual bulk request should wait for shards that are
// unavailable.
// API name: timeout
func (r *Reindex) Timeout(v string) *Reindex {
	r.values.Set("timeout", v)

	return r
}

// WaitForActiveShards Sets the number of shard copies that must be active before proceeding with
// the reindex operation. Defaults to 1, meaning the primary shard only. Set to
// `all` for all shard copies, otherwise set to any non-negative value less than
// or equal to the total number of copies for the shard (number of replicas + 1)
// API name: wait_for_active_shards
func (r *Reindex) WaitForActiveShards(v string) *Reindex {
	r.values.Set("wait_for_active_shards", v)

	return r
}

// WaitForCompletion Should the request should block until the reindex is complete.
// API name: wait_for_completion
func (r *Reindex) WaitForCompletion(b bool) *Reindex {
	r.values.Set("wait_for_completion", strconv.FormatBool(b))

	return r
}

// API name: require_alias
func (r *Reindex) RequireAlias(b bool) *Reindex {
	r.values.Set("require_alias", strconv.FormatBool(b))

	return r
}
