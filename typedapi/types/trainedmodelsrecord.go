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

// TrainedModelsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cat/ml_trained_models/types.ts#L23-L111
type TrainedModelsRecord struct {
	// CreateTime The time the model was created
	CreateTime *DateTime `json:"create_time,omitempty"`
	// CreatedBy who created the model
	CreatedBy *string `json:"created_by,omitempty"`
	// DataFrameAnalysis The analysis used by the data frame to build the model
	DataFrameAnalysis *string `json:"data_frame.analysis,omitempty"`
	// DataFrameCreateTime The time the data frame analytics config was created
	DataFrameCreateTime *string `json:"data_frame.create_time,omitempty"`
	// DataFrameId The data frame analytics config id that created the model (if still
	// available)
	DataFrameId *string `json:"data_frame.id,omitempty"`
	// DataFrameSourceIndex The source index used to train in the data frame analysis
	DataFrameSourceIndex *string `json:"data_frame.source_index,omitempty"`
	// Description The model description
	Description *string `json:"description,omitempty"`
	// HeapSize the estimated heap size to keep the model in memory
	HeapSize *ByteSize `json:"heap_size,omitempty"`
	// Id the trained model id
	Id *string `json:"id,omitempty"`
	// IngestCount The total number of docs processed by the model
	IngestCount *string `json:"ingest.count,omitempty"`
	// IngestCurrent The total documents currently being handled by the model
	IngestCurrent *string `json:"ingest.current,omitempty"`
	// IngestFailed The total count of failed ingest attempts with this model
	IngestFailed *string `json:"ingest.failed,omitempty"`
	// IngestPipelines The number of pipelines referencing the model
	IngestPipelines *string `json:"ingest.pipelines,omitempty"`
	// IngestTime The total time spent processing docs with this model
	IngestTime *string `json:"ingest.time,omitempty"`
	// License The license level of the model
	License *string `json:"license,omitempty"`
	// Operations the estimated number of operations to use the model
	Operations *string `json:"operations,omitempty"`
	Type       *string `json:"type,omitempty"`
	// Version The version of Elasticsearch when the model was created
	Version *string `json:"version,omitempty"`
}

// NewTrainedModelsRecord returns a TrainedModelsRecord.
func NewTrainedModelsRecord() *TrainedModelsRecord {
	r := &TrainedModelsRecord{}

	return r
}
