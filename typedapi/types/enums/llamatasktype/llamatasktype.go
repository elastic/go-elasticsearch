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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

// Package llamatasktype
package llamatasktype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27/specification/inference/_types/CommonTypes.ts#L1758-L1762
type LlamaTaskType struct {
	Name string
}

var (
	Textembedding = LlamaTaskType{"text_embedding"}

	Completion = LlamaTaskType{"completion"}

	Chatcompletion = LlamaTaskType{"chat_completion"}
)

func (l LlamaTaskType) MarshalText() (text []byte, err error) {
	return []byte(l.String()), nil
}

func (l *LlamaTaskType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "text_embedding":
		*l = Textembedding
	case "completion":
		*l = Completion
	case "chat_completion":
		*l = Chatcompletion
	default:
		*l = LlamaTaskType{string(text)}
	}

	return nil
}

func (l LlamaTaskType) String() string {
	return l.Name
}
