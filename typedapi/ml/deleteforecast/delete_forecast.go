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

// Deletes forecasts from a machine learning job.
package deleteforecast

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

	forecastidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteForecast struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	jobid      string
	forecastid string
}

// NewDeleteForecast type alias for index.
type NewDeleteForecast func(jobid string) *DeleteForecast

// NewDeleteForecastFunc returns a new instance of DeleteForecast with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteForecastFunc(tp elastictransport.Interface) NewDeleteForecast {
	return func(jobid string) *DeleteForecast {
		n := New(tp)

		n._jobid(jobid)

		return n
	}
}

// Deletes forecasts from a machine learning job.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-forecast.html
func New(tp elastictransport.Interface) *DeleteForecast {
	r := &DeleteForecast{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *DeleteForecast) HttpRequest(ctx context.Context) (*http.Request, error) {
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
		path.WriteString("/")
		path.WriteString("_forecast")

		method = http.MethodDelete
	case r.paramSet == jobidMask|forecastidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")
		path.WriteString("/")

		path.WriteString(r.jobid)
		path.WriteString("/")
		path.WriteString("_forecast")
		path.WriteString("/")

		path.WriteString(r.forecastid)

		method = http.MethodDelete
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
func (r DeleteForecast) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the DeleteForecast query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a deleteforecast.Response
func (r DeleteForecast) Do(ctx context.Context) (*Response, error) {

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
func (r DeleteForecast) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the DeleteForecast headers map.
func (r *DeleteForecast) Header(key, value string) *DeleteForecast {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the anomaly detection job.
// API Name: jobid
func (r *DeleteForecast) _jobid(jobid string) *DeleteForecast {
	r.paramSet |= jobidMask
	r.jobid = jobid

	return r
}

// ForecastId A comma-separated list of forecast identifiers. If you do not specify
// this optional parameter or if you specify `_all` or `*` the API deletes
// all forecasts from the job.
// API Name: forecastid
func (r *DeleteForecast) ForecastId(forecastid string) *DeleteForecast {
	r.paramSet |= forecastidMask
	r.forecastid = forecastid

	return r
}

// AllowNoForecasts Specifies whether an error occurs when there are no forecasts. In
// particular, if this parameter is set to `false` and there are no
// forecasts associated with the job, attempts to delete all forecasts
// return an error.
// API name: allow_no_forecasts
func (r *DeleteForecast) AllowNoForecasts(allownoforecasts bool) *DeleteForecast {
	r.values.Set("allow_no_forecasts", strconv.FormatBool(allownoforecasts))

	return r
}

// Timeout Specifies the period of time to wait for the completion of the delete
// operation. When this period of time elapses, the API fails and returns an
// error.
// API name: timeout
func (r *DeleteForecast) Timeout(duration string) *DeleteForecast {
	r.values.Set("timeout", duration)

	return r
}
