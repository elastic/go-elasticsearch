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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationexecutionhint"
)

// SignificantTermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/aggregations/bucket.ts#L342-L358
type SignificantTermsAggregation struct {
	BackgroundFilter  *Query                                                       `json:"background_filter,omitempty"`
	ChiSquare         *ChiSquareHeuristic                                          `json:"chi_square,omitempty"`
	Exclude           []string                                                     `json:"exclude,omitempty"`
	ExecutionHint     *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`
	Field             *string                                                      `json:"field,omitempty"`
	Gnd               *GoogleNormalizedDistanceHeuristic                           `json:"gnd,omitempty"`
	Include           TermsInclude                                                 `json:"include,omitempty"`
	Jlh               *EmptyObject                                                 `json:"jlh,omitempty"`
	Meta              Metadata                                                     `json:"meta,omitempty"`
	MinDocCount       *int64                                                       `json:"min_doc_count,omitempty"`
	MutualInformation *MutualInformationHeuristic                                  `json:"mutual_information,omitempty"`
	Name              *string                                                      `json:"name,omitempty"`
	Percentage        *PercentageScoreHeuristic                                    `json:"percentage,omitempty"`
	ScriptHeuristic   *ScriptedHeuristic                                           `json:"script_heuristic,omitempty"`
	ShardMinDocCount  *int64                                                       `json:"shard_min_doc_count,omitempty"`
	ShardSize         *int                                                         `json:"shard_size,omitempty"`
	Size              *int                                                         `json:"size,omitempty"`
}

func (s *SignificantTermsAggregation) UnmarshalJSON(data []byte) error {

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

		case "background_filter":
			if err := dec.Decode(&s.BackgroundFilter); err != nil {
				return err
			}

		case "chi_square":
			if err := dec.Decode(&s.ChiSquare); err != nil {
				return err
			}

		case "exclude":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Exclude = append(s.Exclude, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Exclude); err != nil {
					return err
				}
			}

		case "execution_hint":
			if err := dec.Decode(&s.ExecutionHint); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "gnd":
			if err := dec.Decode(&s.Gnd); err != nil {
				return err
			}

		case "include":
			if err := dec.Decode(&s.Include); err != nil {
				return err
			}

		case "jlh":
			if err := dec.Decode(&s.Jlh); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "min_doc_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MinDocCount = &value
			case float64:
				f := int64(v)
				s.MinDocCount = &f
			}

		case "mutual_information":
			if err := dec.Decode(&s.MutualInformation); err != nil {
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

		case "percentage":
			if err := dec.Decode(&s.Percentage); err != nil {
				return err
			}

		case "script_heuristic":
			if err := dec.Decode(&s.ScriptHeuristic); err != nil {
				return err
			}

		case "shard_min_doc_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ShardMinDocCount = &value
			case float64:
				f := int64(v)
				s.ShardMinDocCount = &f
			}

		case "shard_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ShardSize = &value
			case float64:
				f := int(v)
				s.ShardSize = &f
			}

		case "size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		}
	}
	return nil
}

// NewSignificantTermsAggregation returns a SignificantTermsAggregation.
func NewSignificantTermsAggregation() *SignificantTermsAggregation {
	r := &SignificantTermsAggregation{}

	return r
}
