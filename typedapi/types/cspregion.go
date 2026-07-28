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
// https://github.com/elastic/elasticsearch-specification/tree/7fd0bd13eaf28bd179fc906f57da09e852eb818e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// A cloud service provider region.
//
// https://github.com/elastic/elasticsearch-specification/blob/7fd0bd13eaf28bd179fc906f57da09e852eb818e/specification/inference/_types/RegionPolicy.ts#L22-L34
type CspRegion struct {
	// Csp The cloud service provider, for example `aws`, `gcp`, or `azure`.
	Csp string `json:"csp"`
	// Region The region of the cloud service provider, for example `us-east-1`.
	Region string `json:"region"`
}

func (s *CspRegion) UnmarshalJSON(data []byte) error {

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

		case "csp":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Csp", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Csp = o

		case "region":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Region", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Region = o

		}
	}
	return nil
}

// NewCspRegion returns a CspRegion.
func NewCspRegion() *CspRegion {
	r := &CspRegion{}

	return r
}

type CspRegionVariant interface {
	CspRegionCaster() *CspRegion
}

func (s *CspRegion) CspRegionCaster() *CspRegion {
	return s
}
