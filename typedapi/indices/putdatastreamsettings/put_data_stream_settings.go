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
// https://github.com/elastic/elasticsearch-specification/tree/224e96968e3ab27c2d1d33f015783b44ed183c1f

// Update data stream settings.
//
// This API can be used to override settings on specific data streams. These
// overrides will take precedence over what
// is specified in the template that the data stream matches. To prevent your
// data stream from getting into an invalid state,
// only certain settings are allowed. If possible, the setting change is applied
// to all
// backing indices. Otherwise, it will be applied when the data stream is next
// rolled over.
package putdatastreamsettings

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexcheckonstartup"
)

const (
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutDataStreamSettings struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutDataStreamSettings type alias for index.
type NewPutDataStreamSettings func(name string) *PutDataStreamSettings

// NewPutDataStreamSettingsFunc returns a new instance of PutDataStreamSettings with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutDataStreamSettingsFunc(tp elastictransport.Interface) NewPutDataStreamSettings {
	return func(name string) *PutDataStreamSettings {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Update data stream settings.
//
// This API can be used to override settings on specific data streams. These
// overrides will take precedence over what
// is specified in the template that the data stream matches. To prevent your
// data stream from getting into an invalid state,
// only certain settings are allowed. If possible, the setting change is applied
// to all
// backing indices. Otherwise, it will be applied when the data stream is next
// rolled over.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-data-stream-settings.html
func New(tp elastictransport.Interface) *PutDataStreamSettings {
	r := &PutDataStreamSettings{
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
func (r *PutDataStreamSettings) Raw(raw io.Reader) *PutDataStreamSettings {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutDataStreamSettings) Request(req *Request) *PutDataStreamSettings {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutDataStreamSettings) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutDataStreamSettings: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_data_stream")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", r.name)
		}
		path.WriteString(r.name)
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r PutDataStreamSettings) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.put_data_stream_settings")
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
		instrument.BeforeRequest(req, "indices.put_data_stream_settings")
		if reader := instrument.RecordRequestBody(ctx, "indices.put_data_stream_settings", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.put_data_stream_settings")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutDataStreamSettings query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putdatastreamsettings.Response
func (r PutDataStreamSettings) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.put_data_stream_settings")
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

// Header set a key, value pair in the PutDataStreamSettings headers map.
func (r *PutDataStreamSettings) Header(key, value string) *PutDataStreamSettings {
	r.headers.Set(key, value)

	return r
}

// Name A comma-separated list of data streams or data stream patterns.
// API Name: name
func (r *PutDataStreamSettings) _name(name string) *PutDataStreamSettings {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// DryRun If `true`, the request does not actually change the settings on any data
// streams or indices. Instead, it
// simulates changing the settings and reports back to the user what would have
// happened had these settings
// actually been applied.
// API name: dry_run
func (r *PutDataStreamSettings) DryRun(dryrun bool) *PutDataStreamSettings {
	r.values.Set("dry_run", strconv.FormatBool(dryrun))

	return r
}

// MasterTimeout The period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an
// error.
// API name: master_timeout
func (r *PutDataStreamSettings) MasterTimeout(duration string) *PutDataStreamSettings {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout The period to wait for a response. If no response is received before the
//
//	timeout expires, the request fails and returns an error.
//
// API name: timeout
func (r *PutDataStreamSettings) Timeout(duration string) *PutDataStreamSettings {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutDataStreamSettings) ErrorTrace(errortrace bool) *PutDataStreamSettings {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutDataStreamSettings) FilterPath(filterpaths ...string) *PutDataStreamSettings {
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
func (r *PutDataStreamSettings) Human(human bool) *PutDataStreamSettings {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutDataStreamSettings) Pretty(pretty bool) *PutDataStreamSettings {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// API name: analysis
func (r *PutDataStreamSettings) Analysis(analysis *types.IndexSettingsAnalysis) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Analysis = analysis

	return r
}

// Analyze Settings to define analyzers, tokenizers, token filters and character
// filters.
// API name: analyze
func (r *PutDataStreamSettings) Analyze(analyze *types.SettingsAnalyze) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Analyze = analyze

	return r
}

// API name: auto_expand_replicas
func (r *PutDataStreamSettings) AutoExpandReplicas(autoexpandreplicas any) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.AutoExpandReplicas = autoexpandreplicas

	return r
}

// API name: blocks
func (r *PutDataStreamSettings) Blocks(blocks *types.IndexSettingBlocks) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Blocks = blocks

	return r
}

// API name: check_on_startup
func (r *PutDataStreamSettings) CheckOnStartup(checkonstartup indexcheckonstartup.IndexCheckOnStartup) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.CheckOnStartup = &checkonstartup

	return r
}

// API name: codec
func (r *PutDataStreamSettings) Codec(codec string) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Codec = &codec

	return r
}

