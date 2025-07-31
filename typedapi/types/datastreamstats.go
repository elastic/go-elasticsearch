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

// DataStreamStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/indices/get_data_lifecycle_stats/IndicesGetDataLifecycleStatsResponse.ts#L46-L59
type DataStreamStats struct {
	// BackingIndicesInError The count of the backing indices for the data stream.
	BackingIndicesInError int `json:"backing_indices_in_error"`
	// BackingIndicesInTotal The count of the backing indices for the data stream that have encountered an
	// error.
	BackingIndicesInTotal int `json:"backing_indices_in_total"`
	// Name The name of the data stream.
	Name string `json:"name"`
}

func (s *DataStreamStats) UnmarshalJSON(data []byte) error {

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

		case "backing_indices_in_error":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "BackingIndicesInError", err)
				}
				s.BackingIndicesInError = value
			case float64:
				f := int(v)
				s.BackingIndicesInError = f
			}

		case "backing_indices_in_total":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "BackingIndicesInTotal", err)
				}
				s.BackingIndicesInTotal = value
			case float64:
				f := int(v)
				s.BackingIndicesInTotal = f
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		}
	}
	return nil
}

// NewDataStreamStats returns a DataStreamStats.
func NewDataStreamStats() *DataStreamStats {
	r := &DataStreamStats{}

	return r
}
