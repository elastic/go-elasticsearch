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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package types

import (
	"encoding/json"
	"fmt"
)

// CompositeAggregationSource type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d520d9e8cf14cad487de5e0654007686c395b494/specification/_types/aggregations/bucket.ts#L151-L171
type CompositeAggregationSource struct {
	AdditionalCompositeAggregationSourceProperty map[string]json.RawMessage `json:"-"`
	// DateHistogram A date histogram aggregation.
	DateHistogram *CompositeDateHistogramAggregation `json:"date_histogram,omitempty"`
	// GeotileGrid A geotile grid aggregation.
	GeotileGrid *CompositeGeoTileGridAggregation `json:"geotile_grid,omitempty"`
	// Histogram A histogram aggregation.
	Histogram *CompositeHistogramAggregation `json:"histogram,omitempty"`
	// Terms A terms aggregation.
	Terms *CompositeTermsAggregation `json:"terms,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s CompositeAggregationSource) MarshalJSON() ([]byte, error) {
	type opt CompositeAggregationSource
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalCompositeAggregationSourceProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalCompositeAggregationSourceProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewCompositeAggregationSource returns a CompositeAggregationSource.
func NewCompositeAggregationSource() *CompositeAggregationSource {
	r := &CompositeAggregationSource{
		AdditionalCompositeAggregationSourceProperty: make(map[string]json.RawMessage),
	}

	return r
}

type CompositeAggregationSourceVariant interface {
	CompositeAggregationSourceCaster() *CompositeAggregationSource
}

func (s *CompositeAggregationSource) CompositeAggregationSourceCaster() *CompositeAggregationSource {
	return s
}
