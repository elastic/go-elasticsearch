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

// Reindex documents.
//
// Copy documents from a source to a destination.
// You can copy all documents to the destination index or reindex a subset of
// the documents.
// The source can be any existing index, alias, or data stream.
// The destination must differ from the source.
// For example, you cannot reindex a data stream into itself.
//
// IMPORTANT: Reindex requires `_source` to be enabled for all documents in the
// source.
// The destination should be configured as wanted before calling the reindex
// API.
// Reindex does not copy the settings from the source or its associated
// template.
// Mappings, shard counts, and replicas, for example, must be configured ahead
// of time.
//
// If the Elasticsearch security features are enabled, you must have the
// following security privileges:
//
// * The `read` index privilege for the source data stream, index, or alias.
// * The `write` index privilege for the destination data stream, index, or
// index alias.
// * To automatically create a data stream or index with a reindex API request,
// you must have the `auto_configure`, `create_index`, or `manage` index
// privilege for the destination data stream, index, or alias.
// * If reindexing from a remote cluster, the `source.remote.user` must have the
// `monitor` cluster privilege and the `read` index privilege for the source
// data stream, index, or alias.
//
// If reindexing from a remote cluster, you must explicitly allow the remote
// host in the `reindex.remote.whitelist` setting.
// Automatic data stream creation requires a matching index template with data
// stream enabled.
//
// The `dest` element can be configured like the index API to control optimistic
// concurrency control.
// Omitting `version_type` or setting it to `internal` causes Elasticsearch to
// blindly dump documents into the destination, overwriting any that happen to
// have the same ID.
//
// Setting `version_type` to `external` causes Elasticsearch to preserve the
// `version` from the source, create any documents that are missing, and update
// any documents that have an older version in the destination than they do in
// the source.
//
// Setting `op_type` to `create` causes the reindex API to create only missing
// documents in the destination.
// All existing documents will cause a version conflict.
//
// IMPORTANT: Because data streams are append-only, any reindex request to a
// destination data stream must have an `op_type` of `create`.
// A reindex can only add new documents to a destination data stream.
// It cannot update existing documents in a destination data stream.
//
// By default, version conflicts abort the reindex process.
// To continue reindexing if there are conflicts, set the `conflicts` request
// body property to `proceed`.
// In this case, the response includes a count of the version conflicts that
// were encountered.
// Note that the handling of other error types is unaffected by the `conflicts`
// property.
// Additionally, if you opt to count version conflicts, the operation could
// attempt to reindex more documents from the source than `max_docs` until it
// has successfully indexed `max_docs` documents into the target or it has gone
// through every document in the source query.
//
// NOTE: The reindex API makes no effort to handle ID collisions.
// The last document written will "win" but the order isn't usually predictable
// so it is not a good idea to rely on this behavior.
// Instead, make sure that IDs are unique by using a script.
//
// **Running reindex asynchronously**
//
// If the request contains `wait_for_completion=false`, Elasticsearch performs
// some preflight checks, launches the request, and returns a task you can use
// to cancel or get the status of the task.
// Elasticsearch creates a record of this task as a document at
// `_tasks/<task_id>`.
//
// **Reindex from multiple sources**
//
// If you have many sources to reindex it is generally better to reindex them
// one at a time rather than using a glob pattern to pick up multiple sources.
// That way you can resume the process if there are any errors by removing the
// partially completed source and starting over.
// It also makes parallelizing the process fairly simple: split the list of
// sources to reindex and run each list in parallel.
//
// For example, you can use a bash script like this:
//
// ```
// for index in i1 i2 i3 i4 i5; do
//
//	curl -HContent-Type:application/json -XPOST localhost:9200/_reindex?pretty
//
//	-d'{
//	    "source": {
//	      "index": "'$index'"
//	    },
//	    "dest": {
//	      "index": "'$index'-reindexed"
//	    }
//	  }'
//
// done
// ```
//
// **Throttling**
//
// Set `requests_per_second` to any positive decimal number (`1.4`, `6`, `1000`,
// for example) to throttle the rate at which reindex issues batches of index
// operations.
// Requests are throttled by padding each batch with a wait time.
// To turn off throttling, set `requests_per_second` to `-1`.
//
// The throttling is done by waiting between batches so that the scroll that
// reindex uses internally can be given a timeout that takes into account the
// padding.
// The padding time is the difference between the batch size divided by the
// `requests_per_second` and the time spent writing.
// By default the batch size is `1000`, so if `requests_per_second` is set to
// `500`:
//
// ```
// target_time = 1000 / 500 per second = 2 seconds
// wait_time = target_time - write_time = 2 seconds - .5 seconds = 1.5 seconds
// ```
//
// Since the batch is issued as a single bulk request, large batch sizes cause
// Elasticsearch to create many requests and then wait for a while before
// starting the next set.
// This is "bursty" instead of "smooth".
//
// **Slicing**
//
// Reindex supports sliced scroll to parallelize the reindexing process.
// This parallelization can improve efficiency and provide a convenient way to
// break the request down into smaller parts.
//
// NOTE: Reindexing from remote clusters does not support manual or automatic
// slicing.
//
// You can slice a reindex request manually by providing a slice ID and total
// number of slices to each request.
// You can also let reindex automatically parallelize by using sliced scroll to
// slice on `_id`.
// The `slices` parameter specifies the number of slices to use.
//
// Adding `slices` to the reindex request just automates the manual process,
// creating sub-requests which means it has some quirks:
//
// * You can see these requests in the tasks API. These sub-requests are "child"
// tasks of the task for the request with slices.
// * Fetching the status of the task for the request with `slices` only contains
// the status of completed slices.
// * These sub-requests are individually addressable for things like
// cancellation and rethrottling.
// * Rethrottling the request with `slices` will rethrottle the unfinished
// sub-request proportionally.
// * Canceling the request with `slices` will cancel each sub-request.
// * Due to the nature of `slices`, each sub-request won't get a perfectly even
// portion of the documents. All documents will be addressed, but some slices
// may be larger than others. Expect larger slices to have a more even
// distribution.
// * Parameters like `requests_per_second` and `max_docs` on a request with
// `slices` are distributed proportionally to each sub-request. Combine that
// with the previous point about distribution being uneven and you should
// conclude that using `max_docs` with `slices` might not result in exactly
// `max_docs` documents being reindexed.
// * Each sub-request gets a slightly different snapshot of the source, though
// these are all taken at approximately the same time.
//
// If slicing automatically, setting `slices` to `auto` will choose a reasonable
// number for most indices.
// If slicing manually or otherwise tuning automatic slicing, use the following
// guidelines.
//
// Query performance is most efficient when the number of slices is equal to the
// number of shards in the index.
// If that number is large (for example, `500`), choose a lower number as too
// many slices will hurt performance.
// Setting slices higher than the number of shards generally does not improve
// efficiency and adds overhead.
//
// Indexing performance scales linearly across available resources with the
// number of slices.
//
// Whether query or indexing performance dominates the runtime depends on the
// documents being reindexed and cluster resources.
//
// **Modify documents during reindexing**
//
// Like `_update_by_query`, reindex operations support a script that modifies
// the document.
// Unlike `_update_by_query`, the script is allowed to modify the document's
// metadata.
//
// Just as in `_update_by_query`, you can set `ctx.op` to change the operation
// that is run on the destination.
// For example, set `ctx.op` to `noop` if your script decides that the document
// doesn’t have to be indexed in the destination. This "no operation" will be
// reported in the `noop` counter in the response body.
// Set `ctx.op` to `delete` if your script decides that the document must be
// deleted from the destination.
// The deletion will be reported in the `deleted` counter in the response body.
// Setting `ctx.op` to anything else will return an error, as will setting any
// other field in `ctx`.
//
// Think of the possibilities! Just be careful; you are able to change:
//
// * `_id`
// * `_index`
// * `_version`
// * `_routing`
//
// Setting `_version` to `null` or clearing it from the `ctx` map is just like
// not sending the version in an indexing request.
// It will cause the document to be overwritten in the destination regardless of
// the version on the target or the version type you use in the reindex API.
//
// **Reindex from remote**
//
// Reindex supports reindexing from a remote Elasticsearch cluster.
// The `host` parameter must contain a scheme, host, port, and optional path.
// The `username` and `password` parameters are optional and when they are
// present the reindex operation will connect to the remote Elasticsearch node
// using basic authentication.
// Be sure to use HTTPS when using basic authentication or the password will be
// sent in plain text.
// There are a range of settings available to configure the behavior of the
// HTTPS connection.
//
// When using Elastic Cloud, it is also possible to authenticate against the
// remote cluster through the use of a valid API key.
// Remote hosts must be explicitly allowed with the `reindex.remote.whitelist`
// setting.
// It can be set to a comma delimited list of allowed remote host and port
// combinations.
// Scheme is ignored; only the host and port are used.
// For example:
//
// ```
// reindex.remote.whitelist: [otherhost:9200, another:9200, 127.0.10.*:9200,
// localhost:*"]
// ```
//
// The list of allowed hosts must be configured on any nodes that will
// coordinate the reindex.
// This feature should work with remote clusters of any version of
// Elasticsearch.
// This should enable you to upgrade from any version of Elasticsearch to the
// current version by reindexing from a cluster of the old version.
//
// WARNING: Elasticsearch does not support forward compatibility across major
// versions.
// For example, you cannot reindex from a 7.x cluster into a 6.x cluster.
//
// To enable queries sent to older versions of Elasticsearch, the `query`
// parameter is sent directly to the remote host without validation or
// modification.
//
// NOTE: Reindexing from remote clusters does not support manual or automatic
// slicing.
//
// Reindexing from a remote server uses an on-heap buffer that defaults to a
// maximum size of 100mb.
// If the remote index includes very large documents you'll need to use a
// smaller batch size.
// It is also possible to set the socket read timeout on the remote connection
// with the `socket_timeout` field and the connection timeout with the
// `connect_timeout` field.
// Both default to 30 seconds.
//
// **Configuring SSL parameters**
//
// Reindex from remote supports configurable SSL settings.
// These must be specified in the `elasticsearch.yml` file, with the exception
// of the secure settings, which you add in the Elasticsearch keystore.
// It is not possible to configure SSL in the body of the reindex request.
package reindex

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conflicts"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Reindex struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewReindex type alias for index.
type NewReindex func() *Reindex

