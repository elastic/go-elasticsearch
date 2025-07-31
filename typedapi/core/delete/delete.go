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

// Delete a document.
//
// Remove a JSON document from the specified index.
//
// NOTE: You cannot send deletion requests directly to a data stream.
// To delete a document in a data stream, you must target the backing index
// containing the document.
//
// **Optimistic concurrency control**
//
// Delete operations can be made conditional and only be performed if the last
// modification to the document was assigned the sequence number and primary
// term specified by the `if_seq_no` and `if_primary_term` parameters.
// If a mismatch is detected, the operation will result in a
// `VersionConflictException` and a status code of `409`.
//
// **Versioning**
//
// Each document indexed is versioned.
// When deleting a document, the version can be specified to make sure the
// relevant document you are trying to delete is actually being deleted and it
// has not changed in the meantime.
// Every write operation run on a document, deletes included, causes its version
// to be incremented.
// The version number of a deleted document remains available for a short time
// after deletion to allow for control of concurrent operations.
// The length of time for which a deleted document's version remains available
// is determined by the `index.gc_deletes` index setting.
//
// **Routing**
//
// If routing is used during indexing, the routing value also needs to be
// specified to delete a document.
//
// If the `_routing` mapping is set to `required` and no routing value is
// specified, the delete API throws a `RoutingMissingException` and rejects the
// request.
//
// For example:
//
// ```
// DELETE /my-index-000001/_doc/1?routing=shard-1
// ```
//
// This request deletes the document with ID 1, but it is routed based on the
// user.
// The document is not deleted if the correct routing is not specified.
//
// **Distributed**
//
// The delete operation gets hashed into a specific shard ID.
// It then gets redirected into the primary shard within that ID group and
// replicated (if needed) to shard replicas within that ID group.
package delete

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

const (
	idMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Delete struct {
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

// NewDelete type alias for index.
type NewDelete func(index, id string) *Delete

// NewDeleteFunc returns a new instance of Delete with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteFunc(tp elastictransport.Interface) NewDelete {
	return func(index, id string) *Delete {
		n := New(tp)

		n._id(id)

		n._index(index)

		return n
	}
}

// Delete a document.
//
// Remove a JSON document from the specified index.
//
// NOTE: You cannot send deletion requests directly to a data stream.
// To delete a document in a data stream, you must target the backing index
// containing the document.
//
// **Optimistic concurrency control**
//
// Delete operations can be made conditional and only be performed if the last
// modification to the document was assigned the sequence number and primary
// term specified by the `if_seq_no` and `if_primary_term` parameters.
// If a mismatch is detected, the operation will result in a
// `VersionConflictException` and a status code of `409`.
//
// **Versioning**
//
// Each document indexed is versioned.
// When deleting a document, the version can be specified to make sure the
// relevant document you are trying to delete is actually being deleted and it
// has not changed in the meantime.
// Every write operation run on a document, deletes included, causes its version
// to be incremented.
// The version number of a deleted document remains available for a short time
// after deletion to allow for control of concurrent operations.
// The length of time for which a deleted document's version remains available
// is determined by the `index.gc_deletes` index setting.
//
// **Routing**
//
// If routing is used during indexing, the routing value also needs to be
// specified to delete a document.
//
// If the `_routing` mapping is set to `required` and no routing value is
// specified, the delete API throws a `RoutingMissingException` and rejects the
// request.
//
// For example:
//
// ```
// DELETE /my-index-000001/_doc/1?routing=shard-1
// ```
//
// This request deletes the document with ID 1, but it is routed based on the
// user.
// The document is not deleted if the correct routing is not specified.
//
// **Distributed**
//
// The delete operation gets hashed into a specific shard ID.
// It then gets redirected into the primary shard within that ID group and
// replicated (if needed) to shard replicas within that ID group.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete.html
func New(tp elastictransport.Interface) *Delete {
	r := &Delete{
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
func (r *Delete) HttpRequest(ctx context.Context) (*http.Request, error) {
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

		method = http.MethodDelete
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
func (r Delete) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "delete")
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
		instrument.BeforeRequest(req, "delete")
		if reader := instrument.RecordRequestBody(ctx, "delete", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "delete")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Delete query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a delete.Response
func (r Delete) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "delete")
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
func (r Delete) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "delete")
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
		err := fmt.Errorf("an error happened during the Delete query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the Delete headers map.
func (r *Delete) Header(key, value string) *Delete {
	r.headers.Set(key, value)

	return r
}

// Id A unique identifier for the document.
// API Name: id
func (r *Delete) _id(id string) *Delete {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Index The name of the target index.
// API Name: index
func (r *Delete) _index(index string) *Delete {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// IfPrimaryTerm Only perform the operation if the document has this primary term.
// API name: if_primary_term
func (r *Delete) IfPrimaryTerm(ifprimaryterm string) *Delete {
	r.values.Set("if_primary_term", ifprimaryterm)

	return r
}

// IfSeqNo Only perform the operation if the document has this sequence number.
// API name: if_seq_no
func (r *Delete) IfSeqNo(sequencenumber string) *Delete {
	r.values.Set("if_seq_no", sequencenumber)

	return r
}

// Refresh If `true`, Elasticsearch refreshes the affected shards to make this operation
// visible to search.
// If `wait_for`, it waits for a refresh to make this operation visible to
// search.
// If `false`, it does nothing with refreshes.
// API name: refresh
func (r *Delete) Refresh(refresh refresh.Refresh) *Delete {
	r.values.Set("refresh", refresh.String())

	return r
}

// Routing A custom value used to route operations to a specific shard.
// API name: routing
func (r *Delete) Routing(routing string) *Delete {
	r.values.Set("routing", routing)

	return r
}

// Timeout The period to wait for active shards.
//
// This parameter is useful for situations where the primary shard assigned to
// perform the delete operation might not be available when the delete operation
// runs.
// Some reasons for this might be that the primary shard is currently recovering
// from a store or undergoing relocation.
// By default, the delete operation will wait on the primary shard to become
// available for up to 1 minute before failing and responding with an error.
// API name: timeout
func (r *Delete) Timeout(duration string) *Delete {
	r.values.Set("timeout", duration)

	return r
}

// Version An explicit version number for concurrency control.
// It must match the current version of the document for the request to succeed.
// API name: version
func (r *Delete) Version(versionnumber string) *Delete {
	r.values.Set("version", versionnumber)

	return r
}

// VersionType The version type.
// API name: version_type
func (r *Delete) VersionType(versiontype versiontype.VersionType) *Delete {
	r.values.Set("version_type", versiontype.String())

	return r
}

// WaitForActiveShards The minimum number of shard copies that must be active before proceeding with
// the operation.
// You can set it to `all` or any positive integer up to the total number of
// shards in the index (`number_of_replicas+1`).
// The default value of `1` means it waits for each primary shard to be active.
// API name: wait_for_active_shards
func (r *Delete) WaitForActiveShards(waitforactiveshards string) *Delete {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Delete) ErrorTrace(errortrace bool) *Delete {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Delete) FilterPath(filterpaths ...string) *Delete {
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
func (r *Delete) Human(human bool) *Delete {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Delete) Pretty(pretty bool) *Delete {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