// API name: creation_date
func (r *PutDataStreamSettings) CreationDate(stringifiedepochtimeunitmillis types.StringifiedEpochTimeUnitMillis) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.CreationDate = stringifiedepochtimeunitmillis

	return r
}

// API name: creation_date_string
func (r *PutDataStreamSettings) CreationDateString(datetime types.DateTime) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.CreationDateString = datetime

	return r
}

// API name: default_pipeline
func (r *PutDataStreamSettings) DefaultPipeline(pipelinename string) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.DefaultPipeline = &pipelinename

	return r
}

// API name: final_pipeline
func (r *PutDataStreamSettings) FinalPipeline(pipelinename string) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.FinalPipeline = &pipelinename

	return r
}

// API name: format
func (r *PutDataStreamSettings) Format(format string) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Format = &format

	return r
}

// API name: gc_deletes
func (r *PutDataStreamSettings) GcDeletes(duration types.Duration) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.GcDeletes = duration

	return r
}

// API name: hidden
func (r *PutDataStreamSettings) Hidden(hidden string) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Hidden = &hidden

	return r
}

// API name: highlight
func (r *PutDataStreamSettings) Highlight(highlight *types.SettingsHighlight) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Highlight = highlight

	return r
}

// API name: index
func (r *PutDataStreamSettings) Index(index *types.IndexSettings) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Index = index

	return r
}

// API name: IndexSettings
func (r *PutDataStreamSettings) IndexSettings(indexsettings map[string]json.RawMessage) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexSettings = indexsettings

	return r
}

// IndexingPressure Configure indexing back pressure limits.
// API name: indexing_pressure
func (r *PutDataStreamSettings) IndexingPressure(indexingpressure *types.IndicesIndexingPressure) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexingPressure = indexingpressure

	return r
}

// API name: indexing.slowlog
func (r *PutDataStreamSettings) IndexingSlowlog(indexingslowlog *types.IndexingSlowlogSettings) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexingSlowlog = indexingslowlog

	return r
}

// API name: lifecycle
func (r *PutDataStreamSettings) Lifecycle(lifecycle *types.IndexSettingsLifecycle) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Lifecycle = lifecycle

	return r
}

// API name: load_fixed_bitset_filters_eagerly
func (r *PutDataStreamSettings) LoadFixedBitsetFiltersEagerly(loadfixedbitsetfilterseagerly bool) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.LoadFixedBitsetFiltersEagerly = &loadfixedbitsetfilterseagerly

	return r
}

// Mapping Enable or disable dynamic mapping for an index.
// API name: mapping
func (r *PutDataStreamSettings) Mapping(mapping *types.MappingLimitSettings) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Mapping = mapping

	return r
}

// API name: max_docvalue_fields_search
func (r *PutDataStreamSettings) MaxDocvalueFieldsSearch(maxdocvaluefieldssearch int) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxDocvalueFieldsSearch = &maxdocvaluefieldssearch

	return r
}

// API name: max_inner_result_window
func (r *PutDataStreamSettings) MaxInnerResultWindow(maxinnerresultwindow int) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxInnerResultWindow = &maxinnerresultwindow

	return r
}

// API name: max_ngram_diff
func (r *PutDataStreamSettings) MaxNgramDiff(maxngramdiff int) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxNgramDiff = &maxngramdiff

	return r
}

// API name: max_refresh_listeners
func (r *PutDataStreamSettings) MaxRefreshListeners(maxrefreshlisteners int) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxRefreshListeners = &maxrefreshlisteners

	return r
}

// API name: max_regex_length
func (r *PutDataStreamSettings) MaxRegexLength(maxregexlength int) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxRegexLength = &maxregexlength

	return r
}

// API name: max_rescore_window
func (r *PutDataStreamSettings) MaxRescoreWindow(maxrescorewindow int) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxRescoreWindow = &maxrescorewindow

	return r
}

// API name: max_result_window
func (r *PutDataStreamSettings) MaxResultWindow(maxresultwindow int) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxResultWindow = &maxresultwindow

	return r
}

// API name: max_script_fields
func (r *PutDataStreamSettings) MaxScriptFields(maxscriptfields int) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxScriptFields = &maxscriptfields

	return r
}

