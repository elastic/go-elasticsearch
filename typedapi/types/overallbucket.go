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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// OverallBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/Bucket.ts#L130-L145
type OverallBucket struct {
	// BucketSpan The length of the bucket in seconds. Matches the job with the longest
	// bucket_span value.
	BucketSpan int64 `json:"bucket_span"`
	// IsInterim If true, this is an interim result. In other words, the results are
	// calculated based on partial input data.
	IsInterim bool `json:"is_interim"`
	// Jobs An array of objects that contain the max_anomaly_score per job_id.
	Jobs []OverallBucketJob `json:"jobs"`
	// OverallScore The top_n average of the maximum bucket anomaly_score per job.
	OverallScore Float64 `json:"overall_score"`
	// ResultType Internal. This is always set to overall_bucket.
	ResultType string `json:"result_type"`
	// Timestamp The start time of the bucket for which these results were calculated.
	Timestamp int64 `json:"timestamp"`
	// TimestampString The start time of the bucket for which these results were calculated.
	TimestampString DateTime `json:"timestamp_string"`
}

func (s *OverallBucket) UnmarshalJSON(data []byte) error {

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

		case "bucket_span":
			if err := dec.Decode(&s.BucketSpan); err != nil {
				return err
			}

		case "is_interim":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IsInterim = value
			case bool:
				s.IsInterim = v
			}

		case "jobs":
			if err := dec.Decode(&s.Jobs); err != nil {
				return err
			}

		case "overall_score":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.OverallScore = f
			case float64:
				f := Float64(v)
				s.OverallScore = f
			}

		case "result_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ResultType = o

		case "timestamp":
			if err := dec.Decode(&s.Timestamp); err != nil {
				return err
			}

		case "timestamp_string":
			if err := dec.Decode(&s.TimestampString); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewOverallBucket returns a OverallBucket.
func NewOverallBucket() *OverallBucket {
	r := &OverallBucket{}

	return r
}
