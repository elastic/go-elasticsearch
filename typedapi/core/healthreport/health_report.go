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

// Get the cluster health.
// Get a report with the health status of an Elasticsearch cluster.
// The report contains a list of indicators that compose Elasticsearch
// functionality.
//
// Each indicator has a health status of: green, unknown, yellow or red.
// The indicator will provide an explanation and metadata describing the reason
// for its current health status.
//
// The cluster’s status is controlled by the worst indicator status.
//
// In the event that an indicator’s status is non-green, a list of impacts may
// be present in the indicator result which detail the functionalities that are
// negatively affected by the health issue.
// Each impact carries with it a severity level, an area of the system that is
// affected, and a simple description of the impact on the system.
//
// Some health indicators can determine the root cause of a health problem and
// prescribe a set of steps that can be performed in order to improve the health
// of the system.
// The root cause and remediation steps are encapsulated in a diagnosis.
// A diagnosis contains a cause detailing a root cause analysis, an action
// containing a brief description of the steps to take to fix the problem, the
// list of affected resources (if applicable), and a detailed step-by-step
// troubleshooting guide to fix the diagnosed problem.
//
// NOTE: The health indicators perform root cause analysis of non-green health
// statuses. This can be computationally expensive when called frequently.
// When setting up automated polling of the API for health status, set verbose
// to false to disable the more expensive analysis logic.
package healthreport

import (
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
	featureMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type HealthReport struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	feature string

	spanStarted bool

	instrument elastictransport.Instrumentation
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

// Get the cluster health.
// Get a report with the health status of an Elasticsearch cluster.
// The report contains a list of indicators that compose Elasticsearch
// functionality.
//
// Each indicator has a health status of: green, unknown, yellow or red.
// The indicator will provide an explanation and metadata describing the reason
// for its current health status.
//
// The cluster’s status is controlled by the worst indicator status.
//
// In the event that an indicator’s status is non-green, a list of impacts may
// be present in the indicator result which detail the functionalities that are
// negatively affected by the health issue.
// Each impact carries with it a severity level, an area of the system that is
// affected, and a simple description of the impact on the system.
//
// Some health indicators can determine the root cause of a health problem and
// prescribe a set of steps that can be performed in order to improve the health
// of the system.
// The root cause and remediation steps are encapsulated in a diagnosis.
// A diagnosis contains a cause detailing a root cause analysis, an action
// containing a brief description of the steps to take to fix the problem, the
// list of affected resources (if applicable), and a detailed step-by-step
// troubleshooting guide to fix the diagnosed problem.
//
// NOTE: The health indicators perform root cause analysis of non-green health
// statuses. This can be computationally expensive when called frequently.
// When setting up automated polling of the API for health status, set verbose
// to false to disable the more expensive analysis logic.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/health-api.html
func New(tp elastictransport.Interface) *HealthReport {
	r := &HealthReport{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
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

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "feature", r.feature)
		}
		path.WriteString(r.feature)

		method = http.MethodGet
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r HealthReport) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "health_report")
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
		instrument.BeforeRequest(req, "health_report")
		if reader := instrument.RecordRequestBody(ctx, "health_report", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "health_report")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the HealthReport query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a healthreport.Response
func (r HealthReport) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "health_report")
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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r HealthReport) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "health_report")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(io.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the HealthReport query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
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

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *HealthReport) ErrorTrace(errortrace bool) *HealthReport {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *HealthReport) FilterPath(filterpaths ...string) *HealthReport {
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
func (r *HealthReport) Human(human bool) *HealthReport {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *HealthReport) Pretty(pretty bool) *HealthReport {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
