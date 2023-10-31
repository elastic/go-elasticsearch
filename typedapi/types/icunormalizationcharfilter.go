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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icunormalizationmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icunormalizationtype"
)

// IcuNormalizationCharFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/analysis/icu-plugin.ts#L40-L44
type IcuNormalizationCharFilter struct {
	Mode    *icunormalizationmode.IcuNormalizationMode `json:"mode,omitempty"`
	Name    *icunormalizationtype.IcuNormalizationType `json:"name,omitempty"`
	Type    string                                     `json:"type,omitempty"`
	Version *string                                    `json:"version,omitempty"`
}

func (s *IcuNormalizationCharFilter) UnmarshalJSON(data []byte) error {

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

		case "mode":
			if err := dec.Decode(&s.Mode); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s IcuNormalizationCharFilter) MarshalJSON() ([]byte, error) {
	type innerIcuNormalizationCharFilter IcuNormalizationCharFilter
	tmp := innerIcuNormalizationCharFilter{
		Mode:    s.Mode,
		Name:    s.Name,
		Type:    s.Type,
		Version: s.Version,
	}

	tmp.Type = "icu_normalizer"

	return json.Marshal(tmp)
}

// NewIcuNormalizationCharFilter returns a IcuNormalizationCharFilter.
func NewIcuNormalizationCharFilter() *IcuNormalizationCharFilter {
	r := &IcuNormalizationCharFilter{}

	return r
}
