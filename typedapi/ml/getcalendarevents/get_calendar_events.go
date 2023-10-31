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

// Retrieves information about the scheduled events in calendars.
package getcalendarevents

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
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	calendaridMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetCalendarEvents struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	calendarid string
}

// NewGetCalendarEvents type alias for index.
type NewGetCalendarEvents func(calendarid string) *GetCalendarEvents

// NewGetCalendarEventsFunc returns a new instance of GetCalendarEvents with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetCalendarEventsFunc(tp elastictransport.Interface) NewGetCalendarEvents {
	return func(calendarid string) *GetCalendarEvents {
		n := New(tp)

		n._calendarid(calendarid)

		return n
	}
}

// Retrieves information about the scheduled events in calendars.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-calendar-event.html
func New(tp elastictransport.Interface) *GetCalendarEvents {
	r := &GetCalendarEvents{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetCalendarEvents) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == calendaridMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("calendars")
		path.WriteString("/")

		path.WriteString(r.calendarid)
		path.WriteString("/")
		path.WriteString("events")

		method = http.MethodGet
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
func (r GetCalendarEvents) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetCalendarEvents query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getcalendarevents.Response
func (r GetCalendarEvents) Do(ctx context.Context) (*Response, error) {

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
func (r GetCalendarEvents) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the GetCalendarEvents headers map.
func (r *GetCalendarEvents) Header(key, value string) *GetCalendarEvents {
	r.headers.Set(key, value)

	return r
}

// CalendarId A string that uniquely identifies a calendar. You can get information for
// multiple calendars by using a comma-separated list of ids or a wildcard
// expression. You can get information for all calendars by using `_all` or `*`
// or by omitting the calendar identifier.
// API Name: calendarid
func (r *GetCalendarEvents) _calendarid(calendarid string) *GetCalendarEvents {
	r.paramSet |= calendaridMask
	r.calendarid = calendarid

	return r
}

// End Specifies to get events with timestamps earlier than this time.
// API name: end
func (r *GetCalendarEvents) End(datetime string) *GetCalendarEvents {
	r.values.Set("end", datetime)

	return r
}

// From Skips the specified number of events.
// API name: from
func (r *GetCalendarEvents) From(from int) *GetCalendarEvents {
	r.values.Set("from", strconv.Itoa(from))

	return r
}

// JobId Specifies to get events for a specific anomaly detection job identifier or
// job group. It must be used with a calendar identifier of `_all` or `*`.
// API name: job_id
func (r *GetCalendarEvents) JobId(id string) *GetCalendarEvents {
	r.values.Set("job_id", id)

	return r
}

// Size Specifies the maximum number of events to obtain.
// API name: size
func (r *GetCalendarEvents) Size(size int) *GetCalendarEvents {
	r.values.Set("size", strconv.Itoa(size))

	return r
}

// Start Specifies to get events with timestamps after this time.
// API name: start
func (r *GetCalendarEvents) Start(datetime string) *GetCalendarEvents {
	r.values.Set("start", datetime)

	return r
}
