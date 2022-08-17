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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/trainedmodeltype"
)

// TrainedModelConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/TrainedModel.ts#L156-L188
type TrainedModelConfig struct {
	CompressedDefinition *string `json:"compressed_definition,omitempty"`
	// CreateTime The time when the trained model was created.
	CreateTime *DateTime `json:"create_time,omitempty"`
	// CreatedBy Information on the creator of the trained model.
	CreatedBy *string `json:"created_by,omitempty"`
	// DefaultFieldMap Any field map described in the inference configuration takes precedence.
	DefaultFieldMap map[string]string `json:"default_field_map,omitempty"`
	// Description The free-text description of the trained model.
	Description *string `json:"description,omitempty"`
	// EstimatedHeapMemoryUsageBytes The estimated heap usage in bytes to keep the trained model in memory.
	EstimatedHeapMemoryUsageBytes *int `json:"estimated_heap_memory_usage_bytes,omitempty"`
	// EstimatedOperations The estimated number of operations to use the trained model.
	EstimatedOperations *int `json:"estimated_operations,omitempty"`
	// InferenceConfig The default configuration for inference. This can be either a regression,
	// classification, or one of the many NLP focused configurations. It must match
	// the underlying definition.trained_model's target_type.
	InferenceConfig InferenceConfigCreateContainer `json:"inference_config"`
	// Input The input field names for the model definition.
	Input TrainedModelConfigInput `json:"input"`
	// LicenseLevel The license level of the trained model.
	LicenseLevel *string               `json:"license_level,omitempty"`
	Location     *TrainedModelLocation `json:"location,omitempty"`
	// Metadata An object containing metadata about the trained model. For example, models
	// created by data frame analytics contain analysis_config and input objects.
	Metadata *TrainedModelConfigMetadata `json:"metadata,omitempty"`
	// ModelId Identifier for the trained model.
	ModelId        Id        `json:"model_id"`
	ModelSizeBytes *ByteSize `json:"model_size_bytes,omitempty"`
	// ModelType The model type
	ModelType *trainedmodeltype.TrainedModelType `json:"model_type,omitempty"`
	// Tags A comma delimited string of tags. A trained model can have many tags, or
	// none.
	Tags []string `json:"tags"`
	// Version The Elasticsearch version number in which the trained model was created.
	Version *VersionString `json:"version,omitempty"`
}

// TrainedModelConfigBuilder holds TrainedModelConfig struct and provides a builder API.
type TrainedModelConfigBuilder struct {
	v *TrainedModelConfig
}

// NewTrainedModelConfig provides a builder for the TrainedModelConfig struct.
func NewTrainedModelConfigBuilder() *TrainedModelConfigBuilder {
	r := TrainedModelConfigBuilder{
		&TrainedModelConfig{
			DefaultFieldMap: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelConfig struct
func (rb *TrainedModelConfigBuilder) Build() TrainedModelConfig {
	return *rb.v
}

func (rb *TrainedModelConfigBuilder) CompressedDefinition(compresseddefinition string) *TrainedModelConfigBuilder {
	rb.v.CompressedDefinition = &compresseddefinition
	return rb
}

// CreateTime The time when the trained model was created.

func (rb *TrainedModelConfigBuilder) CreateTime(createtime *DateTimeBuilder) *TrainedModelConfigBuilder {
	v := createtime.Build()
	rb.v.CreateTime = &v
	return rb
}

// CreatedBy Information on the creator of the trained model.

func (rb *TrainedModelConfigBuilder) CreatedBy(createdby string) *TrainedModelConfigBuilder {
	rb.v.CreatedBy = &createdby
	return rb
}

// DefaultFieldMap Any field map described in the inference configuration takes precedence.

func (rb *TrainedModelConfigBuilder) DefaultFieldMap(value map[string]string) *TrainedModelConfigBuilder {
	rb.v.DefaultFieldMap = value
	return rb
}

// Description The free-text description of the trained model.

func (rb *TrainedModelConfigBuilder) Description(description string) *TrainedModelConfigBuilder {
	rb.v.Description = &description
	return rb
}

// EstimatedHeapMemoryUsageBytes The estimated heap usage in bytes to keep the trained model in memory.

func (rb *TrainedModelConfigBuilder) EstimatedHeapMemoryUsageBytes(estimatedheapmemoryusagebytes int) *TrainedModelConfigBuilder {
	rb.v.EstimatedHeapMemoryUsageBytes = &estimatedheapmemoryusagebytes
	return rb
}

// EstimatedOperations The estimated number of operations to use the trained model.

func (rb *TrainedModelConfigBuilder) EstimatedOperations(estimatedoperations int) *TrainedModelConfigBuilder {
	rb.v.EstimatedOperations = &estimatedoperations
	return rb
}

// InferenceConfig The default configuration for inference. This can be either a regression,
// classification, or one of the many NLP focused configurations. It must match
// the underlying definition.trained_model's target_type.

func (rb *TrainedModelConfigBuilder) InferenceConfig(inferenceconfig *InferenceConfigCreateContainerBuilder) *TrainedModelConfigBuilder {
	v := inferenceconfig.Build()
	rb.v.InferenceConfig = v
	return rb
}

// Input The input field names for the model definition.

func (rb *TrainedModelConfigBuilder) Input(input *TrainedModelConfigInputBuilder) *TrainedModelConfigBuilder {
	v := input.Build()
	rb.v.Input = v
	return rb
}

// LicenseLevel The license level of the trained model.

func (rb *TrainedModelConfigBuilder) LicenseLevel(licenselevel string) *TrainedModelConfigBuilder {
	rb.v.LicenseLevel = &licenselevel
	return rb
}

func (rb *TrainedModelConfigBuilder) Location(location *TrainedModelLocationBuilder) *TrainedModelConfigBuilder {
	v := location.Build()
	rb.v.Location = &v
	return rb
}

// Metadata An object containing metadata about the trained model. For example, models
// created by data frame analytics contain analysis_config and input objects.

func (rb *TrainedModelConfigBuilder) Metadata(metadata *TrainedModelConfigMetadataBuilder) *TrainedModelConfigBuilder {
	v := metadata.Build()
	rb.v.Metadata = &v
	return rb
}

// ModelId Identifier for the trained model.

func (rb *TrainedModelConfigBuilder) ModelId(modelid Id) *TrainedModelConfigBuilder {
	rb.v.ModelId = modelid
	return rb
}

func (rb *TrainedModelConfigBuilder) ModelSizeBytes(modelsizebytes *ByteSizeBuilder) *TrainedModelConfigBuilder {
	v := modelsizebytes.Build()
	rb.v.ModelSizeBytes = &v
	return rb
}

// ModelType The model type

func (rb *TrainedModelConfigBuilder) ModelType(modeltype trainedmodeltype.TrainedModelType) *TrainedModelConfigBuilder {
	rb.v.ModelType = &modeltype
	return rb
}

// Tags A comma delimited string of tags. A trained model can have many tags, or
// none.

func (rb *TrainedModelConfigBuilder) Tags(tags ...string) *TrainedModelConfigBuilder {
	rb.v.Tags = tags
	return rb
}

// Version The Elasticsearch version number in which the trained model was created.

func (rb *TrainedModelConfigBuilder) Version(version VersionString) *TrainedModelConfigBuilder {
	rb.v.Version = &version
	return rb
}
