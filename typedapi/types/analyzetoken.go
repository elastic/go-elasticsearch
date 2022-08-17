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

// AnalyzeToken type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/analyze/types.ts#L37-L44
type AnalyzeToken struct {
	EndOffset      int64  `json:"end_offset"`
	Position       int64  `json:"position"`
	PositionLength *int64 `json:"position_length,omitempty"`
	StartOffset    int64  `json:"start_offset"`
	Token          string `json:"token"`
	Type           string `json:"type"`
}

// AnalyzeTokenBuilder holds AnalyzeToken struct and provides a builder API.
type AnalyzeTokenBuilder struct {
	v *AnalyzeToken
}

// NewAnalyzeToken provides a builder for the AnalyzeToken struct.
func NewAnalyzeTokenBuilder() *AnalyzeTokenBuilder {
	r := AnalyzeTokenBuilder{
		&AnalyzeToken{},
	}

	return &r
}

// Build finalize the chain and returns the AnalyzeToken struct
func (rb *AnalyzeTokenBuilder) Build() AnalyzeToken {
	return *rb.v
}

func (rb *AnalyzeTokenBuilder) EndOffset(endoffset int64) *AnalyzeTokenBuilder {
	rb.v.EndOffset = endoffset
	return rb
}

func (rb *AnalyzeTokenBuilder) Position(position int64) *AnalyzeTokenBuilder {
	rb.v.Position = position
	return rb
}

func (rb *AnalyzeTokenBuilder) PositionLength(positionlength int64) *AnalyzeTokenBuilder {
	rb.v.PositionLength = &positionlength
	return rb
}

func (rb *AnalyzeTokenBuilder) StartOffset(startoffset int64) *AnalyzeTokenBuilder {
	rb.v.StartOffset = startoffset
	return rb
}

func (rb *AnalyzeTokenBuilder) Token(token string) *AnalyzeTokenBuilder {
	rb.v.Token = token
	return rb
}

func (rb *AnalyzeTokenBuilder) Type_(type_ string) *AnalyzeTokenBuilder {
	rb.v.Type = type_
	return rb
}
