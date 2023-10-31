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

// Used by the monitoring features to send monitoring data.
package bulk

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
	type_Mask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Bulk struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	type_ string
}

// NewBulk type alias for index.
type NewBulk func() *Bulk

// NewBulkFunc returns a new instance of Bulk with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewBulkFunc(tp elastictransport.Interface) NewBulk {
	return func() *Bulk {
		n := New(tp)

		return n
	}
}

// Used by the monitoring features to send monitoring data.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/monitor-elasticsearch-cluster.html
func New(tp elastictransport.Interface) *Bulk {
	r := &Bulk{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Bulk) Raw(raw io.Reader) *Bulk {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Bulk) Request(req *Request) *Bulk {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Bulk) HttpRequest(ctx context.Context) (*http.Request, error) {
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

		for _, elem := range *r.req {
			data, err := json.Marshal(elem)
			if err != nil {
				return nil, err
			}
			r.buf.Write(data)
			r.buf.Write([]byte("\n"))
		}

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Bulk: %w", err)
		}

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_monitoring")
		path.WriteString("/")
		path.WriteString("bulk")

		method = http.MethodPost
	case r.paramSet == type_Mask:
		path.WriteString("/")
		path.WriteString("_monitoring")
		path.WriteString("/")

		path.WriteString(r.type_)
		path.WriteString("/")
		path.WriteString("bulk")

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
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+x-ndjson;compatible-with=8")
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
func (r Bulk) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Bulk query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a bulk.Response
func (r Bulk) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the Bulk headers map.
func (r *Bulk) Header(key, value string) *Bulk {
	r.headers.Set(key, value)

	return r
}

// Type Default document type for items which don't provide one
// API Name: type_
func (r *Bulk) Type(type_ string) *Bulk {
	r.paramSet |= type_Mask
	r.type_ = type_

	return r
}

// SystemId Identifier of the monitored system
// API name: system_id
func (r *Bulk) SystemId(systemid string) *Bulk {
	r.values.Set("system_id", systemid)

	return r
}

// API name: system_api_version
func (r *Bulk) SystemApiVersion(systemapiversion string) *Bulk {
	r.values.Set("system_api_version", systemapiversion)

	return r
}

// Interval Collection interval (e.g., '10s' or '10000ms') of the payload
// API name: interval
func (r *Bulk) Interval(duration string) *Bulk {
	r.values.Set("interval", duration)

	return r
}
