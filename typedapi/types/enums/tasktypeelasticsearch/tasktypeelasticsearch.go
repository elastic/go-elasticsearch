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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

// Package tasktypeelasticsearch
package tasktypeelasticsearch

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/de4ff9ec1f716256f521d9e30011ad9c284b0dcc/specification/inference/_types/TaskType.ts#L68-L72
type TaskTypeElasticsearch struct {
	Name string
}

var (
	Sparseembedding = TaskTypeElasticsearch{"sparse_embedding"}

	Textembedding = TaskTypeElasticsearch{"text_embedding"}

	Rerank = TaskTypeElasticsearch{"rerank"}
)

func (t TaskTypeElasticsearch) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TaskTypeElasticsearch) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "sparse_embedding":
		*t = Sparseembedding
	case "text_embedding":
		*t = Textembedding
	case "rerank":
		*t = Rerank
	default:
		*t = TaskTypeElasticsearch{string(text)}
	}

	return nil
}

func (t TaskTypeElasticsearch) String() string {
	return t.Name
}
