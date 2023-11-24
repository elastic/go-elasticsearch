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

// Package catdatafeedcolumn
package catdatafeedcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cat/_types/CatBase.ts#L405-L471
type CatDatafeedColumn struct {
	Name string
}

var (
	Ae = CatDatafeedColumn{"ae"}

	Bc = CatDatafeedColumn{"bc"}

	Id = CatDatafeedColumn{"id"}

	Na = CatDatafeedColumn{"na"}

	Ne = CatDatafeedColumn{"ne"}

	Ni = CatDatafeedColumn{"ni"}

	Nn = CatDatafeedColumn{"nn"}

	Sba = CatDatafeedColumn{"sba"}

	Sc = CatDatafeedColumn{"sc"}

	Seah = CatDatafeedColumn{"seah"}

	St = CatDatafeedColumn{"st"}

	S = CatDatafeedColumn{"s"}
)

func (c CatDatafeedColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatDatafeedColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "ae":
		*c = Ae
	case "bc":
		*c = Bc
	case "id":
		*c = Id
	case "na":
		*c = Na
	case "ne":
		*c = Ne
	case "ni":
		*c = Ni
	case "nn":
		*c = Nn
	case "sba":
		*c = Sba
	case "sc":
		*c = Sc
	case "seah":
		*c = Seah
	case "st":
		*c = St
	case "s":
		*c = S
	default:
		*c = CatDatafeedColumn{string(text)}
	}

	return nil
}

func (c CatDatafeedColumn) String() string {
	return c.Name
}
