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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Run a scrolling search.
//
// IMPORTANT: The scroll API is no longer recommend for deep pagination. If you
// need to preserve the index state while paging through more than 10,000 hits,
// use the `search_after` parameter with a point in time (PIT).
//
// The scroll API gets large sets of results from a single scrolling search
// request.
// To get the necessary scroll ID, submit a search API request that includes an
// argument for the `scroll` query parameter.
// The `scroll` parameter indicates how long Elasticsearch should retain the
// search context for the request.
// The search response returns a scroll ID in the `_scroll_id` response body
// parameter.
// You can then use the scroll ID with the scroll API to retrieve the next batch
// of results for the request.
// If the Elasticsearch security features are enabled, the access to the results
// of a specific scroll ID is restricted to the user or API key that submitted
// the search.
//
// You can also use the scroll API to specify a new scroll parameter that
// extends or shortens the retention period for the search context.
//
// IMPORTANT: Results from a scrolling search reflect the state of the index at
// the time of the initial search request. Subsequent indexing or document
// changes only affect later search and scroll requests.
package scroll

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
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Scroll struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	scrollid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewScroll type alias for index.
type NewScroll func() *Scroll

// NewScrollFunc returns a new instance of Scroll with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewScrollFunc(tp elastictransport.Interface) NewScroll {
	return func() *Scroll {
		n := New(tp)

		return n
	}
}

// Run a scrolling search.
//
// IMPORTANT: The scroll API is no longer recommend for deep pagination. If you
// need to preserve the index state while paging through more than 10,000 hits,
// use the `search_after` parameter with a point in time (PIT).
//
// The scroll API gets large sets of results from a single scrolling search
// request.
// To get the necessary scroll ID, submit a search API request that includes an
// argument for the `scroll` query parameter.
// The `scroll` parameter indicates how long Elasticsearch should retain the
// search context for the request.
// The search response returns a scroll ID in the `_scroll_id` response body
// parameter.
// You can then use the scroll ID with the scroll API to retrieve the next batch
// of results for the request.
// If the Elasticsearch security features are enabled, the access to the results
// of a specific scroll ID is restricted to the user or API key that submitted
// the search.
//
// You can also use the scroll API to specify a new scroll parameter that
// extends or shortens the retention period for the search context.
//
// IMPORTANT: Results from a scrolling search reflect the state of the index at
// the time of the initial search request. Subsequent indexing or document
// changes only affect later search and scroll requests.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-scroll
func New(tp elastictransport.Interface) *Scroll {
	r := &Scroll{
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
func (r *Scroll) Raw(raw io.Reader) *Scroll {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Scroll) Request(req *Request) *Scroll {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Scroll) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Scroll: %w", err)
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
		path.WriteString("_search")
		path.WriteString("/")
		path.WriteString("scroll")

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
func (r Scroll) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "scroll")
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
		instrument.BeforeRequest(req, "scroll")
		if reader := instrument.RecordRequestBody(ctx, "scroll", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "scroll")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Scroll query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a scroll.Response
func (r Scroll) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "scroll")
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

// Header set a key, value pair in the Scroll headers map.
func (r *Scroll) Header(key, value string) *Scroll {
	r.headers.Set(key, value)

	return r
}

// RestTotalHitsAsInt If true, the API response’s hit.total property is returned as an integer. If
// false, the API response’s hit.total property is returned as an object.
// API name: rest_total_hits_as_int
func (r *Scroll) RestTotalHitsAsInt(resttotalhitsasint bool) *Scroll {
	r.values.Set("rest_total_hits_as_int", strconv.FormatBool(resttotalhitsasint))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Scroll) ErrorTrace(errortrace bool) *Scroll {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Scroll) FilterPath(filterpaths ...string) *Scroll {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"exists_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *Scroll) Human(human bool) *Scroll {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Scroll) Pretty(pretty bool) *Scroll {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// The period to retain the search context for scrolling.
// API name: scroll
func (r *Scroll) Scroll(duration types.DurationVariant) *Scroll {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Scroll = *duration.DurationCaster()

	return r
}

// The scroll ID of the search.
// API name: scroll_id
func (r *Scroll) ScrollId(scrollid string) *Scroll {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ScrollId = scrollid

	return r
}
