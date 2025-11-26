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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

// Package catcomponentcolumn
package catcomponentcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/aa1459fbdcaf57c653729142b3b6e9982373bb1c/specification/cat/_types/CatBase.ts#L1263-L1300
type CatComponentColumn struct {
	Name string
}

var (
	Name = CatComponentColumn{"name"}

	Version = CatComponentColumn{"version"}

	Aliascount = CatComponentColumn{"alias_count"}

	Mappingcount = CatComponentColumn{"mapping_count"}

	Settingscount = CatComponentColumn{"settings_count"}

	Metadatacount = CatComponentColumn{"metadata_count"}

	Includedin = CatComponentColumn{"included_in"}
)

func (c CatComponentColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatComponentColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "name":
		*c = Name
	case "version":
		*c = Version
	case "alias_count":
		*c = Aliascount
	case "mapping_count":
		*c = Mappingcount
	case "settings_count":
		*c = Settingscount
	case "metadata_count":
		*c = Metadatacount
	case "included_in":
		*c = Includedin
	default:
		*c = CatComponentColumn{string(text)}
	}

	return nil
}

func (c CatComponentColumn) String() string {
	return c.Name
}
