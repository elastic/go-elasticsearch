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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package searchmvt

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gridtype"
)

// Request holds the request body struct for the package searchmvt
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_global/search_mvt/SearchMvtRequest.ts#L33-L164
type Request struct {

	// Aggs Sub-aggregations for the geotile_grid.
	//
	// Supports the following aggregation types:
	// - avg
	// - cardinality
	// - max
	// - min
	// - sum
	Aggs map[string]types.Aggregations `json:"aggs,omitempty"`
	// ExactBounds If false, the meta layer’s feature is the bounding box of the tile.
	// If true, the meta layer’s feature is a bounding box resulting from a
	// geo_bounds aggregation. The aggregation runs on <field> values that intersect
	// the <zoom>/<x>/<y> tile with wrap_longitude set to false. The resulting
	// bounding box may be larger than the vector tile.
	ExactBounds *bool `json:"exact_bounds,omitempty"`
	// Extent Size, in pixels, of a side of the tile. Vector tiles are square with equal
	// sides.
	Extent *int `json:"extent,omitempty"`
	// Fields Fields to return in the `hits` layer. Supports wildcards (`*`).
	// This parameter does not support fields with array values. Fields with array
	// values may return inconsistent results.
	Fields []string `json:"fields,omitempty"`
	// GridPrecision Additional zoom levels available through the aggs layer. For example, if
	// <zoom> is 7
	// and grid_precision is 8, you can zoom in up to level 15. Accepts 0-8. If 0,
	// results
	// don’t include the aggs layer.
	GridPrecision *int `json:"grid_precision,omitempty"`
	// GridType Determines the geometry type for features in the aggs layer. In the aggs
	// layer,
	// each feature represents a geotile_grid cell. If 'grid' each feature is a
	// Polygon
	// of the cells bounding box. If 'point' each feature is a Point that is the
	// centroid
	// of the cell.
	GridType *gridtype.GridType `json:"grid_type,omitempty"`
	// Query Query DSL used to filter documents for the search.
	Query *types.Query `json:"query,omitempty"`
	// RuntimeMappings Defines one or more runtime fields in the search request. These fields take
	// precedence over mapped fields with the same name.
	RuntimeMappings map[string]types.RuntimeField `json:"runtime_mappings,omitempty"`
	// Size Maximum number of features to return in the hits layer. Accepts 0-10000.
	// If 0, results don’t include the hits layer.
	Size *int `json:"size,omitempty"`
	// Sort Sorts features in the hits layer. By default, the API calculates a bounding
	// box for each feature. It sorts features based on this box’s diagonal length,
	// from longest to shortest.
	Sort []types.SortCombinations `json:"sort,omitempty"`
	// TrackTotalHits Number of hits matching the query to count accurately. If `true`, the exact
	// number
	// of hits is returned at the cost of some performance. If `false`, the response
	// does
	// not include the total number of hits matching the query.
	TrackTotalHits *types.TrackHits `json:"track_total_hits,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Aggs: make(map[string]types.Aggregations, 0),
	}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Searchmvt request: %w", err)
	}

	return &req, nil
}
