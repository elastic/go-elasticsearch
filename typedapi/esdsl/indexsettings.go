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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexcheckonstartup"
)

type _indexSettings struct {
	v *types.IndexSettings
}

func NewIndexSettings() *_indexSettings {

	return &_indexSettings{v: types.NewIndexSettings()}

}

func (s *_indexSettings) Analysis(analysis types.IndexSettingsAnalysisVariant) *_indexSettings {

	s.v.Analysis = analysis.IndexSettingsAnalysisCaster()

	return s
}

func (s *_indexSettings) Analyze(analyze types.SettingsAnalyzeVariant) *_indexSettings {

	s.v.Analyze = analyze.SettingsAnalyzeCaster()

	return s
}

func (s *_indexSettings) AutoExpandReplicas(autoexpandreplicas any) *_indexSettings {

	s.v.AutoExpandReplicas = autoexpandreplicas

	return s
}

func (s *_indexSettings) Blocks(blocks types.IndexSettingBlocksVariant) *_indexSettings {

	s.v.Blocks = blocks.IndexSettingBlocksCaster()

	return s
}

func (s *_indexSettings) CheckOnStartup(checkonstartup indexcheckonstartup.IndexCheckOnStartup) *_indexSettings {

	s.v.CheckOnStartup = &checkonstartup
	return s
}

func (s *_indexSettings) Codec(codec string) *_indexSettings {

	s.v.Codec = &codec

	return s
}

func (s *_indexSettings) CreationDate(stringifiedepochtimeunitmillis types.StringifiedEpochTimeUnitMillisVariant) *_indexSettings {

	s.v.CreationDate = *stringifiedepochtimeunitmillis.StringifiedEpochTimeUnitMillisCaster()

	return s
}

func (s *_indexSettings) CreationDateString(datetime types.DateTimeVariant) *_indexSettings {

	s.v.CreationDateString = *datetime.DateTimeCaster()

	return s
}

func (s *_indexSettings) DefaultPipeline(pipelinename string) *_indexSettings {

	s.v.DefaultPipeline = &pipelinename

	return s
}

func (s *_indexSettings) FinalPipeline(pipelinename string) *_indexSettings {

	s.v.FinalPipeline = &pipelinename

	return s
}

func (s *_indexSettings) Format(format string) *_indexSettings {

	s.v.Format = &format

	return s
}

func (s *_indexSettings) GcDeletes(duration types.DurationVariant) *_indexSettings {

	s.v.GcDeletes = *duration.DurationCaster()

	return s
}

func (s *_indexSettings) Hidden(hidden string) *_indexSettings {

	s.v.Hidden = &hidden

	return s
}

func (s *_indexSettings) Highlight(highlight types.SettingsHighlightVariant) *_indexSettings {

	s.v.Highlight = highlight.SettingsHighlightCaster()

	return s
}

func (s *_indexSettings) Index(index types.IndexSettingsVariant) *_indexSettings {

	s.v.Index = index.IndexSettingsCaster()

	return s
}

func (s *_indexSettings) IndexSettings(indexsettings map[string]json.RawMessage) *_indexSettings {

	s.v.IndexSettings = indexsettings
	return s
}

func (s *_indexSettings) AddIndexSetting(key string, value json.RawMessage) *_indexSettings {

	var tmp map[string]json.RawMessage
	if s.v.IndexSettings == nil {
		s.v.IndexSettings = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.IndexSettings
	}

	tmp[key] = value

	s.v.IndexSettings = tmp
	return s
}

func (s *_indexSettings) IndexingPressure(indexingpressure types.IndicesIndexingPressureVariant) *_indexSettings {

	s.v.IndexingPressure = indexingpressure.IndicesIndexingPressureCaster()

	return s
}

func (s *_indexSettings) IndexingSlowlog(indexingslowlog types.IndexingSlowlogSettingsVariant) *_indexSettings {

	s.v.IndexingSlowlog = indexingslowlog.IndexingSlowlogSettingsCaster()

	return s
}

func (s *_indexSettings) Lifecycle(lifecycle types.IndexSettingsLifecycleVariant) *_indexSettings {

	s.v.Lifecycle = lifecycle.IndexSettingsLifecycleCaster()

	return s
}

func (s *_indexSettings) LoadFixedBitsetFiltersEagerly(loadfixedbitsetfilterseagerly bool) *_indexSettings {

	s.v.LoadFixedBitsetFiltersEagerly = &loadfixedbitsetfilterseagerly

	return s
}

func (s *_indexSettings) Mapping(mapping types.MappingLimitSettingsVariant) *_indexSettings {

	s.v.Mapping = mapping.MappingLimitSettingsCaster()

	return s
}

func (s *_indexSettings) MaxDocvalueFieldsSearch(maxdocvaluefieldssearch int) *_indexSettings {

	s.v.MaxDocvalueFieldsSearch = &maxdocvaluefieldssearch

	return s
}

func (s *_indexSettings) MaxInnerResultWindow(maxinnerresultwindow int) *_indexSettings {

	s.v.MaxInnerResultWindow = &maxinnerresultwindow

	return s
}

func (s *_indexSettings) MaxNgramDiff(maxngramdiff int) *_indexSettings {

	s.v.MaxNgramDiff = &maxngramdiff

	return s
}

func (s *_indexSettings) MaxRefreshListeners(maxrefreshlisteners int) *_indexSettings {

	s.v.MaxRefreshListeners = &maxrefreshlisteners

	return s
}

func (s *_indexSettings) MaxRegexLength(maxregexlength int) *_indexSettings {

	s.v.MaxRegexLength = &maxregexlength

	return s
}

func (s *_indexSettings) MaxRescoreWindow(maxrescorewindow int) *_indexSettings {

	s.v.MaxRescoreWindow = &maxrescorewindow

	return s
}

