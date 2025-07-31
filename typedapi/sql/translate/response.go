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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package translate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package translate
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/sql/translate/TranslateSqlResponse.ts#L27-L37
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
				return fmt.Errorf("%s | %w", "Aggregations", err)
			}

		case "fields":
			if err := dec.Decode(&s.Fields); err != nil {
				return fmt.Errorf("%s | %w", "Fields", err)
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}

		case "size":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Size", err)
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
					return fmt.Errorf("%s | %w", "Sort", err)
				}

				s.Sort = append(s.Sort, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Sort); err != nil {
					return fmt.Errorf("%s | %w", "Sort", err)
				}
			}

		case "_source":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Source_", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		source__field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Source_", err)
				}

				switch t {

				case "exclude_vectors", "excludes", "includes":
					o := types.NewSourceFilter()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Source_", err)
					}
					s.Source_ = o
					break source__field

				}
			}
			if s.Source_ == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Source_); err != nil {
					return fmt.Errorf("%s | %w", "Source_", err)
				}
			}

		}
	}
	return nil
}
