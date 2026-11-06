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

// KnnEvalResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/knn_eval/types.ts#L129-L154
type KnnEvalResult struct {
	// ExcludedQueries The number of queries excluded from the mean recall, because the search found
	// a document scoring above the baseline's lowest top `k` score.
	ExcludedQueries int64 `json:"excluded_queries"`
	// IncludedQueries The number of queries included in the mean recall.
	IncludedQueries int64 `json:"included_queries"`
	// KnnSettings The evaluated settings.
	KnnSettings KnnEvalReportedSettings `json:"knn_settings"`
	// Recall The mean recall of the included queries, or `null` if no query is included.
	Recall *Float64 `json:"recall,omitempty"`
	// TookMs The time the searches took, in milliseconds.
	TookMs int64 `json:"took_ms"`
	// VectorOps The number of vector operations of the searches.
	VectorOps int64 `json:"vector_ops"`
}

func (s *KnnEvalResult) UnmarshalJSON(data []byte) error {

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

		case "excluded_queries":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ExcludedQueries", err)
				}
				s.ExcludedQueries = value
			case float64:
				f := int64(v)
				s.ExcludedQueries = f
			}

		case "included_queries":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "IncludedQueries", err)
				}
				s.IncludedQueries = value
			case float64:
				f := int64(v)
				s.IncludedQueries = f
			}

		case "knn_settings":
			if err := dec.Decode(&s.KnnSettings); err != nil {
				return fmt.Errorf("%s | %w", "KnnSettings", err)
			}

		case "recall":
			if err := dec.Decode(&s.Recall); err != nil {
				return fmt.Errorf("%s | %w", "Recall", err)
			}

		case "took_ms":
			if err := dec.Decode(&s.TookMs); err != nil {
				return fmt.Errorf("%s | %w", "TookMs", err)
			}

		case "vector_ops":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "VectorOps", err)
				}
				s.VectorOps = value
			case float64:
				f := int64(v)
				s.VectorOps = f
			}

		}
	}
	return nil
}

// NewKnnEvalResult returns a KnnEvalResult.
func NewKnnEvalResult() *KnnEvalResult {
	r := &KnnEvalResult{}

	return r
}
