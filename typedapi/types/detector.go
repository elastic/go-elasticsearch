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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/excludefrequent"
)

// Detector type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/Detector.ts#L25-L67
type Detector struct {
	// ByFieldName The field used to split the data. In particular, this property is used for
	// analyzing the splits with respect to their own history. It is used for
	// finding unusual values in the context of the split.
	ByFieldName *string `json:"by_field_name,omitempty"`
	// CustomRules Custom rules enable you to customize the way detectors operate. For example,
	// a rule may dictate conditions under which results should be skipped. Kibana
	// refers to custom rules as job rules.
	CustomRules []DetectionRule `json:"custom_rules,omitempty"`
	// DetectorDescription A description of the detector.
	DetectorDescription *string `json:"detector_description,omitempty"`
	// DetectorIndex A unique identifier for the detector. This identifier is based on the order
	// of the detectors in the `analysis_config`, starting at zero. If you specify a
	// value for this property, it is ignored.
	DetectorIndex *int `json:"detector_index,omitempty"`
	// ExcludeFrequent If set, frequent entities are excluded from influencing the anomaly results.
	// Entities can be considered frequent over time or frequent in a population. If
	// you are working with both over and by fields, you can set `exclude_frequent`
	// to `all` for both fields, or to `by` or `over` for those specific fields.
	ExcludeFrequent *excludefrequent.ExcludeFrequent `json:"exclude_frequent,omitempty"`
	// FieldName The field that the detector uses in the function. If you use an event rate
	// function such as count or rare, do not specify this field. The `field_name`
	// cannot contain double quotes or backslashes.
	FieldName *string `json:"field_name,omitempty"`
	// Function The analysis function that is used. For example, `count`, `rare`, `mean`,
	// `min`, `max`, or `sum`.
	Function *string `json:"function,omitempty"`
	// OverFieldName The field used to split the data. In particular, this property is used for
	// analyzing the splits with respect to the history of all splits. It is used
	// for finding unusual values in the population of all splits.
	OverFieldName *string `json:"over_field_name,omitempty"`
	// PartitionFieldName The field used to segment the analysis. When you use this property, you have
	// completely independent baselines for each value of this field.
	PartitionFieldName *string `json:"partition_field_name,omitempty"`
	// UseNull Defines whether a new series is used as the null series when there is no
	// value for the by or partition fields.
	UseNull *bool `json:"use_null,omitempty"`
}

func (s *Detector) UnmarshalJSON(data []byte) error {

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

		case "by_field_name":
			if err := dec.Decode(&s.ByFieldName); err != nil {
				return err
			}

		case "custom_rules":
			if err := dec.Decode(&s.CustomRules); err != nil {
				return err
			}

		case "detector_description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DetectorDescription = &o

		case "detector_index":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.DetectorIndex = &value
			case float64:
				f := int(v)
				s.DetectorIndex = &f
			}

		case "exclude_frequent":
			if err := dec.Decode(&s.ExcludeFrequent); err != nil {
				return err
			}

		case "field_name":
			if err := dec.Decode(&s.FieldName); err != nil {
				return err
			}

		case "function":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Function = &o

		case "over_field_name":
			if err := dec.Decode(&s.OverFieldName); err != nil {
				return err
			}

		case "partition_field_name":
			if err := dec.Decode(&s.PartitionFieldName); err != nil {
				return err
			}

		case "use_null":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.UseNull = &value
			case bool:
				s.UseNull = &v
			}

		}
	}
	return nil
}

// NewDetector returns a Detector.
func NewDetector() *Detector {
	r := &Detector{}

	return r
}
