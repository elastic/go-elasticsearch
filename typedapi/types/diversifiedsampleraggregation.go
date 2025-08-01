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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sampleraggregationexecutionhint"
)

// DiversifiedSamplerAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/bucket.ts#L333-L357
type DiversifiedSamplerAggregation struct {
	// ExecutionHint The type of value used for de-duplication.
	ExecutionHint *sampleraggregationexecutionhint.SamplerAggregationExecutionHint `json:"execution_hint,omitempty"`
	// Field The field used to provide values used for de-duplication.
	Field *string `json:"field,omitempty"`
	// MaxDocsPerValue Limits how many documents are permitted per choice of de-duplicating value.
	MaxDocsPerValue *int    `json:"max_docs_per_value,omitempty"`
	Script          *Script `json:"script,omitempty"`
	// ShardSize Limits how many top-scoring documents are collected in the sample processed
	// on each shard.
	ShardSize *int `json:"shard_size,omitempty"`
}

func (s *DiversifiedSamplerAggregation) UnmarshalJSON(data []byte) error {

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

		case "execution_hint":
			if err := dec.Decode(&s.ExecutionHint); err != nil {
				return fmt.Errorf("%s | %w", "ExecutionHint", err)
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "max_docs_per_value":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxDocsPerValue", err)
				}
				s.MaxDocsPerValue = &value
			case float64:
				f := int(v)
				s.MaxDocsPerValue = &f
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
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

		}
	}
	return nil
}

// NewDiversifiedSamplerAggregation returns a DiversifiedSamplerAggregation.
func NewDiversifiedSamplerAggregation() *DiversifiedSamplerAggregation {
	r := &DiversifiedSamplerAggregation{}

	return r
}

type DiversifiedSamplerAggregationVariant interface {
	DiversifiedSamplerAggregationCaster() *DiversifiedSamplerAggregation
}

func (s *DiversifiedSamplerAggregation) DiversifiedSamplerAggregationCaster() *DiversifiedSamplerAggregation {
	return s
}
