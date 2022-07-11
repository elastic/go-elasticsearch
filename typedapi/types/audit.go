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

// Audit type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/xpack/usage/types.ts#L67-L69
type Audit struct {
	Enabled bool     `json:"enabled"`
	Outputs []string `json:"outputs,omitempty"`
}

// AuditBuilder holds Audit struct and provides a builder API.
type AuditBuilder struct {
	v *Audit
}

// NewAudit provides a builder for the Audit struct.
func NewAuditBuilder() *AuditBuilder {
	r := AuditBuilder{
		&Audit{},
	}

	return &r
}

// Build finalize the chain and returns the Audit struct
func (rb *AuditBuilder) Build() Audit {
	return *rb.v
}

func (rb *AuditBuilder) Enabled(enabled bool) *AuditBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *AuditBuilder) Outputs(outputs ...string) *AuditBuilder {
	rb.v.Outputs = outputs
	return rb
}
