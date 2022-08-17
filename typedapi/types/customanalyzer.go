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

// CustomAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/analyzers.ts#L28-L35
type CustomAnalyzer struct {
	CharFilter           []string `json:"char_filter,omitempty"`
	Filter               []string `json:"filter,omitempty"`
	PositionIncrementGap *int     `json:"position_increment_gap,omitempty"`
	PositionOffsetGap    *int     `json:"position_offset_gap,omitempty"`
	Tokenizer            string   `json:"tokenizer"`
	Type                 string   `json:"type,omitempty"`
}

// CustomAnalyzerBuilder holds CustomAnalyzer struct and provides a builder API.
type CustomAnalyzerBuilder struct {
	v *CustomAnalyzer
}

// NewCustomAnalyzer provides a builder for the CustomAnalyzer struct.
func NewCustomAnalyzerBuilder() *CustomAnalyzerBuilder {
	r := CustomAnalyzerBuilder{
		&CustomAnalyzer{},
	}

	r.v.Type = "custom"

	return &r
}

// Build finalize the chain and returns the CustomAnalyzer struct
func (rb *CustomAnalyzerBuilder) Build() CustomAnalyzer {
	return *rb.v
}

func (rb *CustomAnalyzerBuilder) CharFilter(char_filter ...string) *CustomAnalyzerBuilder {
	rb.v.CharFilter = char_filter
	return rb
}

func (rb *CustomAnalyzerBuilder) Filter(filter ...string) *CustomAnalyzerBuilder {
	rb.v.Filter = filter
	return rb
}

func (rb *CustomAnalyzerBuilder) PositionIncrementGap(positionincrementgap int) *CustomAnalyzerBuilder {
	rb.v.PositionIncrementGap = &positionincrementgap
	return rb
}

func (rb *CustomAnalyzerBuilder) PositionOffsetGap(positionoffsetgap int) *CustomAnalyzerBuilder {
	rb.v.PositionOffsetGap = &positionoffsetgap
	return rb
}

func (rb *CustomAnalyzerBuilder) Tokenizer(tokenizer string) *CustomAnalyzerBuilder {
	rb.v.Tokenizer = tokenizer
	return rb
}
