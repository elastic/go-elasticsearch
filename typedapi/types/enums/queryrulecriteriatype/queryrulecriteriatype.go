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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Package queryrulecriteriatype
package queryrulecriteriatype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/query_rules/_types/QueryRuleset.ts#L95-L108
type QueryRuleCriteriaType struct {
	Name string
}

var (
	Global = QueryRuleCriteriaType{"global"}

	Exact = QueryRuleCriteriaType{"exact"}

	Exactfuzzy = QueryRuleCriteriaType{"exact_fuzzy"}

	Fuzzy = QueryRuleCriteriaType{"fuzzy"}

	Prefix = QueryRuleCriteriaType{"prefix"}

	Suffix = QueryRuleCriteriaType{"suffix"}

	Contains = QueryRuleCriteriaType{"contains"}

	Lt = QueryRuleCriteriaType{"lt"}

	Lte = QueryRuleCriteriaType{"lte"}

	Gt = QueryRuleCriteriaType{"gt"}

	Gte = QueryRuleCriteriaType{"gte"}

	Always = QueryRuleCriteriaType{"always"}
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
	case "fuzzy":
		*q = Fuzzy
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
	case "always":
		*q = Always
	default:
		*q = QueryRuleCriteriaType{string(text)}
	}

	return nil
}

func (q QueryRuleCriteriaType) String() string {
	return q.Name
}
