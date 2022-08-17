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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


// Retrieves information about model snapshots.
package getmodelsnapshots

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
	jobidMask = iota + 1

	snapshotidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetModelSnapshots struct {
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

// NewGetModelSnapshots type alias for index.
type NewGetModelSnapshots func(jobid string) *GetModelSnapshots

// NewGetModelSnapshotsFunc returns a new instance of GetModelSnapshots with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetModelSnapshotsFunc(tp elastictransport.Interface) NewGetModelSnapshots {
	return func(jobid string) *GetModelSnapshots {
		n := New(tp)

		n.JobId(jobid)

		return n
	}
}

// Retrieves information about model snapshots.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-snapshot.html
func New(tp elastictransport.Interface) *GetModelSnapshots {
	r := &GetModelSnapshots{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *GetModelSnapshots) Raw(raw json.RawMessage) *GetModelSnapshots {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *GetModelSnapshots) Request(req *Request) *GetModelSnapshots {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetModelSnapshots) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for GetModelSnapshots: %w", err)
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
		path.WriteString(url.PathEscape(r.jobid))
		path.WriteString("/")
		path.WriteString("model_snapshots")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.snapshotid))

		method = http.MethodPost
	case r.paramSet == jobidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.jobid))
		path.WriteString("/")
		path.WriteString("model_snapshots")

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

	if r.buf.Len() > 0 {
		req.Header.Set("content-type", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	req.Header.Set("accept", "application/vnd.elasticsearch+json;compatible-with=8")

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Do runs the http.Request through the provided transport.
func (r GetModelSnapshots) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetModelSnapshots query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the GetModelSnapshots headers map.
func (r *GetModelSnapshots) Header(key, value string) *GetModelSnapshots {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the anomaly detection job.
// API Name: jobid
func (r *GetModelSnapshots) JobId(v string) *GetModelSnapshots {
	r.paramSet |= jobidMask
	r.jobid = v

	return r
}

// SnapshotId A numerical character string that uniquely identifies the model snapshot. You
// can get information for multiple
// snapshots by using a comma-separated list or a wildcard expression. You can
// get all snapshots by using `_all`,
// by specifying `*` as the snapshot ID, or by omitting the snapshot ID.
// API Name: snapshotid
func (r *GetModelSnapshots) SnapshotId(v string) *GetModelSnapshots {
	r.paramSet |= snapshotidMask
	r.snapshotid = v

	return r
}

// Desc If true, the results are sorted in descending order.
// API name: desc
func (r *GetModelSnapshots) Desc(b bool) *GetModelSnapshots {
	r.values.Set("desc", strconv.FormatBool(b))

	return r
}

// End Returns snapshots with timestamps earlier than this time.
// API name: end
func (r *GetModelSnapshots) End(value string) *GetModelSnapshots {
	r.values.Set("end", value)

	return r
}

// From Skips the specified number of snapshots.
// API name: from
func (r *GetModelSnapshots) From(i int) *GetModelSnapshots {
	r.values.Set("from", strconv.Itoa(i))

	return r
}

// Size Specifies the maximum number of snapshots to obtain.
// API name: size
func (r *GetModelSnapshots) Size(i int) *GetModelSnapshots {
	r.values.Set("size", strconv.Itoa(i))

	return r
}

// Sort Specifies the sort field for the requested snapshots. By default, the
// snapshots are sorted by their timestamp.
// API name: sort
func (r *GetModelSnapshots) Sort(value string) *GetModelSnapshots {
	r.values.Set("sort", value)

	return r
}

// Start Returns snapshots with timestamps after this time.
// API name: start
func (r *GetModelSnapshots) Start(value string) *GetModelSnapshots {
	r.values.Set("start", value)

	return r
}
