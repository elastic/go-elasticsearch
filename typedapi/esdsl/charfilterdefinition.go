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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// This is provide all the types that are part of the union.
type _charFilterDefinition struct {
	v types.CharFilterDefinition
}

func NewCharFilterDefinition() *_charFilterDefinition {
	return &_charFilterDefinition{v: nil}
}

// UnknownCharFilterDefinition is used to set the unknown value of the union.
// Highlited as @non_exhaustive in the specification.
func (u *_charFilterDefinition) UnknownCharFilterDefinition(unknown json.RawMessage) *_charFilterDefinition {
	u.v = unknown
	return u
}

func (u *_charFilterDefinition) HtmlStripCharFilter(htmlstripcharfilter types.HtmlStripCharFilterVariant) *_charFilterDefinition {

	u.v = htmlstripcharfilter.HtmlStripCharFilterCaster()

	return u
}

// Interface implementation for HtmlStripCharFilter in CharFilterDefinition union
func (u *_htmlStripCharFilter) CharFilterDefinitionCaster() *types.CharFilterDefinition {
	t := types.CharFilterDefinition(u.v)
	return &t
}

func (u *_charFilterDefinition) MappingCharFilter(mappingcharfilter types.MappingCharFilterVariant) *_charFilterDefinition {

	u.v = mappingcharfilter.MappingCharFilterCaster()

	return u
}

// Interface implementation for MappingCharFilter in CharFilterDefinition union
func (u *_mappingCharFilter) CharFilterDefinitionCaster() *types.CharFilterDefinition {
	t := types.CharFilterDefinition(u.v)
	return &t
}

func (u *_charFilterDefinition) PatternReplaceCharFilter(patternreplacecharfilter types.PatternReplaceCharFilterVariant) *_charFilterDefinition {

	u.v = patternreplacecharfilter.PatternReplaceCharFilterCaster()

	return u
}

// Interface implementation for PatternReplaceCharFilter in CharFilterDefinition union
func (u *_patternReplaceCharFilter) CharFilterDefinitionCaster() *types.CharFilterDefinition {
	t := types.CharFilterDefinition(u.v)
	return &t
}

func (u *_charFilterDefinition) IcuNormalizationCharFilter(icunormalizationcharfilter types.IcuNormalizationCharFilterVariant) *_charFilterDefinition {

	u.v = icunormalizationcharfilter.IcuNormalizationCharFilterCaster()

	return u
}

// Interface implementation for IcuNormalizationCharFilter in CharFilterDefinition union
func (u *_icuNormalizationCharFilter) CharFilterDefinitionCaster() *types.CharFilterDefinition {
	t := types.CharFilterDefinition(u.v)
	return &t
}

func (u *_charFilterDefinition) KuromojiIterationMarkCharFilter(kuromojiiterationmarkcharfilter types.KuromojiIterationMarkCharFilterVariant) *_charFilterDefinition {

	u.v = kuromojiiterationmarkcharfilter.KuromojiIterationMarkCharFilterCaster()

	return u
}

// Interface implementation for KuromojiIterationMarkCharFilter in CharFilterDefinition union
func (u *_kuromojiIterationMarkCharFilter) CharFilterDefinitionCaster() *types.CharFilterDefinition {
	t := types.CharFilterDefinition(u.v)
	return &t
}

func (u *_charFilterDefinition) CharFilterDefinitionCaster() *types.CharFilterDefinition {
	return &u.v
}
