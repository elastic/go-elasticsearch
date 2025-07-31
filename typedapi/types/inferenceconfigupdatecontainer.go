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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"encoding/json"
	"fmt"
)

// InferenceConfigUpdateContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/inference.ts#L315-L337
type InferenceConfigUpdateContainer struct {
	AdditionalInferenceConfigUpdateContainerProperty map[string]json.RawMessage `json:"-"`
	// Classification Classification configuration for inference.
	Classification *ClassificationInferenceOptions `json:"classification,omitempty"`
	// FillMask Fill mask configuration for inference.
	FillMask *FillMaskInferenceUpdateOptions `json:"fill_mask,omitempty"`
	// Ner Named entity recognition configuration for inference.
	Ner *NerInferenceUpdateOptions `json:"ner,omitempty"`
	// PassThrough Pass through configuration for inference.
	PassThrough *PassThroughInferenceUpdateOptions `json:"pass_through,omitempty"`
	// QuestionAnswering Question answering configuration for inference
	QuestionAnswering *QuestionAnsweringInferenceUpdateOptions `json:"question_answering,omitempty"`
	// Regression Regression configuration for inference.
	Regression *RegressionInferenceOptions `json:"regression,omitempty"`
	// TextClassification Text classification configuration for inference.
	TextClassification *TextClassificationInferenceUpdateOptions `json:"text_classification,omitempty"`
	// TextEmbedding Text embedding configuration for inference.
	TextEmbedding *TextEmbeddingInferenceUpdateOptions `json:"text_embedding,omitempty"`
	// TextExpansion Text expansion configuration for inference.
	TextExpansion *TextExpansionInferenceUpdateOptions `json:"text_expansion,omitempty"`
	// ZeroShotClassification Zeroshot classification configuration for inference.
	ZeroShotClassification *ZeroShotClassificationInferenceUpdateOptions `json:"zero_shot_classification,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s InferenceConfigUpdateContainer) MarshalJSON() ([]byte, error) {
	type opt InferenceConfigUpdateContainer
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalInferenceConfigUpdateContainerProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalInferenceConfigUpdateContainerProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewInferenceConfigUpdateContainer returns a InferenceConfigUpdateContainer.
func NewInferenceConfigUpdateContainer() *InferenceConfigUpdateContainer {
	r := &InferenceConfigUpdateContainer{
		AdditionalInferenceConfigUpdateContainerProperty: make(map[string]json.RawMessage),
	}

	return r
}
