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
// https://github.com/elastic/elasticsearch-specification/tree/24afbdf78c21fde141eb2cad34491d952bd6daa8

// Updates the index settings.
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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutSettings struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *types.IndexSettings
	deferred []func(request *types.IndexSettings) error
	raw      io.Reader

	paramSet int

	index string
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

// Updates the index settings.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-update-settings.html
func New(tp elastictransport.Interface) *PutSettings {
	r := &PutSettings{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
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
func (r *PutSettings) Request(req *types.IndexSettings) *PutSettings {
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutSettings: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_settings")

		method = http.MethodPut
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_settings")

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
func (r PutSettings) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutSettings query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putsettings.Response
func (r PutSettings) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the PutSettings headers map.
func (r *PutSettings) Header(key, value string) *PutSettings {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and aliases used to limit
// the request. Supports wildcards (`*`). To target all data streams and
// indices, omit this parameter or use `*` or `_all`.
// API Name: index
func (r *PutSettings) Index(index string) *PutSettings {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices If `false`, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices. This
// behavior applies even if the request targets other open indices. For
// example, a request targeting `foo*,bar*` returns an error if an index
// starts with `foo` but no index starts with `bar`.
// API name: allow_no_indices
func (r *PutSettings) AllowNoIndices(allownoindices bool) *PutSettings {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match. If the request can target
// data streams, this argument determines whether wildcard expressions match
// hidden data streams. Supports comma-separated values, such as
// `open,hidden`.
// API name: expand_wildcards
func (r *PutSettings) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *PutSettings {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// FlatSettings If `true`, returns settings in flat format.
// API name: flat_settings
func (r *PutSettings) FlatSettings(flatsettings bool) *PutSettings {
	r.values.Set("flat_settings", strconv.FormatBool(flatsettings))

	return r
}

// IgnoreUnavailable If `true`, returns settings in flat format.
// API name: ignore_unavailable
func (r *PutSettings) IgnoreUnavailable(ignoreunavailable bool) *PutSettings {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an
// error.
// API name: master_timeout
func (r *PutSettings) MasterTimeout(duration string) *PutSettings {
	r.values.Set("master_timeout", duration)

	return r
}

// PreserveExisting If `true`, existing index settings remain unchanged.
// API name: preserve_existing
func (r *PutSettings) PreserveExisting(preserveexisting bool) *PutSettings {
	r.values.Set("preserve_existing", strconv.FormatBool(preserveexisting))

	return r
}

// Timeout Period to wait for a response. If no response is received before the
//  timeout expires, the request fails and returns an error.
// API name: timeout
func (r *PutSettings) Timeout(duration string) *PutSettings {
	r.values.Set("timeout", duration)

	return r
}
