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

// Update the cluster settings.
//
// Configure and update dynamic settings on a running cluster.
// You can also configure dynamic settings locally on an unstarted or shut down
// node in `elasticsearch.yml`.
//
// Updates made with this API can be persistent, which apply across cluster
// restarts, or transient, which reset after a cluster restart.
// You can also reset transient or persistent settings by assigning them a null
// value.
//
// If you configure the same setting using multiple methods, Elasticsearch
// applies the settings in following order of precedence: 1) Transient setting;
// 2) Persistent setting; 3) `elasticsearch.yml` setting; 4) Default setting
// value.
// For example, you can apply a transient setting to override a persistent
// setting or `elasticsearch.yml` setting.
// However, a change to an `elasticsearch.yml` setting will not override a
// defined transient or persistent setting.
//
// TIP: In Elastic Cloud, use the user settings feature to configure all cluster
// settings. This method automatically rejects unsafe settings that could break
// your cluster.
// If you run Elasticsearch on your own hardware, use this API to configure
// dynamic cluster settings.
// Only use `elasticsearch.yml` for static cluster settings and node settings.
// The API doesn’t require a restart and ensures a setting’s value is the same
// on all nodes.
//
// WARNING: Transient cluster settings are no longer recommended. Use persistent
// cluster settings instead.
// If a cluster becomes unstable, transient settings can clear unexpectedly,
// resulting in a potentially undesired cluster configuration.
package putsettings

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

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutSettings struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutSettings type alias for index.
type NewPutSettings func() *PutSettings

// NewPutSettingsFunc returns a new instance of PutSettings with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutSettingsFunc(tp elastictransport.Interface) NewPutSettings {
	return func() *PutSettings {
		n := New(tp)

		return n
	}
}

// Update the cluster settings.
//
// Configure and update dynamic settings on a running cluster.
// You can also configure dynamic settings locally on an unstarted or shut down
// node in `elasticsearch.yml`.
//
// Updates made with this API can be persistent, which apply across cluster
// restarts, or transient, which reset after a cluster restart.
// You can also reset transient or persistent settings by assigning them a null
// value.
//
// If you configure the same setting using multiple methods, Elasticsearch
// applies the settings in following order of precedence: 1) Transient setting;
// 2) Persistent setting; 3) `elasticsearch.yml` setting; 4) Default setting
// value.
// For example, you can apply a transient setting to override a persistent
// setting or `elasticsearch.yml` setting.
// However, a change to an `elasticsearch.yml` setting will not override a
// defined transient or persistent setting.
//
// TIP: In Elastic Cloud, use the user settings feature to configure all cluster
// settings. This method automatically rejects unsafe settings that could break
// your cluster.
// If you run Elasticsearch on your own hardware, use this API to configure
// dynamic cluster settings.
// Only use `elasticsearch.yml` for static cluster settings and node settings.
// The API doesn’t require a restart and ensures a setting’s value is the same
// on all nodes.
//
// WARNING: Transient cluster settings are no longer recommended. Use persistent
// cluster settings instead.
// If a cluster becomes unstable, transient settings can clear unexpectedly,
// resulting in a potentially undesired cluster configuration.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-update-settings.html
func New(tp elastictransport.Interface) *PutSettings {
	r := &PutSettings{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *PutSettings) Raw(raw io.Reader) *PutSettings {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutSettings) Request(req *Request) *PutSettings {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutSettings) HttpRequest(ctx context.Context) (*http.Request, error) {
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

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutSettings: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cluster")
		path.WriteString("/")
		path.WriteString("settings")

		method = http.MethodPut
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

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
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
func (r PutSettings) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cluster.put_settings")
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
		instrument.BeforeRequest(req, "cluster.put_settings")
		if reader := instrument.RecordRequestBody(ctx, "cluster.put_settings", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cluster.put_settings")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutSettings query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putsettings.Response
func (r PutSettings) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.put_settings")
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

// Header set a key, value pair in the PutSettings headers map.
func (r *PutSettings) Header(key, value string) *PutSettings {
	r.headers.Set(key, value)

	return r
}

// FlatSettings Return settings in flat format (default: false)
// API name: flat_settings
func (r *PutSettings) FlatSettings(flatsettings bool) *PutSettings {
	r.values.Set("flat_settings", strconv.FormatBool(flatsettings))

	return r
}

// MasterTimeout Explicit operation timeout for connection to master node
// API name: master_timeout
func (r *PutSettings) MasterTimeout(duration string) *PutSettings {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *PutSettings) Timeout(duration string) *PutSettings {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutSettings) ErrorTrace(errortrace bool) *PutSettings {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutSettings) FilterPath(filterpaths ...string) *PutSettings {
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
func (r *PutSettings) Human(human bool) *PutSettings {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutSettings) Pretty(pretty bool) *PutSettings {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Persistent The settings that persist after the cluster restarts.
// API name: persistent
func (r *PutSettings) Persistent(persistent map[string]json.RawMessage) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Persistent = persistent

	return r
}

// Transient The settings that do not persist after the cluster restarts.
// API name: transient
func (r *PutSettings) Transient(transient map[string]json.RawMessage) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Transient = transient

	return r
}
