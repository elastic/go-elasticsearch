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
)

// DenseVectorIndexOptionsRescoreVector type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/mapping/DenseVectorProperty.ts#L215-L223
type DenseVectorIndexOptionsRescoreVector struct {
	// Oversample The oversampling factor to use when searching for the nearest neighbor. This
	// is only applicable to the quantized formats: `bbq_*`, `int4_*`, and `int8_*`.
	// When provided, `oversample * k` vectors will be gathered and then their
	// scores will be re-computed with the original vectors.
	//
	// valid values are between `1.0` and `10.0` (inclusive), or `0` exactly to
	// disable oversampling.
	Oversample float32 `json:"oversample"`
}

func (s *DenseVectorIndexOptionsRescoreVector) UnmarshalJSON(data []byte) error {

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

		case "oversample":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Oversample", err)
				}
				f := float32(value)
				s.Oversample = f
			case float64:
				f := float32(v)
				s.Oversample = f
			}

		}
	}
	return nil
}

// NewDenseVectorIndexOptionsRescoreVector returns a DenseVectorIndexOptionsRescoreVector.
func NewDenseVectorIndexOptionsRescoreVector() *DenseVectorIndexOptionsRescoreVector {
	r := &DenseVectorIndexOptionsRescoreVector{}

	return r
}

type DenseVectorIndexOptionsRescoreVectorVariant interface {
	DenseVectorIndexOptionsRescoreVectorCaster() *DenseVectorIndexOptionsRescoreVector
}

func (s *DenseVectorIndexOptionsRescoreVector) DenseVectorIndexOptionsRescoreVectorCaster() *DenseVectorIndexOptionsRescoreVector {
	return s
}
