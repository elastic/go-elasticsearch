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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Package watsonxtasktype
package watsonxtasktype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/inference/_types/CommonTypes.ts#L1724-L1728
type WatsonxTaskType struct {
	Name string
}

var (
	Textembedding = WatsonxTaskType{"text_embedding"}

	Chatcompletion = WatsonxTaskType{"chat_completion"}

	Completion = WatsonxTaskType{"completion"}
)

func (w WatsonxTaskType) MarshalText() (text []byte, err error) {
	return []byte(w.String()), nil
}

func (w *WatsonxTaskType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "text_embedding":
		*w = Textembedding
	case "chat_completion":
		*w = Chatcompletion
	case "completion":
		*w = Completion
	default:
		*w = WatsonxTaskType{string(text)}
	}

	return nil
}

func (w WatsonxTaskType) String() string {
	return w.Name
}
