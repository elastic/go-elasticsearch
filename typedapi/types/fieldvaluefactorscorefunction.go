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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldvaluefactormodifier"

	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// FieldValueFactorScoreFunction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/query_dsl/compound.ts#L70-L75
type FieldValueFactorScoreFunction struct {
	Factor   *Float64                                           `json:"factor,omitempty"`
	Field    string                                             `json:"field"`
	Missing  *Float64                                           `json:"missing,omitempty"`
	Modifier *fieldvaluefactormodifier.FieldValueFactorModifier `json:"modifier,omitempty"`
}

func (s *FieldValueFactorScoreFunction) UnmarshalJSON(data []byte) error {

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

		case "factor":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Factor = &f
			case float64:
				f := Float64(v)
				s.Factor = &f
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "missing":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Missing = &f
			case float64:
				f := Float64(v)
				s.Missing = &f
			}

		case "modifier":
			if err := dec.Decode(&s.Modifier); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewFieldValueFactorScoreFunction returns a FieldValueFactorScoreFunction.
func NewFieldValueFactorScoreFunction() *FieldValueFactorScoreFunction {
	r := &FieldValueFactorScoreFunction{}

	return r
}
