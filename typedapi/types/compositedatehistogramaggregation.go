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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/valuetype"
)

// CompositeDateHistogramAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/aggregations/bucket.ts#L174-L182
type CompositeDateHistogramAggregation struct {
	// CalendarInterval Either `calendar_interval` or `fixed_interval` must be present
	CalendarInterval *string `json:"calendar_interval,omitempty"`
	// Field Either `field` or `script` must be present
	Field *string `json:"field,omitempty"`
	// FixedInterval Either `calendar_interval` or `fixed_interval` must be present
	FixedInterval *string                    `json:"fixed_interval,omitempty"`
	Format        *string                    `json:"format,omitempty"`
	MissingBucket *bool                      `json:"missing_bucket,omitempty"`
	MissingOrder  *missingorder.MissingOrder `json:"missing_order,omitempty"`
	Offset        Duration                   `json:"offset,omitempty"`
	Order         *sortorder.SortOrder       `json:"order,omitempty"`
	// Script Either `field` or `script` must be present
	Script    Script               `json:"script,omitempty"`
	TimeZone  *string              `json:"time_zone,omitempty"`
	ValueType *valuetype.ValueType `json:"value_type,omitempty"`
}

func (s *CompositeDateHistogramAggregation) UnmarshalJSON(data []byte) error {

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

		case "calendar_interval":
			if err := dec.Decode(&s.CalendarInterval); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "fixed_interval":
			if err := dec.Decode(&s.FixedInterval); err != nil {
				return err
			}

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Format = &o

		case "missing_bucket":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.MissingBucket = &value
			case bool:
				s.MissingBucket = &v
			}

		case "missing_order":
			if err := dec.Decode(&s.MissingOrder); err != nil {
				return err
			}

		case "offset":
			if err := dec.Decode(&s.Offset); err != nil {
				return err
			}

		case "order":
			if err := dec.Decode(&s.Order); err != nil {
				return err
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		case "time_zone":
			if err := dec.Decode(&s.TimeZone); err != nil {
				return err
			}

		case "value_type":
			if err := dec.Decode(&s.ValueType); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewCompositeDateHistogramAggregation returns a CompositeDateHistogramAggregation.
func NewCompositeDateHistogramAggregation() *CompositeDateHistogramAggregation {
	r := &CompositeDateHistogramAggregation{}

	return r
}
