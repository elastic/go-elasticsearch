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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoshaperelation"
)

// GeoShapeFieldQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/_types/query_dsl/geo.ts#L106-L117
type GeoShapeFieldQuery struct {
	// IndexedShape Query using an indexed shape retrieved from the the specified document and
	// path.
	IndexedShape *FieldLookup `json:"indexed_shape,omitempty"`
	// Relation Spatial relation operator used to search a geo field.
	Relation *geoshaperelation.GeoShapeRelation `json:"relation,omitempty"`
	Shape    json.RawMessage                    `json:"shape,omitempty"`
}

func (s *GeoShapeFieldQuery) UnmarshalJSON(data []byte) error {

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

		case "indexed_shape":
			if err := dec.Decode(&s.IndexedShape); err != nil {
				return fmt.Errorf("%s | %w", "IndexedShape", err)
			}

		case "relation":
			if err := dec.Decode(&s.Relation); err != nil {
				return fmt.Errorf("%s | %w", "Relation", err)
			}

		case "shape":
			if err := dec.Decode(&s.Shape); err != nil {
				return fmt.Errorf("%s | %w", "Shape", err)
			}

		}
	}
	return nil
}

// NewGeoShapeFieldQuery returns a GeoShapeFieldQuery.
func NewGeoShapeFieldQuery() *GeoShapeFieldQuery {
	r := &GeoShapeFieldQuery{}

	return r
}
