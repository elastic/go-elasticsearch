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

// Update index settings.
// Changes dynamic index settings in real time.
// For data streams, index setting changes are applied to all backing indices by
// default.
//
// To revert a setting to the default value, use a null value.
// The list of per-index settings that can be updated dynamically on live
// indices can be found in index settings documentation.
// To preserve existing settings from being updated, set the `preserve_existing`
// parameter to `true`.
//
//	There are multiple valid ways to represent index settings in the request
//
// body. You can specify only the setting, for example:
//
// ```
//
//	{
//	  "number_of_replicas": 1
//	}
//
// ```
//
// Or you can use an `index` setting object:
// ```
//
//	{
//	  "index": {
//	    "number_of_replicas": 1
//	  }
//	}
//
// ```
//
// Or you can use dot annotation:
// ```
//
//	{
//	  "index.number_of_replicas": 1
//	}
//
// ```
//
// Or you can embed any of the aforementioned options in a `settings` object.
// For example:
//
// ```
//
//	{
//	  "settings": {
//	    "index": {
//	      "number_of_replicas": 1
//	    }
//	  }
//	}
//
// ```
//
// NOTE: You can only define new analyzers on closed indices.
// To add an analyzer, you must close the index, define the analyzer, and reopen
// the index.
// You cannot close the write index of a data stream.
// To update the analyzer for a data stream's write index and future backing
// indices, update the analyzer in the index template used by the stream.
// Then roll over the data stream to apply the new analyzer to the stream's
// write index and future backing indices.
// This affects searches and any new data added to the stream after the
// rollover.
// However, it does not affect the data stream's backing indices or their
// existing data.
// To change the analyzer for existing backing indices, you must create a new
// data stream and reindex your data into it.
package putsettings

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexcheckonstartup"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutSettings struct {
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

// NewPutSettings type alias for index.
type NewPutSettings func() *PutSettings

// NewPutSettingsFunc returns a new instance of PutSettings with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutSettingsFunc(tp elastictransport.Interface) NewPutSettings {
	return func() *PutSettings {
		n := New(tp)

		return n
	}
}

