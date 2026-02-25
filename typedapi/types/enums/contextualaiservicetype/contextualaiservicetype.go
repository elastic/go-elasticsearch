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

// Package contextualaiservicetype
package contextualaiservicetype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/inference/_types/CommonTypes.ts#L1336-L1338
type ContextualAIServiceType struct {
	Name string
}

var (
	Contextualai = ContextualAIServiceType{"contextualai"}
)

func (c ContextualAIServiceType) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *ContextualAIServiceType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "contextualai":
		*c = Contextualai
	default:
		*c = ContextualAIServiceType{string(text)}
	}

	return nil
}

func (c ContextualAIServiceType) String() string {
	return c.Name
}
