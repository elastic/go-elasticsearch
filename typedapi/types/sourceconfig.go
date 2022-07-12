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

// SourceConfig holds the union for the following types:
//     bool
//     SourceFilter
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_global/search/_types/SourceFilter.ts#L33-L37
type SourceConfig interface{}

// SourceConfigBuilder holds SourceConfig struct and provides a builder API.
type SourceConfigBuilder struct {
	v SourceConfig
}

// NewSourceConfig provides a builder for the SourceConfig struct.
func NewSourceConfigBuilder() *SourceConfigBuilder {
	return &SourceConfigBuilder{}
}

// Build finalize the chain and returns the SourceConfig struct
func (u *SourceConfigBuilder) Build() SourceConfig {
	return u.v
}

func (u *SourceConfigBuilder) Bool(bool bool) *SourceConfigBuilder {
	u.v = &bool
	return u
}

func (u *SourceConfigBuilder) SourceFilter(sourcefilter *SourceFilterBuilder) *SourceConfigBuilder {
	v := sourcefilter.Build()
	u.v = &v
	return u
}
