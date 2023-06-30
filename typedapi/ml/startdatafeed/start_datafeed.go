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

// Starts one or more datafeeds.
package startdatafeed

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
	datafeedidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type StartDatafeed struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw io.Reader

	paramSet int

	datafeedid string
}

// NewStartDatafeed type alias for index.
type NewStartDatafeed func(datafeedid string) *StartDatafeed

// NewStartDatafeedFunc returns a new instance of StartDatafeed with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewStartDatafeedFunc(tp elastictransport.Interface) NewStartDatafeed {
	return func(datafeedid string) *StartDatafeed {
		n := New(tp)

		n.DatafeedId(datafeedid)

		return n
	}
}

// Starts one or more datafeeds.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-start-datafeed.html
func New(tp elastictransport.Interface) *StartDatafeed {
	r := &StartDatafeed{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *StartDatafeed) Raw(raw io.Reader) *StartDatafeed {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *StartDatafeed) Request(req *Request) *StartDatafeed {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *StartDatafeed) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for StartDatafeed: %w", err)
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
		path.WriteString("_start")

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
func (r StartDatafeed) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the StartDatafeed query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a startdatafeed.Response
func (r StartDatafeed) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the StartDatafeed headers map.
func (r *StartDatafeed) Header(key, value string) *StartDatafeed {
	r.headers.Set(key, value)

	return r
}

// DatafeedId A numerical character string that uniquely identifies the datafeed. This
// identifier can contain lowercase
// alphanumeric characters (a-z and 0-9), hyphens, and underscores. It must
// start and end with alphanumeric
// characters.
// API Name: datafeedid
func (r *StartDatafeed) DatafeedId(v string) *StartDatafeed {
	r.paramSet |= datafeedidMask
	r.datafeedid = v

	return r
}

// End The time that the datafeed should end, which can be specified by using one of
// the following formats:
//
// * ISO 8601 format with milliseconds, for example `2017-01-22T06:00:00.000Z`
// * ISO 8601 format without milliseconds, for example
// `2017-01-22T06:00:00+00:00`
// * Milliseconds since the epoch, for example `1485061200000`
//
// Date-time arguments using either of the ISO 8601 formats must have a time
// zone designator, where `Z` is accepted
// as an abbreviation for UTC time. When a URL is expected (for example, in
// browsers), the `+` used in time zone
// designators must be encoded as `%2B`.
// The end time value is exclusive. If you do not specify an end time, the
// datafeed
// runs continuously.
// API name: end
func (r *StartDatafeed) End(v string) *StartDatafeed {
	r.values.Set("end", v)

	return r
}

// Start The time that the datafeed should begin, which can be specified by using the
// same formats as the `end` parameter.
// This value is inclusive.
// If you do not specify a start time and the datafeed is associated with a new
// anomaly detection job, the analysis
// starts from the earliest time for which data is available.
// If you restart a stopped datafeed and specify a start value that is earlier
// than the timestamp of the latest
// processed record, the datafeed continues from 1 millisecond after the
// timestamp of the latest processed record.
// API name: start
func (r *StartDatafeed) Start(v string) *StartDatafeed {
	r.values.Set("start", v)

	return r
}

// Timeout Specifies the amount of time to wait until a datafeed starts.
// API name: timeout
func (r *StartDatafeed) Timeout(v string) *StartDatafeed {
	r.values.Set("timeout", v)

	return r
}
