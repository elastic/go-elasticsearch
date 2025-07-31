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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

// Package googleaistudiotasktype
package googleaistudiotasktype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/86f41834c7bb975159a38a73be8a9d930010d673/specification/inference/_types/CommonTypes.ts#L885-L888
type GoogleAiStudioTaskType struct {
	Name string
}

var (
	Completion = GoogleAiStudioTaskType{"completion"}

	Textembedding = GoogleAiStudioTaskType{"text_embedding"}
)

func (g GoogleAiStudioTaskType) MarshalText() (text []byte, err error) {
	return []byte(g.String()), nil
}

func (g *GoogleAiStudioTaskType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "completion":
		*g = Completion
	case "text_embedding":
		*g = Textembedding
	default:
		*g = GoogleAiStudioTaskType{string(text)}
	}

	return nil
}

func (g GoogleAiStudioTaskType) String() string {
	return g.Name
}
