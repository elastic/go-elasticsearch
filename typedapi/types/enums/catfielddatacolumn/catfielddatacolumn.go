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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

// Package catfielddatacolumn
package catfielddatacolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/6785a6caa1fa3ca5ab3308963d79dce923a3469f/specification/cat/_types/CatBase.ts#L1370-L1400
type CatFieldDataColumn struct {
	Name string
}

var (
	Id = CatFieldDataColumn{"id"}

	Host = CatFieldDataColumn{"host"}

	Ip = CatFieldDataColumn{"ip"}

	Node = CatFieldDataColumn{"node"}

	Field = CatFieldDataColumn{"field"}

	Size = CatFieldDataColumn{"size"}
)

func (c CatFieldDataColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatFieldDataColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "id":
		*c = Id
	case "host":
		*c = Host
	case "ip":
		*c = Ip
	case "node":
		*c = Node
	case "field":
		*c = Field
	case "size":
		*c = Size
	default:
		*c = CatFieldDataColumn{string(text)}
	}

	return nil
}

func (c CatFieldDataColumn) String() string {
	return c.Name
}
