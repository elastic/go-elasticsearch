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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexcheckonstartup"
)

// IndexSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L69-L168
type IndexSettings struct {
	Analysis *IndexSettingsAnalysis `json:"analysis,omitempty"`
	// Analyze Settings to define analyzers, tokenizers, token filters and character
	// filters.
	Analyze            *SettingsAnalyze                         `json:"analyze,omitempty"`
	AutoExpandReplicas *string                                  `json:"auto_expand_replicas,omitempty"`
	Blocks             *IndexSettingBlocks                      `json:"blocks,omitempty"`
	CheckOnStartup     *indexcheckonstartup.IndexCheckOnStartup `json:"check_on_startup,omitempty"`
	Codec              *string                                  `json:"codec,omitempty"`
	CreationDate       *StringifiedEpochTimeUnitMillis          `json:"creation_date,omitempty"`
	CreationDateString *DateTime                                `json:"creation_date_string,omitempty"`
	DefaultPipeline    *PipelineName                            `json:"default_pipeline,omitempty"`
	FinalPipeline      *PipelineName                            `json:"final_pipeline,omitempty"`
	Format             string                                   `json:"format,omitempty"`
	GcDeletes          *Duration                                `json:"gc_deletes,omitempty"`
	Hidden             string                                   `json:"hidden,omitempty"`
	Highlight          *SettingsHighlight                       `json:"highlight,omitempty"`
	Index              *IndexSettings                           `json:"index,omitempty"`
	IndexSettings      map[string]interface{}                   `json:"-"`
	// IndexingPressure Configure indexing back pressure limits.
	IndexingPressure              *IndexingPressure       `json:"indexing_pressure,omitempty"`
	IndexingSlowlog               *SlowlogSettings        `json:"indexing.slowlog,omitempty"`
	Lifecycle                     *IndexSettingsLifecycle `json:"lifecycle,omitempty"`
	LoadFixedBitsetFiltersEagerly *bool                   `json:"load_fixed_bitset_filters_eagerly,omitempty"`
	// Mapping Enable or disable dynamic mapping for an index.
	Mapping                 *MappingLimitSettings `json:"mapping,omitempty"`
	MaxDocvalueFieldsSearch *int                  `json:"max_docvalue_fields_search,omitempty"`
	MaxInnerResultWindow    *int                  `json:"max_inner_result_window,omitempty"`
	MaxNgramDiff            *int                  `json:"max_ngram_diff,omitempty"`
	MaxRefreshListeners     *int                  `json:"max_refresh_listeners,omitempty"`
	MaxRegexLength          *int                  `json:"max_regex_length,omitempty"`
	MaxRescoreWindow        *int                  `json:"max_rescore_window,omitempty"`
	MaxResultWindow         *int                  `json:"max_result_window,omitempty"`
	MaxScriptFields         *int                  `json:"max_script_fields,omitempty"`
	MaxShingleDiff          *int                  `json:"max_shingle_diff,omitempty"`
	MaxSlicesPerScroll      *int                  `json:"max_slices_per_scroll,omitempty"`
	MaxTermsCount           *int                  `json:"max_terms_count,omitempty"`
	Merge                   *Merge                `json:"merge,omitempty"`
	Mode                    *string               `json:"mode,omitempty"`
	NumberOfReplicas        string                `json:"number_of_replicas,omitempty"`
	NumberOfRoutingShards   *int                  `json:"number_of_routing_shards,omitempty"`
	NumberOfShards          string                `json:"number_of_shards,omitempty"`
	Priority                string                `json:"priority,omitempty"`
	ProvidedName            *Name                 `json:"provided_name,omitempty"`
	Queries                 *Queries              `json:"queries,omitempty"`
	QueryString             *SettingsQueryString  `json:"query_string,omitempty"`
	RefreshInterval         *Duration             `json:"refresh_interval,omitempty"`
	Routing                 *IndexRouting         `json:"routing,omitempty"`
	RoutingPartitionSize    *int                  `json:"routing_partition_size,omitempty"`
	RoutingPath             []string              `json:"routing_path,omitempty"`
	Search                  *SettingsSearch       `json:"search,omitempty"`
	Settings                *IndexSettings        `json:"settings,omitempty"`
	Shards                  *int                  `json:"shards,omitempty"`
	// Similarity Configure custom similarity settings to customize how search results are
	// scored.
	Similarity  *SettingsSimilarity `json:"similarity,omitempty"`
	SoftDeletes *SoftDeletes        `json:"soft_deletes,omitempty"`
	Sort        *IndexSegmentSort   `json:"sort,omitempty"`
	// Store The store module allows you to control how index data is stored and accessed
	// on disk.
	Store               *Storage                 `json:"store,omitempty"`
	TimeSeries          *IndexSettingsTimeSeries `json:"time_series,omitempty"`
	TopMetricsMaxSize   *int                     `json:"top_metrics_max_size,omitempty"`
	Translog            *Translog                `json:"translog,omitempty"`
	Uuid                *Uuid                    `json:"uuid,omitempty"`
	VerifiedBeforeClose string                   `json:"verified_before_close,omitempty"`
	Version             *IndexVersioning         `json:"version,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s IndexSettings) MarshalJSON() ([]byte, error) {
	type opt IndexSettings
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.IndexSettings {
		tmp[string(key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// IndexSettingsBuilder holds IndexSettings struct and provides a builder API.
type IndexSettingsBuilder struct {
	v *IndexSettings
}

// NewIndexSettings provides a builder for the IndexSettings struct.
func NewIndexSettingsBuilder() *IndexSettingsBuilder {
	r := IndexSettingsBuilder{
		&IndexSettings{
			IndexSettings: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IndexSettings struct
func (rb *IndexSettingsBuilder) Build() IndexSettings {
	return *rb.v
}

func (rb *IndexSettingsBuilder) Analysis(analysis *IndexSettingsAnalysisBuilder) *IndexSettingsBuilder {
	v := analysis.Build()
	rb.v.Analysis = &v
	return rb
}

// Analyze Settings to define analyzers, tokenizers, token filters and character
// filters.

func (rb *IndexSettingsBuilder) Analyze(analyze *SettingsAnalyzeBuilder) *IndexSettingsBuilder {
	v := analyze.Build()
	rb.v.Analyze = &v
	return rb
}

func (rb *IndexSettingsBuilder) AutoExpandReplicas(autoexpandreplicas string) *IndexSettingsBuilder {
	rb.v.AutoExpandReplicas = &autoexpandreplicas
	return rb
}

func (rb *IndexSettingsBuilder) Blocks(blocks *IndexSettingBlocksBuilder) *IndexSettingsBuilder {
	v := blocks.Build()
	rb.v.Blocks = &v
	return rb
}

func (rb *IndexSettingsBuilder) CheckOnStartup(checkonstartup indexcheckonstartup.IndexCheckOnStartup) *IndexSettingsBuilder {
	rb.v.CheckOnStartup = &checkonstartup
	return rb
}

func (rb *IndexSettingsBuilder) Codec(codec string) *IndexSettingsBuilder {
	rb.v.Codec = &codec
	return rb
}

func (rb *IndexSettingsBuilder) CreationDate(creationdate *StringifiedEpochTimeUnitMillisBuilder) *IndexSettingsBuilder {
	v := creationdate.Build()
	rb.v.CreationDate = &v
	return rb
}

func (rb *IndexSettingsBuilder) CreationDateString(creationdatestring *DateTimeBuilder) *IndexSettingsBuilder {
	v := creationdatestring.Build()
	rb.v.CreationDateString = &v
	return rb
}

func (rb *IndexSettingsBuilder) DefaultPipeline(defaultpipeline PipelineName) *IndexSettingsBuilder {
	rb.v.DefaultPipeline = &defaultpipeline
	return rb
}

func (rb *IndexSettingsBuilder) FinalPipeline(finalpipeline PipelineName) *IndexSettingsBuilder {
	rb.v.FinalPipeline = &finalpipeline
	return rb
}

func (rb *IndexSettingsBuilder) Format(arg string) *IndexSettingsBuilder {
	rb.v.Format = arg
	return rb
}

func (rb *IndexSettingsBuilder) GcDeletes(gcdeletes *DurationBuilder) *IndexSettingsBuilder {
	v := gcdeletes.Build()
	rb.v.GcDeletes = &v
	return rb
}

func (rb *IndexSettingsBuilder) Hidden(arg string) *IndexSettingsBuilder {
	rb.v.Hidden = arg
	return rb
}

func (rb *IndexSettingsBuilder) Highlight(highlight *SettingsHighlightBuilder) *IndexSettingsBuilder {
	v := highlight.Build()
	rb.v.Highlight = &v
	return rb
}

func (rb *IndexSettingsBuilder) Index(index *IndexSettingsBuilder) *IndexSettingsBuilder {
	v := index.Build()
	rb.v.Index = &v
	return rb
}

func (rb *IndexSettingsBuilder) IndexSettings(value map[string]interface{}) *IndexSettingsBuilder {
	rb.v.IndexSettings = value
	return rb
}

// IndexingPressure Configure indexing back pressure limits.

func (rb *IndexSettingsBuilder) IndexingPressure(indexingpressure *IndexingPressureBuilder) *IndexSettingsBuilder {
	v := indexingpressure.Build()
	rb.v.IndexingPressure = &v
	return rb
}

func (rb *IndexSettingsBuilder) IndexingSlowlog(indexingslowlog *SlowlogSettingsBuilder) *IndexSettingsBuilder {
	v := indexingslowlog.Build()
	rb.v.IndexingSlowlog = &v
	return rb
}

func (rb *IndexSettingsBuilder) Lifecycle(lifecycle *IndexSettingsLifecycleBuilder) *IndexSettingsBuilder {
	v := lifecycle.Build()
	rb.v.Lifecycle = &v
	return rb
}

func (rb *IndexSettingsBuilder) LoadFixedBitsetFiltersEagerly(loadfixedbitsetfilterseagerly bool) *IndexSettingsBuilder {
	rb.v.LoadFixedBitsetFiltersEagerly = &loadfixedbitsetfilterseagerly
	return rb
}

// Mapping Enable or disable dynamic mapping for an index.

func (rb *IndexSettingsBuilder) Mapping(mapping *MappingLimitSettingsBuilder) *IndexSettingsBuilder {
	v := mapping.Build()
	rb.v.Mapping = &v
	return rb
}

func (rb *IndexSettingsBuilder) MaxDocvalueFieldsSearch(maxdocvaluefieldssearch int) *IndexSettingsBuilder {
	rb.v.MaxDocvalueFieldsSearch = &maxdocvaluefieldssearch
	return rb
}

func (rb *IndexSettingsBuilder) MaxInnerResultWindow(maxinnerresultwindow int) *IndexSettingsBuilder {
	rb.v.MaxInnerResultWindow = &maxinnerresultwindow
	return rb
}

func (rb *IndexSettingsBuilder) MaxNgramDiff(maxngramdiff int) *IndexSettingsBuilder {
	rb.v.MaxNgramDiff = &maxngramdiff
	return rb
}

func (rb *IndexSettingsBuilder) MaxRefreshListeners(maxrefreshlisteners int) *IndexSettingsBuilder {
	rb.v.MaxRefreshListeners = &maxrefreshlisteners
	return rb
}

func (rb *IndexSettingsBuilder) MaxRegexLength(maxregexlength int) *IndexSettingsBuilder {
	rb.v.MaxRegexLength = &maxregexlength
	return rb
}

func (rb *IndexSettingsBuilder) MaxRescoreWindow(maxrescorewindow int) *IndexSettingsBuilder {
	rb.v.MaxRescoreWindow = &maxrescorewindow
	return rb
}

func (rb *IndexSettingsBuilder) MaxResultWindow(maxresultwindow int) *IndexSettingsBuilder {
	rb.v.MaxResultWindow = &maxresultwindow
	return rb
}

func (rb *IndexSettingsBuilder) MaxScriptFields(maxscriptfields int) *IndexSettingsBuilder {
	rb.v.MaxScriptFields = &maxscriptfields
	return rb
}

func (rb *IndexSettingsBuilder) MaxShingleDiff(maxshinglediff int) *IndexSettingsBuilder {
	rb.v.MaxShingleDiff = &maxshinglediff
	return rb
}

func (rb *IndexSettingsBuilder) MaxSlicesPerScroll(maxslicesperscroll int) *IndexSettingsBuilder {
	rb.v.MaxSlicesPerScroll = &maxslicesperscroll
	return rb
}

func (rb *IndexSettingsBuilder) MaxTermsCount(maxtermscount int) *IndexSettingsBuilder {
	rb.v.MaxTermsCount = &maxtermscount
	return rb
}

func (rb *IndexSettingsBuilder) Merge(merge *MergeBuilder) *IndexSettingsBuilder {
	v := merge.Build()
	rb.v.Merge = &v
	return rb
}

func (rb *IndexSettingsBuilder) Mode(mode string) *IndexSettingsBuilder {
	rb.v.Mode = &mode
	return rb
}

func (rb *IndexSettingsBuilder) NumberOfReplicas(arg string) *IndexSettingsBuilder {
	rb.v.NumberOfReplicas = arg
	return rb
}

func (rb *IndexSettingsBuilder) NumberOfRoutingShards(numberofroutingshards int) *IndexSettingsBuilder {
	rb.v.NumberOfRoutingShards = &numberofroutingshards
	return rb
}

func (rb *IndexSettingsBuilder) NumberOfShards(arg string) *IndexSettingsBuilder {
	rb.v.NumberOfShards = arg
	return rb
}

func (rb *IndexSettingsBuilder) Priority(arg string) *IndexSettingsBuilder {
	rb.v.Priority = arg
	return rb
}

func (rb *IndexSettingsBuilder) ProvidedName(providedname Name) *IndexSettingsBuilder {
	rb.v.ProvidedName = &providedname
	return rb
}

func (rb *IndexSettingsBuilder) Queries(queries *QueriesBuilder) *IndexSettingsBuilder {
	v := queries.Build()
	rb.v.Queries = &v
	return rb
}

func (rb *IndexSettingsBuilder) QueryString(querystring *SettingsQueryStringBuilder) *IndexSettingsBuilder {
	v := querystring.Build()
	rb.v.QueryString = &v
	return rb
}

func (rb *IndexSettingsBuilder) RefreshInterval(refreshinterval *DurationBuilder) *IndexSettingsBuilder {
	v := refreshinterval.Build()
	rb.v.RefreshInterval = &v
	return rb
}

func (rb *IndexSettingsBuilder) Routing(routing *IndexRoutingBuilder) *IndexSettingsBuilder {
	v := routing.Build()
	rb.v.Routing = &v
	return rb
}

func (rb *IndexSettingsBuilder) RoutingPartitionSize(routingpartitionsize int) *IndexSettingsBuilder {
	rb.v.RoutingPartitionSize = &routingpartitionsize
	return rb
}

func (rb *IndexSettingsBuilder) RoutingPath(arg []string) *IndexSettingsBuilder {
	rb.v.RoutingPath = arg
	return rb
}

func (rb *IndexSettingsBuilder) Search(search *SettingsSearchBuilder) *IndexSettingsBuilder {
	v := search.Build()
	rb.v.Search = &v
	return rb
}

func (rb *IndexSettingsBuilder) Settings(settings *IndexSettingsBuilder) *IndexSettingsBuilder {
	v := settings.Build()
	rb.v.Settings = &v
	return rb
}

func (rb *IndexSettingsBuilder) Shards(shards int) *IndexSettingsBuilder {
	rb.v.Shards = &shards
	return rb
}

// Similarity Configure custom similarity settings to customize how search results are
// scored.

func (rb *IndexSettingsBuilder) Similarity(similarity *SettingsSimilarityBuilder) *IndexSettingsBuilder {
	v := similarity.Build()
	rb.v.Similarity = &v
	return rb
}

func (rb *IndexSettingsBuilder) SoftDeletes(softdeletes *SoftDeletesBuilder) *IndexSettingsBuilder {
	v := softdeletes.Build()
	rb.v.SoftDeletes = &v
	return rb
}

func (rb *IndexSettingsBuilder) Sort(sort *IndexSegmentSortBuilder) *IndexSettingsBuilder {
	v := sort.Build()
	rb.v.Sort = &v
	return rb
}

// Store The store module allows you to control how index data is stored and accessed
// on disk.

func (rb *IndexSettingsBuilder) Store(store *StorageBuilder) *IndexSettingsBuilder {
	v := store.Build()
	rb.v.Store = &v
	return rb
}

func (rb *IndexSettingsBuilder) TimeSeries(timeseries *IndexSettingsTimeSeriesBuilder) *IndexSettingsBuilder {
	v := timeseries.Build()
	rb.v.TimeSeries = &v
	return rb
}

func (rb *IndexSettingsBuilder) TopMetricsMaxSize(topmetricsmaxsize int) *IndexSettingsBuilder {
	rb.v.TopMetricsMaxSize = &topmetricsmaxsize
	return rb
}

func (rb *IndexSettingsBuilder) Translog(translog *TranslogBuilder) *IndexSettingsBuilder {
	v := translog.Build()
	rb.v.Translog = &v
	return rb
}

func (rb *IndexSettingsBuilder) Uuid(uuid Uuid) *IndexSettingsBuilder {
	rb.v.Uuid = &uuid
	return rb
}

func (rb *IndexSettingsBuilder) VerifiedBeforeClose(arg string) *IndexSettingsBuilder {
	rb.v.VerifiedBeforeClose = arg
	return rb
}

func (rb *IndexSettingsBuilder) Version(version *IndexVersioningBuilder) *IndexSettingsBuilder {
	v := version.Build()
	rb.v.Version = &v
	return rb
}
