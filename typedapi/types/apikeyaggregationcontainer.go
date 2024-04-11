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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// ApiKeyAggregationContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/security/query_api_keys/types.ts#L64-L121
type ApiKeyAggregationContainer struct {
	// Aggregations Sub-aggregations for this aggregation.
	// Only applies to bucket aggregations.
	Aggregations map[string]ApiKeyAggregationContainer `json:"aggregations,omitempty"`
	// Cardinality A single-value metrics aggregation that calculates an approximate count of
	// distinct values.
	Cardinality *CardinalityAggregation `json:"cardinality,omitempty"`
	// Composite A multi-bucket aggregation that creates composite buckets from different
	// sources.
	// Unlike the other multi-bucket aggregations, you can use the `composite`
	// aggregation to paginate *all* buckets from a multi-level aggregation
	// efficiently.
	Composite *CompositeAggregation `json:"composite,omitempty"`
	// DateRange A multi-bucket value source based aggregation that enables the user to define
	// a set of date ranges - each representing a bucket.
	DateRange *DateRangeAggregation `json:"date_range,omitempty"`
	// Filter A single bucket aggregation that narrows the set of documents to those that
	// match a query.
	Filter *ApiKeyQueryContainer `json:"filter,omitempty"`
	// Filters A multi-bucket aggregation where each bucket contains the documents that
	// match a query.
	Filters *ApiKeyFiltersAggregation `json:"filters,omitempty"`
	Meta    Metadata                  `json:"meta,omitempty"`
	Missing *MissingAggregation       `json:"missing,omitempty"`
	// Range A multi-bucket value source based aggregation that enables the user to define
	// a set of ranges - each representing a bucket.
	Range *RangeAggregation `json:"range,omitempty"`
	// Terms A multi-bucket value source based aggregation where buckets are dynamically
	// built - one per unique value.
	Terms *TermsAggregation `json:"terms,omitempty"`
	// ValueCount A single-value metrics aggregation that counts the number of values that are
	// extracted from the aggregated documents.
	ValueCount *ValueCountAggregation `json:"value_count,omitempty"`
}

func (s *ApiKeyAggregationContainer) UnmarshalJSON(data []byte) error {

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

		case "aggregations", "aggs":
			if s.Aggregations == nil {
				s.Aggregations = make(map[string]ApiKeyAggregationContainer, 0)
			}
			if err := dec.Decode(&s.Aggregations); err != nil {
				return fmt.Errorf("%s | %w", "Aggregations", err)
			}

		case "cardinality":
			if err := dec.Decode(&s.Cardinality); err != nil {
				return fmt.Errorf("%s | %w", "Cardinality", err)
			}

		case "composite":
			if err := dec.Decode(&s.Composite); err != nil {
				return fmt.Errorf("%s | %w", "Composite", err)
			}

		case "date_range":
			if err := dec.Decode(&s.DateRange); err != nil {
				return fmt.Errorf("%s | %w", "DateRange", err)
			}

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return fmt.Errorf("%s | %w", "Filter", err)
			}

		case "filters":
			if err := dec.Decode(&s.Filters); err != nil {
				return fmt.Errorf("%s | %w", "Filters", err)
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return fmt.Errorf("%s | %w", "Meta", err)
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return fmt.Errorf("%s | %w", "Missing", err)
			}

		case "range":
			if err := dec.Decode(&s.Range); err != nil {
				return fmt.Errorf("%s | %w", "Range", err)
			}

		case "terms":
			if err := dec.Decode(&s.Terms); err != nil {
				return fmt.Errorf("%s | %w", "Terms", err)
			}

		case "value_count":
			if err := dec.Decode(&s.ValueCount); err != nil {
				return fmt.Errorf("%s | %w", "ValueCount", err)
			}

		}
	}
	return nil
}

// NewApiKeyAggregationContainer returns a ApiKeyAggregationContainer.
func NewApiKeyAggregationContainer() *ApiKeyAggregationContainer {
	r := &ApiKeyAggregationContainer{
		Aggregations: make(map[string]ApiKeyAggregationContainer, 0),
	}

	return r
}
