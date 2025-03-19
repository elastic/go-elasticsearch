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
// https://github.com/elastic/elasticsearch-specification/tree/0f6f3696eb685db8b944feefb6a209ad7e385b9c

// Package ecscompatibilitytype
package ecscompatibilitytype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/0f6f3696eb685db8b944feefb6a209ad7e385b9c/specification/text_structure/_types/Structure.ts#L40-L43
type EcsCompatibilityType struct {
	Name string
}

var (
	Disabled = EcsCompatibilityType{"disabled"}

	V1 = EcsCompatibilityType{"v1"}
)

func (e EcsCompatibilityType) MarshalText() (text []byte, err error) {
	return []byte(e.String()), nil
}

func (e *EcsCompatibilityType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "disabled":
		*e = Disabled
	case "v1":
		*e = V1
	default:
		*e = EcsCompatibilityType{string(text)}
	}

	return nil
}

func (e EcsCompatibilityType) String() string {
	return e.Name
}
