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

// Limits type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/info/types.ts#L34-L38
type Limits struct {
	EffectiveMaxModelMemoryLimit string  `json:"effective_max_model_memory_limit"`
	MaxModelMemoryLimit          *string `json:"max_model_memory_limit,omitempty"`
	TotalMlMemory                string  `json:"total_ml_memory"`
}

// LimitsBuilder holds Limits struct and provides a builder API.
type LimitsBuilder struct {
	v *Limits
}

// NewLimits provides a builder for the Limits struct.
func NewLimitsBuilder() *LimitsBuilder {
	r := LimitsBuilder{
		&Limits{},
	}

	return &r
}

// Build finalize the chain and returns the Limits struct
func (rb *LimitsBuilder) Build() Limits {
	return *rb.v
}

func (rb *LimitsBuilder) EffectiveMaxModelMemoryLimit(effectivemaxmodelmemorylimit string) *LimitsBuilder {
	rb.v.EffectiveMaxModelMemoryLimit = effectivemaxmodelmemorylimit
	return rb
}

func (rb *LimitsBuilder) MaxModelMemoryLimit(maxmodelmemorylimit string) *LimitsBuilder {
	rb.v.MaxModelMemoryLimit = &maxmodelmemorylimit
	return rb
}

func (rb *LimitsBuilder) TotalMlMemory(totalmlmemory string) *LimitsBuilder {
	rb.v.TotalMlMemory = totalmlmemory
	return rb
}
