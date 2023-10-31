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
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoexecution"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geovalidationmethod"
)

// GeoBoundingBoxQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/query_dsl/geo.ts#L32-L50
type GeoBoundingBoxQuery struct {
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost               *float32             `json:"boost,omitempty"`
	GeoBoundingBoxQuery map[string]GeoBounds `json:"GeoBoundingBoxQuery,omitempty"`
	// IgnoreUnmapped Set to `true` to ignore an unmapped field and not match any documents for
	// this query.
	// Set to `false` to throw an exception if the field is not mapped.
	IgnoreUnmapped *bool                      `json:"ignore_unmapped,omitempty"`
	QueryName_     *string                    `json:"_name,omitempty"`
	Type           *geoexecution.GeoExecution `json:"type,omitempty"`
	// ValidationMethod Set to `IGNORE_MALFORMED` to accept geo points with invalid latitude or
	// longitude.
	// Set to `COERCE` to also try to infer correct latitude or longitude.
	ValidationMethod *geovalidationmethod.GeoValidationMethod `json:"validation_method,omitempty"`
}

func (s *GeoBoundingBoxQuery) UnmarshalJSON(data []byte) error {

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

		case "boost":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "GeoBoundingBoxQuery":
			if s.GeoBoundingBoxQuery == nil {
				s.GeoBoundingBoxQuery = make(map[string]GeoBounds, 0)
			}
			if err := dec.Decode(&s.GeoBoundingBoxQuery); err != nil {
				return err
			}

		case "ignore_unmapped":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IgnoreUnmapped = &value
			case bool:
				s.IgnoreUnmapped = &v
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryName_ = &o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "validation_method":
			if err := dec.Decode(&s.ValidationMethod); err != nil {
				return err
			}

		default:

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s GeoBoundingBoxQuery) MarshalJSON() ([]byte, error) {
	type opt GeoBoundingBoxQuery
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.GeoBoundingBoxQuery {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "GeoBoundingBoxQuery")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewGeoBoundingBoxQuery returns a GeoBoundingBoxQuery.
func NewGeoBoundingBoxQuery() *GeoBoundingBoxQuery {
	r := &GeoBoundingBoxQuery{
		GeoBoundingBoxQuery: make(map[string]GeoBounds, 0),
	}

	return r
}
