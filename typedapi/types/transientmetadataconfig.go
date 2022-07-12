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

// TransientMetadataConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/security/_types/TransientMetadataConfig.ts#L20-L22
type TransientMetadataConfig struct {
	Enabled bool `json:"enabled"`
}

// TransientMetadataConfigBuilder holds TransientMetadataConfig struct and provides a builder API.
type TransientMetadataConfigBuilder struct {
	v *TransientMetadataConfig
}

// NewTransientMetadataConfig provides a builder for the TransientMetadataConfig struct.
func NewTransientMetadataConfigBuilder() *TransientMetadataConfigBuilder {
	r := TransientMetadataConfigBuilder{
		&TransientMetadataConfig{},
	}

	return &r
}

// Build finalize the chain and returns the TransientMetadataConfig struct
func (rb *TransientMetadataConfigBuilder) Build() TransientMetadataConfig {
	return *rb.v
}

func (rb *TransientMetadataConfigBuilder) Enabled(enabled bool) *TransientMetadataConfigBuilder {
	rb.v.Enabled = enabled
	return rb
}
