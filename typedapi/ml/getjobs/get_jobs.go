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

// Retrieves configuration information for anomaly detection jobs.
package getjobs

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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

type GetJobs struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	jobid string
}

// NewGetJobs type alias for index.
type NewGetJobs func() *GetJobs

// NewGetJobsFunc returns a new instance of GetJobs with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetJobsFunc(tp elastictransport.Interface) NewGetJobs {
	return func() *GetJobs {
		n := New(tp)

		return n
	}
}

// Retrieves configuration information for anomaly detection jobs.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job.html
func New(tp elastictransport.Interface) *GetJobs {
	r := &GetJobs{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetJobs) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == jobidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")
		path.WriteString("/")

		path.WriteString(r.jobid)

		method = http.MethodGet
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")

		method = http.MethodGet
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r GetJobs) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetJobs query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getjobs.Response
func (r GetJobs) Do(ctx context.Context) (*Response, error) {

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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetJobs) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	return false, nil
}

// Header set a key, value pair in the GetJobs headers map.
func (r *GetJobs) Header(key, value string) *GetJobs {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the anomaly detection job. It can be a job identifier, a
// group name, or a wildcard expression. If you do not specify one of these
// options, the API returns information for all anomaly detection jobs.
// API Name: jobid
func (r *GetJobs) JobId(jobid string) *GetJobs {
	r.paramSet |= jobidMask
	r.jobid = jobid

	return r
}

// AllowNoMatch Specifies what to do when the request:
//
// 1. Contains wildcard expressions and there are no jobs that match.
// 2. Contains the _all string or no identifiers and there are no matches.
// 3. Contains wildcard expressions and there are only partial matches.
//
// The default value is `true`, which returns an empty `jobs` array when
// there are no matches and the subset of results when there are partial
// matches. If this parameter is `false`, the request returns a `404` status
// code when there are no matches or only partial matches.
// API name: allow_no_match
func (r *GetJobs) AllowNoMatch(allownomatch bool) *GetJobs {
	r.values.Set("allow_no_match", strconv.FormatBool(allownomatch))

	return r
}

// ExcludeGenerated Indicates if certain fields should be removed from the configuration on
// retrieval. This allows the configuration to be in an acceptable format to
// be retrieved and then added to another cluster.
// API name: exclude_generated
func (r *GetJobs) ExcludeGenerated(excludegenerated bool) *GetJobs {
	r.values.Set("exclude_generated", strconv.FormatBool(excludegenerated))

	return r
}
