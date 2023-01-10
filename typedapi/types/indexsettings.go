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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexcheckonstartup"
)

// IndexSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/indices/_types/IndexSettings.ts#L69-L168
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
	DefaultPipeline    *string                                  `json:"default_pipeline,omitempty"`
	FinalPipeline      *string                                  `json:"final_pipeline,omitempty"`
	Format             string                                   `json:"format,omitempty"`
	GcDeletes          *Duration                                `json:"gc_deletes,omitempty"`
	Hidden             string                                   `json:"hidden,omitempty"`
	Highlight          *SettingsHighlight                       `json:"highlight,omitempty"`
	Index              *IndexSettings                           `json:"index,omitempty"`
	IndexSettings      map[string]interface{}                   `json:"-"`
	// IndexingPressure Configure indexing back pressure limits.
	IndexingPressure              *IndicesIndexingPressure `json:"indexing_pressure,omitempty"`
	IndexingSlowlog               *SlowlogSettings         `json:"indexing.slowlog,omitempty"`
	Lifecycle                     *IndexSettingsLifecycle  `json:"lifecycle,omitempty"`
	LoadFixedBitsetFiltersEagerly *bool                    `json:"load_fixed_bitset_filters_eagerly,omitempty"`
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
	ProvidedName            *string               `json:"provided_name,omitempty"`
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
	Uuid                *string                  `json:"uuid,omitempty"`
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
		tmp[fmt.Sprintf("%s", key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewIndexSettings returns a IndexSettings.
func NewIndexSettings() *IndexSettings {
	r := &IndexSettings{
		IndexSettings: make(map[string]interface{}, 0),
	}

	return r
}
