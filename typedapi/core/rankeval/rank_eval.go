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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Evaluate ranked search results.
//
// Evaluate the quality of ranked search results over a set of typical search
// queries.
package rankeval

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type RankEval struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewRankEval type alias for index.
type NewRankEval func() *RankEval

// NewRankEvalFunc returns a new instance of RankEval with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewRankEvalFunc(tp elastictransport.Interface) NewRankEval {
	return func() *RankEval {
		n := New(tp)

		return n
	}
}

// Evaluate ranked search results.
//
// Evaluate the quality of ranked search results over a set of typical search
// queries.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-rank-eval.html
func New(tp elastictransport.Interface) *RankEval {
	r := &RankEval{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *RankEval) Raw(raw io.Reader) *RankEval {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *RankEval) Request(req *Request) *RankEval {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *RankEval) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for RankEval: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_rank_eval")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_rank_eval")

		method = http.MethodPost
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
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
func (r RankEval) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "rank_eval")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "rank_eval")
		if reader := instrument.RecordRequestBody(ctx, "rank_eval", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "rank_eval")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the RankEval query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a rankeval.Response
func (r RankEval) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "rank_eval")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// Header set a key, value pair in the RankEval headers map.
func (r *RankEval) Header(key, value string) *RankEval {
	r.headers.Set(key, value)

	return r
}

// Index A  comma-separated list of data streams, indices, and index aliases used to
// limit the request.
// Wildcard (`*`) expressions are supported.
// To target all data streams and indices in a cluster, omit this parameter or
// use `_all` or `*`.
// API Name: index
func (r *RankEval) Index(index string) *RankEval {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices If `false`, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices. This behavior
// applies even if the request targets other open indices. For example, a
// request targeting `foo*,bar*` returns an error if an index starts with `foo`
// but no index starts with `bar`.
// API name: allow_no_indices
func (r *RankEval) AllowNoIndices(allownoindices bool) *RankEval {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *RankEval) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *RankEval {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreUnavailable If `true`, missing or closed indices are not included in the response.
// API name: ignore_unavailable
func (r *RankEval) IgnoreUnavailable(ignoreunavailable bool) *RankEval {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// SearchType Search operation type
// API name: search_type
func (r *RankEval) SearchType(searchtype string) *RankEval {
	r.values.Set("search_type", searchtype)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *RankEval) ErrorTrace(errortrace bool) *RankEval {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *RankEval) FilterPath(filterpaths ...string) *RankEval {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *RankEval) Human(human bool) *RankEval {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *RankEval) Pretty(pretty bool) *RankEval {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Metric Definition of the evaluation metric to calculate.
// API name: metric
func (r *RankEval) Metric(metric *types.RankEvalMetric) *RankEval {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Metric = metric

	return r
}

// Requests A set of typical search requests, together with their provided ratings.
// API name: requests
func (r *RankEval) Requests(requests ...types.RankEvalRequestItem) *RankEval {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Requests = requests

	return r
}