// NewReindexFunc returns a new instance of Reindex with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewReindexFunc(tp elastictransport.Interface) NewReindex {
	return func() *Reindex {
		n := New(tp)

		return n
	}
}

// Reindex documents.
//
// Copy documents from a source to a destination.
// You can copy all documents to the destination index or reindex a subset of
// the documents.
// The source can be any existing index, alias, or data stream.
// The destination must differ from the source.
// For example, you cannot reindex a data stream into itself.
//
// IMPORTANT: Reindex requires `_source` to be enabled for all documents in the
// source.
// The destination should be configured as wanted before calling the reindex
// API.
// Reindex does not copy the settings from the source or its associated
// template.
// Mappings, shard counts, and replicas, for example, must be configured ahead
// of time.
//
// If the Elasticsearch security features are enabled, you must have the
// following security privileges:
//
// * The `read` index privilege for the source data stream, index, or alias.
// * The `write` index privilege for the destination data stream, index, or
// index alias.
// * To automatically create a data stream or index with a reindex API request,
// you must have the `auto_configure`, `create_index`, or `manage` index
// privilege for the destination data stream, index, or alias.
// * If reindexing from a remote cluster, the `source.remote.user` must have the
// `monitor` cluster privilege and the `read` index privilege for the source
// data stream, index, or alias.
//
// If reindexing from a remote cluster, you must explicitly allow the remote
// host in the `reindex.remote.whitelist` setting.
// Automatic data stream creation requires a matching index template with data
// stream enabled.
//
// The `dest` element can be configured like the index API to control optimistic
// concurrency control.
// Omitting `version_type` or setting it to `internal` causes Elasticsearch to
// blindly dump documents into the destination, overwriting any that happen to
// have the same ID.
//
// Setting `version_type` to `external` causes Elasticsearch to preserve the
// `version` from the source, create any documents that are missing, and update
// any documents that have an older version in the destination than they do in
// the source.
//
// Setting `op_type` to `create` causes the reindex API to create only missing
// documents in the destination.
// All existing documents will cause a version conflict.
//
// IMPORTANT: Because data streams are append-only, any reindex request to a
// destination data stream must have an `op_type` of `create`.
// A reindex can only add new documents to a destination data stream.
// It cannot update existing documents in a destination data stream.
//
// By default, version conflicts abort the reindex process.
// To continue reindexing if there are conflicts, set the `conflicts` request
// body property to `proceed`.
// In this case, the response includes a count of the version conflicts that
// were encountered.
// Note that the handling of other error types is unaffected by the `conflicts`
// property.
// Additionally, if you opt to count version conflicts, the operation could
// attempt to reindex more documents from the source than `max_docs` until it
// has successfully indexed `max_docs` documents into the target or it has gone
// through every document in the source query.
//
// NOTE: The reindex API makes no effort to handle ID collisions.
// The last document written will "win" but the order isn't usually predictable
// so it is not a good idea to rely on this behavior.
// Instead, make sure that IDs are unique by using a script.
//
// **Running reindex asynchronously**
//
// If the request contains `wait_for_completion=false`, Elasticsearch performs
// some preflight checks, launches the request, and returns a task you can use
// to cancel or get the status of the task.
// Elasticsearch creates a record of this task as a document at
// `_tasks/<task_id>`.
//
// **Reindex from multiple sources**
//
// If you have many sources to reindex it is generally better to reindex them
// one at a time rather than using a glob pattern to pick up multiple sources.
// That way you can resume the process if there are any errors by removing the
// partially completed source and starting over.
// It also makes parallelizing the process fairly simple: split the list of
// sources to reindex and run each list in parallel.
//
// For example, you can use a bash script like this:
//
// ```
// for index in i1 i2 i3 i4 i5; do
//
//	curl -HContent-Type:application/json -XPOST localhost:9200/_reindex?pretty
//
//	-d'{
//	    "source": {
//	      "index": "'$index'"
//	    },
//	    "dest": {
//	      "index": "'$index'-reindexed"
//	    }
//	  }'
//
// done
// ```
//
// **Throttling**
//
// Set `requests_per_second` to any positive decimal number (`1.4`, `6`, `1000`,
// for example) to throttle the rate at which reindex issues batches of index
// operations.
// Requests are throttled by padding each batch with a wait time.
// To turn off throttling, set `requests_per_second` to `-1`.
//
// The throttling is done by waiting between batches so that the scroll that
// reindex uses internally can be given a timeout that takes into account the
// padding.
// The padding time is the difference between the batch size divided by the
// `requests_per_second` and the time spent writing.
// By default the batch size is `1000`, so if `requests_per_second` is set to
// `500`:
//
// ```
// target_time = 1000 / 500 per second = 2 seconds
// wait_time = target_time - write_time = 2 seconds - .5 seconds = 1.5 seconds
// ```
//
// Since the batch is issued as a single bulk request, large batch sizes cause
// Elasticsearch to create many requests and then wait for a while before
// starting the next set.
// This is "bursty" instead of "smooth".
//
// **Slicing**
//
// Reindex supports sliced scroll to parallelize the reindexing process.
// This parallelization can improve efficiency and provide a convenient way to
// break the request down into smaller parts.
//
// NOTE: Reindexing from remote clusters does not support manual or automatic
// slicing.
//
// You can slice a reindex request manually by providing a slice ID and total
// number of slices to each request.
// You can also let reindex automatically parallelize by using sliced scroll to
// slice on `_id`.
// The `slices` parameter specifies the number of slices to use.
//
// Adding `slices` to the reindex request just automates the manual process,
// creating sub-requests which means it has some quirks:
//
// * You can see these requests in the tasks API. These sub-requests are "child"
// tasks of the task for the request with slices.
// * Fetching the status of the task for the request with `slices` only contains
// the status of completed slices.
// * These sub-requests are individually addressable for things like
// cancellation and rethrottling.
// * Rethrottling the request with `slices` will rethrottle the unfinished
// sub-request proportionally.
// * Canceling the request with `slices` will cancel each sub-request.
// * Due to the nature of `slices`, each sub-request won't get a perfectly even
// portion of the documents. All documents will be addressed, but some slices
// may be larger than others. Expect larger slices to have a more even
// distribution.
// * Parameters like `requests_per_second` and `max_docs` on a request with
// `slices` are distributed proportionally to each sub-request. Combine that
// with the previous point about distribution being uneven and you should
// conclude that using `max_docs` with `slices` might not result in exactly
// `max_docs` documents being reindexed.
// * Each sub-request gets a slightly different snapshot of the source, though
// these are all taken at approximately the same time.
//
// If slicing automatically, setting `slices` to `auto` will choose a reasonable
// number for most indices.
// If slicing manually or otherwise tuning automatic slicing, use the following
// guidelines.
//
// Query performance is most efficient when the number of slices is equal to the
// number of shards in the index.
// If that number is large (for example, `500`), choose a lower number as too
// many slices will hurt performance.
// Setting slices higher than the number of shards generally does not improve
// efficiency and adds overhead.
//
// Indexing performance scales linearly across available resources with the
// number of slices.
//
// Whether query or indexing performance dominates the runtime depends on the
// documents being reindexed and cluster resources.
//
// **Modify documents during reindexing**
//
// Like `_update_by_query`, reindex operations support a script that modifies
// the document.
// Unlike `_update_by_query`, the script is allowed to modify the document's
// metadata.
//
// Just as in `_update_by_query`, you can set `ctx.op` to change the operation
// that is run on the destination.
// For example, set `ctx.op` to `noop` if your script decides that the document
// doesn’t have to be indexed in the destination. This "no operation" will be
// reported in the `noop` counter in the response body.
// Set `ctx.op` to `delete` if your script decides that the document must be
// deleted from the destination.
// The deletion will be reported in the `deleted` counter in the response body.
// Setting `ctx.op` to anything else will return an error, as will setting any
// other field in `ctx`.
//
// Think of the possibilities! Just be careful; you are able to change:
//
// * `_id`
// * `_index`
// * `_version`
// * `_routing`
//
// Setting `_version` to `null` or clearing it from the `ctx` map is just like
// not sending the version in an indexing request.
// It will cause the document to be overwritten in the destination regardless of
// the version on the target or the version type you use in the reindex API.
//
// **Reindex from remote**
//
// Reindex supports reindexing from a remote Elasticsearch cluster.
// The `host` parameter must contain a scheme, host, port, and optional path.
// The `username` and `password` parameters are optional and when they are
// present the reindex operation will connect to the remote Elasticsearch node
// using basic authentication.
// Be sure to use HTTPS when using basic authentication or the password will be
// sent in plain text.
// There are a range of settings available to configure the behavior of the
// HTTPS connection.
//
// When using Elastic Cloud, it is also possible to authenticate against the
// remote cluster through the use of a valid API key.
// Remote hosts must be explicitly allowed with the `reindex.remote.whitelist`
// setting.
// It can be set to a comma delimited list of allowed remote host and port
// combinations.
// Scheme is ignored; only the host and port are used.
// For example:
//
// ```
// reindex.remote.whitelist: [otherhost:9200, another:9200, 127.0.10.*:9200,
// localhost:*"]
// ```
//
// The list of allowed hosts must be configured on any nodes that will
// coordinate the reindex.
// This feature should work with remote clusters of any version of
// Elasticsearch.
// This should enable you to upgrade from any version of Elasticsearch to the
// current version by reindexing from a cluster of the old version.
//
// WARNING: Elasticsearch does not support forward compatibility across major
// versions.
// For example, you cannot reindex from a 7.x cluster into a 6.x cluster.
//
// To enable queries sent to older versions of Elasticsearch, the `query`
// parameter is sent directly to the remote host without validation or
// modification.
//
// NOTE: Reindexing from remote clusters does not support manual or automatic
// slicing.
//
// Reindexing from a remote server uses an on-heap buffer that defaults to a
// maximum size of 100mb.
// If the remote index includes very large documents you'll need to use a
// smaller batch size.
// It is also possible to set the socket read timeout on the remote connection
// with the `socket_timeout` field and the connection timeout with the
// `connect_timeout` field.
// Both default to 30 seconds.
//
// **Configuring SSL parameters**
//
// Reindex from remote supports configurable SSL settings.
// These must be specified in the `elasticsearch.yml` file, with the exception
// of the secure settings, which you add in the Elasticsearch keystore.
// It is not possible to configure SSL in the body of the reindex request.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
func New(tp elastictransport.Interface) *Reindex {
	r := &Reindex{
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
func (r *Reindex) Raw(raw io.Reader) *Reindex {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Reindex) Request(req *Request) *Reindex {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Reindex) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Reindex: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_reindex")

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
func (r Reindex) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "reindex")
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
		instrument.BeforeRequest(req, "reindex")
		if reader := instrument.RecordRequestBody(ctx, "reindex", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "reindex")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Reindex query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a reindex.Response
func (r Reindex) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "reindex")
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

