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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

// Package googlevertexaitasktype
package googlevertexaitasktype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9/specification/inference/_types/CommonTypes.ts#L860-L863
type GoogleVertexAITaskType struct {
	Name string
}

var (
	Rerank = GoogleVertexAITaskType{"rerank"}

	Textembedding = GoogleVertexAITaskType{"text_embedding"}
)

func (g GoogleVertexAITaskType) MarshalText() (text []byte, err error) {
	return []byte(g.String()), nil
}

func (g *GoogleVertexAITaskType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "rerank":
		*g = Rerank
	case "text_embedding":
		*g = Textembedding
	default:
		*g = GoogleVertexAITaskType{string(text)}
	}

	return nil
}

func (g GoogleVertexAITaskType) String() string {
	return g.Name
}
