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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package sorttype
package sorttype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cluster/stats/types.ts#L176-L182
type SortType struct {
	Name string
}

var (
	Doc = SortType{"_doc"}

	Geodistance = SortType{"_geo_distance"}

	Score = SortType{"_score"}

	Script = SortType{"_script"}

	Fieldsort = SortType{"field_sort"}
)

func (s SortType) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *SortType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "_doc":
		*s = Doc
	case "_geo_distance":
		*s = Geodistance
	case "_score":
		*s = Score
	case "_script":
		*s = Script
	case "field_sort":
		*s = Fieldsort
	default:
		*s = SortType{string(text)}
	}

	return nil
}

func (s SortType) String() string {
	return s.Name
}
