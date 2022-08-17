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

// MigrationFeatureIndexInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/migration/get_feature_upgrade_status/GetFeatureUpgradeStatusResponse.ts#L44-L48
type MigrationFeatureIndexInfo struct {
	FailureCause *ErrorCause   `json:"failure_cause,omitempty"`
	Index        IndexName     `json:"index"`
	Version      VersionString `json:"version"`
}

// MigrationFeatureIndexInfoBuilder holds MigrationFeatureIndexInfo struct and provides a builder API.
type MigrationFeatureIndexInfoBuilder struct {
	v *MigrationFeatureIndexInfo
}

// NewMigrationFeatureIndexInfo provides a builder for the MigrationFeatureIndexInfo struct.
func NewMigrationFeatureIndexInfoBuilder() *MigrationFeatureIndexInfoBuilder {
	r := MigrationFeatureIndexInfoBuilder{
		&MigrationFeatureIndexInfo{},
	}

	return &r
}

// Build finalize the chain and returns the MigrationFeatureIndexInfo struct
func (rb *MigrationFeatureIndexInfoBuilder) Build() MigrationFeatureIndexInfo {
	return *rb.v
}

func (rb *MigrationFeatureIndexInfoBuilder) FailureCause(failurecause *ErrorCauseBuilder) *MigrationFeatureIndexInfoBuilder {
	v := failurecause.Build()
	rb.v.FailureCause = &v
	return rb
}

func (rb *MigrationFeatureIndexInfoBuilder) Index(index IndexName) *MigrationFeatureIndexInfoBuilder {
	rb.v.Index = index
	return rb
}

func (rb *MigrationFeatureIndexInfoBuilder) Version(version VersionString) *MigrationFeatureIndexInfoBuilder {
	rb.v.Version = version
	return rb
}
