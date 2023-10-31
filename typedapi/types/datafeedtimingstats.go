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

// DatafeedTimingStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/Datafeed.ts#L171-L196
type DatafeedTimingStats struct {
	// AverageSearchTimePerBucketMs The average search time per bucket, in milliseconds.
	AverageSearchTimePerBucketMs Float64 `json:"average_search_time_per_bucket_ms,omitempty"`
	// BucketCount The number of buckets processed.
	BucketCount int64 `json:"bucket_count"`
	// ExponentialAverageSearchTimePerHourMs The exponential average search time per hour, in milliseconds.
	ExponentialAverageSearchTimePerHourMs Float64 `json:"exponential_average_search_time_per_hour_ms"`
	// JobId Identifier for the anomaly detection job.
	JobId string `json:"job_id"`
	// SearchCount The number of searches run by the datafeed.
	SearchCount int64 `json:"search_count"`
	// TotalSearchTimeMs The total time the datafeed spent searching, in milliseconds.
	TotalSearchTimeMs Float64 `json:"total_search_time_ms"`
}

func (s *DatafeedTimingStats) UnmarshalJSON(data []byte) error {

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

		case "average_search_time_per_bucket_ms":
			if err := dec.Decode(&s.AverageSearchTimePerBucketMs); err != nil {
				return err
			}

		case "bucket_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.BucketCount = value
			case float64:
				f := int64(v)
				s.BucketCount = f
			}

		case "exponential_average_search_time_per_hour_ms":
			if err := dec.Decode(&s.ExponentialAverageSearchTimePerHourMs); err != nil {
				return err
			}

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return err
			}

		case "search_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.SearchCount = value
			case float64:
				f := int64(v)
				s.SearchCount = f
			}

		case "total_search_time_ms":
			if err := dec.Decode(&s.TotalSearchTimeMs); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDatafeedTimingStats returns a DatafeedTimingStats.
func NewDatafeedTimingStats() *DatafeedTimingStats {
	r := &DatafeedTimingStats{}

	return r
}
