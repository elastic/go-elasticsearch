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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// KeyValueProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ingest/_types/Processors.ts#L286-L298
type KeyValueProcessor struct {
	Description   *string              `json:"description,omitempty"`
	ExcludeKeys   []string             `json:"exclude_keys,omitempty"`
	Field         string               `json:"field"`
	FieldSplit    string               `json:"field_split"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	IncludeKeys   []string             `json:"include_keys,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Prefix        *string              `json:"prefix,omitempty"`
	StripBrackets *bool                `json:"strip_brackets,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   *string              `json:"target_field,omitempty"`
	TrimKey       *string              `json:"trim_key,omitempty"`
	TrimValue     *string              `json:"trim_value,omitempty"`
	ValueSplit    string               `json:"value_split"`
}

// NewKeyValueProcessor returns a KeyValueProcessor.
func NewKeyValueProcessor() *KeyValueProcessor {
	r := &KeyValueProcessor{}

	return r
}
