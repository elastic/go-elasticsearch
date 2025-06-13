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
// https://github.com/elastic/elasticsearch-specification/tree/3a94b6715915b1e9311724a2614c643368eece90

// Package mistraltasktype
package mistraltasktype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/3a94b6715915b1e9311724a2614c643368eece90/specification/inference/_types/CommonTypes.ts#L1078-L1080
type MistralTaskType struct {
	Name string
}

var (
	Textembedding = MistralTaskType{"text_embedding"}
)

func (m MistralTaskType) MarshalText() (text []byte, err error) {
	return []byte(m.String()), nil
}

func (m *MistralTaskType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "text_embedding":
		*m = Textembedding
	default:
		*m = MistralTaskType{string(text)}
	}

	return nil
}

func (m MistralTaskType) String() string {
	return m.Name
}
