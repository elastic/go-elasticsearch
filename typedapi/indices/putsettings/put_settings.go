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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Updates the index settings.
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	index string
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

// Updates the index settings.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-update-settings.html
func New(tp elastictransport.Interface) *PutSettings {
	r := &PutSettings{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutSettings: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_settings")

		method = http.MethodPut
	case r.paramSet == indexMask:
		path.WriteString("/")

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
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
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
func (r PutSettings) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutSettings query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putsettings.Response
func (r PutSettings) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// Header set a key, value pair in the PutSettings headers map.
func (r *PutSettings) Header(key, value string) *PutSettings {
	r.headers.Set(key, value)

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

// Timeout Period to wait for a response. If no response is received before the
//  timeout expires, the request fails and returns an error.
// API name: timeout
func (r *PutSettings) Timeout(duration string) *PutSettings {
	r.values.Set("timeout", duration)

	return r
}

// API name: analysis
func (r *PutSettings) Analysis(analysis *types.IndexSettingsAnalysis) *PutSettings {

	r.req.Analysis = analysis

	return r
}

// Analyze Settings to define analyzers, tokenizers, token filters and character
// filters.
// API name: analyze
func (r *PutSettings) Analyze(analyze *types.SettingsAnalyze) *PutSettings {

	r.req.Analyze = analyze

	return r
}

// API name: auto_expand_replicas
func (r *PutSettings) AutoExpandReplicas(autoexpandreplicas string) *PutSettings {

	r.req.AutoExpandReplicas = &autoexpandreplicas

	return r
}

// API name: blocks
func (r *PutSettings) Blocks(blocks *types.IndexSettingBlocks) *PutSettings {

	r.req.Blocks = blocks

	return r
}

// API name: check_on_startup
func (r *PutSettings) CheckOnStartup(checkonstartup indexcheckonstartup.IndexCheckOnStartup) *PutSettings {
	r.req.CheckOnStartup = &checkonstartup

	return r
}

// API name: codec
func (r *PutSettings) Codec(codec string) *PutSettings {

	r.req.Codec = &codec

	return r
}

// API name: creation_date
func (r *PutSettings) CreationDate(stringifiedepochtimeunitmillis types.StringifiedEpochTimeUnitMillis) *PutSettings {
	r.req.CreationDate = stringifiedepochtimeunitmillis

	return r
}

// API name: creation_date_string
func (r *PutSettings) CreationDateString(datetime types.DateTime) *PutSettings {
	r.req.CreationDateString = datetime

	return r
}

// API name: default_pipeline
func (r *PutSettings) DefaultPipeline(pipelinename string) *PutSettings {
	r.req.DefaultPipeline = &pipelinename

	return r
}

// API name: final_pipeline
func (r *PutSettings) FinalPipeline(pipelinename string) *PutSettings {
	r.req.FinalPipeline = &pipelinename

	return r
}

// API name: format
func (r *PutSettings) Format(format string) *PutSettings {
	r.req.Format = format

	return r
}

// API name: gc_deletes
func (r *PutSettings) GcDeletes(duration types.Duration) *PutSettings {
	r.req.GcDeletes = duration

	return r
}

// API name: hidden
func (r *PutSettings) Hidden(hidden string) *PutSettings {
	r.req.Hidden = hidden

	return r
}

// API name: highlight
func (r *PutSettings) Highlight(highlight *types.SettingsHighlight) *PutSettings {

	r.req.Highlight = highlight

	return r
}

// API name: index
func (r *PutSettings) Index(index *types.IndexSettings) *PutSettings {

	r.req.Index = index

	return r
}

// API name: IndexSettings
func (r *PutSettings) IndexSettings(indexsettings map[string]json.RawMessage) *PutSettings {

	r.req.IndexSettings = indexsettings

	return r
}

// IndexingPressure Configure indexing back pressure limits.
// API name: indexing_pressure
func (r *PutSettings) IndexingPressure(indexingpressure *types.IndicesIndexingPressure) *PutSettings {

	r.req.IndexingPressure = indexingpressure

	return r
}

// API name: indexing.slowlog
func (r *PutSettings) IndexingSlowlog(indexingslowlog *types.IndexingSlowlogSettings) *PutSettings {

	r.req.IndexingSlowlog = indexingslowlog

	return r
}

// API name: lifecycle
func (r *PutSettings) Lifecycle(lifecycle *types.IndexSettingsLifecycle) *PutSettings {

	r.req.Lifecycle = lifecycle

	return r
}

// API name: load_fixed_bitset_filters_eagerly
func (r *PutSettings) LoadFixedBitsetFiltersEagerly(loadfixedbitsetfilterseagerly bool) *PutSettings {
	r.req.LoadFixedBitsetFiltersEagerly = &loadfixedbitsetfilterseagerly

	return r
}

// Mapping Enable or disable dynamic mapping for an index.
// API name: mapping
func (r *PutSettings) Mapping(mapping *types.MappingLimitSettings) *PutSettings {

	r.req.Mapping = mapping

	return r
}

// API name: max_docvalue_fields_search
func (r *PutSettings) MaxDocvalueFieldsSearch(maxdocvaluefieldssearch int) *PutSettings {
	r.req.MaxDocvalueFieldsSearch = &maxdocvaluefieldssearch

	return r
}

// API name: max_inner_result_window
func (r *PutSettings) MaxInnerResultWindow(maxinnerresultwindow int) *PutSettings {
	r.req.MaxInnerResultWindow = &maxinnerresultwindow

	return r
}

// API name: max_ngram_diff
func (r *PutSettings) MaxNgramDiff(maxngramdiff int) *PutSettings {
	r.req.MaxNgramDiff = &maxngramdiff

	return r
}

// API name: max_refresh_listeners
func (r *PutSettings) MaxRefreshListeners(maxrefreshlisteners int) *PutSettings {
	r.req.MaxRefreshListeners = &maxrefreshlisteners

	return r
}

// API name: max_regex_length
func (r *PutSettings) MaxRegexLength(maxregexlength int) *PutSettings {
	r.req.MaxRegexLength = &maxregexlength

	return r
}

// API name: max_rescore_window
func (r *PutSettings) MaxRescoreWindow(maxrescorewindow int) *PutSettings {
	r.req.MaxRescoreWindow = &maxrescorewindow

	return r
}

// API name: max_result_window
func (r *PutSettings) MaxResultWindow(maxresultwindow int) *PutSettings {
	r.req.MaxResultWindow = &maxresultwindow

	return r
}

// API name: max_script_fields
func (r *PutSettings) MaxScriptFields(maxscriptfields int) *PutSettings {
	r.req.MaxScriptFields = &maxscriptfields

	return r
}

// API name: max_shingle_diff
func (r *PutSettings) MaxShingleDiff(maxshinglediff int) *PutSettings {
	r.req.MaxShingleDiff = &maxshinglediff

	return r
}

// API name: max_slices_per_scroll
func (r *PutSettings) MaxSlicesPerScroll(maxslicesperscroll int) *PutSettings {
	r.req.MaxSlicesPerScroll = &maxslicesperscroll

	return r
}

// API name: max_terms_count
func (r *PutSettings) MaxTermsCount(maxtermscount int) *PutSettings {
	r.req.MaxTermsCount = &maxtermscount

	return r
}

// API name: merge
func (r *PutSettings) Merge(merge *types.Merge) *PutSettings {

	r.req.Merge = merge

	return r
}

// API name: mode
func (r *PutSettings) Mode(mode string) *PutSettings {

	r.req.Mode = &mode

	return r
}

// API name: number_of_replicas
func (r *PutSettings) NumberOfReplicas(numberofreplicas string) *PutSettings {
	r.req.NumberOfReplicas = numberofreplicas

	return r
}

// API name: number_of_routing_shards
func (r *PutSettings) NumberOfRoutingShards(numberofroutingshards int) *PutSettings {
	r.req.NumberOfRoutingShards = &numberofroutingshards

	return r
}

// API name: number_of_shards
func (r *PutSettings) NumberOfShards(numberofshards string) *PutSettings {
	r.req.NumberOfShards = numberofshards

	return r
}

// API name: priority
func (r *PutSettings) Priority(priority string) *PutSettings {
	r.req.Priority = priority

	return r
}

// API name: provided_name
func (r *PutSettings) ProvidedName(name string) *PutSettings {
	r.req.ProvidedName = &name

	return r
}

// API name: queries
func (r *PutSettings) Queries(queries *types.Queries) *PutSettings {

	r.req.Queries = queries

	return r
}

// API name: query_string
func (r *PutSettings) QueryString(querystring *types.SettingsQueryString) *PutSettings {

	r.req.QueryString = querystring

	return r
}

// API name: refresh_interval
func (r *PutSettings) RefreshInterval(duration types.Duration) *PutSettings {
	r.req.RefreshInterval = duration

	return r
}

// API name: routing
func (r *PutSettings) Routing(routing *types.IndexRouting) *PutSettings {

	r.req.Routing = routing

	return r
}

// API name: routing_partition_size
func (r *PutSettings) RoutingPartitionSize(stringifiedinteger types.Stringifiedinteger) *PutSettings {
	r.req.RoutingPartitionSize = stringifiedinteger

	return r
}

// API name: routing_path
func (r *PutSettings) RoutingPath(routingpaths ...string) *PutSettings {
	r.req.RoutingPath = routingpaths

	return r
}

// API name: search
func (r *PutSettings) Search(search *types.SettingsSearch) *PutSettings {

	r.req.Search = search

	return r
}

// API name: settings
func (r *PutSettings) Settings(settings *types.IndexSettings) *PutSettings {

	r.req.Settings = settings

	return r
}

// API name: shards
func (r *PutSettings) Shards(shards int) *PutSettings {
	r.req.Shards = &shards

	return r
}

// Similarity Configure custom similarity settings to customize how search results are
// scored.
// API name: similarity
func (r *PutSettings) Similarity(similarity *types.SettingsSimilarity) *PutSettings {

	r.req.Similarity = similarity

	return r
}

// API name: soft_deletes
func (r *PutSettings) SoftDeletes(softdeletes *types.SoftDeletes) *PutSettings {

	r.req.SoftDeletes = softdeletes

	return r
}

// API name: sort
func (r *PutSettings) Sort(sort *types.IndexSegmentSort) *PutSettings {

	r.req.Sort = sort

	return r
}

// Store The store module allows you to control how index data is stored and accessed
// on disk.
// API name: store
func (r *PutSettings) Store(store *types.Storage) *PutSettings {

	r.req.Store = store

	return r
}

// API name: time_series
func (r *PutSettings) TimeSeries(timeseries *types.IndexSettingsTimeSeries) *PutSettings {

	r.req.TimeSeries = timeseries

	return r
}

// API name: top_metrics_max_size
func (r *PutSettings) TopMetricsMaxSize(topmetricsmaxsize int) *PutSettings {
	r.req.TopMetricsMaxSize = &topmetricsmaxsize

	return r
}

// API name: translog
func (r *PutSettings) Translog(translog *types.Translog) *PutSettings {

	r.req.Translog = translog

	return r
}

// API name: uuid
func (r *PutSettings) Uuid(uuid string) *PutSettings {
	r.req.Uuid = &uuid

	return r
}

// API name: verified_before_close
func (r *PutSettings) VerifiedBeforeClose(verifiedbeforeclose string) *PutSettings {
	r.req.VerifiedBeforeClose = verifiedbeforeclose

	return r
}

// API name: version
func (r *PutSettings) Version(version *types.IndexVersioning) *PutSettings {

	r.req.Version = version

	return r
}
