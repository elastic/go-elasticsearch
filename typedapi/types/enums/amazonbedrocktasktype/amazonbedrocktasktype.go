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

// Package amazonbedrocktasktype
package amazonbedrocktasktype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0b0db20330063a6d11f7997ff443fd2a1a827d1/specification/inference/_types/CommonTypes.ts#L436-L439
type AmazonBedrockTaskType struct {
	Name string
}

var (
	Completion = AmazonBedrockTaskType{"completion"}

	Textembedding = AmazonBedrockTaskType{"text_embedding"}
)

func (a AmazonBedrockTaskType) MarshalText() (text []byte, err error) {
	return []byte(a.String()), nil
}

func (a *AmazonBedrockTaskType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "completion":
		*a = Completion
	case "text_embedding":
		*a = Textembedding
	default:
		*a = AmazonBedrockTaskType{string(text)}
	}

	return nil
}

func (a AmazonBedrockTaskType) String() string {
	return a.Name
}
