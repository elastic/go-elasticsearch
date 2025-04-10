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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/licensetype"
)

type _license struct {
	v *types.License
}

func NewLicense(issuedto string, issuer string, signature string, type_ licensetype.LicenseType, uid string) *_license {

	tmp := &_license{v: types.NewLicense()}

	tmp.IssuedTo(issuedto)

	tmp.Issuer(issuer)

	tmp.Signature(signature)

	tmp.Type(type_)

	tmp.Uid(uid)

	return tmp

}

func (s *_license) ExpiryDateInMillis(epochtimeunitmillis int64) *_license {

	s.v.ExpiryDateInMillis = epochtimeunitmillis

	return s
}

func (s *_license) IssueDateInMillis(epochtimeunitmillis int64) *_license {

	s.v.IssueDateInMillis = epochtimeunitmillis

	return s
}

func (s *_license) IssuedTo(issuedto string) *_license {

	s.v.IssuedTo = issuedto

	return s
}

func (s *_license) Issuer(issuer string) *_license {

	s.v.Issuer = issuer

	return s
}

func (s *_license) MaxNodes(maxnodes int64) *_license {

	s.v.MaxNodes = &maxnodes

	return s
}

func (s *_license) MaxResourceUnits(maxresourceunits int64) *_license {

	s.v.MaxResourceUnits = &maxresourceunits

	return s
}

func (s *_license) Signature(signature string) *_license {

	s.v.Signature = signature

	return s
}

func (s *_license) StartDateInMillis(epochtimeunitmillis int64) *_license {

	s.v.StartDateInMillis = &epochtimeunitmillis

	return s
}

func (s *_license) Type(type_ licensetype.LicenseType) *_license {

	s.v.Type = type_
	return s
}

func (s *_license) Uid(uid string) *_license {

	s.v.Uid = uid

	return s
}

func (s *_license) LicenseCaster() *types.License {
	return s.v
}
