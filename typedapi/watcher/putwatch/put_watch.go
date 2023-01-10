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


// Creates a new watch, or updates an existing one.
package putwatch

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
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutWatch struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	id string
}

// NewPutWatch type alias for index.
type NewPutWatch func(id string) *PutWatch

// NewPutWatchFunc returns a new instance of PutWatch with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutWatchFunc(tp elastictransport.Interface) NewPutWatch {
	return func(id string) *PutWatch {
		n := New(tp)

		n.Id(id)

		return n
	}
}

// Creates a new watch, or updates an existing one.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-put-watch.html
func New(tp elastictransport.Interface) *PutWatch {
	r := &PutWatch{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *PutWatch) Raw(raw json.RawMessage) *PutWatch {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutWatch) Request(req *Request) *PutWatch {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutWatch) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutWatch: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_watcher")
		path.WriteString("/")
		path.WriteString("watch")
		path.WriteString("/")

		path.WriteString(r.id)

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
func (r PutWatch) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutWatch query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the PutWatch headers map.
func (r *PutWatch) Header(key, value string) *PutWatch {
	r.headers.Set(key, value)

	return r
}

// Id Watch ID
// API Name: id
func (r *PutWatch) Id(v string) *PutWatch {
	r.paramSet |= idMask
	r.id = v

	return r
}

// Active Specify whether the watch is in/active by default
// API name: active
func (r *PutWatch) Active(b bool) *PutWatch {
	r.values.Set("active", strconv.FormatBool(b))

	return r
}

// IfPrimaryTerm only update the watch if the last operation that has changed the watch has
// the specified primary term
// API name: if_primary_term
func (r *PutWatch) IfPrimaryTerm(value string) *PutWatch {
	r.values.Set("if_primary_term", value)

	return r
}

// IfSeqNo only update the watch if the last operation that has changed the watch has
// the specified sequence number
// API name: if_seq_no
func (r *PutWatch) IfSeqNo(value string) *PutWatch {
	r.values.Set("if_seq_no", value)

	return r
}

// Version Explicit version number for concurrency control
// API name: version
func (r *PutWatch) Version(value string) *PutWatch {
	r.values.Set("version", value)

	return r
}
