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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Restore a snapshot.
// Restore a snapshot of a cluster or data streams and indices.
//
// You can restore a snapshot only to a running cluster with an elected master
// node.
// The snapshot repository must be registered and available to the cluster.
// The snapshot and cluster versions must be compatible.
//
// To restore a snapshot, the cluster's global metadata must be writable. Ensure
// there are't any cluster blocks that prevent writes. The restore operation
// ignores index blocks.
//
// Before you restore a data stream, ensure the cluster contains a matching
// index template with data streams enabled. To check, use the index management
// feature in Kibana or the get index template API:
//
// ```
// GET
// _index_template/*?filter_path=index_templates.name,index_templates.index_template.index_patterns,index_templates.index_template.data_stream
// ```
//
// If no such template exists, you can create one or restore a cluster state
// that contains one. Without a matching index template, a data stream can't
// roll over or create backing indices.
//
// If your snapshot contains data from App Search or Workplace Search, you must
// restore the Enterprise Search encryption key before you restore the snapshot.
package restore

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	repositoryMask = iota + 1

	snapshotMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Restore struct {
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

// NewRestore type alias for index.
type NewRestore func(repository, snapshot string) *Restore

// NewRestoreFunc returns a new instance of Restore with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewRestoreFunc(tp elastictransport.Interface) NewRestore {
	return func(repository, snapshot string) *Restore {
		n := New(tp)

		n._repository(repository)

		n._snapshot(snapshot)

		return n
	}
}

// Restore a snapshot.
// Restore a snapshot of a cluster or data streams and indices.
//
// You can restore a snapshot only to a running cluster with an elected master
// node.
// The snapshot repository must be registered and available to the cluster.
// The snapshot and cluster versions must be compatible.
//
// To restore a snapshot, the cluster's global metadata must be writable. Ensure
// there are't any cluster blocks that prevent writes. The restore operation
// ignores index blocks.
//
// Before you restore a data stream, ensure the cluster contains a matching
// index template with data streams enabled. To check, use the index management
// feature in Kibana or the get index template API:
//
// ```
// GET
// _index_template/*?filter_path=index_templates.name,index_templates.index_template.index_patterns,index_templates.index_template.data_stream
// ```
//
// If no such template exists, you can create one or restore a cluster state
// that contains one. Without a matching index template, a data stream can't
// roll over or create backing indices.
//
// If your snapshot contains data from App Search or Workplace Search, you must
// restore the Enterprise Search encryption key before you restore the snapshot.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-snapshot-restore
func New(tp elastictransport.Interface) *Restore {
	r := &Restore{
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
func (r *Restore) Raw(raw io.Reader) *Restore {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Restore) Request(req *Request) *Restore {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Restore) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Restore: %w", err)
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
		path.WriteString("_restore")

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
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=9")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=9")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r Restore) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "snapshot.restore")
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
		instrument.BeforeRequest(req, "snapshot.restore")
		if reader := instrument.RecordRequestBody(ctx, "snapshot.restore", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "snapshot.restore")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Restore query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a restore.Response
func (r Restore) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "snapshot.restore")
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

// Header set a key, value pair in the Restore headers map.
func (r *Restore) Header(key, value string) *Restore {
	r.headers.Set(key, value)

	return r
}

// Repository The name of the repository to restore a snapshot from.
// API Name: repository
func (r *Restore) _repository(repository string) *Restore {
	r.paramSet |= repositoryMask
	r.repository = repository

	return r
}

// Snapshot The name of the snapshot to restore.
// API Name: snapshot
func (r *Restore) _snapshot(snapshot string) *Restore {
	r.paramSet |= snapshotMask
	r.snapshot = snapshot

	return r
}

// MasterTimeout The period to wait for the master node.
// If the master node is not available before the timeout expires, the request
// fails and returns an error.
// To indicate that the request should never timeout, set it to `-1`.
// API name: master_timeout
func (r *Restore) MasterTimeout(duration string) *Restore {
	r.values.Set("master_timeout", duration)

	return r
}

