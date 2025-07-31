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

// Create a behavioral analytics collection event.
package postbehavioralanalyticsevent

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	collectionnameMask = iota + 1

	eventtypeMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PostBehavioralAnalyticsEvent struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      any
	deferred []func(request any) error
	buf      *gobytes.Buffer

	paramSet int

	collectionname string
	eventtype      string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPostBehavioralAnalyticsEvent type alias for index.
type NewPostBehavioralAnalyticsEvent func(collectionname, eventtype string) *PostBehavioralAnalyticsEvent

// NewPostBehavioralAnalyticsEventFunc returns a new instance of PostBehavioralAnalyticsEvent with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPostBehavioralAnalyticsEventFunc(tp elastictransport.Interface) NewPostBehavioralAnalyticsEvent {
	return func(collectionname, eventtype string) *PostBehavioralAnalyticsEvent {
		n := New(tp)

		n._collectionname(collectionname)

		n._eventtype(eventtype)

		return n
	}
}

// Create a behavioral analytics collection event.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/post-analytics-collection-event.html
func New(tp elastictransport.Interface) *PostBehavioralAnalyticsEvent {
	r := &PostBehavioralAnalyticsEvent{
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
func (r *PostBehavioralAnalyticsEvent) Raw(raw io.Reader) *PostBehavioralAnalyticsEvent {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PostBehavioralAnalyticsEvent) Request(req any) *PostBehavioralAnalyticsEvent {
	r.req = req

	return r
}

// Payload allows to set the request property with the appropriate payload.
func (r *PostBehavioralAnalyticsEvent) Payload(payload any) *PostBehavioralAnalyticsEvent {
	r.req = payload

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PostBehavioralAnalyticsEvent) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PostBehavioralAnalyticsEvent: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == collectionnameMask|eventtypeMask:
		path.WriteString("/")
		path.WriteString("_application")
		path.WriteString("/")
		path.WriteString("analytics")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "collectionname", r.collectionname)
		}
		path.WriteString(r.collectionname)
		path.WriteString("/")
		path.WriteString("event")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "eventtype", r.eventtype)
		}
		path.WriteString(r.eventtype)

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
func (r PostBehavioralAnalyticsEvent) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "search_application.post_behavioral_analytics_event")
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
		instrument.BeforeRequest(req, "search_application.post_behavioral_analytics_event")
		if reader := instrument.RecordRequestBody(ctx, "search_application.post_behavioral_analytics_event", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "search_application.post_behavioral_analytics_event")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PostBehavioralAnalyticsEvent query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a postbehavioralanalyticsevent.Response
func (r PostBehavioralAnalyticsEvent) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "search_application.post_behavioral_analytics_event")
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

	if res.StatusCode < 299 || slices.Contains([]int{202, 400, 404}, res.StatusCode) {

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

// Header set a key, value pair in the PostBehavioralAnalyticsEvent headers map.
func (r *PostBehavioralAnalyticsEvent) Header(key, value string) *PostBehavioralAnalyticsEvent {
	r.headers.Set(key, value)

	return r
}

// CollectionName The name of the behavioral analytics collection.
// API Name: collectionname
func (r *PostBehavioralAnalyticsEvent) _collectionname(collectionname string) *PostBehavioralAnalyticsEvent {
	r.paramSet |= collectionnameMask
	r.collectionname = collectionname

	return r
}

// EventType The analytics event type.
// API Name: eventtype
func (r *PostBehavioralAnalyticsEvent) _eventtype(eventtype string) *PostBehavioralAnalyticsEvent {
	r.paramSet |= eventtypeMask
	r.eventtype = eventtype

	return r
}

// Debug Whether the response type has to include more details
// API name: debug
func (r *PostBehavioralAnalyticsEvent) Debug(debug bool) *PostBehavioralAnalyticsEvent {
	r.values.Set("debug", strconv.FormatBool(debug))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PostBehavioralAnalyticsEvent) ErrorTrace(errortrace bool) *PostBehavioralAnalyticsEvent {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PostBehavioralAnalyticsEvent) FilterPath(filterpaths ...string) *PostBehavioralAnalyticsEvent {
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
func (r *PostBehavioralAnalyticsEvent) Human(human bool) *PostBehavioralAnalyticsEvent {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PostBehavioralAnalyticsEvent) Pretty(pretty bool) *PostBehavioralAnalyticsEvent {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
