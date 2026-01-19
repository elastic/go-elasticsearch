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
type _settingsSimilarity struct {
	v types.SettingsSimilarity
}

func NewSettingsSimilarity() *_settingsSimilarity {
	return &_settingsSimilarity{v: nil}
}

// UnknownSettingsSimilarity is used to set the unknown value of the union.
// Highlited as @non_exhaustive in the specification.
func (u *_settingsSimilarity) UnknownSettingsSimilarity(unknown json.RawMessage) *_settingsSimilarity {
	u.v = unknown
	return u
}

func (u *_settingsSimilarity) SettingsSimilarityBm25(settingssimilaritybm25 types.SettingsSimilarityBm25Variant) *_settingsSimilarity {

	u.v = settingssimilaritybm25.SettingsSimilarityBm25Caster()

	return u
}

// Interface implementation for SettingsSimilarityBm25 in SettingsSimilarity union
func (u *_settingsSimilarityBm25) SettingsSimilarityCaster() *types.SettingsSimilarity {
	t := types.SettingsSimilarity(u.v)
	return &t
}

func (u *_settingsSimilarity) SettingsSimilarityBoolean(settingssimilarityboolean types.SettingsSimilarityBooleanVariant) *_settingsSimilarity {

	u.v = settingssimilarityboolean.SettingsSimilarityBooleanCaster()

	return u
}

// Interface implementation for SettingsSimilarityBoolean in SettingsSimilarity union
func (u *_settingsSimilarityBoolean) SettingsSimilarityCaster() *types.SettingsSimilarity {
	t := types.SettingsSimilarity(u.v)
	return &t
}

func (u *_settingsSimilarity) SettingsSimilarityDfi(settingssimilaritydfi types.SettingsSimilarityDfiVariant) *_settingsSimilarity {

	u.v = settingssimilaritydfi.SettingsSimilarityDfiCaster()

	return u
}

// Interface implementation for SettingsSimilarityDfi in SettingsSimilarity union
func (u *_settingsSimilarityDfi) SettingsSimilarityCaster() *types.SettingsSimilarity {
	t := types.SettingsSimilarity(u.v)
	return &t
}

func (u *_settingsSimilarity) SettingsSimilarityDfr(settingssimilaritydfr types.SettingsSimilarityDfrVariant) *_settingsSimilarity {

	u.v = settingssimilaritydfr.SettingsSimilarityDfrCaster()

	return u
}

// Interface implementation for SettingsSimilarityDfr in SettingsSimilarity union
func (u *_settingsSimilarityDfr) SettingsSimilarityCaster() *types.SettingsSimilarity {
	t := types.SettingsSimilarity(u.v)
	return &t
}

func (u *_settingsSimilarity) SettingsSimilarityIb(settingssimilarityib types.SettingsSimilarityIbVariant) *_settingsSimilarity {

	u.v = settingssimilarityib.SettingsSimilarityIbCaster()

	return u
}

// Interface implementation for SettingsSimilarityIb in SettingsSimilarity union
func (u *_settingsSimilarityIb) SettingsSimilarityCaster() *types.SettingsSimilarity {
	t := types.SettingsSimilarity(u.v)
	return &t
}

func (u *_settingsSimilarity) SettingsSimilarityLmd(settingssimilaritylmd types.SettingsSimilarityLmdVariant) *_settingsSimilarity {

	u.v = settingssimilaritylmd.SettingsSimilarityLmdCaster()

	return u
}

// Interface implementation for SettingsSimilarityLmd in SettingsSimilarity union
func (u *_settingsSimilarityLmd) SettingsSimilarityCaster() *types.SettingsSimilarity {
	t := types.SettingsSimilarity(u.v)
	return &t
}

func (u *_settingsSimilarity) SettingsSimilarityLmj(settingssimilaritylmj types.SettingsSimilarityLmjVariant) *_settingsSimilarity {

	u.v = settingssimilaritylmj.SettingsSimilarityLmjCaster()

	return u
}

// Interface implementation for SettingsSimilarityLmj in SettingsSimilarity union
func (u *_settingsSimilarityLmj) SettingsSimilarityCaster() *types.SettingsSimilarity {
	t := types.SettingsSimilarity(u.v)
	return &t
}

func (u *_settingsSimilarity) SettingsSimilarityScripted(settingssimilarityscripted types.SettingsSimilarityScriptedVariant) *_settingsSimilarity {

	u.v = settingssimilarityscripted.SettingsSimilarityScriptedCaster()

	return u
}

// Interface implementation for SettingsSimilarityScripted in SettingsSimilarity union
func (u *_settingsSimilarityScripted) SettingsSimilarityCaster() *types.SettingsSimilarity {
	t := types.SettingsSimilarity(u.v)
	return &t
}

func (u *_settingsSimilarity) SettingsSimilarityCaster() *types.SettingsSimilarity {
	return &u.v
}
