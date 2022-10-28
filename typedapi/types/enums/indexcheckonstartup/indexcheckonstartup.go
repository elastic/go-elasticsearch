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


// Package indexcheckonstartup
package indexcheckonstartup

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ec3159eb31c62611202a4fb157ea88fa6ff78e1a/specification/indices/_types/IndexSettings.ts#L253-L260
type IndexCheckOnStartup struct {
	name string
}

var (
	True = IndexCheckOnStartup{"true"}

	False = IndexCheckOnStartup{"false"}

	Checksum = IndexCheckOnStartup{"checksum"}
)

func (i IndexCheckOnStartup) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *IndexCheckOnStartup) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "true":
		*i = True
	case "false":
		*i = False
	case "checksum":
		*i = Checksum
	default:
		*i = IndexCheckOnStartup{string(text)}
	}

	return nil
}

func (i IndexCheckOnStartup) String() string {
	return i.name
}
