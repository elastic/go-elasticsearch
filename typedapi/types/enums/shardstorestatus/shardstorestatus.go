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
// https://github.com/elastic/elasticsearch-specification/tree/19027dbdd366978ccae41842a040a636730e7c10

// Package shardstorestatus
package shardstorestatus

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/19027dbdd366978ccae41842a040a636730e7c10/specification/indices/shard_stores/types.ts#L63-L72
type ShardStoreStatus struct {
	Name string
}

var (
	Green = ShardStoreStatus{"green"}

	Yellow = ShardStoreStatus{"yellow"}

	Red = ShardStoreStatus{"red"}

	All = ShardStoreStatus{"all"}
)

func (s ShardStoreStatus) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *ShardStoreStatus) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "green":
		*s = Green
	case "yellow":
		*s = Yellow
	case "red":
		*s = Red
	case "all":
		*s = All
	default:
		*s = ShardStoreStatus{string(text)}
	}

	return nil
}

func (s ShardStoreStatus) String() string {
	return s.Name
}
