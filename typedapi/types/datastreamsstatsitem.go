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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

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
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/indices/data_streams_stats/IndicesDataStreamsStatsResponse.ts#L45-L65
type DataStreamsStatsItem struct {
	// BackingIndices Current number of backing indices for the data stream.
	BackingIndices int `json:"backing_indices"`
	// DataStream Name of the data stream.
	DataStream string `json:"data_stream"`
	// MaximumTimestamp The data stream’s highest `@timestamp` value, converted to milliseconds since
	// the Unix epoch.
	// NOTE: This timestamp is provided as a best effort.
	// The data stream may contain `@timestamp` values higher than this if one or
	// more of the following conditions are met:
	// The stream contains closed backing indices;
	// Backing indices with a lower generation contain higher `@timestamp` values.
	MaximumTimestamp int64 `json:"maximum_timestamp"`
	// StoreSize Total size of all shards for the data stream’s backing indices.
	// This parameter is only returned if the `human` query parameter is `true`.
	StoreSize ByteSize `json:"store_size,omitempty"`
	// StoreSizeBytes Total size, in bytes, of all shards for the data stream’s backing indices.
	StoreSizeBytes int `json:"store_size_bytes"`
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
