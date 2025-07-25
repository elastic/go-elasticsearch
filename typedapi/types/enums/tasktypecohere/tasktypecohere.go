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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

// Package tasktypecohere
package tasktypecohere

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0b0db20330063a6d11f7997ff443fd2a1a827d1/specification/inference/_types/TaskType.ts#L62-L66
type TaskTypeCohere struct {
	Name string
}

var (
	Textembedding = TaskTypeCohere{"text_embedding"}

	Rerank = TaskTypeCohere{"rerank"}

	Completion = TaskTypeCohere{"completion"}
)

func (t TaskTypeCohere) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TaskTypeCohere) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "text_embedding":
		*t = Textembedding
	case "rerank":
		*t = Rerank
	case "completion":
		*t = Completion
	default:
		*t = TaskTypeCohere{string(text)}
	}

	return nil
}

func (t TaskTypeCohere) String() string {
	return t.Name
}
