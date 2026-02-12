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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

// Package catcountcolumn
package catcountcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/e196f9953fa743572ee46884835f1934bce9a16b/specification/cat/_types/CatBase.ts#L1297-L1314
type CatCountColumn struct {
	Name string
}

var (
	Epoch = CatCountColumn{"epoch"}

	Timestamp = CatCountColumn{"timestamp"}

	Count = CatCountColumn{"count"}
)

func (c CatCountColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatCountColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "epoch":
		*c = Epoch
	case "timestamp":
		*c = Timestamp
	case "count":
		*c = Count
	default:
		*c = CatCountColumn{string(text)}
	}

	return nil
}

func (c CatCountColumn) String() string {
	return c.Name
}
