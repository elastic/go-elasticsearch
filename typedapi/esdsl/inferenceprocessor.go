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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _inferenceProcessor struct {
	v *types.InferenceProcessor
}

// Uses a pre-trained data frame analytics model or a model deployed for natural
// language processing tasks to infer against the data that is being ingested in
// the pipeline.
func NewInferenceProcessor() *_inferenceProcessor {

	return &_inferenceProcessor{v: types.NewInferenceProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_inferenceProcessor) Description(description string) *_inferenceProcessor {

	s.v.Description = &description

	return s
}

// Maps the document field names to the known field names of the model.
// This mapping takes precedence over any default mappings provided in the model
// configuration.
func (s *_inferenceProcessor) FieldMap(fieldmap map[string]json.RawMessage) *_inferenceProcessor {

	s.v.FieldMap = fieldmap
	return s
}

func (s *_inferenceProcessor) AddFieldMap(key string, value json.RawMessage) *_inferenceProcessor {

	var tmp map[string]json.RawMessage
	if s.v.FieldMap == nil {
		s.v.FieldMap = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.FieldMap
	}

	tmp[key] = value

	s.v.FieldMap = tmp
	return s
}

// Conditionally execute the processor.
func (s *_inferenceProcessor) If(if_ types.ScriptVariant) *_inferenceProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_inferenceProcessor) IgnoreFailure(ignorefailure bool) *_inferenceProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// Contains the inference type and its options.
func (s *_inferenceProcessor) InferenceConfig(inferenceconfig types.InferenceConfigVariant) *_inferenceProcessor {

	s.v.InferenceConfig = inferenceconfig.InferenceConfigCaster()

	return s
}

// The ID or alias for the trained model, or the ID of the deployment.
func (s *_inferenceProcessor) ModelId(id string) *_inferenceProcessor {

	s.v.ModelId = id

	return s
}

// Handle failures for the processor.
func (s *_inferenceProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_inferenceProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_inferenceProcessor) Tag(tag string) *_inferenceProcessor {

	s.v.Tag = &tag

	return s
}

// Field added to incoming documents to contain results objects.
func (s *_inferenceProcessor) TargetField(field string) *_inferenceProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_inferenceProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Inference = s.v

	return container
}

func (s *_inferenceProcessor) InferenceProcessorCaster() *types.InferenceProcessor {
	return s.v
}
