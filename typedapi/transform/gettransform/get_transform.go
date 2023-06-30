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

// Retrieves configuration information for transforms.
package gettransform

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
)

const (
	transformidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetTransform struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	transformid string
}

// NewGetTransform type alias for index.
type NewGetTransform func() *GetTransform

// NewGetTransformFunc returns a new instance of GetTransform with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetTransformFunc(tp elastictransport.Interface) NewGetTransform {
	return func() *GetTransform {
		n := New(tp)

		return n
	}
}

// Retrieves configuration information for transforms.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform.html
func New(tp elastictransport.Interface) *GetTransform {
	r := &GetTransform{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == transformidMask:
		path.WriteString("/")
		path.WriteString("_transform")
		path.WriteString("/")

		path.WriteString(r.transformid)

		method = http.MethodGet
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_transform")

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
func (r GetTransform) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetTransform query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a gettransform.Response
func (r GetTransform) Do(ctx context.Context) (*Response, error) {

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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetTransform) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the GetTransform headers map.
func (r *GetTransform) Header(key, value string) *GetTransform {
	r.headers.Set(key, value)

	return r
}

// TransformId Identifier for the transform. It can be a transform identifier or a
// wildcard expression. You can get information for all transforms by using
// `_all`, by specifying `*` as the `<transform_id>`, or by omitting the
// `<transform_id>`.
// API Name: transformid
func (r *GetTransform) TransformId(v string) *GetTransform {
	r.paramSet |= transformidMask
	r.transformid = v

	return r
}

// AllowNoMatch Specifies what to do when the request:
//
// 1. Contains wildcard expressions and there are no transforms that match.
// 2. Contains the _all string or no identifiers and there are no matches.
// 3. Contains wildcard expressions and there are only partial matches.
//
// If this parameter is false, the request returns a 404 status code when
// there are no matches or only partial matches.
// API name: allow_no_match
func (r *GetTransform) AllowNoMatch(b bool) *GetTransform {
	r.values.Set("allow_no_match", strconv.FormatBool(b))

	return r
}

// From Skips the specified number of transforms.
// API name: from
func (r *GetTransform) From(i int) *GetTransform {
	r.values.Set("from", strconv.Itoa(i))

	return r
}

// Size Specifies the maximum number of transforms to obtain.
// API name: size
func (r *GetTransform) Size(i int) *GetTransform {
	r.values.Set("size", strconv.Itoa(i))

	return r
}

// ExcludeGenerated Excludes fields that were automatically added when creating the
// transform. This allows the configuration to be in an acceptable format to
// be retrieved and then added to another cluster.
// API name: exclude_generated
func (r *GetTransform) ExcludeGenerated(b bool) *GetTransform {
	r.values.Set("exclude_generated", strconv.FormatBool(b))

	return r
}
