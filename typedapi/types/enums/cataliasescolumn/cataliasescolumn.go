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

// Package cataliasescolumn
package cataliasescolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L1162-L1194
type CatAliasesColumn struct {
	Name string
}

var (

	// Alias The name of the alias.
	Alias = CatAliasesColumn{"alias"}

	// Index The name of the index the alias points to.
	Index = CatAliasesColumn{"index"}

	// Filter The filter applied to the alias.
	Filter = CatAliasesColumn{"filter"}

	// Routingindex Index routing value for the alias.
	Routingindex = CatAliasesColumn{"routing.index"}

	// Routingsearch Search routing value for the alias.
	Routingsearch = CatAliasesColumn{"routing.search"}

	// Iswriteindex Indicates if the index is the write index for the alias.
	Iswriteindex = CatAliasesColumn{"is_write_index"}
)

func (c CatAliasesColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatAliasesColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "alias":
		*c = Alias
	case "index":
		*c = Index
	case "filter":
		*c = Filter
	case "routing.index":
		*c = Routingindex
	case "routing.search":
		*c = Routingsearch
	case "is_write_index":
		*c = Iswriteindex
	default:
		*c = CatAliasesColumn{string(text)}
	}

	return nil
}

func (c CatAliasesColumn) String() string {
	return c.Name
}
