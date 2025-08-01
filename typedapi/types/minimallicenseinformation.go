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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/licensestatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/licensetype"
)

// MinimalLicenseInformation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/xpack/info/types.ts#L34-L40
type MinimalLicenseInformation struct {
	ExpiryDateInMillis int64                       `json:"expiry_date_in_millis"`
	Mode               licensetype.LicenseType     `json:"mode"`
	Status             licensestatus.LicenseStatus `json:"status"`
	Type               licensetype.LicenseType     `json:"type"`
	Uid                string                      `json:"uid"`
}

func (s *MinimalLicenseInformation) UnmarshalJSON(data []byte) error {

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
				return fmt.Errorf("%s | %w", "ExpiryDateInMillis", err)
			}

		case "mode":
			if err := dec.Decode(&s.Mode); err != nil {
				return fmt.Errorf("%s | %w", "Mode", err)
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
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Uid", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Uid = o

		}
	}
	return nil
}

// NewMinimalLicenseInformation returns a MinimalLicenseInformation.
func NewMinimalLicenseInformation() *MinimalLicenseInformation {
	r := &MinimalLicenseInformation{}

	return r
}
