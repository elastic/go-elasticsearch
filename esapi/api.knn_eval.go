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
//
// Code generated from specification version 9.6.0: DO NOT EDIT

package esapi

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func newKnnEvalFunc(t Transport) KnnEval {
	return func(index []string, body io.Reader, o ...func(*KnnEvalRequest)) (*Response, error) {
		var r = KnnEvalRequest{Index: index, Body: body}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.Instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// KnnEval estimate the recall and cost of approximate kNN search configurations on a bbq_disk dense_vector field against a more thorough baseline. Requires the read and view_index_metadata index privileges: the field mapping is validated as the calling user. Refused when search.knn_eval.enabled or the index's index.knn_eval.enabled is false. The bounded default baseline uses 20% visit with rescore_vector.oversample 100; an exact full-precision baseline must be requested explicitly with allow_exact_baseline:true and is limited to 100 million vector-by-query comparisons. Each knn_settings entry reports mean recall of the baseline's top-k, not of relevance judgments, plus included/excluded query counts and cost totals. A query is excluded from mean recall when the candidate finds a document scoring above the baseline's lowest top-k score. A baseline returning fewer than k hits fails that query because recall@k is undefined. A rescore_window_capped flag reports when search_size multiplied by oversample exceeds the 10,000-vector rescore limit. All configurations run against one point-in-time.
//
// This API is experimental.
type KnnEval func(index []string, body io.Reader, o ...func(*KnnEvalRequest)) (*Response, error)

// KnnEvalRequest configures the Knn Eval API request.
type KnnEvalRequest struct {
	Index []string

	Body io.Reader

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreUnavailable *bool
	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r KnnEvalRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "knn_eval")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	if len(r.Index) == 0 {
		return nil, errors.New("index is required and cannot be nil or empty")
	}

	path.Grow(7 + 1 + len(strings.Join(r.Index, ",")) + 1 + len("_knn_eval"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString(strings.Join(r.Index, ","))
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "index", strings.Join(r.Index, ","))
	}
	path.WriteString("/")
	path.WriteString("_knn_eval")

	params = make(map[string]string)

	if r.AllowNoIndices != nil {
		params["allow_no_indices"] = strconv.FormatBool(*r.AllowNoIndices)
	}

	if len(r.ExpandWildcards) > 0 {
		params["expand_wildcards"] = strings.Join(r.ExpandWildcards, ",")
	}

	if r.IgnoreUnavailable != nil {
		params["ignore_unavailable"] = strconv.FormatBool(*r.IgnoreUnavailable)
	}

	if r.WaitForCompletion != nil {
		params["wait_for_completion"] = strconv.FormatBool(*r.WaitForCompletion)
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), r.Body)
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if r.Body != nil && req.Header.Get(headerContentType) == "" {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "knn_eval")
		if reader := instrument.RecordRequestBody(ctx, "knn_eval", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "knn_eval")
	}
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
func (f KnnEval) WithContext(v context.Context) func(*KnnEvalRequest) {
	return func(r *KnnEvalRequest) {
		r.ctx = v
	}
}

// WithAllowNoIndices - whether to allow (1) wildcard index expressions that match no indices and (2) requests where the final resolved set is empty..
func (f KnnEval) WithAllowNoIndices(v bool) func(*KnnEvalRequest) {
	return func(r *KnnEvalRequest) {
		r.AllowNoIndices = &v
	}
}

// WithExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both..
func (f KnnEval) WithExpandWildcards(v ...string) func(*KnnEvalRequest) {
	return func(r *KnnEvalRequest) {
		r.ExpandWildcards = v
	}
}

// WithIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func (f KnnEval) WithIgnoreUnavailable(v bool) func(*KnnEvalRequest) {
	return func(r *KnnEvalRequest) {
		r.IgnoreUnavailable = &v
	}
}

// WithWaitForCompletion - whether to wait for the evaluation to complete. a large sweep can run for minutes, and closing the connection (a client or proxy timeout, for example) cancels it, so use false for those. if false, return a task identifier and store the completed response for retrieval through the task api..
func (f KnnEval) WithWaitForCompletion(v bool) func(*KnnEvalRequest) {
	return func(r *KnnEvalRequest) {
		r.WaitForCompletion = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f KnnEval) WithPretty() func(*KnnEvalRequest) {
	return func(r *KnnEvalRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f KnnEval) WithHuman() func(*KnnEvalRequest) {
	return func(r *KnnEvalRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f KnnEval) WithErrorTrace() func(*KnnEvalRequest) {
	return func(r *KnnEvalRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f KnnEval) WithFilterPath(v ...string) func(*KnnEvalRequest) {
	return func(r *KnnEvalRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f KnnEval) WithHeader(h map[string]string) func(*KnnEvalRequest) {
	return func(r *KnnEvalRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f KnnEval) WithOpaqueID(s string) func(*KnnEvalRequest) {
	return func(r *KnnEvalRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
