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

// Bulk index or delete documents.
// Perform multiple `index`, `create`, `delete`, and `update` actions in a
// single request.
// This reduces overhead and can greatly increase indexing speed.
//
// If the Elasticsearch security features are enabled, you must have the
// following index privileges for the target data stream, index, or index alias:
//
// * To use the `create` action, you must have the `create_doc`, `create`,
// `index`, or `write` index privilege. Data streams support only the `create`
// action.
// * To use the `index` action, you must have the `create`, `index`, or `write`
// index privilege.
// * To use the `delete` action, you must have the `delete` or `write` index
// privilege.
// * To use the `update` action, you must have the `index` or `write` index
// privilege.
// * To automatically create a data stream or index with a bulk API request, you
// must have the `auto_configure`, `create_index`, or `manage` index privilege.
// * To make the result of a bulk operation visible to search using the
// `refresh` parameter, you must have the `maintenance` or `manage` index
// privilege.
//
// Automatic data stream creation requires a matching index template with data
// stream enabled.
//
// The actions are specified in the request body using a newline delimited JSON
// (NDJSON) structure:
//
// ```
// action_and_meta_data\n
// optional_source\n
// action_and_meta_data\n
// optional_source\n
// ....
// action_and_meta_data\n
// optional_source\n
// ```
//
// The `index` and `create` actions expect a source on the next line and have
// the same semantics as the `op_type` parameter in the standard index API.
// A `create` action fails if a document with the same ID already exists in the
// target
// An `index` action adds or replaces a document as necessary.
//
// NOTE: Data streams support only the `create` action.
// To update or delete a document in a data stream, you must target the backing
// index containing the document.
//
// An `update` action expects that the partial doc, upsert, and script and its
// options are specified on the next line.
//
// A `delete` action does not expect a source on the next line and has the same
// semantics as the standard delete API.
//
// NOTE: The final line of data must end with a newline character (`\n`).
// Each newline character may be preceded by a carriage return (`\r`).
// When sending NDJSON data to the `_bulk` endpoint, use a `Content-Type` header
// of `application/json` or `application/x-ndjson`.
// Because this format uses literal newline characters (`\n`) as delimiters,
// make sure that the JSON actions and sources are not pretty printed.
//
// If you provide a target in the request path, it is used for any actions that
// don't explicitly specify an `_index` argument.
//
// A note on the format: the idea here is to make processing as fast as
// possible.
// As some of the actions are redirected to other shards on other nodes, only
// `action_meta_data` is parsed on the receiving node side.
//
// Client libraries using this protocol should try and strive to do something
// similar on the client side, and reduce buffering as much as possible.
//
// There is no "correct" number of actions to perform in a single bulk request.
// Experiment with different settings to find the optimal size for your
// particular workload.
// Note that Elasticsearch limits the maximum size of a HTTP request to 100mb by
// default so clients must ensure that no request exceeds this size.
// It is not possible to index a single document that exceeds the size limit, so
// you must pre-process any such documents into smaller pieces before sending
// them to Elasticsearch.
// For instance, split documents into pages or chapters before indexing them, or
// store raw binary data in a system outside Elasticsearch and replace the raw
// data with a link to the external system in the documents that you send to
// Elasticsearch.
//
// **Client suppport for bulk requests**
//
// Some of the officially supported clients provide helpers to assist with bulk
// requests and reindexing:
//
// * Go: Check out `esutil.BulkIndexer`
// * Perl: Check out `Search::Elasticsearch::Client::5_0::Bulk` and
// `Search::Elasticsearch::Client::5_0::Scroll`
// * Python: Check out `elasticsearch.helpers.*`
// * JavaScript: Check out `client.helpers.*`
// * .NET: Check out `BulkAllObservable`
// * PHP: Check out bulk indexing.
//
// **Submitting bulk requests with cURL**
//
// If you're providing text file input to `curl`, you must use the
// `--data-binary` flag instead of plain `-d`.
// The latter doesn't preserve newlines. For example:
//
// ```
// $ cat requests
// { "index" : { "_index" : "test", "_id" : "1" } }
// { "field1" : "value1" }
// $ curl -s -H "Content-Type: application/x-ndjson" -XPOST localhost:9200/_bulk
// --data-binary "@requests"; echo
// {"took":7, "errors": false,
// "items":[{"index":{"_index":"test","_id":"1","_version":1,"result":"created","forced_refresh":false}}]}
// ```
//
// **Optimistic concurrency control**
//
// Each `index` and `delete` action within a bulk API call may include the
// `if_seq_no` and `if_primary_term` parameters in their respective action and
// meta data lines.
// The `if_seq_no` and `if_primary_term` parameters control how operations are
// run, based on the last modification to existing documents. See Optimistic
// concurrency control for more details.
//
// **Versioning**
//
// Each bulk item can include the version value using the `version` field.
// It automatically follows the behavior of the index or delete operation based
// on the `_version` mapping.
// It also support the `version_type`.
//
// **Routing**
//
// Each bulk item can include the routing value using the `routing` field.
// It automatically follows the behavior of the index or delete operation based
// on the `_routing` mapping.
//
// NOTE: Data streams do not support custom routing unless they were created
// with the `allow_custom_routing` setting enabled in the template.
//
// **Wait for active shards**
//
// When making bulk calls, you can set the `wait_for_active_shards` parameter to
// require a minimum number of shard copies to be active before starting to
// process the bulk request.
//
// **Refresh**
//
// Control when the changes made by this request are visible to search.
//
// NOTE: Only the shards that receive the bulk request will be affected by
// refresh.
// Imagine a `_bulk?refresh=wait_for` request with three documents in it that
// happen to be routed to different shards in an index with five shards.
// The request will only wait for those three shards to refresh.
// The other two shards that make up the index do not participate in the `_bulk`
// request at all.
package bulk

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
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Bulk struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewBulk type alias for index.
type NewBulk func() *Bulk

