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

// Get a document by its ID.
//
// Get a document and its source or stored fields from an index.
//
// By default, this API is realtime and is not affected by the refresh rate of
// the index (when data will become visible for search).
// In the case where stored fields are requested with the `stored_fields`
// parameter and the document has been updated but is not yet refreshed, the API
// will have to parse and analyze the source to extract the stored fields.
// To turn off realtime behavior, set the `realtime` parameter to false.
//
// **Source filtering**
//
// By default, the API returns the contents of the `_source` field unless you
// have used the `stored_fields` parameter or the `_source` field is turned off.
// You can turn off `_source` retrieval by using the `_source` parameter:
//
// ```
// GET my-index-000001/_doc/0?_source=false
// ```
//
// If you only need one or two fields from the `_source`, use the
// `_source_includes` or `_source_excludes` parameters to include or filter out
// particular fields.
// This can be helpful with large documents where partial retrieval can save on
// network overhead
// Both parameters take a comma separated list of fields or wildcard
// expressions.
// For example:
//
// ```
// GET my-index-000001/_doc/0?_source_includes=*.id&_source_excludes=entities
// ```
//
// If you only want to specify includes, you can use a shorter notation:
//
// ```
// GET my-index-000001/_doc/0?_source=*.id
// ```
//
// **Routing**
//
// If routing is used during indexing, the routing value also needs to be
// specified to retrieve a document.
// For example:
//
// ```
// GET my-index-000001/_doc/2?routing=user1
// ```
//
// This request gets the document with ID 2, but it is routed based on the user.
// The document is not fetched if the correct routing is not specified.
//
// **Distributed**
//
// The GET operation is hashed into a specific shard ID.
// It is then redirected to one of the replicas within that shard ID and returns
// the result.
// The replicas are the primary shard and its replicas within that shard ID
// group.
// This means that the more replicas you have, the better your GET scaling will
// be.
//
// **Versioning support**
//
// You can use the `version` parameter to retrieve the document only if its
// current version is equal to the specified one.
//
// Internally, Elasticsearch has marked the old document as deleted and added an
// entirely new document.
// The old version of the document doesn't disappear immediately, although you
// won't be able to access it.
// Elasticsearch cleans up deleted documents in the background as you continue
// to index more data.
package get

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

