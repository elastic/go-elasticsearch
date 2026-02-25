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

// Package catnodeattrscolumn
package catnodeattrscolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L2173-L2214
type CatNodeattrsColumn struct {
	Name string
}

var (

	// Node The node name.
	Node = CatNodeattrsColumn{"node"}

	// Id The unique node ID.
	Id = CatNodeattrsColumn{"id"}

	// Pid The process ID.
	Pid = CatNodeattrsColumn{"pid"}

	// Host The host name.
	Host = CatNodeattrsColumn{"host"}

	// Ip The IP address.
	Ip = CatNodeattrsColumn{"ip"}

	// Port The bound transport port.
	Port = CatNodeattrsColumn{"port"}

	// Attr The attribute description.
	Attr = CatNodeattrsColumn{"attr"}

	// Value The attribute value.
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
