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

// Settings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/transform/_types/Transform.ts#L98-L144
type Settings struct {
	// AlignCheckpoints Specifies whether the transform checkpoint ranges should be optimized for
	// performance. Such optimization can align
	// checkpoint ranges with the date histogram interval when date histogram is
	// specified as a group source in the
	// transform config. As a result, less document updates in the destination index
	// will be performed thus improving
	// overall performance.
	AlignCheckpoints *bool `json:"align_checkpoints,omitempty"`
	// DatesAsEpochMillis Defines if dates in the ouput should be written as ISO formatted string or as
	// millis since epoch. epoch_millis was
	// the default for transforms created before version 7.11. For compatible output
	// set this value to `true`.
	DatesAsEpochMillis *bool `json:"dates_as_epoch_millis,omitempty"`
	// DeduceMappings Specifies whether the transform should deduce the destination index mappings
	// from the transform configuration.
	DeduceMappings *bool `json:"deduce_mappings,omitempty"`
	// DocsPerSecond Specifies a limit on the number of input documents per second. This setting
	// throttles the transform by adding a
	// wait time between search requests. The default value is null, which disables
	// throttling.
	DocsPerSecond *float32 `json:"docs_per_second,omitempty"`
	// MaxPageSearchSize Defines the initial page size to use for the composite aggregation for each
	// checkpoint. If circuit breaker
	// exceptions occur, the page size is dynamically adjusted to a lower value. The
	// minimum value is `10` and the
	// maximum is `65,536`.
	MaxPageSearchSize *int `json:"max_page_search_size,omitempty"`
	// Unattended If `true`, the transform runs in unattended mode. In unattended mode, the
	// transform retries indefinitely in case
	// of an error which means the transform never fails. Setting the number of
	// retries other than infinite fails in
	// validation.
	Unattended *bool `json:"unattended,omitempty"`
}

func (s *Settings) UnmarshalJSON(data []byte) error {

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

		case "align_checkpoints":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AlignCheckpoints = &value
			case bool:
				s.AlignCheckpoints = &v
			}

		case "dates_as_epoch_millis":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.DatesAsEpochMillis = &value
			case bool:
				s.DatesAsEpochMillis = &v
			}

		case "deduce_mappings":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.DeduceMappings = &value
			case bool:
				s.DeduceMappings = &v
			}

		case "docs_per_second":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.DocsPerSecond = &f
			case float64:
				f := float32(v)
				s.DocsPerSecond = &f
			}

		case "max_page_search_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxPageSearchSize = &value
			case float64:
				f := int(v)
				s.MaxPageSearchSize = &f
			}

		case "unattended":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Unattended = &value
			case bool:
				s.Unattended = &v
			}

		}
	}
	return nil
}

// NewSettings returns a Settings.
func NewSettings() *Settings {
	r := &Settings{}

	return r
}
