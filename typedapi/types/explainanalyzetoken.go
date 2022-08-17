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
	"encoding/json"
)

// ExplainAnalyzeToken type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/analyze/types.ts#L52-L64
type ExplainAnalyzeToken struct {
	Bytes               string                 `json:"bytes"`
	EndOffset           int64                  `json:"end_offset"`
	ExplainAnalyzeToken map[string]interface{} `json:"-"`
	Keyword             *bool                  `json:"keyword,omitempty"`
	Position            int64                  `json:"position"`
	PositionLength      int64                  `json:"positionLength"`
	StartOffset         int64                  `json:"start_offset"`
	TermFrequency       int64                  `json:"termFrequency"`
	Token               string                 `json:"token"`
	Type                string                 `json:"type"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s ExplainAnalyzeToken) MarshalJSON() ([]byte, error) {
	type opt ExplainAnalyzeToken
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.ExplainAnalyzeToken {
		tmp[string(key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ExplainAnalyzeTokenBuilder holds ExplainAnalyzeToken struct and provides a builder API.
type ExplainAnalyzeTokenBuilder struct {
	v *ExplainAnalyzeToken
}

// NewExplainAnalyzeToken provides a builder for the ExplainAnalyzeToken struct.
func NewExplainAnalyzeTokenBuilder() *ExplainAnalyzeTokenBuilder {
	r := ExplainAnalyzeTokenBuilder{
		&ExplainAnalyzeToken{
			ExplainAnalyzeToken: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ExplainAnalyzeToken struct
func (rb *ExplainAnalyzeTokenBuilder) Build() ExplainAnalyzeToken {
	return *rb.v
}

func (rb *ExplainAnalyzeTokenBuilder) Bytes(bytes string) *ExplainAnalyzeTokenBuilder {
	rb.v.Bytes = bytes
	return rb
}

func (rb *ExplainAnalyzeTokenBuilder) EndOffset(endoffset int64) *ExplainAnalyzeTokenBuilder {
	rb.v.EndOffset = endoffset
	return rb
}

func (rb *ExplainAnalyzeTokenBuilder) ExplainAnalyzeToken(value map[string]interface{}) *ExplainAnalyzeTokenBuilder {
	rb.v.ExplainAnalyzeToken = value
	return rb
}

func (rb *ExplainAnalyzeTokenBuilder) Keyword(keyword bool) *ExplainAnalyzeTokenBuilder {
	rb.v.Keyword = &keyword
	return rb
}

func (rb *ExplainAnalyzeTokenBuilder) Position(position int64) *ExplainAnalyzeTokenBuilder {
	rb.v.Position = position
	return rb
}

func (rb *ExplainAnalyzeTokenBuilder) PositionLength(positionlength int64) *ExplainAnalyzeTokenBuilder {
	rb.v.PositionLength = positionlength
	return rb
}

func (rb *ExplainAnalyzeTokenBuilder) StartOffset(startoffset int64) *ExplainAnalyzeTokenBuilder {
	rb.v.StartOffset = startoffset
	return rb
}

func (rb *ExplainAnalyzeTokenBuilder) TermFrequency(termfrequency int64) *ExplainAnalyzeTokenBuilder {
	rb.v.TermFrequency = termfrequency
	return rb
}

func (rb *ExplainAnalyzeTokenBuilder) Token(token string) *ExplainAnalyzeTokenBuilder {
	rb.v.Token = token
	return rb
}

func (rb *ExplainAnalyzeTokenBuilder) Type_(type_ string) *ExplainAnalyzeTokenBuilder {
	rb.v.Type = type_
	return rb
}
