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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Package conditiontype
package conditiontype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/watcher/_types/Conditions.ts#L65-L71
type ConditionType struct {
	Name string
}

var (
	Always = ConditionType{"always"}

	Never = ConditionType{"never"}

	Script = ConditionType{"script"}

	Compare = ConditionType{"compare"}

	Arraycompare = ConditionType{"array_compare"}
)

func (c ConditionType) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *ConditionType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "always":
		*c = Always
	case "never":
		*c = Never
	case "script":
		*c = Script
	case "compare":
		*c = Compare
	case "array_compare":
		*c = Arraycompare
	default:
		*c = ConditionType{string(text)}
	}

	return nil
}

func (c ConditionType) String() string {
	return c.Name
}
