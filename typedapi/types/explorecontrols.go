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

// ExploreControls type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/graph/_types/ExploreControls.ts#L24-L49
type ExploreControls struct {
	// SampleDiversity To avoid the top-matching documents sample being dominated by a single source
	// of results, it is sometimes necessary to request diversity in the sample.
	// You can do this by selecting a single-value field and setting a maximum
	// number of documents per value for that field.
	SampleDiversity *SampleDiversity `json:"sample_diversity,omitempty"`
	// SampleSize Each hop considers a sample of the best-matching documents on each shard.
	// Using samples improves the speed of execution and keeps exploration focused
	// on meaningfully-connected terms.
	// Very small values (less than 50) might not provide sufficient
	// weight-of-evidence to identify significant connections between terms.
	// Very large sample sizes can dilute the quality of the results and increase
	// execution times.
	SampleSize *int `json:"sample_size,omitempty"`
	// Timeout The length of time in milliseconds after which exploration will be halted and
	// the results gathered so far are returned.
	// This timeout is honored on a best-effort basis.
	// Execution might overrun this timeout if, for example, a long pause is
	// encountered while FieldData is loaded for a field.
	Timeout Duration `json:"timeout,omitempty"`
	// UseSignificance Filters associated terms so only those that are significantly associated with
	// your query are included.
	UseSignificance bool `json:"use_significance"`
}

func (s *ExploreControls) UnmarshalJSON(data []byte) error {

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

		case "sample_diversity":
			if err := dec.Decode(&s.SampleDiversity); err != nil {
				return err
			}

		case "sample_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.SampleSize = &value
			case float64:
				f := int(v)
				s.SampleSize = &f
			}

		case "timeout":
			if err := dec.Decode(&s.Timeout); err != nil {
				return err
			}

		case "use_significance":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.UseSignificance = value
			case bool:
				s.UseSignificance = v
			}

		}
	}
	return nil
}

// NewExploreControls returns a ExploreControls.
func NewExploreControls() *ExploreControls {
	r := &ExploreControls{}

	return r
}
