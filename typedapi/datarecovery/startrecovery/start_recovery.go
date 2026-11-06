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

// Start a recovery.
//
// Start recovering indices and data streams from a recovery point. This API is
// intended for internal operator use. Depending on the mode, the sources are
// either replaced with their state at the recovery point, over the data that
// currently exists, or recovered under new names. The recovery runs in the
// background, so the API returns with a `202 Accepted` status once it has
// started.
package startrecovery

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/recoverymode"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type StartRecovery struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewStartRecovery type alias for index.
type NewStartRecovery func() *StartRecovery

// NewStartRecoveryFunc returns a new instance of StartRecovery with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewStartRecoveryFunc(tp elastictransport.Interface) NewStartRecovery {
	return func() *StartRecovery {
		n := New(tp)

		return n
	}
}

// Start a recovery.
//
// Start recovering indices and data streams from a recovery point. This API is
// intended for internal operator use. Depending on the mode, the sources are
// either replaced with their state at the recovery point, over the data that
// currently exists, or recovered under new names. The recovery runs in the
// background, so the API returns with a `202 Accepted` status once it has
// started.
func New(tp elastictransport.Interface) *StartRecovery {
	r := &StartRecovery{
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
func (r *StartRecovery) Raw(raw io.Reader) *StartRecovery {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *StartRecovery) Request(req *Request) *StartRecovery {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *StartRecovery) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for StartRecovery: %w", err)
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
		path.WriteString("_data_recovery")
		path.WriteString("/")
		path.WriteString("recoveries")

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
func (r StartRecovery) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx = instrument.Start(providedCtx, "data_recovery.start_recovery")
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
		instrument.BeforeRequest(req, "data_recovery.start_recovery")
		if reader := instrument.RecordRequestBody(ctx, "data_recovery.start_recovery", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "data_recovery.start_recovery")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the StartRecovery query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a startrecovery.Response
func (r StartRecovery) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "data_recovery.start_recovery")
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

// Header set a key, value pair in the StartRecovery headers map.
func (r *StartRecovery) Header(key, value string) *StartRecovery {
	r.headers.Set(key, value)

	return r
}

// MasterTimeout The period to wait for the master node. If the master node is not available
// before the timeout expires, the request fails and returns an error. To
// indicate that the request should never timeout, set it to `-1`.
// API name: master_timeout
func (r *StartRecovery) MasterTimeout(duration string) *StartRecovery {
	r.values.Set("master_timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *StartRecovery) ErrorTrace(errortrace bool) *StartRecovery {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *StartRecovery) FilterPath(filterpaths ...string) *StartRecovery {
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
func (r *StartRecovery) Human(human bool) *StartRecovery {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use this
// option for debugging only.
// API name: pretty
func (r *StartRecovery) Pretty(pretty bool) *StartRecovery {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Custom metadata for the recovery.
// API name: metadata
func (r *StartRecovery) Metadata(metadata types.MetadataVariant) *StartRecovery {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Metadata = *metadata.MetadataCaster()

	return r
}

// How to recover the sources: `replace` recovers them in place, over the data
// that currently exists, and `rename` recovers them under new names.
// API name: mode
func (r *StartRecovery) Mode(mode recoverymode.RecoveryMode) *StartRecovery {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Mode = mode
	return r
}

// The exact end time of the recovery point to recover from, as a date or epoch
// milliseconds. An `end_time` returned when listing recovery points can be
// passed back as it is.
// API name: recovery_point_end_time
func (r *StartRecovery) RecoveryPointEndTime(datetime types.DateTimeVariant) *StartRecovery {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RecoveryPointEndTime = *datetime.DateTimeCaster()

	return r
}

// How to derive the new name of each recovered source from its name. Required
// when `mode` is `rename`, and not allowed when `mode` is `replace`.
// API name: rename_strategy
func (r *StartRecovery) RenameStrategy(renamestrategy types.RenameStrategyVariant) *StartRecovery {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RenameStrategy = renamestrategy.RenameStrategyCaster()

	return r
}

// The sources to recover. Each value is an index or data stream expression.
// Wildcards and exclusions prefixed with `-` are supported. The list must not
// be empty and must not contain blank values.
// API name: sources
func (r *StartRecovery) Sources(sources ...string) *StartRecovery {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range sources {

		r.req.Sources = append(r.req.Sources, v)

	}
	return r
}
