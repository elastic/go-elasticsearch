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

// Package googlemodelgardenprovider
package googlemodelgardenprovider

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/inference/_types/CommonTypes.ts#L1635-L1642
type GoogleModelGardenProvider struct {
	Name string
}

var (
	Google = GoogleModelGardenProvider{"google"}

	Anthropic = GoogleModelGardenProvider{"anthropic"}

	Meta = GoogleModelGardenProvider{"meta"}

	Huggingface = GoogleModelGardenProvider{"hugging_face"}

	Mistral = GoogleModelGardenProvider{"mistral"}

	Ai21 = GoogleModelGardenProvider{"ai21"}
)

func (g GoogleModelGardenProvider) MarshalText() (text []byte, err error) {
	return []byte(g.String()), nil
}

func (g *GoogleModelGardenProvider) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "google":
		*g = Google
	case "anthropic":
		*g = Anthropic
	case "meta":
		*g = Meta
	case "hugging_face":
		*g = Huggingface
	case "mistral":
		*g = Mistral
	case "ai21":
		*g = Ai21
	default:
		*g = GoogleModelGardenProvider{string(text)}
	}

	return nil
}

func (g GoogleModelGardenProvider) String() string {
	return g.Name
}
