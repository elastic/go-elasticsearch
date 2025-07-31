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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tdigestexecutionhint"
)

// MedianAbsoluteDeviationAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/metric.ts#L173-L188
type MedianAbsoluteDeviationAggregation struct {
	// Compression Limits the maximum number of nodes used by the underlying TDigest algorithm
	// to `20 * compression`, enabling control of memory usage and approximation
	// error.
	Compression *Float64 `json:"compression,omitempty"`
	// ExecutionHint The default implementation of TDigest is optimized for performance, scaling
	// to millions or even billions of sample values while maintaining acceptable
	// accuracy levels (close to 1% relative error for millions of samples in some
	// cases).
	// To use an implementation optimized for accuracy, set this parameter to
	// high_accuracy instead.
	ExecutionHint *tdigestexecutionhint.TDigestExecutionHint `json:"execution_hint,omitempty"`
	// Field The field on which to run the aggregation.
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
}

func (s *MedianAbsoluteDeviationAggregation) UnmarshalJSON(data []byte) error {

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

		case "compression":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Compression", err)
				}
				f := Float64(value)
				s.Compression = &f
			case float64:
				f := Float64(v)
				s.Compression = &f
			}

		case "execution_hint":
			if err := dec.Decode(&s.ExecutionHint); err != nil {
				return fmt.Errorf("%s | %w", "ExecutionHint", err)
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Format", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Format = &o

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return fmt.Errorf("%s | %w", "Missing", err)
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		}
	}
	return nil
}

// NewMedianAbsoluteDeviationAggregation returns a MedianAbsoluteDeviationAggregation.
func NewMedianAbsoluteDeviationAggregation() *MedianAbsoluteDeviationAggregation {
	r := &MedianAbsoluteDeviationAggregation{}

	return r
}

type MedianAbsoluteDeviationAggregationVariant interface {
	MedianAbsoluteDeviationAggregationCaster() *MedianAbsoluteDeviationAggregation
}

func (s *MedianAbsoluteDeviationAggregation) MedianAbsoluteDeviationAggregationCaster() *MedianAbsoluteDeviationAggregation {
	return s
}
