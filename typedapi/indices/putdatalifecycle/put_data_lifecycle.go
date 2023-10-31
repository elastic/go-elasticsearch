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

// Updates the data stream lifecycle of the selected data streams.
package putdatalifecycle

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
)

const (
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutDataLifecycle struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	name string
}

// NewPutDataLifecycle type alias for index.
type NewPutDataLifecycle func(name string) *PutDataLifecycle

// NewPutDataLifecycleFunc returns a new instance of PutDataLifecycle with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutDataLifecycleFunc(tp elastictransport.Interface) NewPutDataLifecycle {
	return func(name string) *PutDataLifecycle {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Updates the data stream lifecycle of the selected data streams.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/data-streams-put-lifecycle.html
func New(tp elastictransport.Interface) *PutDataLifecycle {
	r := &PutDataLifecycle{
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
func (r *PutDataLifecycle) Raw(raw io.Reader) *PutDataLifecycle {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutDataLifecycle) Request(req *Request) *PutDataLifecycle {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutDataLifecycle) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutDataLifecycle: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_data_stream")
		path.WriteString("/")

		path.WriteString(r.name)
		path.WriteString("/")
		path.WriteString("_lifecycle")

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
func (r PutDataLifecycle) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutDataLifecycle query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putdatalifecycle.Response
func (r PutDataLifecycle) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the PutDataLifecycle headers map.
func (r *PutDataLifecycle) Header(key, value string) *PutDataLifecycle {
	r.headers.Set(key, value)

	return r
}

// Name Comma-separated list of data streams used to limit the request.
// Supports wildcards (`*`).
// To target all data streams use `*` or `_all`.
// API Name: name
func (r *PutDataLifecycle) _name(name string) *PutDataLifecycle {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// ExpandWildcards Type of data stream that wildcard patterns can match.
// Supports comma-separated values, such as `open,hidden`.
// Valid values are: `all`, `hidden`, `open`, `closed`, `none`.
// API name: expand_wildcards
func (r *PutDataLifecycle) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *PutDataLifecycle {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an
// error.
// API name: master_timeout
func (r *PutDataLifecycle) MasterTimeout(duration string) *PutDataLifecycle {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *PutDataLifecycle) Timeout(duration string) *PutDataLifecycle {
	r.values.Set("timeout", duration)

	return r
}

// DataRetention If defined, every document added to this data stream will be stored at least
// for this time frame.
// Any time after this duration the document could be deleted.
// When empty, every document in this data stream will be stored indefinitely.
// API name: data_retention
func (r *PutDataLifecycle) DataRetention(duration types.Duration) *PutDataLifecycle {
	r.req.DataRetention = duration

	return r
}

// Downsampling If defined, every backing index will execute the configured downsampling
// configuration after the backing
// index is not the data stream write index anymore.
// API name: downsampling
func (r *PutDataLifecycle) Downsampling(downsampling *types.DataStreamLifecycleDownsampling) *PutDataLifecycle {

	r.req.Downsampling = downsampling

	return r
}
