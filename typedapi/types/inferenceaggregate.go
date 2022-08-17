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
	"encoding/json"
)

// InferenceAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L610-L621
type InferenceAggregate struct {
	Data              map[string]interface{}       `json:"-"`
	FeatureImportance []InferenceFeatureImportance `json:"feature_importance,omitempty"`
	Meta              *Metadata                    `json:"meta,omitempty"`
	TopClasses        []InferenceTopClassEntry     `json:"top_classes,omitempty"`
	Value             *FieldValue                  `json:"value,omitempty"`
	Warning           *string                      `json:"warning,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s InferenceAggregate) MarshalJSON() ([]byte, error) {
	type opt InferenceAggregate
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.Data {
		tmp[string(key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// InferenceAggregateBuilder holds InferenceAggregate struct and provides a builder API.
type InferenceAggregateBuilder struct {
	v *InferenceAggregate
}

// NewInferenceAggregate provides a builder for the InferenceAggregate struct.
func NewInferenceAggregateBuilder() *InferenceAggregateBuilder {
	r := InferenceAggregateBuilder{
		&InferenceAggregate{
			Data: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the InferenceAggregate struct
func (rb *InferenceAggregateBuilder) Build() InferenceAggregate {
	return *rb.v
}

func (rb *InferenceAggregateBuilder) Data(value map[string]interface{}) *InferenceAggregateBuilder {
	rb.v.Data = value
	return rb
}

func (rb *InferenceAggregateBuilder) FeatureImportance(feature_importance []InferenceFeatureImportanceBuilder) *InferenceAggregateBuilder {
	tmp := make([]InferenceFeatureImportance, len(feature_importance))
	for _, value := range feature_importance {
		tmp = append(tmp, value.Build())
	}
	rb.v.FeatureImportance = tmp
	return rb
}

func (rb *InferenceAggregateBuilder) Meta(meta *MetadataBuilder) *InferenceAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *InferenceAggregateBuilder) TopClasses(top_classes []InferenceTopClassEntryBuilder) *InferenceAggregateBuilder {
	tmp := make([]InferenceTopClassEntry, len(top_classes))
	for _, value := range top_classes {
		tmp = append(tmp, value.Build())
	}
	rb.v.TopClasses = tmp
	return rb
}

func (rb *InferenceAggregateBuilder) Value(value *FieldValueBuilder) *InferenceAggregateBuilder {
	v := value.Build()
	rb.v.Value = &v
	return rb
}

func (rb *InferenceAggregateBuilder) Warning(warning string) *InferenceAggregateBuilder {
	rb.v.Warning = &warning
	return rb
}
