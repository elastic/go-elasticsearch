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
// https://github.com/elastic/elasticsearch-specification/tree/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// DataStreamLifecycleRetentionStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c/specification/xpack/usage/types.ts#L181-L186
type DataStreamLifecycleRetentionStats struct {
	// AverageMillis The average configured value in milliseconds.
	AverageMillis *Float64 `json:"average_millis,omitempty"`
	// ConfiguredDataStreams The number of data streams for which this value is configured.
	ConfiguredDataStreams int64 `json:"configured_data_streams"`
	// MaximumMillis The largest configured value in milliseconds.
	MaximumMillis *int64 `json:"maximum_millis,omitempty"`
	// MinimumMillis The smallest configured value in milliseconds.
	MinimumMillis *int64 `json:"minimum_millis,omitempty"`
}

func (s *DataStreamLifecycleRetentionStats) UnmarshalJSON(data []byte) error {

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

		case "average_millis":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "AverageMillis", err)
				}
				f := Float64(value)
				s.AverageMillis = &f
			case float64:
				f := Float64(v)
				s.AverageMillis = &f
			}

		case "configured_data_streams":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ConfiguredDataStreams", err)
				}
				s.ConfiguredDataStreams = value
			case float64:
				f := int64(v)
				s.ConfiguredDataStreams = f
			}

		case "maximum_millis":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaximumMillis", err)
				}
				s.MaximumMillis = &value
			case float64:
				f := int64(v)
				s.MaximumMillis = &f
			}

		case "minimum_millis":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinimumMillis", err)
				}
				s.MinimumMillis = &value
			case float64:
				f := int64(v)
				s.MinimumMillis = &f
			}

		}
	}
	return nil
}

// NewDataStreamLifecycleRetentionStats returns a DataStreamLifecycleRetentionStats.
func NewDataStreamLifecycleRetentionStats() *DataStreamLifecycleRetentionStats {
	r := &DataStreamLifecycleRetentionStats{}

	return r
}
