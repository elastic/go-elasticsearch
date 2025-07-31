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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexcheckonstartup"
)

// IndexSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/indices/_types/IndexSettings.ts#L70-L176
type IndexSettings struct {
	Analysis *IndexSettingsAnalysis `json:"analysis,omitempty"`
	// Analyze Settings to define analyzers, tokenizers, token filters and character
	// filters.
	Analyze            *SettingsAnalyze                         `json:"analyze,omitempty"`
	AutoExpandReplicas any                                      `json:"auto_expand_replicas,omitempty"`
	Blocks             *IndexSettingBlocks                      `json:"blocks,omitempty"`
	CheckOnStartup     *indexcheckonstartup.IndexCheckOnStartup `json:"check_on_startup,omitempty"`
	Codec              *string                                  `json:"codec,omitempty"`
	CreationDate       StringifiedEpochTimeUnitMillis           `json:"creation_date,omitempty"`
	CreationDateString DateTime                                 `json:"creation_date_string,omitempty"`
	DefaultPipeline    *string                                  `json:"default_pipeline,omitempty"`
	FinalPipeline      *string                                  `json:"final_pipeline,omitempty"`
	Format             *string                                  `json:"format,omitempty"`
	GcDeletes          Duration                                 `json:"gc_deletes,omitempty"`
	Hidden             *string                                  `json:"hidden,omitempty"`
	Highlight          *SettingsHighlight                       `json:"highlight,omitempty"`
	Index              *IndexSettings                           `json:"index,omitempty"`
	IndexSettings      map[string]json.RawMessage               `json:"-"`
	// IndexingPressure Configure indexing back pressure limits.
	IndexingPressure              *IndicesIndexingPressure `json:"indexing_pressure,omitempty"`
	IndexingSlowlog               *IndexingSlowlogSettings `json:"indexing.slowlog,omitempty"`
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
	NumberOfReplicas        *string               `json:"number_of_replicas,omitempty"`
	NumberOfRoutingShards   *int                  `json:"number_of_routing_shards,omitempty"`
	NumberOfShards          *string               `json:"number_of_shards,omitempty"`
	Priority                *string               `json:"priority,omitempty"`
	ProvidedName            *string               `json:"provided_name,omitempty"`
	Queries                 *Queries              `json:"queries,omitempty"`
	QueryString             *SettingsQueryString  `json:"query_string,omitempty"`
	RefreshInterval         Duration              `json:"refresh_interval,omitempty"`
	Routing                 *IndexRouting         `json:"routing,omitempty"`
	RoutingPartitionSize    Stringifiedinteger    `json:"routing_partition_size,omitempty"`
	RoutingPath             []string              `json:"routing_path,omitempty"`
	Search                  *SettingsSearch       `json:"search,omitempty"`
	Settings                *IndexSettings        `json:"settings,omitempty"`
	// Similarity Configure custom similarity settings to customize how search results are
	// scored.
	Similarity  map[string]SettingsSimilarity `json:"similarity,omitempty"`
	SoftDeletes *SoftDeletes                  `json:"soft_deletes,omitempty"`
	Sort        *IndexSegmentSort             `json:"sort,omitempty"`
	// Store The store module allows you to control how index data is stored and accessed
	// on disk.
	Store               *Storage                 `json:"store,omitempty"`
	TimeSeries          *IndexSettingsTimeSeries `json:"time_series,omitempty"`
	TopMetricsMaxSize   *int                     `json:"top_metrics_max_size,omitempty"`
	Translog            *Translog                `json:"translog,omitempty"`
	Uuid                *string                  `json:"uuid,omitempty"`
	VerifiedBeforeClose *string                  `json:"verified_before_close,omitempty"`
	Version             *IndexVersioning         `json:"version,omitempty"`
}

