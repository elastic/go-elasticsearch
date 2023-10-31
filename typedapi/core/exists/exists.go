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

// Returns information about whether a document exists in an index.
package exists

import (
	gobytes "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

const (
	idMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Exists struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	id    string
	index string
}

// NewExists type alias for index.
type NewExists func(index, id string) *Exists

// NewExistsFunc returns a new instance of Exists with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewExistsFunc(tp elastictransport.Interface) NewExists {
	return func(index, id string) *Exists {
		n := New(tp)

		n._id(id)

		n._index(index)

		return n
	}
}

// Returns information about whether a document exists in an index.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-get.html
func New(tp elastictransport.Interface) *Exists {
	r := &Exists{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Exists) HttpRequest(ctx context.Context) (*http.Request, error) {
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
		path.WriteString("_doc")
		path.WriteString("/")

		path.WriteString(r.id)

		method = http.MethodHead
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
func (r Exists) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Exists query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a exists.Response
func (r Exists) Do(ctx context.Context) (bool, error) {
	return r.IsSuccess(ctx)
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r Exists) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the Exists headers map.
func (r *Exists) Header(key, value string) *Exists {
	r.headers.Set(key, value)

	return r
}

// Id Identifier of the document.
// API Name: id
func (r *Exists) _id(id string) *Exists {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Index Comma-separated list of data streams, indices, and aliases.
// Supports wildcards (`*`).
// API Name: index
func (r *Exists) _index(index string) *Exists {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Preference Specifies the node or shard the operation should be performed on.
// Random by default.
// API name: preference
func (r *Exists) Preference(preference string) *Exists {
	r.values.Set("preference", preference)

	return r
}

// Realtime If `true`, the request is real-time as opposed to near-real-time.
// API name: realtime
func (r *Exists) Realtime(realtime bool) *Exists {
	r.values.Set("realtime", strconv.FormatBool(realtime))

	return r
}

// Refresh If `true`, Elasticsearch refreshes all shards involved in the delete by query
// after the request completes.
// API name: refresh
func (r *Exists) Refresh(refresh bool) *Exists {
	r.values.Set("refresh", strconv.FormatBool(refresh))

	return r
}

// Routing Target the specified primary shard.
// API name: routing
func (r *Exists) Routing(routing string) *Exists {
	r.values.Set("routing", routing)

	return r
}

// Source_ `true` or `false` to return the `_source` field or not, or a list of fields
// to return.
// API name: _source
func (r *Exists) Source_(sourceconfigparam string) *Exists {
	r.values.Set("_source", sourceconfigparam)

	return r
}

// SourceExcludes_ A comma-separated list of source fields to exclude in the response.
// API name: _source_excludes
func (r *Exists) SourceExcludes_(fields ...string) *Exists {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ A comma-separated list of source fields to include in the response.
// API name: _source_includes
func (r *Exists) SourceIncludes_(fields ...string) *Exists {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// StoredFields List of stored fields to return as part of a hit.
// If no fields are specified, no stored fields are included in the response.
// If this field is specified, the `_source` parameter defaults to false.
// API name: stored_fields
func (r *Exists) StoredFields(fields ...string) *Exists {
	r.values.Set("stored_fields", strings.Join(fields, ","))

	return r
}

// Version Explicit version number for concurrency control.
// The specified version must match the current version of the document for the
// request to succeed.
// API name: version
func (r *Exists) Version(versionnumber string) *Exists {
	r.values.Set("version", versionnumber)

	return r
}

// VersionType Specific version type: `external`, `external_gte`.
// API name: version_type
func (r *Exists) VersionType(versiontype versiontype.VersionType) *Exists {
	r.values.Set("version_type", versiontype.String())

	return r
}
