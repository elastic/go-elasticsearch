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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"encoding/json"
	"fmt"
)

// InferenceConfigContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/pipeline.ts#L236-L242
type InferenceConfigContainer struct {
	AdditionalInferenceConfigContainerProperty map[string]json.RawMessage `json:"-"`
	// Classification Classification configuration for inference.
	Classification *ClassificationInferenceOptions `json:"classification,omitempty"`
	// Regression Regression configuration for inference.
	Regression *RegressionInferenceOptions `json:"regression,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s InferenceConfigContainer) MarshalJSON() ([]byte, error) {
	type opt InferenceConfigContainer
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
	for key, value := range s.AdditionalInferenceConfigContainerProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalInferenceConfigContainerProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewInferenceConfigContainer returns a InferenceConfigContainer.
func NewInferenceConfigContainer() *InferenceConfigContainer {
	r := &InferenceConfigContainer{
		AdditionalInferenceConfigContainerProperty: make(map[string]json.RawMessage),
	}

	return r
}

type InferenceConfigContainerVariant interface {
	InferenceConfigContainerCaster() *InferenceConfigContainer
}

func (s *InferenceConfigContainer) InferenceConfigContainerCaster() *InferenceConfigContainer {
	return s
}