func (s *IndexSettings) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "analysis":
			if err := dec.Decode(&s.Analysis); err != nil {
				return fmt.Errorf("%s | %w", "Analysis", err)
			}

		case "analyze":
			if err := dec.Decode(&s.Analyze); err != nil {
				return fmt.Errorf("%s | %w", "Analyze", err)
			}

		case "auto_expand_replicas":
			if err := dec.Decode(&s.AutoExpandReplicas); err != nil {
				return fmt.Errorf("%s | %w", "AutoExpandReplicas", err)
			}

		case "blocks":
			if err := dec.Decode(&s.Blocks); err != nil {
				return fmt.Errorf("%s | %w", "Blocks", err)
			}

		case "check_on_startup":
			if err := dec.Decode(&s.CheckOnStartup); err != nil {
				return fmt.Errorf("%s | %w", "CheckOnStartup", err)
			}

		case "codec":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Codec", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Codec = &o

		case "creation_date":
			if err := dec.Decode(&s.CreationDate); err != nil {
				return fmt.Errorf("%s | %w", "CreationDate", err)
			}

		case "creation_date_string":
			if err := dec.Decode(&s.CreationDateString); err != nil {
				return fmt.Errorf("%s | %w", "CreationDateString", err)
			}

		case "default_pipeline":
			if err := dec.Decode(&s.DefaultPipeline); err != nil {
				return fmt.Errorf("%s | %w", "DefaultPipeline", err)
			}

		case "final_pipeline":
			if err := dec.Decode(&s.FinalPipeline); err != nil {
				return fmt.Errorf("%s | %w", "FinalPipeline", err)
			}

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Format", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Format = &o

		case "gc_deletes":
			if err := dec.Decode(&s.GcDeletes); err != nil {
				return fmt.Errorf("%s | %w", "GcDeletes", err)
			}

		case "hidden":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Hidden", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Hidden = &o

		case "highlight":
			if err := dec.Decode(&s.Highlight); err != nil {
				return fmt.Errorf("%s | %w", "Highlight", err)
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "indexing_pressure":
			if err := dec.Decode(&s.IndexingPressure); err != nil {
				return fmt.Errorf("%s | %w", "IndexingPressure", err)
			}

		case "indexing.slowlog":
			if err := dec.Decode(&s.IndexingSlowlog); err != nil {
				return fmt.Errorf("%s | %w", "IndexingSlowlog", err)
			}

		case "lifecycle":
			if err := dec.Decode(&s.Lifecycle); err != nil {
				return fmt.Errorf("%s | %w", "Lifecycle", err)
			}

		case "load_fixed_bitset_filters_eagerly":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "LoadFixedBitsetFiltersEagerly", err)
				}
				s.LoadFixedBitsetFiltersEagerly = &value
			case bool:
				s.LoadFixedBitsetFiltersEagerly = &v
			}

		case "mapping":
			if err := dec.Decode(&s.Mapping); err != nil {
				return fmt.Errorf("%s | %w", "Mapping", err)
			}

		case "max_docvalue_fields_search":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxDocvalueFieldsSearch", err)
				}
				s.MaxDocvalueFieldsSearch = &value
			case float64:
				f := int(v)
				s.MaxDocvalueFieldsSearch = &f
			}

		case "max_inner_result_window":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxInnerResultWindow", err)
				}
				s.MaxInnerResultWindow = &value
			case float64:
				f := int(v)
				s.MaxInnerResultWindow = &f
			}

		case "max_ngram_diff":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxNgramDiff", err)
				}
				s.MaxNgramDiff = &value
			case float64:
				f := int(v)
				s.MaxNgramDiff = &f
			}

		case "max_refresh_listeners":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxRefreshListeners", err)
				}
				s.MaxRefreshListeners = &value
			case float64:
				f := int(v)
				s.MaxRefreshListeners = &f
			}

		case "max_regex_length":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxRegexLength", err)
				}
				s.MaxRegexLength = &value
			case float64:
				f := int(v)
				s.MaxRegexLength = &f
			}

		case "max_rescore_window":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxRescoreWindow", err)
				}
				s.MaxRescoreWindow = &value
			case float64:
				f := int(v)
				s.MaxRescoreWindow = &f
			}

		case "max_result_window":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxResultWindow", err)
				}
				s.MaxResultWindow = &value
			case float64:
				f := int(v)
				s.MaxResultWindow = &f
			}

		case "max_script_fields":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxScriptFields", err)
				}
				s.MaxScriptFields = &value
			case float64:
				f := int(v)
				s.MaxScriptFields = &f
			}

		case "max_shingle_diff":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxShingleDiff", err)
				}
				s.MaxShingleDiff = &value
			case float64:
				f := int(v)
				s.MaxShingleDiff = &f
			}

		case "max_slices_per_scroll":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxSlicesPerScroll", err)
				}
				s.MaxSlicesPerScroll = &value
			case float64:
				f := int(v)
				s.MaxSlicesPerScroll = &f
			}

		case "max_terms_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxTermsCount", err)
				}
				s.MaxTermsCount = &value
			case float64:
				f := int(v)
				s.MaxTermsCount = &f
			}

		case "merge":
			if err := dec.Decode(&s.Merge); err != nil {
				return fmt.Errorf("%s | %w", "Merge", err)
			}

		case "mode":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Mode", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Mode = &o

		case "number_of_replicas":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "NumberOfReplicas", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NumberOfReplicas = &o

		case "number_of_routing_shards":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfRoutingShards", err)
				}
				s.NumberOfRoutingShards = &value
			case float64:
				f := int(v)
				s.NumberOfRoutingShards = &f
			}

		case "number_of_shards":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "NumberOfShards", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NumberOfShards = &o

		case "priority":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Priority", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Priority = &o

		case "provided_name":
			if err := dec.Decode(&s.ProvidedName); err != nil {
				return fmt.Errorf("%s | %w", "ProvidedName", err)
			}

		case "queries":
			if err := dec.Decode(&s.Queries); err != nil {
				return fmt.Errorf("%s | %w", "Queries", err)
			}

		case "query_string":
			if err := dec.Decode(&s.QueryString); err != nil {
				return fmt.Errorf("%s | %w", "QueryString", err)
			}

		case "refresh_interval":
			if err := dec.Decode(&s.RefreshInterval); err != nil {
				return fmt.Errorf("%s | %w", "RefreshInterval", err)
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return fmt.Errorf("%s | %w", "Routing", err)
			}

		case "routing_partition_size":
			if err := dec.Decode(&s.RoutingPartitionSize); err != nil {
				return fmt.Errorf("%s | %w", "RoutingPartitionSize", err)
			}

		case "routing_path":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "RoutingPath", err)
				}

				s.RoutingPath = append(s.RoutingPath, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.RoutingPath); err != nil {
					return fmt.Errorf("%s | %w", "RoutingPath", err)
				}
			}

		case "search":
			if err := dec.Decode(&s.Search); err != nil {
				return fmt.Errorf("%s | %w", "Search", err)
			}

		case "settings":
			if err := dec.Decode(&s.Settings); err != nil {
				return fmt.Errorf("%s | %w", "Settings", err)
			}

		case "similarity":
			if s.Similarity == nil {
				s.Similarity = make(map[string]SettingsSimilarity, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]any)
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)

				switch kind["type"] {
				case "BM25":
					oo := NewSettingsSimilarityBm25()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Similarity | %w", err)
					}
					s.Similarity[key] = oo
				case "boolean":
					oo := NewSettingsSimilarityBoolean()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Similarity | %w", err)
					}
					s.Similarity[key] = oo
				case "DFI":
					oo := NewSettingsSimilarityDfi()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Similarity | %w", err)
					}
					s.Similarity[key] = oo
				case "DFR":
					oo := NewSettingsSimilarityDfr()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Similarity | %w", err)
					}
					s.Similarity[key] = oo
				case "IB":
					oo := NewSettingsSimilarityIb()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Similarity | %w", err)
					}
					s.Similarity[key] = oo
				case "LMDirichlet":
					oo := NewSettingsSimilarityLmd()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Similarity | %w", err)
					}
					s.Similarity[key] = oo
				case "LMJelinekMercer":
					oo := NewSettingsSimilarityLmj()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Similarity | %w", err)
					}
					s.Similarity[key] = oo
				case "scripted":
					oo := NewSettingsSimilarityScripted()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Similarity | %w", err)
					}
					s.Similarity[key] = oo
				default:
					oo := new(SettingsSimilarity)
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("new(SettingsSimilarity) | %w", err)
					}
					s.Similarity[key] = oo
				}
			}

		case "soft_deletes":
			if err := dec.Decode(&s.SoftDeletes); err != nil {
				return fmt.Errorf("%s | %w", "SoftDeletes", err)
			}

		case "sort":
			if err := dec.Decode(&s.Sort); err != nil {
				return fmt.Errorf("%s | %w", "Sort", err)
			}

		case "store":
			if err := dec.Decode(&s.Store); err != nil {
				return fmt.Errorf("%s | %w", "Store", err)
			}

		case "time_series":
			if err := dec.Decode(&s.TimeSeries); err != nil {
				return fmt.Errorf("%s | %w", "TimeSeries", err)
			}

		case "top_metrics_max_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TopMetricsMaxSize", err)
				}
				s.TopMetricsMaxSize = &value
			case float64:
				f := int(v)
				s.TopMetricsMaxSize = &f
			}

		case "translog":
			if err := dec.Decode(&s.Translog); err != nil {
				return fmt.Errorf("%s | %w", "Translog", err)
			}

		case "uuid":
			if err := dec.Decode(&s.Uuid); err != nil {
				return fmt.Errorf("%s | %w", "Uuid", err)
			}

		case "verified_before_close":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "VerifiedBeforeClose", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.VerifiedBeforeClose = &o

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		default:

			if key, ok := t.(string); ok {
				if s.IndexSettings == nil {
					s.IndexSettings = make(map[string]json.RawMessage, 0)
				}
				raw := new(json.RawMessage)
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "IndexSettings", err)
				}
				s.IndexSettings[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s IndexSettings) MarshalJSON() ([]byte, error) {
	type opt IndexSettings
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

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
	delete(tmp, "IndexSettings")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewIndexSettings returns a IndexSettings.
func NewIndexSettings() *IndexSettings {
	r := &IndexSettings{
		IndexSettings: make(map[string]json.RawMessage),
		Similarity:    make(map[string]SettingsSimilarity),
	}

	return r
}
