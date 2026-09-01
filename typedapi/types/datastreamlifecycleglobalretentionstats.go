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
// https://github.com/elastic/elasticsearch-specification/tree/9665eef0d78c41f20c4c83e69b7c8155efd58f24

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// The `affected_data_streams` and `retention_millis` fields are only present
// when this global retention is defined.
//
// https://github.com/elastic/elasticsearch-specification/blob/9665eef0d78c41f20c4c83e69b7c8155efd58f24/specification/xpack/usage/types.ts#L206-L222
type DataStreamLifecycleGlobalRetentionStats struct {
	// AffectedDataStreams The number of data streams affected by this global retention.
	AffectedDataStreams *int64 `json:"affected_data_streams,omitempty"`
	// Defined Whether this global retention is defined for the cluster.
	Defined bool `json:"defined"`
	// RetentionMillis The global retention period in milliseconds.
	RetentionMillis *int64 `json:"retention_millis,omitempty"`
}

func (s *DataStreamLifecycleGlobalRetentionStats) UnmarshalJSON(data []byte) error {

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

		case "affected_data_streams":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "AffectedDataStreams", err)
				}
				s.AffectedDataStreams = &value
			case float64:
				f := int64(v)
				s.AffectedDataStreams = &f
			}

		case "defined":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Defined", err)
				}
				s.Defined = value
			case bool:
				s.Defined = v
			}

		case "retention_millis":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "RetentionMillis", err)
				}
				s.RetentionMillis = &value
			case float64:
				f := int64(v)
				s.RetentionMillis = &f
			}

		}
	}
	return nil
}

// NewDataStreamLifecycleGlobalRetentionStats returns a DataStreamLifecycleGlobalRetentionStats.
func NewDataStreamLifecycleGlobalRetentionStats() *DataStreamLifecycleGlobalRetentionStats {
	r := &DataStreamLifecycleGlobalRetentionStats{}

	return r
}