// API name: max_shingle_diff
func (r *PutDataStreamSettings) MaxShingleDiff(maxshinglediff int) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxShingleDiff = &maxshinglediff

	return r
}

// API name: max_slices_per_scroll
func (r *PutDataStreamSettings) MaxSlicesPerScroll(maxslicesperscroll int) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxSlicesPerScroll = &maxslicesperscroll

	return r
}

// API name: max_terms_count
func (r *PutDataStreamSettings) MaxTermsCount(maxtermscount int) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxTermsCount = &maxtermscount

	return r
}

// API name: merge
func (r *PutDataStreamSettings) Merge(merge *types.Merge) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Merge = merge

	return r
}

// API name: mode
func (r *PutDataStreamSettings) Mode(mode string) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Mode = &mode

	return r
}

// API name: number_of_replicas
func (r *PutDataStreamSettings) NumberOfReplicas(numberofreplicas string) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.NumberOfReplicas = &numberofreplicas

	return r
}

// API name: number_of_routing_shards
func (r *PutDataStreamSettings) NumberOfRoutingShards(numberofroutingshards int) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.NumberOfRoutingShards = &numberofroutingshards

	return r
}

// API name: number_of_shards
func (r *PutDataStreamSettings) NumberOfShards(numberofshards string) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.NumberOfShards = &numberofshards

	return r
}

// API name: priority
func (r *PutDataStreamSettings) Priority(priority string) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Priority = &priority

	return r
}

// API name: provided_name
func (r *PutDataStreamSettings) ProvidedName(name string) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ProvidedName = &name

	return r
}

// API name: queries
func (r *PutDataStreamSettings) Queries(queries *types.Queries) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Queries = queries

	return r
}

// API name: query_string
func (r *PutDataStreamSettings) QueryString(querystring *types.SettingsQueryString) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.QueryString = querystring

	return r
}

// API name: refresh_interval
func (r *PutDataStreamSettings) RefreshInterval(duration types.Duration) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RefreshInterval = duration

	return r
}

// API name: routing
func (r *PutDataStreamSettings) Routing(routing *types.IndexRouting) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Routing = routing

	return r
}

// API name: routing_partition_size
func (r *PutDataStreamSettings) RoutingPartitionSize(stringifiedinteger types.Stringifiedinteger) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RoutingPartitionSize = stringifiedinteger

	return r
}

// API name: routing_path
func (r *PutDataStreamSettings) RoutingPath(routingpaths ...string) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RoutingPath = routingpaths

	return r
}

// API name: search
func (r *PutDataStreamSettings) Search(search *types.SettingsSearch) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Search = search

	return r
}

// API name: settings
func (r *PutDataStreamSettings) Settings(settings *types.IndexSettings) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Settings = settings

	return r
}

// Similarity Configure custom similarity settings to customize how search results are
// scored.
// API name: similarity
func (r *PutDataStreamSettings) Similarity(similarity map[string]types.SettingsSimilarity) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Similarity = similarity

	return r
}

// API name: soft_deletes
func (r *PutDataStreamSettings) SoftDeletes(softdeletes *types.SoftDeletes) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.SoftDeletes = softdeletes

	return r
}

// API name: sort
func (r *PutDataStreamSettings) Sort(sort *types.IndexSegmentSort) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Sort = sort

	return r
}

// Store The store module allows you to control how index data is stored and accessed
// on disk.
// API name: store
func (r *PutDataStreamSettings) Store(store *types.Storage) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Store = store

	return r
}

// API name: time_series
func (r *PutDataStreamSettings) TimeSeries(timeseries *types.IndexSettingsTimeSeries) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TimeSeries = timeseries

	return r
}

// API name: top_metrics_max_size
func (r *PutDataStreamSettings) TopMetricsMaxSize(topmetricsmaxsize int) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.TopMetricsMaxSize = &topmetricsmaxsize

	return r
}

// API name: translog
func (r *PutDataStreamSettings) Translog(translog *types.Translog) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Translog = translog

	return r
}

// API name: uuid
func (r *PutDataStreamSettings) Uuid(uuid string) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Uuid = &uuid

	return r
}

// API name: verified_before_close
func (r *PutDataStreamSettings) VerifiedBeforeClose(verifiedbeforeclose string) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.VerifiedBeforeClose = &verifiedbeforeclose

	return r
}

// API name: version
func (r *PutDataStreamSettings) Version(version *types.IndexVersioning) *PutDataStreamSettings {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Version = version

	return r
}
