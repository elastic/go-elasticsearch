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
	"strconv"
)

// DataStreamsStatsItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/indices/data_streams_stats/IndicesDataStreamsStatsResponse.ts#L36-L42
type DataStreamsStatsItem struct {
	BackingIndices   int      `json:"backing_indices"`
	DataStream       string   `json:"data_stream"`
	MaximumTimestamp int64    `json:"maximum_timestamp"`
	StoreSize        ByteSize `json:"store_size,omitempty"`
	StoreSizeBytes   int      `json:"store_size_bytes"`
}

func (s *DataStreamsStatsItem) UnmarshalJSON(data []byte) error {

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

		case "backing_indices":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.BackingIndices = value
			case float64:
				f := int(v)
				s.BackingIndices = f
			}

		case "data_stream":
			if err := dec.Decode(&s.DataStream); err != nil {
				return err
			}

		case "maximum_timestamp":
			if err := dec.Decode(&s.MaximumTimestamp); err != nil {
				return err
			}

		case "store_size":
			if err := dec.Decode(&s.StoreSize); err != nil {
				return err
			}

		case "store_size_bytes":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.StoreSizeBytes = value
			case float64:
				f := int(v)
				s.StoreSizeBytes = f
			}

		}
	}
	return nil
}

// NewDataStreamsStatsItem returns a DataStreamsStatsItem.
func NewDataStreamsStatsItem() *DataStreamsStatsItem {
	r := &DataStreamsStatsItem{}

	return r
}
