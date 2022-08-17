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

// CharFilterDefinition holds the union for the following types:
//
//	HtmlStripCharFilter
//	IcuNormalizationCharFilter
//	KuromojiIterationMarkCharFilter
//	MappingCharFilter
//	PatternReplaceCharFilter
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/char_filters.ts#L32-L41
type CharFilterDefinition interface{}

// CharFilterDefinitionBuilder holds CharFilterDefinition struct and provides a builder API.
type CharFilterDefinitionBuilder struct {
	v CharFilterDefinition
}

// NewCharFilterDefinition provides a builder for the CharFilterDefinition struct.
func NewCharFilterDefinitionBuilder() *CharFilterDefinitionBuilder {
	return &CharFilterDefinitionBuilder{}
}

// Build finalize the chain and returns the CharFilterDefinition struct
func (u *CharFilterDefinitionBuilder) Build() CharFilterDefinition {
	return u.v
}

func (u *CharFilterDefinitionBuilder) HtmlStripCharFilter(htmlstripcharfilter *HtmlStripCharFilterBuilder) *CharFilterDefinitionBuilder {
	v := htmlstripcharfilter.Build()
	u.v = &v
	return u
}

func (u *CharFilterDefinitionBuilder) IcuNormalizationCharFilter(icunormalizationcharfilter *IcuNormalizationCharFilterBuilder) *CharFilterDefinitionBuilder {
	v := icunormalizationcharfilter.Build()
	u.v = &v
	return u
}

func (u *CharFilterDefinitionBuilder) KuromojiIterationMarkCharFilter(kuromojiiterationmarkcharfilter *KuromojiIterationMarkCharFilterBuilder) *CharFilterDefinitionBuilder {
	v := kuromojiiterationmarkcharfilter.Build()
	u.v = &v
	return u
}

func (u *CharFilterDefinitionBuilder) MappingCharFilter(mappingcharfilter *MappingCharFilterBuilder) *CharFilterDefinitionBuilder {
	v := mappingcharfilter.Build()
	u.v = &v
	return u
}

func (u *CharFilterDefinitionBuilder) PatternReplaceCharFilter(patternreplacecharfilter *PatternReplaceCharFilterBuilder) *CharFilterDefinitionBuilder {
	v := patternreplacecharfilter.Build()
	u.v = &v
	return u
}
