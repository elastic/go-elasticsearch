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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Package jsonprocessorconflictstrategy
package jsonprocessorconflictstrategy

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ingest/_types/Processors.ts#L1170-L1175
type JsonProcessorConflictStrategy struct {
	Name string
}

var (
	Replace = JsonProcessorConflictStrategy{"replace"}

	Merge = JsonProcessorConflictStrategy{"merge"}
)

func (j JsonProcessorConflictStrategy) MarshalText() (text []byte, err error) {
	return []byte(j.String()), nil
}

func (j *JsonProcessorConflictStrategy) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "replace":
		*j = Replace
	case "merge":
		*j = Merge
	default:
		*j = JsonProcessorConflictStrategy{string(text)}
	}

	return nil
}

func (j JsonProcessorConflictStrategy) String() string {
	return j.Name
}
