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

// Update a data frame analytics job.
package updatedataframeanalytics

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
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateDataFrameAnalytics struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewUpdateDataFrameAnalytics type alias for index.
type NewUpdateDataFrameAnalytics func(id string) *UpdateDataFrameAnalytics

// NewUpdateDataFrameAnalyticsFunc returns a new instance of UpdateDataFrameAnalytics with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpdateDataFrameAnalyticsFunc(tp elastictransport.Interface) NewUpdateDataFrameAnalytics {
	return func(id string) *UpdateDataFrameAnalytics {
		n := New(tp)

		n._id(id)

		return n
	}
}

// Update a data frame analytics job.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-ml-update-data-frame-analytics
func New(tp elastictransport.Interface) *UpdateDataFrameAnalytics {
	r := &UpdateDataFrameAnalytics{
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
func (r *UpdateDataFrameAnalytics) Raw(raw io.Reader) *UpdateDataFrameAnalytics {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *UpdateDataFrameAnalytics) Request(req *Request) *UpdateDataFrameAnalytics {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *UpdateDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for UpdateDataFrameAnalytics: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("data_frame")
		path.WriteString("/")
		path.WriteString("analytics")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)
		path.WriteString("/")
		path.WriteString("_update")

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
func (r UpdateDataFrameAnalytics) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.update_data_frame_analytics")
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
		instrument.BeforeRequest(req, "ml.update_data_frame_analytics")
		if reader := instrument.RecordRequestBody(ctx, "ml.update_data_frame_analytics", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.update_data_frame_analytics")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the UpdateDataFrameAnalytics query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a updatedataframeanalytics.Response
func (r UpdateDataFrameAnalytics) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.update_data_frame_analytics")
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

// Header set a key, value pair in the UpdateDataFrameAnalytics headers map.
func (r *UpdateDataFrameAnalytics) Header(key, value string) *UpdateDataFrameAnalytics {
	r.headers.Set(key, value)

	return r
}

// Id Identifier for the data frame analytics job. This identifier can contain
// lowercase alphanumeric characters (a-z and 0-9), hyphens, and
// underscores. It must start and end with alphanumeric characters.
// API Name: id
func (r *UpdateDataFrameAnalytics) _id(id string) *UpdateDataFrameAnalytics {
	r.paramSet |= idMask
	r.id = id

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *UpdateDataFrameAnalytics) ErrorTrace(errortrace bool) *UpdateDataFrameAnalytics {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *UpdateDataFrameAnalytics) FilterPath(filterpaths ...string) *UpdateDataFrameAnalytics {
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
func (r *UpdateDataFrameAnalytics) Human(human bool) *UpdateDataFrameAnalytics {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *UpdateDataFrameAnalytics) Pretty(pretty bool) *UpdateDataFrameAnalytics {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// AllowLazyStart Specifies whether this job can start when there is insufficient machine
// learning node capacity for it to be immediately assigned to a node.
// API name: allow_lazy_start
func (r *UpdateDataFrameAnalytics) AllowLazyStart(allowlazystart bool) *UpdateDataFrameAnalytics {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.AllowLazyStart = &allowlazystart

	return r
}

// Description A description of the job.
// API name: description
func (r *UpdateDataFrameAnalytics) Description(description string) *UpdateDataFrameAnalytics {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Description = &description

	return r
}

// MaxNumThreads The maximum number of threads to be used by the analysis. Using more
// threads may decrease the time necessary to complete the analysis at the
// cost of using more CPU. Note that the process may use additional threads
// for operational functionality other than the analysis itself.
// API name: max_num_threads
func (r *UpdateDataFrameAnalytics) MaxNumThreads(maxnumthreads int) *UpdateDataFrameAnalytics {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxNumThreads = &maxnumthreads

	return r
}

// ModelMemoryLimit The approximate maximum amount of memory resources that are permitted for
// analytical processing. If your `elasticsearch.yml` file contains an
// `xpack.ml.max_model_memory_limit` setting, an error occurs when you try
// to create data frame analytics jobs that have `model_memory_limit` values
// greater than that setting.
// API name: model_memory_limit
func (r *UpdateDataFrameAnalytics) ModelMemoryLimit(modelmemorylimit string) *UpdateDataFrameAnalytics {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ModelMemoryLimit = &modelmemorylimit

	return r
}
