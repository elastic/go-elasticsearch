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

// SetProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ingest/_types/Processors.ts#L329-L336
type SetProcessor struct {
	CopyFrom         *string              `json:"copy_from,omitempty"`
	Description      *string              `json:"description,omitempty"`
	Field            string               `json:"field"`
	If               *string              `json:"if,omitempty"`
	IgnoreEmptyValue *bool                `json:"ignore_empty_value,omitempty"`
	IgnoreFailure    *bool                `json:"ignore_failure,omitempty"`
	MediaType        *string              `json:"media_type,omitempty"`
	OnFailure        []ProcessorContainer `json:"on_failure,omitempty"`
	Override         *bool                `json:"override,omitempty"`
	Tag              *string              `json:"tag,omitempty"`
	Value            interface{}          `json:"value,omitempty"`
}

// NewSetProcessor returns a SetProcessor.
func NewSetProcessor() *SetProcessor {
	r := &SetProcessor{}

	return r
}
