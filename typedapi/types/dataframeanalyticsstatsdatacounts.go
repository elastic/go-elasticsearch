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

// DataframeAnalyticsStatsDataCounts type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/DataframeAnalytics.ts#L361-L368
type DataframeAnalyticsStatsDataCounts struct {
	// SkippedDocsCount The number of documents that are skipped during the analysis because they
	// contained values that are not supported by the analysis. For example, outlier
	// detection does not support missing fields so it skips documents with missing
	// fields. Likewise, all types of analysis skip documents that contain arrays
	// with more than one element.
	SkippedDocsCount int `json:"skipped_docs_count"`
	// TestDocsCount The number of documents that are not used for training the model and can be
	// used for testing.
	TestDocsCount int `json:"test_docs_count"`
	// TrainingDocsCount The number of documents that are used for training the model.
	TrainingDocsCount int `json:"training_docs_count"`
}

// NewDataframeAnalyticsStatsDataCounts returns a DataframeAnalyticsStatsDataCounts.
func NewDataframeAnalyticsStatsDataCounts() *DataframeAnalyticsStatsDataCounts {
	r := &DataframeAnalyticsStatsDataCounts{}

	return r
}
