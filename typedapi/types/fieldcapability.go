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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeseriesmetrictype"
)

// FieldCapability type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_global/field_caps/types.ts#L23-L81
type FieldCapability struct {
	// Aggregatable Whether this field can be aggregated on all indices.
	Aggregatable bool `json:"aggregatable"`
	// Indices The list of indices where this field has the same type family, or null if all
	// indices have the same type family for the field.
	Indices []string `json:"indices,omitempty"`
	// Meta Merged metadata across all indices as a map of string keys to arrays of
	// values. A value length of 1 indicates that all indices had the same value for
	// this key, while a length of 2 or more indicates that not all indices had the
	// same value for this key.
	Meta Metadata `json:"meta,omitempty"`
	// MetadataField Whether this field is registered as a metadata field.
	MetadataField *bool `json:"metadata_field,omitempty"`
	// MetricConflictsIndices The list of indices where this field is present if these indices
	// don’t have the same `time_series_metric` value for this field.
	MetricConflictsIndices []string `json:"metric_conflicts_indices,omitempty"`
	// NonAggregatableIndices The list of indices where this field is not aggregatable, or null if all
	// indices have the same definition for the field.
	NonAggregatableIndices []string `json:"non_aggregatable_indices,omitempty"`
	// NonDimensionIndices If this list is present in response then some indices have the
	// field marked as a dimension and other indices, the ones in this list, do not.
	NonDimensionIndices []string `json:"non_dimension_indices,omitempty"`
	// NonSearchableIndices The list of indices where this field is not searchable, or null if all
	// indices have the same definition for the field.
	NonSearchableIndices []string `json:"non_searchable_indices,omitempty"`
	// Searchable Whether this field is indexed for search on all indices.
	Searchable bool `json:"searchable"`
	// TimeSeriesDimension Whether this field is used as a time series dimension.
	TimeSeriesDimension *bool `json:"time_series_dimension,omitempty"`
	// TimeSeriesMetric Contains metric type if this fields is used as a time series
	// metrics, absent if the field is not used as metric.
	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type             string                                     `json:"type"`
}

func (s *FieldCapability) UnmarshalJSON(data []byte) error {

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

		case "aggregatable":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Aggregatable = value
			case bool:
				s.Aggregatable = v
			}

		case "indices":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Indices = append(s.Indices, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Indices); err != nil {
					return err
				}
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "metadata_field":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.MetadataField = &value
			case bool:
				s.MetadataField = &v
			}

		case "metric_conflicts_indices":
			if err := dec.Decode(&s.MetricConflictsIndices); err != nil {
				return err
			}

		case "non_aggregatable_indices":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.NonAggregatableIndices = append(s.NonAggregatableIndices, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.NonAggregatableIndices); err != nil {
					return err
				}
			}

		case "non_dimension_indices":
			if err := dec.Decode(&s.NonDimensionIndices); err != nil {
				return err
			}

		case "non_searchable_indices":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.NonSearchableIndices = append(s.NonSearchableIndices, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.NonSearchableIndices); err != nil {
					return err
				}
			}

		case "searchable":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Searchable = value
			case bool:
				s.Searchable = v
			}

		case "time_series_dimension":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.TimeSeriesDimension = &value
			case bool:
				s.TimeSeriesDimension = &v
			}

		case "time_series_metric":
			if err := dec.Decode(&s.TimeSeriesMetric); err != nil {
				return err
			}

		case "type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = o

		}
	}
	return nil
}

// NewFieldCapability returns a FieldCapability.
func NewFieldCapability() *FieldCapability {
	r := &FieldCapability{}

	return r
}
