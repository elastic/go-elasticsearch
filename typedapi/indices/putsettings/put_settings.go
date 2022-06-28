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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


// Updates the index settings.
package putsettings

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
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

	req *types.IndexSettings
	raw json.RawMessage

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
func (r *PutSettings) Raw(raw json.RawMessage) *PutSettings {
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

	if r.raw != nil {
		r.buf.Write(r.raw)
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
		path.WriteString(url.PathEscape(r.index))
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

	if r.buf.Len() > 0 {
		req.Header.Set("content-type", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	req.Header.Set("accept", "application/vnd.elasticsearch+json;compatible-with=8")

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Do runs the http.Request through the provided transport.
func (r PutSettings) Do(ctx context.Context) (*http.Response, error) {
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

// Header set a key, value pair in the PutSettings headers map.
func (r *PutSettings) Header(key, value string) *PutSettings {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names; use `_all` or empty string to perform
// the operation on all indices
// API Name: index
func (r *PutSettings) Index(v string) *PutSettings {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *PutSettings) AllowNoIndices(b bool) *PutSettings {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *PutSettings) ExpandWildcards(value string) *PutSettings {
	r.values.Set("expand_wildcards", value)

	return r
}

// FlatSettings Return settings in flat format (default: false)
// API name: flat_settings
func (r *PutSettings) FlatSettings(b bool) *PutSettings {
	r.values.Set("flat_settings", strconv.FormatBool(b))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *PutSettings) IgnoreUnavailable(b bool) *PutSettings {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// MasterTimeout Specify timeout for connection to master
// API name: master_timeout
func (r *PutSettings) MasterTimeout(value string) *PutSettings {
	r.values.Set("master_timeout", value)

	return r
}

// PreserveExisting Whether to update existing settings. If set to `true` existing settings on an
// index remain unchanged, the default is `false`
// API name: preserve_existing
func (r *PutSettings) PreserveExisting(b bool) *PutSettings {
	r.values.Set("preserve_existing", strconv.FormatBool(b))

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *PutSettings) Timeout(value string) *PutSettings {
	r.values.Set("timeout", value)

	return r
}
