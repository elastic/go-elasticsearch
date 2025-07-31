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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/tdigestexecutionhint"
)

// TDigest type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/metric.ts#L244-L255
type TDigest struct {
	// Compression Limits the maximum number of nodes used by the underlying TDigest algorithm
	// to `20 * compression`, enabling control of memory usage and approximation
	// error.
	Compression *int `json:"compression,omitempty"`
	// ExecutionHint The default implementation of TDigest is optimized for performance, scaling
	// to millions or even billions of sample values while maintaining acceptable
	// accuracy levels (close to 1% relative error for millions of samples in some
	// cases).
	// To use an implementation optimized for accuracy, set this parameter to
	// high_accuracy instead.
	ExecutionHint *tdigestexecutionhint.TDigestExecutionHint `json:"execution_hint,omitempty"`
}

func (s *TDigest) UnmarshalJSON(data []byte) error {

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
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Compression", err)
				}
				s.Compression = &value
			case float64:
				f := int(v)
				s.Compression = &f
			}

		case "execution_hint":
			if err := dec.Decode(&s.ExecutionHint); err != nil {
				return fmt.Errorf("%s | %w", "ExecutionHint", err)
			}

		}
	}
	return nil
}

// NewTDigest returns a TDigest.
func NewTDigest() *TDigest {
	r := &TDigest{}

	return r
}
