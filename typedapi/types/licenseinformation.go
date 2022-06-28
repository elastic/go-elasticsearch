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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/licensestatus"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/licensetype"
)

// LicenseInformation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/license/get/types.ts#L25-L38
type LicenseInformation struct {
	ExpiryDate         *DateString                 `json:"expiry_date,omitempty"`
	ExpiryDateInMillis *EpochMillis                `json:"expiry_date_in_millis,omitempty"`
	IssueDate          DateString                  `json:"issue_date"`
	IssueDateInMillis  EpochMillis                 `json:"issue_date_in_millis"`
	IssuedTo           string                      `json:"issued_to"`
	Issuer             string                      `json:"issuer"`
	MaxNodes           int64                       `json:"max_nodes,omitempty"`
	MaxResourceUnits   int                         `json:"max_resource_units,omitempty"`
	StartDateInMillis  EpochMillis                 `json:"start_date_in_millis"`
	Status             licensestatus.LicenseStatus `json:"status"`
	Type               licensetype.LicenseType     `json:"type"`
	Uid                Uuid                        `json:"uid"`
}

// LicenseInformationBuilder holds LicenseInformation struct and provides a builder API.
type LicenseInformationBuilder struct {
	v *LicenseInformation
}

// NewLicenseInformation provides a builder for the LicenseInformation struct.
func NewLicenseInformationBuilder() *LicenseInformationBuilder {
	r := LicenseInformationBuilder{
		&LicenseInformation{},
	}

	return &r
}

// Build finalize the chain and returns the LicenseInformation struct
func (rb *LicenseInformationBuilder) Build() LicenseInformation {
	return *rb.v
}

func (rb *LicenseInformationBuilder) ExpiryDate(expirydate DateString) *LicenseInformationBuilder {
	rb.v.ExpiryDate = &expirydate
	return rb
}

func (rb *LicenseInformationBuilder) ExpiryDateInMillis(expirydateinmillis *EpochMillisBuilder) *LicenseInformationBuilder {
	v := expirydateinmillis.Build()
	rb.v.ExpiryDateInMillis = &v
	return rb
}

func (rb *LicenseInformationBuilder) IssueDate(issuedate DateString) *LicenseInformationBuilder {
	rb.v.IssueDate = issuedate
	return rb
}

func (rb *LicenseInformationBuilder) IssueDateInMillis(issuedateinmillis *EpochMillisBuilder) *LicenseInformationBuilder {
	v := issuedateinmillis.Build()
	rb.v.IssueDateInMillis = v
	return rb
}

func (rb *LicenseInformationBuilder) IssuedTo(issuedto string) *LicenseInformationBuilder {
	rb.v.IssuedTo = issuedto
	return rb
}

func (rb *LicenseInformationBuilder) Issuer(issuer string) *LicenseInformationBuilder {
	rb.v.Issuer = issuer
	return rb
}

func (rb *LicenseInformationBuilder) MaxNodes(maxnodes int64) *LicenseInformationBuilder {
	rb.v.MaxNodes = maxnodes
	return rb
}

func (rb *LicenseInformationBuilder) MaxResourceUnits(maxresourceunits int) *LicenseInformationBuilder {
	rb.v.MaxResourceUnits = maxresourceunits
	return rb
}

func (rb *LicenseInformationBuilder) StartDateInMillis(startdateinmillis *EpochMillisBuilder) *LicenseInformationBuilder {
	v := startdateinmillis.Build()
	rb.v.StartDateInMillis = v
	return rb
}

func (rb *LicenseInformationBuilder) Status(status licensestatus.LicenseStatus) *LicenseInformationBuilder {
	rb.v.Status = status
	return rb
}

func (rb *LicenseInformationBuilder) Type_(type_ licensetype.LicenseType) *LicenseInformationBuilder {
	rb.v.Type = type_
	return rb
}

func (rb *LicenseInformationBuilder) Uid(uid Uuid) *LicenseInformationBuilder {
	rb.v.Uid = uid
	return rb
}
