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
// https://github.com/elastic/elasticsearch-specification/tree/93ed2b29c9e75f49cd340f06286d6ead5965f900


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icunormalizationmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icunormalizationtype"
)

// IcuAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/analysis/icu-plugin.ts#L67-L71
type IcuAnalyzer struct {
	Method icunormalizationtype.IcuNormalizationType `json:"method"`
	Mode   icunormalizationmode.IcuNormalizationMode `json:"mode"`
	Type   string                                    `json:"type,omitempty"`
}

// IcuAnalyzerBuilder holds IcuAnalyzer struct and provides a builder API.
type IcuAnalyzerBuilder struct {
	v *IcuAnalyzer
}

// NewIcuAnalyzer provides a builder for the IcuAnalyzer struct.
func NewIcuAnalyzerBuilder() *IcuAnalyzerBuilder {
	r := IcuAnalyzerBuilder{
		&IcuAnalyzer{},
	}

	r.v.Type = "icu_analyzer"

	return &r
}

// Build finalize the chain and returns the IcuAnalyzer struct
func (rb *IcuAnalyzerBuilder) Build() IcuAnalyzer {
	return *rb.v
}

func (rb *IcuAnalyzerBuilder) Method(method icunormalizationtype.IcuNormalizationType) *IcuAnalyzerBuilder {
	rb.v.Method = method
	return rb
}

func (rb *IcuAnalyzerBuilder) Mode(mode icunormalizationmode.IcuNormalizationMode) *IcuAnalyzerBuilder {
	rb.v.Mode = mode
	return rb
}
