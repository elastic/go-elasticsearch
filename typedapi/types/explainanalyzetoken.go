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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

import (
	"encoding/json"
	"fmt"
)

// ExplainAnalyzeToken type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/indices/analyze/types.ts#L52-L64
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
		tmp[fmt.Sprintf("%s", key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewExplainAnalyzeToken returns a ExplainAnalyzeToken.
func NewExplainAnalyzeToken() *ExplainAnalyzeToken {
	r := &ExplainAnalyzeToken{
		ExplainAnalyzeToken: make(map[string]interface{}, 0),
	}

	return r
}
