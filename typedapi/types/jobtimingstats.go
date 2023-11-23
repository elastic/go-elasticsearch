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

// JobTimingStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/Job.ts#L332-L341
type JobTimingStats struct {
	AverageBucketProcessingTimeMs                   Float64 `json:"average_bucket_processing_time_ms,omitempty"`
	BucketCount                                     int64   `json:"bucket_count"`
	ExponentialAverageBucketProcessingTimeMs        Float64 `json:"exponential_average_bucket_processing_time_ms,omitempty"`
	ExponentialAverageBucketProcessingTimePerHourMs Float64 `json:"exponential_average_bucket_processing_time_per_hour_ms"`
	JobId                                           string  `json:"job_id"`
	MaximumBucketProcessingTimeMs                   Float64 `json:"maximum_bucket_processing_time_ms,omitempty"`
	MinimumBucketProcessingTimeMs                   Float64 `json:"minimum_bucket_processing_time_ms,omitempty"`
	TotalBucketProcessingTimeMs                     Float64 `json:"total_bucket_processing_time_ms"`
}

func (s *JobTimingStats) UnmarshalJSON(data []byte) error {

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

		case "average_bucket_processing_time_ms":
			if err := dec.Decode(&s.AverageBucketProcessingTimeMs); err != nil {
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

		case "exponential_average_bucket_processing_time_ms":
			if err := dec.Decode(&s.ExponentialAverageBucketProcessingTimeMs); err != nil {
				return err
			}

		case "exponential_average_bucket_processing_time_per_hour_ms":
			if err := dec.Decode(&s.ExponentialAverageBucketProcessingTimePerHourMs); err != nil {
				return err
			}

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return err
			}

		case "maximum_bucket_processing_time_ms":
			if err := dec.Decode(&s.MaximumBucketProcessingTimeMs); err != nil {
				return err
			}

		case "minimum_bucket_processing_time_ms":
			if err := dec.Decode(&s.MinimumBucketProcessingTimeMs); err != nil {
				return err
			}

		case "total_bucket_processing_time_ms":
			if err := dec.Decode(&s.TotalBucketProcessingTimeMs); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewJobTimingStats returns a JobTimingStats.
func NewJobTimingStats() *JobTimingStats {
	r := &JobTimingStats{}

	return r
}
