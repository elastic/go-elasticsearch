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

// DataframeAnalyticsAuthorization type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/Authorization.ts#L45-L57
type DataframeAnalyticsAuthorization struct {
	// ApiKey If an API key was used for the most recent update to the job, its name and
	// identifier are listed in the response.
	ApiKey *ApiKeyAuthorization `json:"api_key,omitempty"`
	// Roles If a user ID was used for the most recent update to the job, its roles at the
	// time of the update are listed in the response.
	Roles []string `json:"roles,omitempty"`
	// ServiceAccount If a service account was used for the most recent update to the job, the
	// account name is listed in the response.
	ServiceAccount *string `json:"service_account,omitempty"`
}

func (s *DataframeAnalyticsAuthorization) UnmarshalJSON(data []byte) error {

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

// NewDataframeAnalyticsAuthorization returns a DataframeAnalyticsAuthorization.
func NewDataframeAnalyticsAuthorization() *DataframeAnalyticsAuthorization {
	r := &DataframeAnalyticsAuthorization{}

	return r
}
