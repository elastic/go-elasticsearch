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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"encoding/json"
)

// CsvProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ingest/_types/Processors.ts#L154-L162
type CsvProcessor struct {
	Description   *string              `json:"description,omitempty"`
	EmptyValue    json.RawMessage      `json:"empty_value,omitempty"`
	Field         string               `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Quote         *string              `json:"quote,omitempty"`
	Separator     *string              `json:"separator,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	TargetFields  []string             `json:"target_fields"`
	Trim          *bool                `json:"trim,omitempty"`
}

// NewCsvProcessor returns a CsvProcessor.
func NewCsvProcessor() *CsvProcessor {
	r := &CsvProcessor{}

	return r
}
