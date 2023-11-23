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

// Package licensetype
package licensetype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/license/_types/License.ts#L23-L33
type LicenseType struct {
	Name string
}

var (
	Missing = LicenseType{"missing"}

	Trial = LicenseType{"trial"}

	Basic = LicenseType{"basic"}

	Standard = LicenseType{"standard"}

	Dev = LicenseType{"dev"}

	Silver = LicenseType{"silver"}

	Gold = LicenseType{"gold"}

	Platinum = LicenseType{"platinum"}

	Enterprise = LicenseType{"enterprise"}
)

func (l LicenseType) MarshalText() (text []byte, err error) {
	return []byte(l.String()), nil
}

func (l *LicenseType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "missing":
		*l = Missing
	case "trial":
		*l = Trial
	case "basic":
		*l = Basic
	case "standard":
		*l = Standard
	case "dev":
		*l = Dev
	case "silver":
		*l = Silver
	case "gold":
		*l = Gold
	case "platinum":
		*l = Platinum
	case "enterprise":
		*l = Enterprise
	default:
		*l = LicenseType{string(text)}
	}

	return nil
}

func (l LicenseType) String() string {
	return l.Name
}
