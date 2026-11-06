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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// KnnEvalSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/knn_eval/types.ts#L70-L91
type KnnEvalSettings struct {
	// Exact If `true`, compares every vector at full precision. Only applies to the
	// `baseline`, requires `allow_exact_baseline`, and can't be combined with other
	// settings.
	Exact *bool `json:"exact,omitempty"`
	// NumCandidates The number of nearest neighbor candidates to consider per shard. It can't be
	// less than `k`, or more than 10000 (9999 with sampling).
	NumCandidates *int `json:"num_candidates,omitempty"`
	// RescoreVector The rescoring of the candidates. Only applies to quantized fields.
	RescoreVector *RescoreVector `json:"rescore_vector,omitempty"`
	// VisitPercentage The percentage of postings to visit, from 0 to 100. `0` selects the automatic
	// visit calculation.
	VisitPercentage *float32 `json:"visit_percentage,omitempty"`
}

func (s *KnnEvalSettings) UnmarshalJSON(data []byte) error {

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

		case "exact":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Exact", err)
				}
				s.Exact = &value
			case bool:
				s.Exact = &v
			}

		case "num_candidates":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumCandidates", err)
				}
				s.NumCandidates = &value
			case float64:
				f := int(v)
				s.NumCandidates = &f
			}

		case "rescore_vector":
			if err := dec.Decode(&s.RescoreVector); err != nil {
				return fmt.Errorf("%s | %w", "RescoreVector", err)
			}

		case "visit_percentage":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "VisitPercentage", err)
				}
				f := float32(value)
				s.VisitPercentage = &f
			case float64:
				f := float32(v)
				s.VisitPercentage = &f
			}

		}
	}
	return nil
}

// NewKnnEvalSettings returns a KnnEvalSettings.
func NewKnnEvalSettings() *KnnEvalSettings {
	r := &KnnEvalSettings{}

	return r
}

type KnnEvalSettingsVariant interface {
	KnnEvalSettingsCaster() *KnnEvalSettings
}

func (s *KnnEvalSettings) KnnEvalSettingsCaster() *KnnEvalSettings {
	return s
}