// Update index settings.
// Changes dynamic index settings in real time.
// For data streams, index setting changes are applied to all backing indices by
// default.
//
// To revert a setting to the default value, use a null value.
// The list of per-index settings that can be updated dynamically on live
// indices can be found in index settings documentation.
// To preserve existing settings from being updated, set the `preserve_existing`
// parameter to `true`.
//
//	There are multiple valid ways to represent index settings in the request
//
// body. You can specify only the setting, for example:
//
// ```
//
//	{
//	  "number_of_replicas": 1
//	}
//
// ```
//
// Or you can use an `index` setting object:
// ```
//
//	{
//	  "index": {
//	    "number_of_replicas": 1
//	  }
//	}
//
// ```
//
// Or you can use dot annotation:
// ```
//
//	{
//	  "index.number_of_replicas": 1
//	}
//
// ```
//
// Or you can embed any of the aforementioned options in a `settings` object.
// For example:
//
// ```
//
//	{
//	  "settings": {
//	    "index": {
//	      "number_of_replicas": 1
//	    }
//	  }
//	}
//
// ```
//
// NOTE: You can only define new analyzers on closed indices.
// To add an analyzer, you must close the index, define the analyzer, and reopen
// the index.
// You cannot close the write index of a data stream.
// To update the analyzer for a data stream's write index and future backing
// indices, update the analyzer in the index template used by the stream.
// Then roll over the data stream to apply the new analyzer to the stream's
// write index and future backing indices.
// This affects searches and any new data added to the stream after the
// rollover.
// However, it does not affect the data stream's backing indices or their
// existing data.
// To change the analyzer for existing backing indices, you must create a new
// data stream and reindex your data into it.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-update-settings.html
func New(tp elastictransport.Interface) *PutSettings {
	r := &PutSettings{
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
func (r *PutSettings) Raw(raw io.Reader) *PutSettings {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutSettings) Request(req *Request) *PutSettings {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutSettings) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutSettings: %w", err)
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
		path.WriteString("_settings")

		method = http.MethodPut
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_settings")

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
func (r PutSettings) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.put_settings")
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
		instrument.BeforeRequest(req, "indices.put_settings")
		if reader := instrument.RecordRequestBody(ctx, "indices.put_settings", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.put_settings")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutSettings query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putsettings.Response
func (r PutSettings) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.put_settings")
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

// Header set a key, value pair in the PutSettings headers map.
func (r *PutSettings) Header(key, value string) *PutSettings {
	r.headers.Set(key, value)

	return r
}

// Indices Comma-separated list of data streams, indices, and aliases used to limit
// the request. Supports wildcards (`*`). To target all data streams and
// indices, omit this parameter or use `*` or `_all`.
// API Name: index
func (r *PutSettings) Indices(index string) *PutSettings {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices If `false`, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices. This
// behavior applies even if the request targets other open indices. For
// example, a request targeting `foo*,bar*` returns an error if an index
// starts with `foo` but no index starts with `bar`.
// API name: allow_no_indices
func (r *PutSettings) AllowNoIndices(allownoindices bool) *PutSettings {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match. If the request can target
// data streams, this argument determines whether wildcard expressions match
// hidden data streams. Supports comma-separated values, such as
// `open,hidden`.
// API name: expand_wildcards
func (r *PutSettings) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *PutSettings {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// FlatSettings If `true`, returns settings in flat format.
// API name: flat_settings
func (r *PutSettings) FlatSettings(flatsettings bool) *PutSettings {
	r.values.Set("flat_settings", strconv.FormatBool(flatsettings))

	return r
}

// IgnoreUnavailable If `true`, returns settings in flat format.
// API name: ignore_unavailable
func (r *PutSettings) IgnoreUnavailable(ignoreunavailable bool) *PutSettings {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an
// error.
// API name: master_timeout
func (r *PutSettings) MasterTimeout(duration string) *PutSettings {
	r.values.Set("master_timeout", duration)

	return r
}

// PreserveExisting If `true`, existing index settings remain unchanged.
// API name: preserve_existing
func (r *PutSettings) PreserveExisting(preserveexisting bool) *PutSettings {
	r.values.Set("preserve_existing", strconv.FormatBool(preserveexisting))

	return r
}

// Reopen Whether to close and reopen the index to apply non-dynamic settings.
// If set to `true` the indices to which the settings are being applied
// will be closed temporarily and then reopened in order to apply the changes.
// API name: reopen
func (r *PutSettings) Reopen(reopen bool) *PutSettings {
	r.values.Set("reopen", strconv.FormatBool(reopen))

	return r
}

// Timeout Period to wait for a response. If no response is received before the
//
//	timeout expires, the request fails and returns an error.
//
// API name: timeout
func (r *PutSettings) Timeout(duration string) *PutSettings {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutSettings) ErrorTrace(errortrace bool) *PutSettings {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutSettings) FilterPath(filterpaths ...string) *PutSettings {
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
func (r *PutSettings) Human(human bool) *PutSettings {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutSettings) Pretty(pretty bool) *PutSettings {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// API name: analysis
func (r *PutSettings) Analysis(analysis *types.IndexSettingsAnalysis) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Analysis = analysis

	return r
}

// Analyze Settings to define analyzers, tokenizers, token filters and character
// filters.
// API name: analyze
func (r *PutSettings) Analyze(analyze *types.SettingsAnalyze) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Analyze = analyze

	return r
}

// API name: auto_expand_replicas
func (r *PutSettings) AutoExpandReplicas(autoexpandreplicas any) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.AutoExpandReplicas = autoexpandreplicas

	return r
}

// API name: blocks
func (r *PutSettings) Blocks(blocks *types.IndexSettingBlocks) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Blocks = blocks

	return r
}

// API name: check_on_startup
func (r *PutSettings) CheckOnStartup(checkonstartup indexcheckonstartup.IndexCheckOnStartup) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.CheckOnStartup = &checkonstartup

	return r
}

// API name: codec
func (r *PutSettings) Codec(codec string) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Codec = &codec

	return r
}

// API name: creation_date
func (r *PutSettings) CreationDate(stringifiedepochtimeunitmillis types.StringifiedEpochTimeUnitMillis) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.CreationDate = stringifiedepochtimeunitmillis

	return r
}

// API name: creation_date_string
func (r *PutSettings) CreationDateString(datetime types.DateTime) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.CreationDateString = datetime

	return r
}

// API name: default_pipeline
func (r *PutSettings) DefaultPipeline(pipelinename string) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.DefaultPipeline = &pipelinename

	return r
}

