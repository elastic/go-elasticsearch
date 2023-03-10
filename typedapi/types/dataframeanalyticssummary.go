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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// DataframeAnalyticsSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/DataframeAnalytics.ts#L306-L322
type DataframeAnalyticsSummary struct {
	AllowLazyStart *bool                            `json:"allow_lazy_start,omitempty"`
	Analysis       DataframeAnalysisContainer       `json:"analysis"`
	AnalyzedFields *DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`
	// Authorization The security privileges that the job uses to run its queries. If Elastic
	// Stack security features were disabled at the time of the most recent update
	// to the job, this property is omitted.
	Authorization    *DataframeAnalyticsAuthorization `json:"authorization,omitempty"`
	CreateTime       *int64                           `json:"create_time,omitempty"`
	Description      *string                          `json:"description,omitempty"`
	Dest             DataframeAnalyticsDestination    `json:"dest"`
	Id               string                           `json:"id"`
	MaxNumThreads    *int                             `json:"max_num_threads,omitempty"`
	ModelMemoryLimit *string                          `json:"model_memory_limit,omitempty"`
	Source           DataframeAnalyticsSource         `json:"source"`
	Version          *string                          `json:"version,omitempty"`
}

// NewDataframeAnalyticsSummary returns a DataframeAnalyticsSummary.
func NewDataframeAnalyticsSummary() *DataframeAnalyticsSummary {
	r := &DataframeAnalyticsSummary{}

	return r
}
