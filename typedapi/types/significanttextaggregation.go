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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationexecutionhint"
)

// SignificantTextAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/bucket.ts#L886-L961
type SignificantTextAggregation struct {
	// BackgroundFilter A background filter that can be used to focus in on significant terms within
	// a narrower context, instead of the entire index.
	BackgroundFilter *Query `json:"background_filter,omitempty"`
	// ChiSquare Use Chi square, as described in "Information Retrieval", Manning et al.,
	// Chapter 13.5.2, as the significance score.
	ChiSquare *ChiSquareHeuristic `json:"chi_square,omitempty"`
	// Exclude Values to exclude.
	Exclude []string `json:"exclude,omitempty"`
	// ExecutionHint Determines whether the aggregation will use field values directly or global
	// ordinals.
	ExecutionHint *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`
	// Field The field from which to return significant text.
	Field *string `json:"field,omitempty"`
	// FilterDuplicateText Whether to out duplicate text to deal with noisy data.
	FilterDuplicateText *bool `json:"filter_duplicate_text,omitempty"`
	// Gnd Use Google normalized distance as described in "The Google Similarity
	// Distance", Cilibrasi and Vitanyi, 2007, as the significance score.
	Gnd *GoogleNormalizedDistanceHeuristic `json:"gnd,omitempty"`
	// Include Values to include.
	Include TermsInclude `json:"include,omitempty"`
	// Jlh Use JLH score as the significance score.
	Jlh *EmptyObject `json:"jlh,omitempty"`
	// MinDocCount Only return values that are found in more than `min_doc_count` hits.
	MinDocCount *int64 `json:"min_doc_count,omitempty"`
	// MutualInformation Use mutual information as described in "Information Retrieval", Manning et
	// al., Chapter 13.5.1, as the significance score.
	MutualInformation *MutualInformationHeuristic `json:"mutual_information,omitempty"`
	// Percentage A simple calculation of the number of documents in the foreground sample with
	// a term divided by the number of documents in the background with the term.
	Percentage *PercentageScoreHeuristic `json:"percentage,omitempty"`
	// ScriptHeuristic Customized score, implemented via a script.
	ScriptHeuristic *ScriptedHeuristic `json:"script_heuristic,omitempty"`
	// ShardMinDocCount Regulates the certainty a shard has if the values should actually be added to
	// the candidate list or not with respect to the min_doc_count.
	// Values will only be considered if their local shard frequency within the set
	// is higher than the `shard_min_doc_count`.
	ShardMinDocCount *int64 `json:"shard_min_doc_count,omitempty"`
	// ShardSize The number of candidate terms produced by each shard.
	// By default, `shard_size` will be automatically estimated based on the number
	// of shards and the `size` parameter.
	ShardSize *int `json:"shard_size,omitempty"`
	// Size The number of buckets returned out of the overall terms list.
	Size *int `json:"size,omitempty"`
	// SourceFields Overrides the JSON `_source` fields from which text will be analyzed.
	SourceFields []string `json:"source_fields,omitempty"`
}

func (s *SignificantTextAggregation) UnmarshalJSON(data []byte) error {

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
				return fmt.Errorf("%s | %w", "BackgroundFilter", err)
			}

		case "chi_square":
			if err := dec.Decode(&s.ChiSquare); err != nil {
				return fmt.Errorf("%s | %w", "ChiSquare", err)
			}

		case "exclude":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Exclude", err)
				}

				s.Exclude = append(s.Exclude, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Exclude); err != nil {
					return fmt.Errorf("%s | %w", "Exclude", err)
				}
			}

		case "execution_hint":
			if err := dec.Decode(&s.ExecutionHint); err != nil {
				return fmt.Errorf("%s | %w", "ExecutionHint", err)
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "filter_duplicate_text":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "FilterDuplicateText", err)
				}
				s.FilterDuplicateText = &value
			case bool:
				s.FilterDuplicateText = &v
			}

		case "gnd":
			if err := dec.Decode(&s.Gnd); err != nil {
				return fmt.Errorf("%s | %w", "Gnd", err)
			}

		case "include":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Include", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		include_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Include", err)
				}

				switch t {

				case "num_partitions", "partition":
					o := NewTermsPartition()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Include", err)
					}
					s.Include = o
					break include_field

				}
			}
			if s.Include == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Include); err != nil {
					return fmt.Errorf("%s | %w", "Include", err)
				}
			}

		case "jlh":
			if err := dec.Decode(&s.Jlh); err != nil {
				return fmt.Errorf("%s | %w", "Jlh", err)
			}

		case "min_doc_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinDocCount", err)
				}
				s.MinDocCount = &value
			case float64:
				f := int64(v)
				s.MinDocCount = &f
			}

		case "mutual_information":
			if err := dec.Decode(&s.MutualInformation); err != nil {
				return fmt.Errorf("%s | %w", "MutualInformation", err)
			}

		case "percentage":
			if err := dec.Decode(&s.Percentage); err != nil {
				return fmt.Errorf("%s | %w", "Percentage", err)
			}

		case "script_heuristic":
			if err := dec.Decode(&s.ScriptHeuristic); err != nil {
				return fmt.Errorf("%s | %w", "ScriptHeuristic", err)
			}

		case "shard_min_doc_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardMinDocCount", err)
				}
				s.ShardMinDocCount = &value
			case float64:
				f := int64(v)
				s.ShardMinDocCount = &f
			}

		case "shard_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardSize", err)
				}
				s.ShardSize = &value
			case float64:
				f := int(v)
				s.ShardSize = &f
			}

		case "size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Size", err)
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		case "source_fields":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "SourceFields", err)
				}

				s.SourceFields = append(s.SourceFields, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.SourceFields); err != nil {
					return fmt.Errorf("%s | %w", "SourceFields", err)
				}
			}

		}
	}
	return nil
}

// NewSignificantTextAggregation returns a SignificantTextAggregation.
func NewSignificantTextAggregation() *SignificantTextAggregation {
	r := &SignificantTextAggregation{}

	return r
}

type SignificantTextAggregationVariant interface {
	SignificantTextAggregationCaster() *SignificantTextAggregation
}

func (s *SignificantTextAggregation) SignificantTextAggregationCaster() *SignificantTextAggregation {
	return s
}
