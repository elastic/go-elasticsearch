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

// SynonymsSetItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/synonyms/get_synonyms_sets/SynonymsSetsGetResponse.ts#L30-L39
type SynonymsSetItem struct {
	// Count Number of synonym rules that the synonym set contains
	Count int `json:"count"`
	// SynonymsSet Synonyms set identifier
	SynonymsSet string `json:"synonyms_set"`
}

func (s *SynonymsSetItem) UnmarshalJSON(data []byte) error {

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

		case "count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Count = value
			case float64:
				f := int(v)
				s.Count = f
			}

		case "synonyms_set":
			if err := dec.Decode(&s.SynonymsSet); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSynonymsSetItem returns a SynonymsSetItem.
func NewSynonymsSetItem() *SynonymsSetItem {
	r := &SynonymsSetItem{}

	return r
}
