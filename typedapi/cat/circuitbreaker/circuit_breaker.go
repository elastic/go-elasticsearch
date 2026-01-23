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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

// Get circuit breakers statistics.
//
// IMPORTANT: CAT APIs are only intended for human consumption using the command
// line or Kibana console. They are not intended for use by applications.
package circuitbreaker

import (
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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catcircuitbreakercolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	circuitbreakerpatternsMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CircuitBreaker struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	circuitbreakerpatterns string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewCircuitBreaker type alias for index.
type NewCircuitBreaker func() *CircuitBreaker

// NewCircuitBreakerFunc returns a new instance of CircuitBreaker with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCircuitBreakerFunc(tp elastictransport.Interface) NewCircuitBreaker {
	return func() *CircuitBreaker {
		n := New(tp)

		return n
	}
}

// Get circuit breakers statistics.
//
// IMPORTANT: CAT APIs are only intended for human consumption using the command
// line or Kibana console. They are not intended for use by applications.
//
// https://www.elastic.co/docs/api/doc/elasticsearch#TODO
func New(tp elastictransport.Interface) *CircuitBreaker {
	r := &CircuitBreaker{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *CircuitBreaker) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("circuit_breaker")

		method = http.MethodGet
	case r.paramSet == circuitbreakerpatternsMask:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("circuit_breaker")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "circuitbreakerpatterns", r.circuitbreakerpatterns)
		}
		path.WriteString(r.circuitbreakerpatterns)

		method = http.MethodGet
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=9")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r CircuitBreaker) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cat.circuit_breaker")
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
		instrument.BeforeRequest(req, "cat.circuit_breaker")
		if reader := instrument.RecordRequestBody(ctx, "cat.circuit_breaker", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cat.circuit_breaker")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the CircuitBreaker query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a circuitbreaker.Response
func (r CircuitBreaker) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cat.circuit_breaker")
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
		err = json.NewDecoder(res.Body).Decode(&response)
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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r CircuitBreaker) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cat.circuit_breaker")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(io.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the CircuitBreaker query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the CircuitBreaker headers map.
func (r *CircuitBreaker) Header(key, value string) *CircuitBreaker {
	r.headers.Set(key, value)

	return r
}

// CircuitBreakerPatterns A comma-separated list of regular-expressions to filter the circuit breakers
// in the output
// API Name: circuitbreakerpatterns
func (r *CircuitBreaker) CircuitBreakerPatterns(circuitbreakerpatterns ...string) *CircuitBreaker {
	r.paramSet |= circuitbreakerpatternsMask
	r.circuitbreakerpatterns = strings.Join(circuitbreakerpatterns, ",")

	return r
}

// H A comma-separated list of columns names to display. It supports simple
// wildcards.
// API name: h
func (r *CircuitBreaker) H(catcircuitbreakercolumns ...catcircuitbreakercolumn.CatCircuitBreakerColumn) *CircuitBreaker {
	tmp := []string{}
	for _, item := range catcircuitbreakercolumns {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// S List of columns that determine how the table should be sorted.
// Sorting defaults to ascending and can be changed by setting `:asc`
// or `:desc` as a suffix to the column name.
// API name: s
func (r *CircuitBreaker) S(names ...string) *CircuitBreaker {
	r.values.Set("s", strings.Join(names, ","))

	return r
}

// Local If `true`, the request computes the list of selected nodes from the
// local cluster state. If `false` the list of selected nodes are computed
// from the cluster state of the master node. In both cases the coordinating
// node will send requests for further information to each selected node.
// API name: local
func (r *CircuitBreaker) Local(local bool) *CircuitBreaker {
	r.values.Set("local", strconv.FormatBool(local))

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// API name: master_timeout
func (r *CircuitBreaker) MasterTimeout(duration string) *CircuitBreaker {
	r.values.Set("master_timeout", duration)

	return r
}

// Bytes Sets the units for columns that contain a byte-size value.
// Note that byte-size value units work in terms of powers of 1024. For instance
// `1kb` means 1024 bytes, not 1000 bytes.
// If omitted, byte-size values are rendered with a suffix such as `kb`, `mb`,
// or `gb`, chosen such that the numeric value of the column is as small as
// possible whilst still being at least `1.0`.
// If given, byte-size values are rendered as an integer with no suffix,
// representing the value of the column in the chosen unit.
// Values that are not an exact multiple of the chosen unit are rounded down.
// API name: bytes
func (r *CircuitBreaker) Bytes(bytes bytes.Bytes) *CircuitBreaker {
	r.values.Set("bytes", bytes.String())

	return r
}

// Format Specifies the format to return the columnar data in, can be set to
// `text`, `json`, `cbor`, `yaml`, or `smile`.
// API name: format
func (r *CircuitBreaker) Format(format string) *CircuitBreaker {
	r.values.Set("format", format)

	return r
}

// Help When set to `true` will output available columns. This option
// can't be combined with any other query string option.
// API name: help
func (r *CircuitBreaker) Help(help bool) *CircuitBreaker {
	r.values.Set("help", strconv.FormatBool(help))

	return r
}

// Time Sets the units for columns that contain a time duration.
// If omitted, time duration values are rendered with a suffix such as `ms`,
// `s`, `m` or `h`, chosen such that the numeric value of the column is as small
// as possible whilst still being at least `1.0`.
// If given, time duration values are rendered as an integer with no suffix.
// Values that are not an exact multiple of the chosen unit are rounded down.
// API name: time
func (r *CircuitBreaker) Time(time timeunit.TimeUnit) *CircuitBreaker {
	r.values.Set("time", time.String())

	return r
}

// V When set to `true` will enable verbose output.
// API name: v
func (r *CircuitBreaker) V(v bool) *CircuitBreaker {
	r.values.Set("v", strconv.FormatBool(v))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *CircuitBreaker) ErrorTrace(errortrace bool) *CircuitBreaker {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *CircuitBreaker) FilterPath(filterpaths ...string) *CircuitBreaker {
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
func (r *CircuitBreaker) Human(human bool) *CircuitBreaker {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *CircuitBreaker) Pretty(pretty bool) *CircuitBreaker {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
