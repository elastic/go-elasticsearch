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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Retrieves overall bucket results that summarize the bucket results of
// multiple anomaly detection jobs.
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	jobid string
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

// Retrieves overall bucket results that summarize the bucket results of
// multiple anomaly detection jobs.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-overall-buckets.html
func New(tp elastictransport.Interface) *GetOverallBuckets {
	r := &GetOverallBuckets{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for GetOverallBuckets: %w", err)
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

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r GetOverallBuckets) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetOverallBuckets query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getoverallbuckets.Response
func (r GetOverallBuckets) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
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

// AllowNoMatch Refer to the description for the `allow_no_match` query parameter.
// API name: allow_no_match
func (r *GetOverallBuckets) AllowNoMatch(allownomatch bool) *GetOverallBuckets {
	r.req.AllowNoMatch = &allownomatch

	return r
}

// BucketSpan Refer to the description for the `bucket_span` query parameter.
// API name: bucket_span
func (r *GetOverallBuckets) BucketSpan(duration types.Duration) *GetOverallBuckets {
	r.req.BucketSpan = duration

	return r
}

// End Refer to the description for the `end` query parameter.
// API name: end
func (r *GetOverallBuckets) End(datetime types.DateTime) *GetOverallBuckets {
	r.req.End = datetime

	return r
}

// ExcludeInterim Refer to the description for the `exclude_interim` query parameter.
// API name: exclude_interim
func (r *GetOverallBuckets) ExcludeInterim(excludeinterim bool) *GetOverallBuckets {
	r.req.ExcludeInterim = &excludeinterim

	return r
}

// OverallScore Refer to the description for the `overall_score` query parameter.
// API name: overall_score
func (r *GetOverallBuckets) OverallScore(overallscore string) *GetOverallBuckets {
	r.req.OverallScore = overallscore

	return r
}

// Start Refer to the description for the `start` query parameter.
// API name: start
func (r *GetOverallBuckets) Start(datetime types.DateTime) *GetOverallBuckets {
	r.req.Start = datetime

	return r
}

// TopN Refer to the description for the `top_n` query parameter.
// API name: top_n
func (r *GetOverallBuckets) TopN(topn int) *GetOverallBuckets {
	r.req.TopN = &topn

	return r
}
