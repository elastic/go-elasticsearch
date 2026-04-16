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

// Package threadtype
package threadtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/836fca874204ca4173ae5c36fb6b5107d28d2fc0/specification/_types/common.ts#L314-L325
type ThreadType struct {
	Name string
}

var (

	// Cpu Threads that consume the most CPU time.
	Cpu = ThreadType{"cpu"}

	// Wait Threads that have been in a waiting state the longest.
	Wait = ThreadType{"wait"}

	// Block Threads that have been blocked the longest.
	Block = ThreadType{"block"}

	// Gpu Threads that consume the most GPU time.
	Gpu = ThreadType{"gpu"}

	// Mem Threads that allocate the most memory.
	Mem = ThreadType{"mem"}
)

func (t ThreadType) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *ThreadType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "cpu":
		*t = Cpu
	case "wait":
		*t = Wait
	case "block":
		*t = Block
	case "gpu":
		*t = Gpu
	case "mem":
		*t = Mem
	default:
		*t = ThreadType{string(text)}
	}

	return nil
}

func (t ThreadType) String() string {
	return t.Name
}
