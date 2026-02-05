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

// Package tasktypenvidia
package tasktypenvidia

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/2514615770f18dbb4e3887cc1a279995dbfd0724/specification/inference/_types/TaskType.ts#L147-L152
type TaskTypeNvidia struct {
	Name string
}

var (
	Chatcompletion = TaskTypeNvidia{"chat_completion"}

	Completion = TaskTypeNvidia{"completion"}

	Rerank = TaskTypeNvidia{"rerank"}

	Textembedding = TaskTypeNvidia{"text_embedding"}
)

func (t TaskTypeNvidia) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TaskTypeNvidia) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "chat_completion":
		*t = Chatcompletion
	case "completion":
		*t = Completion
	case "rerank":
		*t = Rerank
	case "text_embedding":
		*t = Textembedding
	default:
		*t = TaskTypeNvidia{string(text)}
	}

	return nil
}

func (t TaskTypeNvidia) String() string {
	return t.Name
}
