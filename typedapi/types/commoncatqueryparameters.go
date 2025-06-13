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
// https://github.com/elastic/elasticsearch-specification/tree/3a94b6715915b1e9311724a2614c643368eece90

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// CommonCatQueryParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3a94b6715915b1e9311724a2614c643368eece90/specification/_spec_utils/behaviors.ts#L77-L100
type CommonCatQueryParameters struct {
	// Format Specifies the format to return the columnar data in, can be set to
	// `text`, `json`, `cbor`, `yaml`, or `smile`.
	Format *string `json:"format,omitempty"`
	// Help When set to `true` will output available columns. This option
	// can't be combined with any other query string option.
	Help *bool `json:"help,omitempty"`
	// V When set to `true` will enable verbose output.
	V *bool `json:"v,omitempty"`
}

func (s *CommonCatQueryParameters) UnmarshalJSON(data []byte) error {

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
			s.Format = &o

		case "help":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Help", err)
				}
				s.Help = &value
			case bool:
				s.Help = &v
			}

		case "v":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "V", err)
				}
				s.V = &value
			case bool:
				s.V = &v
			}

		}
	}
	return nil
}

// NewCommonCatQueryParameters returns a CommonCatQueryParameters.
func NewCommonCatQueryParameters() *CommonCatQueryParameters {
	r := &CommonCatQueryParameters{}

	return r
}
