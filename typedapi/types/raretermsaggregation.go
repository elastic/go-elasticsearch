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
)

// RareTermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/aggregations/bucket.ts#L687-L717
type RareTermsAggregation struct {
	// Exclude Terms that should be excluded from the aggregation.
	Exclude []string `json:"exclude,omitempty"`
	// Field The field from which to return rare terms.
	Field *string `json:"field,omitempty"`
	// Include Terms that should be included in the aggregation.
	Include TermsInclude `json:"include,omitempty"`
	// MaxDocCount The maximum number of documents a term should appear in.
	MaxDocCount *int64   `json:"max_doc_count,omitempty"`
	Meta        Metadata `json:"meta,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing Missing `json:"missing,omitempty"`
	Name    *string `json:"name,omitempty"`
	// Precision The precision of the internal CuckooFilters.
	// Smaller precision leads to better approximation, but higher memory usage.
	Precision *Float64 `json:"precision,omitempty"`
	ValueType *string  `json:"value_type,omitempty"`
}

func (s *RareTermsAggregation) UnmarshalJSON(data []byte) error {

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

		case "exclude":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Exclude = append(s.Exclude, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Exclude); err != nil {
					return err
				}
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "include":
			if err := dec.Decode(&s.Include); err != nil {
				return err
			}

		case "max_doc_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MaxDocCount = &value
			case float64:
				f := int64(v)
				s.MaxDocCount = &f
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
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

		case "precision":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Precision = &f
			case float64:
				f := Float64(v)
				s.Precision = &f
			}

		case "value_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ValueType = &o

		}
	}
	return nil
}

// NewRareTermsAggregation returns a RareTermsAggregation.
func NewRareTermsAggregation() *RareTermsAggregation {
	r := &RareTermsAggregation{}

	return r
}