const (
	idMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Get struct {
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

// NewGet type alias for index.
type NewGet func(index, id string) *Get

// NewGetFunc returns a new instance of Get with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetFunc(tp elastictransport.Interface) NewGet {
	return func(index, id string) *Get {
		n := New(tp)

		n._id(id)

		n._index(index)

		return n
	}
}

// Get a document by its ID.
//
// Get a document and its source or stored fields from an index.
//
// By default, this API is realtime and is not affected by the refresh rate of
// the index (when data will become visible for search).
// In the case where stored fields are requested with the `stored_fields`
// parameter and the document has been updated but is not yet refreshed, the API
// will have to parse and analyze the source to extract the stored fields.
// To turn off realtime behavior, set the `realtime` parameter to false.
//
// **Source filtering**
//
// By default, the API returns the contents of the `_source` field unless you
// have used the `stored_fields` parameter or the `_source` field is turned off.
// You can turn off `_source` retrieval by using the `_source` parameter:
//
// ```
// GET my-index-000001/_doc/0?_source=false
// ```
//
// If you only need one or two fields from the `_source`, use the
// `_source_includes` or `_source_excludes` parameters to include or filter out
// particular fields.
// This can be helpful with large documents where partial retrieval can save on
// network overhead
// Both parameters take a comma separated list of fields or wildcard
// expressions.
// For example:
//
// ```
// GET my-index-000001/_doc/0?_source_includes=*.id&_source_excludes=entities
// ```
//
// If you only want to specify includes, you can use a shorter notation:
//
// ```
// GET my-index-000001/_doc/0?_source=*.id
// ```
//
// **Routing**
//
// If routing is used during indexing, the routing value also needs to be
// specified to retrieve a document.
// For example:
//
// ```
// GET my-index-000001/_doc/2?routing=user1
// ```
//
// This request gets the document with ID 2, but it is routed based on the user.
// The document is not fetched if the correct routing is not specified.
//
// **Distributed**
//
// The GET operation is hashed into a specific shard ID.
// It is then redirected to one of the replicas within that shard ID and returns
// the result.
// The replicas are the primary shard and its replicas within that shard ID
// group.
// This means that the more replicas you have, the better your GET scaling will
// be.
//
// **Versioning support**
//
// You can use the `version` parameter to retrieve the document only if its
// current version is equal to the specified one.
//
// Internally, Elasticsearch has marked the old document as deleted and added an
// entirely new document.
// The old version of the document doesn't disappear immediately, although you
// won't be able to access it.
// Elasticsearch cleans up deleted documents in the background as you continue
// to index more data.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
func New(tp elastictransport.Interface) *Get {
	r := &Get{
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
func (r *Get) HttpRequest(ctx context.Context) (*http.Request, error) {
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

		method = http.MethodGet
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
func (r Get) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "get")
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
		instrument.BeforeRequest(req, "get")
		if reader := instrument.RecordRequestBody(ctx, "get", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "get")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Get query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a get.Response
func (r Get) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "get")
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

	if res.StatusCode < 299 || slices.Contains([]int{404}, res.StatusCode) {

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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r Get) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "get")
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
		err := fmt.Errorf("an error happened during the Get query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the Get headers map.
func (r *Get) Header(key, value string) *Get {
	r.headers.Set(key, value)

	return r
}

// Id A unique document identifier.
// API Name: id
func (r *Get) _id(id string) *Get {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Index The name of the index that contains the document.
// API Name: index
func (r *Get) _index(index string) *Get {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// ForceSyntheticSource Indicates whether the request forces synthetic `_source`.
// Use this paramater to test if the mapping supports synthetic `_source` and to
// get a sense of the worst case performance.
// Fetches with this parameter enabled will be slower than enabling synthetic
// source natively in the index.
// API name: force_synthetic_source
func (r *Get) ForceSyntheticSource(forcesyntheticsource bool) *Get {
	r.values.Set("force_synthetic_source", strconv.FormatBool(forcesyntheticsource))

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
func (r *Get) Preference(preference string) *Get {
	r.values.Set("preference", preference)

	return r
}

// Realtime If `true`, the request is real-time as opposed to near-real-time.
// API name: realtime
func (r *Get) Realtime(realtime bool) *Get {
	r.values.Set("realtime", strconv.FormatBool(realtime))

	return r
}

// Refresh If `true`, the request refreshes the relevant shards before retrieving the
// document.
// Setting it to `true` should be done after careful thought and verification
// that this does not cause a heavy load on the system (and slow down indexing).
// API name: refresh
func (r *Get) Refresh(refresh bool) *Get {
	r.values.Set("refresh", strconv.FormatBool(refresh))

	return r
}

// Routing A custom value used to route operations to a specific shard.
// API name: routing
func (r *Get) Routing(routing string) *Get {
	r.values.Set("routing", routing)

	return r
}

// Source_ Indicates whether to return the `_source` field (`true` or `false`) or lists
// the fields to return.
// API name: _source
func (r *Get) Source_(sourceconfigparam string) *Get {
	r.values.Set("_source", sourceconfigparam)

	return r
}

// SourceExcludes_ A comma-separated list of source fields to exclude from the response.
// You can also use this parameter to exclude fields from the subset specified
// in `_source_includes` query parameter.
// If the `_source` parameter is `false`, this parameter is ignored.
// API name: _source_excludes
func (r *Get) SourceExcludes_(fields ...string) *Get {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ A comma-separated list of source fields to include in the response.
// If this parameter is specified, only these source fields are returned.
// You can exclude fields from this subset using the `_source_excludes` query
// parameter.
// If the `_source` parameter is `false`, this parameter is ignored.
// API name: _source_includes
func (r *Get) SourceIncludes_(fields ...string) *Get {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// StoredFields A comma-separated list of stored fields to return as part of a hit.
// If no fields are specified, no stored fields are included in the response.
// If this field is specified, the `_source` parameter defaults to `false`.
// Only leaf fields can be retrieved with the `stored_field` option.
// Object fields can't be returned;​if specified, the request fails.
// API name: stored_fields
func (r *Get) StoredFields(fields ...string) *Get {
	r.values.Set("stored_fields", strings.Join(fields, ","))

	return r
}

// Version The version number for concurrency control.
// It must match the current version of the document for the request to succeed.
// API name: version
func (r *Get) Version(versionnumber string) *Get {
	r.values.Set("version", versionnumber)

	return r
}

// VersionType The version type.
// API name: version_type
func (r *Get) VersionType(versiontype versiontype.VersionType) *Get {
	r.values.Set("version_type", versiontype.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Get) ErrorTrace(errortrace bool) *Get {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Get) FilterPath(filterpaths ...string) *Get {
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
func (r *Get) Human(human bool) *Get {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Get) Pretty(pretty bool) *Get {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
