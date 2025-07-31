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

// AdjacencyMatrixAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/bucket.ts#L60-L70
type AdjacencyMatrixAggregation struct {
	// Filters Filters used to create buckets.
	// At least one filter is required.
	Filters map[string]Query `json:"filters,omitempty"`
	// Separator Separator used to concatenate filter names. Defaults to &.
	Separator *string `json:"separator,omitempty"`
}

func (s *AdjacencyMatrixAggregation) UnmarshalJSON(data []byte) error {

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

		case "filters":
			if s.Filters == nil {
				s.Filters = make(map[string]Query, 0)
			}
			if err := dec.Decode(&s.Filters); err != nil {
				return fmt.Errorf("%s | %w", "Filters", err)
			}

		case "separator":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Separator", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Separator = &o

		}
	}
	return nil
}

// NewAdjacencyMatrixAggregation returns a AdjacencyMatrixAggregation.
func NewAdjacencyMatrixAggregation() *AdjacencyMatrixAggregation {
	r := &AdjacencyMatrixAggregation{
		Filters: make(map[string]Query),
	}

	return r
}
