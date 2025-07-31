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

// Create a new document in the index.
//
// You can index a new JSON document with the `/<target>/_doc/` or
// `/<target>/_create/<_id>` APIs
// Using `_create` guarantees that the document is indexed only if it does not
// already exist.
// It returns a 409 response when a document with a same ID already exists in
// the index.
// To update an existing document, you must use the `/<target>/_doc/` API.
//
// If the Elasticsearch security features are enabled, you must have the
// following index privileges for the target data stream, index, or index alias:
//
// * To add a document using the `PUT /<target>/_create/<_id>` or `POST
// /<target>/_create/<_id>` request formats, you must have the `create_doc`,
// `create`, `index`, or `write` index privilege.
// * To automatically create a data stream or index with this API request, you
// must have the `auto_configure`, `create_index`, or `manage` index privilege.
//
// Automatic data stream creation requires a matching index template with data
// stream enabled.
//
// **Automatically create data streams and indices**
//
// If the request's target doesn't exist and matches an index template with a
// `data_stream` definition, the index operation automatically creates the data
// stream.
//
// If the target doesn't exist and doesn't match a data stream template, the
// operation automatically creates the index and applies any matching index
// templates.
//
// NOTE: Elasticsearch includes several built-in index templates. To avoid
// naming collisions with these templates, refer to index pattern documentation.
//
// If no mapping exists, the index operation creates a dynamic mapping.
// By default, new fields and objects are automatically added to the mapping if
// needed.
//
// Automatic index creation is controlled by the `action.auto_create_index`
// setting.
// If it is `true`, any index can be created automatically.
// You can modify this setting to explicitly allow or block automatic creation
// of indices that match specified patterns or set it to `false` to turn off
// automatic index creation entirely.
// Specify a comma-separated list of patterns you want to allow or prefix each
// pattern with `+` or `-` to indicate whether it should be allowed or blocked.
// When a list is specified, the default behaviour is to disallow.
//
// NOTE: The `action.auto_create_index` setting affects the automatic creation
// of indices only.
// It does not affect the creation of data streams.
//
// **Routing**
//
// By default, shard placement — or routing — is controlled by using a hash of
// the document's ID value.
// For more explicit control, the value fed into the hash function used by the
// router can be directly specified on a per-operation basis using the `routing`
// parameter.
//
// When setting up explicit mapping, you can also use the `_routing` field to
// direct the index operation to extract the routing value from the document
// itself.
// This does come at the (very minimal) cost of an additional document parsing
// pass.
// If the `_routing` mapping is defined and set to be required, the index
// operation will fail if no routing value is provided or extracted.
//
// NOTE: Data streams do not support custom routing unless they were created
// with the `allow_custom_routing` setting enabled in the template.
//
// **Distributed**
//
// The index operation is directed to the primary shard based on its route and
// performed on the actual node containing this shard.
// After the primary shard completes the operation, if needed, the update is
// distributed to applicable replicas.
//
// **Active shards**
//
// To improve the resiliency of writes to the system, indexing operations can be
// configured to wait for a certain number of active shard copies before
// proceeding with the operation.
// If the requisite number of active shard copies are not available, then the
// write operation must wait and retry, until either the requisite shard copies
// have started or a timeout occurs.
// By default, write operations only wait for the primary shards to be active
// before proceeding (that is to say `wait_for_active_shards` is `1`).
// This default can be overridden in the index settings dynamically by setting
// `index.write.wait_for_active_shards`.
// To alter this behavior per operation, use the `wait_for_active_shards
// request` parameter.
//
// Valid values are all or any positive integer up to the total number of
// configured copies per shard in the index (which is `number_of_replicas`+1).
// Specifying a negative value or a number greater than the number of shard
// copies will throw an error.
//
// For example, suppose you have a cluster of three nodes, A, B, and C and you
// create an index index with the number of replicas set to 3 (resulting in 4
// shard copies, one more copy than there are nodes).
// If you attempt an indexing operation, by default the operation will only
// ensure the primary copy of each shard is available before proceeding.
// This means that even if B and C went down and A hosted the primary shard
// copies, the indexing operation would still proceed with only one copy of the
// data.
// If `wait_for_active_shards` is set on the request to `3` (and all three nodes
// are up), the indexing operation will require 3 active shard copies before
// proceeding.
// This requirement should be met because there are 3 active nodes in the
// cluster, each one holding a copy of the shard.
// However, if you set `wait_for_active_shards` to `all` (or to `4`, which is
// the same in this situation), the indexing operation will not proceed as you
// do not have all 4 copies of each shard active in the index.
// The operation will timeout unless a new node is brought up in the cluster to
// host the fourth copy of the shard.
//
// It is important to note that this setting greatly reduces the chances of the
// write operation not writing to the requisite number of shard copies, but it
// does not completely eliminate the possibility, because this check occurs
// before the write operation starts.
// After the write operation is underway, it is still possible for replication
// to fail on any number of shard copies but still succeed on the primary.
// The `_shards` section of the API response reveals the number of shard copies
// on which replication succeeded and failed.
package create

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

