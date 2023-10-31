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

// ModelPlotConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/ModelPlot.ts#L23-L42
type ModelPlotConfig struct {
	// AnnotationsEnabled If true, enables calculation and storage of the model change annotations for
	// each entity that is being analyzed.
	AnnotationsEnabled *bool `json:"annotations_enabled,omitempty"`
	// Enabled If true, enables calculation and storage of the model bounds for each entity
	// that is being analyzed.
	Enabled *bool `json:"enabled,omitempty"`
	// Terms Limits data collection to this comma separated list of partition or by field
	// values. If terms are not specified or it is an empty string, no filtering is
	// applied. Wildcards are not supported. Only the specified terms can be viewed
	// when using the Single Metric Viewer.
	Terms *string `json:"terms,omitempty"`
}

func (s *ModelPlotConfig) UnmarshalJSON(data []byte) error {

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

		case "annotations_enabled":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AnnotationsEnabled = &value
			case bool:
				s.AnnotationsEnabled = &v
			}

		case "enabled":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Enabled = &value
			case bool:
				s.Enabled = &v
			}

		case "terms":
			if err := dec.Decode(&s.Terms); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewModelPlotConfig returns a ModelPlotConfig.
func NewModelPlotConfig() *ModelPlotConfig {
	r := &ModelPlotConfig{}

	return r
}