func (s *_indexSettings) MaxResultWindow(maxresultwindow int) *_indexSettings {

	s.v.MaxResultWindow = &maxresultwindow

	return s
}

func (s *_indexSettings) MaxScriptFields(maxscriptfields int) *_indexSettings {

	s.v.MaxScriptFields = &maxscriptfields

	return s
}

func (s *_indexSettings) MaxShingleDiff(maxshinglediff int) *_indexSettings {

	s.v.MaxShingleDiff = &maxshinglediff

	return s
}

func (s *_indexSettings) MaxSlicesPerScroll(maxslicesperscroll int) *_indexSettings {

	s.v.MaxSlicesPerScroll = &maxslicesperscroll

	return s
}

func (s *_indexSettings) MaxTermsCount(maxtermscount int) *_indexSettings {

	s.v.MaxTermsCount = &maxtermscount

	return s
}

func (s *_indexSettings) Merge(merge types.MergeVariant) *_indexSettings {

	s.v.Merge = merge.MergeCaster()

	return s
}

func (s *_indexSettings) Mode(mode string) *_indexSettings {

	s.v.Mode = &mode

	return s
}

func (s *_indexSettings) NumberOfReplicas(numberofreplicas string) *_indexSettings {

	s.v.NumberOfReplicas = &numberofreplicas

	return s
}

func (s *_indexSettings) NumberOfRoutingShards(numberofroutingshards int) *_indexSettings {

	s.v.NumberOfRoutingShards = &numberofroutingshards

	return s
}

func (s *_indexSettings) NumberOfShards(numberofshards string) *_indexSettings {

	s.v.NumberOfShards = &numberofshards

	return s
}

func (s *_indexSettings) Priority(priority string) *_indexSettings {

	s.v.Priority = &priority

	return s
}

func (s *_indexSettings) ProvidedName(name string) *_indexSettings {

	s.v.ProvidedName = &name

	return s
}

func (s *_indexSettings) Queries(queries types.QueriesVariant) *_indexSettings {

	s.v.Queries = queries.QueriesCaster()

	return s
}

func (s *_indexSettings) QueryString(querystring types.SettingsQueryStringVariant) *_indexSettings {

	s.v.QueryString = querystring.SettingsQueryStringCaster()

	return s
}

func (s *_indexSettings) RefreshInterval(duration types.DurationVariant) *_indexSettings {

	s.v.RefreshInterval = *duration.DurationCaster()

	return s
}

func (s *_indexSettings) Routing(routing types.IndexRoutingVariant) *_indexSettings {

	s.v.Routing = routing.IndexRoutingCaster()

	return s
}

func (s *_indexSettings) RoutingPartitionSize(stringifiedinteger types.StringifiedintegerVariant) *_indexSettings {

	s.v.RoutingPartitionSize = *stringifiedinteger.StringifiedintegerCaster()

	return s
}

func (s *_indexSettings) RoutingPath(routingpaths ...string) *_indexSettings {

	s.v.RoutingPath = make([]string, len(routingpaths))
	s.v.RoutingPath = routingpaths

	return s
}

func (s *_indexSettings) Search(search types.SettingsSearchVariant) *_indexSettings {

	s.v.Search = search.SettingsSearchCaster()

	return s
}

func (s *_indexSettings) Settings(settings types.IndexSettingsVariant) *_indexSettings {

	s.v.Settings = settings.IndexSettingsCaster()

	return s
}

func (s *_indexSettings) Similarity(similarity map[string]types.SettingsSimilarity) *_indexSettings {

	s.v.Similarity = similarity
	return s
}

func (s *_indexSettings) AddSimilarity(key string, value types.SettingsSimilarityVariant) *_indexSettings {

	var tmp map[string]types.SettingsSimilarity
	if s.v.Similarity == nil {
		s.v.Similarity = make(map[string]types.SettingsSimilarity)
	} else {
		tmp = s.v.Similarity
	}

	tmp[key] = *value.SettingsSimilarityCaster()

	s.v.Similarity = tmp
	return s
}

func (s *_indexSettings) SoftDeletes(softdeletes types.SoftDeletesVariant) *_indexSettings {

	s.v.SoftDeletes = softdeletes.SoftDeletesCaster()

	return s
}

func (s *_indexSettings) Sort(sort types.IndexSegmentSortVariant) *_indexSettings {

	s.v.Sort = sort.IndexSegmentSortCaster()

	return s
}

func (s *_indexSettings) Store(store types.StorageVariant) *_indexSettings {

	s.v.Store = store.StorageCaster()

	return s
}

func (s *_indexSettings) TimeSeries(timeseries types.IndexSettingsTimeSeriesVariant) *_indexSettings {

	s.v.TimeSeries = timeseries.IndexSettingsTimeSeriesCaster()

	return s
}

func (s *_indexSettings) TopMetricsMaxSize(topmetricsmaxsize int) *_indexSettings {

	s.v.TopMetricsMaxSize = &topmetricsmaxsize

	return s
}

func (s *_indexSettings) Translog(translog types.TranslogVariant) *_indexSettings {

	s.v.Translog = translog.TranslogCaster()

	return s
}

func (s *_indexSettings) Uuid(uuid string) *_indexSettings {

	s.v.Uuid = &uuid

	return s
}

func (s *_indexSettings) VerifiedBeforeClose(verifiedbeforeclose string) *_indexSettings {

	s.v.VerifiedBeforeClose = &verifiedbeforeclose

	return s
}

func (s *_indexSettings) Version(version types.IndexVersioningVariant) *_indexSettings {

	s.v.Version = version.IndexVersioningCaster()

	return s
}

func (s *_indexSettings) IndexSettingsCaster() *types.IndexSettings {
	return s.v
}
