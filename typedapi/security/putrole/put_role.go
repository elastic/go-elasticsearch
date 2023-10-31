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

// Adds and updates roles in the native realm.
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	name string
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

// Adds and updates roles in the native realm.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-role.html
func New(tp elastictransport.Interface) *PutRole {
	r := &PutRole{
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutRole: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("role")
		path.WriteString("/")

		path.WriteString(r.name)

		method = http.MethodPut
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
func (r PutRole) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutRole query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putrole.Response
func (r PutRole) Do(ctx context.Context) (*Response, error) {

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

// Applications A list of application privilege entries.
// API name: applications
func (r *PutRole) Applications(applications ...types.ApplicationPrivileges) *PutRole {
	r.req.Applications = applications

	return r
}

// Cluster A list of cluster privileges. These privileges define the cluster-level
// actions for users with this role.
// API name: cluster
func (r *PutRole) Cluster(clusters ...clusterprivilege.ClusterPrivilege) *PutRole {
	r.req.Cluster = clusters

	return r
}

// Global An object defining global privileges. A global privilege is a form of cluster
// privilege that is request-aware. Support for global privileges is currently
// limited to the management of application privileges.
// API name: global
func (r *PutRole) Global(global map[string]json.RawMessage) *PutRole {

	r.req.Global = global

	return r
}

// Indices A list of indices permissions entries.
// API name: indices
func (r *PutRole) Indices(indices ...types.IndicesPrivileges) *PutRole {
	r.req.Indices = indices

	return r
}

// Metadata Optional metadata. Within the metadata object, keys that begin with an
// underscore (`_`) are reserved for system use.
// API name: metadata
func (r *PutRole) Metadata(metadata types.Metadata) *PutRole {
	r.req.Metadata = metadata

	return r
}

// RunAs A list of users that the owners of this role can impersonate.
// API name: run_as
func (r *PutRole) RunAs(runas ...string) *PutRole {
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
func (r *PutRole) TransientMetadata(transientmetadata *types.TransientMetadataConfig) *PutRole {

	r.req.TransientMetadata = transientmetadata

	return r
}
