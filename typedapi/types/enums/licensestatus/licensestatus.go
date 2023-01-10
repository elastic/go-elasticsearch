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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Package licensestatus
package licensestatus

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/license/_types/License.ts#L35-L40
type LicenseStatus struct {
	Name string
}

var (
	Active = LicenseStatus{"active"}

	Valid = LicenseStatus{"valid"}

	Invalid = LicenseStatus{"invalid"}

	Expired = LicenseStatus{"expired"}
)

func (l LicenseStatus) MarshalText() (text []byte, err error) {
	return []byte(l.String()), nil
}

func (l *LicenseStatus) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "active":
		*l = Active
	case "valid":
		*l = Valid
	case "invalid":
		*l = Invalid
	case "expired":
		*l = Expired
	default:
		*l = LicenseStatus{string(text)}
	}

	return nil
}

func (l LicenseStatus) String() string {
	return l.Name
}
