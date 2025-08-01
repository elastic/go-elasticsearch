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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package delegatepki

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package delegatepki
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/security/delegate_pki/SecurityDelegatePkiRequest.ts#L22-L57
type Request struct {

	// X509CertificateChain The X509Certificate chain, which is represented as an ordered string array.
	// Each string in the array is a base64-encoded (Section 4 of RFC4648 - not
	// base64url-encoded) of the certificate's DER encoding.
	//
	// The first element is the target certificate that contains the subject
	// distinguished name that is requesting access.
	// This may be followed by additional certificates; each subsequent certificate
	// is used to certify the previous one.
	X509CertificateChain []string `json:"x509_certificate_chain"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Delegatepki request: %w", err)
	}

	return &req, nil
}
