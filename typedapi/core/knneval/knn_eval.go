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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

// Evaluate kNN search settings.
//
// Estimate the recall and cost of approximate kNN search configurations on a
// `dense_vector` field against a more thorough baseline. The field must use the
// `bbq_disk` index type, and must not be inside a nested object.
//
// Each `knn_settings` entry reports the mean recall of the baseline's top `k`,
// not of relevance judgments, plus included and excluded query counts and cost
// totals. A query is excluded from the mean recall when the candidate finds a
// document scoring above the baseline's lowest top `k` score. A baseline
// returning fewer than `k` hits fails that query, because recall at `k` is
// undefined. All configurations run against one point in time.
//
// The default baseline uses a `visit_percentage` of 20 and a
// `rescore_vector.oversample` of 100. An exact full-precision baseline must be
// requested explicitly with `allow_exact_baseline`, and is limited to 100
// million vector-by-query comparisons.
//
// The request is refused when the `search.knn_eval.enabled` cluster setting or
// the `index.knn_eval.enabled` index setting is `false`.
package knneval

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type KnnEval struct {
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

// NewKnnEval type alias for index.
type NewKnnEval func(index string) *KnnEval

// NewKnnEvalFunc returns a new instance of KnnEval with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewKnnEvalFunc(tp elastictransport.Interface) NewKnnEval {
	return func(index string) *KnnEval {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Evaluate kNN search settings.
//
// Estimate the recall and cost of approximate kNN search configurations on a
// `dense_vector` field against a more thorough baseline. The field must use the
// `bbq_disk` index type, and must not be inside a nested object.
//
// Each `knn_settings` entry reports the mean recall of the baseline's top `k`,
// not of relevance judgments, plus included and excluded query counts and cost
// totals. A query is excluded from the mean recall when the candidate finds a
// document scoring above the baseline's lowest top `k` score. A baseline
// returning fewer than `k` hits fails that query, because recall at `k` is
// undefined. All configurations run against one point in time.
//
// The default baseline uses a `visit_percentage` of 20 and a
// `rescore_vector.oversample` of 100. An exact full-precision baseline must be
// requested explicitly with `allow_exact_baseline`, and is limited to 100
// million vector-by-query comparisons.
//
// The request is refused when the `search.knn_eval.enabled` cluster setting or
// the `index.knn_eval.enabled` index setting is `false`.
//
// [Elasticsearch] https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-knn-eval
//
// [Serverless] https://www.elastic.co/docs/api/doc/elasticsearch-serverless/operation/operation-knn-eval
func New(tp elastictransport.Interface) *KnnEval {
	r := &KnnEval{
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
func (r *KnnEval) Raw(raw io.Reader) *KnnEval {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *KnnEval) Request(req *Request) *KnnEval {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *KnnEval) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for KnnEval: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_knn_eval")

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
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=9")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=9")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r KnnEval) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx = instrument.Start(providedCtx, "knn_eval")
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
		instrument.BeforeRequest(req, "knn_eval")
		if reader := instrument.RecordRequestBody(ctx, "knn_eval", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "knn_eval")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the KnnEval query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a knneval.Response
func (r KnnEval) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "knn_eval")
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

// Header set a key, value pair in the KnnEval headers map.
func (r *KnnEval) Header(key, value string) *KnnEval {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names to search.
// API Name: index
func (r *KnnEval) _index(index string) *KnnEval {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// IgnoreUnavailable If `false`, the request returns an error if it targets a missing or closed
// index.
// API name: ignore_unavailable
func (r *KnnEval) IgnoreUnavailable(ignoreunavailable bool) *KnnEval {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// AllowNoIndices If `false`, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices. This behavior
// applies even if the request targets other open indices.
// API name: allow_no_indices
func (r *KnnEval) AllowNoIndices(allownoindices bool) *KnnEval {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards The type of index that wildcard patterns can match. It supports
// comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *KnnEval) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *KnnEval {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// WaitForCompletion If `true`, the request waits for the evaluation to complete. A large
// evaluation can run for minutes, and closing the connection, for example
// because of a client or proxy timeout, cancels it. If `false`, the request
// returns a task identifier, and stores the completed response for retrieval
// through the task management API.
// API name: wait_for_completion
func (r *KnnEval) WaitForCompletion(waitforcompletion bool) *KnnEval {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *KnnEval) ErrorTrace(errortrace bool) *KnnEval {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *KnnEval) FilterPath(filterpaths ...string) *KnnEval {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and `"exists_time_in_millis":
// 3600000` for computers. When disabled the human readable values will be
// omitted. This makes sense for responses being consumed only by machines.
// API name: human
func (r *KnnEval) Human(human bool) *KnnEval {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use this
// option for debugging only.
// API name: pretty
func (r *KnnEval) Pretty(pretty bool) *KnnEval {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// If `true`, the `baseline` can be `exact`. An exact baseline compares every
// vector at full precision, and can severely slow down other searches.
// API name: allow_exact_baseline
func (r *KnnEval) AllowExactBaseline(allowexactbaseline bool) *KnnEval {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AllowExactBaseline = &allowexactbaseline

	return r
}

// The settings of the search that the `knn_settings` entries are compared with.
// If not specified, it uses a `visit_percentage` of 20 and a
// `rescore_vector.oversample` of 100.
// API name: baseline
func (r *KnnEval) Baseline(baseline types.KnnEvalSettingsVariant) *KnnEval {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Baseline = baseline.KnnEvalSettingsCaster()

	return r
}

// The `dense_vector` field to evaluate.
// API name: field
func (r *KnnEval) Field(field string) *KnnEval {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Field = field

	return r
}

// The number of nearest neighbors to compute the recall for, from 1 to 1000.
// API name: k
func (r *KnnEval) K(k int) *KnnEval {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.K = k

	return r
}

// The kNN search settings to evaluate, from 1 to 32 unique entries.
// API name: knn_settings
func (r *KnnEval) KnnSettings(knnsettings ...types.KnnEvalSettingsVariant) *KnnEval {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range knnsettings {

		r.req.KnnSettings = append(r.req.KnnSettings, *v.KnnEvalSettingsCaster())

	}
	return r
}

func (r *KnnEval) KnnSettingsValues(knnsettingsvalues []types.KnnEvalSettings) *KnnEval {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.KnnSettings = knnsettingsvalues
	return r
}

// The source of the query vectors.
// API name: query_source
func (r *KnnEval) QuerySource(knnevalquerysource types.KnnEvalQuerySourceVariant) *KnnEval {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.QuerySource = *knnevalquerysource.KnnEvalQuerySourceCaster()

	return r
}
