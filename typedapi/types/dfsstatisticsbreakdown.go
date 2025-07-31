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

// DfsStatisticsBreakdown type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/search/_types/profile.ts#L170-L179
type DfsStatisticsBreakdown struct {
	CollectionStatistics      int64 `json:"collection_statistics"`
	CollectionStatisticsCount int64 `json:"collection_statistics_count"`
	CreateWeight              int64 `json:"create_weight"`
	CreateWeightCount         int64 `json:"create_weight_count"`
	Rewrite                   int64 `json:"rewrite"`
	RewriteCount              int64 `json:"rewrite_count"`
	TermStatistics            int64 `json:"term_statistics"`
	TermStatisticsCount       int64 `json:"term_statistics_count"`
}

func (s *DfsStatisticsBreakdown) UnmarshalJSON(data []byte) error {

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

		case "collection_statistics":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CollectionStatistics", err)
				}
				s.CollectionStatistics = value
			case float64:
				f := int64(v)
				s.CollectionStatistics = f
			}

		case "collection_statistics_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CollectionStatisticsCount", err)
				}
				s.CollectionStatisticsCount = value
			case float64:
				f := int64(v)
				s.CollectionStatisticsCount = f
			}

		case "create_weight":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CreateWeight", err)
				}
				s.CreateWeight = value
			case float64:
				f := int64(v)
				s.CreateWeight = f
			}

		case "create_weight_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CreateWeightCount", err)
				}
				s.CreateWeightCount = value
			case float64:
				f := int64(v)
				s.CreateWeightCount = f
			}

		case "rewrite":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Rewrite", err)
				}
				s.Rewrite = value
			case float64:
				f := int64(v)
				s.Rewrite = f
			}

		case "rewrite_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "RewriteCount", err)
				}
				s.RewriteCount = value
			case float64:
				f := int64(v)
				s.RewriteCount = f
			}

		case "term_statistics":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TermStatistics", err)
				}
				s.TermStatistics = value
			case float64:
				f := int64(v)
				s.TermStatistics = f
			}

		case "term_statistics_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TermStatisticsCount", err)
				}
				s.TermStatisticsCount = value
			case float64:
				f := int64(v)
				s.TermStatisticsCount = f
			}

		}
	}
	return nil
}

// NewDfsStatisticsBreakdown returns a DfsStatisticsBreakdown.
func NewDfsStatisticsBreakdown() *DfsStatisticsBreakdown {
	r := &DfsStatisticsBreakdown{}

	return r
}
