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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

// Returns information about whether a document source exists in an index.
package existssource

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

type ExistsSource struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	id    string
	index string
}

// NewExistsSource type alias for index.
type NewExistsSource func(index, id string) *ExistsSource

// NewExistsSourceFunc returns a new instance of ExistsSource with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewExistsSourceFunc(tp elastictransport.Interface) NewExistsSource {
	return func(index, id string) *ExistsSource {
		n := New(tp)

		n.Id(id)

		n.Index(index)

		return n
	}
}

// Returns information about whether a document source exists in an index.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-get.html
func New(tp elastictransport.Interface) *ExistsSource {
	r := &ExistsSource{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ExistsSource) HttpRequest(ctx context.Context) (*http.Request, error) {
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
func (r ExistsSource) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the ExistsSource query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r ExistsSource) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the ExistsSource headers map.
func (r *ExistsSource) Header(key, value string) *ExistsSource {
	r.headers.Set(key, value)

	return r
}

// Id The document ID
// API Name: id
func (r *ExistsSource) Id(v string) *ExistsSource {
	r.paramSet |= idMask
	r.id = v

	return r
}

// Index The name of the index
// API Name: index
func (r *ExistsSource) Index(v string) *ExistsSource {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// Preference Specify the node or shard the operation should be performed on (default:
// random)
// API name: preference
func (r *ExistsSource) Preference(v string) *ExistsSource {
	r.values.Set("preference", v)

	return r
}

// Realtime Specify whether to perform the operation in realtime or search mode
// API name: realtime
func (r *ExistsSource) Realtime(b bool) *ExistsSource {
	r.values.Set("realtime", strconv.FormatBool(b))

	return r
}

// Refresh Refresh the shard containing the document before performing the operation
// API name: refresh
func (r *ExistsSource) Refresh(b bool) *ExistsSource {
	r.values.Set("refresh", strconv.FormatBool(b))

	return r
}

// Routing Specific routing value
// API name: routing
func (r *ExistsSource) Routing(v string) *ExistsSource {
	r.values.Set("routing", v)

	return r
}

// Source_ True or false to return the _source field or not, or a list of fields to
// return
// API name: _source
func (r *ExistsSource) Source_(v string) *ExistsSource {
	r.values.Set("_source", v)

	return r
}

// SourceExcludes_ A list of fields to exclude from the returned _source field
// API name: _source_excludes
func (r *ExistsSource) SourceExcludes_(v string) *ExistsSource {
	r.values.Set("_source_excludes", v)

	return r
}

// SourceIncludes_ A list of fields to extract and return from the _source field
// API name: _source_includes
func (r *ExistsSource) SourceIncludes_(v string) *ExistsSource {
	r.values.Set("_source_includes", v)

	return r
}

// Version Explicit version number for concurrency control
// API name: version
func (r *ExistsSource) Version(v string) *ExistsSource {
	r.values.Set("version", v)

	return r
}

// VersionType Specific version type
// API name: version_type
func (r *ExistsSource) VersionType(enum versiontype.VersionType) *ExistsSource {
	r.values.Set("version_type", enum.String())

	return r
}
