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
// https://github.com/elastic/elasticsearch-specification/tree/f1932ce6b46a53a8342db522b1a7883bcc9e0996

// Package azureopenaitasktype
package azureopenaitasktype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/f1932ce6b46a53a8342db522b1a7883bcc9e0996/specification/inference/_types/CommonTypes.ts#L563-L566
type AzureOpenAITaskType struct {
	Name string
}

var (
	Completion = AzureOpenAITaskType{"completion"}

	Textembedding = AzureOpenAITaskType{"text_embedding"}
)

func (a AzureOpenAITaskType) MarshalText() (text []byte, err error) {
	return []byte(a.String()), nil
}

func (a *AzureOpenAITaskType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "completion":
		*a = Completion
	case "text_embedding":
		*a = Textembedding
	default:
		*a = AzureOpenAITaskType{string(text)}
	}

	return nil
}

func (a AzureOpenAITaskType) String() string {
	return a.Name
}
