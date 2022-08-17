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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icunormalizationtype"
)

// IcuNormalizationTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/icu-plugin.ts#L35-L38
type IcuNormalizationTokenFilter struct {
	Name    icunormalizationtype.IcuNormalizationType `json:"name"`
	Type    string                                    `json:"type,omitempty"`
	Version *VersionString                            `json:"version,omitempty"`
}

// IcuNormalizationTokenFilterBuilder holds IcuNormalizationTokenFilter struct and provides a builder API.
type IcuNormalizationTokenFilterBuilder struct {
	v *IcuNormalizationTokenFilter
}

// NewIcuNormalizationTokenFilter provides a builder for the IcuNormalizationTokenFilter struct.
func NewIcuNormalizationTokenFilterBuilder() *IcuNormalizationTokenFilterBuilder {
	r := IcuNormalizationTokenFilterBuilder{
		&IcuNormalizationTokenFilter{},
	}

	r.v.Type = "icu_normalizer"

	return &r
}

// Build finalize the chain and returns the IcuNormalizationTokenFilter struct
func (rb *IcuNormalizationTokenFilterBuilder) Build() IcuNormalizationTokenFilter {
	return *rb.v
}

func (rb *IcuNormalizationTokenFilterBuilder) Name(name icunormalizationtype.IcuNormalizationType) *IcuNormalizationTokenFilterBuilder {
	rb.v.Name = name
	return rb
}

func (rb *IcuNormalizationTokenFilterBuilder) Version(version VersionString) *IcuNormalizationTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
