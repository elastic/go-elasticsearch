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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/licensetype"
)

// License type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/license/_types/License.ts#L42-L53
type License struct {
	ExpiryDateInMillis EpochTimeUnitMillis     `json:"expiry_date_in_millis"`
	IssueDateInMillis  EpochTimeUnitMillis     `json:"issue_date_in_millis"`
	IssuedTo           string                  `json:"issued_to"`
	Issuer             string                  `json:"issuer"`
	MaxNodes           int64                   `json:"max_nodes,omitempty"`
	MaxResourceUnits   *int64                  `json:"max_resource_units,omitempty"`
	Signature          string                  `json:"signature"`
	StartDateInMillis  *EpochTimeUnitMillis    `json:"start_date_in_millis,omitempty"`
	Type               licensetype.LicenseType `json:"type"`
	Uid                string                  `json:"uid"`
}

// LicenseBuilder holds License struct and provides a builder API.
type LicenseBuilder struct {
	v *License
}

// NewLicense provides a builder for the License struct.
func NewLicenseBuilder() *LicenseBuilder {
	r := LicenseBuilder{
		&License{},
	}

	return &r
}

// Build finalize the chain and returns the License struct
func (rb *LicenseBuilder) Build() License {
	return *rb.v
}

func (rb *LicenseBuilder) ExpiryDateInMillis(expirydateinmillis *EpochTimeUnitMillisBuilder) *LicenseBuilder {
	v := expirydateinmillis.Build()
	rb.v.ExpiryDateInMillis = v
	return rb
}

func (rb *LicenseBuilder) IssueDateInMillis(issuedateinmillis *EpochTimeUnitMillisBuilder) *LicenseBuilder {
	v := issuedateinmillis.Build()
	rb.v.IssueDateInMillis = v
	return rb
}

func (rb *LicenseBuilder) IssuedTo(issuedto string) *LicenseBuilder {
	rb.v.IssuedTo = issuedto
	return rb
}

func (rb *LicenseBuilder) Issuer(issuer string) *LicenseBuilder {
	rb.v.Issuer = issuer
	return rb
}

func (rb *LicenseBuilder) MaxNodes(maxnodes int64) *LicenseBuilder {
	rb.v.MaxNodes = maxnodes
	return rb
}

func (rb *LicenseBuilder) MaxResourceUnits(maxresourceunits int64) *LicenseBuilder {
	rb.v.MaxResourceUnits = &maxresourceunits
	return rb
}

func (rb *LicenseBuilder) Signature(signature string) *LicenseBuilder {
	rb.v.Signature = signature
	return rb
}

func (rb *LicenseBuilder) StartDateInMillis(startdateinmillis *EpochTimeUnitMillisBuilder) *LicenseBuilder {
	v := startdateinmillis.Build()
	rb.v.StartDateInMillis = &v
	return rb
}

func (rb *LicenseBuilder) Type_(type_ licensetype.LicenseType) *LicenseBuilder {
	rb.v.Type = type_
	return rb
}

func (rb *LicenseBuilder) Uid(uid string) *LicenseBuilder {
	rb.v.Uid = uid
	return rb
}
