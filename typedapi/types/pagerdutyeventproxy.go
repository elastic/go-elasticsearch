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

// PagerDutyEventProxy type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Actions.ts#L56-L59
type PagerDutyEventProxy struct {
	Host *Host `json:"host,omitempty"`
	Port *int  `json:"port,omitempty"`
}

// PagerDutyEventProxyBuilder holds PagerDutyEventProxy struct and provides a builder API.
type PagerDutyEventProxyBuilder struct {
	v *PagerDutyEventProxy
}

// NewPagerDutyEventProxy provides a builder for the PagerDutyEventProxy struct.
func NewPagerDutyEventProxyBuilder() *PagerDutyEventProxyBuilder {
	r := PagerDutyEventProxyBuilder{
		&PagerDutyEventProxy{},
	}

	return &r
}

// Build finalize the chain and returns the PagerDutyEventProxy struct
func (rb *PagerDutyEventProxyBuilder) Build() PagerDutyEventProxy {
	return *rb.v
}

func (rb *PagerDutyEventProxyBuilder) Host(host Host) *PagerDutyEventProxyBuilder {
	rb.v.Host = &host
	return rb
}

func (rb *PagerDutyEventProxyBuilder) Port(port int) *PagerDutyEventProxyBuilder {
	rb.v.Port = &port
	return rb
}
