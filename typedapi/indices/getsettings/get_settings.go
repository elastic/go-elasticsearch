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
// https://github.com/elastic/elasticsearch-specification/tree/76e25d34bff1060e300c95f4be468ef88e4f3465

// Returns settings for one or more indices.
package getsettings

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1

	nameMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetSettings struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	index string
	name  string
}

// NewGetSettings type alias for index.
type NewGetSettings func() *GetSettings

// NewGetSettingsFunc returns a new instance of GetSettings with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetSettingsFunc(tp elastictransport.Interface) NewGetSettings {
	return func() *GetSettings {
		n := New(tp)

		return n
	}
}

// Returns settings for one or more indices.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-get-settings.html
func New(tp elastictransport.Interface) *GetSettings {
	r := &GetSettings{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetSettings) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_settings")

		method = http.MethodGet
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_settings")

		method = http.MethodGet
	case r.paramSet == indexMask|nameMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_settings")
		path.WriteString("/")

		path.WriteString(r.name)

		method = http.MethodGet
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_settings")
		path.WriteString("/")

		path.WriteString(r.name)

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
func (r GetSettings) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetSettings query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getsettings.Response
func (r GetSettings) Do(ctx context.Context) (Response, error) {

	response := NewResponse()

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
func (r GetSettings) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the GetSettings headers map.
func (r *GetSettings) Header(key, value string) *GetSettings {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names; use `_all` or empty string to perform
// the operation on all indices
// API Name: index
func (r *GetSettings) Index(index string) *GetSettings {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Name The name of the settings that should be included
// API Name: name
func (r *GetSettings) Name(name string) *GetSettings {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *GetSettings) AllowNoIndices(allownoindices bool) *GetSettings {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *GetSettings) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *GetSettings {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// FlatSettings Return settings in flat format (default: false)
// API name: flat_settings
func (r *GetSettings) FlatSettings(flatsettings bool) *GetSettings {
	r.values.Set("flat_settings", strconv.FormatBool(flatsettings))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *GetSettings) IgnoreUnavailable(ignoreunavailable bool) *GetSettings {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// IncludeDefaults Whether to return all default setting for each of the indices.
// API name: include_defaults
func (r *GetSettings) IncludeDefaults(includedefaults bool) *GetSettings {
	r.values.Set("include_defaults", strconv.FormatBool(includedefaults))

	return r
}

// Local Return local information, do not retrieve the state from master node
// (default: false)
// API name: local
func (r *GetSettings) Local(local bool) *GetSettings {
	r.values.Set("local", strconv.FormatBool(local))

	return r
}

// MasterTimeout Specify timeout for connection to master
// API name: master_timeout
func (r *GetSettings) MasterTimeout(duration string) *GetSettings {
	r.values.Set("master_timeout", duration)

	return r
}
