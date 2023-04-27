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

// Creates or updates a script.
package putscript

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	idMask = iota + 1

	contextMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutScript struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw io.Reader

	paramSet int

	id      string
	context string
}

// NewPutScript type alias for index.
type NewPutScript func(id string) *PutScript

// NewPutScriptFunc returns a new instance of PutScript with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutScriptFunc(tp elastictransport.Interface) NewPutScript {
	return func(id string) *PutScript {
		n := New(tp)

		n.Id(id)

		return n
	}
}

// Creates or updates a script.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/modules-scripting.html
func New(tp elastictransport.Interface) *PutScript {
	r := &PutScript{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *PutScript) Raw(raw io.Reader) *PutScript {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutScript) Request(req *Request) *PutScript {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutScript) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutScript: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_scripts")
		path.WriteString("/")

		path.WriteString(r.id)

		method = http.MethodPut
	case r.paramSet == idMask|contextMask:
		path.WriteString("/")
		path.WriteString("_scripts")
		path.WriteString("/")

		path.WriteString(r.id)
		path.WriteString("/")

		path.WriteString(r.context)

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
func (r PutScript) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutScript query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putscript.Response
func (r PutScript) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the PutScript headers map.
func (r *PutScript) Header(key, value string) *PutScript {
	r.headers.Set(key, value)

	return r
}

// Id Script ID
// API Name: id
func (r *PutScript) Id(v string) *PutScript {
	r.paramSet |= idMask
	r.id = v

	return r
}

// Context Script context
// API Name: context
func (r *PutScript) Context(v string) *PutScript {
	r.paramSet |= contextMask
	r.context = v

	return r
}

// MasterTimeout Specify timeout for connection to master
// API name: master_timeout
func (r *PutScript) MasterTimeout(v string) *PutScript {
	r.values.Set("master_timeout", v)

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *PutScript) Timeout(v string) *PutScript {
	r.values.Set("timeout", v)

	return r
}
