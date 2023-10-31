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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

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

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

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

		n._id(id)

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

		req: NewRequest(),
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

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// Header set a key, value pair in the PutScript headers map.
func (r *PutScript) Header(key, value string) *PutScript {
	r.headers.Set(key, value)

	return r
}

// Id Identifier for the stored script or search template.
// Must be unique within the cluster.
// API Name: id
func (r *PutScript) _id(id string) *PutScript {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Context Context in which the script or search template should run.
// To prevent errors, the API immediately compiles the script or template in
// this context.
// API Name: context
func (r *PutScript) Context(context string) *PutScript {
	r.paramSet |= contextMask
	r.context = context

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *PutScript) MasterTimeout(duration string) *PutScript {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *PutScript) Timeout(duration string) *PutScript {
	r.values.Set("timeout", duration)

	return r
}

// Script Contains the script or search template, its parameters, and its language.
// API name: script
func (r *PutScript) Script(script *types.StoredScript) *PutScript {

	r.req.Script = *script

	return r
}