// API name: final_pipeline
func (r *PutSettings) FinalPipeline(pipelinename string) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.FinalPipeline = &pipelinename

	return r
}

// API name: format
func (r *PutSettings) Format(format string) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Format = &format

	return r
}

// API name: gc_deletes
func (r *PutSettings) GcDeletes(duration types.Duration) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.GcDeletes = duration

	return r
}

// API name: hidden
func (r *PutSettings) Hidden(hidden string) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Hidden = &hidden

	return r
}

// API name: highlight
func (r *PutSettings) Highlight(highlight *types.SettingsHighlight) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Highlight = highlight

	return r
}

// API name: index
func (r *PutSettings) Index(index *types.IndexSettings) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Index = index

	return r
}

// API name: IndexSettings
func (r *PutSettings) IndexSettings(indexsettings map[string]json.RawMessage) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexSettings = indexsettings

	return r
}

// IndexingPressure Configure indexing back pressure limits.
// API name: indexing_pressure
func (r *PutSettings) IndexingPressure(indexingpressure *types.IndicesIndexingPressure) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexingPressure = indexingpressure

	return r
}

// API name: indexing.slowlog
func (r *PutSettings) IndexingSlowlog(indexingslowlog *types.IndexingSlowlogSettings) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexingSlowlog = indexingslowlog

	return r
}

// API name: lifecycle
func (r *PutSettings) Lifecycle(lifecycle *types.IndexSettingsLifecycle) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Lifecycle = lifecycle

	return r
}

// API name: load_fixed_bitset_filters_eagerly
func (r *PutSettings) LoadFixedBitsetFiltersEagerly(loadfixedbitsetfilterseagerly bool) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.LoadFixedBitsetFiltersEagerly = &loadfixedbitsetfilterseagerly

	return r
}

// Mapping Enable or disable dynamic mapping for an index.
// API name: mapping
func (r *PutSettings) Mapping(mapping *types.MappingLimitSettings) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Mapping = mapping

	return r
}

// API name: max_docvalue_fields_search
func (r *PutSettings) MaxDocvalueFieldsSearch(maxdocvaluefieldssearch int) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxDocvalueFieldsSearch = &maxdocvaluefieldssearch

	return r
}

// API name: max_inner_result_window
func (r *PutSettings) MaxInnerResultWindow(maxinnerresultwindow int) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxInnerResultWindow = &maxinnerresultwindow

	return r
}

// API name: max_ngram_diff
func (r *PutSettings) MaxNgramDiff(maxngramdiff int) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxNgramDiff = &maxngramdiff

	return r
}

// API name: max_refresh_listeners
func (r *PutSettings) MaxRefreshListeners(maxrefreshlisteners int) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxRefreshListeners = &maxrefreshlisteners

	return r
}

// API name: max_regex_length
func (r *PutSettings) MaxRegexLength(maxregexlength int) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxRegexLength = &maxregexlength

	return r
}

// API name: max_rescore_window
func (r *PutSettings) MaxRescoreWindow(maxrescorewindow int) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxRescoreWindow = &maxrescorewindow

	return r
}

// API name: max_result_window
func (r *PutSettings) MaxResultWindow(maxresultwindow int) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxResultWindow = &maxresultwindow

	return r
}

// API name: max_script_fields
func (r *PutSettings) MaxScriptFields(maxscriptfields int) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxScriptFields = &maxscriptfields

	return r
}

