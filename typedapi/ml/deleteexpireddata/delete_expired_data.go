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

// Delete expired ML data.
//
// Delete all job results, model snapshots and forecast data that have exceeded
// their retention days period. Machine learning state documents that are not
// associated with any job are also deleted.
// You can limit the request to a single or set of anomaly detection jobs by
// using a job identifier, a group name, a comma-separated list of jobs, or a
// wildcard expression. You can delete expired data for all anomaly detection
// jobs by using `_all`, by specifying `*` as the `<job_id>`, or by omitting the
// `<job_id>`.
package deleteexpireddata

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
	jobidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteExpiredData struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	jobid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewDeleteExpiredData type alias for index.
type NewDeleteExpiredData func() *DeleteExpiredData

// NewDeleteExpiredDataFunc returns a new instance of DeleteExpiredData with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteExpiredDataFunc(tp elastictransport.Interface) NewDeleteExpiredData {
	return func() *DeleteExpiredData {
		n := New(tp)

		return n
	}
}

// Delete expired ML data.
//
// Delete all job results, model snapshots and forecast data that have exceeded
// their retention days period. Machine learning state documents that are not
// associated with any job are also deleted.
// You can limit the request to a single or set of anomaly detection jobs by
// using a job identifier, a group name, a comma-separated list of jobs, or a
// wildcard expression. You can delete expired data for all anomaly detection
// jobs by using `_all`, by specifying `*` as the `<job_id>`, or by omitting the
// `<job_id>`.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-expired-data.html
func New(tp elastictransport.Interface) *DeleteExpiredData {
	r := &DeleteExpiredData{
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
func (r *DeleteExpiredData) Raw(raw io.Reader) *DeleteExpiredData {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *DeleteExpiredData) Request(req *Request) *DeleteExpiredData {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *DeleteExpiredData) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for DeleteExpiredData: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == jobidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("_delete_expired_data")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "jobid", r.jobid)
		}
		path.WriteString(r.jobid)

		method = http.MethodDelete
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("_delete_expired_data")

		method = http.MethodDelete
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
func (r DeleteExpiredData) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.delete_expired_data")
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
		instrument.BeforeRequest(req, "ml.delete_expired_data")
		if reader := instrument.RecordRequestBody(ctx, "ml.delete_expired_data", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.delete_expired_data")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the DeleteExpiredData query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a deleteexpireddata.Response
func (r DeleteExpiredData) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.delete_expired_data")
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

// Header set a key, value pair in the DeleteExpiredData headers map.
func (r *DeleteExpiredData) Header(key, value string) *DeleteExpiredData {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for an anomaly detection job. It can be a job identifier, a
// group name, or a wildcard expression.
// API Name: jobid
func (r *DeleteExpiredData) JobId(jobid string) *DeleteExpiredData {
	r.paramSet |= jobidMask
	r.jobid = jobid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *DeleteExpiredData) ErrorTrace(errortrace bool) *DeleteExpiredData {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *DeleteExpiredData) FilterPath(filterpaths ...string) *DeleteExpiredData {
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
func (r *DeleteExpiredData) Human(human bool) *DeleteExpiredData {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *DeleteExpiredData) Pretty(pretty bool) *DeleteExpiredData {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// RequestsPerSecond The desired requests per second for the deletion processes. The default
// behavior is no throttling.
// API name: requests_per_second
func (r *DeleteExpiredData) RequestsPerSecond(requestspersecond float32) *DeleteExpiredData {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RequestsPerSecond = &requestspersecond

	return r
}

// Timeout How long can the underlying delete processes run until they are canceled.
// API name: timeout
func (r *DeleteExpiredData) Timeout(duration types.Duration) *DeleteExpiredData {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Timeout = duration

	return r
}
