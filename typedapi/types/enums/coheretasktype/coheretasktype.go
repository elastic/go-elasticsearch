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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

// Package coheretasktype
package coheretasktype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/d520d9e8cf14cad487de5e0654007686c395b494/specification/inference/_types/CommonTypes.ts#L877-L881
type CohereTaskType struct {
	Name string
}

var (
	Completion = CohereTaskType{"completion"}

	Rerank = CohereTaskType{"rerank"}

	Textembedding = CohereTaskType{"text_embedding"}
)

func (c CohereTaskType) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CohereTaskType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "completion":
		*c = Completion
	case "rerank":
		*c = Rerank
	case "text_embedding":
		*c = Textembedding
	default:
		*c = CohereTaskType{string(text)}
	}

	return nil
}

func (c CohereTaskType) String() string {
	return c.Name
}
