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

// Package versiontype
package versiontype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/_types/common.ts#L99-L114
type VersionType struct {
	Name string
}

var (

	// Internal Use internal versioning that starts at 1 and increments with each update or
	// delete.
	Internal = VersionType{"internal"}

	// External Only index the document if the specified version is strictly higher than the
	// version of the stored document or if there is no existing document.
	External = VersionType{"external"}

	// Externalgte Only index the document if the specified version is equal or higher than the
	// version of the stored document or if there is no existing document. NOTE: The
	// `external_gte` version type is meant for special use cases and should be used
	// with care. If used incorrectly, it can result in loss of data.
	Externalgte = VersionType{"external_gte"}
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
	default:
		*v = VersionType{string(text)}
	}

	return nil
}

func (v VersionType) String() string {
	return v.Name
}
