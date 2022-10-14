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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/policytype"
)

// Summary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/enrich/_types/Policy.ts#L23-L25
type Summary struct {
	Config map[policytype.PolicyType]Policy `json:"config"`
}

// SummaryBuilder holds Summary struct and provides a builder API.
type SummaryBuilder struct {
	v *Summary
}

// NewSummary provides a builder for the Summary struct.
func NewSummaryBuilder() *SummaryBuilder {
	r := SummaryBuilder{
		&Summary{
			Config: make(map[policytype.PolicyType]Policy, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Summary struct
func (rb *SummaryBuilder) Build() Summary {
	return *rb.v
}

func (rb *SummaryBuilder) Config(values map[policytype.PolicyType]*PolicyBuilder) *SummaryBuilder {
	tmp := make(map[policytype.PolicyType]Policy, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Config = tmp
	return rb
}
