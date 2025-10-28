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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

// Package catnodeattrscolumn
package catnodeattrscolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/d520d9e8cf14cad487de5e0654007686c395b494/specification/cat/_types/CatBase.ts#L2112-L2153
type CatNodeattrsColumn struct {
	Name string
}

var (
	Node = CatNodeattrsColumn{"node"}

	Id = CatNodeattrsColumn{"id"}

	Pid = CatNodeattrsColumn{"pid"}

	Host = CatNodeattrsColumn{"host"}

	Ip = CatNodeattrsColumn{"ip"}

	Port = CatNodeattrsColumn{"port"}

	Attr = CatNodeattrsColumn{"attr"}

	Value = CatNodeattrsColumn{"value"}
)

func (c CatNodeattrsColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatNodeattrsColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "node":
		*c = Node
	case "id":
		*c = Id
	case "pid":
		*c = Pid
	case "host":
		*c = Host
	case "ip":
		*c = Ip
	case "port":
		*c = Port
	case "attr":
		*c = Attr
	case "value":
		*c = Value
	default:
		*c = CatNodeattrsColumn{string(text)}
	}

	return nil
}

func (c CatNodeattrsColumn) String() string {
	return c.Name
}
