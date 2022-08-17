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

// KuromojiReadingFormTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/kuromoji-plugin.ts#L42-L45
type KuromojiReadingFormTokenFilter struct {
	Type      string         `json:"type,omitempty"`
	UseRomaji bool           `json:"use_romaji"`
	Version   *VersionString `json:"version,omitempty"`
}

// KuromojiReadingFormTokenFilterBuilder holds KuromojiReadingFormTokenFilter struct and provides a builder API.
type KuromojiReadingFormTokenFilterBuilder struct {
	v *KuromojiReadingFormTokenFilter
}

// NewKuromojiReadingFormTokenFilter provides a builder for the KuromojiReadingFormTokenFilter struct.
func NewKuromojiReadingFormTokenFilterBuilder() *KuromojiReadingFormTokenFilterBuilder {
	r := KuromojiReadingFormTokenFilterBuilder{
		&KuromojiReadingFormTokenFilter{},
	}

	r.v.Type = "kuromoji_readingform"

	return &r
}

// Build finalize the chain and returns the KuromojiReadingFormTokenFilter struct
func (rb *KuromojiReadingFormTokenFilterBuilder) Build() KuromojiReadingFormTokenFilter {
	return *rb.v
}

func (rb *KuromojiReadingFormTokenFilterBuilder) UseRomaji(useromaji bool) *KuromojiReadingFormTokenFilterBuilder {
	rb.v.UseRomaji = useromaji
	return rb
}

func (rb *KuromojiReadingFormTokenFilterBuilder) Version(version VersionString) *KuromojiReadingFormTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
