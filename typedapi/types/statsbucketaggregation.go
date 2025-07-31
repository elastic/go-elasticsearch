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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gappolicy"
)

// StatsBucketAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/pipeline.ts#L410-L410
type StatsBucketAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`
	// Format `DecimalFormat` pattern for the output value.
	// If specified, the formatted value is returned in the aggregation’s
	// `value_as_string` property.
	Format *string `json:"format,omitempty"`
	// GapPolicy Policy to apply when gaps are found in the data.
	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
}

func (s *StatsBucketAggregation) UnmarshalJSON(data []byte) error {

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

		}
	}
	return nil
}

// NewStatsBucketAggregation returns a StatsBucketAggregation.
func NewStatsBucketAggregation() *StatsBucketAggregation {
	r := &StatsBucketAggregation{}

	return r
}
