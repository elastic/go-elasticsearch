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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

// Package huggingfacetasktype
package huggingfacetasktype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/2514615770f18dbb4e3887cc1a279995dbfd0724/specification/inference/_types/CommonTypes.ts#L1763-L1768
type HuggingFaceTaskType struct {
	Name string
}

var (
	Chatcompletion = HuggingFaceTaskType{"chat_completion"}

	Completion = HuggingFaceTaskType{"completion"}

	Rerank = HuggingFaceTaskType{"rerank"}

	Textembedding = HuggingFaceTaskType{"text_embedding"}
)

func (h HuggingFaceTaskType) MarshalText() (text []byte, err error) {
	return []byte(h.String()), nil
}

func (h *HuggingFaceTaskType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "chat_completion":
		*h = Chatcompletion
	case "completion":
		*h = Completion
	case "rerank":
		*h = Rerank
	case "text_embedding":
		*h = Textembedding
	default:
		*h = HuggingFaceTaskType{string(text)}
	}

	return nil
}

func (h HuggingFaceTaskType) String() string {
	return h.Name
}
