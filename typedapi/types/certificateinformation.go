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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// CertificateInformation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ssl/certificates/types.ts#L22-L57
type CertificateInformation struct {
	// Alias If the path refers to a container file (a jks keystore, or a PKCS#12 file),
	// it is the alias of the certificate.
	// Otherwise, it is null.
	Alias *string `json:"alias,omitempty"`
	// Expiry The ISO formatted date of the certificate's expiry (not-after) date.
	Expiry DateTime `json:"expiry"`
	// Format The format of the file.
	// Valid values include `jks`, `PKCS12`, and `PEM`.
	Format string `json:"format"`
	// HasPrivateKey Indicates whether Elasticsearch has access to the private key for this
	// certificate.
	HasPrivateKey bool `json:"has_private_key"`
	// Issuer The Distinguished Name of the certificate's issuer.
	Issuer *string `json:"issuer,omitempty"`
	// Path The path to the certificate, as configured in the `elasticsearch.yml` file.
	Path string `json:"path"`
	// SerialNumber The hexadecimal representation of the certificate's serial number.
	SerialNumber string `json:"serial_number"`
	// SubjectDn The Distinguished Name of the certificate's subject.
	SubjectDn string `json:"subject_dn"`
}

func (s *CertificateInformation) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "alias":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Alias", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Alias = &o

		case "expiry":
			if err := dec.Decode(&s.Expiry); err != nil {
				return fmt.Errorf("%s | %w", "Expiry", err)
			}

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Format", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Format = o

		case "has_private_key":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "HasPrivateKey", err)
				}
				s.HasPrivateKey = value
			case bool:
				s.HasPrivateKey = v
			}

		case "issuer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Issuer", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Issuer = &o

		case "path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Path", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Path = o

		case "serial_number":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "SerialNumber", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SerialNumber = o

		case "subject_dn":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "SubjectDn", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SubjectDn = o

		}
	}
	return nil
}

// NewCertificateInformation returns a CertificateInformation.
func NewCertificateInformation() *CertificateInformation {
	r := &CertificateInformation{}

	return r
}
