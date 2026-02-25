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

// Package catmastercolumn
package catmastercolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L2151-L2171
type CatMasterColumn struct {
	Name string
}

var (
	Id = CatMasterColumn{"id"}

	Host = CatMasterColumn{"host"}

	Ip = CatMasterColumn{"ip"}

	Node = CatMasterColumn{"node"}
)

func (c CatMasterColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatMasterColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "id":
		*c = Id
	case "host":
		*c = Host
	case "ip":
		*c = Ip
	case "node":
		*c = Node
	default:
		*c = CatMasterColumn{string(text)}
	}

	return nil
}

func (c CatMasterColumn) String() string {
	return c.Name
}
