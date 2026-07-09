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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// MultipleSynonymGraphFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/cluster/stats/types.ts#L387-L398
type MultipleSynonymGraphFilter struct {
	// AnalyzerCount Number of analyzers across the cluster whose filter chain contains more than
	// one synonym_graph filter.
	AnalyzerCount *int `json:"analyzer_count,omitempty"`
	// IndexCount Number of indices that contain at least one analyzer with more than one
	// synonym_graph filter.
	IndexCount *int `json:"index_count,omitempty"`
}

func (s *MultipleSynonymGraphFilter) UnmarshalJSON(data []byte) error {

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

		case "analyzer_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AnalyzerCount", err)
				}
				s.AnalyzerCount = &value
			case float64:
				f := int(v)
				s.AnalyzerCount = &f
			}

		case "index_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndexCount", err)
				}
				s.IndexCount = &value
			case float64:
				f := int(v)
				s.IndexCount = &f
			}

		}
	}
	return nil
}

// NewMultipleSynonymGraphFilter returns a MultipleSynonymGraphFilter.
func NewMultipleSynonymGraphFilter() *MultipleSynonymGraphFilter {
	r := &MultipleSynonymGraphFilter{}

	return r
}
