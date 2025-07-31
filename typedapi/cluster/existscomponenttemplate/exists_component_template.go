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

// Check component templates.
// Returns information about whether a particular component template exists.
package existscomponenttemplate

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ExistsComponentTemplate struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewExistsComponentTemplate type alias for index.
type NewExistsComponentTemplate func(name string) *ExistsComponentTemplate

// NewExistsComponentTemplateFunc returns a new instance of ExistsComponentTemplate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewExistsComponentTemplateFunc(tp elastictransport.Interface) NewExistsComponentTemplate {
	return func(name string) *ExistsComponentTemplate {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Check component templates.
// Returns information about whether a particular component template exists.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
func New(tp elastictransport.Interface) *ExistsComponentTemplate {
	r := &ExistsComponentTemplate{
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
func (r *ExistsComponentTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_component_template")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", r.name)
		}
		path.WriteString(r.name)

		method = http.MethodHead
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
func (r ExistsComponentTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cluster.exists_component_template")
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
		instrument.BeforeRequest(req, "cluster.exists_component_template")
		if reader := instrument.RecordRequestBody(ctx, "cluster.exists_component_template", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cluster.exists_component_template")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ExistsComponentTemplate query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a existscomponenttemplate.Response
func (r ExistsComponentTemplate) Do(ctx context.Context) (bool, error) {
	return r.IsSuccess(ctx)
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r ExistsComponentTemplate) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.exists_component_template")
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
		err := fmt.Errorf("an error happened during the ExistsComponentTemplate query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the ExistsComponentTemplate headers map.
func (r *ExistsComponentTemplate) Header(key, value string) *ExistsComponentTemplate {
	r.headers.Set(key, value)

	return r
}

// Name Comma-separated list of component template names used to limit the request.
// Wildcard (*) expressions are supported.
// API Name: name
func (r *ExistsComponentTemplate) _name(name string) *ExistsComponentTemplate {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an
// error.
// API name: master_timeout
func (r *ExistsComponentTemplate) MasterTimeout(duration string) *ExistsComponentTemplate {
	r.values.Set("master_timeout", duration)

	return r
}

// Local If true, the request retrieves information from the local node only.
// Defaults to false, which means information is retrieved from the master node.
// API name: local
func (r *ExistsComponentTemplate) Local(local bool) *ExistsComponentTemplate {
	r.values.Set("local", strconv.FormatBool(local))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *ExistsComponentTemplate) ErrorTrace(errortrace bool) *ExistsComponentTemplate {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *ExistsComponentTemplate) FilterPath(filterpaths ...string) *ExistsComponentTemplate {
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
func (r *ExistsComponentTemplate) Human(human bool) *ExistsComponentTemplate {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *ExistsComponentTemplate) Pretty(pretty bool) *ExistsComponentTemplate {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
