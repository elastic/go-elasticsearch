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

// BucketKsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/aggregations/pipeline.ts#L94-L127
type BucketKsAggregation struct {
	// Alternative A list of string values indicating which K-S test alternative to calculate.
	// The valid values
	// are: "greater", "less", "two_sided". This parameter is key for determining
	// the K-S statistic used
	// when calculating the K-S test. Default value is all possible alternative
	// hypotheses.
	Alternative []string `json:"alternative,omitempty"`
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`
	// Fractions A list of doubles indicating the distribution of the samples with which to
	// compare to the `buckets_path` results.
	// In typical usage this is the overall proportion of documents in each bucket,
	// which is compared with the actual
	// document proportions in each bucket from the sibling aggregation counts. The
	// default is to assume that overall
	// documents are uniformly distributed on these buckets, which they would be if
	// one used equal percentiles of a
	// metric to define the bucket end points.
	Fractions []Float64 `json:"fractions,omitempty"`
	Meta      Metadata  `json:"meta,omitempty"`
	Name      *string   `json:"name,omitempty"`
	// SamplingMethod Indicates the sampling methodology when calculating the K-S test. Note, this
	// is sampling of the returned values.
	// This determines the cumulative distribution function (CDF) points used
	// comparing the two samples. Default is
	// `upper_tail`, which emphasizes the upper end of the CDF points. Valid options
	// are: `upper_tail`, `uniform`,
	// and `lower_tail`.
	SamplingMethod *string `json:"sampling_method,omitempty"`
}

func (s *BucketKsAggregation) UnmarshalJSON(data []byte) error {

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

		case "alternative":
			if err := dec.Decode(&s.Alternative); err != nil {
				return err
			}

		case "buckets_path":
			if err := dec.Decode(&s.BucketsPath); err != nil {
				return err
			}

		case "fractions":
			if err := dec.Decode(&s.Fractions); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = &o

		case "sampling_method":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SamplingMethod = &o

		}
	}
	return nil
}

// NewBucketKsAggregation returns a BucketKsAggregation.
func NewBucketKsAggregation() *BucketKsAggregation {
	r := &BucketKsAggregation{}

	return r
}
