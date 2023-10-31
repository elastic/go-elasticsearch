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

package translate

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package translate
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/sql/translate/TranslateSqlResponse.ts#L28-L38

type Response struct {
	Aggregations map[string]types.Aggregations `json:"aggregations,omitempty"`
	Fields       []types.FieldAndFormat        `json:"fields,omitempty"`
	Query        *types.Query                  `json:"query,omitempty"`
	Size         *int64                        `json:"size,omitempty"`
	Sort         []types.SortCombinations      `json:"sort,omitempty"`
	Source_      types.SourceConfig            `json:"_source,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Aggregations: make(map[string]types.Aggregations, 0),
	}
	return r
}

func (s *Response) UnmarshalJSON(data []byte) error {
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

		case "aggregations":
			if s.Aggregations == nil {
				s.Aggregations = make(map[string]types.Aggregations, 0)
			}
			if err := dec.Decode(&s.Aggregations); err != nil {
				return err
			}

		case "fields":
			if err := dec.Decode(&s.Fields); err != nil {
				return err
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return err
			}

		case "size":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Size = &value
			case float64:
				f := int64(v)
				s.Size = &f
			}

		case "sort":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(types.SortCombinations)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Sort = append(s.Sort, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Sort); err != nil {
					return err
				}
			}

		case "_source":
			if err := dec.Decode(&s.Source_); err != nil {
				return err
			}

		}
	}
	return nil
}
