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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Allows to get multiple documents in one request.
package mget

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
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Mget struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	index string
}

// NewMget type alias for index.
type NewMget func() *Mget

// NewMgetFunc returns a new instance of Mget with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewMgetFunc(tp elastictransport.Interface) NewMget {
	return func() *Mget {
		n := New(tp)

		return n
	}
}

// Allows to get multiple documents in one request.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-multi-get.html
func New(tp elastictransport.Interface) *Mget {
	r := &Mget{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Mget) Raw(raw json.RawMessage) *Mget {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Mget) Request(req *Request) *Mget {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Mget) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Mget: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_mget")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_mget")

		method = http.MethodPost
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

// Do runs the http.Request through the provided transport.
func (r Mget) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Mget query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the Mget headers map.
func (r *Mget) Header(key, value string) *Mget {
	r.headers.Set(key, value)

	return r
}

// Index Name of the index to retrieve documents from when `ids` are specified, or
// when a document in the `docs` array does not specify an index.
// API Name: index
func (r *Mget) Index(v string) *Mget {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// Preference Specifies the node or shard the operation should be performed on. Random by
// default.
// API name: preference
func (r *Mget) Preference(value string) *Mget {
	r.values.Set("preference", value)

	return r
}

// Realtime If `true`, the request is real-time as opposed to near-real-time.
// API name: realtime
func (r *Mget) Realtime(b bool) *Mget {
	r.values.Set("realtime", strconv.FormatBool(b))

	return r
}

// Refresh If `true`, the request refreshes relevant shards before retrieving documents.
// API name: refresh
func (r *Mget) Refresh(b bool) *Mget {
	r.values.Set("refresh", strconv.FormatBool(b))

	return r
}

// Routing Custom value used to route operations to a specific shard.
// API name: routing
func (r *Mget) Routing(value string) *Mget {
	r.values.Set("routing", value)

	return r
}

// Source_ True or false to return the `_source` field or not, or a list of fields to
// return.
// API name: _source
func (r *Mget) Source_(value string) *Mget {
	r.values.Set("_source", value)

	return r
}

// SourceExcludes_ A comma-separated list of source fields to exclude from the response.
// You can also use this parameter to exclude fields from the subset specified
// in `_source_includes` query parameter.
// API name: _source_excludes
func (r *Mget) SourceExcludes_(value string) *Mget {
	r.values.Set("_source_excludes", value)

	return r
}

// SourceIncludes_ A comma-separated list of source fields to include in the response.
// If this parameter is specified, only these source fields are returned. You
// can exclude fields from this subset using the `_source_excludes` query
// parameter.
// If the `_source` parameter is `false`, this parameter is ignored.
// API name: _source_includes
func (r *Mget) SourceIncludes_(value string) *Mget {
	r.values.Set("_source_includes", value)

	return r
}

// StoredFields If `true`, retrieves the document fields stored in the index rather than the
// document `_source`.
// API name: stored_fields
func (r *Mget) StoredFields(value string) *Mget {
	r.values.Set("stored_fields", value)

	return r
}
