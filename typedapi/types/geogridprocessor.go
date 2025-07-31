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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geogridtargetformat"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geogridtiletype"
)

// GeoGridProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ingest/_types/Processors.ts#L388-L429
type GeoGridProcessor struct {
	// ChildrenField If specified and children tiles exist, save those tile addresses to this
	// field as an array of strings.
	ChildrenField *string `json:"children_field,omitempty"`
	// Description Description of the processor.
	// Useful for describing the purpose of the processor or its configuration.
	Description *string `json:"description,omitempty"`
	// Field The field to interpret as a geo-tile.=
	// The field format is determined by the `tile_type`.
	Field string `json:"field"`
	// If Conditionally execute the processor.
	If *string `json:"if,omitempty"`
	// IgnoreFailure Ignore failures for the processor.
	IgnoreFailure *bool `json:"ignore_failure,omitempty"`
	// IgnoreMissing If `true` and `field` does not exist, the processor quietly exits without
	// modifying the document.
	IgnoreMissing *bool `json:"ignore_missing,omitempty"`
	// NonChildrenField If specified and intersecting non-child tiles exist, save their addresses to
	// this field as an array of strings.
	NonChildrenField *string `json:"non_children_field,omitempty"`
	// OnFailure Handle failures for the processor.
	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`
	// ParentField If specified and a parent tile exists, save that tile address to this field.
	ParentField *string `json:"parent_field,omitempty"`
	// PrecisionField If specified, save the tile precision (zoom) as an integer to this field.
	PrecisionField *string `json:"precision_field,omitempty"`
	// Tag Identifier for the processor.
	// Useful for debugging and metrics.
	Tag *string `json:"tag,omitempty"`
	// TargetField The field to assign the polygon shape to, by default, the `field` is updated
	// in-place.
	TargetField *string `json:"target_field,omitempty"`
	// TargetFormat Which format to save the generated polygon in.
	TargetFormat *geogridtargetformat.GeoGridTargetFormat `json:"target_format,omitempty"`
	// TileType Three tile formats are understood: geohash, geotile and geohex.
	TileType geogridtiletype.GeoGridTileType `json:"tile_type"`
}

func (s *GeoGridProcessor) UnmarshalJSON(data []byte) error {

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

		case "children_field":
			if err := dec.Decode(&s.ChildrenField); err != nil {
				return fmt.Errorf("%s | %w", "ChildrenField", err)
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "field":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Field = o

		case "if":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "If", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.If = &o

		case "ignore_failure":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreFailure", err)
				}
				s.IgnoreFailure = &value
			case bool:
				s.IgnoreFailure = &v
			}

		case "ignore_missing":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreMissing", err)
				}
				s.IgnoreMissing = &value
			case bool:
				s.IgnoreMissing = &v
			}

		case "non_children_field":
			if err := dec.Decode(&s.NonChildrenField); err != nil {
				return fmt.Errorf("%s | %w", "NonChildrenField", err)
			}

		case "on_failure":
			if err := dec.Decode(&s.OnFailure); err != nil {
				return fmt.Errorf("%s | %w", "OnFailure", err)
			}

		case "parent_field":
			if err := dec.Decode(&s.ParentField); err != nil {
				return fmt.Errorf("%s | %w", "ParentField", err)
			}

		case "precision_field":
			if err := dec.Decode(&s.PrecisionField); err != nil {
				return fmt.Errorf("%s | %w", "PrecisionField", err)
			}

		case "tag":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Tag", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Tag = &o

		case "target_field":
			if err := dec.Decode(&s.TargetField); err != nil {
				return fmt.Errorf("%s | %w", "TargetField", err)
			}

		case "target_format":
			if err := dec.Decode(&s.TargetFormat); err != nil {
				return fmt.Errorf("%s | %w", "TargetFormat", err)
			}

		case "tile_type":
			if err := dec.Decode(&s.TileType); err != nil {
				return fmt.Errorf("%s | %w", "TileType", err)
			}

		}
	}
	return nil
}

// NewGeoGridProcessor returns a GeoGridProcessor.
func NewGeoGridProcessor() *GeoGridProcessor {
	r := &GeoGridProcessor{}

	return r
}
