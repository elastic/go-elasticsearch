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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/snowballlanguage"
)

// SnowballTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L308-L311
type SnowballTokenFilter struct {
	Language snowballlanguage.SnowballLanguage `json:"language"`
	Type     string                            `json:"type,omitempty"`
	Version  *VersionString                    `json:"version,omitempty"`
}

// SnowballTokenFilterBuilder holds SnowballTokenFilter struct and provides a builder API.
type SnowballTokenFilterBuilder struct {
	v *SnowballTokenFilter
}

// NewSnowballTokenFilter provides a builder for the SnowballTokenFilter struct.
func NewSnowballTokenFilterBuilder() *SnowballTokenFilterBuilder {
	r := SnowballTokenFilterBuilder{
		&SnowballTokenFilter{},
	}

	r.v.Type = "snowball"

	return &r
}

// Build finalize the chain and returns the SnowballTokenFilter struct
func (rb *SnowballTokenFilterBuilder) Build() SnowballTokenFilter {
	return *rb.v
}

func (rb *SnowballTokenFilterBuilder) Language(language snowballlanguage.SnowballLanguage) *SnowballTokenFilterBuilder {
	rb.v.Language = language
	return rb
}

func (rb *SnowballTokenFilterBuilder) Version(version VersionString) *SnowballTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
