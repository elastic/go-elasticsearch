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
)

// DatafeedAuthorization type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/Authorization.ts#L31-L43
type DatafeedAuthorization struct {
	// ApiKey If an API key was used for the most recent update to the datafeed, its name
	// and identifier are listed in the response.
	ApiKey *ApiKeyAuthorization `json:"api_key,omitempty"`
	// Roles If a user ID was used for the most recent update to the datafeed, its roles
	// at the time of the update are listed in the response.
	Roles []string `json:"roles,omitempty"`
	// ServiceAccount If a service account was used for the most recent update to the datafeed, the
	// account name is listed in the response.
	ServiceAccount *string `json:"service_account,omitempty"`
}

func (s *DatafeedAuthorization) UnmarshalJSON(data []byte) error {

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

		case "api_key":
			if err := dec.Decode(&s.ApiKey); err != nil {
				return fmt.Errorf("%s | %w", "ApiKey", err)
			}

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
				return fmt.Errorf("%s | %w", "Roles", err)
			}

		case "service_account":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ServiceAccount", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ServiceAccount = &o

		}
	}
	return nil
}

// NewDatafeedAuthorization returns a DatafeedAuthorization.
func NewDatafeedAuthorization() *DatafeedAuthorization {
	r := &DatafeedAuthorization{}

	return r
}
