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
// https://github.com/elastic/elasticsearch-specification/tree/abf9c2c6bb21328339daa197aae15af2ecbc46f0

package putregionpolicy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Response holds the response body struct for the package putregionpolicy
//
// https://github.com/elastic/elasticsearch-specification/blob/abf9c2c6bb21328339daa197aae15af2ecbc46f0/specification/inference/put_region_policy/PutRegionPolicyResponse.ts#L22-L25
type Response struct {
	// CreatedAt The date and time the region policy was created.
	CreatedAt types.DateTime `json:"created_at"`
	// CreatedBy The user who created the region policy.
	CreatedBy    *string            `json:"created_by,omitempty"`
	RegionPolicy types.RegionPolicy `json:"region_policy"`
	// UpdatedAt The date and time the region policy was last updated.
	UpdatedAt types.DateTime `json:"updated_at,omitempty"`
	// UpdatedBy The user who last updated the region policy.
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}

func (s *Response) UnmarshalJSON(data []byte) error {
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

		case "created_at":
			if err := dec.Decode(&s.CreatedAt); err != nil {
				return fmt.Errorf("%s | %w", "CreatedAt", err)
			}

		case "created_by":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "CreatedBy", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CreatedBy = &o

		case "region_policy":
			if err := dec.Decode(&s.RegionPolicy); err != nil {
				return fmt.Errorf("%s | %w", "RegionPolicy", err)
			}

		case "updated_at":
			if err := dec.Decode(&s.UpdatedAt); err != nil {
				return fmt.Errorf("%s | %w", "UpdatedAt", err)
			}

		case "updated_by":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "UpdatedBy", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UpdatedBy = &o

		}
	}
	return nil
}
