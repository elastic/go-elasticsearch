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

// AnomalyDetectors type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/info/types.ts#L44-L50
type AnomalyDetectors struct {
	CategorizationAnalyzer               CategorizationAnalyzer `json:"categorization_analyzer"`
	CategorizationExamplesLimit          int                    `json:"categorization_examples_limit"`
	DailyModelSnapshotRetentionAfterDays int                    `json:"daily_model_snapshot_retention_after_days"`
	ModelMemoryLimit                     string                 `json:"model_memory_limit"`
	ModelSnapshotRetentionDays           int                    `json:"model_snapshot_retention_days"`
}

// AnomalyDetectorsBuilder holds AnomalyDetectors struct and provides a builder API.
type AnomalyDetectorsBuilder struct {
	v *AnomalyDetectors
}

// NewAnomalyDetectors provides a builder for the AnomalyDetectors struct.
func NewAnomalyDetectorsBuilder() *AnomalyDetectorsBuilder {
	r := AnomalyDetectorsBuilder{
		&AnomalyDetectors{},
	}

	return &r
}

// Build finalize the chain and returns the AnomalyDetectors struct
func (rb *AnomalyDetectorsBuilder) Build() AnomalyDetectors {
	return *rb.v
}

func (rb *AnomalyDetectorsBuilder) CategorizationAnalyzer(categorizationanalyzer *CategorizationAnalyzerBuilder) *AnomalyDetectorsBuilder {
	v := categorizationanalyzer.Build()
	rb.v.CategorizationAnalyzer = v
	return rb
}

func (rb *AnomalyDetectorsBuilder) CategorizationExamplesLimit(categorizationexampleslimit int) *AnomalyDetectorsBuilder {
	rb.v.CategorizationExamplesLimit = categorizationexampleslimit
	return rb
}

func (rb *AnomalyDetectorsBuilder) DailyModelSnapshotRetentionAfterDays(dailymodelsnapshotretentionafterdays int) *AnomalyDetectorsBuilder {
	rb.v.DailyModelSnapshotRetentionAfterDays = dailymodelsnapshotretentionafterdays
	return rb
}

func (rb *AnomalyDetectorsBuilder) ModelMemoryLimit(modelmemorylimit string) *AnomalyDetectorsBuilder {
	rb.v.ModelMemoryLimit = modelmemorylimit
	return rb
}

func (rb *AnomalyDetectorsBuilder) ModelSnapshotRetentionDays(modelsnapshotretentiondays int) *AnomalyDetectorsBuilder {
	rb.v.ModelSnapshotRetentionDays = modelsnapshotretentiondays
	return rb
}
