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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoshaperelation"
)

// ShapeFieldQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/query_dsl/specialized.ts#L354-L367
type ShapeFieldQuery struct {
	// IndexedShape Queries using a pre-indexed shape.
	IndexedShape *FieldLookup `json:"indexed_shape,omitempty"`
	// Relation Spatial relation between the query shape and the document shape.
	Relation *geoshaperelation.GeoShapeRelation `json:"relation,omitempty"`
	// Shape Queries using an inline shape definition in GeoJSON or Well Known Text (WKT)
	// format.
	Shape json.RawMessage `json:"shape,omitempty"`
}

func (s *ShapeFieldQuery) UnmarshalJSON(data []byte) error {

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
				return err
			}

		case "relation":
			if err := dec.Decode(&s.Relation); err != nil {
				return err
			}

		case "shape":
			if err := dec.Decode(&s.Shape); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewShapeFieldQuery returns a ShapeFieldQuery.
func NewShapeFieldQuery() *ShapeFieldQuery {
	r := &ShapeFieldQuery{}

	return r
}