const (
	idMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Create struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      any
	deferred []func(request any) error
	buf      *gobytes.Buffer

	paramSet int

	id    string
	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewCreate type alias for index.
type NewCreate func(index, id string) *Create

// NewCreateFunc returns a new instance of Create with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCreateFunc(tp elastictransport.Interface) NewCreate {
	return func(index, id string) *Create {
		n := New(tp)

		n._id(id)

		n._index(index)

		return n
	}
}

// Create a new document in the index.
//
// You can index a new JSON document with the `/<target>/_doc/` or
// `/<target>/_create/<_id>` APIs
// Using `_create` guarantees that the document is indexed only if it does not
// already exist.
// It returns a 409 response when a document with a same ID already exists in
// the index.
// To update an existing document, you must use the `/<target>/_doc/` API.
//
// If the Elasticsearch security features are enabled, you must have the
// following index privileges for the target data stream, index, or index alias:
//
// * To add a document using the `PUT /<target>/_create/<_id>` or `POST
// /<target>/_create/<_id>` request formats, you must have the `create_doc`,
// `create`, `index`, or `write` index privilege.
// * To automatically create a data stream or index with this API request, you
// must have the `auto_configure`, `create_index`, or `manage` index privilege.
//
// Automatic data stream creation requires a matching index template with data
// stream enabled.
//
// **Automatically create data streams and indices**
//
// If the request's target doesn't exist and matches an index template with a
// `data_stream` definition, the index operation automatically creates the data
// stream.
//
// If the target doesn't exist and doesn't match a data stream template, the
// operation automatically creates the index and applies any matching index
// templates.
//
// NOTE: Elasticsearch includes several built-in index templates. To avoid
// naming collisions with these templates, refer to index pattern documentation.
//
// If no mapping exists, the index operation creates a dynamic mapping.
// By default, new fields and objects are automatically added to the mapping if
// needed.
//
// Automatic index creation is controlled by the `action.auto_create_index`
// setting.
// If it is `true`, any index can be created automatically.
// You can modify this setting to explicitly allow or block automatic creation
// of indices that match specified patterns or set it to `false` to turn off
// automatic index creation entirely.
// Specify a comma-separated list of patterns you want to allow or prefix each
// pattern with `+` or `-` to indicate whether it should be allowed or blocked.
// When a list is specified, the default behaviour is to disallow.
//
// NOTE: The `action.auto_create_index` setting affects the automatic creation
// of indices only.
// It does not affect the creation of data streams.
//
// **Routing**
//
// By default, shard placement — or routing — is controlled by using a hash of
// the document's ID value.
// For more explicit control, the value fed into the hash function used by the
// router can be directly specified on a per-operation basis using the `routing`
// parameter.
//
// When setting up explicit mapping, you can also use the `_routing` field to
// direct the index operation to extract the routing value from the document
// itself.
// This does come at the (very minimal) cost of an additional document parsing
// pass.
// If the `_routing` mapping is defined and set to be required, the index
// operation will fail if no routing value is provided or extracted.
//
// NOTE: Data streams do not support custom routing unless they were created
// with the `allow_custom_routing` setting enabled in the template.
//
// **Distributed**
//
// The index operation is directed to the primary shard based on its route and
// performed on the actual node containing this shard.
// After the primary shard completes the operation, if needed, the update is
// distributed to applicable replicas.
//
// **Active shards**
//
// To improve the resiliency of writes to the system, indexing operations can be
// configured to wait for a certain number of active shard copies before
// proceeding with the operation.
// If the requisite number of active shard copies are not available, then the
// write operation must wait and retry, until either the requisite shard copies
// have started or a timeout occurs.
// By default, write operations only wait for the primary shards to be active
// before proceeding (that is to say `wait_for_active_shards` is `1`).
// This default can be overridden in the index settings dynamically by setting
// `index.write.wait_for_active_shards`.
// To alter this behavior per operation, use the `wait_for_active_shards
// request` parameter.
//
// Valid values are all or any positive integer up to the total number of
// configured copies per shard in the index (which is `number_of_replicas`+1).
// Specifying a negative value or a number greater than the number of shard
// copies will throw an error.
//
// For example, suppose you have a cluster of three nodes, A, B, and C and you
// create an index index with the number of replicas set to 3 (resulting in 4
// shard copies, one more copy than there are nodes).
// If you attempt an indexing operation, by default the operation will only
// ensure the primary copy of each shard is available before proceeding.
// This means that even if B and C went down and A hosted the primary shard
// copies, the indexing operation would still proceed with only one copy of the
// data.
// If `wait_for_active_shards` is set on the request to `3` (and all three nodes
// are up), the indexing operation will require 3 active shard copies before
// proceeding.
// This requirement should be met because there are 3 active nodes in the
// cluster, each one holding a copy of the shard.
// However, if you set `wait_for_active_shards` to `all` (or to `4`, which is
// the same in this situation), the indexing operation will not proceed as you
// do not have all 4 copies of each shard active in the index.
// The operation will timeout unless a new node is brought up in the cluster to
// host the fourth copy of the shard.
//
// It is important to note that this setting greatly reduces the chances of the
// write operation not writing to the requisite number of shard copies, but it
// does not completely eliminate the possibility, because this check occurs
// before the write operation starts.
// After the write operation is underway, it is still possible for replication
// to fail on any number of shard copies but still succeed on the primary.
// The `_shards` section of the API response reveals the number of shard copies
// on which replication succeeded and failed.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
func New(tp elastictransport.Interface) *Create {
	r := &Create{
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
func (r *Create) Raw(raw io.Reader) *Create {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Create) Request(req any) *Create {
	r.req = req

	return r
}

// Document allows to set the request property with the appropriate payload.
func (r *Create) Document(document any) *Create {
	r.req = document

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Create) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Create: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|idMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_create")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)

		method = http.MethodPut
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
func (r Create) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "create")
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
		instrument.BeforeRequest(req, "create")
		if reader := instrument.RecordRequestBody(ctx, "create", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "create")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Create query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a create.Response
func (r Create) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "create")
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

