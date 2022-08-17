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

// LifecycleExplainPhaseExecution type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ilm/explain_lifecycle/types.ts#L64-L68
type LifecycleExplainPhaseExecution struct {
	ModifiedDateInMillis EpochTimeUnitMillis `json:"modified_date_in_millis"`
	Policy               Name                `json:"policy"`
	Version              VersionNumber       `json:"version"`
}

// LifecycleExplainPhaseExecutionBuilder holds LifecycleExplainPhaseExecution struct and provides a builder API.
type LifecycleExplainPhaseExecutionBuilder struct {
	v *LifecycleExplainPhaseExecution
}

// NewLifecycleExplainPhaseExecution provides a builder for the LifecycleExplainPhaseExecution struct.
func NewLifecycleExplainPhaseExecutionBuilder() *LifecycleExplainPhaseExecutionBuilder {
	r := LifecycleExplainPhaseExecutionBuilder{
		&LifecycleExplainPhaseExecution{},
	}

	return &r
}

// Build finalize the chain and returns the LifecycleExplainPhaseExecution struct
func (rb *LifecycleExplainPhaseExecutionBuilder) Build() LifecycleExplainPhaseExecution {
	return *rb.v
}

func (rb *LifecycleExplainPhaseExecutionBuilder) ModifiedDateInMillis(modifieddateinmillis *EpochTimeUnitMillisBuilder) *LifecycleExplainPhaseExecutionBuilder {
	v := modifieddateinmillis.Build()
	rb.v.ModifiedDateInMillis = v
	return rb
}

func (rb *LifecycleExplainPhaseExecutionBuilder) Policy(policy Name) *LifecycleExplainPhaseExecutionBuilder {
	rb.v.Policy = policy
	return rb
}

func (rb *LifecycleExplainPhaseExecutionBuilder) Version(version VersionNumber) *LifecycleExplainPhaseExecutionBuilder {
	rb.v.Version = version
	return rb
}
