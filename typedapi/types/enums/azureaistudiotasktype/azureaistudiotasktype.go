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
// https://github.com/elastic/elasticsearch-specification/tree/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f

// Package azureaistudiotasktype
package azureaistudiotasktype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f/specification/inference/_types/CommonTypes.ts#L581-L585
type AzureAiStudioTaskType struct {
	Name string
}

var (
	Completion = AzureAiStudioTaskType{"completion"}

	Rerank = AzureAiStudioTaskType{"rerank"}

	Textembedding = AzureAiStudioTaskType{"text_embedding"}
)

func (a AzureAiStudioTaskType) MarshalText() (text []byte, err error) {
	return []byte(a.String()), nil
}

func (a *AzureAiStudioTaskType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "completion":
		*a = Completion
	case "rerank":
		*a = Rerank
	case "text_embedding":
		*a = Textembedding
	default:
		*a = AzureAiStudioTaskType{string(text)}
	}

	return nil
}

func (a AzureAiStudioTaskType) String() string {
	return a.Name
}
