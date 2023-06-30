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

// AnomalyExplanation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/Anomaly.ts#L156-L197
type AnomalyExplanation struct {
	// AnomalyCharacteristicsImpact Impact from the duration and magnitude of the detected anomaly relative to
	// the historical average.
	AnomalyCharacteristicsImpact *int `json:"anomaly_characteristics_impact,omitempty"`
	// AnomalyLength Length of the detected anomaly in the number of buckets.
	AnomalyLength *int `json:"anomaly_length,omitempty"`
	// AnomalyType Type of the detected anomaly: `spike` or `dip`.
	AnomalyType *string `json:"anomaly_type,omitempty"`
	// HighVariancePenalty Indicates reduction of anomaly score for the bucket with large confidence
	// intervals. If a bucket has large confidence intervals, the score is reduced.
	HighVariancePenalty *bool `json:"high_variance_penalty,omitempty"`
	// IncompleteBucketPenalty If the bucket contains fewer samples than expected, the score is reduced.
	IncompleteBucketPenalty *bool `json:"incomplete_bucket_penalty,omitempty"`
	// LowerConfidenceBound Lower bound of the 95% confidence interval.
	LowerConfidenceBound *Float64 `json:"lower_confidence_bound,omitempty"`
	// MultiBucketImpact Impact of the deviation between actual and typical values in the past 12
	// buckets.
	MultiBucketImpact *int `json:"multi_bucket_impact,omitempty"`
	// SingleBucketImpact Impact of the deviation between actual and typical values in the current
	// bucket.
	SingleBucketImpact *int `json:"single_bucket_impact,omitempty"`
	// TypicalValue Typical (expected) value for this bucket.
	TypicalValue *Float64 `json:"typical_value,omitempty"`
	// UpperConfidenceBound Upper bound of the 95% confidence interval.
	UpperConfidenceBound *Float64 `json:"upper_confidence_bound,omitempty"`
}

func (s *AnomalyExplanation) UnmarshalJSON(data []byte) error {

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

		case "anomaly_characteristics_impact":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.AnomalyCharacteristicsImpact = &value
			case float64:
				f := int(v)
				s.AnomalyCharacteristicsImpact = &f
			}

		case "anomaly_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.AnomalyLength = &value
			case float64:
				f := int(v)
				s.AnomalyLength = &f
			}

		case "anomaly_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AnomalyType = &o

		case "high_variance_penalty":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.HighVariancePenalty = &value
			case bool:
				s.HighVariancePenalty = &v
			}

		case "incomplete_bucket_penalty":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IncompleteBucketPenalty = &value
			case bool:
				s.IncompleteBucketPenalty = &v
			}

		case "lower_confidence_bound":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.LowerConfidenceBound = &f
			case float64:
				f := Float64(v)
				s.LowerConfidenceBound = &f
			}

		case "multi_bucket_impact":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MultiBucketImpact = &value
			case float64:
				f := int(v)
				s.MultiBucketImpact = &f
			}

		case "single_bucket_impact":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.SingleBucketImpact = &value
			case float64:
				f := int(v)
				s.SingleBucketImpact = &f
			}

		case "typical_value":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.TypicalValue = &f
			case float64:
				f := Float64(v)
				s.TypicalValue = &f
			}

		case "upper_confidence_bound":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.UpperConfidenceBound = &f
			case float64:
				f := Float64(v)
				s.UpperConfidenceBound = &f
			}

		}
	}
	return nil
}

// NewAnomalyExplanation returns a AnomalyExplanation.
func NewAnomalyExplanation() *AnomalyExplanation {
	r := &AnomalyExplanation{}

	return r
}
