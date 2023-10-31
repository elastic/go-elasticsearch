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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/ttesttype"
)

// TTestAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/aggregations/metric.ts#L294-L308
type TTestAggregation struct {
	// A Test population A.
	A *TestPopulation `json:"a,omitempty"`
	// B Test population B.
	B    *TestPopulation `json:"b,omitempty"`
	Meta Metadata        `json:"meta,omitempty"`
	Name *string         `json:"name,omitempty"`
	// Type The type of test.
	Type *ttesttype.TTestType `json:"type,omitempty"`
}

func (s *TTestAggregation) UnmarshalJSON(data []byte) error {

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

		case "a":
			if err := dec.Decode(&s.A); err != nil {
				return err
			}

		case "b":
			if err := dec.Decode(&s.B); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = &o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewTTestAggregation returns a TTestAggregation.
func NewTTestAggregation() *TTestAggregation {
	r := &TTestAggregation{}

	return r
}
