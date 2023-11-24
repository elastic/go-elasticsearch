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

// ClusterStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/Stats.ts#L27-L35
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
				return err
			}

		case "failed":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Failed = value
			case float64:
				f := int(v)
				s.Failed = f
			}

		case "partial":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Partial = value
			case float64:
				f := int(v)
				s.Partial = f
			}

		case "running":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Running = value
			case float64:
				f := int(v)
				s.Running = f
			}

		case "skipped":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Skipped = value
			case float64:
				f := int(v)
				s.Skipped = f
			}

		case "successful":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Successful = value
			case float64:
				f := int(v)
				s.Successful = f
			}

		case "total":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
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
		Details: make(map[string]ClusterDetails, 0),
	}

	return r
}