// NewBulkFunc returns a new instance of Bulk with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewBulkFunc(tp elastictransport.Interface) NewBulk {
	return func() *Bulk {
		n := New(tp)

		return n
	}
}

// Bulk index or delete documents.
// Perform multiple `index`, `create`, `delete`, and `update` actions in a
// single request.
// This reduces overhead and can greatly increase indexing speed.
//
// If the Elasticsearch security features are enabled, you must have the
// following index privileges for the target data stream, index, or index alias:
//
// * To use the `create` action, you must have the `create_doc`, `create`,
// `index`, or `write` index privilege. Data streams support only the `create`
// action.
// * To use the `index` action, you must have the `create`, `index`, or `write`
// index privilege.
// * To use the `delete` action, you must have the `delete` or `write` index
// privilege.
// * To use the `update` action, you must have the `index` or `write` index
// privilege.
// * To automatically create a data stream or index with a bulk API request, you
// must have the `auto_configure`, `create_index`, or `manage` index privilege.
// * To make the result of a bulk operation visible to search using the
// `refresh` parameter, you must have the `maintenance` or `manage` index
// privilege.
//
// Automatic data stream creation requires a matching index template with data
// stream enabled.
//
// The actions are specified in the request body using a newline delimited JSON
// (NDJSON) structure:
//
// ```
// action_and_meta_data\n
// optional_source\n
// action_and_meta_data\n
// optional_source\n
// ....
// action_and_meta_data\n
// optional_source\n
// ```
//
// The `index` and `create` actions expect a source on the next line and have
// the same semantics as the `op_type` parameter in the standard index API.
// A `create` action fails if a document with the same ID already exists in the
// target
// An `index` action adds or replaces a document as necessary.
//
// NOTE: Data streams support only the `create` action.
// To update or delete a document in a data stream, you must target the backing
// index containing the document.
//
// An `update` action expects that the partial doc, upsert, and script and its
// options are specified on the next line.
//
// A `delete` action does not expect a source on the next line and has the same
// semantics as the standard delete API.
//
// NOTE: The final line of data must end with a newline character (`\n`).
// Each newline character may be preceded by a carriage return (`\r`).
// When sending NDJSON data to the `_bulk` endpoint, use a `Content-Type` header
// of `application/json` or `application/x-ndjson`.
// Because this format uses literal newline characters (`\n`) as delimiters,
// make sure that the JSON actions and sources are not pretty printed.
//
// If you provide a target in the request path, it is used for any actions that
// don't explicitly specify an `_index` argument.
//
// A note on the format: the idea here is to make processing as fast as
// possible.
// As some of the actions are redirected to other shards on other nodes, only
// `action_meta_data` is parsed on the receiving node side.
//
// Client libraries using this protocol should try and strive to do something
// similar on the client side, and reduce buffering as much as possible.
//
// There is no "correct" number of actions to perform in a single bulk request.
// Experiment with different settings to find the optimal size for your
// particular workload.
// Note that Elasticsearch limits the maximum size of a HTTP request to 100mb by
// default so clients must ensure that no request exceeds this size.
// It is not possible to index a single document that exceeds the size limit, so
// you must pre-process any such documents into smaller pieces before sending
// them to Elasticsearch.
// For instance, split documents into pages or chapters before indexing them, or
// store raw binary data in a system outside Elasticsearch and replace the raw
// data with a link to the external system in the documents that you send to
// Elasticsearch.
//
// **Client suppport for bulk requests**
//
// Some of the officially supported clients provide helpers to assist with bulk
// requests and reindexing:
//
// * Go: Check out `esutil.BulkIndexer`
// * Perl: Check out `Search::Elasticsearch::Client::5_0::Bulk` and
// `Search::Elasticsearch::Client::5_0::Scroll`
// * Python: Check out `elasticsearch.helpers.*`
// * JavaScript: Check out `client.helpers.*`
// * .NET: Check out `BulkAllObservable`
// * PHP: Check out bulk indexing.
//
// **Submitting bulk requests with cURL**
//
// If you're providing text file input to `curl`, you must use the
// `--data-binary` flag instead of plain `-d`.
// The latter doesn't preserve newlines. For example:
//
// ```
// $ cat requests
// { "index" : { "_index" : "test", "_id" : "1" } }
// { "field1" : "value1" }
// $ curl -s -H "Content-Type: application/x-ndjson" -XPOST localhost:9200/_bulk
// --data-binary "@requests"; echo
// {"took":7, "errors": false,
// "items":[{"index":{"_index":"test","_id":"1","_version":1,"result":"created","forced_refresh":false}}]}
// ```
//
// **Optimistic concurrency control**
//
// Each `index` and `delete` action within a bulk API call may include the
// `if_seq_no` and `if_primary_term` parameters in their respective action and
// meta data lines.
// The `if_seq_no` and `if_primary_term` parameters control how operations are
// run, based on the last modification to existing documents. See Optimistic
// concurrency control for more details.
//
// **Versioning**
//
// Each bulk item can include the version value using the `version` field.
// It automatically follows the behavior of the index or delete operation based
// on the `_version` mapping.
// It also support the `version_type`.
//
// **Routing**
//
// Each bulk item can include the routing value using the `routing` field.
// It automatically follows the behavior of the index or delete operation based
// on the `_routing` mapping.
//
// NOTE: Data streams do not support custom routing unless they were created
// with the `allow_custom_routing` setting enabled in the template.
//
// **Wait for active shards**
//
// When making bulk calls, you can set the `wait_for_active_shards` parameter to
// require a minimum number of shard copies to be active before starting to
// process the bulk request.
//
// **Refresh**
//
// Control when the changes made by this request are visible to search.
//
// NOTE: Only the shards that receive the bulk request will be affected by
// refresh.
// Imagine a `_bulk?refresh=wait_for` request with three documents in it that
// happen to be routed to different shards in an index with five shards.
// The request will only wait for those three shards to refresh.
// The other two shards that make up the index do not participate in the `_bulk`
// request at all.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
func New(tp elastictransport.Interface) *Bulk {
	r := &Bulk{
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
func (r *Bulk) Raw(raw io.Reader) *Bulk {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Bulk) Request(req *Request) *Bulk {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Bulk) HttpRequest(ctx context.Context) (*http.Request, error) {
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

		for _, elem := range *r.req {
			data, err := json.Marshal(elem)
			if err != nil {
				return nil, err
			}
			r.buf.Write(data)
			r.buf.Write([]byte("\n"))
		}

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Bulk: %w", err)
		}

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_bulk")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_bulk")

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
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+x-ndjson;compatible-with=8")
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
func (r Bulk) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "bulk")
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
		instrument.BeforeRequest(req, "bulk")
		if reader := instrument.RecordRequestBody(ctx, "bulk", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "bulk")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Bulk query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a bulk.Response
