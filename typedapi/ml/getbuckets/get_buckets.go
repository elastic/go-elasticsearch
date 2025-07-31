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

// Get anomaly detection job results for buckets.
// The API presents a chronological view of the records, grouped by bucket.
package getbuckets

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

	timestampMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetBuckets struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	jobid     string
	timestamp string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewGetBuckets type alias for index.
type NewGetBuckets func(jobid string) *GetBuckets

// NewGetBucketsFunc returns a new instance of GetBuckets with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetBucketsFunc(tp elastictransport.Interface) NewGetBuckets {
	return func(jobid string) *GetBuckets {
		n := New(tp)

		n._jobid(jobid)

		return n
	}
}

// Get anomaly detection job results for buckets.
// The API presents a chronological view of the records, grouped by bucket.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-bucket.html
func New(tp elastictransport.Interface) *GetBuckets {
	r := &GetBuckets{
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
func (r *GetBuckets) Raw(raw io.Reader) *GetBuckets {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *GetBuckets) Request(req *Request) *GetBuckets {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetBuckets) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for GetBuckets: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == jobidMask|timestampMask:
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
		path.WriteString("buckets")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "timestamp", r.timestamp)
		}
		path.WriteString(r.timestamp)

		method = http.MethodPost
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
		path.WriteString("buckets")

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
func (r GetBuckets) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.get_buckets")
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
		instrument.BeforeRequest(req, "ml.get_buckets")
		if reader := instrument.RecordRequestBody(ctx, "ml.get_buckets", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.get_buckets")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GetBuckets query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getbuckets.Response
func (r GetBuckets) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_buckets")
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

// Header set a key, value pair in the GetBuckets headers map.
func (r *GetBuckets) Header(key, value string) *GetBuckets {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the anomaly detection job.
// API Name: jobid
func (r *GetBuckets) _jobid(jobid string) *GetBuckets {
	r.paramSet |= jobidMask
	r.jobid = jobid

	return r
}

// Timestamp The timestamp of a single bucket result. If you do not specify this
// parameter, the API returns information about all buckets.
// API Name: timestamp
func (r *GetBuckets) Timestamp(timestamp string) *GetBuckets {
	r.paramSet |= timestampMask
	r.timestamp = timestamp

	return r
}

// From Skips the specified number of buckets.
// API name: from
func (r *GetBuckets) From(from int) *GetBuckets {
	r.values.Set("from", strconv.Itoa(from))

	return r
}

// Size Specifies the maximum number of buckets to obtain.
// API name: size
func (r *GetBuckets) Size(size int) *GetBuckets {
	r.values.Set("size", strconv.Itoa(size))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *GetBuckets) ErrorTrace(errortrace bool) *GetBuckets {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *GetBuckets) FilterPath(filterpaths ...string) *GetBuckets {
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
func (r *GetBuckets) Human(human bool) *GetBuckets {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *GetBuckets) Pretty(pretty bool) *GetBuckets {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// AnomalyScore Refer to the description for the `anomaly_score` query parameter.
// API name: anomaly_score
func (r *GetBuckets) AnomalyScore(anomalyscore types.Float64) *GetBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AnomalyScore = &anomalyscore

	return r
}

// Desc Refer to the description for the `desc` query parameter.
// API name: desc
func (r *GetBuckets) Desc(desc bool) *GetBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Desc = &desc

	return r
}

// End Refer to the description for the `end` query parameter.
// API name: end
func (r *GetBuckets) End(datetime types.DateTime) *GetBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.End = datetime

	return r
}

// ExcludeInterim Refer to the description for the `exclude_interim` query parameter.
// API name: exclude_interim
func (r *GetBuckets) ExcludeInterim(excludeinterim bool) *GetBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ExcludeInterim = &excludeinterim

	return r
}

// Expand Refer to the description for the `expand` query parameter.
// API name: expand
func (r *GetBuckets) Expand(expand bool) *GetBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Expand = &expand

	return r
}

// API name: page
func (r *GetBuckets) Page(page *types.Page) *GetBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Page = page

	return r
}

// Sort Refer to the desription for the `sort` query parameter.
// API name: sort
func (r *GetBuckets) Sort(field string) *GetBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Sort = &field

	return r
}

// Start Refer to the description for the `start` query parameter.
// API name: start
func (r *GetBuckets) Start(datetime types.DateTime) *GetBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Start = datetime

	return r
}
