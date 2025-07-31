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

// CCSUsageClusterStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/stats/types.ts#L855-L862
type CCSUsageClusterStats struct {
	// Skipped The total number of cross-cluster search requests for which this cluster was
	// skipped.
	Skipped int `json:"skipped"`
	// Took Statistics about the time taken to execute requests against this cluster.
	Took CCSUsageTimeValue `json:"took"`
	// Total The total number of successful (not skipped) cross-cluster search requests
	// that were executed against this cluster. This may include requests where
	// partial results were returned, but not requests in which the cluster has been
	// skipped entirely.
	Total int `json:"total"`
}

func (s *CCSUsageClusterStats) UnmarshalJSON(data []byte) error {

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

		case "skipped":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Skipped", err)
				}
				s.Skipped = value
			case float64:
				f := int(v)
				s.Skipped = f
			}

		case "took":
			if err := dec.Decode(&s.Took); err != nil {
				return fmt.Errorf("%s | %w", "Took", err)
			}

		case "total":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Total", err)
				}
				s.Total = value
			case float64:
				f := int(v)
				s.Total = f
			}

		}
	}
	return nil
}

// NewCCSUsageClusterStats returns a CCSUsageClusterStats.
func NewCCSUsageClusterStats() *CCSUsageClusterStats {
	r := &CCSUsageClusterStats{}

	return r
}