// Header set a key, value pair in the Create headers map.
func (r *Create) Header(key, value string) *Create {
	r.headers.Set(key, value)

	return r
}

// Id A unique identifier for the document.
// To automatically generate a document ID, use the `POST /<target>/_doc/`
// request format.
// API Name: id
func (r *Create) _id(id string) *Create {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Index The name of the data stream or index to target.
// If the target doesn't exist and matches the name or wildcard (`*`) pattern of
// an index template with a `data_stream` definition, this request creates the
// data stream.
// If the target doesn't exist and doesn’t match a data stream template, this
// request creates the index.
// API Name: index
func (r *Create) _index(index string) *Create {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// IncludeSourceOnError True or false if to include the document source in the error message in case
// of parsing errors.
// API name: include_source_on_error
func (r *Create) IncludeSourceOnError(includesourceonerror bool) *Create {
	r.values.Set("include_source_on_error", strconv.FormatBool(includesourceonerror))

	return r
}

// Pipeline The ID of the pipeline to use to preprocess incoming documents.
// If the index has a default ingest pipeline specified, setting the value to
// `_none` turns off the default ingest pipeline for this request.
// If a final pipeline is configured, it will always run regardless of the value
// of this parameter.
// API name: pipeline
func (r *Create) Pipeline(pipeline string) *Create {
	r.values.Set("pipeline", pipeline)

	return r
}

// Refresh If `true`, Elasticsearch refreshes the affected shards to make this operation
// visible to search.
// If `wait_for`, it waits for a refresh to make this operation visible to
// search.
// If `false`, it does nothing with refreshes.
// API name: refresh
func (r *Create) Refresh(refresh refresh.Refresh) *Create {
	r.values.Set("refresh", refresh.String())

	return r
}

// RequireAlias If `true`, the destination must be an index alias.
// API name: require_alias
func (r *Create) RequireAlias(requirealias bool) *Create {
	r.values.Set("require_alias", strconv.FormatBool(requirealias))

	return r
}

// RequireDataStream If `true`, the request's actions must target a data stream (existing or to be
// created).
// API name: require_data_stream
func (r *Create) RequireDataStream(requiredatastream bool) *Create {
	r.values.Set("require_data_stream", strconv.FormatBool(requiredatastream))

	return r
}

// Routing A custom value that is used to route operations to a specific shard.
// API name: routing
func (r *Create) Routing(routing string) *Create {
	r.values.Set("routing", routing)

	return r
}

// Timeout The period the request waits for the following operations: automatic index
// creation, dynamic mapping updates, waiting for active shards.
// Elasticsearch waits for at least the specified timeout period before failing.
// The actual wait time could be longer, particularly when multiple waits occur.
//
// This parameter is useful for situations where the primary shard assigned to
// perform the operation might not be available when the operation runs.
// Some reasons for this might be that the primary shard is currently recovering
// from a gateway or undergoing relocation.
// By default, the operation will wait on the primary shard to become available
// for at least 1 minute before failing and responding with an error.
// The actual wait time could be longer, particularly when multiple waits occur.
// API name: timeout
func (r *Create) Timeout(duration string) *Create {
	r.values.Set("timeout", duration)

	return r
}

// Version The explicit version number for concurrency control.
// It must be a non-negative long number.
// API name: version
func (r *Create) Version(versionnumber string) *Create {
	r.values.Set("version", versionnumber)

	return r
}

// VersionType The version type.
// API name: version_type
func (r *Create) VersionType(versiontype versiontype.VersionType) *Create {
	r.values.Set("version_type", versiontype.String())

	return r
}

// WaitForActiveShards The number of shard copies that must be active before proceeding with the
// operation.
// You can set it to `all` or any positive integer up to the total number of
// shards in the index (`number_of_replicas+1`).
// The default value of `1` means it waits for each primary shard to be active.
// API name: wait_for_active_shards
func (r *Create) WaitForActiveShards(waitforactiveshards string) *Create {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Create) ErrorTrace(errortrace bool) *Create {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Create) FilterPath(filterpaths ...string) *Create {
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
func (r *Create) Human(human bool) *Create {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Create) Pretty(pretty bool) *Create {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
