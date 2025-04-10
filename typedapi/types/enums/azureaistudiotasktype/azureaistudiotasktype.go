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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

// Package azureaistudiotasktype
package azureaistudiotasktype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9/specification/inference/_types/CommonTypes.ts#L499-L502
type AzureAiStudioTaskType struct {
	Name string
}

var (
	Completion = AzureAiStudioTaskType{"completion"}

	Textembedding = AzureAiStudioTaskType{"text_embedding"}
)

func (a AzureAiStudioTaskType) MarshalText() (text []byte, err error) {
	return []byte(a.String()), nil
}

func (a *AzureAiStudioTaskType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "completion":
		*a = Completion
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
