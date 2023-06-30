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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/trainedmodeltype"
)

// TrainedModelConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/TrainedModel.ts#L159-L193
type TrainedModelConfig struct {
	CompressedDefinition *string `json:"compressed_definition,omitempty"`
	// CreateTime The time when the trained model was created.
	CreateTime DateTime `json:"create_time,omitempty"`
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
	InferenceConfig *InferenceConfigCreateContainer `json:"inference_config,omitempty"`
	// Input The input field names for the model definition.
	Input TrainedModelConfigInput `json:"input"`
	// LicenseLevel The license level of the trained model.
	LicenseLevel *string               `json:"license_level,omitempty"`
	Location     *TrainedModelLocation `json:"location,omitempty"`
	// Metadata An object containing metadata about the trained model. For example, models
	// created by data frame analytics contain analysis_config and input objects.
	Metadata *TrainedModelConfigMetadata `json:"metadata,omitempty"`
	// ModelId Identifier for the trained model.
	ModelId        string   `json:"model_id"`
	ModelSizeBytes ByteSize `json:"model_size_bytes,omitempty"`
	// ModelType The model type
	ModelType *trainedmodeltype.TrainedModelType `json:"model_type,omitempty"`
	// Tags A comma delimited string of tags. A trained model can have many tags, or
	// none.
	Tags []string `json:"tags"`
	// Version The Elasticsearch version number in which the trained model was created.
	Version *string `json:"version,omitempty"`
}

func (s *TrainedModelConfig) UnmarshalJSON(data []byte) error {

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

		case "compressed_definition":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CompressedDefinition = &o

		case "create_time":
			if err := dec.Decode(&s.CreateTime); err != nil {
				return err
			}

		case "created_by":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CreatedBy = &o

		case "default_field_map":
			if s.DefaultFieldMap == nil {
				s.DefaultFieldMap = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.DefaultFieldMap); err != nil {
				return err
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "estimated_heap_memory_usage_bytes":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.EstimatedHeapMemoryUsageBytes = &value
			case float64:
				f := int(v)
				s.EstimatedHeapMemoryUsageBytes = &f
			}

		case "estimated_operations":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.EstimatedOperations = &value
			case float64:
				f := int(v)
				s.EstimatedOperations = &f
			}

		case "fully_defined":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.FullyDefined = &value
			case bool:
				s.FullyDefined = &v
			}

		case "inference_config":
			if err := dec.Decode(&s.InferenceConfig); err != nil {
				return err
			}

		case "input":
			if err := dec.Decode(&s.Input); err != nil {
				return err
			}

		case "license_level":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.LicenseLevel = &o

		case "location":
			if err := dec.Decode(&s.Location); err != nil {
				return err
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return err
			}

		case "model_id":
			if err := dec.Decode(&s.ModelId); err != nil {
				return err
			}

		case "model_size_bytes":
			if err := dec.Decode(&s.ModelSizeBytes); err != nil {
				return err
			}

		case "model_type":
			if err := dec.Decode(&s.ModelType); err != nil {
				return err
			}

		case "tags":
			if err := dec.Decode(&s.Tags); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewTrainedModelConfig returns a TrainedModelConfig.
func NewTrainedModelConfig() *TrainedModelConfig {
	r := &TrainedModelConfig{
		DefaultFieldMap: make(map[string]string, 0),
	}

	return r
}
