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
// Code generated from specification version 8.12.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func newSearchApplicationPostBehavioralAnalyticsEventFunc(t Transport) SearchApplicationPostBehavioralAnalyticsEvent {
	return func(body io.Reader, collection_name string, event_type string, o ...func(*SearchApplicationPostBehavioralAnalyticsEventRequest)) (*Response, error) {
		var r = SearchApplicationPostBehavioralAnalyticsEventRequest{Body: body, CollectionName: collection_name, EventType: event_type}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SearchApplicationPostBehavioralAnalyticsEvent creates a behavioral analytics event for existing collection.
//
// This API is experimental.
//
// See full documentation at http://todo.com/tbd.
type SearchApplicationPostBehavioralAnalyticsEvent func(body io.Reader, collection_name string, event_type string, o ...func(*SearchApplicationPostBehavioralAnalyticsEventRequest)) (*Response, error)

// SearchApplicationPostBehavioralAnalyticsEventRequest configures the Search Application Post Behavioral Analytics Event API request.
type SearchApplicationPostBehavioralAnalyticsEventRequest struct {
	Body io.Reader

	CollectionName string
	EventType      string

	Debug *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SearchApplicationPostBehavioralAnalyticsEventRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "search_application.post_behavioral_analytics_event")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_application") + 1 + len("analytics") + 1 + len(r.CollectionName) + 1 + len("event") + 1 + len(r.EventType))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_application")
	path.WriteString("/")
	path.WriteString("analytics")
	path.WriteString("/")
	path.WriteString(r.CollectionName)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "collection_name", r.CollectionName)
	}
	path.WriteString("/")
	path.WriteString("event")
	path.WriteString("/")
	path.WriteString(r.EventType)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "event_type", r.EventType)
	}

	params = make(map[string]string)

	if r.Debug != nil {
		params["debug"] = strconv.FormatBool(*r.Debug)
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
		if instrument, ok := r.instrument.(Instrumentation); ok {
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

	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "search_application.post_behavioral_analytics_event")
		if reader := instrument.RecordRequestBody(ctx, "search_application.post_behavioral_analytics_event", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "search_application.post_behavioral_analytics_event")
	}
	if err != nil {
		if instrument, ok := r.instrument.(Instrumentation); ok {
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
func (f SearchApplicationPostBehavioralAnalyticsEvent) WithContext(v context.Context) func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	return func(r *SearchApplicationPostBehavioralAnalyticsEventRequest) {
		r.ctx = v
	}
}

// WithDebug - if true, returns event information that will be stored.
func (f SearchApplicationPostBehavioralAnalyticsEvent) WithDebug(v bool) func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	return func(r *SearchApplicationPostBehavioralAnalyticsEventRequest) {
		r.Debug = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SearchApplicationPostBehavioralAnalyticsEvent) WithPretty() func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	return func(r *SearchApplicationPostBehavioralAnalyticsEventRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SearchApplicationPostBehavioralAnalyticsEvent) WithHuman() func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	return func(r *SearchApplicationPostBehavioralAnalyticsEventRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SearchApplicationPostBehavioralAnalyticsEvent) WithErrorTrace() func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	return func(r *SearchApplicationPostBehavioralAnalyticsEventRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SearchApplicationPostBehavioralAnalyticsEvent) WithFilterPath(v ...string) func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	return func(r *SearchApplicationPostBehavioralAnalyticsEventRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SearchApplicationPostBehavioralAnalyticsEvent) WithHeader(h map[string]string) func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	return func(r *SearchApplicationPostBehavioralAnalyticsEventRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SearchApplicationPostBehavioralAnalyticsEvent) WithOpaqueID(s string) func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	return func(r *SearchApplicationPostBehavioralAnalyticsEventRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
