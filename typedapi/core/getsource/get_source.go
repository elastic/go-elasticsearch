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

// Returns the source of a document.
package getsource

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

const (
	idMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetSource struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	id    string
	index string
}

// NewGetSource type alias for index.
type NewGetSource func(index, id string) *GetSource

// NewGetSourceFunc returns a new instance of GetSource with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetSourceFunc(tp elastictransport.Interface) NewGetSource {
	return func(index, id string) *GetSource {
		n := New(tp)

		n._id(id)

		n._index(index)

		return n
	}
}

// Returns the source of a document.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
func New(tp elastictransport.Interface) *GetSource {
	r := &GetSource{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetSource) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|idMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_source")
		path.WriteString("/")

		path.WriteString(r.id)

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
func (r GetSource) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetSource query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getsource.Response
func (r GetSource) Do(ctx context.Context) (Response, error) {

	response := new(Response)

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			return nil, err
		}

		return *response, nil
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
func (r GetSource) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the GetSource headers map.
func (r *GetSource) Header(key, value string) *GetSource {
	r.headers.Set(key, value)

	return r
}

// Id Unique identifier of the document.
// API Name: id
func (r *GetSource) _id(id string) *GetSource {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Index Name of the index that contains the document.
// API Name: index
func (r *GetSource) _index(index string) *GetSource {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Preference Specifies the node or shard the operation should be performed on. Random by
// default.
// API name: preference
func (r *GetSource) Preference(preference string) *GetSource {
	r.values.Set("preference", preference)

	return r
}

// Realtime Boolean) If true, the request is real-time as opposed to near-real-time.
// API name: realtime
func (r *GetSource) Realtime(realtime bool) *GetSource {
	r.values.Set("realtime", strconv.FormatBool(realtime))

	return r
}

// Refresh If true, Elasticsearch refreshes the affected shards to make this operation
// visible to search. If false, do nothing with refreshes.
// API name: refresh
func (r *GetSource) Refresh(refresh bool) *GetSource {
	r.values.Set("refresh", strconv.FormatBool(refresh))

	return r
}

// Routing Target the specified primary shard.
// API name: routing
func (r *GetSource) Routing(routing string) *GetSource {
	r.values.Set("routing", routing)

	return r
}

// Source_ True or false to return the _source field or not, or a list of fields to
// return.
// API name: _source
func (r *GetSource) Source_(sourceconfigparam string) *GetSource {
	r.values.Set("_source", sourceconfigparam)

	return r
}

// SourceExcludes_ A comma-separated list of source fields to exclude in the response.
// API name: _source_excludes
func (r *GetSource) SourceExcludes_(fields ...string) *GetSource {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ A comma-separated list of source fields to include in the response.
// API name: _source_includes
func (r *GetSource) SourceIncludes_(fields ...string) *GetSource {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// API name: stored_fields
func (r *GetSource) StoredFields(fields ...string) *GetSource {
	r.values.Set("stored_fields", strings.Join(fields, ","))

	return r
}

// Version Explicit version number for concurrency control. The specified version must
// match the current version of the document for the request to succeed.
// API name: version
func (r *GetSource) Version(versionnumber string) *GetSource {
	r.values.Set("version", versionnumber)

	return r
}

// VersionType Specific version type: internal, external, external_gte.
// API name: version_type
func (r *GetSource) VersionType(versiontype versiontype.VersionType) *GetSource {
	r.values.Set("version_type", versiontype.String())

	return r
}
