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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

// Package tasktypegooglevertexai
package tasktypegooglevertexai

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/6785a6caa1fa3ca5ab3308963d79dce923a3469f/specification/inference/_types/TaskType.ts#L114-L119
type TaskTypeGoogleVertexAI struct {
	Name string
}

var (
	Chatcompletion = TaskTypeGoogleVertexAI{"chat_completion"}

	Completion = TaskTypeGoogleVertexAI{"completion"}

	Textembedding = TaskTypeGoogleVertexAI{"text_embedding"}

	Rerank = TaskTypeGoogleVertexAI{"rerank"}
)

func (t TaskTypeGoogleVertexAI) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TaskTypeGoogleVertexAI) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "chat_completion":
		*t = Chatcompletion
	case "completion":
		*t = Completion
	case "text_embedding":
		*t = Textembedding
	case "rerank":
		*t = Rerank
	default:
		*t = TaskTypeGoogleVertexAI{string(text)}
	}

	return nil
}

func (t TaskTypeGoogleVertexAI) String() string {
	return t.Name
}
