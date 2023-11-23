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

// Returns the health of the cluster.
package healthreport

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
	featureMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type HealthReport struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	feature string
}

// NewHealthReport type alias for index.
type NewHealthReport func() *HealthReport

// NewHealthReportFunc returns a new instance of HealthReport with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewHealthReportFunc(tp elastictransport.Interface) NewHealthReport {
	return func() *HealthReport {
		n := New(tp)

		return n
	}
}

// Returns the health of the cluster.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/health-api.html
func New(tp elastictransport.Interface) *HealthReport {
	r := &HealthReport{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *HealthReport) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_health_report")

		method = http.MethodGet
	case r.paramSet == featureMask:
		path.WriteString("/")
		path.WriteString("_health_report")
		path.WriteString("/")

		path.WriteString(r.feature)

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
func (r HealthReport) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the HealthReport query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a healthreport.Response
func (r HealthReport) Do(ctx context.Context) (*Response, error) {

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
func (r HealthReport) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the HealthReport headers map.
func (r *HealthReport) Header(key, value string) *HealthReport {
	r.headers.Set(key, value)

	return r
}

// Feature A feature of the cluster, as returned by the top-level health report API.
// API Name: feature
func (r *HealthReport) Feature(features ...string) *HealthReport {
	r.paramSet |= featureMask
	r.feature = strings.Join(features, ",")

	return r
}

// Timeout Explicit operation timeout.
// API name: timeout
func (r *HealthReport) Timeout(duration string) *HealthReport {
	r.values.Set("timeout", duration)

	return r
}

// Verbose Opt-in for more information about the health of the system.
// API name: verbose
func (r *HealthReport) Verbose(verbose bool) *HealthReport {
	r.values.Set("verbose", strconv.FormatBool(verbose))

	return r
}

// Size Limit the number of affected resources the health report API returns.
// API name: size
func (r *HealthReport) Size(size int) *HealthReport {
	r.values.Set("size", strconv.Itoa(size))

	return r
}
