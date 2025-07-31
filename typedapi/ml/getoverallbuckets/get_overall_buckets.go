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

// Get overall bucket results.
//
// Retrievs overall bucket results that summarize the bucket results of
// multiple anomaly detection jobs.
//
// The `overall_score` is calculated by combining the scores of all the
// buckets within the overall bucket span. First, the maximum
// `anomaly_score` per anomaly detection job in the overall bucket is
// calculated. Then the `top_n` of those scores are averaged to result in
// the `overall_score`. This means that you can fine-tune the
// `overall_score` so that it is more or less sensitive to the number of
// jobs that detect an anomaly at the same time. For example, if you set
// `top_n` to `1`, the `overall_score` is the maximum bucket score in the
// overall bucket. Alternatively, if you set `top_n` to the number of jobs,
// the `overall_score` is high only when all jobs detect anomalies in that
// overall bucket. If you set the `bucket_span` parameter (to a value
// greater than its default), the `overall_score` is the maximum
// `overall_score` of the overall buckets that have a span equal to the
// jobs' largest bucket span.
package getoverallbuckets

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

type GetOverallBuckets struct {
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

// NewGetOverallBuckets type alias for index.
type NewGetOverallBuckets func(jobid string) *GetOverallBuckets

// NewGetOverallBucketsFunc returns a new instance of GetOverallBuckets with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetOverallBucketsFunc(tp elastictransport.Interface) NewGetOverallBuckets {
	return func(jobid string) *GetOverallBuckets {
		n := New(tp)

		n._jobid(jobid)

		return n
	}
}

// Get overall bucket results.
//
// Retrievs overall bucket results that summarize the bucket results of
// multiple anomaly detection jobs.
//
// The `overall_score` is calculated by combining the scores of all the
// buckets within the overall bucket span. First, the maximum
// `anomaly_score` per anomaly detection job in the overall bucket is
// calculated. Then the `top_n` of those scores are averaged to result in
// the `overall_score`. This means that you can fine-tune the
// `overall_score` so that it is more or less sensitive to the number of
// jobs that detect an anomaly at the same time. For example, if you set
// `top_n` to `1`, the `overall_score` is the maximum bucket score in the
// overall bucket. Alternatively, if you set `top_n` to the number of jobs,
// the `overall_score` is high only when all jobs detect anomalies in that
// overall bucket. If you set the `bucket_span` parameter (to a value
// greater than its default), the `overall_score` is the maximum
// `overall_score` of the overall buckets that have a span equal to the
// jobs' largest bucket span.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-overall-buckets.html
func New(tp elastictransport.Interface) *GetOverallBuckets {
	r := &GetOverallBuckets{
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
func (r *GetOverallBuckets) Raw(raw io.Reader) *GetOverallBuckets {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *GetOverallBuckets) Request(req *Request) *GetOverallBuckets {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetOverallBuckets) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for GetOverallBuckets: %w", err)
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
		path.WriteString("overall_buckets")

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
func (r GetOverallBuckets) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.get_overall_buckets")
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
		instrument.BeforeRequest(req, "ml.get_overall_buckets")
		if reader := instrument.RecordRequestBody(ctx, "ml.get_overall_buckets", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.get_overall_buckets")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GetOverallBuckets query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getoverallbuckets.Response
func (r GetOverallBuckets) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_overall_buckets")
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

// Header set a key, value pair in the GetOverallBuckets headers map.
func (r *GetOverallBuckets) Header(key, value string) *GetOverallBuckets {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the anomaly detection job. It can be a job identifier, a
// group name, a comma-separated list of jobs or groups, or a wildcard
// expression.
//
// You can summarize the bucket results for all anomaly detection jobs by
// using `_all` or by specifying `*` as the `<job_id>`.
// API Name: jobid
func (r *GetOverallBuckets) _jobid(jobid string) *GetOverallBuckets {
	r.paramSet |= jobidMask
	r.jobid = jobid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *GetOverallBuckets) ErrorTrace(errortrace bool) *GetOverallBuckets {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *GetOverallBuckets) FilterPath(filterpaths ...string) *GetOverallBuckets {
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
func (r *GetOverallBuckets) Human(human bool) *GetOverallBuckets {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *GetOverallBuckets) Pretty(pretty bool) *GetOverallBuckets {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// AllowNoMatch Refer to the description for the `allow_no_match` query parameter.
// API name: allow_no_match
func (r *GetOverallBuckets) AllowNoMatch(allownomatch bool) *GetOverallBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.AllowNoMatch = &allownomatch

	return r
}

// BucketSpan Refer to the description for the `bucket_span` query parameter.
// API name: bucket_span
func (r *GetOverallBuckets) BucketSpan(duration types.Duration) *GetOverallBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.BucketSpan = duration

	return r
}

// End Refer to the description for the `end` query parameter.
// API name: end
func (r *GetOverallBuckets) End(datetime types.DateTime) *GetOverallBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.End = datetime

	return r
}

// ExcludeInterim Refer to the description for the `exclude_interim` query parameter.
// API name: exclude_interim
func (r *GetOverallBuckets) ExcludeInterim(excludeinterim bool) *GetOverallBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ExcludeInterim = &excludeinterim

	return r
}

// OverallScore Refer to the description for the `overall_score` query parameter.
// API name: overall_score
func (r *GetOverallBuckets) OverallScore(overallscore string) *GetOverallBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.OverallScore = &overallscore

	return r
}

// Start Refer to the description for the `start` query parameter.
// API name: start
func (r *GetOverallBuckets) Start(datetime types.DateTime) *GetOverallBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Start = datetime

	return r
}

// TopN Refer to the description for the `top_n` query parameter.
// API name: top_n
func (r *GetOverallBuckets) TopN(topn int) *GetOverallBuckets {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.TopN = &topn

	return r
}
