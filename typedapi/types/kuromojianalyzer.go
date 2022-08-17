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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/kuromojitokenizationmode"
)

// KuromojiAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/kuromoji-plugin.ts#L25-L29
type KuromojiAnalyzer struct {
	Mode           kuromojitokenizationmode.KuromojiTokenizationMode `json:"mode"`
	Type           string                                            `json:"type,omitempty"`
	UserDictionary *string                                           `json:"user_dictionary,omitempty"`
}

// KuromojiAnalyzerBuilder holds KuromojiAnalyzer struct and provides a builder API.
type KuromojiAnalyzerBuilder struct {
	v *KuromojiAnalyzer
}

// NewKuromojiAnalyzer provides a builder for the KuromojiAnalyzer struct.
func NewKuromojiAnalyzerBuilder() *KuromojiAnalyzerBuilder {
	r := KuromojiAnalyzerBuilder{
		&KuromojiAnalyzer{},
	}

	r.v.Type = "kuromoji"

	return &r
}

// Build finalize the chain and returns the KuromojiAnalyzer struct
func (rb *KuromojiAnalyzerBuilder) Build() KuromojiAnalyzer {
	return *rb.v
}

func (rb *KuromojiAnalyzerBuilder) Mode(mode kuromojitokenizationmode.KuromojiTokenizationMode) *KuromojiAnalyzerBuilder {
	rb.v.Mode = mode
	return rb
}

func (rb *KuromojiAnalyzerBuilder) UserDictionary(userdictionary string) *KuromojiAnalyzerBuilder {
	rb.v.UserDictionary = &userdictionary
	return rb
}
