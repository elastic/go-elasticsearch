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

	req *Request
	raw io.Reader

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

		n.JobId(jobid)

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

	return nil, errorResponse
}

// Header set a key, value pair in the GetBuckets headers map.
func (r *GetBuckets) Header(key, value string) *GetBuckets {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the anomaly detection job.
// API Name: jobid
func (r *GetBuckets) JobId(v string) *GetBuckets {
	r.paramSet |= jobidMask
	r.jobid = v

	return r
}

// Timestamp The timestamp of a single bucket result. If you do not specify this
// parameter, the API returns information about all buckets.
// API Name: timestamp
func (r *GetBuckets) Timestamp(v string) *GetBuckets {
	r.paramSet |= timestampMask
	r.timestamp = v

	return r
}

// AnomalyScore Returns buckets with anomaly scores greater or equal than this value.
// API name: anomaly_score
func (r *GetBuckets) AnomalyScore(v string) *GetBuckets {
	r.values.Set("anomaly_score", v)

	return r
}

// Desc If `true`, the buckets are sorted in descending order.
// API name: desc
func (r *GetBuckets) Desc(b bool) *GetBuckets {
	r.values.Set("desc", strconv.FormatBool(b))

	return r
}

// End Returns buckets with timestamps earlier than this time. `-1` means it is
// unset and results are not limited to specific timestamps.
// API name: end
func (r *GetBuckets) End(v string) *GetBuckets {
	r.values.Set("end", v)

	return r
}

// ExcludeInterim If `true`, the output excludes interim results.
// API name: exclude_interim
func (r *GetBuckets) ExcludeInterim(b bool) *GetBuckets {
	r.values.Set("exclude_interim", strconv.FormatBool(b))

	return r
}

// Expand If true, the output includes anomaly records.
// API name: expand
func (r *GetBuckets) Expand(b bool) *GetBuckets {
	r.values.Set("expand", strconv.FormatBool(b))

	return r
}

// From Skips the specified number of buckets.
// API name: from
func (r *GetBuckets) From(i int) *GetBuckets {
	r.values.Set("from", strconv.Itoa(i))

	return r
}

// Size Specifies the maximum number of buckets to obtain.
// API name: size
func (r *GetBuckets) Size(i int) *GetBuckets {
	r.values.Set("size", strconv.Itoa(i))

	return r
}

// Sort Specifies the sort field for the requested buckets.
// API name: sort
func (r *GetBuckets) Sort(v string) *GetBuckets {
	r.values.Set("sort", v)

	return r
}

// Start Returns buckets with timestamps after this time. `-1` means it is unset
// and results are not limited to specific timestamps.
// API name: start
func (r *GetBuckets) Start(v string) *GetBuckets {
	r.values.Set("start", v)

	return r
}
