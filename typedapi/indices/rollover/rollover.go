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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

// Updates an alias to point to a new index when the existing index
// is considered to be too large or too old.
package rollover

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
)

const (
	aliasMask = iota + 1

	newindexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Rollover struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw io.Reader

	paramSet int

	alias    string
	newindex string
}

// NewRollover type alias for index.
type NewRollover func(alias string) *Rollover

// NewRolloverFunc returns a new instance of Rollover with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewRolloverFunc(tp elastictransport.Interface) NewRollover {
	return func(alias string) *Rollover {
		n := New(tp)

		n.Alias(alias)

		return n
	}
}

// Updates an alias to point to a new index when the existing index
// is considered to be too large or too old.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-rollover-index.html
func New(tp elastictransport.Interface) *Rollover {
	r := &Rollover{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Rollover) Raw(raw io.Reader) *Rollover {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Rollover) Request(req *Request) *Rollover {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Rollover) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Rollover: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == aliasMask:
		path.WriteString("/")

		path.WriteString(r.alias)
		path.WriteString("/")
		path.WriteString("_rollover")

		method = http.MethodPost
	case r.paramSet == aliasMask|newindexMask:
		path.WriteString("/")

		path.WriteString(r.alias)
		path.WriteString("/")
		path.WriteString("_rollover")
		path.WriteString("/")

		path.WriteString(r.newindex)

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

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r Rollover) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Rollover query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a rollover.Response
func (r Rollover) Do(ctx context.Context) (*Response, error) {

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

	return nil, errorResponse
}

// Header set a key, value pair in the Rollover headers map.
func (r *Rollover) Header(key, value string) *Rollover {
	r.headers.Set(key, value)

	return r
}

// Alias The name of the alias to rollover
// API Name: alias
func (r *Rollover) Alias(v string) *Rollover {
	r.paramSet |= aliasMask
	r.alias = v

	return r
}

// NewIndex The name of the rollover index
// API Name: newindex
func (r *Rollover) NewIndex(v string) *Rollover {
	r.paramSet |= newindexMask
	r.newindex = v

	return r
}

// DryRun If set to true the rollover action will only be validated but not actually
// performed even if a condition matches. The default is false
// API name: dry_run
func (r *Rollover) DryRun(b bool) *Rollover {
	r.values.Set("dry_run", strconv.FormatBool(b))

	return r
}

// MasterTimeout Specify timeout for connection to master
// API name: master_timeout
func (r *Rollover) MasterTimeout(v string) *Rollover {
	r.values.Set("master_timeout", v)

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *Rollover) Timeout(v string) *Rollover {
	r.values.Set("timeout", v)

	return r
}

// WaitForActiveShards Set the number of active shards to wait for on the newly created rollover
// index before the operation returns.
// API name: wait_for_active_shards
func (r *Rollover) WaitForActiveShards(v string) *Rollover {
	r.values.Set("wait_for_active_shards", v)

	return r
}
