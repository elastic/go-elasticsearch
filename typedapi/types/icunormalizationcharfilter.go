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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icunormalizationmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icunormalizationtype"
)

// IcuNormalizationCharFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/icu-plugin.ts#L40-L44
type IcuNormalizationCharFilter struct {
	Mode    *icunormalizationmode.IcuNormalizationMode `json:"mode,omitempty"`
	Name    *icunormalizationtype.IcuNormalizationType `json:"name,omitempty"`
	Type    string                                     `json:"type,omitempty"`
	Version *VersionString                             `json:"version,omitempty"`
}

// IcuNormalizationCharFilterBuilder holds IcuNormalizationCharFilter struct and provides a builder API.
type IcuNormalizationCharFilterBuilder struct {
	v *IcuNormalizationCharFilter
}

// NewIcuNormalizationCharFilter provides a builder for the IcuNormalizationCharFilter struct.
func NewIcuNormalizationCharFilterBuilder() *IcuNormalizationCharFilterBuilder {
	r := IcuNormalizationCharFilterBuilder{
		&IcuNormalizationCharFilter{},
	}

	r.v.Type = "icu_normalizer"

	return &r
}

// Build finalize the chain and returns the IcuNormalizationCharFilter struct
func (rb *IcuNormalizationCharFilterBuilder) Build() IcuNormalizationCharFilter {
	return *rb.v
}

func (rb *IcuNormalizationCharFilterBuilder) Mode(mode icunormalizationmode.IcuNormalizationMode) *IcuNormalizationCharFilterBuilder {
	rb.v.Mode = &mode
	return rb
}

func (rb *IcuNormalizationCharFilterBuilder) Name(name icunormalizationtype.IcuNormalizationType) *IcuNormalizationCharFilterBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *IcuNormalizationCharFilterBuilder) Version(version VersionString) *IcuNormalizationCharFilterBuilder {
	rb.v.Version = &version
	return rb
}
