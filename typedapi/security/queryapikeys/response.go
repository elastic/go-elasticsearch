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

package queryapikeys

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package queryapikeys
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/query_api_keys/QueryApiKeysResponse.ts#L26-L45
type Response struct {

	// Aggregations The aggregations result, if requested.
	Aggregations map[string]types.ApiKeyAggregate `json:"aggregations,omitempty"`
	// ApiKeys A list of API key information.
	ApiKeys []types.ApiKey `json:"api_keys"`
	// Count The number of API keys returned in the response.
	Count int `json:"count"`
	// Total The total number of API keys found.
	Total int `json:"total"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Aggregations: make(map[string]types.ApiKeyAggregate, 0),
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
				s.Aggregations = make(map[string]types.ApiKeyAggregate, 0)
			}

			for dec.More() {
				tt, err := dec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return err
				}
				if value, ok := tt.(string); ok {
					if strings.Contains(value, "#") {
						elems := strings.Split(value, "#")
						if len(elems) == 2 {
							if s.Aggregations == nil {
								s.Aggregations = make(map[string]types.ApiKeyAggregate, 0)
							}
							switch elems[0] {

							case "cardinality":
								o := types.NewCardinalityAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "value_count":
								o := types.NewValueCountAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "sterms":
								o := types.NewStringTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "lterms":
								o := types.NewLongTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "dterms":
								o := types.NewDoubleTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "umterms":
								o := types.NewUnmappedTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "multi_terms":
								o := types.NewMultiTermsAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "missing":
								o := types.NewMissingAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "filter":
								o := types.NewFilterAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "filters":
								o := types.NewFiltersAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "range":
								o := types.NewRangeAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "date_range":
								o := types.NewDateRangeAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							case "composite":
								o := types.NewCompositeAggregate()
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o

							default:
								o := make(map[string]any, 0)
								if err := dec.Decode(&o); err != nil {
									return fmt.Errorf("%s | %w", "Aggregations", err)
								}
								s.Aggregations[elems[1]] = o
							}
						} else {
							return errors.New("cannot decode JSON for field Aggregations")
						}
					} else {
						o := make(map[string]any, 0)
						if err := dec.Decode(&o); err != nil {
							return fmt.Errorf("%s | %w", "Aggregations", err)
						}
						s.Aggregations[value] = o
					}
				}
			}

		case "api_keys":
			if err := dec.Decode(&s.ApiKeys); err != nil {
				return fmt.Errorf("%s | %w", "ApiKeys", err)
			}

		case "count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = value
			case float64:
				f := int(v)
				s.Count = f
			}

		case "total":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Total", err)
				}
				s.Total = value
			case float64:
				f := int(v)
				s.Total = f
			}

		}
	}
	return nil
}
