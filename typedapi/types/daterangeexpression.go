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
// https://github.com/elastic/elasticsearch-specification/tree/accc26662ab4c58f4f6fb0fc1d9fc5249d0de339

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// DateRangeExpression type.
//
// https://github.com/elastic/elasticsearch-specification/blob/accc26662ab4c58f4f6fb0fc1d9fc5249d0de339/specification/_types/aggregations/bucket.ts#L305-L318
type DateRangeExpression struct {
	// From Start of the range (inclusive).
	From FieldDateMath `json:"from,omitempty"`
	// Key Custom key to return the range with.
	Key *string `json:"key,omitempty"`
	// To End of the range (exclusive).
	To FieldDateMath `json:"to,omitempty"`
}

func (s *DateRangeExpression) UnmarshalJSON(data []byte) error {

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

		case "from":
			if err := dec.Decode(&s.From); err != nil {
				return fmt.Errorf("%s | %w", "From", err)
			}

		case "key":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Key", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Key = &o

		case "to":
			if err := dec.Decode(&s.To); err != nil {
				return fmt.Errorf("%s | %w", "To", err)
			}

		}
	}
	return nil
}

// NewDateRangeExpression returns a DateRangeExpression.
func NewDateRangeExpression() *DateRangeExpression {
	r := &DateRangeExpression{}

	return r
}
