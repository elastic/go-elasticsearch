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


// Allow to shrink an existing index into a new index with fewer primary shards.
package shrink

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	indexMask = iota + 1

	targetMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Shrink struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	index  string
	target string
}

// NewShrink type alias for index.
type NewShrink func(index, target string) *Shrink

// NewShrinkFunc returns a new instance of Shrink with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewShrinkFunc(tp elastictransport.Interface) NewShrink {
	return func(index, target string) *Shrink {
		n := New(tp)

		n.Index(index)

		n.Target(target)

		return n
	}
}

// Allow to shrink an existing index into a new index with fewer primary shards.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-shrink-index.html
func New(tp elastictransport.Interface) *Shrink {
	r := &Shrink{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Shrink) Raw(raw json.RawMessage) *Shrink {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Shrink) Request(req *Request) *Shrink {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Shrink) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Shrink: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|targetMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_shrink")
		path.WriteString("/")

		path.WriteString(r.target)

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

// Do runs the http.Request through the provided transport.
func (r Shrink) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Shrink query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the Shrink headers map.
func (r *Shrink) Header(key, value string) *Shrink {
	r.headers.Set(key, value)

	return r
}

// Index The name of the source index to shrink
// API Name: index
func (r *Shrink) Index(v string) *Shrink {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// Target The name of the target index to shrink into
// API Name: target
func (r *Shrink) Target(v string) *Shrink {
	r.paramSet |= targetMask
	r.target = v

	return r
}

// MasterTimeout Specify timeout for connection to master
// API name: master_timeout
func (r *Shrink) MasterTimeout(value string) *Shrink {
	r.values.Set("master_timeout", value)

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *Shrink) Timeout(value string) *Shrink {
	r.values.Set("timeout", value)

	return r
}

// WaitForActiveShards Set the number of active shards to wait for on the shrunken index before the
// operation returns.
// API name: wait_for_active_shards
func (r *Shrink) WaitForActiveShards(value string) *Shrink {
	r.values.Set("wait_for_active_shards", value)

	return r
}
