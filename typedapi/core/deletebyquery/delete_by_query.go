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


// Deletes documents matching the provided query.
package deletebyquery

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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conflicts"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/searchtype"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteByQuery struct {
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

// NewDeleteByQuery type alias for index.
type NewDeleteByQuery func(index string) *DeleteByQuery

// NewDeleteByQueryFunc returns a new instance of DeleteByQuery with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteByQueryFunc(tp elastictransport.Interface) NewDeleteByQuery {
	return func(index string) *DeleteByQuery {
		n := New(tp)

		n.Index(index)

		return n
	}
}

// Deletes documents matching the provided query.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-delete-by-query.html
func New(tp elastictransport.Interface) *DeleteByQuery {
	r := &DeleteByQuery{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *DeleteByQuery) Raw(raw json.RawMessage) *DeleteByQuery {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *DeleteByQuery) Request(req *Request) *DeleteByQuery {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *DeleteByQuery) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for DeleteByQuery: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_delete_by_query")

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
func (r DeleteByQuery) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the DeleteByQuery query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the DeleteByQuery headers map.
func (r *DeleteByQuery) Header(key, value string) *DeleteByQuery {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names to search; use `_all` or empty string
// to perform the operation on all indices
// API Name: index
func (r *DeleteByQuery) Index(v string) *DeleteByQuery {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *DeleteByQuery) AllowNoIndices(b bool) *DeleteByQuery {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// Analyzer The analyzer to use for the query string
// API name: analyzer
func (r *DeleteByQuery) Analyzer(value string) *DeleteByQuery {
	r.values.Set("analyzer", value)

	return r
}

// AnalyzeWildcard Specify whether wildcard and prefix queries should be analyzed (default:
// false)
// API name: analyze_wildcard
func (r *DeleteByQuery) AnalyzeWildcard(b bool) *DeleteByQuery {
	r.values.Set("analyze_wildcard", strconv.FormatBool(b))

	return r
}

// Conflicts What to do when the delete by query hits version conflicts?
// API name: conflicts
func (r *DeleteByQuery) Conflicts(enum conflicts.Conflicts) *DeleteByQuery {
	r.values.Set("conflicts", enum.String())

	return r
}

// DefaultOperator The default operator for query string query (AND or OR)
// API name: default_operator
func (r *DeleteByQuery) DefaultOperator(enum operator.Operator) *DeleteByQuery {
	r.values.Set("default_operator", enum.String())

	return r
}

// Df The field to use as default where no field prefix is given in the query
// string
// API name: df
func (r *DeleteByQuery) Df(value string) *DeleteByQuery {
	r.values.Set("df", value)

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *DeleteByQuery) ExpandWildcards(value string) *DeleteByQuery {
	r.values.Set("expand_wildcards", value)

	return r
}

// From Starting offset (default: 0)
// API name: from
func (r *DeleteByQuery) From(value string) *DeleteByQuery {
	r.values.Set("from", value)

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *DeleteByQuery) IgnoreUnavailable(b bool) *DeleteByQuery {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// Lenient Specify whether format-based query failures (such as providing text to a
// numeric field) should be ignored
// API name: lenient
func (r *DeleteByQuery) Lenient(b bool) *DeleteByQuery {
	r.values.Set("lenient", strconv.FormatBool(b))

	return r
}

// MaxDocs Maximum number of documents to process (default: all documents)
// API name: max_docs
func (r *DeleteByQuery) MaxDocs(value string) *DeleteByQuery {
	r.values.Set("max_docs", value)

	return r
}

// Preference Specify the node or shard the operation should be performed on (default:
// random)
// API name: preference
func (r *DeleteByQuery) Preference(value string) *DeleteByQuery {
	r.values.Set("preference", value)

	return r
}

// Refresh Should the affected indexes be refreshed?
// API name: refresh
func (r *DeleteByQuery) Refresh(b bool) *DeleteByQuery {
	r.values.Set("refresh", strconv.FormatBool(b))

	return r
}

// RequestCache Specify if request cache should be used for this request or not, defaults to
// index level setting
// API name: request_cache
func (r *DeleteByQuery) RequestCache(b bool) *DeleteByQuery {
	r.values.Set("request_cache", strconv.FormatBool(b))

	return r
}

// RequestsPerSecond The throttle for this request in sub-requests per second. -1 means no
// throttle.
// API name: requests_per_second
func (r *DeleteByQuery) RequestsPerSecond(value string) *DeleteByQuery {
	r.values.Set("requests_per_second", value)

	return r
}

// Routing A comma-separated list of specific routing values
// API name: routing
func (r *DeleteByQuery) Routing(value string) *DeleteByQuery {
	r.values.Set("routing", value)

	return r
}

// Q Query in the Lucene query string syntax
// API name: q
func (r *DeleteByQuery) Q(value string) *DeleteByQuery {
	r.values.Set("q", value)

	return r
}

// Scroll Specify how long a consistent view of the index should be maintained for
// scrolled search
// API name: scroll
func (r *DeleteByQuery) Scroll(value string) *DeleteByQuery {
	r.values.Set("scroll", value)

	return r
}

// ScrollSize Size on the scroll request powering the delete by query
// API name: scroll_size
func (r *DeleteByQuery) ScrollSize(value string) *DeleteByQuery {
	r.values.Set("scroll_size", value)

	return r
}

// SearchTimeout Explicit timeout for each search request. Defaults to no timeout.
// API name: search_timeout
func (r *DeleteByQuery) SearchTimeout(value string) *DeleteByQuery {
	r.values.Set("search_timeout", value)

	return r
}

// SearchType Search operation type
// API name: search_type
func (r *DeleteByQuery) SearchType(enum searchtype.SearchType) *DeleteByQuery {
	r.values.Set("search_type", enum.String())

	return r
}

// Slices The number of slices this task should be divided into. Defaults to 1, meaning
// the task isn't sliced into subtasks. Can be set to `auto`.
// API name: slices
func (r *DeleteByQuery) Slices(value string) *DeleteByQuery {
	r.values.Set("slices", value)

	return r
}

// Sort A comma-separated list of <field>:<direction> pairs
// API name: sort
func (r *DeleteByQuery) Sort(value string) *DeleteByQuery {
	r.values.Set("sort", value)

	return r
}

// Stats Specific 'tag' of the request for logging and statistical purposes
// API name: stats
func (r *DeleteByQuery) Stats(value string) *DeleteByQuery {
	r.values.Set("stats", value)

	return r
}

// TerminateAfter The maximum number of documents to collect for each shard, upon reaching
// which the query execution will terminate early.
// API name: terminate_after
func (r *DeleteByQuery) TerminateAfter(value string) *DeleteByQuery {
	r.values.Set("terminate_after", value)

	return r
}

// Timeout Time each individual bulk request should wait for shards that are
// unavailable.
// API name: timeout
func (r *DeleteByQuery) Timeout(value string) *DeleteByQuery {
	r.values.Set("timeout", value)

	return r
}

// Version Specify whether to return document version as part of a hit
// API name: version
func (r *DeleteByQuery) Version(b bool) *DeleteByQuery {
	r.values.Set("version", strconv.FormatBool(b))

	return r
}

// WaitForActiveShards Sets the number of shard copies that must be active before proceeding with
// the delete by query operation. Defaults to 1, meaning the primary shard only.
// Set to `all` for all shard copies, otherwise set to any non-negative value
// less than or equal to the total number of copies for the shard (number of
// replicas + 1)
// API name: wait_for_active_shards
func (r *DeleteByQuery) WaitForActiveShards(value string) *DeleteByQuery {
	r.values.Set("wait_for_active_shards", value)

	return r
}

// WaitForCompletion Should the request should block until the delete by query is complete.
// API name: wait_for_completion
func (r *DeleteByQuery) WaitForCompletion(b bool) *DeleteByQuery {
	r.values.Set("wait_for_completion", strconv.FormatBool(b))

	return r
}
