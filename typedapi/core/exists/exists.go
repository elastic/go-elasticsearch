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

// Check a document.
//
// Verify that a document exists.
// For example, check to see if a document with the `_id` 0 exists:
//
// ```
// HEAD my-index-000001/_doc/0
// ```
//
// If the document exists, the API returns a status code of `200 - OK`.
// If the document doesn’t exist, the API returns `404 - Not Found`.
//
// **Versioning support**
//
// You can use the `version` parameter to check the document only if its current
// version is equal to the specified one.
//
// Internally, Elasticsearch has marked the old document as deleted and added an
// entirely new document.
// The old version of the document doesn't disappear immediately, although you
// won't be able to access it.
// Elasticsearch cleans up deleted documents in the background as you continue
// to index more data.
package exists

import (
	"context"
	"errors"
	"fmt"
	"io"
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

type Exists struct {
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

// NewExists type alias for index.
type NewExists func(index, id string) *Exists

// NewExistsFunc returns a new instance of Exists with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewExistsFunc(tp elastictransport.Interface) NewExists {
	return func(index, id string) *Exists {
		n := New(tp)

		n._id(id)

		n._index(index)

		return n
	}
}

// Check a document.
//
// Verify that a document exists.
// For example, check to see if a document with the `_id` 0 exists:
//
// ```
// HEAD my-index-000001/_doc/0
// ```
//
// If the document exists, the API returns a status code of `200 - OK`.
// If the document doesn’t exist, the API returns `404 - Not Found`.
//
// **Versioning support**
//
// You can use the `version` parameter to check the document only if its current
// version is equal to the specified one.
//
// Internally, Elasticsearch has marked the old document as deleted and added an
// entirely new document.
// The old version of the document doesn't disappear immediately, although you
// won't be able to access it.
// Elasticsearch cleans up deleted documents in the background as you continue
// to index more data.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
func New(tp elastictransport.Interface) *Exists {
	r := &Exists{
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
func (r *Exists) HttpRequest(ctx context.Context) (*http.Request, error) {
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
		path.WriteString("_doc")
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
func (r Exists) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "exists")
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
		instrument.BeforeRequest(req, "exists")
		if reader := instrument.RecordRequestBody(ctx, "exists", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "exists")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Exists query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a exists.Response
func (r Exists) Do(ctx context.Context) (bool, error) {
	return r.IsSuccess(ctx)
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r Exists) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "exists")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(io.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the Exists query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the Exists headers map.
func (r *Exists) Header(key, value string) *Exists {
	r.headers.Set(key, value)

	return r
}

// Id A unique document identifier.
// API Name: id
func (r *Exists) _id(id string) *Exists {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Index A comma-separated list of data streams, indices, and aliases.
// It supports wildcards (`*`).
// API Name: index
func (r *Exists) _index(index string) *Exists {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Preference The node or shard the operation should be performed on.
// By default, the operation is randomized between the shard replicas.
//
// If it is set to `_local`, the operation will prefer to be run on a local
// allocated shard when possible.
// If it is set to a custom value, the value is used to guarantee that the same
// shards will be used for the same custom value.
// This can help with "jumping values" when hitting different shards in
// different refresh states.
// A sample value can be something like the web session ID or the user name.
// API name: preference
func (r *Exists) Preference(preference string) *Exists {
	r.values.Set("preference", preference)

	return r
}

// Realtime If `true`, the request is real-time as opposed to near-real-time.
// API name: realtime
func (r *Exists) Realtime(realtime bool) *Exists {
	r.values.Set("realtime", strconv.FormatBool(realtime))

	return r
}

// Refresh If `true`, the request refreshes the relevant shards before retrieving the
// document.
// Setting it to `true` should be done after careful thought and verification
// that this does not cause a heavy load on the system (and slow down indexing).
// API name: refresh
func (r *Exists) Refresh(refresh bool) *Exists {
	r.values.Set("refresh", strconv.FormatBool(refresh))

	return r
}

// Routing A custom value used to route operations to a specific shard.
// API name: routing
func (r *Exists) Routing(routing string) *Exists {
	r.values.Set("routing", routing)

	return r
}

// Source_ Indicates whether to return the `_source` field (`true` or `false`) or lists
// the fields to return.
// API name: _source
func (r *Exists) Source_(sourceconfigparam string) *Exists {
	r.values.Set("_source", sourceconfigparam)

	return r
}

// SourceExcludes_ A comma-separated list of source fields to exclude from the response.
// You can also use this parameter to exclude fields from the subset specified
// in `_source_includes` query parameter.
// If the `_source` parameter is `false`, this parameter is ignored.
// API name: _source_excludes
func (r *Exists) SourceExcludes_(fields ...string) *Exists {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ A comma-separated list of source fields to include in the response.
// If this parameter is specified, only these source fields are returned.
// You can exclude fields from this subset using the `_source_excludes` query
// parameter.
// If the `_source` parameter is `false`, this parameter is ignored.
// API name: _source_includes
func (r *Exists) SourceIncludes_(fields ...string) *Exists {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// StoredFields A comma-separated list of stored fields to return as part of a hit.
// If no fields are specified, no stored fields are included in the response.
// If this field is specified, the `_source` parameter defaults to `false`.
// API name: stored_fields
func (r *Exists) StoredFields(fields ...string) *Exists {
	r.values.Set("stored_fields", strings.Join(fields, ","))

	return r
}

// Version Explicit version number for concurrency control.
// The specified version must match the current version of the document for the
// request to succeed.
// API name: version
func (r *Exists) Version(versionnumber string) *Exists {
	r.values.Set("version", versionnumber)

	return r
}

// VersionType The version type.
// API name: version_type
func (r *Exists) VersionType(versiontype versiontype.VersionType) *Exists {
	r.values.Set("version_type", versiontype.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Exists) ErrorTrace(errortrace bool) *Exists {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Exists) FilterPath(filterpaths ...string) *Exists {
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
func (r *Exists) Human(human bool) *Exists {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Exists) Pretty(pretty bool) *Exists {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
