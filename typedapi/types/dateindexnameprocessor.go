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

// DateIndexNameProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ingest/_types/Processors.ts#L164-L177
type DateIndexNameProcessor struct {
	DateFormats []string `json:"date_formats"`
	// DateRounding How to round the date when formatting the date into the index name. Valid
	// values are:
	// `y` (year), `M` (month), `w` (week), `d` (day), `h` (hour), `m` (minute) and
	// `s` (second).
	// Supports template snippets.
	DateRounding    string               `json:"date_rounding"`
	Description     *string              `json:"description,omitempty"`
	Field           string               `json:"field"`
	If              *string              `json:"if,omitempty"`
	IgnoreFailure   *bool                `json:"ignore_failure,omitempty"`
	IndexNameFormat *string              `json:"index_name_format,omitempty"`
	IndexNamePrefix *string              `json:"index_name_prefix,omitempty"`
	Locale          *string              `json:"locale,omitempty"`
	OnFailure       []ProcessorContainer `json:"on_failure,omitempty"`
	Tag             *string              `json:"tag,omitempty"`
	Timezone        *string              `json:"timezone,omitempty"`
}

// NewDateIndexNameProcessor returns a DateIndexNameProcessor.
func NewDateIndexNameProcessor() *DateIndexNameProcessor {
	r := &DateIndexNameProcessor{}

	return r
}
