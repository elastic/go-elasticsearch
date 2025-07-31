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
	"encoding/json"
	"fmt"
)

// SmoothingModelContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/search/_types/suggester.ts#L445-L461
type SmoothingModelContainer struct {
	AdditionalSmoothingModelContainerProperty map[string]json.RawMessage `json:"-"`
	// Laplace A smoothing model that uses an additive smoothing where a constant (typically
	// `1.0` or smaller) is added to all counts to balance weights.
	Laplace *LaplaceSmoothingModel `json:"laplace,omitempty"`
	// LinearInterpolation A smoothing model that takes the weighted mean of the unigrams, bigrams, and
	// trigrams based on user supplied weights (lambdas).
	LinearInterpolation *LinearInterpolationSmoothingModel `json:"linear_interpolation,omitempty"`
	// StupidBackoff A simple backoff model that backs off to lower order n-gram models if the
	// higher order count is `0` and discounts the lower order n-gram model by a
	// constant factor.
	StupidBackoff *StupidBackoffSmoothingModel `json:"stupid_backoff,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s SmoothingModelContainer) MarshalJSON() ([]byte, error) {
	type opt SmoothingModelContainer
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
	for key, value := range s.AdditionalSmoothingModelContainerProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalSmoothingModelContainerProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewSmoothingModelContainer returns a SmoothingModelContainer.
func NewSmoothingModelContainer() *SmoothingModelContainer {
	r := &SmoothingModelContainer{
		AdditionalSmoothingModelContainerProperty: make(map[string]json.RawMessage),
	}

	return r
}
