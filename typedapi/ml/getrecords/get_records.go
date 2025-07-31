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

// Get anomaly records for an anomaly detection job.
// Records contain the detailed analytical results. They describe the anomalous
// activity that has been identified in the input data based on the detector
// configuration.
// There can be many anomaly records depending on the characteristics and size
// of the input data. In practice, there are often too many to be able to
// manually process them. The machine learning features therefore perform a
// sophisticated aggregation of the anomaly records into buckets.
// The number of record results depends on the number of anomalies found in each
// bucket, which relates to the number of time series being modeled and the
// number of detectors.
package getrecords

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

type GetRecords struct {
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

// NewGetRecords type alias for index.
type NewGetRecords func(jobid string) *GetRecords

// NewGetRecordsFunc returns a new instance of GetRecords with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetRecordsFunc(tp elastictransport.Interface) NewGetRecords {
	return func(jobid string) *GetRecords {
		n := New(tp)

		n._jobid(jobid)

		return n
	}
}

// Get anomaly records for an anomaly detection job.
// Records contain the detailed analytical results. They describe the anomalous
// activity that has been identified in the input data based on the detector
// configuration.
// There can be many anomaly records depending on the characteristics and size
// of the input data. In practice, there are often too many to be able to
// manually process them. The machine learning features therefore perform a
// sophisticated aggregation of the anomaly records into buckets.
// The number of record results depends on the number of anomalies found in each
// bucket, which relates to the number of time series being modeled and the
// number of detectors.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-record.html
func New(tp elastictransport.Interface) *GetRecords {
	r := &GetRecords{
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
func (r *GetRecords) Raw(raw io.Reader) *GetRecords {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *GetRecords) Request(req *Request) *GetRecords {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetRecords) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for GetRecords: %w", err)
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
		path.WriteString("anomaly_detectors")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "jobid", r.jobid)
		}
		path.WriteString(r.jobid)
		path.WriteString("/")
		path.WriteString("results")
		path.WriteString("/")
		path.WriteString("records")

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
func (r GetRecords) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.get_records")
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
		instrument.BeforeRequest(req, "ml.get_records")
		if reader := instrument.RecordRequestBody(ctx, "ml.get_records", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.get_records")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GetRecords query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getrecords.Response
func (r GetRecords) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_records")
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

// Header set a key, value pair in the GetRecords headers map.
func (r *GetRecords) Header(key, value string) *GetRecords {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the anomaly detection job.
// API Name: jobid
func (r *GetRecords) _jobid(jobid string) *GetRecords {
	r.paramSet |= jobidMask
	r.jobid = jobid

	return r
}

// From Skips the specified number of records.
// API name: from
func (r *GetRecords) From(from int) *GetRecords {
	r.values.Set("from", strconv.Itoa(from))

	return r
}

// Size Specifies the maximum number of records to obtain.
// API name: size
func (r *GetRecords) Size(size int) *GetRecords {
	r.values.Set("size", strconv.Itoa(size))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *GetRecords) ErrorTrace(errortrace bool) *GetRecords {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *GetRecords) FilterPath(filterpaths ...string) *GetRecords {
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
func (r *GetRecords) Human(human bool) *GetRecords {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *GetRecords) Pretty(pretty bool) *GetRecords {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Desc Refer to the description for the `desc` query parameter.
// API name: desc
func (r *GetRecords) Desc(desc bool) *GetRecords {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Desc = &desc

	return r
}

// End Refer to the description for the `end` query parameter.
// API name: end
func (r *GetRecords) End(datetime types.DateTime) *GetRecords {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.End = datetime

	return r
}

// ExcludeInterim Refer to the description for the `exclude_interim` query parameter.
// API name: exclude_interim
func (r *GetRecords) ExcludeInterim(excludeinterim bool) *GetRecords {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ExcludeInterim = &excludeinterim

	return r
}

// API name: page
func (r *GetRecords) Page(page *types.Page) *GetRecords {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Page = page

	return r
}

// RecordScore Refer to the description for the `record_score` query parameter.
// API name: record_score
func (r *GetRecords) RecordScore(recordscore types.Float64) *GetRecords {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RecordScore = &recordscore

	return r
}

// Sort Refer to the description for the `sort` query parameter.
// API name: sort
func (r *GetRecords) Sort(field string) *GetRecords {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Sort = &field

	return r
}

// Start Refer to the description for the `start` query parameter.
// API name: start
func (r *GetRecords) Start(datetime types.DateTime) *GetRecords {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Start = datetime

	return r
}
