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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/licensestatus"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/licensetype"
)

// LicenseInformation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/license/get/types.ts#L25-L38
type LicenseInformation struct {
	ExpiryDate         DateTime                    `json:"expiry_date,omitempty"`
	ExpiryDateInMillis *int64                      `json:"expiry_date_in_millis,omitempty"`
	IssueDate          DateTime                    `json:"issue_date"`
	IssueDateInMillis  int64                       `json:"issue_date_in_millis"`
	IssuedTo           string                      `json:"issued_to"`
	Issuer             string                      `json:"issuer"`
	MaxNodes           *int64                      `json:"max_nodes,omitempty"`
	MaxResourceUnits   *int                        `json:"max_resource_units,omitempty"`
	StartDateInMillis  int64                       `json:"start_date_in_millis"`
	Status             licensestatus.LicenseStatus `json:"status"`
	Type               licensetype.LicenseType     `json:"type"`
	Uid                string                      `json:"uid"`
}

func (s *LicenseInformation) UnmarshalJSON(data []byte) error {

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

		case "expiry_date":
			if err := dec.Decode(&s.ExpiryDate); err != nil {
				return fmt.Errorf("%s | %w", "ExpiryDate", err)
			}

		case "expiry_date_in_millis":
			if err := dec.Decode(&s.ExpiryDateInMillis); err != nil {
				return fmt.Errorf("%s | %w", "ExpiryDateInMillis", err)
			}

		case "issue_date":
			if err := dec.Decode(&s.IssueDate); err != nil {
				return fmt.Errorf("%s | %w", "IssueDate", err)
			}

		case "issue_date_in_millis":
			if err := dec.Decode(&s.IssueDateInMillis); err != nil {
				return fmt.Errorf("%s | %w", "IssueDateInMillis", err)
			}

		case "issued_to":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "IssuedTo", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IssuedTo = o

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
			s.Issuer = o

		case "max_nodes":
			if err := dec.Decode(&s.MaxNodes); err != nil {
				return fmt.Errorf("%s | %w", "MaxNodes", err)
			}

		case "max_resource_units":
			if err := dec.Decode(&s.MaxResourceUnits); err != nil {
				return fmt.Errorf("%s | %w", "MaxResourceUnits", err)
			}

		case "start_date_in_millis":
			if err := dec.Decode(&s.StartDateInMillis); err != nil {
				return fmt.Errorf("%s | %w", "StartDateInMillis", err)
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "uid":
			if err := dec.Decode(&s.Uid); err != nil {
				return fmt.Errorf("%s | %w", "Uid", err)
			}

		}
	}
	return nil
}

// NewLicenseInformation returns a LicenseInformation.
func NewLicenseInformation() *LicenseInformation {
	r := &LicenseInformation{}

	return r
}
