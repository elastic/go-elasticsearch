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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Retrieves anomaly detection job results for one or more buckets.
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	jobid     string
	timestamp string
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

// Retrieves anomaly detection job results for one or more buckets.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-bucket.html
func New(tp elastictransport.Interface) *GetBuckets {
	r := &GetBuckets{
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for GetBuckets: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == jobidMask|timestampMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")
		path.WriteString("/")

		path.WriteString(r.jobid)
		path.WriteString("/")
		path.WriteString("results")
		path.WriteString("/")
		path.WriteString("buckets")
		path.WriteString("/")

		path.WriteString(r.timestamp)

		method = http.MethodPost
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
		path.WriteString("buckets")

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
func (r GetBuckets) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetBuckets query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getbuckets.Response
func (r GetBuckets) Do(ctx context.Context) (*Response, error) {

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

// AnomalyScore Refer to the description for the `anomaly_score` query parameter.
// API name: anomaly_score
func (r *GetBuckets) AnomalyScore(anomalyscore types.Float64) *GetBuckets {

	r.req.AnomalyScore = &anomalyscore

	return r
}

// Desc Refer to the description for the `desc` query parameter.
// API name: desc
func (r *GetBuckets) Desc(desc bool) *GetBuckets {
	r.req.Desc = &desc

	return r
}

// End Refer to the description for the `end` query parameter.
// API name: end
func (r *GetBuckets) End(datetime types.DateTime) *GetBuckets {
	r.req.End = datetime

	return r
}

// ExcludeInterim Refer to the description for the `exclude_interim` query parameter.
// API name: exclude_interim
func (r *GetBuckets) ExcludeInterim(excludeinterim bool) *GetBuckets {
	r.req.ExcludeInterim = &excludeinterim

	return r
}

// Expand Refer to the description for the `expand` query parameter.
// API name: expand
func (r *GetBuckets) Expand(expand bool) *GetBuckets {
	r.req.Expand = &expand

	return r
}

// API name: page
func (r *GetBuckets) Page(page *types.Page) *GetBuckets {

	r.req.Page = page

	return r
}

// Sort Refer to the desription for the `sort` query parameter.
// API name: sort
func (r *GetBuckets) Sort(field string) *GetBuckets {
	r.req.Sort = &field

	return r
}

// Start Refer to the description for the `start` query parameter.
// API name: start
func (r *GetBuckets) Start(datetime types.DateTime) *GetBuckets {
	r.req.Start = datetime

	return r
}
