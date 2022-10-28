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
// https://github.com/elastic/elasticsearch-specification/tree/ec3159eb31c62611202a4fb157ea88fa6ff78e1a


// Package indexmetadatastate
package indexmetadatastate

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ec3159eb31c62611202a4fb157ea88fa6ff78e1a/specification/indices/stats/types.ts#L203-L209
type IndexMetadataState struct {
	name string
}

var (
	Open = IndexMetadataState{"open"}

	Close = IndexMetadataState{"close"}
)

func (i IndexMetadataState) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *IndexMetadataState) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "open":
		*i = Open
	case "close":
		*i = Close
	default:
		*i = IndexMetadataState{string(text)}
	}

	return nil
}

func (i IndexMetadataState) String() string {
	return i.name
}
