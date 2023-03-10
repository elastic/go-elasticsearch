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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

// Package waitforevents
package waitforevents

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/common.ts#L276-L283
type WaitForEvents struct {
	Name string
}

var (
	Immediate = WaitForEvents{"immediate"}

	Urgent = WaitForEvents{"urgent"}

	High = WaitForEvents{"high"}

	Normal = WaitForEvents{"normal"}

	Low = WaitForEvents{"low"}

	Languid = WaitForEvents{"languid"}
)

func (w WaitForEvents) MarshalText() (text []byte, err error) {
	return []byte(w.String()), nil
}

func (w *WaitForEvents) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "immediate":
		*w = Immediate
	case "urgent":
		*w = Urgent
	case "high":
		*w = High
	case "normal":
		*w = Normal
	case "low":
		*w = Low
	case "languid":
		*w = Languid
	default:
		*w = WaitForEvents{string(text)}
	}

	return nil
}

func (w WaitForEvents) String() string {
	return w.Name
}
