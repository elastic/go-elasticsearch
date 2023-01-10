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


// Returns number of documents matching a query.
package count

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Count struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	index string
}

// NewCount type alias for index.
type NewCount func() *Count

// NewCountFunc returns a new instance of Count with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCountFunc(tp elastictransport.Interface) NewCount {
	return func() *Count {
		n := New(tp)

		return n
	}
}

// Returns number of documents matching a query.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/search-count.html
func New(tp elastictransport.Interface) *Count {
	r := &Count{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Count) Raw(raw json.RawMessage) *Count {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Count) Request(req *Request) *Count {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Count) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Count: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_count")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_count")

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

// Do runs the http.Request through the provided transport.
func (r Count) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Count query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the Count headers map.
func (r *Count) Header(key, value string) *Count {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of indices to restrict the results
// API Name: index
func (r *Count) Index(v string) *Count {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *Count) AllowNoIndices(b bool) *Count {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// Analyzer The analyzer to use for the query string
// API name: analyzer
func (r *Count) Analyzer(value string) *Count {
	r.values.Set("analyzer", value)

	return r
}

// AnalyzeWildcard Specify whether wildcard and prefix queries should be analyzed (default:
// false)
// API name: analyze_wildcard
func (r *Count) AnalyzeWildcard(b bool) *Count {
	r.values.Set("analyze_wildcard", strconv.FormatBool(b))

	return r
}

// DefaultOperator The default operator for query string query (AND or OR)
// API name: default_operator
func (r *Count) DefaultOperator(enum operator.Operator) *Count {
	r.values.Set("default_operator", enum.String())

	return r
}

// Df The field to use as default where no field prefix is given in the query
// string
// API name: df
func (r *Count) Df(value string) *Count {
	r.values.Set("df", value)

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *Count) ExpandWildcards(value string) *Count {
	r.values.Set("expand_wildcards", value)

	return r
}

// IgnoreThrottled Whether specified concrete, expanded or aliased indices should be ignored
// when throttled
// API name: ignore_throttled
func (r *Count) IgnoreThrottled(b bool) *Count {
	r.values.Set("ignore_throttled", strconv.FormatBool(b))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *Count) IgnoreUnavailable(b bool) *Count {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// Lenient Specify whether format-based query failures (such as providing text to a
// numeric field) should be ignored
// API name: lenient
func (r *Count) Lenient(b bool) *Count {
	r.values.Set("lenient", strconv.FormatBool(b))

	return r
}

// MinScore Include only documents with a specific `_score` value in the result
// API name: min_score
func (r *Count) MinScore(value string) *Count {
	r.values.Set("min_score", value)

	return r
}

// Preference Specify the node or shard the operation should be performed on (default:
// random)
// API name: preference
func (r *Count) Preference(value string) *Count {
	r.values.Set("preference", value)

	return r
}

// Routing A comma-separated list of specific routing values
// API name: routing
func (r *Count) Routing(value string) *Count {
	r.values.Set("routing", value)

	return r
}

// TerminateAfter The maximum count for each shard, upon reaching which the query execution
// will terminate early
// API name: terminate_after
func (r *Count) TerminateAfter(value string) *Count {
	r.values.Set("terminate_after", value)

	return r
}

// Q Query in the Lucene query string syntax
// API name: q
func (r *Count) Q(value string) *Count {
	r.values.Set("q", value)

	return r
}
