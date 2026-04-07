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
// https://github.com/elastic/elasticsearch-specification/tree/836fca874204ca4173ae5c36fb6b5107d28d2fc0

// Package storageoption
package storageoption

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/836fca874204ca4173ae5c36fb6b5107d28d2fc0/specification/searchable_snapshots/mount/types.ts#L29-L42
type StorageOption struct {
	Name string
}

var (

	// Fullcopy Fully caches the snapshotted index’s shards in the Elasticsearch cluster.
	Fullcopy = StorageOption{"full_copy"}

	// Sharedcache Uses a local cache containing only recently searched parts of the snapshotted
	// index’s data.
	Sharedcache = StorageOption{"shared_cache"}
)

func (s StorageOption) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *StorageOption) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "full_copy":
		*s = Fullcopy
	case "shared_cache":
		*s = Sharedcache
	default:
		*s = StorageOption{string(text)}
	}

	return nil
}

func (s StorageOption) String() string {
	return s.Name
}
