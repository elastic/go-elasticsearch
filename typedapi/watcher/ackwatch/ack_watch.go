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

// Acknowledges a watch, manually throttling the execution of the watch's
// actions.
package ackwatch

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
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	watchidMask = iota + 1

	actionidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type AckWatch struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	watchid  string
	actionid string
}

// NewAckWatch type alias for index.
type NewAckWatch func(watchid string) *AckWatch

// NewAckWatchFunc returns a new instance of AckWatch with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewAckWatchFunc(tp elastictransport.Interface) NewAckWatch {
	return func(watchid string) *AckWatch {
		n := New(tp)

		n.WatchId(watchid)

		return n
	}
}

// Acknowledges a watch, manually throttling the execution of the watch's
// actions.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-ack-watch.html
func New(tp elastictransport.Interface) *AckWatch {
	r := &AckWatch{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *AckWatch) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == watchidMask:
		path.WriteString("/")
		path.WriteString("_watcher")
		path.WriteString("/")
		path.WriteString("watch")
		path.WriteString("/")

		path.WriteString(r.watchid)
		path.WriteString("/")
		path.WriteString("_ack")

		method = http.MethodPut
	case r.paramSet == watchidMask|actionidMask:
		path.WriteString("/")
		path.WriteString("_watcher")
		path.WriteString("/")
		path.WriteString("watch")
		path.WriteString("/")

		path.WriteString(r.watchid)
		path.WriteString("/")
		path.WriteString("_ack")
		path.WriteString("/")

		path.WriteString(r.actionid)

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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r AckWatch) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the AckWatch query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a ackwatch.Response
func (r AckWatch) Do(ctx context.Context) (*Response, error) {

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
func (r AckWatch) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the AckWatch headers map.
func (r *AckWatch) Header(key, value string) *AckWatch {
	r.headers.Set(key, value)

	return r
}

// WatchId Watch ID
// API Name: watchid
func (r *AckWatch) WatchId(v string) *AckWatch {
	r.paramSet |= watchidMask
	r.watchid = v

	return r
}

// ActionId A comma-separated list of the action ids to be acked
// API Name: actionid
func (r *AckWatch) ActionId(v string) *AckWatch {
	r.paramSet |= actionidMask
	r.actionid = v

	return r
}
