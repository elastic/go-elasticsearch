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

// Reload the keystore on nodes in the cluster.
//
// Secure settings are stored in an on-disk keystore. Certain of these settings
// are reloadable.
// That is, you can change them on disk and reload them without restarting any
// nodes in the cluster.
// When you have updated reloadable secure settings in your keystore, you can
// use this API to reload those settings on each node.
//
// When the Elasticsearch keystore is password protected and not simply
// obfuscated, you must provide the password for the keystore when you reload
// the secure settings.
// Reloading the settings for the whole cluster assumes that the keystores for
// all nodes are protected with the same password; this method is allowed only
// when inter-node communications are encrypted.
// Alternatively, you can reload the secure settings on each node by locally
// accessing the API and passing the node-specific Elasticsearch keystore
// password.
package reloadsecuresettings

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
	nodeidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ReloadSecureSettings struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	nodeid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewReloadSecureSettings type alias for index.
type NewReloadSecureSettings func() *ReloadSecureSettings

// NewReloadSecureSettingsFunc returns a new instance of ReloadSecureSettings with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewReloadSecureSettingsFunc(tp elastictransport.Interface) NewReloadSecureSettings {
	return func() *ReloadSecureSettings {
		n := New(tp)

		return n
	}
}

// Reload the keystore on nodes in the cluster.
//
// Secure settings are stored in an on-disk keystore. Certain of these settings
// are reloadable.
// That is, you can change them on disk and reload them without restarting any
// nodes in the cluster.
// When you have updated reloadable secure settings in your keystore, you can
// use this API to reload those settings on each node.
//
// When the Elasticsearch keystore is password protected and not simply
// obfuscated, you must provide the password for the keystore when you reload
// the secure settings.
// Reloading the settings for the whole cluster assumes that the keystores for
// all nodes are protected with the same password; this method is allowed only
// when inter-node communications are encrypted.
// Alternatively, you can reload the secure settings on each node by locally
// accessing the API and passing the node-specific Elasticsearch keystore
// password.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-reload-secure-settings.html
func New(tp elastictransport.Interface) *ReloadSecureSettings {
	r := &ReloadSecureSettings{
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
func (r *ReloadSecureSettings) Raw(raw io.Reader) *ReloadSecureSettings {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *ReloadSecureSettings) Request(req *Request) *ReloadSecureSettings {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ReloadSecureSettings) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for ReloadSecureSettings: %w", err)
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
		path.WriteString("_nodes")
		path.WriteString("/")
		path.WriteString("reload_secure_settings")

		method = http.MethodPost
	case r.paramSet == nodeidMask:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "nodeid", r.nodeid)
		}
		path.WriteString(r.nodeid)
		path.WriteString("/")
		path.WriteString("reload_secure_settings")

		method = http.MethodPost
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
func (r ReloadSecureSettings) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "nodes.reload_secure_settings")
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
		instrument.BeforeRequest(req, "nodes.reload_secure_settings")
		if reader := instrument.RecordRequestBody(ctx, "nodes.reload_secure_settings", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "nodes.reload_secure_settings")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ReloadSecureSettings query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a reloadsecuresettings.Response
func (r ReloadSecureSettings) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "nodes.reload_secure_settings")
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

// Header set a key, value pair in the ReloadSecureSettings headers map.
func (r *ReloadSecureSettings) Header(key, value string) *ReloadSecureSettings {
	r.headers.Set(key, value)

	return r
}

// NodeId The names of particular nodes in the cluster to target.
// API Name: nodeid
func (r *ReloadSecureSettings) NodeId(nodeid string) *ReloadSecureSettings {
	r.paramSet |= nodeidMask
	r.nodeid = nodeid

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *ReloadSecureSettings) Timeout(duration string) *ReloadSecureSettings {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *ReloadSecureSettings) ErrorTrace(errortrace bool) *ReloadSecureSettings {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *ReloadSecureSettings) FilterPath(filterpaths ...string) *ReloadSecureSettings {
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
func (r *ReloadSecureSettings) Human(human bool) *ReloadSecureSettings {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *ReloadSecureSettings) Pretty(pretty bool) *ReloadSecureSettings {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// SecureSettingsPassword The password for the Elasticsearch keystore.
// API name: secure_settings_password
func (r *ReloadSecureSettings) SecureSettingsPassword(password string) *ReloadSecureSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.SecureSettingsPassword = &password

	return r
}
