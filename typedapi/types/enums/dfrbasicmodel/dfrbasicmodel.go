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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Package dfrbasicmodel
package dfrbasicmodel

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/Similarity.ts#L32-L40
type DFRBasicModel struct {
	Name string
}

var (
	Be = DFRBasicModel{"be"}

	D = DFRBasicModel{"d"}

	G = DFRBasicModel{"g"}

	If = DFRBasicModel{"if"}

	In = DFRBasicModel{"in"}

	Ine = DFRBasicModel{"ine"}

	P = DFRBasicModel{"p"}
)

func (d DFRBasicModel) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DFRBasicModel) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "be":
		*d = Be
	case "d":
		*d = D
	case "g":
		*d = G
	case "if":
		*d = If
	case "in":
		*d = In
	case "ine":
		*d = Ine
	case "p":
		*d = P
	default:
		*d = DFRBasicModel{string(text)}
	}

	return nil
}

func (d DFRBasicModel) String() string {
	return d.Name
}
