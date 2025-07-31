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

// ClusterStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/Stats.ts#L27-L35
type ClusterStatistics struct {
	Details    map[string]ClusterDetails `json:"details,omitempty"`
	Failed     int                       `json:"failed"`
	Partial    int                       `json:"partial"`
	Running    int                       `json:"running"`
	Skipped    int                       `json:"skipped"`
	Successful int                       `json:"successful"`
	Total      int                       `json:"total"`
}

func (s *ClusterStatistics) UnmarshalJSON(data []byte) error {

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

		case "details":
			if s.Details == nil {
				s.Details = make(map[string]ClusterDetails, 0)
			}
			if err := dec.Decode(&s.Details); err != nil {
				return fmt.Errorf("%s | %w", "Details", err)
			}

		case "failed":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Failed", err)
				}
				s.Failed = value
			case float64:
				f := int(v)
				s.Failed = f
			}

		case "partial":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Partial", err)
				}
				s.Partial = value
			case float64:
				f := int(v)
				s.Partial = f
			}

		case "running":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Running", err)
				}
				s.Running = value
			case float64:
				f := int(v)
				s.Running = f
			}

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

		case "successful":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Successful", err)
				}
				s.Successful = value
			case float64:
				f := int(v)
				s.Successful = f
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

// NewClusterStatistics returns a ClusterStatistics.
func NewClusterStatistics() *ClusterStatistics {
	r := &ClusterStatistics{
		Details: make(map[string]ClusterDetails),
	}

	return r
}
