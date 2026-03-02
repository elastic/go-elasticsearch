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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

// Package jinaaitasktype
package jinaaitasktype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/e196f9953fa743572ee46884835f1934bce9a16b/specification/inference/_types/CommonTypes.ts#L1658-L1661
type JinaAITaskType struct {
	Name string
}

var (
	Rerank = JinaAITaskType{"rerank"}

	Textembedding = JinaAITaskType{"text_embedding"}
)

func (j JinaAITaskType) MarshalText() (text []byte, err error) {
	return []byte(j.String()), nil
}

func (j *JinaAITaskType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "rerank":
		*j = Rerank
	case "text_embedding":
		*j = Textembedding
	default:
		*j = JinaAITaskType{string(text)}
	}

	return nil
}

func (j JinaAITaskType) String() string {
	return j.Name
}
