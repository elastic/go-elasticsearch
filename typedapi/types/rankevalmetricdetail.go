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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// RankEvalMetricDetail type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/rank_eval/types.ts#L125-L134
type RankEvalMetricDetail struct {
	// Hits The hits section shows a grouping of the search results with their supplied
	// ratings
	Hits []RankEvalHitItem `json:"hits"`
	// MetricDetails The metric_details give additional information about the calculated quality
	// metric (e.g. how many of the retrieved documents were relevant). The content
	// varies for each metric but allows for better interpretation of the results
	MetricDetails map[string]map[string]json.RawMessage `json:"metric_details"`
	// MetricScore The metric_score in the details section shows the contribution of this query
	// to the global quality metric score
	MetricScore Float64 `json:"metric_score"`
	// UnratedDocs The unrated_docs section contains an _index and _id entry for each document
	// in the search result for this query that didn’t have a ratings value. This
	// can be used to ask the user to supply ratings for these documents
	UnratedDocs []UnratedDocument `json:"unrated_docs"`
}

func (s *RankEvalMetricDetail) UnmarshalJSON(data []byte) error {

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

		case "hits":
			if err := dec.Decode(&s.Hits); err != nil {
				return err
			}

		case "metric_details":
			if s.MetricDetails == nil {
				s.MetricDetails = make(map[string]map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.MetricDetails); err != nil {
				return err
			}

		case "metric_score":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.MetricScore = f
			case float64:
				f := Float64(v)
				s.MetricScore = f
			}

		case "unrated_docs":
			if err := dec.Decode(&s.UnratedDocs); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewRankEvalMetricDetail returns a RankEvalMetricDetail.
func NewRankEvalMetricDetail() *RankEvalMetricDetail {
	r := &RankEvalMetricDetail{
		MetricDetails: make(map[string]map[string]json.RawMessage, 0),
	}

	return r
}
