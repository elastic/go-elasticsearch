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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

// BucketSelectorAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/pipeline.ts#L93-L101
type BucketSelectorAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`
	// Format `DecimalFormat` pattern for the output value.
	// If specified, the formatted value is returned in the aggregation’s
	// `value_as_string` property.
	Format *string `json:"format,omitempty"`
	// GapPolicy Policy to apply when gaps are found in the data.
	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	// Script The script to run for this aggregation.
	Script *Script `json:"script,omitempty"`
}

func (s *BucketSelectorAggregation) UnmarshalJSON(data []byte) error {

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

		case "buckets_path":
			if err := dec.Decode(&s.BucketsPath); err != nil {
				return fmt.Errorf("%s | %w", "BucketsPath", err)
			}

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Format", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Format = &o

		case "gap_policy":
			if err := dec.Decode(&s.GapPolicy); err != nil {
				return fmt.Errorf("%s | %w", "GapPolicy", err)
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		}
	}
	return nil
}

// NewBucketSelectorAggregation returns a BucketSelectorAggregation.
func NewBucketSelectorAggregation() *BucketSelectorAggregation {
	r := &BucketSelectorAggregation{}

	return r
}

type BucketSelectorAggregationVariant interface {
	BucketSelectorAggregationCaster() *BucketSelectorAggregation
}

func (s *BucketSelectorAggregation) BucketSelectorAggregationCaster() *BucketSelectorAggregation {
	return s
}
