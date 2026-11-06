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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Usage of the indices with a given `index.mode`.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/xpack/usage/types.ts#L693-L709
type IndexModeUsage struct {
	Available bool `json:"available"`
	// DataStreamsCount The number of data streams in this index mode.
	DataStreamsCount int64 `json:"data_streams_count"`
	// DataStreamsManagedByDlm The number of data streams in this index mode that are managed by data stream
	// lifecycle.
	DataStreamsManagedByDlm int64 `json:"data_streams_managed_by_dlm"`
	// DataStreamsManagedByIlm The number of data streams in this index mode that are managed by index
	// lifecycle management.
	DataStreamsManagedByIlm int64 `json:"data_streams_managed_by_ilm"`
	Enabled                 bool  `json:"enabled"`
	// IndicesCount The number of indices in this index mode.
	IndicesCount int64 `json:"indices_count"`
	// IndicesWithSyntheticSource The number of indices in this index mode that use synthetic `_source`.
	IndicesWithSyntheticSource int64 `json:"indices_with_synthetic_source"`
	// NumDocs The number of documents in the indices in this index mode.
	NumDocs int64 `json:"num_docs"`
	// SizeInBytes The size of the indices in this index mode, in bytes.
	SizeInBytes int64 `json:"size_in_bytes"`
}

func (s *IndexModeUsage) UnmarshalJSON(data []byte) error {

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

		case "available":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Available", err)
				}
				s.Available = value
			case bool:
				s.Available = v
			}

		case "data_streams_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DataStreamsCount", err)
				}
				s.DataStreamsCount = value
			case float64:
				f := int64(v)
				s.DataStreamsCount = f
			}

		case "data_streams_managed_by_dlm":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DataStreamsManagedByDlm", err)
				}
				s.DataStreamsManagedByDlm = value
			case float64:
				f := int64(v)
				s.DataStreamsManagedByDlm = f
			}

		case "data_streams_managed_by_ilm":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "DataStreamsManagedByIlm", err)
				}
				s.DataStreamsManagedByIlm = value
			case float64:
				f := int64(v)
				s.DataStreamsManagedByIlm = f
			}

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		case "indices_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndicesCount", err)
				}
				s.IndicesCount = value
			case float64:
				f := int64(v)
				s.IndicesCount = f
			}

		case "indices_with_synthetic_source":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndicesWithSyntheticSource", err)
				}
				s.IndicesWithSyntheticSource = value
			case float64:
				f := int64(v)
				s.IndicesWithSyntheticSource = f
			}

		case "num_docs":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumDocs", err)
				}
				s.NumDocs = value
			case float64:
				f := int64(v)
				s.NumDocs = f
			}

		case "size_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SizeInBytes", err)
				}
				s.SizeInBytes = value
			case float64:
				f := int64(v)
				s.SizeInBytes = f
			}

		}
	}
	return nil
}

// NewIndexModeUsage returns a IndexModeUsage.
func NewIndexModeUsage() *IndexModeUsage {
	r := &IndexModeUsage{}

	return r
}