// API name: max_shingle_diff
func (r *PutSettings) MaxShingleDiff(maxshinglediff int) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxShingleDiff = &maxshinglediff

	return r
}

// API name: max_slices_per_scroll
func (r *PutSettings) MaxSlicesPerScroll(maxslicesperscroll int) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxSlicesPerScroll = &maxslicesperscroll

	return r
}

// API name: max_terms_count
func (r *PutSettings) MaxTermsCount(maxtermscount int) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxTermsCount = &maxtermscount

	return r
}

// API name: merge
func (r *PutSettings) Merge(merge *types.Merge) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Merge = merge

	return r
}

// API name: mode
func (r *PutSettings) Mode(mode string) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Mode = &mode

	return r
}

// API name: number_of_replicas
func (r *PutSettings) NumberOfReplicas(numberofreplicas string) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.NumberOfReplicas = &numberofreplicas

	return r
}

// API name: number_of_routing_shards
func (r *PutSettings) NumberOfRoutingShards(numberofroutingshards int) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.NumberOfRoutingShards = &numberofroutingshards

	return r
}

// API name: number_of_shards
func (r *PutSettings) NumberOfShards(numberofshards string) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.NumberOfShards = &numberofshards

	return r
}

// API name: priority
func (r *PutSettings) Priority(priority string) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Priority = &priority

	return r
}

// API name: provided_name
func (r *PutSettings) ProvidedName(name string) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ProvidedName = &name

	return r
}

// API name: queries
func (r *PutSettings) Queries(queries *types.Queries) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Queries = queries

	return r
}

// API name: query_string
func (r *PutSettings) QueryString(querystring *types.SettingsQueryString) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.QueryString = querystring

	return r
}

// API name: refresh_interval
func (r *PutSettings) RefreshInterval(duration types.Duration) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RefreshInterval = duration

	return r
}

// API name: routing
func (r *PutSettings) Routing(routing *types.IndexRouting) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Routing = routing

	return r
}

// API name: routing_partition_size
func (r *PutSettings) RoutingPartitionSize(stringifiedinteger types.Stringifiedinteger) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RoutingPartitionSize = stringifiedinteger

	return r
}

// API name: routing_path
func (r *PutSettings) RoutingPath(routingpaths ...string) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RoutingPath = routingpaths

	return r
}

// API name: search
func (r *PutSettings) Search(search *types.SettingsSearch) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Search = search

	return r
}

// API name: settings
func (r *PutSettings) Settings(settings *types.IndexSettings) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Settings = settings

	return r
}

// Similarity Configure custom similarity settings to customize how search results are
// scored.
// API name: similarity
func (r *PutSettings) Similarity(similarity map[string]types.SettingsSimilarity) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Similarity = similarity

	return r
}

// API name: soft_deletes
func (r *PutSettings) SoftDeletes(softdeletes *types.SoftDeletes) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.SoftDeletes = softdeletes

	return r
}

// API name: sort
func (r *PutSettings) Sort(sort *types.IndexSegmentSort) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Sort = sort

	return r
}

// Store The store module allows you to control how index data is stored and accessed
// on disk.
// API name: store
func (r *PutSettings) Store(store *types.Storage) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Store = store

	return r
}

// API name: time_series
func (r *PutSettings) TimeSeries(timeseries *types.IndexSettingsTimeSeries) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TimeSeries = timeseries

	return r
}

// API name: top_metrics_max_size
func (r *PutSettings) TopMetricsMaxSize(topmetricsmaxsize int) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.TopMetricsMaxSize = &topmetricsmaxsize

	return r
}

// API name: translog
func (r *PutSettings) Translog(translog *types.Translog) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Translog = translog

	return r
}

// API name: uuid
func (r *PutSettings) Uuid(uuid string) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Uuid = &uuid

	return r
}

// API name: verified_before_close
func (r *PutSettings) VerifiedBeforeClose(verifiedbeforeclose string) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.VerifiedBeforeClose = &verifiedbeforeclose

	return r
}

// API name: version
func (r *PutSettings) Version(version *types.IndexVersioning) *PutSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Version = version

	return r
}
