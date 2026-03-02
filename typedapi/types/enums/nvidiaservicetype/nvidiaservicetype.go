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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

// Package nvidiaservicetype
package nvidiaservicetype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6/specification/inference/_types/CommonTypes.ts#L1859-L1861
type NvidiaServiceType struct {
	Name string
}

var (
	Nvidia = NvidiaServiceType{"nvidia"}
)

func (n NvidiaServiceType) MarshalText() (text []byte, err error) {
	return []byte(n.String()), nil
}

func (n *NvidiaServiceType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "nvidia":
		*n = Nvidia
	default:
		*n = NvidiaServiceType{string(text)}
	}

	return nil
}

func (n NvidiaServiceType) String() string {
	return n.Name
}
