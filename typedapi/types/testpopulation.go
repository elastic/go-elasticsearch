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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// TestPopulation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/metric.ts#L150-L154
type TestPopulation struct {
	Field  Field           `json:"field"`
	Filter *QueryContainer `json:"filter,omitempty"`
	Script *Script         `json:"script,omitempty"`
}

// TestPopulationBuilder holds TestPopulation struct and provides a builder API.
type TestPopulationBuilder struct {
	v *TestPopulation
}

// NewTestPopulation provides a builder for the TestPopulation struct.
func NewTestPopulationBuilder() *TestPopulationBuilder {
	r := TestPopulationBuilder{
		&TestPopulation{},
	}

	return &r
}

// Build finalize the chain and returns the TestPopulation struct
func (rb *TestPopulationBuilder) Build() TestPopulation {
	return *rb.v
}

func (rb *TestPopulationBuilder) Field(field Field) *TestPopulationBuilder {
	rb.v.Field = field
	return rb
}

func (rb *TestPopulationBuilder) Filter(filter *QueryContainerBuilder) *TestPopulationBuilder {
	v := filter.Build()
	rb.v.Filter = &v
	return rb
}

func (rb *TestPopulationBuilder) Script(script *ScriptBuilder) *TestPopulationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
