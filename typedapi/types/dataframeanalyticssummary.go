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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// DataframeAnalyticsSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/ml/_types/DataframeAnalytics.ts#L305-L317
type DataframeAnalyticsSummary struct {
	AllowLazyStart   *bool                            `json:"allow_lazy_start,omitempty"`
	Analysis         DataframeAnalysisContainer       `json:"analysis"`
	AnalyzedFields   *DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`
	CreateTime       *int64                           `json:"create_time,omitempty"`
	Description      *string                          `json:"description,omitempty"`
	Dest             DataframeAnalyticsDestination    `json:"dest"`
	Id               Id                               `json:"id"`
	MaxNumThreads    *int                             `json:"max_num_threads,omitempty"`
	ModelMemoryLimit *string                          `json:"model_memory_limit,omitempty"`
	Source           DataframeAnalyticsSource         `json:"source"`
	Version          *VersionString                   `json:"version,omitempty"`
}

// DataframeAnalyticsSummaryBuilder holds DataframeAnalyticsSummary struct and provides a builder API.
type DataframeAnalyticsSummaryBuilder struct {
	v *DataframeAnalyticsSummary
}

// NewDataframeAnalyticsSummary provides a builder for the DataframeAnalyticsSummary struct.
func NewDataframeAnalyticsSummaryBuilder() *DataframeAnalyticsSummaryBuilder {
	r := DataframeAnalyticsSummaryBuilder{
		&DataframeAnalyticsSummary{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalyticsSummary struct
func (rb *DataframeAnalyticsSummaryBuilder) Build() DataframeAnalyticsSummary {
	return *rb.v
}

func (rb *DataframeAnalyticsSummaryBuilder) AllowLazyStart(allowlazystart bool) *DataframeAnalyticsSummaryBuilder {
	rb.v.AllowLazyStart = &allowlazystart
	return rb
}

func (rb *DataframeAnalyticsSummaryBuilder) Analysis(analysis *DataframeAnalysisContainerBuilder) *DataframeAnalyticsSummaryBuilder {
	v := analysis.Build()
	rb.v.Analysis = v
	return rb
}

func (rb *DataframeAnalyticsSummaryBuilder) AnalyzedFields(analyzedfields *DataframeAnalysisAnalyzedFieldsBuilder) *DataframeAnalyticsSummaryBuilder {
	v := analyzedfields.Build()
	rb.v.AnalyzedFields = &v
	return rb
}

func (rb *DataframeAnalyticsSummaryBuilder) CreateTime(createtime int64) *DataframeAnalyticsSummaryBuilder {
	rb.v.CreateTime = &createtime
	return rb
}

func (rb *DataframeAnalyticsSummaryBuilder) Description(description string) *DataframeAnalyticsSummaryBuilder {
	rb.v.Description = &description
	return rb
}

func (rb *DataframeAnalyticsSummaryBuilder) Dest(dest *DataframeAnalyticsDestinationBuilder) *DataframeAnalyticsSummaryBuilder {
	v := dest.Build()
	rb.v.Dest = v
	return rb
}

func (rb *DataframeAnalyticsSummaryBuilder) Id(id Id) *DataframeAnalyticsSummaryBuilder {
	rb.v.Id = id
	return rb
}

func (rb *DataframeAnalyticsSummaryBuilder) MaxNumThreads(maxnumthreads int) *DataframeAnalyticsSummaryBuilder {
	rb.v.MaxNumThreads = &maxnumthreads
	return rb
}

func (rb *DataframeAnalyticsSummaryBuilder) ModelMemoryLimit(modelmemorylimit string) *DataframeAnalyticsSummaryBuilder {
	rb.v.ModelMemoryLimit = &modelmemorylimit
	return rb
}

func (rb *DataframeAnalyticsSummaryBuilder) Source(source *DataframeAnalyticsSourceBuilder) *DataframeAnalyticsSummaryBuilder {
	v := source.Build()
	rb.v.Source = v
	return rb
}

func (rb *DataframeAnalyticsSummaryBuilder) Version(version VersionString) *DataframeAnalyticsSummaryBuilder {
	rb.v.Version = &version
	return rb
}
