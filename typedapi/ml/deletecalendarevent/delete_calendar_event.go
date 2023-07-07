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
// https://github.com/elastic/elasticsearch-specification/tree/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c

// Deletes scheduled events from a calendar.
package deletecalendarevent

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
	calendaridMask = iota + 1

	eventidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteCalendarEvent struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	calendarid string
	eventid    string
}

// NewDeleteCalendarEvent type alias for index.
type NewDeleteCalendarEvent func(calendarid, eventid string) *DeleteCalendarEvent

// NewDeleteCalendarEventFunc returns a new instance of DeleteCalendarEvent with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteCalendarEventFunc(tp elastictransport.Interface) NewDeleteCalendarEvent {
	return func(calendarid, eventid string) *DeleteCalendarEvent {
		n := New(tp)

		n.CalendarId(calendarid)

		n.EventId(eventid)

		return n
	}
}

// Deletes scheduled events from a calendar.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/ml-delete-calendar-event.html
func New(tp elastictransport.Interface) *DeleteCalendarEvent {
	r := &DeleteCalendarEvent{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *DeleteCalendarEvent) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == calendaridMask|eventidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("calendars")
		path.WriteString("/")

		path.WriteString(r.calendarid)
		path.WriteString("/")
		path.WriteString("events")
		path.WriteString("/")

		path.WriteString(r.eventid)

		method = http.MethodDelete
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
func (r DeleteCalendarEvent) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the DeleteCalendarEvent query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a deletecalendarevent.Response
func (r DeleteCalendarEvent) Do(ctx context.Context) (*Response, error) {

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
func (r DeleteCalendarEvent) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the DeleteCalendarEvent headers map.
func (r *DeleteCalendarEvent) Header(key, value string) *DeleteCalendarEvent {
	r.headers.Set(key, value)

	return r
}

// CalendarId The ID of the calendar to modify
// API Name: calendarid
func (r *DeleteCalendarEvent) CalendarId(calendarid string) *DeleteCalendarEvent {
	r.paramSet |= calendaridMask
	r.calendarid = calendarid

	return r
}

// EventId The ID of the event to remove from the calendar
// API Name: eventid
func (r *DeleteCalendarEvent) EventId(eventid string) *DeleteCalendarEvent {
	r.paramSet |= eventidMask
	r.eventid = eventid

	return r
}
