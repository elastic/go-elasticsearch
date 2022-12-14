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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icucollationalternate"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icucollationcasefirst"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icucollationdecomposition"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icucollationstrength"
)

// IcuCollationTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/_types/analysis/icu-plugin.ts#L51-L65
type IcuCollationTokenFilter struct {
	Alternate              *icucollationalternate.IcuCollationAlternate         `json:"alternate,omitempty"`
	CaseFirst              *icucollationcasefirst.IcuCollationCaseFirst         `json:"caseFirst,omitempty"`
	CaseLevel              *bool                                                `json:"caseLevel,omitempty"`
	Country                *string                                              `json:"country,omitempty"`
	Decomposition          *icucollationdecomposition.IcuCollationDecomposition `json:"decomposition,omitempty"`
	HiraganaQuaternaryMode *bool                                                `json:"hiraganaQuaternaryMode,omitempty"`
	Language               *string                                              `json:"language,omitempty"`
	Numeric                *bool                                                `json:"numeric,omitempty"`
	Rules                  *string                                              `json:"rules,omitempty"`
	Strength               *icucollationstrength.IcuCollationStrength           `json:"strength,omitempty"`
	Type                   string                                               `json:"type,omitempty"`
	VariableTop            *string                                              `json:"variableTop,omitempty"`
	Variant                *string                                              `json:"variant,omitempty"`
	Version                *string                                              `json:"version,omitempty"`
}

// NewIcuCollationTokenFilter returns a IcuCollationTokenFilter.
func NewIcuCollationTokenFilter() *IcuCollationTokenFilter {
	r := &IcuCollationTokenFilter{}

	r.Type = "icu_collation"

	return r
}
