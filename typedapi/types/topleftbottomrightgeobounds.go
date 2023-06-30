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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// TopLeftBottomRightGeoBounds type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/Geo.ts#L145-L148
type TopLeftBottomRightGeoBounds struct {
	BottomRight GeoLocation `json:"bottom_right"`
	TopLeft     GeoLocation `json:"top_left"`
}

func (s *TopLeftBottomRightGeoBounds) UnmarshalJSON(data []byte) error {

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

		case "bottom_right":
			if err := dec.Decode(&s.BottomRight); err != nil {
				return err
			}

		case "top_left":
			if err := dec.Decode(&s.TopLeft); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewTopLeftBottomRightGeoBounds returns a TopLeftBottomRightGeoBounds.
func NewTopLeftBottomRightGeoBounds() *TopLeftBottomRightGeoBounds {
	r := &TopLeftBottomRightGeoBounds{}

	return r
}
