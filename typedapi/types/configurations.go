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

// Configurations type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ilm/_types/Phase.ts#L46-L50
type Configurations struct {
	Forcemerge *ForceMergeConfiguration `json:"forcemerge,omitempty"`
	Rollover   *RolloverConditions      `json:"rollover,omitempty"`
	Shrink     *ShrinkConfiguration     `json:"shrink,omitempty"`
}

// ConfigurationsBuilder holds Configurations struct and provides a builder API.
type ConfigurationsBuilder struct {
	v *Configurations
}

// NewConfigurations provides a builder for the Configurations struct.
func NewConfigurationsBuilder() *ConfigurationsBuilder {
	r := ConfigurationsBuilder{
		&Configurations{},
	}

	return &r
}

// Build finalize the chain and returns the Configurations struct
func (rb *ConfigurationsBuilder) Build() Configurations {
	return *rb.v
}

func (rb *ConfigurationsBuilder) Forcemerge(forcemerge *ForceMergeConfigurationBuilder) *ConfigurationsBuilder {
	v := forcemerge.Build()
	rb.v.Forcemerge = &v
	return rb
}

func (rb *ConfigurationsBuilder) Rollover(rollover *RolloverConditionsBuilder) *ConfigurationsBuilder {
	v := rollover.Build()
	rb.v.Rollover = &v
	return rb
}

func (rb *ConfigurationsBuilder) Shrink(shrink *ShrinkConfigurationBuilder) *ConfigurationsBuilder {
	v := shrink.Build()
	rb.v.Shrink = &v
	return rb
}
