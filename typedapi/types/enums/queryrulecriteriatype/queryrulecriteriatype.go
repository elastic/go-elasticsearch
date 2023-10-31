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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Package queryrulecriteriatype
package queryrulecriteriatype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/query_ruleset/_types/QueryRuleset.ts#L54-L65
type QueryRuleCriteriaType struct {
	Name string
}

var (
	Global = QueryRuleCriteriaType{"global"}

	Exact = QueryRuleCriteriaType{"exact"}

	Exactfuzzy = QueryRuleCriteriaType{"exact_fuzzy"}

	Prefix = QueryRuleCriteriaType{"prefix"}

	Suffix = QueryRuleCriteriaType{"suffix"}

	Contains = QueryRuleCriteriaType{"contains"}

	Lt = QueryRuleCriteriaType{"lt"}

	Lte = QueryRuleCriteriaType{"lte"}

	Gt = QueryRuleCriteriaType{"gt"}

	Gte = QueryRuleCriteriaType{"gte"}
)

func (q QueryRuleCriteriaType) MarshalText() (text []byte, err error) {
	return []byte(q.String()), nil
}

func (q *QueryRuleCriteriaType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "global":
		*q = Global
	case "exact":
		*q = Exact
	case "exact_fuzzy":
		*q = Exactfuzzy
	case "prefix":
		*q = Prefix
	case "suffix":
		*q = Suffix
	case "contains":
		*q = Contains
	case "lt":
		*q = Lt
	case "lte":
		*q = Lte
	case "gt":
		*q = Gt
	case "gte":
		*q = Gte
	default:
		*q = QueryRuleCriteriaType{string(text)}
	}

	return nil
}

func (q QueryRuleCriteriaType) String() string {
	return q.Name
}