func (r Bulk) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "bulk")
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

// Header set a key, value pair in the Bulk headers map.
func (r *Bulk) Header(key, value string) *Bulk {
	r.headers.Set(key, value)

	return r
}

// Index The name of the data stream, index, or index alias to perform bulk actions
// on.
// API Name: index
func (r *Bulk) Index(index string) *Bulk {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// IncludeSourceOnError True or false if to include the document source in the error message in case
// of parsing errors.
// API name: include_source_on_error
func (r *Bulk) IncludeSourceOnError(includesourceonerror bool) *Bulk {
	r.values.Set("include_source_on_error", strconv.FormatBool(includesourceonerror))

	return r
}

// ListExecutedPipelines If `true`, the response will include the ingest pipelines that were run for
// each index or create.
// API name: list_executed_pipelines
func (r *Bulk) ListExecutedPipelines(listexecutedpipelines bool) *Bulk {
	r.values.Set("list_executed_pipelines", strconv.FormatBool(listexecutedpipelines))

	return r
}

// Pipeline The pipeline identifier to use to preprocess incoming documents.
// If the index has a default ingest pipeline specified, setting the value to
// `_none` turns off the default ingest pipeline for this request.
// If a final pipeline is configured, it will always run regardless of the value
// of this parameter.
// API name: pipeline
func (r *Bulk) Pipeline(pipeline string) *Bulk {
	r.values.Set("pipeline", pipeline)

	return r
}

// Refresh If `true`, Elasticsearch refreshes the affected shards to make this operation
// visible to search.
// If `wait_for`, wait for a refresh to make this operation visible to search.
// If `false`, do nothing with refreshes.
// Valid values: `true`, `false`, `wait_for`.
// API name: refresh
func (r *Bulk) Refresh(refresh refresh.Refresh) *Bulk {
	r.values.Set("refresh", refresh.String())

	return r
}

// Routing A custom value that is used to route operations to a specific shard.
// API name: routing
func (r *Bulk) Routing(routing string) *Bulk {
	r.values.Set("routing", routing)

	return r
}

// Source_ Indicates whether to return the `_source` field (`true` or `false`) or
// contains a list of fields to return.
// API name: _source
func (r *Bulk) Source_(sourceconfigparam string) *Bulk {
	r.values.Set("_source", sourceconfigparam)

	return r
}

// SourceExcludes_ A comma-separated list of source fields to exclude from the response.
// You can also use this parameter to exclude fields from the subset specified
// in `_source_includes` query parameter.
// If the `_source` parameter is `false`, this parameter is ignored.
// API name: _source_excludes
func (r *Bulk) SourceExcludes_(fields ...string) *Bulk {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ A comma-separated list of source fields to include in the response.
// If this parameter is specified, only these source fields are returned.
// You can exclude fields from this subset using the `_source_excludes` query
// parameter.
// If the `_source` parameter is `false`, this parameter is ignored.
// API name: _source_includes
func (r *Bulk) SourceIncludes_(fields ...string) *Bulk {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// Timeout The period each action waits for the following operations: automatic index
// creation, dynamic mapping updates, and waiting for active shards.
// The default is `1m` (one minute), which guarantees Elasticsearch waits for at
// least the timeout before failing.
// The actual wait time could be longer, particularly when multiple waits occur.
// API name: timeout
func (r *Bulk) Timeout(duration string) *Bulk {
	r.values.Set("timeout", duration)

	return r
}

// WaitForActiveShards The number of shard copies that must be active before proceeding with the
// operation.
// Set to `all` or any positive integer up to the total number of shards in the
// index (`number_of_replicas+1`).
// The default is `1`, which waits for each primary shard to be active.
// API name: wait_for_active_shards
func (r *Bulk) WaitForActiveShards(waitforactiveshards string) *Bulk {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// RequireAlias If `true`, the request's actions must target an index alias.
// API name: require_alias
func (r *Bulk) RequireAlias(requirealias bool) *Bulk {
	r.values.Set("require_alias", strconv.FormatBool(requirealias))

	return r
}

// RequireDataStream If `true`, the request's actions must target a data stream (existing or to be
// created).
// API name: require_data_stream
func (r *Bulk) RequireDataStream(requiredatastream bool) *Bulk {
	r.values.Set("require_data_stream", strconv.FormatBool(requiredatastream))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Bulk) ErrorTrace(errortrace bool) *Bulk {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Bulk) FilterPath(filterpaths ...string) *Bulk {
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
func (r *Bulk) Human(human bool) *Bulk {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Bulk) Pretty(pretty bool) *Bulk {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
