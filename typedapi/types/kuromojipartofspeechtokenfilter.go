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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// KuromojiPartOfSpeechTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/kuromoji-plugin.ts#L37-L40
type KuromojiPartOfSpeechTokenFilter struct {
	Stoptags []string       `json:"stoptags"`
	Type     string         `json:"type,omitempty"`
	Version  *VersionString `json:"version,omitempty"`
}

// KuromojiPartOfSpeechTokenFilterBuilder holds KuromojiPartOfSpeechTokenFilter struct and provides a builder API.
type KuromojiPartOfSpeechTokenFilterBuilder struct {
	v *KuromojiPartOfSpeechTokenFilter
}

// NewKuromojiPartOfSpeechTokenFilter provides a builder for the KuromojiPartOfSpeechTokenFilter struct.
func NewKuromojiPartOfSpeechTokenFilterBuilder() *KuromojiPartOfSpeechTokenFilterBuilder {
	r := KuromojiPartOfSpeechTokenFilterBuilder{
		&KuromojiPartOfSpeechTokenFilter{},
	}

	r.v.Type = "kuromoji_part_of_speech"

	return &r
}

// Build finalize the chain and returns the KuromojiPartOfSpeechTokenFilter struct
func (rb *KuromojiPartOfSpeechTokenFilterBuilder) Build() KuromojiPartOfSpeechTokenFilter {
	return *rb.v
}

func (rb *KuromojiPartOfSpeechTokenFilterBuilder) Stoptags(stoptags ...string) *KuromojiPartOfSpeechTokenFilterBuilder {
	rb.v.Stoptags = stoptags
	return rb
}

func (rb *KuromojiPartOfSpeechTokenFilterBuilder) Version(version VersionString) *KuromojiPartOfSpeechTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
