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


// Updates certain properties of a snapshot.
package updatemodelsnapshot

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	jobidMask = iota + 1

	snapshotidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateModelSnapshot struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	jobid      string
	snapshotid string
}

// NewUpdateModelSnapshot type alias for index.
type NewUpdateModelSnapshot func(jobid, snapshotid string) *UpdateModelSnapshot

// NewUpdateModelSnapshotFunc returns a new instance of UpdateModelSnapshot with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpdateModelSnapshotFunc(tp elastictransport.Interface) NewUpdateModelSnapshot {
	return func(jobid, snapshotid string) *UpdateModelSnapshot {
		n := New(tp)

		n.JobId(jobid)

		n.SnapshotId(snapshotid)

		return n
	}
}

// Updates certain properties of a snapshot.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-snapshot.html
func New(tp elastictransport.Interface) *UpdateModelSnapshot {
	r := &UpdateModelSnapshot{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *UpdateModelSnapshot) Raw(raw json.RawMessage) *UpdateModelSnapshot {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *UpdateModelSnapshot) Request(req *Request) *UpdateModelSnapshot {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *UpdateModelSnapshot) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for UpdateModelSnapshot: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == jobidMask|snapshotidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")
		path.WriteString("/")

		path.WriteString(r.jobid)
		path.WriteString("/")
		path.WriteString("model_snapshots")
		path.WriteString("/")

		path.WriteString(r.snapshotid)
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
func (r UpdateModelSnapshot) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the UpdateModelSnapshot query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the UpdateModelSnapshot headers map.
func (r *UpdateModelSnapshot) Header(key, value string) *UpdateModelSnapshot {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the anomaly detection job.
// API Name: jobid
func (r *UpdateModelSnapshot) JobId(v string) *UpdateModelSnapshot {
	r.paramSet |= jobidMask
	r.jobid = v

	return r
}

// SnapshotId Identifier for the model snapshot.
// API Name: snapshotid
func (r *UpdateModelSnapshot) SnapshotId(v string) *UpdateModelSnapshot {
	r.paramSet |= snapshotidMask
	r.snapshotid = v

	return r
}
