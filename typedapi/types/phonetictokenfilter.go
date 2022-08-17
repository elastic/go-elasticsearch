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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/phoneticencoder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/phoneticlanguage"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/phoneticnametype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/phoneticruletype"
)

// PhoneticTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/phonetic-plugin.ts#L64-L72
type PhoneticTokenFilter struct {
	Encoder     phoneticencoder.PhoneticEncoder     `json:"encoder"`
	Languageset []phoneticlanguage.PhoneticLanguage `json:"languageset"`
	MaxCodeLen  *int                                `json:"max_code_len,omitempty"`
	NameType    phoneticnametype.PhoneticNameType   `json:"name_type"`
	Replace     *bool                               `json:"replace,omitempty"`
	RuleType    phoneticruletype.PhoneticRuleType   `json:"rule_type"`
	Type        string                              `json:"type,omitempty"`
	Version     *VersionString                      `json:"version,omitempty"`
}

// PhoneticTokenFilterBuilder holds PhoneticTokenFilter struct and provides a builder API.
type PhoneticTokenFilterBuilder struct {
	v *PhoneticTokenFilter
}

// NewPhoneticTokenFilter provides a builder for the PhoneticTokenFilter struct.
func NewPhoneticTokenFilterBuilder() *PhoneticTokenFilterBuilder {
	r := PhoneticTokenFilterBuilder{
		&PhoneticTokenFilter{},
	}

	r.v.Type = "phonetic"

	return &r
}

// Build finalize the chain and returns the PhoneticTokenFilter struct
func (rb *PhoneticTokenFilterBuilder) Build() PhoneticTokenFilter {
	return *rb.v
}

func (rb *PhoneticTokenFilterBuilder) Encoder(encoder phoneticencoder.PhoneticEncoder) *PhoneticTokenFilterBuilder {
	rb.v.Encoder = encoder
	return rb
}

func (rb *PhoneticTokenFilterBuilder) Languageset(languageset ...phoneticlanguage.PhoneticLanguage) *PhoneticTokenFilterBuilder {
	rb.v.Languageset = languageset
	return rb
}

func (rb *PhoneticTokenFilterBuilder) MaxCodeLen(maxcodelen int) *PhoneticTokenFilterBuilder {
	rb.v.MaxCodeLen = &maxcodelen
	return rb
}

func (rb *PhoneticTokenFilterBuilder) NameType(nametype phoneticnametype.PhoneticNameType) *PhoneticTokenFilterBuilder {
	rb.v.NameType = nametype
	return rb
}

func (rb *PhoneticTokenFilterBuilder) Replace(replace bool) *PhoneticTokenFilterBuilder {
	rb.v.Replace = &replace
	return rb
}

func (rb *PhoneticTokenFilterBuilder) RuleType(ruletype phoneticruletype.PhoneticRuleType) *PhoneticTokenFilterBuilder {
	rb.v.RuleType = ruletype
	return rb
}

func (rb *PhoneticTokenFilterBuilder) Version(version VersionString) *PhoneticTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
