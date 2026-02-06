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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

// Package catpluginscolumn
package catpluginscolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/2514615770f18dbb4e3887cc1a279995dbfd0724/specification/cat/_types/CatBase.ts#L2294-L2320
type CatPluginsColumn struct {
	Name string
}

var (
	Id = CatPluginsColumn{"id"}

	Name = CatPluginsColumn{"name"}

	Component = CatPluginsColumn{"component"}

	Version = CatPluginsColumn{"version"}

	Description = CatPluginsColumn{"description"}
)

func (c CatPluginsColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatPluginsColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "id":
		*c = Id
	case "name":
		*c = Name
	case "component":
		*c = Component
	case "version":
		*c = Version
	case "description":
		*c = Description
	default:
		*c = CatPluginsColumn{string(text)}
	}

	return nil
}

func (c CatPluginsColumn) String() string {
	return c.Name
}
