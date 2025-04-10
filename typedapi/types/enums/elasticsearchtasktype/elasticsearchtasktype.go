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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

// Package elasticsearchtasktype
package elasticsearchtasktype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/beeb1dc688bcc058488dcc45d9cbd2cd364e9943/specification/inference/_types/CommonTypes.ts#L748-L752
type ElasticsearchTaskType struct {
	Name string
}

var (
	Rerank = ElasticsearchTaskType{"rerank"}

	Sparseembedding = ElasticsearchTaskType{"sparse_embedding"}

	Textembedding = ElasticsearchTaskType{"text_embedding"}
)

func (e ElasticsearchTaskType) MarshalText() (text []byte, err error) {
	return []byte(e.String()), nil
}

func (e *ElasticsearchTaskType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "rerank":
		*e = Rerank
	case "sparse_embedding":
		*e = Sparseembedding
	case "text_embedding":
		*e = Textembedding
	default:
		*e = ElasticsearchTaskType{string(text)}
	}

	return nil
}

func (e ElasticsearchTaskType) String() string {
	return e.Name
}
