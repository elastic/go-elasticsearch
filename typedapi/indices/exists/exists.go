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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Returns information about whether a particular index exists.
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
)

const (
	indexMask = iota + 1
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

	index string
}

// NewExists type alias for index.
type NewExists func(index string) *Exists

// NewExistsFunc returns a new instance of Exists with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewExistsFunc(tp elastictransport.Interface) NewExists {
	return func(index string) *Exists {
		n := New(tp)

		n.Index(index)

		return n
	}
}

// Returns information about whether a particular index exists.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-exists.html
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
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)

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

// Index A comma-separated list of index names
// API Name: index
func (r *Exists) Index(v string) *Exists {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// AllowNoIndices Ignore if a wildcard expression resolves to no concrete indices (default:
// false)
// API name: allow_no_indices
func (r *Exists) AllowNoIndices(b bool) *Exists {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// ExpandWildcards Whether wildcard expressions should get expanded to open or closed indices
// (default: open)
// API name: expand_wildcards
func (r *Exists) ExpandWildcards(v string) *Exists {
	r.values.Set("expand_wildcards", v)

	return r
}

// FlatSettings Return settings in flat format (default: false)
// API name: flat_settings
func (r *Exists) FlatSettings(b bool) *Exists {
	r.values.Set("flat_settings", strconv.FormatBool(b))

	return r
}

// IgnoreUnavailable Ignore unavailable indexes (default: false)
// API name: ignore_unavailable
func (r *Exists) IgnoreUnavailable(b bool) *Exists {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// IncludeDefaults Whether to return all default setting for each of the indices.
// API name: include_defaults
func (r *Exists) IncludeDefaults(b bool) *Exists {
	r.values.Set("include_defaults", strconv.FormatBool(b))

	return r
}

// Local Return local information, do not retrieve the state from master node
// (default: false)
// API name: local
func (r *Exists) Local(b bool) *Exists {
	r.values.Set("local", strconv.FormatBool(b))

	return r
}
