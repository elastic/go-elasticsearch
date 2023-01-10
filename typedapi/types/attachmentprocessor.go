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

// AttachmentProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ingest/_types/Processors.ts#L96-L104
type AttachmentProcessor struct {
	Description       *string              `json:"description,omitempty"`
	Field             string               `json:"field"`
	If                *string              `json:"if,omitempty"`
	IgnoreFailure     *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing     *bool                `json:"ignore_missing,omitempty"`
	IndexedChars      *int64               `json:"indexed_chars,omitempty"`
	IndexedCharsField *string              `json:"indexed_chars_field,omitempty"`
	OnFailure         []ProcessorContainer `json:"on_failure,omitempty"`
	Properties        []string             `json:"properties,omitempty"`
	ResourceName      *string              `json:"resource_name,omitempty"`
	Tag               *string              `json:"tag,omitempty"`
	TargetField       *string              `json:"target_field,omitempty"`
}

// NewAttachmentProcessor returns a AttachmentProcessor.
func NewAttachmentProcessor() *AttachmentProcessor {
	r := &AttachmentProcessor{}

	return r
}
