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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package puttrainedmodel

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/trainedmodeltype"
)

// Response holds the response body struct for the package puttrainedmodel
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/put_trained_model/MlPutTrainedModelResponse.ts#L22-L24
type Response struct {
	CompressedDefinition *string `json:"compressed_definition,omitempty"`
	// CreateTime The time when the trained model was created.
	CreateTime types.DateTime `json:"create_time,omitempty"`
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
	// FullyDefined True if the full model definition is present.
	FullyDefined *bool `json:"fully_defined,omitempty"`
	// InferenceConfig The default configuration for inference. This can be either a regression,
	// classification, or one of the many NLP focused configurations. It must match
	// the underlying definition.trained_model's target_type. For pre-packaged
	// models such as ELSER the config is not required.
	InferenceConfig *types.InferenceConfigCreateContainer `json:"inference_config,omitempty"`
	// Input The input field names for the model definition.
	Input types.TrainedModelConfigInput `json:"input"`
	// LicenseLevel The license level of the trained model.
	LicenseLevel *string                     `json:"license_level,omitempty"`
	Location     *types.TrainedModelLocation `json:"location,omitempty"`
	// Metadata An object containing metadata about the trained model. For example, models
	// created by data frame analytics contain analysis_config and input objects.
	Metadata *types.TrainedModelConfigMetadata `json:"metadata,omitempty"`
	// ModelId Identifier for the trained model.
	ModelId        string         `json:"model_id"`
	ModelSizeBytes types.ByteSize `json:"model_size_bytes,omitempty"`
	// ModelType The model type
	ModelType *trainedmodeltype.TrainedModelType `json:"model_type,omitempty"`
	// Tags A comma delimited string of tags. A trained model can have many tags, or
	// none.
	Tags []string `json:"tags"`
	// Version The Elasticsearch version number in which the trained model was created.
	Version *string `json:"version,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		DefaultFieldMap: make(map[string]string, 0),
	}
	return r
}
