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
// https://github.com/elastic/elasticsearch-specification/tree/d63a0e35ee85d84c83d0612ff3c0641a7a1e7e2e


// Stops one or more data frame analytics jobs.
package stopdataframeanalytics

import (
	gobytes "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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

type StopDataFrameAnalytics struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	id string
}

// NewStopDataFrameAnalytics type alias for index.
type NewStopDataFrameAnalytics func(id string) *StopDataFrameAnalytics

// NewStopDataFrameAnalyticsFunc returns a new instance of StopDataFrameAnalytics with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewStopDataFrameAnalyticsFunc(tp elastictransport.Interface) NewStopDataFrameAnalytics {
	return func(id string) *StopDataFrameAnalytics {
		n := New(tp)

		n.Id(id)

		return n
	}
}

// Stops one or more data frame analytics jobs.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-dfanalytics.html
func New(tp elastictransport.Interface) *StopDataFrameAnalytics {
	r := &StopDataFrameAnalytics{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *StopDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("data_frame")
		path.WriteString("/")
		path.WriteString("analytics")
		path.WriteString("/")

		path.WriteString(r.id)
		path.WriteString("/")
		path.WriteString("_stop")

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
func (r StopDataFrameAnalytics) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the StopDataFrameAnalytics query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r StopDataFrameAnalytics) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Do(ctx)

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

// Header set a key, value pair in the StopDataFrameAnalytics headers map.
func (r *StopDataFrameAnalytics) Header(key, value string) *StopDataFrameAnalytics {
	r.headers.Set(key, value)

	return r
}

// Id Identifier for the data frame analytics job. This identifier can contain
// lowercase alphanumeric characters (a-z and 0-9), hyphens, and
// underscores. It must start and end with alphanumeric characters.
// API Name: id
func (r *StopDataFrameAnalytics) Id(v string) *StopDataFrameAnalytics {
	r.paramSet |= idMask
	r.id = v

	return r
}

// AllowNoMatch Specifies what to do when the request:
//
// 1. Contains wildcard expressions and there are no data frame analytics
// jobs that match.
// 2. Contains the _all string or no identifiers and there are no matches.
// 3. Contains wildcard expressions and there are only partial matches.
//
// The default value is true, which returns an empty data_frame_analytics
// array when there are no matches and the subset of results when there are
// partial matches. If this parameter is false, the request returns a 404
// status code when there are no matches or only partial matches.
// API name: allow_no_match
func (r *StopDataFrameAnalytics) AllowNoMatch(b bool) *StopDataFrameAnalytics {
	r.values.Set("allow_no_match", strconv.FormatBool(b))

	return r
}

// Force If true, the data frame analytics job is stopped forcefully.
// API name: force
func (r *StopDataFrameAnalytics) Force(b bool) *StopDataFrameAnalytics {
	r.values.Set("force", strconv.FormatBool(b))

	return r
}

// Timeout Controls the amount of time to wait until the data frame analytics job
// stops. Defaults to 20 seconds.
// API name: timeout
func (r *StopDataFrameAnalytics) Timeout(value string) *StopDataFrameAnalytics {
	r.values.Set("timeout", value)

	return r
}
