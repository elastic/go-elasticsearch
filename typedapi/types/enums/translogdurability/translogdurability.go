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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

// Package translogdurability
package translogdurability

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0b0db20330063a6d11f7997ff443fd2a1a827d1/specification/indices/_types/IndexSettings.ts#L381-L396
type TranslogDurability struct {
	Name string
}

var (
	Request = TranslogDurability{"request"}

	Async = TranslogDurability{"async"}
)

func (t TranslogDurability) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TranslogDurability) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "request":
		*t = Request
	case "async":
		*t = Async
	default:
		*t = TranslogDurability{string(text)}
	}

	return nil
}

func (t TranslogDurability) String() string {
	return t.Name
}
