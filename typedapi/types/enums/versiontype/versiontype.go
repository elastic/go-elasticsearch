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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Package versiontype
package versiontype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/common.ts#L100-L116
type VersionType struct {
	Name string
}

var (
	Internal = VersionType{"internal"}

	External = VersionType{"external"}

	Externalgte = VersionType{"external_gte"}

	Force = VersionType{"force"}
)

func (v VersionType) MarshalText() (text []byte, err error) {
	return []byte(v.String()), nil
}

func (v *VersionType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "internal":
		*v = Internal
	case "external":
		*v = External
	case "external_gte":
		*v = Externalgte
	case "force":
		*v = Force
	default:
		*v = VersionType{string(text)}
	}

	return nil
}

func (v VersionType) String() string {
	return v.Name
}
