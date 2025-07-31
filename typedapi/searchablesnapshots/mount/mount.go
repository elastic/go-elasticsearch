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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Mount a snapshot.
// Mount a snapshot as a searchable snapshot index.
// Do not use this API for snapshots managed by index lifecycle management
// (ILM).
// Manually mounting ILM-managed snapshots can interfere with ILM processes.
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

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	repository string
	snapshot   string

	spanStarted bool

	instrument elastictransport.Instrumentation
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

// Mount a snapshot.
// Mount a snapshot as a searchable snapshot index.
// Do not use this API for snapshots managed by index lifecycle management
// (ILM).
// Manually mounting ILM-managed snapshots can interfere with ILM processes.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/searchable-snapshots-api-mount-snapshot.html
func New(tp elastictransport.Interface) *Mount {
	r := &Mount{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
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

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Mount: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == repositoryMask|snapshotMask:
		path.WriteString("/")
		path.WriteString("_snapshot")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "repository", r.repository)
		}
		path.WriteString(r.repository)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "snapshot", r.snapshot)
		}
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
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
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
func (r Mount) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "searchable_snapshots.mount")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "searchable_snapshots.mount")
		if reader := instrument.RecordRequestBody(ctx, "searchable_snapshots.mount", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "searchable_snapshots.mount")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Mount query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a mount.Response
func (r Mount) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "searchable_snapshots.mount")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// Header set a key, value pair in the Mount headers map.
func (r *Mount) Header(key, value string) *Mount {
	r.headers.Set(key, value)

	return r
}

// Repository The name of the repository containing the snapshot of the index to mount.
// API Name: repository
func (r *Mount) _repository(repository string) *Mount {
	r.paramSet |= repositoryMask
	r.repository = repository

	return r
}

// Snapshot The name of the snapshot of the index to mount.
// API Name: snapshot
func (r *Mount) _snapshot(snapshot string) *Mount {
	r.paramSet |= snapshotMask
	r.snapshot = snapshot

	return r
}

// MasterTimeout The period to wait for the master node.
// If the master node is not available before the timeout expires, the request
// fails and returns an error.
// To indicate that the request should never timeout, set it to `-1`.
// API name: master_timeout
func (r *Mount) MasterTimeout(duration string) *Mount {
	r.values.Set("master_timeout", duration)

	return r
}

// WaitForCompletion If true, the request blocks until the operation is complete.
// API name: wait_for_completion
func (r *Mount) WaitForCompletion(waitforcompletion bool) *Mount {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}

// Storage The mount option for the searchable snapshot index.
// API name: storage
func (r *Mount) Storage(storage string) *Mount {
	r.values.Set("storage", storage)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Mount) ErrorTrace(errortrace bool) *Mount {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Mount) FilterPath(filterpaths ...string) *Mount {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *Mount) Human(human bool) *Mount {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Mount) Pretty(pretty bool) *Mount {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// IgnoreIndexSettings The names of settings that should be removed from the index when it is
// mounted.
// API name: ignore_index_settings
func (r *Mount) IgnoreIndexSettings(ignoreindexsettings ...string) *Mount {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.IgnoreIndexSettings = ignoreindexsettings

	return r
}

// Index The name of the index contained in the snapshot whose data is to be mounted.
// If no `renamed_index` is specified, this name will also be used to create the
// new index.
// API name: index
func (r *Mount) Index(indexname string) *Mount {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Index = indexname

	return r
}

// IndexSettings The settings that should be added to the index when it is mounted.
// API name: index_settings
func (r *Mount) IndexSettings(indexsettings map[string]json.RawMessage) *Mount {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexSettings = indexsettings

	return r
}

// RenamedIndex The name of the index that will be created.
// API name: renamed_index
func (r *Mount) RenamedIndex(indexname string) *Mount {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RenamedIndex = &indexname

	return r
}
