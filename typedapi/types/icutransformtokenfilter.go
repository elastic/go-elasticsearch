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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icutransformdirection"
)

// IcuTransformTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/icu-plugin.ts#L24-L28
type IcuTransformTokenFilter struct {
	Dir     icutransformdirection.IcuTransformDirection `json:"dir"`
	Id      string                                      `json:"id"`
	Type    string                                      `json:"type,omitempty"`
	Version *VersionString                              `json:"version,omitempty"`
}

// IcuTransformTokenFilterBuilder holds IcuTransformTokenFilter struct and provides a builder API.
type IcuTransformTokenFilterBuilder struct {
	v *IcuTransformTokenFilter
}

// NewIcuTransformTokenFilter provides a builder for the IcuTransformTokenFilter struct.
func NewIcuTransformTokenFilterBuilder() *IcuTransformTokenFilterBuilder {
	r := IcuTransformTokenFilterBuilder{
		&IcuTransformTokenFilter{},
	}

	r.v.Type = "icu_transform"

	return &r
}

// Build finalize the chain and returns the IcuTransformTokenFilter struct
func (rb *IcuTransformTokenFilterBuilder) Build() IcuTransformTokenFilter {
	return *rb.v
}

func (rb *IcuTransformTokenFilterBuilder) Dir(dir icutransformdirection.IcuTransformDirection) *IcuTransformTokenFilterBuilder {
	rb.v.Dir = dir
	return rb
}

func (rb *IcuTransformTokenFilterBuilder) Id(id string) *IcuTransformTokenFilterBuilder {
	rb.v.Id = id
	return rb
}

func (rb *IcuTransformTokenFilterBuilder) Version(version VersionString) *IcuTransformTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
