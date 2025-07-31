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
)

// SparseVectorIndexOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/mapping/SparseVectorIndexOptions.ts#L22-L42
type SparseVectorIndexOptions struct {
	// Prune Whether to perform pruning, omitting the non-significant tokens from the
	// query to improve query performance.
	// If prune is true but the pruning_config is not specified, pruning will occur
	// but default values will be used.
	// Default: false
	Prune *bool `json:"prune,omitempty"`
	// PruningConfig Optional pruning configuration.
	// If enabled, this will omit non-significant tokens from the query in order to
	// improve query performance.
	// This is only used if prune is set to true.
	// If prune is set to true but pruning_config is not specified, default values
	// will be used.
	PruningConfig *TokenPruningConfig `json:"pruning_config,omitempty"`
}

func (s *SparseVectorIndexOptions) UnmarshalJSON(data []byte) error {

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

		case "prune":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Prune", err)
				}
				s.Prune = &value
			case bool:
				s.Prune = &v
			}

		case "pruning_config":
			if err := dec.Decode(&s.PruningConfig); err != nil {
				return fmt.Errorf("%s | %w", "PruningConfig", err)
			}

		}
	}
	return nil
}

// NewSparseVectorIndexOptions returns a SparseVectorIndexOptions.
func NewSparseVectorIndexOptions() *SparseVectorIndexOptions {
	r := &SparseVectorIndexOptions{}

	return r
}
