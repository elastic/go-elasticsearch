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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/licensetype"

	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// License type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/license/_types/License.ts#L42-L53
type License struct {
	ExpiryDateInMillis int64                   `json:"expiry_date_in_millis"`
	IssueDateInMillis  int64                   `json:"issue_date_in_millis"`
	IssuedTo           string                  `json:"issued_to"`
	Issuer             string                  `json:"issuer"`
	MaxNodes           int64                   `json:"max_nodes,omitempty"`
	MaxResourceUnits   *int64                  `json:"max_resource_units,omitempty"`
	Signature          string                  `json:"signature"`
	StartDateInMillis  *int64                  `json:"start_date_in_millis,omitempty"`
	Type               licensetype.LicenseType `json:"type"`
	Uid                string                  `json:"uid"`
}

func (s *License) UnmarshalJSON(data []byte) error {

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

		case "expiry_date_in_millis":
			if err := dec.Decode(&s.ExpiryDateInMillis); err != nil {
				return err
			}

		case "issue_date_in_millis":
			if err := dec.Decode(&s.IssueDateInMillis); err != nil {
				return err
			}

		case "issued_to":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.IssuedTo = o

		case "issuer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Issuer = o

		case "max_nodes":
			if err := dec.Decode(&s.MaxNodes); err != nil {
				return err
			}

		case "max_resource_units":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MaxResourceUnits = &value
			case float64:
				f := int64(v)
				s.MaxResourceUnits = &f
			}

		case "signature":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Signature = o

		case "start_date_in_millis":
			if err := dec.Decode(&s.StartDateInMillis); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "uid":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Uid = o

		}
	}
	return nil
}

// NewLicense returns a License.
func NewLicense() *License {
	r := &License{}

	return r
}