// WaitForCompletion If `true`, the request returns a response when the restore operation
// completes.
// The operation is complete when it finishes all attempts to recover primary
// shards for restored indices.
// This applies even if one or more of the recovery attempts fail.
//
// If `false`, the request returns a response when the restore operation
// initializes.
// API name: wait_for_completion
func (r *Restore) WaitForCompletion(waitforcompletion bool) *Restore {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Restore) ErrorTrace(errortrace bool) *Restore {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Restore) FilterPath(filterpaths ...string) *Restore {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"exists_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *Restore) Human(human bool) *Restore {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Restore) Pretty(pretty bool) *Restore {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// The feature states to restore.
// If `include_global_state` is `true`, the request restores all feature states
// in the snapshot by default.
// If `include_global_state` is `false`, the request restores no feature states
// by default.
// Note that specifying an empty array will result in the default behavior.
// To restore no feature states, regardless of the `include_global_state` value,
// specify an array containing only the value `none` (`["none"]`).
// API name: feature_states
func (r *Restore) FeatureStates(featurestates ...string) *Restore {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range featurestates {

		r.req.FeatureStates = append(r.req.FeatureStates, v)

	}
	return r
}

// The index settings to not restore from the snapshot.
// You can't use this option to ignore `index.number_of_shards`.
//
// For data streams, this option applies only to restored backing indices.
// New backing indices are configured using the data stream's matching index
// template.
// API name: ignore_index_settings
func (r *Restore) IgnoreIndexSettings(ignoreindexsettings ...string) *Restore {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range ignoreindexsettings {

		r.req.IgnoreIndexSettings = append(r.req.IgnoreIndexSettings, v)

	}
	return r
}

// If `true`, the request ignores any index or data stream in indices that's
// missing from the snapshot.
// If `false`, the request returns an error for any missing index or data
// stream.
// API name: ignore_unavailable
func (r *Restore) IgnoreUnavailable(ignoreunavailable bool) *Restore {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IgnoreUnavailable = &ignoreunavailable

	return r
}

// If `true`, the request restores aliases for any restored data streams and
// indices.
// If `false`, the request doesn’t restore aliases.
// API name: include_aliases
func (r *Restore) IncludeAliases(includealiases bool) *Restore {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IncludeAliases = &includealiases

	return r
}

// If `true`, restore the cluster state. The cluster state includes:
//
// * Persistent cluster settings
// * Index templates
// * Legacy index templates
// * Ingest pipelines
// * Index lifecycle management (ILM) policies
// * Stored scripts
// * For snapshots taken after 7.12.0, feature states
//
// If `include_global_state` is `true`, the restore operation merges the legacy
// index templates in your cluster with the templates contained in the snapshot,
// replacing any existing ones whose name matches one in the snapshot.
// It completely removes all persistent settings, non-legacy index templates,
// ingest pipelines, and ILM lifecycle policies that exist in your cluster and
// replaces them with the corresponding items from the snapshot.
//
// Use the `feature_states` parameter to configure how feature states are
// restored.
//
// If `include_global_state` is `true` and a snapshot was created without a
// global state then the restore request will fail.
// API name: include_global_state
func (r *Restore) IncludeGlobalState(includeglobalstate bool) *Restore {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IncludeGlobalState = &includeglobalstate

	return r
}

// Index settings to add or change in restored indices, including backing
// indices.
// You can't use this option to change `index.number_of_shards`.
//
// For data streams, this option applies only to restored backing indices.
// New backing indices are configured using the data stream's matching index
// template.
// API name: index_settings
func (r *Restore) IndexSettings(indexsettings types.IndexSettingsVariant) *Restore {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexSettings = indexsettings.IndexSettingsCaster()

	return r
}

// A comma-separated list of indices and data streams to restore.
// It supports a multi-target syntax.
// The default behavior is all regular indices and regular data streams in the
// snapshot.
//
// You can't use this parameter to restore system indices or system data
// streams.
// Use `feature_states` instead.
// API name: indices
func (r *Restore) Indices(indices ...string) *Restore {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Indices = indices

	return r
}

// If `false`, the entire restore operation will fail if one or more indices
// included in the snapshot do not have all primary shards available.
//
// If true, it allows restoring a partial snapshot of indices with unavailable
// shards.
// Only shards that were successfully included in the snapshot will be restored.
// All missing shards will be recreated as empty.
// API name: partial
func (r *Restore) Partial(partial bool) *Restore {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Partial = &partial

	return r
}

// A rename pattern to apply to restored data streams and indices.
// Data streams and indices matching the rename pattern will be renamed
// according to `rename_replacement`.
//
// The rename pattern is applied as defined by the regular expression that
// supports referencing the original text, according to the `appendReplacement`
// logic.
// API name: rename_pattern
func (r *Restore) RenamePattern(renamepattern string) *Restore {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RenamePattern = &renamepattern

	return r
}

// The rename replacement string that is used with the `rename_pattern`.
// API name: rename_replacement
func (r *Restore) RenameReplacement(renamereplacement string) *Restore {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RenameReplacement = &renamereplacement

	return r
}
