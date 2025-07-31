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

// Get component templates.
//
// Get information about component templates in a cluster.
// Component templates are building blocks for constructing index templates that
// specify index mappings, settings, and aliases.
//
// IMPORTANT: CAT APIs are only intended for human consumption using the command
// line or Kibana console.
// They are not intended for use by applications. For application consumption,
// use the get component template API.
package componenttemplates

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
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ComponentTemplates struct {
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

// NewComponentTemplates type alias for index.
type NewComponentTemplates func() *ComponentTemplates

// NewComponentTemplatesFunc returns a new instance of ComponentTemplates with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewComponentTemplatesFunc(tp elastictransport.Interface) NewComponentTemplates {
	return func() *ComponentTemplates {
		n := New(tp)

		return n
	}
}

// Get component templates.
//
// Get information about component templates in a cluster.
// Component templates are building blocks for constructing index templates that
// specify index mappings, settings, and aliases.
//
// IMPORTANT: CAT APIs are only intended for human consumption using the command
// line or Kibana console.
// They are not intended for use by applications. For application consumption,
// use the get component template API.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-component-templates.html
func New(tp elastictransport.Interface) *ComponentTemplates {
	r := &ComponentTemplates{
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
func (r *ComponentTemplates) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("component_templates")

		method = http.MethodGet
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("component_templates")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", r.name)
		}
		path.WriteString(r.name)

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
func (r ComponentTemplates) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cat.component_templates")
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
		instrument.BeforeRequest(req, "cat.component_templates")
		if reader := instrument.RecordRequestBody(ctx, "cat.component_templates", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cat.component_templates")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ComponentTemplates query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a componenttemplates.Response
func (r ComponentTemplates) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cat.component_templates")
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
		err = json.NewDecoder(res.Body).Decode(&response)
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
func (r ComponentTemplates) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cat.component_templates")
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
		err := fmt.Errorf("an error happened during the ComponentTemplates query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the ComponentTemplates headers map.
func (r *ComponentTemplates) Header(key, value string) *ComponentTemplates {
	r.headers.Set(key, value)

	return r
}

// Name The name of the component template.
// It accepts wildcard expressions.
// If it is omitted, all component templates are returned.
// API Name: name
func (r *ComponentTemplates) Name(name string) *ComponentTemplates {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// H List of columns to appear in the response. Supports simple wildcards.
// API name: h
func (r *ComponentTemplates) H(names ...string) *ComponentTemplates {
	r.values.Set("h", strings.Join(names, ","))

	return r
}

// S List of columns that determine how the table should be sorted.
// Sorting defaults to ascending and can be changed by setting `:asc`
// or `:desc` as a suffix to the column name.
// API name: s
func (r *ComponentTemplates) S(names ...string) *ComponentTemplates {
	r.values.Set("s", strings.Join(names, ","))

	return r
}

// Local If `true`, the request computes the list of selected nodes from the
// local cluster state. If `false` the list of selected nodes are computed
// from the cluster state of the master node. In both cases the coordinating
// node will send requests for further information to each selected node.
// API name: local
func (r *ComponentTemplates) Local(local bool) *ComponentTemplates {
	r.values.Set("local", strconv.FormatBool(local))

	return r
}

// MasterTimeout The period to wait for a connection to the master node.
// API name: master_timeout
func (r *ComponentTemplates) MasterTimeout(duration string) *ComponentTemplates {
	r.values.Set("master_timeout", duration)

	return r
}

// Format Specifies the format to return the columnar data in, can be set to
// `text`, `json`, `cbor`, `yaml`, or `smile`.
// API name: format
func (r *ComponentTemplates) Format(format string) *ComponentTemplates {
	r.values.Set("format", format)

	return r
}

// Help When set to `true` will output available columns. This option
// can't be combined with any other query string option.
// API name: help
func (r *ComponentTemplates) Help(help bool) *ComponentTemplates {
	r.values.Set("help", strconv.FormatBool(help))

	return r
}

// V When set to `true` will enable verbose output.
// API name: v
func (r *ComponentTemplates) V(v bool) *ComponentTemplates {
	r.values.Set("v", strconv.FormatBool(v))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *ComponentTemplates) ErrorTrace(errortrace bool) *ComponentTemplates {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *ComponentTemplates) FilterPath(filterpaths ...string) *ComponentTemplates {
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
func (r *ComponentTemplates) Human(human bool) *ComponentTemplates {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *ComponentTemplates) Pretty(pretty bool) *ComponentTemplates {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
