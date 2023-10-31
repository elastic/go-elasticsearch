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

// Closes one or more anomaly detection jobs. A job can be opened and closed
// multiple times throughout its lifecycle.
package closejob

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

type CloseJob struct {
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

// NewCloseJob type alias for index.
type NewCloseJob func(jobid string) *CloseJob

// NewCloseJobFunc returns a new instance of CloseJob with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCloseJobFunc(tp elastictransport.Interface) NewCloseJob {
	return func(jobid string) *CloseJob {
		n := New(tp)

		n._jobid(jobid)

		return n
	}
}

// Closes one or more anomaly detection jobs. A job can be opened and closed
// multiple times throughout its lifecycle.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/ml-close-job.html
func New(tp elastictransport.Interface) *CloseJob {
	r := &CloseJob{
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
func (r *CloseJob) Raw(raw io.Reader) *CloseJob {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *CloseJob) Request(req *Request) *CloseJob {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *CloseJob) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for CloseJob: %w", err)
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
		path.WriteString("_close")

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
func (r CloseJob) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the CloseJob query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a closejob.Response
func (r CloseJob) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the CloseJob headers map.
func (r *CloseJob) Header(key, value string) *CloseJob {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the anomaly detection job. It can be a job identifier, a group
// name, or a wildcard expression. You can close multiple anomaly detection jobs
// in a single API request by using a group name, a comma-separated list of
// jobs, or a wildcard expression. You can close all jobs by using `_all` or by
// specifying `*` as the job identifier.
// API Name: jobid
func (r *CloseJob) _jobid(jobid string) *CloseJob {
	r.paramSet |= jobidMask
	r.jobid = jobid

	return r
}

// AllowNoMatch Refer to the description for the `allow_no_match` query parameter.
// API name: allow_no_match
func (r *CloseJob) AllowNoMatch(allownomatch bool) *CloseJob {
	r.req.AllowNoMatch = &allownomatch

	return r
}

// Force Refer to the descriptiion for the `force` query parameter.
// API name: force
func (r *CloseJob) Force(force bool) *CloseJob {
	r.req.Force = &force

	return r
}

// Timeout Refer to the description for the `timeout` query parameter.
// API name: timeout
func (r *CloseJob) Timeout(duration types.Duration) *CloseJob {
	r.req.Timeout = duration

	return r
}
