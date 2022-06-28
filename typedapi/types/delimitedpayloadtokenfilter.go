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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/delimitedpayloadencoding"
)

// DelimitedPayloadTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/analysis/token_filters.ts#L67-L71
type DelimitedPayloadTokenFilter struct {
	Delimiter string                                            `json:"delimiter"`
	Encoding  delimitedpayloadencoding.DelimitedPayloadEncoding `json:"encoding"`
	Type      string                                            `json:"type,omitempty"`
	Version   *VersionString                                    `json:"version,omitempty"`
}

// DelimitedPayloadTokenFilterBuilder holds DelimitedPayloadTokenFilter struct and provides a builder API.
type DelimitedPayloadTokenFilterBuilder struct {
	v *DelimitedPayloadTokenFilter
}

// NewDelimitedPayloadTokenFilter provides a builder for the DelimitedPayloadTokenFilter struct.
func NewDelimitedPayloadTokenFilterBuilder() *DelimitedPayloadTokenFilterBuilder {
	r := DelimitedPayloadTokenFilterBuilder{
		&DelimitedPayloadTokenFilter{},
	}

	r.v.Type = "delimited_payload"

	return &r
}

// Build finalize the chain and returns the DelimitedPayloadTokenFilter struct
func (rb *DelimitedPayloadTokenFilterBuilder) Build() DelimitedPayloadTokenFilter {
	return *rb.v
}

func (rb *DelimitedPayloadTokenFilterBuilder) Delimiter(delimiter string) *DelimitedPayloadTokenFilterBuilder {
	rb.v.Delimiter = delimiter
	return rb
}

func (rb *DelimitedPayloadTokenFilterBuilder) Encoding(encoding delimitedpayloadencoding.DelimitedPayloadEncoding) *DelimitedPayloadTokenFilterBuilder {
	rb.v.Encoding = encoding
	return rb
}

func (rb *DelimitedPayloadTokenFilterBuilder) Version(version VersionString) *DelimitedPayloadTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
