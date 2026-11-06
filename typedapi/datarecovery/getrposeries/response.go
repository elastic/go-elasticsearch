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

package getrposeries

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Response holds the response body struct for the package getrposeries
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/data_recovery/get_rpo_series/DataRecoveryGetRpoSeriesResponse.ts#L24-L37
type Response struct {
	// BucketCount The number of buckets in the series.
	BucketCount int `json:"bucket_count"`
	// BucketDuration The duration of each bucket.
	BucketDuration types.Duration `json:"bucket_duration"`
	// Buckets The buckets in ascending order by start time.
	Buckets []types.RpoBucket `json:"buckets"`
	// End The exclusive end of the range.
	End types.DateTime `json:"end"`
	// Start The inclusive start of the range.
	Start types.DateTime `json:"start"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}

func (s *Response) UnmarshalJSON(data []byte) error {
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

		case "bucket_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "BucketCount", err)
				}
				s.BucketCount = value
			case float64:
				f := int(v)
				s.BucketCount = f
			}

		case "bucket_duration":
			if err := dec.Decode(&s.BucketDuration); err != nil {
				return fmt.Errorf("%s | %w", "BucketDuration", err)
			}

		case "buckets":
			if err := dec.Decode(&s.Buckets); err != nil {
				return fmt.Errorf("%s | %w", "Buckets", err)
			}

		case "end":
			if err := dec.Decode(&s.End); err != nil {
				return fmt.Errorf("%s | %w", "End", err)
			}

		case "start":
			if err := dec.Decode(&s.Start); err != nil {
				return fmt.Errorf("%s | %w", "Start", err)
			}

		}
	}
	return nil
}
