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

// Create a rollup job.
//
// WARNING: From 8.15.0, calling this API in a cluster with no rollup usage will
// fail with a message about the deprecation and planned removal of rollup
// features. A cluster needs to contain either a rollup job or a rollup index in
// order for this API to be allowed to run.
//
// The rollup job configuration contains all the details about how the job
// should run, when it indexes documents, and what future queries will be able
// to run against the rollup index.
//
// There are three main sections to the job configuration: the logistical
// details about the job (for example, the cron schedule), the fields that are
// used for grouping, and what metrics to collect for each group.
//
// Jobs are created in a `STOPPED` state. You can start them with the start
// rollup jobs API.
package putjob

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

type PutJob struct {
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

// NewPutJob type alias for index.
type NewPutJob func(id string) *PutJob

// NewPutJobFunc returns a new instance of PutJob with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutJobFunc(tp elastictransport.Interface) NewPutJob {
	return func(id string) *PutJob {
		n := New(tp)

		n._id(id)

		return n
	}
}

// Create a rollup job.
//
// WARNING: From 8.15.0, calling this API in a cluster with no rollup usage will
// fail with a message about the deprecation and planned removal of rollup
// features. A cluster needs to contain either a rollup job or a rollup index in
// order for this API to be allowed to run.
//
// The rollup job configuration contains all the details about how the job
// should run, when it indexes documents, and what future queries will be able
// to run against the rollup index.
//
// There are three main sections to the job configuration: the logistical
// details about the job (for example, the cron schedule), the fields that are
// used for grouping, and what metrics to collect for each group.
//
// Jobs are created in a `STOPPED` state. You can start them with the start
// rollup jobs API.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-put-job.html
func New(tp elastictransport.Interface) *PutJob {
	r := &PutJob{
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
func (r *PutJob) Raw(raw io.Reader) *PutJob {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutJob) Request(req *Request) *PutJob {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutJob) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutJob: %w", err)
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
		path.WriteString("_rollup")
		path.WriteString("/")
		path.WriteString("job")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)

		method = http.MethodPut
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
func (r PutJob) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "rollup.put_job")
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
		instrument.BeforeRequest(req, "rollup.put_job")
		if reader := instrument.RecordRequestBody(ctx, "rollup.put_job", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "rollup.put_job")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutJob query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putjob.Response
func (r PutJob) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "rollup.put_job")
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

// Header set a key, value pair in the PutJob headers map.
func (r *PutJob) Header(key, value string) *PutJob {
	r.headers.Set(key, value)

	return r
}

// Id Identifier for the rollup job. This can be any alphanumeric string and
// uniquely identifies the
// data that is associated with the rollup job. The ID is persistent; it is
// stored with the rolled
// up data. If you create a job, let it run for a while, then delete the job,
// the data that the job
// rolled up is still be associated with this job ID. You cannot create a new
// job with the same ID
// since that could lead to problems with mismatched job configurations.
// API Name: id
func (r *PutJob) _id(id string) *PutJob {
	r.paramSet |= idMask
	r.id = id

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutJob) ErrorTrace(errortrace bool) *PutJob {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutJob) FilterPath(filterpaths ...string) *PutJob {
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
func (r *PutJob) Human(human bool) *PutJob {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutJob) Pretty(pretty bool) *PutJob {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Cron A cron string which defines the intervals when the rollup job should be
// executed. When the interval
// triggers, the indexer attempts to rollup the data in the index pattern. The
// cron pattern is unrelated
// to the time interval of the data being rolled up. For example, you may wish
// to create hourly rollups
// of your document but to only run the indexer on a daily basis at midnight, as
// defined by the cron. The
// cron pattern is defined just like a Watcher cron schedule.
// API name: cron
func (r *PutJob) Cron(cron string) *PutJob {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Cron = cron

	return r
}

// Groups Defines the grouping fields and aggregations that are defined for this rollup
// job. These fields will then be
// available later for aggregating into buckets. These aggs and fields can be
// used in any combination. Think of
// the groups configuration as defining a set of tools that can later be used in
// aggregations to partition the
// data. Unlike raw data, we have to think ahead to which fields and
// aggregations might be used. Rollups provide
// enough flexibility that you simply need to determine which fields are needed,
// not in what order they are needed.
// API name: groups
func (r *PutJob) Groups(groups *types.Groupings) *PutJob {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Groups = *groups

	return r
}

// API name: headers
func (r *PutJob) Headers(httpheaders types.HttpHeaders) *PutJob {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Headers = httpheaders

	return r
}

// IndexPattern The index or index pattern to roll up. Supports wildcard-style patterns
// (`logstash-*`). The job attempts to
// rollup the entire index or index-pattern.
// API name: index_pattern
func (r *PutJob) IndexPattern(indexpattern string) *PutJob {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexPattern = indexpattern

	return r
}

// Metrics Defines the metrics to collect for each grouping tuple. By default, only the
// doc_counts are collected for each
// group. To make rollup useful, you will often add metrics like averages, mins,
// maxes, etc. Metrics are defined
// on a per-field basis and for each field you configure which metric should be
// collected.
// API name: metrics
func (r *PutJob) Metrics(metrics ...types.FieldMetric) *PutJob {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Metrics = metrics

	return r
}

// PageSize The number of bucket results that are processed on each iteration of the
// rollup indexer. A larger value tends
// to execute faster, but requires more memory during processing. This value has
// no effect on how the data is
// rolled up; it is merely used for tweaking the speed or memory cost of the
// indexer.
// API name: page_size
func (r *PutJob) PageSize(pagesize int) *PutJob {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.PageSize = pagesize

	return r
}

// RollupIndex The index that contains the rollup results. The index can be shared with
// other rollup jobs. The data is stored so that it doesn’t interfere with
// unrelated jobs.
// API name: rollup_index
func (r *PutJob) RollupIndex(indexname string) *PutJob {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RollupIndex = indexname

	return r
}

// Timeout Time to wait for the request to complete.
// API name: timeout
func (r *PutJob) Timeout(duration types.Duration) *PutJob {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Timeout = duration

	return r
}
