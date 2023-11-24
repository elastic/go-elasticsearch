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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Mount a snapshot as a searchable index.
package mount

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	repositoryMask = iota + 1

	snapshotMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Mount struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	repository string
	snapshot   string
}

// NewMount type alias for index.
type NewMount func(repository, snapshot string) *Mount

// NewMountFunc returns a new instance of Mount with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewMountFunc(tp elastictransport.Interface) NewMount {
	return func(repository, snapshot string) *Mount {
		n := New(tp)

		n._repository(repository)

		n._snapshot(snapshot)

		return n
	}
}

// Mount a snapshot as a searchable index.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/searchable-snapshots-api-mount-snapshot.html
func New(tp elastictransport.Interface) *Mount {
	r := &Mount{
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
func (r *Mount) Raw(raw io.Reader) *Mount {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Mount) Request(req *Request) *Mount {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Mount) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Mount: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == repositoryMask|snapshotMask:
		path.WriteString("/")
		path.WriteString("_snapshot")
		path.WriteString("/")

		path.WriteString(r.repository)
		path.WriteString("/")

		path.WriteString(r.snapshot)
		path.WriteString("/")
		path.WriteString("_mount")

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
func (r Mount) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Mount query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a mount.Response
func (r Mount) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the Mount headers map.
func (r *Mount) Header(key, value string) *Mount {
	r.headers.Set(key, value)

	return r
}

// Repository The name of the repository containing the snapshot of the index to mount
// API Name: repository
func (r *Mount) _repository(repository string) *Mount {
	r.paramSet |= repositoryMask
	r.repository = repository

	return r
}

// Snapshot The name of the snapshot of the index to mount
// API Name: snapshot
func (r *Mount) _snapshot(snapshot string) *Mount {
	r.paramSet |= snapshotMask
	r.snapshot = snapshot

	return r
}

// MasterTimeout Explicit operation timeout for connection to master node
// API name: master_timeout
func (r *Mount) MasterTimeout(duration string) *Mount {
	r.values.Set("master_timeout", duration)

	return r
}

// WaitForCompletion Should this request wait until the operation has completed before returning
// API name: wait_for_completion
func (r *Mount) WaitForCompletion(waitforcompletion bool) *Mount {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}

// Storage Selects the kind of local storage used to accelerate searches. Experimental,
// and defaults to `full_copy`
// API name: storage
func (r *Mount) Storage(storage string) *Mount {
	r.values.Set("storage", storage)

	return r
}

// API name: ignore_index_settings
func (r *Mount) IgnoreIndexSettings(ignoreindexsettings ...string) *Mount {
	r.req.IgnoreIndexSettings = ignoreindexsettings

	return r
}

// API name: index
func (r *Mount) Index(indexname string) *Mount {
	r.req.Index = indexname

	return r
}

// API name: index_settings
func (r *Mount) IndexSettings(indexsettings map[string]json.RawMessage) *Mount {

	r.req.IndexSettings = indexsettings

	return r
}

// API name: renamed_index
func (r *Mount) RenamedIndex(indexname string) *Mount {
	r.req.RenamedIndex = &indexname

	return r
}
