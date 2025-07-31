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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldsortnumerictype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// FieldSort type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/sort.ts#L43-L52
type FieldSort struct {
	Format       *string                                    `json:"format,omitempty"`
	Missing      Missing                                    `json:"missing,omitempty"`
	Mode         *sortmode.SortMode                         `json:"mode,omitempty"`
	Nested       *NestedSortValue                           `json:"nested,omitempty"`
	NumericType  *fieldsortnumerictype.FieldSortNumericType `json:"numeric_type,omitempty"`
	Order        *sortorder.SortOrder                       `json:"order,omitempty"`
	UnmappedType *fieldtype.FieldType                       `json:"unmapped_type,omitempty"`
}

func (s *FieldSort) UnmarshalJSON(data []byte) error {

	if !bytes.HasPrefix(data, []byte(`{`)) {
		err := json.NewDecoder(bytes.NewReader(data)).Decode(&s.Order)
		return err
	}

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

		case "mode":
			if err := dec.Decode(&s.Mode); err != nil {
				return fmt.Errorf("%s | %w", "Mode", err)
			}

		case "nested":
			if err := dec.Decode(&s.Nested); err != nil {
				return fmt.Errorf("%s | %w", "Nested", err)
			}

		case "numeric_type":
			if err := dec.Decode(&s.NumericType); err != nil {
				return fmt.Errorf("%s | %w", "NumericType", err)
			}

		case "order":
			if err := dec.Decode(&s.Order); err != nil {
				return fmt.Errorf("%s | %w", "Order", err)
			}

		case "unmapped_type":
			if err := dec.Decode(&s.UnmappedType); err != nil {
				return fmt.Errorf("%s | %w", "UnmappedType", err)
			}

		}
	}
	return nil
}

// NewFieldSort returns a FieldSort.
func NewFieldSort() *FieldSort {
	r := &FieldSort{}

	return r
}
