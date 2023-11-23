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

// BucketInfluencer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/Bucket.ts#L80-L128
type BucketInfluencer struct {
	// AnomalyScore A normalized score between 0-100, which is calculated for each bucket
	// influencer. This score might be updated as
	// newer data is analyzed.
	AnomalyScore Float64 `json:"anomaly_score"`
	// BucketSpan The length of the bucket in seconds. This value matches the bucket span that
	// is specified in the job.
	BucketSpan int64 `json:"bucket_span"`
	// InfluencerFieldName The field name of the influencer.
	InfluencerFieldName string `json:"influencer_field_name"`
	// InitialAnomalyScore The score between 0-100 for each bucket influencer. This score is the initial
	// value that was calculated at the
	// time the bucket was processed.
	InitialAnomalyScore Float64 `json:"initial_anomaly_score"`
	// IsInterim If true, this is an interim result. In other words, the results are
	// calculated based on partial input data.
	IsInterim bool `json:"is_interim"`
	// JobId Identifier for the anomaly detection job.
	JobId string `json:"job_id"`
	// Probability The probability that the bucket has this behavior, in the range 0 to 1. This
	// value can be held to a high precision
	// of over 300 decimal places, so the `anomaly_score` is provided as a
	// human-readable and friendly interpretation of
	// this.
	Probability Float64 `json:"probability"`
	// RawAnomalyScore Internal.
	RawAnomalyScore Float64 `json:"raw_anomaly_score"`
	// ResultType Internal. This value is always set to `bucket_influencer`.
	ResultType string `json:"result_type"`
	// Timestamp The start time of the bucket for which these results were calculated.
	Timestamp int64 `json:"timestamp"`
	// TimestampString The start time of the bucket for which these results were calculated.
	TimestampString DateTime `json:"timestamp_string,omitempty"`
}

func (s *BucketInfluencer) UnmarshalJSON(data []byte) error {

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

		case "anomaly_score":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.AnomalyScore = f
			case float64:
				f := Float64(v)
				s.AnomalyScore = f
			}

		case "bucket_span":
			if err := dec.Decode(&s.BucketSpan); err != nil {
				return err
			}

		case "influencer_field_name":
			if err := dec.Decode(&s.InfluencerFieldName); err != nil {
				return err
			}

		case "initial_anomaly_score":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.InitialAnomalyScore = f
			case float64:
				f := Float64(v)
				s.InitialAnomalyScore = f
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

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return err
			}

		case "probability":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Probability = f
			case float64:
				f := Float64(v)
				s.Probability = f
			}

		case "raw_anomaly_score":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.RawAnomalyScore = f
			case float64:
				f := Float64(v)
				s.RawAnomalyScore = f
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

// NewBucketInfluencer returns a BucketInfluencer.
func NewBucketInfluencer() *BucketInfluencer {
	r := &BucketInfluencer{}

	return r
}