// Header set a key, value pair in the Reindex headers map.
func (r *Reindex) Header(key, value string) *Reindex {
	r.headers.Set(key, value)

	return r
}

// Refresh If `true`, the request refreshes affected shards to make this operation
// visible to search.
// API name: refresh
func (r *Reindex) Refresh(refresh bool) *Reindex {
	r.values.Set("refresh", strconv.FormatBool(refresh))

	return r
}

// RequestsPerSecond The throttle for this request in sub-requests per second.
// By default, there is no throttle.
// API name: requests_per_second
func (r *Reindex) RequestsPerSecond(requestspersecond string) *Reindex {
	r.values.Set("requests_per_second", requestspersecond)

	return r
}

// Scroll The period of time that a consistent view of the index should be maintained
// for scrolled search.
// API name: scroll
func (r *Reindex) Scroll(duration string) *Reindex {
	r.values.Set("scroll", duration)

	return r
}

// Slices The number of slices this task should be divided into.
// It defaults to one slice, which means the task isn't sliced into subtasks.
//
// Reindex supports sliced scroll to parallelize the reindexing process.
// This parallelization can improve efficiency and provide a convenient way to
// break the request down into smaller parts.
//
// NOTE: Reindexing from remote clusters does not support manual or automatic
// slicing.
//
// If set to `auto`, Elasticsearch chooses the number of slices to use.
// This setting will use one slice per shard, up to a certain limit.
// If there are multiple sources, it will choose the number of slices based on
// the index or backing index with the smallest number of shards.
// API name: slices
func (r *Reindex) Slices(slices string) *Reindex {
	r.values.Set("slices", slices)

	return r
}

