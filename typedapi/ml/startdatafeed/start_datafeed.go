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

// Start datafeeds.
//
// A datafeed must be started in order to retrieve data from Elasticsearch. A
// datafeed can be started and stopped
// multiple times throughout its lifecycle.
//
// Before you can start a datafeed, the anomaly detection job must be open.
// Otherwise, an error occurs.
//
// If you restart a stopped datafeed, it continues processing input data from
// the next millisecond after it was stopped.
// If new data was indexed for that exact millisecond between stopping and
// starting, it will be ignored.
//
// When Elasticsearch security features are enabled, your datafeed remembers
// which roles the last user to create or
// update it had at the time of creation or update and runs the query using
// those same roles. If you provided secondary
// authorization headers when you created or updated the datafeed, those
// credentials are used instead.
package startdatafeed

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

const (
	datafeedidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type StartDatafeed struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	datafeedid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewStartDatafeed type alias for index.
type NewStartDatafeed func(datafeedid string) *StartDatafeed

// NewStartDatafeedFunc returns a new instance of StartDatafeed with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewStartDatafeedFunc(tp elastictransport.Interface) NewStartDatafeed {
	return func(datafeedid string) *StartDatafeed {
		n := New(tp)

		n._datafeedid(datafeedid)

		return n
	}
}

// Start datafeeds.
//
// A datafeed must be started in order to retrieve data from Elasticsearch. A
// datafeed can be started and stopped
// multiple times throughout its lifecycle.
//
// Before you can start a datafeed, the anomaly detection job must be open.
// Otherwise, an error occurs.
//
// If you restart a stopped datafeed, it continues processing input data from
// the next millisecond after it was stopped.
// If new data was indexed for that exact millisecond between stopping and
// starting, it will be ignored.
//
// When Elasticsearch security features are enabled, your datafeed remembers
// which roles the last user to create or
// update it had at the time of creation or update and runs the query using
// those same roles. If you provided secondary
// authorization headers when you created or updated the datafeed, those
// credentials are used instead.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-start-datafeed.html
func New(tp elastictransport.Interface) *StartDatafeed {
	r := &StartDatafeed{
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
func (r *StartDatafeed) Raw(raw io.Reader) *StartDatafeed {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *StartDatafeed) Request(req *Request) *StartDatafeed {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *StartDatafeed) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for StartDatafeed: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == datafeedidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("datafeeds")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "datafeedid", r.datafeedid)
		}
		path.WriteString(r.datafeedid)
		path.WriteString("/")
		path.WriteString("_start")

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
func (r StartDatafeed) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.start_datafeed")
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
		instrument.BeforeRequest(req, "ml.start_datafeed")
		if reader := instrument.RecordRequestBody(ctx, "ml.start_datafeed", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.start_datafeed")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the StartDatafeed query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a startdatafeed.Response
func (r StartDatafeed) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.start_datafeed")
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

// Header set a key, value pair in the StartDatafeed headers map.
func (r *StartDatafeed) Header(key, value string) *StartDatafeed {
	r.headers.Set(key, value)

	return r
}

// DatafeedId A numerical character string that uniquely identifies the datafeed. This
// identifier can contain lowercase
// alphanumeric characters (a-z and 0-9), hyphens, and underscores. It must
// start and end with alphanumeric
// characters.
// API Name: datafeedid
func (r *StartDatafeed) _datafeedid(datafeedid string) *StartDatafeed {
	r.paramSet |= datafeedidMask
	r.datafeedid = datafeedid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *StartDatafeed) ErrorTrace(errortrace bool) *StartDatafeed {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *StartDatafeed) FilterPath(filterpaths ...string) *StartDatafeed {
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
func (r *StartDatafeed) Human(human bool) *StartDatafeed {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *StartDatafeed) Pretty(pretty bool) *StartDatafeed {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// End Refer to the description for the `end` query parameter.
// API name: end
func (r *StartDatafeed) End(datetime types.DateTime) *StartDatafeed {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.End = datetime

	return r
}

// Start Refer to the description for the `start` query parameter.
// API name: start
func (r *StartDatafeed) Start(datetime types.DateTime) *StartDatafeed {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Start = datetime

	return r
}

// Timeout Refer to the description for the `timeout` query parameter.
// API name: timeout
func (r *StartDatafeed) Timeout(duration types.Duration) *StartDatafeed {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Timeout = duration

	return r
}
