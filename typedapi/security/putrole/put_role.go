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

// Create or update roles.
//
// The role management APIs are generally the preferred way to manage roles in
// the native realm, rather than using file-based role management.
// The create or update roles API cannot update roles that are defined in roles
// files.
// File-based role management is not available in Elastic Serverless.
package putrole

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/clusterprivilege"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

const (
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutRole struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutRole type alias for index.
type NewPutRole func(name string) *PutRole

// NewPutRoleFunc returns a new instance of PutRole with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutRoleFunc(tp elastictransport.Interface) NewPutRole {
	return func(name string) *PutRole {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Create or update roles.
//
// The role management APIs are generally the preferred way to manage roles in
// the native realm, rather than using file-based role management.
// The create or update roles API cannot update roles that are defined in roles
// files.
// File-based role management is not available in Elastic Serverless.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-role.html
func New(tp elastictransport.Interface) *PutRole {
	r := &PutRole{
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
func (r *PutRole) Raw(raw io.Reader) *PutRole {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutRole) Request(req *Request) *PutRole {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutRole) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutRole: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("role")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", r.name)
		}
		path.WriteString(r.name)

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
func (r PutRole) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.put_role")
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
		instrument.BeforeRequest(req, "security.put_role")
		if reader := instrument.RecordRequestBody(ctx, "security.put_role", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.put_role")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutRole query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putrole.Response
func (r PutRole) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.put_role")
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

// Header set a key, value pair in the PutRole headers map.
func (r *PutRole) Header(key, value string) *PutRole {
	r.headers.Set(key, value)

	return r
}

// Name The name of the role.
// API Name: name
func (r *PutRole) _name(name string) *PutRole {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// Refresh If `true` (the default) then refresh the affected shards to make this
// operation visible to search, if `wait_for` then wait for a refresh to make
// this operation visible to search, if `false` then do nothing with refreshes.
// API name: refresh
func (r *PutRole) Refresh(refresh refresh.Refresh) *PutRole {
	r.values.Set("refresh", refresh.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutRole) ErrorTrace(errortrace bool) *PutRole {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutRole) FilterPath(filterpaths ...string) *PutRole {
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
func (r *PutRole) Human(human bool) *PutRole {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutRole) Pretty(pretty bool) *PutRole {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Applications A list of application privilege entries.
// API name: applications
func (r *PutRole) Applications(applications ...types.ApplicationPrivileges) *PutRole {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Applications = applications

	return r
}

// Cluster A list of cluster privileges. These privileges define the cluster-level
// actions for users with this role.
// API name: cluster
func (r *PutRole) Cluster(clusters ...clusterprivilege.ClusterPrivilege) *PutRole {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Cluster = clusters

	return r
}

// Description Optional description of the role descriptor
// API name: description
func (r *PutRole) Description(description string) *PutRole {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Description = &description

	return r
}

// Global An object defining global privileges. A global privilege is a form of cluster
// privilege that is request-aware. Support for global privileges is currently
// limited to the management of application privileges.
// API name: global
func (r *PutRole) Global(global map[string]json.RawMessage) *PutRole {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Global = global

	return r
}

// Indices A list of indices permissions entries.
// API name: indices
func (r *PutRole) Indices(indices ...types.IndicesPrivileges) *PutRole {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Indices = indices

	return r
}

// Metadata Optional metadata. Within the metadata object, keys that begin with an
// underscore (`_`) are reserved for system use.
// API name: metadata
func (r *PutRole) Metadata(metadata types.Metadata) *PutRole {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Metadata = metadata

	return r
}

// RemoteCluster A list of remote cluster permissions entries.
// API name: remote_cluster
func (r *PutRole) RemoteCluster(remoteclusters ...types.RemoteClusterPrivileges) *PutRole {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RemoteCluster = remoteclusters

	return r
}

// RemoteIndices A list of remote indices permissions entries.
//
// NOTE: Remote indices are effective for remote clusters configured with the
// API key based model.
// They have no effect for remote clusters configured with the certificate based
// model.
// API name: remote_indices
func (r *PutRole) RemoteIndices(remoteindices ...types.RemoteIndicesPrivileges) *PutRole {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RemoteIndices = remoteindices

	return r
}

// RunAs A list of users that the owners of this role can impersonate. *Note*: in
// Serverless, the run-as feature is disabled. For API compatibility, you can
// still specify an empty `run_as` field, but a non-empty list will be rejected.
// API name: run_as
func (r *PutRole) RunAs(runas ...string) *PutRole {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RunAs = runas

	return r
}

// TransientMetadata Indicates roles that might be incompatible with the current cluster license,
// specifically roles with document and field level security. When the cluster
// license doesn’t allow certain features for a given role, this parameter is
// updated dynamically to list the incompatible features. If `enabled` is
// `false`, the role is ignored, but is still listed in the response from the
// authenticate API.
// API name: transient_metadata
func (r *PutRole) TransientMetadata(transientmetadata map[string]json.RawMessage) *PutRole {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TransientMetadata = transientmetadata

	return r
}
