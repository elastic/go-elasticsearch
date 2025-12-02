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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// CartesianBoundsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/aa1459fbdcaf57c653729142b3b6e9982373bb1c/specification/_types/aggregations/Aggregate.ts#L330-L333
type CartesianBoundsAggregate struct {
	Bounds *TopLeftBottomRightGeoBounds `json:"bounds,omitempty"`
	Meta   Metadata                     `json:"meta,omitempty"`
}

func (s *CartesianBoundsAggregate) UnmarshalJSON(data []byte) error {

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

		case "bounds":
			if err := dec.Decode(&s.Bounds); err != nil {
				return fmt.Errorf("%s | %w", "Bounds", err)
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return fmt.Errorf("%s | %w", "Meta", err)
			}

		}
	}
	return nil
}

// NewCartesianBoundsAggregate returns a CartesianBoundsAggregate.
func NewCartesianBoundsAggregate() *CartesianBoundsAggregate {
	r := &CartesianBoundsAggregate{}

	return r
}
