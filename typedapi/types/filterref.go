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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/filtertype"
)

// FilterRef type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/Filter.ts#L31-L41
type FilterRef struct {
	// FilterId The identifier for the filter.
	FilterId string `json:"filter_id"`
	// FilterType If set to `include`, the rule applies for values in the filter. If set to
	// `exclude`, the rule applies for values not in the filter.
	FilterType *filtertype.FilterType `json:"filter_type,omitempty"`
}

func (s *FilterRef) UnmarshalJSON(data []byte) error {

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

		case "filter_id":
			if err := dec.Decode(&s.FilterId); err != nil {
				return err
			}

		case "filter_type":
			if err := dec.Decode(&s.FilterType); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewFilterRef returns a FilterRef.
func NewFilterRef() *FilterRef {
	r := &FilterRef{}

	return r
}
