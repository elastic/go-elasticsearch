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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Forces any buffered data to be processed by the job.
package flushjob

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	jobidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type FlushJob struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	jobid string
}

// NewFlushJob type alias for index.
type NewFlushJob func(jobid string) *FlushJob

// NewFlushJobFunc returns a new instance of FlushJob with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewFlushJobFunc(tp elastictransport.Interface) NewFlushJob {
	return func(jobid string) *FlushJob {
		n := New(tp)

		n.JobId(jobid)

		return n
	}
}

// Forces any buffered data to be processed by the job.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-flush-job.html
func New(tp elastictransport.Interface) *FlushJob {
	r := &FlushJob{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *FlushJob) Raw(raw json.RawMessage) *FlushJob {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *FlushJob) Request(req *Request) *FlushJob {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *FlushJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for FlushJob: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == jobidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")
		path.WriteString("/")

		path.WriteString(r.jobid)
		path.WriteString("/")
		path.WriteString("_flush")

		method = http.MethodPost
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
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

// Do runs the http.Request through the provided transport.
func (r FlushJob) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the FlushJob query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the FlushJob headers map.
func (r *FlushJob) Header(key, value string) *FlushJob {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the anomaly detection job.
// API Name: jobid
func (r *FlushJob) JobId(v string) *FlushJob {
	r.paramSet |= jobidMask
	r.jobid = v

	return r
}

// AdvanceTime Specifies to advance to a particular time value. Results are generated
// and the model is updated for data from the specified time interval.
// API name: advance_time
func (r *FlushJob) AdvanceTime(value string) *FlushJob {
	r.values.Set("advance_time", value)

	return r
}

// CalcInterim If true, calculates the interim results for the most recent bucket or all
// buckets within the latency period.
// API name: calc_interim
func (r *FlushJob) CalcInterim(b bool) *FlushJob {
	r.values.Set("calc_interim", strconv.FormatBool(b))

	return r
}

// End When used in conjunction with `calc_interim` and `start`, specifies the
// range of buckets on which to calculate interim results.
// API name: end
func (r *FlushJob) End(value string) *FlushJob {
	r.values.Set("end", value)

	return r
}

// SkipTime Specifies to skip to a particular time value. Results are not generated
// and the model is not updated for data from the specified time interval.
// API name: skip_time
func (r *FlushJob) SkipTime(value string) *FlushJob {
	r.values.Set("skip_time", value)

	return r
}

// Start When used in conjunction with `calc_interim`, specifies the range of
// buckets on which to calculate interim results.
// API name: start
func (r *FlushJob) Start(value string) *FlushJob {
	r.values.Set("start", value)

	return r
}
