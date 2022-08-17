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

// TrainedModelsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/ml_trained_models/types.ts#L23-L111
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
	Id *Id `json:"id,omitempty"`
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
	Version *VersionString `json:"version,omitempty"`
}

// TrainedModelsRecordBuilder holds TrainedModelsRecord struct and provides a builder API.
type TrainedModelsRecordBuilder struct {
	v *TrainedModelsRecord
}

// NewTrainedModelsRecord provides a builder for the TrainedModelsRecord struct.
func NewTrainedModelsRecordBuilder() *TrainedModelsRecordBuilder {
	r := TrainedModelsRecordBuilder{
		&TrainedModelsRecord{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelsRecord struct
func (rb *TrainedModelsRecordBuilder) Build() TrainedModelsRecord {
	return *rb.v
}

// CreateTime The time the model was created

func (rb *TrainedModelsRecordBuilder) CreateTime(createtime *DateTimeBuilder) *TrainedModelsRecordBuilder {
	v := createtime.Build()
	rb.v.CreateTime = &v
	return rb
}

// CreatedBy who created the model

func (rb *TrainedModelsRecordBuilder) CreatedBy(createdby string) *TrainedModelsRecordBuilder {
	rb.v.CreatedBy = &createdby
	return rb
}

// DataFrameAnalysis The analysis used by the data frame to build the model

func (rb *TrainedModelsRecordBuilder) DataFrameAnalysis(dataframeanalysis string) *TrainedModelsRecordBuilder {
	rb.v.DataFrameAnalysis = &dataframeanalysis
	return rb
}

// DataFrameCreateTime The time the data frame analytics config was created

func (rb *TrainedModelsRecordBuilder) DataFrameCreateTime(dataframecreatetime string) *TrainedModelsRecordBuilder {
	rb.v.DataFrameCreateTime = &dataframecreatetime
	return rb
}

// DataFrameId The data frame analytics config id that created the model (if still
// available)

func (rb *TrainedModelsRecordBuilder) DataFrameId(dataframeid string) *TrainedModelsRecordBuilder {
	rb.v.DataFrameId = &dataframeid
	return rb
}

// DataFrameSourceIndex The source index used to train in the data frame analysis

func (rb *TrainedModelsRecordBuilder) DataFrameSourceIndex(dataframesourceindex string) *TrainedModelsRecordBuilder {
	rb.v.DataFrameSourceIndex = &dataframesourceindex
	return rb
}

// Description The model description

func (rb *TrainedModelsRecordBuilder) Description(description string) *TrainedModelsRecordBuilder {
	rb.v.Description = &description
	return rb
}

// HeapSize the estimated heap size to keep the model in memory

func (rb *TrainedModelsRecordBuilder) HeapSize(heapsize *ByteSizeBuilder) *TrainedModelsRecordBuilder {
	v := heapsize.Build()
	rb.v.HeapSize = &v
	return rb
}

// Id the trained model id

func (rb *TrainedModelsRecordBuilder) Id(id Id) *TrainedModelsRecordBuilder {
	rb.v.Id = &id
	return rb
}

// IngestCount The total number of docs processed by the model

func (rb *TrainedModelsRecordBuilder) IngestCount(ingestcount string) *TrainedModelsRecordBuilder {
	rb.v.IngestCount = &ingestcount
	return rb
}

// IngestCurrent The total documents currently being handled by the model

func (rb *TrainedModelsRecordBuilder) IngestCurrent(ingestcurrent string) *TrainedModelsRecordBuilder {
	rb.v.IngestCurrent = &ingestcurrent
	return rb
}

// IngestFailed The total count of failed ingest attempts with this model

func (rb *TrainedModelsRecordBuilder) IngestFailed(ingestfailed string) *TrainedModelsRecordBuilder {
	rb.v.IngestFailed = &ingestfailed
	return rb
}

// IngestPipelines The number of pipelines referencing the model

func (rb *TrainedModelsRecordBuilder) IngestPipelines(ingestpipelines string) *TrainedModelsRecordBuilder {
	rb.v.IngestPipelines = &ingestpipelines
	return rb
}

// IngestTime The total time spent processing docs with this model

func (rb *TrainedModelsRecordBuilder) IngestTime(ingesttime string) *TrainedModelsRecordBuilder {
	rb.v.IngestTime = &ingesttime
	return rb
}

// License The license level of the model

func (rb *TrainedModelsRecordBuilder) License(license string) *TrainedModelsRecordBuilder {
	rb.v.License = &license
	return rb
}

// Operations the estimated number of operations to use the model

func (rb *TrainedModelsRecordBuilder) Operations(operations string) *TrainedModelsRecordBuilder {
	rb.v.Operations = &operations
	return rb
}

func (rb *TrainedModelsRecordBuilder) Type_(type_ string) *TrainedModelsRecordBuilder {
	rb.v.Type = &type_
	return rb
}

// Version The version of Elasticsearch when the model was created

func (rb *TrainedModelsRecordBuilder) Version(version VersionString) *TrainedModelsRecordBuilder {
	rb.v.Version = &version
	return rb
}
