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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

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

	buf *gobytes.Buffer

	req *Request
	raw io.Reader

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

		n.JobId(jobid)

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
func (r *GetOverallBuckets) JobId(v string) *GetOverallBuckets {
	r.paramSet |= jobidMask
	r.jobid = v

	return r
}

// AllowNoMatch Specifies what to do when the request:
//
// 1. Contains wildcard expressions and there are no jobs that match.
// 2. Contains the `_all` string or no identifiers and there are no matches.
// 3. Contains wildcard expressions and there are only partial matches.
//
// If `true`, the request returns an empty `jobs` array when there are no
// matches and the subset of results when there are partial matches. If this
// parameter is `false`, the request returns a `404` status code when there
// are no matches or only partial matches.
// API name: allow_no_match
func (r *GetOverallBuckets) AllowNoMatch(b bool) *GetOverallBuckets {
	r.values.Set("allow_no_match", strconv.FormatBool(b))

	return r
}

// BucketSpan The span of the overall buckets. Must be greater or equal to the largest
// bucket span of the specified anomaly detection jobs, which is the default
// value.
//
// By default, an overall bucket has a span equal to the largest bucket span
// of the specified anomaly detection jobs. To override that behavior, use
// the optional `bucket_span` parameter.
// API name: bucket_span
func (r *GetOverallBuckets) BucketSpan(v string) *GetOverallBuckets {
	r.values.Set("bucket_span", v)

	return r
}

// End Returns overall buckets with timestamps earlier than this time.
// API name: end
func (r *GetOverallBuckets) End(v string) *GetOverallBuckets {
	r.values.Set("end", v)

	return r
}

// ExcludeInterim If `true`, the output excludes interim results.
// API name: exclude_interim
func (r *GetOverallBuckets) ExcludeInterim(b bool) *GetOverallBuckets {
	r.values.Set("exclude_interim", strconv.FormatBool(b))

	return r
}

// OverallScore Returns overall buckets with overall scores greater than or equal to this
// value.
// API name: overall_score
func (r *GetOverallBuckets) OverallScore(v string) *GetOverallBuckets {
	r.values.Set("overall_score", v)

	return r
}

// Start Returns overall buckets with timestamps after this time.
// API name: start
func (r *GetOverallBuckets) Start(v string) *GetOverallBuckets {
	r.values.Set("start", v)

	return r
}

// TopN The number of top anomaly detection job bucket scores to be used in the
// `overall_score` calculation.
// API name: top_n
func (r *GetOverallBuckets) TopN(i int) *GetOverallBuckets {
	r.values.Set("top_n", strconv.Itoa(i))

	return r
}
