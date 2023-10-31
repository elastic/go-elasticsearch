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

// Influencer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/Influencer.ts#L31-L83
type Influencer struct {
	// BucketSpan The length of the bucket in seconds. This value matches the bucket span that
	// is specified in the job.
	BucketSpan int64 `json:"bucket_span"`
	// Foo Additional influencer properties are added, depending on the fields being
	// analyzed. For example, if it’s
	// analyzing `user_name` as an influencer, a field `user_name` is added to the
	// result document. This
	// information enables you to filter the anomaly results more easily.
	Foo *string `json:"foo,omitempty"`
	// InfluencerFieldName The field name of the influencer.
	InfluencerFieldName string `json:"influencer_field_name"`
	// InfluencerFieldValue The entity that influenced, contributed to, or was to blame for the anomaly.
	InfluencerFieldValue string `json:"influencer_field_value"`
	// InfluencerScore A normalized score between 0-100, which is based on the probability of the
	// influencer in this bucket aggregated
	// across detectors. Unlike `initial_influencer_score`, this value is updated by
	// a re-normalization process as new
	// data is analyzed.
	InfluencerScore Float64 `json:"influencer_score"`
	// InitialInfluencerScore A normalized score between 0-100, which is based on the probability of the
	// influencer aggregated across detectors.
	// This is the initial value that was calculated at the time the bucket was
	// processed.
	InitialInfluencerScore Float64 `json:"initial_influencer_score"`
	// IsInterim If true, this is an interim result. In other words, the results are
	// calculated based on partial input data.
	IsInterim bool `json:"is_interim"`
	// JobId Identifier for the anomaly detection job.
	JobId string `json:"job_id"`
	// Probability The probability that the influencer has this behavior, in the range 0 to 1.
	// This value can be held to a high
	// precision of over 300 decimal places, so the `influencer_score` is provided
	// as a human-readable and friendly
	// interpretation of this value.
	Probability Float64 `json:"probability"`
	// ResultType Internal. This value is always set to `influencer`.
	ResultType string `json:"result_type"`
	// Timestamp The start time of the bucket for which these results were calculated.
	Timestamp int64 `json:"timestamp"`
}

func (s *Influencer) UnmarshalJSON(data []byte) error {

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

		case "foo":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Foo = &o

		case "influencer_field_name":
			if err := dec.Decode(&s.InfluencerFieldName); err != nil {
				return err
			}

		case "influencer_field_value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.InfluencerFieldValue = o

		case "influencer_score":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.InfluencerScore = f
			case float64:
				f := Float64(v)
				s.InfluencerScore = f
			}

		case "initial_influencer_score":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.InitialInfluencerScore = f
			case float64:
				f := Float64(v)
				s.InitialInfluencerScore = f
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

		}
	}
	return nil
}

// NewInfluencer returns a Influencer.
func NewInfluencer() *Influencer {
	r := &Influencer{}

	return r
}
