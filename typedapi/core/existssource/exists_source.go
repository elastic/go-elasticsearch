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
// https://github.com/elastic/elasticsearch-specification/tree/17ac39c7f9266bc303baa029f90194aecb1c3b7c

// Returns information about whether a document source exists in an index.
package existssource

import (
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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

const (
	idMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ExistsSource struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id    string
	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewExistsSource type alias for index.
type NewExistsSource func(index, id string) *ExistsSource

// NewExistsSourceFunc returns a new instance of ExistsSource with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewExistsSourceFunc(tp elastictransport.Interface) NewExistsSource {
	return func(index, id string) *ExistsSource {
		n := New(tp)

		n._id(id)

		n._index(index)

		return n
	}
}

// Returns information about whether a document source exists in an index.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
func New(tp elastictransport.Interface) *ExistsSource {
	r := &ExistsSource{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ExistsSource) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|idMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_source")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)

		method = http.MethodHead
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r ExistsSource) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "exists_source")
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
		instrument.BeforeRequest(req, "exists_source")
		if reader := instrument.RecordRequestBody(ctx, "exists_source", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "exists_source")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ExistsSource query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a existssource.Response
func (r ExistsSource) Do(ctx context.Context) (bool, error) {
	return r.IsSuccess(ctx)
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r ExistsSource) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "exists_source")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

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

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the ExistsSource query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the ExistsSource headers map.
func (r *ExistsSource) Header(key, value string) *ExistsSource {
	r.headers.Set(key, value)

	return r
}

// Id Identifier of the document.
// API Name: id
func (r *ExistsSource) _id(id string) *ExistsSource {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Index Comma-separated list of data streams, indices, and aliases.
// Supports wildcards (`*`).
// API Name: index
func (r *ExistsSource) _index(index string) *ExistsSource {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Preference Specifies the node or shard the operation should be performed on.
// Random by default.
// API name: preference
func (r *ExistsSource) Preference(preference string) *ExistsSource {
	r.values.Set("preference", preference)

	return r
}

// Realtime If true, the request is real-time as opposed to near-real-time.
// API name: realtime
func (r *ExistsSource) Realtime(realtime bool) *ExistsSource {
	r.values.Set("realtime", strconv.FormatBool(realtime))

	return r
}

// Refresh If `true`, Elasticsearch refreshes all shards involved in the delete by query
// after the request completes.
// API name: refresh
func (r *ExistsSource) Refresh(refresh bool) *ExistsSource {
	r.values.Set("refresh", strconv.FormatBool(refresh))

	return r
}

// Routing Target the specified primary shard.
// API name: routing
func (r *ExistsSource) Routing(routing string) *ExistsSource {
	r.values.Set("routing", routing)

	return r
}

// Source_ `true` or `false` to return the `_source` field or not, or a list of fields
// to return.
// API name: _source
func (r *ExistsSource) Source_(sourceconfigparam string) *ExistsSource {
	r.values.Set("_source", sourceconfigparam)

	return r
}

// SourceExcludes_ A comma-separated list of source fields to exclude in the response.
// API name: _source_excludes
func (r *ExistsSource) SourceExcludes_(fields ...string) *ExistsSource {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ A comma-separated list of source fields to include in the response.
// API name: _source_includes
func (r *ExistsSource) SourceIncludes_(fields ...string) *ExistsSource {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// Version Explicit version number for concurrency control.
// The specified version must match the current version of the document for the
// request to succeed.
// API name: version
func (r *ExistsSource) Version(versionnumber string) *ExistsSource {
	r.values.Set("version", versionnumber)

	return r
}

// VersionType Specific version type: `external`, `external_gte`.
// API name: version_type
func (r *ExistsSource) VersionType(versiontype versiontype.VersionType) *ExistsSource {
	r.values.Set("version_type", versiontype.String())

	return r
}