// Timeout The period each indexing waits for automatic index creation, dynamic mapping
// updates, and waiting for active shards.
// By default, Elasticsearch waits for at least one minute before failing.
// The actual wait time could be longer, particularly when multiple waits occur.
// API name: timeout
func (r *Reindex) Timeout(duration string) *Reindex {
	r.values.Set("timeout", duration)

	return r
}

// WaitForActiveShards The number of shard copies that must be active before proceeding with the
// operation.
// Set it to `all` or any positive integer up to the total number of shards in
// the index (`number_of_replicas+1`).
// The default value is one, which means it waits for each primary shard to be
// active.
// API name: wait_for_active_shards
func (r *Reindex) WaitForActiveShards(waitforactiveshards string) *Reindex {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// WaitForCompletion If `true`, the request blocks until the operation is complete.
// API name: wait_for_completion
func (r *Reindex) WaitForCompletion(waitforcompletion bool) *Reindex {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}

// RequireAlias If `true`, the destination must be an index alias.
// API name: require_alias
func (r *Reindex) RequireAlias(requirealias bool) *Reindex {
	r.values.Set("require_alias", strconv.FormatBool(requirealias))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Reindex) ErrorTrace(errortrace bool) *Reindex {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Reindex) FilterPath(filterpaths ...string) *Reindex {
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
func (r *Reindex) Human(human bool) *Reindex {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Reindex) Pretty(pretty bool) *Reindex {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Conflicts Indicates whether to continue reindexing even when there are conflicts.
// API name: conflicts
func (r *Reindex) Conflicts(conflicts conflicts.Conflicts) *Reindex {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Conflicts = &conflicts

	return r
}

// Dest The destination you are copying to.
// API name: dest
func (r *Reindex) Dest(dest *types.ReindexDestination) *Reindex {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Dest = *dest

	return r
}

// MaxDocs The maximum number of documents to reindex.
// By default, all documents are reindexed.
// If it is a value less then or equal to `scroll_size`, a scroll will not be
// used to retrieve the results for the operation.
//
// If `conflicts` is set to `proceed`, the reindex operation could attempt to
// reindex more documents from the source than `max_docs` until it has
// successfully indexed `max_docs` documents into the target or it has gone
// through every document in the source query.
// API name: max_docs
func (r *Reindex) MaxDocs(maxdocs int64) *Reindex {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.MaxDocs = &maxdocs

	return r
}

// Script The script to run to update the document source or metadata when reindexing.
// API name: script
func (r *Reindex) Script(script *types.Script) *Reindex {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Script = script

	return r
}

// API name: size
func (r *Reindex) Size(size int64) *Reindex {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Size = &size

	return r
}

// Source The source you are copying from.
// API name: source
func (r *Reindex) Source(source *types.ReindexSource) *Reindex {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Source = *source

	return r
}
