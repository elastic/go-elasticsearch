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


// Updates certain properties of a datafeed.
package updatedatafeed

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
	datafeedidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateDatafeed struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	datafeedid string
}

// NewUpdateDatafeed type alias for index.
type NewUpdateDatafeed func(datafeedid string) *UpdateDatafeed

// NewUpdateDatafeedFunc returns a new instance of UpdateDatafeed with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpdateDatafeedFunc(tp elastictransport.Interface) NewUpdateDatafeed {
	return func(datafeedid string) *UpdateDatafeed {
		n := New(tp)

		n.DatafeedId(datafeedid)

		return n
	}
}

// Updates certain properties of a datafeed.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-datafeed.html
func New(tp elastictransport.Interface) *UpdateDatafeed {
	r := &UpdateDatafeed{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *UpdateDatafeed) Raw(raw json.RawMessage) *UpdateDatafeed {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *UpdateDatafeed) Request(req *Request) *UpdateDatafeed {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *UpdateDatafeed) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for UpdateDatafeed: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == datafeedidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("datafeeds")
		path.WriteString("/")

		path.WriteString(r.datafeedid)
		path.WriteString("/")
		path.WriteString("_update")

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
func (r UpdateDatafeed) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the UpdateDatafeed query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the UpdateDatafeed headers map.
func (r *UpdateDatafeed) Header(key, value string) *UpdateDatafeed {
	r.headers.Set(key, value)

	return r
}

// DatafeedId A numerical character string that uniquely identifies the datafeed.
// This identifier can contain lowercase alphanumeric characters (a-z and 0-9),
// hyphens, and underscores.
// It must start and end with alphanumeric characters.
// API Name: datafeedid
func (r *UpdateDatafeed) DatafeedId(v string) *UpdateDatafeed {
	r.paramSet |= datafeedidMask
	r.datafeedid = v

	return r
}

// AllowNoIndices If `true`, wildcard indices expressions that resolve into no concrete indices
// are ignored. This includes the
// `_all` string or when no indices are specified.
// API name: allow_no_indices
func (r *UpdateDatafeed) AllowNoIndices(b bool) *UpdateDatafeed {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match. If the request can target
// data streams, this argument determines
// whether wildcard expressions match hidden data streams. Supports
// comma-separated values. Valid values are:
//
// * `all`: Match any data stream or index, including hidden ones.
// * `closed`: Match closed, non-hidden indices. Also matches any non-hidden
// data stream. Data streams cannot be closed.
// * `hidden`: Match hidden data streams and hidden indices. Must be combined
// with `open`, `closed`, or both.
// * `none`: Wildcard patterns are not accepted.
// * `open`: Match open, non-hidden indices. Also matches any non-hidden data
// stream.
// API name: expand_wildcards
func (r *UpdateDatafeed) ExpandWildcards(value string) *UpdateDatafeed {
	r.values.Set("expand_wildcards", value)

	return r
}

// IgnoreThrottled If `true`, concrete, expanded or aliased indices are ignored when frozen.
// API name: ignore_throttled
func (r *UpdateDatafeed) IgnoreThrottled(b bool) *UpdateDatafeed {
	r.values.Set("ignore_throttled", strconv.FormatBool(b))

	return r
}

// IgnoreUnavailable If `true`, unavailable indices (missing or closed) are ignored.
// API name: ignore_unavailable
func (r *UpdateDatafeed) IgnoreUnavailable(b bool) *UpdateDatafeed {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}
