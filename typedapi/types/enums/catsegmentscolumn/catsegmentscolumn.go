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

// Package catsegmentscolumn
package catsegmentscolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L1026-L1091
type CatSegmentsColumn struct {
	Name string
}

var (

	// Index The name of the index.
	Index = CatSegmentsColumn{"index"}

	// Shard The name of the shard.
	Shard = CatSegmentsColumn{"shard"}

	// Prirep The shard type. Returned values are 'primary' or 'replica'.
	Prirep = CatSegmentsColumn{"prirep"}

	// Ip IP address of the segment’s shard, such as '127.0.1.1'.
	Ip = CatSegmentsColumn{"ip"}

	// Segment The name of the segment, such as '_0'. The segment name is derived from the
	// segment generation and used internally to create file names in the directory
	// of the shard.
	Segment = CatSegmentsColumn{"segment"}

	// Generation Generation number, such as '0'. Elasticsearch increments this generation
	// number for each segment written. Elasticsearch then uses this number to
	// derive the segment name.
	Generation = CatSegmentsColumn{"generation"}

	// Docscount The number of documents as reported by Lucene. This excludes deleted
	// documents and counts any [nested
	// documents](https://www.elastic.co/docs/reference/elasticsearch/mapping-reference/nested)
	// separately from their parents. It also excludes documents which were indexed
	// recently and do not yet belong to a segment.
	Docscount = CatSegmentsColumn{"docs.count"}

	// Docsdeleted The number of deleted documents as reported by Lucene, which may be higher or
	// lower than the number of delete operations you have performed. This number
	// excludes deletes that were performed recently and do not yet belong to a
	// segment. Deleted documents are cleaned up by the [automatic merge
	// process](https://www.elastic.co/docs/reference/elasticsearch/index-settings/merge)
	// if it makes sense to do so. Also, Elasticsearch creates extra deleted
	// documents to internally track the recent history of operations on a shard.
	Docsdeleted = CatSegmentsColumn{"docs.deleted"}

	// Size The disk space used by the segment, such as '50kb'.
	Size = CatSegmentsColumn{"size"}

	// Sizememory The bytes of segment data stored in memory for efficient search, such as
	// '1264'. A value of '-1' indicates Elasticsearch was unable to compute this
	// number.
	Sizememory = CatSegmentsColumn{"size.memory"}

	// Committed If 'true', the segments is synced to disk. Segments that are synced can
	// survive a hard reboot. If 'false', the data from uncommitted segments is also
	// stored in the transaction log so that Elasticsearch is able to replay changes
	// on the next start.
	Committed = CatSegmentsColumn{"committed"}

	// Searchable If 'true', the segment is searchable. If 'false', the segment has most likely
	// been written to disk but needs a
	// [refresh](https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-indices-refresh)
	// to be searchable.
	Searchable = CatSegmentsColumn{"searchable"}

	// Version The version of Lucene used to write the segment.
	Version = CatSegmentsColumn{"version"}

	// Compound If 'true', the segment is stored in a compound file. This means Lucene merged
	// all files from the segment in a single file to save file descriptors.
	Compound = CatSegmentsColumn{"compound"}

	// Id The ID of the node, such as 'k0zy'.
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
