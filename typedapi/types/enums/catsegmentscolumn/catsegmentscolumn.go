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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Package catsegmentscolumn
package catsegmentscolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cat/_types/CatBase.ts#L1147-L1212
type CatSegmentsColumn struct {
	Name string
}

var (
	Index = CatSegmentsColumn{"index"}

	Shard = CatSegmentsColumn{"shard"}

	Prirep = CatSegmentsColumn{"prirep"}

	Ip = CatSegmentsColumn{"ip"}

	Segment = CatSegmentsColumn{"segment"}

	Generation = CatSegmentsColumn{"generation"}

	Docscount = CatSegmentsColumn{"docs.count"}

	Docsdeleted = CatSegmentsColumn{"docs.deleted"}

	Size = CatSegmentsColumn{"size"}

	Sizememory = CatSegmentsColumn{"size.memory"}

	Committed = CatSegmentsColumn{"committed"}

	Searchable = CatSegmentsColumn{"searchable"}

	Version = CatSegmentsColumn{"version"}

	Compound = CatSegmentsColumn{"compound"}

	Id = CatSegmentsColumn{"id"}
)

func (c CatSegmentsColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatSegmentsColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "index":
		*c = Index
	case "shard":
		*c = Shard
	case "prirep":
		*c = Prirep
	case "ip":
		*c = Ip
	case "segment":
		*c = Segment
	case "generation":
		*c = Generation
	case "docs.count":
		*c = Docscount
	case "docs.deleted":
		*c = Docsdeleted
	case "size":
		*c = Size
	case "size.memory":
		*c = Sizememory
	case "committed":
		*c = Committed
	case "searchable":
		*c = Searchable
	case "version":
		*c = Version
	case "compound":
		*c = Compound
	case "id":
		*c = Id
	default:
		*c = CatSegmentsColumn{string(text)}
	}

	return nil
}

func (c CatSegmentsColumn) String() string {
	return c.Name
}
